package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint32; F1 int64; F2 uint16 }
type GenS3 struct { F0 uint; F1 int; F2 string }
type GenS4 struct { F0 uint; F1 int8 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int8; F1 int32 }
type GenS8 struct { F0 int8; F1 uint64; F2 uint32 }
type GenS9 struct { F0 int; F1 uint }
func genFn0(p0 int) int32 {
    if false {
        v6 := GenS1{F0: 1000, F1: 65535, F2: 0}
        _ = v6
    }
    for fi7 := 0; fi7 < 8; fi7++ {
        p0 = fi7
        for fi8 := 0; fi8 < 18; fi8++ {
            p0 = (fi8 | (fi7 << ((255) & 63)))
            var v9 uint16 = 255
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    p0 -= (p0 * p0)
    v10 := GenS2{F0: 4294967295, F1: -804, F2: 2}
    v11 := []int{-2147483648}
    _ = p0
    _ = v10
    _ = v11
    return 42
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{}
    v3 := make(map[int]string)
    v4 := []int{}
    var v5 int64 = -138
    for fi12 := 0; fi12 < 4; fi12++ {
        for fi13 := 0; fi13 < 5; fi13++ {
            v14 := GenS1{F0: 0, F1: 2, F2: 874}
            v5 -= v5
            _ = fi13
            _ = v14
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi12
    }
    delete(v3, -842)
    if (((1000 == -102) || (!(!("\n\t" < "foo bar")))) || (!false)) {
        fmt.Println(v5)
        if len(v4) > 0 {
            v4[0] = 4294967295
        }
        if true {
            v15 := GenS1{F0: -1, F1: 42, F2: 2}
            var v16 int16 = -1
            v16 = 1
            _ = v15
            _ = v16
        } else {
            v5 += (v5 ^ (v5 ^ v5))
        }
    }
    switch -9223372036854775808 {
    case 127:
        v5 -= (v5 + v5)
    case 9223372036854775807:
        v17 := GenS2{F0: 677, F1: -50, F2: 42}
        _ = v17
    case -603:
        v18 := GenS3{F0: 812, F1: 905, F2: ("foo bar" + "\n\t")}
        _ = v18
    }
    if (!true) {
        v5 = (-(v5))
    } else {
        switch 9223372036854775807 {
        case 2147483647:
            if true {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            } else {
                v19 := GenS4{F0: 42, F1: 42}
                if (-1 >= -128) {
                    v20 := GenS1{F0: -32768, F1: 1, F2: 67}
                    var v21 bool = false
                    _ = v20
                    _ = v21
                }
                _ = v19
            }
        }
    }
    if true {
        fmt.Println(v5)
        if (false && false) {
            if false {
                for fi22 := 0; fi22 < 8; fi22++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi22
                }
                switch 1 {
                case 2147483647:
                    var v23 int = 128
                    _ = v23
                case 32768:
                    fmt.Println(v5)
                case 4294967295:
                    for fi24 := 0; fi24 < 20; fi24++ {
                        fmt.Println(v5, fi24)
                        var v25 int8 = 0
                        _ = fi24
                        _ = v25
                    }
                }
            }
            for fi26 := 0; fi26 < 12; fi26++ {
                v5 -= (v5 ^ 1)
                v5 += v5
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi26
            }
            switch 9223372036854775807 {
            case -743:
                v5 += v5
            case 32768:
                v27 := GenS3{F0: 4294967295, F1: -175, F2: ("foo bar" + ("" + "a"))}
                _ = v27
            case 274:
                if false {
                    var v28 int32 = 42
                    var v29 uint32 = 4294967295
                    _ = v28
                    _ = v29
                } else {
                    v30 := GenS5{F0: 2147483648}
                    v31 := GenS6{F0: false}
                    v32 := GenS4{F0: 433, F1: 127}
                    _ = v30
                    _ = v31
                    _ = v32
                }
            }
        }
    }
    if (-758 < 0) {
        for fi33 := 0; fi33 < 14; fi33++ {
            var v34 int = (fi33 * (fi33 / (((fi33 % (((fi33 - 0)) | 1))) | 1)))
            fmt.Println(v5, v34, v34)
            v35 := GenS1{F0: 338, F1: 562, F2: 2}
            _ = fi33
            _ = v34
            _ = v35
        }
    }
    switch 371 {
    case 2147483647:
        v5 = v5
    case -380:
        switch ((("foo bar" + "") + "\n\t") + ("foo bar" + "\n\t")) {
        case "\n\t":
            var v36 int8 = 127
            _ = v36
        case "foo bar":
            switch "" {
            case "foo bar":
                fmt.Println(v5)
            case "a":
                for fi37 := 0; fi37 < 6; fi37++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi37
                }
            case "\n\t":
                v5 -= v5
            default:
                for fi38 := 0; fi38 < 18; fi38++ {
                    for fi39 := 0; fi39 < 13; fi39++ {
                        if true {
                            var v40 int8 = 0
                            delete(v1, 1)
                            if (67 == 255) {
                                v5 = (v5 << ((42) & 63))
                            } else {
                                if (18446744073709551615 >= 1) {
                                    if false {
                                        v40 = (v40 / ((114) | 1))
                                        switch "\n\t" {
                                        case "foo bar":
                                            if ((^(fi39)) < (fi39 - fi38)) {
                                                var v41 int64 = (v5 | (v5 + (v5 >> ((167) & 63))))
                                                _ = v41
                                            }
                                        case "\n\t":
                                            v42 := GenS7{F0: v40, F1: 42}
                                            _ = v42
                                        }
                                        switch "foo bar" {
                                        case "a":
                                            _ = len(v1)
                                        case "\n\t":
                                            v5 = (v5 << ((2) & 63))
                                        case "hello":
                                            v5 -= v5
                                        }
                                    } else {
                                        if (!false) {
                                            v43 := GenS2{F0: 1, F1: (-(v5)), F2: 42}
                                            _ = v43
                                        }
                                        if true {
                                            if ((("\n\t" + "foo bar") + "a") >= ("" + ("a" + ""))) {
                                                fmt.Println(fi38)
                                            } else {
                                                if (0 >= (^(v40))) {
                                                    v5 += (v5 % ((v5) | 1))
                                                    v40 = (v40 / ((v40) | 1))
                                                    v40 -= v40
                                                }
                                                v3[(fi39 % ((fi38) | 1))] = "\n\t"
                                            }
                                            var v44 string = "a"
                                            _ = v44
                                        }
                                    }
                                    _ = len(v2)
                                    var v45 uint64 = 18446744073709551615
                                    _ = v45
                                }
                            }
                            fmt.Println(v40)
                            var v46 uint8 = 22
                            v1[0] = int(v40)
                            _ = v40
                            _ = v46
                        }
                        if len(v2) > 0 {
                            v2[0] = ("hello" + (("hello" + "\n\t") + "\n\t"))
                        }
                        fmt.Println(fi39)
                        _ = fi39
                    }
                    _ = fi38
                }
            }
        default:
            switch 42 {
            case -786:
                var v47 uint64 = 502
                _ = v47
            default:
                switch 882 {
                case 638:
                    v48 := GenS8{F0: -128, F1: 753, F2: 0}
                    _ = v48
                case 32767:
                    switch "hello" {
                    case "foo bar":
                        fmt.Println(v5)
                    case "":
                        var v49 uint16 = 2
                        _ = v49
                    }
                }
            }
        }
    default:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    for fi50 := 0; fi50 < 17; fi50++ {
        var v51 int16 = 0
        if false {
            _ = len(v3)
            switch "foo bar" {
            case "hello":
                var v52 uint16 = 32767
                _ = v52
            case "":
                v51 -= int16(fi50)
            case "foo bar":
                if true {
                    for fi53 := 0; fi53 < 15; fi53++ {
                        v51 += (-(v51))
                        _ = fi53
                    }
                    var v54 float64 = (3.14159 / 0.1)
                    _ = v54
                }
            }
        } else {
            for fi55 := 0; fi55 < 13; fi55++ {
                if len(v2) > 0 {
                    v2[0] = (("hello" + "") + ("a" + ("hello" + "\n\t")))
                }
                v5 += v5
                _ = fi55
            }
            v51 -= v51
            v51 -= (v51 >> ((4294967295) & 63))
        }
        _ = fi50
        _ = v51
    }
    v56 := GenS9{F0: 32767, F1: 935}
    if true {
        var v57 int = 760
        switch v57 {
        case 9223372036854775807:
            var v58 uint8 = 128
            _ = v58
        default:
            fmt.Println(v5)
        }
        _ = v57
    } else {
        for fi59 := 0; fi59 < 4; fi59++ {
            var v60 uint = 42
            switch (((("\n\t" + "hello") + "\n\t") + "hello") + "") {
            case "foo bar":
                v5 = (v5 & (v5 ^ v5))
            case "":
                v5 = v5
            case "a":
                var v61 int16 = 880
                _ = v61
            default:
                v5 -= (-(v5))
            }
            v4 = append(v4, (fi59 << (((v60 % ((769) | 1))) & 63)))
            _ = fi59
            _ = v60
        }
    }
    v5 = v5
    var v62 uint32 = 200
    for fi63 := 0; fi63 < 6; fi63++ {
        for fi64 := 0; fi64 < 20; fi64++ {
            v65 := GenS5{F0: v62}
            if ((2 < v62) || false) {
                _ = len(v3)
                _ = len(v3)
                var v66 int8 = -128
                _ = v66
            }
            _ = fi64
            _ = v65
        }
        _ = fi63
    }
    var v67 float32 = 1e10
    _ = len(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v62:", v62)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v62
    _ = v67
}
