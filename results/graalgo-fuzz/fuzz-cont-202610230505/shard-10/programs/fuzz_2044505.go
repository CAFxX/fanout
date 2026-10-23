package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 uint8; F1 uint; F2 int32 }
type GenS3 struct { F0 uint16 }

func main() {
    v1 := []int{}
    var v2 int = -9223372036854775808
    v3 := make(map[string]int)
    fmt.Println(v2)
    var v4 float32 = float32(((-2.25 * float64(1e10)) - (1.5 * 0.1)))
    if true {
        var v5 float64 = (-2.25 * float64(v4))
        var v6 int16 = 42
        v7 := GenS1{F0: false, F1: ("\n\t" + "hello")}
        _ = v5
        _ = v6
        _ = v7
    } else {
        switch "\n\t" {
        case "hello":
            if (1 == 18446744073709551615) {
                if true {
                    for fi8 := 0; fi8 < 17; fi8++ {
                        v9 := GenS2{F0: 255, F1: 255, F2: -1}
                        _ = fi8
                        _ = v9
                    }
                    v10 := GenS2{F0: 128, F1: 1, F2: -1}
                    _ = v10
                }
                var v11 uint32 = 0
                _ = v11
            }
        default:
            _ = len(v1)
        }
        if ((false || ((397 == 42) || false)) || false) {
            if len(v1) > 0 {
                v1[0] = v2
            }
        } else {
            for fi12 := 0; fi12 < 19; fi12++ {
                var v13 float32 = v4
                _ = fi12
                _ = v13
            }
        }
    }
    v14 := GenS3{F0: 2}
    var v15 float32 = v4
    for fi16 := 0; fi16 < 4; fi16++ {
        switch fi16 {
        case 42:
            var v17 int = v2
            _ = v17
        default:
            var v18 int64 = -566
            _ = v18
        }
        if false {
            v15 -= 3.14159
            fmt.Println(fi16)
            var v19 float32 = (((v15 + float32(-2.25)) - v15) - v15)
            _ = v19
        } else {
            var v20 float64 = -2.25
            var v21 float32 = v4
            _ = v20
            _ = v21
        }
        _ = fi16
    }
    var v22 uint16 = 0
    v4 -= v15
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
