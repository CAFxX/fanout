package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 int }
type GenS2 struct { F0 uint64; F1 uint64; F2 int16 }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 uint16 = 865
    v2 := []string{((("\n\t" + "foo bar") + "\n\t") + (("" + "\n\t") + "a"))}
    v3 := make(map[string]int)
    v4 := make(map[int]uint32)
    v5 := make(map[string]uint32)
    var v6 uint64 = 4294967295
    fmt.Println(v1, v1)
    var v7 string = ((("hello" + ("foo bar" + ("a" + ("foo bar" + "foo bar")))) + "\n\t") + ("foo bar" + "hello"))
    v2 = append(v2, "foo bar")
    switch v7 {
    case "hello":
        for fi8 := 0; fi8 < 7; fi8++ {
            v4[-9223372036854775808] = 2
            v9 := GenS1{F0: -1, F1: -128, F2: (fi8 | (fi8 / ((9223372036854775807) | 1)))}
            _ = fi8
            _ = v9
        }
    case "foo bar":
        v10 := v5[v7]
        _ = v10
    default:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    if (0 != 42) {
        v6 -= 18446744073709551615
        var v11 uint32 = 940
        v12 := v5[""]
        _ = v11
        _ = v12
    }
    if (!false) {
        v6 -= v6
        if false {
            if true {
                var v13 uint8 = 107
                fmt.Println(v6, v13)
                _ = v13
            } else {
                v6 -= 0
                _ = len(v4)
            }
            for fi14 := 0; fi14 < 9; fi14++ {
                v15 := GenS2{F0: (v6 / ((496) | 1)), F1: v6, F2: -357}
                _ = fi14
                _ = v15
            }
        }
    } else {
        if (106 <= 42) {
            v7 = "foo bar"
            var v16 uint = 837
            _ = v16
        }
        var v17 int64 = 1
        if len(v2) > 0 {
            v2[0] = v7
        }
        _ = v17
    }
    v18 := GenS3{F0: v6}
    v6 += v6
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
