package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 string; F1 string; F2 uint32 }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int32; F1 int64; F2 string }
type GenS6 struct { F0 uint16; F1 int32; F2 uint32 }
func genFn0(p0 int8) uint {
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 float32 = -2.25
        _ = fi7
        _ = v8
    }
    var v9 uint = 138
    for fi10 := 0; fi10 < 6; fi10++ {
        var v11 int64 = 192
        switch "hello" {
        case "foo bar":
            switch fi10 {
            case 4294967295:
                v12 := GenS1{F0: 0, F1: 2}
                _ = v12
            default:
                v13 := GenS2{F0: "foo bar", F1: ("hello" + ("hello" + "hello")), F2: 42}
                _ = v13
            }
        }
        _ = fi10
        _ = v11
    }
    var v14 uint16 = 736
    _ = p0
    _ = v14
    _ = v9
    return 18446744073709551615
}
func genFn1() int64 {
    for fi15 := 0; fi15 < 11; fi15++ {
        var v16 int = fi15
        fmt.Println(v16, v16)
        v16 += 4294967295
        _ = fi15
        _ = v16
    }
    var v17 float64 = (0.1 + 0.0)
    v17 = v17
    fmt.Println(v17)
    _ = v17
    return -453
}

func main() {
    var v1 int = -810
    var v2 uint64 = 964
    v3 := []string{("a" + "a"), "\n\t"}
    var v4 float64 = 1e10
    var v5 bool = (("a" + (("\n\t" + "a") + ("\n\t" + ("hello" + ("" + ("" + ("\n\t" + "hello"))))))) != ("a" + "a"))
    var v6 bool = (1 >= 116)
    v5 = (!((("" + ("hello" + ("\n\t" + "\n\t"))) + "hello") == (("foo bar" + "a") + ("hello" + ("a" + "")))))
    v5 = false
    v3 = append(v3, (((("foo bar" + "hello") + "a") + ("a" + ("hello" + "a"))) + "hello"))
    for fi18 := 0; fi18 < 18; fi18++ {
        for fi19 := 0; fi19 < 14; fi19++ {
            v20 := GenS3{F0: 1, F1: genFn0(42)}
            var v21 float64 = v4
            if (0.1 > 3.14159) {
                switch "foo bar" {
                case "\n\t":
                    v6 = v5
                case "":
                    v4 = v4
                }
                var v22 int64 = -9223372036854775808
                fmt.Println(fi18)
                _ = v22
            } else {
                var v23 string = ("" + "\n\t")
                var v24 uint64 = 227
                _ = v23
                _ = v24
            }
            _ = fi19
            _ = v20
            _ = v21
        }
        for fi25 := 0; fi25 < 4; fi25++ {
            _ = len(v3)
            if (!(-1 > -744)) {
                v26 := GenS1{F0: -1, F1: uint16(v2)}
                var v27 float32 = 3.14159
                _ = len(v3)
                _ = v26
                _ = v27
            }
            v5 = (v5 || true)
            _ = fi25
        }
        _ = fi18
    }
    _ = len(v3)
    if ((-1 > -1) && (("hello" + "hello") != "\n\t")) {
        var v28 uint8 = 2
        if (v6 || v5) {
            if (127 > -128) {
                _ = len(v3)
            } else {
                v29 := GenS3{F0: 255, F1: 18446744073709551615}
                if (v6 && (!("\n\t" == (((("\n\t" + ("foo bar" + "\n\t")) + ("hello" + "\n\t")) + ("\n\t" + "a")) + "hello")))) {
                    var v30 int64 = -358
                    v2 -= v2
                    if ((2 < 254) && (0.1 == (0.1 - 3.14159))) {
                        for fi31 := 0; fi31 < 10; fi31++ {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v5 = v5
                            _ = fi31
                        }
                    } else {
                        fmt.Println(v28)
                        v5 = (true || ((((100 > 1) || v6) || (((1e10 * 0.0) == (-2.25 / (1e10 / float32(v4)))) || v5)) && v5))
                    }
                    _ = v30
                }
                if v6 {
                    for fi32 := 0; fi32 < 4; fi32++ {
                        _ = len(v3)
                        fmt.Println(v28, v2, v5)
                        var v33 int32 = -774
                        _ = fi32
                        _ = v33
                    }
                    v28 += v28
                } else {
                    for fi34 := 0; fi34 < 7; fi34++ {
                        fmt.Println(fi34)
                        v35 := GenS4{F0: v6}
                        for fi36 := 0; fi36 < 8; fi36++ {
                            _ = len(v3)
                            if (float32(v4) < 1.5) {
                                fmt.Println(v4, fi36, v1)
                                switch fi34 {
                                case 127:
                                    v37 := GenS1{F0: -12, F1: 1}
                                    _ = v37
                                case 42:
                                    if v6 {
                                        fmt.Println(v4, fi36, fi36)
                                        switch "a" {
                                        case "\n\t":
                                            switch "foo bar" {
                                            case "\n\t":
                                                v38 := GenS5{F0: 2147483647, F1: 1, F2: (("\n\t" + "") + "a")}
                                                _ = v38
                                            case "":
                                                v39 := GenS2{F0: "foo bar", F1: "hello", F2: 203}
                                                _ = v39
                                            default:
                                                v3 = append(v3, "foo bar")
                                            }
                                        case "a":
                                            if (fi34 != (-(v1))) {
                                                var v40 int64 = 627
                                                _ = v40
                                            }
                                        default:
                                            var v41 int64 = 0
                                            _ = v41
                                        }
                                    }
                                default:
                                    var v42 uint8 = 212
                                    _ = v42
                                }
                            }
                            _ = fi36
                        }
                        _ = fi34
                        _ = v35
                    }
                    fmt.Println(v28, v5, v28)
                }
                _ = v29
            }
            if (!false) {
                if (v5 && false) {
                    var v43 float64 = (float64(((1e10 - ((1.5 + float32(float64(1.5))) / 0.1)) - (1.5 + (1.5 - 3.14159)))) + (0.1 / v4))
                    v43 -= 0.0
                    if (v6 || v5) {
                        var v44 int = (v1 / ((v1) | 1))
                        switch (v44 << ((4294967295) & 63)) {
                        case -421:
                            v45 := GenS1{F0: 42, F1: 42}
                            _ = v45
                        default:
                            v46 := GenS1{F0: 1, F1: 251}
                            _ = v46
                        }
                        var v47 float64 = v4
                        _ = v44
                        _ = v47
                    }
                    _ = v43
                } else {
                    v48 := GenS4{F0: (v5 && v5)}
                    fmt.Println(v28, v28)
                    var v49 uint = 0
                    _ = v48
                    _ = v49
                }
                _ = len(v3)
                if (0 != 0) {
                    v5 = (-126 <= 2147483647)
                    switch "\n\t" {
                    case "hello":
                        v6 = true
                    case "foo bar":
                        v28 = v28
                    case "a":
                        v4 = (-2.25 + v4)
                    }
                } else {
                    var v50 float64 = v4
                    var v51 uint8 = (v28 | 2)
                    _ = v50
                    _ = v51
                }
            }
        }
        if v5 {
            v28 = (v28 * 127)
            for fi52 := 0; fi52 < 8; fi52++ {
                for fi53 := 0; fi53 < 15; fi53++ {
                    fmt.Println(v5)
                    v1 -= v1
                    _ = fi53
                }
                _ = fi52
            }
            for fi54 := 0; fi54 < 19; fi54++ {
                v55 := GenS4{F0: false}
                _ = len(v3)
                v28 += (^(v28))
                _ = fi54
                _ = v55
            }
        }
        _ = v28
    }
    v5 = ((!((!((1e10 + 3.14159) != 0.1)) && (18446744073709551615 <= genFn0(-88)))) || (true || v5))
    v4 += ((v4 - (v4 - 0.1)) / (v4 / v4))
    if true {
        if true {
            v3 = append(v3, "hello")
        } else {
            var v56 int64 = -9223372036854775808
            _ = v56
        }
        var v57 bool = v6
        v5 = (!(v57 || v5))
        _ = v57
    }
    fmt.Println(v4)
    v2 = v2
    var v58 uint16 = 1
    v1 += -754
    v59 := GenS6{F0: uint16(v2), F1: 42, F2: 1}
    v60 := GenS6{F0: v58, F1: -1, F2: 2147483648}
    v61 := GenS4{F0: (v1 == v1)}
    var v62 float32 = 1e10
    v5 = (!v6)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v59
    _ = v6
    _ = v60
    _ = v61
    _ = v62
}
