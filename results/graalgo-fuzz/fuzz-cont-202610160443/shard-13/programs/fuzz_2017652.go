package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 uint64; F1 bool; F2 uint64 }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 int8; F1 uint64; F2 string }
type GenS5 struct { F0 uint64; F1 uint32; F2 int }
type GenS6 struct { F0 int16; F1 bool; F2 uint32 }
type GenS7 struct { F0 int16; F1 uint64; F2 int32 }
type GenS8 struct { F0 uint; F1 uint; F2 uint8 }
type GenS9 struct { F0 uint64; F1 int16; F2 uint32 }
type GenS10 struct { F0 uint8; F1 int }
type GenS11 struct { F0 uint; F1 uint; F2 uint8 }
type GenS12 struct { F0 int16 }
func genFn0(p0 bool, p1 int64) uint64 {
    if p0 {
        switch "" {
        case "":
            if true {
                switch "foo bar" {
                case "foo bar":
                    for fi7 := 0; fi7 < 8; fi7++ {
                        for fi8 := 0; fi8 < 11; fi8++ {
                            v9 := []int{(fi7 / ((fi8) | 1))}
                            var v10 uint64 = 1
                            var v11 string = ("foo bar" + (("" + "\n\t") + ("a" + ("foo bar" + ((("foo bar" + ((("" + "hello") + "a") + ("hello" + ("\n\t" + ("\n\t" + ("" + (("foo bar" + (("" + "") + "")) + ("a" + (("hello" + "a") + "hello"))))))))) + "foo bar") + "a")))))
                            _ = fi8
                            _ = v10
                            _ = v11
                            _ = v9
                        }
                        _ = fi7
                    }
                case "a":
                    var v12 int16 = 1000
                    _ = v12
                case "":
                    p1 = p1
                }
            }
        case "a":
            switch -793 {
            case -634:
                fmt.Println(p1)
            case 374:
                var v13 int8 = 4
                _ = v13
            }
        default:
            fmt.Println(p0, p0)
        }
    }
    var v14 int = -9223372036854775808
    for fi15 := 0; fi15 < 14; fi15++ {
        var v16 uint64 = 1
        v14 = (v14 % (((fi15 << ((65535) & 63))) | 1))
        var v17 uint16 = 32767
        _ = fi15
        _ = v16
        _ = v17
    }
    v18 := GenS1{F0: 0, F1: "a"}
    v19 := GenS1{F0: 4294967295, F1: ""}
    _ = p0
    _ = p1
    _ = v14
    _ = v18
    _ = v19
    return 0
}

