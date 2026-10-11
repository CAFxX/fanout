package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
func genFn0(p0 uint8, p1 uint8) int16 {
    p0 = (p0 - p0)
    switch 32767 {
    case 32767:
        fmt.Println(p1, p1)
    case 138:
        v5 := []int{-2147483648, 128, 2147483647}
        _ = v5
    }
    if true {
        var v6 float64 = 3.14159
        _ = v6
    }
    var v7 int = 0
    _ = p0
    _ = p1
    _ = v7
    return 853
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{"a", "", "", "\n\t"}
    var v3 uint = 189
    v4 := make(map[uint32]bool)
    v3 += v3
    v3 = 4294967295
    for fi8 := 0; fi8 < 7; fi8++ {
        for fi9 := 0; fi9 < 10; fi9++ {
            v3 = (v3 * v3)
            switch int(v3) {
            case 1:
                var v10 int = fi8
                _ = v10
            }
            switch (-(fi8)) {
            case 32767:
                v3 -= (v3 + 65535)
            case 9223372036854775807:
                var v11 int16 = 623
                _ = v11
            case 128:
                v3 = v3
            }
            v12 := GenS1{F0: 32767, F1: -1}
            _ = fi9
            _ = v12
        }
        _ = fi8
    }
    v3 -= v3
    var v13 int32 = 2147483647
    if (4294967295 != 2147483647) {
        if false {
            var v14 uint = 220
            _ = v14
        }
    }
    if ((v3 * (v3 ^ (v3 - 18446744073709551615))) < (v3 | (v3 >> ((v3) & 63)))) {
        if (true && (!(42 > -40))) {
            var v15 uint32 = 471
            v15 = v15
            _ = v15
        }
    } else {
        if false {
            var v16 int8 = 42
            _ = v16
        } else {
            fmt.Println(v3)
        }
        v3 += v3
    }
    if (0 >= uint8(v3)) {
        var v17 uint8 = 255
        v3 = (v3 + (^(v3)))
        _ = v17
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
