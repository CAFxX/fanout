package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 int8 }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint64 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{-807, 9223372036854775807, 32768}
    var v3 uint = 546
    v4 := []int{-207, -2147483648}
    var v5 int32 = 2147483647
    var v6 string = "a"
    switch (("\n\t" + "") + v6) {
    case "\n\t":
        fmt.Println(v5, v6, v6)
    }
    v5 += v5
    fmt.Println(v3, v6, v6)
    v3 += v3
    for fi7 := 0; fi7 < 2; fi7++ {
        if false {
            if true {
                v3 += 2
            } else {
                for fi8 := 0; fi8 < 16; fi8++ {
                    v9 := GenS1{F0: fi8, F1: 2, F2: 0}
                    _ = fi8
                    _ = v9
                }
            }
            if (((1e10 - -2.25) * (3.14159 + -2.25)) == (3.14159 / ((0.1 + float64((0.1 / 1e10))) - float64(0.1)))) {
                delete(v1, 931)
            } else {
                v10 := GenS2{F0: "", F1: 109}
                v6 = ("foo bar" + "hello")
                _ = v10
            }
        } else {
            v11 := GenS2{F0: v6, F1: 878}
            switch (fi7 * (fi7 * fi7)) {
            case 9223372036854775807:
                var v12 bool = false
                _ = v12
            case 42:
                switch v6 {
                case "hello":
                    v5 = v5
                default:
                    if ((2147483648 != 944) || true) {
                        if (1 >= 0) {
                            if true {
                                var v13 uint32 = 42
                                var v14 int8 = 1
                                var v15 float64 = 0.1
                                _ = v13
                                _ = v14
                                _ = v15
                            }
                            v6 = v6
                            v6 = v6
                        } else {
                            v5 -= v5
                        }
                    }
                }
            }
            var v16 uint = (v3 | v3)
            _ = v11
            _ = v16
        }
        _ = fi7
    }
    for fi17 := 0; fi17 < 3; fi17++ {
        v5 -= v5
        v3 = v3
        v18 := v1[358]
        _ = fi17
        _ = v18
    }
    v6 = ""
    if (!(32767 < 42)) {
        if ((v3 >> ((v3) & 63)) == (v3 + 537)) {
            v19 := GenS3{F0: 32768}
            fmt.Println(v3, v3, v3)
            v20 := GenS3{F0: 1}
            _ = v19
            _ = v20
        }
        if len(v2) > 0 {
            v2[0] = 0
        }
    }
    switch v6 {
    case "\n\t":
        v6 = v6
    case "hello":
        v3 += v3
    case "":
        var v21 uint = 65535
        _ = v21
    }
    v3 = v3
    fmt.Println(v3, v5)
    for fi22 := 0; fi22 < 18; fi22++ {
        for fi23 := 0; fi23 < 5; fi23++ {
            if len(v4) > 0 {
                v4[0] = int(v5)
            }
            var v24 float64 = ((1e10 * (1e10 * float64(0.0))) - -2.25)
            var v25 float64 = float64(1e10)
            v3 = v3
            _ = fi23
            _ = v24
            _ = v25
        }
        v26 := GenS1{F0: (fi22 % ((fi22) | 1)), F1: (v3 - (v3 >> ((18446744073709551615) & 63))), F2: 54}
        fmt.Println(v6, v5, fi22)
        _ = fi22
        _ = v26
    }
    fmt.Println(v5)
    v27 := GenS4{F0: 662}
    v28 := GenS1{F0: 2147483647, F1: (v3 & (v3 & 255)), F2: 127}
    v6 = ((v6 + (v6 + v6)) + v6)
    v29 := GenS2{F0: "\n\t", F1: uint64(v3)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
