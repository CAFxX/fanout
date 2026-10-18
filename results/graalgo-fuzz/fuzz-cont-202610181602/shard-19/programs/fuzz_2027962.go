package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 uint16; F1 int8; F2 int64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint16; F1 uint32 }
type GenS6 struct { F0 int16; F1 bool; F2 int32 }
type GenS7 struct { F0 int; F1 int64 }
type GenS8 struct { F0 uint8; F1 int32 }
type GenS9 struct { F0 int32; F1 int64 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 uint64; F1 int8 }

func main() {
    var v1 int = 2147483647
    var v2 float64 = 1.5
    var v3 int64 = 413
    var v4 int64 = v3
    v5 := []string{"hello", "hello", ("hello" + "foo bar")}
    _ = len(v5)
    if len(v5) > 0 {
        v5[0] = ("hello" + "foo bar")
    }
    var v6 bool = true
    for fi7 := 0; fi7 < 11; fi7++ {
        for fi8 := 0; fi8 < 4; fi8++ {
            if (0 > 1000) {
                if ((v1 * 987) <= (fi7 & 127)) {
                    fmt.Println(fi8, fi7, fi7)
                    fmt.Println(fi8, fi8)
                    fmt.Println(fi7, fi8, v4)
                }
                v2 = ((((0.0 / v2) - ((float64(0.0) + (v2 * 0.0)) - v2)) * (v2 + (v2 * v2))) + -2.25)
                if (!((18446744073709551615 < 2) || (false && v6))) {
                    var v9 float64 = 0.1
                    _ = v9
                } else {
                    v10 := GenS1{F0: 1, F1: 4294967295}
                    if true {
                        var v11 uint64 = 18446744073709551615
                        v12 := GenS2{F0: "", F1: 99}
                        v13 := GenS1{F0: 0, F1: 774}
                        _ = v11
                        _ = v12
                        _ = v13
                    }
                    fmt.Println(fi8, v2, fi7)
                    _ = v10
                }
                if v6 {
                    if ("foo bar" > "") {
                        v14 := GenS1{F0: 0, F1: 18446744073709551615}
                        if ("a" != "hello") {
                            var v15 float32 = 1e10
                            _ = v15
                        }
                        v1 -= v1
                        if (4294967295 > 789) {
                            switch -2147483648 {
                            case 0:
                                if false {
                                    fmt.Println(v1)
                                } else {
                                    if (142 >= 0) {
                                        if v6 {
                                            v4 = (v3 | v3)
                                            var v16 string = "a"
                                            v17 := GenS3{F0: 65535, F1: 127, F2: v3}
                                            _ = v16
                                            _ = v17
                                        }
                                        switch ("\n\t" + (("a" + "foo bar") + "\n\t")) {
                                        case "\n\t":
                                            v18 := GenS3{F0: 1, F1: -128, F2: (-(v4))}
                                            fmt.Println(v4, fi8)
                                            var v19 float64 = v2
                                            _ = v18
                                            _ = v19
                                        case "a":
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                        case "foo bar":
                                            v6 = (v1 <= v1)
                                        default:
                                            v6 = (!v6)
                                            v4 = v3
                                        }
                                    } else {
                                        v20 := GenS4{F0: v6}
                                        v1 -= (fi8 * 128)
                                        v1 = 32767
                                        _ = v20
                                    }
                                    if (!v6) {
                                        v21 := GenS5{F0: 848, F1: 42}
                                        if (v6 && v6) {
                                            var v22 bool = (!v6)
                                            switch fi8 {
                                            case 513:
                                                switch (fi7 / ((fi7) | 1)) {
                                                case 42:
                                                    var v23 uint64 = 4294967295
                                                    _ = v23
                                                case 0:
                                                    if v22 {
                                                        if len(v5) > 0 {
                                                            v5[0] = ("foo bar" + (((((("foo bar" + ("a" + (("hello" + (("" + "a") + "foo bar")) + ""))) + "\n\t") + ("hello" + "")) + "\n\t") + ("a" + "a")) + (((("" + ("a" + ("foo bar" + ("foo bar" + ("a" + "a"))))) + "a") + ("\n\t" + "foo bar")) + ("hello" + "foo bar"))))
                                                        }
                                                        v22 = (18446744073709551615 > 18446744073709551615)
                                                        if ((0.0 + v2) <= (v2 / 1.0)) {
                                                            v24 := GenS1{F0: -1, F1: 0}
                                                            if len(v5) > 0 {
                                                                v5[0] = ""
                                                            }
                                                            _ = v24
                                                        }
                                                    }
                                                default:
                                                    if (-440 < 1) {
                                                        v4 = (^(v3))
                                                    }
                                                }
                                            case -720:
                                                v6 = false
                                            case 925:
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                            default:
                                                if len(v5) > 0 {
                                                    v5[0] = ""
                                                }
                                            }
                                            var v25 uint64 = 317
                                            var v26 float32 = 1e10
                                            fmt.Println(v1, fi8)
                                            _ = v22
                                            _ = v25
                                            _ = v26
                                        } else {
                                            if v6 {
                                                v2 -= 1e10
                                                var v27 int32 = 1
                                                var v28 uint32 = 254
                                                _ = v27
                                                _ = v28
                                            }
                                            if false {
                                                v3 += (v4 + v4)
                                                var v29 uint16 = 2
                                                var v30 uint16 = 0
                                                _ = v29
                                                _ = v30
                                            }
                                        }
                                        _ = v21
                                    }
                                }
                            case 9223372036854775807:
                                v31 := GenS6{F0: -709, F1: v6, F2: 151}
                                _ = v31
                            case 701:
                                if v6 {
                                    fmt.Println(fi8, fi8, v6)
                                }
                            default:
                                fmt.Println(fi7, v3)
                            }
                        } else {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            var v32 uint64 = 18446744073709551615
                            _ = v32
                        }
                        v4 += 42
                        v33 := GenS7{F0: -243, F1: -1}
                        var v34 uint64 = 1
                        fmt.Println(v1, fi8)
                        fmt.Println(v3, v6)
                        _ = v14
                        _ = v33
                        _ = v34
                    }
                    var v35 float64 = 1e10
                    v5 = append(v5, (((("foo bar" + "") + ("\n\t" + "")) + (("hello" + "hello") + ("hello" + ""))) + "a"))
                    if ((192 > 127) || true) {
                        var v36 uint8 = 255
                        v37 := GenS3{F0: 53, F1: 20, F2: (^(v3))}
                        _ = v36
                        _ = v37
                    } else {
                        var v38 uint64 = 594
                        switch 2147483647 {
                        case 1:
                            v38 = (v38 & 1)
                        case -469:
                            if len(v5) > 0 {
                                v5[0] = (("a" + "foo bar") + "foo bar")
                            }
                        case 706:
                            if len(v5) > 0 {
                                v5[0] = "hello"
                            }
                            v6 = (!false)
                            v2 = v2
                        }
                        fmt.Println(v3, v35)
                        _ = v38
                    }
                    _ = v35
                }
            } else {
                if true {
                    fmt.Println(v1, v6, fi7)
                    var v39 float64 = v2
                    if (v6 && v6) {
                        v40 := GenS7{F0: fi8, F1: 9223372036854775807}
                        v5 = append(v5, (("\n\t" + "foo bar") + "foo bar"))
                        if v6 {
                            _ = len(v5)
                            _ = len(v5)
                            var v41 float32 = 3.14159
                            _ = v41
                        }
                        _ = v40
                    } else {
                        fmt.Println(v3)
                        v39 -= (v2 - v39)
                        v42 := GenS8{F0: 128, F1: 0}
                        _ = v42
                    }
                    _ = v39
                }
                fmt.Println(v1)
            }
            if ((!((fi8 != (v1 >> ((172) & 63))) || (-1 != -71))) && v6) {
                v2 = v2
                if v6 {
                    v2 -= v2
                } else {
                    v43 := GenS4{F0: (-106 != int32(v1))}
                    _ = v43
                }
                v44 := GenS9{F0: -259, F1: (v4 + (v4 >> ((42) & 63)))}
                fmt.Println(v3)
                _ = v44
            } else {
                v1 = 246
                var v45 int32 = 768
                v46 := GenS10{F0: 996}
                _ = v45
                _ = v46
            }
            _ = fi8
        }
        if (("foo bar" + "foo bar") > "\n\t") {
            v3 += v4
            var v47 string = (("\n\t" + "foo bar") + "hello")
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = v47
        } else {
            v48 := GenS11{F0: 18446744073709551615, F1: -108}
            v4 -= (v3 + (v3 >> ((1) & 63)))
            _ = v48
        }
        _ = fi7
    }
    v1 += v1
    fmt.Println(v4, v2)
    switch v1 {
    case 873:
        fmt.Println(v1)
    }
    v49 := GenS2{F0: (("a" + "foo bar") + "hello"), F1: 2}
    fmt.Println(v3, v3, v6)
    fmt.Println(v3, v3, v6)
    v50 := GenS9{F0: -2147483648, F1: v4}
    v3 = -1
    if len(v5) > 0 {
        v5[0] = "a"
    }
    for fi51 := 0; fi51 < 19; fi51++ {
        for fi52 := 0; fi52 < 14; fi52++ {
            fmt.Println(v4)
            _ = fi52
        }
        var v53 uint = 255
        _ = fi51
        _ = v53
    }
    for fi54 := 0; fi54 < 15; fi54++ {
        v55 := GenS9{F0: -1, F1: -9223372036854775808}
        fmt.Println(v1, fi54, v6)
        v5 = append(v5, "\n\t")
        _ = fi54
        _ = v55
    }
    fmt.Println(v1, v1)
    var v56 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v56
    _ = v6
}
