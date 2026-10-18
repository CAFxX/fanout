package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 int64; F1 uint32; F2 int8 }
type GenS4 struct { F0 uint; F1 int16 }
type GenS5 struct { F0 uint64; F1 bool; F2 bool }
type GenS6 struct { F0 bool; F1 uint16 }
type GenS7 struct { F0 uint16; F1 uint8; F2 int64 }
type GenS8 struct { F0 int8 }
func genFn0() uint16 {
    switch (("" + "\n\t") + "\n\t") {
    case "a":
        var v6 string = "foo bar"
        _ = v6
    case "\n\t":
        var v7 uint16 = 284
        _ = v7
    case "hello":
        v8 := []int{905, 127, -268, 0}
        _ = v8
    default:
        if (!((-32768 > 32767) || (353 > 801))) {
            var v9 bool = (!false)
            fmt.Println(v9)
            _ = v9
        } else {
            var v10 float32 = (-2.25 + (0.1 / 3.14159))
            for fi11 := 0; fi11 < 13; fi11++ {
                v12 := []int{fi11, 42, (fi11 << ((42) & 63)), (fi11 >> ((255) & 63))}
                _ = fi11
                _ = v12
            }
            _ = v10
        }
    }
    var v13 float32 = 0.0
    for fi14 := 0; fi14 < 6; fi14++ {
        v15 := []int{-440, 128, 1, fi14}
        v16 := GenS1{F0: 42, F1: 42}
        for fi17 := 0; fi17 < 17; fi17++ {
            if false {
                v13 -= 1e10
                var v18 int8 = 0
                _ = v18
            }
            var v19 int8 = -128
            _ = fi17
            _ = v19
        }
        _ = fi14
        _ = v15
        _ = v16
    }
    var v20 int8 = -41
    v21 := []int{}
    _ = v13
    _ = v20
    _ = v21
    return 0
}

