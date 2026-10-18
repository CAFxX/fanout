package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16; F1 int }
type GenS5 struct { F0 int32; F1 int32; F2 uint64 }
type GenS6 struct { F0 int8; F1 int64; F2 int64 }
type GenS7 struct { F0 int64; F1 uint8; F2 int32 }
type GenS8 struct { F0 int64; F1 uint32 }
type GenS9 struct { F0 int32; F1 uint8; F2 uint64 }
type GenS10 struct { F0 uint16; F1 uint; F2 int16 }
type GenS11 struct { F0 int32; F1 bool; F2 string }
type GenS12 struct { F0 uint32; F1 uint16 }

func main() {
    v1 := []string{}
    var v2 int = 4294967295
    var v3 int32 = -2147483648
    var v4 float32 = (((0.1 / 1.0) / ((-2.25 * -2.25) + 0.1)) * ((0.1 - 1e10) / float32((-2.25 * -2.25))))
    var v5 uint64 = 142
    if ((1000 == -32768) || false) {
        var v6 uint16 = 0
        v7 := GenS1{F0: (v6 >> ((0) & 63)), F1: 0}
        v5 += (v5 - (v5 / ((v5) | 1)))
        _ = v6
        _ = v7
    } else {
        v2 = (v2 / ((42) | 1))
        for fi8 := 0; fi8 < 10; fi8++ {
            v9 := GenS1{F0: 1, F1: 0}
            _ = fi8
            _ = v9
        }
        _ = len(v1)
    }
    var v10 uint = 65535
    var v11 float32 = (v4 * v4)
    if (-693 >= 1000) {
        fmt.Println(v4, v3, v10)
        v5 += 42
        v12 := GenS1{F0: 2, F1: -1}
        _ = v12
    } else {
        if true {
            v13 := GenS1{F0: 388, F1: 127}
            _ = v13
        } else {
            fmt.Println(v10, v10)
            v4 = 0.0
            for fi14 := 0; fi14 < 13; fi14++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v15 := GenS1{F0: 1, F1: 100}
                if (false || ((v5 - (^(v5))) < v5)) {
                    var v16 int32 = (v3 ^ -87)
                    for fi17 := 0; fi17 < 4; fi17++ {
                        if false {
                            var v18 string = ""
                            _ = v18
                        }
                        if (1e10 == (-2.25 - float64(v4))) {
                            v4 = 0.1
                            var v19 uint32 = 127
                            var v20 int16 = 42
                            _ = v19
                            _ = v20
                        } else {
                            if (!false) {
                                switch 631 {
                                case 32767:
                                    fmt.Println(v3, fi17, v3)
                                case -1:
                                    v21 := GenS1{F0: 803, F1: -80}
                                    switch 1 {
                                    case 870:
                                        if true {
                                            var v22 string = "a"
                                            v16 = (v3 >> ((v10) & 63))
                                            var v23 float32 = v4
                                            _ = v22
                                            _ = v23
                                        } else {
                                            v1 = append(v1, "foo bar")
                                        }
                                    case 988:
                                        var v24 uint64 = 729
                                        _ = v24
                                    }
                                    if false {
                                        v1 = append(v1, "\n\t")
                                    }
                                    switch "\n\t" {
                                    case "\n\t":
                                        v2 = (fi14 - 128)
                                    case "hello":
                                        switch v2 {
                                        case -405:
                                            if (9223372036854775807 == 0) {
                                                if (-128 <= 1) {
                                                    var v25 float64 = 3.14159
                                                    _ = v25
                                                }
                                                v26 := GenS2{F0: 0}
                                                if false {
                                                    var v27 uint8 = 1
                                                    if len(v1) > 0 {
                                                        v1[0] = "foo bar"
                                                    }
                                                    var v28 int8 = -1
                                                    _ = v27
                                                    _ = v28
                                                }
                                                v29 := GenS3{F0: 42}
                                                var v30 int16 = 42
                                                _ = v26
                                                _ = v29
                                                _ = v30
                                            }
                                        case 196:
                                            fmt.Println(v2)
                                            var v31 int16 = 42
                                            fmt.Println(v4, v4)
                                            _ = v31
                                        case 0:
                                            if (892 > 42) {
                                                if (((true || false) || true) || ((-1 > 32767) && (0.1 > 1e10))) {
                                                    v32 := GenS3{F0: 2147483648}
                                                    fmt.Println(v10, v3)
                                                    v4 -= 0.1
                                                    var v33 uint8 = 255
                                                    _ = v32
                                                    _ = v33
                                                } else {
                                                    var v34 int64 = -1
                                                    _ = v34
                                                }
                                                if true {
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        v5 = v5
                                                    case "hello":
                                                        if ((true || (v2 >= (v2 / ((fi17) | 1)))) || (true || (true && true))) {
                                                            var v35 bool = ((645 >= v10) || false)
                                                            fmt.Println(v35)
                                                            v36 := GenS3{F0: 938}
                                                            _ = v35
                                                            _ = v36
                                                        }
                                                    case "foo bar":
                                                        v10 = (^(v10))
                                                    default:
                                                        var v37 string = ((("" + ("a" + ("hello" + "hello"))) + (((("\n\t" + "\n\t") + (("\n\t" + "") + "")) + "") + ("" + "foo bar"))) + ("foo bar" + "foo bar"))
                                                        _ = v37
                                                    }
                                                    v4 += v4
                                                } else {
                                                    if false {
                                                        if false {
                                                            fmt.Println(v11, fi14)
                                                        }
                                                    } else {
                                                        v38 := GenS4{F0: 0, F1: fi14}
                                                        if len(v1) > 0 {
                                                            v1[0] = "foo bar"
                                                        }
                                                        var v39 float32 = 0.0
                                                        _ = v38
                                                        _ = v39
                                                    }
                                                }
                                            }
                                        }
                                    case "foo bar":
                                        v40 := GenS4{F0: 32767, F1: 4294967295}
                                        _ = v40
                                    }
                                    v2 = 32767
                                    v1 = append(v1, "hello")
                                    v41 := GenS2{F0: 193}
                                    if true {
                                        _ = len(v1)
                                        var v42 int8 = 42
                                        v4 = v4
                                        _ = v42
                                    } else {
                                        v10 -= v10
                                    }
                                    _ = v21
                                    _ = v41
                                default:
                                    fmt.Println(v11, v10, fi17)
                                }
                                v43 := GenS5{F0: -973, F1: (v3 << (((^(v10))) & 63)), F2: (^(v5))}
                                _ = v43
                            } else {
                                fmt.Println(fi14, v4)
                                v5 += (v5 % (((^(v5))) | 1))
                                switch (fi17 + v2) {
                                case 128:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                case -9223372036854775808:
                                    if (false && false) {
                                        switch ("\n\t" + "hello") {
                                        case "\n\t":
                                            var v44 bool = (!(18446744073709551615 <= v5))
                                            _ = v44
                                        case "a":
                                            var v45 uint = (v10 % ((v10) | 1))
                                            _ = v45
                                        }
                                        v10 += v10
                                        if (2147483647 < 401) {
                                            fmt.Println(fi17, fi17)
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        } else {
                                            v11 -= 0.1
                                            switch (v2 / ((fi14) | 1)) {
                                            case 42:
                                                v1 = append(v1, "")
                                            case -2147483648:
                                                if ("" > "") {
                                                    if len(v1) > 0 {
                                                        v1[0] = "hello"
                                                    }
                                                    var v46 uint16 = 1
                                                    v47 := GenS6{F0: 0, F1: -628, F2: -9223372036854775808}
                                                    _ = v46
                                                    _ = v47
                                                } else {
                                                    v48 := GenS4{F0: 379, F1: (fi14 ^ fi17)}
                                                    var v49 float32 = float32(3.14159)
                                                    _ = v48
                                                    _ = v49
                                                }
                                            case 128:
                                                fmt.Println(v5, v2)
                                            default:
                                                switch (("hello" + "foo bar") + ((("foo bar" + "hello") + "") + "hello")) {
                                                case "hello":
                                                    v3 = int32(v5)
                                                case "foo bar":
                                                    v50 := GenS7{F0: 0, F1: 0, F2: 412}
                                                    _ = v50
                                                case "":
                                                    switch (fi17 / ((fi17) | 1)) {
                                                    case 32767:
                                                        v4 = (float32(((((1.5 + 1e10) + float64((v4 - float32((float64(v4) + 0.0))))) - float64((v11 - (v4 * v11)))) - (((0.0 + -2.25) * float64((v11 + v11))) / 1.0))) + 0.0)
                                                    case 2:
                                                        var v51 int64 = 42
                                                        _ = v51
                                                    case 0:
                                                        if true {
                                                            var v52 uint32 = 2147483648
                                                            var v53 uint = v10
                                                            _ = v52
                                                            _ = v53
                                                        }
                                                    default:
                                                        fmt.Println(v10, v16, v16)
                                                    }
                                                default:
                                                    v2 += (v2 << ((v10) & 63))
                                                }
                                            }
                                            fmt.Println(v11)
                                            switch fi17 {
                                            case 939:
                                                if ((0.1 < (((1.5 * 0.0) + 1.5) / float64(3.14159))) || true) {
                                                    if true {
                                                        if false {
                                                            if (!(v10 != (v10 | (^(v10))))) {
                                                                if (!true) {
                                                                    v10 = v10
                                                                } else {
                                                                    v54 := GenS1{F0: 2, F1: 0}
                                                                    v3 = (v3 & (v16 - (v16 & v3)))
                                                                    switch (v2 & v2) {
                                                                    case -718:
                                                                        var v55 bool = (128 != 255)
                                                                        _ = v55
                                                                    case 9223372036854775807:
                                                                        if (!true) {
                                                                            v4 = float32((3.14159 + 0.0))
                                                                            v2 -= (v2 / (((fi14 / ((128) | 1))) | 1))
                                                                            v5 = v5
                                                                            v56 := GenS8{F0: 9223372036854775807, F1: 871}
                                                                            _ = v56
                                                                        }
                                                                    default:
                                                                        if false {
                                                                            var v57 int64 = 564
                                                                            if false {
                                                                                v10 = (v10 | v10)
                                                                                v58 := GenS9{F0: v16, F1: 42, F2: (v5 + (v5 + v5))}
                                                                                _ = v58
                                                                            }
                                                                            v1 = append(v1, "hello")
                                                                            _ = v57
                                                                        } else {
                                                                            switch "\n\t" {
                                                                            case "hello":
                                                                                var v59 int8 = 42
                                                                                _ = v59
                                                                            case "a":
                                                                                v16 -= v3
                                                                            case "foo bar":
                                                                                var v60 uint16 = 439
                                                                                _ = v60
                                                                            default:
                                                                                if (true && (v10 > (v10 % (((v10 >> (((v10 % ((0) | 1))) & 63))) | 1)))) {
                                                                                    v61 := GenS10{F0: 308, F1: (v10 & 4294967295), F2: 742}
                                                                                    _ = v61
                                                                                } else {
                                                                                    v62 := GenS9{F0: (-(v16)), F1: 165, F2: (^(v5))}
                                                                                    _ = v62
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = v54
                                                                }
                                                                v16 += v16
                                                            } else {
                                                                v16 = v3
                                                                v63 := GenS7{F0: -1, F1: 1, F2: int32(v10)}
                                                                _ = v63
                                                            }
                                                        } else {
                                                            v64 := GenS1{F0: 42, F1: 127}
                                                            v10 = v10
                                                            v11 += ((v11 - 1e10) + v11)
                                                            _ = v64
                                                        }
                                                        var v65 uint8 = 127
                                                        v3 += (v3 >> ((v10) & 63))
                                                        _ = v65
                                                    } else {
                                                        v10 = (v10 - (v10 % (((v10 | (v10 * 2))) | 1)))
                                                        v16 = v3
                                                    }
                                                    var v66 int32 = v3
                                                    fmt.Println(v11, v3, v11)
                                                    _ = v66
                                                } else {
                                                    fmt.Println(v16)
                                                    v67 := GenS11{F0: v3, F1: true, F2: ""}
                                                    var v68 uint64 = (v5 | v5)
                                                    switch "foo bar" {
                                                    case "a":
                                                        v69 := GenS2{F0: 0}
                                                        _ = v69
                                                    case "\n\t":
                                                        if (9223372036854775807 == 382) {
                                                            var v70 bool = false
                                                            var v71 uint32 = 192
                                                            _ = len(v1)
                                                            _ = v70
                                                            _ = v71
                                                        }
                                                    case "foo bar":
                                                        if false {
                                                            _ = len(v1)
                                                        }
                                                    default:
                                                        var v72 int = (v2 ^ -2147483648)
                                                        fmt.Println(v2, v72)
                                                        fmt.Println(fi14, v68)
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        if (true && (255 <= 1)) {
                                                            v16 -= 2147483647
                                                        }
                                                        _ = v72
                                                    }
                                                    var v73 float32 = 1e10
                                                    switch (-(fi17)) {
                                                    case -1:
                                                        var v74 int16 = 1
                                                        _ = v74
                                                    case 32767:
                                                        v16 += (v16 >> (((v10 >> (((v10 * v10)) & 63))) & 63))
                                                    default:
                                                        v11 = -2.25
                                                    }
                                                    switch (fi14 >> (((v10 - 606)) & 63)) {
                                                    case 4294967295:
                                                        v11 = v73
                                                    }
                                                    _ = v67
                                                    _ = v68
                                                    _ = v73
                                                }
                                            }
                                        }
                                    } else {
                                        switch (("" + "foo bar") + "a") {
                                        case "a":
                                            v10 = v10
                                        }
                                        if false {
                                            v3 -= (v16 ^ v16)
                                            fmt.Println(v5)
                                            var v75 float64 = 1.5
                                            v16 = v3
                                            if ("" < "a") {
                                                fmt.Println(v4, v10)
                                                v4 -= ((v4 - -2.25) + (v4 + (v4 * -2.25)))
                                                _ = len(v1)
                                            } else {
                                                fmt.Println(v10)
                                                if len(v1) > 0 {
                                                    v1[0] = "foo bar"
                                                }
                                                v5 += 1
                                            }
                                            if ((true && false) || false) {
                                                var v76 float32 = float32(v75)
                                                var v77 int = -1
                                                _ = v76
                                                _ = v77
                                            } else {
                                                if (124 == 42) {
                                                    v1 = append(v1, ("a" + "foo bar"))
                                                    var v78 string = ""
                                                    _ = v78
                                                } else {
                                                    v11 = 0.0
                                                    v4 -= v4
                                                }
                                            }
                                            _ = v75
                                        } else {
                                            if (325 <= 796) {
                                                switch "foo bar" {
                                                case "foo bar":
                                                    v2 = (fi14 / ((fi17) | 1))
                                                default:
                                                    _ = len(v1)
                                                }
                                                var v79 uint32 = 2147483647
                                                _ = v79
                                            } else {
                                                fmt.Println(v11, fi14)
                                                _ = len(v1)
                                            }
                                        }
                                    }
                                case -329:
                                    v16 = v3
                                }
                            }
                        }
                        v2 = (-(v2))
                        _ = fi17
                    }
                    v16 = (v16 ^ 42)
                    _ = v16
                }
                _ = fi14
                _ = v15
            }
        }
    }
    for fi80 := 0; fi80 < 14; fi80++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if false {
            fmt.Println(v11)
        }
        _ = fi80
    }
    for fi81 := 0; fi81 < 4; fi81++ {
        v82 := GenS2{F0: 0}
        if (!(("a" + "") == ("hello" + ""))) {
            fmt.Println(v4, v11)
            v83 := GenS12{F0: 42, F1: 1}
            if (!false) {
                v3 += v3
                var v84 uint64 = (v5 << (((v10 + v10)) & 63))
                _ = v84
            } else {
                v4 += (float32(0.0) + (((float32(0.1) / (v11 + ((float32(0.0) + float32(float64(v4))) - v4))) - v4) + ((v11 - (1.5 + (1e10 / (float32(((0.1 + -2.25) * ((-2.25 + 0.1) * 0.1))) + v11)))) / float32(((0.1 / (1e10 - (3.14159 * ((float64(v11) - (1e10 - -2.25)) - -2.25)))) / -2.25)))))
                if true {
                    _ = len(v1)
                } else {
                    v10 -= 629
                }
            }
            _ = v83
        } else {
            switch fi81 {
            case 127:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case -167:
                for fi85 := 0; fi85 < 17; fi85++ {
                    var v86 uint16 = 696
                    _ = len(v1)
                    fmt.Println(v10)
                    _ = fi85
                    _ = v86
                }
            case -565:
                var v87 bool = ((3.14159 + 1.5) >= float32(-2.25))
                _ = v87
            default:
                for fi88 := 0; fi88 < 11; fi88++ {
                    v89 := GenS9{F0: (v3 | (v3 % (((v3 & (v3 >> ((0) & 63)))) | 1))), F1: 203, F2: v5}
                    _ = fi88
                    _ = v89
                }
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        v4 -= (1e10 + v4)
        _ = fi81
        _ = v82
    }
    var v90 uint8 = 128
    fmt.Println(v11)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v90
}
