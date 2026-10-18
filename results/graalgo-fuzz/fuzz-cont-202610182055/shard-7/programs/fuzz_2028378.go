package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 string; F2 uint }
type GenS3 struct { F0 int; F1 bool; F2 bool }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint }
func genFn0(p0 bool, p1 uint8, p2 uint64) int8 {
    v7 := GenS1{F0: -290}
    for fi8 := 0; fi8 < 10; fi8++ {
        v9 := GenS1{F0: 42}
        var v10 uint8 = (p1 * p1)
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}

func main() {
    var v1 uint8 = 128
    var v2 int64 = 9223372036854775807
    var v3 uint8 = v1
    var v4 uint = 1
    v5 := make(map[uint32]bool)
    v6 := make(map[string]bool)
    delete(v6, ("foo bar" + "foo bar"))
    for fi11 := 0; fi11 < 8; fi11++ {
        var v12 int = fi11
        for fi13 := 0; fi13 < 7; fi13++ {
            v1 -= v1
            var v14 string = "foo bar"
            if true {
                v15 := []int{42, int(v3), 32768}
                var v16 uint32 = 637
                if true {
                    var v17 float32 = 0.0
                    var v18 int32 = -1
                    _ = v17
                    _ = v18
                } else {
                    switch v14 {
                    case "a":
                        v4 += (v4 - (v4 >> (((v4 / ((1) | 1))) & 63)))
                    }
                    if true {
                        switch fi11 {
                        case 2147483647:
                            fmt.Println(v14, v1, fi11)
                        case -824:
                            if false {
                                var v19 string = v14
                                switch (fi11 >> ((v4) & 63)) {
                                case 186:
                                    v4 = v4
                                case 212:
                                    if false {
                                        if (v16 > v16) {
                                            var v20 string = ((v19 + "hello") + (v14 + v19))
                                            switch v14 {
                                            case "":
                                                if false {
                                                    switch (v14 + v14) {
                                                    case "hello":
                                                        v14 = (v14 + (v14 + ""))
                                                    }
                                                    switch (v20 + ((((v19 + (v19 + v19)) + "foo bar") + v20) + (v14 + v20))) {
                                                    case "\n\t":
                                                        switch (v19 + ((v19 + v19) + v20)) {
                                                        case "a":
                                                            fmt.Println(v1, v19, fi11)
                                                        case "\n\t":
                                                            var v21 int32 = 42
                                                            _ = v21
                                                        case "hello":
                                                            var v22 int64 = (-(v2))
                                                            _ = v22
                                                        default:
                                                            v3 = (v3 + v3)
                                                        }
                                                    }
                                                }
                                            default:
                                                v2 = -585
                                            }
                                            _ = v20
                                        } else {
                                            var v23 int32 = -2147483648
                                            v14 = (v14 + v19)
                                            if true {
                                                if true {
                                                    v6[v19] = (!(217 > v16))
                                                    var v24 uint8 = v3
                                                    _ = v24
                                                }
                                            } else {
                                                var v25 uint8 = (^(v1))
                                                if false {
                                                    if false {
                                                        if true {
                                                            var v26 int64 = v2
                                                            _ = v26
                                                        }
                                                        if len(v15) > 0 {
                                                            v15[0] = 725
                                                        }
                                                    } else {
                                                        if (true || ((1 >= 705) || ((18446744073709551615 != (v4 % (((v4 / ((v4) | 1))) | 1))) || (!(!(v23 >= v23)))))) {
                                                            v27 := GenS1{F0: fi11}
                                                            var v28 int8 = -7
                                                            _ = v27
                                                            _ = v28
                                                        }
                                                        v5[v16] = ((0.0 != 1e10) && (!false))
                                                        _ = len(v15)
                                                    }
                                                    if true {
                                                        v29 := GenS1{F0: 128}
                                                        v30 := GenS2{F0: fi13, F1: "a", F2: (^(v4))}
                                                        v19 = "a"
                                                        _ = v29
                                                        _ = v30
                                                    } else {
                                                        var v31 int64 = -123
                                                        _ = v31
                                                    }
                                                    if true {
                                                        v12 = fi11
                                                        v3 += 127
                                                    } else {
                                                        var v32 uint8 = v3
                                                        v25 = v1
                                                        v33 := GenS1{F0: 0}
                                                        _ = v32
                                                        _ = v33
                                                    }
                                                    v34 := GenS2{F0: int(v2), F1: ("" + ("hello" + "foo bar")), F2: (v4 ^ (v4 & v4))}
                                                    _ = v34
                                                }
                                                _ = v25
                                            }
                                            _ = v23
                                        }
                                        fmt.Println(v2, v1, v19)
                                    } else {
                                        var v35 float32 = -2.25
                                        var v36 bool = true
                                        v16 += (v16 - (v16 >> (((v4 | v4)) & 63)))
                                        if false {
                                            v16 -= v16
                                            v3 = uint8(v4)
                                            v15 = append(v15, 32767)
                                        } else {
                                            fmt.Println(v4, v4, v19)
                                        }
                                        var v37 float32 = v35
                                        v36 = (!(v36 || (!v36)))
                                        fmt.Println(v12)
                                        v38 := GenS3{F0: 2147483647, F1: true, F2: v36}
                                        var v39 float32 = (v37 - v35)
                                        _ = v35
                                        _ = v36
                                        _ = v37
                                        _ = v38
                                        _ = v39
                                    }
                                case 708:
                                    var v40 int16 = 32767
                                    _ = v40
                                }
                                v41 := GenS1{F0: (v12 ^ (fi11 & (-(v12))))}
                                _ = v19
                                _ = v41
                            }
                        case 184:
                            var v42 int8 = -128
                            _ = v42
                        }
                        if true {
                            if false {
                                v14 = v14
                                v14 = ((("hello" + v14) + "a") + "a")
                            }
                            var v43 int64 = 9223372036854775807
                            _ = v43
                        }
                        v44 := GenS1{F0: fi13}
                        v2 = v2
                        _ = v44
                    }
                }
                _ = v15
                _ = v16
            }
            _ = fi13
            _ = v14
        }
        v45 := v6["a"]
        _ = fi11
        _ = v12
        _ = v45
    }
    v46 := GenS2{F0: 464, F1: "\n\t", F2: 42}
    var v47 bool = false
    v48 := GenS4{F0: 950}
    v49 := []int{}
    delete(v5, 4294967295)
    var v50 uint16 = 992
    if true {
        switch 32767 {
        case 630:
            if (v2 >= (v2 + v2)) {
                for fi51 := 0; fi51 < 6; fi51++ {
                    var v52 uint16 = 42
                    var v53 float32 = float32(0.1)
                    _ = fi51
                    _ = v52
                    _ = v53
                }
            } else {
                if v47 {
                    v54 := GenS5{F0: 42}
                    v3 -= v1
                    _ = v54
                } else {
                    for fi55 := 0; fi55 < 12; fi55++ {
                        var v56 float64 = 0.0
                        _ = fi55
                        _ = v56
                    }
                    for fi57 := 0; fi57 < 16; fi57++ {
                        var v58 int32 = -2147483648
                        _ = fi57
                        _ = v58
                    }
                    for fi59 := 0; fi59 < 5; fi59++ {
                        v4 += v4
                        if (false || v47) {
                            v60 := v5[uint32(v4)]
                            v50 = (v50 | (v50 % (((v50 + 2)) | 1)))
                            for fi61 := 0; fi61 < 17; fi61++ {
                                var v62 int = fi59
                                _ = fi61
                                _ = v62
                            }
                            _ = v60
                        } else {
                            v3 = (^(v3))
                            if (v4 <= v4) {
                                v47 = true
                                v1 = 1
                                for fi63 := 0; fi63 < 6; fi63++ {
                                    fmt.Println(v2)
                                    v64 := GenS2{F0: 649, F1: "", F2: 480}
                                    var v65 int16 = -976
                                    if (-2.25 <= -2.25) {
                                        _ = len(v5)
                                        fmt.Println(v2)
                                        v3 -= v1
                                    }
                                    var v66 int = fi63
                                    _ = fi63
                                    _ = v64
                                    _ = v65
                                    _ = v66
                                }
                            }
                        }
                        fmt.Println(v4, v1, v2)
                        _ = fi59
                    }
                }
                var v67 int64 = -504
                _ = v67
            }
        case 4294967295:
            var v68 uint32 = 969
            _ = v68
        }
        fmt.Println(v3, v1, v47)
    } else {
        if len(v49) > 0 {
            v49[0] = 32767
        }
        switch ("a" + (("" + "foo bar") + ("foo bar" + "a"))) {
        case "\n\t":
            for fi69 := 0; fi69 < 14; fi69++ {
                var v70 uint64 = 18446744073709551615
                var v71 int64 = (v2 - v2)
                fmt.Println(v4, v3, v2)
                _ = fi69
                _ = v70
                _ = v71
            }
        default:
            var v72 uint16 = (v50 << ((v4) & 63))
            _ = v72
        }
    }
    v2 = (v2 * 617)
    switch 42 {
    case -401:
        var v73 int32 = 42
        _ = v73
    default:
        var v74 int8 = 42
        _ = v74
    }
    fmt.Println(v47)
    switch 981 {
    case 127:
        v4 -= 4294967295
    case -417:
        switch "\n\t" {
        case "a":
            v1 = (v1 << (((^(v4))) & 63))
        case "hello":
            if v47 {
                v2 = (v2 / ((v2) | 1))
                switch ("foo bar" + (("\n\t" + "") + ("foo bar" + "hello"))) {
                case "a":
                    switch "a" {
                    case "":
                        var v75 uint16 = 2
                        _ = v75
                    case "\n\t":
                        var v76 uint32 = 45
                        _ = v76
                    case "a":
                        if len(v49) > 0 {
                            v49[0] = -2147483648
                        }
                    default:
                        fmt.Println(v47, v1, v1)
                    }
                default:
                    for fi77 := 0; fi77 < 13; fi77++ {
                        var v78 int16 = -1
                        var v79 int = fi77
                        _ = fi77
                        _ = v78
                        _ = v79
                    }
                }
            } else {
                var v80 uint = 65535
                var v81 float64 = float64(3.14159)
                v3 += 127
                _ = v80
                _ = v81
            }
        case "":
            var v82 uint32 = 2
            _ = v82
        }
    }
    var v83 float32 = 1.5
    switch ("a" + (("a" + "\n\t") + "hello")) {
    case "a":
        var v84 float64 = (-2.25 - float64(v83))
        _ = v84
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v49:", len(v49), v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v6
    _ = v83
}
