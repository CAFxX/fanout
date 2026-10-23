package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 bool }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 uint8; F1 uint8 }

func main() {
    var v1 int32 = 2147483647
    var v2 int64 = 9223372036854775807
    var v3 int32 = 1
    v4 := []int{}
    fmt.Println(v2, v3, v3)
    var v5 float32 = 3.14159
    fmt.Println(v1)
    for fi6 := 0; fi6 < 2; fi6++ {
        fmt.Println(v1)
        var v7 int64 = (v2 / ((v2) | 1))
        var v8 int8 = -124
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 0, F1: -128, F2: true}
    switch ((("" + ("" + ("\n\t" + "a"))) + "hello") + ("hello" + "")) {
    case "\n\t":
        v10 := GenS2{F0: -40, F1: 127}
        _ = v10
    case "hello":
        var v11 int32 = (v3 - (v1 - (-(v1))))
        _ = v11
    default:
        v4 = append(v4, -415)
    }
    var v12 int32 = (^(v3))
    var v13 string = ("foo bar" + "foo bar")
    v3 = (v1 << ((42) & 63))
    if len(v4) > 0 {
        v4[0] = 284
    }
    switch -861 {
    case 42:
        var v14 uint8 = 128
        _ = v14
    case 519:
        v15 := GenS2{F0: -50, F1: 91}
        _ = v15
    case -1:
        for fi16 := 0; fi16 < 6; fi16++ {
            var v17 uint = 4294967295
            v12 = v1
            v18 := GenS2{F0: -1, F1: 18}
            _ = fi16
            _ = v17
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 8; fi19++ {
        if false {
            for fi20 := 0; fi20 < 8; fi20++ {
                v21 := GenS1{F0: 0, F1: -4, F2: false}
                switch v13 {
                case "foo bar":
                    v22 := GenS3{F0: 32, F1: 128}
                    _ = v22
                case "hello":
                    if len(v4) > 0 {
                        v4[0] = fi20
                    }
                }
                v3 += (v3 ^ (v12 & v12))
                _ = fi20
                _ = v21
            }
            v23 := GenS1{F0: 42, F1: -128, F2: (!(v2 > 42))}
            _ = v23
        }
        v2 = v2
        _ = len(v4)
        _ = fi19
    }
    if len(v4) > 0 {
        v4[0] = 42
    }
    v13 = ("" + v13)
    v24 := GenS1{F0: 327, F1: -1, F2: (78 > 2147483647)}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
