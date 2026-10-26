package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 int8; F1 uint8; F2 uint32 }
type GenS4 struct { F0 int32; F1 uint8; F2 int8 }
type GenS5 struct { F0 string; F1 int16; F2 uint8 }
type GenS6 struct { F0 int }
func genFn0(p0 int16, p1 uint, p2 uint16) string {
    var v6 string = "a"
    fmt.Println(p1)
    p1 = 255
    fmt.Println(p0, p1, v6)
    if (((2 > (p2 * p2)) && (0.0 <= -2.25)) && (!false)) {
        fmt.Println(p2, p1, p0)
    } else {
        var v7 string = v6
        if (4294967295 > 253) {
            if false {
                fmt.Println(v6)
            }
            for fi8 := 0; fi8 < 18; fi8++ {
                fmt.Println(p2, v7)
                _ = fi8
            }
            switch -788 {
            case -1:
                switch ("hello" + v6) {
                case "hello":
                    var v9 uint8 = 1
                    _ = v9
                case "foo bar":
                    fmt.Println(v7, v6)
                }
            case 128:
                for fi10 := 0; fi10 < 3; fi10++ {
                    v11 := []int{(fi10 - (-(fi10)))}
                    fmt.Println(p1)
                    fmt.Println(v7)
                    _ = fi10
                    _ = v11
                }
            }
        }
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (("\n\t" + "foo bar") + ((v6 + "foo bar") + v6))
}
func genFn1() int64 {
    var v12 uint = 18446744073709551615
    for fi13 := 0; fi13 < 8; fi13++ {
        var v14 int = (fi13 >> ((v12) & 63))
        _ = fi13
        _ = v14
    }
    v12 = 65535
    var v15 int32 = -709
    _ = v12
    _ = v15
    return 9223372036854775807
}

func main() {
    var v1 uint16 = 42
    var v2 int32 = 1
    v3 := make(map[string]string)
    var v4 float32 = float32((float64(3.14159) / (1e10 + (3.14159 * 0.1))))
    var v5 string = ""
    v5 = "a"
    var v16 uint8 = 127
    v1 -= (v1 >> ((4294967295) & 63))
    v17 := []int{9, 4294967295}
    if len(v17) > 0 {
        v17[0] = 32768
    }
    if false {
        fmt.Println(v5, v4)
        var v18 uint8 = (v16 - 2)
        fmt.Println(v1, v2, v4)
        _ = v18
    } else {
        for fi19 := 0; fi19 < 20; fi19++ {
            if false {
                fmt.Println(fi19)
                if (true && false) {
                    for fi20 := 0; fi20 < 2; fi20++ {
                        if false {
                            if (fi20 != fi19) {
                                if false {
                                    if false {
                                        v21 := v3[(v5 + v5)]
                                        var v22 int16 = -32768
                                        var v23 float64 = 1.5
                                        _ = v21
                                        _ = v22
                                        _ = v23
                                    } else {
                                        var v24 int = (fi19 + (fi20 + (fi20 * fi20)))
                                        v1 -= v1
                                        _ = v24
                                    }
                                    fmt.Println(v16, fi20, v4)
                                } else {
                                    v4 += 1.5
                                    fmt.Println(fi20, v5)
                                    var v25 uint32 = 4294967295
                                    _ = v25
                                }
                                var v26 uint32 = 2147483648
                                if len(v17) > 0 {
                                    v17[0] = (fi20 / ((fi19) | 1))
                                }
                                _ = v26
                            }
                            v1 = (v1 | v1)
                            if (1 == genFn1()) {
                                v3[v5] = (("foo bar" + v5) + ("a" + "foo bar"))
                            }
                            switch "" {
                            case "hello":
                                if false {
                                    var v27 uint32 = 2
                                    _ = v27
                                }
                            case "a":
                                v16 -= (v16 | 42)
                                v5 = ("foo bar" + "")
                                switch v5 {
                                case "hello":
                                    fmt.Println(v16, v2, fi20)
                                default:
                                    v5 = v5
                                }
                            }
                        } else {
                            var v28 uint16 = v1
                            switch fi19 {
                            case -576:
                                if (v2 <= (v2 ^ v2)) {
                                    var v29 uint8 = (v16 * (v16 >> ((0) & 63)))
                                    _ = v29
                                } else {
                                    fmt.Println(v28, v2)
                                    if (!(937 >= (v1 >> ((1) & 63)))) {
                                        if true {
                                            fmt.Println(v2, v16, v5)
                                            var v30 int64 = -9223372036854775808
                                            if (!(255 < 0)) {
                                                var v31 float64 = 1e10
                                                _ = v31
                                            } else {
                                                var v32 uint64 = 1
                                                v33 := GenS1{F0: true, F1: 683}
                                                _ = v32
                                                _ = v33
                                            }
                                            _ = v30
                                        } else {
                                            fmt.Println(fi20, v5, v16)
                                            v34 := GenS1{F0: (!false), F1: 42}
                                            v35 := GenS1{F0: false, F1: 18446744073709551615}
                                            _ = v34
                                            _ = v35
                                        }
                                    } else {
                                        _ = len(v17)
                                        var v36 string = ((v5 + v5) + v5)
                                        fmt.Println(v1)
                                        _ = v36
                                    }
                                    var v37 uint16 = v28
                                    v5 = "hello"
                                    switch (("hello" + (v5 + "hello")) + v5) {
                                    case "hello":
                                        fmt.Println(v2, v2)
                                    case "foo bar":
                                        var v38 int = fi19
                                        _ = v38
                                    case "a":
                                        var v39 string = ("foo bar" + v5)
                                        _ = v39
                                    }
                                    _ = v37
                                }
                            case 128:
                                v16 = (^(v16))
                            case 32767:
                                if (false && true) {
                                    v40 := GenS2{F0: false, F1: 2}
                                    _ = v40
                                }
                            default:
                                if (fi20 > (fi19 ^ fi20)) {
                                    _ = len(v17)
                                    v4 = float32((1.5 / 3.14159))
                                    var v41 int32 = v2
                                    _ = v41
                                }
                            }
                            _ = v28
                        }
                        v42 := GenS3{F0: 1, F1: 0, F2: 4294967295}
                        fmt.Println(v5)
                        _ = fi20
                        _ = v42
                    }
                    v2 = (^(v2))
                } else {
                    switch v5 {
                    case "foo bar":
                        fmt.Println(v1, fi19)
                    }
                    var v43 uint8 = v16
                    switch "a" {
                    case "hello":
                        v44 := GenS2{F0: (v4 >= v4), F1: 780}
                        _ = v44
                    case "\n\t":
                        v45 := GenS4{F0: 0, F1: 76, F2: 127}
                        _ = v45
                    default:
                        for fi46 := 0; fi46 < 19; fi46++ {
                            v43 = v16
                            var v47 uint64 = 1
                            if len(v17) >= 2 {
                                v17 = v17[1:len(v17)]
                            }
                            _ = fi46
                            _ = v47
                        }
                    }
                    _ = v43
                }
            } else {
                v48 := GenS3{F0: 8, F1: uint8(v1), F2: 42}
                var v49 float32 = v4
                _ = v48
                _ = v49
            }
            var v50 uint8 = 42
            fmt.Println(v16)
            _ = fi19
            _ = v50
        }
        if ((-1 == 42) && true) {
            var v51 uint8 = uint8(v1)
            switch v5 {
            case "":
                var v52 uint = 42
                _ = v52
            case "hello":
                var v53 string = "foo bar"
                _ = v53
            case "a":
                fmt.Println(v5, v1)
            default:
                v54 := GenS5{F0: v5, F1: -32768, F2: (^(v51))}
                _ = v54
            }
            _ = v51
        }
        v55 := GenS2{F0: (18446744073709551615 <= 18446744073709551615), F1: 4294967295}
        _ = v55
    }
    fmt.Println(v16)
    v1 -= v1
    if false {
        v2 -= v2
        var v56 float64 = 0.0
        _ = v56
    }
    var v57 float32 = v4
    if true {
        if len(v17) >= 2 {
            v17 = v17[1:len(v17)]
        }
        var v58 uint32 = 65535
        _ = v58
    } else {
        switch v5 {
        case "a":
            fmt.Println(v16)
        default:
            var v59 uint = 756
            _ = v59
        }
    }
    v16 = (v16 ^ 183)
    for fi60 := 0; fi60 < 7; fi60++ {
        for fi61 := 0; fi61 < 17; fi61++ {
            if true {
                v57 += 0.0
            }
            var v62 int = fi60
            var v63 string = v5
            v64 := GenS1{F0: false, F1: 1}
            _ = fi61
            _ = v62
            _ = v63
            _ = v64
        }
        v57 = v57
        v4 += float32(0.1)
        _ = fi60
    }
    if len(v17) > 0 {
        v17[0] = 652
    }
    var v65 string = ("" + (v5 + ((v5 + v5) + v5)))
    for fi66 := 0; fi66 < 16; fi66++ {
        if (0 < 4294967295) {
            switch (v65 + "") {
            case "":
                var v67 int = (^(fi66))
                _ = v67
            }
            var v68 int64 = 730
            _ = v68
        }
        _ = fi66
    }
    v69 := GenS6{F0: 127}
    v3[v65] = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v65
    _ = v69
}
