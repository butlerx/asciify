# ascii-logo

ascii-logo converts images too ascii version.

## Example

_Before_

![CoderDojo Octocat](./octocat.png)

_After_

```
     :      @@@@                  @@@@
     :,     @@@@@@              @@@@@@
      ,     @@@@@@########@@@@@@@@@@@@
      ,     @@@@@####@@@@@@@@#@@@@@@@@
      ,     @@@@################@@@@@@
      ,     #@@@@##@############@@@@@@
      ;    #@@#######@##@##########@@@#
      ;   :####@@@@################@@@@@
      :   ##########@@@##########@@@@@@@
      :   #####@####################@@@@
      :, ###@########################@@@@
       , ########@@@@@@@@@@@@@@@@@@##@@@@
       , ##########################@#@@@@
       ;  #####;::::::::::::::::;:;#@@@@@
       :  ####S:::::;::::::::::::::;#@@@
       :  ####;:: ??,::::::::,??? ::#@@@
        , ####:::????:::::::: ??? ::#@@@
        , ####::: ?? :::::::: ??? ::#@@@
        ; :###;:: ?? ::;??;::: ? :::@@@@
   #S   :  #@##::: ,?##@@@@@##, ,::#@@#      *@
        :,  #######@@########@@#####@@?
        ?#   ########@@@@@@@@###@@@@@,
         #    @##@@@########@@@@@@@@
         @@      @@@@@#####@@@@@@
         @#@        #@@@@@@@@@
          #%%      ;######@@#@S
          @%S%     %%##### @@%S
            %%%%%*%%%S#####@S%S
            ?%%%%%%%%S######;%S
             ,%%%%%%S #SS#S# %S
                SS %S+##@#@@ %S
                   %SS#@@##@ %S
                   %S ##@##@ %S
                  ,%SS##@##@,SS,
               ,,,,@@,SSS%%S,@@,,,,
             ,,,,,,@@,%SS%%S,@@,,,,,,
           ,,,,,,:@@@,%S,,%S,@@@S,,,,,,
           ,,,,,,,,,,#SS::%%S,,,,,,,,,,
           ,,,,,,,,,,#?,,,,;#,,,,,,,,,,
            ,,,,,,,,,,,,,,,,,,,,,,,,,,
             ,,,,,,,,,,,,,,,,,,,,,,,,
               ,,,,,,,,,,,,,,,,,,,,
                    ,,,,,,,,,,
```

## Usage

```bash
go get https://github.com/butler/ascii-logo/cmd/asciify
asciify https://octodex.github.com/images/dojocat.jpg -p -w 75
```

## Features

* Both remote and local images
* scaling images based on width
* Injecting into readme(or any markdown file) with the `{{.}}` keyword
* Printing to terminal
