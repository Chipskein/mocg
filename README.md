# MOCG ( Music on console with Go )
<img src="https://images.squarespace-cdn.com/content/v1/5e10bdc20efb8f0d169f85f9/1590751925678-5XVSVXMC2BX38RNKKO19/music.png" style="width:160px;">

## Index
* [Demo](#demo)
   * [Alacritty](#alacritty)
    * [TTY](#tty)
* [Description](#description)
* [Explanation](#explanation)
    * [What's MOC](#what-is-moc)
    * [Keyboard Shortcuts](#keyboard-shortcuts)
    * [Config File](#config-file)
    * [Supported File Formats](#supported-file-formats)
* [How to install](#how-to-install)   
* [How to run](#how-to-run)
* [Music Samples Test](#music-samples-test)
* [Known Bugs](#known-bugs)
* [Dependencies](#dependencies)    
* [Screenshots](#screenshots)
* [License](#license)

## Demo
* #### Alacritty
* #### TTY
## Description
## Explanation
* #### What is MOC
* #### Keyboard Shortcuts
* #### Config File
* #### Supported File Formats
## How to install
## How to run
## Music Samples Test
## Known Bugs
* #### Music Files with Sample's Rate != 480000 will play with strange behavior like speed up or high pitch
  This happens because of beep's package speaker.Init() function,reinit will SIGFAULT on linux systems [Similar issue](https://github.com/faiface/beep/issues/146).A temporary solution is to use a Default Sample Rate and Init Speaker only one time
  
## Dependencies
* [beep](https://github.com/faiface/beep)
* [termui](https://github.com/gizak/termui)
## Screenshots
## License
* [MIT](https://raw.githubusercontent.com/Chipskein/mocg/main/LICENSE?token=GHSAT0AAAAAABXMZE7QAZAU3JB54VRCP754YYMFC7A)




  


