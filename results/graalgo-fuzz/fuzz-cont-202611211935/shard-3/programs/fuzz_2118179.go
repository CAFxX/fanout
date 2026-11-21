package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int16 }
type GenS2 struct { F0 uint64; F1 int32; F2 uint }
type GenS3 struct { F0 int8; F1 int; F2 int32 }
type GenS4 struct { F0 int16; F1 uint8; F2 int8 }
type GenS5 struct { F0 uint16; F1 uint8; F2 bool }
type GenS6 struct { F0 int8; F1 int16; F2 uint8 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int8; F1 int16; F2 int8 }
type GenS9 struct { F0 int16; F1 uint64; F2 int16 }
type GenS10 struct { F0 int32; F1 bool }
type GenS11 struct { F0 int16; F1 string; F2 string }
func genFn0(p0 int64, p1 uint32, p2 bool) uint32 {
    switch -870 {
    case -493:
        v7 := []int{-9223372036854775808, 4294967295, -1, 32767}
        _ = v7
    case 713:
        v8 := GenS1{F0: 2147483647, F1: p0, F2: 596}
        _ = v8
    case 223:
        v9 := GenS1{F0: 1, F1: (p0 * (p0 - (p0 - 42))), F2: -1}
        _ = v9
    default:
        v10 := GenS1{F0: 32768, F1: p0, F2: 1000}
        _ = v10
    }
    fmt.Println(p1, p0, p2)
    v11 := GenS1{F0: -1, F1: (p0 | 672), F2: 60}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return p1
}
func genFn1(p0 int64, p1 int16, p2 bool) int32 {
    if (0 < 42) {
        switch 636 {
        case 624:
            var v12 float64 = 0.0
            _ = v12
        case 128:
            v13 := []int{40, 127, 4294967295}
            _ = v13
        default:
            v14 := GenS2{F0: 921, F1: 331, F2: 445}
            _ = v14
        }
        fmt.Println(p0)
        fmt.Println(p0, p2, p1)
    } else {
        v15 := []int{0, -2147483648, 32767, -402}
        var v16 uint32 = 65535
        var v17 string = ("\n\t" + "\n\t")
        _ = v15
        _ = v16
        _ = v17
    }
    switch -713 {
    case 0:
        v18 := []int{}
        _ = v18
    }
    var v19 bool = (4294967295 > 1)
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    return -2147483648
}

