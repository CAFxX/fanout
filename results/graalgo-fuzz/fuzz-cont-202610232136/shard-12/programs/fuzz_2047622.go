package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64; F1 int16; F2 string }
type GenS4 struct { F0 int; F1 uint32; F2 uint16 }
type GenS5 struct { F0 uint8; F1 uint16 }
type GenS6 struct { F0 uint32; F1 int64; F2 int64 }
func genFn0(p0 int64, p1 bool, p2 int64) uint32 {
    p2 = p2
    fmt.Println(p0, p2)
    _ = p0
    _ = p1
    _ = p2
    return 65535
}
func genFn1(p0 string, p1 int8, p2 int16) int64 {
    v4 := GenS1{F0: (p0 + (p0 + p0))}
    v5 := []int{20}
    var v6 uint16 = 763
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return -1
}

func main() {
    var v1 int32 = -2147483648
    v2 := []int{-448, -1, -9223372036854775808}
    v3 := []string{"foo bar", "foo bar", "hello"}
    var v7 int16 = 949
    var v8 uint16 = 1
    v8 += v8
    for fi9 := 0; fi9 < 19; fi9++ {
        v10 := GenS2{F0: -128}
        var v11 uint = 65535
        _ = fi9
        _ = v10
        _ = v11
    }
    if ((407 != 363) && false) {
        for fi12 := 0; fi12 < 11; fi12++ {
            for fi13 := 0; fi13 < 20; fi13++ {
                var v14 int = (fi12 * 128)
                v15 := GenS3{F0: 464, F1: (^(v7)), F2: ""}
                fmt.Println(v8, v14, v1)
                var v16 uint64 = 1
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
            }
            _ = fi12
        }
        var v17 int32 = (-(v1))
        _ = v17
    } else {
        for fi18 := 0; fi18 < 10; fi18++ {
            if true {
                v19 := GenS4{F0: (fi18 * fi18), F1: 29, F2: v8}
                if (false || (1 <= 280)) {
                    v7 = (v7 % ((0) | 1))
                } else {
                    fmt.Println(v7)
                }
                v8 -= 42
                _ = v19
            } else {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            v20 := GenS5{F0: 42, F1: v8}
            _ = fi18
            _ = v20
        }
        for fi21 := 0; fi21 < 10; fi21++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi21
        }
        if (127 > 245) {
            var v22 float64 = 1e10
            fmt.Println(v7)
            _ = v22
        } else {
            fmt.Println(v7, v8)
            if false {
                fmt.Println(v7)
            }
            v7 = v7
        }
    }
    v23 := GenS4{F0: 32768, F1: 171, F2: v8}
    var v24 uint16 = (v8 % (((v8 >> ((0) & 63))) | 1))
    for fi25 := 0; fi25 < 4; fi25++ {
        if false {
            v1 += (v1 % (((v1 ^ (^(v1)))) | 1))
        }
        _ = fi25
    }
    v26 := GenS2{F0: 1}
    fmt.Println(v7)
    var v27 int = -989
    switch (v27 - (v27 ^ (-(v27)))) {
    case 917:
        var v28 string = "\n\t"
        _ = v28
    case 32768:
        v8 = 2
    case -276:
        var v29 float64 = 0.1
        _ = v29
    default:
        for fi30 := 0; fi30 < 16; fi30++ {
            v31 := GenS5{F0: 1, F1: 255}
            _ = fi30
            _ = v31
        }
    }
    v32 := GenS6{F0: 485, F1: -9223372036854775808, F2: genFn1("foo bar", 87, 932)}
    v24 += (v8 / ((v8) | 1))
    var v33 int64 = 917
    v3 = append(v3, "\n\t")
    for fi34 := 0; fi34 < 5; fi34++ {
        if (!true) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if ((172 < 198) || false) {
                var v35 string = "foo bar"
                if false {
                    fmt.Println(v8, v35, v8)
                }
                _ = v35
            }
        } else {
            switch "foo bar" {
            case "a":
                var v36 float32 = 0.0
                _ = v36
            }
            v8 = v8
            v24 = (^(v24))
        }
        _ = fi34
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v26
    _ = v27
    _ = v3
    _ = v32
    _ = v33
    _ = v7
    _ = v8
}
