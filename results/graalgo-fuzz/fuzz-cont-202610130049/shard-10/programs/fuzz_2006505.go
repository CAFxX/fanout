package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int64 }
type GenS2 struct { F0 int16; F1 uint16; F2 uint }
type GenS3 struct { F0 uint64; F1 uint }
type GenS4 struct { F0 uint8; F1 uint16 }
type GenS5 struct { F0 int; F1 int32; F2 int16 }

func main() {
    var v1 int16 = 664
    v2 := make(map[uint32]bool)
    var v3 uint = 42
    v4 := GenS1{F0: 32767, F1: v1, F2: 533}
    for fi5 := 0; fi5 < 18; fi5++ {
        v6 := GenS1{F0: 602, F1: (v1 & -677), F2: 9223372036854775807}
        fmt.Println(v3, v3)
        v1 = (v1 * 563)
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: 255, F1: (v1 << ((42) & 63)), F2: 481}
    v8 := []int{-9223372036854775808, -940, 0}
    for fi9 := 0; fi9 < 1; fi9++ {
        var v10 uint16 = 57
        _ = fi9
        _ = v10
    }
    var v11 string = ("hello" + "")
    v1 = v1
    v12 := GenS1{F0: uint16(v3), F1: 1000, F2: 866}
    v3 -= 4294967295
    v13 := GenS1{F0: 255, F1: v1, F2: -1}
    if false {
        fmt.Println(v1, v1, v11)
        for fi14 := 0; fi14 < 16; fi14++ {
            var v15 int8 = 100
            var v16 int8 = 100
            var v17 int32 = 42
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
    } else {
        var v18 int32 = 42
        _ = v18
    }
    switch (v11 + "foo bar") {
    case "foo bar":
        v1 += v1
    case "":
        if ((v3 & (^(v3))) >= v3) {
            _ = len(v2)
            for fi19 := 0; fi19 < 4; fi19++ {
                switch (fi19 + fi19) {
                case 2147483647:
                    v11 = v11
                default:
                    var v20 uint = 33
                    _ = v20
                }
                _ = fi19
            }
            if (836 >= -612) {
                for fi21 := 0; fi21 < 7; fi21++ {
                    fmt.Println(v11)
                    _ = fi21
                }
            }
        } else {
            if (false && (1.5 > 3.14159)) {
                var v22 uint8 = 1
                _ = v22
            }
            if (false && (1 == 1)) {
                var v23 uint8 = 128
                switch v11 {
                case "foo bar":
                    v24 := GenS2{F0: (v1 * v1), F1: 65535, F2: (v3 << (((v3 + v3)) & 63))}
                    _ = v24
                default:
                    _ = len(v8)
                }
                _ = v23
            } else {
                v25 := GenS3{F0: 1, F1: (^(v3))}
                v26 := GenS3{F0: 42, F1: (v3 % ((v3) | 1))}
                _ = v25
                _ = v26
            }
            for fi27 := 0; fi27 < 6; fi27++ {
                v11 = v11
                if true {
                    v8 = append(v8, (fi27 / ((fi27) | 1)))
                } else {
                    v28 := GenS4{F0: 2, F1: 461}
                    v29 := GenS5{F0: 128, F1: -1, F2: (^(v1))}
                    var v30 uint8 = 43
                    _ = v28
                    _ = v29
                    _ = v30
                }
                if (!(!true)) {
                    fmt.Println(v3, v1)
                } else {
                    v3 -= (v3 ^ (v3 >> (((v3 >> ((v3) & 63))) & 63)))
                    v31 := GenS1{F0: 32767, F1: v1, F2: 9223372036854775807}
                    fmt.Println(v1, v3, v3)
                    _ = v31
                }
                _ = fi27
            }
        }
    case "hello":
        if false {
            switch -1 {
            case 127:
                var v32 uint64 = 181
                _ = v32
            default:
                fmt.Println(v11, v3)
            }
        }
    default:
        if len(v8) > 0 {
            v8[0] = -1
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
