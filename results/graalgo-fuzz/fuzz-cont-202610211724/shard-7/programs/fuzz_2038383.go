package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64; F1 uint; F2 int16 }
type GenS4 struct { F0 uint32 }
func genFn0(p0 string) int8 {
    fmt.Println(p0)
    var v4 int8 = 42
    var v5 uint32 = 1
    switch (("a" + (p0 + (p0 + "foo bar"))) + (p0 + ("\n\t" + ("a" + ("a" + "a"))))) {
    case "a":
        var v6 uint16 = 42
        _ = v6
    case "":
        fmt.Println(v5, v4, v4)
    case "hello":
        var v7 uint16 = 486
        _ = v7
    }
    _ = p0
    _ = v4
    _ = v5
    return 100
}
func genFn1(p0 int8) uint64 {
    v8 := []int{2147483647}
    p0 -= p0
    if true {
        if (36 >= 17) {
            v9 := GenS1{F0: 839, F1: -1, F2: -1}
            p0 = p0
            _ = v9
        } else {
            if (true && true) {
                v10 := GenS1{F0: -173, F1: 42, F2: 127}
                var v11 bool = true
                if (-983 >= 379) {
                    if len(v8) > 0 {
                        v8[0] = 1
                    }
                    var v12 int = 630
                    _ = v12
                }
                _ = v10
                _ = v11
            }
            var v13 uint8 = 128
            var v14 int16 = 1
            _ = v13
            _ = v14
        }
        v8 = append(v8, -112)
    }
    p0 = (p0 * (-(p0)))
    v15 := GenS1{F0: 42, F1: 1000, F2: 4294967295}
    _ = p0
    _ = v15
    _ = v8
    return 99
}

