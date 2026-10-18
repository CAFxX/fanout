package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 uint8; F2 int }
type GenS3 struct { F0 int64; F1 uint64; F2 int16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int16; F1 int; F2 bool }
func genFn0(p0 uint16, p1 int16) int16 {
    fmt.Println(p0, p1)
    for fi6 := 0; fi6 < 12; fi6++ {
        v7 := GenS1{F0: 0}
        p1 = p1
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    return (p1 + 201)
}
func genFn1(p0 int64, p1 int8, p2 uint) string {
    var v8 float32 = 0.0
    var v9 bool = false
    switch 204 {
    case -209:
        var v10 uint64 = 2
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return "foo bar"
}

func main() {
    v1 := make(map[string]int)
    v2 := []int{0, 4294967295, -1, 696}
    v3 := make(map[string]string)
    var v4 int16 = 0
    v5 := make(map[int]uint32)
    var v11 uint16 = 255
    if len(v2) > 0 {
        v2[0] = 32767
    }
    fmt.Println(v4, v4)
    if (-64 <= 100) {
        v12 := GenS2{F0: "", F1: 127, F2: 929}
        for fi13 := 0; fi13 < 20; fi13++ {
            if false {
                for fi14 := 0; fi14 < 1; fi14++ {
                    var v15 bool = true
                    v4 += 1
                    _ = fi14
                    _ = v15
                }
                var v16 uint = 4294967295
                _ = v16
            } else {
                for fi17 := 0; fi17 < 6; fi17++ {
                    var v18 int8 = 127
                    switch "\n\t" {
                    case "\n\t":
                        if (float64(-2.25) <= 1e10) {
                            var v19 int32 = 1
                            switch (fi17 ^ (fi17 % ((433) | 1))) {
                            case 0:
                                var v20 uint8 = 166
                                _ = v20
                            default:
                                if (0 >= 0) {
                                    switch 957 {
                                    case 1:
                                        fmt.Println(v19, v18, v18)
                                    case 128:
                                        var v21 int64 = -1
                                        _ = v21
                                    default:
                                        v4 = (^(v4))
                                    }
                                    if true {
                                        switch (fi13 << ((65535) & 63)) {
                                        case 128:
                                            _ = len(v2)
                                        }
                                        var v22 uint32 = 42
                                        _ = v22
                                    }
                                    v4 += (^(v4))
                                } else {
                                    if (!(!false)) {
                                        fmt.Println(v18, v11, fi17)
                                        v23 := GenS2{F0: ("" + "a"), F1: 1, F2: fi13}
                                        var v24 uint = 255
                                        _ = v23
                                        _ = v24
                                    }
                                    v18 = (v18 + int8(v19))
                                }
                            }
                            var v25 int8 = (v18 << ((65535) & 63))
                            _ = v19
                            _ = v25
                        }
                        if (true || (true && false)) {
                            v4 += (v4 & genFn0(862, 42))
                            v2 = append(v2, int(v18))
                            v11 += 255
                        } else {
                            if (233 > 2147483647) {
                                delete(v1, "")
                                var v26 uint8 = 42
                                if ((-1 == -681) || false) {
                                    if len(v2) > 0 {
                                        v2[0] = fi17
                                    }
                                    v18 = (v18 * v18)
                                    v27 := GenS2{F0: "foo bar", F1: v26, F2: fi13}
                                    _ = v27
                                } else {
                                    v28 := GenS2{F0: "\n\t", F1: (v26 % (((v26 >> ((4294967295) & 63))) | 1)), F2: (fi17 * (fi13 - fi13))}
                                    v29 := GenS3{F0: 136, F1: 18446744073709551615, F2: (^(v4))}
                                    switch (fi17 ^ fi17) {
                                    case -1:
                                        v30 := GenS1{F0: -1}
                                        _ = v30
                                    case 350:
                                        var v31 int32 = 193
                                        var v32 uint = 18
                                        fmt.Println(v32, v32)
                                        _ = v31
                                        _ = v32
                                    default:
                                        v1[("\n\t" + "foo bar")] = fi17
                                    }
                                    _ = v28
                                    _ = v29
                                }
                                _ = v26
                            } else {
                                var v33 float64 = 3.14159
                                v33 = (v33 * ((v33 * -2.25) / 1.5))
                                var v34 uint64 = 42
                                v4 += v4
                                _ = v33
                                _ = v34
                            }
                            var v35 int8 = (v18 + (v18 >> ((866) & 63)))
                            _ = v35
                        }
                        if len(v2) > 0 {
                            v2[0] = fi17
                        }
                    case "hello":
                        var v36 int16 = genFn0(v11, -32768)
                        _ = v36
                    case "foo bar":
                        v37 := GenS3{F0: int64(v11), F1: 4294967295, F2: v4}
                        _ = v37
                    default:
                        v38 := GenS4{F0: 1}
                        _ = v38
                    }
                    _ = fi17
                    _ = v18
                }
                var v39 int32 = -574
                _ = v39
            }
            v4 += v4
            var v40 uint32 = 0
            _ = fi13
            _ = v40
        }
        _ = v12
    } else {
        var v41 int8 = 100
        _ = v41
    }
    for fi42 := 0; fi42 < 12; fi42++ {
        var v43 uint64 = 0
        _ = fi42
        _ = v43
    }
    switch "hello" {
    case "a":
        v11 -= (v11 % ((v11) | 1))
    }
    for fi44 := 0; fi44 < 4; fi44++ {
        var v45 uint = 1
        _ = fi44
        _ = v45
    }
    fmt.Println(v11, v4)
    v11 += (v11 << ((255) & 63))
    v11 -= (v11 << ((298) & 63))
    for fi46 := 0; fi46 < 10; fi46++ {
        var v47 uint32 = 2147483647
        v5[(fi46 * (fi46 >> ((536) & 63)))] = v47
        if (405 <= 0) {
            for fi48 := 0; fi48 < 10; fi48++ {
                var v49 uint = 2
                _ = fi48
                _ = v49
            }
        } else {
            if false {
                _ = len(v2)
                v11 = 1
                for fi50 := 0; fi50 < 10; fi50++ {
                    fmt.Println(v47)
                    _ = fi50
                }
            } else {
                v51 := GenS3{F0: -1, F1: 852, F2: genFn0((v11 >> ((65535) & 63)), -600)}
                var v52 float32 = (-2.25 / ((0.1 * 3.14159) / 0.1))
                _ = v51
                _ = v52
            }
            var v53 float64 = float64(0.1)
            _ = v53
        }
        _ = fi46
        _ = v47
    }
    if ((1.5 >= 3.14159) || (!false)) {
        v4 = genFn0(v11, v4)
    }
    v54 := GenS3{F0: 1, F1: 2, F2: v4}
    var v55 int64 = 770
    for fi56 := 0; fi56 < 19; fi56++ {
        v57 := GenS5{F0: v4, F1: fi56, F2: (!true)}
        _ = fi56
        _ = v57
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v55:", v55)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v55
}
