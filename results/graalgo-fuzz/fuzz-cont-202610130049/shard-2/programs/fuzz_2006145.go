package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64; F1 uint8; F2 string }
func genFn0() int {
    var v6 uint32 = 2
    fmt.Println(v6)
    _ = v6
    return 0
}
func genFn1() uint64 {
    for fi7 := 0; fi7 < 8; fi7++ {
        for fi8 := 0; fi8 < 8; fi8++ {
            var v9 uint64 = 42
            v9 = v9
            v10 := []int{fi7, 186}
            _ = fi8
            _ = v10
            _ = v9
        }
        var v11 float32 = (1e10 * 1.5)
        var v12 float64 = -2.25
        _ = fi7
        _ = v11
        _ = v12
    }
    var v13 int16 = 263
    _ = v13
    return 2
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]uint32)
    var v3 string = "\n\t"
    var v4 bool = (true || (!(!true)))
    v5 := []int{127, 9223372036854775807}
    var v14 float64 = (3.14159 * -2.25)
    var v15 int8 = -111
    v3 = (v3 + v3)
    var v16 int64 = -1
    v17 := GenS1{F0: 32767}
    v15 -= -44
    v18 := GenS2{F0: 942}
    v19 := GenS2{F0: 18446744073709551615}
    for fi20 := 0; fi20 < 4; fi20++ {
        switch (v3 + v3) {
        case "\n\t":
            v14 = -2.25
        case "":
            v21 := GenS2{F0: genFn1()}
            _ = v21
        case "hello":
            var v22 uint16 = 1
            _ = v22
        }
        fmt.Println(v15, v14, fi20)
        _ = fi20
    }
    for fi23 := 0; fi23 < 17; fi23++ {
        v4 = ((0.1 * ((float32((v14 - (v14 + v14))) / 1.0) - (-2.25 * ((-2.25 - float32(v14)) / 1e10)))) >= 1.5)
        switch v3 {
        case "":
            var v24 int32 = 0
            _ = v24
        default:
            v25 := GenS3{F0: 2, F1: 42, F2: v3}
            _ = v25
        }
        _ = fi23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
