# go-boy
Gameboy emulator in Go 

## To-do

- [ ] Write interpreter for gameboy cpu (some flavor of Z80 family, refer to http://marc.rawer.de/Gameboy/Docs/GBCPUman.pdf)
- [ ] Design how hardware components will interact between CPU
- [ ] Figure out how to output graphics (literally have no idea about this)
- [ ] Figure out how to output sound (no idea about this either, https://gamedev.stackexchange.com/questions/62049/do-games-use-multiple-threads-to-play-music?)
- [ ] Add a real time hex-editor, debugger, and if possible Ghidra integration for dynamic analysis


## Directions 

Project was inspired by the lack of tools to dynamically debug and mark RAM sections for reverse engineering. This emulator hopes to do the patchwork. Unfortunately there is no full-fledged tutorial on how to write gameboy emulator. I'll have to piece together the information needed for this task based on other platforms. 

- https://www.youtube.com/watch?v=nViZg02IMQo&list=PLrOv9FMX8xJHqMvSGB_9G9nZZ_4IgteYf&index=1 (How to make an NES emulator)
- https://www.youtube.com/watch?v=ohATX8biqs4&list=PLEDYD952XcG7RwWeshln0dGYdfwqt24Ps&index=1 (GB simulator in Python)
- https://github.com/gbdev/awesome-gbdev (Compilation of resouces about GB internals)
- http://pepijndevos.nl/sha2017/workshop.pdf (GB hacking)
- https://blog.rekawek.eu/2017/02/09/coffee-gb/ (Some blog post about GB emulation)
- http://www.emulator101.com/ (General tutorial on emulation)
- http://imrannazar.com/GameBoy-Emulation-in-JavaScript (In JS, lol)
- https://bitbucket.org/milosonator/stygiagb/src/master/ (in Python, same thing as above)
