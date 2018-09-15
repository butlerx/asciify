# ascii-logo

ascii-logo converts images too ascii version.

## Example

_Before_

![CoderDojo Octocat](./octocat.png)

_After_

```
             :,            ~=~:                                          :~==
             ~,           ~MMMMN87=                                  ~IODMMMMI
             ~,           ZMMMMMMMMNZ=        ,,::::::::,,        ~7DMMMMMMMMN
             ~,           NMMMMMMMMMMMDII$ZO8DDNNNNNNNNNNNDD8O$7I8MMMMMMMMMMMM=
             ::          :MMMMMMMMMMMNNNNNNNNNNNDDDDDNNNNNNMMMMMMMNMMMMMMMMMMM?
             ::          ~MMMMMMMMMNDDDDDDDNNNNNNNNNNNNNNNNNNNNNNNNNMMMMMMMMMMI
             ,:          ,MMMMMMMMNDDDDDDDDDDDDDDDDDDDDDDDDDDDNNNNNNNNMMMMMMMM*
             ,~           8NNNNNNNNNNNNNNNNDDDDDDDDDDDDDDDDDDDDDNNNNNNNNNNNNMM,
              ~,        ,ZNDNNNNNNNNNNNNNNNNNDDNNNNNDDDDDDDDNNNNNNNNNNNNNNNNNN8:
              ~,       :DNNNNDDDDDDDNNNDDDDDNNNDDDNNNNNNNDDDDDDDDDDDDDDNNNNNNNMN=
              ::      :DNDDDDNNNNNNNNNNNNNNNDDDDDDDDDDDDDDDDDDDDDDDDDDDDNNNNNNNMM=
              ::      ONDDDDDDDDDDDDDDDDNNNNNNNNDDDDDNNNNNNNNNNNNNNNNNNNNNNNNNNNMN,
              ,~     IMDDDDDDDNNNNNNNDDDDDDDDDDDDDDDNNNNNNNNNNNNNNNNNNNNNNNNNNNNNMZ
               ~,   ,DNDDDNNNNNNDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDNNNNNNNM:
               :,   *NDNNNNNDDDDDDDDDDDDDDDDDDDDDDDDDDDDDNNNNNNNNNNNNNDDDDDNNNNNNNMI
               ::   7NNNNDDDDDDDDNNNNNNNMMMMMMMMMMMMMMMMMMMMMMMMMMMNNNNNNNDDNNNNNNM$
               ,~   7MNDDDDDDDDNNNND88OOOOOOOOOOOOOOOOOOOOOOOOOOOOO8DNNNNNNDNNNNNNN:
                ~,  :DNDDDDDDNNN87*=~::::~::::::::::::::::::::::::~~=*IONNNNNNNNNNN,
                ::   ONDNDDDDNN$~:::::::::~~::::::::::::::::::::::~::::~7DNNNNNNNMD
                ,:   ZNDDDDDDN$:::~: ~7I?=~=~:::::::::::::::::~*I77: :~::7NNNNNNNMO
                 ~,  $MDDDDDND=::::  7$$$$I ,:::::::::::::::, I$$7$7 ,~::~8NNNNNNM7
                 :,  ~DDDDDDN8~:::: ,$77777  ~:::::::::::::~  77777$, ::::ZNNNNNNM~
       ,,,:::::~~=*~~:ZNDDDDND=::::  7$77$I ,~:::::::::::::~, I$77$7 ,~:::ONNNNNM8:~~~~~~:::::,,,
  ~====~~~::,,,,,:=:::~NDDDDDNI::::: :7$$I, ::::~~====~~::::: ,I$$7: ::::*NNNNNNM*:::,,,,,,,:::~~~===~
          ,,::~~~~**~~~ZMDDDDN8=::::,  ,,  ~*I$O8DDNNDD8O$I*~,  ,,  :::::ONNNNNM8~~~~~~~~~:::,
  ,:~~===~~::,,   ,:    ZMDDNDN8=~~~~~~~*7ODNMMMNNNNNNNNMMMNDO$?~::~~~~~ZNNNNNMD:        ,,::~~~===~:,
  ::,,             =*~   7NNDDDNDD88DDNNMMMNNNDDDDNNNNNDDDDNNMMMMNDDDDDDNNNNMMZ,                   ,:~
                   $M7    ~ZNNNDDNNNNNNNDDDNNNNNNNNNNNNNNNNNNNDNNNNMMMNNNMMMO*
                   :D8::,   :I8NMNNNDDDNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNMMMMD$~
                   ~DMMM8      ~?$8NNMMMNNNDDDDDDDDDDDDDDDNNNNMMMMMMNDZI~
                   ZMMMM8*         :=?7ZOODNNNNNNNNNNNNNNNNNNNOOZ7?=:
                   ,78NNMMZ~             IDNNNNNNNNMMNMMNNNNNN$,
                     :DNNOOO?           ,ODNNDDDDDDMMND7ONDNNDN~
                      INO$ZZZ?          ?Z$8DNDDDDNNNMI*DMDNOO8$
                       ,:ZOZZZ$*:,    :=ZZZO8NNDNNNNDDDNMNNO$O8D~
                         ~$ZZ$ZZZ$ZZ$$ZZZZZ88ONNNNDDNNNNDNOZZZ8D?
                           ?ZZ$$$ZZZZZ$$ZZZ8OIDDDDDDDDDDDN77ZZ8D7
                            ~7OOOZZZZZOOZ$ZDZ?NDDD8DD8DDDN7?ZZ8D$
                              :*7$ZZZ$7$Z$ZDO~NNNDN8ZNDNNM?IZZ8DI
                                       ~ZZO8O*NDDDMZINDNNMI7OZ8D?
                                       ~Z$Z8O?NDDNMO7NDNNM7IZZ8DI
                                       ~Z$ZDZ*NDDNMZ7NDNNM7?ZZ8DI
                                       =OZO8O~DNNNM7*MNNNM=7OO8DI
                                 ,,,,,,:NMMMD,788DN?~8O8D8,ZMMMM=,,,,,,
                             ,,,,,,,,, *MMMM8 IZO8DI~ZZO8O $MMMM7 ,,,,,,,,,
                          ,,,,,,,,,, ,?MMMMM=,$ZO8D=:ZZZ88::NMMMM7: ,,,,,,,,,,
                         ,,,,,,,,,,,$NMMMNZ= *ZZ887,,*ZZZ8$ :$DMMMMO:,,,,,,,,,,
                        ,,,,,,,,,,,,*??*~: ,78O8DD*,,~O8OZDZ: ,~**??,,,,,,,,,,,,
                        ,,,,,,,,,,,,    ,,,=OOZ$I=:,,,~?7$OO?,,,    ,,,,,,,,,,,,
                         ,,,,,,,,,,,,,,,,,,,,,,  ,,,,,, ,,,,,,,,,,,,,,,,,,,,,,,
                           ,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
                             ,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
                                 ,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,,
                                         ,,,,,,,,,,,,,,,,,,,,,,
```

## Usage

```bash
go get https://github.com/butler/ascii-logo/cmd/asciify
asciify https://octodex.github.com/images/dojocat.jpg -p -w 75
```

## Features

* Both remote and local images
* scaling images based on width
* Injecting into readme(or any markdown file) with the `{{ .Image }}` keyword
* Printing to terminal
