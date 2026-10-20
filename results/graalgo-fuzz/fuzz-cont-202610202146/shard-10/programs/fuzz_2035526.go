package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 int8; F2 int }

func main() {
    var v1 int64 = 865
    var v2 uint8 = 128
    var v3 float64 = -2.25
    var v4 int = 2147483647
    v5 := []int{v4, v4}
    if len(v5) > 0 {
        v5[0] = v4
    }
    v1 += (v1 >> ((930) & 63))
    fmt.Println(v4)
    v5 = append(v5, (v4 & v4))
    v1 = v1
    var v6 uint = uint(v2)
    v4 = (v4 - (v4 - v4))
    if len(v5) > 0 {
        v5[0] = -79
    }
    for fi7 := 0; fi7 < 7; fi7++ {
        v6 += v6
        for fi8 := 0; fi8 < 6; fi8++ {
            fmt.Println(v1, v1, v3)
            if ((true || (4294967295 <= 2147483647)) && (42 == 0)) {
                switch "hello" {
                case "":
                    var v9 int = 718
                    _ = v9
                case "\n\t":
                    v4 -= (fi8 + (fi8 % ((32767) | 1)))
                case "hello":
                    v1 = int64(fi8)
                default:
                    if false {
                        v10 := GenS1{F0: (-(fi7)), F1: 255, F2: (("\n\t" + (("foo bar" + "") + "hello")) + "hello")}
                        v5 = append(v5, fi8)
                        _ = v10
                    }
                }
            } else {
                if (59 < 204) {
                    v3 = 0.0
                    v3 -= v3
                    var v11 int64 = v1
                    _ = v11
                } else {
                    var v12 uint = 18446744073709551615
                    v6 += (v12 >> ((661) & 63))
                    var v13 int = v4
                    _ = v12
                    _ = v13
                }
                var v14 int64 = v1
                switch (fi8 * (v4 / ((32767) | 1))) {
                case 552:
                    v1 = v14
                default:
                    v15 := GenS2{F0: 42}
                    _ = v15
                }
                _ = v14
            }
            var v16 uint16 = 778
            v6 -= (v6 & 42)
            _ = fi8
            _ = v16
        }
        var v17 uint16 = 620
        _ = fi7
        _ = v17
    }
    _ = len(v5)
    if (false || ((v6 - (v6 << ((v6) & 63))) == (v6 | (v6 - (^(v6)))))) {
        v18 := GenS3{F0: (v2 >> ((389) & 63)), F1: -1, F2: v4}
        for fi19 := 0; fi19 < 11; fi19++ {
            for fi20 := 0; fi20 < 12; fi20++ {
                v6 = v6
                _ = fi20
            }
            _ = fi19
        }
        _ = v18
    } else {
        for fi21 := 0; fi21 < 2; fi21++ {
            switch int(v6) {
            case 32767:
                if true {
                    v3 += v3
                }
            case 4294967295:
                var v22 float32 = 1e10
                _ = v22
            case 128:
                v23 := GenS3{F0: 73, F1: 42, F2: 164}
                _ = v23
            }
            for fi24 := 0; fi24 < 16; fi24++ {
                var v25 int32 = 42
                v26 := GenS3{F0: 100, F1: -59, F2: (v4 % (((fi21 % (((fi21 >> ((uint(v2)) & 63))) | 1))) | 1))}
                if true {
                    switch "foo bar" {
                    case "":
                        var v27 float64 = v3
                        _ = v27
                    default:
                        var v28 uint64 = 4294967295
                        _ = v28
                    }
                    v1 = int64(v25)
                    v4 += (fi24 % ((0) | 1))
                }
                fmt.Println(fi21, v25, v3)
                _ = fi24
                _ = v25
                _ = v26
            }
            v29 := GenS2{F0: fi21}
            _ = fi21
            _ = v29
        }
        v30 := GenS1{F0: (v4 - 128), F1: 42, F2: "a"}
        _ = v30
    }
    var v31 uint32 = 113
    if true {
        v3 = float64(1e10)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
