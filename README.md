Avatar

Avatar is a package that allows you to create avatars for pictures and initials.

You can create square avatars, or round avatars.

To use it with a picture:

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

To use it with initials:

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
