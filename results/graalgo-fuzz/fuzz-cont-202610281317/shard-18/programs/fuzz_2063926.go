package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 int64; F1 int8 }

func main() {
    v1 := []string{("foo bar" + "hello"), ""}
    var v2 uint = 0
    v3 := make(map[string]bool)
    v1 = append(v1, "foo bar")
    var v4 int64 = int64(v2)
    v4 -= (^(v4))
    var v5 uint64 = 1
    v6 := GenS1{F0: -1, F1: (v2 - v2)}
    if ((1.5 >= 1.5) && (uint(v5) >= v2)) {
        var v7 int = -1
        for fi8 := 0; fi8 < 17; fi8++ {
            for fi9 := 0; fi9 < 8; fi9++ {
                switch fi9 {
                case 311:
                    _ = len(v3)
                case 128:
                    var v10 int = (fi8 | (fi8 | fi9))
                    _ = v10
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi9
            }
            _ = fi8
        }
        fmt.Println(v5)
        _ = v7
    }
    var v11 uint32 = 65535
    for fi12 := 0; fi12 < 3; fi12++ {
        switch (fi12 * fi12) {
        case 9223372036854775807:
            var v13 uint = v2
            _ = v13
        case 4294967295:
            var v14 string = "hello"
            _ = v14
        case 32768:
            fmt.Println(v11)
        }
        v15 := v3["hello"]
        _ = fi12
        _ = v15
    }
    fmt.Println(v4)
    var v16 int8 = 100
    fmt.Println(v5, v5)
    v17 := GenS2{F0: v4, F1: 42}
    v3["\n\t"] = (false && (!(-1 >= -272)))
    var v18 float32 = float32((1.5 + 1.5))
    var v19 uint16 = 978
    v16 += 100
    v5 = (v5 & 18446744073709551615)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
