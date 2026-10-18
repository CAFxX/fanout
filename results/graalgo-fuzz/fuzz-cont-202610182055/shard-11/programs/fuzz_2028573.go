package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
func genFn0() int8 {
    var v4 uint8 = 42
    var v5 uint16 = 550
    v6 := []int{}
    _ = v4
    _ = v5
    _ = v6
    return 100
}
func genFn1() uint64 {
    v7 := GenS1{F0: 973, F1: 973}
    switch "hello" {
    case "\n\t":
        var v8 int = 32768
        _ = v8
    case "foo bar":
        v9 := []int{1, -9223372036854775808}
        _ = v9
    case "hello":
        var v10 uint32 = 2
        _ = v10
    }
    _ = v7
    return 4294967295
}

func main() {
    var v1 float64 = ((0.0 * 1e10) - 1.5)
    v2 := make(map[string]uint32)
    v3 := make(map[string]bool)
    v1 -= v1
    var v11 string = "\n\t"
    if (!false) {
        v12 := []int{}
        _ = len(v12)
        _ = v12
    }
    if (true || false) {
        v1 += v1
        var v13 int = 892
        _ = v13
    } else {
        for fi14 := 0; fi14 < 3; fi14++ {
            for fi15 := 0; fi15 < 6; fi15++ {
                v16 := []int{(-(fi14)), (fi15 >> ((733) & 63))}
                v11 = ("\n\t" + (((((v11 + v11) + "") + v11) + ((v11 + ((v11 + v11) + (v11 + v11))) + v11)) + v11))
                _ = fi15
                _ = v16
            }
            var v17 int64 = -627
            var v18 int64 = (v17 << ((953) & 63))
            _ = fi14
            _ = v17
            _ = v18
        }
    }
    switch 32767 {
    case -290:
        v19 := []int{-988, 128, 515, 1}
        _ = v19
    case 127:
        v1 = v1
    default:
        var v20 string = v11
        _ = v20
    }
    for fi21 := 0; fi21 < 5; fi21++ {
        v22 := GenS1{F0: 1, F1: 1}
        var v23 int32 = 1
        v24 := GenS1{F0: 0, F1: 0}
        _ = fi21
        _ = v22
        _ = v23
        _ = v24
    }
    var v25 uint = 16
    v25 = (^(v25))
    for fi26 := 0; fi26 < 9; fi26++ {
        v1 = -2.25
        v27 := []int{(fi26 - (-(fi26))), fi26}
        _ = fi26
        _ = v27
    }
    if (!(false && true)) {
        _ = len(v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v2
    _ = v25
    _ = v3
}
