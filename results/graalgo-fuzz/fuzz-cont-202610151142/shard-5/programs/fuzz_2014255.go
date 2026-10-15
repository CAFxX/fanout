package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 int32 }
type GenS2 struct { F0 uint8; F1 int; F2 uint64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int; F1 uint16 }
type GenS5 struct { F0 int64; F1 uint64 }
type GenS6 struct { F0 uint64 }
func genFn0(p0 uint8) uint8 {
    var v7 int16 = -32768
    var v8 string = "foo bar"
    var v9 float64 = (-2.25 / 0.1)
    var v10 float32 = (3.14159 - -2.25)
    _ = p0
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return p0
}

func main() {
    var v1 uint8 = 127
    var v2 float64 = float64(1e10)
    var v3 uint32 = 65535
    v4 := make(map[string]uint32)
    v5 := []int{42}
    v6 := make(map[int]int)
    var v11 uint = 1
    if (42 < (^(v3))) {
        if (true || (!true)) {
            fmt.Println(v11, v1, v3)
        }
    }
    if true {
        fmt.Println(v1, v3, v11)
    }
    if ((((("hello" + "\n\t") + ((("" + "foo bar") + "") + "hello")) + "\n\t") + "hello") >= "a") {
        v2 = 1.5
        if (-9223372036854775808 <= 570) {
            v12 := GenS1{F0: false, F1: true, F2: -578}
            for fi13 := 0; fi13 < 7; fi13++ {
                if ((!true) || true) {
                    switch (fi13 * 1) {
                    case -165:
                        v11 += (^(v11))
                    case 127:
                        if (fi13 <= 1) {
                            delete(v6, fi13)
                            for fi14 := 0; fi14 < 6; fi14++ {
                                var v15 uint = (v11 % (((v11 & (v11 * v11))) | 1))
                                if (-128 > -128) {
                                    var v16 int32 = -952
                                    _ = v16
                                }
                                var v17 int64 = int64(fi13)
                                _ = fi14
                                _ = v15
                                _ = v17
                            }
                            fmt.Println(v1)
                        }
                    case 776:
                        v18 := GenS1{F0: (-1 == (-(fi13))), F1: (!true), F2: -2147483648}
                        _ = v18
                    default:
                        if (false || (!true)) {
                            v11 = v11
                            var v19 uint = (v11 << ((2) & 63))
                            if true {
                                if (0.1 == 1.5) {
                                    v20 := GenS1{F0: (!true), F1: (!false), F2: 2147483647}
                                    if (false && false) {
                                        v19 -= (v19 / ((v11) | 1))
                                    } else {
                                        switch fi13 {
                                        case 42:
                                            delete(v4, "\n\t")
                                        case 4294967295:
                                            if (-1 != -32768) {
                                                v2 += v2
                                                for fi21 := 0; fi21 < 6; fi21++ {
                                                    v19 = 537
                                                    _ = fi21
                                                }
                                                if (!false) {
                                                    for fi22 := 0; fi22 < 13; fi22++ {
                                                        v2 += v2
                                                        v11 = v11
                                                        _ = fi22
                                                    }
                                                }
                                            }
                                        }
                                    }
                                    _ = len(v6)
                                    _ = v20
                                } else {
                                    v11 -= v11
                                    v19 += 703
                                }
                            }
                            _ = v19
                        } else {
                            if true {
                                var v23 float64 = 1e10
                                v1 += 128
                                _ = v23
                            } else {
                                var v24 int = fi13
                                v1 += v1
                                _ = v24
                            }
                            v3 = (v3 + (v3 / (((v3 << ((459) & 63))) | 1)))
                            var v25 float32 = 3.14159
                            _ = v25
                        }
                    }
                    v2 = v2
                }
                for fi26 := 0; fi26 < 8; fi26++ {
                    v27 := GenS2{F0: 86, F1: (fi26 >> (((v11 << ((v11) & 63))) & 63)), F2: 174}
                    v28 := GenS2{F0: (v1 | 0), F1: (fi26 / ((-2147483648) | 1)), F2: 104}
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                for fi29 := 0; fi29 < 19; fi29++ {
                    v1 += (v1 + (v1 ^ v1))
                    v3 -= v3
                    _ = fi29
                }
                _ = fi13
            }
            if ((^(v1)) < v1) {
                for fi30 := 0; fi30 < 14; fi30++ {
                    var v31 string = (((("a" + ("a" + ("a" + ("" + ((("a" + "hello") + "a") + "a"))))) + "hello") + "foo bar") + ("\n\t" + ""))
                    _ = fi30
                    _ = v31
                }
                if ((!(1 > -128)) && (81 >= 81)) {
                    v32 := GenS1{F0: true, F1: false, F2: -2147483648}
                    for fi33 := 0; fi33 < 4; fi33++ {
                        v2 -= v2
                        v2 = (v2 + v2)
                        for fi34 := 0; fi34 < 17; fi34++ {
                            v11 += (v11 & uint(v1))
                            var v35 uint = v11
                            _ = fi34
                            _ = v35
                        }
                        _ = fi33
                    }
                    v36 := GenS3{F0: 2}
                    _ = v32
                    _ = v36
                }
            }
            _ = v12
        } else {
            v37 := GenS4{F0: 1, F1: 1}
            var v38 string = "a"
            var v39 bool = (9223372036854775807 == -1)
            _ = v37
            _ = v38
            _ = v39
        }
        v2 += -2.25
    }
    for fi40 := 0; fi40 < 13; fi40++ {
        v3 = 2
        for fi41 := 0; fi41 < 15; fi41++ {
            fmt.Println(v11)
            v3 = (v3 ^ 955)
            v4[""] = 971
            _ = fi41
        }
        if len(v5) > 0 {
            v5[0] = (fi40 + fi40)
        }
        _ = fi40
    }
    v11 -= 255
    v3 = 1
    if true {
        var v42 int = 127
        v43 := GenS5{F0: 1, F1: 4294967295}
        _ = v42
        _ = v43
    }
    v2 -= ((v2 / v2) / 1.0)
    switch "hello" {
    case "\n\t":
        if len(v5) > 0 {
            v5[0] = 42
        }
    case "foo bar":
        if true {
            v44 := GenS6{F0: 2}
            v3 += v3
            var v45 float32 = (0.1 - float32(v2))
            _ = v44
            _ = v45
        } else {
            var v46 int32 = 2147483647
            var v47 uint8 = (v1 % ((genFn0(v1)) | 1))
            for fi48 := 0; fi48 < 9; fi48++ {
                var v49 uint32 = v3
                _ = fi48
                _ = v49
            }
            _ = v46
            _ = v47
        }
    }
    fmt.Println(v11, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
