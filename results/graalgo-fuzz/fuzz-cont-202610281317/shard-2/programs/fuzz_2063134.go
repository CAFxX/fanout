package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 uint16; F2 int16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int8; F1 bool; F2 uint16 }
type GenS5 struct { F0 int8; F1 int8 }
type GenS6 struct { F0 int32; F1 int }
type GenS7 struct { F0 int16; F1 uint32; F2 bool }
type GenS8 struct { F0 uint32 }
func genFn0() uint32 {
    var v6 bool = (4294967295 == 1)
    var v7 uint64 = 605
    v7 = v7
    _ = v6
    _ = v7
    return 2
}
func genFn1(p0 int16) uint16 {
    v8 := []int{-675}
    v9 := GenS1{F0: 669}
    _ = len(v8)
    _ = p0
    _ = v8
    _ = v9
    return 83
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]string)
    var v3 int64 = -9223372036854775808
    var v4 float64 = 3.14159
    var v5 bool = (0.1 != 3.14159)
    v5 = v5
    fmt.Println(v3, v5, v5)
    for fi10 := 0; fi10 < 12; fi10++ {
        v11 := GenS2{F0: 42, F1: genFn1(829), F2: -1}
        v4 = v4
        _ = fi10
        _ = v11
    }
    if (!true) {
        if (int(v3) <= 127) {
            v12 := GenS2{F0: -2147483648, F1: 2, F2: 0}
            _ = v12
        } else {
            switch "" {
            case "":
                if ((100 < -1) || true) {
                    for fi13 := 0; fi13 < 20; fi13++ {
                        if (475 >= 2147483647) {
                            fmt.Println(fi13, v3)
                            var v14 uint16 = genFn1(42)
                            for fi15 := 0; fi15 < 11; fi15++ {
                                switch (fi13 + (fi13 * fi15)) {
                                case 42:
                                    v16 := GenS1{F0: int32(v14)}
                                    _ = v16
                                case -755:
                                    if v5 {
                                        var v17 uint64 = 4294967295
                                        if (-2.25 >= (0.0 * float32(((v4 * (((1.5 * 1e10) * ((1.5 + ((v4 + v4) / v4)) - v4)) / ((1.5 / (-2.25 * v4)) + v4))) / (v4 / v4))))) {
                                            var v18 uint8 = 112
                                            _ = v18
                                        } else {
                                            v19 := []int{fi13, (^(fi15)), (fi13 + -2147483648), (fi13 ^ (fi15 - (fi13 >> ((uint(v17)) & 63))))}
                                            _ = v19
                                        }
                                        var v20 uint8 = 81
                                        var v21 uint8 = v20
                                        _ = v17
                                        _ = v20
                                        _ = v21
                                    } else {
                                        v14 = v14
                                        if v5 {
                                            v22 := []int{fi13, 128, fi13, (-(fi13))}
                                            v3 += v3
                                            if ("foo bar" != (("\n\t" + "hello") + "")) {
                                                v23 := GenS1{F0: 1}
                                                var v24 uint32 = 65535
                                                _ = v23
                                                _ = v24
                                            }
                                            _ = v22
                                        }
                                    }
                                default:
                                    var v25 string = "a"
                                    _ = v25
                                }
                                switch (fi13 - fi15) {
                                case -577:
                                    v4 += -2.25
                                    if ((3.14159 - 3.14159) > (1e10 + -2.25)) {
                                        var v26 int32 = 2147483647
                                        v1["foo bar"] = "\n\t"
                                        v27 := GenS3{F0: 2}
                                        _ = v26
                                        _ = v27
                                    } else {
                                        v14 += (^(v14))
                                        var v28 int16 = 32767
                                        _ = v28
                                    }
                                default:
                                    v29 := []int{fi13, fi15, fi13}
                                    _ = v29
                                }
                                v30 := GenS4{F0: 1, F1: ((v3 << ((42) & 63)) <= (v3 + v3)), F2: (v14 % ((v14) | 1))}
                                v31 := GenS1{F0: -864}
                                _ = fi15
                                _ = v30
                                _ = v31
                            }
                            _ = v14
                        }
                        _ = len(v2)
                        v3 = (v3 / ((-9223372036854775808) | 1))
                        _ = fi13
                    }
                    fmt.Println(v5)
                }
            case "foo bar":
                v3 = v3
            case "hello":
                var v32 uint32 = 1
                _ = v32
            default:
                v33 := GenS5{F0: 0, F1: -1}
                _ = v33
            }
        }
        var v34 string = "\n\t"
        _ = v34
    } else {
        v35 := GenS3{F0: 2}
        var v36 string = "foo bar"
        v3 -= (v3 + v3)
        _ = v35
        _ = v36
    }
    if (!false) {
        v37 := GenS6{F0: -355, F1: -2147483648}
        for fi38 := 0; fi38 < 17; fi38++ {
            fmt.Println(v3, v3)
            v3 += -9223372036854775808
            _ = fi38
        }
        _ = v37
    } else {
        v39 := GenS6{F0: 1, F1: -506}
        _ = v39
    }
    v40 := GenS7{F0: 719, F1: 42, F2: (1 != 0)}
    v41 := GenS7{F0: 42, F1: 2, F2: true}
    var v42 uint = 0
    v3 += (^(v3))
    switch (("hello" + ("" + "\n\t")) + ((("a" + "foo bar") + ("" + "a")) + "a")) {
    case "\n\t":
        v43 := GenS4{F0: -1, F1: (v42 < 4294967295), F2: 1}
        _ = v43
    default:
        for fi44 := 0; fi44 < 11; fi44++ {
            fmt.Println(v42, v3, v5)
            _ = fi44
        }
    }
    v4 -= (v4 + float64(1.5))
    if false {
        v45 := v1["foo bar"]
        _ = v45
    } else {
        v5 = v5
        if (0 > 18446744073709551615) {
            for fi46 := 0; fi46 < 14; fi46++ {
                v5 = ((^(fi46)) != fi46)
                for fi47 := 0; fi47 < 15; fi47++ {
                    var v48 float64 = v4
                    _ = fi47
                    _ = v48
                }
                _ = fi46
            }
            v49 := GenS8{F0: genFn0()}
            _ = v49
        } else {
            if (v4 == v4) {
                var v50 float64 = ((v4 - ((v4 + -2.25) + -2.25)) + 0.0)
                v2[("foo bar" + "a")] = "a"
                _ = v50
            } else {
                var v51 int32 = 0
                var v52 float32 = 3.14159
                _ = v51
                _ = v52
            }
            switch ("hello" + "") {
            case "hello":
                v53 := GenS7{F0: -877, F1: 0, F2: ((42 >= 447) && (!false))}
                _ = v53
            case "\n\t":
                switch -2147483648 {
                case 32768:
                    v54 := []int{}
                    _ = v54
                default:
                    v55 := v2["foo bar"]
                    _ = v55
                }
            default:
                v42 = v42
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
}
