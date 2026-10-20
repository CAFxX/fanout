package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 uint }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]bool)
    var v4 int16 = -1
    v5 := []uint32{485, 872, 426, 545}
    _ = len(v5)
    var v6 uint16 = 142
    if (175 < 128) {
        if true {
            v6 += v6
        }
    }
    v4 = -1
    if false {
        v4 = (-(v4))
        if false {
            v2["foo bar"] = 2147483648
        }
    } else {
        for fi7 := 0; fi7 < 19; fi7++ {
            delete(v3, 557)
            _ = fi7
        }
    }
    fmt.Println(v4, v6)
    for fi8 := 0; fi8 < 5; fi8++ {
        fmt.Println(v4)
        v9 := GenS1{F0: "foo bar", F1: 2147483647, F2: 255}
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 string = ""
        _ = fi10
        _ = v11
    }
    if true {
        var v12 int32 = 42
        fmt.Println(v12)
        v12 = v12
        _ = v12
    }
    if ((false || (127 == -1)) || false) {
        v6 -= (v6 << ((2) & 63))
        v6 -= 1
    } else {
        var v13 uint32 = 894
        _ = v13
    }
    var v14 uint64 = 2
    for fi15 := 0; fi15 < 18; fi15++ {
        fmt.Println(v6, v14, v4)
        _ = fi15
    }
    _ = len(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
