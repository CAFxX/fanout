package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 string; F1 int16 }
type GenS5 struct { F0 bool; F1 uint16; F2 uint32 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int; F1 int64 }
type GenS8 struct { F0 int32 }

func main() {
    v1 := []string{("" + "foo bar"), "a", ("" + (("hello" + "\n\t") + ("\n\t" + "hello"))), "\n\t"}
    var v2 uint8 = 1
    var v3 int16 = 0
    var v4 string = ("hello" + "hello")
    switch "a" {
    case "foo bar":
        if (false || true) {
            switch 9223372036854775807 {
            case 1:
                v4 = v4
            case 32768:
                v4 = ((v4 + ("foo bar" + (((v4 + v4) + v4) + "foo bar"))) + v4)
            case -639:
                var v5 float64 = 0.0
                _ = v5
            default:
                fmt.Println(v2, v2, v4)
            }
        }
    case "\n\t":
        if false {
            var v6 float64 = float64(1.5)
            v7 := GenS1{F0: 4294967295}
            if false {
                if (858 >= 2) {
                    var v8 float64 = 0.0
                    v2 -= (v2 / ((v2) | 1))
                    for fi9 := 0; fi9 < 7; fi9++ {
                        v8 -= v8
                        v10 := GenS1{F0: 2147483648}
                        for fi11 := 0; fi11 < 18; fi11++ {
                            v8 = v6
                            var v12 uint64 = 873
                            _ = fi11
                            _ = v12
                        }
                        _ = fi9
                        _ = v10
                    }
                    _ = v8
                } else {
                    var v13 bool = (32768 == 32767)
                    _ = v13
                }
                var v14 int8 = 127
                _ = v14
            } else {
                for fi15 := 0; fi15 < 19; fi15++ {
                    if true {
                        var v16 bool = (v3 < 1)
                        if true {
                            fmt.Println(v16, v3)
                            if v16 {
                                v17 := GenS2{F0: 2147483648, F1: v4}
                                switch v4 {
                                case "hello":
                                    if v16 {
                                        if (v16 && v16) {
                                            v18 := GenS1{F0: 4294967295}
                                            _ = v18
                                        } else {
                                            if (((true && v16) && v16) && v16) {
                                                v16 = (!(true && (127 == 0)))
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            } else {
                                                v19 := GenS1{F0: 140}
                                                var v20 float32 = (-2.25 - (3.14159 / -2.25))
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                _ = v19
                                                _ = v20
                                            }
                                            fmt.Println(v2)
                                            for fi21 := 0; fi21 < 6; fi21++ {
                                                fmt.Println(v2, v6, v4)
                                                v22 := GenS2{F0: 288, F1: (v4 + (v4 + (((("a" + v4) + v4) + v4) + "a")))}
                                                switch fi15 {
                                                case 32767:
                                                    v3 = (v3 + -5)
                                                }
                                                if ((v2 * 1) <= (v2 << ((2) & 63))) {
                                                    var v23 int64 = 9223372036854775807
                                                    _ = v23
                                                }
                                                var v24 uint16 = 255
                                                _ = fi21
                                                _ = v22
                                                _ = v24
                                            }
                                        }
                                        var v25 uint8 = v2
                                        _ = v25
                                    } else {
                                        switch v4 {
                                        case "foo bar":
                                            fmt.Println(fi15)
                                        }
                                        if true {
                                            fmt.Println(v16, v3, v2)
                                            var v26 string = ((v4 + (v4 + (v4 + ((("foo bar" + (v4 + "foo bar")) + v4) + v4)))) + ("hello" + ("a" + "")))
                                            var v27 bool = v16
                                            _ = v26
                                            _ = v27
                                        }
                                        fmt.Println(v16, v3, v16)
                                    }
                                }
                                _ = v17
                            } else {
                                var v28 int32 = 994
                                var v29 float64 = (-2.25 - float64(0.0))
                                _ = v28
                                _ = v29
                            }
                            fmt.Println(v4)
                        } else {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            v3 += v3
                        }
                        v30 := GenS2{F0: 513, F1: ((("" + v4) + "\n\t") + v4)}
                        _ = v16
                        _ = v30
                    } else {
                        v1 = append(v1, (v4 + (v4 + (v4 + "foo bar"))))
                    }
                    v3 = v3
                    v3 = -73
                    _ = fi15
                }
                v3 = 1
            }
            _ = v6
            _ = v7
        }
    case "hello":
        if (!(0 != 1)) {
            fmt.Println(v2)
            v4 = v4
        } else {
            fmt.Println(v3)
        }
    }
    v1 = append(v1, "a")
    v3 = (v3 - (v3 % (((v3 << ((820) & 63))) | 1)))
    if false {
        v2 = v2
        fmt.Println(v4, v3)
        fmt.Println(v2, v3, v3)
    } else {
        if ((v2 != v2) && (1.5 > 1e10)) {
            v2 = v2
            for fi31 := 0; fi31 < 4; fi31++ {
                v1 = append(v1, ("a" + v4))
                _ = fi31
            }
            switch v4 {
            case "a":
                if (-116 <= -128) {
                    if (!false) {
                        v3 = (^(v3))
                        fmt.Println(v2)
                        fmt.Println(v4, v2)
                    }
                    var v32 string = v4
                    var v33 uint32 = 0
                    _ = v32
                    _ = v33
                } else {
                    v2 = (v2 / ((v2) | 1))
                }
            default:
                switch (v4 + "a") {
                case "hello":
                    if true {
                        if (true || false) {
                            for fi34 := 0; fi34 < 3; fi34++ {
                                v4 = (v4 + v4)
                                for fi35 := 0; fi35 < 10; fi35++ {
                                    _ = len(v1)
                                    if true {
                                        if len(v1) > 0 {
                                            v1[0] = ("a" + (v4 + v4))
                                        }
                                        var v36 uint32 = 566
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = v36
                                    } else {
                                        if (!true) {
                                            fmt.Println(v2, v3, v2)
                                            if (false && true) {
                                                var v37 uint32 = 838
                                                switch (v4 + (((v4 + v4) + ((("" + (v4 + "foo bar")) + "a") + v4)) + ((("\n\t" + "\n\t") + "") + "hello"))) {
                                                case "\n\t":
                                                    var v38 int32 = -2147483648
                                                    _ = v38
                                                case "a":
                                                    _ = len(v1)
                                                }
                                                _ = v37
                                            }
                                        }
                                    }
                                    var v39 float64 = 3.14159
                                    var v40 uint32 = 65535
                                    v39 = (1.5 / (float64(0.0) + (float64(((1e10 - 1e10) * (float32(v39) - (1.5 * -2.25)))) / v39)))
                                    _ = fi35
                                    _ = v39
                                    _ = v40
                                }
                                _ = fi34
                            }
                        }
                        if len(v1) > 0 {
                            v1[0] = "\n\t"
                        }
                        var v41 bool = ((373 < 42) && false)
                        _ = v41
                    } else {
                        if false {
                            v4 = v4
                        } else {
                            var v42 uint8 = (v2 | v2)
                            v43 := GenS3{F0: -1}
                            if (!(false || (!true))) {
                                v3 = (^(v3))
                            } else {
                                v2 = v2
                            }
                            _ = v42
                            _ = v43
                        }
                    }
                case "a":
                    if (false && true) {
                        v44 := GenS4{F0: "foo bar", F1: -32768}
                        _ = v44
                    }
                }
            }
        } else {
            v45 := GenS5{F0: (!false), F1: 626, F2: 0}
            var v46 uint16 = 32767
            _ = v45
            _ = v46
        }
        fmt.Println(v3, v4, v4)
    }
    for fi47 := 0; fi47 < 9; fi47++ {
        fmt.Println(fi47, v4, v3)
        v2 -= (v2 & v2)
        var v48 float64 = -2.25
        _ = fi47
        _ = v48
    }
    var v49 uint = 0
    var v50 float64 = (3.14159 * 3.14159)
    var v51 uint = 255
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (65535 != 42) {
        for fi52 := 0; fi52 < 5; fi52++ {
            v53 := GenS6{F0: 389}
            v51 = v51
            fmt.Println(fi52, fi52)
            _ = fi52
            _ = v53
        }
        var v54 uint32 = 2
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v54
    } else {
        v55 := GenS7{F0: -2147483648, F1: -288}
        fmt.Println(v4, v50, v49)
        for fi56 := 0; fi56 < 14; fi56++ {
            var v57 int32 = 920
            var v58 uint = v51
            _ = fi56
            _ = v57
            _ = v58
        }
        _ = v55
    }
    v59 := GenS5{F0: false, F1: 2, F2: 42}
    if true {
        var v60 uint16 = 883
        for fi61 := 0; fi61 < 19; fi61++ {
            for fi62 := 0; fi62 < 16; fi62++ {
                var v63 bool = (false && true)
                _ = fi62
                _ = v63
            }
            var v64 int32 = -2147483648
            _ = fi61
            _ = v64
        }
        _ = v60
    }
    if ((((0.1 * ((((3.14159 - 0.0) - -2.25) + 1.5) - float32(((float64(-2.25) + 0.0) - v50)))) - float32(0.0)) >= 3.14159) || true) {
        fmt.Println(v3)
    }
    var v65 uint = 4294967295
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if false {
        switch 1 {
        case 4294967295:
            for fi66 := 0; fi66 < 6; fi66++ {
                var v67 int = fi66
                for fi68 := 0; fi68 < 6; fi68++ {
                    var v69 float32 = 1e10
                    if false {
                        v50 = (v50 / float64(v69))
                    } else {
                        v65 += 380
                    }
                    var v70 int64 = 42
                    _ = fi68
                    _ = v69
                    _ = v70
                }
                v50 = (v50 - v50)
                _ = fi66
                _ = v67
            }
        case 32768:
            v71 := GenS8{F0: 1}
            _ = v71
        }
        var v72 int64 = 0
        _ = v72
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v50
    _ = v51
    _ = v59
    _ = v65
}
