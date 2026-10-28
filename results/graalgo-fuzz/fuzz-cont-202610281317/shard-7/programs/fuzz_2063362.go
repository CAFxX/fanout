package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64; F1 int; F2 uint16 }
type GenS4 struct { F0 uint32; F1 int8; F2 int16 }
func genFn0() uint64 {
    var v6 int = 127
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 uint8 = 114
        _ = fi7
        _ = v8
    }
    var v9 uint = 0
    v10 := []int{(v6 | v6), (v6 - 0)}
    var v11 uint16 = 0
    _ = v10
    _ = v11
    _ = v6
    _ = v9
    return 18446744073709551615
}
func genFn1(p0 int32) int32 {
    var v12 uint = 313
    switch 4294967295 {
    case 32768:
        var v13 float32 = 1e10
        _ = v13
    case 128:
        var v14 uint8 = 127
        _ = v14
    default:
        var v15 uint = (v12 / ((v12) | 1))
        _ = v15
    }
    _ = p0
    _ = v12
    return 360
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{0}
    var v3 string = "hello"
    v4 := make(map[uint32]string)
    var v5 uint8 = 1
    v16 := GenS1{F0: v3}
    for fi17 := 0; fi17 < 7; fi17++ {
        _ = len(v1)
        _ = fi17
    }
    v18 := GenS1{F0: "\n\t"}
    switch v3 {
    case "foo bar":
        var v19 bool = (-242 > 42)
        _ = v19
    case "\n\t":
        v3 = v3
    default:
        v20 := GenS2{F0: v5}
        _ = v20
    }
    for fi21 := 0; fi21 < 7; fi21++ {
        var v22 uint16 = 185
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 7; fi23++ {
        for fi24 := 0; fi24 < 1; fi24++ {
            switch -2147483648 {
            case 127:
                var v25 int32 = 1
                _ = v25
            case -693:
                v26 := GenS3{F0: 1, F1: (fi23 * fi23), F2: 65535}
                _ = v26
            case 1:
                if (42 >= 0) {
                    v27 := GenS4{F0: 0, F1: -128, F2: 32767}
                    v28 := GenS4{F0: 42, F1: 127, F2: 1000}
                    _ = v27
                    _ = v28
                }
            }
            _ = fi24
        }
        _ = fi23
    }
    v5 = 2
    fmt.Println(v5, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
