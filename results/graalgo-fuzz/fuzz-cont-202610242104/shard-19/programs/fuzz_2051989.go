package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 int32; F1 bool; F2 uint16 }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 int8; F1 uint64; F2 int }
func genFn0() int64 {
    if ((true || (!true)) || false) {
        var v5 uint8 = 128
        fmt.Println(v5)
        var v6 uint = 685
        _ = v5
        _ = v6
    }
    switch ("" + "\n\t") {
    case "hello":
        for fi7 := 0; fi7 < 13; fi7++ {
            fmt.Println(fi7)
            switch (^(fi7)) {
            case 127:
                v8 := []int{-2147483648, fi7}
                _ = v8
            case 0:
                fmt.Println(fi7)
            default:
                fmt.Println(fi7)
            }
            var v9 int8 = 39
            _ = fi7
            _ = v9
        }
    default:
        var v10 uint = 495
        _ = v10
    }
    var v11 float64 = float64(((0.0 / 1.5) + -2.25))
    _ = v11
    return 9223372036854775807
}
func genFn1() int32 {
    v12 := []int{}
    var v13 int64 = -74
    fmt.Println(v13)
    v13 = (v13 << ((2) & 63))
    v14 := GenS1{F0: 0, F1: 2}
    _ = v12
    _ = v13
    _ = v14
    return 800
}

