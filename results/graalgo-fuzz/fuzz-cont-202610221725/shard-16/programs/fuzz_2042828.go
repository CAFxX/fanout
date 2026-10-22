package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint8 }
type GenS2 struct { F0 uint8; F1 uint64; F2 uint8 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []uint32{748, 4294967295, 1, 2147483647}
    var v3 int8 = 127
    var v4 int = 0
    v5 := GenS1{F0: -2147483648, F1: "\n\t", F2: 78}
    for fi6 := 0; fi6 < 8; fi6++ {
        v7 := GenS2{F0: 255, F1: 457, F2: 128}
        var v8 uint = 65535
        fmt.Println(v3, v3, v3)
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 uint32 = 0
    for fi10 := 0; fi10 < 18; fi10++ {
        var v11 uint32 = 65535
        _ = fi10
        _ = v11
    }
    v3 = 42
    v12 := GenS2{F0: 127, F1: 2, F2: 42}
    fmt.Println(v3)
    for fi13 := 0; fi13 < 10; fi13++ {
        v3 += (v3 & (v3 & (v3 << ((369) & 63))))
        var v14 float32 = ((1.5 - -2.25) * (-2.25 / (0.1 - (0.0 + 1.5))))
        _ = fi13
        _ = v14
    }
    v9 += 4294967295
    var v15 int64 = 0
    if false {
        var v16 string = "foo bar"
        switch 225 {
        case 32767:
            for fi17 := 0; fi17 < 18; fi17++ {
                fmt.Println(v3, v16, fi17)
                var v18 string = "a"
                fmt.Println(v9)
                _ = fi17
                _ = v18
            }
        case -1:
            v3 = v3
        default:
            var v19 uint8 = 98
            _ = v19
        }
        v16 = v16
        _ = v16
    } else {
        v20 := GenS1{F0: 0, F1: "hello", F2: 128}
        _ = v20
    }
    if len(v2) > 0 {
        v2[0] = (v9 ^ (v9 << ((4294967295) & 63)))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
