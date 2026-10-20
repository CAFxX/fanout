package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 bool; F1 string }
func genFn0(p0 uint8, p1 int16) int32 {
    p0 = (p0 + p0)
    v7 := []int{9223372036854775807, 32768}
    _ = p0
    _ = p1
    _ = v7
    return -2147483648
}
func genFn1(p0 int16) uint16 {
    fmt.Println(p0)
    var v8 int16 = p0
    switch 32767 {
    case 0:
        v8 = (v8 % ((-228) | 1))
    case 543:
        v9 := []int{387}
        _ = v9
    }
    p0 += 1000
    if (1 == 0) {
        v10 := GenS1{F0: (!((!(588 > 32767)) || false)), F1: 726}
        v11 := GenS1{F0: true, F1: 126}
        _ = v10
        _ = v11
    }
    _ = p0
    _ = v8
    return 37
}

func main() {
    var v1 float32 = 0.1
    var v2 float32 = v1
    var v3 int64 = 1
    var v4 uint16 = 42
    var v5 float32 = 1.5
    var v6 uint8 = 13
    v12 := []int{}
    switch 0 {
    case -462:
        if ((-2.25 / (float64(v5) + (((float64((((v5 + -2.25) / float32((((0.0 + 1.5) + 3.14159) / (0.0 - 1e10)))) / v2)) / 1.0) + (1.5 + 1.5)) / (1.5 + 3.14159)))) == float64((float32(1e10) * v2))) {
            fmt.Println(v2)
            for fi13 := 0; fi13 < 17; fi13++ {
                var v14 int = fi13
                _ = fi13
                _ = v14
            }
        }
    case -9223372036854775808:
        _ = len(v12)
    case 0:
        if false {
            v4 = v4
            if len(v12) > 0 {
                v12[0] = 128
            }
            v15 := GenS2{F0: 0, F1: 834}
            _ = v15
        } else {
            var v16 int8 = 72
            _ = v16
        }
    }
    v6 = v6
    v17 := GenS1{F0: ("hello" > "\n\t"), F1: 4294967295}
    for fi18 := 0; fi18 < 9; fi18++ {
        var v19 int64 = 1
        _ = fi18
        _ = v19
    }
    switch 32768 {
    case 32767:
        var v20 bool = (0 == -2147483648)
        _ = v20
    default:
        var v21 int32 = -2147483648
        _ = v21
    }
    var v22 float32 = ((1.5 + ((1.5 - v2) + v1)) - v1)
    var v23 uint8 = (v6 + 129)
    v23 = (v6 ^ (v23 * (v23 ^ v23)))
    var v24 int8 = -88
    v6 = (v23 | v23)
    if true {
        v5 = 1.5
    }
    switch "hello" {
    case "hello":
        if true {
            fmt.Println(v2, v3)
        }
    }
    v4 -= 1
    switch ("hello" + ("\n\t" + "")) {
    case "\n\t":
        fmt.Println(v23, v23, v24)
    case "a":
        v25 := GenS3{F0: false, F1: "a"}
        _ = v25
    default:
        for fi26 := 0; fi26 < 4; fi26++ {
            var v27 float64 = -2.25
            if (((1e10 <= 1e10) || false) && (false || (true && (65535 <= 598)))) {
                v27 += v27
            }
            _ = fi26
            _ = v27
        }
    }
    v6 -= (v6 >> ((1) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
