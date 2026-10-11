package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 string; F1 int8 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 bool; F1 int64; F2 int16 }
type GenS8 struct { F0 int; F1 int32; F2 uint16 }
func genFn0(p0 uint64, p1 int8) int8 {
    fmt.Println(p1, p1)
    var v7 bool = true
    p0 = (p0 % ((p0) | 1))
    var v8 int32 = 2147483647
    for fi9 := 0; fi9 < 7; fi9++ {
        if false {
            var v10 int8 = 1
            v8 -= v8
            _ = v10
        }
        for fi11 := 0; fi11 < 11; fi11++ {
            v12 := GenS1{F0: 65535, F1: (v8 / ((v8) | 1))}
            switch (fi9 >> ((2) & 63)) {
            case -1:
                v7 = true
            case 191:
                v8 -= (v8 << ((255) & 63))
            case 342:
                switch "hello" {
                case "hello":
                    v13 := []int{493, 9223372036854775807, fi11}
                    var v14 uint = 255
                    _ = v13
                    _ = v14
                case "a":
                    switch ("\n\t" + ("a" + ("" + "a"))) {
                    case "\n\t":
                        var v15 bool = (((!true) || (160 == 215)) || (!v7))
                        var v16 int8 = 100
                        fmt.Println(p0, v8, fi9)
                        _ = v15
                        _ = v16
                    case "a":
                        var v17 int64 = -9223372036854775808
                        _ = v17
                    case "foo bar":
                        fmt.Println(p0)
                    default:
                        fmt.Println(fi9, p1)
                    }
                default:
                    if false {
                        var v18 float64 = 0.0
                        var v19 uint32 = 1
                        _ = v18
                        _ = v19
                    } else {
                        v20 := GenS1{F0: 1, F1: (v8 / ((2147483647) | 1))}
                        var v21 int64 = 139
                        _ = v20
                        _ = v21
                    }
                }
            default:
                v22 := []int{899}
                _ = v22
            }
            _ = fi11
            _ = v12
        }
        p0 = p0
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return (p1 | (-(p1)))
}
func genFn1(p0 uint32, p1 uint16) uint64 {
    fmt.Println(p0, p0)
    var v23 int32 = 0
    var v24 float32 = ((0.0 - float32((-2.25 + 1.5))) - (3.14159 * 3.14159))
    v25 := []int{32767, -121}
    _ = p0
    _ = p1
    _ = v23
    _ = v24
    _ = v25
    return 2
}

func main() {
    var v1 int = -9223372036854775808
    v2 := make(map[string]bool)
    var v3 int16 = 552
    v4 := []string{"foo bar", "\n\t"}
    var v5 int16 = (v3 << ((462) & 63))
    var v6 float64 = 0.1
    if (65535 < 2147483647) {
        v1 = 4294967295
        fmt.Println(v1)
        var v26 int32 = 152
        _ = v26
    }
    var v27 uint8 = 127
    v3 = v5
    _ = len(v2)
    var v28 uint8 = (^(v27))
    v29 := GenS1{F0: 65535, F1: 42}
    if (((((^(v27)) == v27) && (-1 >= 784)) && false) || false) {
        v5 = v5
    }
    if (18446744073709551615 != 4294967295) {
        for fi30 := 0; fi30 < 8; fi30++ {
            var v31 uint32 = 856
            if (((620 <= 596) || (v5 <= v5)) || (3.14159 >= 1e10)) {
                if (!(!true)) {
                    var v32 int32 = 2147483647
                    for fi33 := 0; fi33 < 4; fi33++ {
                        v34 := GenS2{F0: (("hello" + "a") + (("" + ("foo bar" + ("hello" + "hello"))) + "foo bar")), F1: 42}
                        _ = fi33
                        _ = v34
                    }
                    v35 := GenS1{F0: 255, F1: (v32 ^ v32)}
                    _ = v32
                    _ = v35
                } else {
                    if (!false) {
                        var v36 float64 = v6
                        _ = v36
                    }
                    v5 = v5
                }
                switch "foo bar" {
                case "":
                    v37 := GenS3{F0: 0}
                    _ = v37
                case "foo bar":
                    var v38 int = (^(fi30))
                    _ = v38
                }
                for fi39 := 0; fi39 < 19; fi39++ {
                    fmt.Println(v31, v31)
                    v31 = (v31 + (v31 / ((651) | 1)))
                    _ = fi39
                }
            } else {
                for fi40 := 0; fi40 < 10; fi40++ {
                    var v41 uint32 = uint32(v28)
                    if false {
                        if true {
                            switch "foo bar" {
                            case "\n\t":
                                if len(v4) > 0 {
                                    v4[0] = ""
                                }
                            case "":
                                fmt.Println(fi40, v28, fi30)
                            default:
                                switch ("hello" + ("\n\t" + (("foo bar" + ("hello" + "")) + ""))) {
                                case "\n\t":
                                    fmt.Println(v6, v6, v5)
                                    fmt.Println(v3, v41, v6)
                                    delete(v2, ("" + "a"))
                                case "":
                                    fmt.Println(v28)
                                }
                            }
                        } else {
                            v42 := GenS4{F0: (("" + "\n\t") + ("hello" + "foo bar")), F1: 46}
                            v6 += v6
                            v43 := GenS5{F0: 107}
                            _ = v42
                            _ = v43
                        }
                    } else {
                        if len(v4) > 0 {
                            v4[0] = ("hello" + "a")
                        }
                        var v44 uint32 = v41
                        var v45 float32 = 1.5
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = v44
                        _ = v45
                    }
                    if false {
                        delete(v2, ("foo bar" + "foo bar"))
                        v46 := GenS6{F0: -1}
                        _ = v46
                    } else {
                        switch ("hello" + ((("\n\t" + ("a" + "hello")) + ("\n\t" + "hello")) + (("a" + "a") + "a"))) {
                        case "":
                            v27 = 42
                            var v47 uint32 = 0
                            _ = v47
                        }
                    }
                    _ = fi40
                    _ = v41
                }
            }
            switch (v1 ^ (v1 * 2147483647)) {
            case 127:
                switch v1 {
                case 2147483647:
                    if true {
                        var v48 int16 = -32768
                        _ = v48
                    }
                }
            }
            _ = fi30
            _ = v31
        }
        v27 = v27
    }
    for fi49 := 0; fi49 < 19; fi49++ {
        if (2 == 0) {
            v50 := GenS5{F0: 0}
            _ = v50
        }
        _ = fi49
    }
    fmt.Println(v5)
    if false {
        if ((("" <= (("foo bar" + "a") + "a")) && (!(0 <= 177))) || (1e10 != 1.5)) {
            v28 -= (v27 | (v27 << ((194) & 63)))
            var v51 int16 = (^(v3))
            switch v1 {
            case 32768:
                v52 := GenS2{F0: "hello", F1: 4294967295}
                _ = v52
            case 491:
                if (v1 >= -736) {
                    for fi53 := 0; fi53 < 5; fi53++ {
                        var v54 bool = (1 > 523)
                        if (3.14159 <= (v6 - v6)) {
                            fmt.Println(v54)
                        } else {
                            for fi55 := 0; fi55 < 12; fi55++ {
                                fmt.Println(v28, v5)
                                _ = fi55
                            }
                            v56 := GenS7{F0: (!v54), F1: int64(fi53), F2: v51}
                            v27 += v27
                            _ = v56
                        }
                        _ = fi53
                        _ = v54
                    }
                    if len(v4) > 0 {
                        v4[0] = ("\n\t" + (("" + "hello") + ""))
                    }
                } else {
                    var v57 int16 = (v5 << ((65535) & 63))
                    for fi58 := 0; fi58 < 11; fi58++ {
                        for fi59 := 0; fi59 < 20; fi59++ {
                            switch fi58 {
                            case -614:
                                fmt.Println(v27, v51)
                            }
                            _ = fi59
                        }
                        _ = fi58
                    }
                    _ = len(v4)
                    _ = v57
                }
            default:
                switch (v1 * v1) {
                case 131:
                    if (42 >= 0) {
                        v2[""] = (!(-377 > 230))
                    } else {
                        var v60 uint = 0
                        _ = v60
                    }
                }
            }
            _ = v51
        } else {
            v5 += -1
            if false {
                if false {
                    v5 = v5
                } else {
                    if (false || (0 > 4294967295)) {
                        v6 = v6
                        var v61 float32 = 0.1
                        _ = v61
                    }
                }
            } else {
                fmt.Println(v27)
                v62 := GenS8{F0: v1, F1: -812, F2: 34}
                _ = v62
            }
            v3 -= (v5 ^ v3)
        }
        v28 = (v28 & v27)
        var v63 bool = true
        _ = v63
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
