package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
func genFn0() uint64 {
    var v6 uint8 = 255
    v6 -= (v6 >> ((1) & 63))
    var v7 uint64 = 691
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 uint64) uint {
    p0 = p0
    fmt.Println(p0)
    p0 = (^(p0))
    _ = p0
    return 18446744073709551615
}

func main() {
    v1 := []int{128}
    var v2 int8 = 100
    v3 := []int{822, 924, 127}
    v4 := make(map[int]uint32)
    v5 := []string{"a", "a", "a"}
    for fi8 := 0; fi8 < 3; fi8++ {
        switch ("a" + "hello") {
        case "foo bar":
            var v9 uint32 = 2
            _ = v9
        case "a":
            v10 := GenS1{F0: 0, F1: 42}
            _ = v10
        case "hello":
            switch "a" {
            case "foo bar":
                v11 := GenS1{F0: 338, F1: 42}
                _ = v11
            case "a":
                v2 = 100
            case "\n\t":
                v2 -= (v2 ^ -37)
            }
        }
        _ = fi8
    }
    v2 = (v2 << ((980) & 63))
    for fi12 := 0; fi12 < 14; fi12++ {
        fmt.Println(v2, v2)
        var v13 uint16 = 504
        v2 = (v2 % ((v2) | 1))
        _ = fi12
        _ = v13
    }
    fmt.Println(v2)
    switch 42 {
    case -9223372036854775808:
        fmt.Println(v2)
    case -1:
        fmt.Println(v2)
    case 4294967295:
        switch -824 {
        case 127:
            v2 -= v2
        default:
            for fi14 := 0; fi14 < 7; fi14++ {
                v15 := GenS1{F0: 515, F1: 0}
                v1 = append(v1, 0)
                fmt.Println(fi14)
                _ = fi14
                _ = v15
            }
        }
    }
    if (true || true) {
        v4[2147483647] = 42
        for fi16 := 0; fi16 < 13; fi16++ {
            v2 = -1
            _ = fi16
        }
        v2 -= (v2 >> ((536) & 63))
    }
    var v17 uint64 = 263
    for fi18 := 0; fi18 < 13; fi18++ {
        var v19 string = (((("\n\t" + "") + "foo bar") + ((("hello" + "\n\t") + "hello") + "\n\t")) + ("" + "\n\t"))
        _ = fi18
        _ = v19
    }
    var v20 float32 = 1.5
    _ = len(v4)
    v1 = append(v1, 1)
    var v21 string = "foo bar"
    var v22 uint32 = 344
    if true {
        if false {
            var v23 uint = 0
            _ = v23
        } else {
            v22 = (v22 * v22)
        }
    } else {
        v22 += v22
        var v24 bool = (128 >= 491)
        v22 += (v22 & (v22 | 985))
        _ = v24
    }
    v25 := GenS1{F0: 0, F1: 82}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
