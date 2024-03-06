

## Git Package Usage
#RT-Datacore
go get git.frontiir.net/sa-dev/rtdatacore@v0.1.0
password : EKJs3ebXT8k9TbUysLTQ

##Local extension Usage
create folder name extension in current dir
##RT-Datacore
go to extension and git clone git@git.frontiir.net:sa-dev/rtdatacore.git
go mod file open 

    module core
    go 1.22.0
    replace git.frontiir.net/sa-dev/rtdatacore => ./extension/rtdatacore

add gitignore 
## ignore env , bin , extension
*.env
extension/
bin/

## SetUp Build and DockerContainer Up
make setup && make build && make up


## Run swagger file without docker 
make run


Usage
Initialize your module
go mod init core

Get the go-lib module
Note that you need to include the v in the version tag.

go get git.frontiir.net/sa-dev/rtdatacore@v0.1.0

Git Usage
Initialize the project for git:
git init
Add the files to git:
Note that you need to include the period at the end of this command:
git add .
Commit the files
git commit -m "init commit"
Create the repo
Go to your git cloud provider (like github) and create a new repo
Follow the instructions of your provider for creating the new repo


…or create a new repository on the command line

git init
git commit -m "init commit"
git add .
git branch -M main
git remote add origin https://github.com/banyar/local.git
git push -u origin main


…or push an existing repository from the command line
git remote add origin https://github.com/banyar/local.git
git branch -M main
git push -u origin main
…or import code from another repository
You can initialize this repository with code from a Subversion, Mercurial, or TFS project


replace git.frontiir.net/sa-dev/rtdatacore => ./extension/rtdatacore


git remote add origin https://git.frontiir.net/sa-dev/rtdatacore.git
git branch -M main
git push -u origin main

git tag -d v0.1.0
git tag v0.1.0 HEAD -m "Version v0.1.0 released"
git push origin --tags


Tesing Ticket ID 3924319



