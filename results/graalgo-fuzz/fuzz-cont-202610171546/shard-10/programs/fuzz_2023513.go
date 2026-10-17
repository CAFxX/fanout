package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 string; F1 int32; F2 string }
type GenS5 struct { F0 string; F1 int8; F2 uint }
type GenS6 struct { F0 int64; F1 uint; F2 int32 }
type GenS7 struct { F0 uint64; F1 int16; F2 int8 }
type GenS8 struct { F0 int8; F1 uint8 }
func genFn0(p0 bool, p1 int) uint32 {
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 uint64 = 0
        var v9 int16 = int16(v8)
        v10 := []int{(p1 << ((4294967295) & 63)), p1}
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        v12 := GenS1{F0: -1}
        switch "" {
        case "":
            var v13 uint32 = 0
            _ = v13
        default:
            p1 -= (p1 + (^(fi11)))
        }
        var v14 uint16 = 0
        _ = fi11
        _ = v12
        _ = v14
    }
    v15 := []int{(p1 + p1)}
    _ = p0
    _ = p1
    _ = v15
    return 363
}
func genFn1(p0 int64, p1 int) int8 {
    p1 = (p1 >> ((511) & 63))
    for fi16 := 0; fi16 < 17; fi16++ {
        switch (p1 / ((fi16) | 1)) {
        case 230:
            if ((fi16 * p1) > fi16) {
                p0 -= (p0 | (p0 + p0))
            }
        case 127:
            switch "hello" {
            case "":
                v17 := []int{0, fi16, fi16}
                _ = v17
            }
        }
        if false {
            v18 := GenS2{F0: 1}
            _ = v18
        } else {
            p1 = 128
            p0 = (p0 + p0)
            v19 := []int{p1, int(p0), fi16}
            _ = v19
        }
        p1 = (p1 | (p1 * p1))
        _ = fi16
    }
    switch "hello" {
    case "hello":
        for fi20 := 0; fi20 < 7; fi20++ {
            p1 = (p1 - p1)
            _ = fi20
        }
    case "foo bar":
        if (((((-2.25 <= 3.14159) || true) && (false || (int64(p1) <= (p0 >> ((822) & 63))))) || (p1 <= (p1 >> ((255) & 63)))) && true) {
            switch (("a" + ("" + "a")) + "foo bar") {
            case "\n\t":
                for fi21 := 0; fi21 < 13; fi21++ {
                    for fi22 := 0; fi22 < 9; fi22++ {
                        v23 := []int{fi22}
                        _ = fi22
                        _ = v23
                    }
                    fmt.Println(fi21, p0, p1)
                    var v24 int16 = -32768
                    _ = fi21
                    _ = v24
                }
            case "foo bar":
                p1 = 4294967295
            case "hello":
                v25 := []int{(p1 - (p1 / ((p1) | 1))), (p1 / ((p1) | 1))}
                _ = v25
            }
            fmt.Println(p0)
            for fi26 := 0; fi26 < 6; fi26++ {
                fmt.Println(fi26)
                for fi27 := 0; fi27 < 18; fi27++ {
                    switch 399 {
                    case 2147483647:
                        v28 := GenS1{F0: p0}
                        _ = v28
                    }
                    _ = fi27
                }
                var v29 uint64 = 18446744073709551615
                _ = fi26
                _ = v29
            }
        } else {
            if (!true) {
                switch "hello" {
                case "\n\t":
                    fmt.Println(p1, p0)
                }
                switch p1 {
                case -9223372036854775808:
                    for fi30 := 0; fi30 < 19; fi30++ {
                        for fi31 := 0; fi31 < 1; fi31++ {
                            v32 := []int{(^(fi31)), (p1 & fi30)}
                            fmt.Println(fi31, fi31, fi31)
                            if false {
                                var v33 uint32 = 42
                                v32 = append(v32, 128)
                                fmt.Println(p1, v33, p1)
                                p0 += (^(p0))
                                var v34 float64 = (1e10 * -2.25)
                                _ = v33
                                _ = v34
                            }
                            _ = fi31
                            _ = v32
                        }
                        v35 := GenS3{F0: true, F1: 18446744073709551615}
                        _ = fi30
                        _ = v35
                    }
                case -2147483648:
                    if true {
                        if ((true && (258 >= 65535)) || (true && (!(0 == p0)))) {
                            v36 := GenS4{F0: "a", F1: -909, F2: "\n\t"}
                            _ = v36
                        }
                        if ((0 >= p0) || (!true)) {
                            fmt.Println(p0, p1)
                            var v37 float64 = 1e10
                            fmt.Println(p1, p0)
                            _ = v37
                        }
                        var v38 uint8 = 214
                        _ = v38
                    }
                case 1:
                    for fi39 := 0; fi39 < 5; fi39++ {
                        if true {
                            p0 -= (p0 - (p0 % ((p0) | 1)))
                        } else {
                            var v40 bool = true
                            for fi41 := 0; fi41 < 12; fi41++ {
                                switch "a" {
                                case "":
                                    if v40 {
                                        if false {
                                            if (503 > 18446744073709551615) {
                                                var v42 int32 = -2147483648
                                                v43 := []int{fi39, p1, fi41, fi39}
                                                v42 = (v42 ^ (v42 << ((65535) & 63)))
                                                _ = v42
                                                _ = v43
                                            }
                                            v44 := GenS1{F0: p0}
                                            var v45 uint = 255
                                            _ = v44
                                            _ = v45
                                        } else {
                                            var v46 uint32 = 4294967295
                                            _ = v46
                                        }
                                        fmt.Println(p1, p0)
                                    }
                                case "foo bar":
                                    var v47 uint64 = 4294967295
                                    _ = v47
                                default:
                                    fmt.Println(fi41)
                                }
                                _ = fi41
                            }
                            var v48 uint8 = 127
                            _ = v40
                            _ = v48
                        }
                        fmt.Println(fi39, p1, p1)
                        fmt.Println(fi39, fi39, fi39)
                        _ = fi39
                    }
                }
            }
            switch p1 {
            case -2147483648:
                var v49 uint64 = 807
                _ = v49
            case -1:
                for fi50 := 0; fi50 < 10; fi50++ {
                    v51 := GenS5{F0: "foo bar", F1: 112, F2: 65535}
                    _ = fi50
                    _ = v51
                }
            }
        }
    case "a":
        var v52 uint = 42
        _ = v52
    }
    switch "a" {
    case "hello":
        p1 -= -801
    }
    _ = p0
    _ = p1
    return -76
}