func main() {
    var v1 int8 = 0
    var v2 uint8 = 115
    var v3 int = 127
    var v4 int = -2147483648
    v5 := []int{(v4 ^ (v4 - v4))}
    if false {
        v3 -= v4
    } else {
        switch v4 {
        case -9223372036854775808:
            if true {
                v22 := GenS1{F0: 1, F1: 630}
                _ = v22
            } else {
                fmt.Println(v3, v3)
            }
        case 42:
            v23 := GenS2{F0: 0, F1: 0}
            _ = v23
        case -744:
            if true {
                if len(v5) > 0 {
                    v5[0] = (v4 * -1)
                }
            } else {
                if true {
                    var v24 uint32 = 2147483648
                    _ = v24
                } else {
                    v25 := GenS2{F0: 330, F1: 32767}
                    for fi26 := 0; fi26 < 7; fi26++ {
                        var v27 uint = 517
                        _ = fi26
                        _ = v27
                    }
                    _ = v25
                }
            }
        }
        var v28 uint = 460
        _ = v28
    }
    fmt.Println(v2, v1, v2)
    fmt.Println(v3, v2)
    for fi29 := 0; fi29 < 3; fi29++ {
        if true {
            if false {
                _ = len(v5)
                if true {
                    var v30 int32 = -336
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    if false {
                        v31 := GenS3{F0: 0, F1: 670, F2: 127}
                        _ = v31
                    }
                    _ = v30
                }
            }
            var v32 uint8 = (v2 | v2)
            _ = v32
        } else {
            for fi33 := 0; fi33 < 10; fi33++ {
                if (!true) {
                    if false {
                        if (!(!false)) {
                            if (v1 != v1) {
                                if len(v5) > 0 {
                                    v5[0] = fi33
                                }
                                var v34 float64 = 1.5
                                _ = v34
                            } else {
                                var v35 int16 = 32767
                                v35 = (-(v35))
                                var v36 uint8 = (v2 * 255)
                                _ = v35
                                _ = v36
                            }
                            v37 := GenS4{F0: 2, F1: 0}
                            if false {
                                v1 = v1
                                var v38 int8 = v1
                                _ = v38
                            } else {
                                v39 := GenS3{F0: 1, F1: 2147483648, F2: v1}
                                v4 = fi33
                                _ = v39
                            }
                            if true {
                                v40 := GenS4{F0: 2, F1: 32767}
                                var v41 int64 = 547
                                switch "hello" {
                                case "hello":
                                    var v42 int = v4
                                    _ = v42
                                case "":
                                    if (("foo bar" < ("hello" + "foo bar")) && true) {
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        v43 := GenS3{F0: v41, F1: 2, F2: int8(v2)}
                                        _ = v43
                                    }
                                case "a":
                                    v3 += fi33
                                default:
                                    v44 := GenS4{F0: 614, F1: 211}
                                    _ = v44
                                }
                                _ = v40
                                _ = v41
                            } else {
                                var v45 int64 = -9223372036854775808
                                v1 = (v1 ^ (v1 & v1))
                                _ = len(v5)
                                _ = v45
                            }
                            _ = v37
                        }
                    } else {
                        fmt.Println(fi33)
                        if false {
                            var v46 int = v4
                            var v47 int16 = 1
                            _ = v46
                            _ = v47
                        } else {
                            v2 -= 240
                        }
                    }
                    var v48 float32 = -2.25
                    _ = v48
                }
                var v49 uint16 = genFn0()
                _ = fi33
                _ = v49
            }
            var v50 int8 = v1
            if true {
                for fi51 := 0; fi51 < 16; fi51++ {
                    var v52 float32 = 1e10
                    if false {
                        var v53 bool = true
                        v3 -= (fi51 >> ((1) & 63))
                        _ = v53
                    } else {
                        v54 := GenS5{F0: 271, F1: (((!(!(v50 > (v1 - (v50 ^ -1))))) && (!false)) && (-246 > 1)), F2: (!(814 != 1))}
                        v2 = (v2 * (v2 & v2))
                        _ = v54
                    }
                    var v55 float32 = v52
                    _ = fi51
                    _ = v52
                    _ = v55
                }
                var v56 int = 9223372036854775807
                v5 = append(v5, v56)
                _ = v56
            } else {
                v57 := GenS6{F0: ("a" < "foo bar"), F1: 864}
                var v58 uint32 = 2
                _ = v57
                _ = v58
            }
            _ = v50
        }
        _ = fi29
    }
    v1 = 113
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if false {
        for fi59 := 0; fi59 < 10; fi59++ {
            var v60 int = fi59
            _ = fi59
            _ = v60
        }
        fmt.Println(v2, v4)
        v61 := GenS7{F0: 674, F1: 42, F2: 9223372036854775807}
        _ = v61
    }
    v1 = -69
    if true {
        if false {
            switch "a" {
            case "hello":
                if false {
                    var v62 bool = false
                    _ = v62
                }
            case "foo bar":
                fmt.Println(v1, v4)
            case "\n\t":
                fmt.Println(v2, v4, v4)
            }
            for fi63 := 0; fi63 < 7; fi63++ {
                for fi64 := 0; fi64 < 16; fi64++ {
                    v5 = append(v5, fi64)
                    _ = fi64
                }
                var v65 string = "hello"
                switch v65 {
                case "foo bar":
                    fmt.Println(v2, v3, v4)
                case "\n\t":
                    v66 := GenS8{F0: (v1 % ((v1) | 1))}
                    _ = v66
                case "":
                    if len(v5) > 0 {
                        v5[0] = v4
                    }
                default:
                    v3 += v3
                }
                _ = fi63
                _ = v65
            }
            v67 := GenS8{F0: int8(v2)}
            _ = v67
        } else {
            for fi68 := 0; fi68 < 20; fi68++ {
                var v69 int16 = int16(v1)
                fmt.Println(v4, v3)
                _ = fi68
                _ = v69
            }
        }
    }
    fmt.Println(v4, v2, v3)
    v1 -= -85
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
