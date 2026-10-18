package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 int8; F1 string; F2 int32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 string; F1 bool }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 bool; F1 int32 }
type GenS8 struct { F0 int16; F1 int32 }

func main() {
    var v1 int32 = 42
    var v2 float64 = 3.14159
    var v3 uint = 725
    v4 := GenS1{F0: 791, F1: v1}
    var v5 uint8 = 155
    v3 -= 932
    v1 += int32(v5)
    var v6 uint8 = (v5 + 16)
    var v7 uint32 = 2147483648
    v8 := GenS1{F0: -506, F1: v1}
    fmt.Println(v2, v1)
    for fi9 := 0; fi9 < 15; fi9++ {
        var v10 uint = 778
        fmt.Println(v6, v2, v1)
        if true {
            if true {
                var v11 int64 = -516
                if ((^(v3)) >= v3) {
                    var v12 uint32 = 2147483648
                    var v13 string = "\n\t"
                    fmt.Println(v5, v1)
                    _ = v12
                    _ = v13
                }
                v14 := []int{}
                _ = v11
                _ = v14
            } else {
                if true {
                    v2 -= v2
                    fmt.Println(v3)
                    var v15 int = -1
                    _ = v15
                }
                v16 := GenS2{F0: 2147483647, F1: 42}
                if false {
                    fmt.Println(fi9)
                    if false {
                        var v17 uint64 = 111
                        for fi18 := 0; fi18 < 14; fi18++ {
                            switch -9223372036854775808 {
                            case 32768:
                                var v19 float32 = 1.5
                                _ = v19
                            default:
                                switch fi9 {
                                case 32768:
                                    fmt.Println(v17, v2)
                                case 9223372036854775807:
                                    fmt.Println(v7)
                                default:
                                    if (v17 == (v17 + v17)) {
                                        v20 := []int{4294967295}
                                        switch (fi18 % ((int(v17)) | 1)) {
                                        case -1:
                                            v10 -= v3
                                        }
                                        _ = v20
                                    } else {
                                        if false {
                                            if true {
                                                fmt.Println(v3, v6)
                                            } else {
                                                v7 = v7
                                                var v21 uint16 = 42
                                                _ = v21
                                            }
                                            v22 := GenS2{F0: fi18, F1: int64(v3)}
                                            switch (fi9 % ((fi18) | 1)) {
                                            case 42:
                                                fmt.Println(v17, v3, v2)
                                            case -9223372036854775808:
                                                if (!((-9223372036854775808 > 0) || false)) {
                                                    fmt.Println(fi9, fi18, v5)
                                                    v23 := GenS3{F0: 100, F1: ("\n\t" + ""), F2: v1}
                                                    var v24 uint16 = 32767
                                                    _ = v23
                                                    _ = v24
                                                }
                                            case 547:
                                                v25 := GenS2{F0: fi9, F1: 1}
                                                _ = v25
                                            default:
                                                var v26 bool = true
                                                _ = v26
                                            }
                                            _ = v22
                                        }
                                        var v27 int = (fi18 ^ fi9)
                                        _ = v27
                                    }
                                }
                            }
                            v6 = v6
                            var v28 int8 = 1
                            _ = fi18
                            _ = v28
                        }
                        _ = v17
                    } else {
                        v29 := []int{}
                        v2 = (3.14159 * v2)
                        v30 := GenS4{F0: 42}
                        _ = v29
                        _ = v30
                    }
                    v31 := GenS5{F0: ("a" + ("foo bar" + "hello")), F1: (-464 == -379)}
                    _ = v31
                } else {
                    v32 := []int{}
                    var v33 uint64 = 42
                    v34 := GenS6{F0: v10}
                    _ = v32
                    _ = v33
                    _ = v34
                }
                _ = v16
            }
            var v35 uint32 = v7
            var v36 int16 = 42
            _ = v35
            _ = v36
        }
        _ = fi9
        _ = v10
    }
    fmt.Println(v3, v2)
    v37 := GenS5{F0: "", F1: false}
    if true {
        fmt.Println(v5)
        for fi38 := 0; fi38 < 8; fi38++ {
            switch "\n\t" {
            case "foo bar":
                var v39 int64 = -896
                _ = v39
            default:
                v7 = uint32(v3)
            }
            v40 := GenS7{F0: false, F1: v1}
            var v41 uint32 = (^(v7))
            _ = fi38
            _ = v40
            _ = v41
        }
    } else {
        if (-979 < -9223372036854775808) {
            for fi42 := 0; fi42 < 12; fi42++ {
                v43 := GenS2{F0: -437, F1: 1}
                _ = fi42
                _ = v43
            }
        } else {
            var v44 int32 = v1
            v5 += 127
            for fi45 := 0; fi45 < 11; fi45++ {
                for fi46 := 0; fi46 < 14; fi46++ {
                    switch (fi46 / (((fi46 * 127)) | 1)) {
                    case 42:
                        v44 = (v44 - (v44 ^ (v1 * (v1 + (v44 - v1)))))
                    }
                    var v47 float32 = float32(v2)
                    _ = fi46
                    _ = v47
                }
                v1 -= v1
                if (-662 > 1) {
                    fmt.Println(v1)
                }
                _ = fi45
            }
            _ = v44
        }
        for fi48 := 0; fi48 < 3; fi48++ {
            v49 := GenS8{F0: int16(v7), F1: 5}
            _ = fi48
            _ = v49
        }
        if ((v3 & v3) <= v3) {
            var v50 uint8 = v5
            fmt.Println(v50)
            _ = v50
        }
    }
    var v51 float64 = (((v2 / float64(((0.1 - 1e10) - 1e10))) / v2) + (v2 + (1.5 * float64(((1e10 * (-2.25 / 1.5)) - (3.14159 - 0.1))))))
    var v52 float32 = (1.5 + 0.1)
    for fi53 := 0; fi53 < 3; fi53++ {
        for fi54 := 0; fi54 < 9; fi54++ {
            if (978 != 1) {
                v55 := []int{-382, -9223372036854775808, fi54, (^(fi54))}
                _ = v55
            } else {
                if (false || ((!("a" >= ("foo bar" + ""))) || true)) {
                    v51 -= (v2 + v51)
                } else {
                    v3 -= (v3 - v3)
                    v56 := GenS3{F0: int8(v1), F1: "\n\t", F2: (^(v1))}
                    _ = v56
                }
                v7 += 2
                switch (("hello" + ((("hello" + ("foo bar" + (("hello" + "") + "a"))) + ("" + "\n\t")) + ((("a" + "a") + ("" + "\n\t")) + ("\n\t" + "foo bar")))) + ("a" + "")) {
                case "":
                    var v57 int64 = int64(fi54)
                    _ = v57
                default:
                    v58 := GenS6{F0: uint(v7)}
                    _ = v58
                }
                var v59 uint16 = 255
                _ = v59
            }
            var v60 uint8 = (v6 % ((255) | 1))
            _ = fi54
            _ = v60
        }
        _ = fi53
    }
    fmt.Println(v5, v2, v3)
    v61 := []int{-854}
    var v62 float32 = ((float32((v51 + (((v2 / ((v51 / 0.1) * v2)) - v2) * v2))) / 1.0) * float32(-2.25))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v61:", len(v61), v61)
    fmt.Println("v62:", v62)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v6
    _ = v61
    _ = v62
    _ = v7
    _ = v8
}
