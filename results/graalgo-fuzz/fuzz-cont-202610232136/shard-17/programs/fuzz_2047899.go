package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint64 }
type GenS2 struct { F0 string; F1 uint32 }
func genFn0() int16 {
    var v5 uint = 1
    var v6 int = 128
    var v7 uint16 = 42
    _ = v5
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 int8 = 99
    v2 := make(map[int]bool)
    v3 := make(map[string]int)
    v4 := make(map[string]bool)
    var v8 int = 32767
    var v9 float64 = 0.0
    v10 := GenS1{F0: 21, F1: -876, F2: 980}
    for fi11 := 0; fi11 < 1; fi11++ {
        v12 := GenS1{F0: -2147483648, F1: 42, F2: 1}
        _ = fi11
        _ = v12
    }
    var v13 uint8 = 128
    for fi14 := 0; fi14 < 2; fi14++ {
        for fi15 := 0; fi15 < 14; fi15++ {
            switch "" {
            case "\n\t":
                v1 += (v1 << ((751) & 63))
            case "hello":
                fmt.Println(v9, v13, v8)
            case "a":
                v9 -= 0.0
            default:
                var v16 string = "hello"
                _ = v16
            }
            _ = fi15
        }
        _ = fi14
    }
    v13 = v13
    v17 := []int{(-(v8)), v8, int(v13), (v8 << ((316) & 63))}
    if len(v17) > 0 {
        v17[0] = v8
    }
    switch "foo bar" {
    case "foo bar":
        delete(v3, "foo bar")
    }
    for fi18 := 0; fi18 < 17; fi18++ {
        var v19 int64 = -239
        v2[(-(fi18))] = false
        _ = fi18
        _ = v19
    }
    var v20 float32 = 1e10
    v21 := GenS2{F0: (("hello" + "hello") + "\n\t"), F1: 2}
    for fi22 := 0; fi22 < 12; fi22++ {
        var v23 int8 = (v1 % (((^(v1))) | 1))
        _ = fi22
        _ = v23
    }
    switch (v8 / ((v8) | 1)) {
    case 1:
        var v24 uint16 = 2
        _ = v24
    case -1:
        v3[("" + "")] = -9223372036854775808
    case 423:
        var v25 int = v8
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
