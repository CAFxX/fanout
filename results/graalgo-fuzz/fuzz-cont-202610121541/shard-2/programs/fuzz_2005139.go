package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint16; F1 bool }
type GenS6 struct { F0 bool; F1 bool }
type GenS7 struct { F0 int8; F1 int8 }
func genFn0(p0 string) int8 {
    v7 := []int{1, 95, -1, -2147483648}
    fmt.Println(p0)
    p0 = p0
    p0 = ""
    fmt.Println(p0)
    _ = p0
    _ = v7
    return -6
}
func genFn1(p0 int) uint32 {
    switch p0 {
    case 1:
        for fi8 := 0; fi8 < 2; fi8++ {
            if true {
                p0 = (fi8 | (p0 ^ (p0 % (((fi8 - 42)) | 1))))
            }
            p0 -= (p0 + (fi8 * (^(p0))))
            _ = fi8
        }
    case 32768:
        p0 = 1
    case 4294967295:
        var v9 uint64 = 2
        _ = v9
    }
    var v10 int64 = 104
    _ = p0
    _ = v10
    return 42
}

func main() {
    v1 := []uint32{1}
    var v2 float32 = 1.5
    var v3 int = 427
    var v4 int16 = 1000
    var v5 uint8 = 55
    v6 := make(map[uint32]int)
    for fi11 := 0; fi11 < 17; fi11++ {
        switch 9223372036854775807 {
        case -9223372036854775808:
            v12 := GenS1{F0: 42}
            _ = v12
        case -1:
            var v13 bool = (v4 == (v4 | -1))
            _ = v13
        case 9223372036854775807:
            v5 -= 0
        }
        v2 = 0.1
        if true {
            v2 -= -2.25
            v14 := GenS2{F0: 75}
            _ = v14
        } else {
            v1 = append(v1, 65535)
        }
        _ = fi11
    }
    if false {
        v2 -= (v2 * v2)
        var v15 uint = 415
        v16 := GenS1{F0: 1}
        _ = v15
        _ = v16
    }
    var v17 bool = (true && true)
    var v18 int8 = 1
    switch v3 {
    case 584:
        v19 := GenS3{F0: 49}
        _ = v19
    case 840:
        var v20 uint64 = 42
        _ = v20
    case 32767:
        switch "foo bar" {
        case "hello":
            if v17 {
                v18 += (v18 * v18)
                fmt.Println(v2, v5, v3)
            } else {
                if (533 >= 1) {
                    var v21 int16 = v4
                    _ = v21
                }
            }
        case "":
            if (("" + "hello") >= ((("\n\t" + ("a" + "foo bar")) + "\n\t") + "hello")) {
                switch ("\n\t" + "\n\t") {
                case "foo bar":
                    var v22 uint16 = 32767
                    _ = v22
                case "\n\t":
                    var v23 int32 = 0
                    _ = v23
                }
            }
        }
    }
    fmt.Println(v5, v5, v17)
    var v24 string = "\n\t"
    var v25 uint32 = 42
    for fi26 := 0; fi26 < 6; fi26++ {
        if (42 == 32767) {
            for fi27 := 0; fi27 < 7; fi27++ {
                if len(v1) > 0 {
                    v1[0] = genFn1(v3)
                }
                if (406 <= 9223372036854775807) {
                    var v28 int = (v3 % ((fi26) | 1))
                    if v17 {
                        var v29 uint = 42
                        _ = v29
                    }
                    switch fi26 {
                    case 4294967295:
                        var v30 string = v24
                        var v31 string = (v24 + ("a" + ""))
                        _ = v30
                        _ = v31
                    case 32768:
                        v32 := GenS2{F0: v18}
                        _ = v32
                    case -9223372036854775808:
                        switch v24 {
                        case "":
                            fmt.Println(fi26)
                        case "foo bar":
                            fmt.Println(v2)
                        case "\n\t":
                            if true {
                                var v33 float64 = (0.1 * 0.1)
                                var v34 int8 = (v18 ^ (v18 + (v18 & genFn0(v24))))
                                _ = v33
                                _ = v34
                            }
                        default:
                            fmt.Println(v25)
                        }
                    }
                    _ = v28
                }
                _ = fi27
            }
            for fi35 := 0; fi35 < 15; fi35++ {
                var v36 uint = 271
                if (0.0 < float64(0.0)) {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    fmt.Println(fi26, v36, v18)
                    _ = len(v1)
                } else {
                    v37 := GenS3{F0: 0}
                    var v38 bool = (!v17)
                    if v17 {
                        v3 += int(v4)
                    } else {
                        if v17 {
                            v39 := v6[(v25 << ((779) & 63))]
                            v17 = ((^(v5)) != (v5 / ((153) | 1)))
                            v40 := GenS4{F0: 128}
                            _ = v39
                            _ = v40
                        }
                    }
                    switch int(v4) {
                    case 42:
                        v41 := GenS3{F0: int8(v5)}
                        _ = v41
                    case 0:
                        v42 := v6[(v25 >> ((664) & 63))]
                        _ = v42
                    case 918:
                        v2 -= -2.25
                    }
                    if false {
                        var v43 int64 = -539
                        _ = v43
                    } else {
                        delete(v6, v25)
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v36 = (v36 + (v36 | (v36 + v36)))
                        switch (v3 | 215) {
                        case -852:
                            v44 := GenS4{F0: fi26}
                            _ = v44
                        case 32768:
                            var v45 uint64 = 0
                            _ = v45
                        case 32767:
                            var v46 int64 = -1
                            _ = v46
                        default:
                            if len(v1) > 0 {
                                v1[0] = uint32(v5)
                            }
                        }
                    }
                    _ = v37
                    _ = v38
                }
                if v17 {
                    if false {
                        v17 = false
                        if v17 {
                            var v47 int16 = (v4 << ((4294967295) & 63))
                            _ = v47
                        }
                    } else {
                        var v48 uint64 = 18446744073709551615
                        switch (fi26 + v3) {
                        case -2147483648:
                            v49 := GenS5{F0: 638, F1: true}
                            _ = v49
                        }
                        if (-2.25 == 1.5) {
                            switch v24 {
                            case "\n\t":
                                var v50 bool = (!(float64(v2) <= (((1e10 * 1e10) - 0.0) * (((float64((v2 + v2)) * 0.1) + ((0.1 + ((0.0 / ((1e10 * (0.1 - (((0.1 / 0.1) / -2.25) - (3.14159 - 3.14159)))) - (float64(v2) - 0.0))) / float64((((v2 - float32(1.5)) * v2) / v2)))) - 3.14159)) * (((1.5 + 1.5) + (1e10 + float64(1.5))) + -2.25)))))
                                _ = v50
                            }
                        }
                        _ = v48
                    }
                    v51 := GenS4{F0: (-(fi35))}
                    var v52 uint32 = uint32(v36)
                    _ = v51
                    _ = v52
                } else {
                    var v53 uint32 = v25
                    _ = v53
                }
                var v54 uint = (v36 % (((v36 >> (((v36 & v36)) & 63))) | 1))
                if (v25 > v25) {
                    var v55 string = (v24 + "foo bar")
                    _ = v55
                }
                _ = fi35
                _ = v36
                _ = v54
            }
            for fi56 := 0; fi56 < 6; fi56++ {
                fmt.Println(v5, v5)
                v3 -= -624
                v57 := GenS1{F0: -1}
                _ = fi56
                _ = v57
            }
        }
        v3 -= (v3 >> ((4294967295) & 63))
        _ = fi26
    }
    var v58 int8 = (-(v18))
    if (!v17) {
        switch (((("foo bar" + (v24 + v24)) + ("\n\t" + v24)) + (v24 + (v24 + (v24 + ((v24 + (("hello" + v24) + "\n\t")) + v24))))) + v24) {
        case "hello":
            if (!((0.0 != 1e10) || (!v17))) {
                var v59 int32 = 1
                _ = v59
            } else {
                v4 += (v4 | 703)
                for fi60 := 0; fi60 < 16; fi60++ {
                    var v61 uint32 = genFn1((fi60 | fi60))
                    v4 = (-(v4))
                    var v62 int = (fi60 >> ((0) & 63))
                    _ = fi60
                    _ = v61
                    _ = v62
                }
            }
        case "a":
            var v63 int8 = 100
            _ = v63
        default:
            v64 := GenS3{F0: v58}
            _ = v64
        }
        var v65 uint16 = 1
        v3 -= (v3 / ((171) | 1))
        _ = v65
    } else {
        v25 += genFn1(v3)
        var v66 int32 = 1
        if (!v17) {
            switch v24 {
            case "":
                fmt.Println(v25)
            case "\n\t":
                for fi67 := 0; fi67 < 6; fi67++ {
                    v1 = append(v1, (v25 * (v25 << ((65535) & 63))))
                    switch -764 {
                    case 1:
                        var v68 int = (^(v3))
                        _ = v68
                    default:
                        fmt.Println(v66)
                    }
                    _ = fi67
                }
            }
            for fi69 := 0; fi69 < 20; fi69++ {
                v70 := v6[v25]
                switch 9223372036854775807 {
                case 127:
                    if (65535 == 2) {
                        if false {
                            fmt.Println(v5, v70, v18)
                            v71 := GenS6{F0: v17, F1: ((v4 - v4) > v4)}
                            _ = len(v1)
                            _ = v71
                        }
                    }
                case 42:
                    fmt.Println(v58, v18, v17)
                }
                if v17 {
                    v17 = v17
                    var v72 uint = 0
                    _ = v72
                } else {
                    var v73 int32 = v66
                    if false {
                        for fi74 := 0; fi74 < 15; fi74++ {
                            var v75 uint16 = 42
                            if ((((v2 - v2) - ((v2 + 3.14159) / v2)) * 1e10) > float32(0.0)) {
                                _ = len(v6)
                                var v76 uint = 329
                                _ = v76
                            } else {
                                var v77 int32 = (v73 << ((0) & 63))
                                _ = v77
                            }
                            if v17 {
                                if (3.14159 > -2.25) {
                                    switch (("" + v24) + v24) {
                                    case "a":
                                        v1 = append(v1, uint32(v5))
                                    }
                                    var v78 float32 = v2
                                    _ = v78
                                } else {
                                    switch (fi69 ^ 4294967295) {
                                    case 4294967295:
                                        v2 = float32(0.0)
                                    }
                                }
                                if len(v1) > 0 {
                                    v1[0] = v25
                                }
                                if ((^(v25)) < (^(v25))) {
                                    fmt.Println(v2)
                                    var v79 int32 = -964
                                    _ = v79
                                } else {
                                    v58 += -1
                                    v1 = append(v1, uint32(v5))
                                }
                            }
                            switch 4294967295 {
                            case -1:
                                delete(v6, (v25 % ((463) | 1)))
                            default:
                                v58 += (v58 ^ (v58 | (^(v18))))
                            }
                            v80 := GenS1{F0: v66}
                            _ = len(v6)
                            _ = fi74
                            _ = v75
                            _ = v80
                        }
                    } else {
                        var v81 uint8 = v5
                        _ = v81
                    }
                    var v82 int = (v70 + (fi69 + (fi69 / ((v70) | 1))))
                    _ = v73
                    _ = v82
                }
                _ = fi69
                _ = v70
            }
            var v83 uint64 = 42
            _ = v83
        } else {
            var v84 int = -831
            v85 := GenS5{F0: 632, F1: v17}
            _ = v84
            _ = v85
        }
        _ = v66
    }
    _ = len(v6)
    for fi86 := 0; fi86 < 9; fi86++ {
        if v17 {
            delete(v6, (v25 + genFn1(v3)))
            v4 -= (-(v4))
            fmt.Println(v5)
        }
        v5 = 2
        _ = fi86
    }
    v24 = v24
    if (v17 || ((v24 + ("hello" + ("a" + ""))) < v24)) {
        var v87 string = v24
        v17 = true
        _ = v87
    } else {
        if (32767 != 255) {
            v88 := GenS7{F0: 5, F1: 89}
            fmt.Println(v2)
            _ = v88
        }
        fmt.Println(v17, v2, v58)
        v89 := GenS7{F0: v58, F1: genFn0(v24)}
        _ = v89
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v6
}
