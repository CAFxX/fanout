package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }

func main() {
    v1 := make(map[string]bool)
    v2 := []int{1, -1}
    var v3 uint8 = 198
    v4 := make(map[int]uint32)
    v5 := make(map[string]int)
    fmt.Println(v3)
    var v6 uint8 = 158
    v1["a"] = (!true)
    var v7 uint16 = 65535
    if (-2.25 != 1e10) {
        v8 := GenS1{F0: "foo bar", F1: (275 != 124)}
        _ = v8
    }
    switch "a" {
    case "foo bar":
        for fi9 := 0; fi9 < 17; fi9++ {
            v10 := GenS1{F0: (("" + "foo bar") + "a"), F1: (1000 != 0)}
            for fi11 := 0; fi11 < 17; fi11++ {
                if len(v2) > 0 {
                    v2[0] = 32768
                }
                _ = fi11
            }
            var v12 string = ("\n\t" + ("\n\t" + ("foo bar" + "\n\t")))
            _ = fi9
            _ = v10
            _ = v12
        }
    case "\n\t":
        if (!true) {
            v13 := GenS1{F0: "", F1: (932 > 497)}
            v14 := GenS1{F0: ((("\n\t" + (("foo bar" + "a") + (("a" + "") + "\n\t"))) + "foo bar") + "\n\t"), F1: false}
            v15 := v4[391]
            _ = v13
            _ = v14
            _ = v15
        }
    default:
        if ((((-93 <= 42) && false) || (!true)) || (9223372036854775807 == 172)) {
            if (0 >= 0) {
                var v16 string = "hello"
                _ = v16
            }
        }
    }
    for fi17 := 0; fi17 < 16; fi17++ {
        var v18 uint16 = v7
        _ = fi17
        _ = v18
    }
    fmt.Println(v7)
    var v19 uint64 = 562
    v20 := GenS1{F0: "\n\t", F1: ((v7 - v7) <= (v7 + v7))}
    v21 := GenS1{F0: "", F1: (!(114 > 42))}
    for fi22 := 0; fi22 < 1; fi22++ {
        var v23 float64 = (-2.25 + 3.14159)
        for fi24 := 0; fi24 < 13; fi24++ {
            if len(v2) > 0 {
                v2[0] = (^(fi24))
            }
            _ = fi24
        }
        var v25 int = (fi22 & (fi22 * fi22))
        _ = fi22
        _ = v23
        _ = v25
    }
    v1[("a" + (("foo bar" + "") + ((("\n\t" + "a") + "hello") + "\n\t")))] = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
