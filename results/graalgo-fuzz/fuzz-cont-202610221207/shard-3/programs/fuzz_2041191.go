package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 uint16 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 uint8; F1 int64; F2 int8 }
func genFn0(p0 bool) int8 {
    v5 := []int{-2147483648, 913, 127}
    v6 := GenS1{F0: 8, F1: (false || p0), F2: 571}
    _ = p0
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 int8) string {
    switch -393 {
    case 1:
        switch -1 {
        case -430:
            if (2 > 1) {
                var v7 uint16 = 255
                for fi8 := 0; fi8 < 19; fi8++ {
                    if (!false) {
                        var v9 uint16 = (^(v7))
                        var v10 uint16 = 65535
                        _ = v10
                        _ = v9
                    } else {
                        for fi11 := 0; fi11 < 8; fi11++ {
                            v12 := []int{9223372036854775807}
                            v13 := GenS1{F0: p0, F1: true, F2: 2}
                            _ = fi11
                            _ = v12
                            _ = v13
                        }
                    }
                    if false {
                        v14 := GenS1{F0: (p0 ^ (^(p0))), F1: (18446744073709551615 < 18446744073709551615), F2: v7}
                        if false {
                            for fi15 := 0; fi15 < 3; fi15++ {
                                var v16 int32 = -1
                                var v17 bool = ((4294967295 <= 852) || (!(42 < int64(fi8))))
                                _ = fi15
                                _ = v16
                                _ = v17
                            }
                            p0 = p0
                            var v18 int32 = 2147483647
                            _ = v18
                        }
                        _ = v14
                    }
                    v7 += v7
                    _ = fi8
                }
                _ = v7
            }
        case 128:
            for fi19 := 0; fi19 < 20; fi19++ {
                if true {
                    p0 = genFn0(false)
                    if (3.14159 < 3.14159) {
                        for fi20 := 0; fi20 < 3; fi20++ {
                            p0 -= p0
                            _ = fi20
                        }
                        fmt.Println(p0, fi19)
                        if true {
                            switch "foo bar" {
                            case "a":
                                p0 += (^(p0))
                            case "":
                                var v21 uint64 = 953
                                _ = v21
                            }
                        }
                    }
                } else {
                    for fi22 := 0; fi22 < 13; fi22++ {
                        fmt.Println(p0)
                        var v23 uint32 = 624
                        var v24 int8 = genFn0((!(((-907 < -1) || ((34 > 165) && true)) && true)))
                        switch (("hello" + "hello") + "a") {
                        case "hello":
                            if ((((0.1 + 0.1) * (float32(-2.25) - 3.14159)) >= -2.25) || true) {
                                v25 := []int{(fi22 / ((128) | 1))}
                                _ = v25
                            } else {
                                switch (fi22 ^ fi19) {
                                case -413:
                                    if true {
                                        if (-2147483648 < 42) {
                                            if (false && true) {
                                                fmt.Println(v24)
                                                v23 += 0
                                            } else {
                                                fmt.Println(v23, v23, v23)
                                            }
                                        } else {
                                            var v26 uint32 = (^(v23))
                                            var v27 uint32 = v26
                                            v28 := GenS2{F0: -9223372036854775808, F1: (v27 >> ((255) & 63))}
                                            _ = v26
                                            _ = v27
                                            _ = v28
                                        }
                                        v29 := GenS3{F0: 50, F1: -1, F2: genFn0(("" <= "\n\t"))}
                                        _ = v29
                                    } else {
                                        var v30 string = "a"
                                        var v31 int8 = (p0 << ((255) & 63))
                                        _ = v30
                                        _ = v31
                                    }
                                default:
                                    var v32 uint64 = 4294967295
                                    _ = v32
                                }
                            }
                        case "foo bar":
                            if (756 < 18446744073709551615) {
                                var v33 float32 = 1.5
                                var v34 int64 = 9223372036854775807
                                _ = v33
                                _ = v34
                            }
                            if true {
                                fmt.Println(v24, fi22)
                                var v35 uint = 18446744073709551615
                                switch fi19 {
                                case 0:
                                    if (0 >= -244) {
                                        fmt.Println(v23, v23, fi19)
                                        var v36 uint32 = v23
                                        _ = v36
                                    }
                                case 231:
                                    v37 := []int{int(v23)}
                                    _ = v37
                                default:
                                    var v38 int32 = -2147483648
                                    _ = v38
                                }
                                _ = v35
                            } else {
                                var v39 int8 = p0
                                _ = v39
                            }
                            v40 := GenS2{F0: -9223372036854775808, F1: 232}
                            _ = v40
                        }
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                }
                _ = fi19
            }
        default:
            var v41 uint = 2
            _ = v41
        }
    }
    fmt.Println(p0)
    fmt.Println(p0)
    if (!false) {
        var v42 uint8 = 2
        _ = v42
    }
    var v43 float64 = -2.25
    _ = p0
    _ = v43
    return "hello"
}

func main() {
    var v1 int64 = 184
    var v2 int = 2147483647
    v3 := make(map[int]int)
    var v4 string = ("foo bar" + "\n\t")
    for fi44 := 0; fi44 < 6; fi44++ {
        v4 = v4
        _ = fi44
    }
    fmt.Println(v4, v2, v1)
    var v45 int8 = 42
    if (!(!(false || ((true || false) && true)))) {
        v4 = v4
    }
    var v46 uint8 = 248
    for fi47 := 0; fi47 < 9; fi47++ {
        fmt.Println(fi47)
        var v48 int16 = -502
        v49 := v3[v2]
        _ = fi47
        _ = v48
        _ = v49
    }
    if ((-409 != (-(v1))) || (false && (false || (!(!true))))) {
        v50 := []int{v2}
        _ = v50
    }
    var v51 bool = (1 > -379)
    switch (v2 | v2) {
    case 2147483647:
        fmt.Println(v1)
    case 9223372036854775807:
        v52 := GenS2{F0: v1, F1: 4294967295}
        _ = v52
    case -336:
        if (!true) {
            v4 = v4
        }
    default:
        fmt.Println(v1)
    }
    for fi53 := 0; fi53 < 2; fi53++ {
        v51 = (65535 < 65535)
        _ = fi53
    }
    for fi54 := 0; fi54 < 4; fi54++ {
        for fi55 := 0; fi55 < 20; fi55++ {
            var v56 uint16 = 0
            _ = fi55
            _ = v56
        }
        var v57 int64 = v1
        v58 := []int{(fi54 + (v2 & (-(v2)))), 127, 42}
        _ = fi54
        _ = v57
        _ = v58
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v51
}
