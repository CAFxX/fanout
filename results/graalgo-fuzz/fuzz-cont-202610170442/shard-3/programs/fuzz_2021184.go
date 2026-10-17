package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 int }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int64; F1 bool; F2 uint }
type GenS7 struct { F0 int8; F1 uint64; F2 uint16 }
func genFn0() uint16 {
    v7 := GenS1{F0: 549}
    var v8 int64 = -9223372036854775808
    v8 += (v8 | v8)
    _ = v7
    _ = v8
    return 255
}
func genFn1(p0 int64, p1 uint32) uint64 {
    v9 := []int{-340, 2147483647}
    for fi10 := 0; fi10 < 14; fi10++ {
        p0 -= p0
        if (1 < 18446744073709551615) {
            p0 -= p0
            v11 := GenS2{F0: 42, F1: (false || false)}
            _ = v11
        } else {
            p0 -= p0
            var v12 int64 = p0
            var v13 uint16 = 1
            _ = v12
            _ = v13
        }
        _ = fi10
    }
    if true {
        var v14 int8 = -1
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v9
    return 42
}

func main() {
    var v1 string = "foo bar"
    v2 := []int{407, -601, -1, 4294967295}
    var v3 bool = ((true || false) || false)
    var v4 int16 = 1000
    var v5 float64 = 1e10
    v6 := make(map[int]uint32)
    for fi15 := 0; fi15 < 2; fi15++ {
        v1 = ""
        _ = fi15
    }
    v2 = append(v2, -2147483648)
    v3 = v3
    for fi16 := 0; fi16 < 3; fi16++ {
        var v17 float32 = 0.1
        _ = fi16
        _ = v17
    }
    fmt.Println(v1)
    v1 = "a"
    for fi18 := 0; fi18 < 4; fi18++ {
        if ((0.0 + (3.14159 / 1.5)) <= 1e10) {
            _ = len(v6)
        } else {
            v5 = (v5 * v5)
        }
        _ = fi18
    }
    for fi19 := 0; fi19 < 16; fi19++ {
        fmt.Println(v4, v1)
        v1 = "a"
        v20 := GenS2{F0: (v4 - v4), F1: v3}
        _ = fi19
        _ = v20
    }
    var v21 bool = (false || (!v3))
    for fi22 := 0; fi22 < 18; fi22++ {
        for fi23 := 0; fi23 < 16; fi23++ {
            var v24 int16 = 81
            var v25 string = (((("a" + ((("" + (("a" + v1) + (v1 + "hello"))) + "foo bar") + (v1 + v1))) + v1) + "foo bar") + (v1 + ("foo bar" + v1)))
            switch fi23 {
            case 4294967295:
                v26 := GenS2{F0: v4, F1: true}
                _ = v26
            case 127:
                v5 = 0.1
            default:
                if v3 {
                    if v3 {
                        v27 := GenS3{F0: 2147483647, F1: ("foo bar" + "foo bar")}
                        v1 = v1
                        switch v25 {
                        case "hello":
                            if (genFn1(-967, 0) > genFn1(745, 65535)) {
                                if v3 {
                                    fmt.Println(v5, fi23)
                                    v25 = "\n\t"
                                }
                                var v28 uint64 = 2
                                var v29 int8 = 108
                                v30 := GenS2{F0: -1, F1: v21}
                                v31 := GenS4{F0: (fi23 * -1)}
                                _ = v28
                                _ = v29
                                _ = v30
                                _ = v31
                            } else {
                                if (!true) {
                                    v32 := v6[fi22]
                                    _ = v32
                                }
                                v2 = append(v2, 0)
                                var v33 bool = v3
                                _ = v33
                            }
                        case "":
                            if true {
                                var v34 uint16 = 774
                                _ = v34
                            } else {
                                v35 := GenS1{F0: 32767}
                                v6[(fi22 % (((-(fi22))) | 1))] = 2147483648
                                v36 := GenS5{F0: v3}
                                v37 := GenS3{F0: 1, F1: v25}
                                v3 = false
                                var v38 uint8 = 153
                                switch ("" + "a") {
                                case "a":
                                    fmt.Println(v24, v3, fi23)
                                    var v39 float64 = ((v5 + v5) / 1.0)
                                    if true {
                                        var v40 uint = 65535
                                        switch v25 {
                                        case "foo bar":
                                            switch "foo bar" {
                                            case "":
                                                v2 = append(v2, (fi23 - -9223372036854775808))
                                            case "\n\t":
                                                fmt.Println(v25, fi22)
                                            default:
                                                var v41 int8 = -1
                                                _ = v41
                                            }
                                        case "hello":
                                            var v42 string = v1
                                            _ = v42
                                        case "\n\t":
                                            fmt.Println(v25, fi23, v24)
                                        default:
                                            v2 = append(v2, -9223372036854775808)
                                        }
                                        var v43 float32 = 0.1
                                        _ = v40
                                        _ = v43
                                    }
                                    _ = v39
                                case "":
                                    if (((0.0 + -2.25) * 1.5) == 3.14159) {
                                        fmt.Println(fi22, fi22, v5)
                                    }
                                case "\n\t":
                                    var v44 bool = (((v38 % ((v38) | 1)) == (^(v38))) && (127 >= 178))
                                    var v45 bool = (0 > 338)
                                    _ = v44
                                    _ = v45
                                }
                                _ = v35
                                _ = v36
                                _ = v37
                                _ = v38
                            }
                        default:
                            v24 = int16(fi22)
                        }
                        _ = v27
                    } else {
                        v46 := GenS6{F0: -1, F1: (1 <= 830), F2: 956}
                        _ = v46
                    }
                } else {
                    fmt.Println(v3, v3)
                }
            }
            fmt.Println(v4, fi22, v3)
            _ = fi23
            _ = v24
            _ = v25
        }
        v3 = (2147483647 == 42)
        _ = fi22
    }
    if (genFn1(42, 754) >= 2) {
        v47 := GenS7{F0: 127, F1: 0, F2: 473}
        _ = v47
    } else {
        v2 = append(v2, 0)
    }
    var v48 int32 = 2147483647
    var v49 string = ""
    for fi50 := 0; fi50 < 20; fi50++ {
        v51 := GenS7{F0: 0, F1: genFn1(1, 2), F2: 2}
        _ = fi50
        _ = v51
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
