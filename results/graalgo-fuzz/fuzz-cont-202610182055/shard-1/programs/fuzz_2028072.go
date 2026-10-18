package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 int8; F1 uint32; F2 int64 }
func genFn0(p0 uint8) uint8 {
    v7 := []int{42, -3, 127}
    switch "hello" {
    case "foo bar":
        var v8 int32 = 1
        _ = v8
    case "hello":
        if (2 <= 2147483648) {
            p0 = (p0 << ((1) & 63))
        }
    }
    switch 1 {
    case 32768:
        p0 -= p0
    case 128:
        p0 -= (p0 - 255)
    case 127:
        if len(v7) > 0 {
            v7[0] = -1
        }
    default:
        v9 := GenS1{F0: -109, F1: -578}
        _ = v9
    }
    switch -9223372036854775808 {
    case 0:
        for fi10 := 0; fi10 < 12; fi10++ {
            if false {
                var v11 int = (fi10 ^ (fi10 | 0))
                fmt.Println(v11)
                _ = v11
            }
            _ = fi10
        }
    }
    v7 = append(v7, -360)
    _ = p0
    _ = v7
    return p0
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]uint32)
    var v3 int = 0
    v4 := make(map[string]bool)
    v5 := make(map[string]uint32)
    v6 := []string{"hello", ("\n\t" + "\n\t")}
    var v12 int8 = 100
    v12 -= -94
    fmt.Println(v3)
    v12 -= (v12 * -60)
    for fi13 := 0; fi13 < 13; fi13++ {
        var v14 int32 = -2147483648
        var v15 int8 = (v12 >> ((2) & 63))
        switch "hello" {
        case "a":
            v2[("foo bar" + "hello")] = 1
        default:
            v6 = append(v6, "\n\t")
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    v12 -= v12
    switch -699 {
    case -1:
        if len(v6) > 0 {
            v6[0] = ""
        }
    }
    switch "\n\t" {
    case "foo bar":
        for fi16 := 0; fi16 < 20; fi16++ {
            var v17 float32 = 1.5
            switch (v3 % ((32767) | 1)) {
            case 927:
                v18 := GenS1{F0: 0, F1: -449}
                _ = v18
            case 0:
                v3 = fi16
            }
            v3 = -546
            _ = fi16
            _ = v17
        }
    case "a":
        var v19 int8 = (v12 >> ((42) & 63))
        _ = v19
    default:
        v3 -= v3
    }
    if (0.1 < (((1e10 + 0.0) + 1.5) / (1e10 + 0.1))) {
        fmt.Println(v3, v3)
        v12 += v12
        switch (-(v3)) {
        case 2147483647:
            v20 := GenS1{F0: 0, F1: v3}
            _ = v20
        case 586:
            delete(v4, "\n\t")
        }
    }
    for fi21 := 0; fi21 < 8; fi21++ {
        var v22 int64 = 1
        v23 := GenS2{F0: v12, F1: 42, F2: v22}
        fmt.Println(v3, v3)
        _ = fi21
        _ = v22
        _ = v23
    }
    v24 := GenS1{F0: 461, F1: (v3 * v3)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
