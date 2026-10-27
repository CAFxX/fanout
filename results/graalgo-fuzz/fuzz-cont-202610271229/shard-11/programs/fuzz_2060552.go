package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 uint8; F1 int8; F2 string }
type GenS4 struct { F0 uint16; F1 int16 }
type GenS5 struct { F0 int8; F1 string; F2 uint64 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int8; F1 bool }
type GenS9 struct { F0 bool; F1 int8; F2 uint16 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint; F1 bool }
type GenS12 struct { F0 int; F1 int32 }
type GenS13 struct { F0 int32; F1 int; F2 uint }
func genFn0() uint {
    var v7 int64 = -187
    var v8 uint = 2
    _ = v7
    _ = v8
    return v8
}

func main() {
    v1 := make(map[int]string)
    var v2 uint = 42
    v3 := []uint32{42, 233}
    var v4 uint8 = 42
    var v5 uint32 = 2147483647
    var v6 bool = ((v2 - v2) < 255)
    v9 := GenS1{F0: "foo bar"}
    switch 2147483647 {
    case 32767:
        fmt.Println(v5, v2)
    case -2147483648:
        v10 := GenS1{F0: "a"}
        _ = v10
    case 9223372036854775807:
        if true {
            var v11 float64 = -2.25
            var v12 float32 = 1e10
            _ = v11
            _ = v12
        } else {
            if false {
                var v13 uint16 = 42
                _ = v13
            }
        }
    default:
        switch "a" {
        case "":
            if (!(uint32(v4) > 2147483648)) {
                v6 = false
                for fi14 := 0; fi14 < 18; fi14++ {
                    v2 = 255
                    v2 += 18446744073709551615
                    _ = fi14
                }
                v15 := GenS1{F0: "foo bar"}
                _ = v15
            } else {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                if (0 != 0) {
                    v16 := GenS2{F0: (v5 * v5), F1: -32768}
                    _ = v16
                } else {
                    _ = len(v3)
                    v4 = (v4 & 0)
                }
            }
        case "\n\t":
            v2 = (v2 / ((4294967295) | 1))
        case "foo bar":
            v4 += 1
        default:
            fmt.Println(v5)
        }
    }
    for fi17 := 0; fi17 < 18; fi17++ {
        if ((3.14159 <= 1e10) && (362 == 0)) {
            for fi18 := 0; fi18 < 2; fi18++ {
                switch ("\n\t" + ("" + "\n\t")) {
                case "a":
                    switch (fi17 << ((genFn0()) & 63)) {
                    case 2147483647:
                        if true {
                            var v19 uint8 = uint8(v5)
                            _ = len(v1)
                            v19 += 148
                            _ = v19
                        } else {
                            v4 += v4
                            if v6 {
                                v20 := v1[fi18]
                                v21 := GenS1{F0: v20}
                                _ = v20
                                _ = v21
                            } else {
                                v22 := GenS3{F0: v4, F1: 1, F2: "foo bar"}
                                if (127 <= 42) {
                                    fmt.Println(fi17)
                                    var v23 uint8 = 127
                                    _ = v23
                                } else {
                                    v24 := GenS2{F0: v5, F1: 519}
                                    _ = v24
                                }
                                _ = v22
                            }
                            if false {
                                if v6 {
                                    var v25 int32 = 42
                                    _ = v25
                                } else {
                                    v26 := GenS2{F0: v5, F1: -11}
                                    _ = len(v1)
                                    var v27 int64 = 9223372036854775807
                                    _ = v26
                                    _ = v27
                                }
                            } else {
                                v4 = 255
                                fmt.Println(v4)
                            }
                        }
                    case 127:
                        var v28 int32 = -86
                        _ = v28
                    case -204:
                        fmt.Println(v6, v4)
                    }
                default:
                    v3 = append(v3, (v5 & 2147483648))
                    v29 := GenS4{F0: uint16(v2), F1: 32767}
                    if true {
                        v30 := GenS2{F0: (v5 * v5), F1: int16(v4)}
                        delete(v1, fi18)
                        v31 := GenS5{F0: 127, F1: ("foo bar" + "hello"), F2: 2}
                        _ = v30
                        _ = v31
                    } else {
                        v5 = 42
                        if ((uint32(v2) != 1) && false) {
                            v2 = genFn0()
                            fmt.Println(v5)
                        }
                        v32 := GenS5{F0: 127, F1: ("hello" + ""), F2: 235}
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        if true {
                            fmt.Println(v4, fi17)
                            switch (fi18 | int(v5)) {
                            case 9223372036854775807:
                                fmt.Println(fi18, v6)
                            }
                            v33 := GenS4{F0: 42, F1: -32768}
                            var v34 int = (-(fi17))
                            v35 := GenS6{F0: -994}
                            _ = v33
                            _ = v34
                            _ = v35
                        }
                        _ = v32
                    }
                    var v36 int16 = -1
                    _ = v29
                    _ = v36
                }
                v6 = (v2 < (v2 | v2))
                _ = fi18
            }
            if v6 {
                v37 := GenS7{F0: 4294967295}
                var v38 uint = v2
                var v39 float64 = 0.0
                _ = v37
                _ = v38
                _ = v39
            } else {
                if len(v3) > 0 {
                    v3[0] = v5
                }
                fmt.Println(fi17)
                if (v6 && (2 != 965)) {
                    var v40 string = ("\n\t" + ((("a" + ("a" + "\n\t")) + "\n\t") + "a"))
                    v41 := GenS2{F0: (v5 ^ (^(v5))), F1: 1000}
                    _ = v40
                    _ = v41
                } else {
                    for fi42 := 0; fi42 < 10; fi42++ {
                        var v43 bool = (580 < 1)
                        _ = fi42
                        _ = v43
                    }
                    for fi44 := 0; fi44 < 3; fi44++ {
                        if v6 {
                            if (-393 >= (fi17 & (fi44 / ((fi17) | 1)))) {
                                var v45 int32 = 0
                                var v46 int64 = 347
                                switch (fi44 % ((fi17) | 1)) {
                                case -9223372036854775808:
                                    v46 -= v46
                                }
                                _ = v45
                                _ = v46
                            } else {
                                v5 -= v5
                                v47 := GenS4{F0: 2, F1: -676}
                                var v48 float64 = 0.1
                                _ = len(v1)
                                v49 := GenS2{F0: (v5 * v5), F1: 32767}
                                switch fi44 {
                                case 42:
                                    v5 -= v5
                                case 121:
                                    switch (-(fi17)) {
                                    case 331:
                                        var v50 int32 = -649
                                        _ = v50
                                    }
                                case 127:
                                    switch ((((("\n\t" + "a") + ("hello" + ("hello" + (("\n\t" + "") + (("hello" + "\n\t") + "foo bar"))))) + ("" + ("foo bar" + "hello"))) + ("a" + "a")) + "a") {
                                    case "foo bar":
                                        v51 := GenS5{F0: 1, F1: ("\n\t" + "foo bar"), F2: 42}
                                        _ = v51
                                    }
                                    v52 := GenS8{F0: 127, F1: true}
                                    if true {
                                        switch "a" {
                                        case "hello":
                                            if len(v3) > 0 {
                                                v3[0] = v5
                                            }
                                        case "":
                                            fmt.Println(v48, fi17)
                                        }
                                        v48 -= v48
                                        fmt.Println(fi17, v5, fi44)
                                    } else {
                                        var v53 bool = (1 != 2147483647)
                                        v4 = 127
                                        _ = v53
                                    }
                                    _ = v52
                                default:
                                    switch fi17 {
                                    case 4294967295:
                                        if (1 != 2) {
                                            _ = len(v1)
                                            _ = len(v1)
                                            var v54 int16 = 1
                                            _ = v54
                                        } else {
                                            switch ("foo bar" + "foo bar") {
                                            case "hello":
                                                var v55 string = (("a" + ("a" + "hello")) + (("a" + "foo bar") + ((("\n\t" + "hello") + ("hello" + "\n\t")) + "\n\t")))
                                                _ = v55
                                            default:
                                                if v6 {
                                                    if (!((("\n\t" + "foo bar") + "hello") >= "foo bar")) {
                                                        if (fi44 <= -1) {
                                                            v6 = ((v5 < (v5 & (v5 ^ v5))) || (973 < v2))
                                                            switch fi44 {
                                                            case -724:
                                                                if v6 {
                                                                    _ = len(v3)
                                                                    var v56 uint32 = v5
                                                                    _ = v56
                                                                } else {
                                                                    var v57 int64 = 9223372036854775807
                                                                    _ = v57
                                                                }
                                                            case 521:
                                                                var v58 uint = (v2 / ((v2) | 1))
                                                                _ = v58
                                                            }
                                                        } else {
                                                            var v59 uint64 = 0
                                                            switch fi17 {
                                                            case -1:
                                                                switch "a" {
                                                                case "foo bar":
                                                                    var v60 int = (fi17 & (fi17 & fi44))
                                                                    _ = v60
                                                                case "a":
                                                                    var v61 int16 = 42
                                                                    _ = v61
                                                                case "hello":
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                default:
                                                                    if v6 {
                                                                        v62 := GenS4{F0: 2, F1: 42}
                                                                        v63 := GenS9{F0: v6, F1: 1, F2: 744}
                                                                        _ = v62
                                                                        _ = v63
                                                                    } else {
                                                                        switch ("" + "") {
                                                                        case "hello":
                                                                            v48 = v48
                                                                        }
                                                                        v64 := GenS10{F0: (fi17 | (fi17 & (fi17 + fi44)))}
                                                                        _ = v64
                                                                    }
                                                                }
                                                                v2 -= (^(v2))
                                                            }
                                                            var v65 int32 = 1
                                                            _ = v59
                                                            _ = v65
                                                        }
                                                        if v6 {
                                                            v66 := GenS10{F0: (fi44 * 745)}
                                                            v6 = (true && v6)
                                                            if v6 {
                                                                switch (fi44 - fi17) {
                                                                case -80:
                                                                    var v67 int = fi44
                                                                    _ = v67
                                                                case 127:
                                                                    var v68 float32 = 1.5
                                                                    _ = v68
                                                                }
                                                                var v69 int32 = 1
                                                                v48 = v48
                                                                _ = v69
                                                            } else {
                                                                v4 -= (v4 * (v4 & 1))
                                                                v2 = v2
                                                                v48 -= v48
                                                            }
                                                            _ = v66
                                                        }
                                                        var v70 string = "foo bar"
                                                        _ = v70
                                                    }
                                                    v4 -= v4
                                                    delete(v1, fi44)
                                                }
                                            }
                                            var v71 int = (fi44 << (((^(v2))) & 63))
                                            var v72 int16 = -58
                                            _ = v71
                                            _ = v72
                                        }
                                    default:
                                        v3 = append(v3, v5)
                                        v73 := GenS11{F0: v2, F1: (!true)}
                                        fmt.Println(fi44, v5)
                                        _ = v73
                                    }
                                }
                                _ = v47
                                _ = v48
                                _ = v49
                            }
                            var v74 uint16 = 0
                            _ = v74
                        } else {
                            if ((v2 + genFn0()) >= v2) {
                                var v75 bool = (42 == 2)
                                _ = v75
                            } else {
                                v4 -= v4
                            }
                        }
                        if false {
                            _ = len(v3)
                            var v76 uint = v2
                            v6 = (fi17 < (fi17 - (fi44 ^ (fi17 >> (((v76 % ((4294967295) | 1))) & 63)))))
                            _ = v76
                        }
                        if (261 > 1) {
                            fmt.Println(v2)
                        } else {
                            var v77 uint32 = v5
                            _ = v77
                        }
                        _ = fi44
                    }
                }
            }
            v78 := GenS12{F0: (fi17 >> ((v2) & 63)), F1: 239}
            _ = v78
        } else {
            v79 := GenS12{F0: fi17, F1: 0}
            v4 -= v4
            _ = v79
        }
        _ = fi17
    }
    v2 = (v2 % ((v2) | 1))
    v80 := GenS13{F0: 2147483647, F1: 127, F2: v2}
    fmt.Println(v2, v4)
    v4 = v4
    delete(v1, 4294967295)
    if (("" + "a") != ("foo bar" + "hello")) {
        if false {
            v5 = uint32(v2)
            for fi81 := 0; fi81 < 19; fi81++ {
                var v82 float64 = 0.0
                if v6 {
                    fmt.Println(v6)
                    var v83 bool = (((!true) || v6) && v6)
                    v5 = (^(v5))
                    _ = v83
                } else {
                    v4 -= (v4 << (((v2 / ((genFn0()) | 1))) & 63))
                }
                for fi84 := 0; fi84 < 19; fi84++ {
                    v82 -= v82
                    if v6 {
                        fmt.Println(v5)
                        v85 := v1[(fi81 | fi81)]
                        _ = v85
                    }
                    if (!(0 <= 4294967295)) {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        switch (-(fi84)) {
                        case 127:
                            v2 += v2
                        case 677:
                            v4 = v4
                        }
                        v86 := GenS3{F0: 255, F1: 0, F2: "a"}
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v87 := GenS4{F0: 255, F1: -1}
                        _ = v86
                        _ = v87
                    }
                    _ = fi84
                }
                _ = fi81
                _ = v82
            }
        }
    }
    switch "" {
    case "a":
        for fi88 := 0; fi88 < 4; fi88++ {
            v2 = (v2 % ((4294967295) | 1))
            _ = fi88
        }
    }
    var v89 uint = (^(v2))
    v5 -= (v5 - v5)
    v2 = (v89 >> (((v89 - v2)) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v80
    _ = v89
    _ = v9
}
