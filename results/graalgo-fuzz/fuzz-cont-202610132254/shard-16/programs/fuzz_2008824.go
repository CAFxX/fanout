package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 int32; F1 uint16; F2 int16 }
type GenS6 struct { F0 int16; F1 int16; F2 int32 }
type GenS7 struct { F0 int32; F1 uint64; F2 uint }

func main() {
    var v1 int64 = 1
    var v2 float64 = float64(1e10)
    var v3 uint32 = 924
    var v4 int64 = v1
    var v5 bool = true
    for fi6 := 0; fi6 < 15; fi6++ {
        for fi7 := 0; fi7 < 15; fi7++ {
            if (!v5) {
                switch fi7 {
                case -492:
                    fmt.Println(v5, fi7, v3)
                case -776:
                    fmt.Println(v4, v1)
                case -795:
                    switch (("a" + "a") + "a") {
                    case "\n\t":
                        v1 = v4
                    }
                default:
                    v8 := []int{fi7, int(v4)}
                    _ = v8
                }
            } else {
                switch -332 {
                case 127:
                    var v9 uint8 = uint8(v3)
                    _ = v9
                case -2147483648:
                    if true {
                        v10 := []int{fi6}
                        v10 = append(v10, 0)
                        _ = v10
                    } else {
                        if (3.14159 < -2.25) {
                            fmt.Println(v5, v3)
                            v11 := GenS1{F0: 18446744073709551615, F1: 0}
                            var v12 uint32 = 65535
                            v4 = (^(v1))
                            if (0 >= -52) {
                                if false {
                                    if ((((((v2 - 0.1) / -2.25) / 3.14159) + v2) > (v2 - v2)) || false) {
                                        fmt.Println(v5)
                                        var v13 int64 = (v1 + v4)
                                        _ = v13
                                    } else {
                                        if v5 {
                                            fmt.Println(v4, v2, v4)
                                            switch -215 {
                                            case 601:
                                                var v14 uint32 = (v3 & (v12 >> ((4294967295) & 63)))
                                                _ = v14
                                            }
                                            fmt.Println(v5, v1, v5)
                                        }
                                    }
                                    if (0 < 1000) {
                                        v15 := []int{-914, fi6}
                                        _ = v15
                                    } else {
                                        v16 := []int{fi6, fi6}
                                        _ = v16
                                    }
                                    v17 := []int{(fi7 | 42), (fi6 / ((32768) | 1))}
                                    _ = v17
                                }
                                var v18 string = "hello"
                                _ = v18
                            }
                            v19 := GenS2{F0: ((-1 > -500) && false)}
                            if (!(!(132 < 42))) {
                                var v20 uint16 = 42
                                _ = v20
                            }
                            _ = v11
                            _ = v12
                            _ = v19
                        }
                        switch "" {
                        case "a":
                            fmt.Println(v5, v3, v3)
                        }
                        v21 := []int{fi6, (fi7 ^ (fi6 >> ((282) & 63))), 2147483647, (-(fi7))}
                        _ = v21
                    }
                case 505:
                    var v22 uint = uint(v3)
                    _ = v22
                default:
                    fmt.Println(v5, v4, v1)
                }
            }
            _ = fi7
        }
        v3 -= v3
        v23 := GenS1{F0: 42, F1: 2}
        _ = fi6
        _ = v23
    }
    v24 := []int{2147483647, 127}
    fmt.Println(v1, v5, v5)
    v25 := GenS1{F0: 18446744073709551615, F1: 42}
    if (137 >= 128) {
        for fi26 := 0; fi26 < 10; fi26++ {
            if len(v24) >= 2 {
                v24 = v24[1:len(v24)]
            }
            v27 := GenS2{F0: v5}
            _ = fi26
            _ = v27
        }
        v28 := GenS1{F0: 4294967295, F1: 0}
        _ = len(v24)
        _ = v28
    } else {
        for fi29 := 0; fi29 < 3; fi29++ {
            switch (("a" + "\n\t") + ("foo bar" + "foo bar")) {
            case "hello":
                for fi30 := 0; fi30 < 20; fi30++ {
                    var v31 int32 = 919
                    _ = fi30
                    _ = v31
                }
            default:
                if (fi29 > 32768) {
                    for fi32 := 0; fi32 < 4; fi32++ {
                        v24 = append(v24, (fi32 ^ (fi32 | (fi32 - (fi32 + fi32)))))
                        v33 := GenS2{F0: (1 <= 32767)}
                        switch (fi32 * 32767) {
                        case 4294967295:
                            if false {
                                _ = len(v24)
                                if false {
                                    v34 := GenS3{F0: -1, F1: fi32}
                                    fmt.Println(v3, v2, v3)
                                    v4 = (v1 + (-(v1)))
                                    var v35 float64 = -2.25
                                    v3 = v3
                                    _ = len(v24)
                                    _ = v34
                                    _ = v35
                                }
                            } else {
                                v24 = append(v24, (fi32 | (fi32 ^ (fi32 | (fi32 - fi29)))))
                                var v36 int64 = (v4 & (v1 * v4))
                                v3 -= v3
                                v37 := GenS2{F0: false}
                                _ = v36
                                _ = v37
                            }
                        }
                        v38 := GenS4{F0: 65535, F1: (896 != v4)}
                        _ = fi32
                        _ = v33
                        _ = v38
                    }
                    fmt.Println(fi29)
                    var v39 uint = 24
                    _ = v39
                } else {
                    v40 := GenS2{F0: ((365 == fi29) || true)}
                    v41 := GenS5{F0: 42, F1: 2, F2: -89}
                    _ = v40
                    _ = v41
                }
            }
            switch ((("a" + "foo bar") + "\n\t") + ("a" + ("a" + "foo bar"))) {
            case "\n\t":
                v2 -= (v2 * 1e10)
            case "foo bar":
                var v42 int16 = 32767
                _ = v42
            default:
                v43 := GenS2{F0: ((-690 >= -2147483648) || v5)}
                _ = v43
            }
            _ = fi29
        }
        if (v2 != v2) {
            v1 -= v1
            fmt.Println(v3, v4, v4)
            for fi44 := 0; fi44 < 20; fi44++ {
                v45 := GenS2{F0: (false && v5)}
                if len(v24) > 0 {
                    v24[0] = fi44
                }
                _ = fi44
                _ = v45
            }
        } else {
            var v46 int64 = 854
            var v47 int16 = 32767
            _ = v46
            _ = v47
        }
    }
    v4 -= (v4 % (((v4 >> ((2) & 63))) | 1))
    fmt.Println(v1)
    v4 += -9223372036854775808
    fmt.Println(v4, v3)
    v48 := GenS4{F0: 42, F1: v5}
    v49 := GenS6{F0: -1, F1: -32768, F2: 0}
    v50 := GenS7{F0: -1, F1: 2, F2: 1}
    for fi51 := 0; fi51 < 15; fi51++ {
        if (!v5) {
            if ("hello" < ("\n\t" + "foo bar")) {
                var v52 int = 797
                for fi53 := 0; fi53 < 14; fi53++ {
                    fmt.Println(v1, v3, v5)
                    var v54 uint64 = 1
                    _ = fi53
                    _ = v54
                }
                v55 := GenS6{F0: 1000, F1: 980, F2: 0}
                _ = v52
                _ = v55
            }
            if len(v24) > 0 {
                v24[0] = 1
            }
        }
        var v56 uint64 = 2
        _ = fi51
        _ = v56
    }
    var v57 uint16 = 695
    v1 += v1
    var v58 int8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v57
    _ = v58
}
