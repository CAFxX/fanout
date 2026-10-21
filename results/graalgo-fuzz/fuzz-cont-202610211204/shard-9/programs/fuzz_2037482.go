package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 bool; F2 int }
type GenS4 struct { F0 uint64; F1 uint16; F2 int64 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint8; F1 uint16; F2 bool }
type GenS8 struct { F0 int; F1 uint64; F2 int8 }
func genFn0(p0 int64, p1 string, p2 uint) int16 {
    var v4 int32 = -2147483648
    v5 := GenS1{F0: 150, F1: 42, F2: p0}
    v6 := GenS2{F0: 2}
    var v7 int16 = 32767
    switch ((("hello" + p1) + "") + (p1 + p1)) {
    case "\n\t":
        p2 -= p2
    case "hello":
        var v8 string = ""
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return v7
}

func main() {
    v1 := make(map[int]int)
    var v2 float64 = float64(3.14159)
    var v3 int32 = 752
    fmt.Println(v2, v3)
    var v9 string = (("" + (((((("foo bar" + ((((("\n\t" + ("hello" + "")) + "a") + "") + "") + "foo bar")) + "hello") + "a") + ("a" + "\n\t")) + "\n\t") + "")) + "")
    v9 = ((v9 + v9) + "")
    switch "\n\t" {
    case "a":
        v10 := []int{int(v3)}
        _ = v10
    default:
        if true {
            v11 := GenS1{F0: 128, F1: 2147483647, F2: -502}
            _ = v11
        }
    }
    v12 := v1[32768]
    v13 := GenS3{F0: "foo bar", F1: true, F2: (v12 - (v12 + 128))}
    var v14 int8 = 100
    v12 = v12
    for fi15 := 0; fi15 < 1; fi15++ {
        for fi16 := 0; fi16 < 20; fi16++ {
            if (32767 >= 42) {
                switch v9 {
                case "a":
                    if (v3 >= 202) {
                        if (false && true) {
                            fmt.Println(v9, v14)
                        }
                        var v17 int64 = 1
                        _ = v17
                    } else {
                        var v18 int8 = v14
                        switch (v12 << ((255) & 63)) {
                        case -1:
                            switch v12 {
                            case 0:
                                v19 := []int{fi16}
                                v14 = (^(v14))
                                v9 = v9
                                v20 := v1[(fi16 ^ (fi16 + fi16))]
                                var v21 int64 = 0
                                fmt.Println(v21, v20, v21)
                                _ = v19
                                _ = v20
                                _ = v21
                            case -596:
                                var v22 int64 = 1
                                var v23 int64 = v22
                                _ = v22
                                _ = v23
                            case 42:
                                if true {
                                    v1[(v12 & (fi15 ^ 316))] = -236
                                    v24 := GenS2{F0: 4294967295}
                                    if true {
                                        v25 := GenS4{F0: 720, F1: 65535, F2: -9223372036854775808}
                                        switch fi15 {
                                        case 32768:
                                            v9 = ""
                                        default:
                                            if (-1 <= 9223372036854775807) {
                                                v26 := []int{0, (fi15 + (fi16 >> ((4294967295) & 63))), (v12 - fi15), (fi16 << ((65535) & 63))}
                                                v12 = fi16
                                                var v27 float64 = v2
                                                v27 = v27
                                                _ = v26
                                                _ = v27
                                            } else {
                                                v28 := GenS3{F0: (v9 + v9), F1: (true && ((-9223372036854775808 >= 1) || true)), F2: (^(v12))}
                                                if true {
                                                    switch 42 {
                                                    case 222:
                                                        v29 := []int{-9223372036854775808, v12, 86}
                                                        _ = v29
                                                    }
                                                } else {
                                                    if true {
                                                        var v30 string = "foo bar"
                                                        if (18446744073709551615 == 1) {
                                                            if false {
                                                                fmt.Println(v12)
                                                            }
                                                            var v31 int8 = (v18 >> ((65535) & 63))
                                                            _ = len(v1)
                                                            _ = v31
                                                        } else {
                                                            fmt.Println(v3)
                                                        }
                                                        var v32 int = v12
                                                        _ = v30
                                                        _ = v32
                                                    }
                                                }
                                                _ = v28
                                            }
                                        }
                                        v33 := GenS5{F0: int32(v14)}
                                        _ = v25
                                        _ = v33
                                    } else {
                                        var v34 int8 = (v18 % (((^(v14))) | 1))
                                        fmt.Println(v14, v3)
                                        _ = v34
                                    }
                                    fmt.Println(v18)
                                    var v35 uint32 = 2147483647
                                    _ = v24
                                    _ = v35
                                } else {
                                    if false {
                                        v36 := v1[(fi16 & v12)]
                                        v18 += 100
                                        var v37 float64 = 0.0
                                        _ = v36
                                        _ = v37
                                    }
                                    if false {
                                        v38 := []int{}
                                        _ = v38
                                    } else {
                                        v39 := []int{(v12 ^ (^(fi16))), 153, (v12 / (((fi15 + fi15)) | 1)), (v12 & (^(fi16)))}
                                        switch v12 {
                                        case 32768:
                                            v40 := GenS6{F0: (true || false)}
                                            _ = v40
                                        case 127:
                                            fmt.Println(v3)
                                        default:
                                            v2 = v2
                                        }
                                        _ = v39
                                    }
                                }
                                switch (fi16 % ((4294967295) | 1)) {
                                case -1:
                                    _ = len(v1)
                                default:
                                    v9 = v9
                                }
                                switch v9 {
                                case "":
                                    v1[(fi16 & (fi16 / ((32767) | 1)))] = fi15
                                case "\n\t":
                                    v12 = 127
                                    v41 := GenS5{F0: -2147483648}
                                    v12 -= fi16
                                    _ = v41
                                }
                            }
                        case 476:
                            if (18446744073709551615 > 1) {
                                fmt.Println(v3)
                                var v42 int = fi15
                                var v43 uint = 82
                                if false {
                                    if (2147483647 == (v42 << (((v43 << ((v43) & 63))) & 63))) {
                                        if false {
                                            var v44 bool = false
                                            delete(v1, 1)
                                            _ = v44
                                        } else {
                                            switch "a" {
                                            case "\n\t":
                                                var v45 uint8 = 255
                                                _ = v45
                                            case "a":
                                                var v46 float32 = (3.14159 * 1.5)
                                                _ = v46
                                            default:
                                                if (v9 == "a") {
                                                    if (false && true) {
                                                        var v47 float64 = ((0.0 * 1e10) * 1.5)
                                                        if (true || (true || (2147483647 > 0))) {
                                                            var v48 uint8 = 22
                                                            _ = v48
                                                        }
                                                        _ = v47
                                                    }
                                                    v49 := GenS1{F0: 127, F1: 42, F2: -9223372036854775808}
                                                    _ = v49
                                                }
                                            }
                                        }
                                        var v50 uint64 = 0
                                        switch -1 {
                                        case 4294967295:
                                            v3 = (v3 >> ((v43) & 63))
                                        default:
                                            var v51 int16 = 814
                                            var v52 uint32 = 65535
                                            fmt.Println(v12, v2, v43)
                                            _ = v51
                                            _ = v52
                                        }
                                        _ = v50
                                    }
                                    v53 := GenS6{F0: true}
                                    fmt.Println(v2)
                                    _ = v53
                                } else {
                                    v9 = v9
                                }
                                _ = v42
                                _ = v43
                            } else {
                                fmt.Println(v3)
                            }
                        case 980:
                            if true {
                                v18 += -128
                                var v54 uint64 = 0
                                _ = v54
                            } else {
                                v55 := GenS7{F0: 128, F1: 541, F2: (1e10 >= 0.0)}
                                v3 = (v3 ^ v3)
                                fmt.Println(fi16, v18)
                                _ = v55
                            }
                        default:
                            v56 := GenS3{F0: (v9 + v9), F1: ((true && (0.0 != float32((((v2 - (float64(0.1) * v2)) * 0.1) * v2)))) || ((float64(0.0) + -2.25) < (v2 / v2))), F2: (^(v12))}
                            _ = v56
                        }
                        switch (fi15 ^ fi16) {
                        case -535:
                            fmt.Println(v18, fi16)
                        case 42:
                            v18 -= v18
                        }
                        _ = v18
                    }
                case "hello":
                    v3 += (v3 + 2147483647)
                }
                delete(v1, (fi16 >> ((18446744073709551615) & 63)))
                var v57 int16 = 1000
                _ = v57
            } else {
                var v58 uint32 = 1
                _ = v58
            }
            var v59 int64 = 509
            _ = fi16
            _ = v59
        }
        for fi60 := 0; fi60 < 6; fi60++ {
            v2 = ((float64(-2.25) - v2) - 1.5)
            if (true || false) {
                var v61 bool = false
                _ = v61
            }
            var v62 uint16 = 42
            _ = fi60
            _ = v62
        }
        _ = fi15
    }
    v63 := GenS8{F0: -161, F1: 232, F2: 1}
    v64 := GenS2{F0: 930}
    if (1 > 32767) {
        var v65 float32 = 1.5
        for fi66 := 0; fi66 < 6; fi66++ {
            v67 := []int{(v12 << ((4294967295) & 63)), (-(v12)), v12}
            _ = fi66
            _ = v67
        }
        v9 = ((v9 + "foo bar") + "hello")
        _ = v65
    } else {
        switch "\n\t" {
        case "foo bar":
            fmt.Println(v9, v2, v14)
        case "":
            if false {
                if true {
                    for fi68 := 0; fi68 < 1; fi68++ {
                        v9 = v9
                        _ = fi68
                    }
                    fmt.Println(v2, v12)
                }
            } else {
                v1[v12] = -103
            }
        default:
            var v69 uint32 = 0
            _ = v69
        }
    }
    for fi70 := 0; fi70 < 8; fi70++ {
        switch v9 {
        case "foo bar":
            v2 -= (0.1 + (float64(-2.25) * v2))
        case "\n\t":
            switch v9 {
            case "hello":
                v12 -= fi70
            default:
                v71 := GenS1{F0: 250, F1: 0, F2: -1}
                _ = v71
            }
        }
        switch v9 {
        case "foo bar":
            var v72 uint = 65535
            _ = v72
        default:
            if false {
                var v73 string = "foo bar"
                var v74 int64 = -263
                _ = v73
                _ = v74
            }
        }
        var v75 int64 = -599
        _ = fi70
        _ = v75
    }
    delete(v1, 128)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v63
    _ = v64
    _ = v9
}
