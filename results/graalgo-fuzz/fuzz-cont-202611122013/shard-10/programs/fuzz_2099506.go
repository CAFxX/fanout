package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 int }
type GenS2 struct { F0 int; F1 uint16; F2 uint }
type GenS3 struct { F0 int8; F1 int32; F2 uint16 }
type GenS4 struct { F0 int64; F1 uint32 }
type GenS5 struct { F0 int8; F1 uint32 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 string; F1 uint8; F2 uint8 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int64; F1 uint64 }
func genFn0(p0 int16, p1 uint) int {
    for fi4 := 0; fi4 < 4; fi4++ {
        v5 := GenS1{F0: 42, F1: 0, F2: fi4}
        _ = fi4
        _ = v5
    }
    p1 += p1
    _ = p0
    _ = p1
    return 1
}
func genFn1(p0 int8) int {
    v6 := GenS2{F0: 32768, F1: 807, F2: 1}
    v7 := []int{-830, genFn0(1000, 725), 32768, 42}
    var v8 bool = (true && true)
    for fi9 := 0; fi9 < 8; fi9++ {
        fmt.Println(v8)
        for fi10 := 0; fi10 < 14; fi10++ {
            if (413 == 2147483647) {
                v8 = (!(false || (!v8)))
                var v11 float32 = 0.1
                if false {
                    fmt.Println(fi10, v11)
                }
                _ = v11
            } else {
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            }
            if false {
                var v12 uint = 255
                p0 = (p0 ^ p0)
                _ = v12
            } else {
                var v13 uint = 42
                if true {
                    var v14 uint32 = 861
                    var v15 bool = true
                    v15 = true
                    _ = v14
                    _ = v15
                } else {
                    v7 = append(v7, (^(fi10)))
                    v7 = append(v7, fi10)
                    v16 := GenS1{F0: 0, F1: 0, F2: genFn0(0, 4294967295)}
                    _ = v16
                }
                if (403 == -2147483648) {
                    if true {
                        v13 -= v13
                        v17 := GenS1{F0: 9223372036854775807, F1: 729, F2: -1}
                        if v8 {
                            fmt.Println(v13)
                        } else {
                            v18 := GenS2{F0: (-(fi9)), F1: 255, F2: 1}
                            _ = v18
                        }
                        _ = v17
                    }
                } else {
                    switch ("a" + "foo bar") {
                    case "hello":
                        if (993 <= (v13 % ((130) | 1))) {
                            v19 := GenS3{F0: 1, F1: 42, F2: 65535}
                            p0 -= p0
                            if false {
                                var v20 float32 = (1e10 + 0.0)
                                v7 = append(v7, genFn0(59, (v13 + 255)))
                                _ = v20
                            } else {
                                var v21 int16 = -32768
                                var v22 int32 = 42
                                _ = v21
                                _ = v22
                            }
                            v8 = v8
                            switch genFn0(-32768, v13) {
                            case 9223372036854775807:
                                v13 += 2
                            case 0:
                                v23 := GenS3{F0: 88, F1: 488, F2: 65535}
                                _ = v23
                            }
                            _ = v19
                        }
                    }
                    v8 = (58 >= p0)
                    v24 := GenS4{F0: 665, F1: 1}
                    v25 := GenS4{F0: 0, F1: 556}
                    v26 := GenS5{F0: -1, F1: 4294967295}
                    fmt.Println(v13)
                    switch (fi9 % ((4294967295) | 1)) {
                    case 1:
                        var v27 bool = false
                        _ = v27
                    case 42:
                        var v28 uint = 1
                        _ = v28
                    }
                    _ = v24
                    _ = v25
                    _ = v26
                }
                _ = v13
            }
            fmt.Println(fi9, fi10, fi9)
            _ = fi10
        }
        switch -9223372036854775808 {
        case 0:
            var v29 int = (^(fi9))
            _ = v29
        default:
            if true {
                if v8 {
                    var v30 uint8 = 128
                    _ = v30
                }
                v8 = true
            }
        }
        _ = fi9
    }
    for fi31 := 0; fi31 < 15; fi31++ {
        v32 := GenS3{F0: 17, F1: 2147483647, F2: 2}
        p0 = -75
        if (0.0 >= (1.5 / float64((float32(((-2.25 + 1.5) / -2.25)) * 0.1)))) {
            switch "" {
            case "hello":
                if len(v7) > 0 {
                    v7[0] = (fi31 ^ fi31)
                }
            case "":
                var v33 bool = (true || true)
                _ = v33
            case "\n\t":
                fmt.Println(fi31, fi31)
            default:
                v34 := GenS6{F0: (fi31 | -281)}
                _ = v34
            }
            v7 = append(v7, (-(fi31)))
        } else {
            for fi35 := 0; fi35 < 12; fi35++ {
                v8 = v8
                _ = fi35
            }
            for fi36 := 0; fi36 < 10; fi36++ {
                v37 := GenS7{F0: (("\n\t" + ((("a" + "foo bar") + "hello") + ((("\n\t" + "") + "foo bar") + "\n\t"))) + "foo bar"), F1: 2, F2: 1}
                switch (fi36 / (((fi31 * genFn0(-1, 1))) | 1)) {
                case 95:
                    v38 := GenS8{F0: 850}
                    _ = v38
                case 128:
                    v8 = v8
                case 1:
                    switch "a" {
                    case "\n\t":
                        v39 := GenS7{F0: "\n\t", F1: 128, F2: 127}
                        _ = v39
                    case "a":
                        p0 += 0
                    case "hello":
                        var v40 bool = v8
                        _ = v40
                    }
                }
                _ = fi36
                _ = v37
            }
            fmt.Println(v8, fi31, p0)
        }
        _ = fi31
        _ = v32
    }
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    v1 := []string{}
    var v2 float64 = 0.1
    v3 := make(map[string]int)
    v41 := GenS8{F0: 4294967295}
    v2 -= (v2 + v2)
    v3["a"] = 4294967295
    var v42 uint16 = 0
    if false {
        v43 := GenS9{F0: 9223372036854775807, F1: 18}
        if (v2 != v2) {
            v2 -= ((1.5 + float64(1e10)) * ((((((v2 * (v2 - v2)) - ((float64(1e10) - (0.1 * (1e10 + v2))) * float64(0.1))) - -2.25) / v2) + ((v2 + v2) / (v2 - v2))) + (float64(0.1) + (v2 * (((3.14159 + (v2 + v2)) * float64(((1.5 / ((-2.25 * (-2.25 + ((-2.25 / -2.25) + 0.0))) / 1e10)) * (1e10 + (1.5 / 1.0))))) * (1.5 / (((v2 + ((v2 - 1.5) / float64(0.1))) - ((v2 * 1.5) - v2)) / -2.25)))))))
            var v44 int32 = 595
            _ = v44
        } else {
            v42 = v42
        }
        if (!true) {
            switch "a" {
            case "":
                fmt.Println(v42, v2)
            default:
                v45 := v3["\n\t"]
                _ = v45
            }
            var v46 uint = 0
            v47 := GenS2{F0: 674, F1: v42, F2: v46}
            _ = v46
            _ = v47
        }
        _ = v43
    } else {
        fmt.Println(v42, v2)
    }
    v48 := GenS5{F0: 1, F1: 1}
    fmt.Println(v2, v2)
    v2 = (((((1.5 / float64(1e10)) * v2) + v2) / v2) + (v2 / (0.1 / -2.25)))
    var v49 int8 = -1
    if false {
        var v50 uint16 = 384
        _ = v50
    }
    v2 += (v2 * 0.1)
    v49 = (-(v49))
    v42 -= v42
    _ = len(v1)
    if (false || true) {
        switch "\n\t" {
        case "\n\t":
            v49 -= (v49 / ((-115) | 1))
        case "a":
            fmt.Println(v49)
        default:
            v2 = 1e10
        }
        for fi51 := 0; fi51 < 19; fi51++ {
            fmt.Println(v49, v49)
            if (!("a" != "a")) {
                var v52 int16 = 32767
                _ = v52
            }
            fmt.Println(fi51)
            _ = fi51
        }
    } else {
        v49 += (-(v49))
        var v53 float64 = -2.25
        _ = v53
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v42:", v42)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v41
    _ = v42
    _ = v48
    _ = v49
}
