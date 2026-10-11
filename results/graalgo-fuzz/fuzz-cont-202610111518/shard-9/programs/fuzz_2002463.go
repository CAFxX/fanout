package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 string; F1 int8; F2 int16 }
type GenS3 struct { F0 bool; F1 uint8; F2 uint8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int16; F1 bool }
type GenS6 struct { F0 int16; F1 int64 }

func main() {
    var v1 int64 = -277
    var v2 int32 = int32(v1)
    v3 := []int{-173, -9223372036854775808, 42}
    var v4 int16 = 42
    var v5 int32 = int32(v1)
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 int64 = v1
        _ = fi6
        _ = v7
    }
    if len(v3) > 0 {
        v3[0] = 537
    }
    for fi8 := 0; fi8 < 13; fi8++ {
        var v9 int32 = v2
        if len(v3) > 0 {
            v3[0] = (fi8 + (fi8 % ((fi8) | 1)))
        }
        var v10 uint8 = 42
        _ = fi8
        _ = v10
        _ = v9
    }
    if (((0.1 / 1.0) + -2.25) == 3.14159) {
        v11 := GenS1{F0: v5, F1: 0}
        _ = v11
    }
    v1 -= v1
    var v12 int16 = v4
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if true {
        if true {
            v3 = append(v3, 32768)
            var v13 uint = 768
            v2 = (v2 | -536)
            _ = v13
        } else {
            var v14 uint16 = 162
            switch 167 {
            case 127:
                switch 127 {
                case 128:
                    v5 -= (v2 + (v2 * -2147483648))
                default:
                    v2 = (v2 % ((v5) | 1))
                }
            default:
                var v15 float64 = 0.1
                _ = v15
            }
            _ = v14
        }
        if false {
            v3 = append(v3, 42)
        }
        if ((("\n\t" + "") == "\n\t") && (!false)) {
            if (true || false) {
                v3 = append(v3, 0)
                if len(v3) > 0 {
                    v3[0] = 1
                }
            } else {
                fmt.Println(v5)
                for fi16 := 0; fi16 < 8; fi16++ {
                    v17 := GenS1{F0: (v2 / ((-1) | 1)), F1: 692}
                    switch "" {
                    case "":
                        v18 := GenS2{F0: "a", F1: 127, F2: -1}
                        _ = v18
                    }
                    if (!((0.0 + (0.0 / 1.5)) > -2.25)) {
                        v2 = v2
                    } else {
                        if (true || true) {
                            if false {
                                if (((-2.25 + (0.1 + 3.14159)) + (((3.14159 - float32(((float64((0.1 - (float32(0.0) / -2.25))) * -2.25) + -2.25))) - -2.25) * 1.5)) > (-2.25 / 1.5)) {
                                    var v19 string = "a"
                                    if (0 != int8(v1)) {
                                        var v20 int8 = 127
                                        for fi21 := 0; fi21 < 18; fi21++ {
                                            var v22 uint64 = 414
                                            var v23 int64 = (v1 & v1)
                                            _ = fi21
                                            _ = v22
                                            _ = v23
                                        }
                                        for fi24 := 0; fi24 < 14; fi24++ {
                                            v25 := GenS3{F0: false, F1: 255, F2: 202}
                                            _ = fi24
                                            _ = v25
                                        }
                                        _ = v20
                                    }
                                    v4 -= v12
                                    _ = v19
                                } else {
                                    v3 = append(v3, fi16)
                                }
                            }
                            for fi26 := 0; fi26 < 2; fi26++ {
                                var v27 int16 = -1
                                v2 = v2
                                _ = fi26
                                _ = v27
                            }
                        }
                    }
                    _ = fi16
                    _ = v17
                }
            }
            v5 = 589
        } else {
            switch (("a" + "\n\t") + ("" + "")) {
            case "foo bar":
                if true {
                    v28 := GenS4{F0: 1}
                    _ = v28
                }
            case "a":
                fmt.Println(v1)
            }
            var v29 int8 = 100
            v30 := GenS2{F0: ((("" + "hello") + "hello") + "\n\t"), F1: (v29 + (v29 % ((v29) | 1))), F2: (v4 % ((v4) | 1))}
            _ = v29
            _ = v30
        }
    }
    v31 := GenS4{F0: 0}
    v3 = append(v3, -1)
    if ((42 < 42) && ("a" != "\n\t")) {
        v4 = (-(v12))
    }
    v12 -= (v4 % (((^(v12))) | 1))
    if false {
        if len(v3) > 0 {
            v3[0] = -174
        }
    } else {
        var v32 int8 = 91
        for fi33 := 0; fi33 < 2; fi33++ {
            fmt.Println(v12, v2)
            var v34 float32 = 0.1
            _ = fi33
            _ = v34
        }
        var v35 uint16 = 338
        _ = v32
        _ = v35
    }
    v2 = v2
    if ((v5 << ((42) & 63)) < (-(v2))) {
        v36 := GenS5{F0: 1000, F1: true}
        if (!false) {
            for fi37 := 0; fi37 < 9; fi37++ {
                for fi38 := 0; fi38 < 13; fi38++ {
                    var v39 uint = 86
                    v40 := GenS2{F0: "a", F1: 1, F2: (v4 ^ -597)}
                    _ = fi38
                    _ = v39
                    _ = v40
                }
                _ = fi37
            }
        } else {
            if len(v3) > 0 {
                v3[0] = -771
            }
            v41 := GenS4{F0: 42}
            _ = v41
        }
        v5 = (v2 + (v2 << ((18446744073709551615) & 63)))
        _ = v36
    }
    v42 := GenS6{F0: v12, F1: 9223372036854775807}
    if false {
        var v43 uint8 = 255
        _ = v43
    }
    if (-128 > -128) {
        fmt.Println(v5, v12, v1)
        v5 -= -2147483648
    } else {
        v5 += 2147483647
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v42
    _ = v5
}
