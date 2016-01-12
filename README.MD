Sample in go
============

## Get this sample on your computer (Linux/Mac):

**Install and configure your environment**
* Install GIT (recommended, already installed on MacOS)
* Install go 1.5 or later (type `go version` to check it)
* Create your working folder and source : `mkdir -p ~/work/src`
* Declare your working folder to Go : `export GOPATH=~/work`
* You need vendoring activated (Go1.5 only) : `export GO15VENDOREXPERIMENT=1`

**Install the package**
* Go into your workspace (cd $GOPATH)
* You have GIT : `go get github.com/ludovicdeluna/go_tests/...`  
or
* You haven't GIT :
  * Download file from : https://github.com/ludovicdeluna/go_tests/archive/master.zip
  * Create folder path : `mkdir -p $GOPATH/src/github.com/ludovicdeluna`
  * Unpack the zip into it, this should create a new folder. Rename it to go_tests.
  * Compile and install : `go install github.com/ludovicdeluna/go_tests/...`

Test any binaries in $GOPATH/bin

### What it is for...

In this sample, we demonstrate how :
* Use package name ready to be published on public DCVS (like github)
* Use Vendored package for the "hello" command (need go1.5)
* Make versioned packages as we do with ui.v1 et ui.v2. Easy, thanks to Go.
* Use build directive into go source :
  * Say to only build hello.go when you have a Go version 1.5 or later.
  * Build specific platform code in ui.v2 package : Linux / Windows.

Developments
============

For developments :
* Fork this repository
* Declare your working folder to Go :  
  `export GOPATH=~/work`
* Get your local copy using GIT :  
  `git clone git_path_to_your_fork $GOPATH/src/github.com/ludovicdeluna/go_tests`
* If you want : Add upstream to track my changes :  
  `cd $GOPATH/src/github.com/ludovicdelun/go_tests`  
  `git remote add -t master upstream https://github.com/ludovicdeluna/go_tests.git`

For reuse :
* Consider perhaps a dual GOPATH (your work goes to go/work) :  
  `export GOPATH=~/go/deps:~/go/work`
* Get the package without install it:  
  `go get -d github.com/ludovicdeluna/go_tests/...`
* Move the package into vendor, or simply sync deps with vendor folder :
  * Sync : `rsync -av ~/go/deps/src ~/go/work/src/ROOT_OF_YOUR_PROJECT/vendor`
  * Reset deps folder : `rm -rf ~/go/deps/src`

> To embed any package, always use vendor folder into the root of your project,
> and activate vendor flag (Go1.5 only). Clean your vendor : Keep only
> packages you actively use (or required by them).


Other package techniques for Go
===============================

Sample Redirect :
See go-import, realy useful

curl https://golang.org/x/tools/godoc
```html
<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta name="go-import" content="golang.org/x/tools git https://go.googlesource.com/tools">
<meta name="go-source" content="golang.org/x/tools https://github.com/golang/tools/ https://github.com/golang/tools/tree/master{/dir} https://github.com/golang/tools/blob/master{/dir}/{file}#L{line}">
<meta http-equiv="refresh" content="0; url=https://godoc.org/golang.org/x/tools/godoc">
</head>
<body>
Nothing to see here; <a href="https://godoc.org/golang.org/x/tools/godoc">move along</a>.
</body>
</html>
```

curl https://golang.org/x/tools/
```html
<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta name="go-import" content="golang.org/x/tools git https://go.googlesource.com/tools">
<meta name="go-source" content="golang.org/x/tools https://github.com/golang/tools/ https://github.com/golang/tools/tree/master{/dir} https://github.com/golang/tools/blob/master{/dir}/{file}#L{line}">
<meta http-equiv="refresh" content="0; url=https://godoc.org/golang.org/x/tools/">
</head>
<body>
Nothing to see here; <a href="https://godoc.org/golang.org/x/tools/">move along</a>.
</body>
</html>
```


import "package-name" ->folder to store package

Know DCVS host:  
import "github.com/toto/project/sub" ->github.com/toto/project/sub  
==> Fetch github.com/toto/project in https mode and get "sub".

Know DCVS host & Versioned:  
import "github.com/toto/project.v1/sub" ->github.com/toto/project.v1/sub  
==> Fetch github.com/toto/project.v1 in https mode and get "sub".

Unknow DCVS host:  
import "toto.com/project.git/sub" ->toto.com/project/sub  
==> Fetch toto.com/project in https mode and get "sub".

Unknow DCVS host & Versioned:  
import "toto.com/project.v1.git/sub" ->toto.com/project.v1/sub  
==> Fetch toto.com/project.v1 in https mode and get "sub".

Proxied & Versioned:  
import "toto.com/project.v1/sub" ->toto.com/project.v1/sub  
==> Get HTML from "https://toto.com/project.v1/sub" and follow the Metatag go-import.

```
    <meta name="go-import" content="toto.com/project.v1 git bitbucket.com/toto/project.v1">
```
==> Fetch bitbucket.com/toto/project.v1 in https mode and get "sub"

```
    <meta name="go-import" content="toto.com/project.v1 git toto.com/~ludo/project.v1">
```
==> Fetch toto.com/~ludo/project.v1 in https mode and get "sub"

Be aware Go will try also the url "https://toto.com/project.v1", because it's the root repo
and the current url is only a sub-directory of it.  
==> Need to get exactly the same Metatag (case of the last sample):

```
    <meta name="go-import" content="toto.com/project.v1 git toto.com/project.v1">
```

==> In this case, Go had confirmation of the root namespace, and everything works.

Cas d'un package contenant une reference externe (vendor), dev externe :
go get pour récupérer la source et ensuite supprimer le depôt embarqué (et tout
sous-dépot trouvé). Ne pas forker les packages dans vendor (inutile). 2 choix
pour gérer le source :

- Si c'est unique au projet: L'envoyer tel quel dans le depot du programme.
- Si c'est partageable entre plusieurs projets : Créer un depot git depuis le
dossier du module "vendorisé "+ commit initial et le référencer via un module
git depuis le projet. Les modules git sont gérés par go get depuis 1.5 (si flag
vendor activé).
