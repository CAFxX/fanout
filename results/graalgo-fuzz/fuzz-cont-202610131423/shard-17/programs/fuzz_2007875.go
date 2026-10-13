package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 int16; F1 int64; F2 uint64 }
type GenS4 struct { F0 uint16; F1 uint16 }
type GenS5 struct { F0 int16; F1 int32 }
type GenS6 struct { F0 bool; F1 uint64; F2 int64 }
type GenS7 struct { F0 uint64; F1 int16; F2 uint32 }
func genFn0(p0 int32, p1 uint64) uint16 {
    var v7 uint16 = 0
    v8 := []int{int(p0), -9223372036854775808}
    var v9 uint = 987
    p0 = 0
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return (^(v7))
}
func genFn1(p0 int8, p1 uint32) uint32 {
    var v10 float32 = 1.5
    var v11 float64 = 3.14159
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return p1
}

func main() {
    v1 := []int{128, -1, 32767, 42}
    var v2 bool = false
    var v3 uint16 = 784
    v4 := make(map[int]bool)
    v5 := make(map[uint32]string)
    var v6 float32 = 0.1
    v12 := GenS1{F0: -1, F1: 264}
    if v2 {
        v2 = (-9223372036854775808 > 9223372036854775807)
    } else {
        if v2 {
            v13 := GenS2{F0: 2, F1: 2}
            v4[32768] = (3.14159 <= 0.0)
            _ = v13
        } else {
            var v14 bool = false
            _ = v14
        }
    }
    v15 := GenS3{F0: 32767, F1: -1, F2: 264}
    v16 := GenS1{F0: 42, F1: 458}
    v2 = (-267 >= -9223372036854775808)
    _ = len(v1)
    fmt.Println(v3)
    for fi17 := 0; fi17 < 6; fi17++ {
        if v2 {
            v18 := GenS4{F0: v3, F1: v3}
            _ = v18
        }
        var v19 float64 = (0.0 / ((3.14159 - (0.0 / float64((1e10 * (v6 - 3.14159))))) * ((0.0 / 1e10) + 0.1)))
        _ = fi17
        _ = v19
    }
    var v20 uint8 = 95
    var v21 uint16 = (v3 >> ((28) & 63))
    var v22 int = 127
    if (false && (18446744073709551615 >= 2)) {
        switch "foo bar" {
        case "a":
            for fi23 := 0; fi23 < 13; fi23++ {
                v24 := GenS4{F0: (v21 & (^(v21))), F1: (v3 - (^(v3)))}
                if (v20 == 200) {
                    v3 += 255
                } else {
                    var v25 uint8 = (v20 - v20)
                    switch (^(v22)) {
                    case 127:
                        v26 := GenS1{F0: v22, F1: 65535}
                        _ = v26
                    case 34:
                        fmt.Println(v2, v6, v25)
                    default:
                        var v27 uint8 = 128
                        _ = v27
                    }
                    _ = v25
                }
                if ((float64(((1e10 / (v6 - float32(1e10))) + v6)) / 1.0) >= 1.5) {
                    var v28 uint64 = 42
                    _ = v28
                }
                _ = fi23
                _ = v24
            }
        case "hello":
            v29 := GenS5{F0: 210, F1: -1}
            _ = v29
        case "":
            switch v22 {
            case 454:
                fmt.Println(v21, v3)
            case 539:
                var v30 int64 = 42
                _ = v30
            case 520:
                v31 := GenS2{F0: uint(v21), F1: 939}
                _ = v31
            }
        default:
            var v32 uint8 = (^(v20))
            _ = v32
        }
        if (65535 <= 2147483648) {
            fmt.Println(v22)
            var v33 int64 = -771
            _ = v33
        }
    } else {
        if false {
            v34 := GenS3{F0: 155, F1: 0, F2: 42}
            _ = v34
        }
        v35 := GenS6{F0: (!((("a" + "\n\t") != "foo bar") || false)), F1: 42, F2: 1}
        _ = v35
    }
    var v36 int = (v22 ^ v22)
    v37 := GenS7{F0: 576, F1: -32768, F2: 1}
    v38 := GenS6{F0: true, F1: 2, F2: 595}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
