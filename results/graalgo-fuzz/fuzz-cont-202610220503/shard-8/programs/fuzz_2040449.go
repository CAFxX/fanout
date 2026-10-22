package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 int32 }
type GenS2 struct { F0 bool; F1 int16 }

func main() {
    var v1 int8 = -128
    var v2 int8 = v1
    v3 := make(map[int]string)
    var v4 uint = 552
    var v5 uint = v4
    v4 += v4
    var v6 uint8 = 0
    if (!true) {
        var v7 uint16 = 355
        var v8 float32 = 0.1
        v9 := GenS1{F0: 1, F1: "\n\t", F2: -1}
        _ = v7
        _ = v8
        _ = v9
    } else {
        for fi10 := 0; fi10 < 10; fi10++ {
            var v11 int = (-(fi10))
            v1 -= 1
            fmt.Println(v1, v5, v6)
            _ = fi10
            _ = v11
        }
        v2 -= (v2 & (v1 << (((v4 * v4)) & 63)))
        v6 += (v6 / ((255) | 1))
    }
    v12 := []int{9223372036854775807, 0, 4294967295, 998}
    v3[869] = ("foo bar" + ("a" + "foo bar"))
    v13 := GenS1{F0: 528, F1: (("hello" + "\n\t") + "foo bar"), F2: 367}
    v4 -= v5
    if false {
        v14 := GenS2{F0: (3.14159 < (0.0 + (1e10 / (0.1 + ((1e10 * float32((-2.25 * 1e10))) - 3.14159))))), F1: 1000}
        _ = v14
    } else {
        var v15 int = 32767
        _ = v15
    }
    fmt.Println(v1, v1, v5)
    for fi16 := 0; fi16 < 14; fi16++ {
        var v17 uint = (v4 | v4)
        var v18 uint = v4
        v19 := GenS2{F0: false, F1: 160}
        _ = fi16
        _ = v17
        _ = v18
        _ = v19
    }
    var v20 uint = (^(v4))
    v2 -= (v1 ^ (v2 * int8(v6)))
    var v21 uint = v20
    for fi22 := 0; fi22 < 15; fi22++ {
        v4 = 255
        var v23 uint32 = 42
        if (42 > int(v21)) {
            for fi24 := 0; fi24 < 5; fi24++ {
                v2 = 0
                _ = fi24
            }
        }
        _ = fi22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
