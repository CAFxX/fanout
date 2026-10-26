package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 string; F1 uint; F2 uint32 }
type GenS3 struct { F0 int32; F1 uint; F2 uint16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int8; F1 uint8 }
type GenS7 struct { F0 int16; F1 uint8 }
type GenS8 struct { F0 int8; F1 uint64; F2 uint16 }
type GenS9 struct { F0 uint32; F1 uint8; F2 uint }
type GenS10 struct { F0 uint8; F1 uint16; F2 uint32 }
type GenS11 struct { F0 uint64 }
func genFn0(p0 int8, p1 uint16) uint64 {
    fmt.Println(p0)
    var v6 uint16 = (p1 & (p1 % (((p1 & 42)) | 1)))
    _ = p0
    _ = p1
    _ = v6
    return 446
}
func genFn1(p0 uint8, p1 int64, p2 int) int64 {
    p2 -= p2
    if (((("hello" + "\n\t") + "a") + (("" + "a") + "")) >= (("foo bar" + ("foo bar" + (("foo bar" + "") + ("a" + "hello")))) + "foo bar")) {
        for fi7 := 0; fi7 < 2; fi7++ {
            var v8 uint = 7
            for fi9 := 0; fi9 < 11; fi9++ {
                var v10 int32 = -1
                fmt.Println(p2, v10)
                _ = fi9
                _ = v10
            }
            _ = fi7
            _ = v8
        }
    }
    for fi11 := 0; fi11 < 16; fi11++ {
        p0 = p0
        switch p2 {
        case -949:
            v12 := GenS1{F0: 269, F1: 1}
            _ = v12
        case -823:
            if false {
                var v13 uint32 = 944
                p0 = (p0 & p0)
                _ = v13
            }
        }
        v14 := []int{(p2 - 127)}
        _ = fi11
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    return (p1 / (((^(p1))) | 1))
}

func main() {
    v1 := []int{4294967295, 0, 4294967295}
    v2 := make(map[uint32]int)
    v3 := []uint32{2147483647, 65535}
    var v4 int32 = 941
    var v5 int8 = 0
    switch ("foo bar" + ((((("" + "\n\t") + (("" + "hello") + "\n\t")) + "hello") + "\n\t") + "hello")) {
    case "foo bar":
        v4 = 1
    }
    var v15 string = "foo bar"
    for fi16 := 0; fi16 < 17; fi16++ {
        v17 := GenS1{F0: 1, F1: 654}
        if true {
            v18 := GenS2{F0: "hello", F1: 2, F2: 4294967295}
            for fi19 := 0; fi19 < 9; fi19++ {
                switch fi16 {
                case 2147483647:
                    fmt.Println(fi19, fi16)
                case 1:
                    v5 -= v5
                case 920:
                    v15 = v15
                default:
                    if ((-(v5)) >= -128) {
                        switch fi16 {
                        case -1:
                            var v20 uint = 65535
                            v21 := GenS3{F0: v4, F1: (v20 * (v20 | v20)), F2: 32767}
                            if true {
                                fmt.Println(fi16, v4, v15)
                                if false {
                                    v4 = (v4 % ((885) | 1))
                                    if (1.5 == 3.14159) {
                                        if (true || (0.0 != (-2.25 / 1e10))) {
                                            if (593 < 577) {
                                                if (!(false || false)) {
                                                    switch v15 {
                                                    case "\n\t":
                                                        switch fi19 {
                                                        case 32768:
                                                            if false {
                                                                var v22 float32 = (((1e10 / 3.14159) * 0.1) + 1.5)
                                                                _ = v22
                                                            }
                                                            v23 := GenS3{F0: (v4 ^ (v4 - v4)), F1: 0, F2: 65535}
                                                            _ = v23
                                                        case 9223372036854775807:
                                                            if false {
                                                                if false {
                                                                    var v24 int64 = -9223372036854775808
                                                                    var v25 string = ""
                                                                    var v26 string = "foo bar"
                                                                    v24 = -663
                                                                    if (!false) {
                                                                        if false {
                                                                            if (!(!((v25 + "") != v26))) {
                                                                                var v27 float32 = 0.0
                                                                                _ = v27
                                                                            } else {
                                                                                var v28 int = fi19
                                                                                v5 -= (v5 >> ((v20) & 63))
                                                                                v26 = (("foo bar" + v25) + v25)
                                                                                _ = v28
                                                                            }
                                                                            if len(v1) > 0 {
                                                                                v1[0] = -1
                                                                            }
                                                                            v20 = (^(v20))
                                                                        }
                                                                        v29 := GenS4{F0: 128}
                                                                        _ = len(v3)
                                                                        _ = v29
                                                                    } else {
                                                                        var v30 int16 = 42
                                                                        var v31 bool = ((v20 + (v20 + v20)) < (v20 - v20))
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                        fmt.Println(v5, v15)
                                                                        _ = v30
                                                                        _ = v31
                                                                    }
                                                                    v26 = "hello"
                                                                    v32 := GenS5{F0: (^(v24))}
                                                                    v25 = "a"
                                                                    _ = v24
                                                                    _ = v25
                                                                    _ = v26
                                                                    _ = v32
                                                                } else {
                                                                    if len(v3) > 0 {
                                                                        v3[0] = 941
                                                                    }
                                                                    var v33 uint16 = 255
                                                                    _ = len(v2)
                                                                    _ = v33
                                                                }
                                                                v15 = v15
                                                            }
                                                        default:
                                                            if ((fi16 - (fi19 >> ((18446744073709551615) & 63))) <= 127) {
                                                                var v34 uint32 = 0
                                                                _ = v34
                                                            }
                                                        }
                                                    case "hello":
                                                        v35 := GenS5{F0: -806}
                                                        _ = v35
                                                    default:
                                                        v15 = (v15 + "foo bar")
                                                    }
                                                } else {
                                                    switch fi16 {
                                                    case 4294967295:
                                                        var v36 uint8 = 127
                                                        _ = v36
                                                    case 32768:
                                                        delete(v2, 869)
                                                    }
                                                    if (1 > 0) {
                                                        switch (("\n\t" + v15) + v15) {
                                                        case "a":
                                                            var v37 float64 = 0.1
                                                            _ = v37
                                                        case "\n\t":
                                                            v15 = ""
                                                        }
                                                        fmt.Println(fi16)
                                                    } else {
                                                        if (true && false) {
                                                            v2[834] = (fi19 / ((fi16) | 1))
                                                        } else {
                                                            var v38 float64 = 3.14159
                                                            if false {
                                                                v3 = append(v3, 65535)
                                                                v38 -= (v38 * 0.1)
                                                                _ = len(v2)
                                                            } else {
                                                                if (0.1 <= (((0.0 / float64((3.14159 / 3.14159))) * (v38 * ((v38 + v38) + ((v38 - float64(0.0)) - v38)))) - v38)) {
                                                                    v5 = int8(fi19)
                                                                    delete(v2, 65535)
                                                                    v39 := GenS3{F0: v4, F1: 65535, F2: 255}
                                                                    _ = v39
                                                                }
                                                            }
                                                            fmt.Println(v20)
                                                            switch int(v5) {
                                                            case 450:
                                                                if false {
                                                                    v4 = v4
                                                                } else {
                                                                    if len(v1) > 0 {
                                                                        v1[0] = (fi16 - fi19)
                                                                    }
                                                                }
                                                            case 127:
                                                                v38 -= v38
                                                            }
                                                            _ = v38
                                                        }
                                                        if true {
                                                            var v40 int32 = (v4 | v4)
                                                            _ = v40
                                                        }
                                                        v4 += v4
                                                    }
                                                }
                                                if ((fi16 / ((128) | 1)) != fi19) {
                                                    v15 = v15
                                                    v4 = (v4 >> ((v20) & 63))
                                                }
                                            } else {
                                                v41 := GenS1{F0: 0, F1: 471}
                                                v15 = "foo bar"
                                                switch v15 {
                                                case "":
                                                    switch v15 {
                                                    case "foo bar":
                                                        switch 128 {
                                                        case 42:
                                                            var v42 string = v15
                                                            _ = v42
                                                        case 2147483647:
                                                            if false {
                                                                v4 = v4
                                                                v43 := GenS6{F0: v5, F1: 0}
                                                                switch 4294967295 {
                                                                case 4294967295:
                                                                    var v44 int8 = v5
                                                                    _ = v44
                                                                case 127:
                                                                    v45 := GenS6{F0: (v5 + v5), F1: 215}
                                                                    _ = v45
                                                                }
                                                                _ = v43
                                                            }
                                                        }
                                                    case "\n\t":
                                                        switch v15 {
                                                        case "hello":
                                                            v46 := v2[0]
                                                            _ = v46
                                                        case "foo bar":
                                                            var v47 string = "hello"
                                                            _ = v47
                                                        }
                                                    default:
                                                        v20 -= v20
                                                    }
                                                case "\n\t":
                                                    v4 -= (v4 - 2147483647)
                                                case "hello":
                                                    if ((!true) && (51 != 2)) {
                                                        v48 := GenS7{F0: 1, F1: 0}
                                                        var v49 uint64 = 690
                                                        fmt.Println(v5, v15)
                                                        v15 = "hello"
                                                        v4 -= 131
                                                        v50 := GenS8{F0: (v5 / (((v5 | -114)) | 1)), F1: genFn0(-1, 0), F2: 2}
                                                        _ = v48
                                                        _ = v49
                                                        _ = v50
                                                    } else {
                                                        if false {
                                                            switch -879 {
                                                            case 0:
                                                                v51 := GenS2{F0: v15, F1: 687, F2: 308}
                                                                _ = v51
                                                            case 2147483647:
                                                                var v52 uint32 = 4294967295
                                                                _ = v52
                                                            case 1:
                                                                v20 = (v20 >> ((4294967295) & 63))
                                                            }
                                                            v53 := GenS9{F0: 65535, F1: 174, F2: (v20 << ((677) & 63))}
                                                            _ = v53
                                                        } else {
                                                            if true {
                                                                v15 = ("foo bar" + v15)
                                                                var v54 string = v15
                                                                var v55 uint = (v20 | (v20 % ((v20) | 1)))
                                                                _ = v54
                                                                _ = v55
                                                            }
                                                            v56 := GenS3{F0: v4, F1: v20, F2: 2}
                                                            _ = v56
                                                        }
                                                        v2[425] = -707
                                                        v57 := GenS4{F0: 128}
                                                        _ = v57
                                                    }
                                                }
                                                _ = v41
                                            }
                                            switch 42 {
                                            case -9223372036854775808:
                                                delete(v2, 1)
                                            case -1:
                                                var v58 int = (-(fi19))
                                                _ = v58
                                            case 42:
                                                var v59 uint64 = 2
                                                _ = v59
                                            default:
                                                var v60 bool = false
                                                _ = v60
                                            }
                                            var v61 int32 = 42
                                            _ = v61
                                        }
                                    }
                                }
                            } else {
                                fmt.Println(v15, v15)
                                var v62 int64 = genFn1(42, 1, (-(fi16)))
                                v63 := GenS9{F0: 209, F1: 127, F2: 65535}
                                _ = v62
                                _ = v63
                            }
                            delete(v2, 42)
                            switch 932 {
                            case 421:
                                v5 = (v5 ^ (v5 ^ v5))
                            case 2147483647:
                                v64 := GenS10{F0: 0, F1: 32767, F2: 142}
                                _ = v64
                            case -1:
                                if len(v1) > 0 {
                                    v1[0] = fi19
                                }
                                v65 := GenS7{F0: 32767, F1: 1}
                                fmt.Println(v20, fi16)
                                if len(v1) > 0 {
                                    v1[0] = (fi19 * (fi19 ^ (fi19 & fi19)))
                                }
                                if true {
                                    var v66 bool = true
                                    v5 = v5
                                    v4 = (v4 ^ (-(v4)))
                                    _ = v66
                                }
                                _ = v65
                            default:
                                v15 = "\n\t"
                            }
                            var v67 uint32 = 42
                            _ = v20
                            _ = v21
                            _ = v67
                        case 9223372036854775807:
                            switch v15 {
                            case "\n\t":
                                fmt.Println(v5, v15)
                            }
                        }
                        v15 = "foo bar"
                        _ = len(v2)
                    } else {
                        var v68 int8 = (v5 << ((18446744073709551615) & 63))
                        v69 := v2[750]
                        v68 += (v68 >> ((65535) & 63))
                        _ = v68
                        _ = v69
                    }
                }
                v15 = v15
                _ = fi19
            }
            _ = v18
        }
        if true {
            var v70 uint = 918
            _ = v70
        }
        _ = fi16
        _ = v17
    }
    v15 = (v15 + "")
    switch -979 {
    case 32767:
        fmt.Println(v5)
    case 128:
        if false {
            v15 = v15
            fmt.Println(v15)
        }
    case -157:
        v4 = v4
    }
    v71 := GenS5{F0: 42}
    var v72 uint64 = genFn0((^(v5)), 65535)
    v5 -= (v5 >> ((65) & 63))
    var v73 int = 188
    var v74 uint16 = 2
    if (false || false) {
        v74 += v74
        fmt.Println(v4, v4, v74)
        v5 -= 38
    } else {
        if len(v3) > 0 {
            v3[0] = uint32(v74)
        }
        for fi75 := 0; fi75 < 8; fi75++ {
            var v76 uint8 = 0
            v1 = append(v1, fi75)
            _ = fi75
            _ = v76
        }
        for fi77 := 0; fi77 < 12; fi77++ {
            v78 := GenS2{F0: (v15 + v15), F1: 4294967295, F2: 65535}
            _ = fi77
            _ = v78
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    _ = len(v2)
    fmt.Println(v5, v72, v4)
    v79 := GenS11{F0: v72}
    for fi80 := 0; fi80 < 7; fi80++ {
        var v81 uint = 65535
        _ = fi80
        _ = v81
    }
    for fi82 := 0; fi82 < 5; fi82++ {
        var v83 int64 = -9223372036854775808
        _ = fi82
        _ = v83
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v72:", v72)
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v71
    _ = v72
    _ = v73
    _ = v74
    _ = v79
}
