package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 string }

func main() {
    v1 := make(map[int]string)
    var v2 uint64 = 300
    v3 := make(map[uint32]bool)
    var v4 uint32 = 1
    v5 := []uint32{}
    var v6 uint8 = 148
    var v7 bool = true
    for fi8 := 0; fi8 < 18; fi8++ {
        switch fi8 {
        case -818:
            var v9 uint16 = 2
            _ = v9
        case -2147483648:
            v6 = 127
        default:
            if ((!(false && true)) || (v2 > (v2 * (v2 + (v2 % ((151) | 1)))))) {
                v10 := GenS1{F0: 407, F1: int32(v6), F2: ""}
                _ = v10
            }
        }
        v11 := GenS1{F0: 1000, F1: 42, F2: ("a" + "")}
        _ = fi8
        _ = v11
    }
    v12 := GenS1{F0: 1000, F1: 647, F2: "\n\t"}
    v6 -= 72
    var v13 int16 = 42
    v14 := GenS1{F0: (v13 << ((4294967295) & 63)), F1: 679, F2: ("" + "\n\t")}
    var v15 uint8 = (v6 * 128)
    var v16 int64 = -702
    v17 := GenS1{F0: (-(v13)), F1: 2147483647, F2: "foo bar"}
    var v18 int8 = 127
    var v19 bool = (v4 <= v4)
    fmt.Println(v18, v6, v4)
    switch ("" + "foo bar") {
    case "":
        var v20 int32 = 366
        _ = v20
    case "foo bar":
        v2 += v2
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
