package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool; F1 uint64; F2 string }
type GenS4 struct { F0 uint; F1 bool }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int; F1 uint32 }
func genFn0(p0 bool, p1 uint64) bool {
    v6 := GenS1{F0: -9223372036854775808}
    v7 := GenS2{F0: 2147483647}
    for fi8 := 0; fi8 < 18; fi8++ {
        for fi9 := 0; fi9 < 7; fi9++ {
            var v10 uint = 18446744073709551615
            _ = fi9
            _ = v10
        }
        for fi11 := 0; fi11 < 11; fi11++ {
            var v12 uint16 = 116
            v13 := []int{}
            var v14 bool = (-9223372036854775808 < -1)
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return (-1 >= 128)
}
func genFn1(p0 bool, p1 int, p2 bool) bool {
    if (p0 || p2) {
        switch "a" {
        case "":
            if p0 {
                for fi15 := 0; fi15 < 19; fi15++ {
                    if (!p0) {
                        var v16 float32 = (((((1.5 + 1.5) / 1e10) / -2.25) + -2.25) - 1.5)
                        if true {
                            for fi17 := 0; fi17 < 12; fi17++ {
                                switch (p1 + (p1 - (p1 >> ((65535) & 63)))) {
                                case 353:
                                    v18 := GenS2{F0: 4294967295}
                                    _ = v18
                                case 4294967295:
                                    v19 := []int{}
                                    _ = v19
                                case 2147483647:
                                    v20 := GenS3{F0: (208 != 127), F1: 0, F2: (("\n\t" + ((("foo bar" + ((("a" + ("\n\t" + "a")) + ("foo bar" + "\n\t")) + ((("\n\t" + "hello") + "hello") + "\n\t"))) + "") + "a")) + "hello")}
                                    _ = v20
                                }
                                _ = fi17
                            }
                            v21 := GenS4{F0: 2, F1: (42 != 9223372036854775807)}
                            v22 := GenS4{F0: 18446744073709551615, F1: (!(2 != 886))}
                            _ = v21
                            _ = v22
                        }
                        v23 := []int{766, (fi15 << ((0) & 63)), 0, 4294967295}
                        _ = v16
                        _ = v23
                    }
                    var v24 int16 = -1
                    _ = fi15
                    _ = v24
                }
            } else {
                p2 = false
            }
        case "foo bar":
            fmt.Println(p2, p1, p2)
        case "hello":
            v25 := []int{(p1 + (p1 & 0)), p1}
            _ = v25
        }
        var v26 uint64 = 2
        v27 := GenS4{F0: 1, F1: (!(-2.25 <= -2.25))}
        _ = v26
        _ = v27
    }
    var v28 int8 = 100
    var v29 string = ""
    v30 := []int{(p1 >> ((941) & 63)), (p1 / ((128) | 1)), -677, p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v28
    _ = v29
    _ = v30
    return p0
}

func main() {
    var v1 int32 = 42
    var v2 string = ("hello" + "hello")
    var v3 bool = false
    var v4 bool = false
    var v5 int32 = -2147483648
    v5 += v5
    var v31 int32 = v1
    if false {
        v32 := []int{32768}
        for fi33 := 0; fi33 < 7; fi33++ {
            var v34 int8 = -69
            v31 = (v1 ^ -310)
            _ = fi33
            _ = v34
        }
        _ = v32
    }
    if v4 {
        if v4 {
            var v35 uint8 = 117
            for fi36 := 0; fi36 < 5; fi36++ {
                v37 := GenS1{F0: -9223372036854775808}
                if v3 {
                    v38 := GenS5{F0: 885}
                    v4 = v3
                    _ = v38
                } else {
                    for fi39 := 0; fi39 < 18; fi39++ {
                        v40 := []int{fi36, fi36, fi39}
                        if false {
                            v41 := GenS6{F0: 32768, F1: 117}
                            fmt.Println(v2, v2, v5)
                            if (2 == 32767) {
                                var v42 int8 = 127
                                v43 := GenS2{F0: 1}
                                if v4 {
                                    var v44 float32 = (float32(float64(((3.14159 - ((-2.25 + (0.1 - 0.0)) - (1.5 + (float32((3.14159 / (-2.25 * 3.14159))) - ((1e10 / ((float32(3.14159) + (float32((((0.0 / -2.25) - 0.1) - (float64((((1e10 / 1e10) - (float32((0.1 + 1.5)) * 3.14159)) + -2.25)) * (float64(float32((1e10 * float64((-2.25 * 1.5))))) - 0.0)))) + 3.14159)) - 1.5)) * 0.1))))) * (0.1 / 3.14159)))) * ((((0.0 / float32(-2.25)) / -2.25) - 3.14159) / 1.5))
                                    _ = v44
                                }
                                _ = v42
                                _ = v43
                            } else {
                                var v45 int = 1
                                if (false && (v4 || (!false))) {
                                    v40 = append(v40, fi39)
                                    v46 := GenS6{F0: (v45 - 0), F1: 0}
                                    fmt.Println(v4, v35, v35)
                                    _ = v46
                                }
                                fmt.Println(fi39, v2, v31)
                                _ = v45
                            }
                            _ = v41
                        }
                        _ = fi39
                        _ = v40
                    }
                    v47 := []int{127, fi36, fi36, 4294967295}
                    for fi48 := 0; fi48 < 12; fi48++ {
                        switch "\n\t" {
                        case "hello":
                            if (!(176 != 18446744073709551615)) {
                                v49 := GenS4{F0: 4294967295, F1: (!v4)}
                                switch (v2 + (v2 + (v2 + v2))) {
                                case "\n\t":
                                    if (!(100 >= -46)) {
                                        v50 := GenS6{F0: (fi48 << ((18446744073709551615) & 63)), F1: 2147483647}
                                        v51 := GenS6{F0: fi36, F1: 4294967295}
                                        v5 += v31
                                        _ = v50
                                        _ = v51
                                    }
                                }
                                fmt.Println(fi36)
                                _ = v49
                            }
                        case "":
                            var v52 uint16 = uint16(v35)
                            _ = v52
                        default:
                            v2 = ("" + "a")
                        }
                        var v53 int8 = -128
                        _ = fi48
                        _ = v53
                    }
                    _ = v47
                }
                fmt.Println(v4, v2, v5)
                _ = fi36
                _ = v37
            }
            _ = v35
        }
        fmt.Println(v31)
        if (9223372036854775807 <= 463) {
            v5 -= v1
            if (511 > 397) {
                if v4 {
                    var v54 bool = (v3 && true)
                    _ = v54
                } else {
                    var v55 float64 = 0.1
                    v55 = (v55 / v55)
                    for fi56 := 0; fi56 < 2; fi56++ {
                        v57 := []int{(fi56 + fi56), 987, (fi56 % ((fi56) | 1)), fi56}
                        _ = fi56
                        _ = v57
                    }
                    _ = v55
                }
                v58 := []int{32768, -364}
                v59 := GenS6{F0: 127, F1: 89}
                _ = v58
                _ = v59
            }
        }
    }
    var v60 int32 = 2147483647
    v60 += -2147483648
    v3 = (4294967295 != 1)
    for fi61 := 0; fi61 < 13; fi61++ {
        fmt.Println(v31, v5)
        for fi62 := 0; fi62 < 13; fi62++ {
            var v63 int8 = 127
            v64 := []int{(fi62 << ((2) & 63)), 42, fi62}
            if len(v64) >= 2 {
                v64 = v64[1:len(v64)]
            }
            fmt.Println(v5)
            _ = fi62
            _ = v63
            _ = v64
        }
        _ = fi61
    }
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v60
}
