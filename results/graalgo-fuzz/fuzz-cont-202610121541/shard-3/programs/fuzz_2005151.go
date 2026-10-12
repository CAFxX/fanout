package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int; F1 uint16 }
type GenS5 struct { F0 int64; F1 int }
type GenS6 struct { F0 uint16; F1 uint32 }
type GenS7 struct { F0 bool; F1 int16; F2 int16 }
func genFn0(p0 uint, p1 int) int {
    p1 = p1
    var v5 uint32 = 2147483648
    v6 := []int{42, 1}
    if false {
        _ = len(v6)
        for fi7 := 0; fi7 < 20; fi7++ {
            var v8 int64 = 9223372036854775807
            v9 := GenS1{F0: true, F1: -32768}
            _ = fi7
            _ = v8
            _ = v9
        }
        v5 = (v5 - (v5 << ((301) & 63)))
    } else {
        if len(v6) > 0 {
            v6[0] = (^(p1))
        }
        var v10 uint = p0
        _ = v10
    }
    switch ("hello" + "\n\t") {
    case "a":
        var v11 uint16 = 350
        _ = v11
    case "hello":
        if (!(!false)) {
            var v12 float64 = 1.5
            fmt.Println(p0, v12)
            if false {
                var v13 int64 = 860
                var v14 int16 = -217
                _ = v13
                _ = v14
            } else {
                for fi15 := 0; fi15 < 7; fi15++ {
                    for fi16 := 0; fi16 < 19; fi16++ {
                        v6 = append(v6, -767)
                        p0 = p0
                        _ = fi16
                    }
                    switch fi15 {
                    case 32767:
                        var v17 int = (p1 + fi15)
                        _ = v17
                    case -610:
                        var v18 uint32 = v5
                        _ = v18
                    }
                    _ = fi15
                }
            }
            _ = v12
        }
    default:
        if (!(!(false || false))) {
            v19 := GenS1{F0: ((!false) || ((!(-1 == -1)) || false)), F1: -32768}
            fmt.Println(v5)
            _ = v19
        } else {
            for fi20 := 0; fi20 < 1; fi20++ {
                var v21 int64 = 208
                _ = fi20
                _ = v21
            }
            var v22 int = (p1 & 127)
            switch -546 {
            case 32768:
                var v23 int64 = 133
                _ = v23
            }
            _ = v22
        }
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (-(p1))
}

func main() {
    var v1 int = -817
    var v2 uint = 2
    var v3 bool = (100 == -128)
    var v4 uint16 = 255
    if v3 {
        var v24 float64 = -2.25
        switch "hello" {
        case "hello":
            for fi25 := 0; fi25 < 4; fi25++ {
                v3 = (9223372036854775807 >= fi25)
                v24 -= v24
                v26 := GenS2{F0: -1, F1: 1000}
                _ = fi25
                _ = v26
            }
        case "\n\t":
            fmt.Println(v24, v4)
        case "foo bar":
            if ("" != ("\n\t" + "")) {
                v1 -= (^(v1))
                for fi27 := 0; fi27 < 6; fi27++ {
                    v28 := []int{}
                    var v29 int8 = 85
                    for fi30 := 0; fi30 < 18; fi30++ {
                        if v3 {
                            var v31 int32 = 446
                            if v3 {
                                fmt.Println(fi30, fi27, v1)
                                if (172 <= 1) {
                                    if v3 {
                                        if v3 {
                                            switch genFn0(65535, (fi30 << ((v2) & 63))) {
                                            case -1:
                                                switch (fi27 << ((365) & 63)) {
                                                case 32767:
                                                    v29 = (v29 & (-(v29)))
                                                }
                                            case -9223372036854775808:
                                                if ((^(v31)) != 2147483647) {
                                                    var v32 float64 = v24
                                                    v1 += -366
                                                    v32 += -2.25
                                                    _ = v32
                                                } else {
                                                    if ("\n\t" > ("" + ("foo bar" + "foo bar"))) {
                                                        v33 := GenS1{F0: false, F1: 42}
                                                        fmt.Println(v2)
                                                        var v34 uint32 = 348
                                                        _ = v33
                                                        _ = v34
                                                    }
                                                    v35 := GenS1{F0: true, F1: 1000}
                                                    var v36 string = "foo bar"
                                                    v2 = 2
                                                    _ = v35
                                                    _ = v36
                                                }
                                            case 32767:
                                                v3 = false
                                            default:
                                                if v3 {
                                                    var v37 uint = (v2 & (v2 << (((v2 | (v2 * (v2 + v2)))) & 63)))
                                                    _ = v37
                                                }
                                            }
                                        } else {
                                            var v38 bool = (v31 == v31)
                                            _ = v38
                                        }
                                        if len(v28) >= 2 {
                                            v28 = v28[1:len(v28)]
                                        }
                                        if len(v28) > 0 {
                                            v28[0] = int(v4)
                                        }
                                        v1 = (fi30 & 2147483647)
                                    } else {
                                        var v39 int = -260
                                        v40 := GenS3{F0: -1}
                                        if v3 {
                                            v24 = 0.0
                                            var v41 int16 = -32768
                                            _ = v41
                                        } else {
                                            v42 := GenS4{F0: fi30, F1: v4}
                                            if len(v28) >= 2 {
                                                v28 = v28[1:len(v28)]
                                            }
                                            var v43 uint16 = (v4 % (((v4 * 42)) | 1))
                                            _ = v42
                                            _ = v43
                                        }
                                        _ = v39
                                        _ = v40
                                    }
                                } else {
                                    var v44 uint64 = 42
                                    v4 = v4
                                    var v45 int = -807
                                    _ = v44
                                    _ = v45
                                }
                                fmt.Println(fi30, v4)
                            }
                            _ = v31
                        } else {
                            var v46 int8 = v29
                            switch "" {
                            case "a":
                                v24 = float64(0.0)
                            case "hello":
                                var v47 uint16 = v4
                                _ = v47
                            case "":
                                var v48 uint = 42
                                var v49 uint = v48
                                v2 = (^(v2))
                                _ = v48
                                _ = v49
                            default:
                                var v50 uint = 1
                                _ = v50
                            }
                            fmt.Println(v1, v24)
                            _ = v46
                        }
                        switch "a" {
                        case "hello":
                            if v3 {
                                if v3 {
                                    v4 = (^(v4))
                                    var v51 float32 = 0.1
                                    if (((fi27 > fi30) || (!v3)) && (false && (!false))) {
                                        var v52 int32 = 2147483647
                                        v4 = v4
                                        _ = v52
                                    }
                                    _ = v51
                                } else {
                                    var v53 float64 = float64(1.5)
                                    var v54 float64 = float64((0.1 * 3.14159))
                                    v28 = append(v28, 217)
                                    _ = v53
                                    _ = v54
                                }
                                v29 = (v29 / (((v29 >> (((v2 >> ((4294967295) & 63))) & 63))) | 1))
                            } else {
                                fmt.Println(fi27, v3, fi30)
                            }
                        }
                        v55 := GenS1{F0: false, F1: -32768}
                        fmt.Println(v1, v2)
                        _ = fi30
                        _ = v55
                    }
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                var v56 int16 = 83
                _ = v56
            }
        }
        fmt.Println(v1, v1)
        _ = v24
    }
    v57 := []int{v1, genFn0(v2, (v1 >> (((v2 >> ((v2) & 63))) & 63))), (v1 ^ (v1 << (((v2 / (((v2 * (v2 & 4294967295))) | 1))) & 63)))}
    v1 -= (v1 / ((genFn0(v2, 32768)) | 1))
    var v58 uint32 = 65535
    var v59 int64 = -9223372036854775808
    var v60 float32 = 1e10
    var v61 uint8 = 128
    if len(v57) >= 2 {
        v57 = v57[1:len(v57)]
    }
    v2 = (v2 % ((990) | 1))
    for fi62 := 0; fi62 < 15; fi62++ {
        fmt.Println(fi62, v3)
        _ = fi62
    }
    if ((42 > uint64(v4)) || v3) {
        for fi63 := 0; fi63 < 14; fi63++ {
            if v3 {
                for fi64 := 0; fi64 < 6; fi64++ {
                    if ((!v3) || v3) {
                        if v3 {
                            switch "hello" {
                            case "foo bar":
                                var v65 float64 = float64(v60)
                                _ = v65
                            case "\n\t":
                                v60 = v60
                            case "hello":
                                var v66 int64 = (v59 & v59)
                                _ = v66
                            }
                            switch 4294967295 {
                            case -2147483648:
                                v67 := GenS4{F0: 2147483647, F1: 466}
                                _ = v67
                            case 970:
                                var v68 bool = true
                                _ = v68
                            }
                        } else {
                            if v3 {
                                if v3 {
                                    v69 := GenS2{F0: 0, F1: 42}
                                    v2 += v2
                                    switch "hello" {
                                    case "a":
                                        fmt.Println(v58, v2, v2)
                                    case "\n\t":
                                        if (!((v1 % ((fi64) | 1)) != 1)) {
                                            v57 = append(v57, (^(fi64)))
                                        }
                                    }
                                    _ = v69
                                }
                                switch (fi64 ^ fi63) {
                                case -2147483648:
                                    fmt.Println(v1)
                                default:
                                    v70 := GenS5{F0: v59, F1: (v1 >> ((v2) & 63))}
                                    _ = v70
                                }
                                v58 += 2147483648
                            }
                            v1 = genFn0((v2 + (v2 | 255)), (v1 + 42))
                        }
                        var v71 uint64 = 42
                        v57 = append(v57, 0)
                        _ = v71
                    }
                    v4 += (^(v4))
                    _ = fi64
                }
                var v72 int16 = -699
                var v73 int64 = (^(v59))
                _ = v72
                _ = v73
            }
            v74 := GenS6{F0: (v4 / ((v4) | 1)), F1: (v58 * 1)}
            for fi75 := 0; fi75 < 17; fi75++ {
                v1 = (fi75 % ((fi75) | 1))
                _ = fi75
            }
            _ = fi63
            _ = v74
        }
    } else {
        v76 := GenS7{F0: (((v60 / 1.5) - 1.5) < -2.25), F1: -32768, F2: 32767}
        _ = v76
    }
    fmt.Println(v58)
    var v77 uint32 = v58
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v57:", len(v57), v57)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v57
    _ = v58
    _ = v59
    _ = v60
    _ = v61
    _ = v77
}
