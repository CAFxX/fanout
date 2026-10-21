package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8; F1 uint; F2 uint8 }
type GenS5 struct { F0 int8; F1 uint64 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int8; F1 int16; F2 uint32 }
func genFn0() int16 {
    v7 := []int{-9223372036854775808, -1}
    v8 := GenS1{F0: 4294967295}
    var v9 uint32 = 4294967295
    var v10 bool = (float32((((float64(0.0) - 0.0) + 1.5) - (0.1 * 1.5))) < (float32(-2.25) + (1e10 + 0.0)))
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 uint = 0
    var v2 uint16 = 255
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]string)
    v5 := make(map[int]string)
    var v6 int16 = 438
    var v11 float64 = 0.0
    switch -74 {
    case 116:
        v12 := GenS2{F0: (^(v1))}
        _ = v12
    case 2147483647:
        v13 := GenS3{F0: 255}
        _ = v13
    case -9223372036854775808:
        v14 := GenS3{F0: 0}
        _ = v14
    default:
        if true {
            if (762 < 0) {
                fmt.Println(v1, v11, v11)
                v4[2147483647] = ("hello" + ("hello" + (("hello" + "hello") + "hello")))
                if true {
                    v11 += ((3.14159 + (v11 - (v11 / (v11 * 0.1)))) + 0.1)
                    var v15 uint16 = (v2 * (v2 >> ((v1) & 63)))
                    _ = v15
                }
            }
        }
    }
    switch "a" {
    case "":
        var v16 uint8 = 145
        _ = v16
    case "\n\t":
        var v17 uint32 = 307
        _ = v17
    case "foo bar":
        if false {
            if ((v6 / ((v6) | 1)) != 1) {
                v1 += v1
                var v18 int16 = v6
                _ = v18
            }
        }
    default:
        if (!true) {
            if (true || true) {
                v2 = (^(v2))
                v19 := GenS4{F0: -1, F1: (^(v1)), F2: 68}
                v20 := GenS4{F0: -1, F1: v1, F2: 2}
                _ = v19
                _ = v20
            }
        }
    }
    v21 := GenS4{F0: 60, F1: 2, F2: 128}
    var v22 int32 = -2147483648
    delete(v5, 668)
    if ((v1 <= v1) && (127 > 0)) {
        v23 := GenS5{F0: 0, F1: 42}
        var v24 uint32 = 659
        _ = v23
        _ = v24
    }
    v25 := GenS6{F0: 9223372036854775807}
    fmt.Println(v22, v6, v1)
    v6 = 42
    fmt.Println(v1, v1)
    v26 := GenS7{F0: 120, F1: genFn0(), F2: 2147483648}
    var v27 float32 = (1e10 / (1.5 * float32(float64(((((0.1 * 1e10) / 0.1) + -2.25) * (((-2.25 * -2.25) + (float32((v11 * 0.0)) / 1e10)) * float32(v11)))))))
    var v28 int64 = 586
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
