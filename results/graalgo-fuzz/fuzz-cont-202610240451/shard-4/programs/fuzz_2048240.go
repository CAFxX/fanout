package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 uint64; F2 int32 }
type GenS3 struct { F0 int64; F1 bool; F2 uint8 }
type GenS4 struct { F0 int64; F1 int64; F2 uint32 }
type GenS5 struct { F0 int; F1 uint32 }
type GenS6 struct { F0 uint64; F1 int32 }
type GenS7 struct { F0 uint64 }
func genFn0() int16 {
    v7 := []int{-691, -1, 0}
    if true {
        v8 := GenS1{F0: 999}
        var v9 float64 = -2.25
        for fi10 := 0; fi10 < 7; fi10++ {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = fi10
        }
        _ = v8
        _ = v9
    } else {
        v11 := GenS1{F0: 1}
        _ = v11
    }
    v7 = append(v7, 189)
    var v12 int16 = -1
    fmt.Println(v12)
    _ = v12
    _ = v7
    return 1
}
func genFn1(p0 int8) int32 {
    switch "foo bar" {
    case "":
        for fi13 := 0; fi13 < 19; fi13++ {
            v14 := []int{}
            if (508 <= 286) {
                var v15 int = (fi13 & (fi13 + 1))
                v14 = append(v14, (fi13 | v15))
                _ = v15
            }
            _ = fi13
            _ = v14
        }
    }
    var v16 int = 2147483647
    switch 2147483647 {
    case -871:
        var v17 string = "foo bar"
        _ = v17
    case -13:
        v18 := []int{(v16 ^ (-(v16))), v16, v16}
        _ = v18
    }
    _ = p0
    _ = v16
    return 1
}

