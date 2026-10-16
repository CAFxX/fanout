package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint8; F1 string; F2 bool }
func genFn0(p0 int) uint64 {
    v7 := []int{}
    fmt.Println(p0)
    if (!false) {
        p0 += -9223372036854775808
    }
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 uint32 = 65535
        var v10 string = ""
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = v7
    return 42
}
func genFn1() int16 {
    var v11 bool = true
    v12 := []int{-1}
    v11 = (42 < 42)
    _ = v11
    _ = v12
    return 32767
}

func main() {
    var v1 bool = false
    var v2 int64 = -1
    var v3 bool = (0 < 1)
    v4 := make(map[int]int)
    var v5 int16 = -32768
    var v6 uint32 = 880
    var v13 float32 = (float32(((((((1.5 - 3.14159) * 1e10) + 0.1) * 1e10) / 3.14159) / -2.25)) / 1.0)
    v14 := GenS1{F0: 255, F1: 388, F2: 9223372036854775807}
    v15 := GenS2{F0: (!true), F1: v2}
    v13 = float32((1e10 / float64((0.1 - ((0.1 + ((float32(0.0) / (v13 + v13)) / (v13 + ((-2.25 - float32((-2.25 / (-2.25 / 1.0)))) + (1.5 / (1e10 + -2.25)))))) + float32(0.1))))))
    switch 0 {
    case -9223372036854775808:
        if (!v3) {
            delete(v4, -9223372036854775808)
            v16 := GenS3{F0: v3, F1: 1}
            _ = v16
        }
    case -642:
        for fi17 := 0; fi17 < 16; fi17++ {
            if true {
                var v18 bool = true
                v19 := []int{}
                _ = v18
                _ = v19
            } else {
                var v20 int32 = 973
                var v21 bool = v1
                var v22 int8 = 1
                _ = v20
                _ = v21
                _ = v22
            }
            v3 = ((v6 >> ((255) & 63)) > (v6 ^ (v6 & v6)))
            _ = fi17
        }
    default:
        v23 := GenS2{F0: (("foo bar" + "foo bar") < "a"), F1: (v2 >> ((0) & 63))}
        _ = v23
    }
    for fi24 := 0; fi24 < 4; fi24++ {
        fmt.Println(v1, v13, v2)
        _ = fi24
    }
    fmt.Println(v2, v13)
    var v25 int8 = -128
    v2 -= (^(v2))
    v4[661] = 32768
    v26 := GenS1{F0: 2, F1: 4294967295, F2: -482}
    var v27 int32 = 2147483647
    v4[213] = 778
    fmt.Println(v13, v5, v5)
    v27 = (v27 ^ (^(v27)))
    switch 9223372036854775807 {
    case 32767:
        if v1 {
            if (!(v3 || (!(int(v5) == 967)))) {
                fmt.Println(v25, v1)
                if (v13 != v13) {
                    v13 += v13
                }
            }
            switch 32767 {
            case 127:
                v28 := GenS2{F0: v3, F1: -838}
                _ = v28
            case 2147483647:
                for fi29 := 0; fi29 < 6; fi29++ {
                    v30 := GenS2{F0: false, F1: v2}
                    v31 := GenS4{F0: v27}
                    fmt.Println(v3, v2, v2)
                    _ = fi29
                    _ = v30
                    _ = v31
                }
            case 32767:
                v4[138] = 0
            default:
                v32 := GenS3{F0: v1, F1: (v25 >> ((0) & 63))}
                _ = v32
            }
        }
    default:
        v33 := GenS5{F0: 42, F1: ("a" + (("hello" + "foo bar") + "a")), F2: v3}
        _ = v33
    }
    if ((v27 * (v27 * (v27 - (v27 | (v27 + v27))))) == v27) {
        for fi34 := 0; fi34 < 15; fi34++ {
            fmt.Println(v1)
            _ = fi34
        }
        for fi35 := 0; fi35 < 16; fi35++ {
            for fi36 := 0; fi36 < 4; fi36++ {
                v25 = -69
                switch -288 {
                case 128:
                    fmt.Println(v2, v5)
                    fmt.Println(v5, fi35, v1)
                default:
                    var v37 uint32 = v6
                    _ = v37
                }
                _ = fi36
            }
            _ = fi35
        }
        v4[1] = -441
    } else {
        v38 := []int{537, 32767}
        fmt.Println(v1)
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
