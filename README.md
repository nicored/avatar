Avatar: Generate Avatars for the web
=====================================

[![Build Status](https://travis-ci.org/nxtvibe/avatar.svg)](https://travis-ci.org/nxtvibe/avatar) [![Go Report Card](https://goreportcard.com/badge/github.com/nxtvibe/avatar)](https://goreportcard.com/report/github.com/nxtvibe/avatar) [![Coverage Status](https://coveralls.io/repos/github/nxtvibe/avatar/badge.svg?branch=master)](https://coveralls.io/github/nxtvibe/avatar?branch=master) [![GoDoc](https://godoc.org/github.com/nxtvibe/avatar?status.svg)](https://godoc.org/github.com/nxtvibe/avatar)

Avatar is a package that allows you to create avatars for pictures and initials.

You can create square avatars, or round avatars.

# Picture:

```go
    size := 200

    fileBytes, _ := ioutil.ReadFile("./test_data/super_mascot.jpg")
    newAvatar, _ := NewAvatarFromPic(fileBytes, &PictureOptions{
        Size: size, // default 300
    })

    round, err := newAvatar.Circle()
    roundFile, _ := os.Create("./output/round_super_mascot.png")
    roundFile.Write(round)

    square, err := newAvatar.Square()
    squareFile, _ := os.Create("./output/square_super_mascot.png")
    roundFile.Write(square)
```

# Initials:

```go
    size := 200
    newAvatar, err := NewAvatarFromInitials([]byte("John Smith"), &InitialsOptions{
        FontPath:  "./test_data/Arial.ttf",    // Required
        Size:      size,                       // default 300
        NInitials: 2,                          // default 1 - If 0, the whole text will be printed
        TextColor: color.White,                // Default White
        BgColor:   color.RGBA{0, 0, 255, 255}, // Default color.RGBA{215, 0, 255, 255} (purple)
    })

    square, _ := newAvatar.Square()
    squareFile, _ := os.Create("./output/square_john_smith_initials.png")
    defer squareFile.Close()
    squareFile.Write(square)

    round, _ := newAvatar.Circle()
    roundFile, _ := os.Create("./output/round_john_smith_initials.png")
    defer roundFile.Close()
    roundFile.Write(round)
```

The input for initials can be any text, and the package identifies the best candidates for the initials.

If you give it an email address, it will try to identify the initials in the string preceding the @.

Not specifying NInitials in options would print the whole text as is.

And the font path is required, so maybe it should not be put in InitialsOptions, we'll see what the community thinks.



# What you get:

## From our original picture:

![Square: Super Mascot](./test_data/super_mascot.jpg "Square: Super Mascot")

## Our outputs
![Square: John Smith](./output/square_john_smith_initials.png "Square: John Smith") ![Round: John Smith](./output/round_john_smith_initials.png "Round: John Smith") ![Square: Mascot](./output/square_super_mascot.png "Square: Mascot") ![Round: Mascot](./output/round_super_mascot.png "Round: Mascot")
