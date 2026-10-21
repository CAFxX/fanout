package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 string; F2 int16 }
func genFn0(p0 bool, p1 uint16) uint8 {
    v6 := GenS1{F0: p1, F1: 1}
    var v7 uint = 2
    v8 := GenS1{F0: (p1 * p1), F1: 0}
    v7 -= v7
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 uint64 = 42
    v2 := make(map[string]bool)
    var v3 int16 = 546
    var v4 int32 = 0
    v5 := make(map[string]bool)
    fmt.Println(v3, v3, v4)
    var v9 int16 = (v3 % ((int16(v4)) | 1))
    if false {
        v10 := GenS2{F0: 140}
        _ = v10
    } else {
        v1 = (^(v1))
    }
    var v11 int16 = int16(v1)
    for fi12 := 0; fi12 < 8; fi12++ {
        v13 := GenS1{F0: 32767, F1: v11}
        v2["foo bar"] = true
        _ = fi12
        _ = v13
    }
    v14 := GenS3{F0: 0, F1: "hello", F2: (v9 + 42)}
    var v15 int8 = 78
    var v16 float32 = (((0.0 / 3.14159) - 0.0) - 3.14159)
    v3 -= (v3 + -1)
    if false {
        v17 := GenS1{F0: 65535, F1: (v9 - (v11 >> ((1) & 63)))}
        v18 := GenS2{F0: 0}
        if (-473 == 4294967295) {
            v4 -= v4
        } else {
            var v19 string = ("foo bar" + "a")
            v3 = -32768
            _ = v19
        }
        _ = v17
        _ = v18
    }
    switch 32768 {
    case -1:
        v11 = v11
    default:
        var v20 string = ("" + "")
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
