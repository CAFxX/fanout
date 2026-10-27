package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32; F1 string }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint16; F1 int8; F2 int }
type GenS8 struct { F0 string; F1 uint64; F2 bool }
func genFn0(p0 uint8) uint16 {
    v4 := GenS1{F0: -1}
    p0 -= p0
    var v5 bool = (!true)
    v6 := []int{}
    for fi7 := 0; fi7 < 10; fi7++ {
        if (763 < 65535) {
            v6 = append(v6, 42)
        }
        _ = fi7
    }
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 uint64, p1 int, p2 int8) uint64 {
    fmt.Println(p0, p2)
    fmt.Println(p1, p0)
    var v8 uint32 = 4294967295
    p2 += (p2 >> ((525) & 63))
    var v9 uint64 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return v9
}

func main() {
    var v1 uint32 = 1
    v2 := make(map[string]uint32)
    v3 := []string{(("\n\t" + "\n\t") + "a"), "", (("foo bar" + ("foo bar" + "\n\t")) + ("\n\t" + "foo bar")), ("" + ("\n\t" + ("a" + "foo bar")))}
    fmt.Println(v1)
    switch ("a" + "hello") {
    case "hello":
        delete(v2, "")
    case "a":
        for fi10 := 0; fi10 < 1; fi10++ {
            v11 := GenS1{F0: 1}
            v1 = (v1 ^ (v1 / ((65535) | 1)))
            if false {
                fmt.Println(v1, v1)
                v1 += v1
            }
            _ = fi10
            _ = v11
        }
    default:
        v12 := GenS2{F0: uint64(v1), F1: 67}
        _ = v12
    }
    v1 = 2147483647
    var v13 int64 = 0
    v13 = 0
    for fi14 := 0; fi14 < 16; fi14++ {
        if false {
            v1 = v1
        }
        fmt.Println(v1)
        _ = fi14
    }
    var v15 int = 128
    delete(v2, ("\n\t" + ""))
    for fi16 := 0; fi16 < 7; fi16++ {
        if true {
            v17 := GenS3{F0: 42}
            if false {
                for fi18 := 0; fi18 < 5; fi18++ {
                    var v19 bool = true
                    _ = fi18
                    _ = v19
                }
            }
            if (v13 < (v13 * (^(v13)))) {
                fmt.Println(fi16)
                for fi20 := 0; fi20 < 2; fi20++ {
                    v21 := GenS1{F0: (v13 - (v13 >> ((0) & 63)))}
                    fmt.Println(v15, fi20)
                    var v22 uint = 690
                    var v23 int = (fi20 * (^(fi16)))
                    var v24 int = (-(v23))
                    _ = fi20
                    _ = v21
                    _ = v22
                    _ = v23
                    _ = v24
                }
                v15 += v15
            } else {
                v25 := GenS2{F0: 370, F1: 127}
                _ = v25
            }
            _ = v17
        }
        v26 := GenS3{F0: 2}
        _ = fi16
        _ = v26
    }
    v27 := GenS3{F0: 0}
    for fi28 := 0; fi28 < 16; fi28++ {
        if false {
            v1 = v1
            fmt.Println(v13)
        } else {
            v15 = (fi28 >> ((294) & 63))
            v29 := GenS1{F0: (v13 & -1)}
            switch (fi28 * 32768) {
            case 127:
                var v30 uint8 = 255
                _ = v30
            case 450:
                var v31 uint16 = 465
                _ = v31
            case 128:
                v15 += 2147483647
            default:
                if true {
                    v15 += fi28
                    v32 := GenS3{F0: 42}
                    v3 = append(v3, (("hello" + ("hello" + "foo bar")) + ""))
                    _ = v32
                }
            }
            _ = v29
        }
        for fi33 := 0; fi33 < 19; fi33++ {
            var v34 int = fi33
            v3 = append(v3, "foo bar")
            v2["a"] = v1
            _ = fi33
            _ = v34
        }
        for fi35 := 0; fi35 < 6; fi35++ {
            var v36 int16 = 118
            v37 := GenS1{F0: (v13 - v13)}
            if (((v13 > (v13 + v13)) && true) && false) {
                switch fi35 {
                case 32767:
                    if false {
                        if (true && true) {
                            var v38 uint64 = 1
                            if false {
                                if ("\n\t" >= ("" + "\n\t")) {
                                    v13 -= (-(v13))
                                    var v39 uint = 182
                                    v40 := v2[("a" + "a")]
                                    _ = v39
                                    _ = v40
                                } else {
                                    if (1 == 2147483647) {
                                        var v41 int8 = 1
                                        v42 := GenS3{F0: 2}
                                        _ = v41
                                        _ = v42
                                    }
                                    v36 = v36
                                }
                                _ = len(v3)
                                if true {
                                    v3 = append(v3, "\n\t")
                                    var v43 int16 = v36
                                    _ = v43
                                }
                            }
                            _ = v38
                        } else {
                            v15 += 2147483647
                        }
                        v44 := GenS4{F0: (v1 + v1), F1: ""}
                        v45 := GenS5{F0: (-1 != 0)}
                        _ = v44
                        _ = v45
                    } else {
                        if false {
                            var v46 uint64 = 725
                            _ = v46
                        } else {
                            _ = len(v2)
                            v47 := GenS5{F0: (false || true)}
                            v48 := GenS1{F0: v13}
                            fmt.Println(v15)
                            _ = v47
                            _ = v48
                        }
                        var v49 uint64 = 558
                        var v50 int64 = v13
                        fmt.Println(v1, v13, fi35)
                        _ = v49
                        _ = v50
                    }
                case 382:
                    fmt.Println(v15, v13, fi35)
                case -2147483648:
                    fmt.Println(v36)
                default:
                    v15 -= fi28
                }
            }
            _ = fi35
            _ = v36
            _ = v37
        }
        _ = fi28
    }
    v13 = v13
    var v51 uint8 = 255
    v13 -= (^(v13))
    if true {
        var v52 float64 = float64(1.5)
        v53 := GenS6{F0: (("a" + ("a" + "hello")) + "\n\t")}
        var v54 int32 = 61
        _ = v52
        _ = v53
        _ = v54
    }
    v55 := GenS3{F0: 2}
    if false {
        v15 -= -2147483648
        switch (v15 ^ v15) {
        case 128:
            for fi56 := 0; fi56 < 3; fi56++ {
                _ = len(v2)
                v57 := GenS2{F0: 975, F1: 255}
                var v58 int8 = 1
                _ = fi56
                _ = v57
                _ = v58
            }
        case 127:
            var v59 uint8 = v51
            _ = v59
        default:
            _ = len(v2)
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    switch v15 {
    case 2147483647:
        switch ((("foo bar" + ("a" + ("hello" + "hello"))) + ((("foo bar" + "a") + "hello") + "foo bar")) + ("a" + "a")) {
        case "":
            v3 = append(v3, "")
        case "\n\t":
            var v60 string = "\n\t"
            _ = v60
        case "a":
            if (42 != 907) {
                v13 += (v13 / ((1) | 1))
            }
        default:
            switch "" {
            case "":
                if len(v3) > 0 {
                    v3[0] = "a"
                }
            case "\n\t":
                if (!(100 >= -1)) {
                    var v61 uint32 = uint32(v51)
                    var v62 float32 = 3.14159
                    switch "" {
                    case "":
                        for fi63 := 0; fi63 < 9; fi63++ {
                            if ((false || ((false || false) && false)) || false) {
                                v3 = append(v3, ("foo bar" + (((("foo bar" + "\n\t") + "") + "foo bar") + "a")))
                            } else {
                                fmt.Println(v1)
                                var v64 uint8 = v51
                                v1 -= (v61 << ((836) & 63))
                                _ = v64
                            }
                            v62 += (0.1 - v62)
                            _ = fi63
                        }
                    case "a":
                        v2[("hello" + "hello")] = 2147483647
                    case "\n\t":
                        var v65 uint32 = uint32(v51)
                        _ = v65
                    }
                    _ = v61
                    _ = v62
                } else {
                    v13 = v13
                    switch "a" {
                    case "a":
                        var v66 int64 = v13
                        _ = v66
                    case "\n\t":
                        v67 := GenS7{F0: 255, F1: -22, F2: v15}
                        _ = v67
                    case "":
                        if false {
                            for fi68 := 0; fi68 < 9; fi68++ {
                                for fi69 := 0; fi69 < 1; fi69++ {
                                    if ((!(18446744073709551615 > 1)) || false) {
                                        var v70 float64 = 1e10
                                        _ = v70
                                    }
                                    v71 := GenS6{F0: "\n\t"}
                                    _ = fi69
                                    _ = v71
                                }
                                var v72 uint16 = genFn0(v51)
                                _ = fi68
                                _ = v72
                            }
                        }
                    default:
                        if (-32768 == -32768) {
                            var v73 float32 = 0.1
                            if true {
                                v74 := GenS8{F0: "a", F1: genFn1(1, v15, 1), F2: (255 != 104)}
                                v13 = (-(v13))
                                _ = v74
                            } else {
                                v73 += (float32(float64(v73)) * 1e10)
                                v75 := GenS2{F0: 939, F1: (v51 << ((0) & 63))}
                                fmt.Println(v15, v1)
                                _ = v75
                            }
                            _ = v73
                        } else {
                            var v76 string = ""
                            v15 += (-(v15))
                            v15 -= -1
                            _ = v76
                        }
                    }
                    v2["foo bar"] = (v1 * (v1 >> ((0) & 63)))
                }
            default:
                var v77 uint64 = 4294967295
                _ = v77
            }
        }
    case -608:
        for fi78 := 0; fi78 < 10; fi78++ {
            if ("foo bar" >= "a") {
                fmt.Println(fi78)
            } else {
                var v79 string = "\n\t"
                _ = v79
            }
            if (-651 >= 0) {
                var v80 uint64 = 18446744073709551615
                _ = v80
            }
            _ = fi78
        }
    default:
        fmt.Println(v13, v13)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v13
    _ = v15
    _ = v2
    _ = v27
    _ = v3
    _ = v51
    _ = v55
}
