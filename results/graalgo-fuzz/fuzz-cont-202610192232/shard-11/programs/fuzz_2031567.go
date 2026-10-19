package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 uint64; F1 uint32; F2 int16 }
type GenS4 struct { F0 uint8; F1 uint8 }
type GenS5 struct { F0 int16; F1 string; F2 int32 }
type GenS6 struct { F0 uint8; F1 uint64 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int64; F1 int64; F2 string }
type GenS9 struct { F0 uint8; F1 int16 }
type GenS10 struct { F0 int64; F1 int }
func genFn0(p0 bool) int16 {
    v5 := GenS1{F0: 0, F1: 0}
    var v6 uint16 = 255
    var v7 uint32 = 2147483648
    v8 := GenS2{F0: "a", F1: 34}
    p0 = (!true)
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 1
}

func main() {
    v1 := []uint32{537, 4294967295, 766}
    v2 := []int{32768, 717}
    var v3 int = 4294967295
    var v4 string = "\n\t"
    v9 := GenS3{F0: 85, F1: 2147483647, F2: genFn0(true)}
    v10 := GenS4{F0: 127, F1: 2}
    if len(v2) > 0 {
        v2[0] = v3
    }
    if false {
        var v11 bool = true
        var v12 bool = v11
        _ = v11
        _ = v12
    } else {
        var v13 string = "\n\t"
        _ = v13
    }
    var v14 uint64 = 42
    v2 = append(v2, v3)
    switch -2147483648 {
    case -1:
        if (-1 >= -441) {
            if (!false) {
                v1 = append(v1, 65535)
                fmt.Println(v4)
            }
            if len(v1) > 0 {
                v1[0] = 0
            }
            _ = len(v1)
        }
    case -209:
        v14 -= (^(v14))
    case 32768:
        var v15 uint = 2
        _ = v15
    default:
        if false {
            var v16 int32 = 539
            for fi17 := 0; fi17 < 9; fi17++ {
                for fi18 := 0; fi18 < 1; fi18++ {
                    v19 := GenS5{F0: int16(fi18), F1: v4, F2: (v16 | (^(v16)))}
                    var v20 int32 = (v16 | v16)
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                if (true && (("hello" + (v4 + ((v4 + "") + v4))) == (v4 + ""))) {
                    var v21 float32 = (1e10 - 1.5)
                    _ = v21
                } else {
                    v4 = v4
                    for fi22 := 0; fi22 < 6; fi22++ {
                        if len(v1) > 0 {
                            v1[0] = 1
                        }
                        v1 = append(v1, 708)
                        if (909 == uint(v14)) {
                            if true {
                                if true {
                                    if true {
                                        var v23 uint8 = 255
                                        if (!true) {
                                            var v24 int = int(v16)
                                            v16 += (v16 >> ((255) & 63))
                                            if (289 >= uint16(v23)) {
                                                v4 = "a"
                                                switch ("foo bar" + v4) {
                                                case "a":
                                                    if (65535 == 4294967295) {
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        v2 = append(v2, (v24 | (^(v24))))
                                                    } else {
                                                        fmt.Println(v24, v16)
                                                    }
                                                case "":
                                                    if ("" != (("foo bar" + (v4 + v4)) + (("" + v4) + v4))) {
                                                        if (!false) {
                                                            if ((false || true) && (42 > 2)) {
                                                                v1 = append(v1, 4294967295)
                                                            }
                                                            var v25 uint8 = v23
                                                            _ = len(v1)
                                                            _ = v25
                                                        } else {
                                                            if false {
                                                                if true {
                                                                    v3 = -253
                                                                    v14 += (v14 >> ((18446744073709551615) & 63))
                                                                    var v26 int = fi17
                                                                    _ = v26
                                                                } else {
                                                                    fmt.Println(v23)
                                                                    var v27 int = v3
                                                                    var v28 int = v3
                                                                    _ = v27
                                                                    _ = v28
                                                                }
                                                                v29 := GenS6{F0: (v23 & 255), F1: v14}
                                                                v16 = -2147483648
                                                                v24 = -1
                                                                _ = v29
                                                            } else {
                                                                if true {
                                                                    if true {
                                                                        v24 = (fi22 * v24)
                                                                    }
                                                                    var v30 float32 = (float32(1e10) * (3.14159 / 0.1))
                                                                    _ = v30
                                                                } else {
                                                                    if (int32(v24) > int32(v23)) {
                                                                        var v31 int8 = -1
                                                                        _ = v31
                                                                    }
                                                                }
                                                                _ = len(v1)
                                                            }
                                                            var v32 float64 = 1e10
                                                            if len(v1) > 0 {
                                                                v1[0] = 994
                                                            }
                                                            _ = v32
                                                        }
                                                        var v33 uint16 = 42
                                                        _ = v33
                                                    } else {
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                    }
                                                }
                                                if (!false) {
                                                    v4 = ("" + v4)
                                                } else {
                                                    var v34 int32 = 2147483647
                                                    v34 += (-(v16))
                                                    _ = v34
                                                }
                                                if true {
                                                    var v35 uint32 = uint32(v14)
                                                    v3 = (v3 / ((-143) | 1))
                                                    var v36 int8 = 1
                                                    v16 += v16
                                                    v37 := GenS5{F0: 1, F1: "", F2: -2147483648}
                                                    _ = v35
                                                    _ = v36
                                                    _ = v37
                                                } else {
                                                    v38 := GenS3{F0: (v14 / (((v14 + v14)) | 1)), F1: 561, F2: -32768}
                                                    v39 := GenS7{F0: true}
                                                    _ = v38
                                                    _ = v39
                                                }
                                            } else {
                                                v40 := GenS7{F0: (int64(v23) >= 42)}
                                                v4 = (v4 + ((("a" + v4) + v4) + "foo bar"))
                                                if ((!false) && true) {
                                                    v41 := GenS5{F0: 0, F1: v4, F2: int32(v24)}
                                                    fmt.Println(v24, fi22, fi17)
                                                    v42 := GenS8{F0: 9223372036854775807, F1: -811, F2: v4}
                                                    _ = len(v2)
                                                    _ = v41
                                                    _ = v42
                                                }
                                                _ = v40
                                            }
                                            _ = v24
                                        } else {
                                            var v43 int16 = -1
                                            _ = v43
                                        }
                                        _ = v23
                                    } else {
                                        v44 := GenS5{F0: 1000, F1: (v4 + v4), F2: (v16 % ((v16) | 1))}
                                        _ = v44
                                    }
                                    var v45 float64 = 0.0
                                    _ = v45
                                }
                                v3 = 9223372036854775807
                            }
                        }
                        _ = fi22
                    }
                    fmt.Println(fi17)
                }
                _ = fi17
            }
            _ = v16
        } else {
            v14 += (v14 + v14)
        }
    }
    if len(v2) > 0 {
        v2[0] = (v3 << ((917) & 63))
    }
    if ((true || (false || true)) && true) {
        var v46 int = 42
        v47 := GenS9{F0: 160, F1: -567}
        var v48 int32 = 0
        _ = v46
        _ = v47
        _ = v48
    } else {
        v14 -= 2
        for fi49 := 0; fi49 < 5; fi49++ {
            v1 = append(v1, 65535)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi49
        }
    }
    if (!true) {
        var v50 uint64 = 0
        fmt.Println(v3)
        _ = v50
    } else {
        var v51 int32 = -144
        fmt.Println(v3)
        _ = v51
    }
    var v52 float64 = 1e10
    if (!(((-2.25 + ((1e10 / 1e10) / 1e10)) + 1.5) <= 3.14159)) {
        v14 = (v14 | 1)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        if false {
            if (0 >= -32768) {
                switch "" {
                case "\n\t":
                    var v53 float64 = v52
                    _ = v53
                case "hello":
                    var v54 uint64 = 4294967295
                    _ = v54
                case "":
                    v14 += 381
                }
            }
            switch v4 {
            case "a":
                var v55 uint16 = 919
                _ = v55
            case "hello":
                var v56 uint64 = (v14 << ((42) & 63))
                _ = v56
            case "\n\t":
                var v57 int8 = -70
                _ = v57
            }
            if ((!(false && true)) && false) {
                for fi58 := 0; fi58 < 17; fi58++ {
                    var v59 float32 = (((-2.25 * (float32(v52) - (((((3.14159 * 1.5) * -2.25) - float32(v52)) - 0.0) * ((-2.25 * 3.14159) + (-2.25 * (0.1 / 1.5)))))) * 1e10) / 1.0)
                    _ = fi58
                    _ = v59
                }
                var v60 uint = 65535
                _ = v60
            }
        } else {
            var v61 uint16 = 32767
            v62 := GenS10{F0: 0, F1: (-(v3))}
            if false {
                v4 = v4
                var v63 uint = 2
                _ = v63
            }
            _ = v61
            _ = v62
        }
        fmt.Println(v4, v52, v52)
    }
    fmt.Println(v4)
    var v64 float32 = -2.25
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v52:", v52)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v52
    _ = v64
    _ = v9
}
