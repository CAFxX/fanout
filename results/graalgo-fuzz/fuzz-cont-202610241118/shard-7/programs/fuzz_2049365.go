package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 bool; F1 uint8 }

func main() {
    v1 := make(map[string]bool)
    v2 := []uint32{42, 2, 891, 1}
    v3 := []string{"", "a"}
    var v4 int64 = -1
    v5 := []string{"hello", (("foo bar" + "") + ""), "foo bar", ("foo bar" + "\n\t")}
    v6 := []uint32{329, 437}
    for fi7 := 0; fi7 < 10; fi7++ {
        for fi8 := 0; fi8 < 2; fi8++ {
            fmt.Println(fi7, fi7)
            _ = fi8
        }
        _ = fi7
    }
    switch ("" + "foo bar") {
    case "\n\t":
        v9 := GenS1{F0: ((0 != 1) && true), F1: 140}
        _ = v9
    case "a":
        v6 = append(v6, 2)
    default:
        v10 := GenS1{F0: true, F1: 42}
        _ = v10
    }
    v4 -= (v4 / (((^(v4))) | 1))
    fmt.Println(v4)
    v3 = append(v3, "a")
    v11 := GenS2{F0: false, F1: 148}
    for fi12 := 0; fi12 < 7; fi12++ {
        delete(v1, "foo bar")
        _ = fi12
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        var v14 int8 = 0
        if (32767 != 667) {
            var v15 float32 = -2.25
            switch ("hello" + "a") {
            case "a":
                _ = len(v1)
            case "":
                v4 -= v4
            default:
                v16 := v1[((((("" + "\n\t") + "foo bar") + "a") + ("\n\t" + ("\n\t" + ""))) + "a")]
                _ = v16
            }
            _ = v15
        }
        v4 -= v4
        _ = fi13
        _ = v14
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
