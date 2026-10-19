package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 string; F1 int32; F2 bool }
type GenS4 struct { F0 int8 }
func genFn0(p0 uint16, p1 uint16, p2 int8) uint8 {
    switch -681 {
    case 128:
        fmt.Println(p2, p1)
    default:
        v5 := []int{670, 127}
        _ = v5
    }
    switch (("foo bar" + "a") + "\n\t") {
    case "foo bar":
        if false {
            var v6 int8 = -1
            _ = v6
        }
    }
    p2 = -1
    switch ("hello" + "a") {
    case "":
        var v7 uint32 = 2
        _ = v7
    }
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 string = "foo bar"
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return 255
}
func genFn1(p0 int32, p1 bool, p2 int32) int32 {
    if (255 != 66) {
        v10 := GenS1{F0: 1000, F1: -1}
        for fi11 := 0; fi11 < 7; fi11++ {
            for fi12 := 0; fi12 < 13; fi12++ {
                var v13 uint32 = 751
                _ = fi12
                _ = v13
            }
            if (!p1) {
                fmt.Println(p1)
                var v14 int16 = -1
                v14 = v14
                _ = v14
            }
            _ = fi11
        }
        _ = v10
    }
    if p1 {
        var v15 uint64 = 1
        v16 := GenS1{F0: 443, F1: 400}
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 3; fi17++ {
        p0 = p0
        _ = fi17
    }
    _ = p0
    _ = p1
    _ = p2
    return p2
}

func main() {
    var v1 uint64 = 36
    v2 := make(map[string]bool)
    v3 := []string{"hello", "\n\t", "a", "a"}
    v4 := make(map[int]int)
    _ = len(v3)
    var v18 int = 32768
    v18 = (^(v18))
    v19 := GenS1{F0: 32767, F1: 1}
    for fi20 := 0; fi20 < 17; fi20++ {
        var v21 int8 = 100
        _ = len(v4)
        var v22 float64 = 1.5
        _ = fi20
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 12; fi23++ {
        for fi24 := 0; fi24 < 4; fi24++ {
            v1 += (v1 | (v1 % ((v1) | 1)))
            var v25 uint8 = 213
            if len(v3) > 0 {
                v3[0] = ("foo bar" + "")
            }
            switch (v18 - fi24) {
            case 9223372036854775807:
                switch "hello" {
                case "":
                    switch "foo bar" {
                    case "":
                        if (0 != 65535) {
                            if len(v3) > 0 {
                                v3[0] = ("a" + "")
                            }
                            v3 = append(v3, ("a" + ""))
                        } else {
                            v4[v18] = (^(v18))
                        }
                    case "foo bar":
                        switch "hello" {
                        case "foo bar":
                            switch fi23 {
                            case 9223372036854775807:
                                var v26 int8 = 24
                                var v27 uint = 255
                                delete(v2, ("\n\t" + ""))
                                _ = v26
                                _ = v27
                            }
                        case "a":
                            switch fi23 {
                            case 127:
                                switch ("\n\t" + "a") {
                                case "foo bar":
                                    v18 = 2147483647
                                default:
                                    switch -2147483648 {
                                    case 128:
                                        if (v18 > fi24) {
                                            v28 := GenS2{F0: 0, F1: 0}
                                            var v29 float32 = ((float32(((0.0 * 0.0) - (1e10 / 3.14159))) * 1e10) - (float32((-2.25 * 1.5)) - 0.0))
                                            _ = v28
                                            _ = v29
                                        } else {
                                            var v30 uint8 = genFn0(164, 0, 1)
                                            switch fi24 {
                                            case 156:
                                                fmt.Println(v18, v30)
                                            case 23:
                                                v31 := GenS1{F0: -796, F1: 9223372036854775807}
                                                _ = v31
                                            default:
                                                _ = len(v4)
                                            }
                                            _ = v30
                                        }
                                    default:
                                        var v32 uint64 = (v1 & 18446744073709551615)
                                        _ = v32
                                    }
                                }
                            case 32768:
                                var v33 float32 = -2.25
                                _ = v33
                            }
                        case "":
                            v1 -= v1
                        default:
                            var v34 string = ("\n\t" + ("\n\t" + "hello"))
                            _ = v34
                        }
                    default:
                        if len(v3) > 0 {
                            v3[0] = ("a" + "")
                        }
                    }
                case "\n\t":
                    v25 += v25
                case "foo bar":
                    if false {
                        fmt.Println(fi23, v25, fi24)
                        v4[fi24] = fi23
                    }
                }
            case 32768:
                switch "foo bar" {
                case "\n\t":
                    switch fi23 {
                    case -1:
                        var v35 string = ((("a" + "") + "") + "a")
                        v25 -= v25
                        _ = v35
                    }
                case "foo bar":
                    switch (fi24 >> ((uint(v25)) & 63)) {
                    case 128:
                        fmt.Println(v25)
                    case -1:
                        if (283 >= 18446744073709551615) {
                            v1 = v1
                        }
                    case -734:
                        var v36 bool = (false && false)
                        _ = v36
                    default:
                        delete(v4, (-(v18)))
                    }
                case "":
                    v18 = fi24
                }
            case -2147483648:
                v2["a"] = (genFn0(787, 2, -72) < 127)
            }
            _ = fi24
            _ = v25
        }
        _ = fi23
    }
    var v37 float64 = 3.14159
    for fi38 := 0; fi38 < 16; fi38++ {
        v2["\n\t"] = true
        for fi39 := 0; fi39 < 7; fi39++ {
            fmt.Println(v1, v18)
            var v40 uint8 = genFn0(255, 0, 42)
            _ = fi39
            _ = v40
        }
        var v41 uint64 = v1
        _ = fi38
        _ = v41
    }
    var v42 int64 = -9223372036854775808
    if (-1 > 42) {
        var v43 int64 = (v42 % (((v42 * (^(v42)))) | 1))
        _ = v43
    } else {
        if (float32(v37) != -2.25) {
            var v44 float64 = v37
            v37 = 1.5
            _ = v44
        }
    }
    if (37 != 59) {
        var v45 float64 = v37
        if ((564 == -937) && (false && (42 <= 892))) {
            if (("foo bar" >= "foo bar") || false) {
                fmt.Println(v18, v18, v37)
                switch "a" {
                case "":
                    switch (v18 >> ((4294967295) & 63)) {
                    case -9223372036854775808:
                        v18 -= v18
                    case 718:
                        fmt.Println(v1, v42, v1)
                    case 128:
                        var v46 uint8 = 2
                        _ = v46
                    default:
                        v18 -= (v18 ^ int(v1))
                    }
                case "a":
                    if false {
                        v45 = v45
                    } else {
                        v1 -= (v1 >> ((4294967295) & 63))
                        var v47 int16 = 0
                        if false {
                            var v48 uint64 = (v1 & v1)
                            _ = v48
                        } else {
                            switch v18 {
                            case 32768:
                                _ = len(v3)
                            default:
                                if (v18 != v18) {
                                    var v49 int8 = -8
                                    fmt.Println(v18)
                                    _ = v49
                                } else {
                                    v50 := GenS2{F0: 1, F1: (-(v18))}
                                    if (-1 <= -2147483648) {
                                        var v51 int16 = v47
                                        var v52 uint32 = 2147483648
                                        _ = v51
                                        _ = v52
                                    }
                                    var v53 int16 = v47
                                    _ = v50
                                    _ = v53
                                }
                            }
                            for fi54 := 0; fi54 < 6; fi54++ {
                                switch (^(v18)) {
                                case 0:
                                    var v55 uint16 = 0
                                    _ = v55
                                }
                                v37 = v37
                                for fi56 := 0; fi56 < 14; fi56++ {
                                    fmt.Println(fi54, fi54, fi56)
                                    var v57 uint64 = 2
                                    var v58 int32 = genFn1(-1, (v45 > (v45 - (float64((3.14159 * 3.14159)) / (0.1 + v45)))), 42)
                                    _ = fi56
                                    _ = v57
                                    _ = v58
                                }
                                _ = fi54
                            }
                            v59 := GenS3{F0: (("hello" + "") + "a"), F1: 721, F2: ((v42 % ((9223372036854775807) | 1)) >= (v42 % ((v42) | 1)))}
                            _ = v59
                        }
                        _ = v47
                    }
                case "foo bar":
                    var v60 uint64 = v1
                    _ = v60
                }
            }
            if true {
                var v61 float32 = 3.14159
                var v62 string = ("\n\t" + "hello")
                fmt.Println(v42)
                _ = v61
                _ = v62
            }
            v18 += (-(v18))
        }
        v63 := GenS4{F0: -1}
        _ = v45
        _ = v63
    } else {
        if true {
            var v64 int8 = -119
            var v65 int32 = 42
            _ = v64
            _ = v65
        } else {
            for fi66 := 0; fi66 < 5; fi66++ {
                for fi67 := 0; fi67 < 16; fi67++ {
                    switch (-(fi66)) {
                    case 128:
                        var v68 float32 = float32((1.5 + (v37 * v37)))
                        _ = v68
                    case 32768:
                        v42 = v42
                    default:
                        v69 := GenS1{F0: 42, F1: (v42 * v42)}
                        _ = v69
                    }
                    _ = fi67
                }
                var v70 uint64 = 4294967295
                _ = fi66
                _ = v70
            }
            v1 = 42
            fmt.Println(v42, v42, v18)
        }
    }
    var v71 int = (v18 & v18)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v42
    _ = v71
}
