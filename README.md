# Nintaco Go API - Castlevania Weapons Example

### About

[The Nintaco NES/Famicom emulator](https://nintaco.com/) provides [a Go API](https://github.com/meatfighter/nintaco-go-api) that enables externally running programs to control the emulator at a very granular level. This example enhances *Castlevania*, providing the ability to change the subweapon at the press of a button.

### Launch

1. Start Nintaco and open *Castlevania*.
2. Open the Start Program Server window via Tools | Start Program Server...
3. Press Start Server.
4. From the command-line, launch this Go program.
5. Begin the game by pressing the Start gamepad button (by default its mapped to the Enter key).
6. While playing, change the subweapon at will by pressing Select (by default its mapped to the apostrophe key).

Every time Select is pressed, the subweapon will cycle to a different one. Plus, you'll gain the triple shot, 99 hearts, maximum hit points, the long chain whip and 99 lives. Not only that, one of the available subweapons is the rosary that will destroy all enemies on the screen. However, since it wasn't designed as a subweapon, it may cause the game to crash if activated on certain screens.   

