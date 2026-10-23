package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 uint64 }
type GenS2 struct { F0 bool; F1 int }
func genFn0(p0 uint16, p1 uint8) int64 {
    p0 = (p0 ^ (p0 & (p0 + p0)))
    v4 := []int{}
    _ = p0
    _ = p1
    _ = v4
    return 0
}
func genFn1(p0 int32, p1 int8, p2 string) bool {
    var v5 uint16 = 32767
    p1 = p1
    for fi6 := 0; fi6 < 16; fi6++ {
        p0 = p0
        v5 = (v5 | (v5 % (((^(v5))) | 1)))
        _ = fi6
    }
    if (255 >= 2) {
        var v7 uint8 = 22
        if ((v7 | (v7 | v7)) >= v7) {
            for fi8 := 0; fi8 < 9; fi8++ {
                p0 += -1
                v9 := GenS1{F0: -63, F1: 42, F2: 212}
                _ = fi8
                _ = v9
            }
            v10 := GenS1{F0: 9223372036854775807, F1: 1, F2: 42}
            var v11 uint8 = v7
            _ = v10
            _ = v11
        } else {
            p2 = p2
        }
        v12 := GenS1{F0: 842, F1: 1, F2: 921}
        _ = v12
        _ = v7
    } else {
        v13 := GenS1{F0: int64(v5), F1: int64(p0), F2: 18446744073709551615}
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (true || ((!(172 <= 4294967295)) && (!true)))
}

func main() {
    v1 := make(map[int]string)
    var v2 float64 = 1e10
    v3 := []string{"hello", "\n\t", "hello", "a"}
    switch -1 {
    case 32767:
        for fi14 := 0; fi14 < 20; fi14++ {
            delete(v1, 9223372036854775807)
            _ = fi14
        }
    case 127:
        var v15 uint8 = 2
        _ = v15
    default:
        switch "hello" {
        case "":
            v2 -= v2
        default:
            var v16 int16 = 0
            _ = v16
        }
    }
    var v17 float32 = 1.5
    _ = len(v3)
    v2 += ((((v2 * v2) - 0.1) / v2) * ((1e10 / v2) / 1.0))
    switch 32767 {
    case 138:
        var v18 int64 = -135
        _ = v18
    }
    switch 1 {
    case 32767:
        var v19 int = -9223372036854775808
        _ = v19
    case -9223372036854775808:
        for fi20 := 0; fi20 < 13; fi20++ {
            var v21 bool = (true || (-1 <= -110))
            if false {
                for fi22 := 0; fi22 < 7; fi22++ {
                    fmt.Println(v17, fi20)
                    if (!(-9223372036854775808 <= 611)) {
                        var v23 bool = (v2 < v2)
                        var v24 float64 = (v2 / (v2 / 1e10))
                        if v21 {
                            fmt.Println(v24, fi20, v23)
                            var v25 int8 = 127
                            switch ("\n\t" + "a") {
                            case "\n\t":
                                if v21 {
                                    v1[-2147483648] = ""
                                    v26 := v1[fi22]
                                    _ = v26
                                } else {
                                    if (!true) {
                                        v27 := GenS1{F0: int64(fi22), F1: genFn0(42, 1), F2: 767}
                                        v28 := GenS1{F0: -295, F1: 42, F2: 42}
                                        _ = v27
                                        _ = v28
                                    }
                                    v21 = v23
                                }
                            case "foo bar":
                                var v29 int32 = 1
                                _ = v29
                            }
                            _ = v25
                        }
                        _ = v23
                        _ = v24
                    } else {
                        if v21 {
                            fmt.Println(v17, v17, v21)
                        } else {
                            delete(v1, -427)
                            if true {
                                var v30 int = 743
                                _ = len(v3)
                                _ = v30
                            } else {
                                v21 = (!false)
                                switch fi22 {
                                case -569:
                                    v17 = v17
                                case -889:
                                    fmt.Println(v17, fi22, fi20)
                                }
                                v1[fi22] = "hello"
                            }
                        }
                        v31 := GenS2{F0: true, F1: fi22}
                        _ = v31
                    }
                    _ = fi22
                }
                v32 := GenS1{F0: -974, F1: -9223372036854775808, F2: 849}
                var v33 int64 = -301
                _ = v32
                _ = v33
            }
            _ = fi20
            _ = v21
        }
    case -616:
        var v34 int32 = -332
        _ = v34
    }
    var v35 uint16 = 1
    var v36 int8 = -128
    v36 -= (-(v36))
    for fi37 := 0; fi37 < 2; fi37++ {
        v35 = 255
        var v38 bool = true
        _ = fi37
        _ = v38
    }
    fmt.Println(v2, v17)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v35
    _ = v36
}
