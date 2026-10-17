package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 uint64; F1 uint; F2 string }
type GenS4 struct { F0 uint; F1 uint64 }
type GenS5 struct { F0 int8; F1 uint }
type GenS6 struct { F0 int; F1 int32; F2 int8 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint8; F1 uint32 }

func main() {
    v1 := []string{("\n\t" + (("foo bar" + "foo bar") + "hello")), "hello", "a", ""}
    v2 := make(map[int]string)
    v3 := make(map[uint32]string)
    var v4 bool = (!false)
    v5 := []string{"foo bar", (("a" + ((("foo bar" + "foo bar") + "a") + (("" + ("\n\t" + (((((("foo bar" + "") + "\n\t") + "foo bar") + "a") + "") + ("" + "a")))) + ""))) + ("hello" + "\n\t"))}
    v6 := make(map[string]bool)
    v4 = v4
    for fi7 := 0; fi7 < 9; fi7++ {
        if len(v1) > 0 {
            v1[0] = "foo bar"
        }
        _ = fi7
    }
    var v8 float32 = (1e10 - float32(0.0))
    if (!false) {
        var v9 int8 = 0
        _ = v9
    } else {
        var v10 int32 = -77
        if false {
            _ = len(v1)
            delete(v6, "a")
        } else {
            switch "foo bar" {
            case "foo bar":
                for fi11 := 0; fi11 < 20; fi11++ {
                    _ = len(v5)
                    var v12 string = "foo bar"
                    _ = fi11
                    _ = v12
                }
            default:
                for fi13 := 0; fi13 < 10; fi13++ {
                    var v14 int32 = v10
                    v14 -= (v14 & (v10 ^ 2147483647))
                    v10 = v10
                    _ = fi13
                    _ = v14
                }
            }
            var v15 int = 1
            _ = v15
        }
        _ = v10
    }
    var v16 string = "a"
    for fi17 := 0; fi17 < 9; fi17++ {
        if (-75 < -5) {
            if false {
                if v4 {
                    v8 += v8
                } else {
                    for fi18 := 0; fi18 < 4; fi18++ {
                        v19 := v3[2147483648]
                        _ = fi18
                        _ = v19
                    }
                    var v20 float32 = v8
                    _ = v20
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            } else {
                v16 = v16
                for fi21 := 0; fi21 < 7; fi21++ {
                    v16 = "\n\t"
                    var v22 uint64 = 18446744073709551615
                    _ = fi21
                    _ = v22
                }
            }
            var v23 uint = 1
            _ = v23
        } else {
            fmt.Println(v16)
            for fi24 := 0; fi24 < 13; fi24++ {
                var v25 uint16 = 56
                v25 = 32767
                fmt.Println(v8, v8)
                _ = fi24
                _ = v25
            }
            if v4 {
                for fi26 := 0; fi26 < 9; fi26++ {
                    var v27 int16 = -1
                    var v28 uint32 = 312
                    v16 = (v16 + "foo bar")
                    v29 := GenS1{F0: -1, F1: "foo bar"}
                    var v30 uint16 = 349
                    fmt.Println(v8, v4, v8)
                    var v31 int16 = (v27 % (((v27 - 1)) | 1))
                    _ = len(v2)
                    _ = fi26
                    _ = v27
                    _ = v28
                    _ = v29
                    _ = v30
                    _ = v31
                }
                v32 := GenS2{F0: 0, F1: -914}
                _ = v32
            }
        }
        _ = fi17
    }
    switch 32767 {
    case 1:
        if true {
            switch 32768 {
            case -1:
                if v4 {
                    for fi33 := 0; fi33 < 20; fi33++ {
                        if v4 {
                            fmt.Println(v8, v8)
                        } else {
                            v8 += (-2.25 - (v8 - (-2.25 + (v8 - ((float32((3.14159 * 0.1)) - 1.5) - float32(1e10))))))
                        }
                        if true {
                            if v4 {
                                v4 = true
                            } else {
                                v8 -= v8
                                var v34 string = ((v16 + v16) + v16)
                                v35 := GenS3{F0: 42, F1: 0, F2: v34}
                                _ = v34
                                _ = v35
                            }
                            switch (fi33 - (^(fi33))) {
                            case -901:
                                fmt.Println(v16, v4, fi33)
                            case 32767:
                                for fi36 := 0; fi36 < 15; fi36++ {
                                    v8 = (((3.14159 + 0.1) + 1.5) + v8)
                                    v1 = append(v1, v16)
                                    _ = fi36
                                }
                            default:
                                v4 = false
                            }
                            if true {
                                v37 := GenS2{F0: 42, F1: 535}
                                _ = v37
                            } else {
                                v38 := GenS3{F0: 305, F1: 822, F2: (v16 + "\n\t")}
                                _ = v38
                            }
                        } else {
                            v16 = (("" + v16) + (v16 + (("hello" + (v16 + v16)) + v16)))
                            if v4 {
                                var v39 float64 = (3.14159 + 1.5)
                                var v40 float64 = v39
                                _ = v39
                                _ = v40
                            }
                        }
                        var v41 int8 = -43
                        _ = fi33
                        _ = v41
                    }
                    if true {
                        var v42 int = 9223372036854775807
                        fmt.Println(v16, v8)
                        _ = v42
                    }
                    fmt.Println(v16, v16)
                } else {
                    if false {
                        v4 = v4
                        v5 = append(v5, (v16 + ((v16 + v16) + "hello")))
                        for fi43 := 0; fi43 < 6; fi43++ {
                            switch ("foo bar" + "hello") {
                            case "":
                                if v4 {
                                    v4 = (!false)
                                    var v44 uint = 2
                                    _ = v44
                                } else {
                                    if v4 {
                                        for fi45 := 0; fi45 < 1; fi45++ {
                                            if (!(!v4)) {
                                                v4 = v4
                                                v4 = (697 >= -32768)
                                            } else {
                                                fmt.Println(v8, v4)
                                                fmt.Println(fi43, fi43, fi45)
                                                fmt.Println(v4, fi43, v16)
                                                if ((0 != 1) || (!(!true))) {
                                                    v16 = "a"
                                                    v4 = v4
                                                    fmt.Println(fi43, v16)
                                                    fmt.Println(v16)
                                                }
                                                v4 = (!(977 == -1))
                                                if (!(v4 && v4)) {
                                                    v46 := GenS3{F0: 0, F1: 18446744073709551615, F2: v16}
                                                    _ = v46
                                                } else {
                                                    fmt.Println(v8, v8)
                                                    switch v16 {
                                                    case "hello":
                                                        fmt.Println(fi45, v4)
                                                        var v47 uint16 = 255
                                                        _ = v47
                                                    case "a":
                                                        v48 := GenS2{F0: 180, F1: 42}
                                                        _ = v48
                                                    case "foo bar":
                                                        fmt.Println(fi43, v16, v16)
                                                    default:
                                                        v1 = append(v1, (v16 + v16))
                                                        v49 := GenS4{F0: 65535, F1: 18446744073709551615}
                                                        _ = v49
                                                    }
                                                }
                                                v50 := GenS4{F0: 2, F1: 2}
                                                v16 = v16
                                                fmt.Println(v16)
                                                _ = v50
                                            }
                                            v8 = (v8 / 0.1)
                                            _ = fi45
                                        }
                                        fmt.Println(v8, fi43, v4)
                                    } else {
                                        v5 = append(v5, (v16 + "hello"))
                                        if ((false && (v4 && (v16 <= (v16 + v16)))) || v4) {
                                            v8 = v8
                                        } else {
                                            var v51 int32 = 8
                                            v8 += (float32(0.0) + -2.25)
                                            v4 = false
                                            _ = v51
                                        }
                                    }
                                }
                            default:
                                var v52 float32 = float32(((3.14159 + float64(v8)) - 1e10))
                                _ = v52
                            }
                            _ = fi43
                        }
                    }
                    if (1 > 830) {
                        v53 := GenS5{F0: 42, F1: 255}
                        _ = v53
                    } else {
                        _ = len(v6)
                    }
                    if (v4 || true) {
                        for fi54 := 0; fi54 < 15; fi54++ {
                            var v55 uint64 = 42
                            switch (fi54 >> ((4294967295) & 63)) {
                            case 1:
                                if v4 {
                                    var v56 string = v16
                                    var v57 bool = (764 >= 42)
                                    _ = v56
                                    _ = v57
                                }
                            case -914:
                                fmt.Println(v8, fi54)
                            case 4294967295:
                                v4 = (0 >= 165)
                            default:
                                var v58 uint = uint(v55)
                                _ = v58
                            }
                            _ = fi54
                            _ = v55
                        }
                    }
                }
            }
            v16 = v16
        } else {
            for fi59 := 0; fi59 < 14; fi59++ {
                v4 = (-2147483648 >= -2147483648)
                v16 = "\n\t"
                _ = fi59
            }
            if (!(1 != -357)) {
                if false {
                    v60 := GenS4{F0: 0, F1: 42}
                    var v61 string = ""
                    var v62 uint32 = 4294967295
                    _ = v60
                    _ = v61
                    _ = v62
                }
                var v63 uint32 = 166
                _ = v63
            }
            var v64 int32 = 644
            _ = v64
        }
    case -405:
        for fi65 := 0; fi65 < 18; fi65++ {
            for fi66 := 0; fi66 < 19; fi66++ {
                v16 = ("\n\t" + ("" + "foo bar"))
                v67 := GenS6{F0: (fi65 % ((fi65) | 1)), F1: 2147483647, F2: -19}
                if v4 {
                    var v68 int8 = 100
                    if v4 {
                        if ("a" != (((v16 + (v16 + "a")) + ((v16 + ((v16 + ("\n\t" + v16)) + "")) + (((v16 + (("hello" + (v16 + v16)) + v16)) + ((v16 + v16) + v16)) + v16))) + "")) {
                            var v69 uint8 = 1
                            _ = v69
                        } else {
                            var v70 int = (fi66 + 0)
                            _ = v70
                        }
                    }
                    v4 = ((false || v4) || ((798 <= -1) || v4))
                    _ = v68
                } else {
                    v8 = (v8 / 0.1)
                    v4 = true
                    fmt.Println(v8)
                }
                _ = fi66
                _ = v67
            }
            for fi71 := 0; fi71 < 2; fi71++ {
                if (16 != 4294967295) {
                    fmt.Println(v16)
                    switch fi65 {
                    case -365:
                        v8 += v8
                    default:
                        v4 = (v4 || true)
                    }
                    v72 := GenS1{F0: 120, F1: v16}
                    _ = v72
                }
                if len(v5) > 0 {
                    v5[0] = v16
                }
                _ = fi71
            }
            var v73 uint64 = 399
            _ = fi65
            _ = v73
        }
    default:
        v74 := GenS7{F0: (573 <= 2147483647)}
        _ = v74
    }
    for fi75 := 0; fi75 < 5; fi75++ {
        v4 = true
        switch 636 {
        case 567:
            if true {
                v8 += 3.14159
            } else {
                if (583 >= 18446744073709551615) {
                    if (-1 > (fi75 >> ((65535) & 63))) {
                        if true {
                            v76 := GenS8{F0: 2, F1: 2147483648}
                            for fi77 := 0; fi77 < 7; fi77++ {
                                if true {
                                    switch (v16 + v16) {
                                    case "foo bar":
                                        if true {
                                            v16 = (v16 + v16)
                                            var v78 bool = false
                                            if (0 == -9223372036854775808) {
                                                v4 = v78
                                            }
                                            _ = v78
                                        }
                                    }
                                    v16 = v16
                                    fmt.Println(fi75, v16, fi75)
                                }
                                switch fi77 {
                                case -2147483648:
                                    v79 := v2[fi77]
                                    _ = v79
                                case 4294967295:
                                    v4 = (!v4)
                                    switch v16 {
                                    case "\n\t":
                                        fmt.Println(v4)
                                    case "hello":
                                        fmt.Println(v8)
                                        v4 = v4
                                        v80 := GenS8{F0: 255, F1: 65535}
                                        _ = v80
                                    case "foo bar":
                                        var v81 uint16 = 65535
                                        _ = v81
                                    default:
                                        switch v16 {
                                        case "a":
                                            if (177 >= 743) {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                if (775 < 4294967295) {
                                                    if v4 {
                                                        if true {
                                                            fmt.Println(v16, v16, fi77)
                                                            var v82 int64 = 0
                                                            _ = v82
                                                        }
                                                        v8 = (v8 + v8)
                                                    }
                                                }
                                            }
                                        default:
                                            var v83 float64 = 0.0
                                            delete(v6, v16)
                                            _ = v83
                                        }
                                    }
                                case -957:
                                    var v84 uint = 42
                                    _ = v84
                                default:
                                    if (2147483648 <= 2) {
                                        var v85 uint8 = 124
                                        var v86 float64 = 3.14159
                                        v86 -= ((v86 / 1.0) * v86)
                                        _ = v85
                                        _ = v86
                                    }
                                }
                                _ = fi77
                            }
                            v8 -= float32(float64(v8))
                            _ = v76
                        } else {
                            for fi87 := 0; fi87 < 2; fi87++ {
                                v88 := GenS3{F0: 0, F1: 323, F2: "a"}
                                if false {
                                    var v89 float32 = 1e10
                                    _ = v89
                                }
                                var v90 uint = 918
                                _ = fi87
                                _ = v88
                                _ = v90
                            }
                        }
                        switch 128 {
                        case 9223372036854775807:
                            v16 = (("hello" + v16) + v16)
                        default:
                            v8 = v8
                        }
                        var v91 float32 = 3.14159
                        _ = v91
                    }
                    for fi92 := 0; fi92 < 16; fi92++ {
                        v93 := v2[(-(fi75))]
                        fmt.Println(v8)
                        _ = fi92
                        _ = v93
                    }
                }
                switch fi75 {
                case -575:
                    fmt.Println(fi75, v16, v16)
                }
                var v94 int = (fi75 ^ (fi75 & fi75))
                _ = v94
            }
        }
        _ = fi75
    }
    fmt.Println(v8)
    v95 := v6[v16]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v95
}
