package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int16; F1 string; F2 bool }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint64; F1 uint16 }
type GenS10 struct { F0 int16 }
func genFn0() int64 {
    v6 := GenS1{F0: 127, F1: -182}
    var v7 int8 = 42
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 uint, p1 int) uint64 {
    v8 := GenS2{F0: -457}
    p0 = p0
    for fi9 := 0; fi9 < 6; fi9++ {
        var v10 int64 = 138
        _ = fi9
        _ = v10
    }
    v11 := GenS2{F0: 1}
    _ = p0
    _ = p1
    _ = v11
    _ = v8
    return 577
}

func main() {
    var v1 float64 = 3.14159
    var v2 float64 = v1
    var v3 int16 = -1
    var v4 uint32 = 65535
    v5 := []string{"", "", "hello"}
    var v12 uint64 = genFn1(135, 9223372036854775807)
    v2 = (((v2 / (((1.5 + -2.25) + v2) + v2)) * v1) / 1e10)
    var v13 uint16 = 42
    for fi14 := 0; fi14 < 17; fi14++ {
        _ = len(v5)
        for fi15 := 0; fi15 < 3; fi15++ {
            v16 := GenS1{F0: 177, F1: genFn0()}
            v17 := GenS2{F0: 42}
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = fi14
    }
    if false {
        for fi18 := 0; fi18 < 8; fi18++ {
            if (!(!true)) {
                if len(v5) > 0 {
                    v5[0] = ("\n\t" + "hello")
                }
            }
            for fi19 := 0; fi19 < 7; fi19++ {
                fmt.Println(v12, v3)
                _ = fi19
            }
            _ = len(v5)
            _ = fi18
        }
        for fi20 := 0; fi20 < 4; fi20++ {
            v2 -= ((-2.25 / 1.0) - 0.1)
            if true {
                if (1 >= 1) {
                    var v21 int64 = 1
                    fmt.Println(v3, fi20)
                    _ = v21
                } else {
                    switch (("foo bar" + "hello") + "foo bar") {
                    case "":
                        switch ("hello" + "\n\t") {
                        case "foo bar":
                            for fi22 := 0; fi22 < 10; fi22++ {
                                v23 := GenS3{F0: 2147483647, F1: 42}
                                _ = fi22
                                _ = v23
                            }
                        }
                    case "foo bar":
                        var v24 uint64 = (v12 | v12)
                        _ = v24
                    case "\n\t":
                        var v25 int32 = -1
                        _ = v25
                    default:
                        if (true || (int32(fi20) > -708)) {
                            switch "foo bar" {
                            case "":
                                switch ("\n\t" + ("" + "foo bar")) {
                                case "hello":
                                    switch "" {
                                    case "hello":
                                        if true {
                                            if false {
                                                var v26 float32 = 0.0
                                                var v27 int8 = 1
                                                if false {
                                                    fmt.Println(v2, fi20, v4)
                                                    for fi28 := 0; fi28 < 6; fi28++ {
                                                        v29 := GenS4{F0: 186}
                                                        fmt.Println(v13)
                                                        if (-371 >= -1) {
                                                            var v30 bool = (!true)
                                                            _ = v30
                                                        } else {
                                                            if (v27 < (v27 << ((416) & 63))) {
                                                                if len(v5) > 0 {
                                                                    v5[0] = (("foo bar" + "") + ("" + ""))
                                                                }
                                                            } else {
                                                                if ((^(v12)) != genFn1(258, (-(fi20)))) {
                                                                    v13 = (^(v13))
                                                                } else {
                                                                    v3 = (v3 & v3)
                                                                    switch ("hello" + (("a" + ("foo bar" + "\n\t")) + "hello")) {
                                                                    case "a":
                                                                        if len(v5) >= 2 {
                                                                            v5 = v5[1:len(v5)]
                                                                        }
                                                                    case "\n\t":
                                                                        fmt.Println(fi28, v4)
                                                                        v12 += genFn1(2, (fi28 << ((2) & 63)))
                                                                    default:
                                                                        switch (fi28 >> ((uint(v13)) & 63)) {
                                                                        case 0:
                                                                            var v31 uint16 = (v13 - uint16(v12))
                                                                            _ = v31
                                                                        }
                                                                    }
                                                                }
                                                                v27 += v27
                                                                v27 = (v27 * v27)
                                                            }
                                                            v4 = v4
                                                            var v32 uint = 255
                                                            _ = v32
                                                        }
                                                        _ = fi28
                                                        _ = v29
                                                    }
                                                }
                                                _ = v26
                                                _ = v27
                                            }
                                        }
                                    case "\n\t":
                                        var v33 int8 = 0
                                        _ = v33
                                    }
                                case "\n\t":
                                    var v34 uint = 1
                                    _ = v34
                                case "":
                                    var v35 uint = 0
                                    _ = v35
                                default:
                                    for fi36 := 0; fi36 < 13; fi36++ {
                                        var v37 uint64 = v12
                                        var v38 uint64 = (v37 / ((genFn1(437, (fi36 / ((fi36) | 1)))) | 1))
                                        switch -475 {
                                        case 267:
                                            v13 = (v13 % ((v13) | 1))
                                        case -2147483648:
                                            var v39 uint64 = uint64(v4)
                                            var v40 uint16 = 0
                                            _ = v39
                                            _ = v40
                                        }
                                        _ = fi36
                                        _ = v37
                                        _ = v38
                                    }
                                }
                            }
                        } else {
                            var v41 int16 = -32768
                            for fi42 := 0; fi42 < 1; fi42++ {
                                switch -2147483648 {
                                case 9223372036854775807:
                                    v43 := GenS5{F0: 1, F1: "hello", F2: true}
                                    _ = v43
                                case 32767:
                                    if (((!false) && ((fi42 == (fi20 | (fi42 - (fi20 / ((int(v13)) | 1))))) && false)) && true) {
                                        v12 = (^(v12))
                                        if ((true && ("" > "foo bar")) || (v4 >= (^(v4)))) {
                                            v4 = (^(v4))
                                            if len(v5) > 0 {
                                                v5[0] = "foo bar"
                                            }
                                        } else {
                                            if ((v12 == (v12 / ((v12) | 1))) || (!(((v1 * float64(0.0)) - v1) < 3.14159))) {
                                                var v44 string = "\n\t"
                                                _ = v44
                                            }
                                            v41 -= (v41 << ((4294967295) & 63))
                                        }
                                    }
                                case -192:
                                    var v45 uint16 = (v13 & v13)
                                    _ = v45
                                }
                                v46 := GenS1{F0: 2, F1: -549}
                                switch "foo bar" {
                                case "a":
                                    if false {
                                        switch ((("foo bar" + "") + "a") + ("" + "")) {
                                        case "\n\t":
                                            if ((320 <= (v4 & v4)) && false) {
                                                fmt.Println(v2, v4)
                                                switch "" {
                                                case "foo bar":
                                                    var v47 int64 = 0
                                                    _ = v47
                                                case "\n\t":
                                                    v48 := GenS6{F0: -1}
                                                    _ = v48
                                                }
                                                if false {
                                                    v49 := GenS4{F0: 127}
                                                    _ = v49
                                                }
                                            } else {
                                                if (1 < 945) {
                                                    v2 -= v1
                                                    v50 := GenS4{F0: 161}
                                                    var v51 int32 = 42
                                                    v52 := GenS1{F0: uint8(v4), F1: 0}
                                                    var v53 uint64 = (^(v12))
                                                    _ = len(v5)
                                                    var v54 uint = 68
                                                    _ = v50
                                                    _ = v51
                                                    _ = v52
                                                    _ = v53
                                                    _ = v54
                                                }
                                                v55 := GenS7{F0: ((3.14159 != float32(float64(-2.25))) || true)}
                                                _ = v55
                                            }
                                        case "":
                                            var v56 float32 = 1e10
                                            _ = v56
                                        default:
                                            v2 = ((3.14159 - ((v1 - v2) - (3.14159 * v2))) * v1)
                                        }
                                        v1 -= -2.25
                                        var v57 int = fi42
                                        _ = v57
                                    }
                                case "\n\t":
                                    v4 += v4
                                    if false {
                                        v41 -= 32767
                                        v5 = append(v5, ("\n\t" + ("\n\t" + "foo bar")))
                                        v3 = -798
                                    } else {
                                        v3 += 1
                                        if (634 > v41) {
                                            v3 = 542
                                            if (0 >= genFn0()) {
                                                fmt.Println(v3)
                                                if false {
                                                    v41 -= (v3 | -1)
                                                    v58 := GenS8{F0: 1}
                                                    _ = len(v5)
                                                    _ = v58
                                                }
                                                if true {
                                                    v13 = 552
                                                    v1 -= v1
                                                }
                                            } else {
                                                v3 = 0
                                                v41 -= (v41 % ((722) | 1))
                                            }
                                        } else {
                                            _ = len(v5)
                                        }
                                        var v59 uint = 2
                                        _ = v59
                                    }
                                }
                                _ = fi42
                                _ = v46
                            }
                            _ = v41
                        }
                    }
                }
            } else {
                v5 = append(v5, "\n\t")
                fmt.Println(v3)
                v5 = append(v5, "\n\t")
            }
            var v60 uint32 = (v4 * v4)
            _ = fi20
            _ = v60
        }
        v2 += (float64(-2.25) + float64(0.0))
    } else {
        v4 += 812
        v1 = (((0.0 - float64(0.1)) / 1.5) * v1)
    }
    switch 844 {
    case 2147483647:
        switch 42 {
        case -2147483648:
            var v61 int32 = 0
            _ = v61
        default:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
    default:
        v62 := GenS9{F0: v12, F1: (v13 ^ v13)}
        _ = v62
    }
    v12 += 18446744073709551615
    v63 := GenS10{F0: (v3 - (v3 + v3))}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v63
}
