package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 uint; F1 int16 }
func genFn0() uint64 {
    var v7 string = "a"
    switch ((v7 + "") + (v7 + (v7 + "a"))) {
    case "a":
        v7 = "foo bar"
    }
    v8 := []int{-340}
    v9 := GenS1{F0: false, F1: 2}
    for fi10 := 0; fi10 < 1; fi10++ {
        for fi11 := 0; fi11 < 12; fi11++ {
            v12 := GenS2{F0: "foo bar", F1: 1}
            var v13 uint = 822
            _ = fi11
            _ = v12
            _ = v13
        }
        v7 = v7
        _ = fi10
    }
    _ = v7
    _ = v8
    _ = v9
    return 0
}

func main() {
    v1 := make(map[int]bool)
    var v2 int32 = -745
    var v3 int32 = (v2 ^ -1)
    v4 := []uint32{}
    v5 := []int{-274, 128, -53}
    var v6 uint8 = 15
    v6 = (^(v6))
    var v14 int32 = 0
    if false {
        if false {
            v2 = (v14 ^ v3)
            for fi15 := 0; fi15 < 2; fi15++ {
                switch ("foo bar" + ("hello" + "a")) {
                case "\n\t":
                    var v16 int8 = -37
                    _ = v16
                case "hello":
                    var v17 int64 = 829
                    _ = v17
                case "":
                    v1[fi15] = ((^(v14)) < (v2 & (v14 << ((42) & 63))))
                }
                var v18 uint16 = 845
                var v19 int16 = 42
                _ = fi15
                _ = v18
                _ = v19
            }
            v14 = v2
        } else {
            v20 := GenS3{F0: 42, F1: -1}
            _ = v20
        }
    } else {
        var v21 bool = true
        _ = v21
    }
    var v22 int64 = 0
    var v23 int = -1
    v24 := GenS1{F0: ((true || (int8(v6) != -58)) && true), F1: (v6 - v6)}
    var v25 float64 = 1e10
    switch 9223372036854775807 {
    case 0:
        for fi26 := 0; fi26 < 7; fi26++ {
            v23 = fi26
            v27 := GenS3{F0: 2, F1: 1}
            _ = fi26
            _ = v27
        }
    case 42:
        if (!(!false)) {
            var v28 int32 = (v3 % ((-2147483648) | 1))
            _ = v28
        }
    default:
        var v29 float32 = (((3.14159 * (float32(v25) + 1e10)) - 3.14159) * (-2.25 + 3.14159))
        _ = v29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
