package main

import (
    "fmt"
)


func main() {
    var v1 int16 = -32768
    v2 := make(map[uint32]bool)
    var v3 int8 = 42
    v4 := []uint32{54, 42}
    v2[4294967295] = (542 > -1)
    if (!(v1 < -126)) {
        var v5 string = ("foo bar" + ("\n\t" + ""))
        _ = v5
    }
    switch 42 {
    case 128:
        for fi6 := 0; fi6 < 1; fi6++ {
            var v7 int16 = v1
            _ = fi6
            _ = v7
        }
    case -1:
        var v8 uint64 = 222
        _ = v8
    default:
        _ = len(v2)
    }
    fmt.Println(v1, v3)
    var v9 int = 32768
    fmt.Println(v3)
    fmt.Println(v3, v3)
    v9 = v9
    var v10 int = v9
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
