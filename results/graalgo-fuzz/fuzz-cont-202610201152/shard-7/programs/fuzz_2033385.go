package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 bool }
type GenS2 struct { F0 int32; F1 uint16; F2 int64 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 int64, p1 uint16) bool {
    v4 := []int{9223372036854775807, 165}
    var v5 int8 = 0
    if true {
        p1 += 0
    } else {
        var v6 string = "foo bar"
        for fi7 := 0; fi7 < 9; fi7++ {
            for fi8 := 0; fi8 < 16; fi8++ {
                if (p1 != (p1 - 678)) {
                    var v9 int16 = -77
                    _ = v9
                } else {
                    if true {
                        v10 := GenS1{F0: "\n\t", F1: 18446744073709551615, F2: (42 <= 4294967295)}
                        _ = v10
                    }
                    switch "foo bar" {
                    case "\n\t":
                        var v11 int16 = -1
                        _ = v11
                    case "a":
                        v12 := GenS1{F0: v6, F1: 2, F2: true}
                        _ = v12
                    }
                }
                var v13 uint = 65535
                _ = fi8
                _ = v13
            }
            v5 = (-(v5))
            _ = fi7
        }
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return true
}
func genFn1(p0 uint, p1 int8, p2 int64) int {
    v14 := GenS1{F0: "foo bar", F1: (p0 << (((^(p0))) & 63)), F2: false}
    p1 -= 0
    var v15 uint32 = 2147483648
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    return 2147483647
}

func main() {
    var v1 int32 = 2147483647
    var v2 string = ("a" + ("hello" + "a"))
    v3 := []string{}
    switch (v2 + v2) {
    case "hello":
        var v16 float64 = 0.0
        _ = v16
    default:
        for fi17 := 0; fi17 < 15; fi17++ {
            fmt.Println(v1, fi17, v2)
            _ = fi17
        }
    }
    for fi18 := 0; fi18 < 13; fi18++ {
        _ = len(v3)
        v1 -= v1
        _ = fi18
    }
    v19 := GenS1{F0: v2, F1: 508, F2: (!(41 < -128))}
    v2 = v2
    var v20 uint16 = 220
    var v21 uint = 4294967295
    v2 = v2
    v22 := GenS2{F0: (v1 ^ 1), F1: 448, F2: -663}
    v23 := GenS1{F0: (("a" + v2) + v2), F1: (v21 / (((^(v21))) | 1)), F2: (false && (!false))}
    v24 := GenS2{F0: v1, F1: (v20 * 0), F2: 0}
    switch v2 {
    case "foo bar":
        v2 = v2
    }
    if true {
        var v25 int = 2147483647
        _ = v25
    }
    v26 := GenS3{F0: 2}
    var v27 string = "a"
    v20 = (v20 % (((v20 * (v20 + 734))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v26
    _ = v27
    _ = v3
}
