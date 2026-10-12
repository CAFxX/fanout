package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int8; F1 string; F2 int16 }
type GenS7 struct { F0 int; F1 uint32; F2 int16 }
type GenS8 struct { F0 uint; F1 uint64 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint64; F1 int8; F2 uint16 }
type GenS11 struct { F0 uint; F1 bool; F2 uint64 }
type GenS12 struct { F0 uint16 }
func genFn0() uint64 {
    if true {
        var v7 float64 = 1e10
        if false {
            switch 32768 {
            case 2147483647:
                var v8 uint32 = 993
                _ = v8
            case -9223372036854775808:
                v9 := GenS1{F0: -656, F1: 151}
                _ = v9
            default:
                v10 := []int{32767, -1, 1}
                _ = v10
            }
        }
        _ = v7
    } else {
        var v11 int32 = 438
        var v12 int64 = 42
        _ = v11
        _ = v12
    }
    var v13 uint = 255
    v14 := []int{4294967295, 609, 9223372036854775807, 32768}
    for fi15 := 0; fi15 < 10; fi15++ {
        v13 = (v13 * (v13 << (((^(v13))) & 63)))
        var v16 int8 = 1
        _ = fi15
        _ = v16
    }
    _ = v13
    _ = v14
    return 1
}
func genFn1(p0 int8, p1 uint64, p2 uint16) int64 {
    fmt.Println(p1)
    v17 := GenS2{F0: 2, F1: (p1 % (((p1 >> ((255) & 63))) | 1))}
    for fi18 := 0; fi18 < 6; fi18++ {
        var v19 uint16 = p2
        if ((true || false) || (366 == 2147483647)) {
            v20 := GenS3{F0: ""}
            var v21 int64 = 9223372036854775807
            _ = v20
            _ = v21
        } else {
            v19 = v19
            fmt.Println(v19, p1, fi18)
        }
        _ = fi18
        _ = v19
    }
    if true {
        for fi22 := 0; fi22 < 9; fi22++ {
            if false {
                fmt.Println(p1)
            } else {
                v23 := []int{-9223372036854775808, (fi22 + (^(fi22))), fi22}
                _ = v23
            }
            v24 := GenS4{F0: 2}
            _ = fi22
            _ = v24
        }
    } else {
        var v25 float32 = ((((((-2.25 - 1e10) + (0.0 / 1.0)) + (1.5 / 0.1)) / 0.1) * -2.25) - float32(((0.1 / (3.14159 * (1e10 - 1.5))) + 3.14159)))
        var v26 int64 = -9223372036854775808
        v27 := GenS5{F0: ""}
        _ = v25
        _ = v26
        _ = v27
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return 1
}

func main() {
    v1 := make(map[string]int)
    var v2 int = 9223372036854775807
    var v3 int64 = 42
    var v4 uint16 = 32767
    v5 := make(map[string]bool)
    v6 := []uint32{2147483648, 42}
    v3 -= v3
    fmt.Println(v3)
    var v28 int32 = -2147483648
    if true {
        if (18446744073709551615 == genFn0()) {
            v2 = v2
            var v29 uint32 = 0
            var v30 int32 = v28
            _ = v29
            _ = v30
        } else {
            var v31 bool = true
            for fi32 := 0; fi32 < 11; fi32++ {
                var v33 int32 = v28
                v34 := v5["a"]
                if (1.5 > float32(0.1)) {
                    if v34 {
                        if v34 {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        }
                        if true {
                            v28 += -5
                            v31 = (false && false)
                            for fi35 := 0; fi35 < 16; fi35++ {
                                v33 += (v33 * -1)
                                v34 = (genFn0() < 1)
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                _ = fi35
                            }
                        } else {
                            if (true || v34) {
                                var v36 uint16 = (v4 ^ v4)
                                if false {
                                    var v37 int64 = int64(v2)
                                    v37 = (v3 | 1)
                                    fmt.Println(v28, v3)
                                    _ = v37
                                }
                                _ = v36
                            }
                        }
                    }
                    _ = len(v5)
                }
                _ = fi32
                _ = v33
                _ = v34
            }
            _ = v31
        }
    } else {
        v4 = (v4 | 42)
    }
    if ((3.14159 / 0.1) != 0.0) {
        fmt.Println(v2)
        v4 -= v4
    } else {
        v4 = 65535
        v2 -= (v2 >> ((42) & 63))
        v4 = (v4 ^ 837)
    }
    var v38 float32 = float32(1.5)
    if true {
        var v39 float32 = v38
        _ = v39
    }
    for fi40 := 0; fi40 < 17; fi40++ {
        v41 := GenS6{F0: 127, F1: "", F2: 0}
        _ = fi40
        _ = v41
    }
    v6 = append(v6, 1)
    switch v2 {
    case 42:
        for fi42 := 0; fi42 < 11; fi42++ {
            var v43 int64 = 42
            v3 += (v3 % ((int64(v28)) | 1))
            _ = fi42
            _ = v43
        }
    case 32767:
        v28 -= v28
    default:
        switch 709 {
        case -447:
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        case 127:
            switch "foo bar" {
            case "\n\t":
                var v44 uint64 = 4294967295
                _ = v44
            default:
                if false {
                    v2 += (v2 | (v2 | (-(v2))))
                    if len(v6) > 0 {
                        v6[0] = 2147483648
                    }
                } else {
                    v45 := GenS7{F0: int(v4), F1: 42, F2: 1}
                    _ = v45
                }
            }
        case -9223372036854775808:
            if false {
                for fi46 := 0; fi46 < 1; fi46++ {
                    if (!false) {
                        if false {
                            if false {
                                switch "hello" {
                                case "foo bar":
                                    fmt.Println(v4, v3)
                                default:
                                    v1[("" + "a")] = fi46
                                }
                            }
                            if (1e10 <= 1.5) {
                                if (2147483647 >= (v28 | 1)) {
                                    for fi47 := 0; fi47 < 14; fi47++ {
                                        if false {
                                            v3 = (-(v3))
                                            v48 := GenS8{F0: 18446744073709551615, F1: 2}
                                            _ = v48
                                        }
                                        fmt.Println(fi47, v28, fi46)
                                        var v49 uint = 224
                                        _ = fi47
                                        _ = v49
                                    }
                                    var v50 float64 = (float64(v38) - 0.1)
                                    _ = v50
                                } else {
                                    v4 = v4
                                }
                            }
                            var v51 uint32 = 42
                            _ = v51
                        }
                        v52 := GenS3{F0: "\n\t"}
                        _ = v52
                    } else {
                        var v53 string = ""
                        _ = v53
                    }
                    v4 = v4
                    _ = fi46
                }
                if len(v6) > 0 {
                    v6[0] = 65535
                }
            }
        }
    }
    switch "" {
    case "\n\t":
        var v54 string = ("\n\t" + "foo bar")
        _ = v54
    case "foo bar":
        v6 = append(v6, 2147483648)
    default:
        fmt.Println(v4, v28)
    }
    v4 = 255
    if true {
        if true {
            v55 := GenS1{F0: -416, F1: 249}
            for fi56 := 0; fi56 < 11; fi56++ {
                if len(v6) > 0 {
                    v6[0] = 70
                }
                _ = fi56
            }
            _ = v55
        } else {
            for fi57 := 0; fi57 < 12; fi57++ {
                for fi58 := 0; fi58 < 5; fi58++ {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    fmt.Println(v28)
                    var v59 uint = 431
                    _ = fi58
                    _ = v59
                }
                switch (v2 << ((255) & 63)) {
                case 9223372036854775807:
                    v2 = (^(v2))
                case 1:
                    switch (v2 << ((uint(v4)) & 63)) {
                    case 2147483647:
                        for fi60 := 0; fi60 < 13; fi60++ {
                            switch "hello" {
                            case "a":
                                v28 = v28
                            case "hello":
                                if true {
                                    if (true || (1 >= 48)) {
                                        if (!true) {
                                            var v61 string = (("hello" + ("a" + "a")) + ("hello" + ("hello" + ("\n\t" + ""))))
                                            v3 -= v3
                                            _ = v61
                                        }
                                        var v62 uint = 4294967295
                                        fmt.Println(v28)
                                        _ = v62
                                    } else {
                                        v63 := GenS2{F0: 42, F1: 4294967295}
                                        if true {
                                            v64 := GenS9{F0: 100}
                                            if true {
                                                var v65 uint8 = 147
                                                _ = v65
                                            } else {
                                                v66 := GenS1{F0: 0, F1: 2}
                                                v67 := v1["\n\t"]
                                                _ = v66
                                                _ = v67
                                            }
                                            _ = v64
                                        }
                                        if ((fi60 >> ((4294967295) & 63)) > fi57) {
                                            v68 := GenS10{F0: 42, F1: 100, F2: v4}
                                            _ = v68
                                        }
                                        _ = v63
                                    }
                                    v69 := GenS11{F0: 543, F1: (true && false), F2: 2}
                                    if len(v6) > 0 {
                                        v6[0] = 4294967295
                                    }
                                    _ = v69
                                }
                            default:
                                var v70 int32 = 774
                                _ = v70
                            }
                            _ = len(v5)
                            var v71 bool = (true || (4294967295 >= 421))
                            _ = fi60
                            _ = v71
                        }
                    case 0:
                        fmt.Println(fi57, v2)
                    case -9223372036854775808:
                        var v72 uint32 = 4294967295
                        _ = v72
                    default:
                        v4 += (^(v4))
                    }
                case -459:
                    fmt.Println(fi57, v2, v2)
                }
                v1[("" + "")] = (v2 >> ((255) & 63))
                _ = fi57
            }
        }
        var v73 float32 = v38
        _ = v73
    } else {
        v4 = (v4 | (v4 >> ((921) & 63)))
        v2 += (v2 & v2)
        v74 := GenS5{F0: ("\n\t" + "foo bar")}
        _ = v74
    }
    v38 += (0.1 / (v38 - float32((((((3.14159 + 0.1) + (((0.1 / (1.5 / 1e10)) - (-2.25 * float64(3.14159))) * 0.0)) / ((1.5 - float64((0.0 * 3.14159))) * 1.5)) * 3.14159) + float64(v38)))))
    v3 = (v3 >> ((255) & 63))
    var v75 int64 = (v3 | (v3 * 404))
    v76 := GenS12{F0: (v4 | 65535)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v75
    _ = v76
}
