package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int; F1 uint8 }
type GenS5 struct { F0 int; F1 uint16 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint8; F1 uint; F2 string }
type GenS8 struct { F0 int32; F1 int8; F2 uint64 }
type GenS9 struct { F0 uint32 }

func main() {
    var v1 int = -9223372036854775808
    var v2 int16 = 32767
    var v3 float64 = 0.1
    var v4 int64 = 0
    var v5 uint8 = 2
    if true {
        for fi6 := 0; fi6 < 13; fi6++ {
            for fi7 := 0; fi7 < 4; fi7++ {
                v3 += v3
                v4 = v4
                var v8 uint32 = 74
                if false {
                    var v9 bool = true
                    var v10 int16 = v2
                    if v9 {
                        v10 += 1000
                    } else {
                        if (42 != 4294967295) {
                            v5 = 42
                        } else {
                            var v11 uint = 1
                            _ = v11
                        }
                        v12 := GenS1{F0: 918, F1: 0}
                        fmt.Println(v9, fi6, v10)
                        _ = v12
                    }
                    var v13 uint = 0
                    v14 := []int{(fi7 + v1)}
                    if v9 {
                        var v15 uint = 42
                        v16 := GenS1{F0: 0, F1: -799}
                        _ = v15
                        _ = v16
                    } else {
                        v17 := GenS1{F0: 0, F1: -1}
                        v10 = -1
                        fmt.Println(v3, fi7, v1)
                        if len(v14) > 0 {
                            v14[0] = int(v10)
                        }
                        _ = v17
                    }
                    _ = v10
                    _ = v13
                    _ = v14
                    _ = v9
                } else {
                    v18 := []int{(v1 << ((0) & 63))}
                    fmt.Println(v3, v3)
                    v4 = (v4 << ((65535) & 63))
                    _ = v18
                }
                switch v1 {
                case 32768:
                    fmt.Println(fi6, fi7, v3)
                default:
                    switch fi6 {
                    case 32768:
                        var v19 int16 = (-(v2))
                        _ = v19
                    }
                }
                _ = fi7
                _ = v8
            }
            fmt.Println(fi6)
            _ = fi6
        }
    } else {
        v20 := []int{(v1 & v1), v1, (^(v1)), (^(v1))}
        v21 := GenS1{F0: 2, F1: 888}
        if true {
            v2 -= (v2 << ((270) & 63))
            if (-33 != 42) {
                var v22 uint32 = 65535
                v3 += (float64(3.14159) - v3)
                if (-32768 == -202) {
                    if true {
                        v20 = append(v20, 32768)
                        fmt.Println(v2, v2)
                        v23 := GenS2{F0: 1}
                        _ = v23
                    } else {
                        var v24 uint32 = v22
                        var v25 float32 = 3.14159
                        if ((34 < (v5 - (v5 - (v5 ^ (v5 >> ((776) & 63)))))) && (0 < 912)) {
                            _ = len(v20)
                            v4 += (v4 / ((42) | 1))
                        } else {
                            if (!(((true && (!true)) && (0 < v2)) || false)) {
                                v4 -= (^(v4))
                                if true {
                                    var v26 uint32 = v22
                                    fmt.Println(v26, v25)
                                    _ = v26
                                } else {
                                    for fi27 := 0; fi27 < 2; fi27++ {
                                        v28 := GenS2{F0: 535}
                                        var v29 int64 = (v4 - v4)
                                        if false {
                                            var v30 bool = ((v1 / ((fi27) | 1)) >= (^(fi27)))
                                            _ = v30
                                        } else {
                                            v1 -= (^(v1))
                                        }
                                        _ = fi27
                                        _ = v28
                                        _ = v29
                                    }
                                }
                                v24 = (v24 - v24)
                            } else {
                                var v31 uint32 = v24
                                _ = len(v20)
                                v32 := GenS3{F0: v4}
                                _ = v31
                                _ = v32
                            }
                            var v33 uint64 = 18446744073709551615
                            var v34 int64 = v4
                            _ = v33
                            _ = v34
                        }
                        _ = v24
                        _ = v25
                    }
                } else {
                    v20 = append(v20, int(v2))
                }
                _ = v22
            } else {
                v4 = v4
                v3 += v3
                for fi35 := 0; fi35 < 11; fi35++ {
                    if false {
                        var v36 int8 = int8(v2)
                        _ = v36
                    }
                    _ = fi35
                }
            }
        }
        _ = v20
        _ = v21
    }
    switch ("hello" + "\n\t") {
    case "":
        if false {
            for fi37 := 0; fi37 < 13; fi37++ {
                var v38 int32 = -952
                v39 := GenS4{F0: 1, F1: (v5 * 44)}
                _ = fi37
                _ = v38
                _ = v39
            }
            for fi40 := 0; fi40 < 15; fi40++ {
                var v41 float32 = -2.25
                if (0 <= (v4 << ((42) & 63))) {
                    for fi42 := 0; fi42 < 10; fi42++ {
                        v1 -= -285
                        switch -2147483648 {
                        case 128:
                            switch "hello" {
                            case "":
                                v43 := []int{fi40, 976, (fi42 ^ v1), int(v4)}
                                v2 = (-(v2))
                                var v44 uint8 = 35
                                _ = v43
                                _ = v44
                            case "\n\t":
                                v45 := GenS2{F0: 42}
                                _ = v45
                            case "hello":
                                if true {
                                    v46 := GenS4{F0: (v1 >> ((65535) & 63)), F1: v5}
                                    var v47 float64 = v3
                                    v47 -= 1e10
                                    _ = v46
                                    _ = v47
                                } else {
                                    if (false && (!true)) {
                                        if false {
                                            var v48 uint = 0
                                            v49 := []int{(fi42 % (((fi42 / ((fi40) | 1))) | 1)), 0, (^(v1))}
                                            _ = v48
                                            _ = v49
                                        }
                                        v50 := GenS5{F0: (fi40 ^ fi42), F1: 255}
                                        _ = v50
                                    } else {
                                        v1 += fi40
                                    }
                                    if (374 < 32) {
                                        v51 := []int{4294967295}
                                        v2 = 32767
                                        _ = v51
                                    } else {
                                        fmt.Println(v2, v2)
                                    }
                                    fmt.Println(v1, v5)
                                }
                            }
                        case -2147483648:
                            if ("" == (("\n\t" + (("a" + "") + "hello")) + ("hello" + "a"))) {
                                if true {
                                    var v52 uint8 = 50
                                    var v53 int32 = -2147483648
                                    if false {
                                        v41 += 0.0
                                        if (203 > v52) {
                                            if (!(false || false)) {
                                                var v54 uint16 = 65535
                                                _ = v54
                                            }
                                            v55 := []int{}
                                            _ = v55
                                        } else {
                                            var v56 float32 = v41
                                            _ = v56
                                        }
                                    } else {
                                        fmt.Println(v1, fi40)
                                        if ((true && false) && false) {
                                            var v57 uint16 = 0
                                            _ = v57
                                        }
                                    }
                                    _ = v52
                                    _ = v53
                                }
                                if true {
                                    var v58 int64 = (v4 - v4)
                                    var v59 int16 = -32768
                                    v5 += v5
                                    _ = v58
                                    _ = v59
                                } else {
                                    v60 := []int{fi40}
                                    v5 = (^(v5))
                                    var v61 bool = (!(1.5 != 3.14159))
                                    _ = v60
                                    _ = v61
                                }
                            }
                        default:
                            v2 = -32768
                        }
                        _ = fi42
                    }
                }
                _ = fi40
                _ = v41
            }
            switch "foo bar" {
            case "foo bar":
                switch "foo bar" {
                case "hello":
                    fmt.Println(v2)
                case "\n\t":
                    v62 := GenS6{F0: 127}
                    _ = v62
                case "a":
                    if ("hello" == "hello") {
                        v2 += (v2 | 1)
                        var v63 int64 = v4
                        _ = v63
                    }
                default:
                    if (!((87 >= 0) && false)) {
                        v4 = v4
                        v64 := GenS7{F0: (v5 | v5), F1: 2, F2: ("" + ("a" + (("" + "\n\t") + "\n\t")))}
                        v65 := GenS1{F0: 0, F1: 42}
                        _ = v64
                        _ = v65
                    }
                }
            case "a":
                v2 = v2
            default:
                if (323 >= uint16(v5)) {
                    var v66 uint8 = v5
                    v3 += (3.14159 / 1e10)
                    _ = v66
                }
            }
        }
    case "hello":
        v67 := []int{(v1 ^ (v1 / ((-1) | 1))), (v1 / ((v1) | 1)), (v1 % ((v1) | 1)), (^(v1))}
        _ = v67
    default:
        switch "\n\t" {
        case "\n\t":
            var v68 int16 = 0
            _ = v68
        default:
            switch v1 {
            case 976:
                if (86 > (^(v5))) {
                    v69 := []int{v1, -48}
                    v70 := GenS8{F0: 2147483647, F1: 0, F2: 2}
                    _ = v69
                    _ = v70
                } else {
                    v71 := GenS9{F0: 1}
                    switch "a" {
                    case "":
                        v3 = v3
                    default:
                        v5 -= (^(v5))
                    }
                    v5 = 255
                    _ = v71
                }
            default:
                v4 -= (v4 ^ v4)
            }
        }
    }
    if (917 != 42) {
        for fi72 := 0; fi72 < 2; fi72++ {
            v73 := GenS9{F0: 1}
            _ = fi72
            _ = v73
        }
        fmt.Println(v2, v4, v5)
        v74 := []int{(^(v1))}
        _ = v74
    } else {
        if (((((v3 - 0.0) / (v3 + v3)) + 3.14159) * ((float64((3.14159 / 1.5)) / 1.0) / 1e10)) < 1e10) {
            if true {
                v75 := GenS7{F0: 130, F1: 65535, F2: "hello"}
                v76 := []int{(v1 | (^(v1))), (^(v1)), -581}
                v4 += v4
                _ = v75
                _ = v76
            }
        } else {
            v4 = v4
            var v77 string = ""
            v78 := GenS4{F0: (v1 & v1), F1: v5}
            _ = v77
            _ = v78
        }
        var v79 bool = (true && (2147483647 > 42))
        _ = v79
    }
    v80 := []int{(v1 << ((42) & 63)), 127, (v1 >> ((1) & 63)), v1}
    var v81 int32 = 0
    _ = len(v80)
    v2 = v2
    if len(v80) >= 2 {
        v80 = v80[1:len(v80)]
    }
    v1 -= 736
    if len(v80) > 0 {
        v80[0] = (v1 * (v1 / (((v1 % ((v1) | 1))) | 1)))
    }
    if ((v1 / ((v1) | 1)) < 128) {
        v4 = v4
    }
    v81 -= -2147483648
    var v82 int = (v1 * (v1 / ((v1) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v80:", len(v80), v80)
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v80
    _ = v81
    _ = v82
}
