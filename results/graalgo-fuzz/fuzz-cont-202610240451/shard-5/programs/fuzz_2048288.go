package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 uint8; F2 int16 }
type GenS4 struct { F0 int16; F1 int; F2 uint32 }
type GenS5 struct { F0 int; F1 uint }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int8; F1 int32; F2 int32 }
type GenS8 struct { F0 int32; F1 uint8 }
type GenS9 struct { F0 string; F1 string }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint = 42
    v3 := []uint32{42, 1, 2147483648}
    var v4 int8 = 42
    if true {
        v1["hello"] = 2147483647
    }
    switch -1 {
    case 42:
        fmt.Println(v2)
    case 32767:
        var v5 int = 234
        _ = v5
    default:
        var v6 uint32 = 166
        _ = v6
    }
    switch -956 {
    case 128:
        v7 := GenS1{F0: 106}
        _ = v7
    }
    if true {
        switch 42 {
        case 42:
            for fi8 := 0; fi8 < 10; fi8++ {
                var v9 uint = 255
                v10 := GenS1{F0: 1}
                _ = fi8
                _ = v10
                _ = v9
            }
        default:
            if false {
                fmt.Println(v4, v4)
                for fi11 := 0; fi11 < 19; fi11++ {
                    delete(v1, "foo bar")
                    v2 = (v2 * v2)
                    if (!(("hello" > "hello") && true)) {
                        if true {
                            v2 = v2
                        }
                        switch 32767 {
                        case -882:
                            v4 += -7
                        case 0:
                            v4 = (v4 ^ 56)
                        }
                        v2 += v2
                    } else {
                        for fi12 := 0; fi12 < 8; fi12++ {
                            if true {
                                var v13 uint32 = 38
                                v2 += v2
                                if (42 > 0) {
                                    fmt.Println(v4, fi12, v13)
                                    v14 := GenS2{F0: 255}
                                    fmt.Println(v2, fi11, v4)
                                    var v15 string = "\n\t"
                                    var v16 uint = v2
                                    if true {
                                        switch fi11 {
                                        case 127:
                                            if false {
                                                if true {
                                                    var v17 int16 = 1
                                                    var v18 uint8 = 2
                                                    v18 += 151
                                                    _ = v17
                                                    _ = v18
                                                }
                                                var v19 float64 = 3.14159
                                                _ = v19
                                            } else {
                                                v20 := GenS2{F0: 702}
                                                _ = v20
                                            }
                                        case 32767:
                                            v13 += (^(v13))
                                        default:
                                            fmt.Println(fi11, fi11, fi11)
                                        }
                                        if true {
                                            v16 -= v2
                                        } else {
                                            switch fi11 {
                                            case 32767:
                                                switch fi12 {
                                                case -9223372036854775808:
                                                    if (255 != 128) {
                                                        v15 = v15
                                                        var v21 int64 = 1
                                                        _ = v21
                                                    }
                                                }
                                                v16 = (v16 << (((v2 ^ 183)) & 63))
                                                var v22 int64 = 0
                                                _ = v22
                                            case 127:
                                                switch "\n\t" {
                                                case "foo bar":
                                                    if (!false) {
                                                        v13 = v13
                                                    } else {
                                                        var v23 bool = false
                                                        _ = v23
                                                    }
                                                }
                                            }
                                            v24 := GenS3{F0: 1, F1: 190, F2: 2}
                                            v13 += (v13 % ((v13) | 1))
                                            _ = v24
                                        }
                                        switch v15 {
                                        case "":
                                            v25 := GenS4{F0: -32768, F1: (fi12 | 42), F2: (v13 + (^(v13)))}
                                            _ = v25
                                        }
                                    }
                                    _ = v14
                                    _ = v15
                                    _ = v16
                                } else {
                                    v26 := GenS4{F0: 1000, F1: (fi12 + fi12), F2: 616}
                                    var v27 uint8 = 57
                                    if (2 >= 436) {
                                        var v28 uint = 65535
                                        fmt.Println(fi11, v2, v27)
                                        var v29 uint64 = 192
                                        _ = v28
                                        _ = v29
                                    }
                                    _ = len(v3)
                                    _ = v26
                                    _ = v27
                                }
                                _ = v13
                            } else {
                                if true {
                                    if (236 == 42) {
                                        switch fi11 {
                                        case -545:
                                            if true {
                                                v4 += v4
                                            } else {
                                                var v30 int = (fi11 | 128)
                                                _ = v30
                                            }
                                        case -646:
                                            v31 := GenS4{F0: -1, F1: (-(fi11)), F2: 51}
                                            _ = v31
                                        case 127:
                                            var v32 uint8 = 0
                                            _ = v32
                                        }
                                        fmt.Println(v4)
                                        var v33 uint32 = 557
                                        _ = v33
                                    } else {
                                        delete(v1, "foo bar")
                                        if true {
                                            var v34 int16 = -1
                                            _ = v34
                                        }
                                    }
                                    switch "a" {
                                    case "":
                                        switch ("hello" + (("a" + "a") + "")) {
                                        case "hello":
                                            _ = len(v3)
                                        case "\n\t":
                                            var v35 uint64 = 18446744073709551615
                                            _ = v35
                                        }
                                    }
                                    v2 = 255
                                }
                                var v36 uint8 = 255
                                fmt.Println(v4, fi11, v4)
                                _ = v36
                            }
                            var v37 uint16 = 32767
                            switch "hello" {
                            case "\n\t":
                                var v38 uint16 = (v37 - uint16(v2))
                                _ = v38
                            case "hello":
                                fmt.Println(v4)
                            default:
                                if false {
                                    if len(v3) > 0 {
                                        v3[0] = 2
                                    }
                                    switch "" {
                                    case "a":
                                        fmt.Println(v4, v4)
                                    case "":
                                        switch "a" {
                                        case "foo bar":
                                            v37 = v37
                                        default:
                                            var v39 uint8 = 128
                                            _ = v39
                                        }
                                    case "foo bar":
                                        if true {
                                            v40 := GenS1{F0: 128}
                                            _ = v40
                                        } else {
                                            v4 -= v4
                                            delete(v1, "a")
                                            v2 = 65535
                                        }
                                        v37 -= (v37 & v37)
                                    }
                                }
                            }
                            switch "hello" {
                            case "":
                                if true {
                                    if true {
                                        if (v37 <= (v37 * v37)) {
                                            if false {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                if false {
                                                    var v41 int64 = -9223372036854775808
                                                    v42 := GenS2{F0: 65535}
                                                    _ = v41
                                                    _ = v42
                                                }
                                            } else {
                                                delete(v1, ("hello" + (("\n\t" + (("a" + ("hello" + ("foo bar" + ("foo bar" + ("foo bar" + "foo bar"))))) + "hello")) + "\n\t")))
                                                v37 -= v37
                                            }
                                            v43 := GenS5{F0: (fi11 + (fi12 | fi12)), F1: v2}
                                            v44 := GenS6{F0: uint32(v2)}
                                            _ = v43
                                            _ = v44
                                        } else {
                                            fmt.Println(v2, v4)
                                            var v45 uint16 = (v37 % ((v37) | 1))
                                            _ = v45
                                        }
                                        if false {
                                            switch (fi11 * 32767) {
                                            case 4294967295:
                                                var v46 uint64 = 0
                                                _ = v46
                                            case -2147483648:
                                                v37 = v37
                                            }
                                            if false {
                                                _ = len(v1)
                                            } else {
                                                v37 = 2
                                                if false {
                                                    v37 = (v37 / ((v37) | 1))
                                                    if (-635 < 187) {
                                                        v1["a"] = 1
                                                        switch fi12 {
                                                        case 373:
                                                            var v47 uint32 = 2
                                                            _ = v47
                                                        case -366:
                                                            if ((!(!true)) && ((((v37 + v37) <= 255) || (-29 <= -32768)) && true)) {
                                                                v4 -= v4
                                                                v4 = (-(v4))
                                                                var v48 uint64 = 4294967295
                                                                _ = v48
                                                            } else {
                                                                if true {
                                                                    _ = len(v3)
                                                                    var v49 float64 = 1.5
                                                                    _ = v49
                                                                } else {
                                                                    v50 := v1[(("hello" + "a") + (("" + "") + "foo bar"))]
                                                                    v51 := GenS2{F0: v37}
                                                                    _ = v50
                                                                    _ = v51
                                                                }
                                                            }
                                                        default:
                                                            v2 -= v2
                                                        }
                                                        v52 := v1["foo bar"]
                                                        _ = v52
                                                    } else {
                                                        fmt.Println(v2)
                                                        v2 = (v2 % ((18446744073709551615) | 1))
                                                        var v53 float64 = ((0.1 + float64(-2.25)) - float64((float32((1.5 + 0.1)) / 0.1)))
                                                        _ = v53
                                                    }
                                                    v54 := GenS7{F0: v4, F1: 0, F2: -580}
                                                    v4 -= v4
                                                    var v55 uint64 = 746
                                                    var v56 string = "foo bar"
                                                    var v57 float64 = -2.25
                                                    var v58 int16 = -32768
                                                    var v59 int8 = v4
                                                    _ = v54
                                                    _ = v55
                                                    _ = v56
                                                    _ = v57
                                                    _ = v58
                                                    _ = v59
                                                } else {
                                                    v37 = uint16(v2)
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                }
                                                var v60 int16 = -903
                                                _ = v60
                                            }
                                            v2 += (v2 ^ (v2 + v2))
                                        } else {
                                            var v61 uint = (v2 >> ((1) & 63))
                                            var v62 bool = (false && ((v61 >> ((18446744073709551615) & 63)) < (v2 % (((^(v61))) | 1))))
                                            var v63 bool = (!true)
                                            _ = v61
                                            _ = v62
                                            _ = v63
                                        }
                                        _ = len(v1)
                                    }
                                }
                            case "hello":
                                var v64 int32 = 42
                                _ = v64
                            case "\n\t":
                                v37 -= v37
                            default:
                                var v65 uint16 = (^(v37))
                                _ = v65
                            }
                            _ = fi12
                            _ = v37
                        }
                        fmt.Println(fi11, v2, v2)
                    }
                    _ = fi11
                }
                delete(v1, (("hello" + "") + "\n\t"))
            }
        }
    } else {
        v4 = v4
        v4 = -128
    }
    for fi66 := 0; fi66 < 17; fi66++ {
        v2 = 435
        var v67 int8 = (v4 << ((0) & 63))
        _ = fi66
        _ = v67
    }
    v2 = 0
    v2 += 18446744073709551615
    v4 -= v4
    v68 := GenS8{F0: -1, F1: 163}
    v4 = v4
    v3 = append(v3, 65535)
    v3 = append(v3, 10)
    for fi69 := 0; fi69 < 16; fi69++ {
        var v70 float32 = float32((1.5 - 3.14159))
        var v71 uint = (v2 - (v2 ^ v2))
        for fi72 := 0; fi72 < 11; fi72++ {
            if false {
                var v73 string = (((("a" + "\n\t") + ("foo bar" + "a")) + "foo bar") + "foo bar")
                v71 = (v2 + v2)
                _ = v73
            } else {
                var v74 uint = 1
                v2 = 1
                v75 := GenS9{F0: ("hello" + "hello"), F1: "\n\t"}
                var v76 int16 = 1
                var v77 uint64 = 4294967295
                switch ((("hello" + ("a" + "")) + "\n\t") + ("hello" + "\n\t")) {
                case "a":
                    fmt.Println(v4)
                case "foo bar":
                    if true {
                        if (472 == (^(v71))) {
                            var v78 uint8 = 1
                            _ = v78
                        }
                        v70 = v70
                    }
                case "":
                    var v79 int8 = v4
                    _ = v79
                }
                _ = v74
                _ = v75
                _ = v76
                _ = v77
            }
            _ = fi72
        }
        _ = fi69
        _ = v70
        _ = v71
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v68
}
