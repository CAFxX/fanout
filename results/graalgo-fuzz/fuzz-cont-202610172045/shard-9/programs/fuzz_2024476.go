package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 int8; F2 uint64 }
type GenS4 struct { F0 uint64; F1 uint32; F2 int8 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 bool }

func main() {
    var v1 int16 = 982
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]string)
    v4 := make(map[string]bool)
    var v5 int64 = 9223372036854775807
    var v6 int64 = -28
    v1 += -956
    v7 := GenS1{F0: "hello"}
    var v8 uint16 = 1
    if true {
        v6 -= 9223372036854775807
    }
    for fi9 := 0; fi9 < 12; fi9++ {
        var v10 int8 = 100
        _ = fi9
        _ = v10
    }
    v11 := GenS2{F0: "foo bar"}
    if ("\n\t" != "") {
        fmt.Println(v8, v6)
        if true {
            switch 4294967295 {
            case 995:
                var v12 uint16 = v8
                _ = v12
            case 128:
                fmt.Println(v5, v6, v1)
            default:
                var v13 uint = 550
                _ = v13
            }
            if (42 < 2147483647) {
                for fi14 := 0; fi14 < 5; fi14++ {
                    var v15 int = 550
                    var v16 int64 = (v6 % ((v5) | 1))
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                var v17 uint = 631
                _ = v17
            }
        } else {
            var v18 int8 = 79
            v19 := []int{}
            _ = v18
            _ = v19
        }
    }
    v20 := GenS3{F0: 1, F1: -128, F2: 2}
    for fi21 := 0; fi21 < 17; fi21++ {
        v22 := GenS3{F0: 42, F1: 0, F2: 18446744073709551615}
        _ = fi21
        _ = v22
    }
    v23 := []int{-1, 9223372036854775807, 9223372036854775807}
    v6 = v6
    var v24 int64 = v6
    if false {
        if (!(18446744073709551615 >= 0)) {
            v25 := GenS4{F0: 0, F1: 2147483648, F2: 1}
            v8 -= (v8 / (((v8 & 2)) | 1))
            fmt.Println(v8, v1, v8)
            _ = v25
        } else {
            fmt.Println(v6)
        }
        v26 := GenS5{F0: 2147483647}
        var v27 uint8 = 180
        _ = v26
        _ = v27
    } else {
        v28 := GenS6{F0: (127 > 0)}
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
