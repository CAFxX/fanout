package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 int64; F1 int8; F2 uint }
type GenS3 struct { F0 uint; F1 int; F2 int16 }
type GenS4 struct { F0 int8; F1 int16; F2 uint }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint64 }
func genFn0(p0 uint64, p1 uint8) int8 {
    v7 := []int{2147483647}
    p0 = (p0 + 2)
    for fi8 := 0; fi8 < 7; fi8++ {
        _ = len(v7)
        var v9 uint64 = 782
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: p0, F1: 65535}
    var v11 float64 = 1.5
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v7
    return 127
}

func main() {
    v1 := make(map[string]string)
    var v2 uint64 = 788
    v3 := make(map[string]uint32)
    var v4 float32 = (3.14159 + (-2.25 * 0.0))
    v5 := make(map[uint32]string)
    v6 := []string{"foo bar", "foo bar", "a", ""}
    v12 := GenS2{F0: -9223372036854775808, F1: -51, F2: 2}
    v4 = v4
    for fi13 := 0; fi13 < 1; fi13++ {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        var v14 uint8 = 42
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 7; fi15++ {
        switch "" {
        case "":
            v2 = 2
        case "hello":
            for fi16 := 0; fi16 < 9; fi16++ {
                var v17 uint64 = (v2 + (v2 - 891))
                v2 = v17
                _ = fi16
                _ = v17
            }
        default:
            var v18 uint64 = v2
            _ = v18
        }
        if false {
            var v19 uint64 = (v2 | (v2 | (v2 >> ((867) & 63))))
            v19 = (v19 % ((v19) | 1))
            _ = v19
        } else {
            v6 = append(v6, "")
            if true {
                if false {
                    if ((true || (false && (!false))) || true) {
                        delete(v5, 532)
                        v20 := GenS2{F0: 42, F1: 1, F2: 1}
                        _ = v20
                    }
                } else {
                    _ = len(v3)
                    if ((float32(3.14159) != v4) || true) {
                        v2 = v2
                        v21 := GenS3{F0: 18446744073709551615, F1: (fi15 / ((-312) | 1)), F2: 1000}
                        if true {
                            var v22 float64 = float64((((1.5 / ((3.14159 - v4) + v4)) * (v4 * v4)) + 1e10))
                            var v23 uint = uint(v2)
                            _ = v22
                            _ = v23
                        } else {
                            v2 -= v2
                            if false {
                                var v24 string = "\n\t"
                                _ = v24
                            } else {
                                if true {
                                    v2 += 2
                                    if (840 < 0) {
                                        if (!true) {
                                            v25 := GenS4{F0: 100, F1: 55, F2: 18446744073709551615}
                                            _ = v25
                                        } else {
                                            var v26 bool = (629 < 18446744073709551615)
                                            for fi27 := 0; fi27 < 14; fi27++ {
                                                if v26 {
                                                    v4 += 1.5
                                                    var v28 bool = v26
                                                    _ = v28
                                                }
                                                _ = fi27
                                            }
                                            var v29 int8 = 42
                                            _ = v26
                                            _ = v29
                                        }
                                        v30 := GenS5{F0: ""}
                                        v2 = 0
                                        _ = v30
                                    }
                                    v1["\n\t"] = ("foo bar" + "")
                                }
                            }
                            v2 -= v2
                        }
                        _ = v21
                    } else {
                        v4 += v4
                    }
                }
                v4 = (float32(((float64(v4) * ((-2.25 + 0.0) + (-2.25 / 1.0))) / 0.1)) * (v4 / 1e10))
            } else {
                v2 += 1
            }
        }
        v2 -= v2
        _ = fi15
    }
    switch (("foo bar" + ("foo bar" + ("foo bar" + "a"))) + "foo bar") {
    case "\n\t":
        v4 = ((v4 - 1e10) / v4)
    }
    var v31 uint16 = 676
    v4 = v4
    var v32 int64 = 42
    v4 += v4
    if true {
        switch "" {
        case "foo bar":
            v4 = v4
        case "":
            for fi33 := 0; fi33 < 8; fi33++ {
                v31 -= v31
                _ = fi33
            }
        default:
            v34 := GenS3{F0: 584, F1: -1, F2: 550}
            _ = v34
        }
        for fi35 := 0; fi35 < 2; fi35++ {
            var v36 float32 = ((v4 - ((((v4 + (v4 / ((v4 * 1e10) + 0.0))) - v4) * v4) * 1.5)) - 3.14159)
            if (true || (v2 >= v2)) {
                var v37 uint64 = (v2 ^ v2)
                _ = v37
            }
            for fi38 := 0; fi38 < 10; fi38++ {
                var v39 string = "\n\t"
                fmt.Println(v36, v39, v2)
                fmt.Println(v39)
                _ = fi38
                _ = v39
            }
            _ = fi35
            _ = v36
        }
    } else {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    v40 := GenS4{F0: 127, F1: int16(v2), F2: 255}
    if false {
        v31 = (v31 | (v31 * 255))
    }
    v41 := GenS6{F0: 714}
    v31 = (v31 * 32767)
    var v42 uint16 = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
