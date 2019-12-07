# Spokesman

command line tool that synthesizes speech from text using Google Cloud Text-to-Speech.
 
### set up

```
export GOOGLE_APPLICATION_CREDENTIALS=~/credential.json
```

### Usage
 
```
$ ./spokenman <any string>
```

```
Usage of spokesman:
  -o string
        output audio file (support format of the audio: LINEAR16, MP3)
  -pitch float
        speaking pitch (-20.0 ~ 20.0)
  -rate float
        speech rate (0.25 ~ 4.0) (default 1)
  -text string
        text to speech
  -voice string
        speaker's voice name (default "stand-a")

```

##### speaker's voice name

| speaker's voice name                                                         | Description                                                                                                                                         |
|:-------------------------------------------------------------|:----------------------------------------------------------------------------------------------------------------------------------------------------|
| [stand-a](https://cloud.google.com/text-to-speech/docs/voices)   | Voice name: ja-JP-Standard-A, SSML Gender: FEMALE                                                                                                                                         |
| [stand-b](https://cloud.google.com/text-to-speech/docs/voices)   | Voice name: ja-JP-Standard-B, SSML Gender: FEMALE                                                                                                                                         |
| [stand-c](https://cloud.google.com/text-to-speech/docs/voices)   | Voice name: ja-JP-Standard-C, SSML Gender: MALE                                                                           |
| [stand-d](https://cloud.google.com/text-to-speech/docs/voices)   | Voice name: ja-JP-Standard-D, SSML Gender: MALE                                                                           |
| [wave-a](https://cloud.google.com/text-to-speech/docs/voices)   | Voice name: ja-JP-Wavenet-A, SSML Gender: FEMALE                                                                                                                           |
| [wave-b](https://cloud.google.com/text-to-speech/docs/voices)   | Voice name: ja-JP-Wavenet-B, SSML Gender: FEMALE                                                                             |
| [wave-c](https://cloud.google.com/text-to-speech/docs/voices)   | Voice name: ja-JP-Wavenet-C, SSML Gender: MALE                                                                             |
| [wave-d](https://cloud.google.com/text-to-speech/docs/voices)   | Voice name: ja-JP-Wavenet-D, SSML Gender: MALE                                                                              |
