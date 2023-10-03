# configure a gitlab runner

Using `docker logs gitlab.local`{{exec}}, you should see a lot of logs and the following error from the runner:  
```shell
runner.gitlab.local | ERROR: Failed to load config stat /etc/gitlab-runner/config.toml: no such file or directory  builds=0
```

The configuration of our Runner is missing, we need to **connect** it to the GitLab server to generate a new one. 
Thus, we must register our runner:  

We have to use the web interface to finalize the GitLab setup and then **get a registration token** for our runner. 
Follow the following steps:  

## Get the `root` GitLab password.

Open a shell on the gitlab container: `docker exec -it gitlab.local bash`{{exec}}  
In this shell, retrieve initial root password : `cat /etc/gitlab/initial_root_password`{{exec}}  

## Let's play with gitlab
- Go to {{TRAFFIC_HOST1_80}}

- Connect to the `root` user using the previously retrieved password.
![Scan results](../img/login-root.png)

- Create a new `Project` using the `New Project` button.
![Scan results](../img/home-page.png)
![Scan results](../img/new-repo.png)

- Go to `Settings`/`CI / CD` in your `Project`.
![Scan results](../img/repo-page-hover-settings-ci-cd.png)

- Expand the runner section.
![Scan results](../img/repo-settings-ci-cd-runners-page.png)

- Copy the registration token.
![Scan results](../img/repo-settings-ci-cd-runners-page-token.png)

# Next
You can move on to the next step by clicking the 'Check' button.
The check button uses a clone of the script verify_step2.sh to check that you successfully
