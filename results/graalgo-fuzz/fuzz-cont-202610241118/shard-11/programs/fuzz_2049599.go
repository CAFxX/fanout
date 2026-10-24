package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 uint16; F2 uint64 }
type GenS4 struct { F0 int32; F1 string; F2 int16 }
type GenS5 struct { F0 uint64; F1 string; F2 int32 }
type GenS6 struct { F0 uint16 }
func genFn0() uint32 {
    v6 := []int{-1}
    for fi7 := 0; fi7 < 6; fi7++ {
        switch (fi7 ^ fi7) {
        case 393:
            var v8 bool = (!("" < "a"))
            _ = v8
        case 0:
            if len(v6) > 0 {
                v6[0] = 127
            }
        default:
            var v9 uint = 778
            _ = v9
        }
        _ = fi7
    }
    v6 = append(v6, 128)
    v10 := GenS1{F0: -578}
    for fi11 := 0; fi11 < 12; fi11++ {
        var v12 uint32 = 2147483648
        _ = fi11
        _ = v12
    }
    _ = v10
    _ = v6
    return 2147483647
}
func genFn1() uint16 {
    if (false && (2 < 615)) {
        var v13 int = -9223372036854775808
        var v14 float64 = -2.25
        _ = v13
        _ = v14
    } else {
        var v15 bool = true
        v16 := []int{128, 127, -9223372036854775808, 127}
        v17 := GenS2{F0: 32767}
        _ = v15
        _ = v16
        _ = v17
    }
    var v18 bool = (0.1 >= (1e10 - 0.0))
    v19 := []int{722}
    var v20 uint8 = 67
    _ = v18
    _ = v19
    _ = v20
    return 255
}