func main() {
    var v1 int = 32767
    v2 := []int{v1, (v1 % (((v1 & 499)) | 1))}
    var v3 uint64 = 151
    v4 := make(map[string]bool)
    var v5 uint8 = 12
    var v6 string = "foo bar"
    for fi53 := 0; fi53 < 20; fi53++ {
        if false {
            for fi54 := 0; fi54 < 4; fi54++ {
                v6 = v6
                v6 = v6
                _ = fi54
            }
            switch "a" {
            case "a":
                for fi55 := 0; fi55 < 6; fi55++ {
                    v5 = (v5 << ((0) & 63))
                    v5 = 255
                    _ = fi55
                }
            case "":
                v4["a"] = true
            default:
                for fi56 := 0; fi56 < 1; fi56++ {
                    v57 := GenS6{F0: 1, F1: 4294967295, F2: 42}
                    v58 := GenS6{F0: 102, F1: 2, F2: 2147483647}
                    _ = fi56
                    _ = v57
                    _ = v58
                }
            }
            if (v5 > 128) {
                for fi59 := 0; fi59 < 11; fi59++ {
                    var v60 int16 = -32768
                    var v61 int32 = 42
                    switch "hello" {
                    case "foo bar":
                        v6 = (v6 + "hello")
                    case "\n\t":
                        var v62 uint = 2
                        _ = v62
                    }
                    var v63 int8 = -8
                    var v64 float32 = 0.1
                    var v65 uint16 = 116
                    switch -1 {
                    case 42:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    v63 = (v63 + v63)
                    _ = fi59
                    _ = v60
                    _ = v61
                    _ = v63
                    _ = v64
                    _ = v65
                }
            } else {
                fmt.Println(v6, fi53)
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v66 := GenS7{F0: (v3 * v3), F1: 1000, F2: 42}
                _ = v66
            }
        }
        var v67 bool = (((v5 * (^(v5))) < (v5 | v5)) || false)
        v4[(("\n\t" + (((((v6 + ("foo bar" + (v6 + (v6 + v6)))) + "\n\t") + v6) + "a") + ((("\n\t" + (v6 + ((v6 + v6) + (v6 + v6)))) + v6) + ("foo bar" + (v6 + ""))))) + (v6 + v6))] = false
        _ = fi53
        _ = v67
    }
    var v68 int16 = -182
    v4["hello"] = (false || (v6 < v6))
    fmt.Println(v5, v6, v6)
    fmt.Println(v1)
    v6 = (v6 + v6)
    var v69 uint64 = 1
    var v70 uint8 = 127
    switch "foo bar" {
    case "foo bar":
        var v71 int16 = 42
        _ = v71
    case "a":
        v72 := GenS8{F0: 127, F1: v70}
        _ = v72
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v68
    _ = v69
    _ = v70
}
