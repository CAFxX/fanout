package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 uint16 }
type GenS2 struct { F0 uint16; F1 int32; F2 bool }
type GenS3 struct { F0 bool; F1 int32 }
func genFn0(p0 uint64, p1 string) uint {
    var v7 uint64 = (p0 | p0)
    v7 = v7
    _ = p0
    _ = p1
    _ = v7
    return 18446744073709551615
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = float64(3.14159)
    v3 := make(map[string]int)
    var v4 float64 = v2
    var v5 string = "\n\t"
    var v6 bool = true
    v8 := GenS1{F0: 8, F1: 137, F2: 1}
    for fi9 := 0; fi9 < 3; fi9++ {
        if (false && true) {
            v4 -= v2
        } else {
            v10 := v3[(((v5 + ((("\n\t" + "foo bar") + ((v5 + ("foo bar" + v5)) + v5)) + ("hello" + v5))) + "foo bar") + "foo bar")]
            var v11 string = v5
            _ = v10
            _ = v11
        }
        var v12 string = ("foo bar" + (("hello" + (v5 + "")) + "foo bar"))
        v13 := GenS2{F0: 0, F1: 2147483647, F2: (v2 < (v4 / (0.1 * ((float64(1.5) / v4) / 1.0))))}
        _ = fi9
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 16; fi14++ {
        v2 += v4
        v6 = (42 > 1)
        _ = fi14
    }
    var v15 int32 = 225
    if false {
        v2 += ((float64(1e10) / v2) + v4)
    }
    v15 = (-(v15))
    if v6 {
        v4 = v4
        for fi16 := 0; fi16 < 20; fi16++ {
            v2 += v2
            if (49 > 42) {
                var v17 int64 = -314
                var v18 int16 = 1
                var v19 bool = false
                _ = v17
                _ = v18
                _ = v19
            } else {
                v4 = (float64(0.1) / v4)
            }
            switch fi16 {
            case 127:
                v15 = (v15 % (((v15 << ((504) & 63))) | 1))
            case -9223372036854775808:
                for fi20 := 0; fi20 < 13; fi20++ {
                    if ((v15 ^ v15) < v15) {
                        fmt.Println(v4, v4, v2)
                        v21 := GenS1{F0: 100, F1: 1, F2: 962}
                        switch 1 {
                        case -969:
                            var v22 int64 = -9223372036854775808
                            _ = v22
                        }
                        fmt.Println(fi20, v5)
                        if (v6 && (1 >= -198)) {
                            v23 := []int{fi20, (fi20 + fi16)}
                            if (v4 == 3.14159) {
                                var v24 uint64 = 4294967295
                                var v25 float64 = v2
                                _ = v24
                                _ = v25
                            } else {
                                var v26 float32 = 0.0
                                _ = v26
                            }
                            _ = v23
                        }
                        _ = v21
                    } else {
                        var v27 float64 = v4
                        v2 = 0.0
                        var v28 float64 = -2.25
                        _ = v27
                        _ = v28
                    }
                    _ = fi20
                }
            }
            _ = fi16
        }
    } else {
        fmt.Println(v6, v15)
        switch v5 {
        case "":
            delete(v3, (("a" + v5) + "\n\t"))
        case "\n\t":
            var v29 uint64 = 18446744073709551615
            _ = v29
        case "a":
            fmt.Println(v4, v6, v5)
        default:
            if v6 {
                if v6 {
                    switch "hello" {
                    case "hello":
                        v30 := GenS3{F0: v6, F1: (v15 * v15)}
                        _ = v30
                    case "a":
                        v6 = v6
                    default:
                        v1[78] = ((3 < 28) && ((!(722 <= 42)) || v6))
                    }
                }
            } else {
                _ = len(v1)
                if (121 != 42) {
                    for fi31 := 0; fi31 < 17; fi31++ {
                        var v32 int8 = 34
                        v33 := []int{fi31, (fi31 & fi31)}
                        _ = fi31
                        _ = v32
                        _ = v33
                    }
                    v34 := v3[((v5 + (v5 + ("a" + "foo bar"))) + "\n\t")]
                    var v35 uint16 = 2
                    _ = v34
                    _ = v35
                }
                v36 := GenS1{F0: -111, F1: 255, F2: 356}
                _ = v36
            }
        }
    }
    v15 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
