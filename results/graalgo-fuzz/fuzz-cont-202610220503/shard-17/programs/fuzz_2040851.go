package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 int8 }
type GenS2 struct { F0 int64; F1 int8; F2 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8; F1 uint32; F2 string }
func genFn0() int64 {
    var v6 int = 127
    if true {
        v6 = v6
    }
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 uint8 = 255
        _ = fi7
        _ = v8
    }
    _ = v6
    return -9223372036854775808
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]bool)
    var v3 float64 = ((0.1 / -2.25) - 1e10)
    var v4 int64 = -440
    var v5 float64 = v3
    switch ("hello" + "a") {
    case "\n\t":
        var v9 int32 = -1
        _ = v9
    case "hello":
        var v10 int32 = -1
        _ = v10
    default:
        v11 := GenS1{F0: v4, F1: 5, F2: 127}
        _ = v11
    }
    v12 := []int{-9223372036854775808, -9223372036854775808}
    for fi13 := 0; fi13 < 11; fi13++ {
        _ = len(v1)
        fmt.Println(v5, v3, fi13)
        var v14 uint64 = 1
        _ = fi13
        _ = v14
    }
    switch "a" {
    case "a":
        var v15 uint = 2
        _ = v15
    default:
        for fi16 := 0; fi16 < 10; fi16++ {
            if (0 != fi16) {
                fmt.Println(v5, v3, v5)
            }
            _ = fi16
        }
    }
    v17 := GenS2{F0: 9223372036854775807, F1: 127, F2: 255}
    if (1000 != int16(v4)) {
        fmt.Println(v4, v5)
    } else {
        v18 := GenS3{F0: 2}
        _ = v18
    }
    var v19 int64 = (v4 | genFn0())
    v1["hello"] = 2
    if false {
        var v20 uint32 = 0
        for fi21 := 0; fi21 < 10; fi21++ {
            for fi22 := 0; fi22 < 6; fi22++ {
                var v23 float32 = float32(v5)
                v3 = v5
                v24 := GenS4{F0: 127, F1: (v20 >> ((1) & 63)), F2: ("\n\t" + "a")}
                _ = fi22
                _ = v23
                _ = v24
            }
            _ = fi21
        }
        _ = v20
    }
    for fi25 := 0; fi25 < 20; fi25++ {
        var v26 string = ((("hello" + "foo bar") + "hello") + "hello")
        var v27 uint16 = 1
        _ = fi25
        _ = v26
        _ = v27
    }
    var v28 uint8 = 127
    v19 = 514
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v17
    _ = v19
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
