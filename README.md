# CSGO Demo Tick Parser

[![Open In Colab](https://colab.research.google.com/assets/colab-badge.svg)](https://colab.research.google.com/drive/1CAh6jaZsfgrQEON3hMH_yB7Y_Fnx8PFP?usp=sharing]


The [awpy](https://github.com/pnxenopoulos/awpy) parser allows one to parse a CSGO demofile. However, it is not written to parse _every_ tick due to high memory requirements. To address this issue, this repository provides a standalone Go program that can be run to parse every tick in a CSGO demo. 

Beware that this can produce _very_ large files (around 1GB). The script should work in Google's Colab in a normal instance (see [this link](https://colab.research.google.com/drive/1CAh6jaZsfgrQEON3hMH_yB7Y_Fnx8PFP?usp=sharing)). 

To run the script, run

```go
go run parse_player_frames.go -demo this-is-your-demo.dem -filename player-frames.csv
```

Be sure to have Go installed.