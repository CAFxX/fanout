package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 int64 }
type GenS2 struct { F0 string; F1 bool; F2 uint32 }
type GenS3 struct { F0 uint32; F1 uint64; F2 bool }
type GenS4 struct { F0 string; F1 uint64; F2 bool }
func genFn0() uint16 {
    if false {
        var v6 int16 = 68
        for fi7 := 0; fi7 < 5; fi7++ {
            v6 -= v6
            v6 -= v6
            v6 = v6
            _ = fi7
        }
        v8 := GenS1{F0: -1, F1: 2, F2: -1}
        _ = v6
        _ = v8
    } else {
        for fi9 := 0; fi9 < 11; fi9++ {
            switch "" {
            case "":
                v10 := GenS1{F0: 127, F1: 2147483647, F2: -234}
                _ = v10
            case "a":
                fmt.Println(fi9)
            case "foo bar":
                var v11 string = ""
                _ = v11
            default:
                v12 := GenS1{F0: -112, F1: 1, F2: 42}
                _ = v12
            }
            v13 := []int{0, (fi9 >> ((0) & 63)), 0, -869}
            _ = fi9
            _ = v13
        }
        v14 := GenS2{F0: "foo bar", F1: ((3.14159 / ((((1e10 / 1.5) - 0.1) - -2.25) + 0.1)) >= 1.5), F2: 2147483648}
        var v15 int32 = 979
        _ = v14
        _ = v15
    }
    v16 := GenS1{F0: 42, F1: 65535, F2: 47}
    var v17 int64 = 42
    v18 := GenS2{F0: ("\n\t" + ("hello" + "\n\t")), F1: true, F2: 2147483647}
    v17 = -1
    _ = v16
    _ = v17
    _ = v18
    return 244
}
func genFn1(p0 int64) uint8 {
    var v19 uint64 = 18446744073709551615
    v20 := GenS3{F0: 4294967295, F1: (v19 - 18446744073709551615), F2: (0 <= 0)}
    p0 -= (^(p0))
    _ = p0
    _ = v19
    _ = v20
    return 255
}

func main() {
    var v1 string = ("a" + "foo bar")
    var v2 int32 = 42
    v3 := make(map[int]uint32)
    var v4 uint32 = 42
    var v5 uint64 = 850
    v21 := v3[0]
    switch -9223372036854775808 {
    case 778:
        for fi22 := 0; fi22 < 5; fi22++ {
            switch 2147483647 {
            case 0:
                for fi23 := 0; fi23 < 3; fi23++ {
                    v24 := GenS1{F0: 75, F1: (v4 & 2147483647), F2: -825}
                    _ = fi23
                    _ = v24
                }
            case -867:
                switch (fi22 ^ fi22) {
                case -75:
                    switch fi22 {
                    case -1:
                        if (v2 < 1) {
                            v4 = 65535
                            switch ("foo bar" + "hello") {
                            case "":
                                v25 := GenS4{F0: v1, F1: (v5 ^ uint64(v4)), F2: (0 > 985)}
                                _ = v25
                            case "hello":
                                v2 -= int32(fi22)
                            case "\n\t":
                                v3[fi22] = 4294967295
                            }
                            v4 = 123
                        }
                    case 417:
                        fmt.Println(fi22, v2)
                    }
                case 1:
                    v26 := []int{(fi22 / ((128) | 1)), (fi22 / ((fi22) | 1)), fi22, -350}
                    _ = v26
                case 182:
                    v27 := []int{(fi22 >> ((65535) & 63)), fi22, -730, fi22}
                    _ = v27
                default:
                    v28 := GenS2{F0: (("" + "a") + ((v1 + v1) + "a")), F1: (v2 != (v2 / ((566) | 1))), F2: v21}
                    _ = v28
                }
            case 587:
                for fi29 := 0; fi29 < 15; fi29++ {
                    v30 := []int{509, (fi29 * fi22), -484}
                    v2 = (-(v2))
                    if (255 >= 2) {
                        fmt.Println(v5, v1)
                    } else {
                        fmt.Println(v4, v4, v21)
                        var v31 int32 = v2
                        var v32 string = v1
                        if len(v30) >= 2 {
                            v30 = v30[1:len(v30)]
                        }
                        _ = v31
                        _ = v32
                    }
                    v1 = "a"
                    delete(v3, fi22)
                    switch "\n\t" {
                    case "a":
                        if ((v2 | (v2 * v2)) > (v2 ^ v2)) {
                            var v33 int32 = (v2 * (v2 & v2))
                            _ = v33
                        }
                    }
                    _ = fi29
                    _ = v30
                }
            default:
                v34 := []int{}
                _ = v34
            }
            v5 -= v5
            v21 -= (v4 - (v21 - v21))
            _ = fi22
        }
    }
    var v35 uint = 42
    v21 = 900
    var v36 float64 = -2.25
    fmt.Println(v4, v2, v1)
    v37 := GenS4{F0: (v1 + ("" + (v1 + v1))), F1: (v5 & (^(v5))), F2: false}
    delete(v3, -2147483648)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
