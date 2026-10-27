package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int32 }
type GenS2 struct { F0 uint; F1 uint8; F2 bool }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32; F1 uint8 }
type GenS5 struct { F0 uint }
func genFn0(p0 int64, p1 int8, p2 uint64) int64 {
    for fi7 := 0; fi7 < 6; fi7++ {
        if ((("\n\t" + ("foo bar" + "a")) + "foo bar") <= "\n\t") {
            v8 := []int{}
            for fi9 := 0; fi9 < 11; fi9++ {
                p2 = 0
                _ = fi9
            }
            _ = v8
        } else {
            p1 = p1
            var v10 int = 128
            _ = v10
        }
        if true {
            p2 -= p2
        }
        _ = fi7
    }
    var v11 string = ("a" + "\n\t")
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 1
}

func main() {
    v1 := make(map[int]bool)
    var v2 float64 = ((float64((0.1 * (3.14159 - 3.14159))) * -2.25) * -2.25)
    var v3 uint64 = 18446744073709551615
    var v4 int8 = 100
    v5 := make(map[string]bool)
    v6 := []int{}
    if (42 == 2147483647) {
        if false {
            v12 := GenS1{F0: 42, F1: 4294967295, F2: 42}
            v3 = (v3 << ((255) & 63))
            fmt.Println(v3)
            _ = v12
        }
        var v13 int32 = 1
        _ = v13
    } else {
        var v14 int64 = 0
        var v15 float32 = -2.25
        var v16 uint32 = 4294967295
        _ = v14
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 19; fi17++ {
        for fi18 := 0; fi18 < 12; fi18++ {
            v19 := GenS1{F0: 1, F1: 277, F2: 2147483647}
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    v2 -= (v2 - v2)
    var v20 uint = 2
    v3 += v3
    for fi21 := 0; fi21 < 6; fi21++ {
        v22 := GenS2{F0: v20, F1: 145, F2: ((2147483648 == 2147483647) && ((-2.25 / v2) < 3.14159))}
        if true {
            for fi23 := 0; fi23 < 8; fi23++ {
                switch (fi23 % ((fi21) | 1)) {
                case 32768:
                    v2 -= (((0.1 + float64(0.1)) / v2) * 1e10)
                case -546:
                    v24 := GenS3{F0: 2}
                    _ = v24
                default:
                    v6 = append(v6, 918)
                }
                v25 := v1[(fi21 / ((fi23) | 1))]
                _ = fi23
                _ = v25
            }
            v4 += v4
        }
        _ = fi21
        _ = v22
    }
    for fi26 := 0; fi26 < 10; fi26++ {
        fmt.Println(v4)
        if false {
            for fi27 := 0; fi27 < 11; fi27++ {
                _ = len(v6)
                v28 := GenS2{F0: (v20 << (((v20 ^ 684)) & 63)), F1: 2, F2: (2 < 921)}
                if ((!(true || (!false))) && (42 >= 1000)) {
                    var v29 int = (fi26 | (fi26 % (((fi26 & (fi26 ^ (^(fi27))))) | 1)))
                    _ = v29
                } else {
                    var v30 uint64 = v3
                    v20 -= (v20 % ((42) | 1))
                    _ = v30
                }
                _ = fi27
                _ = v28
            }
            if true {
                v20 = v20
                var v31 int = -741
                _ = v31
            } else {
                var v32 string = "\n\t"
                fmt.Println(fi26, v20)
                _ = v32
            }
            for fi33 := 0; fi33 < 18; fi33++ {
                switch -1 {
                case 2147483647:
                    fmt.Println(fi26, fi26)
                case 32768:
                    v20 = (v20 & v20)
                case 9223372036854775807:
                    if (!(!true)) {
                        fmt.Println(v3, v2, v20)
                        fmt.Println(fi33)
                    }
                    v34 := GenS4{F0: -668, F1: 128}
                    v2 += float64(0.0)
                    _ = v34
                }
                _ = fi33
            }
        } else {
            switch "a" {
            case "foo bar":
                var v35 int64 = 0
                _ = v35
            case "\n\t":
                for fi36 := 0; fi36 < 17; fi36++ {
                    v37 := GenS5{F0: v20}
                    _ = fi36
                    _ = v37
                }
            default:
                v20 += (v20 / ((v20) | 1))
            }
            var v38 int = (fi26 % ((fi26) | 1))
            _ = v38
        }
        var v39 int32 = 42
        _ = fi26
        _ = v39
    }
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
