package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 uint8 }

func main() {
    v1 := []string{("foo bar" + "\n\t")}
    var v2 float32 = -2.25
    var v3 int = 4294967295
    v4 := make(map[string]bool)
    var v5 bool = ((((((0.1 - float64((3.14159 - (v2 * (1e10 - v2))))) - ((0.0 * -2.25) * 0.1)) - 0.0) - 0.1) != float64(0.1)) || true)
    fmt.Println(v2, v5, v2)
    v1 = append(v1, ("" + ("foo bar" + "")))
    fmt.Println(v5)
    for fi6 := 0; fi6 < 14; fi6++ {
        v3 -= (fi6 + (fi6 ^ 42))
        v5 = (0 > -2147483648)
        v4["foo bar"] = v5
        _ = fi6
    }
    if len(v1) > 0 {
        v1[0] = "a"
    }
    var v7 int = (v3 * v3)
    if (3.14159 != ((3.14159 * -2.25) * 1.5)) {
        fmt.Println(v3, v2)
    } else {
        switch v7 {
        case 9223372036854775807:
            var v8 int = v7
            _ = v8
        }
        for fi9 := 0; fi9 < 16; fi9++ {
            switch fi9 {
            case 32767:
                v10 := GenS1{F0: -128, F1: 42, F2: 196}
                _ = v10
            }
            v7 -= -1
            if v5 {
                v2 -= -2.25
            }
            _ = fi9
        }
    }
    if true {
        var v11 float32 = v2
        v4[""] = true
        _ = v11
    } else {
        var v12 int64 = -1
        _ = v12
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
