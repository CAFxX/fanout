package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 int8 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 uint32; F1 int32 }
type GenS4 struct { F0 uint32; F1 uint8 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int16; F1 int8 }
type GenS7 struct { F0 int64; F1 int8 }
type GenS8 struct { F0 int16; F1 int64 }
type GenS9 struct { F0 string; F1 int64; F2 int32 }
type GenS10 struct { F0 uint; F1 uint8; F2 int32 }
type GenS11 struct { F0 int }
func genFn0() uint {
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 uint8 = 143
        v7 = v7
        for fi8 := 0; fi8 < 17; fi8++ {
            v7 = (v7 - v7)
            _ = fi8
        }
        _ = fi6
        _ = v7
    }
    switch "\n\t" {
    case "a":
        var v9 uint8 = 2
        _ = v9
    }
    var v10 int64 = 0
    _ = v10
    return 209
}

func main() {
    v1 := []int{32767}
    var v2 int16 = 0
    var v3 bool = (18446744073709551615 > 2)
    var v4 uint64 = 1
    var v5 bool = v3
    var v11 uint8 = 29
    v12 := GenS1{F0: 128, F1: 2147483647, F2: 100}
    var v13 float64 = 1e10
    fmt.Println(v13, v3, v11)
    v1 = append(v1, -808)
    if v3 {
        switch "" {
        case "a":
            fmt.Println(v4, v5, v5)
        case "hello":
            var v14 string = "foo bar"
            _ = v14
        }
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        for fi15 := 0; fi15 < 11; fi15++ {
            if (!true) {
                for fi16 := 0; fi16 < 6; fi16++ {
                    fmt.Println(v2)
                    v3 = (3.14159 != (0.1 + 1.5))
                    _ = fi16
                }
                var v17 int16 = v2
                _ = v17
            } else {
                v2 += (-(v2))
                var v18 int32 = 0
                for fi19 := 0; fi19 < 11; fi19++ {
                    v3 = v5
                    if true {
                        var v20 int32 = int32(v11)
                        v5 = (!(v11 >= (^(v11))))
                        var v21 uint8 = v11
                        _ = v20
                        _ = v21
                    } else {
                        fmt.Println(v2, v13, v11)
                    }
                    _ = fi19
                }
                _ = v18
            }
            _ = fi15
        }
        var v22 uint8 = v11
        _ = v22
    }
    for fi23 := 0; fi23 < 8; fi23++ {
        if v3 {
            for fi24 := 0; fi24 < 1; fi24++ {
                if (!(v4 != 0)) {
                    if ((((221 >= 0) || v5) && (v13 <= (-2.25 / -2.25))) && (false || (!(-1 > 0)))) {
                        v25 := GenS2{F0: 42, F1: -104}
                        if true {
                            switch fi23 {
                            case 291:
                                fmt.Println(v11, v5)
                            case 32768:
                                if v3 {
                                    if v3 {
                                        v3 = (255 >= 32767)
                                        v5 = true
                                    }
                                }
                            case 2147483647:
                                v26 := GenS3{F0: 0, F1: 1}
                                _ = v26
                            default:
                                v4 -= 42
                            }
                            v27 := GenS3{F0: 881, F1: -1}
                            var v28 uint32 = 2
                            _ = v27
                            _ = v28
                        } else {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        var v29 uint16 = 2
                        v30 := GenS4{F0: 2, F1: uint8(v4)}
                        _ = v25
                        _ = v29
                        _ = v30
                    } else {
                        if v3 {
                            switch 42 {
                            case 32768:
                                v2 -= (v2 % ((v2) | 1))
                            case 72:
                                v5 = ("foo bar" >= ("hello" + ""))
                            case 42:
                                v3 = (-9223372036854775808 < 9223372036854775807)
                            default:
                                if ((("foo bar" + "hello") + ("\n\t" + "\n\t")) > ("" + ("\n\t" + "foo bar"))) {
                                    v31 := GenS5{F0: 783}
                                    v32 := GenS6{F0: (-(v2)), F1: -128}
                                    v1 = append(v1, fi23)
                                    _ = v31
                                    _ = v32
                                } else {
                                    v33 := GenS7{F0: 1, F1: 127}
                                    v34 := GenS8{F0: (v2 << ((1) & 63)), F1: -1}
                                    switch (fi24 - (fi24 - (fi24 / (((fi24 | fi23)) | 1)))) {
                                    case 128:
                                        v11 -= (v11 | (^(v11)))
                                    default:
                                        v3 = (true || false)
                                    }
                                    v4 += (^(v4))
                                    _ = v33
                                    _ = v34
                                }
                            }
                            var v35 uint16 = 42
                            var v36 float32 = float32((v13 * ((-2.25 * v13) - v13)))
                            switch fi24 {
                            case 51:
                                switch fi23 {
                                case 32767:
                                    v37 := GenS2{F0: v35, F1: 100}
                                    _ = v37
                                case 655:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                case -1:
                                    fmt.Println(v36, v4, v36)
                                default:
                                    var v38 int = 32768
                                    v1 = append(v1, fi23)
                                    _ = v38
                                }
                                var v39 float64 = v13
                                _ = v39
                            case -386:
                                switch "\n\t" {
                                case "foo bar":
                                    if ((238 > uint(v4)) && (v11 != v11)) {
                                        fmt.Println(v11, v13, v11)
                                    }
                                }
                            case 9223372036854775807:
                                switch "a" {
                                case "":
                                    v40 := GenS7{F0: 9223372036854775807, F1: 0}
                                    _ = v40
                                }
                                v1 = append(v1, 2147483647)
                                v1 = append(v1, fi24)
                                v13 = v13
                                v5 = ((!false) || false)
                            default:
                                v11 = v11
                            }
                            _ = v35
                            _ = v36
                        }
                        if false {
                            if ((v2 - 1000) != v2) {
                                if false {
                                    var v41 int = fi23
                                    v11 = (v11 * (v11 * v11))
                                    _ = v41
                                } else {
                                    var v42 int64 = -9223372036854775808
                                    _ = v42
                                }
                            }
                            if false {
                                if false {
                                    fmt.Println(v2, v13)
                                    var v43 int32 = -2147483648
                                    _ = v43
                                }
                            }
                        }
                        if (1 == -22) {
                            v2 = v2
                        }
                    }
                }
                _ = fi24
            }
            fmt.Println(v5, v2)
            v4 = (v4 | 4294967295)
        } else {
            switch "" {
            case "hello":
                if ((v4 % ((v4) | 1)) == v4) {
                    var v44 float64 = (((-2.25 * v13) - v13) - float64(0.1))
                    _ = v44
                } else {
                    if v5 {
                        switch ("\n\t" + "foo bar") {
                        case "a":
                            fmt.Println(v13, v2, v3)
                        default:
                            v45 := GenS9{F0: "a", F1: 1, F2: 2147483647}
                            _ = v45
                        }
                        var v46 float64 = 0.0
                        _ = v46
                    } else {
                        var v47 bool = ((v4 == v4) || v5)
                        if v5 {
                            if len(v1) > 0 {
                                v1[0] = (fi23 >> ((255) & 63))
                            }
                            for fi48 := 0; fi48 < 9; fi48++ {
                                if v47 {
                                    if (!v5) {
                                        var v49 uint16 = 911
                                        v50 := GenS10{F0: 1, F1: (v11 ^ v11), F2: int32(v49)}
                                        if true {
                                            v49 = uint16(v11)
                                        } else {
                                            _ = len(v1)
                                            v4 = v4
                                        }
                                        _ = v49
                                        _ = v50
                                    } else {
                                        v51 := GenS11{F0: fi48}
                                        if v3 {
                                            switch fi23 {
                                            case 32767:
                                                fmt.Println(fi23)
                                            default:
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            }
                                        } else {
                                            var v52 string = "hello"
                                            if (3.14159 >= 1e10) {
                                                var v53 uint32 = 0
                                                v2 += 528
                                                _ = v53
                                            } else {
                                                v52 = v52
                                                var v54 string = (v52 + v52)
                                                var v55 float32 = (((-2.25 + (-2.25 + float32(0.1))) * 0.0) + float32(((v13 + v13) + v13)))
                                                _ = v54
                                                _ = v55
                                            }
                                            if (!(!((true && v5) || v3))) {
                                                var v56 uint16 = 1
                                                v1 = append(v1, 42)
                                                _ = v56
                                            }
                                            _ = v52
                                        }
                                        _ = v51
                                    }
                                } else {
                                    v13 -= v13
                                    v47 = v3
                                    v2 += 1
                                }
                                v11 = v11
                                _ = fi48
                            }
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = v47
                    }
                }
            case "foo bar":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        }
        v57 := GenS11{F0: int(v4)}
        _ = fi23
        _ = v57
    }
    if (v5 || (v3 || (!v3))) {
        var v58 uint = 8
        switch "a" {
        case "a":
            var v59 int = 201
            _ = v59
        }
        v3 = (!true)
        _ = v58
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v1 = append(v1, 42)
    fmt.Println(v2, v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
