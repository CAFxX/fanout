package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 uint8 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 uint16; F1 int8; F2 uint8 }
func genFn0(p0 int32, p1 int) bool {
    var v5 int64 = 0
    p0 = 0
    var v6 bool = true
    var v7 float32 = (3.14159 - 3.14159)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return (!v6)
}

func main() {
    v1 := make(map[uint32]string)
    var v2 float32 = 0.0
    var v3 uint = 1
    var v4 float64 = -2.25
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 int64 = 0
        v9 = (-(v9))
        _ = fi8
        _ = v9
    }
    var v10 int64 = 1
    switch "a" {
    case "hello":
        var v11 uint64 = 373
        _ = v11
    case "a":
        v12 := GenS1{F0: v3, F1: 1, F2: 1}
        _ = v12
    case "":
        v13 := GenS2{F0: 42, F1: 2}
        _ = v13
    default:
        fmt.Println(v10, v10, v3)
    }
    var v14 int8 = 42
    v3 -= 65535
    v14 = v14
    var v15 uint = 4294967295
    if true {
        var v16 float64 = (float64(0.0) + 1e10)
        v17 := GenS2{F0: 2, F1: 245}
        v18 := GenS3{F0: 65535, F1: (v14 & (-(v14))), F2: 127}
        _ = v16
        _ = v17
        _ = v18
    }
    v19 := []int{0, -180, 32768, -2147483648}
    var v20 uint = (v15 / ((511) | 1))
    v14 = -43
    var v21 uint32 = 2147483647
    switch ("hello" + "\n\t") {
    case "a":
        if (!true) {
            v1[(v21 << ((v20) & 63))] = "hello"
        } else {
            var v22 string = ("a" + ("\n\t" + ("hello" + "")))
            v21 = (v21 << (((v3 >> ((42) & 63))) & 63))
            _ = v22
        }
    case "":
        for fi23 := 0; fi23 < 13; fi23++ {
            for fi24 := 0; fi24 < 15; fi24++ {
                v2 = float32(v4)
                if true {
                    var v25 float64 = v4
                    _ = v25
                }
                switch fi24 {
                case 1:
                    v26 := GenS1{F0: 255, F1: (v21 & (v21 + v21)), F2: 128}
                    _ = v26
                case 4294967295:
                    v14 += (v14 << (((v20 * (^(v15)))) & 63))
                }
                _ = fi24
            }
            v10 = 412
            var v27 uint64 = 2
            _ = fi23
            _ = v27
        }
    case "hello":
        v3 += v3
    }
    v14 -= 0
    var v28 float64 = (float64(0.0) + (v4 / v4))
    switch "" {
    case "foo bar":
        var v29 bool = (uint8(v15) < 42)
        _ = v29
    case "":
        v3 = (v20 / ((v20) | 1))
    default:
        v30 := v1[(^(v21))]
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v28
    _ = v3
    _ = v4
}
