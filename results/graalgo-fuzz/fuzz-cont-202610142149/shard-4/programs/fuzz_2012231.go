package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 uint8; F1 uint16 }

func main() {
    v1 := []int{9223372036854775807, 348, 2147483647, 285}
    v2 := make(map[int]int)
    var v3 uint32 = 65535
    var v4 int8 = 42
    v5 := make(map[int]bool)
    for fi6 := 0; fi6 < 17; fi6++ {
        for fi7 := 0; fi7 < 19; fi7++ {
            if true {
                v3 += v3
                v8 := GenS1{F0: "\n\t", F1: true}
                v2[(fi6 % ((fi6) | 1))] = (fi7 - -9223372036854775808)
                _ = v8
            } else {
                var v9 uint64 = 4294967295
                v9 = v9
                delete(v2, 0)
                _ = v9
            }
            _ = fi7
        }
        _ = len(v1)
        _ = fi6
    }
    var v10 int = 9223372036854775807
    switch v10 {
    case -264:
        var v11 uint16 = 0
        _ = v11
    case 4294967295:
        var v12 int32 = -1
        _ = v12
    case -29:
        var v13 uint = 42
        _ = v13
    default:
        if true {
            v4 -= v4
            fmt.Println(v4, v3, v3)
        }
    }
    for fi14 := 0; fi14 < 13; fi14++ {
        for fi15 := 0; fi15 < 16; fi15++ {
            switch fi15 {
            case 32768:
                v3 = (v3 ^ v3)
            case 897:
                v16 := GenS2{F0: 2, F1: 32767}
                _ = v16
            default:
                if (float64(0.0) != 0.0) {
                    _ = len(v1)
                }
            }
            var v17 uint = 18446744073709551615
            fmt.Println(v3, fi14, v3)
            _ = fi15
            _ = v17
        }
        _ = fi14
    }
    v18 := v2[(v10 - (v10 ^ 32767))]
    v18 = v10
    var v19 int32 = 643
    v20 := GenS2{F0: 0, F1: uint16(v3)}
    if true {
        switch (v18 & v10) {
        case 42:
            var v21 uint32 = (^(v3))
            _ = v21
        case 32767:
            for fi22 := 0; fi22 < 19; fi22++ {
                v23 := GenS1{F0: "foo bar", F1: false}
                _ = fi22
                _ = v23
            }
        default:
            var v24 uint16 = 997
            _ = v24
        }
        v1 = append(v1, (v18 ^ (v10 + (v10 | 70))))
    } else {
        v4 -= v4
        v3 += 445
    }
    if (!true) {
        var v25 uint32 = v3
        v10 = (-(v10))
        _ = v25
    }
    v26 := GenS1{F0: ("hello" + "\n\t"), F1: (157 >= 248)}
    var v27 bool = true
    switch "" {
    case "":
        if (true && (!(v3 != v3))) {
            for fi28 := 0; fi28 < 12; fi28++ {
                v1 = append(v1, v10)
                for fi29 := 0; fi29 < 6; fi29++ {
                    switch "foo bar" {
                    case "":
                        v30 := GenS2{F0: 36, F1: 2}
                        _ = v30
                    case "a":
                        var v31 int = 0
                        _ = v31
                    default:
                        v32 := v2[v10]
                        _ = v32
                    }
                    if v27 {
                        v3 += (v3 << ((2) & 63))
                        var v33 bool = (v4 < 0)
                        _ = v33
                    } else {
                        v10 += 196
                    }
                    switch fi29 {
                    case -9223372036854775808:
                        v3 -= (v3 / (((v3 - (v3 * v3))) | 1))
                    default:
                        fmt.Println(fi29)
                    }
                    _ = fi29
                }
                var v34 uint16 = 32767
                _ = fi28
                _ = v34
            }
        }
    case "a":
        if (v19 < v19) {
            v4 = (v4 >> ((255) & 63))
            v27 = v27
        }
    case "foo bar":
        if v27 {
            v27 = v27
        }
    default:
        v2[42] = v10
    }
    v10 = 690
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
