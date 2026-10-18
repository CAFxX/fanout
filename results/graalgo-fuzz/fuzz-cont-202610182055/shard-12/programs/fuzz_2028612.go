package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 string; F1 int8; F2 uint32 }
type GenS5 struct { F0 uint32; F1 uint32 }
type GenS6 struct { F0 uint8; F1 bool; F2 int32 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 string; F1 string; F2 int16 }
type GenS10 struct { F0 uint8; F1 string }

func main() {
    v1 := []uint32{2, 2147483647, 1, 2}
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]int)
    var v4 float32 = 1e10
    var v5 float32 = v4
    if true {
        switch -127 {
        case 42:
            fmt.Println(v5, v4)
        case 2147483647:
            if ((211 < 4294967295) && (false || false)) {
                if false {
                    v6 := GenS1{F0: 976}
                    for fi7 := 0; fi7 < 20; fi7++ {
                        v5 -= v5
                        var v8 int = fi7
                        _ = fi7
                        _ = v8
                    }
                    _ = v6
                }
                switch "\n\t" {
                case "a":
                    v5 = (-2.25 * 0.1)
                default:
                    switch ("a" + ("\n\t" + "a")) {
                    case "\n\t":
                        for fi9 := 0; fi9 < 14; fi9++ {
                            v10 := GenS1{F0: 2}
                            v11 := GenS1{F0: 2}
                            _ = fi9
                            _ = v10
                            _ = v11
                        }
                    case "hello":
                        if false {
                            var v12 uint64 = 42
                            _ = v12
                        } else {
                            if true {
                                if (!true) {
                                    if false {
                                        v5 = v5
                                    } else {
                                        if (1e10 < 1.5) {
                                            for fi13 := 0; fi13 < 4; fi13++ {
                                                for fi14 := 0; fi14 < 16; fi14++ {
                                                    v4 += (float32(0.1) * v5)
                                                    var v15 uint64 = 1
                                                    var v16 int16 = -32768
                                                    v17 := v3[2147483648]
                                                    v18 := GenS2{F0: 2, F1: 42}
                                                    _ = fi14
                                                    _ = v15
                                                    _ = v16
                                                    _ = v17
                                                    _ = v18
                                                }
                                                v1 = append(v1, 260)
                                                for fi19 := 0; fi19 < 7; fi19++ {
                                                    if false {
                                                        switch 1 {
                                                        case 42:
                                                            var v20 int = fi19
                                                            fmt.Println(v4)
                                                            if true {
                                                                switch "foo bar" {
                                                                case "hello":
                                                                    var v21 int32 = -2147483648
                                                                    _ = v21
                                                                default:
                                                                    v22 := GenS3{F0: 594}
                                                                    switch "a" {
                                                                    case "foo bar":
                                                                        v5 += v5
                                                                    case "\n\t":
                                                                        if true {
                                                                            fmt.Println(fi13, v20)
                                                                            var v23 bool = false
                                                                            var v24 bool = v23
                                                                            _ = v23
                                                                            _ = v24
                                                                        }
                                                                        switch "\n\t" {
                                                                        case "foo bar":
                                                                            v4 += v4
                                                                        }
                                                                        v20 = fi19
                                                                        v20 = fi13
                                                                        fmt.Println(v5, fi19)
                                                                        var v25 uint32 = 4294967295
                                                                        v5 = (((3.14159 * v4) / (float32((1.5 + 0.0)) / 0.1)) * float32((0.0 / 0.1)))
                                                                        _ = v25
                                                                    default:
                                                                        switch -770 {
                                                                        case -844:
                                                                            v20 -= 0
                                                                        }
                                                                    }
                                                                    _ = v22
                                                                }
                                                            }
                                                            _ = v20
                                                        case 27:
                                                            var v26 float32 = v4
                                                            _ = v26
                                                        }
                                                        if (1 > 0) {
                                                            var v27 float64 = (0.1 + 1e10)
                                                            v27 = (-2.25 - v27)
                                                            v5 = (1.5 + 3.14159)
                                                            _ = v27
                                                        } else {
                                                            switch 675 {
                                                            case 127:
                                                                if true {
                                                                    if (((0.0 * 0.1) >= float64((v4 * v4))) || (!(!true))) {
                                                                        var v28 int16 = 174
                                                                        var v29 int32 = 42
                                                                        var v30 float32 = v4
                                                                        _ = v28
                                                                        _ = v29
                                                                        _ = v30
                                                                    }
                                                                    var v31 uint8 = 42
                                                                    _ = v31
                                                                }
                                                            case 32767:
                                                                var v32 float64 = 3.14159
                                                                _ = v32
                                                            case -515:
                                                                if false {
                                                                    var v33 int16 = 1000
                                                                    _ = v33
                                                                }
                                                            default:
                                                                v5 = (1.5 - 3.14159)
                                                                if (4294967295 >= 187) {
                                                                    var v34 string = "foo bar"
                                                                    var v35 bool = (!true)
                                                                    _ = v34
                                                                    _ = v35
                                                                }
                                                            }
                                                            var v36 uint32 = 116
                                                            switch ("\n\t" + (((("hello" + "") + ("a" + ("a" + ((("foo bar" + "") + "hello") + "a")))) + "hello") + "hello")) {
                                                            case "a":
                                                                delete(v3, v36)
                                                            case "hello":
                                                                var v37 float64 = 0.0
                                                                _ = v37
                                                            }
                                                            _ = v36
                                                        }
                                                    }
                                                    if (18446744073709551615 >= 4294967295) {
                                                        v4 -= (v4 + 3.14159)
                                                        v38 := GenS4{F0: "", F1: 1, F2: 629}
                                                        _ = v38
                                                    }
                                                    v39 := v2[190]
                                                    _ = fi19
                                                    _ = v39
                                                }
                                                _ = fi13
                                            }
                                        }
                                        v5 = -2.25
                                    }
                                } else {
                                    v40 := GenS2{F0: 255, F1: -1}
                                    for fi41 := 0; fi41 < 13; fi41++ {
                                        switch ("\n\t" + ("hello" + ("" + "\n\t"))) {
                                        case "\n\t":
                                            v42 := GenS5{F0: 0, F1: 685}
                                            _ = v42
                                        case "hello":
                                            for fi43 := 0; fi43 < 18; fi43++ {
                                                fmt.Println(v4, fi41, fi41)
                                                _ = fi43
                                            }
                                        case "a":
                                            v4 = v5
                                        default:
                                            var v44 float64 = (1e10 * 0.1)
                                            _ = v44
                                        }
                                        for fi45 := 0; fi45 < 19; fi45++ {
                                            switch fi41 {
                                            case 704:
                                                var v46 float64 = (1e10 - ((-2.25 + float64(1.5)) + 0.1))
                                                _ = v46
                                            case 32768:
                                                v47 := GenS6{F0: 128, F1: true, F2: 61}
                                                _ = v47
                                            case 128:
                                                v1 = append(v1, 1)
                                            }
                                            var v48 int8 = 42
                                            _ = fi45
                                            _ = v48
                                        }
                                        v49 := GenS7{F0: ""}
                                        _ = fi41
                                        _ = v49
                                    }
                                    _ = v40
                                }
                                v4 = (v5 - v5)
                            } else {
                                switch "a" {
                                case "a":
                                    v5 -= 1e10
                                case "foo bar":
                                    v5 -= v5
                                }
                                switch ("foo bar" + ("a" + "a")) {
                                case "\n\t":
                                    fmt.Println(v5, v5)
                                case "hello":
                                    var v50 int64 = -9223372036854775808
                                    _ = v50
                                default:
                                    var v51 int64 = 42
                                    _ = v51
                                }
                            }
                        }
                    case "a":
                        var v52 int16 = 0
                        _ = v52
                    default:
                        v5 += (v5 * v5)
                    }
                }
                for fi53 := 0; fi53 < 13; fi53++ {
                    var v54 int32 = -2147483648
                    _ = fi53
                    _ = v54
                }
            }
        case 32768:
            var v55 float64 = 1e10
            _ = v55
        }
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v4 = 1.5
    _ = len(v3)
    v56 := GenS8{F0: 32767}
    v2[840] = 71
    fmt.Println(v5)
    if false {
        for fi57 := 0; fi57 < 9; fi57++ {
            if ((fi57 < fi57) || (0 != 1)) {
                v1 = append(v1, 890)
            } else {
                v58 := GenS2{F0: 255, F1: 42}
                var v59 int = fi57
                v1 = append(v1, 0)
                _ = v58
                _ = v59
            }
            v4 = ((v5 - (v4 - (1e10 * v5))) * float32((0.0 + 0.0)))
            _ = fi57
        }
        switch ("foo bar" + "foo bar") {
        case "":
            v60 := GenS9{F0: "", F1: "hello", F2: 821}
            _ = v60
        default:
            v61 := GenS2{F0: 2, F1: 529}
            _ = v61
        }
    } else {
        v5 = v4
    }
    var v62 string = "\n\t"
    if false {
        for fi63 := 0; fi63 < 17; fi63++ {
            v5 += -2.25
            var v64 float32 = 1e10
            for fi65 := 0; fi65 < 10; fi65++ {
                v66 := GenS10{F0: 80, F1: (v62 + v62)}
                _ = fi65
                _ = v66
            }
            _ = fi63
            _ = v64
        }
    } else {
        if false {
            v67 := GenS6{F0: 1, F1: (true && (-518 > 32767)), F2: 1}
            fmt.Println(v62, v5)
            v68 := GenS7{F0: ((v62 + ((v62 + v62) + ("hello" + v62))) + "")}
            _ = v67
            _ = v68
        } else {
            fmt.Println(v4)
            v4 = ((v4 - 0.0) * (0.0 + 1.5))
        }
        var v69 int64 = 9223372036854775807
        _ = v69
    }
    fmt.Println(v4)
    v62 = "a"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v62
}
