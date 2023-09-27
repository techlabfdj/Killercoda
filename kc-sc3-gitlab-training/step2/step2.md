# configure a gitlab runner

You should see a lot of logs and the following error from the runner:  
```shell
runner.gitlab.local | ERROR: Failed to load config stat /etc/gitlab-runner/config.toml: no such file or directory  builds=0
```


The configuration of our Runner is missing, we need to **connect** it to the GitLab server to generate a new one. 
Thus, we must register our runner:  

We have to use the web interface to finalize the GitLab setup and then **get a registration token** for our runner. 
Follow the following steps:  


- Go to http://localhost/ .
- Set the `root` GitLab password.
![gitlab register root page](img/register-root.png)
- Connect to the `root` user using the previously registered password.
![gitlab login page](img/login-root.png)
- Create a new `Project` using the `New Project` button.
![gitlab home page](img/home-page.png)
![gitlab home page](img/new-repo.png)
- Go to `Settings`/`CI / CD` in your `Project`.
![gitlab repo page](img/repo-page-hover-settings-ci-cd.png)
- Expand the runner section.
![gitlab repo ci/cd settings](img/repo-settings-ci-cd-runners-page.png)
- Copy the registration token.
![gitlab repo ci/cd settings](img/repo-settings-ci-cd-runners-page-token.png)

If you are running the environment on a **remote machine** and don't have access to the HTTP port. You can't access http://localhost and thus, you may want to **setup port forwarding** to your remote machine (see instructions [here](https://gitlab.techlabfdj.io/techlab/training/ssh/-/blob/master/PortForwarding.md)).

