# acceptance-tests
Acceptance tests setup et al for GoReleaser.

Currently we test `homebrew-taps` for `gitlab` and `gitea` with a local
setup and some manual steps. To simplify this process we provide as much
automation as possible to set up a local instance.

The setup assumes that you have the actual [goreleaser](https://github.com/goreleaser/goreleaser)
repository in the same directory level

```console
.
├── acceptance-tests
├── goreleaser
```

## gitea

Follow the steps to setup the local instance:
```sh
# starts and initializes a local server
./local-gitea.sh
# rename the git remote
git remote rename gitea-http origin
# the the gitea UI and login with USER PASSWORD from 'local-gitea.sh
open http://localhost:3000/
# build the goreleaser binary you want to test
cd ../goreleaser && make build && mv goreleaser ../acceptance-tests && cd ../acceptance-tests
# tag the repo
git tag 0.1.0
# create an access token and save as env GITEA_TOKEN
http://localhost:3000/user/settings/applications
# export it
export GITEA_TOKEN="abc"
# publish the binaries to the local gitea
./goreleaser --config=./goreleaser-gitea-local.yml --debug --rm-dist
# if it fails 
git tag -d 0.1.0
# delete the tag by hand at
http://localhost:3000/goreleaser/acceptance-tests/releases

# shutdown when your finished
./shutdown.sh
```

## gitlab
For the gitlab installation we currently have only manual steps. We suggest keeping the 
directories in the `$GITLAB_HOME` directory and remove then only when using a new image tag.
Also due to the fact that the full setup takes around 12 minutes.

Follow the steps to setup the local instance:
```sh
# The path structure must not be too deep or gitaly cannot create a socket
# see https://gitlab.com/gitlab-org/gitaly/-/issues/2311
export GITLAB_HOME=$HOME/gitlab
# cleanup local folders
rm -rf $GITLAB_HOME/config $GITLAB_HOME/data $GITLAB_HOME/logs
# start gitlab
docker-compose -f docker-compose-gitlab.yml up -d
# follow log and wait about 12 minutes until every
# while [[ "$(curl -s -o /dev/null -w ''%{http_code}'' localhost:10080)" != "302" ]]; do echo "zzz..."; sleep 5; done
docker-compose -f docker-compose-gitlab.yml logs -tf
# terminal 2
# curl -v localhost:10080/api/v4
# register new root user -> set password root123!
http://localhost:10080/
# login as root
http://localhost:10080/users/sign_in
# create a new user and check the 'admin' box
## goreleaser
## goreleaser@acme.com
http://localhost:10080/admin/users/new
# set the password cuz setting via email is not possible
# set to 'testpwd123!'
http://localhost:10080/admin/users/goreleaser/edit
# logout and login as goreleaser and maybe set a new password if requested
http://localhost:10080/users/sign_in
# create an access token with API scope
http://localhost:10080/profile/personal_access_tokens
# export it
export GITLAB_TOKEN="abc"
# create 'acceptance-tests' (private) and 'homebrew-tap' (public) repository 
http://localhost:10080/projects/new
# build the goreleaser binary you want to test
cd ../goreleaser && make build && mv -f goreleaser ../acceptance-tests && cd ../acceptance-tests
# tag the repo
git tag 0.1.0
# clean remotes
git remote rename origin xyz
# add the new origin
git remote add origin http://goreleaser:testpwd123\!@localhost:10080/goreleaser/acceptance-tests.git
# push it 
git push origin main
# publish locally
./goreleaser --config=./goreleaser-gitlab-local.yml --debug --rm-dist
# see the new release
http://localhost:10080/goreleaser/acceptance-tests/-/releases
```

