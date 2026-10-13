package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint16; F1 string }
type GenS5 struct { F0 int64; F1 uint8 }
func genFn0(p0 uint, p1 uint8, p2 bool) uint {
    var v6 uint16 = 1
    fmt.Println(v6)
    fmt.Println(p0)
    for fi7 := 0; fi7 < 6; fi7++ {
        for fi8 := 0; fi8 < 4; fi8++ {
            var v9 uint16 = uint16(p0)
            v10 := []int{(fi8 >> ((p0) & 63)), fi8, (-(fi7)), fi7}
            _ = fi8
            _ = v10
            _ = v9
        }
        var v11 uint32 = 1
        var v12 uint64 = 0
        _ = fi7
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p0
}
func genFn1(p0 uint32, p1 uint, p2 uint64) int8 {
    fmt.Println(p0, p1)
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    return 68
}

func main() {
    var v1 float32 = 1e10
    v2 := []int{}
    var v3 uint8 = 127
    var v4 uint8 = v3
    var v5 string = "\n\t"
    v4 += (^(v4))
    v4 -= 215
    v2 = append(v2, 4294967295)
    switch 127 {
    case -967:
        fmt.Println(v3, v5)
    case -846:
        for fi13 := 0; fi13 < 13; fi13++ {
            if false {
                v3 = 127
                var v14 uint = 4294967295
                v3 = 42
                _ = v14
            }
            switch v5 {
            case "\n\t":
                switch v5 {
                case "hello":
                    if false {
                        v2 = append(v2, (fi13 >> ((160) & 63)))
                    }
                case "":
                    _ = len(v2)
                }
            }
            _ = len(v2)
            _ = fi13
        }
    default:
        v5 = ("a" + v5)
    }
    var v15 float32 = v1
    switch "a" {
    case "foo bar":
        switch 2147483647 {
        case 883:
            var v16 bool = false
            _ = v16
        case 128:
            var v17 uint32 = 2147483648
            _ = v17
        case 1:
            v18 := GenS1{F0: 0, F1: (v4 * 85)}
            _ = v18
        }
    case "a":
        var v19 uint = 18446744073709551615
        _ = v19
    }
    fmt.Println(v5, v15)
    if true {
        for fi20 := 0; fi20 < 18; fi20++ {
            if (false && false) {
                switch fi20 {
                case -2147483648:
                    if (!(4294967295 <= 42)) {
                        var v21 uint64 = 4294967295
                        if (uint8(v21) != v3) {
                            switch (fi20 * fi20) {
                            case 32768:
                                var v22 int32 = 2147483647
                                _ = v22
                            case -2147483648:
                                var v23 uint16 = 353
                                _ = v23
                            }
                            for fi24 := 0; fi24 < 4; fi24++ {
                                var v25 uint8 = v4
                                _ = fi24
                                _ = v25
                            }
                        }
                        _ = v21
                    } else {
                        v15 = (3.14159 + v1)
                        v1 = (1e10 / 1.5)
                        if (65535 >= 554) {
                            var v26 uint32 = 703
                            v27 := GenS2{F0: 100, F1: false}
                            v15 -= (1e10 / v15)
                            _ = v26
                            _ = v27
                        }
                    }
                case 127:
                    for fi28 := 0; fi28 < 3; fi28++ {
                        var v29 int64 = 42
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v30 int16 = -1
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                case -627:
                    v5 = ((v5 + ("" + (v5 + ""))) + "")
                default:
                    _ = len(v2)
                }
            } else {
                if (127 <= -128) {
                    if true {
                        v15 -= (-2.25 - v1)
                        switch fi20 {
                        case -2147483648:
                            fmt.Println(v1, v4)
                        case -9223372036854775808:
                            switch (("" + "") + (v5 + "hello")) {
                            case "a":
                                for fi31 := 0; fi31 < 8; fi31++ {
                                    v4 = 127
                                    _ = fi31
                                }
                            case "\n\t":
                                if false {
                                    v5 = "a"
                                    var v32 uint64 = 42
                                    _ = v32
                                } else {
                                    switch ("\n\t" + v5) {
                                    case "\n\t":
                                        v33 := GenS3{F0: fi20}
                                        _ = v33
                                    case "a":
                                        if (fi20 != (fi20 - (fi20 / (((^(fi20))) | 1)))) {
                                            var v34 string = ("foo bar" + ("" + v5))
                                            v3 = v3
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            _ = v34
                                        }
                                    }
                                }
                            case "foo bar":
                                fmt.Println(v15)
                            }
                        default:
                            v35 := GenS4{F0: 287, F1: ((("a" + ((v5 + "\n\t") + (("" + v5) + (((("\n\t" + (v5 + ("foo bar" + "\n\t"))) + ("foo bar" + ("\n\t" + v5))) + "hello") + v5)))) + v5) + ("foo bar" + (v5 + (v5 + (("" + "") + "\n\t")))))}
                            _ = v35
                        }
                        for fi36 := 0; fi36 < 8; fi36++ {
                            if false {
                                v37 := GenS3{F0: (fi36 * (-(fi20)))}
                                v38 := GenS3{F0: -1}
                                v15 = v15
                                _ = v37
                                _ = v38
                            } else {
                                v3 += (v4 >> ((42) & 63))
                                var v39 uint64 = 18446744073709551615
                                v40 := GenS2{F0: genFn1(65535, 0, v39), F1: false}
                                _ = v39
                                _ = v40
                            }
                            if true {
                                if ((false || (0 <= 1)) && ((603 >= 992) || (((float32((3.14159 * 0.1)) * v15) * v1) < ((((v15 - v15) + (v15 + v1)) + 1e10) / v15)))) {
                                    fmt.Println(v1, v15)
                                    var v41 int8 = 1
                                    if true {
                                        if len(v2) > 0 {
                                            v2[0] = fi20
                                        }
                                    } else {
                                        var v42 uint8 = 127
                                        _ = v42
                                    }
                                    _ = v41
                                } else {
                                    var v43 uint8 = (v3 >> ((42) & 63))
                                    var v44 bool = false
                                    _ = v43
                                    _ = v44
                                }
                                switch (v5 + v5) {
                                case "\n\t":
                                    switch v5 {
                                    case "\n\t":
                                        v5 = (v5 + v5)
                                    case "":
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    case "a":
                                        v45 := GenS3{F0: fi36}
                                        _ = v45
                                    }
                                    fmt.Println(fi20, v4, fi20)
                                    var v46 uint16 = 874
                                    _ = v46
                                case "foo bar":
                                    var v47 int16 = 1000
                                    _ = v47
                                }
                                fmt.Println(fi20, fi20, fi20)
                                if false {
                                    v48 := GenS1{F0: 1000, F1: 2}
                                    v4 += 128
                                    var v49 uint16 = 200
                                    _ = v48
                                    _ = v49
                                } else {
                                    switch "a" {
                                    case "":
                                        if (fi36 <= -161) {
                                            v50 := GenS1{F0: 0, F1: v3}
                                            _ = v50
                                        } else {
                                            if ((fi20 % (((fi36 & 127)) | 1)) != fi36) {
                                                _ = len(v2)
                                                v51 := GenS4{F0: 255, F1: ((((v5 + (v5 + v5)) + v5) + (v5 + ((v5 + v5) + v5))) + v5)}
                                                fmt.Println(fi36, fi20)
                                                _ = v51
                                            }
                                        }
                                    }
                                    v4 = 128
                                }
                                v52 := GenS5{F0: 42, F1: v4}
                                _ = v52
                            }
                            switch "" {
                            case "\n\t":
                                if true {
                                    var v53 string = v5
                                    _ = v53
                                } else {
                                    v54 := GenS1{F0: -1, F1: v3}
                                    var v55 uint32 = 2147483648
                                    _ = v54
                                    _ = v55
                                }
                            default:
                                v1 -= -2.25
                            }
                            if (!false) {
                                var v56 bool = ((true && false) && true)
                                var v57 int = fi20
                                _ = v56
                                _ = v57
                            }
                            if (42 < 42) {
                                switch (^(fi36)) {
                                case -391:
                                    switch v5 {
                                    case "a":
                                        v15 = (float32(1e10) / ((-2.25 - (v1 + 1e10)) - v1))
                                    case "foo bar":
                                        switch (v5 + "hello") {
                                        case "foo bar":
                                            v3 += 1
                                        case "":
                                            v58 := GenS4{F0: 613, F1: "\n\t"}
                                            _ = v58
                                        }
                                        var v59 int64 = -756
                                        v4 = v4
                                        v1 = (v15 * (0.1 * v1))
                                        var v60 int64 = v59
                                        _ = v59
                                        _ = v60
                                    }
                                case 128:
                                    v5 = "foo bar"
                                default:
                                    fmt.Println(v15)
                                }
                            }
                            _ = fi36
                        }
                    } else {
                        v3 = (v4 + (v4 % ((128) | 1)))
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    v2 = append(v2, 2147483647)
                }
                switch (("a" + (v5 + "foo bar")) + "\n\t") {
                case "":
                    if (!false) {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                case "hello":
                    v3 = (v4 - v4)
                case "\n\t":
                    _ = len(v2)
                default:
                    var v61 uint64 = 1
                    _ = v61
                }
                fmt.Println(v3)
            }
            v5 = v5
            fmt.Println(v15, v3, v3)
            _ = fi20
        }
    } else {
        var v62 string = v5
        _ = v62
    }
    v63 := GenS5{F0: -1, F1: v3}
    v64 := GenS4{F0: 2, F1: "hello"}
    var v65 string = v5
    if true {
        switch v5 {
        case "foo bar":
            fmt.Println(v3, v15)
        case "hello":
            for fi66 := 0; fi66 < 8; fi66++ {
                var v67 int32 = 873
                _ = fi66
                _ = v67
            }
        }
    }
    var v68 uint8 = (v4 ^ v3)
    switch (v5 + (((v65 + v65) + "foo bar") + "\n\t")) {
    case "foo bar":
        for fi69 := 0; fi69 < 8; fi69++ {
            var v70 int8 = -1
            _ = fi69
            _ = v70
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v65:", v65)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v63
    _ = v64
    _ = v65
    _ = v68
}
