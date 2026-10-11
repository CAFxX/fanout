package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 int }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]int)
    var v3 uint32 = 65535
    fmt.Println(v3)
    v3 += (v3 - (v3 + (v3 % (((v3 ^ (v3 << ((597) & 63)))) | 1))))
    v4 := []int{32767, 0}
    var v5 uint16 = 0
    for fi6 := 0; fi6 < 12; fi6++ {
        for fi7 := 0; fi7 < 18; fi7++ {
            fmt.Println(v3, fi6)
            _ = fi7
        }
        if ((^(v5)) == 0) {
            for fi8 := 0; fi8 < 15; fi8++ {
                v5 = v5
                v5 = (v5 ^ (v5 >> ((18446744073709551615) & 63)))
                _ = fi8
            }
            v9 := GenS1{F0: 65535, F1: (fi6 << ((656) & 63)), F2: 4294967295}
            var v10 uint8 = 127
            _ = v10
            _ = v9
        } else {
            _ = len(v1)
            v3 -= v3
            if true {
                v11 := GenS1{F0: 839, F1: (fi6 * fi6), F2: (fi6 / ((fi6) | 1))}
                _ = v11
            }
        }
        _ = fi6
    }
    fmt.Println(v5)
    v5 = 69
    if true {
        fmt.Println(v5)
    } else {
        if true {
            var v12 int16 = 1000
            v3 += (v3 * (v3 ^ 65535))
            switch "foo bar" {
            case "foo bar":
                var v13 int8 = 0
                _ = v13
            case "\n\t":
                var v14 float64 = (((0.0 - 1e10) - -2.25) - -2.25)
                _ = v14
            }
            _ = v12
        } else {
            if (114 != v3) {
                v3 -= v3
                v2["foo bar"] = 906
            }
            var v15 int8 = int8(v5)
            _ = v15
        }
    }
    _ = len(v4)
    var v16 int = 9223372036854775807
    var v17 uint32 = 988
    _ = len(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
