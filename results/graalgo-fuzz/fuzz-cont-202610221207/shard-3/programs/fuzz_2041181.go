package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 string; F1 int8 }
type GenS5 struct { F0 uint64; F1 int64 }
func genFn0(p0 int, p1 uint16, p2 uint) uint {
    v6 := []int{p0, p0, (p0 | -642)}
    p1 -= (p1 << ((uint(p1)) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p2
}
func genFn1() uint32 {
    var v7 uint = genFn0(4294967295, 42, 65535)
    for fi8 := 0; fi8 < 5; fi8++ {
        for fi9 := 0; fi9 < 3; fi9++ {
            switch fi9 {
            case 104:
                if (true && true) {
                    v7 += 255
                    var v10 float32 = (3.14159 / 1.0)
                    var v11 float32 = (0.1 * v10)
                    _ = v10
                    _ = v11
                } else {
                    v7 = v7
                    v7 += genFn0(fi8, 42, v7)
                    var v12 uint16 = uint16(v7)
                    _ = v12
                }
            case 992:
                fmt.Println(fi9, fi9, fi8)
            case 128:
                fmt.Println(fi8, v7)
            }
            v13 := []int{(fi9 << ((v7) & 63)), 42, fi9}
            _ = fi9
            _ = v13
        }
        v7 = genFn0(fi8, 13, v7)
        var v14 uint8 = 255
        _ = fi8
        _ = v14
    }
    _ = v7
    return 0
}

func main() {
    var v1 int64 = 1
    var v2 int32 = 151
    var v3 int16 = -32768
    v4 := make(map[int]string)
    var v5 string = "foo bar"
    var v15 int16 = 32767
    var v16 int32 = (v2 | (v2 - v2))
    for fi17 := 0; fi17 < 5; fi17++ {
        v5 = v5
        v15 -= v15
        _ = fi17
    }
    var v18 int = -522
    v19 := GenS1{F0: 762}
    if true {
        delete(v4, 376)
    } else {
        v20 := []int{v18}
        _ = v20
    }
    for fi21 := 0; fi21 < 18; fi21++ {
        if false {
            v16 += v2
            switch (v18 & (fi21 << ((0) & 63))) {
            case 575:
                var v22 uint = 65535
                _ = v22
            case 138:
                if false {
                    if false {
                        var v23 uint32 = genFn1()
                        _ = v23
                    } else {
                        for fi24 := 0; fi24 < 5; fi24++ {
                            v25 := []int{32767, (fi24 % ((fi24) | 1))}
                            _ = fi24
                            _ = v25
                        }
                    }
                    for fi26 := 0; fi26 < 10; fi26++ {
                        var v27 int64 = -9223372036854775808
                        switch v18 {
                        case 2147483647:
                            if false {
                                var v28 uint16 = 725
                                _ = len(v4)
                                fmt.Println(v3, v2)
                                _ = v28
                            }
                        case 32768:
                            if true {
                                fmt.Println(v16, v2)
                            } else {
                                var v29 int16 = v3
                                _ = v29
                            }
                        case 462:
                            v16 += (v2 + v16)
                        }
                        _ = fi26
                        _ = v27
                    }
                    v30 := GenS2{F0: 220}
                    _ = v30
                } else {
                    if true {
                        var v31 float64 = 0.0
                        _ = v31
                    } else {
                        _ = len(v4)
                        var v32 float64 = 1e10
                        v33 := []int{4294967295, v18, fi21, fi21}
                        _ = v32
                        _ = v33
                    }
                    v4[(fi21 - fi21)] = v5
                    v16 = 792
                }
            case 0:
                v34 := GenS3{F0: 48, F1: 0}
                _ = v34
            }
        } else {
            for fi35 := 0; fi35 < 19; fi35++ {
                var v36 float64 = float64((1.5 / 1e10))
                _ = fi35
                _ = v36
            }
            var v37 uint32 = 42
            _ = v37
        }
        _ = fi21
    }
    switch v18 {
    case 903:
        fmt.Println(v1, v16, v3)
    default:
        var v38 string = (("foo bar" + ("\n\t" + v5)) + ("" + ""))
        _ = v38
    }
    if true {
        switch ("hello" + v5) {
        case "\n\t":
            v39 := GenS2{F0: 65535}
            _ = v39
        case "a":
            var v40 uint32 = 308
            _ = v40
        default:
            for fi41 := 0; fi41 < 15; fi41++ {
                if (true && ((18446744073709551615 < 42) && false)) {
                    fmt.Println(fi41, v1, v16)
                    v42 := GenS4{F0: ("\n\t" + ("a" + (("foo bar" + v5) + (v5 + v5)))), F1: 100}
                    v18 = (v18 >> ((genFn0(32767, 2, 628)) & 63))
                    _ = v42
                } else {
                    var v43 uint32 = 397
                    v1 += v1
                    v44 := GenS5{F0: 546, F1: v1}
                    _ = v43
                    _ = v44
                }
                var v45 int16 = (v3 % (((v3 % ((1) | 1))) | 1))
                v45 += (v45 >> ((4294967295) & 63))
                _ = fi41
                _ = v45
            }
        }
        var v46 int32 = v2
        v4[546] = ((v5 + (("foo bar" + "\n\t") + ((v5 + "hello") + v5))) + v5)
        _ = v46
    }
    switch "foo bar" {
    case "":
        v47 := []int{(v18 * 9223372036854775807)}
        _ = v47
    case "hello":
        var v48 int = v18
        _ = v48
    default:
        v49 := v4[520]
        _ = v49
    }
    for fi50 := 0; fi50 < 8; fi50++ {
        if (-282 <= (^(v15))) {
            var v51 int16 = (v15 >> ((237) & 63))
            if false {
                var v52 int32 = v2
                _ = v52
            } else {
                v1 = v1
            }
            _ = v51
        } else {
            v53 := GenS3{F0: 1, F1: v1}
            switch v5 {
            case "a":
                switch v5 {
                case "hello":
                    if true {
                        v15 -= (v3 * (v3 & -1))
                        var v54 int32 = (v2 % (((v2 / ((-1) | 1))) | 1))
                        _ = v54
                    }
                case "":
                    v18 = v18
                }
            case "":
                v18 = (fi50 % (((v18 | int(v1))) | 1))
            }
            _ = v53
        }
        var v55 int64 = 601
        for fi56 := 0; fi56 < 8; fi56++ {
            if false {
                fmt.Println(v5, fi56)
            } else {
                var v57 int = (fi56 & v18)
                _ = v57
            }
            fmt.Println(v55)
            _ = fi56
        }
        _ = fi50
        _ = v55
    }
    if ((v18 % ((v18) | 1)) <= v18) {
        var v58 int64 = (v1 & (v1 / ((v1) | 1)))
        var v59 int32 = v2
        for fi60 := 0; fi60 < 9; fi60++ {
            if ((!((!false) || (("\n\t" != v5) && (!(false || true))))) && (false || false)) {
                var v61 bool = true
                if (v3 <= v15) {
                    v3 -= (v3 / ((v3) | 1))
                    v3 = 0
                } else {
                    var v62 float32 = 0.1
                    switch "\n\t" {
                    case "\n\t":
                        switch v5 {
                        case "hello":
                            fmt.Println(v58, v16, v61)
                        }
                    }
                    _ = v62
                }
                var v63 uint64 = 1
                _ = v61
                _ = v63
            } else {
                var v64 uint64 = 42
                _ = v64
            }
            fmt.Println(fi60)
            v2 -= v16
            _ = fi60
        }
        _ = v58
        _ = v59
    }
    var v65 string = v5
    switch "\n\t" {
    case "a":
        var v66 float64 = float64(float32(-2.25))
        _ = v66
    default:
        switch 364 {
        case 4294967295:
            fmt.Println(v3, v2)
        case 32768:
            fmt.Println(v5, v16, v2)
        case 591:
            for fi67 := 0; fi67 < 15; fi67++ {
                v68 := []int{(fi67 ^ (fi67 + 2147483647)), (fi67 % ((2147483647) | 1)), 2147483647, int(v2)}
                if true {
                    v18 += fi67
                    var v69 string = ((v5 + "") + "")
                    _ = v69
                } else {
                    var v70 int64 = (v1 / ((v1) | 1))
                    _ = v70
                }
                _ = fi67
                _ = v68
            }
        default:
            fmt.Println(v5, v65)
        }
    }
    switch ("hello" + v65) {
    case "hello":
        var v71 bool = ((false && false) || (42 != 858))
        _ = v71
    default:
        for fi72 := 0; fi72 < 13; fi72++ {
            if false {
                v2 -= v2
                v73 := GenS3{F0: 716, F1: (v1 / ((v1) | 1))}
                if (true && false) {
                    v2 += (v2 - v16)
                }
                _ = v73
            }
            switch (^(v18)) {
            case -567:
                var v74 uint = 42
                _ = v74
            case -558:
                v16 = (-(v16))
            case 2147483647:
                fmt.Println(v1, v65, v18)
            default:
                _ = len(v4)
            }
            var v75 int8 = 100
            _ = fi72
            _ = v75
        }
    }
    if false {
        v3 = -844
        delete(v4, 2147483647)
    } else {
        var v76 string = v5
        v3 = 42
        fmt.Println(v1, v2)
        _ = v76
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v65
}
