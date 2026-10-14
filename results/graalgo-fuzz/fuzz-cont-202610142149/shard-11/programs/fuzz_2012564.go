package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 uint16; F1 int64; F2 int8 }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 uint8; F1 int64 }
type GenS5 struct { F0 uint32; F1 bool }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 bool; F1 uint32; F2 int32 }
func genFn0(p0 int32) uint16 {
    v4 := []int{9223372036854775807, 380}
    var v5 float64 = 3.14159
    _ = p0
    _ = v4
    _ = v5
    return 2
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 255
    v3 := []int{1, 0, -2147483648}
    _ = len(v3)
    _ = len(v1)
    v1[4294967295] = -9223372036854775808
    v6 := GenS1{F0: v2, F1: 32767}
    v7 := GenS1{F0: v2, F1: 1}
    if false {
        var v8 uint32 = 4294967295
        if (0 > 0) {
            v9 := v1[v8]
            _ = v9
        } else {
            for fi10 := 0; fi10 < 18; fi10++ {
                _ = len(v1)
                if true {
                    for fi11 := 0; fi11 < 1; fi11++ {
                        if (!(true || (false && (!false)))) {
                            var v12 float64 = -2.25
                            v12 += 0.0
                            var v13 int = 127
                            _ = v12
                            _ = v13
                        }
                        if false {
                            v2 += (^(v2))
                            switch (("foo bar" + "hello") + ("" + "foo bar")) {
                            case "foo bar":
                                switch (fi10 - (fi11 % (((fi11 << ((0) & 63))) | 1))) {
                                case 42:
                                    var v14 bool = false
                                    _ = v14
                                case 2147483647:
                                    v8 = v8
                                default:
                                    delete(v1, v8)
                                }
                            default:
                                if (42 > int64(fi10)) {
                                    v8 = v8
                                    var v15 int8 = -1
                                    v1[2] = (fi11 | fi11)
                                    _ = v15
                                }
                            }
                            var v16 int8 = -1
                            _ = v16
                        } else {
                            fmt.Println(fi11)
                            v8 = (v8 & (v8 | v8))
                            v17 := GenS2{F0: v2, F1: 1, F2: -1}
                            _ = v17
                        }
                        if (v8 == 2) {
                            v8 = v8
                            fmt.Println(fi11, fi10)
                        }
                        _ = fi11
                    }
                    for fi18 := 0; fi18 < 8; fi18++ {
                        switch "" {
                        case "a":
                            var v19 uint32 = (v8 & v8)
                            _ = v19
                        case "foo bar":
                            if false {
                                if (-221 <= -121) {
                                    var v20 uint8 = 255
                                    _ = v20
                                }
                                var v21 int64 = -64
                                v2 += 32767
                                _ = v21
                            } else {
                                if true {
                                    if ((!(1000 > 32767)) || false) {
                                        v1[v8] = (fi10 << ((1) & 63))
                                        switch "" {
                                        case "":
                                            v2 = (^(v2))
                                        case "hello":
                                            v22 := GenS3{F0: (fi10 / (((fi10 / ((127) | 1))) | 1)), F1: ((1.5 <= (0.0 + (3.14159 - 0.1))) || true)}
                                            _ = v22
                                        default:
                                            var v23 int64 = 1
                                            _ = v23
                                        }
                                    } else {
                                        var v24 uint32 = 42
                                        _ = v24
                                    }
                                }
                            }
                        case "hello":
                            if (1.5 <= (float64((float32((0.0 / 0.1)) + 1.5)) * ((1e10 / 1.5) - 0.0))) {
                                if (255 < 1) {
                                    v25 := GenS1{F0: v2, F1: 676}
                                    if len(v3) > 0 {
                                        v3[0] = (fi10 % (((^(fi18))) | 1))
                                    }
                                    v26 := GenS2{F0: genFn0(760), F1: 0, F2: -1}
                                    if false {
                                        fmt.Println(v2, v2, v8)
                                        switch "hello" {
                                        case "foo bar":
                                            var v27 int32 = 0
                                            _ = v27
                                        case "\n\t":
                                            fmt.Println(fi10)
                                        default:
                                            switch ("" + "") {
                                            case "\n\t":
                                                v28 := GenS4{F0: 128, F1: 1}
                                                if false {
                                                    v8 += 65535
                                                    v8 -= 948
                                                } else {
                                                    v8 = (v8 | 198)
                                                }
                                                v2 += 65535
                                                _ = v28
                                            case "foo bar":
                                                delete(v1, v8)
                                            }
                                        }
                                    }
                                    var v29 int16 = 1
                                    if len(v3) > 0 {
                                        v3[0] = (fi18 - fi18)
                                    }
                                    _ = v25
                                    _ = v26
                                    _ = v29
                                } else {
                                    var v30 int = -1
                                    _ = v30
                                }
                                if false {
                                    if true {
                                        var v31 int64 = 1
                                        var v32 uint16 = (v2 * v2)
                                        _ = v31
                                        _ = v32
                                    }
                                } else {
                                    var v33 uint8 = 2
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    if true {
                                        var v34 float64 = (0.0 - 0.1)
                                        var v35 int64 = -9223372036854775808
                                        var v36 string = ("a" + "hello")
                                        _ = v34
                                        _ = v35
                                        _ = v36
                                    }
                                    _ = v33
                                }
                                var v37 uint32 = (v8 + (v8 << ((65535) & 63)))
                                _ = v37
                            } else {
                                v1[2147483647] = fi10
                            }
                        }
                        _ = fi18
                    }
                    v8 -= v8
                } else {
                    fmt.Println(v8, v8, v2)
                    for fi38 := 0; fi38 < 7; fi38++ {
                        var v39 string = "\n\t"
                        var v40 int16 = -762
                        _ = fi38
                        _ = v39
                        _ = v40
                    }
                    v8 = v8
                }
                switch "\n\t" {
                case "":
                    var v41 int16 = 0
                    _ = v41
                case "foo bar":
                    v8 = (v8 * v8)
                case "\n\t":
                    v42 := v1[v8]
                    _ = v42
                }
                _ = fi10
            }
            fmt.Println(v2, v2)
        }
        switch 882 {
        case 4294967295:
            fmt.Println(v2, v2)
        case 1:
            v43 := GenS1{F0: (v2 | (v2 >> ((661) & 63))), F1: 0}
            _ = v43
        default:
            v2 = 747
        }
        _ = v8
    } else {
        for fi44 := 0; fi44 < 14; fi44++ {
            for fi45 := 0; fi45 < 8; fi45++ {
                switch "" {
                case "foo bar":
                    fmt.Println(fi45)
                }
                _ = fi45
            }
            v2 = genFn0(-1)
            var v46 uint64 = 42
            _ = fi44
            _ = v46
        }
        v2 -= (v2 % ((2) | 1))
    }
    var v47 int64 = -998
    v48 := v1[300]
    var v49 uint8 = 128
    switch 1 {
    case 32768:
        switch 32768 {
        case 32767:
            for fi50 := 0; fi50 < 14; fi50++ {
                var v51 int32 = 42
                for fi52 := 0; fi52 < 14; fi52++ {
                    switch 128 {
                    case 373:
                        v48 -= fi52
                    case -490:
                        if false {
                            v2 = v2
                            v53 := GenS5{F0: 276, F1: false}
                            _ = v53
                        } else {
                            var v54 string = (((("foo bar" + ((("" + "foo bar") + ("" + (("foo bar" + ("a" + "")) + ("a" + ("a" + ("hello" + "")))))) + "")) + "\n\t") + ("\n\t" + (((("hello" + ("hello" + ("a" + "hello"))) + ("foo bar" + "a")) + "foo bar") + "\n\t"))) + "")
                            if (!false) {
                                if false {
                                    fmt.Println(v2, v47)
                                    v54 = ""
                                    switch v54 {
                                    case "hello":
                                        v1[42] = 1
                                    case "foo bar":
                                        if false {
                                            fmt.Println(v49, v48, v54)
                                            v55 := GenS6{F0: 1}
                                            v51 += (^(v51))
                                            _ = v55
                                        }
                                    case "\n\t":
                                        var v56 string = v54
                                        _ = v56
                                    }
                                } else {
                                    var v57 int = fi50
                                    v3 = append(v3, 0)
                                    v58 := GenS5{F0: 65535, F1: (850 >= (v47 % ((v47) | 1)))}
                                    _ = v57
                                    _ = v58
                                }
                            }
                            fmt.Println(v48, fi52, v54)
                            if false {
                                v54 = (v54 + v54)
                            }
                            var v59 int64 = 0
                            _ = v54
                            _ = v59
                        }
                    default:
                        if (!(42 <= 37)) {
                            var v60 int32 = -2147483648
                            _ = v60
                        } else {
                            v61 := GenS5{F0: 2147483647, F1: false}
                            fmt.Println(v51, v47, v47)
                            _ = v61
                        }
                    }
                    var v62 int8 = -128
                    _ = fi52
                    _ = v62
                }
                _ = fi50
                _ = v51
            }
        case 128:
            for fi63 := 0; fi63 < 7; fi63++ {
                v47 = (v47 / ((v47) | 1))
                for fi64 := 0; fi64 < 7; fi64++ {
                    v65 := v1[2]
                    _ = fi64
                    _ = v65
                }
                _ = fi63
            }
        case 0:
            fmt.Println(v49, v48, v2)
        default:
            var v66 uint16 = v2
            _ = v66
        }
    case -1:
        if (255 <= 18446744073709551615) {
            var v67 int16 = 0
            v68 := GenS1{F0: 42, F1: 32767}
            switch ((("hello" + ("foo bar" + "foo bar")) + ("" + "foo bar")) + "\n\t") {
            case "":
                for fi69 := 0; fi69 < 2; fi69++ {
                    v70 := GenS1{F0: (v2 + (v2 >> ((842) & 63))), F1: v67}
                    if false {
                        if false {
                            v71 := v1[142]
                            for fi72 := 0; fi72 < 5; fi72++ {
                                var v73 int64 = (^(v47))
                                _ = fi72
                                _ = v73
                            }
                            var v74 int = -154
                            _ = v71
                            _ = v74
                        }
                        for fi75 := 0; fi75 < 8; fi75++ {
                            v67 -= 0
                            _ = fi75
                        }
                    } else {
                        v1[2147483647] = -1
                        v47 = int64(v48)
                    }
                    _ = fi69
                    _ = v70
                }
            default:
                if true {
                    if true {
                        v67 += (v67 & (v67 & v67))
                    }
                } else {
                    v2 -= v2
                }
            }
            _ = v67
            _ = v68
        }
    case 128:
        var v76 int16 = -32768
        _ = v76
    default:
        v77 := GenS7{F0: true, F1: 1, F2: -2147483648}
        _ = v77
    }
    var v78 uint = uint(v49)
    switch v48 {
    case 32767:
        fmt.Println(v49)
    }
    var v79 bool = (!((v48 & v48) == 171))
    if (!false) {
        delete(v1, 1)
    }
    if ((v78 / ((0) | 1)) <= 42) {
        v47 = -1
    } else {
        for fi80 := 0; fi80 < 8; fi80++ {
            v48 = int(v78)
            v47 = v47
            fmt.Println(fi80, v79, v2)
            _ = fi80
        }
        switch v48 {
        case -299:
            v47 = (v47 * (v47 << (((v78 | v78)) & 63)))
        case 2147483647:
            v81 := GenS1{F0: v2, F1: 1}
            _ = v81
        case 127:
            v48 = 387
        default:
            fmt.Println(v79, v49, v47)
        }
    }
    var v82 float32 = 3.14159
    var v83 uint16 = (v2 >> (((^(v78))) & 63))
    switch "" {
    case "foo bar":
        var v84 int = v48
        _ = v84
    default:
        var v85 uint8 = (v49 / (((^(v49))) | 1))
        _ = v85
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    fmt.Println("v82:", v82)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v47
    _ = v48
    _ = v49
    _ = v6
    _ = v7
    _ = v78
    _ = v79
    _ = v82
    _ = v83
}
