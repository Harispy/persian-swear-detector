# persian-swear-detector

A very simple, high performance, ready to go Persian swear detector for golang.
also support some, not very complicated, normalization specially for persian

## Installation

```bash
go get github.com/harispy/persian-swear-detector
```


## Usage

### use with custom list of swear words
```golang
swearDetector := detector.NewCustomSwearWordDetector([]string{"fuck","احمق"})
hasSwear := swearDetector.ContainsSwearWord("تو احمق هستی") // returns true
hasSwear := swearDetector.ContainsSwearWord("عالی بود") // returns false
hasSwear := swearDetector.ContainsSwearWord("f*u*c*k you") // returns true
```