func main() {
    v1 := []string{"a", ("" + "hello"), ("hello" + "a")}
    v2 := make(map[string]string)
    var v3 int8 = 52
    var v4 bool = (!true)
    switch 127 {
    case 128:
        v3 = -128
    case 425:
        if (!v4) {
            for fi15 := 0; fi15 < 20; fi15++ {
                var v16 string = ("hello" + "foo bar")
                switch v16 {
                case "":
                    if (!(255 >= 42)) {
                        if v4 {
                            var v17 string = "a"
                            switch fi15 {
                            case 147:
                                var v18 uint8 = 150
                                _ = v18
                            case 32768:
                                v19 := GenS2{F0: 209, F1: v4, F2: 0}
                                _ = v19
                            default:
                                if v4 {
                                    if (v4 && true) {
                                        if (0 <= 0) {
                                            var v20 uint16 = 2
                                            _ = v20
                                        }
                                        delete(v2, (("foo bar" + ("a" + "\n\t")) + ((v17 + v17) + "a")))
                                        fmt.Println(v17)
                                    }
                                    v1 = append(v1, (v16 + v17))
                                    var v21 float32 = 1.5
                                    _ = v21
                                }
                            }
                            fmt.Println(v4)
                            _ = v17
                        } else {
                            v22 := GenS1{F0: v3, F1: 97}
                            if (514 >= -1) {
                                v16 = ((("foo bar" + ("" + v16)) + v16) + ("a" + ((((v16 + ((v16 + (v16 + v16)) + (v16 + (((v16 + v16) + (v16 + v16)) + "hello")))) + ((v16 + (v16 + "")) + v16)) + ("a" + "a")) + v16)))
                            }
                            var v23 int8 = v3
                            _ = v22
                            _ = v23
                        }
                    } else {
                        v24 := GenS3{F0: v16, F1: 708}
                        _ = v24
                    }
                case "foo bar":
                    var v25 float64 = (-2.25 / 1.0)
                    _ = v25
                case "a":
                    if (v4 && false) {
                        switch (fi15 - (fi15 >> ((42) & 63))) {
                        case 0:
                            for fi26 := 0; fi26 < 2; fi26++ {
                                fmt.Println(v3)
                                v16 = ""
                                v3 = (^(v3))
                                _ = fi26
                            }
                        case 32768:
                            fmt.Println(fi15)
                        }
                        for fi27 := 0; fi27 < 10; fi27++ {
                            v1 = append(v1, ("" + v16))
                            _ = fi27
                        }
                    } else {
                        v4 = ((!v4) || (0 == 1))
                    }
                default:
                    delete(v2, "")
                }
                _ = fi15
                _ = v16
            }
        } else {
            var v28 float32 = 1e10
            v3 = 100
            if v4 {
                switch "" {
                case "":
                    v3 = -1
                }
                if (771 >= 164) {
                    v29 := GenS2{F0: -1, F1: (1 != 4294967295), F2: 255}
                    switch 2147483647 {
                    case 721:
                        if (205 < 958) {
                            var v30 uint = 822
                            _ = v30
                        } else {
                            var v31 bool = v4
                            for fi32 := 0; fi32 < 3; fi32++ {
                                v31 = v4
                                _ = fi32
                            }
                            var v33 uint8 = 178
                            _ = v31
                            _ = v33
                        }
                    case -510:
                        v2["\n\t"] = "foo bar"
                    default:
                        var v34 uint8 = 6
                        _ = v34
                    }
                    _ = v29
                }
            } else {
                switch -2147483648 {
                case -2147483648:
                    var v35 string = "hello"
                    _ = v35
                case 0:
                    v28 = v28
                case -408:
                    v3 -= 0
                }
            }
            _ = v28
        }
    default:
        for fi36 := 0; fi36 < 16; fi36++ {
            var v37 uint16 = 255
            v3 = v3
            v38 := GenS3{F0: "hello", F1: 2147483648}
            _ = fi36
            _ = v37
            _ = v38
        }
    }
    if v4 {
        var v39 uint32 = 310
        v39 = (v39 | 1)
        var v40 int32 = 0
        _ = v39
        _ = v40
    } else {
        fmt.Println(v3)
        for fi41 := 0; fi41 < 17; fi41++ {
            var v42 string = ""
            _ = fi41
            _ = v42
        }
        for fi43 := 0; fi43 < 9; fi43++ {
            for fi44 := 0; fi44 < 14; fi44++ {
                v45 := GenS2{F0: 0, F1: (v4 && v4), F2: 407}
                v46 := GenS1{F0: (v3 >> ((653) & 63)), F1: 0}
                fmt.Println(v4)
                _ = fi44
                _ = v45
                _ = v46
            }
            fmt.Println(fi43, v4)
            _ = fi43
        }
    }
    var v47 uint32 = 552
    v47 += v47
    for fi48 := 0; fi48 < 7; fi48++ {
        v49 := v2["foo bar"]
        delete(v2, v49)
        var v50 string = (((v49 + v49) + v49) + v49)
        _ = fi48
        _ = v49
        _ = v50
    }
    for fi51 := 0; fi51 < 6; fi51++ {
        for fi52 := 0; fi52 < 1; fi52++ {
            var v53 uint64 = 42
            if ((-2.25 != 1e10) || ((fi52 >= -9223372036854775808) || (((float64(3.14159) - 0.1) + (-2.25 * -2.25)) >= 1e10))) {
                v4 = v4
                v4 = (!(!v4))
            }
            _ = fi52
            _ = v53
        }
        _ = fi51
    }
    v54 := v2[((((("foo bar" + "a") + "foo bar") + ("a" + "\n\t")) + "") + ("\n\t" + "a"))]
    var v55 uint64 = 1
    v56 := GenS4{F0: v3, F1: (v55 | 2), F2: 32767}
    if (v47 >= v47) {
        if (false || true) {
            var v57 int32 = 1
            if (65535 <= 0) {
                var v58 float64 = -2.25
                v4 = v4
                _ = v58
            }
            var v59 uint16 = 65535
            _ = v57
            _ = v59
        }
        var v60 string = v54
        _ = v60
    } else {
        fmt.Println(v47, v55)
    }
    v54 = ""
    v54 = v54
    v3 = v3
    v4 = (127 >= 0)
    for fi61 := 0; fi61 < 14; fi61++ {
        switch v54 {
        case "hello":
            v55 += v55
        case "":
            v55 += (^(v55))
        default:
            for fi62 := 0; fi62 < 13; fi62++ {
                var v63 bool = ((^(v3)) != v3)
                v64 := GenS4{F0: (v3 + v3), F1: 2, F2: 32768}
                _ = fi62
                _ = v63
                _ = v64
            }
        }
        fmt.Println(v3, v3)
        v65 := GenS3{F0: "a", F1: 4294967295}
        _ = fi61
        _ = v65
    }
    if true {
        v3 -= v3
    } else {
        var v66 uint16 = 2
        _ = v66
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v54
    _ = v55
    _ = v56
}
