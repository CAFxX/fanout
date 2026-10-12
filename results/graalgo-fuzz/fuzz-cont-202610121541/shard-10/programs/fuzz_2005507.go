package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint8 }
type GenS2 struct { F0 uint16; F1 string }
func genFn0() int64 {
    switch "\n\t" {
    case "\n\t":
        var v4 uint32 = 809
        _ = v4
    case "":
        for fi5 := 0; fi5 < 6; fi5++ {
            var v6 bool = false
            var v7 uint = 18446744073709551615
            _ = fi5
            _ = v6
            _ = v7
        }
    default:
        var v8 uint8 = 42
        _ = v8
    }
    var v9 uint8 = 163
    _ = v9
    return 42
}
func genFn1() int32 {
    var v10 uint32 = 2
    v11 := GenS1{F0: 4294967295, F1: (v10 % (((v10 + (v10 - (v10 % (((v10 % ((1) | 1))) | 1))))) | 1)), F2: 253}
    fmt.Println(v10)
    if (!true) {
        v12 := GenS1{F0: 766, F1: (v10 / (((v10 * 918)) | 1)), F2: 2}
        _ = v12
    } else {
        var v13 uint8 = 2
        v13 = (v13 + v13)
        _ = v13
    }
    _ = v10
    _ = v11
    return int32(v10)
}

func main() {
    var v1 int16 = -1
    v2 := make(map[string]uint32)
    v3 := make(map[string]int)
    fmt.Println(v1)
    v1 = -671
    v1 += v1
    var v14 uint16 = 2
    v15 := GenS1{F0: 862, F1: 1, F2: 42}
    var v16 uint = uint(v14)
    v16 -= v16
    v17 := GenS2{F0: v14, F1: ""}
    v1 -= (v1 | 61)
    fmt.Println(v16, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
