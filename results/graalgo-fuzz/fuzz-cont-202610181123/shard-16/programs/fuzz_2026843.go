package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 int32; F1 uint64; F2 int }
type GenS3 struct { F0 int64; F1 uint16 }
type GenS4 struct { F0 string; F1 uint16 }
type GenS5 struct { F0 int32; F1 int64; F2 int16 }
type GenS6 struct { F0 uint }

func main() {
    v1 := []string{"a", "a", ("hello" + ("foo bar" + ((("\n\t" + ("\n\t" + "\n\t")) + (("\n\t" + "a") + "\n\t")) + ("a" + (("" + ("foo bar" + "")) + "\n\t"))))), ("\n\t" + ("a" + ("foo bar" + ((("hello" + "\n\t") + "a") + "foo bar"))))}
    var v2 uint = 65535
    v3 := make(map[string]uint32)
    var v4 uint16 = 2
    for fi5 := 0; fi5 < 6; fi5++ {
        if (true || false) {
            v4 -= 1
        } else {
            for fi6 := 0; fi6 < 10; fi6++ {
                var v7 int = fi6
                if true {
                    switch "hello" {
                    case "":
                        v8 := v3[""]
                        _ = v8
                    case "a":
                        var v9 float32 = 1.5
                        _ = v9
                    case "hello":
                        v10 := GenS1{F0: v2, F1: v2}
                        _ = v10
                    default:
                        var v11 int32 = -192
                        var v12 int8 = 1
                        _ = v11
                        _ = v12
                    }
                    switch "a" {
                    case "foo bar":
                        v13 := GenS2{F0: -1, F1: 749, F2: v7}
                        _ = v13
                    case "\n\t":
                        if ((2147483647 == 2147483647) || false) {
                            fmt.Println(v4, v2)
                        } else {
                            switch fi6 {
                            case 1:
                                if false {
                                    v7 = (v7 << (((v2 % (((v2 / (((^(v2))) | 1))) | 1))) & 63))
                                    v2 += (v2 + (v2 - (v2 & 512)))
                                } else {
                                    var v14 uint32 = 995
                                    var v15 int = 893
                                    var v16 float64 = 0.1
                                    _ = v14
                                    _ = v15
                                    _ = v16
                                }
                            case 32767:
                                switch (^(fi5)) {
                                case 2147483647:
                                    var v17 int = fi5
                                    _ = v17
                                default:
                                    var v18 uint8 = 255
                                    v19 := GenS3{F0: 410, F1: (^(v4))}
                                    var v20 int16 = 0
                                    if (((v18 & v18) > 1) && false) {
                                        if true {
                                            if (v18 > v18) {
                                                var v21 uint = v2
                                                v21 -= v21
                                                var v22 int = (fi5 + (^(v7)))
                                                v23 := GenS2{F0: 0, F1: 4294967295, F2: fi6}
                                                var v24 uint32 = 2147483647
                                                var v25 bool = false
                                                delete(v3, ("a" + ""))
                                                if len(v1) > 0 {
                                                    v1[0] = ((("\n\t" + "\n\t") + "") + "\n\t")
                                                }
                                                _ = v21
                                                _ = v22
                                                _ = v23
                                                _ = v24
                                                _ = v25
                                            }
                                            var v26 bool = (!true)
                                            v26 = (-2.25 == (1e10 * -2.25))
                                            if v26 {
                                                if (uint64(v18) > 2) {
                                                    var v27 int = (fi6 >> ((v2) & 63))
                                                    var v28 uint8 = uint8(v2)
                                                    if len(v1) > 0 {
                                                        v1[0] = (("a" + "") + "a")
                                                    }
                                                    _ = v27
                                                    _ = v28
                                                }
                                                if (346 >= -984) {
                                                    var v29 float64 = 0.1
                                                    v30 := v3["hello"]
                                                    var v31 int64 = -9223372036854775808
                                                    _ = v29
                                                    _ = v30
                                                    _ = v31
                                                } else {
                                                    var v32 float32 = (1e10 * -2.25)
                                                    var v33 int64 = 9223372036854775807
                                                    var v34 uint64 = uint64(v4)
                                                    var v35 int64 = 366
                                                    var v36 uint64 = (v34 | v34)
                                                    _ = v32
                                                    _ = v33
                                                    _ = v34
                                                    _ = v35
                                                    _ = v36
                                                }
                                                var v37 int8 = 127
                                                _ = v37
                                            } else {
                                                _ = len(v3)
                                            }
                                            var v38 float64 = float64(((1.5 - 0.0) * 3.14159))
                                            _ = v26
                                            _ = v38
                                        }
                                    }
                                    v39 := GenS3{F0: 42, F1: (v4 + v4)}
                                    _ = v18
                                    _ = v19
                                    _ = v20
                                    _ = v39
                                }
                            }
                            v7 = fi6
                        }
                        v40 := GenS4{F0: ("hello" + "\n\t"), F1: (^(v4))}
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = v40
                    case "hello":
                        var v41 int = v7
                        _ = v41
                    }
                    if (true && false) {
                        fmt.Println(v7)
                        var v42 int64 = 42
                        _ = len(v1)
                        _ = v42
                    }
                }
                _ = fi6
                _ = v7
            }
        }
        var v43 int16 = 1
        _ = fi5
        _ = v43
    }
    v44 := v3[("a" + "hello")]
    for fi45 := 0; fi45 < 20; fi45++ {
        delete(v3, ("hello" + ("" + "hello")))
        if true {
            if false {
                switch ("foo bar" + "foo bar") {
                case "hello":
                    if true {
                        v46 := GenS3{F0: 9223372036854775807, F1: 42}
                        _ = v46
                    }
                case "\n\t":
                    switch (fi45 << ((4294967295) & 63)) {
                    case 130:
                        v4 += (v4 - v4)
                    case 1:
                        v2 += 42
                    default:
                        v3["a"] = v44
                    }
                case "a":
                    v47 := GenS4{F0: "\n\t", F1: (v4 | uint16(v2))}
                    _ = v47
                }
                for fi48 := 0; fi48 < 17; fi48++ {
                    v44 = (v44 & (v44 << ((v2) & 63)))
                    switch "foo bar" {
                    case "\n\t":
                        var v49 uint64 = 2
                        v2 += v2
                        _ = v49
                    case "":
                        v4 += (v4 / ((v4) | 1))
                    default:
                        switch "a" {
                        case "hello":
                            var v50 uint16 = 749
                            _ = v50
                        case "a":
                            v2 = 937
                        default:
                            fmt.Println(v2)
                        }
                    }
                    _ = fi48
                }
                _ = len(v1)
            } else {
                v44 = v44
                var v51 float64 = ((0.0 - 0.1) + 0.1)
                _ = v51
            }
            switch "\n\t" {
            case "":
                v2 = v2
            case "\n\t":
                if (v2 <= (v2 - 65535)) {
                    v52 := GenS5{F0: 42, F1: 42, F2: 1}
                    if false {
                        v2 += v2
                    } else {
                        var v53 int16 = -289
                        if false {
                            var v54 float32 = 1e10
                            _ = v54
                        }
                        _ = v53
                    }
                    var v55 string = "foo bar"
                    _ = v52
                    _ = v55
                }
            case "hello":
                var v56 uint32 = v44
                _ = v56
            }
        }
        var v57 int64 = 0
        _ = fi45
        _ = v57
    }
    var v58 string = (("" + "\n\t") + ("" + "hello"))
    var v59 uint32 = (v44 - uint32(v2))
    if (!(127 > -128)) {
        var v60 uint = v2
        _ = v60
    } else {
        var v61 uint16 = (v4 - 340)
        v62 := GenS6{F0: v2}
        _ = v61
        _ = v62
    }
    switch ("hello" + v58) {
    case "\n\t":
        switch 2147483647 {
        case 4294967295:
            var v63 uint32 = v59
            _ = v63
        case -2147483648:
            if ((42 > 127) || ((!true) || (0 < -1))) {
                switch "hello" {
                case "foo bar":
                    fmt.Println(v4)
                case "\n\t":
                    var v64 int64 = 9223372036854775807
                    _ = v64
                case "hello":
                    var v65 uint64 = 0
                    _ = v65
                default:
                    v4 += 1
                }
            }
        case 42:
            v1 = append(v1, "\n\t")
        }
    case "":
        for fi66 := 0; fi66 < 13; fi66++ {
            for fi67 := 0; fi67 < 17; fi67++ {
                var v68 uint16 = (^(v4))
                fmt.Println(v44, v4, fi67)
                var v69 bool = (!false)
                _ = fi67
                _ = v68
                _ = v69
            }
            var v70 bool = (true || true)
            _ = fi66
            _ = v70
        }
    case "hello":
        v71 := v3[v58]
        _ = v71
    default:
        fmt.Println(v58, v58)
    }
    v44 -= 1
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (97 > 2) {
        v44 -= v44
        _ = len(v3)
    } else {
        var v72 uint64 = 0
        var v73 int16 = 42
        v4 -= (v4 | (^(v4)))
        _ = v72
        _ = v73
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v58
    _ = v59
}
