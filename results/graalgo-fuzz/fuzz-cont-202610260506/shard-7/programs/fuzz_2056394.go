package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 bool; F1 uint }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 bool; F1 int; F2 int32 }
type GenS8 struct { F0 int8; F1 uint64; F2 uint64 }

func main() {
    var v1 int8 = 100
    var v2 int = 916
    v3 := []int{v2, v2, v2, v2}
    var v4 float32 = 1e10
    v5 := make(map[uint32]string)
    fmt.Println(v2, v4, v1)
    for fi6 := 0; fi6 < 18; fi6++ {
        v1 = (v1 * 127)
        for fi7 := 0; fi7 < 12; fi7++ {
            v1 += 111
            if ((32767 != 1) && true) {
                v1 -= 100
            } else {
                v1 += -70
                if (false || false) {
                    v2 = -339
                    if (0 <= 2) {
                        v4 += 0.1
                        v8 := v5[956]
                        _ = v8
                    }
                    v9 := GenS1{F0: "a", F1: 65535}
                    _ = v9
                } else {
                    if true {
                        if false {
                            var v10 uint32 = 837
                            _ = v10
                        } else {
                            v1 += v1
                        }
                    } else {
                        switch "hello" {
                        case "foo bar":
                            var v11 bool = false
                            _ = v11
                        default:
                            v1 -= 0
                        }
                        fmt.Println(v4)
                    }
                    var v12 int64 = -412
                    var v13 float32 = ((v4 / (-2.25 / v4)) / 0.1)
                    _ = v12
                    _ = v13
                }
            }
            _ = fi7
        }
        fmt.Println(fi6)
        _ = fi6
    }
    if false {
        fmt.Println(v2)
    }
    if true {
        for fi14 := 0; fi14 < 15; fi14++ {
            v15 := GenS1{F0: ("foo bar" + "foo bar"), F1: 1}
            _ = fi14
            _ = v15
        }
        v16 := GenS2{F0: true}
        switch v2 {
        case 9223372036854775807:
            var v17 uint64 = 18446744073709551615
            _ = v17
        case 32767:
            for fi18 := 0; fi18 < 9; fi18++ {
                v4 -= v4
                for fi19 := 0; fi19 < 1; fi19++ {
                    v1 = v1
                    var v20 int8 = -57
                    switch "a" {
                    case "a":
                        if (true && (true || true)) {
                            var v21 float64 = (0.1 * 1e10)
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v4 -= (float32(0.0) - (v4 * v4))
                            _ = v21
                        } else {
                            fmt.Println(fi19)
                            switch ("\n\t" + "foo bar") {
                            case "foo bar":
                                v4 += (v4 + v4)
                            case "":
                                switch (("hello" + "hello") + "foo bar") {
                                case "\n\t":
                                    var v22 string = "foo bar"
                                    _ = v22
                                case "":
                                    switch 0 {
                                    case 127:
                                        if false {
                                            switch (fi19 | 32768) {
                                            case 1:
                                                v2 = (fi19 % ((fi19) | 1))
                                            default:
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            }
                                            var v23 int32 = 583
                                            v5[0] = "foo bar"
                                            v23 -= v23
                                            var v24 uint32 = 2
                                            _ = v23
                                            _ = v24
                                        } else {
                                            v25 := GenS1{F0: ("hello" + "hello"), F1: 232}
                                            switch ((("a" + "\n\t") + ((("\n\t" + "") + ("foo bar" + "foo bar")) + "foo bar")) + "\n\t") {
                                            case "foo bar":
                                                var v26 bool = (-32768 != 42)
                                                _ = v26
                                            case "a":
                                                fmt.Println(fi19, fi18)
                                            default:
                                                var v27 int = v2
                                                _ = v27
                                            }
                                            _ = v25
                                        }
                                    case -2147483648:
                                        if (!(true || true)) {
                                            v28 := GenS1{F0: "a", F1: 4294967295}
                                            _ = v28
                                        } else {
                                            fmt.Println(fi19, v4, fi19)
                                        }
                                    }
                                default:
                                    switch "foo bar" {
                                    case "hello":
                                        var v29 int64 = 42
                                        _ = v29
                                    default:
                                        var v30 int32 = -1
                                        _ = v30
                                    }
                                }
                            case "\n\t":
                                var v31 float64 = (((-2.25 - (float64((1.5 - ((float32(3.14159) + float32((-2.25 / (1e10 / 3.14159)))) + (3.14159 * (v4 + v4))))) - (3.14159 - 3.14159))) / ((0.1 - ((((3.14159 + 0.0) + 1e10) + -2.25) * 1.5)) - (1.5 * 3.14159))) * (0.1 + 0.0))
                                _ = v31
                            default:
                                v1 = (v20 * 99)
                            }
                        }
                    case "hello":
                        v4 -= 3.14159
                        if false {
                            if len(v3) > 0 {
                                v3[0] = v2
                            }
                        } else {
                            v32 := GenS2{F0: false}
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = v32
                        }
                    default:
                        _ = len(v5)
                    }
                    _ = fi19
                    _ = v20
                }
                var v33 int64 = -47
                _ = fi18
                _ = v33
            }
        case 73:
            if false {
                fmt.Println(v2, v2)
                v34 := GenS3{F0: true, F1: 4294967295}
                for fi35 := 0; fi35 < 20; fi35++ {
                    v1 += (v1 + v1)
                    v36 := GenS4{F0: 892}
                    _ = fi35
                    _ = v36
                }
                _ = v34
            } else {
                v4 = v4
                var v37 uint = 65535
                _ = v37
            }
        }
        _ = v16
    } else {
        v38 := GenS2{F0: (!(((0.1 + -2.25) * 3.14159) <= (float64(v4) - (float64(v4) / (1.5 + 1e10)))))}
        if true {
            for fi39 := 0; fi39 < 4; fi39++ {
                v2 += v2
                _ = fi39
            }
            fmt.Println(v1, v2)
            _ = len(v3)
        } else {
            var v40 int16 = 0
            if false {
                for fi41 := 0; fi41 < 4; fi41++ {
                    v42 := GenS5{F0: (true && (0 != 42)), F1: 0}
                    switch "hello" {
                    case "":
                        v40 -= 486
                    case "hello":
                        if (100 < 100) {
                            var v43 int64 = 42
                            switch fi41 {
                            case -90:
                                var v44 int = (v2 * 62)
                                _ = v44
                            case 127:
                                var v45 int16 = 32767
                                _ = v45
                            case 9223372036854775807:
                                v46 := GenS6{F0: 42}
                                _ = v46
                            }
                            _ = v43
                        } else {
                            if ((107 <= v1) || false) {
                                switch ("\n\t" + (((("hello" + "\n\t") + ("\n\t" + (((("foo bar" + "foo bar") + "hello") + "foo bar") + ("hello" + "\n\t")))) + "a") + ((("" + ("\n\t" + "\n\t")) + "") + ("\n\t" + (("" + "foo bar") + (("\n\t" + "foo bar") + "foo bar")))))) {
                                case "a":
                                    var v47 int = fi41
                                    _ = v47
                                case "":
                                    v1 = 127
                                case "foo bar":
                                    for fi48 := 0; fi48 < 8; fi48++ {
                                        var v49 uint16 = 65535
                                        v5[4294967295] = "hello"
                                        _ = fi48
                                        _ = v49
                                    }
                                default:
                                    for fi50 := 0; fi50 < 18; fi50++ {
                                        v51 := GenS4{F0: 2}
                                        v2 = (v2 / ((fi41) | 1))
                                        _ = fi50
                                        _ = v51
                                    }
                                }
                            } else {
                                switch (("\n\t" + "foo bar") + "hello") {
                                case "a":
                                    v52 := v5[4294967295]
                                    _ = v52
                                case "hello":
                                    v1 = (v1 + -128)
                                case "\n\t":
                                    for fi53 := 0; fi53 < 3; fi53++ {
                                        v1 -= v1
                                        _ = fi53
                                    }
                                default:
                                    v54 := GenS3{F0: ((v1 + v1) < 127), F1: 42}
                                    _ = v54
                                }
                                for fi55 := 0; fi55 < 4; fi55++ {
                                    if false {
                                        v4 = 1e10
                                        switch (v2 - 42) {
                                        case -300:
                                            v56 := GenS7{F0: false, F1: (fi41 - fi41), F2: 2147483647}
                                            _ = v56
                                        default:
                                            var v57 float64 = 1e10
                                            _ = v57
                                        }
                                        v58 := GenS8{F0: 5, F1: 381, F2: 802}
                                        v59 := GenS2{F0: (v40 != (v40 << ((822) & 63)))}
                                        _ = v58
                                        _ = v59
                                    }
                                    _ = fi55
                                }
                                var v60 int = fi41
                                _ = v60
                            }
                            switch ((("foo bar" + (("hello" + "hello") + ("foo bar" + ""))) + "") + ("a" + ("hello" + ("foo bar" + ((("" + "hello") + ("foo bar" + (("\n\t" + "foo bar") + "foo bar"))) + "a"))))) {
                            case "foo bar":
                                for fi61 := 0; fi61 < 20; fi61++ {
                                    var v62 int8 = v1
                                    _ = fi61
                                    _ = v62
                                }
                            case "\n\t":
                                var v63 int = v2
                                _ = v63
                            case "a":
                                var v64 uint8 = 1
                                _ = v64
                            default:
                                if (9223372036854775807 == 42) {
                                    v1 += v1
                                }
                            }
                        }
                    case "a":
                        if false {
                            if (!(!(255 != 112))) {
                                v65 := GenS1{F0: "\n\t", F1: 695}
                                if (true || false) {
                                    v66 := GenS8{F0: v1, F1: 4294967295, F2: 0}
                                    v67 := v5[329]
                                    _ = v66
                                    _ = v67
                                }
                                _ = v65
                            } else {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v68 int32 = 57
                                var v69 string = ("" + "\n\t")
                                _ = v68
                                _ = v69
                            }
                        }
                    }
                    _ = fi41
                    _ = v42
                }
                switch v2 {
                case 1:
                    v1 -= v1
                default:
                    switch "hello" {
                    case "hello":
                        var v70 int = (v2 / ((v2) | 1))
                        _ = v70
                    case "a":
                        var v71 uint = 104
                        _ = v71
                    case "\n\t":
                        v72 := GenS4{F0: 1}
                        _ = v72
                    }
                }
            } else {
                _ = len(v3)
            }
            for fi73 := 0; fi73 < 6; fi73++ {
                _ = len(v5)
                _ = fi73
            }
            _ = v40
        }
        _ = v38
    }
    if false {
        v1 = -42
    } else {
        var v74 uint32 = 807
        fmt.Println(v1, v2)
        var v75 int8 = v1
        _ = v74
        _ = v75
    }
    v76 := GenS1{F0: "foo bar", F1: 1}
    v1 = (v1 >> ((42) & 63))
    v77 := v5[65535]
    if (65535 > 42) {
        for fi78 := 0; fi78 < 15; fi78++ {
            fmt.Println(fi78, fi78, fi78)
            _ = fi78
        }
        v4 = 1e10
        for fi79 := 0; fi79 < 9; fi79++ {
            v77 = v77
            _ = fi79
        }
    } else {
        fmt.Println(v1)
        v80 := GenS3{F0: false, F1: 399}
        _ = v80
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v76
    _ = v77
}
