# Blurbstone

Blurbstone – A bit blurry on the destination, but it’ll probably get you somewhere cozy.


Design brainstorm:

- two players play card game in duel style combat
- each player has deck of 30 cards
- each player has health and mana crystals
- at the beginnig of the match system throws a coin
- timer is ticking every round till player cannot react and next player will play
- players take turns
- time clock says the same for each round during the whole game
- players can chose in a data file (json) what type of cards they would like
- one player spins up a game and gets unique code for other player to join
- it spins up a server so two players can play against eachother
- players operate a game with command line commands
- players 2 initiates a game with python script join + ip + secret code that was generated during
- initiate by player one
- it probably doesnt need to be secure because code is interpreted and this game is going to be
- played only privately ... question is how to implement p2p connection
- players draw cards
- it should have config file for current game all nescessary data like timer, chosen hero, player name,
- health, cards chosen in deck and so on
- domain files with all classes, cards, descriptions and more
- snapshot state at the end of each command (play of the player), 
- not every command yields snapshot of the game ... because some commands are only informational
- i need to figure out UI and syncing mechanism
- syncing and snapshoting mechanism should probably exist on some server
- or you can do it with zero tier and host it on your local computer
- single threaded running server with nohup so if one player accidentaly cancels or has some form of
- signal loss he can still join a game
- game persists even if both players disconnect

## potential classes


player
game
deck
card
table
turn
mana
mana-crystal
hero
hero-health
hero-ability
hero-weapon
timer
spell
minion
minion-type
minon-attack
minion-health
minon-ability


## list of commands

## general commands
get-help - prints list of commands
start-game
leave-game
join-game "game number"
stop-game "game number"
forfeit-game
choose-deck "name of the deck"

## in-game commands
peek-hand
inspect-opponent
show-board
inspect-minion "number of minion"
minon-attack "number of minon" -> then it requests target number (opponent hero is 0, then his minions are indexed and then yours)
play-card "number of card" -> then it plays a card or requests a targe number
hero-power -> "or requersts target number if its like a mage or priest"
weapon-attack "target number"
