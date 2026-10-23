package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint64 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 bool; F1 int16 }
type GenS6 struct { F0 uint; F1 int16; F2 int64 }
type GenS7 struct { F0 int16; F1 string }

func main() {
    var v1 int8 = 126
    var v2 int32 = 2147483647
    v3 := make(map[string]bool)
    var v4 float32 = float32(1e10)
    v5 := GenS1{F0: 42, F1: 2, F2: 18446744073709551615}
    v1 = (v1 * (v1 >> ((42) & 63)))
    var v6 uint64 = 18446744073709551615
    var v7 uint8 = 2
    switch "" {
    case "foo bar":
        v1 = -4
    default:
        v2 = 483
    }
    if ((false && false) || (1.5 <= (1e10 / 1.0))) {
        fmt.Println(v4, v1, v4)
    } else {
        switch ("foo bar" + ("hello" + ("foo bar" + "a"))) {
        case "\n\t":
            switch ("" + (("a" + ("a" + "a")) + "a")) {
            case "hello":
                v7 -= v7
            }
        default:
            if true {
                for fi8 := 0; fi8 < 7; fi8++ {
                    if true {
                        for fi9 := 0; fi9 < 3; fi9++ {
                            if true {
                                if true {
                                    fmt.Println(fi9)
                                    if true {
                                        v10 := []int{(fi9 + (fi8 | (fi8 % (((^(fi8))) | 1)))), (-(fi9)), fi8, (fi9 % (((fi8 << ((18446744073709551615) & 63))) | 1))}
                                        var v11 float64 = (1.5 + 0.1)
                                        var v12 int8 = (-(v1))
                                        v2 += (v2 & v2)
                                        if len(v10) > 0 {
                                            v10[0] = fi8
                                        }
                                        _ = v10
                                        _ = v11
                                        _ = v12
                                    } else {
                                        v13 := GenS2{F0: 4294967295, F1: (v1 + v1)}
                                        if (!(!(!false))) {
                                            v14 := GenS3{F0: 4294967295, F1: 711}
                                            if false {
                                                switch (^(fi9)) {
                                                case 0:
                                                    var v15 float64 = (float64(float32(0.1)) - 1.5)
                                                    _ = v15
                                                default:
                                                    v1 = (-(v1))
                                                }
                                                v1 = v1
                                                var v16 uint32 = 2
                                                _ = v16
                                            }
                                            _ = v14
                                        }
                                        _ = v13
                                    }
                                    var v17 int = fi8
                                    _ = v17
                                }
                                v18 := GenS3{F0: 2, F1: 757}
                                if true {
                                    if false {
                                        if (false && (false && ((!false) && true))) {
                                            fmt.Println(fi8, v1)
                                            if ((v4 / v4) == (v4 / 3.14159)) {
                                                fmt.Println(v1, v6, v1)
                                                v19 := GenS1{F0: 2147483648, F1: 2, F2: v6}
                                                v7 += v7
                                                _ = v19
                                            } else {
                                                var v20 int32 = v2
                                                _ = v20
                                            }
                                            var v21 uint16 = 0
                                            _ = v21
                                        } else {
                                            if false {
                                                v6 -= (v6 + v6)
                                            }
                                        }
                                        v2 = 42
                                        switch fi9 {
                                        case 992:
                                            var v22 uint64 = v6
                                            _ = v22
                                        case 4294967295:
                                            v6 = 2
                                        case 32768:
                                            v23 := GenS3{F0: 65535, F1: 266}
                                            _ = v23
                                        default:
                                            if false {
                                                v24 := GenS4{F0: false}
                                                _ = v24
                                            } else {
                                                if true {
                                                    v25 := GenS2{F0: 347, F1: v1}
                                                    fmt.Println(v6, v2)
                                                    _ = v25
                                                } else {
                                                    v4 = (v4 * float32(0.0))
                                                }
                                                fmt.Println(v6, fi9, v6)
                                                var v26 string = ("\n\t" + "hello")
                                                switch "foo bar" {
                                                case "foo bar":
                                                    if false {
                                                        var v27 int = fi8
                                                        _ = v27
                                                    }
                                                case "":
                                                    var v28 uint32 = 1
                                                    _ = v28
                                                case "\n\t":
                                                    var v29 float32 = v4
                                                    _ = v29
                                                }
                                                _ = v26
                                            }
                                        }
                                    }
                                    if false {
                                        v30 := GenS1{F0: 1, F1: (^(v6)), F2: (v6 - 138)}
                                        v4 = float32(float64(((v4 * 1.5) * 0.0)))
                                        v31 := []int{}
                                        v2 += int32(fi9)
                                        _ = v30
                                        _ = v31
                                    } else {
                                        if false {
                                            var v32 bool = true
                                            v33 := []int{-9223372036854775808, (^(fi9)), (fi8 - -243)}
                                            _ = v32
                                            _ = v33
                                        } else {
                                            v3["a"] = ((v6 * v6) == (v6 * (^(v6))))
                                        }
                                    }
                                    fmt.Println(fi8, v4, v2)
                                }
                                switch "\n\t" {
                                case "a":
                                    v34 := []int{42, 127, (fi9 / (((fi8 & (-(fi8)))) | 1)), fi9}
                                    _ = v34
                                case "foo bar":
                                    if ("\n\t" < "\n\t") {
                                        v35 := GenS1{F0: 1, F1: 516, F2: v6}
                                        v6 = (v6 + (v6 << ((42) & 63)))
                                        _ = v35
                                    } else {
                                        v1 = v1
                                        if false {
                                            switch (^(fi9)) {
                                            case 651:
                                                v4 += 0.0
                                            default:
                                                v36 := v3[(("\n\t" + "\n\t") + "\n\t")]
                                                _ = v36
                                            }
                                            var v37 uint64 = v6
                                            var v38 float32 = ((1e10 + 0.0) / v4)
                                            _ = v37
                                            _ = v38
                                        } else {
                                            var v39 bool = (true || false)
                                            _ = v39
                                        }
                                    }
                                }
                                _ = v18
                            }
                            v40 := []int{fi9, (^(fi9))}
                            switch "" {
                            case "\n\t":
                                v41 := GenS5{F0: (!false), F1: 42}
                                _ = v41
                            default:
                                v42 := GenS5{F0: true, F1: -466}
                                _ = v42
                            }
                            _ = fi9
                            _ = v40
                        }
                        v43 := GenS4{F0: (!true)}
                        _ = v43
                    }
                    v44 := GenS4{F0: true}
                    _ = fi8
                    _ = v44
                }
            }
        }
        v45 := GenS3{F0: 255, F1: 0}
        _ = v45
    }
    v3["\n\t"] = ((v6 >> ((332) & 63)) < v6)
    if (!(false && true)) {
        v7 = 0
    }
    v46 := GenS6{F0: 504, F1: 0, F2: -544}
    v2 -= v2
    var v47 int16 = -1
    for fi48 := 0; fi48 < 14; fi48++ {
        v49 := []int{(fi48 & (fi48 - (fi48 % (((fi48 * 0)) | 1)))), 1, fi48}
        for fi50 := 0; fi50 < 10; fi50++ {
            if (!(true || true)) {
                v6 -= v6
            }
            var v51 uint = 989
            v4 = float32(0.1)
            _ = fi50
            _ = v51
        }
        v47 += v47
        _ = fi48
        _ = v49
    }
    for fi52 := 0; fi52 < 13; fi52++ {
        var v53 uint32 = 0
        for fi54 := 0; fi54 < 5; fi54++ {
            var v55 uint16 = 721
            switch ("a" + "\n\t") {
            case "\n\t":
                var v56 uint = 65535
                _ = v56
            case "foo bar":
                fmt.Println(fi54, v7)
            default:
                if false {
                    var v57 int64 = 1
                    fmt.Println(v6)
                    v55 += (^(v55))
                    if ((true && (true && true)) || true) {
                        if true {
                            switch (fi54 >> ((255) & 63)) {
                            case 42:
                                v6 = 436
                            default:
                                v4 -= (v4 + 1e10)
                            }
                            v58 := []int{42}
                            v59 := GenS7{F0: (^(v47)), F1: "foo bar"}
                            _ = v58
                            _ = v59
                        }
                    }
                    _ = v57
                }
            }
            _ = fi54
            _ = v55
        }
        _ = fi52
        _ = v53
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v6
    _ = v7
}
