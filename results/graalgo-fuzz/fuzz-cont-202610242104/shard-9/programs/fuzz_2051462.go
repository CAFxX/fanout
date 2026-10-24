package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint16 }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 bool; F1 int64 }
type GenS4 struct { F0 int16; F1 uint; F2 string }
type GenS5 struct { F0 uint64; F1 bool }
type GenS6 struct { F0 int; F1 int16 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint8; F1 bool; F2 int64 }
type GenS11 struct { F0 uint32; F1 uint32; F2 int8 }
type GenS12 struct { F0 uint16; F1 int64; F2 uint8 }
type GenS13 struct { F0 uint32 }
func genFn0(p0 int) uint {
    fmt.Println(p0)
    p0 += p0
    p0 = -9223372036854775808
    v5 := GenS1{F0: -1, F1: 255, F2: 65535}
    p0 -= (p0 >> ((2) & 63))
    _ = p0
    _ = v5
    return 42
}

func main() {
    var v1 uint64 = 0
    v2 := []string{"\n\t", "\n\t"}
    v3 := []int{0, -487, 9223372036854775807}
    var v4 uint32 = uint32(v1)
    if false {
        if true {
            if ((1e10 * float32(0.0)) == (1.5 + (float32(1e10) - 1.5))) {
                var v6 int32 = 42
                var v7 uint16 = 255
                for fi8 := 0; fi8 < 10; fi8++ {
                    v9 := GenS2{F0: 1, F1: false}
                    if (!true) {
                        if (false || (18446744073709551615 < genFn0(127))) {
                            v6 += -1
                        }
                        v4 = (v4 % ((2147483648) | 1))
                        var v10 uint16 = (^(v7))
                        _ = v10
                    }
                    _ = fi8
                    _ = v9
                }
                _ = v6
                _ = v7
            } else {
                fmt.Println(v4)
                v3 = append(v3, 2147483647)
                v4 = (v4 ^ v4)
            }
            var v11 int16 = -32768
            _ = v11
        }
        var v12 bool = (57 > 934)
        _ = v12
    }
    if true {
        fmt.Println(v1, v4)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v13 := GenS2{F0: 62, F1: (-9223372036854775808 >= 9223372036854775807)}
    var v14 int16 = 1
    v15 := GenS3{F0: (42 == 1), F1: -9223372036854775808}
    for fi16 := 0; fi16 < 5; fi16++ {
        for fi17 := 0; fi17 < 15; fi17++ {
            v2 = append(v2, "foo bar")
            if ("a" < "hello") {
                if true {
                    v18 := GenS4{F0: 147, F1: 255, F2: ""}
                    var v19 bool = false
                    if true {
                        switch "foo bar" {
                        case "a":
                            v19 = v19
                        case "\n\t":
                            switch "" {
                            case "a":
                                v14 = (v14 * -1)
                            case "hello":
                                if true {
                                    v20 := GenS2{F0: 42, F1: (128 >= 255)}
                                    v4 = (v4 | (v4 * 144))
                                    _ = v20
                                } else {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    var v21 uint16 = 65535
                                    fmt.Println(v14, v21, fi16)
                                    _ = v21
                                }
                            case "":
                                fmt.Println(fi16, v14)
                            default:
                                if true {
                                    v22 := GenS5{F0: v1, F1: v19}
                                    v4 = v4
                                    fmt.Println(v1, v14, fi16)
                                    fmt.Println(fi17, fi17)
                                    _ = v22
                                }
                            }
                            v23 := GenS3{F0: v19, F1: 9223372036854775807}
                            _ = v23
                        default:
                            fmt.Println(v14, fi17, v1)
                        }
                        v24 := GenS6{F0: 9223372036854775807, F1: -1}
                        v19 = v19
                        _ = v24
                    } else {
                        fmt.Println(v19, v4, v19)
                        fmt.Println(fi17)
                        v3 = append(v3, (fi16 % ((int(v14)) | 1)))
                        v25 := GenS7{F0: -1}
                        var v26 int32 = 2147483647
                        _ = v25
                        _ = v26
                    }
                    _ = v18
                    _ = v19
                } else {
                    _ = len(v2)
                    v27 := GenS8{F0: "\n\t"}
                    fmt.Println(fi16, v14)
                    _ = v27
                }
                if true {
                    v28 := GenS8{F0: "\n\t"}
                    _ = v28
                } else {
                    var v29 int = (fi16 * 32767)
                    fmt.Println(v29)
                    _ = v29
                }
            } else {
                v14 = 300
                if false {
                    switch (-(fi17)) {
                    case 9223372036854775807:
                        if (false && false) {
                            if (9223372036854775807 == 1) {
                                v30 := GenS9{F0: fi16}
                                v3 = append(v3, (^(fi16)))
                                _ = len(v2)
                                if ((true || false) || (1 >= 255)) {
                                    if false {
                                        if (0.1 <= 1.5) {
                                            v1 += v1
                                            v1 -= (v1 + v1)
                                            v4 -= v4
                                            var v31 int32 = -1
                                            v32 := GenS10{F0: 1, F1: (0.1 > (((float64(1.5) - float64(0.0)) / (1.5 - 0.0)) / 1.0)), F2: 9223372036854775807}
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            _ = v31
                                            _ = v32
                                        } else {
                                            if true {
                                                fmt.Println(v4)
                                            }
                                            v1 += (v1 ^ v1)
                                        }
                                        var v33 float64 = 0.1
                                        if (false || (229 >= 1)) {
                                            if false {
                                                v34 := GenS8{F0: "\n\t"}
                                                _ = v34
                                            }
                                            if true {
                                                if len(v2) > 0 {
                                                    v2[0] = "foo bar"
                                                }
                                                switch "hello" {
                                                case "a":
                                                    if true {
                                                        var v35 int32 = 665
                                                        _ = v35
                                                    } else {
                                                        v1 -= (v1 / ((v1) | 1))
                                                        var v36 uint32 = v4
                                                        _ = v36
                                                    }
                                                case "hello":
                                                    v14 = v14
                                                case "foo bar":
                                                    var v37 uint16 = 32767
                                                    _ = v37
                                                }
                                            } else {
                                                var v38 int8 = 0
                                                var v39 uint64 = (v1 - v1)
                                                _ = v38
                                                _ = v39
                                            }
                                            v14 = (v14 | v14)
                                        } else {
                                            fmt.Println(v1)
                                            fmt.Println(v33, v33, v33)
                                            fmt.Println(v14)
                                            v4 = 2147483647
                                            var v40 float64 = v33
                                            var v41 int = 0
                                            if true {
                                                v40 = v40
                                                if len(v2) > 0 {
                                                    v2[0] = (("\n\t" + "") + "foo bar")
                                                }
                                                v33 -= v40
                                            }
                                            var v42 uint32 = 533
                                            if false {
                                                v42 -= (v4 % (((v42 >> ((4294967295) & 63))) | 1))
                                                v43 := GenS11{F0: (v42 - (v42 | 27)), F1: v42, F2: -1}
                                                _ = v43
                                            }
                                            var v44 int8 = 0
                                            switch (fi17 ^ (^(fi16))) {
                                            case -237:
                                                v33 = float64(3.14159)
                                            default:
                                                v42 = 268
                                            }
                                            _ = v40
                                            _ = v41
                                            _ = v42
                                            _ = v44
                                        }
                                        _ = v33
                                    } else {
                                        v45 := GenS12{F0: 0, F1: 0, F2: 128}
                                        switch -936 {
                                        case -2147483648:
                                            var v46 int = -9223372036854775808
                                            _ = v46
                                        }
                                        _ = v45
                                    }
                                }
                                _ = v30
                            }
                        }
                    default:
                        v2 = append(v2, "\n\t")
                    }
                } else {
                    v47 := GenS13{F0: 2147483647}
                    _ = v47
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            v14 = 1
            _ = fi17
        }
        if (v1 < 2) {
            fmt.Println(v4)
            v1 += 240
        }
        for fi48 := 0; fi48 < 13; fi48++ {
            v1 = v1
            _ = len(v3)
            if true {
                v49 := GenS13{F0: (v4 ^ 2147483647)}
                _ = v49
            } else {
                var v50 string = "\n\t"
                v51 := GenS7{F0: 127}
                var v52 uint8 = 173
                var v53 float64 = (((((3.14159 + -2.25) / 1.5) * (((0.0 * 3.14159) / 1.5) * (-2.25 / float64((0.0 + (-2.25 - (-2.25 - 3.14159))))))) * (((1e10 + 0.1) * (0.0 + 3.14159)) / (-2.25 - float64((float32(0.1) + 0.0))))) + 3.14159)
                _ = v50
                _ = v51
                _ = v52
                _ = v53
            }
            var v54 int32 = -1
            _ = fi48
            _ = v54
        }
        _ = fi16
    }
    var v55 uint32 = (v4 % (((v4 ^ 1)) | 1))
    for fi56 := 0; fi56 < 17; fi56++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi56
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v55
}
