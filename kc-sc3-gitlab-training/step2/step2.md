# configure a gitlab runner

You should see a lot of logs and the following error from the runner:  
```shell
runner.gitlab.local | ERROR: Failed to load config stat /etc/gitlab-runner/config.toml: no such file or directory  builds=0
```


The configuration of our Runner is missing, we need to **connect** it to the GitLab server to generate a new one. 
Thus, we must register our runner:  

We have to use the web interface to finalize the GitLab setup and then **get a registration token** for our runner. 
Follow the following steps:  


- Go to {{TRAFFIC_HOST1_80}}
- Set the `root` GitLab password.
![](./img/register-root.png)
- Connect to the `root` user using the previously registered password.
![Scan results](./img/login-root.png)
- Create a new `Project` using the `New Project` button.
![Scan results](./img/home-page.png)
![Scan results](./img/new-repo.png)
- Go to `Settings`/`CI / CD` in your `Project`.
![Scan results](./img/repo-page-hover-settings-ci-cd.png)
- Expand the runner section.
![Scan results](./img/repo-settings-ci-cd-runners-page.png)
- Copy the registration token.
![Scan results](./img/repo-settings-ci-cd-runners-page-token.png)


