package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int }
type GenS5 struct { F0 bool; F1 uint; F2 uint32 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int32; F1 uint32; F2 int64 }
type GenS8 struct { F0 int; F1 uint }
type GenS9 struct { F0 uint16; F1 uint16 }
type GenS10 struct { F0 int; F1 bool; F2 bool }
func genFn0(p0 string) string {
    v4 := []int{-475, 2147483647}
    fmt.Println(p0)
    _ = p0
    _ = v4
    return p0
}

func main() {
    var v1 float32 = -2.25
    var v2 uint64 = 0
    var v3 int16 = 1000
    for fi5 := 0; fi5 < 5; fi5++ {
        v6 := GenS1{F0: -128, F1: 2}
        _ = fi5
        _ = v6
    }
    fmt.Println(v2)
    v7 := []int{}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    fmt.Println(v2)
    for fi8 := 0; fi8 < 4; fi8++ {
        fmt.Println(v1, fi8, v1)
        _ = fi8
    }
    if true {
        var v9 uint8 = 2
        fmt.Println(v2)
        _ = v9
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        v1 += 3.14159
        for fi11 := 0; fi11 < 8; fi11++ {
            var v12 int = fi10
            _ = fi11
            _ = v12
        }
        v13 := GenS2{F0: v3, F1: 2}
        _ = fi10
        _ = v13
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        if len(v7) > 0 {
            v7[0] = (fi14 >> ((537) & 63))
        }
        _ = fi14
    }
    v2 = v2
    if true {
        var v15 int8 = 1
        _ = v15
    }
    v7 = append(v7, 4294967295)
    if true {
        v1 = (v1 * v1)
    } else {
        for fi16 := 0; fi16 < 8; fi16++ {
            for fi17 := 0; fi17 < 3; fi17++ {
                var v18 uint64 = (v2 * 1)
                switch "foo bar" {
                case "a":
                    fmt.Println(fi16, fi16, fi16)
                case "foo bar":
                    if false {
                        _ = len(v7)
                    } else {
                        if false {
                            var v19 int64 = 253
                            _ = len(v7)
                            _ = v19
                        } else {
                            v1 = (v1 - 3.14159)
                            v20 := GenS1{F0: 1, F1: 42}
                            if (int(v3) > (fi17 - fi17)) {
                                v21 := GenS3{F0: 1}
                                _ = v21
                            } else {
                                if (867 > v18) {
                                    var v22 uint32 = 393
                                    _ = v22
                                }
                            }
                            _ = v20
                        }
                        if false {
                            fmt.Println(fi16, fi16)
                            v3 -= v3
                        }
                        fmt.Println(v3, v18)
                        if (42 != 65535) {
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                            if ((v1 + v1) < (v1 - -2.25)) {
                                switch fi16 {
                                case 2147483647:
                                    v2 = v18
                                case 127:
                                    _ = len(v7)
                                case -969:
                                    v18 = (^(v2))
                                default:
                                    if (("foo bar" + "a") > "foo bar") {
                                        var v23 bool = true
                                        _ = v23
                                    } else {
                                        var v24 uint64 = v2
                                        var v25 uint = 957
                                        var v26 int16 = 32767
                                        var v27 float32 = 0.1
                                        if ((fi17 ^ fi16) > fi16) {
                                            v28 := GenS4{F0: (fi16 / ((fi16) | 1))}
                                            v3 = 32767
                                            var v29 float64 = float64((v27 + v1))
                                            _ = v28
                                            _ = v29
                                        } else {
                                            switch (fi17 | int(v25)) {
                                            case 250:
                                                v30 := GenS1{F0: -49, F1: 1}
                                                _ = v30
                                            case 128:
                                                var v31 float64 = 0.1
                                                _ = v31
                                            case -707:
                                                v2 = (v18 ^ (v24 + 18446744073709551615))
                                            default:
                                                if false {
                                                    _ = len(v7)
                                                    _ = len(v7)
                                                } else {
                                                    v3 = int16(v25)
                                                }
                                            }
                                        }
                                        v26 += (v3 * v3)
                                        _ = v24
                                        _ = v25
                                        _ = v26
                                        _ = v27
                                    }
                                }
                                var v32 float32 = v1
                                _ = v32
                            }
                            v7 = append(v7, int(v2))
                        } else {
                            fmt.Println(fi17)
                        }
                    }
                case "\n\t":
                    fmt.Println(v1, fi17, v1)
                }
                if ((v1 / v1) >= (float32(float64(1e10)) / v1)) {
                    if false {
                        switch "foo bar" {
                        case "\n\t":
                            var v33 int = (fi16 + (fi16 + -9223372036854775808))
                            _ = v33
                        case "foo bar":
                            var v34 uint32 = 65535
                            _ = v34
                        }
                        v35 := GenS3{F0: 18446744073709551615}
                        if true {
                            if len(v7) > 0 {
                                v7[0] = fi16
                            }
                            v3 = -209
                            v36 := GenS5{F0: false, F1: 530, F2: 42}
                            _ = v36
                        } else {
                            switch "" {
                            case "a":
                                if false {
                                    v3 = v3
                                }
                            case "":
                                fmt.Println(v1, fi16, fi16)
                            default:
                                v37 := GenS2{F0: (v3 % (((v3 - (v3 % (((-(v3))) | 1)))) | 1)), F1: 266}
                                _ = v37
                            }
                        }
                        _ = v35
                    } else {
                        _ = len(v7)
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                    }
                    v7 = append(v7, fi16)
                    v38 := GenS4{F0: fi17}
                    if ((fi17 | (^(fi17))) <= fi17) {
                        if false {
                            var v39 int64 = -9223372036854775808
                            fmt.Println(v39, v3)
                            _ = v39
                        } else {
                            var v40 int8 = -1
                            var v41 uint64 = (v18 / ((2) | 1))
                            switch fi16 {
                            case 9223372036854775807:
                                if false {
                                    if (4294967295 <= 48) {
                                        var v42 int64 = 42
                                        switch (fi16 >> ((4294967295) & 63)) {
                                        case 32768:
                                            if false {
                                                v42 = 1
                                                var v43 int32 = -1
                                                _ = v43
                                            }
                                        case 0:
                                            v44 := GenS6{F0: 532}
                                            _ = v44
                                        case 127:
                                            var v45 float32 = ((v1 + v1) * 0.1)
                                            _ = v45
                                        }
                                        v46 := GenS5{F0: true, F1: 42, F2: 4294967295}
                                        _ = v42
                                        _ = v46
                                    } else {
                                        var v47 int8 = (v40 | v40)
                                        switch "" {
                                        case "hello":
                                            var v48 uint8 = 0
                                            _ = v48
                                        }
                                        v7 = append(v7, (fi17 ^ (fi16 >> ((18446744073709551615) & 63))))
                                        _ = v47
                                    }
                                    if (false && false) {
                                        if (!(-2147483648 == -2147483648)) {
                                            v49 := GenS7{F0: 0, F1: 2147483647, F2: 1}
                                            v41 += v2
                                            _ = v49
                                        } else {
                                            if (!(true || false)) {
                                                var v50 int8 = v40
                                                _ = v50
                                            }
                                            var v51 int8 = v40
                                            _ = v51
                                        }
                                        if (true || false) {
                                            v40 -= v40
                                            switch (fi16 / ((fi16) | 1)) {
                                            case 1:
                                                fmt.Println(v40, fi17, v1)
                                            case 9223372036854775807:
                                                v52 := GenS8{F0: (^(fi17)), F1: 714}
                                                v40 += v40
                                                fmt.Println(v41, v41)
                                                _ = v52
                                            default:
                                                if false {
                                                    var v53 int = (fi17 - fi16)
                                                    fmt.Println(v2)
                                                    if len(v7) > 0 {
                                                        v7[0] = 32767
                                                    }
                                                    v54 := GenS6{F0: 236}
                                                    _ = v53
                                                    _ = v54
                                                } else {
                                                    if (!(((fi17 << ((696) & 63)) != -2147483648) || true)) {
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            if len(v7) > 0 {
                                                                v7[0] = fi16
                                                            }
                                                        }
                                                        switch fi17 {
                                                        case -533:
                                                            if (65535 > 0) {
                                                                if (0 != 65535) {
                                                                    _ = len(v7)
                                                                }
                                                                var v55 string = "\n\t"
                                                                v41 = 2
                                                                _ = v55
                                                            }
                                                        }
                                                        if len(v7) > 0 {
                                                            v7[0] = (fi17 | (fi17 - (^(fi16))))
                                                        }
                                                    } else {
                                                        var v56 float32 = v1
                                                        _ = v56
                                                    }
                                                    if len(v7) > 0 {
                                                        v7[0] = (fi17 % ((int(v40)) | 1))
                                                    }
                                                }
                                            }
                                            var v57 int8 = -128
                                            var v58 uint32 = 344
                                            v7 = append(v7, fi16)
                                            var v59 uint8 = 0
                                            _ = v57
                                            _ = v58
                                            _ = v59
                                        }
                                    } else {
                                        _ = len(v7)
                                        v18 = 42
                                        v60 := GenS9{F0: 1, F1: 255}
                                        _ = v60
                                    }
                                    v41 += (v2 + v2)
                                } else {
                                    if true {
                                        if (42 < 42) {
                                            v61 := GenS5{F0: (!(2147483647 > 65535)), F1: 2, F2: 775}
                                            _ = v61
                                        }
                                    }
                                }
                            case 128:
                                if ((fi16 >> ((42) & 63)) == -218) {
                                    fmt.Println(v18)
                                    _ = len(v7)
                                } else {
                                    if false {
                                        var v62 float32 = (0.0 + v1)
                                        v1 += ((v1 - v1) - (v1 * v1))
                                        v40 = v40
                                        _ = v62
                                    }
                                }
                            default:
                                var v63 uint8 = 135
                                _ = v63
                            }
                            if len(v7) > 0 {
                                v7[0] = fi16
                            }
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                            _ = v40
                            _ = v41
                        }
                    }
                    if len(v7) > 0 {
                        v7[0] = 32767
                    }
                    var v64 string = "a"
                    var v65 uint32 = 4294967295
                    _ = v38
                    _ = v64
                    _ = v65
                }
                _ = fi17
                _ = v18
            }
            _ = fi16
        }
        v1 += v1
        v66 := GenS10{F0: 1, F1: true, F2: false}
        _ = v66
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
}
