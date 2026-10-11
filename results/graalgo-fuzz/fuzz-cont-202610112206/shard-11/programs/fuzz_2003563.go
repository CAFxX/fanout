package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 uint32; F2 uint64 }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 bool; F1 uint64; F2 int64 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 1
    v3 := []int{-9223372036854775808, -781, 42, -326}
    v4 := []string{"\n\t", "hello", ("\n\t" + "\n\t")}
    var v5 uint32 = 2147483647
    _ = len(v1)
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 uint64 = 4294967295
        fmt.Println(v7, v2, v2)
        for fi8 := 0; fi8 < 8; fi8++ {
            if false {
                fmt.Println(v5, fi8, v2)
                var v9 int8 = 0
                var v10 string = ("\n\t" + ("" + ""))
                _ = v10
                _ = v9
            }
            v11 := GenS1{F0: (^(v2))}
            _ = fi8
            _ = v11
        }
        _ = fi6
        _ = v7
    }
    fmt.Println(v2)
    fmt.Println(v2)
    v5 -= (v5 & (v5 / ((2147483648) | 1)))
    if true {
        v12 := GenS1{F0: 42}
        v13 := GenS2{F0: 42, F1: v5, F2: 18446744073709551615}
        _ = v12
        _ = v13
    }
    if ((!(42 >= 1)) && ((0.0 / 1.0) == float32(float64((1.5 + -2.25))))) {
        switch ("foo bar" + (("" + "a") + "\n\t")) {
        case "":
            v2 += (v2 >> ((65535) & 63))
        case "a":
            switch -273 {
            case -9223372036854775808:
                v14 := GenS2{F0: 127, F1: (^(v5)), F2: 4294967295}
                _ = v14
            default:
                switch int(v5) {
                case -753:
                    var v15 uint8 = (v2 >> ((308) & 63))
                    _ = v15
                case 127:
                    if (!false) {
                        var v16 float32 = 1.5
                        var v17 uint32 = 95
                        v5 += v5
                        _ = v16
                        _ = v17
                    }
                case 1:
                    v18 := GenS3{F0: 359, F1: -319}
                    _ = v18
                default:
                    var v19 uint = 42
                    _ = v19
                }
            }
        default:
            v2 -= 235
        }
        switch -9223372036854775808 {
        case 901:
            fmt.Println(v5)
        case -416:
            v5 = v5
        default:
            for fi20 := 0; fi20 < 4; fi20++ {
                v2 -= (v2 & 128)
                if len(v3) > 0 {
                    v3[0] = (-(fi20))
                }
                for fi21 := 0; fi21 < 15; fi21++ {
                    var v22 float64 = 1.5
                    var v23 int64 = -900
                    var v24 int32 = -2147483648
                    _ = fi21
                    _ = v22
                    _ = v23
                    _ = v24
                }
                _ = fi20
            }
        }
        var v25 uint64 = 633
        _ = v25
    } else {
        fmt.Println(v5)
        if (v2 >= (v2 & 170)) {
            if (!true) {
                v26 := GenS2{F0: 32768, F1: 1, F2: 2}
                v27 := GenS1{F0: v2}
                _ = v26
                _ = v27
            } else {
                v28 := GenS3{F0: 97, F1: 42}
                _ = v28
            }
        } else {
            var v29 float64 = 3.14159
            v29 = -2.25
            _ = v29
        }
    }
    fmt.Println(v2)
    if ((("foo bar" + "") < "hello") || true) {
        v30 := GenS4{F0: false, F1: 42, F2: -1}
        _ = v30
    }
    switch -2147483648 {
    case -242:
        if false {
            if false {
                var v31 string = "\n\t"
                _ = v31
            }
            var v32 int16 = -1
            _ = v32
        }
    case -1:
        var v33 int8 = 1
        _ = v33
    case 190:
        for fi34 := 0; fi34 < 10; fi34++ {
            fmt.Println(v5, v5, v5)
            _ = fi34
        }
    }
    v5 = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
