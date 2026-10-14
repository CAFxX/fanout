package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 int8; F2 uint16 }
type GenS3 struct { F0 int16; F1 int8; F2 int8 }
type GenS4 struct { F0 uint64; F1 uint32 }
type GenS5 struct { F0 uint16; F1 int16 }
type GenS6 struct { F0 uint16; F1 uint32 }
type GenS7 struct { F0 int32; F1 uint }
func genFn0() int32 {
    v7 := GenS1{F0: 4294967295}
    var v8 int = 822
    switch "a" {
    case "hello":
        var v9 uint32 = 65535
        _ = v9
    case "a":
        v10 := []int{v8, v8, v8, v8}
        _ = v10
    }
    switch "\n\t" {
    case "a":
        fmt.Println(v8)
    case "\n\t":
        if ((!(1 < 4294967295)) && false) {
            for fi11 := 0; fi11 < 6; fi11++ {
                v8 = v8
                if false {
                    v12 := GenS1{F0: 2}
                    fmt.Println(fi11)
                    _ = v12
                }
                if true {
                    if true {
                        if true {
                            v13 := []int{(^(fi11)), 4294967295}
                            for fi14 := 0; fi14 < 2; fi14++ {
                                var v15 uint8 = 0
                                _ = fi14
                                _ = v15
                            }
                            _ = v13
                        }
                        v16 := GenS1{F0: 42}
                        v17 := GenS1{F0: 813}
                        _ = v16
                        _ = v17
                    } else {
                        for fi18 := 0; fi18 < 20; fi18++ {
                            var v19 int64 = -1
                            v20 := GenS2{F0: 42, F1: 0, F2: 42}
                            var v21 float64 = 1.5
                            v21 -= 1e10
                            v22 := GenS3{F0: 0, F1: -128, F2: -90}
                            _ = fi18
                            _ = v19
                            _ = v20
                            _ = v21
                            _ = v22
                        }
                        v8 = 315
                    }
                } else {
                    fmt.Println(v8)
                }
                _ = fi11
            }
        }
    case "foo bar":
        v8 = (v8 >> ((4294967295) & 63))
    default:
        for fi23 := 0; fi23 < 10; fi23++ {
            fmt.Println(v8, v8)
            _ = fi23
        }
    }
    _ = v7
    _ = v8
    return 296
}
func genFn1(p0 uint64, p1 uint64, p2 uint) uint64 {
    for fi24 := 0; fi24 < 12; fi24++ {
        p0 -= 0
        _ = fi24
    }
    v25 := []int{-935}
    fmt.Println(p0, p2, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v25
    return (p1 - (p0 % ((p1) | 1)))
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{}
    v3 := []string{"\n\t", (("foo bar" + ("foo bar" + "\n\t")) + "hello"), "\n\t", ("foo bar" + "")}
    v4 := []int{}
    var v5 float64 = (0.0 / float64(0.1))
    var v6 uint8 = 127
    for fi26 := 0; fi26 < 9; fi26++ {
        var v27 bool = (718 < 1)
        _ = fi26
        _ = v27
    }
    for fi28 := 0; fi28 < 11; fi28++ {
        fmt.Println(v6, v5, fi28)
        fmt.Println(fi28)
        _ = fi28
    }
    fmt.Println(v6)
    switch 841 {
    case -1:
        var v29 int32 = -2147483648
        _ = v29
    case 0:
        v30 := GenS4{F0: 42, F1: 493}
        _ = v30
    default:
        var v31 int16 = -94
        _ = v31
    }
    fmt.Println(v5, v5)
    for fi32 := 0; fi32 < 20; fi32++ {
        if false {
            v6 -= (^(v6))
            v6 = v6
            if (1 != 2) {
                v5 -= v5
                switch -536 {
                case 42:
                    for fi33 := 0; fi33 < 3; fi33++ {
                        switch ("\n\t" + (("\n\t" + "foo bar") + "")) {
                        case "hello":
                            var v34 float32 = (0.0 / -2.25)
                            _ = v34
                        case "\n\t":
                            var v35 uint16 = 42
                            _ = v35
                        }
                        _ = fi33
                    }
                }
            }
        } else {
            var v36 uint8 = v6
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = v36
        }
        switch ("foo bar" + "a") {
        case "hello":
            var v37 int16 = -32768
            _ = v37
        case "\n\t":
            for fi38 := 0; fi38 < 2; fi38++ {
                fmt.Println(v6, v6)
                _ = len(v1)
                switch "" {
                case "foo bar":
                    if (0 != genFn0()) {
                        v5 = v5
                        var v39 uint32 = 197
                        _ = v39
                    } else {
                        v5 -= ((v5 / v5) * v5)
                    }
                default:
                    v40 := GenS3{F0: 0, F1: 1, F2: 100}
                    _ = v40
                }
                fmt.Println(fi38)
                var v41 string = ""
                _ = fi38
                _ = v41
            }
        case "":
            for fi42 := 0; fi42 < 11; fi42++ {
                if true {
                    v43 := GenS1{F0: 18446744073709551615}
                    switch (fi32 << ((978) & 63)) {
                    case -124:
                        var v44 int64 = 141
                        _ = v44
                    case 1:
                        v45 := GenS2{F0: 0, F1: 1, F2: 255}
                        _ = v45
                    case 128:
                        fmt.Println(v6)
                    default:
                        var v46 uint64 = 1
                        _ = v46
                    }
                    var v47 float32 = 1.5
                    _ = v43
                    _ = v47
                }
                _ = fi42
            }
        }
        v6 = v6
        _ = fi32
    }
    _ = len(v1)
    switch -266 {
    case 32767:
        var v48 int8 = 127
        _ = v48
    default:
        switch ("\n\t" + "foo bar") {
        case "\n\t":
            v5 += (float64(-2.25) - ((1e10 - float64(1.5)) / v5))
        case "a":
            v5 += (1e10 - ((((v5 * (v5 / v5)) * v5) + float64(1e10)) / (0.1 + (v5 / (v5 - v5)))))
        case "foo bar":
            if (-118 < 1) {
                var v49 uint64 = 4294967295
                var v50 uint = 0
                _ = v49
                _ = v50
            }
        default:
            fmt.Println(v5, v6)
        }
    }
    var v51 uint64 = 18446744073709551615
    var v52 uint32 = 65535
    v5 += (v5 + v5)
    delete(v1, v52)
    v51 = v51
    v52 += v52
    if ((1000 >= -32768) || ("hello" != "hello")) {
        fmt.Println(v51, v5)
        v53 := GenS3{F0: -545, F1: 42, F2: -128}
        switch ("foo bar" + "a") {
        case "a":
            for fi54 := 0; fi54 < 1; fi54++ {
                switch ("" + "foo bar") {
                case "a":
                    var v55 uint = 275
                    _ = v55
                default:
                    var v56 uint32 = v52
                    _ = v56
                }
                var v57 float64 = 3.14159
                _ = fi54
                _ = v57
            }
        case "hello":
            v58 := GenS1{F0: 0}
            _ = v58
        default:
            if (983 != -2147483648) {
                v6 -= (v6 - 127)
                v5 += 0.0
                var v59 uint16 = 1
                _ = v59
            } else {
                switch 888 {
                case -89:
                    for fi60 := 0; fi60 < 15; fi60++ {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        switch fi60 {
                        case 1:
                            v52 = v52
                        case 968:
                            var v61 int32 = -2147483648
                            _ = v61
                        }
                        if (!true) {
                            var v62 uint64 = (v51 ^ (^(v51)))
                            switch "a" {
                            case "a":
                                v6 -= 106
                            case "hello":
                                var v63 int16 = int16(v51)
                                _ = v63
                            }
                            fmt.Println(v62, v5)
                            _ = v62
                        }
                        _ = fi60
                    }
                case 128:
                    var v64 bool = true
                    _ = v64
                case 9223372036854775807:
                    v65 := GenS5{F0: uint16(v51), F1: 32767}
                    _ = v65
                }
            }
        }
        _ = v53
    } else {
        var v66 uint16 = uint16(v52)
        var v67 uint8 = v6
        _ = v66
        _ = v67
    }
    switch 2147483647 {
    case 1:
        for fi68 := 0; fi68 < 2; fi68++ {
            var v69 uint = 18446744073709551615
            _ = fi68
            _ = v69
        }
    case 42:
        v5 = v5
    case 32768:
        for fi70 := 0; fi70 < 3; fi70++ {
            if true {
                if len(v4) > 0 {
                    v4[0] = 127
                }
                if false {
                    v4 = append(v4, 32768)
                    var v71 uint = 125
                    _ = v71
                } else {
                    v72 := GenS3{F0: -32768, F1: -30, F2: -47}
                    switch "" {
                    case "foo bar":
                        v52 += (v52 % ((1) | 1))
                    case "\n\t":
                        var v73 bool = (1 <= 42)
                        _ = v73
                    case "":
                        var v74 int = (fi70 ^ 0)
                        _ = v74
                    default:
                        _ = len(v4)
                    }
                    _ = v72
                }
                var v75 int = (^(fi70))
                _ = v75
            } else {
                var v76 int32 = -1
                fmt.Println(v6, v51, fi70)
                for fi77 := 0; fi77 < 3; fi77++ {
                    fmt.Println(fi77, v76, v52)
                    var v78 int = fi77
                    if ((v52 <= (v52 ^ (v52 & v52))) && true) {
                        fmt.Println(v5, fi70, v76)
                        if (fi70 > fi77) {
                            switch 9223372036854775807 {
                            case 4294967295:
                                fmt.Println(v6, fi70, v6)
                                v76 = v76
                                var v79 uint32 = (^(v52))
                                _ = v79
                            }
                            fmt.Println(v6)
                        }
                    } else {
                        var v80 uint8 = (v6 >> ((42) & 63))
                        v76 = 249
                        _ = v80
                    }
                    var v81 uint64 = v51
                    switch (("foo bar" + ("hello" + "\n\t")) + "a") {
                    case "foo bar":
                        var v82 string = ("foo bar" + "foo bar")
                        _ = v82
                    case "hello":
                        v52 += (v52 & v52)
                    }
                    _ = fi77
                    _ = v78
                    _ = v81
                }
                _ = v76
            }
            for fi83 := 0; fi83 < 3; fi83++ {
                if true {
                    if (uint(v51) > 279) {
                        if (!(0 > 32767)) {
                            v84 := GenS6{F0: 2, F1: (v52 & (v52 | (v52 << ((65535) & 63))))}
                            _ = v84
                        } else {
                            fmt.Println(fi83)
                        }
                    } else {
                        _ = len(v1)
                        v1[(v52 / (((v52 << ((1) & 63))) | 1))] = v52
                        var v85 int64 = -40
                        _ = v85
                    }
                } else {
                    v86 := GenS7{F0: 1, F1: 982}
                    v52 -= v52
                    var v87 float32 = 1.5
                    _ = v86
                    _ = v87
                }
                var v88 uint32 = v52
                var v89 uint8 = v6
                _ = len(v1)
                _ = fi83
                _ = v88
                _ = v89
            }
            _ = fi70
        }
    default:
        v6 += 128
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v6
}
