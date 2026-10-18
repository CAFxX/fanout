package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64; F1 bool }
type GenS4 struct { F0 uint8; F1 int8; F2 uint }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint32 = 650
    var v3 int64 = 917
    v4 := make(map[string]bool)
    v5 := []int{32767, 1}
    var v6 float64 = (0.0 - (3.14159 / 1e10))
    v3 += -45
    if false {
        for fi7 := 0; fi7 < 16; fi7++ {
            v8 := GenS1{F0: 64, F1: 18446744073709551615, F2: 493}
            var v9 int = (fi7 >> ((488) & 63))
            for fi10 := 0; fi10 < 4; fi10++ {
                v3 += v3
                fmt.Println(v3, v9)
                var v11 int32 = 1
                _ = fi10
                _ = v11
            }
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v12 uint32 = v2
    v13 := v4["foo bar"]
    for fi14 := 0; fi14 < 5; fi14++ {
        for fi15 := 0; fi15 < 14; fi15++ {
            fmt.Println(fi14, v13)
            _ = fi15
        }
        _ = fi14
    }
    v13 = false
    if ((((v3 & v3) >= 42) && (!v13)) && (488 >= 726)) {
        v13 = ((!true) && (-69 != 588))
        switch 88 {
        case -691:
            if v13 {
                if v13 {
                    v12 = v2
                    for fi16 := 0; fi16 < 12; fi16++ {
                        fmt.Println(v2)
                        _ = fi16
                    }
                    var v17 uint16 = 0
                    _ = v17
                } else {
                    if true {
                        var v18 string = "hello"
                        var v19 bool = v13
                        _ = v18
                        _ = v19
                    } else {
                        v12 = (^(v2))
                        var v20 uint = uint(v2)
                        if true {
                            switch "\n\t" {
                            case "hello":
                                var v21 int64 = v3
                                _ = v21
                            }
                            if (!v13) {
                                switch "foo bar" {
                                case "foo bar":
                                    v13 = v13
                                case "\n\t":
                                    v22 := GenS2{F0: 11}
                                    _ = v22
                                default:
                                    v2 -= v2
                                }
                                _ = len(v5)
                            }
                        } else {
                            for fi23 := 0; fi23 < 12; fi23++ {
                                var v24 uint8 = 0
                                var v25 float64 = v6
                                _ = len(v1)
                                _ = fi23
                                _ = v24
                                _ = v25
                            }
                            v26 := GenS1{F0: 127, F1: 1, F2: v12}
                            v5 = append(v5, 32767)
                            _ = v26
                        }
                        _ = v20
                    }
                    for fi27 := 0; fi27 < 14; fi27++ {
                        switch 0 {
                        case 686:
                            var v28 uint64 = 4294967295
                            _ = v28
                        case 0:
                            if v13 {
                                for fi29 := 0; fi29 < 14; fi29++ {
                                    var v30 int = 127
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    if ((((452 != v12) || v13) || ((v6 / (1e10 + v6)) > 1.5)) || true) {
                                        if v13 {
                                            if v13 {
                                                var v31 float64 = v6
                                                var v32 int64 = 0
                                                var v33 string = ("a" + ("hello" + "foo bar"))
                                                _ = v31
                                                _ = v32
                                                _ = v33
                                            }
                                            v34 := GenS1{F0: 2, F1: 18446744073709551615, F2: (v2 / ((v12) | 1))}
                                            v2 = v12
                                            _ = v34
                                        } else {
                                            v35 := v1[(("hello" + "") + "\n\t")]
                                            fmt.Println(v30)
                                            if false {
                                                v36 := GenS2{F0: 127}
                                                if true {
                                                    v37 := GenS3{F0: 404, F1: (100 >= 0)}
                                                    _ = v37
                                                } else {
                                                    if (!(true || ((float32(float64(1.5)) + float32(v6)) == 1e10))) {
                                                        var v38 int32 = 0
                                                        v3 -= v3
                                                        switch "hello" {
                                                        case "":
                                                            v35 -= 42
                                                        default:
                                                            if true {
                                                                v30 += 4294967295
                                                                if true {
                                                                    fmt.Println(v2)
                                                                    _ = len(v1)
                                                                    v39 := GenS2{F0: 109}
                                                                    v3 = 9223372036854775807
                                                                    v40 := GenS3{F0: 752, F1: (v2 <= v12)}
                                                                    v41 := GenS4{F0: 127, F1: -128, F2: 427}
                                                                    var v42 uint16 = 255
                                                                    _ = v39
                                                                    _ = v40
                                                                    _ = v41
                                                                    _ = v42
                                                                }
                                                                v30 = (-(v30))
                                                                var v43 bool = (v13 && v13)
                                                                v3 += (v3 | 865)
                                                                if (100 != 0) {
                                                                    fmt.Println(fi27)
                                                                    v30 = (fi27 + (fi29 >> ((0) & 63)))
                                                                } else {
                                                                    v44 := GenS1{F0: 127, F1: 18446744073709551615, F2: (v2 / (((v12 / (((^(v2))) | 1))) | 1))}
                                                                    if ((v43 || v43) || (v13 || false)) {
                                                                        fmt.Println(v2, v13)
                                                                    } else {
                                                                        if len(v5) > 0 {
                                                                            v5[0] = 759
                                                                        }
                                                                        if (42 <= v38) {
                                                                            _ = len(v1)
                                                                        } else {
                                                                            var v45 uint32 = 819
                                                                            v4["\n\t"] = v43
                                                                            _ = v45
                                                                        }
                                                                    }
                                                                    _ = v44
                                                                }
                                                                _ = v43
                                                            } else {
                                                                var v46 uint8 = 2
                                                                var v47 uint = 255
                                                                _ = v46
                                                                _ = v47
                                                            }
                                                        }
                                                        _ = v38
                                                    } else {
                                                        fmt.Println(fi29, v12, v13)
                                                        switch "\n\t" {
                                                        case "foo bar":
                                                            if v13 {
                                                                var v48 uint32 = v35
                                                                v49 := GenS1{F0: 0, F1: 4294967295, F2: (v12 + (v2 * (v35 << ((42) & 63))))}
                                                                _ = v48
                                                                _ = v49
                                                            }
                                                        default:
                                                            var v50 int16 = -32768
                                                            _ = v50
                                                        }
                                                        if false {
                                                            switch ("\n\t" + "\n\t") {
                                                            case "a":
                                                                if v13 {
                                                                    v35 = v2
                                                                    var v51 float32 = 0.0
                                                                    _ = v51
                                                                } else {
                                                                    var v52 int16 = -839
                                                                    _ = v52
                                                                }
                                                                fmt.Println(fi29)
                                                                switch "\n\t" {
                                                                case "":
                                                                    v53 := v4["hello"]
                                                                    _ = v53
                                                                case "foo bar":
                                                                    fmt.Println(v30, v2, fi27)
                                                                default:
                                                                    if (127 < 255) {
                                                                        if false {
                                                                            v6 = 0.0
                                                                            fmt.Println(fi29, v13)
                                                                            v54 := GenS2{F0: 1}
                                                                            _ = v54
                                                                        }
                                                                    } else {
                                                                        var v55 int16 = 32767
                                                                        _ = v55
                                                                    }
                                                                }
                                                            case "\n\t":
                                                                if v13 {
                                                                    var v56 int = (fi29 | (fi29 / ((fi27) | 1)))
                                                                    if v13 {
                                                                        var v57 int = (-(v30))
                                                                        _ = v57
                                                                    } else {
                                                                        if len(v5) > 0 {
                                                                            v5[0] = (v30 % ((fi27) | 1))
                                                                        }
                                                                        v30 += 32768
                                                                        var v58 uint = 747
                                                                        var v59 uint8 = 42
                                                                        _ = v58
                                                                        _ = v59
                                                                    }
                                                                    _ = v56
                                                                } else {
                                                                    if v13 {
                                                                        v4[("a" + "\n\t")] = (false || false)
                                                                        v1["a"] = (v12 + v12)
                                                                        v60 := GenS3{F0: 288, F1: v13}
                                                                        _ = v60
                                                                    } else {
                                                                        if v13 {
                                                                            var v61 int16 = 1
                                                                            _ = v61
                                                                        } else {
                                                                            v2 -= 2147483648
                                                                        }
                                                                        var v62 bool = v13
                                                                        _ = v62
                                                                    }
                                                                    v35 -= (^(v35))
                                                                    v13 = true
                                                                    fmt.Println(v2)
                                                                    fmt.Println(v2, fi27, v30)
                                                                }
                                                            case "foo bar":
                                                                v3 = v3
                                                            default:
                                                                var v63 uint8 = 255
                                                                _ = v63
                                                            }
                                                            v3 = v3
                                                        }
                                                        v64 := GenS1{F0: 23, F1: 245, F2: 4294967295}
                                                        v65 := GenS4{F0: 255, F1: 25, F2: 846}
                                                        _ = v64
                                                        _ = v65
                                                    }
                                                }
                                                _ = v36
                                            }
                                            _ = v35
                                        }
                                    }
                                    _ = fi29
                                    _ = v30
                                }
                            } else {
                                if v13 {
                                    v12 -= 1
                                } else {
                                    var v66 int = fi27
                                    _ = v66
                                }
                            }
                        case 2147483647:
                            var v67 uint = 2
                            _ = v67
                        }
                        _ = fi27
                    }
                }
                var v68 uint8 = 128
                switch "" {
                case "\n\t":
                    v12 += 65535
                default:
                    fmt.Println(v68, v3)
                }
                _ = v68
            }
        default:
            if false {
                for fi69 := 0; fi69 < 13; fi69++ {
                    var v70 int = fi69
                    _ = fi69
                    _ = v70
                }
                if len(v5) > 0 {
                    v5[0] = -2147483648
                }
                if v13 {
                    var v71 string = "foo bar"
                    for fi72 := 0; fi72 < 15; fi72++ {
                        switch fi72 {
                        case -942:
                            v12 = (v12 % (((v12 & (v2 ^ v12))) | 1))
                        }
                        var v73 string = ("hello" + v71)
                        if v13 {
                            fmt.Println(v73, v6)
                        }
                        _ = fi72
                        _ = v73
                    }
                    _ = v71
                }
            } else {
                v2 = 0
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
