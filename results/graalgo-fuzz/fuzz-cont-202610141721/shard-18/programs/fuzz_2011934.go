package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 bool; F1 uint64 }
func genFn0(p0 uint, p1 uint8, p2 string) string {
    var v6 int32 = 2147483647
    v7 := GenS1{F0: true, F1: 2}
    v6 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return p2
}
func genFn1() uint16 {
    var v8 uint8 = 42
    switch -9223372036854775808 {
    case 631:
        for fi9 := 0; fi9 < 8; fi9++ {
            for fi10 := 0; fi10 < 9; fi10++ {
                v11 := []int{(fi10 / ((128) | 1)), fi9, (fi10 % (((fi10 | fi10)) | 1)), 128}
                _ = fi10
                _ = v11
            }
            _ = fi9
        }
    case -1:
        var v12 uint8 = v8
        _ = v12
    case 127:
        v13 := []int{492, -9223372036854775808}
        _ = v13
    default:
        v8 = v8
    }
    v14 := GenS1{F0: true, F1: 138}
    for fi15 := 0; fi15 < 19; fi15++ {
        v8 -= v8
        _ = fi15
    }
    _ = v14
    _ = v8
    return 905
}

func main() {
    var v1 uint32 = 42
    v2 := make(map[int]uint32)
    var v3 int64 = 699
    v4 := make(map[int]uint32)
    var v5 uint8 = 94
    v16 := []int{-2147483648, 873, 875}
    v1 -= 946
    fmt.Println(v1, v1, v5)
    switch ("\n\t" + "hello") {
    case "a":
        v4[2147483647] = (^(v1))
    case "":
        for fi17 := 0; fi17 < 18; fi17++ {
            v1 += (v1 + (^(v1)))
            var v18 uint8 = v5
            _ = fi17
            _ = v18
        }
    }
    v2[1] = v1
    for fi19 := 0; fi19 < 8; fi19++ {
        if false {
            for fi20 := 0; fi20 < 19; fi20++ {
                var v21 uint16 = 65535
                fmt.Println(fi19, fi20)
                _ = fi20
                _ = v21
            }
            fmt.Println(v3, v5, fi19)
            _ = len(v4)
        }
        _ = fi19
    }
    if (-566 >= 1) {
        var v22 uint32 = v1
        _ = len(v2)
        v23 := GenS1{F0: false, F1: 0}
        _ = v22
        _ = v23
    } else {
        var v24 int64 = 1
        if false {
            fmt.Println(v24)
            _ = len(v16)
        }
        _ = v24
    }
    v25 := GenS2{F0: false, F1: 18446744073709551615}
    var v26 int64 = -1
    fmt.Println(v26, v5)
    for fi27 := 0; fi27 < 4; fi27++ {
        var v28 bool = (v5 == 170)
        _ = fi27
        _ = v28
    }
    var v29 uint32 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
