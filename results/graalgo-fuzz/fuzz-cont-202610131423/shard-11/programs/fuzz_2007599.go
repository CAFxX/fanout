package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 uint64; F2 int8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int64; F1 uint64 }
type GenS5 struct { F0 int64; F1 bool; F2 bool }
type GenS6 struct { F0 int64; F1 uint }
type GenS7 struct { F0 int8; F1 int64 }
type GenS8 struct { F0 uint16; F1 int8 }
type GenS9 struct { F0 uint8; F1 uint64; F2 int8 }
type GenS10 struct { F0 int; F1 int32; F2 uint8 }
type GenS11 struct { F0 uint; F1 string }
type GenS12 struct { F0 uint32; F1 bool }

func main() {
    var v1 uint16 = 65535
    v2 := make(map[string]bool)
    v3 := []string{""}
    var v4 uint16 = v1
    v5 := make(map[int]int)
    v6 := []string{"hello"}
    switch (("foo bar" + "") + "") {
    case "a":
        v7 := GenS1{F0: 458}
        _ = v7
    case "\n\t":
        v4 = (v1 ^ 427)
    case "foo bar":
        v4 = v1
    default:
        if (v1 > (v1 >> ((391) & 63))) {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        } else {
            var v8 int16 = -1
            var v9 string = ("" + "foo bar")
            fmt.Println(v8)
            _ = v8
            _ = v9
        }
    }
    if len(v6) > 0 {
        v6[0] = "\n\t"
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        for fi11 := 0; fi11 < 3; fi11++ {
            if false {
                fmt.Println(fi10, v1, v1)
                v4 -= v4
            }
            v12 := GenS1{F0: 2}
            switch ("foo bar" + "foo bar") {
            case "hello":
                if (!((uint(v4) != 18446744073709551615) || ("foo bar" < ""))) {
                    fmt.Println(v1, fi10, fi10)
                    v1 += 913
                } else {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v1 += (^(v1))
                }
                var v13 int64 = int64(v1)
                v1 = v4
                _ = v13
            case "a":
                v14 := GenS1{F0: 602}
                _ = v14
            }
            var v15 int8 = 0
            if (!false) {
                fmt.Println(fi10, fi10, v15)
            } else {
                fmt.Println(v15, v1, v1)
            }
            if (176 > 4294967295) {
                v16 := GenS1{F0: 234}
                _ = len(v3)
                switch (-(fi10)) {
                case 128:
                    fmt.Println(fi10, v4, v15)
                case 1:
                    var v17 uint16 = 65535
                    _ = v17
                default:
                    switch fi11 {
                    case 32767:
                        var v18 uint64 = 308
                        _ = v18
                    case 0:
                        if false {
                            v19 := v2[(("foo bar" + "") + "foo bar")]
                            _ = v19
                        } else {
                            if true {
                                if true {
                                    if true {
                                        var v20 string = "hello"
                                        if (46 == 165) {
                                            if true {
                                                var v21 float32 = (3.14159 - 1e10)
                                                v22 := GenS1{F0: 263}
                                                switch v20 {
                                                case "foo bar":
                                                    switch "foo bar" {
                                                    case "":
                                                        var v23 int16 = 1000
                                                        _ = v23
                                                    case "\n\t":
                                                        var v24 int16 = -561
                                                        _ = v24
                                                    case "foo bar":
                                                        v20 = (v20 + v20)
                                                    }
                                                case "a":
                                                    fmt.Println(v4, v4, fi10)
                                                    v15 -= -79
                                                    var v25 float32 = 0.1
                                                    _ = v25
                                                }
                                                _ = v21
                                                _ = v22
                                            }
                                        }
                                        if true {
                                            if ((!(v20 > v20)) || (18446744073709551615 != 18446744073709551615)) {
                                                if true {
                                                    if (false || false) {
                                                        v4 -= v4
                                                    } else {
                                                        if true {
                                                            v4 = (v4 >> ((4294967295) & 63))
                                                            var v26 int32 = 2147483647
                                                            v20 = ((v20 + ((v20 + ((v20 + "\n\t") + "a")) + (v20 + "hello"))) + "\n\t")
                                                            _ = v26
                                                        } else {
                                                            var v27 int32 = 2147483647
                                                            v3 = append(v3, v20)
                                                            _ = v27
                                                        }
                                                    }
                                                }
                                            }
                                            if true {
                                                v28 := GenS2{F0: 0, F1: 4294967295, F2: v15}
                                                switch fi11 {
                                                case -2147483648:
                                                    v20 = v20
                                                }
                                                var v29 float64 = ((0.0 * 0.1) / 0.1)
                                                _ = v28
                                                _ = v29
                                            } else {
                                                fmt.Println(v20, v1)
                                            }
                                            fmt.Println(v4)
                                        } else {
                                            if true {
                                                _ = len(v5)
                                                var v30 int8 = -128
                                                var v31 int8 = -128
                                                _ = v30
                                                _ = v31
                                            }
                                            v32 := GenS3{F0: 728}
                                            _ = v32
                                        }
                                        _ = v20
                                    }
                                } else {
                                    if false {
                                        var v33 int32 = -1
                                        v34 := GenS1{F0: 42}
                                        if (1 != 935) {
                                            if len(v6) > 0 {
                                                v6[0] = "hello"
                                            }
                                            var v35 uint16 = v1
                                            if true {
                                                if (v33 >= (v33 * v33)) {
                                                    v35 += (^(v1))
                                                    v36 := GenS1{F0: 18446744073709551615}
                                                    v4 = (v1 & (v4 | 187))
                                                    v1 += 309
                                                    v15 = v15
                                                    _ = v36
                                                }
                                            } else {
                                                var v37 int16 = -32768
                                                if (false && ((42 >= 571) || ((false || (1 > 4294967295)) || ((true && true) || (v1 > (v4 << ((547) & 63))))))) {
                                                    var v38 int64 = -1
                                                    switch "foo bar" {
                                                    case "a":
                                                        if false {
                                                            var v39 int16 = (v37 - (v37 & v37))
                                                            v37 = v39
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            _ = v39
                                                        } else {
                                                            var v40 int32 = (v33 * (v33 + -1))
                                                            if false {
                                                                if (128 >= 128) {
                                                                    v3 = append(v3, "")
                                                                    v41 := GenS4{F0: (-(v38)), F1: 685}
                                                                    _ = v41
                                                                } else {
                                                                    v42 := v5[fi10]
                                                                    _ = v42
                                                                }
                                                            }
                                                            var v43 int32 = 136
                                                            _ = v40
                                                            _ = v43
                                                        }
                                                        v44 := GenS5{F0: int64(fi10), F1: false, F2: false}
                                                        _ = v44
                                                    }
                                                    switch "hello" {
                                                    case "foo bar":
                                                        if true {
                                                            var v45 int8 = v15
                                                            if true {
                                                                if true {
                                                                    v46 := GenS6{F0: (v38 << ((65535) & 63)), F1: 2}
                                                                    v6 = append(v6, "hello")
                                                                    _ = v46
                                                                }
                                                                v47 := GenS7{F0: v45, F1: (v38 << ((42) & 63))}
                                                                _ = v47
                                                            }
                                                            fmt.Println(v38, v1)
                                                            var v48 int64 = 1
                                                            _ = v45
                                                            _ = v48
                                                        } else {
                                                            switch fi11 {
                                                            case 2147483647:
                                                                var v49 int8 = (v15 % (((v15 | 42)) | 1))
                                                                _ = v49
                                                            case 896:
                                                                var v50 int64 = 0
                                                                _ = v50
                                                            }
                                                            v51 := GenS8{F0: (v1 >> ((1) & 63)), F1: (v15 - v15)}
                                                            var v52 int16 = (v37 & (v37 - v37))
                                                            _ = v51
                                                            _ = v52
                                                        }
                                                    case "\n\t":
                                                        v53 := GenS8{F0: (v1 & (v1 + (v1 >> ((87) & 63)))), F1: -1}
                                                        _ = v53
                                                    default:
                                                        v3 = append(v3, ("a" + "hello"))
                                                    }
                                                    _ = v38
                                                } else {
                                                    if (false && (true && ((!true) || false))) {
                                                        v54 := GenS9{F0: 42, F1: 14, F2: v15}
                                                        fmt.Println(fi11)
                                                        _ = v54
                                                    }
                                                    v35 += (v35 & (v1 - 103))
                                                    var v55 int64 = 0
                                                    _ = v55
                                                }
                                                if (!true) {
                                                    fmt.Println(v1, v4)
                                                    var v56 int8 = v15
                                                    _ = v56
                                                } else {
                                                    fmt.Println(v37, v35, v4)
                                                }
                                                fmt.Println(v1, fi10)
                                                var v57 float32 = 0.0
                                                _ = v37
                                                _ = v57
                                            }
                                            v58 := GenS10{F0: fi10, F1: (v33 * (v33 + (v33 << ((387) & 63)))), F2: 2}
                                            var v59 int = 866
                                            if false {
                                                switch "foo bar" {
                                                case "\n\t":
                                                    var v60 int = fi10
                                                    _ = v60
                                                }
                                                v61 := GenS11{F0: 65535, F1: ("a" + ("a" + "foo bar"))}
                                                _ = v61
                                            } else {
                                                v62 := GenS8{F0: v4, F1: (v15 >> ((361) & 63))}
                                                _ = v62
                                            }
                                            var v63 int64 = 98
                                            _ = v35
                                            _ = v58
                                            _ = v59
                                            _ = v63
                                        }
                                        _ = v33
                                        _ = v34
                                    } else {
                                        switch fi11 {
                                        case 128:
                                            var v64 int = (fi10 / ((fi10) | 1))
                                            _ = v64
                                        default:
                                            v65 := GenS5{F0: -1, F1: true, F2: true}
                                            _ = v65
                                        }
                                        var v66 float32 = -2.25
                                        var v67 uint8 = 61
                                        _ = v66
                                        _ = v67
                                    }
                                    _ = len(v5)
                                    fmt.Println(v1, v4)
                                }
                            }
                        }
                    case 293:
                        var v68 uint64 = 1
                        _ = v68
                    default:
                        v69 := GenS12{F0: 2, F1: false}
                        _ = v69
                    }
                }
                _ = v16
            } else {
                if len(v6) > 0 {
                    v6[0] = ""
                }
                v2["\n\t"] = true
                v1 -= v1
            }
            var v70 int64 = -9223372036854775808
            fmt.Println(v1)
            v4 = 16
            delete(v2, ("" + "hello"))
            _ = fi11
            _ = v12
            _ = v15
            _ = v70
        }
        v1 -= v1
        v1 = v1
        _ = fi10
    }
    if false {
        for fi71 := 0; fi71 < 14; fi71++ {
            switch "a" {
            case "\n\t":
                if (1 < -32768) {
                    v4 = 1
                    v4 += 255
                }
            case "foo bar":
                for fi72 := 0; fi72 < 19; fi72++ {
                    v73 := GenS2{F0: 1, F1: 821, F2: 35}
                    if ((1 <= 32767) && false) {
                        var v74 uint32 = 2147483647
                        if ((^(v74)) == v74) {
                            v1 += v4
                        } else {
                            v4 = 255
                            var v75 uint64 = 877
                            _ = v75
                        }
                        var v76 int16 = 42
                        _ = v74
                        _ = v76
                    }
                    _ = fi72
                    _ = v73
                }
            case "a":
                if false {
                    v1 = 1
                    v4 = (v1 | v4)
                }
            default:
                var v77 int64 = -12
                _ = v77
            }
            _ = fi71
        }
    } else {
        v1 += v4
    }
    v78 := v5[-710]
    delete(v2, "hello")
    switch "" {
    case "foo bar":
        v79 := GenS6{F0: 0, F1: 2}
        _ = v79
    default:
        v2[(("" + "foo bar") + "a")] = true
    }
    var v80 uint = 1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v78:", v78)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v78
    _ = v80
}
