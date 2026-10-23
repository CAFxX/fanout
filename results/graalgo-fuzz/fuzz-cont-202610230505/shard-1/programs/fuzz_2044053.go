package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 uint; F1 string; F2 int32 }
type GenS5 struct { F0 int16; F1 uint; F2 int64 }
func genFn0(p0 int8) uint {
    fmt.Println(p0)
    p0 += (p0 >> ((18446744073709551615) & 63))
    v4 := []int{9223372036854775807}
    v5 := GenS1{F0: 65535}
    for fi6 := 0; fi6 < 2; fi6++ {
        switch fi6 {
        case -604:
            for fi7 := 0; fi7 < 15; fi7++ {
                switch ("hello" + "\n\t") {
                case "\n\t":
                    p0 = p0
                case "":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                default:
                    switch ("\n\t" + "foo bar") {
                    case "foo bar":
                        var v8 bool = (fi6 == 1)
                        _ = v8
                    case "":
                        var v9 uint64 = 74
                        _ = v9
                    case "hello":
                        if true {
                            v10 := GenS1{F0: 96}
                            v11 := GenS2{F0: 42}
                            _ = v10
                            _ = v11
                        }
                    default:
                        v4 = append(v4, 4294967295)
                    }
                }
                p0 = p0
                _ = fi7
            }
        default:
            var v12 float32 = 0.0
            _ = v12
        }
        _ = fi6
    }
    _ = p0
    _ = v4
    _ = v5
    return 1
}

