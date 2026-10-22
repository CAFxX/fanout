package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 bool; F2 uint }
type GenS3 struct { F0 int64; F1 uint; F2 uint32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int32 }
func genFn0(p0 bool, p1 int64, p2 int) string {
    for fi5 := 0; fi5 < 3; fi5++ {
        if p0 {
            var v6 uint32 = 2
            _ = v6
        }
        _ = fi5
    }
    fmt.Println(p2, p1)
    p0 = (1000 == 42)
    _ = p0
    _ = p1
    _ = p2
    return ("hello" + "")
}
func genFn1(p0 int32, p1 int32, p2 int16) uint64 {
    if (!(1 >= 0)) {
        if (false && (65535 != 379)) {
            v7 := []int{840, 332, 32767}
            _ = v7
        } else {
            v8 := GenS1{F0: 970}
            _ = v8
        }
        p0 -= (p1 >> ((0) & 63))
    }
    p0 = (p0 ^ p0)
    v9 := GenS1{F0: 0}
    for fi10 := 0; fi10 < 13; fi10++ {
        switch (fi10 << ((65535) & 63)) {
        case -2147483648:
            for fi11 := 0; fi11 < 6; fi11++ {
                if true {
                    fmt.Println(fi11, p0)
                    v12 := GenS2{F0: 255, F1: false, F2: 897}
                    _ = v12
                } else {
                    var v13 int64 = 9223372036854775807
                    v14 := GenS3{F0: v13, F1: 65535, F2: 1}
                    v15 := []int{fi11, (fi10 << ((42) & 63)), fi10, fi10}
                    v15 = append(v15, fi10)
                    _ = v13
                    _ = v14
                    _ = v15
                }
                _ = fi11
            }
        case -610:
            if (true && (32767 == 949)) {
                for fi16 := 0; fi16 < 6; fi16++ {
                    v17 := []int{(^(fi16)), (fi16 ^ (fi10 | 4294967295))}
                    _ = fi16
                    _ = v17
                }
                p1 = (p1 * (p0 >> ((774) & 63)))
            } else {
                p1 = p0
                switch (fi10 + 9223372036854775807) {
                case 9223372036854775807:
                    v18 := GenS3{F0: 505, F1: 18446744073709551615, F2: 65535}
                    _ = v18
                case 128:
                    v19 := []int{}
                    _ = v19
                case 0:
                    v20 := []int{1, -9223372036854775808, (fi10 + (^(fi10)))}
                    _ = v20
                }
            }
        }
        var v21 float64 = 3.14159
        v22 := GenS4{F0: 47}
        _ = fi10
        _ = v21
        _ = v22
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 1
}

func main() {
    var v1 bool = false
    var v2 bool = ((0 != -2147483648) || v1)
    var v3 float64 = 1.5
    var v4 uint8 = 244
    for fi23 := 0; fi23 < 19; fi23++ {
        v24 := []int{}
        v25 := GenS1{F0: 18446744073709551615}
        _ = fi23
        _ = v24
        _ = v25
    }
    switch -2147483648 {
    case 313:
        if false {
            switch "hello" {
            case "\n\t":
                var v26 uint8 = v4
                _ = v26
            case "":
                switch -224 {
                case 0:
                    if (!(65535 != 94)) {
                        for fi27 := 0; fi27 < 18; fi27++ {
                            var v28 int64 = 1
                            v1 = v1
                            _ = fi27
                            _ = v28
                        }
                    } else {
                        if v2 {
                            v4 = (v4 ^ (v4 * v4))
                            v1 = (32767 >= 32767)
                            switch (("a" + "foo bar") + "foo bar") {
                            case "":
                                v29 := GenS5{F0: -1}
                                _ = v29
                            case "foo bar":
                                fmt.Println(v1, v3)
                            default:
                                if (v1 || v2) {
                                    if v1 {
                                        v30 := []int{42, 9223372036854775807}
                                        _ = v30
                                    }
                                    v31 := GenS4{F0: 255}
                                    _ = v31
                                } else {
                                    if true {
                                        if v2 {
                                            if v1 {
                                                var v32 int32 = int32(v4)
                                                _ = v32
                                            }
                                            v3 += v3
                                        } else {
                                            v33 := []int{1}
                                            _ = v33
                                        }
                                        v34 := GenS4{F0: (v4 / ((127) | 1))}
                                        v1 = ((-71 > 1) || (42 < 42))
                                        _ = v34
                                    }
                                    v1 = v2
                                }
                            }
                        } else {
                            if true {
                                for fi35 := 0; fi35 < 19; fi35++ {
                                    if (v2 || ((v1 || v2) || false)) {
                                        var v36 uint8 = 126
                                        _ = v36
                                    }
                                    switch 127 {
                                    case 78:
                                        var v37 int16 = 518
                                        _ = v37
                                    }
                                    _ = fi35
                                }
                            }
                        }
                        var v38 float32 = 3.14159
                        _ = v38
                    }
                case 64:
                    switch "" {
                    case "a":
                        v39 := []int{}
                        _ = v39
                    case "\n\t":
                        var v40 int = 32768
                        _ = v40
                    default:
                        var v41 uint = 2
                        _ = v41
                    }
                default:
                    v4 -= (v4 % ((v4) | 1))
                }
            default:
                fmt.Println(v3, v1, v2)
            }
        }
    case 4294967295:
        v42 := GenS4{F0: (v4 * v4)}
        _ = v42
    case 32767:
        switch "a" {
        case "foo bar":
            switch "a" {
            case "foo bar":
                v3 -= v3
            default:
                v43 := []int{}
                _ = v43
            }
        case "hello":
            for fi44 := 0; fi44 < 13; fi44++ {
                v45 := []int{(fi44 << ((506) & 63)), (fi44 / ((fi44) | 1)), (fi44 >> ((123) & 63))}
                var v46 uint32 = 2
                _ = fi44
                _ = v45
                _ = v46
            }
        }
    default:
        var v47 int8 = 1
        _ = v47
    }
    switch ("\n\t" + ("hello" + "\n\t")) {
    case "foo bar":
        if (v1 || false) {
            v48 := []int{0}
            v4 = v4
            _ = v48
        }
    case "\n\t":
        var v49 uint = 754
        _ = v49
    case "a":
        v4 -= (v4 & v4)
    default:
        if ((!v1) || v2) {
            fmt.Println(v2, v3, v2)
        } else {
            v3 -= v3
        }
    }
    if v1 {
        if v1 {
            switch "" {
            case "foo bar":
                if (!v2) {
                    for fi50 := 0; fi50 < 4; fi50++ {
                        var v51 uint = 0
                        _ = fi50
                        _ = v51
                    }
                    fmt.Println(v1, v4)
                    v3 = 0.1
                }
            case "hello":
                for fi52 := 0; fi52 < 17; fi52++ {
                    v53 := GenS2{F0: 255, F1: v2, F2: 65535}
                    for fi54 := 0; fi54 < 12; fi54++ {
                        var v55 uint32 = 511
                        v3 = float64(3.14159)
                        if v1 {
                            v56 := GenS6{F0: -1}
                            v57 := []int{(fi54 & (fi52 << ((18446744073709551615) & 63))), (fi54 & (fi54 + -49)), (fi54 << ((924) & 63)), fi52}
                            _ = v56
                            _ = v57
                        }
                        _ = fi54
                        _ = v55
                    }
                    fmt.Println(fi52, fi52, fi52)
                    _ = fi52
                    _ = v53
                }
            case "":
                v58 := []int{}
                _ = v58
            default:
                var v59 int64 = -1
                _ = v59
            }
            var v60 bool = true
            _ = v60
        }
        var v61 float64 = ((v3 + (float64(3.14159) * 0.0)) + (1e10 / 1.5))
        fmt.Println(v1, v2)
        _ = v61
    }
    switch -294 {
    case 4294967295:
        v2 = ("" >= "\n\t")
    case 127:
        v4 = v4
    default:
        v62 := []int{}
        _ = v62
    }
    v63 := GenS4{F0: v4}
    var v64 uint8 = v4
    for fi65 := 0; fi65 < 12; fi65++ {
        var v66 float64 = -2.25
        _ = fi65
        _ = v66
    }
    var v67 uint32 = 2147483647
    var v68 bool = (!v2)
    fmt.Println(v68)
    v69 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v64:", v64)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    fmt.Println("slicelen v69:", len(v69), v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v63
    _ = v64
    _ = v67
    _ = v68
    _ = v69
}
