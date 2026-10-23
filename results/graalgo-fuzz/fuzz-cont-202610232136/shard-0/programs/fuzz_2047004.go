package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string; F1 int16; F2 string }
type GenS4 struct { F0 uint; F1 uint; F2 int16 }
func genFn0(p0 uint8) uint32 {
    p0 = p0
    var v7 float64 = 1e10
    v8 := []int{}
    if true {
        if len(v8) > 0 {
            v8[0] = 127
        }
        p0 = (p0 & (p0 ^ p0))
        fmt.Println(v7, p0)
    }
    _ = p0
    _ = v7
    _ = v8
    return 4294967295
}
func genFn1() int {
    var v9 int16 = 32767
    v9 = -719
    if true {
        v9 -= v9
    } else {
        var v10 string = (("\n\t" + "a") + "a")
        _ = v10
    }
    _ = v9
    return 9223372036854775807
}

func main() {
    var v1 string = ""
    v2 := make(map[int]string)
    v3 := []uint32{2147483647, 980, 0}
    var v4 uint64 = 128
    var v5 int = -204
    var v6 float64 = -2.25
    fmt.Println(v4, v4)
    v11 := GenS1{F0: false, F1: -1, F2: ((v1 + ("\n\t" + v1)) + v1)}
    var v12 bool = false
    fmt.Println(v4, v5, v12)
    var v13 uint32 = 386
    var v14 int32 = -1
    if (v14 > -1) {
        fmt.Println(v12, v5, v1)
    }
    var v15 float32 = 1.5
    if (!v12) {
        for fi16 := 0; fi16 < 6; fi16++ {
            if len(v3) > 0 {
                v3[0] = v13
            }
            if (29 >= 4294967295) {
                v17 := v2[(v5 + v5)]
                var v18 string = "\n\t"
                for fi19 := 0; fi19 < 12; fi19++ {
                    if ((float64(0.1) < 3.14159) || v12) {
                        v20 := GenS1{F0: v12, F1: -273, F2: v17}
                        if (v12 || (float64((((3.14159 / float32(3.14159)) - ((v15 / v15) * 0.1)) + v15)) < 0.1)) {
                            fmt.Println(v14, v1)
                            v6 += -2.25
                        }
                        switch (v18 + ((("\n\t" + v18) + v17) + v18)) {
                        case "":
                            var v21 uint64 = (^(v4))
                            _ = v21
                        default:
                            v1 = "foo bar"
                        }
                        _ = v20
                    }
                    v22 := GenS1{F0: ((v4 | 1) <= (v4 % ((989) | 1))), F1: 42, F2: v1}
                    _ = fi19
                    _ = v22
                }
                _ = v17
                _ = v18
            } else {
                v23 := GenS2{F0: -113}
                for fi24 := 0; fi24 < 14; fi24++ {
                    v6 += ((v6 + 0.0) * v6)
                    if v12 {
                        v4 -= (v4 | (v4 & (v4 << ((uint(v4)) & 63))))
                        v15 += 3.14159
                    }
                    if (2 == 0) {
                        if (false && false) {
                            _ = len(v3)
                            delete(v2, v5)
                        }
                        v25 := GenS2{F0: 0}
                        _ = v25
                    }
                    _ = fi24
                }
                for fi26 := 0; fi26 < 5; fi26++ {
                    v5 -= (fi26 ^ fi16)
                    var v27 uint8 = 128
                    var v28 uint16 = 1
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                _ = v23
            }
            _ = fi16
        }
        var v29 int64 = 0
        _ = v29
    }
    v13 = 0
    var v30 string = v1
    if false {
        for fi31 := 0; fi31 < 7; fi31++ {
            switch -9223372036854775808 {
            case 926:
                switch "" {
                case "a":
                    if len(v3) > 0 {
                        v3[0] = (v13 / ((1) | 1))
                    }
                }
            case 1:
                v15 -= (v15 * (1e10 * 3.14159))
            default:
                for fi32 := 0; fi32 < 17; fi32++ {
                    var v33 string = v1
                    v34 := GenS3{F0: "", F1: -1, F2: ("hello" + v1)}
                    v6 -= v6
                    _ = fi32
                    _ = v33
                    _ = v34
                }
            }
            v35 := GenS1{F0: v12, F1: -359, F2: "hello"}
            var v36 string = (v30 + ("\n\t" + (("\n\t" + "\n\t") + v30)))
            _ = fi31
            _ = v35
            _ = v36
        }
        _ = len(v3)
    } else {
        v37 := GenS4{F0: 42, F1: 183, F2: 950}
        if (0 > v14) {
            fmt.Println(v5)
            if true {
                if (v15 > 1e10) {
                    var v38 int32 = 2147483647
                    _ = v38
                } else {
                    var v39 float64 = 1.5
                    var v40 uint64 = v4
                    _ = v39
                    _ = v40
                }
                var v41 int32 = (v14 | v14)
                fmt.Println(v1)
                _ = v41
            } else {
                switch 32767 {
                case 127:
                    switch v1 {
                    case "\n\t":
                        for fi42 := 0; fi42 < 15; fi42++ {
                            v3 = append(v3, (v13 % ((genFn0(1)) | 1)))
                            _ = fi42
                        }
                    case "a":
                        v30 = ("" + v30)
                    case "foo bar":
                        v3 = append(v3, (v13 - genFn0(128)))
                    default:
                        fmt.Println(v15, v5)
                    }
                }
                if (!((v13 & (v13 % ((v13) | 1))) < genFn0(243))) {
                    switch "a" {
                    case "hello":
                        var v43 int = 14
                        _ = v43
                    }
                }
                if v12 {
                    for fi44 := 0; fi44 < 3; fi44++ {
                        v12 = ("foo bar" == v1)
                        _ = fi44
                    }
                    var v45 float64 = (3.14159 - v6)
                    var v46 int8 = -128
                    _ = v45
                    _ = v46
                }
            }
            delete(v2, (-(v5)))
        } else {
            if ((v12 || (v12 && false)) || v12) {
                fmt.Println(v1)
            } else {
                switch (v5 | v5) {
                case -248:
                    var v47 float64 = (0.1 * 1e10)
                    _ = v47
                default:
                    var v48 uint8 = 42
                    _ = v48
                }
                v30 = "foo bar"
            }
            fmt.Println(v15)
        }
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
