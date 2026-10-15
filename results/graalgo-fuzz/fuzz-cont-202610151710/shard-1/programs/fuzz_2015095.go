package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 uint32; F1 string; F2 int64 }
type GenS5 struct { F0 uint; F1 int16; F2 uint8 }
func genFn0(p0 uint) int {
    var v7 int = 127
    var v8 int = (v7 / (((v7 << ((255) & 63))) | 1))
    _ = p0
    _ = v7
    _ = v8
    return (-(v8))
}
func genFn1(p0 bool) int64 {
    p0 = (p0 && (399 < -9223372036854775808))
    if false {
        var v9 uint64 = 324
        v10 := GenS1{F0: ("\n\t" + ""), F1: p0}
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        var v12 uint16 = 983
        _ = fi11
        _ = v12
    }
    if p0 {
        p0 = (1 == 255)
        switch "" {
        case "\n\t":
            var v13 float64 = (3.14159 + 0.0)
            _ = v13
        case "a":
            var v14 uint64 = 1
            _ = v14
        default:
            for fi15 := 0; fi15 < 8; fi15++ {
                p0 = (924 < 0)
                switch "\n\t" {
                case "hello":
                    var v16 uint32 = 2147483647
                    _ = v16
                case "a":
                    switch ("foo bar" + "hello") {
                    case "foo bar":
                        var v17 uint32 = 438
                        _ = v17
                    }
                case "\n\t":
                    for fi18 := 0; fi18 < 12; fi18++ {
                        v19 := GenS1{F0: ("" + ("\n\t" + "foo bar")), F1: p0}
                        _ = fi18
                        _ = v19
                    }
                default:
                    if true {
                        if true {
                            p0 = (!p0)
                            var v20 uint = 173
                            _ = v20
                        }
                    } else {
                        v21 := []int{(fi15 | genFn0(0)), genFn0(1), genFn0(255)}
                        v22 := GenS2{F0: 146, F1: 44}
                        _ = v21
                        _ = v22
                    }
                }
                _ = fi15
            }
        }
    } else {
        for fi23 := 0; fi23 < 3; fi23++ {
            fmt.Println(fi23, p0)
            for fi24 := 0; fi24 < 9; fi24++ {
                var v25 uint16 = 2
                v26 := GenS3{F0: 681, F1: -1}
                var v27 bool = true
                _ = fi24
                _ = v25
                _ = v26
                _ = v27
            }
            v28 := GenS4{F0: 0, F1: ("a" + "\n\t"), F2: -9223372036854775808}
            _ = fi23
            _ = v28
        }
        v29 := []int{-2147483648, 615, 32767}
        p0 = false
        _ = v29
    }
    p0 = (1 < 1)
    _ = p0
    return 42
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]bool)
    v3 := []int{551, 182, 127, 42}
    v4 := []int{}
    v5 := []int{2147483647, -2147483648, 32768, 1}
    var v6 int64 = -193
    var v30 float32 = (-2.25 * (-2.25 * 0.0))
    var v31 int8 = -88
    for fi32 := 0; fi32 < 6; fi32++ {
        for fi33 := 0; fi33 < 19; fi33++ {
            v30 = ((v30 * v30) - (0.0 * v30))
            v31 = (v31 & (v31 / ((v31) | 1)))
            var v34 string = "\n\t"
            _ = fi33
            _ = v34
        }
        _ = fi32
    }
    switch -1 {
    case 42:
        var v35 bool = true
        _ = v35
    case -2147483648:
        v36 := GenS5{F0: 0, F1: 0, F2: 0}
        _ = v36
    case 0:
        for fi37 := 0; fi37 < 15; fi37++ {
            var v38 string = "a"
            if true {
                var v39 float32 = (v30 / (v30 * float32((1.5 / 1e10))))
                _ = v39
            } else {
                switch v38 {
                case "\n\t":
                    if ((!true) || (-1 > 697)) {
                        if len(v4) > 0 {
                            v4[0] = fi37
                        }
                        switch fi37 {
                        case 1:
                            switch fi37 {
                            case 796:
                                if (v30 < (v30 / (0.1 + 3.14159))) {
                                    var v40 uint8 = 42
                                    for fi41 := 0; fi41 < 17; fi41++ {
                                        switch v38 {
                                        case "":
                                            v42 := v2[fi41]
                                            _ = v42
                                        default:
                                            var v43 int64 = v6
                                            _ = v43
                                        }
                                        fmt.Println(fi41)
                                        _ = fi41
                                    }
                                    _ = v40
                                }
                            case 452:
                                fmt.Println(fi37, v30)
                            case -1:
                                if (4294967295 > fi37) {
                                    fmt.Println(fi37, v30, fi37)
                                }
                            default:
                                v6 -= v6
                            }
                        case 411:
                            fmt.Println(v31, fi37, v30)
                        case 2147483647:
                            for fi44 := 0; fi44 < 10; fi44++ {
                                v4 = append(v4, -2147483648)
                                _ = fi44
                            }
                        }
                    }
                case "foo bar":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                default:
                    if ((1e10 + 1.5) == 1.5) {
                        if true {
                            if (!(42 < 90)) {
                                v6 = (v6 ^ v6)
                                v38 = ((("a" + v38) + v38) + (v38 + (v38 + (("foo bar" + "foo bar") + (v38 + ("" + (v38 + (v38 + (v38 + v38)))))))))
                            }
                            v6 = (v6 & 42)
                        } else {
                            fmt.Println(v30, v38)
                        }
                    } else {
                        v38 = (v38 + "\n\t")
                        var v45 uint32 = 96
                        var v46 int64 = (v6 | v6)
                        _ = v45
                        _ = v46
                    }
                }
                if (true || (true || true)) {
                    var v47 int8 = (^(v31))
                    var v48 uint32 = 2147483647
                    _ = v47
                    _ = v48
                } else {
                    v49 := v2[genFn0(90)]
                    var v50 bool = (true && (224 == 1))
                    v51 := GenS3{F0: 32767, F1: (v31 - int8(fi37))}
                    _ = v49
                    _ = v50
                    _ = v51
                }
                switch genFn0(234) {
                case -2147483648:
                    var v52 int16 = -587
                    _ = v52
                case 861:
                    v31 = v31
                case 580:
                    switch "hello" {
                    case "foo bar":
                        var v53 int = -25
                        _ = v53
                    case "hello":
                        var v54 int = fi37
                        _ = v54
                    case "\n\t":
                        if (0 >= 180) {
                            var v55 int32 = 2147483647
                            var v56 string = "hello"
                            _ = v55
                            _ = v56
                        }
                    }
                }
            }
            _ = fi37
            _ = v38
        }
    }
    v4 = append(v4, 487)
    v30 -= 3.14159
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(v31)
    fmt.Println(v31)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
