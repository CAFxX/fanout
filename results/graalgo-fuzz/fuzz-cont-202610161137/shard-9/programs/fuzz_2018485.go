package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 int16 }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 uint32 }

func main() {
    var v1 uint32 = 1
    var v2 uint32 = (v1 & v1)
    v3 := []uint32{(^(v2)), v1}
    v4 := make(map[int]bool)
    var v5 float64 = 0.1
    v6 := []uint32{(v2 % ((v1) | 1)), (v2 % ((v2) | 1)), (v1 << ((0) & 63)), v2}
    v7 := GenS1{F0: v2, F1: -1, F2: -32768}
    var v8 uint16 = 606
    if (!true) {
        var v9 int = 9223372036854775807
        var v10 int8 = 100
        v11 := GenS2{F0: (v2 >> ((255) & 63)), F1: ((!(-258 == -2147483648)) || ((v9 | v9) < (v9 << ((4294967295) & 63))))}
        _ = v10
        _ = v11
        _ = v9
    }
    fmt.Println(v8, v2)
    for fi12 := 0; fi12 < 9; fi12++ {
        var v13 uint16 = 0
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 9; fi14++ {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        if (v5 > v5) {
            if false {
                var v15 uint64 = 422
                for fi16 := 0; fi16 < 20; fi16++ {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    var v17 uint = 255
                    var v18 int64 = -128
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                if true {
                    v19 := GenS2{F0: v1, F1: true}
                    _ = v19
                }
                _ = v15
            }
        }
        _ = fi14
    }
    fmt.Println(v5)
    if true {
        var v20 int16 = -32768
        _ = v20
    }
    var v21 string = ("hello" + ("" + ("\n\t" + ((("hello" + "a") + "") + "\n\t"))))
    fmt.Println(v2)
    delete(v4, 32767)
    var v22 int = 32768
    var v23 uint16 = (v8 ^ (v8 | (v8 & (v8 + 65535))))
    switch v21 {
    case "foo bar":
        var v24 uint8 = 117
        _ = v24
    default:
        v22 -= v22
    }
    v25 := GenS3{F0: (v2 + (v2 & 425))}
    fmt.Println(v21, v23, v21)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
