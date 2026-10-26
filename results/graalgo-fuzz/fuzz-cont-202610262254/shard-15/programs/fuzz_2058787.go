package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 uint64; F2 string }
type GenS3 struct { F0 int8; F1 int64; F2 int8 }

func main() {
    v1 := []int{-9223372036854775808, 1}
    var v2 int32 = 1
    var v3 uint32 = 0
    var v4 bool = false
    v5 := []int{128, 800}
    var v6 string = "a"
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 uint = 255
        _ = fi7
        _ = v8
    }
    if v4 {
        var v9 int8 = 127
        _ = v9
    }
    if (-1 != -1) {
        v2 += -1
    }
    v10 := GenS1{F0: 1}
    v2 = (v2 / ((v2) | 1))
    var v11 int32 = -2147483648
    if v4 {
        v12 := GenS2{F0: 4294967295, F1: 0, F2: "\n\t"}
        v5 = append(v5, 693)
        _ = v12
    }
    var v13 int8 = -128
    _ = len(v5)
    switch 961 {
    case 1:
        var v14 int32 = v2
        _ = v14
    case 212:
        v4 = true
    case 4294967295:
        var v15 int64 = 9223372036854775807
        _ = v15
    default:
        _ = len(v5)
    }
    var v16 int64 = 42
    v17 := GenS1{F0: 32767}
    switch v6 {
    case "a":
        if v4 {
            if true {
                v6 = ""
                for fi18 := 0; fi18 < 16; fi18++ {
                    v1 = append(v1, 127)
                    _ = fi18
                }
                v3 = (^(v3))
            }
        }
    case "hello":
        var v19 int64 = 1
        _ = v19
    default:
        if (18446744073709551615 >= 2) {
            for fi20 := 0; fi20 < 14; fi20++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi20
            }
        } else {
            fmt.Println(v2, v13)
        }
    }
    for fi21 := 0; fi21 < 17; fi21++ {
        _ = len(v1)
        var v22 float64 = (1.5 / (1e10 / (3.14159 - ((0.1 - 1e10) * 1.5))))
        v23 := GenS3{F0: v13, F1: v16, F2: (v13 % (((v13 | v13)) | 1))}
        _ = fi21
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 8; fi24++ {
        v13 = v13
        v3 = (v3 - (v3 & (v3 - 2)))
        fmt.Println(fi24)
        _ = fi24
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
