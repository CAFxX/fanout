package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 bool }
type GenS2 struct { F0 string; F1 int8 }
func genFn0() int32 {
    var v7 int = -827
    var v8 uint32 = 42
    var v9 int = -912
    if false {
        v10 := []int{42, (v7 >> ((334) & 63)), (^(v9)), (v9 | v9)}
        v11 := GenS1{F0: 3, F1: "foo bar", F2: true}
        if (18446744073709551615 >= 2) {
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            _ = len(v10)
            switch "hello" {
            case "foo bar":
                fmt.Println(v8, v8)
            default:
                v7 += 0
            }
        }
        _ = v10
        _ = v11
    }
    var v12 int64 = 42
    _ = v12
    _ = v7
    _ = v8
    _ = v9
    return 955
}
func genFn1(p0 uint8) string {
    v13 := GenS2{F0: "hello", F1: 1}
    p0 = 57
    _ = p0
    _ = v13
    return ("hello" + ("foo bar" + (((("\n\t" + "hello") + ("" + "")) + ("\n\t" + ("\n\t" + ("a" + "a")))) + "hello")))
}

func main() {
    v1 := []string{}
    v2 := make(map[int]string)
    var v3 uint8 = 255
    var v4 int8 = 42
    v5 := make(map[int]uint32)
    var v6 float64 = ((float64(1e10) - (0.1 + float64((1e10 / 1e10)))) + 1e10)
    if false {
        fmt.Println(v3)
    } else {
        v3 -= v3
        v4 = (v4 % ((v4) | 1))
        switch "foo bar" {
        case "hello":
            for fi14 := 0; fi14 < 8; fi14++ {
                v4 = (-(v4))
                v15 := GenS1{F0: 255, F1: "\n\t", F2: true}
                for fi16 := 0; fi16 < 15; fi16++ {
                    var v17 bool = (-1 >= 42)
                    v18 := GenS2{F0: "foo bar", F1: 0}
                    v17 = (2 <= 18446744073709551615)
                    var v19 uint32 = 2
                    if (65535 != 32767) {
                        if (-2.25 != v6) {
                            if (1 <= 398) {
                                var v20 uint32 = v19
                                _ = v20
                            } else {
                                fmt.Println(v3, v4, v4)
                                v3 -= 127
                                v3 -= (^(v3))
                            }
                            var v21 bool = ((fi16 == (fi16 >> ((18446744073709551615) & 63))) && v17)
                            v4 += 0
                            _ = v21
                        } else {
                            v22 := GenS1{F0: (^(v3)), F1: "hello", F2: v17}
                            v23 := GenS2{F0: "foo bar", F1: -1}
                            if v17 {
                                v4 += 14
                            } else {
                                v6 += v6
                            }
                            _ = v22
                            _ = v23
                        }
                        var v24 int32 = -810
                        v25 := GenS2{F0: "", F1: (v4 / (((v4 - -128)) | 1))}
                        _ = v24
                        _ = v25
                    } else {
                        var v26 bool = true
                        _ = v26
                    }
                    var v27 float32 = (-2.25 - 0.1)
                    _ = fi16
                    _ = v17
                    _ = v18
                    _ = v19
                    _ = v27
                }
                _ = fi14
                _ = v15
            }
        case "a":
            var v28 uint64 = 18446744073709551615
            _ = v28
        }
    }
    v4 = -17
    v6 += v6
    if false {
        var v29 int16 = 32767
        _ = v29
    }
    var v30 int = 9223372036854775807
    for fi31 := 0; fi31 < 8; fi31++ {
        if len(v1) > 0 {
            v1[0] = ""
        }
        _ = fi31
    }
    var v32 int32 = 1
    var v33 uint8 = v3
    v34 := v5[v30]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
