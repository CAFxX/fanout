package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 int64 }

func main() {
    var v1 float32 = ((((3.14159 - 1.5) + (3.14159 / 1e10)) - -2.25) * (1.5 / 0.1))
    v2 := make(map[int]bool)
    v3 := make(map[uint32]bool)
    switch ("\n\t" + "") {
    case "\n\t":
        var v4 float64 = 0.1
        _ = v4
    case "foo bar":
        var v5 uint = 4294967295
        _ = v5
    case "a":
        for fi6 := 0; fi6 < 17; fi6++ {
            var v7 int64 = -1
            _ = fi6
            _ = v7
        }
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 float32 = v1
        _ = fi8
        _ = v9
    }
    _ = len(v2)
    fmt.Println(v1)
    v10 := []int{-1, -361, -1, 679}
    if (!true) {
        if (-49 >= 120) {
            for fi11 := 0; fi11 < 4; fi11++ {
                v1 += ((v1 + 1e10) * v1)
                v12 := v2[fi11]
                _ = fi11
                _ = v12
            }
            v13 := GenS1{F0: 253, F1: "hello", F2: -1}
            _ = v13
        }
        var v14 int = 362
        _ = v14
    }
    switch 1 {
    case -2147483648:
        var v15 int = 32768
        _ = v15
    case 409:
        v1 = ((float32(3.14159) * v1) - v1)
    }
    v2[9223372036854775807] = false
    for fi16 := 0; fi16 < 12; fi16++ {
        for fi17 := 0; fi17 < 12; fi17++ {
            var v18 uint32 = 1
            _ = fi17
            _ = v18
        }
        v1 += float32(-2.25)
        _ = fi16
    }
    var v19 int32 = 2147483647
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    for fi20 := 0; fi20 < 5; fi20++ {
        v21 := GenS1{F0: 252, F1: ("\n\t" + "foo bar"), F2: 650}
        _ = fi20
        _ = v21
    }
    v22 := GenS1{F0: 524, F1: ("a" + ""), F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v22
    _ = v3
}
