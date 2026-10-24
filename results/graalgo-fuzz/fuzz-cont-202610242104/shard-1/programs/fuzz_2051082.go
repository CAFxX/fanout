package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 uint32; F1 bool; F2 int8 }
type GenS4 struct { F0 int32; F1 bool }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int = 369
    var v3 float64 = 3.14159
    v4 := make(map[string]uint32)
    v2 = 2147483647
    delete(v4, "")
    if (true && true) {
        switch "\n\t" {
        case "":
            var v5 uint32 = 426
            _ = v5
        default:
            var v6 float64 = (v3 + v3)
            _ = v6
        }
        v7 := []int{v2}
        if true {
            for fi8 := 0; fi8 < 6; fi8++ {
                for fi9 := 0; fi9 < 16; fi9++ {
                    delete(v4, "hello")
                    var v10 string = ("a" + "foo bar")
                    if (-2.25 <= 0.1) {
                        if false {
                            fmt.Println(fi8, v10)
                            v11 := GenS1{F0: v10}
                            _ = v11
                        } else {
                            v12 := GenS1{F0: v10}
                            _ = v12
                        }
                    }
                    _ = fi9
                    _ = v10
                }
                if len(v7) > 0 {
                    v7[0] = -686
                }
                _ = fi8
            }
        }
        _ = v7
    }
    var v13 uint8 = 1
    for fi14 := 0; fi14 < 14; fi14++ {
        var v15 uint16 = 2
        if false {
            var v16 uint16 = (v15 % ((42) | 1))
            _ = v16
        } else {
            var v17 int = v2
            _ = v17
        }
        if (!(!true)) {
            if (("hello" + "\n\t") == "foo bar") {
                for fi18 := 0; fi18 < 17; fi18++ {
                    v13 -= (^(v13))
                    v19 := GenS1{F0: ("hello" + "\n\t")}
                    switch "" {
                    case "foo bar":
                        var v20 int16 = -85
                        _ = v20
                    case "a":
                        var v21 bool = (523 <= 477)
                        v22 := GenS1{F0: ("hello" + ("hello" + ((("a" + "hello") + "foo bar") + ((("" + "") + "foo bar") + ("hello" + "hello")))))}
                        delete(v1, 2147483648)
                        if (!(((((float32(float64(3.14159)) - float32(float64((((1.5 - 1e10) / 1.0) * 0.0)))) != 3.14159) && ((v15 & v15) > (v15 >> ((255) & 63)))) && false) || v21)) {
                            v21 = false
                            if (!(!((float32(((v3 / v3) + 3.14159)) - ((((3.14159 / 0.1) - 1.5) - 3.14159) + 3.14159)) <= (1.5 * 1e10)))) {
                                if v21 {
                                    var v23 string = "\n\t"
                                    _ = v23
                                } else {
                                    var v24 bool = false
                                    _ = v24
                                }
                                v3 = (float64(float32(1e10)) * v3)
                                var v25 int8 = 0
                                var v26 uint16 = (v15 >> ((52) & 63))
                                _ = v25
                                _ = v26
                            } else {
                                v15 = 42
                                if false {
                                    v2 += (v2 & 9223372036854775807)
                                }
                                switch (("\n\t" + ("hello" + "foo bar")) + "") {
                                case "":
                                    v27 := GenS1{F0: ((("hello" + ((("foo bar" + ((("" + "a") + "") + "hello")) + "") + "foo bar")) + "\n\t") + "foo bar")}
                                    _ = v27
                                case "\n\t":
                                    var v28 int32 = 42
                                    _ = v28
                                default:
                                    v13 = 214
                                }
                            }
                        }
                        _ = v21
                        _ = v22
                    case "\n\t":
                        v13 = uint8(v15)
                    default:
                        v3 = (float64((1.5 / 0.1)) - 0.1)
                    }
                    _ = fi18
                    _ = v19
                }
                v29 := GenS1{F0: (("hello" + "foo bar") + "hello")}
                v13 = v13
                _ = v29
            } else {
                v3 = (0.1 + v3)
                switch (^(v2)) {
                case 2147483647:
                    v13 = 120
                case 128:
                    v30 := GenS1{F0: "foo bar"}
                    _ = v30
                case 32768:
                    var v31 int64 = -9223372036854775808
                    _ = v31
                }
            }
            if ((v15 ^ 65535) <= (^(v15))) {
                if false {
                    v15 -= v15
                } else {
                    for fi32 := 0; fi32 < 4; fi32++ {
                        fmt.Println(v3, v2)
                        switch "" {
                        case "hello":
                            if (false || (((("" + "\n\t") + "a") <= ("hello" + "hello")) && false)) {
                                v33 := GenS1{F0: ""}
                                _ = v33
                            }
                        }
                        _ = fi32
                    }
                }
                switch fi14 {
                case 9223372036854775807:
                    for fi34 := 0; fi34 < 6; fi34++ {
                        v35 := GenS1{F0: ("hello" + "\n\t")}
                        _ = fi34
                        _ = v35
                    }
                case -2147483648:
                    v15 += (v15 * 259)
                case 764:
                    if (v15 != 299) {
                        v36 := GenS2{F0: "foo bar", F1: "\n\t"}
                        _ = v36
                    }
                default:
                    var v37 uint8 = (v13 - 255)
                    _ = v37
                }
            } else {
                var v38 int16 = 8
                v38 += v38
                _ = v38
            }
        }
        _ = fi14
        _ = v15
    }
    v3 = v3
    for fi39 := 0; fi39 < 15; fi39++ {
        v13 = v13
        _ = fi39
    }
    if false {
        switch "a" {
        case "":
            for fi40 := 0; fi40 < 7; fi40++ {
                switch 128 {
                case 127:
                    v41 := GenS3{F0: 2147483648, F1: false, F2: 22}
                    _ = v41
                default:
                    v42 := GenS3{F0: 65535, F1: true, F2: 23}
                    _ = v42
                }
                v43 := []int{32768, fi40, fi40, (v2 ^ -1)}
                _ = fi40
                _ = v43
            }
        case "hello":
            if false {
                var v44 uint = 2
                _ = v44
            } else {
                var v45 int16 = -32768
                v2 += v2
                var v46 bool = false
                _ = v45
                _ = v46
            }
        default:
            if false {
                v3 -= v3
                v13 = (v13 << ((255) & 63))
            } else {
                var v47 int = (v2 << ((0) & 63))
                _ = v47
            }
        }
        v48 := GenS1{F0: ""}
        for fi49 := 0; fi49 < 11; fi49++ {
            v13 = 255
            switch "hello" {
            case "\n\t":
                if (-1 >= 42) {
                    fmt.Println(fi49)
                    for fi50 := 0; fi50 < 16; fi50++ {
                        var v51 uint = 18446744073709551615
                        var v52 float64 = (v3 + (v3 + v3))
                        var v53 int = 9223372036854775807
                        _ = fi50
                        _ = v51
                        _ = v52
                        _ = v53
                    }
                }
            case "":
                v13 -= (v13 << ((uint(v13)) & 63))
            case "foo bar":
                v54 := []int{v2}
                _ = v54
            }
            var v55 bool = ((766 > 4294967295) || false)
            _ = fi49
            _ = v55
        }
        _ = v48
    } else {
        var v56 int16 = -504
        _ = v56
    }
    v57 := GenS1{F0: ("hello" + "foo bar")}
    v3 += (v3 * -2.25)
    if (!true) {
        if (false && (true || (true && (128 == v2)))) {
            switch "" {
            case "a":
                fmt.Println(v3, v13, v3)
            case "hello":
                v58 := []int{-1, (v2 & v2), 677}
                _ = v58
            case "":
                var v59 int = v2
                _ = v59
            }
        } else {
            var v60 uint8 = (^(v13))
            _ = v60
        }
    } else {
        for fi61 := 0; fi61 < 6; fi61++ {
            if ((-2 <= 0) || (2147483647 == 2147483648)) {
                v3 -= 0.1
            } else {
                for fi62 := 0; fi62 < 10; fi62++ {
                    fmt.Println(fi61)
                    v13 = (v13 - 42)
                    delete(v4, "")
                    _ = fi62
                }
            }
            v3 += float64((0.0 - (3.14159 / 1.5)))
            fmt.Println(fi61, v3, v13)
            _ = fi61
        }
    }
    var v63 float64 = 0.0
    if (2 < 389) {
        if true {
            var v64 float32 = 0.1
            fmt.Println(v2, v3)
            v2 = v2
            _ = v64
        } else {
            fmt.Println(v63, v63, v3)
            v65 := GenS4{F0: 540, F1: ((!false) || true)}
            _ = v65
        }
        var v66 bool = (!false)
        _ = v66
    }
    var v67 uint = 255
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v63:", v63)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v57
    _ = v63
    _ = v67
}
