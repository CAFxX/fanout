package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool; F1 string }
type GenS5 struct { F0 int32; F1 int8; F2 uint }
type GenS6 struct { F0 int; F1 uint8 }

func main() {
    var v1 string = "a"
    var v2 bool = (1 == 255)
    v3 := []int{0, 4294967295, 2147483647}
    var v4 uint64 = 225
    if (0 >= 256) {
        if v2 {
            v2 = (2147483647 < 2147483647)
            v1 = v1
            for fi5 := 0; fi5 < 11; fi5++ {
                var v6 int16 = 1000
                _ = len(v3)
                _ = fi5
                _ = v6
            }
        } else {
            var v7 uint32 = 11
            _ = v7
        }
    } else {
        if (0.1 == float32(float64(-2.25))) {
            fmt.Println(v4)
            if false {
                v8 := GenS1{F0: 718}
                _ = v8
            } else {
                var v9 uint32 = 42
                if (!(113 >= 891)) {
                    switch 128 {
                    case 0:
                        for fi10 := 0; fi10 < 8; fi10++ {
                            var v11 uint = 1
                            if false {
                                if len(v3) > 0 {
                                    v3[0] = 1
                                }
                                for fi12 := 0; fi12 < 14; fi12++ {
                                    if false {
                                        v11 -= (^(v11))
                                        v11 += v11
                                    }
                                    _ = fi12
                                }
                                if (261 > 32767) {
                                    if ((!v2) && false) {
                                        v3 = append(v3, (fi10 / (((fi10 >> ((907) & 63))) | 1)))
                                        v3 = append(v3, (fi10 % ((fi10) | 1)))
                                        var v13 uint64 = (^(v4))
                                        _ = v13
                                    } else {
                                        v14 := GenS2{F0: 981}
                                        var v15 int8 = 80
                                        _ = v14
                                        _ = v15
                                    }
                                    v1 = "foo bar"
                                    switch (fi10 + (fi10 / (((fi10 | (fi10 & -1))) | 1))) {
                                    case -729:
                                        var v16 int16 = 478
                                        _ = v16
                                    }
                                }
                            }
                            v17 := GenS2{F0: fi10}
                            _ = fi10
                            _ = v11
                            _ = v17
                        }
                    default:
                        var v18 float64 = (-2.25 * (3.14159 - (((-2.25 + 0.0) + 1e10) - ((float64(3.14159) - 1e10) + 1.5))))
                        _ = v18
                    }
                    fmt.Println(v2)
                    var v19 uint = 2
                    _ = v19
                } else {
                    v20 := GenS2{F0: 9223372036854775807}
                    _ = v20
                }
                _ = v9
            }
        } else {
            v3 = append(v3, 680)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        v4 = v4
        v2 = false
    }
    v3 = append(v3, 32767)
    fmt.Println(v4, v2, v4)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch v1 {
    case "\n\t":
        var v21 uint64 = v4
        _ = v21
    case "a":
        v2 = v2
    default:
        v22 := GenS2{F0: 768}
        _ = v22
    }
    fmt.Println(v1, v4)
    switch 0 {
    case -9223372036854775808:
        v23 := GenS2{F0: 32768}
        _ = v23
    case 127:
        var v24 string = v1
        _ = v24
    }
    var v25 int32 = -1
    for fi26 := 0; fi26 < 14; fi26++ {
        for fi27 := 0; fi27 < 12; fi27++ {
            v28 := GenS2{F0: fi26}
            var v29 bool = (v1 != ((v1 + v1) + (v1 + "a")))
            switch (fi27 * fi26) {
            case 42:
                if (65535 <= 423) {
                    if true {
                        v29 = ((!false) && (42 > 0))
                        if (true && v2) {
                            if (!((!(2 >= 2)) || (-128 == 0))) {
                                fmt.Println(v25)
                            }
                            v30 := GenS1{F0: 135}
                            v4 = (v4 / (((v4 + v4)) | 1))
                            _ = v30
                        }
                    }
                } else {
                    if true {
                        var v31 uint = 1
                        _ = v31
                    } else {
                        v29 = false
                        var v32 uint64 = v4
                        var v33 int = (fi26 % ((32768) | 1))
                        _ = v32
                        _ = v33
                    }
                    var v34 bool = v29
                    if (!v29) {
                        v1 = ("foo bar" + v1)
                        switch v1 {
                        case "hello":
                            v29 = (-2.25 > 3.14159)
                        default:
                            switch 32768 {
                            case 121:
                                var v35 float32 = 1.5
                                _ = v35
                            case 2147483647:
                                if v2 {
                                    fmt.Println(fi27, v1)
                                    v36 := GenS3{F0: ("hello" + "hello")}
                                    switch 32768 {
                                    case 0:
                                        var v37 uint64 = 0
                                        _ = v37
                                    case -1:
                                        v38 := GenS2{F0: (fi26 / (((fi26 / ((32768) | 1))) | 1))}
                                        _ = v38
                                    default:
                                        var v39 float32 = -2.25
                                        _ = v39
                                    }
                                    _ = v36
                                }
                            case 32768:
                                v40 := GenS4{F0: (0 != 351), F1: v1}
                                _ = v40
                            }
                        }
                    } else {
                        var v41 uint = 319
                        var v42 bool = v2
                        var v43 int64 = 465
                        _ = v41
                        _ = v42
                        _ = v43
                    }
                    _ = v34
                }
            default:
                v4 -= 18446744073709551615
            }
            _ = fi27
            _ = v28
            _ = v29
        }
        for fi44 := 0; fi44 < 2; fi44++ {
            var v45 int32 = 2147483647
            _ = fi44
            _ = v45
        }
        _ = fi26
    }
    for fi46 := 0; fi46 < 13; fi46++ {
        v1 = v1
        fmt.Println(v2, v1, v2)
        if v2 {
            var v47 uint = 2
            _ = v47
        } else {
            if (v2 && v2) {
                for fi48 := 0; fi48 < 6; fi48++ {
                    if false {
                        if (!true) {
                            fmt.Println(v2, fi46)
                            v49 := GenS5{F0: v25, F1: 3, F2: 0}
                            fmt.Println(v25, v2, v25)
                            _ = v49
                        } else {
                            v50 := GenS2{F0: 32767}
                            _ = len(v3)
                            _ = v50
                        }
                    } else {
                        fmt.Println(v1, v4)
                        switch (fi46 ^ fi46) {
                        case 2147483647:
                            var v51 int32 = v25
                            _ = v51
                        case 1:
                            v1 = (v1 + ("foo bar" + v1))
                        case -831:
                            v52 := GenS2{F0: fi46}
                            _ = v52
                        default:
                            v4 = v4
                        }
                        var v53 int = 32767
                        _ = v53
                    }
                    _ = fi48
                }
                v4 = 4294967295
            } else {
                fmt.Println(fi46, fi46)
                for fi54 := 0; fi54 < 18; fi54++ {
                    v1 = v1
                    fmt.Println(v25, fi54)
                    _ = fi54
                }
            }
            v3 = append(v3, fi46)
        }
        _ = fi46
    }
    switch -942 {
    case 128:
        fmt.Println(v4, v2, v4)
    case 527:
        v55 := GenS5{F0: v25, F1: 100, F2: 42}
        _ = v55
    case 98:
        for fi56 := 0; fi56 < 18; fi56++ {
            var v57 int64 = -1
            _ = fi56
            _ = v57
        }
    default:
        var v58 int = -1
        _ = v58
    }
    switch (("" + (v1 + "\n\t")) + v1) {
    case "":
        _ = len(v3)
    case "foo bar":
        v25 = (v25 / (((-(v25))) | 1))
    case "a":
        v59 := GenS6{F0: -941, F1: 178}
        _ = v59
    }
    for fi60 := 0; fi60 < 8; fi60++ {
        var v61 int32 = (v25 << ((217) & 63))
        v3 = append(v3, fi60)
        var v62 uint = 2
        _ = fi60
        _ = v61
        _ = v62
    }
    for fi63 := 0; fi63 < 9; fi63++ {
        for fi64 := 0; fi64 < 1; fi64++ {
            if v2 {
                v1 = "\n\t"
            } else {
                v2 = v2
            }
            _ = fi64
        }
        _ = fi63
    }
    var v65 uint = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v65
}
