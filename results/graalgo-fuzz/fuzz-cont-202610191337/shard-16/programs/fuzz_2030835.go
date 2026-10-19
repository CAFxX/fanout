package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 uint64 }
func genFn0(p0 uint8) int8 {
    p0 = 142
    var v7 uint = 21
    _ = p0
    _ = v7
    return 8
}

func main() {
    var v1 uint8 = 132
    var v2 float64 = 1.5
    var v3 uint32 = 303
    v4 := make(map[int]bool)
    var v5 uint8 = v1
    v6 := make(map[uint32]bool)
    for fi8 := 0; fi8 < 17; fi8++ {
        v5 = v1
        v4[fi8] = true
        _ = fi8
    }
    v2 += v2
    for fi9 := 0; fi9 < 9; fi9++ {
        v3 = v3
        _ = fi9
    }
    v10 := GenS1{F0: 127, F1: false, F2: 27}
    if false {
        var v11 int64 = 1
        _ = v11
    }
    var v12 int8 = 42
    fmt.Println(v3, v1, v1)
    var v13 uint8 = v5
    v1 = 3
    var v14 float64 = v2
    v14 -= (float64(0.0) / v2)
    for fi15 := 0; fi15 < 2; fi15++ {
        if (0 < 500) {
            switch ("\n\t" + "a") {
            case "hello":
                var v16 float32 = 0.1
                _ = v16
            case "":
                v13 -= 36
            case "\n\t":
                v17 := []int{}
                _ = v17
            }
            fmt.Println(v3)
        }
        for fi18 := 0; fi18 < 11; fi18++ {
            v19 := v4[fi15]
            _ = fi18
            _ = v19
        }
        _ = fi15
    }
    var v20 int32 = 2147483647
    switch (("foo bar" + "") + "foo bar") {
    case "hello":
        v14 = (1e10 * float64(1e10))
    case "":
        var v21 uint = 4294967295
        _ = v21
    }
    switch 2147483647 {
    case -537:
        v1 = (v1 | (v1 / ((v1) | 1)))
    case -2147483648:
        fmt.Println(v13, v13, v14)
    case 32768:
        for fi22 := 0; fi22 < 20; fi22++ {
            var v23 uint32 = v3
            _ = fi22
            _ = v23
        }
    default:
        fmt.Println(v2)
    }
    if false {
        fmt.Println(v1)
        switch -436 {
        case -2147483648:
            for fi24 := 0; fi24 < 12; fi24++ {
                fmt.Println(v14)
                _ = fi24
            }
        default:
            v12 = 1
        }
    }
    fmt.Println(v12, v1, v14)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
