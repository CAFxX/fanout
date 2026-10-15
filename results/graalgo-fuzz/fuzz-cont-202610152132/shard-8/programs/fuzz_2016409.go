package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 uint32; F1 int8; F2 string }
type GenS3 struct { F0 int64; F1 int32; F2 int32 }
type GenS4 struct { F0 int; F1 uint64; F2 bool }
type GenS5 struct { F0 int; F1 uint32; F2 string }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint16; F1 uint16; F2 string }

func main() {
    var v1 string = ((("\n\t" + "foo bar") + "foo bar") + "foo bar")
    var v2 uint64 = 1
    v3 := []int{-87, 4294967295, 32767, 128}
    v4 := []uint32{42, 42}
    v2 = v2
    if len(v4) > 0 {
        v4[0] = 1
    }
    var v5 uint64 = 0
    var v6 bool = true
    for fi7 := 0; fi7 < 8; fi7++ {
        for fi8 := 0; fi8 < 14; fi8++ {
            v9 := GenS1{F0: 0, F1: 80}
            v1 = v1
            v10 := GenS1{F0: uint16(v5), F1: 100}
            _ = fi8
            _ = v10
            _ = v9
        }
        for fi11 := 0; fi11 < 12; fi11++ {
            switch v1 {
            case "\n\t":
                var v12 int8 = 0
                _ = v12
            case "foo bar":
                fmt.Println(fi7)
            default:
                v13 := GenS2{F0: 65535, F1: 0, F2: (v1 + v1)}
                _ = v13
            }
            fmt.Println(v5)
            v5 += (v5 ^ 4294967295)
            _ = fi11
        }
        _ = fi7
    }
    v6 = ((-928 == 1) && false)
    fmt.Println(v1)
    fmt.Println(v1, v5)
    v2 = v2
    if true {
        var v14 uint = 42
        _ = v14
    }
    v1 = v1
    if false {
        for fi15 := 0; fi15 < 8; fi15++ {
            v3 = append(v3, (fi15 / (((fi15 ^ 422)) | 1)))
            var v16 uint32 = 1
            _ = fi15
            _ = v16
        }
        v6 = v6
    } else {
        for fi17 := 0; fi17 < 3; fi17++ {
            for fi18 := 0; fi18 < 6; fi18++ {
                if v6 {
                    v1 = "hello"
                } else {
                    switch (fi18 + (-(fi17))) {
                    case -568:
                        v5 = 482
                    case 9223372036854775807:
                        var v19 int8 = 69
                        _ = v19
                    }
                }
                v20 := GenS1{F0: 32767, F1: 0}
                var v21 int32 = 1
                _ = fi18
                _ = v20
                _ = v21
            }
            if (-776 < 75) {
                if v6 {
                    v22 := GenS2{F0: 773, F1: -90, F2: v1}
                    if (1 < 609) {
                        var v23 float64 = 3.14159
                        _ = v23
                    } else {
                        if true {
                            switch (fi17 >> ((559) & 63)) {
                            case 0:
                                v1 = (v1 + "")
                            case 9223372036854775807:
                                for fi24 := 0; fi24 < 15; fi24++ {
                                    var v25 float64 = (0.0 + 1e10)
                                    _ = fi24
                                    _ = v25
                                }
                            case 42:
                                for fi26 := 0; fi26 < 9; fi26++ {
                                    if v6 {
                                        var v27 int16 = -1
                                        switch (fi26 << ((42) & 63)) {
                                        case 128:
                                            v3 = append(v3, (-(fi17)))
                                        case 1:
                                            v28 := GenS3{F0: 42, F1: -85, F2: -2147483648}
                                            _ = v28
                                        }
                                        switch v1 {
                                        case "a":
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        case "\n\t":
                                            var v29 bool = v6
                                            _ = v29
                                        default:
                                            var v30 float64 = float64(0.0)
                                            _ = v30
                                        }
                                        if v6 {
                                            switch "foo bar" {
                                            case "foo bar":
                                                fmt.Println(v27)
                                            case "a":
                                                var v31 int32 = 881
                                                _ = v31
                                            }
                                        } else {
                                            var v32 int32 = 42
                                            v33 := GenS4{F0: fi26, F1: v2, F2: v6}
                                            if ((!(v6 || v6)) || v6) {
                                                v34 := GenS1{F0: 0, F1: -1}
                                                _ = v34
                                            }
                                            var v35 uint = 255
                                            if len(v3) > 0 {
                                                v3[0] = fi26
                                            }
                                            _ = v32
                                            _ = v33
                                            _ = v35
                                        }
                                        _ = v27
                                    }
                                    var v36 float64 = 0.0
                                    _ = fi26
                                    _ = v36
                                }
                            }
                            var v37 int64 = -466
                            if v6 {
                                if (1 != -912) {
                                    switch (fi17 | 1) {
                                    case 43:
                                        var v38 uint8 = 42
                                        _ = v38
                                    default:
                                        var v39 float64 = (3.14159 - 1e10)
                                        _ = v39
                                    }
                                } else {
                                    if (fi17 >= (fi17 << ((42) & 63))) {
                                        if (!("a" <= v1)) {
                                            var v40 int = 128
                                            _ = v40
                                        }
                                        v2 = 0
                                    } else {
                                        if (!v6) {
                                            var v41 uint16 = 32767
                                            _ = v41
                                        } else {
                                            v42 := GenS1{F0: 42, F1: -11}
                                            _ = len(v3)
                                            if (1e10 >= -2.25) {
                                                for fi43 := 0; fi43 < 7; fi43++ {
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    _ = fi43
                                                }
                                            } else {
                                                switch (((v1 + ("foo bar" + ("foo bar" + v1))) + "\n\t") + "a") {
                                                case "hello":
                                                    if (true && (!false)) {
                                                        for fi44 := 0; fi44 < 9; fi44++ {
                                                            v1 = "hello"
                                                            v45 := GenS5{F0: (fi44 * fi44), F1: 426, F2: ("foo bar" + v1)}
                                                            v2 = 2
                                                            _ = fi44
                                                            _ = v45
                                                        }
                                                    }
                                                case "":
                                                    v46 := GenS6{F0: -945}
                                                    _ = v46
                                                default:
                                                    for fi47 := 0; fi47 < 10; fi47++ {
                                                        var v48 float32 = 1.5
                                                        v6 = (224 == 772)
                                                        _ = fi47
                                                        _ = v48
                                                    }
                                                }
                                                v49 := GenS5{F0: (fi17 * (fi17 * fi17)), F1: 0, F2: (v1 + "a")}
                                                v37 -= (-(v37))
                                                _ = v49
                                            }
                                            _ = v42
                                        }
                                    }
                                    var v50 int16 = 1
                                    _ = v50
                                }
                                if v6 {
                                    v6 = (75 != 1)
                                } else {
                                    v1 = "a"
                                    _ = len(v4)
                                }
                            } else {
                                if len(v3) > 0 {
                                    v3[0] = -9223372036854775808
                                }
                                v51 := GenS6{F0: 852}
                                _ = v51
                            }
                            _ = v37
                        } else {
                            var v52 int16 = -449
                            _ = v52
                        }
                        if len(v3) > 0 {
                            v3[0] = (^(fi17))
                        }
                    }
                    _ = v22
                } else {
                    switch fi17 {
                    case -667:
                        var v53 uint32 = 1
                        _ = v53
                    default:
                        if len(v3) > 0 {
                            v3[0] = fi17
                        }
                    }
                    var v54 float32 = 1e10
                    _ = v54
                }
                for fi55 := 0; fi55 < 3; fi55++ {
                    var v56 int32 = 42
                    switch v1 {
                    case "\n\t":
                        switch (fi55 / (((fi55 * (fi55 - (fi55 / (((fi17 - (fi17 % (((fi17 - (fi55 & fi17))) | 1)))) | 1))))) | 1)) {
                        case 4294967295:
                            v57 := GenS4{F0: fi55, F1: v2, F2: v6}
                            _ = v57
                        default:
                            if v6 {
                                var v58 float32 = 1.5
                                var v59 uint32 = 0
                                v3 = append(v3, (fi55 * fi55))
                                _ = v58
                                _ = v59
                            } else {
                                var v60 uint = 255
                                if len(v4) > 0 {
                                    v4[0] = 75
                                }
                                _ = v60
                            }
                        }
                    case "hello":
                        if (1 != -1) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            if ((-1 > -697) && v6) {
                                var v61 uint8 = 107
                                v56 -= (v56 * v56)
                                _ = v61
                            }
                        } else {
                            v4 = append(v4, 0)
                        }
                    case "a":
                        switch fi55 {
                        case -552:
                            var v62 uint = 0
                            _ = v62
                        case 9223372036854775807:
                            if len(v4) > 0 {
                                v4[0] = 42
                            }
                        default:
                            v5 = (v5 / ((v2) | 1))
                        }
                    default:
                        var v63 int8 = 100
                        _ = v63
                    }
                    v1 = v1
                    _ = fi55
                    _ = v56
                }
            }
            _ = fi17
        }
        for fi64 := 0; fi64 < 15; fi64++ {
            var v65 float32 = 1e10
            _ = fi64
            _ = v65
        }
    }
    switch 2147483647 {
    case 32767:
        v1 = (((v1 + "\n\t") + "a") + "foo bar")
    case 1:
        v6 = (255 <= 1)
    }
    for fi66 := 0; fi66 < 17; fi66++ {
        if v6 {
            v1 = v1
        }
        if true {
            v67 := GenS7{F0: v6}
            _ = len(v4)
            for fi68 := 0; fi68 < 18; fi68++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v69 int16 = 0
                fmt.Println(fi68)
                _ = fi68
                _ = v69
            }
            _ = v67
        } else {
            switch v1 {
            case "":
                fmt.Println(fi66, v1, v5)
            case "foo bar":
                var v70 int64 = 678
                _ = v70
            default:
                v71 := GenS1{F0: 1, F1: 100}
                _ = v71
            }
            for fi72 := 0; fi72 < 2; fi72++ {
                var v73 int8 = -1
                if len(v4) > 0 {
                    v4[0] = 24
                }
                var v74 uint = 65535
                v74 -= (v74 / ((v74) | 1))
                _ = fi72
                _ = v73
                _ = v74
            }
        }
        _ = fi66
    }
    _ = len(v3)
    if (((v6 && (42 >= 0)) || v6) && (!v6)) {
        for fi75 := 0; fi75 < 18; fi75++ {
            fmt.Println(v6, v6, v6)
            _ = fi75
        }
        if (v6 || (!false)) {
            v1 = (("a" + "a") + "\n\t")
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        } else {
            v6 = (!(true && ((1 > 17) || (!true))))
            var v76 uint64 = v2
            v3 = append(v3, 230)
            _ = v76
        }
    } else {
        var v77 int8 = -128
        fmt.Println(v1)
        var v78 uint8 = 60
        _ = v77
        _ = v78
    }
    v5 = v2
    v79 := GenS8{F0: 65535, F1: 1, F2: v1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v79
}