func main() {
    var v1 uint = 0
    var v2 bool = (true && false)
    v3 := []string{"a"}
    var v4 int = -863
    var v5 uint64 = 867
    var v21 uint32 = genFn0()
    for fi22 := 0; fi22 < 15; fi22++ {
        for fi23 := 0; fi23 < 6; fi23++ {
            v24 := GenS1{F0: (-(fi23))}
            fmt.Println(v2, v21)
            v5 += v5
            _ = fi23
            _ = v24
        }
        switch ("a" + "\n\t") {
        case "\n\t":
            v21 = v21
        case "a":
            var v25 bool = ((0.0 < (-2.25 * ((-2.25 / 1.0) - (0.0 - ((3.14159 - (float32(((float64(0.0) - 3.14159) * 1e10)) - 1.5)) - 0.1))))) || v2)
            _ = v25
        case "foo bar":
            var v26 int16 = -573
            _ = v26
        }
        _ = fi22
    }
    fmt.Println(v5)
    v1 -= (v1 % ((4294967295) | 1))
    var v27 uint = 154
    fmt.Println(v21)
    if (-1 > 412) {
        fmt.Println(v4, v5, v1)
    }
    if (false && false) {
        switch 77 {
        case 753:
            v27 += (v1 / ((v1) | 1))
        case -9223372036854775808:
            v21 = 42
        }
        v28 := GenS3{F0: (v4 | v4), F1: 32767, F2: 18446744073709551615}
        _ = v28
    } else {
        if (!v2) {
            v2 = true
            var v29 uint32 = (v21 & genFn0())
            _ = v29
        } else {
            v30 := GenS1{F0: v4}
            _ = v30
        }
        var v31 float64 = 3.14159
        for fi32 := 0; fi32 < 19; fi32++ {
            switch fi32 {
            case -1:
                var v33 uint = v27
                _ = v33
            }
            _ = fi32
        }
        _ = v31
    }
    switch "\n\t" {
    case "":
        for fi34 := 0; fi34 < 20; fi34++ {
            v35 := GenS4{F0: -764, F1: "", F2: 42}
            for fi36 := 0; fi36 < 16; fi36++ {
                v37 := GenS3{F0: 6, F1: 2, F2: 42}
                _ = fi36
                _ = v37
            }
            var v38 uint = (v27 ^ v27)
            _ = fi34
            _ = v35
            _ = v38
        }
    case "foo bar":
        v1 += v27
    case "\n\t":
        var v39 uint64 = 4294967295
        _ = v39
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v4 -= (v4 + 128)
    if true {
        v40 := GenS5{F0: uint64(v1), F1: "hello", F2: 88}
        for fi41 := 0; fi41 < 4; fi41++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            switch ("a" + "") {
            case "":
                v27 = v1
            default:
                var v42 uint16 = 32767
                _ = v42
            }
            _ = fi41
        }
        _ = v40
    } else {
        if (42 >= 0) {
            var v43 uint16 = 166
            var v44 uint64 = uint64(v43)
            _ = v43
            _ = v44
        } else {
            var v45 uint = v27
            _ = v45
        }
    }
    v5 = (v5 ^ v5)
    for fi46 := 0; fi46 < 14; fi46++ {
        switch v4 {
        case -635:
            v1 = v1
        default:
            var v47 int = fi46
            _ = v47
        }
        _ = fi46
    }
    v21 += 2147483647
    v3 = append(v3, "hello")
    switch "\n\t" {
    case "foo bar":
        v2 = ((v21 - genFn0()) == v21)
    default:
        if v2 {
            if ((9223372036854775807 > (v4 << ((v1) & 63))) && false) {
                if (!v2) {
                    var v48 float64 = 1.5
                    _ = v48
                } else {
                    if true {
                        v3 = append(v3, ("a" + "a"))
                    } else {
                        var v49 int16 = 42
                        v5 -= uint64(v21)
                        _ = v49
                    }
                    for fi50 := 0; fi50 < 6; fi50++ {
                        var v51 float32 = (float32(0.0) + 3.14159)
                        var v52 uint64 = 166
                        _ = fi50
                        _ = v51
                        _ = v52
                    }
                }
            } else {
                v53 := GenS1{F0: -2147483648}
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                switch "hello" {
                case "":
                    for fi54 := 0; fi54 < 11; fi54++ {
                        if (v2 || (v1 == 42)) {
                            var v55 uint16 = 379
                            v2 = (v2 || v2)
                            for fi56 := 0; fi56 < 12; fi56++ {
                                var v57 string = "\n\t"
                                var v58 string = "foo bar"
                                _ = fi56
                                _ = v57
                                _ = v58
                            }
                            _ = v55
                        } else {
                            var v59 uint32 = (v21 | v21)
                            v60 := GenS6{F0: 65535}
                            fmt.Println(v1, v4, v21)
                            _ = v59
                            _ = v60
                        }
                        if ((v4 & (fi54 - v4)) > -9223372036854775808) {
                            var v61 uint32 = (v21 << ((v27) & 63))
                            _ = v61
                        }
                        var v62 uint8 = 0
                        _ = fi54
                        _ = v62
                    }
                case "foo bar":
                    fmt.Println(v1, v1)
                case "a":
                    var v63 uint8 = 229
                    _ = v63
                default:
                    if (32767 != (-(v4))) {
                        fmt.Println(v4, v1)
                        var v64 float64 = 0.1
                        for fi65 := 0; fi65 < 13; fi65++ {
                            fmt.Println(v2, v4, fi65)
                            var v66 float64 = float64((0.1 * (-2.25 / 1.0)))
                            _ = fi65
                            _ = v66
                        }
                        _ = v64
                    }
                }
                _ = v53
            }
            switch v4 {
            case 32768:
                switch (v4 >> ((v1) & 63)) {
                case -540:
                    for fi67 := 0; fi67 < 1; fi67++ {
                        var v68 uint = (v27 >> ((1) & 63))
                        _ = fi67
                        _ = v68
                    }
                case -1:
                    v1 -= (v27 - v1)
                case -535:
                    v3 = append(v3, ("foo bar" + ""))
                default:
                    if (!(!(-884 > 421))) {
                        v69 := GenS6{F0: 42}
                        if len(v3) > 0 {
                            v3[0] = "foo bar"
                        }
                        _ = v69
                    } else {
                        v2 = false
                    }
                }
            case 32767:
                _ = len(v3)
            }
            v5 += v5
        }
    }
    var v70 uint8 = 160
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v21
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v70
}
