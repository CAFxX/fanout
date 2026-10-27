package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = ((float64(0.1) + (-2.25 / 0.1)) * (1e10 + 0.0))
    v3 := []string{""}
    var v4 int32 = -262
    var v5 int32 = 0
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 int32 = (v5 | v5)
        v8 := GenS1{F0: 583, F1: 1}
        _ = fi6
        _ = v7
        _ = v8
    }
    switch -953 {
    case 127:
        fmt.Println(v4, v4)
    case 32767:
        _ = len(v1)
    }
    if (827 >= 127) {
        v5 += (^(v4))
        v2 = 0.1
    } else {
        fmt.Println(v2, v2)
        for fi9 := 0; fi9 < 5; fi9++ {
            v4 = -2147483648
            v2 += v2
            v2 = v2
            _ = fi9
        }
    }
    if true {
        var v10 int32 = (v4 % ((954) | 1))
        _ = v10
    }
    v4 = (v4 >> ((42) & 63))
    v5 -= (v5 | v4)
    for fi11 := 0; fi11 < 20; fi11++ {
        v12 := GenS1{F0: 1, F1: 1}
        var v13 bool = (false || (2 == 2))
        _ = fi11
        _ = v12
        _ = v13
    }
    v3 = append(v3, "foo bar")
    if ((("" + "") + "foo bar") < "foo bar") {
        if true {
            var v14 uint = 4294967295
            var v15 int32 = v5
            v4 -= 512
            _ = v14
            _ = v15
        }
    } else {
        v1[42] = (1 != 2)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
