package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 int16; F2 uint64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 string }
func genFn0(p0 int8, p1 uint32) string {
    if false {
        v7 := []int{-2147483648, 217, 550}
        _ = v7
    }
    v8 := []int{-1}
    _ = p0
    _ = p1
    _ = v8
    return "foo bar"
}
func genFn1(p0 int32, p1 int, p2 bool) uint8 {
    if (!p2) {
        p2 = (4294967295 == 65535)
        if true {
            v9 := []int{-2147483648}
            var v10 int8 = -92
            for fi11 := 0; fi11 < 16; fi11++ {
                for fi12 := 0; fi12 < 1; fi12++ {
                    v13 := GenS1{F0: (p0 & p0)}
                    _ = fi12
                    _ = v13
                }
                p0 = p0
                _ = fi11
            }
            _ = v10
            _ = v9
        }
    } else {
        p1 = p1
        for fi14 := 0; fi14 < 15; fi14++ {
            v15 := []int{(p1 / ((p1) | 1)), 124}
            var v16 int = (-(fi14))
            _ = fi14
            _ = v15
            _ = v16
        }
        v17 := GenS2{F0: 29, F1: 42, F2: 944}
        _ = v17
    }
    var v18 uint8 = 185
    fmt.Println(p1)
    p0 = (p0 * p0)
    var v19 float32 = 0.1
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v19
    return v18
}

func main() {
    v1 := []uint32{4294967295, 65535, 285}
    var v2 uint32 = 2147483648
    var v3 int32 = -1
    var v4 uint32 = (v2 ^ (v2 % ((42) | 1)))
    var v5 uint64 = 626
    var v6 uint8 = 0
    switch 2147483647 {
    case -1:
        fmt.Println(v4)
    case 32767:
        if false {
            var v20 int64 = 0
            _ = v20
        }
    }
    var v21 uint32 = v4
    for fi22 := 0; fi22 < 15; fi22++ {
        switch "" {
        case "\n\t":
            if (!false) {
                for fi23 := 0; fi23 < 17; fi23++ {
                    v5 += 4294967295
                    var v24 float32 = (3.14159 + 3.14159)
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi23
                    _ = v24
                }
                v3 = (v3 >> ((1) & 63))
            } else {
                var v25 int = -1
                _ = v25
            }
        default:
            v4 = 65535
        }
        v6 = (v6 / ((genFn1(v3, 329, false)) | 1))
        if ("foo bar" >= "\n\t") {
            var v26 uint16 = 824
            v3 = -1
            fmt.Println(v4)
            _ = v26
        } else {
            fmt.Println(v3, v21)
            if (false || (32767 == 0)) {
                v2 -= v4
                if (1 > 709) {
                    if false {
                        if (128 == uint8(v5)) {
                            var v27 int8 = 127
                            for fi28 := 0; fi28 < 8; fi28++ {
                                v29 := GenS1{F0: 0}
                                var v30 int16 = 42
                                _ = fi28
                                _ = v29
                                _ = v30
                            }
                            _ = v27
                        } else {
                            var v31 int64 = -9223372036854775808
                            v21 += v2
                            var v32 float32 = (1e10 / 0.1)
                            _ = v31
                            _ = v32
                        }
                    }
                    fmt.Println(v21, v3)
                    var v33 int32 = 158
                    _ = v33
                } else {
                    fmt.Println(v3, v2)
                    v4 += 0
                    var v34 uint32 = v4
                    _ = v34
                }
                for fi35 := 0; fi35 < 15; fi35++ {
                    if ((("" + "a") == ("a" + (("a" + "hello") + ("a" + "")))) && true) {
                        v36 := GenS3{F0: 4294967295}
                        v37 := GenS2{F0: 0, F1: 42, F2: (v5 - (v5 ^ (^(v5))))}
                        _ = v36
                        _ = v37
                    } else {
                        v38 := GenS4{F0: ("foo bar" + ("foo bar" + ("foo bar" + "foo bar")))}
                        var v39 int16 = -735
                        if len(v1) > 0 {
                            v1[0] = 4294967295
                        }
                        _ = v38
                        _ = v39
                    }
                    _ = fi35
                }
            }
        }
        _ = fi22
    }
    if ((1e10 + 3.14159) == (1e10 - 0.1)) {
        for fi40 := 0; fi40 < 13; fi40++ {
            v4 = 1
            if (4294967295 != 849) {
                v41 := GenS1{F0: 1}
                _ = v41
            }
            v3 = int32(fi40)
            _ = fi40
        }
        var v42 string = ("\n\t" + "")
        _ = v42
    }
    var v43 uint8 = (v6 / (((v6 | 158)) | 1))
    v43 = (v43 * genFn1((v3 / (((v3 >> ((164) & 63))) | 1)), 42, (!(!(-462 <= 9223372036854775807)))))
    switch 2147483647 {
    case 0:
        v44 := GenS4{F0: ""}
        _ = v44
    case 4294967295:
        var v45 float64 = 1e10
        _ = v45
    }
    var v46 int64 = int64(v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v43
    _ = v46
    _ = v5
    _ = v6
}
