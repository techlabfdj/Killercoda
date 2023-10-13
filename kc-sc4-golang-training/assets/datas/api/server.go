package api

import (
	"context"
	"fmt"
	"gitlab-techlab/techlab/training/golang/gin-samples/datas/stores"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
)

const listenAddr = "127.0.0.1:8080"

const (
	// Version holds the tag string for binary (loaded from git by Makefile)
	Version = "develop"
	// Build holds the build version (loaded from git by Makfile)
	Build = "N/A"
	// API is path under which all admin services are exposed
	APIPath = "/datas"
)

// Server holds the admin server
type Server struct {
	sync.RWMutex
	logger     io.Writer
	httpServer *http.Server
	listener   net.Listener
	router     *gin.Engine
}

// Init initializes serverserver using provided:
// * listening address (no default: must be given)
func (s *Server) Init(listenAddr string, logDir string) error {
	s.Lock()
	defer s.Unlock()

	//gin.SetMode(gin.ReleaseMode)

	s.router = gin.New()
	s.router.HandleMethodNotAllowed = true

	log.Printf("\nDatas web server - Version: %s, Build %s\n", Version, Build)
	log.Printf("\nPowered by gin framework version %s\n", gin.Version)
	
	//check if the Log directory already exist
	// If not, exit with fatal error.
	dir, err := os.Stat(logDir)
	if os.IsNotExist(err) || !dir.IsDir() {
		log.Fatalf("specified log directory : \"%s\" does not exist : %s", logDir, err)
	} 
	log.Printf("\naccess and application log directory => \"%s\" \n", logDir)
	
	accesslogWriter, err := rotatelogs.New(
		logDir+"/access_log.%Y-%m-%d",
		rotatelogs.WithRotationCount(14),
		rotatelogs.WithLinkName(logDir+"/accss_log"),
	)
	
	if err != nil {
		log.Fatalf("Failed to Initialize Access Log File:  %s", err)
	}

	applogWriter, err := rotatelogs.New(
		logDir+"application_log.%Y-%m-%d",
		rotatelogs.WithRotationCount(14),
		rotatelogs.WithLinkName(logDir+"/appication_log"),
	)
	//put stdout first in order to ensure that the logs are written even if the applogWriter file does not exist
	gin.DefaultWriter = io.MultiWriter(os.Stdout, applogWriter)
	s.router.Use(gin.Recovery())

	// LoggerWithFormater middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	s.router.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Output: accesslogWriter,
		Formatter: func(param gin.LogFormatterParams) string {
			// your custom format
			return fmt.Sprintf("%s - [%s] \"%s %s %s\" %d %dms \"%s\" %s\n",
				param.ClientIP,
				param.TimeStamp.Format(time.RFC3339Nano),
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				int64(param.Latency/time.Millisecond),
				param.Request.UserAgent(),
				param.ErrorMessage,
			)
		},
	}))

	s.router.Use(gin.Recovery())

	//store creation
	store := stores.NewInMemory()
	store.SetCapacity(100)
	
	//api rooting
	var dataAPI = New(store)
	s.router.GET("/datas", dataAPI.GetDatas)
	s.router.GET("/datas/:id", dataAPI.GetData)
	s.router.PUT("/datas/:id", dataAPI.UpdateData)
	s.router.DELETE("/datas/:id", dataAPI.DeleteData)
	s.router.POST("/datas", dataAPI.CreateData)

	s.httpServer = &http.Server{
		Addr:    listenAddr,
		Handler: s.router,
	}

	if listenAddr != "" {
		var err error
		if s.listener, err = net.Listen("tcp", listenAddr); err != nil {
			return fmt.Errorf("unable to initialize admin listener on %s: %w", listenAddr, err)
		}
	}
	return nil
}

// Startup runs the server
func (s *Server) Startup() int {
	s.RLock()
	if s.httpServer == nil {
		fmt.Fprintln(gin.DefaultWriter, "can't start api : not initialized")
	}
	s.RUnlock()
	fmt.Fprintln(gin.DefaultWriter, "api available at http://"+s.listener.Addr().String())

	// service connections

	exitCode := 0
	if err := s.httpServer.Serve(s.listener); err != nil && err != http.ErrServerClosed {
		fmt.Fprintln(gin.DefaultWriter, "api server exited with error:")
		exitCode = 1
	} else {
		fmt.Fprintln(gin.DefaultWriter, "api server exited")
	}
	return exitCode
}

// Shutdown allows t
func (s *Server) Shutdown(gracePeriod time.Duration) (err error) {
	s.Lock()
	defer s.Unlock()
	if s.httpServer != nil {
		ctx, cancel := context.WithTimeout(context.Background(), gracePeriod)
		defer cancel()
		if err = s.httpServer.Shutdown(ctx); err != nil {
			fmt.Fprintln(gin.DefaultWriter, "api server shutdown failed:", err)
		} else {
			s.httpServer = nil
			fmt.Fprintln(gin.DefaultWriter, "api server shutdown completed")
		}
	}
	return
}
