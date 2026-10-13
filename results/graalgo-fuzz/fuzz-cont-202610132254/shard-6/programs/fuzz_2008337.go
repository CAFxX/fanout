package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 string; F1 int }
type GenS5 struct { F0 int16; F1 bool }
func genFn0(p0 int16, p1 uint) bool {
    fmt.Println(p0)
    if true {
        if (1.5 >= float64(3.14159)) {
            var v6 float32 = 1.5
            v7 := []int{2147483647, -236, 1}
            _ = v6
            _ = v7
        }
        v8 := []int{471, 9223372036854775807, 0}
        _ = v8
    }
    _ = p0
    _ = p1
    return ("" != "a")
}
func genFn1() int64 {
    var v9 int = -676
    v9 -= v9
    var v10 bool = ((v9 == v9) || ((!(-2.25 <= float64(0.1))) || false))
    v11 := GenS1{F0: "a", F1: -1, F2: 999}
    v9 -= v9
    _ = v10
    _ = v11
    _ = v9
    return 1
}

func main() {
    var v1 float32 = 1.5
    v2 := make(map[string]bool)
    var v3 bool = true
    var v4 int64 = 0
    var v5 bool = v3
    if (false || (int(v4) < 42)) {
        v4 = v4
        v12 := v2["\n\t"]
        _ = v12
    } else {
        if false {
            for fi13 := 0; fi13 < 13; fi13++ {
                var v14 float32 = (float32((1e10 * 0.1)) - (v1 + v1))
                if v3 {
                    if (fi13 >= (fi13 ^ -2147483648)) {
                        var v15 float64 = 1.5
                        v4 = v4
                        if (v5 || (481 < 42)) {
                            for fi16 := 0; fi16 < 16; fi16++ {
                                var v17 uint64 = 18446744073709551615
                                if v5 {
                                    var v18 uint = 1
                                    var v19 float64 = v15
                                    v20 := GenS2{F0: (v4 / ((-1) | 1))}
                                    _ = v18
                                    _ = v19
                                    _ = v20
                                }
                                v21 := GenS3{F0: 122}
                                _ = fi16
                                _ = v17
                                _ = v21
                            }
                        }
                        _ = v15
                    } else {
                        v3 = v3
                        var v22 int32 = 1
                        v23 := GenS4{F0: "foo bar", F1: fi13}
                        _ = v22
                        _ = v23
                    }
                } else {
                    v14 += -2.25
                    v5 = true
                    var v24 float64 = 0.0
                    _ = v24
                }
                v25 := v2["a"]
                _ = fi13
                _ = v14
                _ = v25
            }
            var v26 int16 = 42
            v2["hello"] = v3
            _ = v26
        }
    }
    fmt.Println(v4, v3, v4)
    switch "\n\t" {
    case "":
        var v27 bool = (0.0 != 0.0)
        _ = v27
    case "foo bar":
        if false {
            var v28 float32 = float32((float64(0.0) - 3.14159))
            _ = v28
        }
    }
    var v29 int32 = 522
    var v30 uint64 = 1
    v29 = v29
    v4 -= (v4 / (((v4 * 703)) | 1))
    v31 := []int{32767, 9223372036854775807, 2147483647}
    for fi32 := 0; fi32 < 11; fi32++ {
        var v33 int8 = 100
        if ((573 >= genFn1()) || (!false)) {
            v29 -= (v29 + (v29 / (((-(v29))) | 1)))
        }
        _ = fi32
        _ = v33
    }
    var v34 uint16 = 1
    v34 = v34
    var v35 int8 = -128
    v3 = (!false)
    if v3 {
        for fi36 := 0; fi36 < 20; fi36++ {
            if (((((-2.25 / (1.5 * (3.14159 * 1.5))) / 0.1) * 0.0) + ((3.14159 / 1e10) / (0.0 - 3.14159))) <= 3.14159) {
                _ = len(v2)
            }
            _ = fi36
        }
    } else {
        var v37 string = ("a" + "\n\t")
        _ = v37
    }
    switch "" {
    case "hello":
        v34 = 0
    case "a":
        switch 42 {
        case -2147483648:
            var v38 uint16 = v34
            _ = v38
        case 127:
            var v39 uint64 = (v30 | (v30 % ((v30) | 1)))
            _ = v39
        case -771:
            v29 += v29
        default:
            var v40 string = "hello"
            _ = v40
        }
    case "\n\t":
        for fi41 := 0; fi41 < 16; fi41++ {
            v30 += 2
            for fi42 := 0; fi42 < 18; fi42++ {
                v3 = true
                v30 += (v30 ^ 60)
                _ = fi42
            }
            _ = fi41
        }
    default:
        v43 := GenS5{F0: 0, F1: v3}
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
