package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 int }
type GenS2 struct { F0 int }
func genFn0(p0 int8, p1 int) uint32 {
    if (0 < 465) {
        p0 = p0
    }
    if (1000 >= 1) {
        v5 := []int{(-(p1))}
        v6 := GenS1{F0: (^(p0)), F1: -9223372036854775808, F2: 127}
        v7 := GenS1{F0: (p0 << ((0) & 63)), F1: 0, F2: -57}
        _ = v5
        _ = v6
        _ = v7
    } else {
        for fi8 := 0; fi8 < 5; fi8++ {
            switch p1 {
            case 42:
                p0 = (p0 / (((p0 * (p0 >> ((18446744073709551615) & 63)))) | 1))
            }
            v9 := []int{}
            _ = fi8
            _ = v9
        }
        fmt.Println(p0)
        switch "a" {
        case "foo bar":
            v10 := []int{p1, (-(p1))}
            _ = v10
        case "hello":
            var v11 uint16 = 65535
            _ = v11
        case "\n\t":
            switch "\n\t" {
            case "a":
                p1 = (p1 % (((-(p1))) | 1))
            default:
                var v12 uint8 = 0
                _ = v12
            }
        }
    }
    for fi13 := 0; fi13 < 16; fi13++ {
        var v14 bool = true
        for fi15 := 0; fi15 < 11; fi15++ {
            v14 = v14
            _ = fi15
        }
        _ = fi13
        _ = v14
    }
    v16 := []int{}
    _ = p0
    _ = p1
    _ = v16
    return 0
}

func main() {
    var v1 int8 = 1
    v2 := []string{"\n\t", "a"}
    var v3 float64 = (0.1 / 1.0)
    v4 := make(map[uint32]bool)
    switch "hello" {
    case "":
        var v17 uint8 = 183
        _ = v17
    default:
        v3 = v3
    }
    v3 = v3
    var v18 int16 = 42
    var v19 uint8 = 1
    var v20 uint16 = 0
    fmt.Println(v1)
    var v21 int8 = 0
    if (1 >= 255) {
        if (-403 != v18) {
            v22 := GenS1{F0: (v1 * -128), F1: 42, F2: 128}
            _ = v22
        } else {
            if ((((v20 % (((v20 & v20)) | 1)) <= (v20 + (v20 & 1))) && (0 == 255)) && true) {
                for fi23 := 0; fi23 < 8; fi23++ {
                    if ("foo bar" != "hello") {
                        var v24 float64 = float64(0.0)
                        _ = v24
                    }
                    var v25 uint8 = v19
                    for fi26 := 0; fi26 < 14; fi26++ {
                        fmt.Println(fi23, fi23)
                        switch "a" {
                        case "":
                            v3 += v3
                        default:
                            v4[2147483647] = (genFn0(v21, (fi23 * (fi23 + 863))) == 42)
                        }
                        if (0 > 937) {
                            switch "" {
                            case "a":
                                v1 = v21
                            }
                            var v27 float64 = v3
                            _ = v27
                        }
                        _ = fi26
                    }
                    _ = fi23
                    _ = v25
                }
            } else {
                v19 = 42
                var v28 int8 = (v21 >> ((4294967295) & 63))
                _ = v28
            }
        }
    }
    v29 := GenS2{F0: 1}
    v19 = (v19 ^ v19)
    var v30 int = 0
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
