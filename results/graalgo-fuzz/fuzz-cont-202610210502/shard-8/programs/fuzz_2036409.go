package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 int64 }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint; F1 uint8 }

func main() {
    var v1 int32 = 42
    v2 := make(map[uint32]uint32)
    var v3 string = "hello"
    v4 := []int{-1}
    v5 := GenS1{F0: 1, F1: 65535, F2: 9223372036854775807}
    v6 := GenS2{F0: 65535, F1: 255}
    for fi7 := 0; fi7 < 16; fi7++ {
        v8 := GenS3{F0: 815}
        _ = fi7
        _ = v8
    }
    v3 = (("hello" + (((v3 + (v3 + (v3 + (v3 + ("\n\t" + v3))))) + v3) + ((v3 + "a") + v3))) + v3)
    var v9 int8 = -1
    if len(v4) > 0 {
        v4[0] = -538
    }
    if len(v4) > 0 {
        v4[0] = -1
    }
    _ = len(v4)
    if (-9223372036854775808 < 42) {
        fmt.Println(v9, v9)
    } else {
        var v10 uint32 = 42
        _ = v10
    }
    var v11 uint32 = 2
    var v12 int64 = 42
    v13 := GenS1{F0: -1, F1: 32767, F2: v12}
    if true {
        v12 += v12
        v3 = (v3 + (v3 + "hello"))
        fmt.Println(v11, v1, v11)
    }
    var v14 float32 = -2.25
    var v15 int = -2147483648
    switch 4294967295 {
    case 491:
        v16 := v2[v11]
        _ = v16
    case -447:
        delete(v2, v11)
    case -590:
        v17 := GenS4{F0: 65535, F1: 128}
        _ = v17
    default:
        for fi18 := 0; fi18 < 14; fi18++ {
            v19 := GenS1{F0: 32767, F1: 1, F2: (v12 | 102)}
            _ = fi18
            _ = v19
        }
    }
    var v20 uint16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
