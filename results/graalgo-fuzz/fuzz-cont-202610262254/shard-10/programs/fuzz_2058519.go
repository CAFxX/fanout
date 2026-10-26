package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 bool; F2 int16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int8; F1 uint; F2 uint16 }
func genFn0(p0 uint8, p1 string) int32 {
    for fi7 := 0; fi7 < 11; fi7++ {
        switch p1 {
        case "hello":
            p1 = "foo bar"
        case "\n\t":
            fmt.Println(p0)
        case "foo bar":
            if (1 != 18446744073709551615) {
                v8 := GenS1{F0: -379, F1: 490, F2: 42}
                v9 := []int{(fi7 + -190), fi7, 867}
                fmt.Println(p1, fi7, p0)
                _ = v8
                _ = v9
            } else {
                var v10 string = (p1 + "foo bar")
                v11 := []int{}
                _ = v10
                _ = v11
            }
        }
        if true {
            if true {
                switch p1 {
                case "":
                    fmt.Println(p1, p0, p0)
                }
                var v12 float64 = (3.14159 - -2.25)
                var v13 bool = (1e10 < 1e10)
                _ = v12
                _ = v13
            } else {
                fmt.Println(p1)
                var v14 float64 = -2.25
                _ = v14
            }
        }
        switch p1 {
        case "foo bar":
            var v15 string = (p1 + p1)
            _ = v15
        default:
            switch ("a" + p1) {
            case "\n\t":
                var v16 uint16 = 1
                _ = v16
            case "":
                var v17 int = (fi7 % (((fi7 & fi7)) | 1))
                _ = v17
            case "hello":
                v18 := []int{}
                _ = v18
            }
        }
        _ = fi7
    }
    v19 := GenS2{F0: -128}
    _ = p0
    _ = p1
    _ = v19
    return 2147483647
}

func main() {
    var v1 uint32 = 2147483648
    var v2 int8 = -128
    var v3 uint32 = v1
    var v4 uint16 = 796
    var v5 uint8 = 255
    v6 := []uint32{}
    fmt.Println(v1, v1, v3)
    for fi20 := 0; fi20 < 7; fi20++ {
        for fi21 := 0; fi21 < 20; fi21++ {
            var v22 int16 = -32768
            if false {
                var v23 int8 = v2
                if true {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    v2 = v2
                    fmt.Println(v22, v22)
                }
                _ = len(v6)
                _ = v23
            } else {
                fmt.Println(v5, v3, v3)
                v5 += v5
            }
            _ = fi21
            _ = v22
        }
        _ = fi20
    }
    fmt.Println(v1)
    var v24 int64 = -145
    var v25 string = "a"
    if ((!(!false)) || ((("" + "a") + v25) == v25)) {
        for fi26 := 0; fi26 < 16; fi26++ {
            v1 -= (v1 | (v1 << ((295) & 63)))
            if ((v1 < (^(v3))) && false) {
                var v27 uint16 = 1
                for fi28 := 0; fi28 < 16; fi28++ {
                    _ = len(v6)
                    var v29 bool = (912 >= 0)
                    _ = len(v6)
                    _ = fi28
                    _ = v29
                }
                _ = v27
            }
            _ = fi26
        }
        if (2 == 0) {
            var v30 float32 = (1e10 / 1.0)
            v3 = v3
            _ = v30
        } else {
            for fi31 := 0; fi31 < 17; fi31++ {
                if (966 > 42) {
                    fmt.Println(v5)
                }
                _ = fi31
            }
            for fi32 := 0; fi32 < 8; fi32++ {
                v33 := GenS2{F0: v2}
                _ = len(v6)
                var v34 int32 = 2147483647
                _ = fi32
                _ = v33
                _ = v34
            }
        }
        fmt.Println(v3, v1, v1)
    }
    var v35 uint64 = 0
    v36 := GenS3{F0: (((v24 * (v24 / (((v24 | 0)) | 1))) < (-(v24))) && true), F1: false, F2: -1}
    var v37 float64 = 0.1
    if false {
        var v38 int64 = v24
        fmt.Println(v4, v5, v3)
        _ = v38
    }
    switch 32768 {
    case 2147483647:
        var v39 float64 = ((v37 - 1e10) * 1.5)
        _ = v39
    default:
        var v40 uint64 = (v35 ^ 18446744073709551615)
        _ = v40
    }
    v41 := GenS4{F0: (v5 ^ 42)}
    switch ("foo bar" + v25) {
    case "":
        if true {
            if len(v6) > 0 {
                v6[0] = (v1 | 2147483648)
            }
        }
    case "foo bar":
        var v42 uint8 = (v5 / (((v5 << ((42) & 63))) | 1))
        _ = v42
    case "\n\t":
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    default:
        v43 := GenS5{F0: v2, F1: 255, F2: v4}
        _ = v43
    }
    v5 = 127
    if (v37 < 3.14159) {
        for fi44 := 0; fi44 < 8; fi44++ {
            var v45 int16 = 42
            if false {
                switch "a" {
                case "a":
                    fmt.Println(v3, fi44, v1)
                case "":
                    var v46 bool = true
                    _ = v46
                case "foo bar":
                    fmt.Println(v35, v35)
                }
            } else {
                var v47 int = (fi44 - (fi44 - fi44))
                fmt.Println(v35, v3)
                switch (fi44 / (((v47 / ((v47) | 1))) | 1)) {
                case 833:
                    for fi48 := 0; fi48 < 3; fi48++ {
                        if len(v6) > 0 {
                            v6[0] = v3
                        }
                        var v49 uint64 = (v35 - (v35 & (v35 * (v35 | (v35 & (v35 * (v35 + (v35 / ((0) | 1)))))))))
                        switch fi44 {
                        case -2147483648:
                            v2 = v2
                        case 32767:
                            var v50 int8 = 100
                            _ = v50
                        case 1:
                            if (false && ((2147483647 > 0) || (!true))) {
                                if (false && ("\n\t" != ("a" + v25))) {
                                    _ = len(v6)
                                    var v51 int16 = (^(v45))
                                    _ = v51
                                }
                                v2 += (v2 + v2)
                                var v52 float32 = (1.5 / (((((((1.5 / float32((0.1 / v37))) * (0.0 + 1.5)) + 0.0) * 1e10) + 1e10) - -2.25) / (1e10 - float32(0.0))))
                                var v53 int8 = v2
                                _ = v52
                                _ = v53
                            }
                        default:
                            var v54 float64 = v37
                            _ = v54
                        }
                        _ = fi48
                        _ = v49
                    }
                case 9223372036854775807:
                    _ = len(v6)
                }
                _ = v47
            }
            v55 := GenS3{F0: (((!((fi44 < fi44) || true)) && ((1e10 != 0.1) || (false || true))) && true), F1: false, F2: v45}
            _ = fi44
            _ = v45
            _ = v55
        }
    }
    fmt.Println(v35, v3)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if (!(127 > (-(v2)))) {
        for fi56 := 0; fi56 < 3; fi56++ {
            v57 := GenS1{F0: (^(fi56)), F1: 1, F2: 0}
            _ = fi56
            _ = v57
        }
    } else {
        var v58 uint8 = (^(v5))
        _ = len(v6)
        _ = v58
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
