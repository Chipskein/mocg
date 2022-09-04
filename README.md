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
    <img src="https://github.com/Chipskein/mocg/blob/main/docs/demo.gif" style="width:650px"/>
  </div>
  
  
  
## Description
  Music on console with Go. Terminal music player inspired in MOC player
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
  | Q               	| Exit MOCG UI       	                    |
  | H               	| Toggle Switch to Show Hidden FIles      |
  | M               	| Mute or Unmute                          |
  | UP,K              | Move Up File Selector                   |
  | DOWN,J            | Move Down File Selector                 |
  | Home              | Move To Top List                        |
  | End               | Move To Bottom List                     |
  | PageUp            | Move Up Half Page                       |
  | PageDown          | Move Down Half Page                     |
  
* #### Supported File Formats
  * ##### FLAC
  * ##### OGG
  * ##### MP3
  * ##### WAV
## How to install
## How to run
## Music Samples Test
|                     **Music Name**                                  |   **Artist Name**    |
|---------------------------------------------------------------------|----------------------|
| [City of Drones](https://www.youtube.com/watch?v=qYTHZCBpycg)       |    WhiteBat      | 
| [Inamorata](https://www.youtube.com/watch?v=WzWSIvxEBrA)            |     Maruex       |  
| [Ninguém me ama](https://www.youtube.com/watch?v=iYENyuka2NQ)       |  Quarteto em Cy  |  
| [The Perfect Girl](https://www.youtube.com/watch?v=W5Sq71VTJ9Q)     |      Maruex      |  
| [りんごの唄](https://www.youtube.com/watch?v=OFXIXF_RYyw)             |     並木路子      |  
| [就寝御礼](https://www.youtube.com/watch?v=mEQZNRT6Pqk)               |     PSYQUI      |    

## Known Bugs
* #### Music Files with Sample's Rate != 48000 will play with strange behavior like speed up or high pitch
  This happens because of beep's package speaker.Init() function,reinit will PANIC on linux systems [Similar issue](https://github.com/faiface/beep/issues/146).A temporary solution is to use a Default Sample Rate and Init Speaker only one time
  
## Dependencies
* [beep](https://github.com/faiface/beep)
* [termui](https://github.com/gizak/termui)
## Screenshots
## License
* [MIT](https://raw.githubusercontent.com/Chipskein/mocg/main/LICENSE?token=GHSAT0AAAAAABXMZE7Q6FPS4YQZUACJXPSGYYNICUA)
