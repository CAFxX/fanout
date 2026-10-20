package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 int; F1 uint16; F2 string }
type GenS5 struct { F0 bool; F1 int }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint64; F1 int16 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint16; F1 uint }
type GenS10 struct { F0 int8; F1 uint; F2 int8 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 uint64 }
func genFn0(p0 uint16, p1 int16, p2 string) string {
    var v5 string = p2
    if true {
        v6 := GenS1{F0: "a", F1: 264, F2: -261}
        _ = v6
    } else {
        var v7 float64 = 3.14159
        _ = v7
    }
    if (false || ((1 != 0) || false)) {
        var v8 float32 = ((1.5 - 0.1) - 1e10)
        for fi9 := 0; fi9 < 3; fi9++ {
            for fi10 := 0; fi10 < 19; fi10++ {
                var v11 int16 = p1
                fmt.Println(p2, v11)
                var v12 int32 = -2147483648
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = fi9
        }
        _ = v8
    }
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return v5
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{"a", "\n\t", "a", "foo bar"}
    var v3 float64 = float64(0.0)
    v4 := []string{"foo bar", (("foo bar" + "hello") + "a"), "hello"}
    v3 = v3
    var v13 uint8 = 127
    if ((!(0 >= 0)) && ((4294967295 > 70) && false)) {
        if ((2147483647 != 2147483647) || ("" == ("foo bar" + "\n\t"))) {
            if (42 == 65535) {
                var v14 bool = true
                _ = v14
            }
        } else {
            switch 0 {
            case 32767:
                var v15 bool = true
                _ = v15
            case 1:
                var v16 uint64 = 578
                _ = v16
            }
            if ((18446744073709551615 <= 4294967295) || false) {
                switch "a" {
                case "foo bar":
                    if (4 == -92) {
                        if (887 <= 18446744073709551615) {
                            v13 = 255
                            v17 := GenS2{F0: -1}
                            _ = v17
                        } else {
                            if (1.5 > (float32(v3) / -2.25)) {
                                if true {
                                    v18 := GenS2{F0: -162}
                                    v19 := v1[2147483648]
                                    _ = v18
                                    _ = v19
                                }
                                if true {
                                    switch ("\n\t" + ("foo bar" + "hello")) {
                                    case "":
                                        v2 = append(v2, "foo bar")
                                    default:
                                        switch ("" + "\n\t") {
                                        case "foo bar":
                                            for fi20 := 0; fi20 < 11; fi20++ {
                                                for fi21 := 0; fi21 < 4; fi21++ {
                                                    if true {
                                                        v22 := GenS2{F0: 0}
                                                        if ("\n\t" != "\n\t") {
                                                            v13 += v13
                                                            var v23 int32 = 42
                                                            var v24 int8 = -128
                                                            _ = v23
                                                            _ = v24
                                                        } else {
                                                            if (true || (255 >= 0)) {
                                                                v25 := GenS3{F0: 1, F1: 18446744073709551615}
                                                                if false {
                                                                    v26 := v1[28]
                                                                    v27 := GenS4{F0: fi20, F1: 0, F2: "\n\t"}
                                                                    _ = v26
                                                                    _ = v27
                                                                }
                                                                if (false && false) {
                                                                    v13 = v13
                                                                    var v28 int8 = -128
                                                                    _ = v28
                                                                }
                                                                _ = v25
                                                            }
                                                            var v29 uint64 = 42
                                                            v1[2147483648] = "\n\t"
                                                            _ = v29
                                                        }
                                                        _ = v22
                                                    }
                                                    if true {
                                                        v30 := v1[2]
                                                        switch (fi21 % ((128) | 1)) {
                                                        case 1:
                                                            v30 = v30
                                                        default:
                                                            v31 := GenS5{F0: false, F1: 154}
                                                            _ = v31
                                                        }
                                                        _ = v30
                                                    } else {
                                                        if false {
                                                            v32 := GenS6{F0: 100}
                                                            v13 -= (v13 & v13)
                                                            _ = v32
                                                        } else {
                                                            v33 := GenS5{F0: true, F1: 42}
                                                            _ = v33
                                                        }
                                                        fmt.Println(fi21, fi20, v13)
                                                        fmt.Println(fi20, fi21, v13)
                                                    }
                                                    v3 -= float64(float32(v3))
                                                    _ = fi21
                                                }
                                                if true {
                                                    v4 = append(v4, "")
                                                    v13 = (v13 & v13)
                                                }
                                                _ = fi20
                                            }
                                        case "a":
                                            fmt.Println(v3)
                                        case "\n\t":
                                            for fi34 := 0; fi34 < 20; fi34++ {
                                                var v35 int16 = -32768
                                                _ = fi34
                                                _ = v35
                                            }
                                        }
                                    }
                                    var v36 int = 9223372036854775807
                                    switch v36 {
                                    case 0:
                                        var v37 int = 128
                                        _ = v37
                                    case 42:
                                        v38 := GenS7{F0: 4294967295, F1: -77}
                                        _ = v38
                                    }
                                    _ = v36
                                }
                            }
                            if true {
                                v13 -= 42
                                _ = len(v2)
                            } else {
                                v3 += (v3 - ((((v3 + v3) - v3) / float64((float32((v3 + 1e10)) * 1e10))) / ((v3 - 3.14159) - -2.25)))
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                            v3 += 1.5
                        }
                    }
                }
                delete(v1, 1)
                switch 127 {
                case -234:
                    v39 := v1[4294967295]
                    _ = v39
                default:
                    for fi40 := 0; fi40 < 1; fi40++ {
                        for fi41 := 0; fi41 < 17; fi41++ {
                            var v42 float64 = 0.0
                            v3 = v42
                            _ = fi41
                            _ = v42
                        }
                        v43 := GenS2{F0: -1}
                        var v44 float32 = 1.5
                        _ = fi40
                        _ = v43
                        _ = v44
                    }
                }
            }
            for fi45 := 0; fi45 < 10; fi45++ {
                v13 += v13
                _ = fi45
            }
        }
    }
    var v46 uint32 = 979
    v47 := GenS8{F0: 7}
    if true {
        if (125 != (v13 + (v13 / ((v13) | 1)))) {
            v13 += v13
        } else {
            var v48 string = "\n\t"
            var v49 float64 = 0.0
            for fi50 := 0; fi50 < 2; fi50++ {
                if (!true) {
                    v51 := GenS4{F0: 32768, F1: 255, F2: v48}
                    var v52 uint64 = 2
                    switch (-(fi50)) {
                    case 919:
                        if true {
                            switch (^(fi50)) {
                            case -465:
                                switch int(v13) {
                                case 9223372036854775807:
                                    if (false && (!(v49 == float64(((1e10 - (1.5 + 1.5)) - (0.1 - (3.14159 + -2.25))))))) {
                                        fmt.Println(v52)
                                    } else {
                                        var v53 uint = 1
                                        _ = v53
                                    }
                                case 128:
                                    for fi54 := 0; fi54 < 13; fi54++ {
                                        v13 -= (v13 << ((42) & 63))
                                        v13 += (v13 - 255)
                                        var v55 uint64 = 42
                                        if (true && false) {
                                            v56 := GenS1{F0: "a", F1: 18446744073709551615, F2: 0}
                                            _ = len(v2)
                                            _ = v56
                                        }
                                        v57 := GenS6{F0: 42}
                                        _ = fi54
                                        _ = v55
                                        _ = v57
                                    }
                                default:
                                    v13 += 1
                                }
                            default:
                                for fi58 := 0; fi58 < 9; fi58++ {
                                    v3 = -2.25
                                    var v59 int64 = -609
                                    var v60 uint64 = 4294967295
                                    _ = fi58
                                    _ = v59
                                    _ = v60
                                }
                            }
                            switch 841 {
                            case 9223372036854775807:
                                var v61 int8 = -1
                                _ = v61
                            case 127:
                                if false {
                                    fmt.Println(v13, v52, v48)
                                    if ((^(v52)) == v52) {
                                        v48 = ("a" + "hello")
                                        if (fi50 > fi50) {
                                            v3 += v3
                                        } else {
                                            if (true || true) {
                                                for fi62 := 0; fi62 < 16; fi62++ {
                                                    var v63 float64 = (v49 - 0.1)
                                                    _ = fi62
                                                    _ = v63
                                                }
                                            }
                                            v64 := GenS1{F0: "a", F1: 42, F2: -2147483648}
                                            _ = v64
                                        }
                                        for fi65 := 0; fi65 < 4; fi65++ {
                                            var v66 int8 = 42
                                            _ = fi65
                                            _ = v66
                                        }
                                    } else {
                                        fmt.Println(v48, v46)
                                        v67 := GenS9{F0: 32767, F1: 255}
                                        v49 = v3
                                        _ = v67
                                    }
                                }
                            case 42:
                                v68 := GenS10{F0: 42, F1: 957, F2: int8(v13)}
                                _ = v68
                            }
                            v49 = v3
                        } else {
                            delete(v1, uint32(v52))
                            if (42 < 520) {
                                if true {
                                    var v69 float32 = 0.0
                                    v3 -= v49
                                    _ = v69
                                } else {
                                    v49 += ((0.1 + v49) / 1.5)
                                    _ = len(v1)
                                }
                                var v70 int16 = 42
                                switch v48 {
                                case "\n\t":
                                    v13 = v13
                                }
                                _ = v70
                            }
                        }
                    case 42:
                        var v71 float64 = v3
                        _ = v71
                    case 606:
                        var v72 bool = true
                        _ = v72
                    default:
                        v46 -= (^(v46))
                    }
                    _ = v51
                    _ = v52
                } else {
                    fmt.Println(fi50, fi50, v13)
                    var v73 bool = true
                    _ = v73
                }
                _ = fi50
            }
            _ = v48
            _ = v49
        }
        switch (("hello" + ("foo bar" + (("foo bar" + "\n\t") + "foo bar"))) + "foo bar") {
        case "\n\t":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    }
    v74 := GenS11{F0: ((false && true) && false)}
    if true {
        var v75 uint = 42
        _ = v75
    } else {
        if (42 >= 745) {
            for fi76 := 0; fi76 < 2; fi76++ {
                fmt.Println(v13)
                if true {
                    if true {
                        v4 = append(v4, ((("hello" + ("foo bar" + "hello")) + ("\n\t" + "hello")) + ("hello" + "foo bar")))
                        var v77 float32 = 3.14159
                        _ = v77
                    }
                }
                _ = fi76
            }
            for fi78 := 0; fi78 < 10; fi78++ {
                for fi79 := 0; fi79 < 4; fi79++ {
                    var v80 bool = true
                    delete(v1, 2147483647)
                    _ = fi79
                    _ = v80
                }
                v46 = (v46 % (((^(v46))) | 1))
                for fi81 := 0; fi81 < 6; fi81++ {
                    switch fi78 {
                    case 0:
                        if true {
                            var v82 float32 = ((0.0 * 1.5) * (1e10 * ((0.1 - float32(float64(-2.25))) - (float32(0.1) + 0.0))))
                            fmt.Println(fi81, v46)
                            _ = v82
                        }
                    case 4294967295:
                        v83 := GenS12{F0: 640}
                        _ = v83
                    default:
                        v3 = v3
                    }
                    _ = fi81
                }
                _ = fi78
            }
            fmt.Println(v13)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v74
}
