package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
func genFn0(p0 bool, p1 uint64) bool {
    var v4 string = "a"
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = v4
    return p0
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int8 = 100
    v3 := []string{"", "a", ""}
    var v5 string = ("foo bar" + ("hello" + "\n\t"))
    v6 := GenS1{F0: 32767, F1: ((0.0 > 0.0) || true)}
    var v7 int16 = -32768
    if (198 > 133) {
        v7 -= (v7 | v7)
    }
    v2 -= (v2 | (v2 << ((65535) & 63)))
    for fi8 := 0; fi8 < 20; fi8++ {
        fmt.Println(fi8, fi8)
        v5 = v5
        _ = fi8
    }
    if false {
        delete(v1, v5)
    } else {
        for fi9 := 0; fi9 < 1; fi9++ {
            v2 += 127
            v5 = v5
            v5 = v5
            _ = fi9
        }
        v2 -= v2
    }
    switch v5 {
    case "":
        var v10 uint32 = 1
        _ = v10
    case "a":
        v2 = 90
    case "foo bar":
        v11 := GenS1{F0: (v7 ^ -1), F1: true}
        _ = v11
    default:
        switch 42 {
        case -9223372036854775808:
            v2 -= (v2 * (v2 % ((v2) | 1)))
        case 966:
            for fi12 := 0; fi12 < 11; fi12++ {
                for fi13 := 0; fi13 < 12; fi13++ {
                    var v14 uint64 = 4294967295
                    v14 += (v14 / ((37) | 1))
                    _ = fi13
                    _ = v14
                }
                _ = fi12
            }
        case -1:
            fmt.Println(v5, v7)
        }
    }
    fmt.Println(v7, v7, v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
}
