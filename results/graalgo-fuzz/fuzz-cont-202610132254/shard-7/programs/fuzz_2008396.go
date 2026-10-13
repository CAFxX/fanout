package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint }
type GenS2 struct { F0 int16; F1 uint16; F2 string }
func genFn0(p0 uint) int64 {
    var v4 uint16 = 65535
    p0 -= p0
    for fi5 := 0; fi5 < 20; fi5++ {
        var v6 int16 = 42
        for fi7 := 0; fi7 < 8; fi7++ {
            var v8 int32 = 0
            v9 := GenS1{F0: "", F1: false, F2: 435}
            v10 := []int{(fi7 / (((^(fi7))) | 1)), (fi7 ^ (^(fi5)))}
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
        var v11 int16 = (v6 + (v6 + v6))
        _ = fi5
        _ = v11
        _ = v6
    }
    switch (("\n\t" + "hello") + "\n\t") {
    case "a":
        if true {
            fmt.Println(p0)
            if true {
                p0 = (p0 ^ p0)
                var v12 uint = (p0 & 65535)
                v13 := GenS1{F0: ("" + "foo bar"), F1: true, F2: (v12 % ((550) | 1))}
                _ = v12
                _ = v13
            }
        } else {
            var v14 uint64 = 4294967295
            var v15 int16 = 42
            _ = v14
            _ = v15
        }
    case "hello":
        var v16 uint32 = 856
        _ = v16
    case "foo bar":
        var v17 uint32 = 244
        _ = v17
    }
    _ = p0
    _ = v4
    return -9223372036854775808
}

func main() {
    var v1 uint = 365
    v2 := []uint32{}
    v3 := make(map[int]int)
    v18 := GenS1{F0: ("a" + "foo bar"), F1: (!false), F2: v1}
    for fi19 := 0; fi19 < 9; fi19++ {
        v1 = (v1 * (v1 | (^(v1))))
        v20 := GenS1{F0: "a", F1: true, F2: (v1 >> (((^(v1))) & 63))}
        _ = fi19
        _ = v20
    }
    v2 = append(v2, 260)
    if len(v2) > 0 {
        v2[0] = 42
    }
    v21 := GenS2{F0: 1000, F1: 150, F2: "\n\t"}
    v22 := GenS2{F0: -353, F1: 42, F2: "\n\t"}
    switch -1 {
    case 0:
        v1 = v1
    default:
        v1 = (^(v1))
    }
    var v23 int16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
