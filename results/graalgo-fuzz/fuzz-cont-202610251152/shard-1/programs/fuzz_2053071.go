package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 int64 }
type GenS2 struct { F0 uint64; F1 uint8; F2 uint64 }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 string; F1 int; F2 int32 }
type GenS5 struct { F0 int8; F1 uint32; F2 uint32 }
type GenS6 struct { F0 bool; F1 bool; F2 int8 }
func genFn0(p0 int32, p1 string) int {
    fmt.Println(p0)
    for fi4 := 0; fi4 < 20; fi4++ {
        fmt.Println(p1, fi4)
        var v5 bool = ((p1 + (p1 + p1)) > (p1 + "a"))
        _ = fi4
        _ = v5
    }
    _ = p0
    _ = p1
    return -2147483648
}

func main() {
    var v1 int8 = -126
    var v2 uint8 = 255
    v3 := make(map[int]int)
    v2 = v2
    if (!(true || (!false))) {
        v6 := GenS1{F0: 1000, F1: 112, F2: 1}
        var v7 int16 = -32768
        switch ("a" + "a") {
        case "":
            v8 := GenS1{F0: (v7 / (((v7 - v7)) | 1)), F1: 754, F2: 1}
            _ = v8
        case "a":
            var v9 int16 = 356
            _ = v9
        }
        _ = v6
        _ = v7
    }
    var v10 uint64 = 42
    v2 -= (^(v2))
    v11 := GenS2{F0: (^(v10)), F1: (^(v2)), F2: 42}
    _ = len(v3)
    if true {
        var v12 int = 42
        v13 := GenS1{F0: 42, F1: v12, F2: -9223372036854775808}
        _ = v12
        _ = v13
    }
    fmt.Println(v10)
    v2 = 0
    if true {
        for fi14 := 0; fi14 < 12; fi14++ {
            var v15 uint = 0
            if true {
                if ((-1 <= 42) || false) {
                    if (false && true) {
                        for fi16 := 0; fi16 < 4; fi16++ {
                            if true {
                                var v17 uint32 = 42
                                fmt.Println(v2)
                                v1 = (^(v1))
                                _ = v17
                            } else {
                                v18 := []int{(fi16 / ((-1) | 1))}
                                if false {
                                    if len(v18) >= 2 {
                                        v18 = v18[1:len(v18)]
                                    }
                                    var v19 int8 = (v1 >> ((705) & 63))
                                    _ = v19
                                }
                                _ = v18
                            }
                            _ = fi16
                        }
                    } else {
                        var v20 uint8 = (v2 | (^(v2)))
                        var v21 uint64 = v10
                        _ = v20
                        _ = v21
                    }
                    v3[(^(fi14))] = (fi14 >> (((^(v15))) & 63))
                } else {
                    fmt.Println(v15, v2)
                    if false {
                        v2 = (v2 + 128)
                    } else {
                        v2 = 42
                        var v22 float32 = 1e10
                        for fi23 := 0; fi23 < 13; fi23++ {
                            v22 = (float32(1e10) + 0.1)
                            v2 = v2
                            v24 := []int{}
                            _ = fi23
                            _ = v24
                        }
                        _ = v22
                    }
                    var v25 bool = true
                    _ = v25
                }
                v10 = v10
                v26 := GenS2{F0: v10, F1: (^(v2)), F2: (^(v10))}
                _ = v26
            } else {
                v10 += (^(v10))
                fmt.Println(v15, fi14, v10)
            }
            if (!(!(fi14 != (fi14 * int(v2))))) {
                if ((v2 != (v2 - v2)) && true) {
                    var v27 uint64 = (v10 & (v10 + v10))
                    var v28 int16 = 42
                    if true {
                        if (625 < 0) {
                            v28 += v28
                        } else {
                            v10 = 18446744073709551615
                            if ((v1 ^ (v1 * 1)) == v1) {
                                v15 = v15
                                v15 += (^(v15))
                            }
                            switch fi14 {
                            case 2147483647:
                                var v29 float32 = 0.1
                                _ = v29
                            default:
                                v10 = v27
                            }
                        }
                        v1 -= 127
                    }
                    _ = v27
                    _ = v28
                }
                v30 := []int{fi14}
                _ = v30
            } else {
                v1 = (^(v1))
                if (0.1 >= (0.0 / 0.1)) {
                    for fi31 := 0; fi31 < 5; fi31++ {
                        v15 = v15
                        switch "foo bar" {
                        case "hello":
                            v2 = v2
                        case "foo bar":
                            var v32 uint = v15
                            _ = v32
                        case "a":
                            var v33 uint = v15
                            _ = v33
                        default:
                            var v34 bool = true
                            _ = v34
                        }
                        _ = fi31
                    }
                    if true {
                        switch "foo bar" {
                        case "\n\t":
                            var v35 uint32 = 65535
                            _ = v35
                        case "":
                            switch ("" + "") {
                            case "hello":
                                v36 := []int{fi14, genFn0(425, ("a" + ""))}
                                _ = v36
                            case "a":
                                var v37 int64 = 9223372036854775807
                                _ = v37
                            case "":
                                var v38 int8 = v1
                                _ = v38
                            default:
                                var v39 bool = false
                                _ = v39
                            }
                        }
                        if true {
                            v40 := GenS2{F0: 4294967295, F1: (v2 - (v2 & v2)), F2: (^(v10))}
                            v2 = 61
                            _ = v40
                        } else {
                            switch fi14 {
                            case 982:
                                if (!(681 == 657)) {
                                    v1 += v1
                                    delete(v3, fi14)
                                    var v41 uint8 = 166
                                    _ = v41
                                }
                            case 128:
                                v42 := GenS2{F0: (v10 << ((v15) & 63)), F1: v2, F2: v10}
                                _ = v42
                            default:
                                v43 := []int{}
                                _ = v43
                            }
                            fmt.Println(v1)
                            v44 := GenS3{F0: -505, F1: (v15 + v15)}
                            _ = v44
                        }
                    } else {
                        v2 += (^(v2))
                        if true {
                            switch (^(fi14)) {
                            case 32768:
                                v1 -= (v1 << ((v15) & 63))
                            case -840:
                                for fi45 := 0; fi45 < 12; fi45++ {
                                    v15 += 18446744073709551615
                                    _ = fi45
                                }
                            case -994:
                                var v46 uint32 = 2147483648
                                _ = v46
                            default:
                                switch (("" + "foo bar") + "hello") {
                                case "foo bar":
                                    if false {
                                        v1 = (-(v1))
                                        fmt.Println(fi14, fi14)
                                    }
                                case "":
                                    v47 := GenS4{F0: "\n\t", F1: fi14, F2: 534}
                                    _ = v47
                                default:
                                    v48 := GenS5{F0: (v1 >> ((18446744073709551615) & 63)), F1: 2, F2: 2}
                                    _ = v48
                                }
                            }
                        }
                    }
                }
                if false {
                    v15 = v15
                } else {
                    v15 = v15
                    switch (^(fi14)) {
                    case -837:
                        for fi49 := 0; fi49 < 17; fi49++ {
                            var v50 uint8 = 244
                            if ("" != "") {
                                if (482 == 989) {
                                    var v51 bool = true
                                    v52 := GenS5{F0: 42, F1: 1, F2: 2}
                                    var v53 int8 = (v1 / (((v1 ^ v1)) | 1))
                                    _ = v51
                                    _ = v52
                                    _ = v53
                                }
                                fmt.Println(v10, v10)
                            } else {
                                var v54 int32 = 1
                                fmt.Println(v50)
                                _ = v54
                            }
                            _ = fi49
                            _ = v50
                        }
                    case 1:
                        var v55 string = ("hello" + ("a" + "\n\t"))
                        _ = v55
                    case -579:
                        var v56 uint64 = (^(v10))
                        _ = v56
                    default:
                        switch "foo bar" {
                        case "\n\t":
                            var v57 bool = (!false)
                            _ = v57
                        case "foo bar":
                            v58 := GenS6{F0: (42 >= 0), F1: (32768 == fi14), F2: v1}
                            _ = v58
                        }
                    }
                }
            }
            _ = fi14
            _ = v15
        }
        v59 := GenS5{F0: (v1 << ((18446744073709551615) & 63)), F1: 988, F2: 65535}
        _ = v59
    }
    for fi60 := 0; fi60 < 15; fi60++ {
        var v61 int64 = -225
        if (false || true) {
            var v62 uint16 = uint16(v2)
            fmt.Println(v2, v2, v62)
            _ = v62
        } else {
            v3[int(v10)] = genFn0(0, "a")
            if ((^(v10)) < 547) {
                var v63 bool = false
                var v64 uint8 = v2
                v63 = ((!((2147483647 == 2147483647) || v63)) && true)
                _ = v63
                _ = v64
            } else {
                v2 -= (v2 - (^(v2)))
                delete(v3, genFn0(42, "hello"))
                for fi65 := 0; fi65 < 17; fi65++ {
                    var v66 int16 = 1000
                    var v67 int16 = v66
                    _ = fi65
                    _ = v66
                    _ = v67
                }
            }
        }
        _ = fi60
        _ = v61
    }
    v68 := v3[32767]
    v1 = v1
    if (v68 != (v68 ^ (-(v68)))) {
        v69 := GenS3{F0: 42, F1: 42}
        _ = v69
    } else {
        var v70 string = "foo bar"
        if (((-2.25 - 1.5) * 1.5) >= 0.1) {
            for fi71 := 0; fi71 < 13; fi71++ {
                fmt.Println(v68, v1, v10)
                _ = fi71
            }
        }
        v1 += 127
        _ = v70
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v68:", v68)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v68
}
