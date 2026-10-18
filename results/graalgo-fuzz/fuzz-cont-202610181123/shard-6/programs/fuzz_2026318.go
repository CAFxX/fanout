package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint8 }
type GenS2 struct { F0 int; F1 uint64; F2 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint16; F1 bool }
type GenS5 struct { F0 uint16; F1 int; F2 int32 }
type GenS6 struct { F0 uint64; F1 bool; F2 uint8 }
func genFn0() uint8 {
    var v4 uint32 = 42
    fmt.Println(v4)
    var v5 uint32 = (v4 >> ((2) & 63))
    v5 -= v5
    _ = v4
    _ = v5
    return 1
}

func main() {
    var v1 int = 943
    var v2 int8 = 42
    v3 := []int{}
    fmt.Println(v2, v2)
    fmt.Println(v1)
    var v6 uint = 409
    var v7 float32 = -2.25
    var v8 float64 = 0.1
    var v9 string = "foo bar"
    var v10 uint16 = 294
    var v11 int64 = 42
    var v12 int8 = (^(v2))
    var v13 int64 = -9223372036854775808
    switch (v1 | v1) {
    case 127:
        var v14 int64 = 1
        _ = v14
    case 32767:
        var v15 int32 = -255
        _ = v15
    case -1:
        for fi16 := 0; fi16 < 10; fi16++ {
            if true {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            } else {
                var v17 string = "a"
                v18 := GenS1{F0: int64(v12), F1: 1, F2: 2}
                _ = v17
                _ = v18
            }
            fmt.Println(v7, v11, v1)
            _ = fi16
        }
    default:
        var v19 float64 = v8
        _ = v19
    }
    v20 := GenS2{F0: (^(v1)), F1: 0, F2: v1}
    if (v8 < -2.25) {
        v21 := GenS3{F0: (^(v12))}
        switch v9 {
        case "foo bar":
            v22 := GenS4{F0: 65535, F1: true}
            _ = v22
        default:
            fmt.Println(v7, v13, v8)
        }
        _ = v21
    } else {
        _ = len(v3)
        if (((^(v12)) == (v2 * (-(v2)))) || false) {
            v1 += v1
            switch "\n\t" {
            case "foo bar":
                var v23 uint32 = 0
                _ = v23
            case "a":
                for fi24 := 0; fi24 < 16; fi24++ {
                    _ = len(v3)
                    for fi25 := 0; fi25 < 12; fi25++ {
                        v13 += (v11 & (v13 | (v13 >> (((^(v6))) & 63))))
                        v12 = -128
                        _ = fi25
                    }
                    v9 = ("foo bar" + v9)
                    _ = fi24
                }
            }
        } else {
            v26 := GenS5{F0: 255, F1: (^(v1)), F2: -1}
            _ = v26
        }
    }
    v27 := GenS6{F0: 18446744073709551615, F1: ((v6 >= (^(v6))) && false), F2: 14}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
