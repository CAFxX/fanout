package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int16; F1 uint64; F2 int8 }
func genFn0(p0 int32, p1 uint) int8 {
    fmt.Println(p0, p1)
    if (!true) {
        fmt.Println(p1, p0)
    }
    p1 -= p1
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 float32 = 0.0
    var v2 int = 29
    var v3 bool = true
    for fi4 := 0; fi4 < 19; fi4++ {
        var v5 uint64 = 42
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 6; fi6++ {
        v2 += fi6
        fmt.Println(v3, v2)
        _ = fi6
    }
    v7 := GenS1{F0: genFn0(0, 2), F1: 2, F2: (-(v2))}
    fmt.Println(v3)
    if true {
        fmt.Println(v1, v3)
        if (!true) {
            if (1 > 4294967295) {
                switch "a" {
                case "a":
                    if false {
                        if v3 {
                            for fi8 := 0; fi8 < 15; fi8++ {
                                v9 := []int{(v2 % (((fi8 & -358)) | 1)), (v2 | v2), v2, 725}
                                _ = fi8
                                _ = v9
                            }
                            fmt.Println(v1, v1)
                        }
                        fmt.Println(v3, v3)
                    }
                case "foo bar":
                    v2 += v2
                default:
                    if v3 {
                        var v10 int8 = 100
                        var v11 uint8 = 1
                        _ = v10
                        _ = v11
                    } else {
                        if false {
                            v12 := GenS1{F0: 1, F1: 0, F2: v2}
                            _ = v12
                        } else {
                            v13 := GenS2{F0: 886}
                            v1 = -2.25
                            var v14 uint64 = 1
                            _ = v13
                            _ = v14
                        }
                    }
                }
                v15 := GenS2{F0: 255}
                if false {
                    var v16 uint8 = 2
                    if false {
                        var v17 uint64 = 1
                        if (!("foo bar" == "hello")) {
                            v18 := []int{(v2 / (((v2 % ((v2) | 1))) | 1)), v2, v2}
                            if v3 {
                                var v19 uint = 1
                                v2 -= v2
                                _ = v19
                            } else {
                                var v20 string = "foo bar"
                                _ = v20
                            }
                            if len(v18) >= 2 {
                                v18 = v18[1:len(v18)]
                            }
                            _ = v18
                        } else {
                            v21 := GenS2{F0: 0}
                            if (912 > 42) {
                                var v22 int32 = 42
                                switch ("foo bar" + "") {
                                case "a":
                                    var v23 float32 = 0.1
                                    _ = v23
                                case "\n\t":
                                    v24 := []int{(v2 / ((9223372036854775807) | 1)), 292, 128, -2147483648}
                                    _ = v24
                                case "hello":
                                    v22 -= -2147483648
                                }
                                _ = v22
                            }
                            _ = v21
                        }
                        v25 := []int{v2, 128}
                        _ = v17
                        _ = v25
                    } else {
                        var v26 int = (v2 * v2)
                        _ = v26
                    }
                    _ = v16
                } else {
                    v1 -= v1
                }
                _ = v15
            }
            v27 := []int{(v2 | (v2 & (^(v2)))), (^(v2))}
            _ = len(v27)
            _ = v27
        }
    } else {
        v3 = true
    }
    var v28 int8 = 110
    for fi29 := 0; fi29 < 9; fi29++ {
        if (9 > 42) {
            var v30 float32 = 1e10
            switch "hello" {
            case "":
                var v31 int8 = v28
                _ = v31
            case "\n\t":
                var v32 int = (v2 ^ 128)
                _ = v32
            case "a":
                fmt.Println(v30, v30, v30)
            }
            _ = v30
        } else {
            if ((v3 && v3) || ((0 >= 1) || (v3 && v3))) {
                switch "foo bar" {
                case "hello":
                    var v33 float64 = (((1e10 + 3.14159) + 1.5) - float64(v1))
                    _ = v33
                }
            } else {
                switch fi29 {
                case 127:
                    fmt.Println(v28, v3, v1)
                case 9223372036854775807:
                    v34 := GenS1{F0: (^(v28)), F1: 0, F2: (fi29 & (^(fi29)))}
                    _ = v34
                case -1:
                    for fi35 := 0; fi35 < 14; fi35++ {
                        fmt.Println(fi29, fi29, v1)
                        fmt.Println(v3, v28, v2)
                        _ = fi35
                    }
                }
            }
            v28 -= (v28 - v28)
        }
        v28 += -128
        _ = fi29
    }
    v36 := GenS3{F0: (v28 - 42), F1: "a"}
    switch (^(v2)) {
    case 42:
        fmt.Println(v3)
    }
    var v37 int = v2
    if (!v3) {
        if false {
            var v38 float64 = 3.14159
            switch "hello" {
            case "a":
                var v39 float32 = float32(v38)
                _ = v39
            }
            _ = v38
        }
    }
    for fi40 := 0; fi40 < 15; fi40++ {
        v41 := GenS4{F0: v28}
        var v42 float32 = v1
        v42 += v1
        _ = fi40
        _ = v41
        _ = v42
    }
    for fi43 := 0; fi43 < 18; fi43++ {
        v44 := []int{v2, (v37 + 282), fi43}
        v45 := GenS5{F0: -632, F1: 1, F2: genFn0(0, 941)}
        var v46 int = fi43
        _ = fi43
        _ = v44
        _ = v45
        _ = v46
    }
    switch -83 {
    case 9223372036854775807:
        switch "\n\t" {
        case "foo bar":
            fmt.Println(v1)
        default:
            fmt.Println(v1, v1, v3)
        }
    case 42:
        if (!(genFn0(1, 18446744073709551615) == 1)) {
            if v3 {
                v47 := GenS2{F0: 762}
                v48 := []int{(v2 / (((v37 % (((-(v2))) | 1))) | 1))}
                for fi49 := 0; fi49 < 6; fi49++ {
                    v48 = append(v48, (v2 >> ((0) & 63)))
                    _ = fi49
                }
                _ = v47
                _ = v48
            } else {
                switch "\n\t" {
                case "\n\t":
                    v1 -= (v1 - ((v1 * 0.1) / v1))
                case "a":
                    fmt.Println(v3)
                case "hello":
                    v50 := []int{-1, v2, -87, int(v28)}
                    _ = v50
                }
                var v51 uint16 = 42
                if v3 {
                    if false {
                        v52 := []int{}
                        var v53 uint32 = 1
                        var v54 float64 = 1.5
                        _ = v52
                        _ = v53
                        _ = v54
                    }
                }
                _ = v51
            }
            var v55 uint16 = 65535
            _ = v55
        }
    case 0:
        fmt.Println(v28, v3, v3)
    default:
        if v3 {
            v28 = (v28 ^ (v28 + v28))
            v37 -= v2
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v36
    _ = v37
    _ = v7
}
