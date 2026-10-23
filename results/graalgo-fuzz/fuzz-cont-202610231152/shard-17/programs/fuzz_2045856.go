package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 int; F1 uint64; F2 uint8 }
type GenS3 struct { F0 int; F1 uint32; F2 uint }
type GenS4 struct { F0 int32; F1 uint64 }
type GenS5 struct { F0 int; F1 uint8 }

func main() {
    var v1 int64 = -1
    var v2 uint16 = 592
    var v3 int8 = 1
    var v4 int32 = 222
    var v5 int16 = -32768
    v2 = v2
    var v6 uint = 471
    var v7 float32 = 1.5
    v8 := []int{9223372036854775807, -2147483648}
    v9 := GenS1{F0: 128, F1: (v4 << (((v6 % ((1) | 1))) & 63))}
    v10 := GenS1{F0: 42, F1: 0}
    v5 = -32768
    for fi11 := 0; fi11 < 12; fi11++ {
        v6 += v6
        var v12 string = "hello"
        _ = fi11
        _ = v12
    }
    _ = len(v8)
    switch 32768 {
    case 32768:
        fmt.Println(v6, v5, v3)
    case 42:
        var v13 bool = false
        _ = v13
    default:
        fmt.Println(v3)
    }
    v14 := GenS1{F0: 1, F1: (v4 ^ (^(v4)))}
    _ = len(v8)
    v1 = (-(v1))
    if true {
        var v15 float32 = v7
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        if ((v3 >> ((v6) & 63)) == v3) {
            v16 := GenS1{F0: 128, F1: v4}
            _ = v16
        } else {
            var v17 int8 = v3
            v18 := GenS2{F0: 32768, F1: 18446744073709551615, F2: 210}
            var v19 int8 = v17
            _ = v17
            _ = v18
            _ = v19
        }
        _ = v15
    }
    var v20 uint16 = v2
    if true {
        if true {
            v21 := GenS3{F0: 127, F1: 371, F2: v6}
            v22 := GenS4{F0: v4, F1: 4294967295}
            _ = v21
            _ = v22
        }
        if (1 != 940) {
            v23 := GenS1{F0: 255, F1: v4}
            _ = v23
        }
        v3 += (v3 | (v3 ^ 100))
    } else {
        for fi24 := 0; fi24 < 18; fi24++ {
            var v25 float32 = (1.5 - v7)
            for fi26 := 0; fi26 < 3; fi26++ {
                var v27 int = fi24
                var v28 int8 = (v3 >> ((v6) & 63))
                _ = fi26
                _ = v27
                _ = v28
            }
            if false {
                var v29 uint16 = (^(v2))
                if true {
                    v30 := GenS4{F0: v4, F1: 18446744073709551615}
                    _ = v30
                }
                _ = v29
            } else {
                v31 := GenS1{F0: 42, F1: -2147483648}
                v32 := GenS5{F0: fi24, F1: 42}
                for fi33 := 0; fi33 < 7; fi33++ {
                    fmt.Println(v6)
                    v4 += (v4 ^ v4)
                    v6 -= v6
                    _ = fi33
                }
                _ = v31
                _ = v32
            }
            _ = fi24
            _ = v25
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
