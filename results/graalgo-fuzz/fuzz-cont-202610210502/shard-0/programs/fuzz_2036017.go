package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 int16; F2 int8 }
type GenS4 struct { F0 uint16 }
func genFn0(p0 uint8) uint8 {
    var v7 uint8 = 1
    var v8 int8 = 127
    var v9 int16 = 42
    v7 += p0
    if (!(18446744073709551615 == 2)) {
        switch 42 {
        case -155:
            v10 := []int{}
            _ = v10
        case 127:
            var v11 int16 = v9
            _ = v11
        case 573:
            if (2 == 602) {
                v7 -= (p0 / ((p0) | 1))
                v7 -= (v7 % ((0) | 1))
            }
        }
    } else {
        v8 += (v8 / ((v8) | 1))
        v12 := GenS1{F0: (v9 + (v9 ^ (v9 << ((uint(v7)) & 63))))}
        _ = v12
    }
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return (v7 >> ((4294967295) & 63))
}

func main() {
    var v1 float32 = 1.5
    v2 := []uint32{680, 42}
    var v3 float32 = (v1 * 1e10)
    v4 := []int{-731}
    var v5 int8 = 24
    var v6 int16 = -1
    var v13 uint = 42
    var v14 uint8 = 80
    v15 := GenS1{F0: 894}
    if false {
        v4 = append(v4, 0)
    }
    if true {
        var v16 int64 = 9223372036854775807
        _ = v16
    }
    switch 32767 {
    case 0:
        v13 -= 0
    case 32768:
        fmt.Println(v14)
    case 45:
        var v17 uint = (v13 & (v13 | (v13 | 1)))
        _ = v17
    }
    fmt.Println(v13)
    switch "" {
    case "a":
        var v18 string = ("\n\t" + "hello")
        _ = v18
    case "foo bar":
        v19 := GenS2{F0: v13}
        _ = v19
    case "\n\t":
        v20 := GenS3{F0: -337, F1: v6, F2: v5}
        _ = v20
    default:
        v14 += v14
    }
    fmt.Println(v6)
    var v21 uint32 = 0
    if (32767 == 947) {
        var v22 float64 = (3.14159 / (((3.14159 - (0.0 / float64(v3))) + ((float64(0.1) - (1.5 - 1e10)) - 3.14159)) + 1.5))
        v1 = v3
        v23 := GenS1{F0: (v6 * (v6 - (v6 ^ v6)))}
        _ = v22
        _ = v23
    } else {
        if false {
            var v24 float64 = -2.25
            _ = v24
        } else {
            v25 := GenS2{F0: 0}
            _ = v25
        }
        var v26 int64 = 42
        v5 = 42
        _ = v26
    }
    if true {
        var v27 uint = 1
        _ = v27
    } else {
        var v28 uint32 = v21
        for fi29 := 0; fi29 < 2; fi29++ {
            if (731 >= 0) {
                _ = len(v2)
                v28 = (v21 << ((uint(v14)) & 63))
            }
            var v30 uint32 = 0
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi29
            _ = v30
        }
        _ = v28
    }
    fmt.Println(v3, v3, v21)
    v31 := GenS4{F0: 42}
    switch "foo bar" {
    case "foo bar":
        switch 128 {
        case 128:
            var v32 int64 = -1
            _ = v32
        case -2147483648:
            v33 := GenS2{F0: v13}
            _ = v33
        case 574:
            var v34 string = ""
            _ = v34
        default:
            v14 = genFn0(v14)
        }
    case "":
        var v35 int = -1
        _ = v35
    case "hello":
        if true {
            v5 += v5
        }
    default:
        v36 := GenS1{F0: v6}
        _ = v36
    }
    if true {
        v37 := GenS3{F0: 128, F1: (-(v6)), F2: v5}
        var v38 uint = 195
        _ = v37
        _ = v38
    }
    v39 := GenS3{F0: 549, F1: int16(v13), F2: (v5 >> ((v13) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v31
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
