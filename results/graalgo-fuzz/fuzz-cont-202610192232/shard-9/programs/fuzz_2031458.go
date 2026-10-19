package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 string; F2 uint64 }
type GenS3 struct { F0 uint64; F1 string; F2 uint }
type GenS4 struct { F0 int16; F1 int64 }
type GenS5 struct { F0 uint64; F1 int16 }
type GenS6 struct { F0 int8; F1 string }
type GenS7 struct { F0 int8; F1 string }
type GenS8 struct { F0 int64; F1 int32 }

func main() {
    var v1 uint64 = 513
    v2 := []uint32{}
    v3 := make(map[uint32]int)
    v4 := GenS1{F0: true}
    _ = len(v2)
    for fi5 := 0; fi5 < 10; fi5++ {
        switch (("hello" + "") + "foo bar") {
        case "":
            fmt.Println(v1, fi5)
        default:
            for fi6 := 0; fi6 < 10; fi6++ {
                switch fi6 {
                case 2147483647:
                    if (15 > 216) {
                        fmt.Println(fi5)
                        if true {
                            var v7 string = "a"
                            v8 := GenS1{F0: (100 >= -1)}
                            _ = v7
                            _ = v8
                        }
                    }
                case 127:
                    var v9 uint32 = 4294967295
                    fmt.Println(v1, v1)
                    v9 = v9
                    delete(v3, v9)
                    if (846 <= -1) {
                        v3[(^(v9))] = fi5
                        fmt.Println(v9, v9)
                    } else {
                        v2 = append(v2, 745)
                    }
                    _ = v9
                case -554:
                    v10 := GenS1{F0: (false || true)}
                    _ = v10
                default:
                    v11 := GenS1{F0: true}
                    _ = v11
                }
                var v12 int32 = -2147483648
                v12 += (v12 << ((2) & 63))
                _ = fi6
                _ = v12
            }
        }
        if false {
            if true {
                var v13 int32 = -1
                v14 := GenS1{F0: (-2.25 <= 0.1)}
                _ = v13
                _ = v14
            }
            switch (("\n\t" + "a") + "foo bar") {
            case "":
                if (128 != 255) {
                    if ((1.5 + -2.25) == ((0.1 * (-2.25 + -2.25)) * (-2.25 - (0.0 / 0.1)))) {
                        var v15 uint = 1
                        switch 1 {
                        case 2147483647:
                            v16 := GenS1{F0: (!(false || true))}
                            _ = v16
                        case 32768:
                            var v17 int = -9223372036854775808
                            _ = v17
                        }
                        _ = v15
                    } else {
                        for fi18 := 0; fi18 < 11; fi18++ {
                            var v19 int32 = -1
                            v1 = (v1 + (v1 + 4294967295))
                            _ = fi18
                            _ = v19
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v20 bool = (!false)
                        _ = v20
                    }
                    fmt.Println(v1, v1)
                    var v21 int64 = 9223372036854775807
                    _ = v21
                }
            case "hello":
                switch "hello" {
                case "":
                    for fi22 := 0; fi22 < 17; fi22++ {
                        if true {
                            v1 = (v1 >> ((18446744073709551615) & 63))
                            var v23 uint32 = 65535
                            _ = v23
                        } else {
                            v24 := GenS2{F0: "a", F1: ("a" + ("a" + "")), F2: (v1 % ((18446744073709551615) | 1))}
                            fmt.Println(fi5)
                            if true {
                                switch (fi5 / (((fi22 << ((42) & 63))) | 1)) {
                                case -9223372036854775808:
                                    fmt.Println(fi22, fi5)
                                case -231:
                                    if false {
                                        switch (fi22 >> ((18446744073709551615) & 63)) {
                                        case 352:
                                            switch (("hello" + "") + "a") {
                                            case "foo bar":
                                                if true {
                                                    var v25 uint16 = 1
                                                    v26 := GenS1{F0: true}
                                                    fmt.Println(v25, fi5)
                                                    _ = v25
                                                    _ = v26
                                                } else {
                                                    switch ("a" + (("foo bar" + "foo bar") + "foo bar")) {
                                                    case "hello":
                                                        if (true || (!false)) {
                                                            if (-355 < 42) {
                                                                switch ("a" + "hello") {
                                                                case "hello":
                                                                    if true {
                                                                        var v27 uint16 = 630
                                                                        if len(v2) > 0 {
                                                                            v2[0] = 239
                                                                        }
                                                                        _ = v27
                                                                    } else {
                                                                        var v28 int8 = 100
                                                                        if ((-1 <= 2147483647) && (0.0 < (1e10 / ((((0.0 - 1.5) * 0.0) * (3.14159 * 3.14159)) + (3.14159 - (1.5 / (-2.25 * 3.14159))))))) {
                                                                            v29 := GenS3{F0: (v1 * 0), F1: "a", F2: uint(v1)}
                                                                            _ = v29
                                                                        }
                                                                        _ = v28
                                                                    }
                                                                case "\n\t":
                                                                    _ = len(v2)
                                                                case "":
                                                                    if (("\n\t" > ((("\n\t" + ((("" + ("foo bar" + "hello")) + "a") + "")) + "") + "a")) || false) {
                                                                        if true {
                                                                            switch "hello" {
                                                                            case "a":
                                                                                v1 = v1
                                                                            case "foo bar":
                                                                                if ((^(v1)) != v1) {
                                                                                    v1 -= v1
                                                                                    fmt.Println(fi5, v1, v1)
                                                                                }
                                                                            default:
                                                                                switch ("a" + "a") {
                                                                                case "":
                                                                                    v1 = v1
                                                                                case "foo bar":
                                                                                    v1 = v1
                                                                                case "\n\t":
                                                                                    if true {
                                                                                        var v30 int = fi22
                                                                                        var v31 string = ("\n\t" + ("hello" + ("foo bar" + "\n\t")))
                                                                                        _ = v30
                                                                                        _ = v31
                                                                                    } else {
                                                                                        var v32 int = fi5
                                                                                        _ = v32
                                                                                    }
                                                                                }
                                                                            }
                                                                            v1 -= v1
                                                                            delete(v3, 493)
                                                                        }
                                                                    }
                                                                }
                                                                switch (fi5 << ((255) & 63)) {
                                                                case 2147483647:
                                                                    var v33 uint64 = (^(v1))
                                                                    _ = v33
                                                                default:
                                                                    if false {
                                                                        v34 := GenS3{F0: (v1 * v1), F1: "a", F2: 65535}
                                                                        switch fi5 {
                                                                        case 4294967295:
                                                                            v1 += (^(v1))
                                                                        case -449:
                                                                            fmt.Println(v1, fi5, fi5)
                                                                        }
                                                                        v1 -= (v1 << ((1) & 63))
                                                                        _ = v34
                                                                    } else {
                                                                        switch 32767 {
                                                                        case 2147483647:
                                                                            var v35 int16 = 0
                                                                            _ = v35
                                                                        }
                                                                        fmt.Println(v1, fi22, v1)
                                                                        v1 = (v1 % (((v1 / ((823) | 1))) | 1))
                                                                    }
                                                                }
                                                                if true {
                                                                    var v36 uint64 = 1
                                                                    _ = v36
                                                                }
                                                                delete(v3, 0)
                                                                v37 := GenS4{F0: 1, F1: 1}
                                                                _ = v37
                                                            } else {
                                                                fmt.Println(fi22)
                                                                switch "a" {
                                                                case "foo bar":
                                                                    v38 := GenS3{F0: 1, F1: ("\n\t" + "foo bar"), F2: 42}
                                                                    _ = v38
                                                                default:
                                                                    var v39 uint64 = 0
                                                                    _ = v39
                                                                }
                                                                v1 = v1
                                                                fmt.Println(fi22, fi22, v1)
                                                            }
                                                            v1 = v1
                                                            v1 -= v1
                                                        }
                                                        switch fi5 {
                                                        case -1:
                                                            v40 := GenS5{F0: (v1 | 395), F1: -32768}
                                                            _ = v40
                                                        case 9223372036854775807:
                                                            if len(v2) > 0 {
                                                                v2[0] = 2147483647
                                                            }
                                                            var v41 bool = false
                                                            switch (fi5 - (fi5 - 934)) {
                                                            case -81:
                                                                v41 = (367 < 2147483647)
                                                            default:
                                                                var v42 int = (fi5 >> ((65535) & 63))
                                                                fmt.Println(v41, fi5, v42)
                                                                _ = v42
                                                            }
                                                            _ = v41
                                                        case 4294967295:
                                                            switch (("foo bar" + "a") + "\n\t") {
                                                            case "hello":
                                                                v1 = 188
                                                            default:
                                                                fmt.Println(fi22, v1)
                                                            }
                                                        }
                                                    }
                                                    v43 := v3[42]
                                                    fmt.Println(fi5)
                                                    _ = v43
                                                }
                                            }
                                        default:
                                            delete(v3, 4294967295)
                                        }
                                        fmt.Println(fi22, fi22)
                                        if (!false) {
                                            v44 := GenS6{F0: 1, F1: "\n\t"}
                                            var v45 uint32 = 65535
                                            v46 := GenS4{F0: 550, F1: -128}
                                            if true {
                                                _ = len(v2)
                                                v45 += uint32(v1)
                                            }
                                            _ = v44
                                            _ = v45
                                            _ = v46
                                        } else {
                                            v1 = v1
                                        }
                                        var v47 uint = 255
                                        _ = v47
                                    }
                                }
                                v1 = 820
                            }
                            var v48 uint = 861
                            v1 = 1
                            fmt.Println(v1)
                            _ = v24
                            _ = v48
                        }
                        var v49 uint16 = 65535
                        _ = fi22
                        _ = v49
                    }
                case "a":
                    v1 = 4294967295
                default:
                    var v50 uint32 = 42
                    _ = v50
                }
            case "a":
                v51 := GenS4{F0: 907, F1: -1}
                _ = v51
            default:
                v52 := GenS4{F0: 1000, F1: 9223372036854775807}
                _ = v52
            }
            var v53 uint32 = 275
            _ = v53
        } else {
            var v54 int8 = 95
            v55 := GenS7{F0: v54, F1: "hello"}
            switch "" {
            case "hello":
                switch ("a" + "a") {
                case "":
                    var v56 uint32 = 1
                    _ = v56
                case "\n\t":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case "a":
                    v54 += 54
                }
            case "a":
                fmt.Println(v1, v54, v54)
            default:
                v1 += v1
            }
            _ = v54
            _ = v55
        }
        switch (fi5 | -1) {
        case 2147483647:
            v1 = 18446744073709551615
        case 228:
            var v57 uint8 = 98
            _ = v57
        default:
            v1 = v1
        }
        _ = fi5
    }
    var v58 uint = 677
    if false {
        v3[2] = 0
    } else {
        v1 -= (v1 >> ((30) & 63))
    }
    var v59 int8 = -123
    v60 := GenS8{F0: -190, F1: 600}
    fmt.Println(v58)
    if true {
        fmt.Println(v58)
    } else {
        if ((!(0 != 0)) || (!(-32768 > -734))) {
            var v61 float32 = 1e10
            v62 := v3[2147483647]
            switch int(v58) {
            case 32767:
                v63 := GenS1{F0: true}
                _ = v63
            }
            _ = v61
            _ = v62
        }
        v59 += (v59 + (^(v59)))
        v3[2147483647] = 0
    }
    var v64 bool = (false && (255 < 65535))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v58
    _ = v59
    _ = v60
    _ = v64
}
