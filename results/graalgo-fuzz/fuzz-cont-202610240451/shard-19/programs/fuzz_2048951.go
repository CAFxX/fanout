package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 int; F2 uint64 }
type GenS4 struct { F0 uint; F1 int32; F2 int64 }
type GenS5 struct { F0 uint32; F1 int8; F2 uint16 }
type GenS6 struct { F0 uint8; F1 int8; F2 uint8 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint64; F1 uint32 }
type GenS9 struct { F0 int8 }
func genFn0(p0 string, p1 bool, p2 int64) int64 {
    for fi7 := 0; fi7 < 9; fi7++ {
        fmt.Println(p0, p1)
        p1 = ((p2 / ((p2) | 1)) <= (p2 >> ((18446744073709551615) & 63)))
        if (!p1) {
            for fi8 := 0; fi8 < 14; fi8++ {
                if p1 {
                    switch fi8 {
                    case -270:
                        v9 := []int{(fi8 >> ((42) & 63))}
                        _ = v9
                    case 1:
                        p1 = (!true)
                    }
                } else {
                    p0 = (p0 + "foo bar")
                    var v10 float64 = 1e10
                    p2 = -106
                    _ = v10
                }
                _ = fi8
            }
            var v11 uint = 425
            if p1 {
                for fi12 := 0; fi12 < 5; fi12++ {
                    v13 := GenS1{F0: 1}
                    _ = fi12
                    _ = v13
                }
                v14 := []int{fi7, 128, fi7, fi7}
                _ = v14
            }
            _ = v11
        }
        _ = fi7
    }
    var v15 int = -1
    if (1.5 > 1e10) {
        v16 := GenS2{F0: 401}
        var v17 float32 = float32(float64(1.5))
        v18 := []int{}
        _ = v16
        _ = v17
        _ = v18
    } else {
        switch p0 {
        case "foo bar":
            v15 = (v15 + v15)
        }
        p0 = ("hello" + "\n\t")
        v15 = -882
    }
    if true {
        for fi19 := 0; fi19 < 3; fi19++ {
            var v20 int = 128
            _ = fi19
            _ = v20
        }
        p2 -= p2
    } else {
        for fi21 := 0; fi21 < 6; fi21++ {
            fmt.Println(p0)
            v22 := GenS2{F0: 18446744073709551615}
            if (true && p1) {
                p2 = p2
            } else {
                v23 := GenS1{F0: (p2 - (p2 & p2))}
                _ = v23
            }
            _ = fi21
            _ = v22
        }
        if (!true) {
            p0 = ""
            p2 = p2
            var v24 uint = 4294967295
            _ = v24
        } else {
            var v25 int = v15
            v26 := []int{}
            _ = v25
            _ = v26
        }
    }
    var v27 float32 = ((0.1 * -2.25) - 1.5)
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v27
    return p2
}

