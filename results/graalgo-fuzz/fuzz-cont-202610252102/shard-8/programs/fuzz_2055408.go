package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 uint32 }
type GenS2 struct { F0 uint64; F1 int32; F2 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int8; F1 int16 }
type GenS7 struct { F0 int; F1 uint; F2 uint16 }
type GenS8 struct { F0 string; F1 bool; F2 int16 }
type GenS9 struct { F0 int64; F1 int64; F2 uint8 }
func genFn0() int16 {
    var v7 float64 = 1.5
    v8 := GenS1{F0: 2, F1: 42, F2: 2}
    switch 42 {
    case -334:
        var v9 int32 = 594
        _ = v9
    default:
        switch -81 {
        case 4294967295:
            if (18446744073709551615 >= 255) {
                var v10 string = "hello"
                for fi11 := 0; fi11 < 10; fi11++ {
                    if false {
                        v12 := []int{fi11, (fi11 * fi11), 885, (fi11 << ((0) & 63))}
                        for fi13 := 0; fi13 < 16; fi13++ {
                            var v14 float64 = v7
                            _ = fi13
                            _ = v14
                        }
                        v15 := GenS1{F0: 55, F1: -1, F2: 1}
                        _ = v12
                        _ = v15
                    }
                    _ = fi11
                }
                v16 := GenS1{F0: 671, F1: -32768, F2: 4294967295}
                _ = v10
                _ = v16
            }
        case 42:
            v17 := GenS2{F0: 654, F1: -2147483648, F2: -215}
            _ = v17
        case -969:
            switch 4294967295 {
            case 32767:
                var v18 uint64 = 1
                _ = v18
            case -9223372036854775808:
                var v19 uint64 = 364
                _ = v19
            }
        }
    }
    v20 := GenS2{F0: 4294967295, F1: -1, F2: 42}
    switch "\n\t" {
    case "":
        if (127 <= 4294967295) {
            for fi21 := 0; fi21 < 2; fi21++ {
                if false {
                    v22 := []int{}
                    var v23 bool = (32767 != -148)
                    _ = v22
                    _ = v23
                } else {
                    v7 = (-2.25 * v7)
                    v7 += v7
                    v24 := GenS3{F0: 729}
                    _ = v24
                }
                v7 = float64((0.1 + float32(((0.0 / v7) * (1e10 + v7)))))
                var v25 uint16 = 32767
                _ = fi21
                _ = v25
            }
        }
    default:
        v26 := GenS3{F0: 32767}
        _ = v26
    }
    _ = v20
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := []uint32{2, 65535, 4294967295, 659}
    var v2 int8 = 127
    var v3 bool = (-711 != 0)
    var v4 int64 = 1
    var v5 int16 = 605
    v6 := make(map[uint32]string)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi27 := 0; fi27 < 6; fi27++ {
        if (((v2 * (v2 - (v2 >> ((2) & 63)))) < v2) && ((v3 && ("hello" == "foo bar")) || true)) {
            for fi28 := 0; fi28 < 15; fi28++ {
                var v29 uint8 = 136
                _ = fi28
                _ = v29
            }
            for fi30 := 0; fi30 < 9; fi30++ {
                if true {
                    switch (fi30 + (-(fi30))) {
                    case -9223372036854775808:
                        v31 := GenS4{F0: 849}
                        _ = v31
                    }
                } else {
                    if (0.0 <= float64(-2.25)) {
                        v4 += -855
                        _ = len(v1)
                    } else {
                        if (42 <= 970) {
                            if v3 {
                                var v32 int = fi30
                                fmt.Println(fi27, v32)
                                fmt.Println(v5, v4)
                                var v33 int16 = (-(v5))
                                _ = v32
                                _ = v33
                            } else {
                                var v34 int8 = (^(v2))
                                if false {
                                    if ((false && (!v3)) && (1e10 <= 3.14159)) {
                                        v6[1] = ""
                                        fmt.Println(fi27, fi30, v5)
                                        fmt.Println(v2, fi30)
                                        fmt.Println(v4)
                                        v3 = (-2147483648 > 1)
                                        var v35 uint32 = 2147483647
                                        fmt.Println(v3, v5, v34)
                                        _ = v35
                                    } else {
                                        if v3 {
                                            if true {
                                                fmt.Println(v34, v4)
                                            } else {
                                                if false {
                                                    if (0 >= 65535) {
                                                        var v36 string = "a"
                                                        _ = v36
                                                    }
                                                    fmt.Println(fi27)
                                                } else {
                                                    if v3 {
                                                        if (v3 || ((^(v34)) < 42)) {
                                                            var v37 string = ("hello" + "hello")
                                                            _ = v37
                                                        }
                                                        var v38 uint64 = 18446744073709551615
                                                        _ = v38
                                                    } else {
                                                        var v39 int16 = (v5 & v5)
                                                        _ = v39
                                                    }
                                                    v4 = v4
                                                }
                                            }
                                            fmt.Println(v2, fi27, v2)
                                        }
                                        if ((((("foo bar" + "foo bar") + "hello") + "") == ("a" + ("foo bar" + ""))) && ((1 == 128) || v3)) {
                                            v5 += (-(v5))
                                            var v40 uint = 0
                                            _ = v40
                                        }
                                        var v41 uint32 = 2147483648
                                        _ = v41
                                    }
                                    if v3 {
                                        switch fi30 {
                                        case 32767:
                                            if true {
                                                var v42 int16 = (v5 * genFn0())
                                                if ((2 < 4294967295) || (true || v3)) {
                                                    if (0 > fi27) {
                                                        var v43 uint32 = 269
                                                        var v44 uint8 = 212
                                                        _ = v43
                                                        _ = v44
                                                    }
                                                } else {
                                                    var v45 int8 = (v2 % ((v2) | 1))
                                                    _ = v45
                                                }
                                                _ = v42
                                            }
                                        default:
                                            v46 := GenS2{F0: 2, F1: 42, F2: fi30}
                                            _ = v46
                                        }
                                        switch (fi30 / ((fi27) | 1)) {
                                        case 892:
                                            fmt.Println(v2, v3)
                                        }
                                        v47 := GenS5{F0: (v5 * -1)}
                                        v48 := GenS6{F0: 0, F1: (v5 << ((4294967295) & 63))}
                                        if v3 {
                                            delete(v6, 155)
                                        } else {
                                            fmt.Println(fi27, v2, fi30)
                                        }
                                        v5 = (v5 / (((v5 << ((1) & 63))) | 1))
                                        _ = v47
                                        _ = v48
                                    } else {
                                        fmt.Println(fi27, v3)
                                        fmt.Println(v3)
                                        v4 = (v4 & -1)
                                        v3 = (!(int32(v4) == 42))
                                        fmt.Println(fi27, v5)
                                    }
                                    var v49 uint = 255
                                    if (0 >= 529) {
                                        if ((v2 != (v34 * v34)) || (159 == 128)) {
                                            v5 -= v5
                                            fmt.Println(v5, v34)
                                        }
                                    }
                                    _ = v49
                                } else {
                                    var v50 int32 = -1
                                    fmt.Println(v5)
                                    if false {
                                        if len(v1) > 0 {
                                            v1[0] = 590
                                        }
                                    }
                                    _ = v50
                                }
                                _ = v34
                            }
                        }
                        switch (fi30 ^ fi27) {
                        case -1:
                            fmt.Println(v4, v5, fi27)
                        case -250:
                            v51 := GenS7{F0: fi30, F1: 400, F2: 43}
                            _ = v51
                        default:
                            if (!((v2 + 0) < (v2 + -128))) {
                                v4 += 1
                            }
                        }
                    }
                    var v52 int8 = v2
                    _ = v52
                }
                if v3 {
                    if false {
                        fmt.Println(fi27, v5)
                    } else {
                        var v53 uint32 = 42
                        _ = v53
                    }
                    var v54 int = fi27
                    _ = v54
                }
                v5 -= int16(fi27)
                _ = fi30
            }
            fmt.Println(v4, v4, v5)
        } else {
            var v55 uint16 = 42
            if len(v1) > 0 {
                v1[0] = 1
            }
            _ = v55
        }
        _ = fi27
    }
    if v3 {
        var v56 string = "\n\t"
        _ = v56
    }
    switch "hello" {
    case "a":
        fmt.Println(v2)
    case "hello":
        var v57 bool = ("\n\t" < ("\n\t" + (("foo bar" + ("" + "\n\t")) + ("" + "a"))))
        _ = v57
    case "foo bar":
        switch 128 {
        case -732:
            if (0 > 2) {
                for fi58 := 0; fi58 < 2; fi58++ {
                    v4 = v4
                    _ = fi58
                }
            } else {
                v2 += (v2 % ((42) | 1))
                _ = len(v1)
            }
        case 42:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case -2147483648:
            if (v3 || v3) {
                for fi59 := 0; fi59 < 16; fi59++ {
                    switch (fi59 | 128) {
                    case -1:
                        if len(v1) > 0 {
                            v1[0] = 65535
                        }
                    }
                    switch (("foo bar" + (("foo bar" + "hello") + "hello")) + "") {
                    case "":
                        v5 += v5
                    case "\n\t":
                        if false {
                            v2 += (v2 ^ v2)
                            v6[65535] = ("a" + "foo bar")
                        }
                    }
                    v60 := GenS2{F0: 402, F1: -929, F2: (fi59 + (-(fi59)))}
                    _ = fi59
                    _ = v60
                }
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v61 bool = (v4 != -14)
                _ = v61
            }
        default:
            v1 = append(v1, 629)
        }
    default:
        v62 := GenS8{F0: "a", F1: (!(!true)), F2: v5}
        _ = v62
    }
    v63 := GenS9{F0: (v4 >> ((1) & 63)), F1: v4, F2: 128}
    var v64 bool = false
    var v65 string = (("foo bar" + "foo bar") + ("hello" + "\n\t"))
    if false {
        v66 := GenS4{F0: 2}
        _ = v66
    }
    for fi67 := 0; fi67 < 18; fi67++ {
        v68 := GenS5{F0: (v5 << ((2) & 63))}
        switch v65 {
        case "a":
            if v64 {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                for fi69 := 0; fi69 < 10; fi69++ {
                    v70 := GenS8{F0: v65, F1: (((-128 <= (v2 & (v2 | v2))) || true) || v3), F2: genFn0()}
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi69
                    _ = v70
                }
            } else {
                v65 = (v65 + "hello")
            }
        }
        _ = fi67
        _ = v68
    }
    switch (v65 + v65) {
    case "a":
        v71 := GenS1{F0: 232, F1: (v5 - v5), F2: 843}
        _ = v71
    case "":
        v3 = (true && v64)
    default:
        switch int(v4) {
        case 924:
            var v72 uint8 = 0
            _ = v72
        }
    }
    switch -8 {
    case 523:
        v73 := GenS3{F0: 32767}
        _ = v73
    case 9223372036854775807:
        _ = len(v6)
    }
    var v74 uint16 = 32767
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v63
    _ = v64
    _ = v65
    _ = v74
}
