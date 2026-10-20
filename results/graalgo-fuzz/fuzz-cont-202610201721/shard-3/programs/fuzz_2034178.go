package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 int32; F1 uint32; F2 bool }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 int16; F1 string; F2 uint }
type GenS5 struct { F0 int16; F1 bool }
type GenS6 struct { F0 int16; F1 uint16 }
func genFn0() string {
    for fi4 := 0; fi4 < 7; fi4++ {
        v5 := []int{fi4}
        _ = fi4
        _ = v5
    }
    if true {
        v6 := GenS1{F0: (!false), F1: 100}
        v7 := GenS1{F0: false, F1: 1}
        var v8 uint32 = 222
        _ = v6
        _ = v7
        _ = v8
    }
    switch "" {
    case "hello":
        var v9 float32 = ((1e10 / float32(1.5)) * 0.0)
        _ = v9
    case "\n\t":
        var v10 uint64 = 0
        _ = v10
    }
    switch -547 {
    case 4294967295:
        v11 := []int{1, 128}
        _ = v11
    case 329:
        v12 := []int{0}
        _ = v12
    default:
        var v13 uint32 = 2147483647
        _ = v13
    }
    return "a"
}
func genFn1(p0 string, p1 bool) int8 {
    v14 := []int{-1, -1}
    var v15 uint16 = 2
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    return 100
}

