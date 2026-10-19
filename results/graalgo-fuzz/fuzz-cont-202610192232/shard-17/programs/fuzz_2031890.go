package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int8; F1 string }
type GenS6 struct { F0 int16; F1 int64; F2 int64 }
func genFn0(p0 int16) uint {
    var v7 uint16 = 1
    if false {
        var v8 int32 = 0
        _ = v8
    } else {
        for fi9 := 0; fi9 < 11; fi9++ {
            if true {
                v10 := GenS1{F0: -46}
                for fi11 := 0; fi11 < 15; fi11++ {
                    switch fi9 {
                    case -1:
                        var v12 uint = 99
                        _ = v12
                    }
                    fmt.Println(fi11, fi9)
                    _ = fi11
                }
                _ = v10
            } else {
                var v13 bool = ((53 == 5) || (!(false && false)))
                _ = v13
            }
            _ = fi9
        }
        if (163 == 14) {
            var v14 uint = 919
            v15 := []int{-2147483648, 0, 9223372036854775807}
            _ = v14
            _ = v15
        } else {
            var v16 bool = (false || (-1 == -128))
            _ = v16
        }
    }
    var v17 bool = (!(((3.14159 != (-2.25 + 3.14159)) || (807 <= (v7 * 255))) && ((!true) && true)))
    if (!(!(("\n\t" != "") && v17))) {
        v7 = v7
    } else {
        v18 := GenS2{F0: 2147483647}
        if (1.5 >= ((0.1 + -2.25) * 0.1)) {
            if ((1.5 != 0.1) && true) {
                v19 := GenS3{F0: -866}
                _ = v19
            } else {
                if v17 {
                    if v17 {
                        v7 -= (v7 | (v7 - (v7 >> ((4294967295) & 63))))
                        if (923 != 65535) {
                            v20 := GenS3{F0: 1}
                            _ = v20
                        } else {
                            v7 = v7
                        }
                    } else {
                        var v21 string = "\n\t"
                        _ = v21
                    }
                    switch "" {
                    case "foo bar":
                        for fi22 := 0; fi22 < 19; fi22++ {
                            var v23 float32 = ((1.5 / float32((0.1 + -2.25))) + 1.5)
                            _ = fi22
                            _ = v23
                        }
                    case "a":
                        v17 = (v17 && ((100 >= -96) && v17))
                    case "hello":
                        var v24 uint64 = 264
                        _ = v24
                    default:
                        v25 := []int{32767, 127, 32767}
                        _ = v25
                    }
                    fmt.Println(v17, v17, v7)
                }
                v7 += 255
            }
            if (!((v7 & (^(v7))) <= (v7 - 661))) {
                v17 = false
                var v26 uint8 = 211
                _ = v26
            } else {
                v27 := []int{}
                _ = v27
            }
            if v17 {
                switch 9223372036854775807 {
                case -315:
                    switch 32767 {
                    case -9223372036854775808:
                        if v17 {
                            v17 = (4294967295 <= 42)
                            var v28 uint = uint(v7)
                            _ = v28
                        }
                    }
                case -544:
                    if v17 {
                        var v29 bool = true
                        _ = v29
                    }
                case 989:
                    p0 += (p0 * p0)
                }
                if ((1 > 122) || ((!(863 < (p0 + (p0 << ((599) & 63))))) && (945 >= 42))) {
                    p0 += p0
                    var v30 float32 = -2.25
                    if true {
                        v31 := []int{-1, int(p0), 765}
                        if len(v31) >= 2 {
                            v31 = v31[1:len(v31)]
                        }
                        _ = v31
                    }
                    _ = v30
                } else {
                    v32 := []int{128, 324}
                    if (9223372036854775807 <= 42) {
                        for fi33 := 0; fi33 < 13; fi33++ {
                            if false {
                                v7 = 2
                            }
                            var v34 uint8 = 1
                            _ = fi33
                            _ = v34
                        }
                        var v35 uint8 = 42
                        fmt.Println(v7, v7)
                        _ = v35
                    }
                    if len(v32) > 0 {
                        v32[0] = 131
                    }
                    _ = v32
                }
                p0 = p0
            }
        } else {
            for fi36 := 0; fi36 < 16; fi36++ {
                v37 := GenS2{F0: -2147483648}
                p0 -= p0
                _ = fi36
                _ = v37
            }
            p0 -= (p0 - (p0 % (((p0 >> ((255) & 63))) | 1)))
            for fi38 := 0; fi38 < 8; fi38++ {
                for fi39 := 0; fi39 < 10; fi39++ {
                    v17 = ((p0 * p0) < 1000)
                    v40 := []int{}
                    _ = fi39
                    _ = v40
                }
                _ = fi38
            }
        }
        v41 := GenS4{F0: 1}
        _ = v18
        _ = v41
    }
    var v42 int64 = 0
    _ = p0
    _ = v17
    _ = v42
    _ = v7
    return 18446744073709551615
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int16 = -32768
    var v3 string = "a"
    var v4 uint64 = 512
    var v5 uint = 210
    var v6 int64 = -1
    v1[0] = "foo bar"
    var v43 int16 = v2
    fmt.Println(v2, v4, v43)
    if false {
        var v44 bool = false
        v45 := []int{42, 805, 84, 127}
        _ = v44
        _ = v45
    } else {
        v3 = ("a" + ("a" + ""))
        if (v6 > (v6 >> ((v5) & 63))) {
            if ((!false) || true) {
                var v46 uint64 = v4
                v47 := GenS2{F0: -14}
                _ = v46
                _ = v47
            } else {
                v48 := []int{-681, 2147483647}
                _ = v48
            }
            v49 := GenS4{F0: 4294967295}
            var v50 int = -417
            _ = v49
            _ = v50
        } else {
            var v51 int = 4294967295
            _ = v51
        }
    }
    v2 = v2
    v2 = v43
    var v52 uint = genFn0(v43)
    var v53 uint16 = 2
    v54 := GenS5{F0: -40, F1: "a"}
    if (((-2.25 / 1.0) / 1e10) <= (0.0 * -2.25)) {
        v55 := GenS4{F0: 4294967295}
        for fi56 := 0; fi56 < 2; fi56++ {
            for fi57 := 0; fi57 < 19; fi57++ {
                if (50 > 0) {
                    fmt.Println(v5, v52)
                    if true {
                        v3 = v3
                    } else {
                        v3 = ""
                        if (0 < 656) {
                            v58 := GenS1{F0: int32(v4)}
                            v59 := []int{(fi57 * (fi56 << (((^(v52))) & 63))), (fi57 - fi57)}
                            v60 := GenS6{F0: v43, F1: v6, F2: v6}
                            _ = v58
                            _ = v59
                            _ = v60
                        }
                        if ((1 >= 255) || (true && true)) {
                            var v61 int8 = 1
                            _ = v61
                        } else {
                            v62 := []int{fi56, 9223372036854775807}
                            if ((^(v43)) != v43) {
                                if true {
                                    if false {
                                        v63 := GenS3{F0: fi57}
                                        v3 = "\n\t"
                                        _ = v63
                                    }
                                    if false {
                                        v43 = v2
                                        var v64 uint8 = 2
                                        v2 = (v43 - (^(v2)))
                                        _ = v64
                                    }
                                    var v65 float64 = 0.1
                                    _ = v65
                                }
                                var v66 float32 = float32(0.0)
                                _ = v66
                            }
                            _ = v62
                        }
                        switch (v3 + v3) {
                        case "\n\t":
                            v67 := []int{-2147483648, (fi57 + fi56), (^(fi56)), -892}
                            _ = v67
                        case "a":
                            fmt.Println(v53)
                        default:
                            fmt.Println(v2, v3, v43)
                        }
                    }
                    v68 := GenS6{F0: v43, F1: 0, F2: (-(v6))}
                    _ = v68
                } else {
                    switch v3 {
                    case "foo bar":
                        v69 := GenS3{F0: fi57}
                        _ = v69
                    case "\n\t":
                        if (v5 <= v5) {
                            if (!false) {
                                v52 = 2
                                if ((!(true && (387 == -428))) || false) {
                                    if (true && false) {
                                        var v70 int8 = 0
                                        _ = v70
                                    } else {
                                        fmt.Println(v3, v53, v43)
                                    }
                                }
                            }
                        }
                    default:
                        v5 = (v52 - (v5 ^ v52))
                    }
                }
                if (938 < v53) {
                    v1[414] = (((v3 + ("\n\t" + "a")) + "hello") + "hello")
                }
                var v71 string = "\n\t"
                _ = fi57
                _ = v71
            }
            switch -72 {
            case 9223372036854775807:
                v43 += v2
            case 32767:
                _ = len(v1)
            }
            v72 := []int{1, fi56, 1}
            _ = fi56
            _ = v72
        }
        _ = v55
    } else {
        var v73 uint64 = v4
        fmt.Println(v3, v4, v53)
        _ = v73
    }
    var v74 int8 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v52
    _ = v53
    _ = v54
    _ = v6
    _ = v74
}
