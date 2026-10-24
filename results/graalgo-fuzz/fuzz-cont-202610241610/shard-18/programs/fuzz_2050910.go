package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int; F1 uint8; F2 int8 }
type GenS6 struct { F0 uint; F1 string }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int32; F1 int32 }
type GenS9 struct { F0 uint64 }

func main() {
    var v1 uint8 = 171
    v2 := make(map[int]string)
    v3 := make(map[int]bool)
    v4 := []uint32{4294967295, 4294967295, 183, 4294967295}
    var v5 uint32 = 2
    for fi6 := 0; fi6 < 11; fi6++ {
        _ = len(v4)
        for fi7 := 0; fi7 < 3; fi7++ {
            if true {
                var v8 int8 = 100
                _ = v8
            }
            if (!true) {
                _ = len(v4)
                v9 := v3[789]
                if ((!false) && true) {
                    var v10 uint64 = 667
                    _ = v10
                }
                var v11 float64 = (0.0 * (float64((3.14159 + 1.5)) - 0.1))
                _ = v11
                _ = v9
            } else {
                var v12 int8 = 48
                v13 := GenS1{F0: (v5 / ((2147483647) | 1))}
                _ = v12
                _ = v13
            }
            if (32767 > 1000) {
                v14 := v2[fi6]
                _ = len(v3)
                fmt.Println(fi7, v1, fi7)
                v14 = ("" + ("hello" + (v14 + "hello")))
                if ((((-2.25 - (float64(1.5) - 1.5)) - 1e10) - 0.0) == 0.0) {
                    switch v14 {
                    case "hello":
                        v15 := GenS1{F0: 58}
                        _ = v15
                    case "\n\t":
                        v14 = ("hello" + v14)
                    }
                    var v16 uint32 = v5
                    _ = v16
                } else {
                    switch (v14 + "\n\t") {
                    case "hello":
                        v5 = 670
                        v3[fi7] = (-1 <= 28)
                        v17 := GenS1{F0: (v5 << ((65535) & 63))}
                        if (18446744073709551615 == 4294967295) {
                            var v18 float64 = (1e10 - 0.0)
                            v19 := v2[(fi7 - (fi6 - 32768))]
                            switch v19 {
                            case "\n\t":
                                v14 = "hello"
                            case "a":
                                var v20 float64 = (v18 / (v18 - v18))
                                _ = v20
                            case "hello":
                                v21 := v3[fi7]
                                _ = v21
                            }
                            _ = v18
                            _ = v19
                        } else {
                            switch fi7 {
                            case 4294967295:
                                v3[fi7] = ((true || false) || (v5 == (^(v5))))
                            case 128:
                                v1 += v1
                            default:
                                var v22 int32 = -2147483648
                                _ = v22
                            }
                        }
                        _ = v17
                    case "a":
                        var v23 uint64 = 175
                        _ = v23
                    case "foo bar":
                        var v24 float32 = (1e10 - ((0.1 - float32(0.0)) / float32(-2.25)))
                        _ = v24
                    default:
                        fmt.Println(v14)
                    }
                    switch v14 {
                    case "a":
                        v3[9223372036854775807] = ((v1 + v1) == (v1 % ((v1) | 1)))
                    default:
                        v5 -= 2147483647
                    }
                }
                _ = v14
            }
            _ = fi7
        }
        _ = fi6
    }
    v5 = v5
    switch ("hello" + ("foo bar" + "a")) {
    case "foo bar":
        v1 = v1
    case "hello":
        var v25 int = 413
        _ = v25
    case "\n\t":
        v5 = 367
    }
    if true {
        for fi26 := 0; fi26 < 14; fi26++ {
            v27 := GenS1{F0: (v5 & 65535)}
            _ = fi26
            _ = v27
        }
        v28 := GenS2{F0: 1, F1: 32767}
        var v29 int16 = 0
        _ = v28
        _ = v29
    } else {
        if (false && true) {
            if ((-32768 >= -813) && (408 > uint64(v5))) {
                v30 := GenS3{F0: (v1 ^ uint8(v5)), F1: "hello"}
                for fi31 := 0; fi31 < 5; fi31++ {
                    var v32 int16 = 379
                    if ("\n\t" > "") {
                        v5 = v5
                    } else {
                        if true {
                            for fi33 := 0; fi33 < 14; fi33++ {
                                v32 += (v32 << ((150) & 63))
                                _ = fi33
                            }
                            for fi34 := 0; fi34 < 19; fi34++ {
                                switch "a" {
                                case "":
                                    if false {
                                        fmt.Println(v32, v1, v5)
                                        fmt.Println(fi34, v1)
                                        v32 -= -32768
                                    } else {
                                        if len(v4) > 0 {
                                            v4[0] = (v5 << ((1) & 63))
                                        }
                                        fmt.Println(v1, fi31)
                                        if (-26 >= -128) {
                                            v35 := GenS4{F0: (("\n\t" + "\n\t") + ("" + ("hello" + ("a" + (("foo bar" + "a") + (("hello" + (("\n\t" + "a") + "a")) + "foo bar"))))))}
                                            switch 938 {
                                            case 127:
                                                switch fi34 {
                                                case 2147483647:
                                                    v1 += v1
                                                case -739:
                                                    var v36 uint = 2
                                                    _ = v36
                                                }
                                            default:
                                                v5 -= 536
                                            }
                                            _ = v35
                                        }
                                    }
                                default:
                                    switch ("" + "a") {
                                    case "a":
                                        if true {
                                            v37 := GenS3{F0: (^(v1)), F1: ("a" + "hello")}
                                            switch "foo bar" {
                                            case "\n\t":
                                                v38 := v3[42]
                                                _ = v38
                                            case "hello":
                                                var v39 uint16 = 919
                                                _ = v39
                                            }
                                            _ = len(v3)
                                            if (true || (!true)) {
                                                v5 = 470
                                                fmt.Println(fi31, fi34)
                                            }
                                            if (!(255 > 891)) {
                                                var v40 uint64 = 528
                                                v41 := GenS5{F0: 128, F1: 0, F2: 127}
                                                v5 += 1
                                                _ = v40
                                                _ = v41
                                            }
                                            _ = v37
                                        }
                                    case "":
                                        switch (fi34 * -9223372036854775808) {
                                        case 1:
                                            v32 += (-(v32))
                                        case -9223372036854775808:
                                            v1 -= (v1 ^ (v1 / ((v1) | 1)))
                                            if (!(1 != 632)) {
                                                v32 = v32
                                            }
                                            switch 127 {
                                            case 682:
                                                v32 = 1000
                                            case 32768:
                                                v1 -= v1
                                            case 32767:
                                                var v42 uint = 1
                                                _ = v42
                                            default:
                                                switch "" {
                                                case "foo bar":
                                                    if ("" >= (("\n\t" + "\n\t") + "a")) {
                                                        switch 2147483647 {
                                                        case 9223372036854775807:
                                                            v1 -= v1
                                                            var v43 string = "hello"
                                                            fmt.Println(v32, fi31)
                                                            _ = v43
                                                        }
                                                        v44 := GenS6{F0: 150, F1: ((((("hello" + "") + (("hello" + "hello") + "\n\t")) + "\n\t") + "a") + (("a" + "a") + ""))}
                                                        _ = v44
                                                    }
                                                case "\n\t":
                                                    if (!true) {
                                                        if true {
                                                            var v45 int = (fi31 - (fi31 - 230))
                                                            _ = v45
                                                        } else {
                                                            switch fi31 {
                                                            case 32767:
                                                                var v46 float64 = 0.1
                                                                _ = v46
                                                            default:
                                                                var v47 int32 = 0
                                                                v48 := GenS7{F0: 0}
                                                                _ = v47
                                                                _ = v48
                                                            }
                                                            v5 = v5
                                                        }
                                                    } else {
                                                        var v49 uint = 255
                                                        v49 -= 767
                                                        v50 := GenS6{F0: (^(v49)), F1: "\n\t"}
                                                        var v51 uint32 = v5
                                                        _ = v49
                                                        _ = v50
                                                        _ = v51
                                                    }
                                                }
                                            }
                                        case -172:
                                            var v52 uint = 255
                                            _ = v52
                                        default:
                                            _ = len(v4)
                                        }
                                    default:
                                        v32 += (v32 << ((42) & 63))
                                    }
                                }
                                fmt.Println(v1, v5)
                                _ = fi34
                            }
                        } else {
                            var v53 uint64 = 18446744073709551615
                            fmt.Println(fi31)
                            v3[(fi31 - fi31)] = true
                            _ = v53
                        }
                        for fi54 := 0; fi54 < 4; fi54++ {
                            var v55 int = (fi31 ^ (fi54 / (((fi31 + (fi54 | fi31))) | 1)))
                            _ = fi54
                            _ = v55
                        }
                        for fi56 := 0; fi56 < 6; fi56++ {
                            v32 -= v32
                            v32 = v32
                            _ = fi56
                        }
                    }
                    _ = fi31
                    _ = v32
                }
                switch "a" {
                case "a":
                    var v57 int64 = 9223372036854775807
                    _ = v57
                case "foo bar":
                    if len(v4) > 0 {
                        v4[0] = uint32(v1)
                    }
                case "\n\t":
                    v58 := GenS8{F0: -1, F1: -1}
                    _ = v58
                }
                _ = v30
            }
        }
    }
    var v59 string = ("\n\t" + "a")
    var v60 string = v59
    v1 -= v1
    v61 := GenS9{F0: 4294967295}
    var v62 int8 = 0
    if true {
        delete(v3, 42)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v60
    _ = v61
    _ = v62
}
