package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 int32 }
func genFn0(p0 uint16, p1 string, p2 int) int64 {
    var v5 int64 = 406
    var v6 uint16 = p0
    var v7 string = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return v5
}
func genFn1(p0 int64, p1 uint16, p2 int) uint8 {
    for fi8 := 0; fi8 < 1; fi8++ {
        var v9 uint8 = 20
        switch "hello" {
        case "hello":
            for fi10 := 0; fi10 < 5; fi10++ {
                v11 := []int{}
                fmt.Println(fi10, p0)
                p0 -= (p0 | p0)
                p0 += (p0 - genFn0(227, "", (fi10 - (fi10 * fi8))))
                fmt.Println(fi8, fi10)
                _ = fi10
                _ = v11
            }
        case "foo bar":
            if (true || false) {
                for fi12 := 0; fi12 < 14; fi12++ {
                    p2 -= (fi12 ^ (fi8 % ((p2) | 1)))
                    _ = fi12
                }
            } else {
                var v13 float32 = 3.14159
                _ = v13
            }
        case "\n\t":
            var v14 uint8 = uint8(p1)
            _ = v14
        default:
            v15 := []int{(p2 ^ p2), 32768}
            _ = v15
        }
        if true {
            v16 := GenS1{F0: ((0 < -112) || (((false && false) && false) || true)), F1: 2, F2: 446}
            _ = v16
        } else {
            p0 = p0
        }
        _ = fi8
        _ = v9
    }
    if (true || (1 == 42)) {
        var v17 bool = false
        var v18 int8 = -24
        for fi19 := 0; fi19 < 11; fi19++ {
            var v20 uint = 327
            _ = fi19
            _ = v20
        }
        _ = v17
        _ = v18
    } else {
        fmt.Println(p2, p1)
    }
    _ = p0
    _ = p1
    _ = p2
    return 127
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int32 = -2147483648
    var v3 uint8 = 127
    var v4 int16 = -993
    switch "hello" {
    case "hello":
        var v21 float64 = 3.14159
        _ = v21
    case "\n\t":
        if ((true || (((v3 / ((v3) | 1)) != genFn1(505, 32767, 9223372036854775807)) && false)) || (false && (127 < 9223372036854775807))) {
            if true {
                var v22 uint16 = 65535
                v2 = v2
                _ = v22
            }
            fmt.Println(v2, v2, v2)
        }
    case "":
        v23 := []int{-249, 32768}
        _ = v23
    }
    var v24 float32 = -2.25
    if ((v2 ^ v2) < (v2 / (((v2 & v2)) | 1))) {
        delete(v1, 546)
    } else {
        delete(v1, 1)
        var v25 uint64 = 688
        switch 127 {
        case -469:
            for fi26 := 0; fi26 < 10; fi26++ {
                _ = len(v1)
                v3 -= 46
                _ = fi26
            }
        case -2147483648:
            if (1e10 <= float32(float64(((v24 + 1.5) - (float32(float64(((v24 + (1e10 + (float32(0.0) / 1e10))) / (0.0 * v24)))) - (v24 * 0.0)))))) {
                if true {
                    fmt.Println(v25, v24, v25)
                }
                for fi27 := 0; fi27 < 8; fi27++ {
                    var v28 bool = (!(0 < 127))
                    v29 := []int{9223372036854775807, fi27, -2147483648, fi27}
                    _ = fi27
                    _ = v28
                    _ = v29
                }
            } else {
                v30 := GenS1{F0: ("hello" > "a"), F1: 880, F2: (v2 ^ (v2 ^ (v2 | (v2 - 0))))}
                _ = v30
            }
        case 32767:
            v31 := []int{268, 782, -1}
            _ = v31
        }
        _ = v25
    }
    v24 = v24
    var v32 uint64 = 2
    fmt.Println(v4, v32, v4)
    v32 = v32
    var v33 int = -2147483648
    fmt.Println(v24)
    fmt.Println(v32, v2)
    for fi34 := 0; fi34 < 12; fi34++ {
        for fi35 := 0; fi35 < 10; fi35++ {
            var v36 float64 = 1e10
            v32 = (v32 / ((1) | 1))
            _ = fi35
            _ = v36
        }
        _ = fi34
    }
    for fi37 := 0; fi37 < 7; fi37++ {
        var v38 int8 = 0
        v38 += (v38 / (((v38 ^ v38)) | 1))
        _ = fi37
        _ = v38
    }
    v2 = (v2 ^ v2)
    fmt.Println(v2, v33)
    fmt.Println(v24, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v32
    _ = v33
    _ = v4
}
