package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 string }

func main() {
    v1 := []uint32{1, 662, 2147483647}
    v2 := make(map[string]string)
    var v3 int32 = -459
    var v4 int8 = 1
    var v5 int8 = 42
    for fi6 := 0; fi6 < 19; fi6++ {
        v2["foo bar"] = ("hello" + "")
        _ = len(v1)
        _ = fi6
    }
    fmt.Println(v3, v3)
    v3 -= v3
    switch -361 {
    case -466:
        v7 := v2[("\n\t" + "foo bar")]
        _ = v7
    }
    var v8 uint16 = 147
    var v9 uint = 467
    for fi10 := 0; fi10 < 17; fi10++ {
        var v11 string = "a"
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 2; fi12++ {
        switch 32768 {
        case 32768:
            var v13 uint32 = 2147483648
            _ = v13
        default:
            for fi14 := 0; fi14 < 10; fi14++ {
                v15 := GenS1{F0: v3, F1: true, F2: ("a" + "hello")}
                _ = fi14
                _ = v15
            }
        }
        var v16 uint16 = 2
        var v17 int64 = 0
        _ = fi12
        _ = v16
        _ = v17
    }
    fmt.Println(v5, v3, v5)
    for fi18 := 0; fi18 < 4; fi18++ {
        v4 = (v4 ^ 100)
        var v19 uint64 = 320
        fmt.Println(v4, v9, v8)
        _ = fi18
        _ = v19
    }
    var v20 uint64 = 949
    var v21 uint16 = (v8 << ((388) & 63))
    if len(v1) > 0 {
        v1[0] = 2
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
