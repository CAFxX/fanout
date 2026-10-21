package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 bool }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 uint32; F1 int; F2 uint16 }

func main() {
    var v1 int64 = -9223372036854775808
    v2 := []int{-2147483648, 2147483647, 1, 728}
    var v3 string = ("\n\t" + "a")
    var v4 uint16 = 627
    v5 := GenS1{F0: 1000, F1: -285, F2: (false || true)}
    var v6 uint16 = v4
    _ = len(v2)
    if false {
        switch v3 {
        case "":
            var v7 bool = (884 > 1000)
            _ = v7
        case "a":
            for fi8 := 0; fi8 < 1; fi8++ {
                var v9 float32 = 1e10
                switch v3 {
                case "foo bar":
                    var v10 int16 = -583
                    _ = v10
                case "\n\t":
                    fmt.Println(fi8, v3)
                case "":
                    v11 := GenS1{F0: 349, F1: 42, F2: (fi8 < 773)}
                    _ = v11
                }
                fmt.Println(v1, v1, v9)
                _ = fi8
                _ = v9
            }
        default:
            v12 := GenS2{F0: -2147483648, F1: uint64(v4)}
            _ = v12
        }
        var v13 int16 = -167
        _ = v13
    }
    v14 := GenS1{F0: 42, F1: 42, F2: true}
    v15 := GenS1{F0: 0, F1: 42, F2: false}
    v1 -= v1
    var v16 uint16 = 255
    if false {
        var v17 uint32 = 65535
        if false {
            v18 := GenS2{F0: 42, F1: 18446744073709551615}
            var v19 int16 = 1000
            _ = v18
            _ = v19
        } else {
            for fi20 := 0; fi20 < 18; fi20++ {
                fmt.Println(fi20, v6)
                v17 = v17
                v2 = append(v2, fi20)
                _ = fi20
            }
            v21 := GenS3{F0: 218, F1: 36, F2: v4}
            v6 = v4
            _ = v21
        }
        fmt.Println(v4)
        _ = v17
    } else {
        var v22 float64 = 0.0
        if (true || true) {
            if len(v2) > 0 {
                v2[0] = -431
            }
        }
        _ = v22
    }
    v4 = v4
    var v23 uint = 1
    if len(v2) > 0 {
        v2[0] = -683
    }
    var v24 string = ((("\n\t" + (v3 + (v3 + ""))) + ("\n\t" + "hello")) + v3)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
