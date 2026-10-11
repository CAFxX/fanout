package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 uint16; F2 bool }
type GenS3 struct { F0 uint32; F1 bool; F2 uint8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int16; F1 int; F2 int16 }
type GenS6 struct { F0 uint8; F1 uint8 }
type GenS7 struct { F0 int32; F1 uint }
func genFn0() int16 {
    var v7 uint8 = 1
    for fi8 := 0; fi8 < 18; fi8++ {
        if true {
            var v9 float32 = 1e10
            _ = v9
        }
        var v10 int32 = -466
        _ = fi8
        _ = v10
    }
    _ = v7
    return 32767
}
func genFn1(p0 uint32) bool {
    if false {
        fmt.Println(p0)
        p0 = p0
        v11 := []int{-868, 947, -84}
        _ = v11
    }
    if true {
        if true {
            v12 := GenS1{F0: 1000}
            _ = v12
        }
        fmt.Println(p0)
    } else {
        var v13 uint8 = 197
        switch "" {
        case "hello":
            for fi14 := 0; fi14 < 12; fi14++ {
                switch 42 {
                case -9223372036854775808:
                    switch -1 {
                    case 2147483647:
                        fmt.Println(v13)
                    default:
                        for fi15 := 0; fi15 < 3; fi15++ {
                            if (100 == 127) {
                                v16 := []int{(fi15 * (fi15 >> ((255) & 63))), fi14}
                                _ = v16
                            }
                            p0 += p0
                            v17 := []int{32767, fi15, fi14, fi15}
                            p0 = (p0 / ((p0) | 1))
                            p0 = (p0 % ((p0) | 1))
                            var v18 float64 = ((0.1 * (((float64(((-2.25 / (-2.25 + float32(1.5))) * (3.14159 - 0.0))) - 0.0) / (1.5 / -2.25)) * 1.5)) - 3.14159)
                            fmt.Println(fi15, v18)
                            switch ((("\n\t" + "hello") + ("a" + ("hello" + ("\n\t" + ("a" + "a"))))) + (("hello" + "foo bar") + "a")) {
                            case "":
                                var v19 bool = (("" + (((("a" + "") + "a") + "hello") + "")) < ("a" + "\n\t"))
                                _ = v19
                            case "\n\t":
                                v17 = append(v17, -346)
                            default:
                                v20 := GenS2{F0: fi14, F1: 1, F2: (5 > 100)}
                                _ = v20
                            }
                            _ = len(v17)
                            _ = fi15
                            _ = v17
                            _ = v18
                        }
                    }
                case 32767:
                    switch (("" + ("hello" + "a")) + "hello") {
                    case "\n\t":
                        if (("\n\t" + "foo bar") < (("" + ("a" + "")) + ("foo bar" + ("\n\t" + "")))) {
                            if ((((("" + (("a" + "") + ("hello" + ""))) + "\n\t") + "hello") > "foo bar") && false) {
                                v13 = 2
                            }
                        } else {
                            v21 := []int{(fi14 + 32768), -578}
                            var v22 uint8 = (v13 | (v13 | (v13 | v13)))
                            _ = v21
                            _ = v22
                        }
                    }
                }
                if (false && false) {
                    v23 := GenS2{F0: (^(fi14)), F1: 65535, F2: false}
                    v13 = (v13 >> ((4294967295) & 63))
                    _ = v23
                }
                _ = fi14
            }
        default:
            v13 -= (v13 % ((1) | 1))
        }
        var v24 int64 = 42
        _ = v13
        _ = v24
    }
    _ = p0
    return (false && false)
}

