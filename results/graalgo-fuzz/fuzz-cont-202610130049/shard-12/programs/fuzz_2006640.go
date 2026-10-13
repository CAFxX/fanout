package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 string; F2 uint }
type GenS4 struct { F0 int; F1 uint64; F2 bool }
type GenS5 struct { F0 uint16; F1 int; F2 int8 }
type GenS6 struct { F0 uint8 }
func genFn0(p0 uint) int8 {
    v7 := GenS1{F0: 465}
    fmt.Println(p0)
    _ = p0
    _ = v7
    return -128
}
func genFn1(p0 int16) int8 {
    switch "" {
    case "":
        v8 := GenS2{F0: (42 != -1)}
        _ = v8
    case "a":
        var v9 bool = (737 == 92)
        _ = v9
    default:
        v10 := []int{1}
        _ = v10
    }
    v11 := GenS3{F0: 4294967295, F1: (((("foo bar" + "\n\t") + "foo bar") + ("a" + ("" + ""))) + ("a" + ("foo bar" + ""))), F2: 4294967295}
    _ = p0
    _ = v11
    return -128
}

func main() {
    var v1 uint64 = 4294967295
    var v2 uint = 42
    var v3 float32 = 0.1
    v4 := make(map[uint32]int)
    var v5 string = ""
    v6 := []uint32{4294967295, 856}
    for fi12 := 0; fi12 < 4; fi12++ {
        if false {
            for fi13 := 0; fi13 < 2; fi13++ {
                fmt.Println(v3, v3)
                _ = fi13
            }
            fmt.Println(fi12, v3)
            switch fi12 {
            case 766:
                v14 := GenS4{F0: (fi12 % (((fi12 << ((18446744073709551615) & 63))) | 1)), F1: (v1 << ((v2) & 63)), F2: (true || ((!(0 > (fi12 + (^(fi12))))) || (4294967295 > 607)))}
                _ = v14
            }
        }
        _ = fi12
    }
    switch ("foo bar" + (v5 + "foo bar")) {
    case "a":
        v3 -= (((v3 / 1.0) + v3) / (0.0 - v3))
    case "foo bar":
        if false {
            var v15 bool = (42 > 128)
            v6 = append(v6, 42)
            var v16 uint8 = 42
            _ = v15
            _ = v16
        }
    }
    fmt.Println(v5, v1)
    var v17 uint8 = 1
    v17 = (v17 << ((42) & 63))
    v1 = 2
    var v18 string = v5
    var v19 int32 = 2147483647
    v20 := GenS2{F0: (!((v19 | v19) == v19))}
    if false {
        v21 := GenS5{F0: 0, F1: 0, F2: int8(v2)}
        _ = v21
    }
    v17 = 1
    var v22 uint16 = 42
    v23 := GenS6{F0: 255}
    v18 = v18
    var v24 int32 = (v19 << (((v2 | 4294967295)) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
