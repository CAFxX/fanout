package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string; F1 bool; F2 uint32 }
type GenS4 struct { F0 int16; F1 uint8 }
type GenS5 struct { F0 uint16; F1 uint8 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint8; F1 int; F2 int8 }
type GenS8 struct { F0 string; F1 int8; F2 uint }
type GenS9 struct { F0 int16; F1 uint8; F2 uint8 }
func genFn0(p0 int8) uint64 {
    var v6 int16 = 1000
    if (1 == 2147483648) {
        fmt.Println(p0, p0)
    }
    _ = p0
    _ = v6
    return 42
}
func genFn1(p0 bool) uint16 {
    if p0 {
        fmt.Println(p0)
        p0 = (true && (false || (p0 || (!false))))
        var v7 float64 = float64(0.0)
        _ = v7
    } else {
        var v8 uint64 = 2
        _ = v8
    }
    for fi9 := 0; fi9 < 1; fi9++ {
        var v10 uint64 = genFn0(100)
        var v11 int64 = 0
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = p0
    return 888
}

func main() {
    var v1 int8 = 0
    var v2 int64 = -1
    var v3 uint = 42
    v4 := []uint32{}
    var v5 uint32 = 2
    for fi12 := 0; fi12 < 8; fi12++ {
        v13 := GenS1{F0: 620, F1: 0, F2: 0}
        for fi14 := 0; fi14 < 20; fi14++ {
            fmt.Println(v1, v2)
            _ = fi14
        }
        switch "foo bar" {
        case "":
            for fi15 := 0; fi15 < 4; fi15++ {
                v3 = v3
                _ = fi15
            }
        case "hello":
            v1 = int8(v2)
        case "foo bar":
            fmt.Println(v3, v2, fi12)
        default:
            v3 = (v3 * (^(v3)))
        }
        _ = fi12
        _ = v13
    }
    v2 = (v2 * (v2 + (-(v2))))
    if true {
        if len(v4) > 0 {
            v4[0] = 42
        }
    } else {
        if true {
            for fi16 := 0; fi16 < 7; fi16++ {
                var v17 int32 = -946
                _ = fi16
                _ = v17
            }
            v18 := GenS2{F0: 0}
            v3 -= 1
            _ = v18
        }
        var v19 int32 = 1
        v20 := GenS1{F0: 0, F1: 32767, F2: 617}
        _ = v19
        _ = v20
    }
    v21 := GenS1{F0: 4294967295, F1: 1, F2: 0}
    v1 = 1
    if (0 != -32768) {
        v22 := GenS3{F0: (("\n\t" + "hello") + "hello"), F1: (!(true || true)), F2: v5}
        _ = v22
    } else {
        v5 = (v5 / ((v5) | 1))
        for fi23 := 0; fi23 < 10; fi23++ {
            v24 := GenS3{F0: "hello", F1: (42 > 32767), F2: 2147483648}
            _ = fi23
            _ = v24
        }
    }
    if false {
        var v25 int16 = -32768
        switch "hello" {
        case "":
            if true {
                v1 = (v1 | -104)
            }
        case "\n\t":
            v26 := GenS1{F0: 341, F1: 255, F2: -2147483648}
            _ = v26
        case "foo bar":
            v27 := GenS3{F0: ("foo bar" + ""), F1: (-2.25 < 1.5), F2: (v5 << ((v3) & 63))}
            _ = v27
        }
        for fi28 := 0; fi28 < 19; fi28++ {
            v29 := GenS2{F0: int32(v5)}
            _ = fi28
            _ = v29
        }
        _ = v25
    }
    v5 += 457
    fmt.Println(v1, v1)
    fmt.Println(v1, v5, v1)
    var v30 int64 = 0
    v2 -= (-(v30))
    v31 := GenS4{F0: 1, F1: 127}
    if ((1 != -1) || false) {
        for fi32 := 0; fi32 < 20; fi32++ {
            switch "hello" {
            case "\n\t":
                var v33 string = "hello"
                _ = v33
            case "hello":
                if ((v1 >> (((v3 % (((v3 % (((v3 * v3)) | 1))) | 1))) & 63)) > -57) {
                    if (181 >= 2) {
                        var v34 int16 = -925
                        var v35 int = (-(fi32))
                        var v36 int32 = 1
                        _ = v34
                        _ = v35
                        _ = v36
                    }
                    v5 = v5
                    for fi37 := 0; fi37 < 14; fi37++ {
                        if (1000 < -367) {
                            v30 += (v30 & v2)
                            var v38 uint16 = 65535
                            fmt.Println(v38, v2)
                            _ = v38
                        } else {
                            if true {
                                fmt.Println(fi37)
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                v2 = -9223372036854775808
                                var v39 float64 = 3.14159
                                fmt.Println(v2)
                                var v40 string = "hello"
                                switch ((v40 + ((("\n\t" + ("a" + ((v40 + v40) + (v40 + "hello")))) + "a") + "a")) + "") {
                                case "\n\t":
                                    v2 += 42
                                default:
                                    fmt.Println(v3, v1, v2)
                                }
                                _ = v39
                                _ = v40
                            } else {
                                if (false || (v1 > 100)) {
                                    var v41 bool = ((v3 ^ (v3 | v3)) != v3)
                                    _ = v41
                                } else {
                                    v42 := GenS5{F0: 42, F1: 0}
                                    var v43 float32 = 0.1
                                    if (!(true || (fi37 < fi32))) {
                                        v3 -= v3
                                    } else {
                                        switch ("a" + "") {
                                        case "hello":
                                            var v44 uint = 65535
                                            _ = v44
                                        case "a":
                                            if false {
                                                fmt.Println(v2, fi37, v30)
                                            } else {
                                                v45 := GenS6{F0: (v2 * (v2 - (^(v30))))}
                                                v46 := GenS4{F0: 42, F1: 0}
                                                v47 := GenS1{F0: 0, F1: 42, F2: 706}
                                                _ = v45
                                                _ = v46
                                                _ = v47
                                            }
                                        default:
                                            var v48 uint64 = 1
                                            _ = v48
                                        }
                                        v4 = append(v4, (v5 >> ((0) & 63)))
                                    }
                                    v49 := GenS7{F0: 255, F1: (-(fi37)), F2: (v1 * (v1 & (-(v1))))}
                                    _ = v42
                                    _ = v43
                                    _ = v49
                                }
                                switch fi37 {
                                case -956:
                                    v50 := GenS1{F0: 42, F1: genFn1((!(true || true))), F2: 2147483647}
                                    v30 = v2
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    var v51 int64 = (^(v2))
                                    _ = v50
                                    _ = v51
                                case 127:
                                    var v52 int = 4294967295
                                    _ = v52
                                case -594:
                                    v1 += (v1 + v1)
                                default:
                                    switch ("foo bar" + "hello") {
                                    case "a":
                                        var v53 int = (fi32 - 2147483647)
                                        _ = v53
                                    case "":
                                        v30 -= (v2 ^ -171)
                                    default:
                                        v54 := GenS4{F0: 32767, F1: 8}
                                        _ = v54
                                    }
                                }
                                v55 := GenS8{F0: ("a" + ""), F1: int8(v2), F2: v3}
                                v4 = append(v4, (v5 + (v5 + v5)))
                                _ = v55
                            }
                            var v56 int64 = v2
                            _ = v56
                        }
                        var v57 uint64 = 193
                        if true {
                            fmt.Println(v5, v2, v57)
                            if (!(!true)) {
                                if (true && false) {
                                    v58 := GenS3{F0: "\n\t", F1: true, F2: (v5 >> ((v3) & 63))}
                                    _ = v58
                                } else {
                                    v59 := GenS1{F0: (^(v57)), F1: 42, F2: 0}
                                    v57 = v57
                                    var v60 float32 = 3.14159
                                    _ = v59
                                    _ = v60
                                }
                                var v61 string = "foo bar"
                                _ = v61
                            }
                            var v62 float32 = float32(1e10)
                            v5 = 2147483648
                            _ = v62
                        } else {
                            v63 := GenS1{F0: genFn0(v1), F1: 42, F2: -229}
                            if len(v4) > 0 {
                                v4[0] = v5
                            }
                            _ = v63
                        }
                        _ = fi37
                        _ = v57
                    }
                }
            }
            v64 := GenS2{F0: -209}
            _ = fi32
            _ = v64
        }
    } else {
        v65 := GenS5{F0: 0, F1: 0}
        if true {
            for fi66 := 0; fi66 < 8; fi66++ {
                fmt.Println(v5)
                switch ((("" + "") + ("a" + ("" + "hello"))) + ("" + "foo bar")) {
                case "hello":
                    for fi67 := 0; fi67 < 8; fi67++ {
                        var v68 int16 = 255
                        _ = fi67
                        _ = v68
                    }
                }
                if len(v4) > 0 {
                    v4[0] = 65535
                }
                _ = fi66
            }
            v4 = append(v4, (v5 ^ uint32(v3)))
            if len(v4) > 0 {
                v4[0] = (^(v5))
            }
        }
        if true {
            for fi69 := 0; fi69 < 12; fi69++ {
                for fi70 := 0; fi70 < 12; fi70++ {
                    if ((true || (-2147483648 >= 0)) && (((float32(0.1) / (3.14159 - ((0.0 * -2.25) * 1.5))) + 0.0) < 1.5)) {
                        v5 -= v5
                    }
                    var v71 uint32 = 42
                    v72 := GenS1{F0: genFn0(v1), F1: 1, F2: 537}
                    _ = fi70
                    _ = v71
                    _ = v72
                }
                for fi73 := 0; fi73 < 10; fi73++ {
                    v1 -= 100
                    _ = fi73
                }
                _ = fi69
            }
        }
        _ = v65
    }
    v74 := GenS9{F0: 42, F1: 131, F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v74
}
