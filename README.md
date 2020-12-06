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
# create an access token andsave as env GITEA_TOKEN
# publish locally
./goreleaser --config=./goreleaser-gitea-local.yml --debug --rm-dist
# if it fails 
git tag -d 0.1.0
# delete the tag by hand at
http://localhost:3000/goreleaser/acceptance-tests/releases
```

## gitlab
```sh
# start gitlab
docker-compose -f docker-compose-gitlab.yml up -d
# follow log and wait about 12 minutes until every
docker-compose -f docker-compose-gitlab.yml logs -tf
# terminal 2
# curl -v localhost:8080/api/v4
# register new root user -> set password root123!
http://localhost:8080/
# login as root
http://localhost:8080/users/sign_in
# create a new user
## goreleaser
## goreleaser@acme.com
http://localhost:8080/admin/users/new
# set the password cuz setting via email is not possible
# set to 'testpwd123!'
http://localhost:8080/admin/users/goreleaser/edit
# logout and login as goreleaser
http://localhost:8080/users/sign_in
# create 'acceptance-tests' (private) and 'homebrew-tap' (public) repository 
http://localhost:8080/projects/new
# build the goreleaser binary you want to test
cd ../goreleaser && make build && mv ../goreleaser/goreleaser .
# tag the repo
git tag 0.1.0
# create an access token andsave as env GITEA_TOKEN
# publish locally
./goreleaser --config=./goreleaser-gitlab-local.yml --debug --rm-dist
```

