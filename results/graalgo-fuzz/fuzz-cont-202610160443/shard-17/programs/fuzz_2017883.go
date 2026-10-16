package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int16; F1 uint8 }
type GenS5 struct { F0 uint8; F1 uint16; F2 uint64 }
type GenS6 struct { F0 int32; F1 uint; F2 int }
func genFn0() uint {
    var v7 uint64 = 2
    switch "\n\t" {
    case "":
        fmt.Println(v7)
    default:
        for fi8 := 0; fi8 < 1; fi8++ {
            v9 := []int{(fi8 % ((42) | 1))}
            _ = fi8
            _ = v9
        }
    }
    var v10 string = ("hello" + "foo bar")
    _ = v10
    _ = v7
    return 0
}
func genFn1(p0 uint64, p1 bool, p2 uint16) string {
    p1 = ((p2 | 2) > 147)
    p0 -= p0
    v11 := []int{-87, 1, 2147483647}
    v11 = append(v11, 127)
    switch "hello" {
    case "":
        var v12 bool = (!p1)
        _ = v12
    case "a":
        var v13 uint16 = 2
        _ = v13
    default:
        for fi14 := 0; fi14 < 13; fi14++ {
            switch "hello" {
            case "a":
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
            }
            _ = fi14
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return ""
}

func main() {
    var v1 uint16 = 800
    var v2 uint = 0
    var v3 int = 9223372036854775807
    var v4 uint64 = uint64(v2)
    v5 := make(map[string]bool)
    v6 := make(map[string]int)
    fmt.Println(v2)
    for fi15 := 0; fi15 < 13; fi15++ {
        for fi16 := 0; fi16 < 13; fi16++ {
            var v17 uint8 = 90
            _ = fi16
            _ = v17
        }
        if ("foo bar" <= "\n\t") {
            for fi18 := 0; fi18 < 12; fi18++ {
                delete(v6, "foo bar")
                var v19 bool = true
                v5["foo bar"] = (-712 < -2147483648)
                _ = fi18
                _ = v19
            }
        }
        v1 = (v1 / ((v1) | 1))
        _ = fi15
    }
    for fi20 := 0; fi20 < 4; fi20++ {
        v2 += (v2 | 594)
        for fi21 := 0; fi21 < 17; fi21++ {
            switch (v3 + (fi20 - fi20)) {
            case 1:
                if (fi20 < (fi20 ^ fi20)) {
                    v4 += uint64(v2)
                }
            case -9223372036854775808:
                v22 := GenS1{F0: ("hello" + ("a" + (("foo bar" + (("foo bar" + ("\n\t" + "\n\t")) + "")) + "a")))}
                _ = v22
            case -1:
                var v23 int32 = 0
                _ = v23
            default:
                if (v2 > genFn0()) {
                    var v24 int16 = -51
                    v25 := GenS1{F0: ""}
                    _ = v24
                    _ = v25
                } else {
                    v26 := GenS1{F0: "a"}
                    switch "hello" {
                    case "foo bar":
                        v27 := v6[("a" + ("" + ""))]
                        _ = v27
                    }
                    _ = v26
                }
            }
            fmt.Println(v4, v4, v1)
            v2 += (v2 / ((genFn0()) | 1))
            _ = fi21
        }
        _ = fi20
    }
    if false {
        if true {
            v3 = v3
            v28 := []int{(v3 & v3), v3}
            _ = v28
        }
        var v29 string = (("\n\t" + "\n\t") + "foo bar")
        _ = v29
    }
    if (!false) {
        v4 = v4
    }
    for fi30 := 0; fi30 < 8; fi30++ {
        v31 := GenS1{F0: ""}
        fmt.Println(v4, v2, v3)
        _ = fi30
        _ = v31
    }
    if (float32(((-2.25 * 0.0) * 1.5)) == ((float32(-2.25) * (float32(((1.5 / ((((-2.25 - (float64((3.14159 - 1.5)) - (0.0 + 0.1))) * (1.5 - 0.1)) * -2.25) / 3.14159)) - 1.5)) + 0.0)) * float32(-2.25))) {
        v32 := v5[("hello" + "\n\t")]
        _ = v32
    }
    fmt.Println(v4)
    var v33 uint32 = 0
    v4 += 18446744073709551615
    if (42 != 32767) {
        v1 += uint16(v33)
        v2 = (v2 + (v2 * genFn0()))
    } else {
        v3 = (v3 - (-(v3)))
        v1 += v1
    }
    if false {
        v34 := []int{(v3 >> ((v2) & 63)), 42, (v3 - v3)}
        for fi35 := 0; fi35 < 19; fi35++ {
            _ = len(v5)
            var v36 float64 = 1.5
            _ = fi35
            _ = v36
        }
        _ = v34
    } else {
        var v37 uint32 = 2147483648
        var v38 int = (v3 % ((-331) | 1))
        if false {
            for fi39 := 0; fi39 < 10; fi39++ {
                if (true || true) {
                    if false {
                        v40 := GenS1{F0: ("" + "hello")}
                        var v41 int32 = 713
                        _ = v40
                        _ = v41
                    }
                    v42 := v5[("a" + "a")]
                    _ = v42
                } else {
                    for fi43 := 0; fi43 < 8; fi43++ {
                        if true {
                            v44 := GenS1{F0: "a"}
                            if true {
                                v45 := GenS2{F0: "hello"}
                                _ = v45
                            }
                            v38 = int(v1)
                            _ = v44
                        } else {
                            if (false && ((float64(0.1) >= 1e10) && (863 == 1))) {
                                v1 -= uint16(v4)
                                v3 = v38
                                fmt.Println(v1, v2, v33)
                            }
                            if (-128 > -1) {
                                var v46 uint8 = 1
                                if false {
                                    v38 = fi39
                                    var v47 uint64 = (v4 * 2)
                                    delete(v6, "\n\t")
                                    v48 := GenS2{F0: "a"}
                                    v49 := []int{}
                                    _ = v47
                                    _ = v48
                                    _ = v49
                                }
                                _ = v46
                            } else {
                                var v50 uint16 = (v1 | (^(v1)))
                                v4 -= (^(v4))
                                _ = v50
                            }
                        }
                        v37 += uint32(v1)
                        v51 := GenS2{F0: "foo bar"}
                        _ = fi43
                        _ = v51
                    }
                    v1 -= (v1 | (v1 ^ v1))
                    delete(v6, ("\n\t" + "foo bar"))
                }
                var v52 uint32 = v37
                if ("a" != ("hello" + "")) {
                    v4 = 18446744073709551615
                } else {
                    switch (v3 & -459) {
                    case 127:
                        switch (^(v38)) {
                        case 1:
                            var v53 uint64 = (v4 & 0)
                            _ = v53
                        case 32767:
                            switch 50 {
                            case 125:
                                var v54 int8 = int8(v38)
                                _ = v54
                            case -155:
                                var v55 int64 = -1
                                _ = v55
                            case 587:
                                v56 := GenS1{F0: (("hello" + "") + ("\n\t" + "\n\t"))}
                                _ = v56
                            default:
                                fmt.Println(v37)
                            }
                        case 128:
                            v57 := GenS3{F0: false}
                            _ = v57
                        default:
                            _ = len(v6)
                        }
                    case 947:
                        if (1000 > 342) {
                            var v58 uint8 = uint8(v2)
                            switch ("foo bar" + "foo bar") {
                            case "a":
                                var v59 uint32 = uint32(v4)
                                _ = v59
                            case "hello":
                                var v60 float32 = (((float32(3.14159) - 0.0) / (1.5 - 3.14159)) / ((((1e10 - -2.25) + 1.5) - 1e10) + float32(float64((3.14159 * -2.25)))))
                                _ = v60
                            }
                            _ = v58
                        } else {
                            var v61 uint = (v2 ^ (v2 | v2))
                            var v62 uint8 = 255
                            if ("a" >= ("foo bar" + "hello")) {
                                delete(v5, "hello")
                                v63 := GenS4{F0: 42, F1: 227}
                                _ = v63
                            } else {
                                var v64 uint = 0
                                v1 = v1
                                _ = v64
                            }
                            _ = v61
                            _ = v62
                        }
                    case 9223372036854775807:
                        _ = len(v5)
                    default:
                        switch "" {
                        case "hello":
                            v37 -= (v37 / ((2147483648) | 1))
                        case "foo bar":
                            v65 := []int{fi39}
                            _ = v65
                        case "":
                            v3 = (-(v3))
                        }
                    }
                    for fi66 := 0; fi66 < 2; fi66++ {
                        v52 += 65535
                        _ = fi66
                    }
                    v37 = 54
                }
                _ = fi39
                _ = v52
            }
            for fi67 := 0; fi67 < 12; fi67++ {
                v68 := GenS3{F0: true}
                if false {
                    for fi69 := 0; fi69 < 6; fi69++ {
                        if false {
                            switch "a" {
                            case "\n\t":
                                var v70 int8 = -101
                                _ = v70
                            }
                        }
                        _ = fi69
                    }
                }
                var v71 int = v38
                _ = fi67
                _ = v68
                _ = v71
            }
            v72 := GenS1{F0: "a"}
            _ = v72
        } else {
            v73 := GenS4{F0: 1000, F1: 127}
            v74 := GenS5{F0: 210, F1: v1, F2: (v4 << ((v2) & 63))}
            v75 := GenS6{F0: -2147483648, F1: uint(v37), F2: (v3 - (-(v38)))}
            _ = v73
            _ = v74
            _ = v75
        }
        _ = v37
        _ = v38
    }
    fmt.Println(v2, v3)
    var v76 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v76
}
