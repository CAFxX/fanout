package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 int; F1 uint64; F2 int8 }

func main() {
    v1 := make(map[int]int)
    var v2 int64 = 9223372036854775807
    var v3 uint = 2
    if (0 != 4294967295) {
        var v4 uint = (v3 - v3)
        v3 = 0
        _ = v4
    }
    if (1 != 42) {
        var v5 uint64 = 2
        _ = v5
    }
    var v6 uint = (v3 & v3)
    v6 = 299
    fmt.Println(v3, v6)
    switch 864 {
    case 42:
        var v7 int8 = 0
        _ = v7
    case -9223372036854775808:
        v2 += v2
    default:
        v8 := GenS1{F0: 4294967295, F1: "hello"}
        _ = v8
    }
    for fi9 := 0; fi9 < 19; fi9++ {
        v3 += (v3 >> ((42) & 63))
        for fi10 := 0; fi10 < 9; fi10++ {
            var v11 uint8 = 10
            var v12 bool = (false && (!true))
            v13 := v1[fi9]
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
        for fi14 := 0; fi14 < 18; fi14++ {
            var v15 uint16 = 255
            var v16 uint32 = 1
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi9
    }
    for fi17 := 0; fi17 < 10; fi17++ {
        v3 = (^(v3))
        v2 -= -1
        if (0.0 < float32((3.14159 * -2.25))) {
            var v18 float64 = -2.25
            _ = v18
        }
        _ = fi17
    }
    switch "hello" {
    case "\n\t":
        v2 -= (v2 | v2)
    default:
        if (!false) {
            var v19 uint = v6
            _ = v19
        }
    }
    switch ("\n\t" + "a") {
    case "\n\t":
        v20 := GenS1{F0: 2, F1: "a"}
        _ = v20
    case "":
        var v21 uint8 = 1
        _ = v21
    case "foo bar":
        v3 += (v6 | v3)
    default:
        for fi22 := 0; fi22 < 18; fi22++ {
            for fi23 := 0; fi23 < 18; fi23++ {
                var v24 uint32 = 65535
                _ = fi23
                _ = v24
            }
            _ = fi22
        }
    }
    v25 := []int{42}
    var v26 uint8 = 181
    v27 := GenS1{F0: 0, F1: "hello"}
    v28 := GenS2{F0: 702, F1: 1, F2: 124}
    v29 := GenS2{F0: 32768, F1: 47, F2: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v6
}
