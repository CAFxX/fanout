package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 uint }
func genFn0() int32 {
    var v7 int = 1
    v8 := []int{v7, v7, 32768, v7}
    if true {
        if false {
            fmt.Println(v7)
        }
    }
    if (!(!(((-2147483648 == -2147483648) && false) || (634 < 0)))) {
        var v9 int32 = 259
        var v10 int = v7
        _ = v10
        _ = v9
    } else {
        var v11 string = "hello"
        _ = v11
    }
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 int32 = 1
    var v2 uint16 = 95
    v3 := []uint32{4294967295, 65535, 407, 365}
    var v4 uint = 709
    v5 := make(map[uint32]string)
    v6 := []string{"\n\t", "foo bar"}
    var v12 int64 = 0
    switch 215 {
    case 490:
        for fi13 := 0; fi13 < 20; fi13++ {
            v3 = append(v3, 2147483647)
            _ = fi13
        }
    case 2147483647:
        v14 := GenS1{F0: v1, F1: -58, F2: v4}
        _ = v14
    default:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    _ = len(v3)
    var v15 int = 2147483647
    var v16 uint8 = 83
    var v17 uint64 = 151
    var v18 int32 = genFn0()
    for fi19 := 0; fi19 < 18; fi19++ {
        v12 = (-(v12))
        _ = fi19
    }
    fmt.Println(v4)
    switch v15 {
    case 127:
        var v20 uint64 = (v17 * 2)
        _ = v20
    case 9223372036854775807:
        var v21 int8 = 23
        _ = v21
    default:
        for fi22 := 0; fi22 < 12; fi22++ {
            v16 += 2
            for fi23 := 0; fi23 < 16; fi23++ {
                var v24 uint64 = (v17 - v17)
                _ = fi23
                _ = v24
            }
            var v25 int8 = 8
            _ = fi22
            _ = v25
        }
    }
    var v26 int64 = -596
    if ((v18 - int32(v16)) > 1) {
        var v27 uint = 42
        if (!(9223372036854775807 != 933)) {
            v15 = (v15 & (v15 << (((v4 % (((^(v27))) | 1))) & 63)))
        }
        v18 = v1
        _ = v27
    }
    var v28 int16 = -32768
    switch "foo bar" {
    case "foo bar":
        _ = len(v6)
    case "a":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "\n\t":
        var v29 float32 = 1.5
        _ = v29
    default:
        var v30 uint64 = (v17 % ((uint64(v2)) | 1))
        _ = v30
    }
    fmt.Println(v4)
    v12 -= (^(v26))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v26
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
