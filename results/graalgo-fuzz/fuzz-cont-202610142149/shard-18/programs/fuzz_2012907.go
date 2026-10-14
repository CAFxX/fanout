package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 bool }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 uint; F1 int8; F2 int8 }
type GenS4 struct { F0 uint; F1 int; F2 int64 }
type GenS5 struct { F0 int16; F1 int8; F2 bool }
type GenS6 struct { F0 uint16; F1 int; F2 int }
type GenS7 struct { F0 int32 }
func genFn0() int32 {
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := []int{-843, (-(fi7))}
        var v9 float32 = -2.25
        var v10 uint32 = 600
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    var v11 uint32 = 2147483648
    var v12 int16 = 42
    v11 = 767
    v13 := GenS1{F0: 42, F1: 968, F2: false}
    _ = v11
    _ = v12
    _ = v13
    return -1
}
func genFn1() uint8 {
    switch 9223372036854775807 {
    case 32768:
        var v14 uint8 = 255
        _ = v14
    case 128:
        if false {
            for fi15 := 0; fi15 < 6; fi15++ {
                for fi16 := 0; fi16 < 17; fi16++ {
                    var v17 uint = 95
                    fmt.Println(fi16, fi15, fi16)
                    var v18 int16 = 1
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                for fi19 := 0; fi19 < 17; fi19++ {
                    fmt.Println(fi15, fi15)
                    fmt.Println(fi19, fi15)
                    if false {
                        if false {
                            v20 := GenS1{F0: 1, F1: 32767, F2: false}
                            v21 := GenS1{F0: -848, F1: 197, F2: (!true)}
                            if (true || (!(458 == 65535))) {
                                var v22 uint8 = 151
                                var v23 uint = 922
                                v24 := GenS1{F0: 9223372036854775807, F1: 42, F2: false}
                                _ = v22
                                _ = v23
                                _ = v24
                            } else {
                                v25 := GenS2{F0: 32767, F1: "a"}
                                switch "foo bar" {
                                case "\n\t":
                                    var v26 float32 = 1e10
                                    _ = v26
                                case "a":
                                    var v27 int = 42
                                    _ = v27
                                default:
                                    switch (fi19 & fi15) {
                                    case -2147483648:
                                        v28 := GenS2{F0: 255, F1: "hello"}
                                        _ = v28
                                    case -1:
                                        var v29 int = (-(fi15))
                                        _ = v29
                                    case 0:
                                        if true {
                                            v30 := []int{}
                                            v31 := GenS3{F0: 449, F1: 0, F2: -1}
                                            _ = v30
                                            _ = v31
                                        } else {
                                            if (((4294967295 == (fi15 * fi15)) && (1.5 >= 0.1)) && true) {
                                                v32 := GenS4{F0: 2, F1: 32768, F2: 193}
                                                var v33 int = -1
                                                v33 += 127
                                                _ = v32
                                                _ = v33
                                            }
                                            fmt.Println(fi15)
                                        }
                                    default:
                                        switch "" {
                                        case "\n\t":
                                            v34 := []int{(fi15 / ((32767) | 1)), fi15, fi19, fi19}
                                            _ = v34
                                        }
                                    }
                                }
                                _ = v25
                            }
                            _ = v20
                            _ = v21
                        }
                        if (780 > 532) {
                            v35 := GenS5{F0: 1000, F1: 0, F2: (("\n\t" != "hello") && true)}
                            fmt.Println(fi19)
                            var v36 uint64 = 18446744073709551615
                            _ = v35
                            _ = v36
                        } else {
                            var v37 int32 = 2147483647
                            _ = v37
                        }
                        var v38 uint8 = 42
                        switch int(v38) {
                        case -9223372036854775808:
                            if (4294967295 <= fi15) {
                                fmt.Println(fi15, fi15)
                                v39 := GenS3{F0: 42, F1: -4, F2: 100}
                                _ = v39
                            } else {
                                var v40 bool = false
                                v41 := GenS1{F0: -9223372036854775808, F1: 1, F2: true}
                                _ = v40
                                _ = v41
                            }
                        case 42:
                            if (4294967295 >= 42) {
                                v38 = (v38 / ((128) | 1))
                            }
                        case 911:
                            var v42 uint8 = 1
                            _ = v42
                        }
                        _ = v38
                    }
                    _ = fi19
                }
                _ = fi15
            }
            var v43 uint32 = 1
            _ = v43
        } else {
            v44 := GenS6{F0: 608, F1: 32767, F2: 1}
            var v45 uint16 = 92
            _ = v44
            _ = v45
        }
    }
    var v46 uint8 = 128
    var v47 float64 = ((1.5 * ((float64(1.5) + (((3.14159 + (((1e10 + 0.0) + (((1.5 * (float64(1e10) - 0.1)) / (1e10 / float64(1.5))) + (float64(((1.5 * 3.14159) - 1e10)) - 0.1))) + 0.1)) - -2.25) / float64(1.5))) + -2.25)) - float64(1e10))
    var v48 int8 = -1
    fmt.Println(v47, v47)
    _ = v46
    _ = v47
    _ = v48
    return 255
}

func main() {
    var v1 bool = (!(23 != 127))
    var v2 bool = true
    v3 := make(map[string]string)
    var v4 string = "foo bar"
    var v5 uint8 = 42
    var v6 uint16 = 263
    switch (v4 + (((v4 + v4) + "\n\t") + v4)) {
    case "foo bar":
        v49 := GenS4{F0: 255, F1: 9223372036854775807, F2: 0}
        _ = v49
    case "hello":
        for fi50 := 0; fi50 < 4; fi50++ {
            var v51 int8 = 127
            for fi52 := 0; fi52 < 5; fi52++ {
                v53 := v3[v4]
                switch int(v51) {
                case 2147483647:
                    v6 -= 335
                case 128:
                    var v54 int32 = 2147483647
                    _ = v54
                case 127:
                    var v55 uint8 = (v5 * v5)
                    _ = v55
                default:
                    v56 := []int{fi50}
                    _ = v56
                }
                if ((v53 + "hello") <= v53) {
                    var v57 bool = (-1 < 42)
                    v1 = v2
                    _ = v57
                } else {
                    v6 -= v6
                    var v58 int8 = int8(v5)
                    _ = v58
                }
                _ = fi52
                _ = v53
            }
            v1 = ((-943 <= 9223372036854775807) && v2)
            _ = fi50
            _ = v51
        }
    default:
        switch 128 {
        case 32768:
            for fi59 := 0; fi59 < 16; fi59++ {
                v4 = v4
                v60 := GenS4{F0: 1, F1: fi59, F2: int64(v6)}
                fmt.Println(v1, fi59)
                _ = fi59
                _ = v60
            }
        case -1:
            var v61 uint16 = v6
            _ = v61
        default:
            var v62 int = 127
            _ = v62
        }
    }
    if v1 {
        if ((v6 - v6) < (v6 + v6)) {
            v63 := GenS4{F0: uint(v6), F1: 2147483647, F2: -9223372036854775808}
            _ = v63
        } else {
            v64 := GenS2{F0: v6, F1: ("\n\t" + (((v4 + (v4 + (v4 + v4))) + v4) + v4))}
            if v2 {
                for fi65 := 0; fi65 < 18; fi65++ {
                    v4 = ((v4 + v4) + "")
                    v66 := GenS1{F0: -1, F1: v6, F2: v1}
                    for fi67 := 0; fi67 < 3; fi67++ {
                        var v68 uint64 = 588
                        fmt.Println(v4, v1, v1)
                        v68 = 4294967295
                        _ = fi67
                        _ = v68
                    }
                    _ = fi65
                    _ = v66
                }
                v69 := []int{-9223372036854775808, 127, 9223372036854775807}
                _ = v69
            } else {
                if v2 {
                    var v70 bool = ((float64((3.14159 + (3.14159 - -2.25))) * 3.14159) > 3.14159)
                    fmt.Println(v1, v6)
                    _ = v70
                } else {
                    v71 := GenS2{F0: v6, F1: "foo bar"}
                    _ = v71
                }
                var v72 int16 = 112
                _ = v72
            }
            _ = v64
        }
        var v73 string = v4
        _ = v73
    }
    fmt.Println(v2, v2, v4)
    v74 := []int{}
    v1 = v2
    if (v1 && (((61 >= -1) || false) || v2)) {
        if (4294967295 >= 860) {
            v75 := GenS4{F0: 42, F1: 79, F2: 709}
            _ = v75
        }
    }
    v2 = v2
    switch v4 {
    case "hello":
        for fi76 := 0; fi76 < 7; fi76++ {
            v6 += 42
            v2 = (!((!v2) || ((fi76 ^ fi76) == (fi76 + -366))))
            _ = fi76
        }
    case "\n\t":
        var v77 uint = 65535
        _ = v77
    case "a":
        fmt.Println(v4, v6)
    default:
        if (1 >= 104) {
            fmt.Println(v2, v1)
            _ = len(v74)
            var v78 int = 127
            _ = v78
        } else {
            v6 += (v6 ^ v6)
            fmt.Println(v5, v6, v4)
        }
    }
    fmt.Println(v4)
    v4 = v4
    v79 := GenS7{F0: 2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v74:", len(v74), v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v74
    _ = v79
}
