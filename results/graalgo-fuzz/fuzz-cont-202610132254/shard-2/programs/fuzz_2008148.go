package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 uint }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 uint8; F1 uint; F2 uint16 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int64; F1 int16; F2 uint }

func main() {
    var v1 uint16 = 846
    v2 := []string{"a", ""}
    v3 := make(map[string]uint32)
    if (0.1 <= float32(0.0)) {
        if true {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if (0 != -861) {
                var v4 uint8 = 0
                var v5 float32 = float32(3.14159)
                _ = v4
                _ = v5
            } else {
                v6 := GenS1{F0: 127, F1: 32768, F2: 682}
                _ = v6
            }
        } else {
            if false {
                var v7 uint8 = 1
                switch -9223372036854775808 {
                case -1:
                    var v8 uint16 = (^(v1))
                    _ = v8
                case 2147483647:
                    v9 := GenS2{F0: 0, F1: v1}
                    _ = v9
                case 1:
                    if (!true) {
                        fmt.Println(v1)
                        if false {
                            var v10 int8 = -53
                            if (!(9223372036854775807 <= -794)) {
                                if true {
                                    for fi11 := 0; fi11 < 11; fi11++ {
                                        switch "hello" {
                                        case "hello":
                                            if false {
                                                v1 += 32767
                                            }
                                        case "a":
                                            if false {
                                                v1 += (v1 - v1)
                                                switch "hello" {
                                                case "\n\t":
                                                    var v12 int32 = 526
                                                    _ = v12
                                                default:
                                                    v2 = append(v2, "hello")
                                                }
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            }
                                        case "\n\t":
                                            v10 = 127
                                        }
                                        _ = fi11
                                    }
                                    v10 = (v10 % ((v10) | 1))
                                } else {
                                    v1 = 42
                                    v1 -= (v1 % ((v1) | 1))
                                    v13 := GenS1{F0: (v10 << ((1) & 63)), F1: int(v1), F2: 255}
                                    _ = v13
                                }
                                for fi14 := 0; fi14 < 3; fi14++ {
                                    switch ("hello" + ("\n\t" + (((("" + "foo bar") + "foo bar") + (("hello" + ((("a" + ("a" + ("a" + ("hello" + "foo bar")))) + "hello") + "")) + "a")) + ("a" + "foo bar")))) {
                                    case "\n\t":
                                        for fi15 := 0; fi15 < 6; fi15++ {
                                            fmt.Println(fi14, v7)
                                            _ = fi15
                                        }
                                    case "foo bar":
                                        if false {
                                            fmt.Println(v1, fi14)
                                            var v16 int8 = v10
                                            for fi17 := 0; fi17 < 10; fi17++ {
                                                switch "\n\t" {
                                                case "hello":
                                                    var v18 int16 = 1000
                                                    _ = v18
                                                case "a":
                                                    var v19 int = 4294967295
                                                    _ = v19
                                                case "foo bar":
                                                    v20 := GenS1{F0: v16, F1: (fi14 << ((1) & 63)), F2: 4294967295}
                                                    _ = v20
                                                default:
                                                    v1 += v1
                                                }
                                                _ = fi17
                                            }
                                            _ = v16
                                        }
                                    case "a":
                                        switch (fi14 >> ((915) & 63)) {
                                        case -824:
                                            fmt.Println(v1)
                                        case 2147483647:
                                            switch "hello" {
                                            case "foo bar":
                                                v21 := GenS2{F0: uint(v1), F1: v1}
                                                _ = v21
                                            default:
                                                var v22 uint8 = 208
                                                _ = v22
                                            }
                                        }
                                    }
                                    fmt.Println(fi14)
                                    _ = fi14
                                }
                                var v23 float64 = 1e10
                                _ = v23
                            } else {
                                if (!(1 == 718)) {
                                    v7 = v7
                                } else {
                                    var v24 uint64 = 42
                                    v3[((("" + ("" + "")) + (("a" + ("hello" + "foo bar")) + "hello")) + "foo bar")] = 42
                                    var v25 uint16 = 65535
                                    _ = v24
                                    _ = v25
                                }
                            }
                            _ = v10
                        }
                        v26 := GenS3{F0: -1, F1: 180}
                        _ = v26
                    }
                }
                fmt.Println(v1)
                _ = v7
            } else {
                for fi27 := 0; fi27 < 20; fi27++ {
                    v1 = v1
                    _ = fi27
                }
                if (("foo bar" + "hello") > "") {
                    v3[("\n\t" + "hello")] = 73
                    v28 := GenS4{F0: "", F1: 181}
                    var v29 string = (("\n\t" + "\n\t") + "foo bar")
                    _ = v28
                    _ = v29
                }
                v2 = append(v2, "a")
            }
            v1 = 968
            var v30 string = ("\n\t" + "hello")
            _ = v30
        }
    } else {
        var v31 string = "foo bar"
        v32 := GenS5{F0: 127, F1: 70, F2: (^(v1))}
        v33 := GenS6{F0: v1}
        _ = v31
        _ = v32
        _ = v33
    }
    if true {
        v1 -= (^(v1))
    }
    var v34 uint = 42
    v1 = v1
    switch (("foo bar" + ("a" + "hello")) + "") {
    case "\n\t":
        switch 1 {
        case 4294967295:
            var v35 uint = 1
            _ = v35
        case 42:
            v36 := GenS2{F0: (v34 % ((42) | 1)), F1: (v1 / ((2) | 1))}
            _ = v36
        }
    case "":
        v34 = (v34 * 65535)
    default:
        switch "foo bar" {
        case "hello":
            var v37 float64 = 0.1
            _ = v37
        case "a":
            v34 -= (^(v34))
        case "\n\t":
            var v38 int64 = 0
            _ = v38
        default:
            v39 := GenS5{F0: 124, F1: 42, F2: (v1 / ((v1) | 1))}
            _ = v39
        }
    }
    if (((("hello" + "foo bar") + "foo bar") + "hello") < "hello") {
        v34 -= v34
        if (!false) {
            var v40 bool = true
            switch -9223372036854775808 {
            case 32767:
                if true {
                    fmt.Println(v34, v40, v34)
                    var v41 bool = v40
                    v42 := GenS5{F0: 255, F1: 2, F2: v1}
                    _ = v41
                    _ = v42
                }
            case 0:
                switch 0 {
                case -1:
                    var v43 uint8 = 4
                    _ = v43
                }
            case -284:
                v44 := v3["foo bar"]
                _ = v44
            default:
                _ = len(v3)
            }
            _ = v40
        } else {
            for fi45 := 0; fi45 < 13; fi45++ {
                fmt.Println(v1, fi45)
                v1 += v1
                _ = fi45
            }
            fmt.Println(v1, v34)
            switch ("hello" + ("a" + "hello")) {
            case "":
                for fi46 := 0; fi46 < 13; fi46++ {
                    var v47 uint = (v34 % ((2) | 1))
                    for fi48 := 0; fi48 < 10; fi48++ {
                        switch (("foo bar" + "\n\t") + ("foo bar" + ("hello" + "foo bar"))) {
                        case "foo bar":
                            if false {
                                _ = len(v3)
                            }
                        default:
                            var v49 float32 = 0.0
                            _ = v49
                        }
                        v47 += v47
                        v3["hello"] = 2147483647
                        _ = fi48
                    }
                    _ = fi46
                    _ = v47
                }
            case "foo bar":
                v2 = append(v2, "\n\t")
            case "\n\t":
                if true {
                    v34 = (v34 + v34)
                    v50 := GenS3{F0: 1, F1: -514}
                    var v51 int16 = 1000
                    _ = v50
                    _ = v51
                } else {
                    var v52 bool = true
                    _ = v52
                }
            }
        }
    } else {
        v53 := GenS4{F0: "\n\t", F1: uint8(v1)}
        var v54 uint = v34
        _ = v53
        _ = v54
    }
    v1 = (^(v1))
    if true {
        v55 := v3["foo bar"]
        for fi56 := 0; fi56 < 6; fi56++ {
            if len(v2) > 0 {
                v2[0] = ("a" + ("foo bar" + ("hello" + "a")))
            }
            v57 := GenS7{F0: 695, F1: 42, F2: 42}
            _ = fi56
            _ = v57
        }
        if true {
            fmt.Println(v55, v34, v55)
        }
        _ = v55
    } else {
        if (-9223372036854775808 != 127) {
            var v58 int32 = -1
            v58 = (v58 - v58)
            var v59 int8 = 1
            _ = v58
            _ = v59
        } else {
            v60 := v3["foo bar"]
            v60 += (v60 + (v60 + v60))
            _ = v60
        }
        v61 := GenS2{F0: v34, F1: (v1 >> ((65535) & 63))}
        var v62 uint32 = 2147483648
        _ = v61
        _ = v62
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
}
