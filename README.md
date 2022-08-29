<div align="center">
  <h1>MOCG ( Music on console with Go )</h1>
  <img src="https://images.squarespace-cdn.com/content/v1/5e10bdc20efb8f0d169f85f9/1590751925678-5XVSVXMC2BX38RNKKO19/music.png" style="width:160px;">
</div>

## Index
* [Demo](#demo)
* [Description](#description)
* [Explanation](#explanation)
    * [What's MOC](#what-is-moc)
    * [Keyboard Shortcuts](#keyboard-shortcuts)
    * [Supported File Formats](#supported-file-formats)
* [How to install](#how-to-install)   
* [How to run](#how-to-run)
* [Music Samples Test](#music-samples-test)
* [Known Bugs](#known-bugs)
* [Dependencies](#dependencies)    
* [Screenshots](#screenshots)
* [License](#license)

## Demo
  
  <div align="center">
    <img src="https://raw.githubusercontent.com/Chipskein/mocg/main/screenshots/demo.gif">
  </div>
  
  
## Description
  
## Explanation
* #### What is MOC
  [Music On Console (MOC)](https://github.com/jonsafari/mocp) is an ncurses-based console audio player for Linux/UNIX.It is designed to be powerful and easy to use, with an interface inspired by the Midnight Commander console file manager. [[See more]](https://en.wikipedia.org/wiki/Music_on_Console)
* #### Keyboard Shortcuts
  | **Key** 	        | **Description**                         |
  |-----------------	|-----------------------------------------|
  | ENTER            	| Play Music File or Access Directory     |
  | SPACE            	| Pause or Resume Music Track             |
  | .               	| Volume Up       	                      |
  | ,               	| Volume Down                             |
  | q               	| Exit MOCG UI       	                    |
  | m               	| Mute or Unmute                          |
  | UP               	| Move Up File Selector                   |
  | DOWN              | Move Down File Selector                 |
  
* #### Supported File Formats
  * ##### FLAC
  * ##### OGG
  * ##### MP3
  * ##### WAV
## How to install
## How to run
## Music Samples Test
## Known Bugs
* #### Music Files with Sample's Rate != 480000 will play with strange behavior like speed up or high pitch
  This happens because of beep's package speaker.Init() function,reinit will PANIC on linux systems [Similar issue](https://github.com/faiface/beep/issues/146).A temporary solution is to use a Default Sample Rate and Init Speaker only one time
  
## Dependencies
* [beep](https://github.com/faiface/beep)
* [termui](https://github.com/gizak/termui)
## Screenshots
## License
* [MIT](https://raw.githubusercontent.com/Chipskein/mocg/main/LICENSE?token=GHSAT0AAAAAABXMZE7Q6FPS4YQZUACJXPSGYYNICUA)
