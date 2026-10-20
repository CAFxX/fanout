package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 uint64; F1 uint32; F2 uint64 }
type GenS3 struct { F0 uint16; F1 int16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint; F1 uint16 }
type GenS6 struct { F0 int8; F1 int; F2 int8 }
type GenS7 struct { F0 uint64; F1 uint64; F2 int64 }
func genFn0(p0 uint) int {
    for fi4 := 0; fi4 < 4; fi4++ {
        if false {
            p0 = (p0 + 0)
            v5 := GenS1{F0: 65535, F1: 4294967295}
            for fi6 := 0; fi6 < 4; fi6++ {
                v7 := []int{32768}
                _ = fi6
                _ = v7
            }
            _ = v5
        }
        _ = fi4
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 float64 = (0.1 / 1.5)
        for fi10 := 0; fi10 < 15; fi10++ {
            switch (^(fi8)) {
            case 128:
                v9 -= (v9 - -2.25)
            case -1:
                if (false || (false || false)) {
                    if (-1 < -448) {
                        if (86 == 42) {
                            switch fi8 {
                            case 2147483647:
                                fmt.Println(v9)
                            case 42:
                                v11 := GenS2{F0: 18446744073709551615, F1: 2, F2: 2}
                                _ = v11
                            default:
                                v12 := []int{}
                                _ = v12
                            }
                            var v13 bool = false
                            var v14 bool = v13
                            switch ("foo bar" + "hello") {
                            case "\n\t":
                                switch fi8 {
                                case 32768:
                                    if v13 {
                                        v13 = false
                                        fmt.Println(fi8, v13, v13)
                                        var v15 int8 = -1
                                        _ = v15
                                    }
                                case -9223372036854775808:
                                    v13 = (!(("" + ("foo bar" + (("foo bar" + ("foo bar" + ("a" + "\n\t"))) + "a"))) > "\n\t"))
                                default:
                                    v16 := GenS1{F0: 0, F1: 93}
                                    _ = v16
                                }
                            default:
                                v17 := []int{-1, 32767, 127}
                                _ = v17
                            }
                            _ = v13
                            _ = v14
                        } else {
                            var v18 string = "\n\t"
                            v9 = float64(float32(v9))
                            v18 = "foo bar"
                            _ = v18
                        }
                        p0 += p0
                    } else {
                        var v19 int8 = -128
                        switch ("a" + ("hello" + ("\n\t" + "foo bar"))) {
                        case "":
                            v20 := GenS3{F0: 42, F1: 1000}
                            _ = v20
                        case "\n\t":
                            switch ((("hello" + "foo bar") + "") + ("a" + ("hello" + "a"))) {
                            case "":
                                if (!(false && (("\n\t" == ("a" + "\n\t")) && (p0 > p0)))) {
                                    var v21 int8 = v19
                                    v19 += (v21 / (((^(v21))) | 1))
                                    var v22 uint16 = 32767
                                    _ = v21
                                    _ = v22
                                } else {
                                    p0 += p0
                                    if (65535 != 18446744073709551615) {
                                        v9 -= v9
                                        var v23 string = ""
                                        fmt.Println(v19)
                                        _ = v23
                                    }
                                }
                            case "foo bar":
                                var v24 float64 = 0.0
                                _ = v24
                            default:
                                if false {
                                    var v25 int8 = v19
                                    _ = v25
                                } else {
                                    v26 := GenS1{F0: 32767, F1: 2}
                                    v27 := GenS2{F0: uint64(p0), F1: 2, F2: 42}
                                    _ = v26
                                    _ = v27
                                }
                            }
                        }
                        _ = v19
                    }
                }
                var v28 float64 = v9
                var v29 uint = (p0 ^ p0)
                _ = v28
                _ = v29
            case 32767:
                var v30 bool = (!((-32768 > -172) || (!(-2147483648 >= 2147483647))))
                _ = v30
            default:
                v9 -= float64((3.14159 - -2.25))
            }
            _ = fi10
        }
        v31 := []int{(fi8 << (((p0 / ((p0) | 1))) & 63)), (fi8 ^ (-(fi8))), 9223372036854775807, 622}
        _ = fi8
        _ = v31
        _ = v9
    }
    fmt.Println(p0)
    p0 = p0
    _ = p0
    return -1
}

func main() {
    var v1 uint8 = 127
    v2 := make(map[int]bool)
    v3 := []uint32{69, 233, 2, 2}
    v1 += v1
    for fi32 := 0; fi32 < 1; fi32++ {
        var v33 int16 = -1
        for fi34 := 0; fi34 < 17; fi34++ {
            v35 := GenS4{F0: -303}
            fmt.Println(fi32, v33, fi32)
            switch (("\n\t" + "foo bar") + "foo bar") {
            case "foo bar":
                v36 := GenS5{F0: 4294967295, F1: 1}
                _ = v36
            case "":
                v33 = v33
            case "hello":
                switch genFn0(407) {
                case 32767:
                    v37 := GenS6{F0: -128, F1: (fi34 % ((fi32) | 1)), F2: -128}
                    _ = v37
                }
            }
            _ = fi34
            _ = v35
        }
        var v38 float32 = float32((1.5 + (1e10 - 3.14159)))
        _ = fi32
        _ = v33
        _ = v38
    }
    v1 = (v1 ^ (v1 - (v1 | 128)))
    v1 = (v1 ^ (v1 / (((v1 + (v1 % (((v1 >> ((4294967295) & 63))) | 1)))) | 1)))
    v1 -= (v1 - (v1 >> ((65535) & 63)))
    var v39 uint = 0
    v40 := GenS4{F0: -1}
    var v41 float32 = (-2.25 - ((1e10 * -2.25) * 0.1))
    fmt.Println(v41)
    v1 += (v1 << (((v39 - 1)) & 63))
    v42 := GenS7{F0: 2, F1: 18446744073709551615, F2: 9223372036854775807}
    v39 = 838
    for fi43 := 0; fi43 < 17; fi43++ {
        v41 -= v41
        _ = fi43
    }
    v39 -= (v39 + 42)
    fmt.Println(v41, v1)
    var v44 float64 = -2.25
    if (false && (-1 == -43)) {
        switch 871 {
        case 671:
            v41 = 0.0
        default:
            v45 := GenS4{F0: 0}
            _ = v45
        }
    }
    fmt.Println(v39)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v41:", v41)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v40
    _ = v41
    _ = v42
    _ = v44
}
