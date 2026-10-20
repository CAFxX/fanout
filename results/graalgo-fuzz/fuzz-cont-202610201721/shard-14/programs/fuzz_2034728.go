package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 uint16; F2 uint }
type GenS4 struct { F0 int8; F1 int16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 string; F1 int32 }
type GenS7 struct { F0 uint32; F1 uint16; F2 int16 }
func genFn0() uint8 {
    var v4 uint8 = 1
    v4 += (v4 % ((v4) | 1))
    _ = v4
    return 2
}
func genFn1(p0 uint8, p1 uint32, p2 int16) int {
    var v5 uint8 = (^(p0))
    p0 -= (v5 & p0)
    p2 -= (p2 + p2)
    fmt.Println(p0, p0, p1)
    if true {
        switch ((((("" + "hello") + (("foo bar" + ("foo bar" + "\n\t")) + "hello")) + "hello") + ("hello" + "a")) + "a") {
        case "a":
            if (false || (!((p2 - (^(p2))) != (p2 % ((p2) | 1))))) {
                v6 := GenS1{F0: 42, F1: 9223372036854775807, F2: (!((false || (((float32((0.0 * 1e10)) + 0.1) > 3.14159) && ((0.1 > -2.25) && (false || (((255 != v5) && false) && (410 <= 2147483647)))))) && (42 >= 42)))}
                var v7 int8 = 127
                switch "" {
                case "":
                    var v8 int16 = -500
                    _ = v8
                case "\n\t":
                    fmt.Println(v7)
                case "foo bar":
                    var v9 int64 = -9223372036854775808
                    _ = v9
                }
                _ = v6
                _ = v7
            }
        }
        v5 = (v5 % (((^(v5))) | 1))
    } else {
        v10 := GenS2{F0: -128}
        fmt.Println(v5, p0, p1)
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 197
}

func main() {
    var v1 int32 = 728
    var v2 bool = (-329 != -1)
    v3 := make(map[string]bool)
    if false {
        fmt.Println(v1, v1)
        switch 42 {
        case -797:
            fmt.Println(v2, v2)
        case 127:
            var v11 int32 = (v1 + v1)
            _ = v11
        }
    } else {
        fmt.Println(v1, v2)
        for fi12 := 0; fi12 < 19; fi12++ {
            if (v2 || v2) {
                if v2 {
                    v2 = v2
                }
            } else {
                for fi13 := 0; fi13 < 9; fi13++ {
                    if ((1.5 - (-2.25 + 0.0)) > 3.14159) {
                        v2 = (genFn1(190, 4294967295, -32768) == (^(fi13)))
                        v14 := GenS3{F0: (fi13 + fi13), F1: 543, F2: 2}
                        _ = v14
                    }
                    var v15 int = 128
                    _ = fi13
                    _ = v15
                }
            }
            _ = fi12
        }
        v16 := GenS1{F0: 61, F1: 0, F2: (float32(0.1) > 1e10)}
        _ = v16
    }
    var v17 uint = 42
    if (v1 == 2147483647) {
        v1 -= (v1 / ((385) | 1))
        var v18 bool = v2
        _ = v18
    }
    v19 := GenS4{F0: 117, F1: 42}
    v20 := []int{1}
    v17 += 4294967295
    if true {
        v21 := v3[("" + ((("\n\t" + "a") + "a") + "\n\t"))]
        _ = v21
    } else {
        v1 = (v1 * v1)
        var v22 uint64 = 4294967295
        _ = v22
    }
    if (("foo bar" + "\n\t") >= ("a" + "hello")) {
        v2 = (v2 && (v17 == (v17 * 0)))
    }
    for fi23 := 0; fi23 < 9; fi23++ {
        for fi24 := 0; fi24 < 14; fi24++ {
            var v25 int = 405
            var v26 bool = (v17 >= 625)
            v27 := GenS4{F0: 127, F1: 53}
            _ = fi24
            _ = v25
            _ = v26
            _ = v27
        }
        if ((!v2) || ((((1e10 - 1.5) / 1e10) * 0.1) <= ((0.1 / -2.25) / 1.5))) {
            var v28 float64 = -2.25
            for fi29 := 0; fi29 < 13; fi29++ {
                switch "" {
                case "":
                    switch (-(fi29)) {
                    case 923:
                        var v30 uint = 65535
                        _ = v30
                    case -904:
                        v17 -= v17
                    }
                    v17 = 709
                    v17 += (v17 >> (((v17 % (((v17 / (((v17 | 1)) | 1))) | 1))) & 63))
                    v2 = false
                    v2 = v2
                default:
                    if false {
                        v28 += v28
                    } else {
                        if v2 {
                            v2 = v2
                            if ((v2 || ((v28 * (((v28 * (v28 / (((v28 + (v28 / 1.0)) * v28) / v28))) / v28) / 1.0)) >= v28)) && (-1 != -128)) {
                                v31 := v3["\n\t"]
                                _ = v31
                            } else {
                                var v32 int8 = 1
                                var v33 int = (fi29 & 32767)
                                if (127 < genFn0()) {
                                    fmt.Println(v33)
                                }
                                _ = v32
                                _ = v33
                            }
                        } else {
                            var v34 string = "foo bar"
                            _ = v34
                        }
                        if (2147483648 <= 0) {
                            v1 = (v1 / (((v1 - 2147483647)) | 1))
                            v1 = v1
                        }
                        v35 := GenS5{F0: genFn1(0, 292, 42)}
                        _ = v35
                    }
                }
                _ = fi29
            }
            var v36 uint8 = 169
            _ = v28
            _ = v36
        } else {
            v2 = false
        }
        switch (^(fi23)) {
        case 986:
            v37 := GenS6{F0: ("\n\t" + "foo bar"), F1: 1}
            _ = v37
        case 292:
            v17 = (v17 + v17)
        case 1:
            var v38 float64 = ((3.14159 - (1e10 + 1e10)) / (0.1 + 1.5))
            _ = v38
        default:
            v20 = append(v20, (fi23 & genFn1(9, 0, 76)))
        }
        _ = fi23
    }
    v39 := GenS2{F0: -128}
    if (!v2) {
        var v40 uint16 = 275
        var v41 uint8 = 68
        _ = v40
        _ = v41
    }
    v42 := GenS7{F0: 110, F1: 719, F2: 0}
    v3["\n\t"] = false
    fmt.Println(v2, v17)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v39
    _ = v42
}
