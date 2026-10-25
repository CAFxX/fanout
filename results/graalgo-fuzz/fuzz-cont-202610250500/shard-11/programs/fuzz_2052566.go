package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 int; F1 int32; F2 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int16 }
func genFn0(p0 uint32, p1 int64) int8 {
    if (-28 > 4294967295) {
        v4 := GenS1{F0: "hello", F1: 18446744073709551615}
        v5 := []int{32, 42, 32768, -1}
        _ = v4
        _ = v5
    } else {
        switch "foo bar" {
        case "a":
            if false {
                p0 += 2147483647
                var v6 int64 = p1
                _ = v6
            }
        case "hello":
            v7 := []int{42}
            _ = v7
        case "foo bar":
            v8 := []int{0, 9223372036854775807}
            _ = v8
        default:
            p0 = (p0 ^ p0)
        }
        var v9 int32 = 1
        _ = v9
    }
    var v10 int8 = 0
    _ = p0
    _ = p1
    _ = v10
    return -1
}
func genFn1() int8 {
    v11 := []int{4294967295, 774, -455, 128}
    var v12 uint64 = 1
    var v13 int = 143
    _ = v11
    _ = v12
    _ = v13
    return 49
}

func main() {
    v1 := []int{}
    v2 := make(map[int]bool)
    var v3 string = (("a" + ("\n\t" + ((("" + "foo bar") + "a") + ("foo bar" + ("a" + ""))))) + "a")
    var v14 string = v3
    v3 = (((((v3 + v14) + v14) + v14) + "a") + (v14 + "a"))
    v3 = "foo bar"
    switch 128 {
    case -9223372036854775808:
        for fi15 := 0; fi15 < 13; fi15++ {
            v16 := GenS2{F0: (fi15 | fi15), F1: -997, F2: 0}
            v1 = append(v1, fi15)
            var v17 int32 = -2147483648
            _ = fi15
            _ = v16
            _ = v17
        }
    default:
        for fi18 := 0; fi18 < 11; fi18++ {
            var v19 float32 = 0.0
            v20 := v2[fi18]
            v14 = ((v3 + v3) + "")
            _ = fi18
            _ = v19
            _ = v20
        }
    }
    var v21 int64 = 836
    var v22 int32 = 42
    v2[0] = (((v21 ^ v21) >= 1) || false)
    if (true && (0 <= 42)) {
        if false {
            v23 := GenS1{F0: (v14 + ("hello" + (v14 + ("\n\t" + (v14 + v3))))), F1: 0}
            switch "" {
            case "":
                v21 = (-(v21))
            case "foo bar":
                if (false && ((v21 << ((0) & 63)) > (v21 >> ((2) & 63)))) {
                    if (3.14159 <= 0.0) {
                        v24 := GenS3{F0: 255}
                        _ = v24
                    } else {
                        if (true || false) {
                            if len(v1) > 0 {
                                v1[0] = 4294967295
                            }
                            if (425 >= -1) {
                                var v25 int64 = 439
                                v25 += 1
                                _ = v25
                            } else {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                        }
                    }
                    var v26 uint32 = 302
                    var v27 uint = 0
                    _ = v26
                    _ = v27
                } else {
                    var v28 float32 = float32(float64(1e10))
                    _ = v28
                }
            }
            v21 = v21
            _ = v23
        }
    }
    v29 := GenS4{F0: 1000}
    switch 127 {
    case 1:
        v14 = v3
    default:
        if len(v1) > 0 {
            v1[0] = 9223372036854775807
        }
    }
    v21 = 1
    fmt.Println(v3)
    v21 += 240
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v29
    _ = v3
}
