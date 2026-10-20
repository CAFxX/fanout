package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint16 }
type GenS2 struct { F0 string; F1 uint8; F2 int }
type GenS3 struct { F0 uint16; F1 uint; F2 bool }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint64; F1 string; F2 bool }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint64; F1 int16; F2 int32 }
type GenS9 struct { F0 uint8; F1 string }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint8 = 127
    var v3 int64 = 1
    v4 := make(map[int]string)
    var v5 uint16 = 2
    v3 -= v3
    for fi6 := 0; fi6 < 5; fi6++ {
        for fi7 := 0; fi7 < 5; fi7++ {
            v3 = -9223372036854775808
            switch "hello" {
            case "a":
                var v8 bool = (1.5 != ((0.1 * -2.25) + 3.14159))
                var v9 float64 = (1e10 / ((1e10 / (0.1 * -2.25)) * (3.14159 + (3.14159 + (-2.25 / -2.25)))))
                _ = v8
                _ = v9
            case "hello":
                var v10 string = ""
                _ = v10
            case "foo bar":
                v11 := []int{fi6, 9223372036854775807}
                _ = v11
            default:
                v3 = -101
            }
            v3 = (v3 - 42)
            _ = fi7
        }
        v12 := GenS1{F0: 78, F1: 10, F2: v5}
        _ = fi6
        _ = v12
    }
    for fi13 := 0; fi13 < 11; fi13++ {
        var v14 bool = (-710 != 273)
        var v15 float64 = (1.5 * 1.5)
        for fi16 := 0; fi16 < 13; fi16++ {
            var v17 int = fi13
            fmt.Println(fi16)
            v18 := GenS1{F0: 805, F1: 561, F2: (^(v5))}
            fmt.Println(v2, v14, fi13)
            if false {
                if (1e10 <= v15) {
                    v2 -= 42
                } else {
                    switch "a" {
                    case "":
                        var v19 uint32 = 42
                        _ = v19
                    case "a":
                        if (!v14) {
                            v20 := GenS1{F0: 42, F1: 209, F2: 65535}
                            var v21 uint8 = (v2 << ((647) & 63))
                            switch fi16 {
                            case 2147483647:
                                v22 := GenS2{F0: ("\n\t" + ""), F1: v2, F2: (fi16 + 616)}
                                _ = v22
                            }
                            _ = v20
                            _ = v21
                        } else {
                            if (v2 != v2) {
                                v14 = v14
                            } else {
                                v23 := GenS2{F0: (("a" + ("" + ("hello" + ("hello" + "a")))) + "a"), F1: 87, F2: fi13}
                                _ = v23
                            }
                        }
                    case "foo bar":
                        var v24 int8 = -68
                        _ = v24
                    }
                }
                fmt.Println(v3, v5, v2)
            }
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    v25 := GenS2{F0: "hello", F1: v2, F2: -301}
    v5 = v5
    v5 -= (v5 & v5)
    var v26 uint8 = 42
    v5 = (v5 / ((648) | 1))
    if (-9223372036854775808 > (^(v3))) {
        v27 := GenS3{F0: 65535, F1: 42, F2: true}
        switch -1 {
        case 128:
            switch -1 {
            case 2147483647:
                if false {
                    fmt.Println(v26, v2)
                } else {
                    fmt.Println(v26)
                    switch "a" {
                    case "a":
                        var v28 int8 = 29
                        _ = v28
                    default:
                        fmt.Println(v26, v26)
                    }
                }
            case 128:
                switch "hello" {
                case "foo bar":
                    if (("\n\t" + "") == ("\n\t" + "a")) {
                        v2 += v26
                    }
                }
            case 0:
                if false {
                    if false {
                        v4[-510] = ("hello" + "\n\t")
                        var v29 uint64 = 0
                        _ = v29
                    } else {
                        v2 = v26
                    }
                    if false {
                        v26 -= 207
                        v30 := GenS4{F0: 0}
                        _ = v30
                    } else {
                        v31 := GenS5{F0: 4294967295, F1: "\n\t", F2: true}
                        v5 = (v5 | (v5 - v5))
                        _ = v31
                    }
                }
            default:
                v32 := GenS3{F0: (v5 << ((18446744073709551615) & 63)), F1: 42, F2: false}
                _ = v32
            }
        case 32768:
            switch -767 {
            case 9223372036854775807:
                for fi33 := 0; fi33 < 2; fi33++ {
                    if false {
                        v34 := v4[fi33]
                        v34 = v34
                        for fi35 := 0; fi35 < 13; fi35++ {
                            var v36 string = v34
                            fmt.Println(fi33, v36)
                            _ = fi35
                            _ = v36
                        }
                        _ = v34
                    } else {
                        v1[(-(fi33))] = 65535
                        delete(v1, 42)
                        v4[(fi33 + (fi33 - (fi33 - 2147483647)))] = ""
                    }
                    if true {
                        v37 := GenS6{F0: 2147483648}
                        _ = v37
                    } else {
                        for fi38 := 0; fi38 < 14; fi38++ {
                            var v39 int16 = 1000
                            _ = fi38
                            _ = v39
                        }
                        fmt.Println(v5, v5)
                        v40 := []int{(fi33 + fi33)}
                        _ = v40
                    }
                    if (true && false) {
                        v2 = (v2 % ((v26) | 1))
                    } else {
                        if (!true) {
                            v2 = v2
                        } else {
                            v3 += v3
                            if (!(1000 >= 42)) {
                                delete(v1, fi33)
                                for fi41 := 0; fi41 < 2; fi41++ {
                                    v26 = v2
                                    if true {
                                        v2 = 1
                                    }
                                    if (v3 != v3) {
                                        v2 = v26
                                        var v42 string = ""
                                        var v43 int64 = -1
                                        v43 += (^(v3))
                                        v44 := v1[fi41]
                                        _ = v42
                                        _ = v43
                                        _ = v44
                                    } else {
                                        var v45 string = "foo bar"
                                        v2 -= v2
                                        v46 := GenS2{F0: v45, F1: 255, F2: fi41}
                                        _ = v45
                                        _ = v46
                                    }
                                    _ = fi41
                                }
                            } else {
                                v26 += 0
                                for fi47 := 0; fi47 < 1; fi47++ {
                                    var v48 uint64 = 42
                                    _ = fi47
                                    _ = v48
                                }
                            }
                            var v49 uint32 = 2147483648
                            _ = v49
                        }
                        fmt.Println(v5, fi33)
                        var v50 uint8 = (v26 + (v2 * v26))
                        _ = v50
                    }
                    _ = fi33
                }
            default:
                for fi51 := 0; fi51 < 13; fi51++ {
                    switch ("foo bar" + "\n\t") {
                    case "foo bar":
                        var v52 bool = ((float64(0.1) / 3.14159) <= 0.0)
                        _ = v52
                    }
                    switch "hello" {
                    case "\n\t":
                        for fi53 := 0; fi53 < 13; fi53++ {
                            var v54 bool = (!false)
                            v55 := GenS7{F0: (v3 >> ((42) & 63))}
                            switch "foo bar" {
                            case "foo bar":
                                fmt.Println(v3, v54)
                            default:
                                var v56 string = "a"
                                _ = v56
                            }
                            _ = fi53
                            _ = v54
                            _ = v55
                        }
                    case "a":
                        v57 := GenS5{F0: 42, F1: "foo bar", F2: true}
                        _ = v57
                    case "":
                        v58 := GenS8{F0: 1, F1: 0, F2: 335}
                        _ = v58
                    default:
                        if false {
                            v2 -= v26
                            var v59 int16 = 1000
                            _ = v59
                        } else {
                            if false {
                                for fi60 := 0; fi60 < 12; fi60++ {
                                    if (!true) {
                                        var v61 bool = (!((float32(1.5) * float32(1.5)) == float32(3.14159)))
                                        if true {
                                            var v62 int64 = (v3 << ((4294967295) & 63))
                                            var v63 float64 = 0.0
                                            _ = v62
                                            _ = v63
                                        } else {
                                            v64 := v4[fi60]
                                            _ = v64
                                        }
                                        v61 = (!(false && (fi60 <= fi60)))
                                        _ = v61
                                    }
                                    _ = fi60
                                }
                                var v65 float64 = 1.5
                                v3 += (v3 ^ (^(v3)))
                                _ = v65
                            } else {
                                var v66 uint32 = 42
                                v67 := []int{2147483647, (fi51 / (((fi51 | fi51)) | 1)), (fi51 * 2147483647), fi51}
                                v5 -= v5
                                _ = v66
                                _ = v67
                            }
                            var v68 uint32 = uint32(v5)
                            _ = v68
                        }
                    }
                    if false {
                        for fi69 := 0; fi69 < 2; fi69++ {
                            var v70 uint16 = v5
                            _ = fi69
                            _ = v70
                        }
                        var v71 uint64 = 42
                        _ = v71
                    } else {
                        for fi72 := 0; fi72 < 16; fi72++ {
                            fmt.Println(v3, v3, v5)
                            switch (fi51 % (((fi51 - -759)) | 1)) {
                            case 574:
                                fmt.Println(v3, v5)
                            default:
                                v2 -= (v2 >> ((871) & 63))
                            }
                            if ("\n\t" > "foo bar") {
                                v73 := GenS4{F0: -1}
                                var v74 float64 = float64(((-2.25 * (float32(1e10) * (((0.0 + 0.0) + (((-2.25 - (0.1 + (-2.25 - 1e10))) * 1e10) - 1.5)) + ((0.1 + (float32(-2.25) / ((0.1 / 3.14159) + 0.1))) + 3.14159)))) - float32(float64((1e10 / -2.25)))))
                                fmt.Println(fi51)
                                _ = v73
                                _ = v74
                            } else {
                                var v75 int16 = int16(v3)
                                var v76 uint16 = 42
                                _ = v75
                                _ = v76
                            }
                            switch "" {
                            case "a":
                                var v77 float32 = 3.14159
                                _ = v77
                            default:
                                v78 := GenS7{F0: (v3 | (^(v3)))}
                                _ = v78
                            }
                            _ = fi72
                        }
                        for fi79 := 0; fi79 < 4; fi79++ {
                            v5 += (v5 | (v5 - (v5 - (v5 * (v5 * (v5 ^ 42))))))
                            _ = fi79
                        }
                    }
                    _ = fi51
                }
            }
        case 4294967295:
            v80 := GenS6{F0: 1}
            _ = v80
        default:
            switch -903 {
            case 1:
                fmt.Println(v2, v2, v26)
            case 32767:
                v5 = 42
            case 4294967295:
                for fi81 := 0; fi81 < 19; fi81++ {
                    v82 := v4[int(v26)]
                    v83 := GenS9{F0: (v26 * (v26 | 128)), F1: v82}
                    _ = fi81
                    _ = v82
                    _ = v83
                }
            }
        }
        v2 += v2
        _ = v27
    }
    v1[0] = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
