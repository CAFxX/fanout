package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 int8 }
func genFn0(p0 string, p1 int16) int8 {
    p1 += (p1 << ((4294967295) & 63))
    p1 -= (p1 & p1)
    v4 := []int{32767}
    p1 = (p1 ^ p1)
    var v5 int = 127
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 117
}

func main() {
    v1 := make(map[int]int)
    v2 := []int{2147483647}
    var v3 uint = 116
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 uint64 = 303
        _ = fi6
        _ = v7
    }
    switch ("a" + "\n\t") {
    case "\n\t":
        for fi8 := 0; fi8 < 7; fi8++ {
            var v9 uint32 = 65535
            var v10 uint16 = 0
            var v11 int64 = 1
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
    case "foo bar":
        v12 := GenS1{F0: false, F1: 2, F2: 127}
        _ = v12
    case "":
        fmt.Println(v3)
    default:
        switch -2147483648 {
        case -1:
            v3 += v3
        case 32767:
            for fi13 := 0; fi13 < 20; fi13++ {
                fmt.Println(v3, fi13)
                v3 += v3
                _ = fi13
            }
        }
    }
    v3 -= 651
    switch 85 {
    case 2147483647:
        var v14 int = 164
        _ = v14
    }
    switch ("hello" + "\n\t") {
    case "hello":
        v1[4294967295] = 859
    }
    var v15 int = 128
    v2 = append(v2, (-(v15)))
    v16 := GenS1{F0: ((3.14159 - float32(1e10)) >= 0.0), F1: 2, F2: genFn0(("" + "foo bar"), 0)}
    _ = len(v2)
    v17 := GenS1{F0: ((1e10 + -2.25) != 0.0), F1: 883, F2: -97}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
