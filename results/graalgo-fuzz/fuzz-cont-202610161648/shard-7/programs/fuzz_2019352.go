package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 int64 }
type GenS2 struct { F0 uint; F1 bool }
func genFn0(p0 int32) bool {
    var v6 int32 = p0
    v6 = (p0 & v6)
    v7 := []int{42, 146, 4294967295}
    _ = p0
    _ = v6
    _ = v7
    return true
}

func main() {
    v1 := []int{205, 2147483647, 42, 2147483647}
    v2 := make(map[uint32]string)
    var v3 string = "\n\t"
    var v4 int16 = 1
    v5 := []uint32{}
    fmt.Println(v4, v3)
    var v8 string = (v3 + v3)
    var v9 uint64 = 817
    v8 = "a"
    for fi10 := 0; fi10 < 18; fi10++ {
        fmt.Println(v9, v9, fi10)
        v11 := GenS1{F0: (((v8 + (v3 + (v3 + "hello"))) + "foo bar") + v8), F1: (0.1 != (1.5 + (3.14159 * 1e10))), F2: 42}
        v8 = ""
        _ = fi10
        _ = v11
    }
    v8 = v3
    v12 := GenS1{F0: v8, F1: true, F2: -1}
    for fi13 := 0; fi13 < 11; fi13++ {
        if (!true) {
            for fi14 := 0; fi14 < 19; fi14++ {
                v3 = ""
                v8 = "\n\t"
                v15 := v2[2147483648]
                fmt.Println(fi13)
                _ = fi14
                _ = v15
            }
            var v16 float32 = 1e10
            var v17 uint32 = 0
            _ = v16
            _ = v17
        }
        switch ("foo bar" + ("a" + ("foo bar" + (((v8 + v8) + v3) + "\n\t")))) {
        case "foo bar":
            var v18 int32 = 42
            _ = v18
        case "":
            if (297 <= 756) {
                for fi19 := 0; fi19 < 1; fi19++ {
                    var v20 int32 = 661
                    v2[42] = v8
                    if (42 > 0) {
                        if (1 > (v9 - (v9 * 297))) {
                            v20 -= 0
                            v3 = v3
                            var v21 int16 = -259
                            _ = v21
                        } else {
                            v22 := GenS1{F0: (((v3 + v8) + "") + "\n\t"), F1: (true && (((0.1 <= 3.14159) || (341 != 4294967295)) && (42 != v9))), F2: -1}
                            v1 = append(v1, fi19)
                            if true {
                                v5 = append(v5, 1)
                                var v23 uint64 = v9
                                v3 = ("foo bar" + v8)
                                v24 := GenS2{F0: 65535, F1: false}
                                _ = v23
                                _ = v24
                            }
                            v25 := GenS1{F0: (v3 + v8), F1: ((1e10 / 1.0) <= -2.25), F2: 42}
                            if false {
                                v4 -= -1
                                v8 = ("foo bar" + v8)
                            }
                            v9 += v9
                            _ = v22
                            _ = v25
                        }
                        v3 = "\n\t"
                        switch fi13 {
                        case 602:
                            fmt.Println(v3, v3, fi13)
                        case -981:
                            v1 = append(v1, (fi13 / ((-9223372036854775808) | 1)))
                        default:
                            _ = len(v2)
                        }
                        var v26 uint = 18446744073709551615
                        v20 -= v20
                        _ = v26
                    }
                    _ = fi19
                    _ = v20
                }
                for fi27 := 0; fi27 < 6; fi27++ {
                    var v28 uint64 = v9
                    if true {
                        fmt.Println(v28, v3, v8)
                        var v29 uint = 1
                        var v30 int16 = v4
                        _ = v29
                        _ = v30
                    }
                    _ = fi27
                    _ = v28
                }
                fmt.Println(fi13, fi13)
            }
        }
        _ = fi13
    }
    v4 = v4
    for fi31 := 0; fi31 < 12; fi31++ {
        var v32 uint = 2
        var v33 int16 = (v4 ^ v4)
        if (true && (((!((765 != -1) || (2 != 42))) || (-128 <= 100)) || true)) {
            v5 = append(v5, 712)
            var v34 float32 = (0.1 - float32(-2.25))
            _ = v34
        }
        _ = fi31
        _ = v32
        _ = v33
    }
    switch 1 {
    case -61:
        switch 8 {
        case 127:
            var v35 int16 = v4
            _ = v35
        }
    case -9223372036854775808:
        switch v8 {
        case "a":
            var v36 int = -2147483648
            _ = v36
        case "hello":
            var v37 uint = 18446744073709551615
            _ = v37
        case "":
            for fi38 := 0; fi38 < 2; fi38++ {
                v4 -= (^(v4))
                var v39 int64 = 9223372036854775807
                _ = fi38
                _ = v39
            }
        }
    case -1:
        switch 128 {
        case 4294967295:
            v9 = (^(v9))
        case -786:
            var v40 float64 = 0.0
            _ = v40
        }
    default:
        delete(v2, 2)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
