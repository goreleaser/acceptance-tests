# acceptance-tests
Acceptance tests setup et al for GoReleaser.

Currently we test `homebrew-taps` for `gitlab` and `gitea` with a local
setup and some manual steps. To simplify this process we provide as much
automation as possible to set up a local instance.

## gitea
```sh
# starts and initializes a local server
./local-gitea.sh
# shutdown
./shutdown.sh
# build the goreleaser binary you want to test
cd ../goreleaser && make build && mv ../goreleaser/goreleaser .
# tag the repo
git tag 0.1.0
# create an access token and save as env GITEA_TOKEN
# publish locally
./goreleaser --config=./goreleaser-gitea-local.yml --debug --rm-dist
# if it fails 
git tag -d 0.1.0
# delete the tag by hand at
http://localhost:3000/goreleaser/acceptance-tests/releases
```

## gitlab
See https://docs.gitlab.com/omnibus/docker
```sh
# see https://gitlab.com/gitlab-org/gitaly/-/issues/2311
export GITLAB_HOME=$HOME/gitlab
# cleanup local folders
rm -rf gitlab/config gitlab/data gitlab/logs
# start gitlab
docker-compose -f docker-compose-gitlab.yml up -d
# follow log and wait about 12 minutes until every
# while [[ "$(curl -s -o /dev/null -w ''%{http_code}'' localhost:8080)" != "302" ]]; do echo "zzz..."; sleep 5; done
docker-compose -f docker-compose-gitlab.yml logs -tf
# terminal 2
# curl -v localhost:8080/api/v4
# register new root user -> set password root123!
http://localhost:8080/
# login as root
http://localhost:8080/users/sign_in
# create a new user and check the 'admin' box
## goreleaser
## goreleaser@acme.com
http://localhost:8080/admin/users/new
# set the password cuz setting via email is not possible
# set to 'testpwd123!'
http://localhost:8080/admin/users/goreleaser/edit
# logout and login as goreleaser and maybe set a new password if requested
http://localhost:8080/users/sign_in
# create an access token with API scope
http://localhost:8080/profile/personal_access_tokens
# export it
export GITLAB_TOKEN="abc"
# create 'acceptance-tests' (private) and 'homebrew-tap' (public) repository 
http://localhost:8080/projects/new
# build the goreleaser binary you want to test
cd ../goreleaser && make build && mv goreleaser ../acceptance-tests && cd ../acceptance-tests
# tag the repo
git tag 0.1.0
# clean remotes
git remote rename origin xyz
# add the new origin
git remote add origin http://goreleaser:testpwd123\!@localhost:8080/goreleaser/acceptance-tests.git
# push it 
git push origin chore-initial-setup
# publish locally
./goreleaser --config=./goreleaser-gitlab-local.yml --debug --rm-dist
```

