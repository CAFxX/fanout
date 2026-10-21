package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 string; F2 uint16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int; F1 string; F2 int }
type GenS5 struct { F0 uint32; F1 int64 }
type GenS6 struct { F0 uint16; F1 uint8; F2 uint32 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 string; F1 int64 }
func genFn0() uint32 {
    var v6 int32 = -1
    var v7 int = 32768
    var v8 float64 = 3.14159
    var v9 uint = 4294967295
    v10 := []int{v7, v7, (v7 % ((v7) | 1)), (v7 - v7)}
    _ = v10
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 2147483647
}
func genFn1(p0 uint, p1 bool, p2 uint32) bool {
    var v11 float32 = (0.0 / 0.1)
    v12 := []int{0, 282, -9223372036854775808, 928}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return false
}

func main() {
    var v1 string = "a"
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]bool)
    v4 := make(map[int]bool)
    var v5 int16 = 687
    for fi13 := 0; fi13 < 8; fi13++ {
        var v14 float64 = (0.1 + (0.0 - 1.5))
        v15 := GenS1{F0: 77}
        _ = fi13
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 17; fi16++ {
        for fi17 := 0; fi17 < 17; fi17++ {
            v5 = v5
            v18 := []int{fi17, fi16}
            _ = fi17
            _ = v18
        }
        var v19 uint16 = 255
        _ = fi16
        _ = v19
    }
    for fi20 := 0; fi20 < 19; fi20++ {
        switch (fi20 % ((133) | 1)) {
        case -9223372036854775808:
            v21 := GenS1{F0: 128}
            _ = v21
        case 32767:
            fmt.Println(fi20, fi20, v5)
        case 988:
            v22 := GenS1{F0: 212}
            _ = v22
        }
        var v23 int = (-(fi20))
        switch v23 {
        case 805:
            for fi24 := 0; fi24 < 14; fi24++ {
                v25 := GenS2{F0: -1, F1: (v1 + (((v1 + v1) + (v1 + ((v1 + ("hello" + v1)) + v1))) + ((v1 + (((("a" + v1) + "\n\t") + (((("hello" + (v1 + v1)) + v1) + v1) + (v1 + v1))) + v1)) + v1))), F2: 312}
                v5 = 0
                v26 := GenS3{F0: genFn0()}
                _ = fi24
                _ = v25
                _ = v26
            }
        case 128:
            for fi27 := 0; fi27 < 5; fi27++ {
                fmt.Println(v23)
                var v28 int64 = -9223372036854775808
                _ = fi27
                _ = v28
            }
        default:
            for fi29 := 0; fi29 < 10; fi29++ {
                if ((!(((!false) && true) && (!(1 > 1)))) || false) {
                    v30 := GenS4{F0: (fi29 << ((618) & 63)), F1: v1, F2: 30}
                    switch fi29 {
                    case 301:
                        if (628 <= 255) {
                            if true {
                                if (0 == 127) {
                                    v1 = v1
                                    v5 = (v5 % (((v5 - (v5 + (v5 - 0)))) | 1))
                                } else {
                                    v23 -= v23
                                    v31 := GenS3{F0: 65535}
                                    _ = v31
                                }
                            }
                            if (-128 != 25) {
                                v23 += (fi20 * (fi20 << ((622) & 63)))
                                v1 = "hello"
                                var v32 bool = (!(true || true))
                                v33 := GenS5{F0: genFn0(), F1: int64(fi20)}
                                _ = len(v4)
                                _ = v32
                                _ = v33
                            }
                            delete(v4, fi29)
                        }
                    case -805:
                        fmt.Println(v5, v23, fi29)
                    default:
                        v1 = "hello"
                    }
                    fmt.Println(v5)
                    _ = v30
                } else {
                    var v34 int16 = v5
                    _ = v34
                }
                v23 += fi29
                _ = fi29
            }
        }
        _ = fi20
        _ = v23
    }
    var v35 int8 = -120
    switch v1 {
    case "":
        switch 0 {
        case -9223372036854775808:
            fmt.Println(v35, v1)
        case 32767:
            var v36 uint64 = 42
            _ = v36
        case -371:
            v37 := GenS5{F0: genFn0(), F1: 0}
            _ = v37
        default:
            fmt.Println(v1, v1)
        }
    case "a":
        v38 := []int{-929, 2147483647}
        _ = v38
    }
    for fi39 := 0; fi39 < 9; fi39++ {
        v40 := GenS2{F0: 1, F1: v1, F2: 1}
        _ = fi39
        _ = v40
    }
    var v41 int64 = 9223372036854775807
    v42 := v3[4294967295]
    fmt.Println(v5)
    var v43 uint32 = 592
    if v42 {
        var v44 int8 = v35
        delete(v3, (v43 & (v43 / (((v43 & 0)) | 1))))
        _ = v44
    } else {
        v45 := []int{9223372036854775807, -8, 773}
        _ = v45
    }
    v41 += v41
    fmt.Println(v43, v41, v1)
    if v42 {
        switch (v1 + (v1 + v1)) {
        case "a":
            for fi46 := 0; fi46 < 7; fi46++ {
                switch (fi46 % ((4294967295) | 1)) {
                case -1:
                    var v47 bool = false
                    _ = v47
                case 1:
                    v5 -= (v5 & 32767)
                default:
                    if (((31 > 159) && (uint64(v43) < 2)) && (!(888 != 42))) {
                        var v48 uint = 0
                        _ = v48
                    }
                }
                if true {
                    var v49 uint = 0
                    if (int8(v5) >= (^(v35))) {
                        fmt.Println(v49)
                    }
                    _ = v49
                } else {
                    switch "foo bar" {
                    case "a":
                        switch "\n\t" {
                        case "foo bar":
                            if (int16(v43) > v5) {
                                v50 := []int{fi46}
                                var v51 int64 = (v41 * -342)
                                _ = v50
                                _ = v51
                            } else {
                                v41 += (v41 + 246)
                            }
                        }
                    case "hello":
                        v2[v43] = v43
                    case "":
                        var v52 int = (fi46 >> ((65535) & 63))
                        _ = v52
                    default:
                        switch int(v43) {
                        case 623:
                            if (true && v42) {
                                v53 := v4[(fi46 * fi46)]
                                v1 = v1
                                var v54 float32 = 1e10
                                _ = v53
                                _ = v54
                            } else {
                                switch (fi46 & fi46) {
                                case -586:
                                    var v55 float32 = -2.25
                                    _ = v55
                                default:
                                    v56 := []int{(-(fi46)), fi46, fi46, 128}
                                    _ = v56
                                }
                                fmt.Println(v41)
                                if v42 {
                                    var v57 float64 = 1e10
                                    var v58 int8 = int8(v43)
                                    _ = v57
                                    _ = v58
                                } else {
                                    if v42 {
                                        var v59 uint16 = 32767
                                        switch 9223372036854775807 {
                                        case 45:
                                            switch (v1 + v1) {
                                            case "":
                                                v43 = v43
                                            }
                                        case -295:
                                            v60 := []int{fi46, (-(fi46))}
                                            _ = v60
                                        }
                                        var v61 uint16 = (v59 - (v59 - (v59 % (((v59 | 65535)) | 1))))
                                        _ = v59
                                        _ = v61
                                    }
                                    var v62 int = fi46
                                    _ = v62
                                }
                            }
                        case 32767:
                            switch "\n\t" {
                            case "a":
                                for fi63 := 0; fi63 < 17; fi63++ {
                                    switch (fi46 + (fi63 >> ((0) & 63))) {
                                    case 42:
                                        var v64 float32 = 1e10
                                        _ = v64
                                    case -304:
                                        v3[v43] = (!v42)
                                    case 2147483647:
                                        v41 += v41
                                    default:
                                        fmt.Println(v42, v1, v35)
                                        v43 = v43
                                    }
                                    var v65 int = fi63
                                    var v66 int32 = -2147483648
                                    var v67 float32 = 1e10
                                    var v68 int32 = int32(v35)
                                    _ = fi63
                                    _ = v65
                                    _ = v66
                                    _ = v67
                                    _ = v68
                                }
                            default:
                                v3[462] = v42
                            }
                        case 4294967295:
                            for fi69 := 0; fi69 < 14; fi69++ {
                                v70 := []int{}
                                v43 = (^(v43))
                                v71 := GenS6{F0: 42, F1: 237, F2: 42}
                                _ = fi69
                                _ = v70
                                _ = v71
                            }
                        }
                    }
                    var v72 uint64 = 18446744073709551615
                    v72 -= 1
                    _ = v72
                }
                v73 := GenS6{F0: 261, F1: 27, F2: (^(v43))}
                _ = fi46
                _ = v73
            }
        case "\n\t":
            v1 = ((v1 + "") + v1)
        case "hello":
            v5 -= (v5 / (((v5 << ((374) & 63))) | 1))
        default:
            v74 := GenS7{F0: v1}
            _ = v74
        }
    } else {
        switch ((v1 + ("foo bar" + (v1 + ("a" + "hello")))) + ("hello" + v1)) {
        case "\n\t":
            v75 := GenS1{F0: 60}
            _ = v75
        case "hello":
            v76 := GenS8{F0: v1, F1: v41}
            _ = v76
        case "foo bar":
            v42 = (181 != 864)
        }
    }
    v4[-1] = v42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
}
