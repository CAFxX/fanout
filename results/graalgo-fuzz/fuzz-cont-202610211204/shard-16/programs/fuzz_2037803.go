package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 uint; F1 uint16; F2 int }
type GenS3 struct { F0 uint32; F1 uint64; F2 uint32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int32; F1 uint }
type GenS6 struct { F0 int16 }

func main() {
    v1 := []int{948}
    var v2 uint32 = 2
    var v3 int = 32767
    if len(v1) > 0 {
        v1[0] = (v3 - 32768)
    }
    var v4 uint = 649
    v5 := GenS1{F0: 2147483647, F1: "hello"}
    if (!(0 != 1)) {
        fmt.Println(v2, v2, v2)
    }
    if false {
        if len(v1) > 0 {
            v1[0] = v3
        }
        for fi6 := 0; fi6 < 5; fi6++ {
            v7 := GenS1{F0: 1, F1: ("\n\t" + "a")}
            var v8 float64 = 3.14159
            _ = fi6
            _ = v7
            _ = v8
        }
    } else {
        var v9 int8 = -128
        if (false || false) {
            var v10 bool = false
            if true {
                v3 -= int(v2)
                if (v3 <= (v3 | (v3 + 4294967295))) {
                    _ = len(v1)
                    v11 := GenS1{F0: 0, F1: ""}
                    var v12 uint = (v4 & v4)
                    _ = v11
                    _ = v12
                } else {
                    v2 -= (v2 ^ (v2 >> ((v4) & 63)))
                    v13 := GenS2{F0: v4, F1: 0, F2: v3}
                    _ = v13
                }
            } else {
                fmt.Println(v4, v9)
                v1 = append(v1, v3)
                v14 := GenS3{F0: (v2 - v2), F1: 42, F2: (v2 / ((v2) | 1))}
                _ = v14
            }
            var v15 float32 = 1.5
            _ = v10
            _ = v15
        }
        if true {
            if ((!((float32((float64(1.5) / float64((float32(0.1) * (1.5 - 0.0))))) <= float32(1e10)) || (!true))) || (-2.25 != 1e10)) {
                for fi16 := 0; fi16 < 16; fi16++ {
                    var v17 int64 = -9223372036854775808
                    v18 := GenS3{F0: v2, F1: 672, F2: 4294967295}
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                v4 = (v4 + 1)
                v2 = (v2 >> ((65535) & 63))
            } else {
                v19 := GenS4{F0: false}
                _ = len(v1)
                switch v3 {
                case 42:
                    for fi20 := 0; fi20 < 7; fi20++ {
                        fmt.Println(fi20, v3)
                        _ = fi20
                    }
                case 128:
                    if len(v1) > 0 {
                        v1[0] = (v3 + -2147483648)
                    }
                default:
                    for fi21 := 0; fi21 < 17; fi21++ {
                        v22 := GenS3{F0: (^(v2)), F1: 4294967295, F2: v2}
                        _ = fi21
                        _ = v22
                    }
                }
                _ = v19
            }
            _ = len(v1)
            var v23 uint = v4
            _ = v23
        }
        _ = v9
    }
    switch "" {
    case "a":
        var v24 int64 = 42
        _ = v24
    }
    v25 := GenS5{F0: 1, F1: (v4 ^ (v4 >> ((uint(v2)) & 63)))}
    v2 -= 4294967295
    var v26 uint16 = 42
    v1 = append(v1, (v3 | 0))
    switch "hello" {
    case "a":
        switch "hello" {
        case "hello":
            fmt.Println(v4, v2)
        case "a":
            var v27 uint8 = 128
            _ = v27
        }
    case "hello":
        var v28 int = 115
        _ = v28
    default:
        var v29 float32 = ((3.14159 - -2.25) / 1e10)
        _ = v29
    }
    v30 := GenS2{F0: v4, F1: v26, F2: 2147483647}
    v4 = v4
    if false {
        v3 = v3
        fmt.Println(v26, v4, v26)
        v26 = (v26 * 2)
    }
    switch "" {
    case "a":
        for fi31 := 0; fi31 < 20; fi31++ {
            v1 = append(v1, v3)
            _ = fi31
        }
    case "foo bar":
        switch "a" {
        case "foo bar":
            v2 -= v2
        case "":
            v32 := GenS6{F0: 32767}
            _ = v32
        }
    case "":
        v3 = v3
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
