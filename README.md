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
# publish locally
./goreleaser --config=./goreleaser-gitea-local.yml --debug --rm-dist
# if it fails 
git tag -d 0.1.0
# delete the tag by hand at
http://localhost:3000/goreleaser/acceptance-tests/releases
```

## gitlab
TBD