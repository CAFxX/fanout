package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 uint16; F1 int32; F2 uint16 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = -932
    v3 := make(map[int]int)
    var v4 int8 = -7
    var v5 float64 = (((0.0 * (((1e10 / 1.0) - (0.0 / (-2.25 - 0.0))) - float64(1.5))) - ((1.5 / -2.25) - -2.25)) + float64(1e10))
    v6 := []string{((((("\n\t" + "a") + "a") + "hello") + "hello") + "hello"), "foo bar", "\n\t", "foo bar"}
    v5 = 1.5
    v6 = append(v6, "a")
    var v7 string = "\n\t"
    if true {
        v8 := v1[983]
        if false {
            if len(v6) > 0 {
                v6[0] = v7
            }
            var v9 bool = true
            fmt.Println(v8, v4, v9)
            _ = v9
        }
        for fi10 := 0; fi10 < 17; fi10++ {
            v8 -= v8
            fmt.Println(v5)
            delete(v1, v8)
            _ = fi10
        }
        _ = v8
    } else {
        v11 := GenS1{F0: 18446744073709551615, F1: v7}
        _ = v11
    }
    for fi12 := 0; fi12 < 16; fi12++ {
        var v13 uint32 = 42
        _ = fi12
        _ = v13
    }
    if (!(94 <= 4)) {
        var v14 string = ((v7 + "") + "hello")
        _ = v14
    } else {
        var v15 string = "a"
        _ = v15
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        fmt.Println(v5)
        _ = fi16
    }
    v17 := GenS2{F0: 2, F1: v2, F2: 0}
    switch ("foo bar" + v7) {
    case "hello":
        var v18 bool = (false || false)
        _ = v18
    case "":
        v4 += v4
    }
    if true {
        _ = len(v1)
    }
    v19 := GenS2{F0: 2, F1: v2, F2: 32767}
    v2 -= (v2 - v2)
    fmt.Println(v7, v7)
    for fi20 := 0; fi20 < 17; fi20++ {
        _ = len(v6)
        var v21 bool = true
        _ = fi20
        _ = v21
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v6 = append(v6, v7)
    fmt.Println(v2, v7, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
