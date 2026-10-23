package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 uint32 }
func genFn0(p0 bool, p1 uint8) uint8 {
    if true {
        var v7 string = "a"
        _ = v7
    } else {
        fmt.Println(p0, p0)
    }
    p1 -= (p1 % (((p1 >> ((0) & 63))) | 1))
    p0 = (!p0)
    _ = p0
    _ = p1
    return (p1 - p1)
}
func genFn1(p0 bool, p1 string) int64 {
    v8 := GenS1{F0: 18446744073709551615, F1: 127}
    v9 := []int{296}
    for fi10 := 0; fi10 < 18; fi10++ {
        p0 = p0
        _ = fi10
    }
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return -140
}

func main() {
    var v1 int8 = 1
    v2 := make(map[string]int)
    var v3 bool = ((false && (1 != 1000)) && ((-2147483648 != -1) && true))
    v4 := make(map[uint32]string)
    var v5 int16 = 1
    var v6 int8 = (v1 % (((v1 % (((v1 ^ (-(v1)))) | 1))) | 1))
    v3 = ((((3.14159 * ((float64(0.0) + (0.1 - 0.1)) - 3.14159)) + 0.0) - (0.0 / 1.0)) >= ((0.1 + 0.1) + (float64((1.5 + 3.14159)) / 1.5)))
    var v11 int16 = v5
    v4[963] = "foo bar"
    v12 := []int{949}
    var v13 int = 9223372036854775807
    v12 = append(v12, (v13 ^ v13))
    if (false || (v3 || (-1 != -801))) {
        if true {
            for fi14 := 0; fi14 < 20; fi14++ {
                if true {
                    v11 = v11
                    var v15 int32 = -2147483648
                    _ = v15
                } else {
                    if (!true) {
                        v16 := GenS2{F0: "\n\t", F1: v6}
                        delete(v4, 682)
                        var v17 float64 = 0.0
                        _ = v16
                        _ = v17
                    } else {
                        if (!(v3 || (1.5 <= (0.1 * float64((-2.25 / 3.14159)))))) {
                            var v18 int64 = 9223372036854775807
                            if false {
                                v19 := v4[59]
                                v20 := GenS3{F0: 2, F1: 4294967295}
                                _ = v19
                                _ = v20
                            } else {
                                fmt.Println(v11, v13, v11)
                                var v21 uint64 = 42
                                _ = v21
                            }
                            var v22 uint16 = 2
                            _ = v18
                            _ = v22
                        }
                    }
                }
                _ = fi14
            }
        } else {
            v13 += v13
            v6 = v6
            v2["foo bar"] = (v13 / (((v13 + 4294967295)) | 1))
        }
        var v23 float32 = -2.25
        _ = v23
    }
    switch "foo bar" {
    case "foo bar":
        v3 = v3
    case "a":
        v3 = ((!false) || ((-(v13)) > (v13 / (((v13 & v13)) | 1))))
    default:
        fmt.Println(v5)
    }
    for fi24 := 0; fi24 < 18; fi24++ {
        switch fi24 {
        case 32768:
            var v25 uint16 = 65535
            _ = v25
        default:
            v26 := GenS4{F0: 409}
            _ = v26
        }
        v27 := GenS1{F0: 4294967295, F1: (v13 - (fi24 / ((v13) | 1)))}
        _ = fi24
        _ = v27
    }
    v13 = v13
    var v28 float64 = -2.25
    var v29 uint16 = 131
    var v30 int32 = -814
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