func main() {
    var v1 float64 = ((3.14159 / 1.5) / float64(float32((0.1 * 3.14159))))
    var v2 bool = (-971 <= 137)
    var v3 float32 = 1.5
    var v4 float64 = 3.14159
    var v5 uint = 24
    var v6 uint8 = 0
    if false {
        if ((v6 >> ((v5) & 63)) != (v6 & (v6 * 185))) {
            v3 -= (float32(((v4 / v1) - (v4 - v4))) - 1e10)
            switch "\n\t" {
            case "hello":
                fmt.Println(v4)
            case "foo bar":
                v6 -= v6
            default:
                v19 := []int{960, 0, 2147483647}
                _ = v19
            }
            for fi20 := 0; fi20 < 6; fi20++ {
                v4 = v4
                v21 := []int{fi20, 32767}
                _ = fi20
                _ = v21
            }
        } else {
            for fi22 := 0; fi22 < 11; fi22++ {
                fmt.Println(v1, v6, v5)
                _ = fi22
            }
        }
    }
    switch 920 {
    case 128:
        if false {
            if v2 {
                for fi23 := 0; fi23 < 3; fi23++ {
                    v24 := []int{127}
                    var v25 float32 = float32(v1)
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            }
            if (!v2) {
                var v26 int8 = 0
                v6 += v6
                _ = v26
            } else {
                var v27 uint16 = 32767
                _ = v27
            }
        } else {
            v4 += (0.0 / 1.5)
            fmt.Println(v1, v3, v6)
            switch 32768 {
            case 2147483647:
                var v28 int32 = 42
                _ = v28
            case -208:
                fmt.Println(v6, v4)
            case 127:
                v4 = v4
            }
        }
    case 127:
        fmt.Println(v6, v5)
    }
    v29 := GenS2{F0: 1, F1: 1, F2: 42}
    if (42 == 0) {
        for fi30 := 0; fi30 < 7; fi30++ {
            var v31 bool = (v2 && (false && v2))
            v6 += 128
            _ = fi30
            _ = v31
        }
        if (869 != 2) {
            if v2 {
                for fi32 := 0; fi32 < 17; fi32++ {
                    fmt.Println(v6, fi32, v1)
                    switch "a" {
                    case "":
                        v33 := []int{(fi32 - (fi32 ^ (fi32 % ((fi32) | 1)))), (-(fi32)), (fi32 * fi32), (fi32 % ((fi32) | 1))}
                        _ = v33
                    case "hello":
                        if v2 {
                            var v34 float32 = v3
                            _ = v34
                        } else {
                            var v35 bool = ((fi32 % ((fi32) | 1)) < fi32)
                            for fi36 := 0; fi36 < 20; fi36++ {
                                v37 := GenS3{F0: -9223372036854775808, F1: v2, F2: (^(v6))}
                                var v38 uint64 = 0
                                _ = fi36
                                _ = v37
                                _ = v38
                            }
                            v39 := GenS3{F0: -595, F1: (!v35), F2: v6}
                            _ = v35
                            _ = v39
                        }
                    default:
                        v40 := GenS1{F0: 9223372036854775807}
                        _ = v40
                    }
                    var v41 int8 = 127
                    _ = fi32
                    _ = v41
                }
            } else {
                switch "foo bar" {
                case "":
                    var v42 bool = v2
                    _ = v42
                default:
                    v43 := GenS3{F0: 0, F1: (0 != 42), F2: (v6 * 84)}
                    _ = v43
                }
                var v44 float64 = (1e10 - ((0.0 - (float64(3.14159) - v1)) - 1e10))
                v45 := GenS4{F0: -1, F1: -9223372036854775808, F2: 2}
                _ = v44
                _ = v45
            }
        } else {
            switch ("a" + (("" + "hello") + "a")) {
            case "\n\t":
                fmt.Println(v6, v4)
            case "foo bar":
                var v46 float32 = float32((0.0 / float64(float32(float64(v3)))))
                _ = v46
            }
            for fi47 := 0; fi47 < 8; fi47++ {
                v48 := []int{127, (fi47 / (((fi47 >> (((v5 + (^(v5)))) & 63))) | 1))}
                _ = fi47
                _ = v48
            }
        }
        v4 = v1
    }
    v3 = (float32(v1) * v3)
    for fi49 := 0; fi49 < 10; fi49++ {
        for fi50 := 0; fi50 < 13; fi50++ {
            switch 128 {
            case 556:
                if (-2147483648 == (fi50 - -430)) {
                    var v51 uint8 = 2
                    fmt.Println(v4, v1, v3)
                    _ = v51
                }
            case 58:
                if v2 {
                    v2 = ((v2 || (true || (0 >= -246))) || v2)
                    var v52 float64 = v1
                    _ = v52
                } else {
                    v53 := GenS4{F0: 0, F1: 42, F2: 247}
                    _ = v53
                }
            case 128:
                switch 4294967295 {
                case -430:
                    switch (^(fi50)) {
                    case -1:
                        var v54 uint64 = 0
                        _ = v54
                    }
                    var v55 int64 = -515
                    v3 = v3
                    _ = v55
                case 32767:
                    v6 += v6
                case 882:
                    fmt.Println(v2)
                }
            }
            if v2 {
                v5 -= (v5 | (v5 >> (((v5 + 42)) & 63)))
                if (-1 <= 2147483647) {
                    v56 := []int{32768}
                    _ = v56
                } else {
                    var v57 uint = (v5 * (^(v5)))
                    v58 := GenS5{F0: 32768, F1: 0}
                    v59 := GenS1{F0: 9223372036854775807}
                    _ = v57
                    _ = v58
                    _ = v59
                }
                v60 := []int{fi50, fi49, fi50}
                _ = v60
            } else {
                var v61 uint = (v5 - (^(v5)))
                v62 := GenS6{F0: 904, F1: 0}
                _ = v61
                _ = v62
            }
            v1 = v1
            _ = fi50
        }
        v63 := GenS7{F0: 947}
        _ = fi49
        _ = v63
    }
    var v64 uint32 = 324
    switch -481 {
    case 624:
        var v65 string = (("\n\t" + "hello") + ("foo bar" + "foo bar"))
        _ = v65
    case -1:
        v5 = v5
    case 9223372036854775807:
        v1 = (((v1 / v1) / 3.14159) / ((1.5 * v1) + (v1 / 1.0)))
    default:
        if v2 {
            if v2 {
                for fi66 := 0; fi66 < 20; fi66++ {
                    var v67 int16 = 1000
                    _ = fi66
                    _ = v67
                }
            } else {
                v5 = 1
                var v68 uint16 = 1
                v69 := []int{}
                _ = v68
                _ = v69
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v64
}