func main() {
    v1 := make(map[int]int)
    var v2 int32 = -2147483648
    v3 := make(map[uint32]uint32)
    v4 := []uint32{0, 65535, 2147483648, 42}
    var v5 bool = false
    v6 := make(map[int]int)
    fmt.Println(v5)
    if (!(-1 != v2)) {
        var v20 int32 = 0
        if (v5 && (v5 || true)) {
            switch 4294967295 {
            case 32767:
                fmt.Println(v5)
            }
        }
        _ = v20
    } else {
        switch ("foo bar" + "hello") {
        case "hello":
            var v21 float64 = (((1.5 * (3.14159 / -2.25)) * (1e10 + (0.0 / -2.25))) + (-2.25 + (1e10 * (0.0 - 3.14159))))
            _ = v21
        case "foo bar":
            switch "foo bar" {
            case "a":
                var v22 float64 = float64(1e10)
                _ = v22
            case "":
                delete(v6, -9223372036854775808)
            case "hello":
                var v23 bool = (v5 || (v2 != (^(v2))))
                _ = v23
            default:
                var v24 int = 42
                _ = v24
            }
        default:
            switch ("\n\t" + ((("a" + ("hello" + "\n\t")) + (("hello" + "hello") + ("" + ("foo bar" + ("" + ("foo bar" + ("" + ""))))))) + "foo bar")) {
            case "foo bar":
                v5 = true
            case "":
                fmt.Println(v5, v5)
            default:
                if (true && v5) {
                    v25 := GenS2{F0: 4294967295, F1: (!(0 <= 1)), F2: 0}
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v5 = (0.1 != 3.14159)
                    _ = v25
                }
            }
        }
        v26 := GenS2{F0: genFn0(true, -1), F1: (421 > 2), F2: 475}
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v26
    }
    v27 := GenS3{F0: 155, F1: 100}
    for fi28 := 0; fi28 < 2; fi28++ {
        var v29 int = (fi28 & fi28)
        v5 = (127 != 128)
        _ = fi28
        _ = v29
    }
    v30 := GenS1{F0: genFn0((!((-32768 <= 1000) && ("\n\t" <= "hello"))), -799), F1: "hello"}
    if (42 < 127) {
        v31 := GenS4{F0: -70, F1: 0, F2: ""}
        if (127 > -1) {
            if false {
                var v32 int = -1
                v33 := GenS2{F0: 18446744073709551615, F1: false, F2: 871}
                var v34 float64 = float64((3.14159 * 0.1))
                _ = v32
                _ = v33
                _ = v34
            } else {
                var v35 string = "\n\t"
                _ = v35
            }
            v5 = v5
            if ((((float32(0.1) / 1e10) - (((3.14159 * ((0.1 * 0.1) - 0.0)) / 1e10) + (1e10 * ((1.5 + 1.5) + 3.14159)))) < 1e10) && v5) {
                v2 = v2
                v5 = v5
            }
        }
        _ = v31
    } else {
        for fi36 := 0; fi36 < 14; fi36++ {
            v2 = 42
            _ = fi36
        }
        v37 := GenS5{F0: 0, F1: 42, F2: int(v2)}
        _ = v37
    }
    var v38 bool = v5
    fmt.Println(v5, v2)
    fmt.Println(v5)
    fmt.Println(v2, v2)
    switch ("a" + "foo bar") {
    case "\n\t":
        var v39 int32 = v2
        _ = v39
    case "":
        v40 := GenS2{F0: genFn0((0 != 42), -1), F1: false, F2: 4294967295}
        _ = v40
    default:
        switch (("foo bar" + "hello") + "a") {
        case "\n\t":
            if (v5 || true) {
                v38 = ((0.0 * 1e10) > 0.1)
                for fi41 := 0; fi41 < 9; fi41++ {
                    for fi42 := 0; fi42 < 10; fi42++ {
                        v43 := GenS6{F0: -1, F1: (v38 && (83 <= 756)), F2: 2147483647}
                        _ = fi42
                        _ = v43
                    }
                    _ = fi41
                }
                for fi44 := 0; fi44 < 17; fi44++ {
                    for fi45 := 0; fi45 < 6; fi45++ {
                        if true {
                            var v46 float32 = -2.25
                            _ = v46
                        } else {
                            v47 := GenS7{F0: 0, F1: 18446744073709551615, F2: (-(v2))}
                            switch "" {
                            case "":
                                _ = len(v4)
                            default:
                                fmt.Println(fi44)
                            }
                            fmt.Println(v5)
                            v48 := GenS8{F0: 682, F1: 818, F2: 0}
                            var v49 float32 = float32(1.5)
                            switch fi44 {
                            case 4294967295:
                                if v38 {
                                    v50 := GenS9{F0: genFn0((549 != 545), 9223372036854775807), F1: -896, F2: 0}
                                    if v5 {
                                        if true {
                                            var v51 uint64 = 42
                                            v52 := GenS5{F0: 4294967295, F1: 4294967295, F2: fi45}
                                            _ = v51
                                            _ = v52
                                        }
                                        var v53 int = 590
                                        _ = v53
                                    } else {
                                        v54 := GenS6{F0: -1, F1: ((fi45 * fi45) < 32768), F2: 0}
                                        if (255 != 2) {
                                            v55 := GenS10{F0: 128, F1: 42}
                                            _ = v55
                                        } else {
                                            v56 := GenS1{F0: 623, F1: ""}
                                            var v57 int32 = (v2 / (((v2 & v2)) | 1))
                                            var v58 float32 = v49
                                            _ = v56
                                            _ = v57
                                            _ = v58
                                        }
                                        switch (fi44 | 32767) {
                                        case 4294967295:
                                            switch ("" + ((("" + "\n\t") + "hello") + "foo bar")) {
                                            case "\n\t":
                                                v5 = (v5 || (-746 >= (^(fi44))))
                                            case "foo bar":
                                                v59 := GenS4{F0: 0, F1: 0, F2: ("foo bar" + "foo bar")}
                                                v60 := GenS4{F0: -1, F1: genFn0((!false), 177), F2: "foo bar"}
                                                var v61 float32 = v49
                                                _ = v59
                                                _ = v60
                                                _ = v61
                                            default:
                                                v62 := GenS5{F0: 42, F1: 1, F2: (^(fi45))}
                                                _ = v62
                                            }
                                        case 127:
                                            var v63 int64 = 997
                                            _ = v63
                                        }
                                        _ = v54
                                    }
                                    _ = v50
                                } else {
                                    v64 := GenS11{F0: 1, F1: 611, F2: 0}
                                    var v65 bool = (122 > 249)
                                    _ = v64
                                    _ = v65
                                }
                            }
                            fmt.Println(v2)
                            var v66 string = ("a" + "a")
                            v49 -= v49
                            if ((v66 + (("a" + v66) + v66)) < v66) {
                                var v67 int = (fi45 % (((fi44 - fi45)) | 1))
                                _ = v67
                            } else {
                                v49 += 1e10
                                var v68 int = fi45
                                _ = v68
                            }
                            _ = v47
                            _ = v48
                            _ = v49
                            _ = v66
                        }
                        _ = fi45
                    }
                    if v38 {
                        var v69 uint32 = 945
                        fmt.Println(v2, v69, v5)
                        _ = v69
                    } else {
                        switch ("\n\t" + "\n\t") {
                        case "foo bar":
                            var v70 uint8 = 128
                            _ = v70
                        }
                    }
                    _ = fi44
                }
            } else {
                var v71 bool = true
                fmt.Println(v5)
                _ = v71
            }
        default:
            switch "a" {
            case "":
                v72 := GenS7{F0: 1000, F1: 42, F2: v2}
                _ = v72
            case "\n\t":
                if (v38 && true) {
                    var v73 float32 = (float32((((1e10 + ((0.0 / float64((0.1 + 1.5))) - 1.5)) + (-2.25 + 3.14159)) / 3.14159)) - 0.1)
                    v73 = (((v73 / (v73 - (((0.0 * (v73 - (v73 + (((v73 - -2.25) + v73) * float32(1e10))))) / 1.5) - v73))) / float32(1e10)) + v73)
                    _ = v73
                }
            case "a":
                for fi74 := 0; fi74 < 8; fi74++ {
                    v75 := GenS12{F0: 1000}
                    delete(v6, fi74)
                    _ = fi74
                    _ = v75
                }
            default:
                fmt.Println(v5, v2)
            }
        }
    }
    var v76 uint = 0
    v76 += (v76 & (v76 * (^(v76))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v76
}
