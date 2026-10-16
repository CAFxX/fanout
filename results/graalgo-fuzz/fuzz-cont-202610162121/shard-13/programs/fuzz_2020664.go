package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 uint16; F1 uint32; F2 uint8 }
type GenS4 struct { F0 int32; F1 int64 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint8; F1 string; F2 int16 }
func genFn0(p0 int8, p1 uint64, p2 int8) uint16 {
    p2 = (p2 % ((p0) | 1))
    p1 += 18446744073709551615
    _ = p0
    _ = p1
    _ = p2
    return 32767
}

func main() {
    var v1 uint64 = 42
    v2 := make(map[string]string)
    v3 := make(map[string]string)
    if ((941 < 126) && true) {
        var v4 uint = 1
        _ = v4
    } else {
        fmt.Println(v1)
        if (42 <= 525) {
            if (!(!("\n\t" == "\n\t"))) {
                v5 := GenS1{F0: 65535}
                switch ("hello" + "hello") {
                case "foo bar":
                    switch ("a" + (("\n\t" + "foo bar") + ("a" + ("foo bar" + "foo bar")))) {
                    case "hello":
                        v6 := GenS2{F0: -32768, F1: 9223372036854775807}
                        _ = v6
                    case "\n\t":
                        for fi7 := 0; fi7 < 11; fi7++ {
                            v8 := GenS2{F0: 438, F1: -758}
                            if false {
                                v9 := []int{}
                                var v10 uint = 65535
                                _ = v10
                                _ = v9
                            }
                            _ = fi7
                            _ = v8
                        }
                    default:
                        for fi11 := 0; fi11 < 9; fi11++ {
                            v1 += v1
                            v12 := GenS3{F0: 2, F1: 42, F2: 237}
                            var v13 float64 = 1e10
                            _ = fi11
                            _ = v12
                            _ = v13
                        }
                    }
                case "a":
                    var v14 uint16 = 65535
                    _ = v14
                case "\n\t":
                    var v15 int8 = -117
                    _ = v15
                }
                switch "a" {
                case "a":
                    v1 -= 4294967295
                }
                _ = v5
            } else {
                for fi16 := 0; fi16 < 4; fi16++ {
                    v17 := []int{}
                    _ = fi16
                    _ = v17
                }
            }
        }
        if (236 <= v1) {
            for fi18 := 0; fi18 < 10; fi18++ {
                switch ("hello" + "") {
                case "":
                    v19 := GenS3{F0: genFn0(35, 18446744073709551615, 1), F1: 2, F2: 42}
                    _ = v19
                case "a":
                    switch (fi18 / (((fi18 + (^(fi18)))) | 1)) {
                    case 32768:
                        var v20 uint64 = v1
                        _ = v20
                    case -2147483648:
                        switch 9223372036854775807 {
                        case 933:
                            v21 := v3["\n\t"]
                            _ = v21
                        case 42:
                            var v22 bool = true
                            _ = v22
                        default:
                            v1 = 0
                        }
                    case 32767:
                        var v23 int16 = 42
                        _ = v23
                    }
                case "foo bar":
                    for fi24 := 0; fi24 < 15; fi24++ {
                        switch "a" {
                        case "a":
                            if true {
                                v1 = 18446744073709551615
                                var v25 int8 = 0
                                _ = v25
                            }
                        case "":
                            fmt.Println(fi18)
                        }
                        if false {
                            switch fi18 {
                            case 96:
                                v1 += (^(v1))
                            case 127:
                                if (4294967295 <= 42) {
                                    v26 := []int{(fi18 / (((fi24 * fi24)) | 1)), fi24, fi24, (fi24 * (fi18 & (-(fi24))))}
                                    v27 := GenS4{F0: -1, F1: 42}
                                    v1 += v1
                                    _ = v26
                                    _ = v27
                                }
                            case 32767:
                                v1 = v1
                            }
                            v1 -= v1
                        } else {
                            var v28 int64 = 0
                            v29 := []int{(fi18 & (-(fi24)))}
                            _ = v28
                            _ = v29
                        }
                        _ = fi24
                    }
                default:
                    var v30 uint = 65535
                    _ = v30
                }
                v1 += (v1 % ((v1) | 1))
                _ = fi18
            }
        }
    }
    delete(v3, ("foo bar" + ("\n\t" + "\n\t")))
    var v31 bool = false
    switch 4294967295 {
    case -516:
        v32 := GenS5{F0: 193}
        _ = v32
    case -271:
        for fi33 := 0; fi33 < 17; fi33++ {
            if v31 {
                v34 := []int{-915, fi33}
                _ = len(v34)
                v35 := GenS6{F0: 128}
                _ = v34
                _ = v35
            } else {
                if v31 {
                    if ((false || (fi33 < 0)) || ((((3.14159 + (1e10 / (-2.25 * (0.0 + 1e10)))) - (0.1 * 0.0)) - 0.0) < 0.0)) {
                        for fi36 := 0; fi36 < 10; fi36++ {
                            var v37 uint8 = 40
                            _ = fi36
                            _ = v37
                        }
                    } else {
                        v38 := GenS3{F0: 255, F1: 2147483647, F2: 1}
                        for fi39 := 0; fi39 < 2; fi39++ {
                            var v40 int16 = 42
                            if false {
                                if v31 {
                                    switch (fi39 + fi39) {
                                    case 32767:
                                        var v41 bool = true
                                        _ = v41
                                    case 2147483647:
                                        v1 += v1
                                    }
                                }
                                if (!(v31 || v31)) {
                                    v40 -= (v40 & v40)
                                    fmt.Println(v40, v31, v1)
                                    var v42 uint16 = genFn0(-128, (v1 + v1), 42)
                                    _ = v42
                                } else {
                                    switch fi39 {
                                    case -9223372036854775808:
                                        v43 := GenS7{F0: 2}
                                        _ = v43
                                    case 1:
                                        var v44 int16 = v40
                                        _ = v44
                                    case 61:
                                        fmt.Println(v40, v40)
                                    default:
                                        v45 := GenS8{F0: 128, F1: ("\n\t" + (("\n\t" + ((("foo bar" + "foo bar") + "a") + "\n\t")) + "")), F2: -32768}
                                        _ = v45
                                    }
                                }
                                var v46 uint = 0
                                _ = v46
                            }
                            _ = fi39
                            _ = v40
                        }
                        _ = v38
                    }
                    var v47 uint64 = 70
                    v31 = (3.14159 != 1.5)
                    _ = v47
                }
            }
            switch ("hello" + "a") {
            case "hello":
                v1 += v1
            default:
                v31 = ((v31 || false) || false)
            }
            _ = fi33
        }
    default:
        switch ("" + "a") {
        case "hello":
            v31 = ((true && (18446744073709551615 != 1)) || (false && v31))
        case "foo bar":
            var v48 uint32 = 0
            _ = v48
        case "":
            v31 = ((32768 < -908) && v31)
        }
    }
    v1 -= (v1 << ((1) & 63))
    switch -802 {
    case 0:
        v49 := GenS5{F0: 1}
        _ = v49
    default:
        switch 10 {
        case 395:
            v50 := []int{32767, 127, int(v1), 32767}
            _ = v50
        default:
            _ = len(v2)
        }
    }
    for fi51 := 0; fi51 < 12; fi51++ {
        var v52 float32 = (-2.25 + (1e10 * ((((0.1 * (1e10 + (0.1 * ((3.14159 + 1e10) + 1e10)))) / 0.1) * 0.1) * 1e10)))
        _ = fi51
        _ = v52
    }
    v31 = false
    var v53 int32 = 2147483647
    v54 := GenS2{F0: -98, F1: -1}
    if (!(uint(v1) < 486)) {
        v31 = ("" >= "hello")
        var v55 float32 = 0.1
        for fi56 := 0; fi56 < 10; fi56++ {
            for fi57 := 0; fi57 < 5; fi57++ {
                v31 = (18446744073709551615 >= 157)
                v58 := []int{fi57}
                v1 -= v1
                v55 = v55
                _ = fi57
                _ = v58
            }
            _ = fi56
        }
        _ = v55
    }
    v1 -= (v1 % ((v1) | 1))
    for fi59 := 0; fi59 < 2; fi59++ {
        if ((^(v1)) != 1) {
            v53 = (-(v53))
            var v60 uint = 2
            v53 += (^(v53))
            _ = v60
        }
        _ = fi59
    }
    v31 = true
    v1 += (^(v1))
    switch 127 {
    case -2147483648:
        v61 := []int{-458, 128, 42}
        _ = v61
    case 42:
        v53 += (v53 ^ (v53 + (v53 * v53)))
    }
    if (516 != 32768) {
        if (914 > -768) {
            if v31 {
                var v62 int64 = 209
                var v63 uint64 = 0
                var v64 uint64 = v63
                _ = v62
                _ = v63
                _ = v64
            }
        }
        v53 = (v53 & (v53 * v53))
        v1 += v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v53
    _ = v54
}