func main() {
    var v1 float64 = 0.1
    v2 := make(map[int]int)
    v3 := make(map[string]bool)
    var v13 uint = 254
    v13 = (v13 >> ((0) & 63))
    var v14 int64 = -1
    v15 := GenS3{F0: 571, F1: 42}
    switch "\n\t" {
    case "\n\t":
        for fi16 := 0; fi16 < 11; fi16++ {
            var v17 uint32 = 42
            v18 := v3[(("a" + "hello") + "hello")]
            if v18 {
                for fi19 := 0; fi19 < 13; fi19++ {
                    v1 += (v1 / float64(1.5))
                    var v20 uint64 = 2
                    _ = fi19
                    _ = v20
                }
                if v18 {
                    v17 = v17
                    v1 = (0.1 - v1)
                } else {
                    for fi21 := 0; fi21 < 18; fi21++ {
                        if (false || v18) {
                            v22 := []int{fi16, 2147483647}
                            _ = v22
                        } else {
                            var v23 uint64 = 198
                            fmt.Println(v18, v18)
                            v2[fi21] = fi16
                            _ = v23
                        }
                        var v24 uint = v13
                        var v25 uint64 = 4294967295
                        if (((!false) || false) && ((fi16 & (fi21 << ((v13) & 63))) <= (^(fi16)))) {
                            v26 := []int{fi16, int(v17), int(v17), (fi16 ^ fi16)}
                            if (-1 != 0) {
                                if v18 {
                                    var v27 uint = v24
                                    _ = v27
                                } else {
                                    if false {
                                        var v28 int8 = -65
                                        fmt.Println(v14)
                                        _ = v28
                                    } else {
                                        switch 2147483647 {
                                        case 2147483647:
                                            var v29 uint32 = (v17 ^ v17)
                                            _ = v29
                                        }
                                        fmt.Println(v18)
                                    }
                                    v30 := GenS1{F0: v17}
                                    _ = v30
                                }
                            } else {
                                var v31 int = (-(fi21))
                                _ = v31
                            }
                            _ = v26
                        }
                        _ = fi21
                        _ = v24
                        _ = v25
                    }
                    v32 := GenS4{F0: 590, F1: "hello", F2: 0}
                    v13 = genFn0(1)
                    _ = v32
                }
                v33 := GenS5{F0: 1000, F1: v13, F2: -300}
                _ = v33
            } else {
                switch fi16 {
                case 0:
                    v34 := []int{}
                    _ = v34
                case 691:
                    v17 -= (v17 / ((65535) | 1))
                case -164:
                    v13 = genFn0(12)
                }
                for fi35 := 0; fi35 < 6; fi35++ {
                    if ((("foo bar" + "\n\t") + "a") >= "") {
                        fmt.Println(fi16)
                        v13 += 401
                    } else {
                        var v36 int = (fi35 / (((fi35 % ((fi35) | 1))) | 1))
                        _ = v36
                    }
                    _ = len(v3)
                    var v37 float64 = v1
                    _ = fi35
                    _ = v37
                }
            }
            _ = fi16
            _ = v17
            _ = v18
        }
    case "":
        var v38 int8 = 123
        _ = v38
    case "foo bar":
        v1 += (-2.25 / 1e10)
    }
    v13 -= genFn0(127)
    switch 9223372036854775807 {
    case 4294967295:
        var v39 bool = true
        _ = v39
    }
    switch "a" {
    case "":
        _ = len(v2)
    case "a":
        if true {
            fmt.Println(v14, v1)
        }
    default:
        fmt.Println(v1, v1, v1)
    }
    delete(v3, "a")
    if true {
        if ((4294967295 != 572) && (!true)) {
            var v40 uint64 = 988
            _ = v40
        }
        var v41 uint64 = 2
        switch ("foo bar" + "a") {
        case "\n\t":
            var v42 uint8 = 1
            _ = v42
        }
        _ = v41
    } else {
        v2[int(v13)] = 1
        v1 += ((1e10 / v1) - 3.14159)
    }
    for fi43 := 0; fi43 < 16; fi43++ {
        delete(v2, (fi43 / ((-848) | 1)))
        _ = fi43
    }
    var v44 int64 = v14
    switch -9223372036854775808 {
    case 42:
        v13 = (v13 / ((v13) | 1))
    case 49:
        var v45 uint16 = 65535
        _ = v45
    case 0:
        v46 := GenS1{F0: 2147483647}
        _ = v46
    }
    v14 = (v44 + (v14 % ((v44) | 1)))
    if true {
        v14 = (-(v14))
        if false {
            if (0 >= 0) {
                v47 := []int{-2147483648, 1, 2147483647, 363}
                v14 -= (v44 ^ (v14 % ((v14) | 1)))
                _ = v47
            }
            for fi48 := 0; fi48 < 20; fi48++ {
                fmt.Println(v14)
                if (18446744073709551615 == 0) {
                    switch "a" {
                    case "foo bar":
                        var v49 uint32 = 65535
                        _ = v49
                    }
                } else {
                    v50 := GenS1{F0: 2147483647}
                    v51 := []int{128, 2147483647, 127}
                    if ((^(v44)) > v44) {
                        v13 -= (^(v13))
                    } else {
                        if false {
                            v3["\n\t"] = true
                        }
                        var v52 uint64 = 18446744073709551615
                        _ = v52
                    }
                    _ = v50
                    _ = v51
                }
                _ = fi48
            }
        } else {
            for fi53 := 0; fi53 < 5; fi53++ {
                v13 = (v13 ^ (v13 * v13))
                for fi54 := 0; fi54 < 18; fi54++ {
                    v55 := []int{}
                    v13 -= (^(v13))
                    switch ("\n\t" + ("hello" + "a")) {
                    case "":
                        fmt.Println(v14, v44)
                    case "foo bar":
                        fmt.Println(v44, fi54, fi54)
                    case "\n\t":
                        var v56 uint64 = 2
                        _ = v56
                    }
                    var v57 uint32 = 42
                    _ = fi54
                    _ = v55
                    _ = v57
                }
                for fi58 := 0; fi58 < 14; fi58++ {
                    fmt.Println(v44)
                    var v59 int = 128
                    var v60 uint8 = 255
                    _ = fi58
                    _ = v59
                    _ = v60
                }
                _ = fi53
            }
        }
    }
    var v61 int16 = -1
    if (2147483648 > 42) {
        v14 = v44
        fmt.Println(v1)
    } else {
        for fi62 := 0; fi62 < 16; fi62++ {
            v13 = genFn0(-128)
            _ = fi62
        }
        v14 -= v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v44:", v44)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v44
    _ = v61
}
