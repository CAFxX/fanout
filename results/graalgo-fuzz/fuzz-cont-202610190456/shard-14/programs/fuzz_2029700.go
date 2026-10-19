package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 int8; F2 int64 }
type GenS3 struct { F0 uint16; F1 bool; F2 uint32 }
type GenS4 struct { F0 bool; F1 int; F2 uint }
func genFn0(p0 uint8, p1 uint, p2 int16) int32 {
    p0 += (^(p0))
    var v7 float32 = (1e10 - 0.1)
    var v8 int16 = p2
    p0 -= p0
    v9 := GenS1{F0: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return -1
}

func main() {
    v1 := []uint32{2147483648}
    var v2 uint64 = 2
    var v3 float32 = ((0.1 * 0.0) / 3.14159)
    v4 := []string{"a"}
    var v5 float32 = 0.1
    var v6 uint64 = (^(v2))
    if true {
        var v10 int32 = -2147483648
        switch 994 {
        case -973:
            v2 = v2
        case 2147483647:
            switch 127 {
            case -1:
                var v11 uint16 = 287
                _ = v11
            default:
                v10 = (v10 - 2147483647)
            }
        default:
            v2 -= 288
        }
        var v12 int32 = (v10 / ((v10) | 1))
        _ = v10
        _ = v12
    } else {
        var v13 bool = (true && (((v2 + 1) != (v6 >> ((625) & 63))) || (0 < 40)))
        v14 := GenS1{F0: 62}
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 1; fi15++ {
        var v16 int = (fi15 >> ((uint(v6)) & 63))
        if true {
            for fi17 := 0; fi17 < 14; fi17++ {
                switch "" {
                case "a":
                    switch (v16 % ((128) | 1)) {
                    case 127:
                        v16 = v16
                    case -1:
                        v16 = (fi17 * -2147483648)
                    case -503:
                        v5 -= v5
                    default:
                        v18 := GenS1{F0: 178}
                        _ = v18
                    }
                case "hello":
                    v19 := GenS2{F0: -926, F1: -113, F2: 9223372036854775807}
                    _ = v19
                default:
                    if false {
                        if len(v4) > 0 {
                            v4[0] = ""
                        }
                        v2 += (v6 | 4294967295)
                    }
                }
                fmt.Println(v16, fi17)
                switch "" {
                case "\n\t":
                    v20 := GenS1{F0: 26}
                    _ = v20
                case "foo bar":
                    var v21 uint32 = 101
                    _ = v21
                case "a":
                    v22 := GenS3{F0: 32767, F1: (true && ((("hello" + "") + "\n\t") <= "a")), F2: 960}
                    _ = v22
                default:
                    v3 -= ((3.14159 - 0.0) - v3)
                }
                if false {
                    v4 = append(v4, "hello")
                    v23 := GenS3{F0: 397, F1: false, F2: 1}
                    _ = v23
                }
                _ = fi17
            }
            if ((fi15 << ((42) & 63)) == (^(v16))) {
                fmt.Println(v2)
                var v24 uint16 = 2
                _ = v24
            } else {
                switch v16 {
                case -643:
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                case 761:
                    if (1e10 != v3) {
                        fmt.Println(v16)
                    }
                case -502:
                    _ = len(v4)
                default:
                    v4 = append(v4, "foo bar")
                }
            }
        } else {
            if (!false) {
                var v25 uint16 = 255
                _ = len(v4)
                v3 = v5
                _ = v25
            }
            for fi26 := 0; fi26 < 15; fi26++ {
                v2 -= (v6 & (v6 << ((371) & 63)))
                _ = fi26
            }
            _ = len(v1)
        }
        if (true && (-1 >= -1)) {
            for fi27 := 0; fi27 < 18; fi27++ {
                var v28 uint8 = 1
                _ = fi27
                _ = v28
            }
            for fi29 := 0; fi29 < 5; fi29++ {
                v5 = v3
                if false {
                    v30 := GenS2{F0: 0, F1: -1, F2: 697}
                    _ = v30
                } else {
                    if ((((-1 < 42) && false) || true) && false) {
                        var v31 int16 = -856
                        var v32 string = ("foo bar" + "foo bar")
                        if (42 != 7) {
                            v33 := GenS3{F0: 42, F1: false, F2: 2147483648}
                            _ = v33
                        }
                        _ = v31
                        _ = v32
                    } else {
                        v34 := GenS4{F0: (-2.25 == 3.14159), F1: fi29, F2: 1}
                        if (!(!(!(true || (true || (!(0 >= 65535))))))) {
                            fmt.Println(v3)
                            v35 := GenS1{F0: 128}
                            _ = v35
                        } else {
                            v4 = append(v4, "a")
                            var v36 int16 = 387
                            _ = v36
                        }
                        _ = v34
                    }
                    if (v6 >= (v2 ^ v2)) {
                        if (true && (1 != 983)) {
                            v37 := GenS3{F0: 32767, F1: (false && true), F2: 2147483648}
                            fmt.Println(v6)
                            v38 := GenS3{F0: 65535, F1: ("" > ""), F2: 1}
                            _ = v37
                            _ = v38
                        }
                        if (!((fi29 / (((v16 ^ fi15)) | 1)) <= 0)) {
                            var v39 string = ""
                            var v40 uint8 = 1
                            _ = v39
                            _ = v40
                        }
                    }
                    v4 = append(v4, (("foo bar" + "foo bar") + (("" + "\n\t") + "")))
                }
                v16 = (fi29 + (^(fi29)))
                switch v16 {
                case 1:
                    var v41 int16 = -916
                    _ = v41
                }
                v5 += float32(0.0)
                _ = fi29
            }
            for fi42 := 0; fi42 < 14; fi42++ {
                var v43 int = fi15
                _ = fi42
                _ = v43
            }
        } else {
            fmt.Println(v2, v6)
            switch (fi15 ^ v16) {
            case 1:
                for fi44 := 0; fi44 < 3; fi44++ {
                    var v45 int8 = 62
                    var v46 int16 = -927
                    if (float32((float64(0.0) / float64((v3 / ((((v5 * v3) * v3) / 1e10) * -2.25))))) >= 0.1) {
                        v2 += (v6 * v6)
                        var v47 int16 = (v46 * -214)
                        if true {
                            fmt.Println(v47, v16)
                            v48 := GenS3{F0: 0, F1: (uint16(v2) == 408), F2: 4294967295}
                            v45 = v45
                            _ = v48
                        }
                        _ = v47
                    } else {
                        v3 = (v3 + -2.25)
                        _ = len(v1)
                        v49 := GenS1{F0: 127}
                        _ = v49
                    }
                    _ = fi44
                    _ = v45
                    _ = v46
                }
            case 42:
                _ = len(v1)
            case 255:
                if false {
                    if len(v1) > 0 {
                        v1[0] = 2147483648
                    }
                    var v50 uint32 = 42
                    _ = v50
                }
            default:
                v16 -= 42
            }
        }
        _ = fi15
        _ = v16
    }
    v3 += 3.14159
    v6 = 2
    var v51 int64 = -9223372036854775808
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v52 float64 = ((1.5 - 0.0) * float64(float32(float64(v5))))
    v53 := GenS1{F0: 0}
    var v54 float64 = float64(float32(1e10))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v6
}
