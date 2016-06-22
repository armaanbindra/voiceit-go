package voiceit-go
import "fmt"
import "net/http"
import "crypto/sha256"
import "io"
import "encoding/hex"
import "io/ioutil"
import "bytes"

type VoiceIt struct {
    devID string
}

func new(developerID string) *VoiceIt{
    return &VoiceIt{
        devID: developerID,
    }
}

func (v *VoiceIt) createUser(mail string, passwd string, firstName string, lastName string, phones ...string) string{
    hasher := sha256.New()
    p1,p2,p3 := "", "", ""
    if len(phones) > 0 {
        p1 = phones[0]
    } else {
        p1 =""
    }
    if len(phones) > 1 {
        p2 = phones[1]
    } else{
        p2 =""
    }
    if len(phones) > 2 {
        p3 = phones[2]
    } else{
        p3 =""
    }
    client := &http.Client{}
    io.WriteString(hasher, passwd)
    shapass := hex.EncodeToString(hasher.Sum(nil))
    req, err := http.NewRequest("POST", "https://siv.voiceprintportal.com/sivservice/api/users", nil)
    req.Header.Add("Accept" , "application/json")
    req.Header.Add("VsitEmail" , mail)
    req.Header.Add("VsitPassword" , shapass)
    req.Header.Add("VsitDeveloperId" , v.devID)
    req.Header.Add("VsitFirstName" , firstName)
    req.Header.Add("VsitLastName" , lastName)
    req.Header.Add("VsitPhone1" , p1)
    req.Header.Add("VsitPhone2" , p2)
    req.Header.Add("VsitPhone3" , p3)
    req.Header.Add("PlatformID" , "24")
    resp, err := client.Do(req)
    if err != nil {
        fmt.Printf("%s\n", "ERROR!")
    }
    defer resp.Body.Close()
    reply, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("%s\n", "ERROR!")
    }
    result := string(reply[:len(reply)])
    return result
}

func (v *VoiceIt) getUser(mail string, passwd string) string{
    hasher := sha256.New()
    client := &http.Client{}
    io.WriteString(hasher, passwd)
    shapass := hex.EncodeToString(hasher.Sum(nil))
    req, err := http.NewRequest("GET", "https://siv.voiceprintportal.com/sivservice/api/users", nil)
    req.Header.Add("Accept" , "application/json")
    req.Header.Add("VsitEmail" , mail)
    req.Header.Add("VsitPassword" , shapass)
    req.Header.Add("VsitDeveloperId" , v.devID)
    req.Header.Add("PlatformID" , "24")
    resp, err := client.Do(req)
    if err != nil {
        fmt.Printf("%s\n", "ERROR!")
    }
    defer resp.Body.Close()
    reply, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("%s\n", "ERROR!")
    }
    result := string(reply[:len(reply)])
    return result
}

func (v *VoiceIt) setUser(mail string, passwd string, firstName string, lastName string, phones ...string) string{
    hasher := sha256.New()
    p1,p2,p3 := "", "", ""
    if len(phones) > 0 {
        p1 = phones[0]
    } else {
        p1 =""
    }
    if len(phones) > 1 {
        p2 = phones[1]
    } else{
        p2 =""
    }
    if len(phones) > 2 {
        p3 = phones[2]
    } else{
        p3 =""
    }
    client := &http.Client{}
    io.WriteString(hasher, passwd)
    shapass := hex.EncodeToString(hasher.Sum(nil))
    req, err := http.NewRequest("PUT", "https://siv.voiceprintportal.com/sivservice/api/users", nil)
    req.Header.Add("Accept" , "application/json")
    req.Header.Add("VsitEmail" , mail)
    req.Header.Add("VsitPassword" , shapass)
    req.Header.Add("VsitDeveloperId" , v.devID)
    req.Header.Add("VsitFirstName" , firstName)
    req.Header.Add("VsitLastName" , lastName)
    req.Header.Add("VsitPhone1" , p1)
    req.Header.Add("VsitPhone2" , p2)
    req.Header.Add("VsitPhone3" , p3)
    req.Header.Add("PlatformID" , "24")
    resp, err := client.Do(req)
    if err != nil {
        fmt.Printf("%s\n", "ERROR!")
    }
    defer resp.Body.Close()
    reply, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("%s\n", "ERROR!")
    }
    result := string(reply[:len(reply)])
    return result

}

