package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 int; F1 uint; F2 uint8 }
type GenS4 struct { F0 int16; F1 int8; F2 uint8 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 bool; F1 int32; F2 string }
func genFn0(p0 uint64, p1 string, p2 uint8) uint32 {
    v6 := []int{-149, int(p2), -1}
    switch p1 {
    case "hello":
        v7 := GenS1{F0: uint16(p2), F1: 2}
        _ = v7
    case "a":
        v6 = append(v6, -666)
    case "":
        if len(v6) > 0 {
            v6[0] = 32767
        }
    default:
        var v8 float64 = 1e10
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 1
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 string = ("a" + "hello")
    var v3 bool = (3 == 1)
    var v4 bool = (false || false)
    v5 := make(map[uint32]string)
    if (((v4 || v4) && (611 != 4294967295)) && (v4 || v3)) {
        v2 = "\n\t"
    } else {
        v2 = "foo bar"
        if (!false) {
            v3 = true
            v1 -= (v1 % ((v1) | 1))
        }
        v3 = false
    }
    switch -2147483648 {
    case 0:
        var v9 float64 = float64((1.5 / 1.0))
        _ = v9
    case 9223372036854775807:
        v2 = v2
    case -1:
        var v10 bool = (-1 > 9)
        _ = v10
    default:
        v1 += (v1 >> (((v1 | v1)) & 63))
    }
    v11 := GenS2{F0: "hello", F1: 2}
    switch -9223372036854775808 {
    case 42:
        switch (v2 + v2) {
        case "a":
            v1 -= 493
        case "\n\t":
            var v12 uint = (^(v1))
            _ = v12
        default:
            v1 -= (v1 & 597)
        }
    case -367:
        if v3 {
            v13 := GenS3{F0: 521, F1: 908, F2: 128}
            _ = v13
        }
    default:
        fmt.Println(v2)
    }
    if true {
        if v4 {
            v14 := GenS4{F0: 1000, F1: 39, F2: 42}
            var v15 uint64 = 1
            _ = v14
            _ = v15
        } else {
            v16 := GenS2{F0: v2, F1: 63}
            _ = v16
        }
    } else {
        for fi17 := 0; fi17 < 19; fi17++ {
            v18 := []int{}
            v19 := GenS4{F0: 1000, F1: -1, F2: 2}
            var v20 float32 = (((1e10 - 1.5) - 3.14159) - ((1e10 * (float32(0.1) + 0.1)) * (3.14159 / (float32(0.1) + -2.25))))
            _ = fi17
            _ = v18
            _ = v19
            _ = v20
        }
        var v21 int = -2147483648
        _ = v21
    }
    if (128 >= 0) {
        var v22 int8 = 127
        if v3 {
            var v23 bool = false
            v24 := []int{32767}
            if false {
                var v25 uint16 = 255
                if (1 <= 127) {
                    if v3 {
                        v25 = 0
                        v26 := GenS5{F0: (v4 || false)}
                        _ = v26
                    }
                    var v27 int8 = 75
                    _ = v27
                }
                _ = v25
            } else {
                v28 := GenS1{F0: 255, F1: 65535}
                fmt.Println(v23, v4, v23)
                _ = v28
            }
            _ = v23
            _ = v24
        }
        _ = v22
    }
    if (v3 || (!false)) {
        var v29 int64 = 1
        if (127 <= 42) {
            _ = len(v5)
            v29 += 158
            v4 = (uint32(v1) >= uint32(v1))
        }
        for fi30 := 0; fi30 < 15; fi30++ {
            v4 = v3
            fmt.Println(v1, v3, v4)
            if (!(820 == 758)) {
                if true {
                    v31 := GenS3{F0: (fi30 ^ (fi30 * (fi30 - (fi30 << (((^(v1))) & 63))))), F1: v1, F2: 128}
                    _ = v31
                }
            } else {
                var v32 uint8 = 128
                _ = len(v5)
                _ = v32
            }
            _ = fi30
        }
        _ = v29
    } else {
        if (v2 <= (((v2 + v2) + "foo bar") + v2)) {
            v1 -= (v1 / ((232) | 1))
            v4 = (!((0.0 - 0.0) > (-2.25 / 1.0)))
            v33 := GenS6{F0: (-43 <= 0), F1: 42, F2: (("foo bar" + (v2 + v2)) + v2)}
            _ = v33
        }
        for fi34 := 0; fi34 < 10; fi34++ {
            delete(v5, 1)
            _ = len(v5)
            _ = fi34
        }
        v35 := GenS5{F0: (!false)}
        _ = v35
    }
    if true {
        if (-478 == 1) {
            v1 -= v1
            var v36 int16 = 32767
            fmt.Println(v2, v36)
            _ = v36
        } else {
            var v37 uint32 = 42
            _ = len(v5)
            var v38 uint8 = 0
            _ = v37
            _ = v38
        }
    }
    v4 = ((4294967295 > genFn0(2, "\n\t", 239)) || v4)
    v39 := GenS1{F0: 409, F1: 2}
    if v4 {
        fmt.Println(v2, v4)
        for fi40 := 0; fi40 < 7; fi40++ {
            delete(v5, genFn0(511, v2, uint8(v1)))
            _ = fi40
        }
        v41 := []int{}
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
}
