package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 int16 }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int32; F1 bool; F2 uint8 }

func main() {
    v1 := []int{-2147483648, 0}
    var v2 string = "hello"
    var v3 uint = 65535
    for fi4 := 0; fi4 < 19; fi4++ {
        switch "foo bar" {
        case "":
            v5 := GenS1{F0: 9223372036854775807, F1: 0, F2: 32767}
            _ = v5
        }
        for fi6 := 0; fi6 < 8; fi6++ {
            switch fi4 {
            case 1:
                if false {
                    _ = len(v1)
                    var v7 int64 = 9223372036854775807
                    _ = v7
                }
            }
            var v8 int = (^(fi6))
            _ = fi6
            _ = v8
        }
        _ = fi4
    }
    var v9 int16 = 1000
    if (((v2 + "foo bar") + v2) != "foo bar") {
        if len(v1) > 0 {
            v1[0] = -280
        }
        _ = len(v1)
    }
    v10 := GenS2{F0: (^(v3)), F1: 0}
    v3 = 255
    _ = len(v1)
    v3 -= (v3 & 4294967295)
    if true {
        v11 := GenS1{F0: 466, F1: 925, F2: (v9 << (((v3 & 2)) & 63))}
        var v12 int64 = 1
        _ = v11
        _ = v12
    }
    var v13 uint16 = 65535
    switch "" {
    case "":
        switch 128 {
        case 4294967295:
            var v14 uint16 = 2
            _ = v14
        }
    default:
        if len(v1) > 0 {
            v1[0] = -2147483648
        }
    }
    v15 := GenS3{F0: -207}
    v16 := GenS2{F0: 65535, F1: 18446744073709551615}
    var v17 int16 = -923
    for fi18 := 0; fi18 < 13; fi18++ {
        fmt.Println(v3)
        var v19 uint32 = 192
        var v20 string = ((v2 + v2) + "a")
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 uint = 365
    var v22 string = "a"
    for fi23 := 0; fi23 < 1; fi23++ {
        v24 := GenS4{F0: 2147483647, F1: (1.5 >= ((3.14159 / (0.1 - (-2.25 / 1e10))) + (1e10 - (0.0 + ((((-2.25 / -2.25) / 1.0) - 3.14159) * ((-2.25 + (0.1 - (-2.25 / 1e10))) * (1e10 * (1.5 - 1e10)))))))), F2: 32}
        _ = fi23
        _ = v24
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v9
}