func main() {
    var v1 int16 = 1000
    var v2 uint = 42
    var v3 float32 = -2.25
    v4 := make(map[string]bool)
    var v5 float64 = 0.0
    v6 := make(map[uint32]int)
    v25 := v4["foo bar"]
    v25 = (!true)
    switch 42 {
    case -164:
        v25 = (!false)
    case -9223372036854775808:
        switch -1 {
        case 2147483647:
            var v26 string = "hello"
            _ = v26
        case 128:
            if v25 {
                for fi27 := 0; fi27 < 11; fi27++ {
                    var v28 int16 = v1
                    _ = fi27
                    _ = v28
                }
                v2 -= (^(v2))
                var v29 uint32 = 2147483647
                _ = v29
            } else {
                if v25 {
                    v5 -= ((-2.25 + v5) / (v5 - float64(v3)))
                    var v30 uint32 = 65535
                    var v31 int64 = 137
                    _ = v30
                    _ = v31
                } else {
                    var v32 int32 = 820
                    v33 := GenS1{F0: v1}
                    _ = v32
                    _ = v33
                }
                if false {
                    v34 := []int{}
                    _ = v34
                } else {
                    _ = len(v4)
                    _ = len(v4)
                }
                v3 = 0.1
            }
        default:
            for fi35 := 0; fi35 < 4; fi35++ {
                switch (("a" + "foo bar") + "") {
                case "foo bar":
                    v5 = v5
                default:
                    fmt.Println(v2, v1)
                }
                _ = fi35
            }
        }
    case -344:
        for fi36 := 0; fi36 < 14; fi36++ {
            switch "a" {
            case "":
                if true {
                    if (686 < 2147483647) {
                        v37 := GenS3{F0: 2147483648, F1: (false || v25), F2: 158}
                        _ = v37
                    }
                    var v38 string = ""
                    _ = v38
                }
            default:
                v39 := GenS1{F0: v1}
                _ = v39
            }
            delete(v4, "\n\t")
            v1 = 42
            _ = fi36
        }
    }
    for fi40 := 0; fi40 < 5; fi40++ {
        switch "a" {
        case "":
            switch (fi40 & fi40) {
            case -198:
                if ((!(0.1 != 0.1)) || v25) {
                    fmt.Println(v2, v3, v5)
                    var v41 uint64 = 18446744073709551615
                    _ = v41
                }
            case 9223372036854775807:
                var v42 int16 = (v1 | (v1 << (((v2 & v2)) & 63)))
                _ = v42
            case 0:
                v5 = float64(v3)
            default:
                switch (fi40 % (((fi40 * 42)) | 1)) {
                case 4294967295:
                    delete(v6, 1)
                case 42:
                    if false {
                        var v43 int16 = v1
                        var v44 uint = v2
                        if (!(2147483648 == 4294967295)) {
                            v45 := GenS4{F0: uint32(v2)}
                            var v46 uint16 = 32767
                            _ = v45
                            _ = v46
                        }
                        _ = v43
                        _ = v44
                    } else {
                        if (v25 && v25) {
                            v25 = v25
                        }
                        v47 := GenS3{F0: 65535, F1: (((v3 * v3) + (v3 - v3)) <= v3), F2: 42}
                        _ = v47
                    }
                }
            }
        case "hello":
            if (708 <= 65535) {
                fmt.Println(v1, v3, v25)
            } else {
                v48 := GenS5{F0: (v1 + genFn0()), F1: 128, F2: genFn0()}
                var v49 bool = (-9223372036854775808 >= -1)
                _ = v48
                _ = v49
            }
        case "foo bar":
            v25 = (983 == 42)
        }
        var v50 int8 = 100
        var v51 uint32 = 42
        _ = fi40
        _ = v50
        _ = v51
    }
    v52 := GenS4{F0: 42}
    if (2 >= 338) {
        for fi53 := 0; fi53 < 13; fi53++ {
            var v54 uint = v2
            v3 += ((float32(v5) + float32(0.0)) + 0.0)
            var v55 int = fi53
            _ = fi53
            _ = v54
            _ = v55
        }
        var v56 int16 = (v1 / (((v1 / ((32767) | 1))) | 1))
        if v25 {
            var v57 uint = (^(v2))
            var v58 int8 = 0
            switch 32767 {
            case 601:
                _ = len(v4)
            case -318:
                for fi59 := 0; fi59 < 18; fi59++ {
                    v60 := []int{fi59, (fi59 + -460)}
                    _ = fi59
                    _ = v60
                }
            default:
                var v61 uint = (v57 / (((v2 << ((v2) & 63))) | 1))
                _ = v61
            }
            _ = v57
            _ = v58
        } else {
            var v62 uint8 = 127
            v25 = false
            _ = v62
        }
        _ = v56
    } else {
        v63 := GenS2{F0: 128, F1: 626, F2: (true || v25)}
        _ = v63
    }
    if v25 {
        var v64 uint8 = 75
        v65 := v4[(("" + ("a" + "\n\t")) + "hello")]
        _ = v64
        _ = v65
    } else {
        for fi66 := 0; fi66 < 6; fi66++ {
            v67 := GenS6{F0: 42, F1: uint8(v2)}
            v2 = (v2 / ((65535) | 1))
            _ = fi66
            _ = v67
        }
        if ((v5 * v5) == v5) {
            v3 += 3.14159
        } else {
            v5 = (((v5 - (float64(v3) + -2.25)) + v5) * ((v5 - 1.5) / 1e10))
            switch "foo bar" {
            case "hello":
                switch "hello" {
                case "a":
                    if ("" >= ("\n\t" + "foo bar")) {
                        fmt.Println(v2, v3, v3)
                    } else {
                        v5 += v5
                        v1 = (^(v1))
                        v68 := GenS7{F0: -107, F1: (^(v2))}
                        _ = v68
                    }
                default:
                    v25 = v25
                }
            case "foo bar":
                var v69 uint32 = 773
                _ = v69
            case "a":
                var v70 uint16 = 42
                _ = v70
            default:
                var v71 uint64 = 42
                _ = v71
            }
        }
    }
    v1 = 1
    for fi72 := 0; fi72 < 20; fi72++ {
        if true {
            if false {
                v3 -= float32((((1.5 / v5) / v5) / (v5 + (-2.25 + (0.0 / v5)))))
                delete(v6, 4294967295)
            }
            v73 := GenS5{F0: (v1 << ((v2) & 63)), F1: fi72, F2: (v1 / ((-1) | 1))}
            _ = v73
        } else {
            var v74 string = "hello"
            _ = v74
        }
        if v25 {
            fmt.Println(v1)
            if (((!(!v25)) || v25) || (!v25)) {
                if v25 {
                    for fi75 := 0; fi75 < 17; fi75++ {
                        v1 = 468
                        var v76 float64 = v5
                        _ = fi75
                        _ = v76
                    }
                    fmt.Println(v2, v5)
                } else {
                    for fi77 := 0; fi77 < 12; fi77++ {
                        v25 = v25
                        _ = fi77
                    }
                    v5 += 0.0
                }
                var v78 int = 32768
                _ = v78
            } else {
                fmt.Println(v5, fi72)
            }
            var v79 float32 = v3
            _ = v79
        }
        var v80 float64 = v5
        _ = fi72
        _ = v80
    }
    var v81 uint8 = 128
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v6
    _ = v81
}
