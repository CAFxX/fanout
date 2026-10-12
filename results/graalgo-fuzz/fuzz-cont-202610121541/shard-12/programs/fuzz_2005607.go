package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 bool; F1 uint; F2 bool }
type GenS5 struct { F0 int16; F1 uint16 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int64; F1 string }

func main() {
    var v1 int32 = -123
    v2 := []string{"a"}
    var v3 uint8 = 2
    var v4 uint64 = 4294967295
    var v5 int32 = v1
    if ((-(v5)) == 400) {
        v6 := GenS1{F0: "", F1: 150}
        _ = v6
    }
    fmt.Println(v4)
    v7 := GenS2{F0: 0}
    if ((!((true || true) || true)) && (!((797 >= 2) && (!false)))) {
        if (("hello" + "a") >= "foo bar") {
            switch (("hello" + "hello") + ("\n\t" + "a")) {
            case "a":
                if true {
                    v4 -= v4
                } else {
                    v4 += v4
                }
            case "\n\t":
                switch ("foo bar" + "a") {
                case "a":
                    fmt.Println(v5, v1)
                case "hello":
                    fmt.Println(v5)
                case "foo bar":
                    var v8 bool = false
                    _ = v8
                default:
                    var v9 uint8 = (v3 | v3)
                    _ = v9
                }
            case "foo bar":
                for fi10 := 0; fi10 < 2; fi10++ {
                    v2 = append(v2, "")
                    _ = fi10
                }
            default:
                fmt.Println(v4)
            }
            v5 -= (-(v5))
            if true {
                var v11 bool = false
                v11 = v11
                _ = v11
            }
        } else {
            if ((v3 * v3) == v3) {
                for fi12 := 0; fi12 < 10; fi12++ {
                    if len(v2) > 0 {
                        v2[0] = (("hello" + ("hello" + "foo bar")) + ("a" + ("\n\t" + "hello")))
                    }
                    _ = fi12
                }
                if (747 < int(v4)) {
                    switch ("\n\t" + ("\n\t" + "foo bar")) {
                    case "a":
                        switch 2147483647 {
                        case -9223372036854775808:
                            v4 = (^(v4))
                        case 128:
                            for fi13 := 0; fi13 < 11; fi13++ {
                                _ = len(v2)
                                for fi14 := 0; fi14 < 9; fi14++ {
                                    if (((true || true) && true) && true) {
                                        var v15 uint = 65535
                                        var v16 int = fi14
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        _ = v15
                                        _ = v16
                                    } else {
                                        v1 -= 230
                                        if len(v2) > 0 {
                                            v2[0] = ""
                                        }
                                        var v17 float32 = 1e10
                                        _ = v17
                                    }
                                    var v18 float32 = (0.0 / 1.0)
                                    _ = fi14
                                    _ = v18
                                }
                                _ = fi13
                            }
                        case 1:
                            v3 = (v3 % (((v3 * v3)) | 1))
                        default:
                            v3 = (v3 % ((v3) | 1))
                        }
                    case "hello":
                        for fi19 := 0; fi19 < 2; fi19++ {
                            var v20 string = ""
                            _ = fi19
                            _ = v20
                        }
                    default:
                        if true {
                            var v21 uint32 = 2147483648
                            var v22 float32 = (((1e10 - 3.14159) + 0.1) / (-2.25 - float32((float64((3.14159 - -2.25)) + (0.0 / (float64((1e10 / 1e10)) + ((0.1 / 1.0) + (float64(0.1) + -2.25))))))))
                            var v23 uint64 = (v4 % ((v4) | 1))
                            _ = v21
                            _ = v22
                            _ = v23
                        }
                    }
                    for fi24 := 0; fi24 < 14; fi24++ {
                        switch (^(fi24)) {
                        case -1:
                            for fi25 := 0; fi25 < 11; fi25++ {
                                if true {
                                    if true {
                                        if ("" >= "\n\t") {
                                            var v26 string = ""
                                            var v27 uint16 = 42
                                            fmt.Println(v4, v3, v3)
                                            _ = v26
                                            _ = v27
                                        }
                                    } else {
                                        v28 := GenS1{F0: "hello", F1: 42}
                                        if ((v4 << ((4294967295) & 63)) < v4) {
                                            if true {
                                                if true {
                                                    v29 := GenS3{F0: 544, F1: (18446744073709551615 >= v4)}
                                                    switch "" {
                                                    case "\n\t":
                                                        var v30 int64 = 386
                                                        _ = v30
                                                    }
                                                    _ = v29
                                                }
                                                v4 += (v4 + (v4 | v4))
                                                v2 = append(v2, "hello")
                                            } else {
                                                if true {
                                                    _ = len(v2)
                                                    var v31 bool = false
                                                    _ = v31
                                                } else {
                                                    var v32 uint = 52
                                                    v1 = v5
                                                    v33 := GenS4{F0: (((535 <= 42) && (v4 != 241)) && true), F1: v32, F2: false}
                                                    _ = v32
                                                    _ = v33
                                                }
                                                switch "" {
                                                case "\n\t":
                                                    if true {
                                                        v2 = append(v2, "hello")
                                                    } else {
                                                        v1 += 672
                                                        var v34 bool = (false || false)
                                                        _ = v34
                                                    }
                                                case "":
                                                    _ = len(v2)
                                                case "a":
                                                    var v35 float64 = 0.1
                                                    _ = v35
                                                default:
                                                    if true {
                                                        v1 -= -2147483648
                                                        if false {
                                                            var v36 uint16 = 42
                                                            if true {
                                                                v2 = append(v2, "a")
                                                                v4 -= (v4 + v4)
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                v3 = (v3 % ((127) | 1))
                                                            }
                                                            fmt.Println(v4, fi25, v5)
                                                            _ = v36
                                                        }
                                                    } else {
                                                        _ = len(v2)
                                                    }
                                                }
                                                v2 = append(v2, ("hello" + "a"))
                                            }
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        } else {
                                            if true {
                                                var v37 int = fi24
                                                v1 -= -2147483648
                                                if ((1e10 / 1e10) >= 3.14159) {
                                                    switch (^(fi24)) {
                                                    case -41:
                                                        var v38 float32 = 1e10
                                                        _ = v38
                                                    case -154:
                                                        v3 = v3
                                                    default:
                                                        if false {
                                                            v4 = (v4 / (((v4 + 0)) | 1))
                                                        } else {
                                                            var v39 float64 = (((1.5 + ((((((3.14159 - (float64(0.1) * (1e10 * 0.0))) / 1.5) + 0.0) * (-2.25 / ((((0.1 - float64(1e10)) * ((1.5 - float64(0.1)) + -2.25)) / 1e10) / (1e10 * ((float64(3.14159) + float64(0.1)) * 0.1))))) + (1.5 * float64(1e10))) / (((3.14159 + 0.1) - (1.5 / (1.5 - -2.25))) / (float64(0.0) - 1.5)))) - (1e10 - 0.0)) * (1.5 + 1e10))
                                                            var v40 uint64 = (^(v4))
                                                            _ = v39
                                                            _ = v40
                                                        }
                                                    }
                                                    switch fi24 {
                                                    case 2147483647:
                                                        v41 := GenS5{F0: -1, F1: 1}
                                                        _ = v41
                                                    }
                                                }
                                                _ = v37
                                            }
                                        }
                                        v42 := GenS4{F0: (((-2.25 / 3.14159) + ((1.5 - 1.5) / 0.1)) < float32(-2.25)), F1: uint(v4), F2: true}
                                        _ = v28
                                        _ = v42
                                    }
                                    switch "hello" {
                                    case "\n\t":
                                        if true {
                                            v43 := GenS6{F0: 2}
                                            fmt.Println(fi24, v1)
                                            v44 := GenS7{F0: -1, F1: (("\n\t" + "\n\t") + (("a" + ("" + "foo bar")) + "foo bar"))}
                                            _ = v43
                                            _ = v44
                                        } else {
                                            if len(v2) > 0 {
                                                v2[0] = "a"
                                            }
                                            v5 = v1
                                        }
                                    case "hello":
                                        v4 = v4
                                    case "foo bar":
                                        switch (("a" + "") + "foo bar") {
                                        case "\n\t":
                                            var v45 uint16 = 150
                                            _ = v45
                                        case "foo bar":
                                            var v46 int = (^(fi24))
                                            _ = v46
                                        default:
                                            switch fi25 {
                                            case -3:
                                                v4 = 896
                                            }
                                        }
                                    default:
                                        v47 := GenS4{F0: true, F1: 65535, F2: false}
                                        _ = v47
                                    }
                                } else {
                                    _ = len(v2)
                                    var v48 uint8 = (v3 - (v3 - (v3 >> ((2) & 63))))
                                    fmt.Println(v48, v3)
                                    _ = v48
                                }
                                v4 = (v4 >> ((65535) & 63))
                                var v49 int8 = 42
                                _ = fi25
                                _ = v49
                            }
                        case 128:
                            v1 = 986
                        case 9223372036854775807:
                            fmt.Println(v1, fi24, fi24)
                        }
                        for fi50 := 0; fi50 < 12; fi50++ {
                            switch 9223372036854775807 {
                            case -1:
                                var v51 int8 = -128
                                _ = v51
                            case -821:
                                fmt.Println(fi24)
                            default:
                                v52 := GenS7{F0: 9223372036854775807, F1: ("hello" + "hello")}
                                _ = v52
                            }
                            var v53 bool = (false || (!(9223372036854775807 <= -9223372036854775808)))
                            _ = fi50
                            _ = v53
                        }
                        _ = fi24
                    }
                    var v54 float32 = -2.25
                    _ = v54
                }
                v5 += v5
            } else {
                switch 128 {
                case 127:
                    var v55 uint32 = 2147483648
                    _ = v55
                case 128:
                    var v56 int16 = 1000
                    _ = v56
                default:
                    v1 = 2147483647
                }
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v57 int64 = 970
            _ = v57
        }
    }
    v4 += v4
    var v58 bool = true
    if (v3 < v3) {
        var v59 float32 = 1e10
        var v60 string = ("hello" + "\n\t")
        var v61 int64 = 9223372036854775807
        _ = v59
        _ = v60
        _ = v61
    }
    v62 := GenS7{F0: -14, F1: (("foo bar" + ("hello" + "\n\t")) + (("\n\t" + (("hello" + "a") + (("" + "\n\t") + "hello"))) + ("\n\t" + "hello")))}
    if (42 <= 127) {
        fmt.Println(v58)
    }
    for fi63 := 0; fi63 < 17; fi63++ {
        for fi64 := 0; fi64 < 14; fi64++ {
            v5 += (v5 - (v5 << ((2) & 63)))
            var v65 uint16 = 0
            var v66 int8 = 127
            _ = fi64
            _ = v65
            _ = v66
        }
        var v67 uint = 0
        var v68 uint = v67
        _ = fi63
        _ = v67
        _ = v68
    }
    if v58 {
        if v58 {
            switch ("" + "hello") {
            case "hello":
                v1 = (v1 ^ v5)
            default:
                v3 = (v3 << ((0) & 63))
            }
            if false {
                fmt.Println(v58, v1, v1)
                var v69 uint16 = 65535
                if len(v2) > 0 {
                    v2[0] = (("a" + "hello") + "hello")
                }
                _ = v69
            } else {
                if v58 {
                    var v70 int16 = 42
                    switch "\n\t" {
                    case "hello":
                        v5 -= v1
                    case "a":
                        v1 += (v1 & v1)
                    case "":
                        fmt.Println(v3, v70, v5)
                    }
                    _ = v70
                }
            }
        }
        var v71 float64 = (((((0.1 - (((float64(((1e10 - (1e10 - 1e10)) * (0.1 + float32(0.1)))) * 0.1) + 1e10) + (-2.25 - 3.14159))) + 0.0) * 1e10) + 0.1) / (3.14159 / 3.14159))
        _ = v71
    } else {
        v72 := GenS7{F0: 42, F1: "foo bar"}
        var v73 bool = ((^(v3)) >= 127)
        _ = v72
        _ = v73
    }
    v5 += (v5 - (v1 | v1))
    var v74 bool = v58
    var v75 int = -2147483648
    v76 := GenS2{F0: -9223372036854775808}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v62
    _ = v7
    _ = v74
    _ = v75
    _ = v76
}