func main() {
    v1 := []string{"foo bar"}
    v2 := []uint32{65535, 65535}
    v3 := []int{9223372036854775807, 756}
    var v4 int = -2147483648
    var v5 bool = ((-22 != 42) && ((42 < 42) && true))
    var v6 uint8 = 0
    if v5 {
        fmt.Println(v4)
        if ((v5 && v5) || (!("hello" >= "\n\t"))) {
            if (301 == 42) {
                fmt.Println(v4, v5, v6)
                v20 := GenS1{F0: 982, F1: -1, F2: 32767}
                if (0.1 <= float64(0.0)) {
                    fmt.Println(v6, v4, v5)
                }
                _ = v20
            }
            switch 0 {
            case 42:
                v3 = append(v3, (v4 * (v4 | (v4 << ((42) & 63)))))
            case 32768:
                v21 := GenS1{F0: (v4 | (v4 ^ 619)), F1: 9223372036854775807, F2: -1}
                _ = v21
            }
            if len(v2) > 0 {
                v2[0] = 4294967295
            }
        }
    }
    v22 := GenS3{F0: -6, F1: -816, F2: -174}
    for fi23 := 0; fi23 < 16; fi23++ {
        var v24 int32 = 42
        var v25 bool = (566 == 420)
        _ = fi23
        _ = v24
        _ = v25
    }
    fmt.Println(v4, v4, v5)
    v5 = v5
    v6 -= 0
    v26 := GenS4{F0: -32768, F1: (v6 | v6), F2: -128}
    if (false || v5) {
        var v27 uint8 = 68
        v28 := GenS5{F0: 210, F1: 1, F2: true}
        _ = v27
        _ = v28
    }
    switch 1 {
    case -9223372036854775808:
        if len(v3) > 0 {
            v3[0] = (v4 >> ((0) & 63))
        }
    }
    if (-55 == -1) {
        v29 := GenS2{F0: 0, F1: 1, F2: 755}
        var v30 uint = 65535
        switch v4 {
        case -1:
            v31 := GenS6{F0: 127, F1: 32767, F2: uint8(v30)}
            _ = v31
        case -877:
            v4 -= v4
        default:
            for fi32 := 0; fi32 < 18; fi32++ {
                var v33 int8 = 1
                switch "hello" {
                case "hello":
                    switch (^(v4)) {
                    case -9223372036854775808:
                        if len(v3) > 0 {
                            v3[0] = fi32
                        }
                    }
                default:
                    if (!(245 > 0)) {
                        v34 := GenS7{F0: 9223372036854775807}
                        v5 = (!true)
                        fmt.Println(v5)
                        _ = v34
                    } else {
                        if (!v5) {
                            v30 = (v30 ^ v30)
                            for fi35 := 0; fi35 < 7; fi35++ {
                                v36 := GenS8{F0: (-(v33)), F1: -241, F2: (v33 & v33)}
                                v37 := GenS5{F0: 2, F1: v6, F2: ((v33 << (((v30 | (v30 - v30))) & 63)) != -128)}
                                v38 := GenS9{F0: -70, F1: 657, F2: 664}
                                v30 = v30
                                _ = fi35
                                _ = v36
                                _ = v37
                                _ = v38
                            }
                        }
                    }
                }
                _ = fi32
                _ = v33
            }
        }
        _ = v29
        _ = v30
    }
    if v5 {
        v3 = append(v3, (v4 / ((-956) | 1)))
        _ = len(v3)
    }
    if (!(((v4 ^ 1) != (v4 % ((int(v6)) | 1))) || false)) {
        var v39 uint32 = 2
        v2 = append(v2, (v39 | v39))
        for fi40 := 0; fi40 < 16; fi40++ {
            if true {
                if v5 {
                    var v41 int64 = -9223372036854775808
                    _ = v41
                }
                if (("\n\t" + "foo bar") == ("a" + "hello")) {
                    var v42 int32 = -2147483648
                    v43 := GenS4{F0: 0, F1: v6, F2: -1}
                    _ = v42
                    _ = v43
                } else {
                    if false {
                        switch ("a" + "") {
                        case "a":
                            var v44 float64 = -2.25
                            _ = v44
                        case "":
                            for fi45 := 0; fi45 < 12; fi45++ {
                                if v5 {
                                    var v46 uint8 = 127
                                    var v47 int16 = 1
                                    var v48 uint = 65535
                                    _ = v46
                                    _ = v47
                                    _ = v48
                                }
                                _ = fi45
                            }
                        }
                        v6 -= v6
                        var v49 uint = 255
                        _ = v49
                    }
                }
                var v50 string = "foo bar"
                _ = v50
            } else {
                switch (fi40 * 2147483647) {
                case 394:
                    _ = len(v1)
                }
                var v51 int = v4
                _ = v51
            }
            v4 += (fi40 / ((v4) | 1))
            _ = fi40
        }
        _ = v39
    } else {
        if (0 > 42) {
            v52 := GenS8{F0: 127, F1: 1, F2: 1}
            v53 := GenS10{F0: -2147483648, F1: v5}
            _ = v52
            _ = v53
        }
        for fi54 := 0; fi54 < 15; fi54++ {
            v2 = append(v2, 2147483648)
            v4 = fi54
            _ = fi54
        }
        v55 := GenS11{F0: 42, F1: "a", F2: "hello"}
        _ = v55
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
