package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8; F1 int8 }
type GenS5 struct { F0 int64 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]string)
    v3 := []int{-2147483648, 133, -9223372036854775808, 127}
    var v4 int64 = 481
    var v5 float64 = float64(0.0)
    var v6 uint16 = 0
    v7 := GenS1{F0: 2147483647}
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 int = fi8
        if true {
            v9 += v9
            v6 += (v6 & v6)
        } else {
            switch (fi8 ^ fi8) {
            case 128:
                fmt.Println(v4, fi8)
            case -957:
                if len(v3) > 0 {
                    v3[0] = (fi8 ^ fi8)
                }
            case -2147483648:
                if ((false && false) && (724 < 2147483648)) {
                    if true {
                        switch -9223372036854775808 {
                        case 319:
                            if (-33 < 100) {
                                fmt.Println(v9, v4, fi8)
                                var v10 bool = false
                                v5 = v5
                                _ = v10
                            }
                        case 32768:
                            fmt.Println(fi8, v6, v5)
                        case -1:
                            var v11 uint16 = v6
                            _ = v11
                        default:
                            _ = len(v3)
                        }
                        v9 += (fi8 & (v9 - fi8))
                        v5 += v5
                    }
                    if ("" > "") {
                        v9 += int(v4)
                        fmt.Println(v6, v5, fi8)
                        var v12 uint32 = 1
                        _ = v12
                    } else {
                        var v13 string = ("hello" + ("a" + ""))
                        v13 = v13
                        _ = v13
                    }
                    v4 = (^(v4))
                }
            }
        }
        _ = fi8
        _ = v9
    }
    var v14 uint8 = 255
    if (1.5 > 0.1) {
        v6 -= 0
    } else {
        for fi15 := 0; fi15 < 20; fi15++ {
            v16 := GenS2{F0: (fi15 << ((1) & 63)), F1: 42}
            _ = fi15
            _ = v16
        }
        for fi17 := 0; fi17 < 14; fi17++ {
            v18 := v1[32767]
            if (!v18) {
                v19 := GenS1{F0: 2147483648}
                _ = v19
            }
            var v20 float64 = v5
            _ = fi17
            _ = v18
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 11; fi21++ {
        if true {
            var v22 uint64 = 203
            for fi23 := 0; fi23 < 10; fi23++ {
                switch "" {
                case "hello":
                    var v24 uint8 = (v14 - (^(v14)))
                    _ = v24
                }
                if (!false) {
                    v14 = (v14 - (v14 % ((v14) | 1)))
                    v6 = v6
                }
                v25 := GenS2{F0: int(v4), F1: -125}
                _ = fi23
                _ = v25
            }
            fmt.Println(v6)
            _ = v22
        } else {
            if (((v5 > (v5 + 3.14159)) || true) && (42 >= -2147483648)) {
                var v26 int32 = 2147483647
                _ = v26
            }
            for fi27 := 0; fi27 < 9; fi27++ {
                v6 = v6
                if false {
                    var v28 float32 = 0.1
                    fmt.Println(v5)
                    if true {
                        v29 := GenS1{F0: 2}
                        _ = len(v1)
                        fmt.Println(v28, v28)
                        _ = v29
                    } else {
                        var v30 uint16 = 581
                        v14 = (v14 | v14)
                        if true {
                            if ((-(fi27)) > fi21) {
                                var v31 float32 = v28
                                if false {
                                    fmt.Println(v5)
                                    if true {
                                        v32 := GenS3{F0: v14}
                                        switch (fi21 & (fi27 * (fi21 | int(v30)))) {
                                        case -579:
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        }
                                        var v33 bool = false
                                        _ = v32
                                        _ = v33
                                    }
                                    var v34 int64 = -9223372036854775808
                                    _ = v34
                                }
                                _ = v31
                            } else {
                                fmt.Println(fi27, v30, fi21)
                            }
                            v35 := GenS1{F0: 4294967295}
                            _ = v35
                        } else {
                            switch fi27 {
                            case -732:
                                var v36 string = "hello"
                                _ = v36
                            case 32768:
                                switch (fi21 << ((0) & 63)) {
                                case 32768:
                                    if true {
                                        v37 := GenS2{F0: -351, F1: 127}
                                        if (0 < 261) {
                                            fmt.Println(v4)
                                            v5 += 0.0
                                        }
                                        _ = v37
                                    } else {
                                        if len(v3) > 0 {
                                            v3[0] = (^(fi27))
                                        }
                                    }
                                }
                                v38 := GenS1{F0: 2147483647}
                                _ = v38
                            }
                            fmt.Println(v4, v28, v30)
                            v6 = v6
                        }
                        _ = v30
                    }
                    _ = v28
                }
                fmt.Println(v14, fi27)
                _ = fi27
            }
            if false {
                fmt.Println(fi21, v5)
                var v39 int64 = 42
                _ = v39
            } else {
                if len(v3) > 0 {
                    v3[0] = (fi21 & fi21)
                }
                switch ("foo bar" + ("\n\t" + "\n\t")) {
                case "hello":
                    var v40 int = -2147483648
                    _ = v40
                default:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                v41 := GenS1{F0: 65535}
                _ = v41
            }
        }
        for fi42 := 0; fi42 < 9; fi42++ {
            var v43 float32 = (3.14159 + ((0.0 / (0.1 - -2.25)) - 0.0))
            v6 -= (v6 + (v6 ^ (^(v6))))
            _ = fi42
            _ = v43
        }
        _ = fi21
    }
    for fi44 := 0; fi44 < 12; fi44++ {
        v45 := GenS2{F0: (^(fi44)), F1: 127}
        _ = fi44
        _ = v45
    }
    if false {
        v4 = -977
        var v46 float32 = (1e10 * (0.0 + (1.5 * 0.1)))
        if (v6 > 1) {
            for fi47 := 0; fi47 < 5; fi47++ {
                var v48 uint8 = (^(v14))
                _ = fi47
                _ = v48
            }
            if false {
                fmt.Println(v5, v6)
                v49 := GenS4{F0: 0, F1: -71}
                _ = v49
            }
        } else {
            for fi50 := 0; fi50 < 11; fi50++ {
                if (v14 != (v14 - v14)) {
                    fmt.Println(v6, v5, v6)
                    for fi51 := 0; fi51 < 11; fi51++ {
                        var v52 int64 = v4
                        _ = fi51
                        _ = v52
                    }
                }
                _ = fi50
            }
            for fi53 := 0; fi53 < 6; fi53++ {
                v54 := GenS3{F0: v14}
                v55 := GenS2{F0: 761, F1: int8(v4)}
                var v56 string = "a"
                _ = fi53
                _ = v54
                _ = v55
                _ = v56
            }
            v14 -= v14
        }
        _ = v46
    } else {
        v4 -= (v4 ^ (^(v4)))
        var v57 int32 = 42
        _ = v57
    }
    _ = len(v3)
    if false {
        _ = len(v3)
        var v58 uint32 = 326
        _ = v58
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v59 int32 = 42
    for fi60 := 0; fi60 < 8; fi60++ {
        for fi61 := 0; fi61 < 2; fi61++ {
            v62 := GenS5{F0: int64(v6)}
            var v63 uint16 = v6
            _ = fi61
            _ = v62
            _ = v63
        }
        v4 -= -1
        _ = fi60
    }
    _ = len(v3)
    var v64 float64 = 3.14159
    v65 := GenS4{F0: -42, F1: 127}
    v59 -= 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v6
    _ = v64
    _ = v65
    _ = v7
}
