## Train (your algo)

*work in progress*

Train your machine learning algo with games that can be played in parallel. Or, you know, if you wanted the challenge of playing 1000 games of knock-off minesweeper simultaneously, you could do that too. :smile:

All games speak the lingua franca that is `[]byte`, so your ML algo must generate some msg in `[]byte` in order to play. I have chosen this data structure for quite a few reasons, and would love to discuss it further if you have any ideas on the topic. Each game is fairly customizable, particularly if you want to scale difficulty i.e. a 4x4 minesweeper game with a single bomb, or poker with a deck of 1000 cards and complex matching system. 

Please feel free to open an issue for any reason 