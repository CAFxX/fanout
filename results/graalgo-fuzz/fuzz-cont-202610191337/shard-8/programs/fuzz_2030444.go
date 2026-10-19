package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32; F1 int16; F2 string }
func genFn0(p0 uint8, p1 uint) bool {
    for fi5 := 0; fi5 < 2; fi5++ {
        v6 := []int{fi5, (fi5 ^ fi5), (^(fi5))}
        p1 = (^(p1))
        p1 = 655
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 8; fi7++ {
        for fi8 := 0; fi8 < 18; fi8++ {
            v9 := GenS1{F0: 1, F1: (2 == 157), F2: p0}
            switch fi7 {
            case 127:
                var v10 int8 = 42
                _ = v10
            default:
                v11 := []int{(fi8 | 127)}
                _ = v11
            }
            p0 -= (^(p0))
            fmt.Println(p0, p0, fi7)
            _ = fi8
            _ = v9
        }
        switch (("a" + ("hello" + "foo bar")) + "") {
        case "hello":
            var v12 float64 = -2.25
            _ = v12
        case "\n\t":
            if (!("\n\t" <= "hello")) {
                var v13 uint8 = p0
                for fi14 := 0; fi14 < 10; fi14++ {
                    var v15 float64 = 0.0
                    _ = fi14
                    _ = v15
                }
                _ = v13
            }
        case "foo bar":
            for fi16 := 0; fi16 < 6; fi16++ {
                fmt.Println(p0, p0, fi7)
                v17 := []int{fi16, (^(fi7)), 997, -1}
                _ = fi16
                _ = v17
            }
        default:
            switch "hello" {
            case "a":
                if false {
                    v18 := []int{(fi7 >> ((42) & 63)), fi7, -770, fi7}
                    var v19 int16 = -259
                    p1 += p1
                    _ = v18
                    _ = v19
                } else {
                    p0 -= p0
                    v20 := GenS1{F0: 309, F1: (false && true), F2: p0}
                    _ = v20
                }
            case "foo bar":
                if false {
                    if false {
                        var v21 int16 = 32767
                        var v22 float64 = ((0.0 * (0.1 + -2.25)) - (-2.25 - 0.1))
                        _ = v21
                        _ = v22
                    } else {
                        for fi23 := 0; fi23 < 6; fi23++ {
                            switch ("" + "") {
                            case "":
                                v24 := []int{816}
                                _ = v24
                            case "hello":
                                p0 += (p0 / ((p0) | 1))
                            case "foo bar":
                                v25 := GenS1{F0: 65535, F1: false, F2: (p0 * (p0 << (((p1 % (((p1 / (((^(p1))) | 1))) | 1))) & 63)))}
                                _ = v25
                            default:
                                if true {
                                    var v26 int64 = -1
                                    _ = v26
                                }
                            }
                            var v27 int16 = -410
                            if true {
                                v28 := GenS2{F0: fi7}
                                p0 -= 7
                                _ = v28
                            } else {
                                v27 += -761
                                p0 = 125
                                v29 := []int{}
                                _ = v29
                            }
                            _ = fi23
                            _ = v27
                        }
                        v30 := []int{(fi7 / (((fi7 << (((p1 << (((p1 + (p1 + (p1 - 941)))) & 63))) & 63))) | 1))}
                        switch "foo bar" {
                        case "hello":
                            p0 += (p0 + (p0 & 137))
                        default:
                            p1 -= uint(p0)
                        }
                        _ = v30
                    }
                    if (2147483647 < -2147483648) {
                        fmt.Println(p1, p0, p1)
                        var v31 uint8 = 255
                        v31 = (p0 & (^(p0)))
                        _ = v31
                    }
                    switch "foo bar" {
                    case "foo bar":
                        for fi32 := 0; fi32 < 14; fi32++ {
                            var v33 float32 = (((3.14159 + (float32(0.0) / ((((((0.1 * 1.5) * 0.0) - 1e10) / (0.0 - 1e10)) + 3.14159) / (1e10 * (1e10 / -2.25))))) * 1e10) - (0.0 * 0.0))
                            v34 := GenS1{F0: 2147483647, F1: true, F2: (p0 + 2)}
                            _ = fi32
                            _ = v33
                            _ = v34
                        }
                    default:
                        switch 9223372036854775807 {
                        case -261:
                            if true {
                                var v35 int32 = 2147483647
                                for fi36 := 0; fi36 < 13; fi36++ {
                                    v35 = v35
                                    var v37 float64 = (3.14159 - 3.14159)
                                    _ = fi36
                                    _ = v37
                                }
                                fmt.Println(fi7, p1)
                                _ = v35
                            }
                        default:
                            var v38 uint16 = 340
                            _ = v38
                        }
                    }
                }
            case "":
                v39 := GenS2{F0: 4294967295}
                _ = v39
            }
        }
        v40 := GenS1{F0: 720, F1: ((^(p0)) < p0), F2: p0}
        _ = fi7
        _ = v40
    }
    _ = p0
    _ = p1
    return true
}
func genFn1() string {
    var v41 int16 = -32768
    v42 := []int{884, -313}
    var v43 int16 = v41
    v41 += v43
    _ = v41
    _ = v42
    _ = v43
    return ""
}

func main() {
    var v1 string = "a"
    v2 := []string{}
    var v3 int8 = -128
    v4 := make(map[int]string)
    v44 := GenS3{F0: 1, F1: 42, F2: ""}
    v45 := GenS1{F0: 4294967295, F1: (1 <= 4294967295), F2: 128}
    for fi46 := 0; fi46 < 18; fi46++ {
        if false {
            var v47 int16 = 42
            if len(v2) > 0 {
                v2[0] = (("\n\t" + (v1 + "\n\t")) + ((("" + "a") + (v1 + v1)) + v1))
            }
            var v48 float64 = 1.5
            _ = v47
            _ = v48
        }
        fmt.Println(v1, v1, fi46)
        _ = fi46
    }
    var v49 float32 = -2.25
    var v50 uint16 = 255
    v50 += (v50 | v50)
    if true {
        v3 -= 127
        v4[-2147483648] = v1
        v1 = (v1 + v1)
    } else {
        switch "hello" {
        case "\n\t":
            for fi51 := 0; fi51 < 8; fi51++ {
                v50 += v50
                v1 = ""
                for fi52 := 0; fi52 < 17; fi52++ {
                    v1 = (v1 + v1)
                    var v53 int = -9223372036854775808
                    switch v1 {
                    case "foo bar":
                        switch 32768 {
                        case -905:
                            if len(v2) > 0 {
                                v2[0] = ((v1 + "hello") + (v1 + ("hello" + (("hello" + "foo bar") + ""))))
                            }
                        default:
                            var v54 uint32 = 65535
                            _ = v54
                        }
                    case "":
                        v1 = "a"
                    case "a":
                        _ = len(v2)
                    default:
                        var v55 float64 = -2.25
                        _ = v55
                    }
                    v56 := GenS1{F0: 2147483647, F1: (32767 <= -1), F2: 0}
                    _ = fi52
                    _ = v53
                    _ = v56
                }
                _ = fi51
            }
        case "a":
            v57 := GenS3{F0: 493, F1: -390, F2: ("\n\t" + "\n\t")}
            _ = v57
        }
        var v58 int16 = -990
        _ = v58
    }
    var v59 uint = 2
    var v60 float64 = float64(1.5)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v49
    _ = v50
    _ = v59
    _ = v60
}
