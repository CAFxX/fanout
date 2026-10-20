package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 string }

func main() {
    v1 := make(map[string]string)
    v2 := []uint32{488, 463, 2}
    var v3 string = ("foo bar" + "\n\t")
    var v4 int64 = 9223372036854775807
    var v5 uint = 2
    for fi6 := 0; fi6 < 14; fi6++ {
        v4 = v4
        if (false && true) {
            v7 := v1[v3]
            delete(v1, "\n\t")
            _ = v7
        } else {
            for fi8 := 0; fi8 < 18; fi8++ {
                v4 = (v4 % (((v4 + (v4 >> (((v5 << (((v5 ^ v5)) & 63))) & 63)))) | 1))
                _ = fi8
            }
            var v9 float64 = float64((((0.1 / (-2.25 - 1e10)) / ((float32(1.5) / (3.14159 * (1.5 / (0.1 * -2.25)))) / 1.0)) - 1.5))
            v10 := GenS1{F0: 42, F1: v3, F2: (v3 + (v3 + ("hello" + v3)))}
            _ = v10
            _ = v9
        }
        var v11 int16 = 32767
        _ = fi6
        _ = v11
    }
    var v12 bool = true
    var v13 uint32 = 42
    fmt.Println(v3, v3, v13)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if v12 {
        v2 = append(v2, 42)
    } else {
        v4 = v4
        fmt.Println(v5, v12)
    }
    v4 = v4
    for fi14 := 0; fi14 < 13; fi14++ {
        v5 -= (^(v5))
        var v15 float64 = ((0.1 * 0.1) - ((1e10 + float64(1e10)) + -2.25))
        for fi16 := 0; fi16 < 18; fi16++ {
            v13 = v13
            var v17 int64 = v4
            switch (v3 + "") {
            case "a":
                v12 = false
                var v18 uint64 = 18446744073709551615
                _ = v18
            case "\n\t":
                v15 -= v15
            }
            _ = len(v2)
            v4 = v17
            if (("a" + "hello") > "") {
                _ = len(v2)
                var v19 int8 = 0
                var v20 uint32 = v13
                _ = v19
                _ = v20
            }
            _ = fi16
            _ = v17
        }
        _ = fi14
        _ = v15
    }
    v21 := v1[((v3 + (v3 + v3)) + v3)]
    v12 = ((!(!((-59 != 100) && true))) || (127 != 2147483647))
    v21 = v3
    _ = len(v2)
    v3 = ("\n\t" + ("foo bar" + v3))
    fmt.Println(v13)
    v22 := GenS1{F0: v5, F1: ("foo bar" + ("" + ((v3 + ("foo bar" + v21)) + v21))), F2: v3}
    delete(v1, ("\n\t" + v3))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
