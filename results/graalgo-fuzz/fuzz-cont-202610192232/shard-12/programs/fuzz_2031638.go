package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 string }
func genFn0(p0 string, p1 string) uint32 {
    fmt.Println(p1)
    for fi7 := 0; fi7 < 20; fi7++ {
        for fi8 := 0; fi8 < 13; fi8++ {
            v9 := GenS1{F0: 711, F1: 1, F2: p1}
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    return 42
}
func genFn1() int64 {
    var v10 int8 = 100
    v10 = 1
    switch -1 {
    case 734:
        var v11 int8 = v10
        _ = v11
    case 2147483647:
        v12 := GenS1{F0: 65535, F1: 1, F2: ("foo bar" + "foo bar")}
        _ = v12
    default:
        switch 127 {
        case 32768:
            var v13 int64 = 690
            _ = v13
        case -429:
            v10 -= (v10 >> ((1) & 63))
        }
    }
    switch -700 {
    case 128:
        var v14 int8 = (v10 << ((862) & 63))
        _ = v14
    case 1:
        if true {
            v10 = (v10 ^ -97)
        } else {
            if false {
                v10 += (v10 << ((0) & 63))
                var v15 int8 = v10
                fmt.Println(v10)
                _ = v15
            } else {
                v10 -= (-(v10))
                fmt.Println(v10)
            }
            fmt.Println(v10)
        }
    }
    _ = v10
    return -1
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]string)
    var v3 int64 = -279
    var v4 uint16 = 255
    v5 := make(map[int]bool)
    var v6 uint = 0
    for fi16 := 0; fi16 < 11; fi16++ {
        var v17 int64 = v3
        _ = fi16
        _ = v17
    }
    v18 := v1[0]
    switch "\n\t" {
    case "a":
        var v19 uint64 = 42
        _ = v19
    case "\n\t":
        var v20 uint64 = 18446744073709551615
        _ = v20
    case "foo bar":
        switch "" {
        case "foo bar":
            var v21 float32 = (1e10 - ((0.0 * (-2.25 * (-2.25 - (((3.14159 / 3.14159) / -2.25) * -2.25)))) + 1.5))
            _ = v21
        }
    }
    for fi22 := 0; fi22 < 18; fi22++ {
        _ = len(v1)
        var v23 uint = 2
        _ = fi22
        _ = v23
    }
    fmt.Println(v6, v3)
    for fi24 := 0; fi24 < 9; fi24++ {
        v18 += v18
        delete(v5, -9223372036854775808)
        _ = fi24
    }
    fmt.Println(v4, v18)
    v3 = (-(v3))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