func main() {
    var v1 uint32 = 0
    v2 := make(map[uint32]bool)
    var v3 bool = true
    var v16 float32 = -2.25
    v17 := GenS1{F0: (65535 < 42), F1: -126}
    if (!true) {
        var v18 int16 = 37
        v18 += (v18 * v18)
        _ = v18
    }
    if true {
        if v3 {
            if (false || false) {
                var v19 uint32 = v1
                _ = v19
            }
        }
        v20 := GenS2{F0: 0, F1: 806, F2: (42 > -2147483648)}
        switch -928 {
        case 1:
            if (1 != 401) {
                if (1 <= 42) {
                    var v21 string = "hello"
                    v16 = v16
                    _ = v21
                } else {
                    if (782 <= 0) {
                        switch ("hello" + "foo bar") {
                        case "\n\t":
                            v2[(v1 << ((1) & 63))] = v3
                        default:
                            v22 := GenS1{F0: (false || (true || (44 < 100))), F1: 42}
                            _ = v22
                        }
                        var v23 int16 = 42
                        _ = v23
                    } else {
                        var v24 uint32 = 176
                        _ = v24
                    }
                    if ((!v3) && v3) {
                        v25 := GenS2{F0: 42, F1: (v1 & 4294967295), F2: true}
                        _ = v25
                    } else {
                        for fi26 := 0; fi26 < 19; fi26++ {
                            fmt.Println(v16, v3, v3)
                            fmt.Println(fi26, v16)
                            var v27 int16 = 0
                            _ = fi26
                            _ = v27
                        }
                        fmt.Println(v1, v1, v16)
                    }
                    if (!(((!(-2147483648 == 1)) && (!false)) || (!v3))) {
                        for fi28 := 0; fi28 < 20; fi28++ {
                            if v3 {
                                fmt.Println(v1)
                                for fi29 := 0; fi29 < 11; fi29++ {
                                    v30 := GenS1{F0: v3, F1: -1}
                                    v31 := GenS3{F0: 1, F1: -1}
                                    _ = fi29
                                    _ = v30
                                    _ = v31
                                }
                            }
                            fmt.Println(fi28, fi28, v1)
                            v16 -= v16
                            _ = fi28
                        }
                        var v32 string = (("" + "\n\t") + "hello")
                        switch 2147483647 {
                        case 514:
                            var v33 uint = 4294967295
                            _ = v33
                        case 42:
                            if false {
                                if (!(1000 < 89)) {
                                    for fi34 := 0; fi34 < 19; fi34++ {
                                        v35 := GenS1{F0: (914 < 0), F1: 0}
                                        v36 := GenS4{F0: 218, F1: (v32 + "a"), F2: 0}
                                        _ = fi34
                                        _ = v35
                                        _ = v36
                                    }
                                    v1 -= (v1 % (((v1 & 1)) | 1))
                                } else {
                                    v37 := []int{}
                                    _ = v37
                                }
                                var v38 int64 = 1
                                _ = v38
                            }
                        case -73:
                            delete(v2, v1)
                        default:
                            v1 = v1
                        }
                        _ = v32
                    } else {
                        v39 := []int{2147483647}
                        _ = v39
                    }
                }
                switch 682 {
                case 32767:
                    if v3 {
                        if v3 {
                            var v40 string = ("\n\t" + ("hello" + ""))
                            _ = v40
                        } else {
                            switch -1 {
                            case 0:
                                v2[v1] = true
                            case 32767:
                                switch 711 {
                                case -2147483648:
                                    v16 += 1.5
                                default:
                                    v16 += v16
                                }
                            case 4294967295:
                                var v41 float32 = 1.5
                                _ = v41
                            }
                            var v42 uint32 = v1
                            var v43 uint8 = 0
                            _ = v42
                            _ = v43
                        }
                    }
                case 32768:
                    v16 -= v16
                }
            }
        }
        _ = v20
    } else {
        var v44 float64 = (1e10 * ((((-2.25 + (-2.25 / 1e10)) / 1.0) - 1e10) / 0.1))
        _ = v44
    }
    if v3 {
        for fi45 := 0; fi45 < 15; fi45++ {
            v46 := GenS3{F0: 24, F1: 2147483647}
            _ = fi45
            _ = v46
        }
        fmt.Println(v1)
        v47 := []int{9223372036854775807, -815, -737}
        _ = v47
    } else {
        switch 434 {
        case 1:
            switch 32767 {
            case 127:
                delete(v2, v1)
            case -45:
                var v48 uint16 = 2
                _ = v48
            default:
                v2[v1] = v3
            }
        }
        v3 = false
    }
    v2[(v1 & (^(v1)))] = ((2 >= 255) && (!(!true)))
    var v49 uint16 = 695
    if (100 < -74) {
        var v50 int8 = 100
        var v51 int64 = -9223372036854775808
        var v52 int = 380
        _ = v50
        _ = v51
        _ = v52
    } else {
        delete(v2, (v1 | 964))
    }
    v49 -= v49
    if ("hello" == ("\n\t" + "a")) {
        if v3 {
            var v53 int64 = 0
            if v3 {
                v54 := GenS5{F0: 438, F1: ((0.0 - 0.1) < (3.14159 + 0.0))}
                for fi55 := 0; fi55 < 13; fi55++ {
                    v56 := GenS5{F0: -1, F1: (!(!(47 != 42)))}
                    _ = fi55
                    _ = v56
                }
                _ = v54
            } else {
                fmt.Println(v16, v16)
                var v57 bool = (!(4294967295 > 765))
                for fi58 := 0; fi58 < 14; fi58++ {
                    switch "" {
                    case "\n\t":
                        var v59 bool = ((v49 % ((2) | 1)) == v49)
                        _ = v59
                    case "":
                        if (false && (!v3)) {
                            v3 = v3
                            v49 += 255
                        }
                    default:
                        for fi60 := 0; fi60 < 12; fi60++ {
                            v61 := GenS6{F0: -1, F1: (v49 * uint16(v1))}
                            var v62 uint16 = (v49 & v49)
                            _ = fi60
                            _ = v61
                            _ = v62
                        }
                    }
                    _ = fi58
                }
                _ = v57
            }
            fmt.Println(v53, v53)
            _ = v53
        }
        fmt.Println(v16, v16, v16)
        if v3 {
            v1 -= 65535
            for fi63 := 0; fi63 < 13; fi63++ {
                v64 := []int{(fi63 << ((1) & 63)), -774}
                if v3 {
                    for fi65 := 0; fi65 < 16; fi65++ {
                        if false {
                            fmt.Println(fi63)
                            fmt.Println(v3, v1, fi65)
                            var v66 int64 = -484
                            _ = v66
                        } else {
                            if v3 {
                                v67 := GenS5{F0: -494, F1: true}
                                _ = v67
                            }
                            var v68 float32 = v16
                            _ = v68
                        }
                        var v69 uint32 = v1
                        var v70 bool = false
                        v16 -= float32(-2.25)
                        _ = fi65
                        _ = v69
                        _ = v70
                    }
                    var v71 float32 = 0.0
                    _ = v71
                }
                _ = fi63
                _ = v64
            }
        } else {
            v2[(v1 % (((v1 * 88)) | 1))] = v3
            v1 = 575
        }
    }
    for fi72 := 0; fi72 < 9; fi72++ {
        var v73 uint16 = 42
        _ = fi72
        _ = v73
    }
    var v74 uint32 = 186
    v1 += (v74 << ((684) & 63))
    v3 = v3
    for fi75 := 0; fi75 < 9; fi75++ {
        if (!(23 > 526)) {
            v74 = (v1 / (((v1 ^ v1)) | 1))
            v16 -= 1.5
        }
        fmt.Println(v3)
        _ = fi75
    }
    v16 += 0.1
    v1 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v49:", v49)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v49
    _ = v74
}
