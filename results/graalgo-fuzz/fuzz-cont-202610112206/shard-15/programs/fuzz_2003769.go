package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 int32; F1 int16 }
func genFn0() int8 {
    if false {
        v7 := []int{-1}
        _ = v7
    } else {
        v8 := []int{127, 127, -1}
        var v9 float32 = (1.5 - float32(0.0))
        _ = v8
        _ = v9
    }
    var v10 uint64 = 4294967295
    for fi11 := 0; fi11 < 10; fi11++ {
        switch fi11 {
        case 742:
            var v12 uint = 18446744073709551615
            _ = v12
        case 963:
            for fi13 := 0; fi13 < 18; fi13++ {
                v14 := []int{}
                var v15 string = "hello"
                _ = fi13
                _ = v14
                _ = v15
            }
        default:
            var v16 bool = (1 != 479)
            _ = v16
        }
        _ = fi11
    }
    for fi17 := 0; fi17 < 13; fi17++ {
        fmt.Println(v10, fi17)
        _ = fi17
    }
    var v18 int = -9223372036854775808
    _ = v10
    _ = v18
    return -1
}

func main() {
    var v1 uint64 = 2
    v2 := make(map[int]int)
    var v3 string = "a"
    v4 := make(map[int]uint32)
    var v5 int = 1
    v6 := make(map[int]bool)
    v19 := GenS1{F0: -9223372036854775808, F1: v1}
    var v20 string = v3
    for fi21 := 0; fi21 < 14; fi21++ {
        v3 = (((v3 + "\n\t") + (v3 + v3)) + (v3 + v20))
        _ = fi21
    }
    switch "\n\t" {
    case "\n\t":
        v22 := GenS1{F0: 42, F1: v1}
        _ = v22
    case "foo bar":
        v20 = v20
    default:
        v23 := GenS1{F0: 9223372036854775807, F1: v1}
        _ = v23
    }
    v1 += v1
    v24 := []int{}
    if (1 <= 255) {
        v25 := GenS2{F0: -1, F1: 127}
        v3 = v20
        if (false || false) {
            for fi26 := 0; fi26 < 18; fi26++ {
                fmt.Println(fi26, v5, v20)
                _ = fi26
            }
            fmt.Println(v5, v1)
        }
        _ = v25
    } else {
        if false {
            switch (-(v5)) {
            case 2147483647:
                v27 := GenS3{F0: 546, F1: 521}
                _ = v27
            case -647:
                v20 = v3
            }
            if (v1 != (v1 | (v1 ^ 42))) {
                switch v5 {
                case -395:
                    fmt.Println(v3, v3, v1)
                case -305:
                    if len(v24) > 0 {
                        v24[0] = (v5 << ((495) & 63))
                    }
                default:
                    v28 := GenS3{F0: 42, F1: 32767}
                    _ = v28
                }
            } else {
                v3 = "a"
                fmt.Println(v5)
            }
        }
    }
    if true {
        v3 = ("" + v20)
        v29 := GenS2{F0: 9223372036854775807, F1: 0}
        v20 = v3
        _ = v29
    }
    v3 = ((v3 + "") + ("hello" + (((("a" + "hello") + v3) + v3) + "\n\t")))
    v1 = (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
