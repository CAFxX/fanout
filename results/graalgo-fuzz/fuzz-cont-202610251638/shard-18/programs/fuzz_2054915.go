package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 uint64; F1 bool; F2 uint16 }
type GenS3 struct { F0 int64; F1 int; F2 int64 }
type GenS4 struct { F0 uint; F1 uint }

func main() {
    v1 := make(map[string]bool)
    var v2 uint16 = 65535
    v3 := make(map[string]uint32)
    v4 := make(map[int]int)
    var v5 int16 = -576
    switch 42 {
    case 127:
        v2 -= v2
    case -766:
        var v6 int32 = -652
        _ = v6
    case 9223372036854775807:
        if (-1 == 32767) {
            for fi7 := 0; fi7 < 7; fi7++ {
                var v8 uint8 = 128
                if (!true) {
                    v5 = (v5 + v5)
                }
                switch "foo bar" {
                case "hello":
                    v9 := GenS1{F0: v8, F1: 455}
                    _ = v9
                case "foo bar":
                    for fi10 := 0; fi10 < 18; fi10++ {
                        var v11 uint8 = 2
                        var v12 uint16 = v2
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                case "\n\t":
                    v8 += 255
                default:
                    switch 42 {
                    case 2147483647:
                        if (-9223372036854775808 <= int64(v8)) {
                            var v13 uint32 = 241
                            _ = v13
                        }
                    case 9223372036854775807:
                        var v14 float32 = ((0.1 * 0.1) / ((1e10 * 3.14159) - float32(0.0)))
                        _ = v14
                    }
                }
                _ = fi7
                _ = v8
            }
            v5 -= (v5 >> ((20) & 63))
        }
    default:
        fmt.Println(v2)
    }
    v15 := GenS1{F0: 58, F1: 65535}
    v16 := GenS1{F0: 42, F1: 807}
    if (v2 > (v2 + 65535)) {
        for fi17 := 0; fi17 < 5; fi17++ {
            if false {
                v2 = v2
            } else {
                fmt.Println(fi17)
            }
            v18 := GenS1{F0: 42, F1: 397}
            _ = fi17
            _ = v18
        }
        v19 := GenS1{F0: 1, F1: 65535}
        v20 := GenS2{F0: 1, F1: (48 == 1), F2: v2}
        _ = v19
        _ = v20
    }
    fmt.Println(v2, v5)
    fmt.Println(v2, v5)
    if false {
        v21 := GenS2{F0: 1, F1: ("hello" == ""), F2: v2}
        _ = v21
    }
    v5 = 1
    switch -880 {
    case 142:
        var v22 uint16 = (v2 + (^(v2)))
        _ = v22
    case 128:
        var v23 uint8 = 127
        _ = v23
    default:
        v5 = (^(v5))
    }
    if true {
        v24 := []int{-9223372036854775808, 42, 9223372036854775807, 2147483647}
        v5 -= -756
        v25 := v1[""]
        _ = v24
        _ = v25
    } else {
        var v26 int8 = 42
        _ = v26
    }
    var v27 int32 = -1
    if true {
        v2 += v2
        switch -9223372036854775808 {
        case 1:
            if true {
                v5 = v5
                switch ("foo bar" + "") {
                case "foo bar":
                    v28 := GenS3{F0: 9223372036854775807, F1: -1, F2: -236}
                    _ = v28
                }
                switch "\n\t" {
                case "\n\t":
                    var v29 int64 = 1
                    _ = v29
                }
            } else {
                v30 := GenS4{F0: 4294967295, F1: 65535}
                _ = v30
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