func (v *VoiceIt) deleteUser(mail string, passwd string) string{
    hasher := sha256.New()
    client := &http.Client{}
    io.WriteString(hasher, passwd)
    shapass := hex.EncodeToString(hasher.Sum(nil))
    req, err := http.NewRequest("DELETE", "https://siv.voiceprintportal.com/sivservice/api/users", nil)
    req.Header.Add("Accept" , "application/json")
    req.Header.Add("VsitEmail" , mail)
    req.Header.Add("VsitPassword" , shapass)
    req.Header.Add("VsitDeveloperId" , v.devID)
    req.Header.Add("PlatformID" , "24")
    resp, err := client.Do(req)
    if err != nil {
        fmt.Printf("%s\n", "ERROR!")
    }
    defer resp.Body.Close()
    reply, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("%s\n", "ERROR!")
    }
    result := string(reply[:len(reply)])
    return result

}

func (v *VoiceIt) createEnrollment(mail string, passwd string, pathToEnrollmentWav string, contentLanguage ... string) string {
    hasher := sha256.New()
    contentLang :=""
    if len(contentLanguage) > 0 {
        contentLang = contentLanguage[0]
    } else {
        contentLang =""
    }
    wavData, err := ioutil.ReadFile(pathToEnrollmentWav)
    if err != nil {
        fmt.Printf("%s\n", "ERROR!")
    }
    client := &http.Client{}
    io.WriteString(hasher, passwd)
    shapass := hex.EncodeToString(hasher.Sum(nil))
    req, err := http.NewRequest("POST", "https://siv.voiceprintportal.com/sivservice/api/enrollments", bytes.NewReader(wavData))
    req.Header.Add("Accept" , "application/json")
    req.Header.Add("VsitEmail" , mail)
    req.Header.Add("VsitPassword" , shapass)
    req.Header.Add("VsitDeveloperId" , v.devID)
    req.Header.Add("ContentLanguage" , contentLang)
    req.Header.Add("PlatformID" , "24")
    resp, err := client.Do(req)
    if err != nil {
        fmt.Printf("%s\n", "ERROR!")
    }
    defer resp.Body.Close()
    reply, err:= ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("%s\n", "ERROR!")
    }
    result := string(reply[:len(reply)])
    return result
}

func (v *VoiceIt) createEnrollmentByWavURL(mail string, passwd string, urlToEnrollmentWav string, contentLanguage ... string) string {
    hasher := sha256.New()
    contentLang :=""
    if len(contentLanguage) > 0 {
        contentLang = contentLanguage[0]
    } else {
        contentLang =""
    }
    client := &http.Client{}
    io.WriteString(hasher, passwd)
    shapass := hex.EncodeToString(hasher.Sum(nil))
    req, err := http.NewRequest("POST", "https://siv.voiceprintportal.com/sivservice/api/enrollments/bywavurl", nil)
    req.Header.Add("Accept" , "application/json")
    req.Header.Add("VsitEmail" , mail)
    req.Header.Add("VsitPassword" , shapass)
    req.Header.Add("VsitDeveloperId" , v.devID)
    req.Header.Add("VsitwavURL", urlToEnrollmentWav)
    req.Header.Add("ContentLanguage" , contentLang)
    req.Header.Add("PlatformID" , "24")
    resp, err := client.Do(req)
    if err != nil {
        fmt.Printf("%s\n", "ERROR!")
    }
    defer resp.Body.Close()
    reply, err:= ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("%s\n", "ERROR!")
    }
    result := string(reply[:len(reply)])
    return result
}

