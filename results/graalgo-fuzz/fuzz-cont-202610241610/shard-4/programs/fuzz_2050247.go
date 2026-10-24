package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64; F1 int8 }
type GenS5 struct { F0 uint16; F1 int16; F2 int64 }
type GenS6 struct { F0 uint; F1 int64 }

func main() {
    var v1 uint = 4294967295
    var v2 float64 = (float64((((3.14159 + 0.0) * -2.25) * 0.1)) + ((1e10 + 0.0) / 3.14159))
    var v3 float64 = 1.5
    v4 := []string{"\n\t", ("foo bar" + "a"), ""}
    for fi5 := 0; fi5 < 3; fi5++ {
        v6 := GenS1{F0: 0, F1: 25, F2: v1}
        v1 += (v1 & v1)
        var v7 float32 = 1.5
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 uint = (v1 + (^(v1)))
    var v9 int16 = 42
    v10 := GenS2{F0: ("\n\t" + "")}
    if true {
        if true {
            v1 = v8
            var v11 float64 = (float64((-2.25 + (-2.25 / (((1e10 - 1e10) / (1e10 + 0.1)) - float32(0.1))))) / (v2 - v2))
            _ = v11
        }
        var v12 int8 = 0
        switch ("\n\t" + "foo bar") {
        case "a":
            if len(v4) > 0 {
                v4[0] = ("foo bar" + (("" + "") + ("foo bar" + "\n\t")))
            }
        case "\n\t":
            if true {
                var v13 uint32 = 4294967295
                fmt.Println(v8)
                v12 = (v12 % (((v12 << ((v1) & 63))) | 1))
                _ = v13
            }
        case "hello":
            v14 := GenS1{F0: 295, F1: 100, F2: v8}
            _ = v14
        default:
            var v15 float64 = v2
            _ = v15
        }
        _ = v12
    }
    for fi16 := 0; fi16 < 12; fi16++ {
        fmt.Println(v1, fi16, fi16)
        fmt.Println(v9)
        if true {
            v8 += v1
            for fi17 := 0; fi17 < 12; fi17++ {
                if ((((("hello" + "") + "\n\t") + ((("a" + ("" + "\n\t")) + ("a" + (("" + "\n\t") + ""))) + "")) <= "\n\t") || true) {
                    v18 := GenS1{F0: 351, F1: -1, F2: (v1 & 18446744073709551615)}
                    if (!((-883 > 1) || (("hello" + "\n\t") > ""))) {
                        if ("a" != "") {
                            v2 = v2
                            fmt.Println(v3, fi16)
                        }
                        if false {
                            v19 := GenS2{F0: "foo bar"}
                            var v20 int = 151
                            _ = v19
                            _ = v20
                        } else {
                            v3 = 0.0
                            var v21 float32 = -2.25
                            _ = v21
                        }
                        var v22 float32 = (1.5 / 3.14159)
                        _ = v22
                    }
                    _ = v18
                }
                _ = fi17
            }
            v8 = v8
        } else {
            for fi23 := 0; fi23 < 15; fi23++ {
                v9 += v9
                _ = fi23
            }
            if true {
                var v24 string = ""
                v2 = (((v2 * ((((0.0 + 0.0) / v3) - v2) + v3)) - float64(((0.0 / -2.25) + (-2.25 - -2.25)))) - v3)
                _ = v24
            }
        }
        _ = fi16
    }
    v25 := GenS3{F0: 884}
    fmt.Println(v3, v8)
    for fi26 := 0; fi26 < 14; fi26++ {
        var v27 int16 = v9
        _ = fi26
        _ = v27
    }
    var v28 float32 = -2.25
    _ = len(v4)
    for fi29 := 0; fi29 < 10; fi29++ {
        fmt.Println(v9, v2, v2)
        switch "" {
        case "":
            v2 -= ((v3 + float64(v28)) * v3)
        case "a":
            var v30 uint = 18446744073709551615
            _ = v30
        }
        v2 = (0.0 - v2)
        _ = fi29
    }
    switch -199 {
    case 9223372036854775807:
        v31 := GenS1{F0: 0, F1: 0, F2: v8}
        _ = v31
    case 127:
        if (((false || true) && true) && (!(0.1 < v28))) {
            var v32 float32 = 1e10
            _ = v32
        }
    case 2147483647:
        if (uint32(v1) <= 1) {
            v2 = (float64(v28) / v3)
            for fi33 := 0; fi33 < 1; fi33++ {
                switch (("\n\t" + (("foo bar" + "foo bar") + "hello")) + "a") {
                case "foo bar":
                    if ("hello" >= "foo bar") {
                        v34 := GenS4{F0: 18446744073709551615, F1: -39}
                        if false {
                            var v35 bool = true
                            if ((v9 << ((224) & 63)) != 42) {
                                for fi36 := 0; fi36 < 20; fi36++ {
                                    fmt.Println(v35, fi33)
                                    if len(v4) > 0 {
                                        v4[0] = "\n\t"
                                    }
                                    v8 += (^(v1))
                                    _ = fi36
                                }
                            }
                            v8 += (v1 - 4294967295)
                            _ = v35
                        }
                        _ = v34
                    }
                case "a":
                    if len(v4) > 0 {
                        v4[0] = "a"
                    }
                case "\n\t":
                    for fi37 := 0; fi37 < 11; fi37++ {
                        var v38 uint32 = 0
                        if true {
                            switch (("\n\t" + "\n\t") + "a") {
                            case "a":
                                if (fi33 != (fi33 << ((255) & 63))) {
                                    switch "a" {
                                    case "a":
                                        v28 += (v28 - (v28 + 3.14159))
                                    default:
                                        var v39 float64 = 0.1
                                        v40 := GenS1{F0: 255, F1: 1, F2: v1}
                                        var v41 string = "foo bar"
                                        _ = v39
                                        _ = v40
                                        _ = v41
                                    }
                                    if true {
                                        var v42 int8 = -128
                                        var v43 int16 = v9
                                        _ = v42
                                        _ = v43
                                    } else {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    }
                                    v4 = append(v4, "a")
                                }
                            case "\n\t":
                                var v44 int32 = 42
                                _ = v44
                            case "foo bar":
                                v3 += ((-2.25 + (v2 * (float64(v28) - 3.14159))) / v2)
                            }
                            switch -250 {
                            case 32768:
                                if (0 != 1) {
                                    if (!((-(v9)) == v9)) {
                                        if (v28 >= ((v28 * v28) + v28)) {
                                            v45 := GenS1{F0: 860, F1: 36, F2: 814}
                                            fmt.Println(v8)
                                            _ = v45
                                        }
                                        var v46 int16 = (v9 + v9)
                                        _ = v46
                                    }
                                } else {
                                    v47 := GenS5{F0: 2, F1: 42, F2: -1}
                                    _ = v47
                                }
                            default:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                            if true {
                                v48 := GenS5{F0: 1, F1: 1, F2: -9223372036854775808}
                                v3 = (0.1 - v3)
                                _ = v48
                            } else {
                                v38 = (v38 + v38)
                                switch ("a" + "a") {
                                case "":
                                    v49 := GenS1{F0: 293, F1: 127, F2: v1}
                                    _ = v49
                                case "hello":
                                    var v50 float64 = v2
                                    _ = v50
                                case "foo bar":
                                    if (2147483647 > -177) {
                                        switch "foo bar" {
                                        case "":
                                            fmt.Println(v8)
                                        case "a":
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        default:
                                            var v51 int64 = 42
                                            _ = v51
                                        }
                                        if ((fi37 / (((fi33 << (((^(v1))) & 63))) | 1)) > (fi33 / ((fi37) | 1))) {
                                            fmt.Println(v3, v9, v38)
                                            v52 := GenS1{F0: 65535, F1: 0, F2: v8}
                                            var v53 uint = v8
                                            _ = v52
                                            _ = v53
                                        } else {
                                            var v54 bool = (!true)
                                            _ = v54
                                        }
                                        if false {
                                            v28 = v28
                                            var v55 int8 = -93
                                            fmt.Println(v8)
                                            _ = v55
                                        }
                                    }
                                }
                            }
                        }
                        _ = fi37
                        _ = v38
                    }
                }
                var v56 string = ""
                for fi57 := 0; fi57 < 2; fi57++ {
                    var v58 bool = (-497 > 918)
                    _ = fi57
                    _ = v58
                }
                _ = fi33
                _ = v56
            }
        }
    }
    if false {
        var v59 int = -981
        _ = v59
    } else {
        var v60 int64 = -1
        v8 = (v8 & 1)
        if true {
            v1 -= 18446744073709551615
        }
        _ = v60
    }
    var v61 int32 = 391
    v62 := GenS6{F0: (v1 % ((v8) | 1)), F1: -687}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v61:", v61)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v25
    _ = v28
    _ = v3
    _ = v4
    _ = v61
    _ = v62
    _ = v8
    _ = v9
}
