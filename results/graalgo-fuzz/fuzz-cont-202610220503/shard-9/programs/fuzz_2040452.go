package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint8; F1 int }
type GenS5 struct { F0 int8; F1 int64; F2 uint }
type GenS6 struct { F0 string; F1 int; F2 int }
func genFn0(p0 bool) uint64 {
    fmt.Println(p0)
    for fi4 := 0; fi4 < 16; fi4++ {
        var v5 uint32 = 42
        v5 = (v5 - (v5 & 743))
        v6 := []int{fi4, fi4}
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 int8 = -128
    v8 := []int{127, -93, 32767, -496}
    _ = p0
    _ = v7
    _ = v8
    return 115
}
func genFn1(p0 int8, p1 uint16) int {
    if false {
        v9 := []int{-594, 32768}
        if len(v9) > 0 {
            v9[0] = 4294967295
        }
        _ = v9
    }
    v10 := []int{32768, 127}
    _ = p0
    _ = p1
    _ = v10
    return 9223372036854775807
}

func main() {
    var v1 uint64 = 0
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]int)
    for fi11 := 0; fi11 < 18; fi11++ {
        v1 = (v1 % ((genFn0(true)) | 1))
        for fi12 := 0; fi12 < 18; fi12++ {
            var v13 int = genFn1(-128, 0)
            fmt.Println(v1)
            fmt.Println(fi12, fi12)
            _ = fi12
            _ = v13
        }
        v14 := []int{genFn1(127, 255), fi11, (fi11 - genFn1(-80, 156))}
        _ = fi11
        _ = v14
    }
    var v15 uint = 2
    v16 := GenS1{F0: v1}
    delete(v2, 4294967295)
    if (!(4294967295 >= 0)) {
        fmt.Println(v15, v1)
        v15 += 138
        if true {
            var v17 int = 32767
            _ = v17
        } else {
            for fi18 := 0; fi18 < 8; fi18++ {
                var v19 uint8 = 42
                switch ("\n\t" + (("hello" + (("\n\t" + "hello") + "a")) + "hello")) {
                case "hello":
                    if ("foo bar" != ("a" + ("foo bar" + "hello"))) {
                        var v20 uint = v15
                        v1 -= (v1 & v1)
                        _ = len(v3)
                        _ = v20
                    } else {
                        v2[4294967295] = 42
                        v19 -= uint8(v1)
                        for fi21 := 0; fi21 < 8; fi21++ {
                            if false {
                                v22 := GenS2{F0: false, F1: 127}
                                var v23 int8 = 42
                                if (!(-637 != -1)) {
                                    v23 = v23
                                    var v24 int16 = int16(fi21)
                                    fmt.Println(v19, v19)
                                    _ = v24
                                }
                                _ = v22
                                _ = v23
                            }
                            _ = fi21
                        }
                    }
                case "":
                    for fi25 := 0; fi25 < 12; fi25++ {
                        fmt.Println(v15)
                        v26 := GenS1{F0: (v1 ^ v1)}
                        delete(v3, "\n\t")
                        _ = fi25
                        _ = v26
                    }
                default:
                    v1 = v1
                }
                _ = fi18
                _ = v19
            }
            v27 := GenS3{F0: -22}
            _ = v27
        }
    } else {
        for fi28 := 0; fi28 < 18; fi28++ {
            fmt.Println(fi28, fi28)
            _ = fi28
        }
    }
    var v29 string = "foo bar"
    for fi30 := 0; fi30 < 17; fi30++ {
        var v31 uint32 = 42
        var v32 float32 = (0.0 * 1.5)
        _ = fi30
        _ = v31
        _ = v32
    }
    v33 := GenS1{F0: 18446744073709551615}
    if (161 < 41) {
        var v34 string = v29
        _ = v34
    }
    if true {
        v35 := GenS4{F0: 255, F1: 2147483647}
        var v36 float64 = 0.0
        _ = v35
        _ = v36
    }
    if ((!(!((genFn1(-1, 65535) > -877) && true))) || false) {
        v37 := []int{4294967295, -541, 42}
        v38 := GenS4{F0: 0, F1: 0}
        if (79 >= 15) {
            if true {
                var v39 bool = false
                for fi40 := 0; fi40 < 6; fi40++ {
                    if (!false) {
                        var v41 uint = (v15 + (v15 * v15))
                        switch v29 {
                        case "hello":
                            switch fi40 {
                            case 294:
                                fmt.Println(v15, v41, v15)
                            case -2147483648:
                                v41 -= (v15 ^ (v41 | (^(v15))))
                            default:
                                if len(v37) >= 2 {
                                    v37 = v37[1:len(v37)]
                                }
                            }
                        default:
                            if (v39 || ((0 <= 592) && ("" >= v29))) {
                                if (v39 || ((!v39) && (131 >= 23))) {
                                    v42 := GenS3{F0: -1}
                                    if len(v37) >= 2 {
                                        v37 = v37[1:len(v37)]
                                    }
                                    _ = v42
                                } else {
                                    v37 = append(v37, fi40)
                                    _ = len(v37)
                                }
                            } else {
                                v43 := GenS1{F0: (v1 + v1)}
                                _ = v43
                            }
                        }
                        _ = v41
                    } else {
                        fmt.Println(v39)
                        var v44 int8 = -1
                        _ = v44
                    }
                    v15 += v15
                    _ = fi40
                }
                _ = v39
            } else {
                for fi45 := 0; fi45 < 7; fi45++ {
                    v3[(("foo bar" + ((v29 + v29) + (v29 + ("" + v29)))) + v29)] = (fi45 / ((fi45) | 1))
                    if (v29 > v29) {
                        v2[2147483648] = 42
                        v46 := GenS2{F0: ((0 < (v1 << ((44) & 63))) && ((-693 >= -953) && ((1 <= 1) && true))), F1: (fi45 | fi45)}
                        switch fi45 {
                        case 127:
                            v1 = genFn0((!(!false)))
                        }
                        _ = v46
                    }
                    fmt.Println(v15, v1)
                    _ = fi45
                }
                v29 = v29
            }
            var v47 uint = v15
            var v48 uint32 = 415
            _ = v47
            _ = v48
        } else {
            switch (((v29 + v29) + ((v29 + (("foo bar" + "foo bar") + "hello")) + v29)) + "\n\t") {
            case "a":
                if true {
                    var v49 int32 = -961
                    var v50 int64 = -9223372036854775808
                    if (!true) {
                        switch 32767 {
                        case -9223372036854775808:
                            var v51 int = genFn1(127, 32767)
                            _ = v51
                        }
                        v15 = (v15 >> ((v15) & 63))
                    } else {
                        var v52 float32 = -2.25
                        switch (v29 + v29) {
                        case "\n\t":
                            fmt.Println(v1, v52)
                        default:
                            switch 127 {
                            case 127:
                                switch 307 {
                                case 968:
                                    var v53 float32 = (v52 / 0.1)
                                    _ = v53
                                case 0:
                                    v29 = v29
                                default:
                                    fmt.Println(v1, v29)
                                }
                            }
                        }
                        _ = len(v2)
                        _ = v52
                    }
                    _ = v49
                    _ = v50
                }
            default:
                v15 += v15
            }
            v54 := GenS5{F0: -128, F1: -565, F2: v15}
            if true {
                v15 = 18446744073709551615
                v29 = v29
                var v55 int = genFn1(126, 2)
                _ = v55
            } else {
                var v56 uint16 = 875
                if true {
                    v57 := GenS6{F0: (("\n\t" + v29) + "\n\t"), F1: 127, F2: 42}
                    _ = v57
                }
                fmt.Println(v56, v56)
                _ = v56
            }
            _ = v54
        }
        _ = v37
        _ = v38
    } else {
        var v58 int32 = 42
        _ = v58
    }
    v59 := []int{32767, 78}
    v60 := GenS1{F0: (v1 + (v1 % ((v1) | 1)))}
    var v61 string = v29
    v15 = v15
    for fi62 := 0; fi62 < 18; fi62++ {
        if (false && (false || true)) {
            var v63 int64 = -1
            _ = v63
        } else {
            v59 = append(v59, 32767)
            v64 := GenS1{F0: genFn0(((false || (!false)) && ((!(false && (1 == 32767))) && (!(((3.14159 <= 3.14159) || true) && (255 == 2))))))}
            _ = v64
        }
        _ = fi62
    }
    v61 = (v61 + ((v61 + ((("foo bar" + ("a" + "hello")) + v29) + v29)) + "\n\t"))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v59:", len(v59), v59)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v29
    _ = v3
    _ = v33
    _ = v59
    _ = v60
    _ = v61
}
