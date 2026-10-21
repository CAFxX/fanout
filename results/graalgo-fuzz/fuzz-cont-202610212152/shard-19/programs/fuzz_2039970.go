package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 uint32; F2 int32 }
type GenS3 struct { F0 int64; F1 int16; F2 string }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint32 }
func genFn0(p0 uint8, p1 uint64, p2 int32) int8 {
    var v6 float64 = 0.0
    switch (("hello" + "a") + "hello") {
    case "foo bar":
        p1 = (p1 | (p1 - p1))
    default:
        if (!false) {
            if (uint8(p1) >= 2) {
                p1 -= p1
                v6 = ((0.0 + (v6 - v6)) * 0.0)
            } else {
                var v7 uint64 = (p1 | (p1 >> ((4294967295) & 63)))
                _ = v7
            }
            var v8 float32 = 0.0
            if ((p2 ^ 1) >= int32(p0)) {
                for fi9 := 0; fi9 < 9; fi9++ {
                    v10 := GenS1{F0: 2147483648}
                    _ = fi9
                    _ = v10
                }
                v11 := GenS2{F0: -829, F1: 65535, F2: p2}
                p0 = (p0 ^ p0)
                _ = v11
            } else {
                p0 -= 1
                for fi12 := 0; fi12 < 17; fi12++ {
                    switch "foo bar" {
                    case "\n\t":
                        if (!false) {
                            v13 := []int{(fi12 * (fi12 % ((fi12) | 1))), int(p2)}
                            if ("foo bar" == "a") {
                                v8 = (v8 / (v8 / v8))
                                v6 += 1.5
                                v13 = append(v13, fi12)
                            } else {
                                for fi14 := 0; fi14 < 9; fi14++ {
                                    v6 += 1.5
                                    fmt.Println(p1, p1)
                                    fmt.Println(p2, fi14, fi12)
                                    _ = fi14
                                }
                                p1 -= (^(p1))
                            }
                            _ = v13
                        } else {
                            v15 := []int{fi12, fi12, fi12, fi12}
                            if false {
                                var v16 uint64 = (p1 & (p1 - (p1 / ((647) | 1))))
                                v6 = (((v6 + -2.25) - float64(v8)) - 1e10)
                                _ = v16
                            }
                            if len(v15) > 0 {
                                v15[0] = 32767
                            }
                            _ = v15
                        }
                    case "":
                        var v17 int16 = 1
                        _ = v17
                    case "foo bar":
                        fmt.Println(p2, fi12)
                    }
                    var v18 bool = false
                    _ = fi12
                    _ = v18
                }
            }
            _ = v8
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -32
}
func genFn1() int {
    var v19 uint = 2
    v20 := GenS1{F0: 4294967295}
    if ((130 == 42) || (2 > 566)) {
        v21 := []int{2}
        v22 := GenS3{F0: 0, F1: -93, F2: ("a" + ("\n\t" + "\n\t"))}
        _ = v21
        _ = v22
    }
    v23 := GenS4{F0: (v19 << (((^(v19))) & 63))}
    _ = v19
    _ = v20
    _ = v23
    return 32767
}

func main() {
    var v1 int32 = -2147483648
    v2 := []uint32{}
    var v3 int8 = 42
    var v4 uint32 = 1
    var v5 int8 = v3
    _ = len(v2)
    var v24 uint = 1
    var v25 int64 = 343
    fmt.Println(v3, v1, v4)
    v2 = append(v2, (^(v4)))
    switch genFn1() {
    case -121:
        if true {
            v26 := GenS3{F0: (v25 & (v25 ^ v25)), F1: -1, F2: ("\n\t" + ((("a" + ("" + "\n\t")) + ("" + ("a" + ("hello" + ("\n\t" + "a"))))) + "hello"))}
            _ = v26
        } else {
            _ = len(v2)
            _ = len(v2)
        }
    case 551:
        v2 = append(v2, 868)
    default:
        v3 = (^(v5))
    }
    if len(v2) > 0 {
        v2[0] = v4
    }
    var v27 int32 = v1
    var v28 uint8 = 0
    var v29 uint16 = 2
    for fi30 := 0; fi30 < 10; fi30++ {
        for fi31 := 0; fi31 < 2; fi31++ {
            v29 += v29
            fmt.Println(v29, v25, v5)
            _ = fi31
        }
        var v32 float64 = -2.25
        if len(v2) > 0 {
            v2[0] = (v4 & (^(v4)))
        }
        _ = fi30
        _ = v32
    }
    var v33 int = genFn1()
    v34 := GenS5{F0: 4294967295}
    var v35 float32 = (1.5 / 0.1)
    v29 = v29
    v35 -= 3.14159
    v36 := GenS4{F0: (v24 % ((v24) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