func main() {
    var v1 int = 32767
    var v2 uint16 = 65535
    var v3 uint = 18446744073709551615
    v4 := []int{}
    var v5 int32 = 2147483647
    var v6 uint16 = v2
    fmt.Println(v2)
    v28 := GenS1{F0: 1}
    v5 += 1
    fmt.Println(v3)
    if ("hello" != ("" + ("\n\t" + ((("hello" + "\n\t") + "foo bar") + "\n\t")))) {
        for fi29 := 0; fi29 < 4; fi29++ {
            v6 += v6
            for fi30 := 0; fi30 < 3; fi30++ {
                var v31 uint = (v3 << ((7) & 63))
                var v32 int32 = 662
                _ = fi30
                _ = v31
                _ = v32
            }
            _ = fi29
        }
        for fi33 := 0; fi33 < 8; fi33++ {
            v34 := GenS3{F0: v3, F1: 0, F2: 960}
            switch 128 {
            case 128:
                v35 := GenS4{F0: (^(v3)), F1: (v5 * v5), F2: 42}
                _ = v35
            }
            v36 := GenS4{F0: 631, F1: (v5 ^ 126), F2: 42}
            _ = fi33
            _ = v34
            _ = v36
        }
        v3 = (^(v3))
    } else {
        switch (v1 ^ (v1 * 4294967295)) {
        case 0:
            v2 += 255
        default:
            v2 = uint16(v3)
        }
        switch 128 {
        case -938:
            if false {
                v37 := GenS4{F0: (v3 + (v3 | (^(v3)))), F1: v5, F2: -9223372036854775808}
                _ = v37
            }
        }
    }
    if (0 < 191) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v4 = append(v4, 32768)
    } else {
        v38 := GenS5{F0: 950, F1: 104, F2: v2}
        _ = v38
    }
    fmt.Println(v3)
    fmt.Println(v2, v3)
    switch "a" {
    case "foo bar":
        v2 = (v6 | (^(v2)))
    }
    v39 := GenS1{F0: 5}
    switch "\n\t" {
    case "foo bar":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case "a":
        v6 += v6
    case "":
        v2 -= (v2 * 255)
    }
    switch "hello" {
    case "foo bar":
        var v40 uint = v3
        _ = v40
    case "":
        v6 -= 2
    default:
        if false {
            var v41 uint64 = 244
            if false {
                if true {
                    fmt.Println(v3, v1, v1)
                    var v42 uint64 = (v41 % ((v41) | 1))
                    if (v1 == 230) {
                        if ((v1 + (-(v1))) < (-(v1))) {
                            for fi43 := 0; fi43 < 11; fi43++ {
                                v44 := GenS6{F0: 42, F1: 46, F2: 22}
                                var v45 uint16 = 392
                                _ = fi43
                                _ = v44
                                _ = v45
                            }
                            v46 := GenS4{F0: v3, F1: v5, F2: 1}
                            _ = v46
                        } else {
                            _ = len(v4)
                            var v47 int = v1
                            for fi48 := 0; fi48 < 4; fi48++ {
                                v49 := GenS7{F0: uint64(v2)}
                                for fi50 := 0; fi50 < 3; fi50++ {
                                    v1 -= (-(fi48))
                                    switch (("a" + "") + ("hello" + ("\n\t" + "a"))) {
                                    case "\n\t":
                                        v51 := GenS8{F0: (v42 + v42), F1: 481}
                                        _ = v51
                                    case "hello":
                                        var v52 int8 = -24
                                        _ = v52
                                    default:
                                        var v53 int16 = 0
                                        switch (("hello" + "foo bar") + ("" + "hello")) {
                                        case "":
                                            if ((v2 ^ (v6 & v2)) == (v6 / ((v6) | 1))) {
                                                v41 += v41
                                            } else {
                                                if true {
                                                    var v54 uint16 = 255
                                                    _ = v54
                                                } else {
                                                    var v55 int64 = 9223372036854775807
                                                    _ = v55
                                                }
                                                if (!true) {
                                                    var v56 uint16 = v6
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    _ = v56
                                                } else {
                                                    if len(v4) > 0 {
                                                        v4[0] = (fi48 & (v47 * (^(v47))))
                                                    }
                                                    var v57 uint64 = v42
                                                    _ = v57
                                                }
                                            }
                                        }
                                        _ = v53
                                    }
                                    _ = fi50
                                }
                                for fi58 := 0; fi58 < 6; fi58++ {
                                    var v59 int8 = 0
                                    switch (v1 * -2147483648) {
                                    case -523:
                                        v2 = 32767
                                    default:
                                        var v60 float64 = (1.5 - (float64(0.0) * 0.0))
                                        _ = v60
                                    }
                                    var v61 uint8 = 0
                                    var v62 uint32 = 2147483648
                                    if false {
                                        v6 -= (v6 + v6)
                                        v4 = append(v4, (v47 * v47))
                                    }
                                    _ = fi58
                                    _ = v59
                                    _ = v61
                                    _ = v62
                                }
                                _ = fi48
                                _ = v49
                            }
                            _ = v47
                        }
                        var v63 int64 = 0
                        _ = v63
                    } else {
                        var v64 int32 = v5
                        if (false || false) {
                            var v65 float64 = 3.14159
                            var v66 int32 = (v64 >> (((v3 - (^(v3)))) & 63))
                            _ = v65
                            _ = v66
                        }
                        for fi67 := 0; fi67 < 2; fi67++ {
                            v68 := GenS7{F0: (v41 & (v42 - v42))}
                            for fi69 := 0; fi69 < 10; fi69++ {
                                _ = len(v4)
                                _ = fi69
                            }
                            _ = fi67
                            _ = v68
                        }
                        _ = v64
                    }
                    _ = v42
                }
                var v70 uint16 = 255
                _ = v70
            } else {
                v71 := GenS8{F0: v41, F1: 2}
                if (71 >= 246) {
                    if true {
                        var v72 int16 = 0
                        _ = v72
                    } else {
                        fmt.Println(v3)
                        v73 := GenS9{F0: 42}
                        _ = v73
                    }
                } else {
                    switch v1 {
                    case 2147483647:
                        for fi74 := 0; fi74 < 1; fi74++ {
                            v3 -= v3
                            var v75 float64 = float64(float32(1e10))
                            fmt.Println(fi74, v41, v1)
                            _ = fi74
                            _ = v75
                        }
                    case 127:
                        var v76 int16 = int16(v3)
                        _ = v76
                    }
                    v77 := GenS1{F0: 0}
                    _ = v77
                }
                _ = v71
            }
            _ = v41
        }
    }
    v78 := GenS6{F0: 128, F1: 1, F2: 42}
    fmt.Println(v6, v2)
    if false {
        v79 := GenS1{F0: 42}
        _ = v79
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v78
}
