package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 uint }
type GenS2 struct { F0 uint16; F1 uint64; F2 int64 }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 int32; F1 string; F2 bool }
type GenS5 struct { F0 uint32; F1 string; F2 uint16 }
type GenS6 struct { F0 uint64; F1 bool }
type GenS7 struct { F0 int8; F1 int64; F2 uint8 }
type GenS8 struct { F0 int; F1 uint16; F2 int16 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int8; F1 string }
func genFn0(p0 bool, p1 uint8) string {
    p0 = p0
    var v6 float32 = -2.25
    if true {
        v6 = 1.5
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v6
    return ""
}
func genFn1(p0 uint64) uint64 {
    var v7 uint32 = 2147483647
    fmt.Println(v7)
    _ = p0
    _ = v7
    return (p0 << ((4294967295) & 63))
}

func main() {
    v1 := []string{""}
    var v2 uint32 = 1
    v3 := make(map[int]int)
    var v4 bool = false
    var v5 string = "foo bar"
    if (2 >= 0) {
        switch 0 {
        case 2147483647:
            v3[-1] = 179
        case -1:
            for fi8 := 0; fi8 < 5; fi8++ {
                var v9 bool = v4
                var v10 bool = ((!(1 <= 42)) && true)
                var v11 int32 = 0
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
        }
    } else {
        var v12 uint = 18446744073709551615
        _ = v12
    }
    v13 := GenS1{F0: 1, F1: -1, F2: 18446744073709551615}
    var v14 int32 = -1
    if (339 >= 32767) {
        v14 = 0
        var v15 bool = v4
        _ = v15
    }
    delete(v3, 386)
    for fi16 := 0; fi16 < 17; fi16++ {
        v4 = (!true)
        _ = fi16
    }
    for fi17 := 0; fi17 < 1; fi17++ {
        v2 = (^(v2))
        v4 = true
        _ = fi17
    }
    for fi18 := 0; fi18 < 14; fi18++ {
        if (1 < 2147483647) {
            v5 = "foo bar"
            if len(v1) > 0 {
                v1[0] = (v5 + v5)
            }
        } else {
            v5 = "a"
            switch (fi18 * 0) {
            case 2147483647:
                if v4 {
                    v2 = (v2 * (v2 ^ 2147483647))
                    var v19 int = int(v14)
                    for fi20 := 0; fi20 < 11; fi20++ {
                        if true {
                            v14 -= (v14 | v14)
                            v2 -= (v2 | v2)
                            if v4 {
                                v19 = fi20
                                if true {
                                    var v21 int64 = 42
                                    v22 := GenS1{F0: -92, F1: -32, F2: 1}
                                    delete(v3, v19)
                                    _ = v21
                                    _ = v22
                                } else {
                                    v23 := GenS1{F0: -128, F1: -128, F2: 977}
                                    v19 = v19
                                    _ = v23
                                }
                            } else {
                                switch 899 {
                                case 128:
                                    var v24 int = fi20
                                    _ = v24
                                }
                                v25 := v3[fi20]
                                _ = v25
                            }
                        }
                        v14 = (v14 ^ (v14 / (((v14 >> ((979) & 63))) | 1)))
                        if v4 {
                            var v26 int32 = 2147483647
                            v27 := GenS1{F0: 127, F1: -128, F2: 319}
                            v28 := GenS1{F0: 1, F1: 121, F2: 0}
                            _ = v26
                            _ = v27
                            _ = v28
                        } else {
                            if v4 {
                                v5 = "a"
                                v2 += 366
                                var v29 int16 = 0
                                _ = v29
                            }
                            v5 = (v5 + v5)
                        }
                        _ = fi20
                    }
                    _ = v19
                } else {
                    fmt.Println(v2, v14, fi18)
                    fmt.Println(v2, v4)
                }
            default:
                if (-1 == -9223372036854775808) {
                    if v4 {
                        for fi30 := 0; fi30 < 13; fi30++ {
                            v31 := GenS2{F0: 65535, F1: 18446744073709551615, F2: -1}
                            v14 = 74
                            v32 := GenS3{F0: v4, F1: v2}
                            _ = fi30
                            _ = v31
                            _ = v32
                        }
                        switch (v5 + v5) {
                        case "\n\t":
                            v33 := GenS4{F0: v14, F1: (v5 + (v5 + "foo bar")), F2: (v4 || (!(v4 && (!v4))))}
                            _ = v33
                        case "hello":
                            for fi34 := 0; fi34 < 20; fi34++ {
                                var v35 uint8 = 0
                                _ = fi34
                                _ = v35
                            }
                        }
                        v5 = "a"
                    }
                } else {
                    v14 -= (v14 % (((v14 % (((v14 + (v14 ^ 2147483647))) | 1))) | 1))
                    if false {
                        v2 -= v2
                        if len(v1) > 0 {
                            v1[0] = (v5 + v5)
                        }
                        if (42 >= 0) {
                            v2 += 1
                        } else {
                            var v36 int16 = 208
                            var v37 int64 = 42
                            _ = v36
                            _ = v37
                        }
                    }
                }
            }
            var v38 int8 = 42
            _ = v38
        }
        _ = fi18
    }
    if false {
        for fi39 := 0; fi39 < 16; fi39++ {
            switch (fi39 & fi39) {
            case 32767:
                fmt.Println(v4, v14)
            case 359:
                switch (v5 + v5) {
                case "":
                    if (!v4) {
                        var v40 uint = 255
                        switch (-(fi39)) {
                        case 127:
                            for fi41 := 0; fi41 < 14; fi41++ {
                                var v42 uint64 = uint64(v40)
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = fi41
                                _ = v42
                            }
                        case -825:
                            if v4 {
                                v5 = ("" + "\n\t")
                            } else {
                                v43 := GenS5{F0: (v2 + 4294967295), F1: "hello", F2: 0}
                                switch fi39 {
                                case 330:
                                    for fi44 := 0; fi44 < 1; fi44++ {
                                        fmt.Println(v14, fi39, fi44)
                                        v45 := GenS5{F0: 969, F1: (v5 + v5), F2: 255}
                                        fmt.Println(fi39, v14, v4)
                                        _ = fi44
                                        _ = v45
                                    }
                                case 2147483647:
                                    fmt.Println(v5)
                                }
                                var v46 uint = 0
                                _ = v43
                                _ = v46
                            }
                        }
                        _ = v40
                    } else {
                        var v47 uint64 = 42
                        _ = v47
                    }
                }
            case -381:
                v14 = v14
            default:
                fmt.Println(v2, v4)
            }
            v48 := GenS6{F0: 2, F1: false}
            _ = fi39
            _ = v48
        }
    } else {
        fmt.Println(v14, v4)
        var v49 uint16 = 255
        _ = len(v3)
        _ = v49
    }
    v5 = v5
    var v50 uint32 = v2
    for fi51 := 0; fi51 < 19; fi51++ {
        if v4 {
            var v52 uint32 = v50
            for fi53 := 0; fi53 < 10; fi53++ {
                var v54 uint16 = 42
                if v4 {
                    if v4 {
                        fmt.Println(v54)
                        fmt.Println(v14)
                    }
                    var v55 uint32 = (v50 & (v2 + (v52 % ((v50) | 1))))
                    _ = v55
                } else {
                    v2 = 2147483647
                    v56 := GenS6{F0: 2, F1: v4}
                    _ = v56
                }
                var v57 float64 = 3.14159
                _ = fi53
                _ = v54
                _ = v57
            }
            _ = v52
        } else {
            switch (fi51 - fi51) {
            case 1:
                _ = len(v3)
            }
            if false {
                v58 := GenS7{F0: -124, F1: 1, F2: 1}
                var v59 int32 = (^(v14))
                v60 := GenS8{F0: fi51, F1: 100, F2: 0}
                _ = v58
                _ = v59
                _ = v60
            } else {
                var v61 int32 = v14
                v61 = -2147483648
                v5 = v5
                _ = v61
            }
            if v4 {
                for fi62 := 0; fi62 < 9; fi62++ {
                    switch fi62 {
                    case -928:
                        if false {
                            v1 = append(v1, v5)
                            v5 = v5
                            fmt.Println(fi62)
                        } else {
                            v4 = true
                            _ = len(v1)
                            switch (fi51 % ((fi51) | 1)) {
                            case 32767:
                                v63 := v3[127]
                                _ = v63
                            case 2147483647:
                                if true {
                                    switch (("foo bar" + ("hello" + (((((v5 + "foo bar") + "foo bar") + v5) + (v5 + ("foo bar" + (v5 + v5)))) + "foo bar"))) + v5) {
                                    case "":
                                        var v64 uint = 0
                                        _ = v64
                                    case "a":
                                        v65 := GenS6{F0: 42, F1: (!false)}
                                        _ = v65
                                    }
                                    v14 = (v14 >> ((4294967295) & 63))
                                }
                            default:
                                var v66 bool = true
                                _ = v66
                            }
                        }
                    default:
                        switch (v5 + (v5 + v5)) {
                        case "\n\t":
                            switch v5 {
                            case "a":
                                _ = len(v1)
                            }
                        case "":
                            v4 = (((fi51 << ((255) & 63)) < fi51) && ((0.1 / 1.0) != 0.0))
                        }
                    }
                    if false {
                        var v67 int = fi62
                        var v68 int = (fi62 ^ fi62)
                        var v69 int = v67
                        _ = v67
                        _ = v68
                        _ = v69
                    } else {
                        v1 = append(v1, v5)
                    }
                    fmt.Println(v4)
                    var v70 int32 = (v14 | (v14 >> ((18446744073709551615) & 63)))
                    _ = fi62
                    _ = v70
                }
                for fi71 := 0; fi71 < 4; fi71++ {
                    if len(v1) > 0 {
                        v1[0] = v5
                    }
                    _ = fi71
                }
            } else {
                var v72 int8 = -67
                switch "foo bar" {
                case "a":
                    if v4 {
                        v73 := v3[(fi51 - fi51)]
                        if v4 {
                            fmt.Println(v14, fi51)
                        }
                        _ = v73
                    }
                }
                v74 := GenS9{F0: 65535}
                _ = v72
                _ = v74
            }
        }
        var v75 int64 = -9223372036854775808
        _ = fi51
        _ = v75
    }
    v76 := GenS10{F0: 127, F1: v5}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v76
}
