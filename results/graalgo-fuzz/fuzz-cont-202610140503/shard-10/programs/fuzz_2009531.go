package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 uint8; F1 uint64; F2 int32 }
type GenS6 struct { F0 uint; F1 uint; F2 int64 }
type GenS7 struct { F0 uint32; F1 uint8 }
type GenS8 struct { F0 int64; F1 bool }
func genFn0(p0 uint64) uint64 {
    for fi4 := 0; fi4 < 6; fi4++ {
        var v5 string = ("" + "a")
        if true {
            v5 = "foo bar"
            fmt.Println(fi4)
            switch v5 {
            case "":
                for fi6 := 0; fi6 < 18; fi6++ {
                    v7 := []int{fi4, (fi4 & fi4)}
                    _ = fi6
                    _ = v7
                }
            case "a":
                v5 = ("hello" + ((v5 + v5) + "foo bar"))
            case "\n\t":
                fmt.Println(fi4)
            }
        } else {
            for fi8 := 0; fi8 < 11; fi8++ {
                p0 = (p0 / ((18446744073709551615) | 1))
                switch fi8 {
                case 1:
                    v5 = (v5 + ((("hello" + v5) + v5) + (v5 + ((v5 + v5) + (v5 + v5)))))
                case 42:
                    v9 := GenS1{F0: 100, F1: 360}
                    _ = v9
                case 127:
                    fmt.Println(fi8, fi4)
                }
                _ = fi8
            }
            var v10 uint = 18446744073709551615
            _ = v10
        }
        v5 = (("" + v5) + v5)
        _ = fi4
        _ = v5
    }
    p0 = p0
    fmt.Println(p0)
    _ = p0
    return (p0 << ((18446744073709551615) & 63))
}
func genFn1() uint {
    for fi11 := 0; fi11 < 16; fi11++ {
        var v12 float32 = 0.1
        _ = fi11
        _ = v12
    }
    var v13 int8 = -1
    v13 = v13
    var v14 bool = (-1 != -901)
    _ = v13
    _ = v14
    return 65535
}

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 22
    var v3 int32 = 2147483647
    fmt.Println(v2, v3)
    v3 = v3
    var v15 int32 = (v3 ^ (v3 ^ (^(v3))))
    if (!(!false)) {
        v15 += v3
        for fi16 := 0; fi16 < 12; fi16++ {
            var v17 uint16 = 65535
            for fi18 := 0; fi18 < 14; fi18++ {
                if false {
                    if ((!true) && true) {
                        v19 := GenS1{F0: 100, F1: 9223372036854775807}
                        _ = v19
                    } else {
                        var v20 int = -1
                        _ = v20
                    }
                    v21 := GenS2{F0: 113}
                    fmt.Println(fi16, fi16)
                    _ = v21
                } else {
                    var v22 uint8 = v2
                    _ = v22
                }
                _ = fi18
            }
            switch ("foo bar" + "\n\t") {
            case "a":
                for fi23 := 0; fi23 < 3; fi23++ {
                    fmt.Println(fi23)
                    _ = fi23
                }
            case "foo bar":
                fmt.Println(v2)
            }
            _ = fi16
            _ = v17
        }
    } else {
        v2 = v2
    }
    if false {
        for fi24 := 0; fi24 < 15; fi24++ {
            for fi25 := 0; fi25 < 4; fi25++ {
                var v26 float32 = (-2.25 * (0.0 * 1e10))
                _ = fi25
                _ = v26
            }
            fmt.Println(fi24)
            _ = fi24
        }
        switch "" {
        case "foo bar":
            v3 = 42
        case "":
            var v27 int16 = -594
            _ = v27
        case "a":
            v28 := GenS3{F0: 345, F1: (((("" + ("\n\t" + (("\n\t" + ("a" + "foo bar")) + "\n\t"))) + (((("\n\t" + "hello") + "") + "foo bar") + ("hello" + ("" + "foo bar")))) + "\n\t") + "")}
            _ = v28
        }
    } else {
        delete(v1, -9223372036854775808)
        v2 = v2
    }
    var v29 uint8 = 127
    v30 := GenS4{F0: genFn0(42), F1: 227}
    if false {
        if (!true) {
            v1[-2147483648] = "\n\t"
        } else {
            v31 := []int{879, -993, -174, 127}
            _ = v31
        }
        switch -9223372036854775808 {
        case 888:
            switch ("foo bar" + ("a" + "a")) {
            case "a":
                v32 := GenS5{F0: v2, F1: 1, F2: v3}
                _ = v32
            case "hello":
                if true {
                    if false {
                        if (18446744073709551615 != 42) {
                            for fi33 := 0; fi33 < 13; fi33++ {
                                fmt.Println(fi33)
                                if false {
                                    for fi34 := 0; fi34 < 2; fi34++ {
                                        v2 = v29
                                        fmt.Println(v15)
                                        var v35 float32 = (3.14159 + (0.0 - 0.0))
                                        v35 += v35
                                        _ = fi34
                                        _ = v35
                                    }
                                } else {
                                    v36 := []int{(fi33 + (fi33 % (((-(fi33))) | 1))), fi33, (fi33 >> ((genFn1()) & 63))}
                                    _ = len(v36)
                                    _ = v36
                                }
                                _ = fi33
                            }
                            if true {
                                switch 126 {
                                case 851:
                                    v29 -= (v29 | 0)
                                case 4294967295:
                                    for fi37 := 0; fi37 < 17; fi37++ {
                                        if true {
                                            var v38 int = fi37
                                            v15 = (v3 ^ v3)
                                            var v39 uint64 = 634
                                            _ = v38
                                            _ = v39
                                        } else {
                                            if true {
                                                v40 := GenS1{F0: 100, F1: -1}
                                                _ = v40
                                            }
                                            v2 = v29
                                            if (0 < 255) {
                                                for fi41 := 0; fi41 < 8; fi41++ {
                                                    delete(v1, fi41)
                                                    _ = fi41
                                                }
                                            } else {
                                                delete(v1, (fi37 * fi37))
                                            }
                                        }
                                        _ = len(v1)
                                        _ = fi37
                                    }
                                case -9223372036854775808:
                                    switch (("" + (("hello" + "hello") + ("foo bar" + "foo bar"))) + "\n\t") {
                                    case "foo bar":
                                        var v42 float32 = 3.14159
                                        _ = v42
                                    }
                                }
                                switch ("a" + "foo bar") {
                                case "a":
                                    var v43 int64 = 9223372036854775807
                                    _ = v43
                                case "foo bar":
                                    for fi44 := 0; fi44 < 20; fi44++ {
                                        var v45 uint = 0
                                        fmt.Println(v29, v29, v29)
                                        _ = fi44
                                        _ = v45
                                    }
                                case "hello":
                                    v46 := GenS4{F0: genFn0(0), F1: (v29 >> ((2) & 63))}
                                    _ = v46
                                }
                            } else {
                                fmt.Println(v29)
                                v47 := GenS2{F0: -1}
                                _ = v47
                            }
                        } else {
                            v3 = (^(v15))
                        }
                    } else {
                        v15 -= v3
                        var v48 uint8 = v29
                        v49 := v1[395]
                        _ = v48
                        _ = v49
                    }
                    if (!((0.0 / (0.1 * (3.14159 / 1.0))) != (0.1 / 1.5))) {
                        if true {
                            var v50 int64 = 1
                            _ = v50
                        } else {
                            var v51 int8 = 127
                            fmt.Println(v3)
                            _ = v51
                        }
                        v52 := v1[-9223372036854775808]
                        _ = v52
                    }
                    v53 := []int{}
                    _ = v53
                } else {
                    if ((!true) || (-256 == -1)) {
                        var v54 string = "a"
                        for fi55 := 0; fi55 < 6; fi55++ {
                            var v56 uint = genFn1()
                            if true {
                                var v57 int16 = 894
                                _ = v57
                            } else {
                                switch -937 {
                                case 42:
                                    var v58 uint = v56
                                    _ = v58
                                case 32767:
                                    v3 -= v15
                                }
                                var v59 uint32 = 2147483648
                                _ = v59
                            }
                            for fi60 := 0; fi60 < 16; fi60++ {
                                v61 := []int{(fi60 & (fi60 >> (((v56 / (((v56 % (((v56 - v56)) | 1))) | 1))) & 63))), fi55, fi55, (fi55 + (fi60 & (fi60 % ((fi55) | 1))))}
                                v62 := GenS4{F0: 2, F1: (v29 + v29)}
                                v61 = append(v61, -197)
                                _ = fi60
                                _ = v61
                                _ = v62
                            }
                            _ = fi55
                            _ = v56
                        }
                        _ = v54
                    } else {
                        fmt.Println(v15, v29, v3)
                        fmt.Println(v29, v3)
                        var v63 uint16 = 32767
                        _ = v63
                    }
                    switch 1 {
                    case 4294967295:
                        var v64 uint = uint(v2)
                        _ = v64
                    default:
                        v65 := GenS6{F0: genFn1(), F1: 255, F2: -1}
                        _ = v65
                    }
                    if false {
                        delete(v1, -995)
                    } else {
                        var v66 int8 = -50
                        for fi67 := 0; fi67 < 19; fi67++ {
                            if true {
                                v68 := []int{(-(fi67))}
                                _ = v68
                            }
                            v29 = 2
                            if (!true) {
                                switch "\n\t" {
                                case "":
                                    for fi69 := 0; fi69 < 12; fi69++ {
                                        var v70 bool = (!true)
                                        _ = fi69
                                        _ = v70
                                    }
                                case "a":
                                    v15 -= 0
                                default:
                                    v71 := []int{(^(fi67)), fi67, fi67, (fi67 | (fi67 >> ((4294967295) & 63)))}
                                    _ = v71
                                }
                                delete(v1, (-(fi67)))
                                var v72 int16 = 42
                                _ = v72
                            } else {
                                if (1 <= -1) {
                                    v73 := v1[fi67]
                                    v74 := GenS2{F0: (^(v66))}
                                    _ = v73
                                    _ = v74
                                }
                            }
                            _ = fi67
                        }
                        _ = v66
                    }
                }
            case "\n\t":
                var v75 int = -9223372036854775808
                _ = v75
            }
        case 127:
            v76 := GenS3{F0: -2147483648, F1: "a"}
            _ = v76
        case -74:
            v2 = (v29 ^ 1)
        default:
            v77 := GenS7{F0: 2, F1: v2}
            _ = v77
        }
        v78 := GenS8{F0: -80, F1: true}
        _ = v78
    } else {
        v15 = v15
        var v79 string = "\n\t"
        _ = v79
    }
    fmt.Println(v3, v15, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v2
    _ = v29
    _ = v3
    _ = v30
}
