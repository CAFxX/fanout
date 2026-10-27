package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32; F1 int8; F2 uint64 }
type GenS5 struct { F0 int; F1 uint64; F2 int64 }
type GenS6 struct { F0 bool; F1 int8 }
type GenS7 struct { F0 uint8; F1 int64 }
func genFn0() bool {
    var v5 int64 = -9223372036854775808
    v5 = (v5 << ((42) & 63))
    v6 := []int{-2147483648, -405}
    if true {
        switch 38 {
        case 947:
            switch 42 {
            case 32768:
                v5 += v5
            case 42:
                v5 -= -156
            case 32767:
                v7 := GenS1{F0: 1, F1: v5, F2: true}
                _ = v7
            default:
                switch "" {
                case "a":
                    var v8 int = 326
                    _ = v8
                case "hello":
                    v9 := GenS1{F0: 0, F1: (v5 ^ (v5 | v5)), F2: false}
                    _ = v9
                case "foo bar":
                    v10 := GenS2{F0: 654}
                    _ = v10
                default:
                    v11 := GenS1{F0: -1, F1: v5, F2: true}
                    _ = v11
                }
            }
        default:
            var v12 uint32 = 65535
            _ = v12
        }
        _ = len(v6)
    } else {
        fmt.Println(v5)
    }
    _ = v5
    _ = v6
    return (true || (1 != 0))
}
func genFn1() uint8 {
    v13 := []int{128}
    var v14 string = ("a" + "hello")
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    switch -659 {
    case 42:
        switch 9223372036854775807 {
        case 665:
            v14 = (v14 + "hello")
        default:
            var v15 uint16 = 754
            _ = v15
        }
    default:
        v14 = ""
    }
    var v16 int8 = -42
    _ = v13
    _ = v14
    _ = v16
    return 255
}

func main() {
    v1 := make(map[string]string)
    var v2 uint = 65535
    var v3 uint = 42
    var v4 uint16 = 255
    var v17 float32 = 1e10
    v2 = (v3 - (v3 % ((18446744073709551615) | 1)))
    for fi18 := 0; fi18 < 18; fi18++ {
        fmt.Println(fi18)
        v19 := GenS3{F0: "foo bar"}
        _ = fi18
        _ = v19
    }
    v20 := GenS4{F0: 2147483647, F1: 1, F2: 0}
    var v21 float32 = v17
    v22 := []int{557, 9223372036854775807, 127}
    v23 := GenS4{F0: 42, F1: 127, F2: 4294967295}
    if true {
        if len(v22) >= 2 {
            v22 = v22[1:len(v22)]
        }
    }
    var v24 string = ((("hello" + "a") + "\n\t") + ("foo bar" + "a"))
    switch v24 {
    case "a":
        var v25 uint32 = 65535
        _ = v25
    case "foo bar":
        if len(v22) > 0 {
            v22[0] = 32767
        }
    case "":
        if (-522 <= 955) {
            v26 := GenS4{F0: -2147483648, F1: -76, F2: 2}
            _ = v26
        } else {
            for fi27 := 0; fi27 < 12; fi27++ {
                switch "" {
                case "":
                    v28 := GenS5{F0: fi27, F1: 42, F2: 1}
                    _ = v28
                default:
                    v29 := GenS4{F0: 2147483647, F1: -1, F2: 42}
                    _ = v29
                }
                v30 := GenS6{F0: false, F1: -1}
                _ = fi27
                _ = v30
            }
            var v31 bool = (false && (!false))
            _ = v31
        }
    default:
        if len(v22) > 0 {
            v22[0] = -103
        }
    }
    if len(v22) >= 2 {
        v22 = v22[1:len(v22)]
    }
    v22 = append(v22, 32767)
    for fi32 := 0; fi32 < 20; fi32++ {
        var v33 int64 = 1
        v1[v24] = ("hello" + v24)
        _ = fi32
        _ = v33
    }
    v34 := GenS7{F0: 128, F1: 42}
    if ((v4 ^ (v4 >> ((v3) & 63))) <= v4) {
        var v35 uint16 = 266
        v3 = 162
        var v36 uint = (^(v2))
        _ = v35
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v34
    _ = v4
}
