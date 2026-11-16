package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int16 }
type GenS2 struct { F0 int8; F1 uint32; F2 uint32 }
type GenS3 struct { F0 int16; F1 uint16; F2 bool }
type GenS4 struct { F0 int8 }
func genFn0() int64 {
    var v6 string = (("hello" + (("hello" + "\n\t") + (("a" + "a") + ""))) + "")
    v6 = (v6 + "a")
    v6 = (v6 + v6)
    switch v6 {
    case "foo bar":
        for fi7 := 0; fi7 < 18; fi7++ {
            var v8 float64 = (0.1 * 0.1)
            if (-112 == 48) {
                if false {
                    for fi9 := 0; fi9 < 16; fi9++ {
                        v6 = "\n\t"
                        v6 = (v6 + v6)
                        v10 := []int{(^(fi7)), 941}
                        _ = fi9
                        _ = v10
                    }
                    fmt.Println(fi7, v6)
                } else {
                    var v11 string = v6
                    _ = v11
                }
                var v12 int16 = 1
                var v13 float64 = ((v8 / (3.14159 / ((0.1 / 0.1) / (-2.25 / float64(3.14159))))) - v8)
                _ = v12
                _ = v13
            } else {
                var v14 uint = 1
                v15 := GenS1{F0: 32767, F1: -479, F2: -1}
                v8 += (v8 / 1.5)
                _ = v14
                _ = v15
            }
            _ = fi7
            _ = v8
        }
    default:
        switch ((v6 + "hello") + v6) {
        case "foo bar":
            v6 = ((v6 + "hello") + (v6 + v6))
        case "":
            var v16 bool = (-2.25 > (float64((float32((1.5 * 0.1)) - 0.0)) + 0.1))
            _ = v16
        default:
            switch 128 {
            case 9223372036854775807:
                v17 := GenS2{F0: 108, F1: 323, F2: 600}
                _ = v17
            default:
                switch 42 {
                case 32767:
                    v6 = "foo bar"
                case 9223372036854775807:
                    v18 := GenS2{F0: 0, F1: 720, F2: 4294967295}
                    _ = v18
                default:
                    v19 := GenS3{F0: 671, F1: 0, F2: false}
                    _ = v19
                }
            }
        }
    }
    _ = v6
    return 0
}

func main() {
    var v1 uint16 = 1
    var v2 int8 = 0
    var v3 uint16 = (v1 | (v1 & 42))
    v4 := make(map[uint32]uint32)
    v5 := make(map[int]int)
    v3 = v1
    if (!true) {
        for fi20 := 0; fi20 < 11; fi20++ {
            var v21 int = -9223372036854775808
            _ = fi20
            _ = v21
        }
        var v22 int8 = (v2 / ((v2) | 1))
        _ = v22
    }
    fmt.Println(v1, v1, v1)
    fmt.Println(v2, v1, v3)
    v23 := GenS4{F0: -128}
    v1 -= v3
    v24 := v4[914]
    var v25 uint = 1
    if false {
        if ((^(v24)) > v24) {
            v1 += v3
            v26 := GenS4{F0: v2}
            _ = v26
        }
        v27 := GenS1{F0: v3, F1: -1, F2: 1}
        for fi28 := 0; fi28 < 3; fi28++ {
            if (!(42 <= 0)) {
                var v29 float32 = ((float32(0.0) + float32(0.0)) - (0.0 / 0.1))
                _ = v29
            } else {
                v3 = v3
            }
            _ = fi28
        }
        _ = v27
    } else {
        v3 = 32767
        var v30 int32 = 150
        v5[-2147483648] = -449
        _ = v30
    }
    _ = len(v5)
    v3 = (v1 % ((255) | 1))
    var v31 float32 = (0.1 / 1.0)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
