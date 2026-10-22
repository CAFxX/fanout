package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint8 }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint; F1 int16 }
type GenS5 struct { F0 uint8; F1 uint; F2 uint }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int16; F1 int; F2 int8 }
type GenS8 struct { F0 bool; F1 uint16 }
type GenS9 struct { F0 uint32; F1 uint16; F2 int32 }
func genFn0(p0 uint64, p1 int16, p2 int) int64 {
    for fi5 := 0; fi5 < 19; fi5++ {
        for fi6 := 0; fi6 < 13; fi6++ {
            var v7 int32 = 713
            fmt.Println(v7, p1)
            _ = fi6
            _ = v7
        }
        var v8 float64 = 3.14159
        _ = fi5
        _ = v8
    }
    v9 := []int{int(p1), (p2 - (p2 + p2))}
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return -1
}
func genFn1() int8 {
    for fi10 := 0; fi10 < 11; fi10++ {
        for fi11 := 0; fi11 < 9; fi11++ {
            var v12 int16 = -541
            var v13 float64 = float64(-2.25)
            switch "\n\t" {
            case "\n\t":
                v14 := GenS1{F0: 42, F1: 4294967295, F2: 2}
                _ = v14
            case "a":
                fmt.Println(v12)
            case "":
                fmt.Println(fi11, v12)
                v15 := GenS1{F0: 42, F1: 2, F2: 184}
                switch (fi10 % ((1) | 1)) {
                case -9223372036854775808:
                    fmt.Println(v13, v12)
                case 1:
                    v12 = v12
                case 0:
                    v16 := []int{-2147483648, (fi10 % ((127) | 1))}
                    _ = v16
                default:
                    var v17 uint32 = 1
                    _ = v17
                }
                v12 = v12
                _ = v15
            }
            v18 := GenS2{F0: false, F1: 128}
            _ = fi11
            _ = v12
            _ = v13
            _ = v18
        }
        if (!(true && false)) {
            for fi19 := 0; fi19 < 1; fi19++ {
                v20 := GenS3{F0: 756}
                v21 := GenS4{F0: 255, F1: -564}
                var v22 uint32 = 444
                switch "foo bar" {
                case "hello":
                    fmt.Println(fi19)
                case "\n\t":
                    v23 := GenS5{F0: 2, F1: 895, F2: 355}
                    if true {
                        var v24 float32 = (3.14159 - (1.5 - 3.14159))
                        _ = v24
                    } else {
                        v25 := GenS2{F0: false, F1: 32}
                        v22 = v22
                        _ = v25
                    }
                    _ = v23
                case "foo bar":
                    var v26 int8 = 42
                    if true {
                        if (false && false) {
                            v22 += (v22 * v22)
                        } else {
                            var v27 int = -125
                            v28 := []int{(^(v27))}
                            var v29 int64 = genFn0(806, 848, v27)
                            _ = v27
                            _ = v28
                            _ = v29
                        }
                    }
                    _ = v26
                default:
                    v30 := []int{32767, fi10, (fi19 * (fi19 * fi19))}
                    _ = v30
                }
                var v31 int32 = -1
                _ = fi19
                _ = v20
                _ = v21
                _ = v22
                _ = v31
            }
        }
        _ = fi10
    }
    for fi32 := 0; fi32 < 13; fi32++ {
        v33 := GenS6{F0: 42}
        _ = fi32
        _ = v33
    }
    return -1
}

func main() {
    v1 := make(map[string]int)
    var v2 float32 = (1e10 + 1.5)
    var v3 uint8 = 42
    v4 := make(map[int]bool)
    for fi34 := 0; fi34 < 3; fi34++ {
        for fi35 := 0; fi35 < 11; fi35++ {
            var v36 uint = 1
            _ = fi35
            _ = v36
        }
        _ = fi34
    }
    for fi37 := 0; fi37 < 3; fi37++ {
        v38 := []int{-695, fi37}
        if true {
            var v39 float32 = v2
            for fi40 := 0; fi40 < 10; fi40++ {
                if (!true) {
                    v39 = -2.25
                    var v41 uint = 376
                    _ = len(v1)
                    _ = v41
                } else {
                    if len(v38) >= 2 {
                        v38 = v38[1:len(v38)]
                    }
                }
                var v42 float64 = -2.25
                _ = fi40
                _ = v42
            }
            for fi43 := 0; fi43 < 1; fi43++ {
                v2 = v2
                fmt.Println(fi37, fi37)
                switch 321 {
                case 127:
                    v3 += (v3 % (((^(v3))) | 1))
                    var v44 uint64 = uint64(v3)
                    var v45 uint16 = 32767
                    _ = v44
                    _ = v45
                case 848:
                    v39 = -2.25
                case -2147483648:
                    var v46 uint32 = 1
                    _ = v46
                }
                if false {
                    v1[("" + "a")] = fi37
                    if (250 <= 94) {
                        if true {
                            if false {
                                var v47 int = (fi37 << ((18446744073709551615) & 63))
                                _ = v47
                            }
                            if (true || (true || (!(-1 >= 0)))) {
                                v48 := v4[(fi37 << ((1) & 63))]
                                _ = v48
                            }
                        }
                    } else {
                        var v49 int = 128
                        var v50 uint16 = 0
                        var v51 int8 = 12
                        v52 := GenS7{F0: 281, F1: 128, F2: -128}
                        _ = len(v4)
                        _ = v49
                        _ = v50
                        _ = v51
                        _ = v52
                    }
                    var v53 string = ("foo bar" + "a")
                    var v54 int = (-(fi37))
                    _ = v53
                    _ = v54
                }
                _ = fi43
            }
            _ = v39
        }
        _ = fi37
        _ = v38
    }
    var v55 bool = true
    v55 = (v55 || v55)
    v56 := []int{1, -1, -2147483648}
    var v57 uint = 255
    v58 := v1[("foo bar" + "\n\t")]
    var v59 int = v58
    v57 = v57
    v60 := GenS8{F0: v55, F1: 1}
    var v61 int8 = genFn1()
    switch v58 {
    case 128:
        v62 := GenS1{F0: uint32(v3), F1: 2, F2: (v3 << ((65535) & 63))}
        _ = v62
    default:
        switch (v59 >> ((v57) & 63)) {
        case -2147483648:
            v3 -= v3
        default:
            switch (^(v59)) {
            case 390:
                if v55 {
                    fmt.Println(v57, v57)
                } else {
                    var v63 int64 = genFn0(1, 42, (v58 % ((-9223372036854775808) | 1)))
                    fmt.Println(v3, v58, v58)
                    v59 += 4294967295
                    _ = v63
                }
            case 46:
                v58 = 412
            case 943:
                if ((783 > 1) && true) {
                    v4[v59] = (-1 == 1)
                    var v64 int = v58
                    if v55 {
                        delete(v1, "a")
                        if (!true) {
                            v58 = (v64 | (^(v64)))
                            for fi65 := 0; fi65 < 6; fi65++ {
                                v57 = 42
                                _ = fi65
                            }
                        } else {
                            v66 := GenS9{F0: 2147483647, F1: 2, F2: -2147483648}
                            v64 -= (v64 - v59)
                            _ = v66
                        }
                        fmt.Println(v61)
                    }
                    _ = v64
                }
            default:
                v3 += (^(v3))
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v55:", v55)
    fmt.Println("slicelen v56:", len(v56), v56)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v55
    _ = v56
    _ = v57
    _ = v58
    _ = v59
    _ = v60
    _ = v61
}
