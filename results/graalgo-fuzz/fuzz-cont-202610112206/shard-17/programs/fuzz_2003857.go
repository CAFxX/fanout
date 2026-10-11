package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
func genFn0(p0 int32) string {
    var v4 uint32 = 616
    var v5 bool = (!true)
    _ = p0
    _ = v4
    _ = v5
    return ("a" + ("a" + ("a" + "foo bar")))
}
func genFn1() uint16 {
    var v6 int = 365
    v6 -= (^(v6))
    v6 += v6
    for fi7 := 0; fi7 < 17; fi7++ {
        if ("\n\t" != "foo bar") {
            var v8 int8 = 16
            v9 := []int{32768}
            _ = v8
            _ = v9
        } else {
            v10 := []int{v6, (fi7 % ((v6) | 1)), fi7}
            if true {
                if true {
                    v6 += (^(fi7))
                    var v11 uint = 2
                    _ = v11
                } else {
                    v6 = fi7
                }
                if ((false || false) || false) {
                    var v12 int8 = -128
                    _ = v12
                } else {
                    for fi13 := 0; fi13 < 16; fi13++ {
                        switch "" {
                        case "":
                            v6 += 127
                        case "foo bar":
                            var v14 int = v6
                            _ = v14
                        case "a":
                            v6 -= (-(fi13))
                        }
                        _ = fi13
                    }
                    var v15 uint16 = 2
                    v16 := GenS1{F0: 519, F1: (v15 >> ((2) & 63))}
                    _ = v15
                    _ = v16
                }
                if (((0.0 * -2.25) - (0.1 * 1e10)) >= 0.0) {
                    for fi17 := 0; fi17 < 19; fi17++ {
                        var v18 uint = 4294967295
                        if (4294967295 != 259) {
                            var v19 float64 = float64(float32(0.0))
                            _ = v19
                        }
                        _ = fi17
                        _ = v18
                    }
                    v6 = (v6 << ((0) & 63))
                }
            }
            for fi20 := 0; fi20 < 13; fi20++ {
                if (774 <= 42) {
                    fmt.Println(fi7, fi7)
                    v6 -= 940
                    fmt.Println(fi7, fi7)
                }
                _ = fi20
            }
            _ = v10
        }
        _ = fi7
    }
    var v21 uint8 = 128
    _ = v21
    _ = v6
    return 255
}

func main() {
    var v1 uint64 = 824
    var v2 int16 = 858
    var v3 uint = 4294967295
    v22 := []int{-729, 42, 32767, -9223372036854775808}
    v2 -= v2
    switch "foo bar" {
    case "hello":
        if (816 >= -1) {
            for fi23 := 0; fi23 < 17; fi23++ {
                var v24 float64 = (1.5 + 1e10)
                if false {
                    v24 += float64(0.0)
                    if len(v22) >= 2 {
                        v22 = v22[1:len(v22)]
                    }
                }
                var v25 bool = false
                _ = fi23
                _ = v24
                _ = v25
            }
        }
    case "foo bar":
        var v26 uint = v3
        _ = v26
    case "a":
        var v27 int32 = -783
        _ = v27
    default:
        v28 := GenS1{F0: -1, F1: 102}
        _ = v28
    }
    v2 += (v2 | v2)
    var v29 string = "a"
    v1 = (v1 % (((^(v1))) | 1))
    v3 = 894
    v30 := GenS1{F0: -948, F1: genFn1()}
    for fi31 := 0; fi31 < 11; fi31++ {
        if len(v22) > 0 {
            v22[0] = (fi31 + fi31)
        }
        var v32 uint = v3
        _ = fi31
        _ = v32
    }
    var v33 uint64 = (v1 / ((v1) | 1))
    v3 = (v3 - (v3 + v3))
    v34 := GenS1{F0: -9223372036854775808, F1: uint16(v33)}
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v22
    _ = v29
    _ = v3
    _ = v30
    _ = v33
    _ = v34
}
