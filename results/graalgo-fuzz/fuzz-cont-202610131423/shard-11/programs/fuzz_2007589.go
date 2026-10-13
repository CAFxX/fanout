package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }

func main() {
    v1 := []uint32{423}
    v2 := []uint32{446, 4294967295, 807}
    var v3 float32 = 0.1
    v4 := []string{}
    v5 := make(map[int]bool)
    var v6 int8 = 0
    switch 127 {
    case 32768:
        if true {
            var v7 uint16 = 32767
            _ = v7
        }
    case 2147483647:
        for fi8 := 0; fi8 < 15; fi8++ {
            for fi9 := 0; fi9 < 20; fi9++ {
                v10 := GenS1{F0: 678, F1: (float64(-2.25) < (float64(-2.25) / 3.14159))}
                var v11 int64 = 9223372036854775807
                fmt.Println(v3, fi8)
                _ = fi9
                _ = v10
                _ = v11
            }
            v6 = v6
            _ = fi8
        }
    case 4294967295:
        v3 = float32(((0.1 + 1.5) - 1e10))
    default:
        var v12 uint16 = 255
        _ = v12
    }
    _ = len(v1)
    _ = len(v2)
    var v13 int16 = 49
    for fi14 := 0; fi14 < 6; fi14++ {
        var v15 string = ("a" + "foo bar")
        v13 = 42
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 16; fi16++ {
        v17 := GenS1{F0: 1, F1: false}
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi16
        _ = v17
    }
    fmt.Println(v3)
    _ = len(v5)
    var v18 string = "\n\t"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
