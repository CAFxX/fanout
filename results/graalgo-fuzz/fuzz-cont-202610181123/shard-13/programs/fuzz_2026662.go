package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int16 }
type GenS2 struct { F0 uint8; F1 int }
func genFn0() uint32 {
    switch "hello" {
    case "hello":
        var v4 string = "hello"
        _ = v4
    case "a":
        var v5 bool = ((2 < 0) || (-631 <= 2147483647))
        _ = v5
    default:
        var v6 int8 = -36
        _ = v6
    }
    v7 := GenS1{F0: 1000, F1: 127, F2: 1}
    var v8 uint8 = 0
    var v9 int64 = 9223372036854775807
    _ = v7
    _ = v8
    _ = v9
    return 527
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{1}
    v3 := []int{4294967295, 42, 1}
    var v10 int32 = -749
    for fi11 := 0; fi11 < 17; fi11++ {
        v10 += (v10 | v10)
        var v12 uint16 = 32767
        _ = fi11
        _ = v12
    }
    v10 -= -2147483648
    if true {
        var v13 int16 = 32767
        fmt.Println(v10, v13)
        _ = v13
    }
    switch -1 {
    case 128:
        _ = len(v2)
    }
    v10 += v10
    if false {
        fmt.Println(v10)
        v10 = 2147483647
    }
    for fi14 := 0; fi14 < 9; fi14++ {
        v10 = -528
        _ = fi14
    }
    for fi15 := 0; fi15 < 10; fi15++ {
        for fi16 := 0; fi16 < 19; fi16++ {
            if (!false) {
                var v17 string = ""
                v17 = (v17 + v17)
                _ = v17
            }
            fmt.Println(fi16, fi15, fi16)
            _ = fi16
        }
        if false {
            for fi18 := 0; fi18 < 1; fi18++ {
                var v19 uint64 = 783
                _ = fi18
                _ = v19
            }
            var v20 bool = (true && true)
            var v21 uint16 = 42
            _ = v20
            _ = v21
        }
        _ = fi15
    }
    if (0 == 1) {
        var v22 int64 = 9223372036854775807
        var v23 float32 = float32(0.1)
        _ = v22
        _ = v23
    } else {
        v24 := GenS2{F0: 127, F1: 0}
        _ = v24
    }
    var v25 uint16 = 0
    switch ("" + "hello") {
    case "hello":
        v10 += v10
    default:
        v10 -= v10
    }
    v25 -= (v25 ^ (v25 * 94))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v25
    _ = v3
}
