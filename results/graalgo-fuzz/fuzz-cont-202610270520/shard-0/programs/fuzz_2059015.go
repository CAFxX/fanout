package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 int; F2 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int16; F1 int64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint; F1 bool }
func genFn0() uint32 {
    var v7 int32 = 885
    switch 42 {
    case 128:
        var v8 uint16 = 32767
        _ = v8
    case 2147483647:
        fmt.Println(v7)
    case -445:
        v9 := []int{32768, 32767}
        _ = v9
    default:
        v10 := []int{-157}
        _ = v10
    }
    switch (("a" + "foo bar") + ("\n\t" + "")) {
    case "a":
        for fi11 := 0; fi11 < 8; fi11++ {
            v7 = v7
            v12 := []int{(fi11 | fi11), (fi11 + 891)}
            switch (fi11 ^ (fi11 << ((0) & 63))) {
            case 127:
                v13 := GenS1{F0: 42}
                _ = v13
            case -9223372036854775808:
                var v14 uint8 = 2
                _ = v14
            }
            _ = fi11
            _ = v12
        }
    default:
        v15 := GenS1{F0: -9223372036854775808}
        _ = v15
    }
    v7 += v7
    fmt.Println(v7)
    _ = v7
    return 345
}
func genFn1(p0 uint) int32 {
    p0 = p0
    v16 := []int{}
    v17 := GenS1{F0: -9223372036854775808}
    _ = p0
    _ = v16
    _ = v17
    return 42
}

func main() {
    v1 := []uint32{529, 137, 0, 0}
    var v2 int = 128
    v3 := make(map[int]bool)
    var v4 uint64 = 4294967295
    var v5 int16 = 728
    var v6 int = v2
    v18 := GenS1{F0: -9223372036854775808}
    var v19 int32 = 42
    var v20 uint16 = 397
    v3[-9223372036854775808] = ("" <= "foo bar")
    switch "hello" {
    case "\n\t":
        switch (v2 - int(v19)) {
        case -1:
            v21 := GenS2{F0: 2147483648, F1: 0, F2: 1}
            _ = v21
        case -9223372036854775808:
            fmt.Println(v4)
        case -201:
            v22 := GenS3{F0: (v19 + v19)}
            _ = v22
        default:
            switch "hello" {
            case "\n\t":
                v23 := GenS4{F0: (v5 * int16(v2)), F1: -9223372036854775808}
                _ = v23
            case "foo bar":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        }
    }
    v3[v6] = false
    v24 := GenS2{F0: 1, F1: (v6 * 579), F2: 0}
    var v25 int = v6
    v2 += (v25 - v25)
    var v26 int64 = int64(v19)
    fmt.Println(v26, v20)
    if true {
        v27 := GenS1{F0: v26}
        if len(v1) > 0 {
            v1[0] = 535
        }
        _ = v27
    } else {
        v1 = append(v1, 170)
        if false {
            v5 += v5
        }
    }
    if false {
        if (0 < 638) {
            if false {
                _ = len(v1)
                if (false || (204 != (v20 | 32767))) {
                    var v28 string = ((("\n\t" + "hello") + "") + "a")
                    for fi29 := 0; fi29 < 17; fi29++ {
                        var v30 int16 = (v5 & v5)
                        _ = fi29
                        _ = v30
                    }
                    _ = v28
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v31 := GenS4{F0: v5, F1: (v26 & v26)}
                    for fi32 := 0; fi32 < 15; fi32++ {
                        v25 += fi32
                        _ = fi32
                    }
                    _ = v31
                }
                if (true || false) {
                    v33 := GenS1{F0: v26}
                    fmt.Println(v19)
                    _ = v33
                } else {
                    delete(v3, v6)
                }
            } else {
                switch ("foo bar" + "a") {
                case "":
                    if (42 < -15) {
                        var v34 uint = 388
                        _ = v34
                    } else {
                        var v35 uint32 = 483
                        if (false && true) {
                            if (!false) {
                                var v36 uint16 = 2
                                if true {
                                    var v37 int8 = 0
                                    _ = v37
                                } else {
                                    for fi38 := 0; fi38 < 12; fi38++ {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = fi38
                                    }
                                    v25 -= v2
                                    for fi39 := 0; fi39 < 12; fi39++ {
                                        var v40 int32 = -89
                                        var v41 uint64 = 18446744073709551615
                                        if false {
                                            _ = len(v3)
                                            v3[(^(v6))] = false
                                            for fi42 := 0; fi42 < 11; fi42++ {
                                                v2 = int(v26)
                                                fmt.Println(v35, v35)
                                                var v43 uint64 = (^(v41))
                                                _ = fi42
                                                _ = v43
                                            }
                                        }
                                        _ = fi39
                                        _ = v40
                                        _ = v41
                                    }
                                }
                                for fi44 := 0; fi44 < 16; fi44++ {
                                    var v45 uint64 = 42
                                    var v46 uint8 = 27
                                    _ = fi44
                                    _ = v45
                                    _ = v46
                                }
                                _ = v36
                            }
                            fmt.Println(v4, v20)
                        }
                        var v47 uint64 = (^(v4))
                        _ = v35
                        _ = v47
                    }
                case "foo bar":
                    fmt.Println(v26, v20, v25)
                case "hello":
                    for fi48 := 0; fi48 < 20; fi48++ {
                        fmt.Println(v4)
                        v49 := GenS3{F0: (v19 | (v19 * v19))}
                        _ = fi48
                        _ = v49
                    }
                default:
                    for fi50 := 0; fi50 < 16; fi50++ {
                        v2 += int(v20)
                        _ = fi50
                    }
                }
                for fi51 := 0; fi51 < 14; fi51++ {
                    var v52 int8 = -1
                    _ = fi51
                    _ = v52
                }
                switch v2 {
                case -2147483648:
                    var v53 int = (v6 - 32768)
                    _ = v53
                case 32768:
                    if true {
                        fmt.Println(v6, v19, v2)
                    } else {
                        switch (v25 << ((554) & 63)) {
                        case 9223372036854775807:
                            for fi54 := 0; fi54 < 11; fi54++ {
                                var v55 int = (-(v2))
                                _ = fi54
                                _ = v55
                            }
                        case 4294967295:
                            v3[32767] = (true && true)
                        case 128:
                            fmt.Println(v6)
                        default:
                            for fi56 := 0; fi56 < 3; fi56++ {
                                fmt.Println(v2, v19, fi56)
                                for fi57 := 0; fi57 < 6; fi57++ {
                                    var v58 bool = true
                                    v26 += v26
                                    if ((!(255 != 42)) && true) {
                                        var v59 float64 = float64(0.1)
                                        var v60 uint32 = 2
                                        _ = v59
                                        _ = v60
                                    }
                                    _ = fi57
                                    _ = v58
                                }
                                v61 := GenS5{F0: 4294967295}
                                _ = fi56
                                _ = v61
                            }
                        }
                        v3[v6] = false
                    }
                case -9223372036854775808:
                    switch "\n\t" {
                    case "hello":
                        for fi62 := 0; fi62 < 3; fi62++ {
                            fmt.Println(v19, v20)
                            _ = fi62
                        }
                    case "\n\t":
                        v19 = 2147483647
                    default:
                        v25 += (-(v2))
                    }
                }
            }
        } else {
            v63 := GenS1{F0: v26}
            _ = v63
        }
    } else {
        v64 := GenS6{F0: 18446744073709551615, F1: (-2.25 <= (1e10 + -2.25))}
        var v65 uint8 = 11
        _ = v64
        _ = v65
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
