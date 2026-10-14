package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 int16 }
type GenS2 struct { F0 uint64; F1 int32; F2 int }
type GenS3 struct { F0 bool }
func genFn0() string {
    v4 := []int{127}
    _ = len(v4)
    var v5 uint8 = 42
    if len(v4) > 0 {
        v4[0] = 917
    }
    _ = v4
    _ = v5
    return "foo bar"
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]int)
    var v3 uint = 48
    if (uint32(v3) > 781) {
        v6 := v2[32767]
        for fi7 := 0; fi7 < 7; fi7++ {
            if (!false) {
                fmt.Println(v3)
            } else {
                var v8 uint8 = 25
                if true {
                    v3 -= v3
                }
                var v9 uint8 = (v8 | (v8 >> (((v3 + 1)) & 63)))
                _ = v8
                _ = v9
            }
            v10 := GenS1{F0: 2, F1: false, F2: 1000}
            _ = fi7
            _ = v10
        }
        if false {
            v3 -= (v3 * 255)
            fmt.Println(v6)
        } else {
            for fi11 := 0; fi11 < 10; fi11++ {
                var v12 uint16 = 65535
                _ = fi11
                _ = v12
            }
            v6 -= (v6 - 32767)
        }
        _ = v6
    }
    fmt.Println(v3)
    v13 := GenS1{F0: 0, F1: (-159 == -9223372036854775808), F2: 32767}
    v3 -= (v3 & v3)
    switch "a" {
    case "hello":
        v3 -= 2
    case "\n\t":
        fmt.Println(v3)
    default:
        if false {
            v14 := []int{-279}
            _ = v14
        }
    }
    var v15 float32 = -2.25
    var v16 uint32 = 65535
    if false {
        v17 := GenS2{F0: 815, F1: 2147483647, F2: 4294967295}
        fmt.Println(v16, v16)
        _ = v17
    } else {
        fmt.Println(v15)
    }
    switch "" {
    case "foo bar":
        switch -2147483648 {
        case 502:
            v18 := GenS3{F0: true}
            _ = v18
        case -9223372036854775808:
            var v19 int8 = 100
            _ = v19
        default:
            var v20 int16 = 0
            _ = v20
        }
    case "a":
        v21 := v2[628]
        _ = v21
    default:
        if (179 == 0) {
            if (128 != -647) {
                fmt.Println(v16)
                fmt.Println(v15, v15)
                var v22 float64 = (1e10 / (1e10 - 3.14159))
                _ = v22
            }
            fmt.Println(v15, v16, v3)
            var v23 uint16 = 65535
            _ = v23
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
