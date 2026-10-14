package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 int8; F1 int32; F2 int32 }
type GenS3 struct { F0 uint32; F1 uint64; F2 uint64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 bool; F1 int; F2 int16 }
type GenS7 struct { F0 int16 }
func genFn0() uint64 {
    if ((!(520 != 87)) || true) {
        if (42 != 32767) {
            if (42 > 2) {
                var v5 string = ""
                var v6 uint16 = 255
                _ = v5
                _ = v6
            } else {
                var v7 int8 = -1
                if ((false && (2147483647 == -671)) || true) {
                    v7 += 0
                } else {
                    if false {
                        v7 += v7
                        switch "hello" {
                        case "foo bar":
                            var v8 int = 689
                            _ = v8
                        case "a":
                            fmt.Println(v7)
                        default:
                            v9 := GenS1{F0: 255, F1: 2147483647}
                            _ = v9
                        }
                        var v10 bool = (765 > 4294967295)
                        _ = v10
                    }
                    var v11 uint8 = 148
                    var v12 string = ""
                    _ = v11
                    _ = v12
                }
                v13 := []int{-472}
                _ = v13
                _ = v7
            }
            var v14 float64 = ((-2.25 - 1e10) - (1e10 * (1e10 * 1.5)))
            _ = v14
        }
        var v15 string = "\n\t"
        var v16 int32 = 42
        _ = v15
        _ = v16
    }
    v17 := []int{2147483647}
    if (64 > 2) {
        switch "a" {
        case "foo bar":
            if false {
                _ = len(v17)
                if true {
                    var v18 int64 = 9223372036854775807
                    var v19 int32 = -2147483648
                    if false {
                        var v20 int8 = -13
                        _ = v20
                    }
                    _ = v18
                    _ = v19
                }
            } else {
                for fi21 := 0; fi21 < 2; fi21++ {
                    v22 := GenS2{F0: 42, F1: 518, F2: 1}
                    if len(v17) >= 2 {
                        v17 = v17[1:len(v17)]
                    }
                    if (-32768 > -32768) {
                        var v23 int32 = 2147483647
                        if (18446744073709551615 > 255) {
                            var v24 bool = true
                            v24 = false
                            switch "a" {
                            case "foo bar":
                                v24 = v24
                            default:
                                v25 := GenS2{F0: 42, F1: v23, F2: (v23 << ((0) & 63))}
                                _ = v25
                            }
                            _ = v24
                        } else {
                            var v26 float32 = 3.14159
                            var v27 uint8 = 0
                            var v28 uint8 = (v27 - (v27 & 139))
                            _ = v26
                            _ = v27
                            _ = v28
                        }
                        var v29 bool = (((0.0 - 3.14159) / (0.0 + ((1e10 - float32((((-2.25 / (3.14159 + (((((-2.25 - 1e10) / 1e10) + 1.5) - float64((0.1 * 0.0))) * 1.5))) - (1.5 * float64((-2.25 - 0.1)))) - (-2.25 * 1e10)))) * 3.14159))) == (((3.14159 / 3.14159) / 3.14159) / (1.5 - (0.1 - 0.0))))
                        _ = v23
                        _ = v29
                    }
                    _ = fi21
                    _ = v22
                }
            }
        case "\n\t":
            var v30 int = 42
            _ = v30
        }
        if ((("foo bar" + "hello") + "foo bar") >= ("\n\t" + "\n\t")) {
            for fi31 := 0; fi31 < 6; fi31++ {
                for fi32 := 0; fi32 < 2; fi32++ {
                    fmt.Println(fi31)
                    var v33 float32 = 3.14159
                    _ = fi32
                    _ = v33
                }
                fmt.Println(fi31)
                var v34 uint = 716
                _ = fi31
                _ = v34
            }
        } else {
            var v35 uint16 = 128
            fmt.Println(v35)
            _ = v35
        }
        if true {
            var v36 string = "\n\t"
            for fi37 := 0; fi37 < 7; fi37++ {
                var v38 int16 = 42
                switch 2147483647 {
                case 9223372036854775807:
                    v38 -= v38
                case 596:
                    var v39 string = "hello"
                    _ = v39
                default:
                    for fi40 := 0; fi40 < 2; fi40++ {
                        v17 = append(v17, (fi40 + fi40))
                        var v41 uint = 259
                        var v42 string = (v36 + (v36 + v36))
                        _ = fi40
                        _ = v41
                        _ = v42
                    }
                }
                if len(v17) >= 2 {
                    v17 = v17[1:len(v17)]
                }
                _ = fi37
                _ = v38
            }
            v43 := GenS3{F0: 979, F1: 154, F2: 26}
            _ = v36
            _ = v43
        }
    }
    if len(v17) > 0 {
        v17[0] = -95
    }
    if len(v17) >= 2 {
        v17 = v17[1:len(v17)]
    }
    _ = v17
    return 0
}

func main() {
    v1 := []string{("hello" + "\n\t"), "a"}
    var v2 string = (("\n\t" + ("hello" + "")) + (("\n\t" + "a") + "\n\t"))
    var v3 string = v2
    v4 := make(map[string]uint32)
    v3 = v3
    fmt.Println(v3)
    fmt.Println(v3, v2)
    for fi44 := 0; fi44 < 4; fi44++ {
        v2 = ("" + v2)
        v45 := GenS4{F0: 0}
        var v46 int8 = 100
        _ = fi44
        _ = v45
        _ = v46
    }
    var v47 uint8 = 255
    v1 = append(v1, "foo bar")
    v47 = v47
    var v48 string = ""
    if false {
        if true {
            var v49 int = -1
            fmt.Println(v47, v47)
            if (!((1 >= -253) && true)) {
                v50 := GenS5{F0: ""}
                _ = v50
            }
            _ = v49
        } else {
            var v51 uint8 = v47
            switch v3 {
            case "":
                v3 = v2
            case "a":
                var v52 int = -373
                _ = v52
            case "foo bar":
                for fi53 := 0; fi53 < 4; fi53++ {
                    v3 = v2
                    _ = fi53
                }
            }
            _ = v51
        }
    } else {
        v54 := GenS6{F0: true, F1: -861, F2: -1}
        _ = v54
    }
    for fi55 := 0; fi55 < 2; fi55++ {
        var v56 uint16 = 266
        v57 := GenS7{F0: -1}
        _ = fi55
        _ = v56
        _ = v57
    }
    delete(v4, ("" + v2))
    if false {
        for fi58 := 0; fi58 < 14; fi58++ {
            if true {
                v59 := GenS7{F0: 0}
                var v60 uint8 = v47
                _ = v59
                _ = v60
            }
            _ = fi58
        }
    }
    v61 := GenS2{F0: 127, F1: 2147483647, F2: -1}
    if ((0 < 1) && (((v47 ^ 55) > (v47 << ((475) & 63))) || true)) {
        v62 := GenS1{F0: (v47 / ((v47) | 1)), F1: 1}
        switch (("foo bar" + "a") + v3) {
        case "foo bar":
            var v63 float32 = 0.1
            _ = v63
        case "a":
            if false {
                var v64 float64 = (1e10 + -2.25)
                _ = v64
            }
        }
        if true {
            var v65 uint = 0
            _ = v65
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v66 := GenS2{F0: 0, F1: -1, F2: -1}
            v67 := GenS4{F0: v47}
            _ = v66
            _ = v67
        }
        _ = v62
    }
    var v68 int = 4294967295
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v61
    _ = v68
}
