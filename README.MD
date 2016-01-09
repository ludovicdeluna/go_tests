Welcome to Hello Sample
=======================

This is a simple markdown file for eventually goes into Github !

Thanks.

Sample Redirect :
See go-import, realy useful

curl https://golang.org/x/tools/godoc
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


curl https://golang.org/x/tools/
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

    <meta name="go-import" content="toto.com/project.v1 git bitbucket.com/toto/project.v1">
    ==> Fetch bitbucket.com/toto/project.v1 in https mode and get "sub"

    <meta name="go-import" content="toto.com/project.v1 git toto.com/~ludo/project.v1">
    ==> Fetch toto.com/~ludo/project.v1 in https mode and get "sub"

    Be aware Go will try also the url "https://toto.com/project.v1", because it's the root repo
    and the current url is only a sub-directory of it.
    ==> Need to get exactly the same Metatag (case of the last sample):
    <meta name="go-import" content="toto.com/project.v1 git toto.com/project.v1">
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