func main() {
    var v1 uint32 = 713
    v2 := make(map[int]string)
    var v3 int32 = int32(v1)
    for fi16 := 0; fi16 < 20; fi16++ {
        if false {
            v3 = v3
        }
        v17 := []int{}
        _ = fi16
        _ = v17
    }
    var v18 uint8 = 2
    v1 -= (v1 - (v1 - (^(v1))))
    fmt.Println(v3)
    v1 += (v1 | (v1 >> ((65535) & 63)))
    switch -1 {
    case 32768:
        for fi19 := 0; fi19 < 6; fi19++ {
            for fi20 := 0; fi20 < 4; fi20++ {
                var v21 uint32 = v1
                if ((42 > 1) || (false || true)) {
                    var v22 uint = 571
                    switch fi20 {
                    case -1:
                        v3 += v3
                    case -2147483648:
                        var v23 uint16 = 32767
                        v2[(fi20 / (((^(fi20))) | 1))] = (("foo bar" + "a") + "foo bar")
                        _ = v23
                    default:
                        if (!true) {
                            if (2 != 736) {
                                v22 += (v22 | v22)
                                switch "a" {
                                case "foo bar":
                                    if (!false) {
                                        v24 := GenS1{F0: 977, F1: 0, F2: -551}
                                        v25 := []int{}
                                        v3 -= (^(v3))
                                        _ = v24
                                        _ = v25
                                    } else {
                                        switch fi20 {
                                        case 128:
                                            v3 -= (v3 | v3)
                                        case 32768:
                                            if (830 <= 32767) {
                                                switch "foo bar" {
                                                case "hello":
                                                    v26 := v2[(fi19 + (-(fi19)))]
                                                    _ = v26
                                                case "foo bar":
                                                    v27 := GenS1{F0: -9223372036854775808, F1: -70, F2: 32767}
                                                    _ = v27
                                                case "":
                                                    v18 = 232
                                                }
                                                _ = len(v2)
                                                var v28 uint64 = genFn1(100)
                                                _ = v28
                                            } else {
                                                if ((!false) || ((^(v18)) >= uint8(v22))) {
                                                    fmt.Println(fi19)
                                                    var v29 uint = (v22 * v22)
                                                    v30 := []int{(fi19 & fi20)}
                                                    _ = v29
                                                    _ = v30
                                                }
                                                v21 += (v1 >> (((v22 + (^(v22)))) & 63))
                                                v31 := GenS2{F0: (-(fi19))}
                                                _ = v31
                                            }
                                        case 32767:
                                            _ = len(v2)
                                        default:
                                            var v32 string = ""
                                            _ = v32
                                        }
                                    }
                                case "":
                                    var v33 int32 = 0
                                    _ = v33
                                case "hello":
                                    v22 = (v22 << (((v22 << (((v22 ^ (v22 * 302))) & 63))) & 63))
                                }
                            } else {
                                v34 := []int{179, fi20, -1, fi20}
                                var v35 uint = 48
                                if ((v3 / (((v3 | (v3 / (((^(v3))) | 1)))) | 1)) <= v3) {
                                    v36 := GenS3{F0: -1, F1: (v35 % ((v22) | 1)), F2: -335}
                                    switch "\n\t" {
                                    case "\n\t":
                                        switch -37 {
                                        case 42:
                                            v3 = -305
                                        }
                                    case "":
                                        v18 = (v18 >> ((v22) & 63))
                                    case "a":
                                        v37 := GenS3{F0: 941, F1: v22, F2: 42}
                                        _ = v37
                                    default:
                                        v38 := GenS4{F0: (v1 % (((v21 >> (((^(v35))) & 63))) | 1))}
                                        _ = v38
                                    }
                                    _ = v36
                                }
                                _ = v34
                                _ = v35
                            }
                        }
                    }
                    if ((1e10 < 3.14159) || (1 == -489)) {
                        v39 := GenS3{F0: -453, F1: (v22 & v22), F2: 219}
                        var v40 string = "a"
                        if false {
                            v41 := GenS3{F0: -874, F1: (v22 / (((v22 - v22)) | 1)), F2: 42}
                            v1 += (v21 * (^(v1)))
                            _ = v41
                        } else {
                            var v42 int16 = 32767
                            _ = v42
                        }
                        v43 := []int{4294967295}
                        if true {
                            v3 -= v3
                            v44 := GenS3{F0: -1, F1: uint(v18), F2: -580}
                            _ = v44
                        } else {
                            v3 = v3
                            var v45 int16 = -132
                            _ = v45
                        }
                        fmt.Println(fi19, v18, v22)
                        var v46 uint = 1
                        _ = v39
                        _ = v40
                        _ = v43
                        _ = v46
                    }
                    v47 := []int{0, 32768, fi19, (fi20 - 2147483647)}
                    fmt.Println(v22)
                    _ = v22
                    _ = v47
                } else {
                    var v48 uint = 878
                    var v49 uint = (^(v48))
                    _ = v48
                    _ = v49
                }
                v50 := []int{2147483647, 0}
                _ = fi20
                _ = v21
                _ = v50
            }
            var v51 float32 = 3.14159
            _ = fi19
            _ = v51
        }
    }
    fmt.Println(v1, v1, v1)
    for fi52 := 0; fi52 < 19; fi52++ {
        v2[fi52] = (("hello" + ("foo bar" + "")) + "")
        v53 := []int{(fi52 * (fi52 ^ (-(fi52)))), 4294967295, fi52}
        if (!false) {
            for fi54 := 0; fi54 < 14; fi54++ {
                v1 = (v1 >> ((2) & 63))
                fmt.Println(v1, v3, v18)
                var v55 uint8 = (v18 * v18)
                _ = fi54
                _ = v55
            }
            fmt.Println(v3, fi52)
            _ = len(v2)
        }
        _ = fi52
        _ = v53
    }
    v18 = (v18 / (((v18 % ((v18) | 1))) | 1))
    var v56 uint64 = 459
    if (1 >= 42) {
        v57 := []int{}
        _ = v57
    } else {
        var v58 int64 = 42
        v58 = (-(v58))
        _ = v58
    }
    fmt.Println(v56, v1, v56)
    v59 := GenS1{F0: 174, F1: -1, F2: 42}
    fmt.Println(v3, v3, v18)
    fmt.Println(v1, v56)
    delete(v2, 4294967295)
    var v60 uint = 0
    var v61 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v56:", v56)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v56
    _ = v59
    _ = v60
    _ = v61
}
