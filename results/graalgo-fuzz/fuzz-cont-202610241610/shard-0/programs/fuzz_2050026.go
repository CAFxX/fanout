package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 int }
func genFn0() uint {
    v4 := GenS1{F0: 657, F1: "a"}
    v5 := []int{0, 32768}
    v6 := GenS1{F0: 267, F1: ("hello" + ("\n\t" + ""))}
    _ = v4
    _ = v5
    _ = v6
    return 122
}

func main() {
    var v1 uint16 = 752
    v2 := make(map[string]bool)
    v3 := make(map[uint32]bool)
    v7 := []int{2147483647}
    _ = len(v7)
    for fi8 := 0; fi8 < 6; fi8++ {
        v1 = 1
        switch ("\n\t" + "a") {
        case "":
            v1 -= (v1 | (v1 - 387))
        case "\n\t":
            v9 := GenS2{F0: 888}
            _ = v9
        case "foo bar":
            var v10 int = (fi8 / (((fi8 % ((fi8) | 1))) | 1))
            _ = v10
        }
        _ = fi8
    }
    if false {
        var v11 float64 = -2.25
        _ = v11
    } else {
        var v12 bool = (!(18446744073709551615 >= 932))
        var v13 bool = (v12 && (!v12))
        _ = v12
        _ = v13
    }
    var v14 uint64 = 4294967295
    var v15 uint8 = 128
    if len(v7) > 0 {
        v7[0] = -2147483648
    }
    v1 -= (v1 % (((v1 & 555)) | 1))
    var v16 string = "hello"
    var v17 string = "\n\t"
    switch 128 {
    case 127:
        v2[(v17 + v17)] = (float64((1e10 + 0.0)) != (0.1 / 3.14159))
    case -6:
        switch v17 {
        case "\n\t":
            var v18 int = 4294967295
            _ = v18
        }
    case 128:
        var v19 uint = 0
        _ = v19
    }
    v20 := GenS2{F0: 32767}
    v14 -= v14
    delete(v3, 2147483648)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v7
}
