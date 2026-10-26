package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 int32; F2 uint32 }
type GenS3 struct { F0 uint64; F1 uint32; F2 string }
func genFn0(p0 uint16, p1 int64) uint {
    v5 := GenS1{F0: 255}
    var v6 int8 = 100
    v7 := GenS1{F0: 2}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 65535
}
func genFn1() int16 {
    var v8 uint64 = 162
    var v9 uint64 = v8
    v10 := []int{835}
    v11 := GenS2{F0: 24, F1: -2147483648, F2: 2147483648}
    if len(v10) > 0 {
        v10[0] = 2147483647
    }
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return 594
}

func main() {
    var v1 int32 = -1
    var v2 int16 = 285
    var v3 uint16 = 0
    v4 := make(map[int]string)
    if false {
        var v12 uint = genFn0(v3, 731)
        switch -9223372036854775808 {
        case 9223372036854775807:
            for fi13 := 0; fi13 < 17; fi13++ {
                var v14 uint16 = 0
                var v15 int16 = v2
                _ = fi13
                _ = v14
                _ = v15
            }
        case 2147483647:
            v1 = (v1 / ((v1) | 1))
        case 127:
            var v16 uint64 = 0
            _ = v16
        default:
            switch "\n\t" {
            case "hello":
                if (!(-1 > -195)) {
                    v3 = v3
                    v17 := []int{173, 1, 623, -448}
                    _ = v17
                } else {
                    var v18 string = ("\n\t" + "a")
                    _ = v18
                }
            case "\n\t":
                v19 := GenS3{F0: 4294967295, F1: 4294967295, F2: "foo bar"}
                _ = v19
            }
        }
        var v20 float64 = 0.0
        _ = v12
        _ = v20
    } else {
        var v21 int32 = (v1 | v1)
        _ = v21
    }
    switch (((("\n\t" + "a") + (("foo bar" + "\n\t") + "hello")) + "a") + ("a" + (("foo bar" + "\n\t") + ""))) {
    case "hello":
        for fi22 := 0; fi22 < 8; fi22++ {
            switch "foo bar" {
            case "foo bar":
                v23 := GenS1{F0: 1}
                _ = v23
            case "":
                var v24 int16 = v2
                _ = v24
            case "a":
                var v25 uint = 0
                _ = v25
            default:
                v1 = (v1 - (v1 | 977))
            }
            _ = fi22
        }
    case "a":
        switch -1 {
        case 0:
            v3 = (v3 | (v3 + v3))
        case 1:
            var v26 float64 = (-2.25 - 0.1)
            _ = v26
        default:
            var v27 uint8 = 1
            _ = v27
        }
    case "\n\t":
        for fi28 := 0; fi28 < 19; fi28++ {
            for fi29 := 0; fi29 < 5; fi29++ {
                v2 = v2
                v1 -= v1
                fmt.Println(v1, fi29)
                _ = fi29
            }
            v3 -= 262
            _ = fi28
        }
    default:
        for fi30 := 0; fi30 < 15; fi30++ {
            v2 -= (v2 ^ v2)
            _ = fi30
        }
    }
    if ((^(v1)) > 657) {
        var v31 int8 = 42
        _ = v31
    }
    var v32 uint = 42
    var v33 uint = v32
    v34 := []int{-9223372036854775808, -2147483648}
    var v35 int8 = int8(v3)
    if len(v34) > 0 {
        v34[0] = 32767
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
}
