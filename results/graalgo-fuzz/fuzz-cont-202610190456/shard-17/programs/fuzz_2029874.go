package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 string; F2 int64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint8; F1 string; F2 bool }
type GenS5 struct { F0 int32; F1 uint8; F2 int16 }
type GenS6 struct { F0 bool }
func genFn0(p0 uint16, p1 uint64, p2 uint32) uint64 {
    switch "foo bar" {
    case "":
        if false {
            if (!true) {
                fmt.Println(p2, p0, p0)
                if true {
                    v7 := GenS1{F0: (p2 / ((p2) | 1))}
                    _ = v7
                } else {
                    fmt.Println(p1, p1, p0)
                    if (int32(p2) < -2147483648) {
                        var v8 string = ""
                        fmt.Println(p1)
                        v9 := []int{0}
                        _ = v8
                        _ = v9
                    } else {
                        v10 := GenS1{F0: 821}
                        _ = v10
                    }
                }
            } else {
                p0 += 42
                for fi11 := 0; fi11 < 10; fi11++ {
                    if (!true) {
                        p0 -= 32767
                        v12 := GenS2{F0: p0, F1: "\n\t", F2: -1}
                        _ = v12
                    } else {
                        if false {
                            switch fi11 {
                            case -768:
                                if (!(fi11 != int(p0))) {
                                    switch (fi11 - fi11) {
                                    case 9223372036854775807:
                                        var v13 int64 = 9223372036854775807
                                        _ = v13
                                    case 4294967295:
                                        var v14 float64 = (0.1 / 3.14159)
                                        _ = v14
                                    default:
                                        p0 = p0
                                    }
                                    var v15 uint16 = (p0 ^ 32767)
                                    _ = v15
                                } else {
                                    p0 -= 1
                                    v16 := []int{fi11, (fi11 << ((42) & 63)), -209}
                                    _ = v16
                                }
                            default:
                                p2 = p2
                            }
                            fmt.Println(p0, p0)
                        }
                        v17 := GenS3{F0: (fi11 ^ (-(fi11)))}
                        var v18 uint32 = p2
                        _ = v17
                        _ = v18
                    }
                    _ = fi11
                }
            }
            v19 := []int{-1, -2147483648, 9223372036854775807}
            _ = v19
        }
    case "\n\t":
        p1 -= (^(p1))
    case "foo bar":
        if true {
            for fi20 := 0; fi20 < 10; fi20++ {
                var v21 int = fi20
                _ = fi20
                _ = v21
            }
            v22 := []int{-563, 42}
            fmt.Println(p1, p1)
            _ = v22
        } else {
            p1 = (p1 / ((4294967295) | 1))
        }
    default:
        var v23 int8 = 4
        _ = v23
    }
    var v24 int64 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    return 0
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]bool)
    var v3 float32 = 0.1
    var v4 uint32 = 636
    var v5 uint8 = 51
    v6 := []uint32{v4}
    v5 = 0
    v6 = append(v6, v4)
    v25 := GenS3{F0: 42}
    for fi26 := 0; fi26 < 20; fi26++ {
        v27 := GenS1{F0: (v4 / (((v4 ^ 2147483648)) | 1))}
        _ = fi26
        _ = v27
    }
    for fi28 := 0; fi28 < 1; fi28++ {
        v5 -= 69
        v4 = (v4 - 2147483647)
        _ = fi28
    }
    switch (("foo bar" + "") + "\n\t") {
    case "hello":
        var v29 uint32 = v4
        _ = v29
    case "":
        for fi30 := 0; fi30 < 7; fi30++ {
            switch "foo bar" {
            case "foo bar":
                v2[(v4 << ((42) & 63))] = (true && ((v4 != 65535) || false))
            case "hello":
                _ = len(v6)
            case "\n\t":
                if (("" + ("" + "a")) > "hello") {
                    if false {
                        v1[(^(v4))] = ""
                        var v31 int32 = 0
                        v32 := GenS1{F0: (v4 + (v4 | v4))}
                        _ = v31
                        _ = v32
                    } else {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        var v33 uint64 = 1
                        if true {
                            var v34 int8 = 42
                            v3 = 0.0
                            _ = v34
                        } else {
                            v4 = (^(v4))
                        }
                        _ = v33
                    }
                    switch fi30 {
                    case -1:
                        if len(v6) > 0 {
                            v6[0] = 484
                        }
                    case 0:
                        for fi35 := 0; fi35 < 19; fi35++ {
                            if ((42 != 1) && true) {
                                var v36 string = (("\n\t" + "\n\t") + ("\n\t" + "foo bar"))
                                switch v36 {
                                case "a":
                                    v37 := GenS2{F0: 119, F1: v36, F2: 0}
                                    _ = v37
                                case "\n\t":
                                    v5 -= 113
                                case "hello":
                                    v3 = v3
                                }
                                _ = v36
                            } else {
                                v3 = v3
                            }
                            if false {
                                var v38 bool = true
                                v5 += v5
                                var v39 uint64 = genFn0(460, 0, v4)
                                _ = v38
                                _ = v39
                            }
                            var v40 bool = false
                            _ = fi35
                            _ = v40
                        }
                    case 32767:
                        v4 -= 2147483648
                    default:
                        v41 := GenS3{F0: (fi30 ^ fi30)}
                        _ = v41
                    }
                    var v42 uint16 = 339
                    _ = v42
                }
            }
            _ = fi30
        }
    case "\n\t":
        if true {
            _ = len(v2)
        }
    default:
        var v43 uint16 = 255
        _ = v43
    }
    if false {
        for fi44 := 0; fi44 < 7; fi44++ {
            v5 -= 2
            var v45 int16 = 1000
            _ = fi44
            _ = v45
        }
        delete(v1, 401)
        delete(v1, v4)
    } else {
        var v46 float32 = v3
        _ = v46
    }
    for fi47 := 0; fi47 < 9; fi47++ {
        v4 += (v4 % ((v4) | 1))
        v3 = v3
        for fi48 := 0; fi48 < 15; fi48++ {
            switch "a" {
            case "a":
                v49 := GenS2{F0: 255, F1: "foo bar", F2: 0}
                _ = v49
            case "":
                _ = len(v6)
            case "foo bar":
                if false {
                    v5 = (v5 & (v5 << ((255) & 63)))
                }
            default:
                var v50 string = ""
                _ = v50
            }
            var v51 int = 32767
            _ = fi48
            _ = v51
        }
        _ = fi47
    }
    v5 = 128
    v3 = v3
    var v52 int = int(v4)
    if (!(163 != v5)) {
        v4 = (v4 - (v4 | v4))
        v53 := GenS4{F0: 0, F1: "\n\t", F2: ((!(!(1 != 65535))) && (18446744073709551615 >= 964))}
        _ = v53
    } else {
        if (("" + "\n\t") <= ("" + ("hello" + ""))) {
            if ((true && false) && true) {
                var v54 int16 = -1
                for fi55 := 0; fi55 < 20; fi55++ {
                    _ = len(v6)
                    var v56 uint16 = 175
                    var v57 float32 = (0.0 * -2.25)
                    _ = fi55
                    _ = v56
                    _ = v57
                }
                v58 := GenS5{F0: -266, F1: v5, F2: 1000}
                _ = v54
                _ = v58
            } else {
                _ = len(v2)
                if (2 == 0) {
                    var v59 float64 = 0.1
                    var v60 int64 = 485
                    _ = v59
                    _ = v60
                } else {
                    if true {
                        var v61 uint8 = v5
                        var v62 bool = ((1e10 / 1.0) == -2.25)
                        _ = v61
                        _ = v62
                    }
                    v5 -= v5
                }
            }
            v4 -= v4
            for fi63 := 0; fi63 < 20; fi63++ {
                var v64 int = fi63
                _ = fi63
                _ = v64
            }
        } else {
            switch "hello" {
            case "hello":
                var v65 int64 = 0
                _ = v65
            case "a":
                v6 = append(v6, (v4 & (v4 + (v4 % ((v4) | 1)))))
            default:
                v66 := GenS1{F0: (v4 / (((v4 % (((v4 | (v4 - (v4 << ((938) & 63))))) | 1))) | 1))}
                _ = v66
            }
        }
        fmt.Println(v5, v4)
        var v67 int16 = 357
        _ = v67
    }
    v3 = (v3 + (0.1 / (v3 - v3)))
    fmt.Println(v3, v52)
    for fi68 := 0; fi68 < 14; fi68++ {
        if true {
            var v69 uint32 = (v4 | (v4 + 2))
            _ = v69
        } else {
            v70 := GenS6{F0: true}
            v4 = 2147483647
            _ = v70
        }
        _ = fi68
    }
    v4 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v6
}
