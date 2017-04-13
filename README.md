Avatar

Avatar is a package that allows you to create avatars for pictures and initials.

You can create square avatars, or round avatars.

To use it with a picture:

```go
    size := 300
    av, _ := avatar.NewAvatarForPic([]byte("/path/to/your/pic.jpg", size)
    mySquaredAvatar, _ := av.Square()
    myRoundAvatar, _ := av.Circle()
    
    squareFile, _ := os.Create("/path/to/avatar_square.png")
    defer squareFile.Close()
    squareFile.Write(mySquareAvatar)

    roundFile, _ := os.Create("/path/to/avatar_circle.png")
    defer roundFile.Close()
    roundFile.Write(myRoundAvatar)

To use it with initials:

```go
    size := 300
    nCharacters := 2
    av, _ := avatar.NewAvatarForInitials([]byte("John Smith", size, nCharacters, "/path/to/your/font.ttf")
    mySquaredAvatar, _ := av.Square()
    myRoundAvatar, _ := av.Circle()
    
    squareFile, _ := os.Create("/path/to/avatar_square.png")
    defer squareFile.Close()
    squareFile.Write(mySquareAvatar)

    roundFile, _ := os.Create("/path/to/avatar_circle.png")
    defer roundFile.Close()
    roundFile.Write(myRoundAvatar)
