package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint8; F1 int16 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int; F1 int32; F2 int16 }
type GenS9 struct { F0 uint }
func genFn0(p0 uint16) string {
    p0 -= p0
    for fi7 := 0; fi7 < 1; fi7++ {
        switch fi7 {
        case 42:
            if (156 != 182) {
                p0 = (p0 | (p0 << ((4294967295) & 63)))
                p0 += 265
            }
        case 443:
            v8 := GenS1{F0: -85, F1: ((!(!((-1 != 1) && (18446744073709551615 == 324)))) && true), F2: (fi7 & 32768)}
            _ = v8
        default:
            v9 := []int{fi7}
            _ = v9
        }
        _ = fi7
    }
    p0 = 108
    var v10 float32 = 0.0
    _ = p0
    _ = v10
    return ("" + ("" + ("foo bar" + "a")))
}

func main() {
    var v1 string = ""
    var v2 int8 = 42
    var v3 uint = 18446744073709551615
    var v4 uint32 = 1
    var v5 uint = v3
    v6 := make(map[string]string)
    for fi11 := 0; fi11 < 20; fi11++ {
        for fi12 := 0; fi12 < 3; fi12++ {
            var v13 uint32 = (v4 >> ((65535) & 63))
            fmt.Println(v3, fi12)
            var v14 int8 = (v2 | (v2 << ((v3) & 63)))
            switch "hello" {
            case "foo bar":
                if false {
                    if (!(false || (-943 <= 138))) {
                        v4 -= v13
                        v15 := GenS2{F0: 42}
                        var v16 float64 = 3.14159
                        _ = v15
                        _ = v16
                    }
                    fmt.Println(fi12, v4, v14)
                } else {
                    if (-1 != -1) {
                        var v17 int32 = -825
                        _ = v17
                    } else {
                        var v18 float32 = -2.25
                        v3 += (v3 % ((v3) | 1))
                        _ = v18
                    }
                    if true {
                        v19 := GenS2{F0: -1}
                        fmt.Println(fi11, v3, fi11)
                        _ = v19
                    }
                    v5 = (v5 / ((v3) | 1))
                }
            case "hello":
                v20 := GenS3{F0: (v4 | (v13 & v4)), F1: false}
                _ = v20
            default:
                if ("a" >= v1) {
                    if (!false) {
                        v3 += (v5 & v3)
                    }
                    v13 += (v13 << ((42) & 63))
                    var v21 bool = true
                    _ = v21
                } else {
                    v22 := GenS4{F0: 53}
                    v23 := GenS5{F0: 0, F1: -32768}
                    _ = v22
                    _ = v23
                }
            }
            v4 = v4
            _ = fi12
            _ = v13
            _ = v14
        }
        switch v1 {
        case "\n\t":
            var v24 int64 = 9223372036854775807
            _ = v24
        case "":
            var v25 int64 = -943
            _ = v25
        case "foo bar":
            if true {
                if (!((v1 + (("\n\t" + (("hello" + "foo bar") + "hello")) + (((v1 + "") + v1) + ""))) != v1)) {
                    if (false || true) {
                        v26 := GenS2{F0: -570}
                        _ = v26
                    } else {
                        v27 := GenS3{F0: 4294967295, F1: (!((fi11 >= int(v2)) || (1000 < 32767)))}
                        var v28 uint32 = v4
                        for fi29 := 0; fi29 < 3; fi29++ {
                            if (false || false) {
                                if (!false) {
                                    var v30 uint = v3
                                    var v31 string = "foo bar"
                                    _ = v30
                                    _ = v31
                                }
                                fmt.Println(fi29, v5)
                            } else {
                                v28 = (v4 << ((v3) & 63))
                                switch 133 {
                                case 42:
                                    var v32 uint16 = uint16(v28)
                                    _ = v32
                                case -1:
                                    if (true && ((^(v3)) >= (v5 % ((v3) | 1)))) {
                                        v3 += (v5 + v5)
                                    }
                                case 128:
                                    var v33 int = (fi29 >> (((^(v5))) & 63))
                                    _ = v33
                                }
                            }
                            if ((v5 >> (((v5 % ((v3) | 1))) & 63)) <= 770) {
                                v34 := GenS5{F0: 42, F1: -492}
                                if (false && true) {
                                    if true {
                                        var v35 int8 = (v2 * 100)
                                        _ = v35
                                    }
                                }
                                switch (^(fi29)) {
                                case 123:
                                    v36 := GenS2{F0: -157}
                                    _ = v36
                                case 4294967295:
                                    v28 = v28
                                }
                                _ = v34
                            }
                            _ = fi29
                        }
                        _ = v27
                        _ = v28
                    }
                    v37 := GenS6{F0: v4}
                    v38 := GenS1{F0: 1, F1: false, F2: fi11}
                    _ = v37
                    _ = v38
                }
                v1 = v1
                switch 797 {
                case 226:
                    var v39 int8 = (v2 % ((v2) | 1))
                    _ = v39
                case 32767:
                    v40 := GenS6{F0: v4}
                    _ = v40
                case 549:
                    for fi41 := 0; fi41 < 5; fi41++ {
                        v42 := GenS2{F0: 2147483647}
                        v1 = ((v1 + ("a" + (v1 + "a"))) + v1)
                        var v43 int64 = -62
                        var v44 uint32 = v4
                        _ = fi41
                        _ = v42
                        _ = v43
                        _ = v44
                    }
                }
            }
        default:
            var v45 uint = (v3 << ((v3) & 63))
            _ = v45
        }
        if true {
            fmt.Println(v4, v1, fi11)
            switch "hello" {
            case "foo bar":
                v4 += (v4 / (((^(v4))) | 1))
            case "a":
                v2 = v2
            case "hello":
                v46 := GenS7{F0: 1}
                _ = v46
            default:
                fmt.Println(v5)
            }
            var v47 int16 = 1000
            _ = v47
        }
        _ = fi11
    }
    fmt.Println(v3)
    for fi48 := 0; fi48 < 8; fi48++ {
        v49 := GenS8{F0: (fi48 ^ fi48), F1: -694, F2: int16(v2)}
        var v50 int16 = 42
        _ = fi48
        _ = v49
        _ = v50
    }
    switch 9223372036854775807 {
    case 1:
        v1 = (v1 + "foo bar")
    }
    for fi51 := 0; fi51 < 13; fi51++ {
        switch (^(fi51)) {
        case 42:
            for fi52 := 0; fi52 < 13; fi52++ {
                v53 := []int{fi51, fi51}
                v5 -= (v3 >> (((v5 * v3)) & 63))
                var v54 int16 = 302
                _ = fi52
                _ = v53
                _ = v54
            }
        case 0:
            switch "\n\t" {
            case "a":
                if false {
                    var v55 uint32 = v4
                    fmt.Println(fi51, v5)
                    _ = v55
                } else {
                    switch (("a" + "foo bar") + v1) {
                    case "\n\t":
                        switch fi51 {
                        case 32767:
                            v56 := GenS7{F0: -2147483648}
                            _ = v56
                        case 0:
                            v1 = (v1 + v1)
                        case 127:
                            if ((^(v5)) >= uint(v4)) {
                                v5 = 255
                            }
                        }
                    default:
                        fmt.Println(fi51)
                    }
                    switch -98 {
                    case 32768:
                        var v57 bool = (2 >= 0)
                        _ = v57
                    case 128:
                        for fi58 := 0; fi58 < 17; fi58++ {
                            v6["foo bar"] = ""
                            _ = len(v6)
                            _ = fi58
                        }
                    case -126:
                        for fi59 := 0; fi59 < 14; fi59++ {
                            var v60 int = fi59
                            v60 = fi59
                            _ = fi59
                            _ = v60
                        }
                    default:
                        fmt.Println(v3)
                    }
                    fmt.Println(v3, v2)
                }
            default:
                switch v1 {
                case "foo bar":
                    if true {
                        v2 = (v2 - v2)
                    }
                case "a":
                    v61 := GenS6{F0: v4}
                    _ = v61
                default:
                    var v62 uint = v3
                    _ = v62
                }
            }
        default:
            if false {
                v4 += (v4 ^ 2147483647)
                switch "" {
                case "a":
                    var v63 int = fi51
                    _ = v63
                }
                for fi64 := 0; fi64 < 14; fi64++ {
                    v3 += (^(v5))
                    switch -1 {
                    case 1:
                        _ = len(v6)
                    }
                    var v65 float64 = 0.0
                    _ = fi64
                    _ = v65
                }
            } else {
                v66 := GenS7{F0: -2147483648}
                for fi67 := 0; fi67 < 1; fi67++ {
                    v6[v1] = "foo bar"
                    _ = fi67
                }
                _ = v66
            }
        }
        _ = fi51
    }
    switch v1 {
    case "a":
        if (42 != 42) {
            switch v1 {
            case "":
                v68 := GenS9{F0: 907}
                _ = v68
            case "\n\t":
                if true {
                    switch "a" {
                    case "a":
                        v4 = 42
                    case "":
                        v69 := GenS2{F0: 42}
                        _ = v69
                    case "foo bar":
                        for fi70 := 0; fi70 < 20; fi70++ {
                            v1 = ""
                            var v71 int = fi70
                            _ = fi70
                            _ = v71
                        }
                    default:
                        fmt.Println(v1, v3, v3)
                    }
                }
            }
            v5 -= v3
            var v72 int32 = 42
            _ = v72
        } else {
            v5 = v3
        }
    case "\n\t":
        v2 -= v2
    default:
        v73 := v6["\n\t"]
        _ = v73
    }
    var v74 float32 = 0.0
    if false {
        v75 := []int{32767}
        v2 += (-(v2))
        if false {
            v74 -= v74
        }
        _ = v75
    }
    for fi76 := 0; fi76 < 1; fi76++ {
        v1 = v1
        _ = fi76
    }
    for fi77 := 0; fi77 < 2; fi77++ {
        var v78 uint = (v5 | (^(v5)))
        var v79 string = "foo bar"
        _ = fi77
        _ = v78
        _ = v79
    }
    var v80 uint = 0
    v4 = 2147483647
    var v81 string = "foo bar"
    var v82 uint8 = uint8(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v74:", v74)
    fmt.Println("v80:", v80)
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v74
    _ = v80
    _ = v81
    _ = v82
}
