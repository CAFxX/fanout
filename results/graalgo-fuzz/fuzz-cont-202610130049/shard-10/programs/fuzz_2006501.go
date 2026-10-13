package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64; F1 uint64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint; F1 string }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint32; F1 uint32; F2 uint16 }
type GenS10 struct { F0 int64 }
func genFn0(p0 int, p1 uint16) int {
    var v5 bool = false
    v6 := GenS1{F0: (p1 - (p1 + p1)), F1: 127}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (-(p0))
}
func genFn1(p0 uint, p1 uint, p2 uint16) int16 {
    v7 := []int{int(p0)}
    switch -837 {
    case 0:
        p2 -= (p2 % (((p2 & (p2 / ((42) | 1)))) | 1))
    case -9223372036854775808:
        v8 := GenS1{F0: p2, F1: 128}
        _ = v8
    default:
        var v9 int16 = -879
        _ = v9
    }
    v10 := GenS2{F0: -1}
    if (127 < 42) {
        if len(v7) > 0 {
            v7[0] = 127
        }
    } else {
        for fi11 := 0; fi11 < 1; fi11++ {
            var v12 uint = 1
            _ = fi11
            _ = v12
        }
    }
    if false {
        if (false || (true && false)) {
            fmt.Println(p0, p0, p0)
        }
        p2 = (^(p2))
        var v13 int32 = 2147483647
        _ = v13
    } else {
        _ = len(v7)
        var v14 float64 = float64((3.14159 * 1.5))
        var v15 bool = (!false)
        _ = v14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    return -1
}

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{}
    var v3 uint32 = 4294967295
    v4 := []int{127}
    if false {
        switch 127 {
        case 32768:
            if (v3 != v3) {
                switch "a" {
                case "hello":
                    for fi16 := 0; fi16 < 19; fi16++ {
                        v2 = append(v2, 0)
                        for fi17 := 0; fi17 < 10; fi17++ {
                            fmt.Println(v3, fi16)
                            v3 = 1
                            _ = fi17
                        }
                        _ = fi16
                    }
                case "\n\t":
                    for fi18 := 0; fi18 < 9; fi18++ {
                        var v19 string = "hello"
                        v20 := GenS2{F0: 9223372036854775807}
                        switch v19 {
                        case "hello":
                            fmt.Println(fi18)
                        case "\n\t":
                            if false {
                                fmt.Println(fi18)
                                if ((v19 + v19) >= (v19 + "")) {
                                    var v21 string = v19
                                    var v22 int64 = -5
                                    for fi23 := 0; fi23 < 2; fi23++ {
                                        var v24 uint32 = 1
                                        v24 += v3
                                        _ = fi23
                                        _ = v24
                                    }
                                    _ = v21
                                    _ = v22
                                } else {
                                    v25 := GenS1{F0: 32767, F1: 136}
                                    _ = v25
                                }
                            } else {
                                if (true && (true && true)) {
                                    v19 = v19
                                    for fi26 := 0; fi26 < 10; fi26++ {
                                        var v27 int32 = -1
                                        v28 := GenS3{F0: 40}
                                        var v29 float64 = -2.25
                                        var v30 int = fi18
                                        var v31 uint32 = 2147483647
                                        v31 = (v31 + (v31 | (^(v3))))
                                        _ = fi26
                                        _ = v27
                                        _ = v28
                                        _ = v29
                                        _ = v30
                                        _ = v31
                                    }
                                    delete(v1, (v19 + v19))
                                }
                                var v32 int32 = -2147483648
                                v33 := GenS1{F0: 32767, F1: 1}
                                _ = v32
                                _ = v33
                            }
                        }
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                case "foo bar":
                    v34 := GenS4{F0: 0, F1: 0}
                    _ = v34
                default:
                    v3 = 867
                }
                v35 := GenS5{F0: 18446744073709551615}
                _ = v35
            }
        default:
            switch ("hello" + (("\n\t" + "") + "\n\t")) {
            case "":
                for fi36 := 0; fi36 < 17; fi36++ {
                    switch (("" + ((("a" + "foo bar") + "foo bar") + "hello")) + "") {
                    case "foo bar":
                        if len(v2) > 0 {
                            v2[0] = (v3 / ((v3) | 1))
                        }
                    case "hello":
                        var v37 int = (-(fi36))
                        _ = v37
                    }
                    _ = len(v1)
                    _ = fi36
                }
            case "\n\t":
                v38 := GenS6{F0: 1}
                _ = v38
            case "foo bar":
                for fi39 := 0; fi39 < 19; fi39++ {
                    v40 := GenS4{F0: -9223372036854775808, F1: 0}
                    if false {
                        if true {
                            var v41 int64 = 9223372036854775807
                            _ = v41
                        }
                    }
                    v42 := GenS2{F0: 650}
                    _ = fi39
                    _ = v40
                    _ = v42
                }
            default:
                var v43 uint16 = 42
                _ = v43
            }
        }
        var v44 uint32 = 2147483648
        fmt.Println(v3)
        _ = v44
    }
    var v45 uint8 = 100
    switch 686 {
    case 995:
        for fi46 := 0; fi46 < 8; fi46++ {
            v47 := GenS7{F0: 318, F1: "\n\t"}
            _ = fi46
            _ = v47
        }
    case 2147483647:
        v45 = v45
    }
    var v48 uint16 = 727
    fmt.Println(v3)
    v48 += 0
    var v49 uint64 = 1
    var v50 int32 = 2147483647
    if ("hello" == ("hello" + ("foo bar" + (("hello" + "\n\t") + (("\n\t" + "hello") + ""))))) {
        if (0 == 390) {
            if (!true) {
                v51 := GenS3{F0: -1}
                v2 = append(v2, 2)
                var v52 float64 = 3.14159
                _ = v51
                _ = v52
            }
            v3 = v3
        }
        var v53 float64 = 0.1
        v48 -= v48
        _ = v53
    }
    for fi54 := 0; fi54 < 13; fi54++ {
        if true {
            var v55 int32 = v50
            if len(v4) > 0 {
                v4[0] = (fi54 * genFn0(genFn0((fi54 - int(v3)), (v48 >> ((1) & 63))), 32767))
            }
            _ = v55
        } else {
            var v56 int8 = 27
            _ = v56
        }
        _ = fi54
    }
    for fi57 := 0; fi57 < 6; fi57++ {
        for fi58 := 0; fi58 < 2; fi58++ {
            fmt.Println(v50, v50, v45)
            _ = fi58
        }
        v50 -= (v50 | (^(v50)))
        _ = fi57
    }
    for fi59 := 0; fi59 < 12; fi59++ {
        if true {
            v60 := GenS8{F0: (v48 * 857)}
            if (18446744073709551615 < (v49 >> ((2) & 63))) {
                v61 := v1[((((("" + "\n\t") + "a") + "hello") + "") + "")]
                var v62 uint = 255
                _ = v61
                _ = v62
            } else {
                fmt.Println(v50)
                if (!false) {
                    var v63 uint = 1
                    var v64 bool = true
                    switch "a" {
                    case "foo bar":
                        v65 := GenS2{F0: -1}
                        _ = v65
                    case "":
                        var v66 float32 = 0.0
                        _ = v66
                    default:
                        var v67 float64 = 0.0
                        _ = v67
                    }
                    _ = v63
                    _ = v64
                } else {
                    if len(v2) > 0 {
                        v2[0] = v3
                    }
                    if (1.5 >= float64(1.5)) {
                        v4 = append(v4, fi59)
                        fmt.Println(fi59, v48, v3)
                    } else {
                        v50 = v50
                    }
                    switch "a" {
                    case "foo bar":
                        var v68 int16 = -851
                        _ = v68
                    case "":
                        for fi69 := 0; fi69 < 17; fi69++ {
                            if false {
                                v70 := GenS6{F0: -1}
                                var v71 int8 = -1
                                fmt.Println(v71, fi69, v49)
                                _ = v70
                                _ = v71
                            }
                            fmt.Println(v45, v48)
                            _ = fi69
                        }
                    }
                }
            }
            if false {
                v72 := GenS9{F0: (^(v3)), F1: 335, F2: v48}
                _ = v72
            } else {
                if false {
                    v1[("foo bar" + "a")] = (fi59 >> ((uint(v45)) & 63))
                } else {
                    var v73 uint16 = v48
                    v3 = 2147483648
                    switch "foo bar" {
                    case "":
                        v50 -= int32(v45)
                    case "foo bar":
                        v74 := GenS10{F0: -138}
                        _ = v74
                    default:
                        var v75 string = ""
                        _ = v75
                    }
                    _ = v73
                }
                var v76 uint8 = (v45 / ((uint8(v3)) | 1))
                _ = v76
            }
            _ = v60
        }
        fmt.Println(v50, fi59)
        var v77 int32 = int32(v45)
        _ = fi59
        _ = v77
    }
    v49 = 888
    v45 = 2
    v78 := GenS6{F0: 100}
    v3 += (v3 / ((v3) | 1))
    fmt.Println(v3, v3, v45)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v48
    _ = v49
    _ = v50
    _ = v78
}
