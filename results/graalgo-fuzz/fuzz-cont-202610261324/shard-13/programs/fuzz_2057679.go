package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int32; F1 uint16; F2 int }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint }
func genFn0(p0 uint64, p1 uint) uint64 {
    var v6 int32 = 0
    v6 += (v6 ^ (v6 * 2147483647))
    _ = p0
    _ = p1
    _ = v6
    return (p0 >> ((p1) & 63))
}

func main() {
    var v1 bool = true
    var v2 int32 = -1
    var v3 float64 = 0.0
    v4 := make(map[int]int)
    var v5 string = ""
    if (v1 && (1 > 0)) {
        v1 = (false || false)
        v7 := GenS1{F0: 42, F1: -504}
        _ = v7
    } else {
        if true {
            var v8 int16 = 32767
            v2 -= v2
            if true {
                fmt.Println(v5, v1)
            } else {
                v9 := GenS2{F0: v5, F1: 2147483648}
                if true {
                    fmt.Println(v5, v8)
                    var v10 float32 = float32((v3 - 1.5))
                    _ = v10
                } else {
                    v11 := GenS3{F0: (0 == 42), F1: 83}
                    switch 32767 {
                    case 42:
                        v12 := []int{32767, 1, 942}
                        _ = v12
                    default:
                        if ("foo bar" == (("" + v5) + "\n\t")) {
                            var v13 int16 = (v8 * (v8 * v8))
                            delete(v4, 2147483647)
                            _ = v13
                        }
                    }
                    _ = v11
                }
                v14 := []int{4294967295, 42}
                _ = v14
                _ = v9
            }
            _ = v8
        }
        v2 = 757
        _ = len(v4)
    }
    if (true || true) {
        v15 := []int{1, 127, 24, -921}
        v3 = 3.14159
        _ = v15
    }
    v2 -= (v2 << ((18446744073709551615) & 63))
    v5 = (("a" + v5) + "foo bar")
    v4[127] = 9223372036854775807
    if ((0 != 42) && false) {
        v5 = v5
    } else {
        switch "a" {
        case "foo bar":
            _ = len(v4)
        }
        if (-29 == 127) {
            switch (v5 + v5) {
            case "":
                switch "foo bar" {
                case "":
                    if (42 <= 417) {
                        v1 = (0 > 0)
                    }
                case "hello":
                    v16 := []int{-376, 9223372036854775807, 0, -9223372036854775808}
                    _ = v16
                case "a":
                    if v1 {
                        v17 := []int{int(v2)}
                        var v18 int64 = int64(v2)
                        switch 13 {
                        case 729:
                            v4[300] = 0
                        case 609:
                            if len(v17) >= 2 {
                                v17 = v17[1:len(v17)]
                            }
                        }
                        _ = v17
                        _ = v18
                    } else {
                        switch v5 {
                        case "foo bar":
                            if v1 {
                                v19 := GenS4{F0: ((v2 << ((913) & 63)) > v2)}
                                v20 := GenS2{F0: v5, F1: 622}
                                _ = v19
                                _ = v20
                            } else {
                                if (!v1) {
                                    v2 = 0
                                    if (true && ((637 < 4294967295) && (v1 || (42 != int16(v2))))) {
                                        if true {
                                            if (42 < 42) {
                                                v3 += v3
                                            }
                                            for fi21 := 0; fi21 < 12; fi21++ {
                                                var v22 int = (^(fi21))
                                                v2 = (^(v2))
                                                _ = fi21
                                                _ = v22
                                            }
                                        }
                                    }
                                }
                                for fi23 := 0; fi23 < 14; fi23++ {
                                    v3 = (v3 / v3)
                                    fmt.Println(v2)
                                    fmt.Println(fi23)
                                    _ = fi23
                                }
                            }
                        case "hello":
                            v24 := GenS4{F0: false}
                            _ = v24
                        case "a":
                            var v25 int64 = -985
                            _ = v25
                        default:
                            _ = len(v4)
                        }
                        switch 128 {
                        case 2147483647:
                            for fi26 := 0; fi26 < 4; fi26++ {
                                v27 := []int{fi26}
                                _ = fi26
                                _ = v27
                            }
                        }
                        switch v5 {
                        case "a":
                            v28 := GenS4{F0: ((3.14159 <= float32(1.5)) || (!(v2 > (v2 << ((4294967295) & 63)))))}
                            _ = v28
                        case "":
                            if (!(1 >= 42)) {
                                var v29 float32 = (3.14159 + -2.25)
                                if true {
                                    v30 := v4[4294967295]
                                    if false {
                                        if (659 > 110) {
                                            v1 = (!(553 == 42))
                                            var v31 uint16 = 42
                                            _ = v31
                                        }
                                    } else {
                                        for fi32 := 0; fi32 < 19; fi32++ {
                                            if (((!(1 != 533)) || ((float32((1e10 * 0.0)) * v29) < (((v29 / v29) - v29) / (1e10 * (v29 * -2.25))))) || v1) {
                                                switch (v30 & (fi32 & (^(fi32)))) {
                                                case -515:
                                                    v33 := GenS5{F0: v2, F1: 32767, F2: v30}
                                                    _ = v33
                                                default:
                                                    var v34 float32 = v29
                                                    _ = v34
                                                }
                                                var v35 bool = v1
                                                _ = v35
                                            } else {
                                                v5 = v5
                                                v36 := GenS6{F0: 0}
                                                _ = v36
                                            }
                                            fmt.Println(v1)
                                            _ = fi32
                                        }
                                        v37 := []int{v30, v30}
                                        _ = v37
                                    }
                                    _ = v30
                                }
                                _ = v29
                            } else {
                                var v38 float32 = -2.25
                                if ("hello" > v5) {
                                    fmt.Println(v5, v5)
                                    v5 = "foo bar"
                                    if false {
                                        for fi39 := 0; fi39 < 16; fi39++ {
                                            _ = len(v4)
                                            _ = fi39
                                        }
                                        fmt.Println(v3, v38)
                                    } else {
                                        v40 := v4[42]
                                        for fi41 := 0; fi41 < 15; fi41++ {
                                            fmt.Println(v3, v1)
                                            var v42 uint = 1
                                            _ = fi41
                                            _ = v42
                                        }
                                        if true {
                                            switch v40 {
                                            case 2147483647:
                                                v43 := []int{(v40 | (v40 * v40))}
                                                _ = v43
                                            case 128:
                                                v1 = ((((-(v2)) != v2) || v1) || v1)
                                            default:
                                                var v44 int8 = 120
                                                _ = v44
                                            }
                                            v4[(v40 / ((42) | 1))] = 4294967295
                                        }
                                        _ = v40
                                    }
                                }
                                _ = v38
                            }
                        }
                    }
                default:
                    var v45 int32 = v2
                    _ = v45
                }
            case "a":
                if (723 != 2) {
                    v46 := []int{4294967295, -9223372036854775808, 32768}
                    _ = v46
                }
            }
            for fi47 := 0; fi47 < 17; fi47++ {
                for fi48 := 0; fi48 < 18; fi48++ {
                    _ = len(v4)
                    delete(v4, 4294967295)
                    fmt.Println(v1, v3, v5)
                    var v49 int16 = -1
                    v50 := GenS4{F0: v1}
                    v51 := []int{-9223372036854775808}
                    _ = fi48
                    _ = v49
                    _ = v50
                    _ = v51
                }
                if false {
                    if (!v1) {
                        var v52 uint16 = 42
                        fmt.Println(v52, v5)
                        _ = v52
                    } else {
                        for fi53 := 0; fi53 < 12; fi53++ {
                            if (984 <= 255) {
                                if (v1 && v1) {
                                    switch 9223372036854775807 {
                                    case 42:
                                        if ((v1 && v1) && v1) {
                                            if ((2 == 2147483647) && (v1 || v1)) {
                                                v54 := GenS2{F0: v5, F1: 2}
                                                v55 := GenS3{F0: true, F1: 0}
                                                _ = v54
                                                _ = v55
                                            } else {
                                                var v56 uint8 = 255
                                                v1 = (!(!(534 >= (fi47 / (((fi53 >> ((255) & 63))) | 1)))))
                                                _ = v56
                                            }
                                            v1 = false
                                            v2 = 1
                                        } else {
                                            var v57 int32 = -909
                                            _ = v57
                                        }
                                        v58 := GenS3{F0: v1, F1: 399}
                                        _ = v58
                                    default:
                                        fmt.Println(fi53, fi47, fi53)
                                    }
                                    v59 := GenS3{F0: true, F1: 42}
                                    _ = v59
                                }
                                v60 := GenS7{F0: 818}
                                fmt.Println(v5, v1, fi53)
                                var v61 float64 = 0.1
                                _ = v60
                                _ = v61
                            }
                            _ = fi53
                        }
                    }
                    if false {
                        var v62 uint32 = 2147483648
                        v62 = 4294967295
                        switch (("foo bar" + (v5 + v5)) + ((("\n\t" + ("\n\t" + "\n\t")) + v5) + (v5 + "a"))) {
                        case "foo bar":
                            fmt.Println(v1)
                        case "hello":
                            var v63 string = ""
                            _ = v63
                        default:
                            var v64 int = (-(fi47))
                            _ = v64
                        }
                        _ = v62
                    }
                } else {
                    v2 -= (v2 ^ (v2 / ((v2) | 1)))
                    v65 := []int{(fi47 | fi47), fi47}
                    _ = v65
                }
                _ = fi47
            }
            for fi66 := 0; fi66 < 2; fi66++ {
                for fi67 := 0; fi67 < 14; fi67++ {
                    _ = len(v4)
                    switch v5 {
                    case "foo bar":
                        v68 := GenS1{F0: -1, F1: 9223372036854775807}
                        _ = v68
                    case "a":
                        var v69 uint8 = 42
                        _ = v69
                    case "\n\t":
                        v70 := v4[-509]
                        _ = v70
                    default:
                        v1 = (v2 <= 42)
                    }
                    _ = fi67
                }
                var v71 uint = 2
                _ = fi66
                _ = v71
            }
        } else {
            for fi72 := 0; fi72 < 8; fi72++ {
                var v73 float64 = (float64((0.1 / (3.14159 * (3.14159 + (1e10 + 0.1))))) * ((3.14159 + (v3 * ((1.5 * float64(float32((float64(0.1) + (3.14159 * v3))))) / (v3 + (0.1 / v3))))) * v3))
                _ = fi72
                _ = v73
            }
            v74 := GenS2{F0: ((v5 + v5) + v5), F1: 0}
            _ = v74
        }
    }
    for fi75 := 0; fi75 < 15; fi75++ {
        v76 := GenS8{F0: ((v5 + v5) + v5)}
        _ = fi75
        _ = v76
    }
    v77 := GenS9{F0: 42}
    v78 := []int{}
    if (!(v1 && ((("a" + v5) + v5) < ("\n\t" + v5)))) {
        fmt.Println(v3, v2, v5)
        v79 := GenS1{F0: -1, F1: 988}
        _ = v79
    } else {
        v80 := GenS5{F0: v2, F1: 65535, F2: 10}
        if len(v78) > 0 {
            v78[0] = 1
        }
        _ = v80
    }
    var v81 string = (v5 + ("foo bar" + "\n\t"))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v78:", len(v78), v78)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v77
    _ = v78
    _ = v81
}