func (v *VoiceIt) getEnrollments(mail string, passwd string) string{
    hasher := sha256.New()
    client := &http.Client{}
    io.WriteString(hasher, passwd)
    shapass := hex.EncodeToString(hasher.Sum(nil))
    req, err := http.NewRequest("GET", "https://siv.voiceprintportal.com/sivservice/api/enrollments", nil)
    req.Header.Add("Accept" , "application/json")
    req.Header.Add("VsitEmail" , mail)
    req.Header.Add("VsitPassword" , shapass)
    req.Header.Add("VsitDeveloperId" , v.devID)
    req.Header.Add("PlatformID" , "24")
    resp, err := client.Do(req)
    if err != nil {
        fmt.Printf("%s\n", "ERROR!")
    }
    defer resp.Body.Close()
    reply, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("%s\n", "ERROR!")
    }
    result := string(reply[:len(reply)])
    return result
}

func (v *VoiceIt) deleteEnrollment(mail string, passwd string, enrollmentId string) string{
    hasher := sha256.New()
    client := &http.Client{}
    io.WriteString(hasher, passwd)
    shapass := hex.EncodeToString(hasher.Sum(nil))
    req, err := http.NewRequest("DELETE", "https://siv.voiceprintportal.com/sivservice/api/enrollments/"+enrollmentId, nil)
    req.Header.Add("Accept" , "application/json")
    req.Header.Add("VsitEmail" , mail)
    req.Header.Add("VsitPassword" , shapass)
    req.Header.Add("VsitDeveloperId" , v.devID)
    req.Header.Add("PlatformID" , "24")
    resp, err := client.Do(req)
    if err != nil {
        fmt.Printf("%s\n", "ERROR!")
    }
    defer resp.Body.Close()
    reply, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("%s\n", "ERROR!")
    }
    result := string(reply[:len(reply)])
    return result

}

func (v *VoiceIt) authentication(mail string, passwd string, pathToAuthenticationWav string, confidence string,contentLanguage ... string) string {
    hasher := sha256.New()
    contentLang :=""
    if len(contentLanguage) > 0 {
        contentLang = contentLanguage[0]
    } else {
        contentLang =""
    }
    wavData, err := ioutil.ReadFile(pathToAuthenticationWav)
    if err != nil {
        fmt.Printf("%s\n", "ERROR!")
    }
    client := &http.Client{}
    io.WriteString(hasher, passwd)
    shapass := hex.EncodeToString(hasher.Sum(nil))
    req, err := http.NewRequest("POST", "https://siv.voiceprintportal.com/sivservice/api/authentications", bytes.NewReader(wavData))
    req.Header.Add("Accept" , "application/json")
    req.Header.Add("VsitEmail" , mail)
    req.Header.Add("VsitPassword" , shapass)
    req.Header.Add("VsitDeveloperId" , v.devID)
    req.Header.Add("VsitConfidence", confidence)
    req.Header.Add("ContentLanguage" , contentLang)
    req.Header.Add("PlatformID" , "24")
    resp, err := client.Do(req)
    if err != nil {
        fmt.Printf("%s\n", "ERROR!")
    }
    defer resp.Body.Close()
    reply, err:= ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("%s\n", "ERROR!")
    }
    result := string(reply[:len(reply)])
    return result
}

func (v *VoiceIt) authenticationByWavURL(mail string, passwd string, urlToAuthenticationWav string, confidence string, contentLanguage ... string) string {
    hasher := sha256.New()
    contentLang :=""
    if len(contentLanguage) > 0 {
        contentLang = contentLanguage[0]
    } else {
        contentLang =""
    }
    client := &http.Client{}
    io.WriteString(hasher, passwd)
    shapass := hex.EncodeToString(hasher.Sum(nil))
    req, err := http.NewRequest("POST", "https://siv.voiceprintportal.com/sivservice/api/authentications/bywavurl", nil)
    req.Header.Add("Accept" , "application/json")
    req.Header.Add("VsitEmail" , mail)
    req.Header.Add("VsitPassword" , shapass)
    req.Header.Add("VsitDeveloperId" , v.devID)
    req.Header.Add("VsitwavURL", urlToAuthenticationWav)
    req.Header.Add("VsitConfidence", confidence)
    req.Header.Add("ContentLanguage" , contentLang)
    req.Header.Add("PlatformID" , "24")
    resp, err := client.Do(req)
    if err != nil {
        fmt.Printf("%s\n", "ERROR!")
    }
    defer resp.Body.Close()
    reply, err:= ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("%s\n", "ERROR!")
    }
    result := string(reply[:len(reply)])
    return result
}
