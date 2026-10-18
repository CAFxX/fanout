package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 uint; F1 uint8 }

func main() {
    var v1 uint8 = 128
    v2 := []string{}
    v3 := make(map[string]bool)
    v4 := []uint32{0, 65535}
    var v5 int8 = 76
    v3["foo bar"] = true
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 bool = false
        var v8 uint = 647
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 int = int(v5)
    for fi10 := 0; fi10 < 3; fi10++ {
        fmt.Println(v1)
        v11 := GenS1{F0: 2, F1: v5}
        _ = fi10
        _ = v11
    }
    if false {
        for fi12 := 0; fi12 < 19; fi12++ {
            v1 = v1
            _ = fi12
        }
    }
    var v13 int32 = -2147483648
    fmt.Println(v13, v5)
    for fi14 := 0; fi14 < 13; fi14++ {
        var v15 int = (fi14 / ((v9) | 1))
        delete(v3, "")
        _ = fi14
        _ = v15
    }
    if false {
        fmt.Println(v1)
        _ = len(v2)
    } else {
        if (!((!true) && false)) {
            fmt.Println(v5)
            if true {
                fmt.Println(v9, v5, v9)
                var v16 bool = (!((v5 ^ v5) >= v5))
                _ = v16
            } else {
                v13 = (v13 << ((42) & 63))
                _ = len(v3)
            }
            var v17 int32 = 866
            _ = v17
        }
    }
    var v18 int64 = 502
    v19 := GenS2{F0: 255, F1: (v1 | (^(v1)))}
    if false {
        v20 := GenS2{F0: 42, F1: v1}
        var v21 uint = 42
        _ = v20
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
