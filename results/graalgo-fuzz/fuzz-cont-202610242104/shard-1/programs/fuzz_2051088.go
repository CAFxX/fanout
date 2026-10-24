package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64; F1 int; F2 int16 }
type GenS4 struct { F0 int16; F1 uint64 }
type GenS5 struct { F0 int; F1 bool }
func genFn0() int {
    var v7 uint32 = 479
    switch "a" {
    case "hello":
        v8 := GenS1{F0: v7, F1: "foo bar"}
        _ = v8
    default:
        var v9 uint64 = uint64(v7)
        _ = v9
    }
    if false {
        v7 = (v7 + (^(v7)))
        var v10 int32 = -2147483648
        var v11 uint = 648
        _ = v10
        _ = v11
    }
    _ = v7
    return 4294967295
}

func main() {
    var v1 uint = 42
    var v2 bool = (false || false)
    var v3 uint32 = 65535
    var v4 uint32 = (^(v3))
    var v5 int8 = 127
    var v6 bool = true
    if (v1 == (v1 | 58)) {
        var v12 int8 = v5
        v13 := []int{127, -2147483648, -1}
        for fi14 := 0; fi14 < 4; fi14++ {
            var v15 uint64 = 42
            var v16 int = (fi14 | genFn0())
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = v12
        _ = v13
    }
    var v17 float64 = -2.25
    var v18 int32 = 0
    fmt.Println(v1)
    v6 = (true || (v6 || false))
    v19 := GenS1{F0: v4, F1: ("" + "foo bar")}
    v20 := []int{}
    if ((float64(((0.0 / 0.1) / 1.5)) / v17) != v17) {
        v21 := GenS2{F0: -1}
        if len(v20) >= 2 {
            v20 = v20[1:len(v20)]
        }
        var v22 int8 = v5
        _ = v21
        _ = v22
    }
    if true {
        if (!(2 > 569)) {
            if len(v20) > 0 {
                v20[0] = -736
            }
            for fi23 := 0; fi23 < 19; fi23++ {
                _ = len(v20)
                v2 = ("a" >= "foo bar")
                v24 := GenS3{F0: 2, F1: (fi23 << ((v1) & 63)), F2: -366}
                _ = fi23
                _ = v24
            }
            v5 += 1
        } else {
            for fi25 := 0; fi25 < 20; fi25++ {
                v26 := GenS4{F0: 1000, F1: 18446744073709551615}
                if (!false) {
                    var v27 uint32 = v4
                    _ = v27
                }
                _ = fi25
                _ = v26
            }
            v4 -= (^(v4))
        }
        v28 := GenS5{F0: 9223372036854775807, F1: (v6 || false)}
        var v29 uint = v1
        _ = v28
        _ = v29
    }
    fmt.Println(v6, v6, v17)
    v5 -= v5
    var v30 float32 = 0.1
    v4 = 0
    v17 = v17
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
