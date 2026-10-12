package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 uint64; F1 bool }
type GenS4 struct { F0 uint8; F1 int32 }
type GenS5 struct { F0 uint64; F1 uint }
type GenS6 struct { F0 uint64; F1 string }
type GenS7 struct { F0 bool; F1 uint8 }
type GenS8 struct { F0 int32; F1 string; F2 int8 }
type GenS9 struct { F0 int32; F1 uint32 }
func genFn0(p0 bool, p1 uint8, p2 bool) int32 {
    p2 = false
    for fi6 := 0; fi6 < 1; fi6++ {
        fmt.Println(p0, p0)
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 uint16 = 32767
    v2 := make(map[uint32]int)
    var v3 int64 = -9223372036854775808
    var v4 int32 = -2147483648
    v5 := []string{"", "\n\t", "a"}
    var v7 uint16 = (v1 & v1)
    fmt.Println(v7, v3)
    v1 -= v7
    for fi8 := 0; fi8 < 15; fi8++ {
        for fi9 := 0; fi9 < 16; fi9++ {
            switch ("a" + "") {
            case "hello":
                var v10 int64 = (v3 | 9223372036854775807)
                _ = v10
            case "foo bar":
                v11 := GenS1{F0: 42}
                _ = v11
            default:
                fmt.Println(v1, fi8, v7)
            }
            if ((v3 & 9223372036854775807) == (v3 | 0)) {
                switch "" {
                case "foo bar":
                    var v12 uint8 = 42
                    _ = v12
                case "":
                    var v13 int64 = (^(v3))
                    _ = v13
                case "\n\t":
                    if false {
                        v14 := GenS1{F0: 2}
                        _ = v14
                    } else {
                        if false {
                            if false {
                                if false {
                                    switch "" {
                                    case "a":
                                        v15 := GenS2{F0: false, F1: 1}
                                        _ = v15
                                    case "":
                                        var v16 uint32 = 1
                                        _ = v16
                                    case "foo bar":
                                        switch "\n\t" {
                                        case "\n\t":
                                            v17 := GenS3{F0: 42, F1: false}
                                            _ = v17
                                        }
                                    }
                                    v18 := GenS2{F0: true, F1: 4294967295}
                                    v19 := GenS2{F0: (((3.14159 > (0.0 * float64(0.0))) && true) || false), F1: 899}
                                    _ = v18
                                    _ = v19
                                } else {
                                    v3 = (v3 >> ((245) & 63))
                                    switch "\n\t" {
                                    case "":
                                        v1 = v7
                                    }
                                    v20 := GenS4{F0: 54, F1: 1}
                                    if false {
                                        fmt.Println(fi9, v1)
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        v21 := GenS4{F0: 204, F1: (-(v4))}
                                        _ = v21
                                    }
                                    _ = v20
                                }
                                if false {
                                    v22 := GenS1{F0: 198}
                                    _ = v22
                                } else {
                                    v23 := v2[563]
                                    _ = v23
                                }
                            }
                            switch (fi9 >> ((4294967295) & 63)) {
                            case -484:
                                if false {
                                    var v24 int32 = 791
                                    _ = v24
                                } else {
                                    _ = len(v2)
                                    v1 += (^(v1))
                                }
                            case 128:
                                var v25 uint8 = 1
                                _ = v25
                            case 127:
                                var v26 uint16 = (^(v1))
                                _ = v26
                            }
                        }
                    }
                default:
                    fmt.Println(v3, v1, v1)
                }
                v7 += v7
                v3 = (v3 * v3)
            }
            v5 = append(v5, "\n\t")
            if false {
                v5 = append(v5, "foo bar")
                v7 = (v7 % (((v1 >> ((1) & 63))) | 1))
            } else {
                v1 = v1
            }
            var v27 int = int(v1)
            _ = fi9
            _ = v27
        }
        v28 := GenS5{F0: 0, F1: 254}
        v4 -= -1
        _ = fi8
        _ = v28
    }
    if len(v5) > 0 {
        v5[0] = ""
    }
    v2[1] = int(v4)
    if (4294967295 < 0) {
        switch "\n\t" {
        case "":
            _ = len(v2)
        case "foo bar":
            v4 -= v4
        case "a":
            if false {
                for fi29 := 0; fi29 < 5; fi29++ {
                    if (65535 > 1) {
                        if false {
                            if len(v5) > 0 {
                                v5[0] = ("a" + "foo bar")
                            }
                        }
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        v7 = v1
                    }
                    _ = fi29
                }
            } else {
                v30 := GenS5{F0: 781, F1: 255}
                var v31 uint8 = 184
                _ = v30
                _ = v31
            }
        }
        var v32 uint64 = 4294967295
        _ = v32
    }
    if false {
        var v33 bool = ((^(v4)) < (v4 + (v4 + (v4 | (v4 - int32(v7))))))
        _ = v33
    }
    for fi34 := 0; fi34 < 7; fi34++ {
        switch (-(fi34)) {
        case 949:
            var v35 bool = (0 < 0)
            _ = v35
        }
        _ = fi34
    }
    var v36 string = (("\n\t" + "a") + "")
    fmt.Println(v36, v4, v36)
    v1 = 32767
    if (!(1 != 4294967295)) {
        if true {
            v37 := GenS1{F0: 1}
            _ = v37
        }
        var v38 int16 = 1
        var v39 int = 4294967295
        _ = v38
        _ = v39
    } else {
        var v40 uint64 = 80
        if true {
            switch v36 {
            case "":
                for fi41 := 0; fi41 < 13; fi41++ {
                    for fi42 := 0; fi42 < 9; fi42++ {
                        if ((v3 % ((488) | 1)) >= v3) {
                            var v43 int16 = -205
                            if (!true) {
                                switch 9223372036854775807 {
                                case 4294967295:
                                    switch (fi42 - fi42) {
                                    case -163:
                                        v44 := GenS6{F0: v40, F1: "a"}
                                        _ = v44
                                    }
                                default:
                                    if true {
                                        switch v36 {
                                        case "hello":
                                            v1 = (^(v1))
                                        }
                                    } else {
                                        v45 := GenS7{F0: true, F1: 210}
                                        if ((v1 & 42) <= (v1 * (v7 + (v7 + 255)))) {
                                            v4 -= (v4 + v4)
                                        }
                                        v43 = v43
                                        v4 += (v4 << ((42) & 63))
                                        _ = v45
                                    }
                                }
                            }
                            var v46 float64 = -2.25
                            _ = len(v5)
                            _ = v43
                            _ = v46
                        }
                        v47 := v2[208]
                        v40 = v40
                        _ = fi42
                        _ = v47
                    }
                    _ = fi41
                }
            case "hello":
                v48 := GenS8{F0: (^(v4)), F1: v36, F2: -1}
                _ = v48
            default:
                var v49 float64 = 3.14159
                _ = v49
            }
        }
        _ = v40
    }
    var v50 int16 = 0
    if false {
        if len(v5) > 0 {
            v5[0] = "hello"
        }
        var v51 uint16 = 0
        _ = len(v2)
        _ = v51
    } else {
        for fi52 := 0; fi52 < 13; fi52++ {
            if (!true) {
                v36 = v36
                if len(v5) > 0 {
                    v5[0] = (v36 + v36)
                }
            } else {
                switch (v36 + ((v36 + v36) + v36)) {
                case "":
                    v50 = 360
                case "\n\t":
                    var v53 int32 = (-(v4))
                    _ = v53
                default:
                    v5 = append(v5, v36)
                }
                v2[249] = 1
                for fi54 := 0; fi54 < 15; fi54++ {
                    if (false && (-54 > 108)) {
                        v3 += v3
                        var v55 float32 = (0.0 + 3.14159)
                        v50 -= -312
                        switch "foo bar" {
                        case "":
                            v56 := GenS6{F0: 86, F1: (v36 + v36)}
                            _ = v56
                        default:
                            switch v36 {
                            case "":
                                v57 := GenS9{F0: (-(v4)), F1: 1}
                                _ = v57
                            }
                        }
                        v58 := GenS9{F0: (v4 >> ((854) & 63)), F1: 1}
                        _ = v55
                        _ = v58
                    }
                    _ = fi54
                }
            }
            _ = fi52
        }
        if (v50 != 971) {
            if false {
                v36 = (("foo bar" + v36) + (v36 + ((v36 + "a") + v36)))
            }
            _ = len(v5)
            var v59 string = ""
            _ = v59
        }
        fmt.Println(v7, v4, v36)
    }
    for fi60 := 0; fi60 < 20; fi60++ {
        _ = len(v5)
        var v61 string = "hello"
        v50 = v50
        _ = fi60
        _ = v61
    }
    var v62 uint8 = 2
    for fi63 := 0; fi63 < 15; fi63++ {
        switch "\n\t" {
        case "\n\t":
            v3 += v3
        default:
            fmt.Println(v7, v36)
        }
        _ = fi63
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("v62:", v62)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v50
    _ = v62
    _ = v7
}
