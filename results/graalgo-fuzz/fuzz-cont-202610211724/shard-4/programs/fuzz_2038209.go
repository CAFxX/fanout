package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool; F1 int32 }
func genFn0(p0 int, p1 bool) uint32 {
    var v4 uint16 = 65535
    if (88 == 1) {
        var v5 uint8 = 1
        var v6 int32 = 618
        v4 = (v4 >> ((444) & 63))
        _ = v5
        _ = v6
    }
    var v7 int32 = 1
    _ = p0
    _ = p1
    _ = v4
    _ = v7
    return 722
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int16 = -1
    v3 := make(map[int]string)
    var v8 int64 = 1
    var v9 bool = (-32768 <= v2)
    fmt.Println(v2)
    fmt.Println(v8)
    if v9 {
        var v10 float64 = 0.1
        _ = v10
    } else {
        v11 := []int{9223372036854775807, -2147483648}
        v12 := GenS1{F0: 255, F1: (32767 >= 2), F2: 527}
        _ = v11
        _ = v12
    }
    var v13 uint64 = 42
    switch "a" {
    case "":
        v13 -= 0
    }
    var v14 uint = 42
    fmt.Println(v9)
    var v15 int8 = 42
    switch -2147483648 {
    case 127:
        var v16 uint16 = 42
        _ = v16
    case -1:
        fmt.Println(v15, v13, v14)
    case 1:
        var v17 int16 = v2
        _ = v17
    default:
        v9 = v9
    }
    v13 -= v13
    var v18 float32 = -2.25
    if v9 {
        for fi19 := 0; fi19 < 8; fi19++ {
            if ((!((42 == 0) && true)) || ((v15 + (v15 * v15)) <= v15)) {
                v2 -= v2
                v18 -= float32(1.5)
                for fi20 := 0; fi20 < 13; fi20++ {
                    v13 -= (v13 / (((^(v13))) | 1))
                    var v21 bool = ((!(2 <= 42)) && (661 != (v13 & uint64(v14))))
                    _ = fi20
                    _ = v21
                }
            } else {
                var v22 int8 = v15
                v22 = (v15 % (((^(v22))) | 1))
                _ = v22
            }
            _ = fi19
        }
        var v23 uint16 = 501
        v24 := GenS1{F0: v23, F1: true, F2: 0}
        _ = v23
        _ = v24
    }
    fmt.Println(v13, v13)
    switch "foo bar" {
    case "a":
        _ = len(v3)
    case "":
        var v25 float64 = -2.25
        _ = v25
    case "\n\t":
        switch -9223372036854775808 {
        case 127:
            v26 := GenS2{F0: v14}
            _ = v26
        case 9223372036854775807:
            if v9 {
                if (v9 || v9) {
                    v27 := v3[9223372036854775807]
                    v15 = v15
                    _ = v27
                } else {
                    v14 -= 237
                    switch "\n\t" {
                    case "hello":
                        v28 := GenS3{F0: v13}
                        _ = v28
                    default:
                        v29 := GenS4{F0: ((v18 / 0.1) != (v18 * v18)), F1: -2147483648}
                        _ = v29
                    }
                    var v30 float64 = 0.1
                    _ = v30
                }
            }
        default:
            var v31 int8 = 1
            _ = v31
        }
    }
    fmt.Println(v8, v9)
    v2 -= v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
