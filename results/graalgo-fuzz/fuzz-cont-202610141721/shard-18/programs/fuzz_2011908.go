package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 int64; F1 uint16; F2 uint32 }
type GenS3 struct { F0 int; F1 bool; F2 string }
type GenS4 struct { F0 uint; F1 bool; F2 string }
type GenS5 struct { F0 uint8; F1 uint8; F2 uint8 }
type GenS6 struct { F0 uint64; F1 uint64 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint8; F1 uint32 }
type GenS9 struct { F0 bool; F1 bool }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 bool }
func genFn0() int16 {
    var v4 string = ((("foo bar" + ((((("hello" + "foo bar") + "") + "a") + "\n\t") + ("hello" + "hello"))) + ("" + "foo bar")) + "")
    var v5 uint = 796
    fmt.Println(v4, v5)
    _ = v4
    _ = v5
    return 570
}
func genFn1() uint64 {
    var v6 float32 = 3.14159
    for fi7 := 0; fi7 < 18; fi7++ {
        for fi8 := 0; fi8 < 4; fi8++ {
            v9 := GenS1{F0: (fi7 / ((fi7) | 1)), F1: fi7}
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 6; fi10++ {
            if true {
                var v11 uint64 = 18446744073709551615
                v6 += v6
                _ = v11
            }
            var v12 int64 = 42
            var v13 uint32 = 42
            _ = fi10
            _ = v12
            _ = v13
        }
        var v14 uint64 = 2
        _ = fi7
        _ = v14
    }
    switch 2147483647 {
    case 416:
        for fi15 := 0; fi15 < 3; fi15++ {
            v6 = v6
            _ = fi15
        }
    case 823:
        v16 := GenS2{F0: 1, F1: 720, F2: 1}
        _ = v16
    case 282:
        if true {
            var v17 uint = 1
            if (v17 >= v17) {
                for fi18 := 0; fi18 < 10; fi18++ {
                    var v19 uint64 = 1
                    _ = fi18
                    _ = v19
                }
            } else {
                for fi20 := 0; fi20 < 8; fi20++ {
                    if true {
                        v17 = v17
                        fmt.Println(v6, v17, fi20)
                    } else {
                        for fi21 := 0; fi21 < 17; fi21++ {
                            var v22 bool = (0 <= 1)
                            _ = fi21
                            _ = v22
                        }
                        for fi23 := 0; fi23 < 12; fi23++ {
                            v17 = (v17 >> (((v17 + (v17 >> (((v17 + 474)) & 63)))) & 63))
                            var v24 bool = ((!(660 < 0)) || false)
                            fmt.Println(fi20, fi23, fi23)
                            _ = fi23
                            _ = v24
                        }
                        if (true && (0.0 == 3.14159)) {
                            v6 = (float32(float64((1e10 - (1e10 * v6)))) / 1e10)
                        }
                    }
                    _ = fi20
                }
            }
            _ = v17
        } else {
            v25 := GenS3{F0: 9223372036854775807, F1: true, F2: ("a" + ("a" + "foo bar"))}
            var v26 int64 = 896
            var v27 int16 = -18
            _ = v25
            _ = v26
            _ = v27
        }
    }
    _ = v6
    return 42
}

func main() {
    v1 := make(map[string]string)
    v2 := []int{}
    v3 := []uint32{}
    switch 42 {
    case -9223372036854775808:
        v28 := GenS4{F0: 311, F1: (false && false), F2: "hello"}
        _ = v28
    case -74:
        v29 := GenS5{F0: 0, F1: 127, F2: 238}
        _ = v29
    case -145:
        var v30 int64 = -1
        _ = v30
    }
    var v31 int8 = -80
    for fi32 := 0; fi32 < 9; fi32++ {
        v31 = 42
        _ = fi32
    }
    var v33 uint8 = 255
    for fi34 := 0; fi34 < 6; fi34++ {
        if (104 >= 122) {
            var v35 int64 = 513
            fmt.Println(v33)
            _ = v35
        }
        v36 := GenS6{F0: 18446744073709551615, F1: 2}
        switch fi34 {
        case -9223372036854775808:
            delete(v1, "")
        case -2147483648:
            v37 := GenS5{F0: 2, F1: (v33 << ((1) & 63)), F2: (v33 * v33)}
            _ = v37
        default:
            var v38 int8 = v31
            _ = v38
        }
        _ = fi34
        _ = v36
    }
    var v39 int16 = 500
    v33 -= 0
    if ((v33 | v33) >= v33) {
        var v40 uint8 = v33
        v41 := GenS7{F0: 0}
        _ = v40
        _ = v41
    } else {
        switch 128 {
        case 9223372036854775807:
            var v42 bool = false
            _ = v42
        case -2147483648:
            switch ("foo bar" + "a") {
            case "foo bar":
                var v43 string = "foo bar"
                _ = v43
            case "a":
                v39 -= genFn0()
            case "\n\t":
                var v44 uint32 = 1
                _ = v44
            default:
                if (true || (true || false)) {
                    if true {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v45 int32 = 264
                        _ = v45
                    }
                    if false {
                        if (true && false) {
                            var v46 uint16 = 32767
                            var v47 uint = 115
                            v47 -= v47
                            _ = v46
                            _ = v47
                        } else {
                            for fi48 := 0; fi48 < 10; fi48++ {
                                v49 := GenS8{F0: (v33 % ((v33) | 1)), F1: 2147483647}
                                _ = len(v3)
                                _ = fi48
                                _ = v49
                            }
                            if false {
                                fmt.Println(v31, v39, v31)
                            } else {
                                v50 := GenS9{F0: true, F1: false}
                                var v51 int = -584
                                _ = v50
                                _ = v51
                            }
                        }
                    }
                    if (118 != (^(v31))) {
                        for fi52 := 0; fi52 < 2; fi52++ {
                            switch "" {
                            case "foo bar":
                                if true {
                                    if (true || false) {
                                        var v53 uint8 = (v33 >> ((538) & 63))
                                        v54 := GenS8{F0: (^(v33)), F1: 2147483648}
                                        _ = v53
                                        _ = v54
                                    } else {
                                        var v55 int8 = (v31 + 127)
                                        _ = v55
                                    }
                                    v56 := GenS10{F0: 363}
                                    if false {
                                        switch ((((("a" + "\n\t") + "a") + "foo bar") + ("hello" + (((("" + "") + "a") + ("" + "hello")) + "a"))) + ("foo bar" + ("" + "hello"))) {
                                        case "hello":
                                            switch "\n\t" {
                                            case "hello":
                                                if true {
                                                    v57 := GenS11{F0: 2147483647}
                                                    var v58 int = (fi52 >> ((18446744073709551615) & 63))
                                                    _ = v57
                                                    _ = v58
                                                }
                                            case "\n\t":
                                                if (330 > uint(v33)) {
                                                    fmt.Println(fi52, v31, v31)
                                                    var v59 int64 = 0
                                                    _ = v59
                                                }
                                            case "a":
                                                var v60 string = "\n\t"
                                                _ = v60
                                            }
                                        default:
                                            v61 := GenS12{F0: (!(v39 <= (-(v39))))}
                                            _ = v61
                                        }
                                    }
                                    _ = v56
                                } else {
                                    fmt.Println(v33, fi52)
                                    var v62 uint = 1
                                    _ = v62
                                }
                            default:
                                v63 := v1[(("a" + "") + "hello")]
                                _ = v63
                            }
                            fmt.Println(v39, v33, fi52)
                            var v64 float32 = (((-2.25 * -2.25) - 3.14159) - 0.0)
                            _ = fi52
                            _ = v64
                        }
                    }
                }
            }
        case 32767:
            var v65 uint16 = 419
            _ = v65
        default:
            var v66 int32 = 42
            _ = v66
        }
        v39 += (v39 % (((v39 / (((v39 * int16(v31))) | 1))) | 1))
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    var v67 int = -2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v39:", v39)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v39
    _ = v67
}
