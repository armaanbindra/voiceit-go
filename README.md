#VoiceIt Go Package
A Package for using the VoiceIt Rest API.

##Installation
To add the package to your Go Workspace, first make sure that your GOPATH and workspace are properly set up.  Then, go to your terminal and run
```
go get github.com/voiceittech/voiceit-go/voiceit
```

##Usage
To use the VoiceIt Go package, import it at the beginning of your program and initialize a VoiceIt object as shown below:

```go
import "github.com/voiceittech/voiceit-go/voiceit"

myVoiceIt := voiceit.New("1d6361f81f3047ca8b0c0332ac0fb17d")
```

Finally use all other API Calls as documented on the [VoiceIt API Documentation](https://siv.voiceprintportal.com/getstarted.jsp#apidocs) page.
