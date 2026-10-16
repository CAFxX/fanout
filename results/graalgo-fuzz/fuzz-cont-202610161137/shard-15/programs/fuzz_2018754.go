package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 int = 1
    v2 := make(map[int]uint32)
    var v3 int32 = 1
    var v4 uint64 = 2
    var v5 uint16 = 2
    var v6 string = ((("a" + "foo bar") + "hello") + "")
    var v7 int8 = 1
    for fi8 := 0; fi8 < 3; fi8++ {
        if false {
            var v9 uint = 35
            _ = v9
        }
        fmt.Println(v4, v5)
        v10 := GenS1{F0: (v3 + (v3 + -1))}
        _ = fi8
        _ = v10
    }
    switch v6 {
    case "":
        var v11 uint = uint(v4)
        _ = v11
    }
    var v12 int16 = 0
    var v13 int32 = (v3 & 967)
    v14 := GenS1{F0: (v3 % ((979) | 1))}
    _ = len(v2)
    var v15 string = (v6 + "")
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
