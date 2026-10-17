package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 bool; F2 uint64 }
type GenS4 struct { F0 int32; F1 uint16; F2 uint16 }
type GenS5 struct { F0 bool; F1 int8; F2 int16 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int; F1 int16 }
type GenS8 struct { F0 uint32; F1 string; F2 bool }
func genFn0(p0 int8) uint8 {
    fmt.Println(p0)
    for fi6 := 0; fi6 < 11; fi6++ {
        fmt.Println(fi6, fi6)
        for fi7 := 0; fi7 < 11; fi7++ {
            switch (fi7 % (((fi6 % ((fi6) | 1))) | 1)) {
            case -285:
                fmt.Println(fi6)
                v8 := []int{2147483647, 4294967295}
                _ = v8
            case 2147483647:
                if false {
                    fmt.Println(fi6)
                }
            case 32767:
                var v9 uint16 = 255
                _ = v9
            }
            if (!false) {
                if true {
                    var v10 uint = 2
                    _ = v10
                }
                var v11 int64 = 42
                var v12 int16 = 742
                _ = v11
                _ = v12
            } else {
                v13 := GenS1{F0: 821, F1: 820}
                v14 := GenS1{F0: 2, F1: 255}
                p0 = (^(p0))
                _ = v13
                _ = v14
            }
            var v15 uint8 = 2
            _ = fi7
            _ = v15
        }
        _ = fi6
    }
    _ = p0
    return 1
}
func genFn1(p0 int32, p1 int32, p2 int16) int32 {
    p0 = p1
    p1 += (p0 / ((2147483647) | 1))
    for fi16 := 0; fi16 < 8; fi16++ {
        switch "a" {
        case "":
            v17 := []int{fi16}
            _ = v17
        case "hello":
            for fi18 := 0; fi18 < 10; fi18++ {
                var v19 int = 2147483647
                v20 := []int{fi16, (fi18 - 0), (-(fi16)), v19}
                _ = fi18
                _ = v19
                _ = v20
            }
        default:
            v21 := GenS2{F0: "a"}
            _ = v21
        }
        if (!false) {
            fmt.Println(p2, p1)
        } else {
            switch (^(fi16)) {
            case -961:
                for fi22 := 0; fi22 < 9; fi22++ {
                    fmt.Println(p1, p1)
                    if true {
                        if false {
                            var v23 uint = 254
                            p1 = (p1 ^ p0)
                            _ = v23
                        } else {
                            p1 += 42
                            fmt.Println(p0, fi22, fi16)
                            if true {
                                v24 := GenS3{F0: 127, F1: (!true), F2: 1}
                                v25 := GenS1{F0: 1, F1: 363}
                                _ = v24
                                _ = v25
                            }
                        }
                    } else {
                        v26 := []int{}
                        if len(v26) > 0 {
                            v26[0] = (fi16 << ((1) & 63))
                        }
                        _ = v26
                    }
                    var v27 uint32 = 42
                    p0 = (p0 | (p1 ^ 42))
                    _ = fi22
                    _ = v27
                }
            case 4294967295:
                fmt.Println(p2)
            case 2147483647:
                fmt.Println(p1, p2, p2)
            }
            if ((-(p2)) > (-(p2))) {
                fmt.Println(p2, p1)
                p1 = p0
            } else {
                if false {
                    p0 += (p0 + -649)
                    var v28 int8 = -128
                    _ = v28
                }
                fmt.Println(p1, fi16)
                v29 := GenS4{F0: -1, F1: 809, F2: 2}
                _ = v29
            }
        }
        _ = fi16
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    return p0
}

func main() {
    v1 := []int{}
    v2 := []int{32767, 0, 128}
    v3 := make(map[uint32]string)
    v4 := []uint32{509}
    var v5 float32 = 1.5
    v30 := GenS3{F0: 88, F1: (true && false), F2: 42}
    var v31 string = "hello"
    v31 = ("a" + v31)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi32 := 0; fi32 < 1; fi32++ {
        v33 := GenS5{F0: (!false), F1: -34, F2: 1000}
        if false {
            if (474 != 32767) {
                _ = len(v3)
                _ = len(v3)
            }
        } else {
            v5 -= (1.5 + v5)
            if false {
                if false {
                    v5 = v5
                    for fi34 := 0; fi34 < 12; fi34++ {
                        v35 := GenS1{F0: 1, F1: 255}
                        v36 := GenS6{F0: 32767}
                        v31 = "hello"
                        fmt.Println(v31, v5)
                        v5 += float32(1.5)
                        _ = fi34
                        _ = v35
                        _ = v36
                    }
                    v37 := GenS6{F0: -836}
                    _ = v37
                } else {
                    var v38 float64 = (3.14159 - 1.5)
                    _ = v38
                }
                for fi39 := 0; fi39 < 1; fi39++ {
                    v5 -= (3.14159 - 1e10)
                    v5 += float32(-2.25)
                    v40 := v3[36]
                    if len(v2) > 0 {
                        v2[0] = fi32
                    }
                    if (false || (!true)) {
                        switch ("" + (v31 + (v31 + "\n\t"))) {
                        case "":
                            delete(v3, 1)
                        case "hello":
                            var v41 int32 = 0
                            _ = v41
                        case "a":
                            if false {
                                var v42 int8 = -119
                                if (-936 < 2147483647) {
                                    var v43 int64 = 672
                                    v42 += (v42 ^ v42)
                                    _ = v43
                                }
                                v44 := GenS3{F0: 97, F1: (42 < 42), F2: 2}
                                _ = v42
                                _ = v44
                            }
                        }
                        v40 = "a"
                        var v45 uint64 = 4294967295
                        _ = v45
                    }
                    _ = fi39
                    _ = v40
                }
            } else {
                fmt.Println(v31, v5)
                v46 := GenS7{F0: (fi32 * (-(fi32))), F1: 42}
                _ = v46
            }
        }
        _ = fi32
        _ = v33
    }
    for fi47 := 0; fi47 < 5; fi47++ {
        for fi48 := 0; fi48 < 17; fi48++ {
            fmt.Println(fi47, v5, v31)
            var v49 bool = (254 <= 163)
            if (710 <= 4294967295) {
                if (826 >= 2) {
                    var v50 int = 32768
                    fmt.Println(v31, fi47, v5)
                    var v51 int32 = genFn1(2147483647, -2147483648, 32767)
                    _ = v50
                    _ = v51
                } else {
                    v31 = ((v31 + v31) + v31)
                    var v52 bool = true
                    switch 650 {
                    case 1:
                        var v53 uint64 = 746
                        _ = v53
                    default:
                        v54 := GenS8{F0: 2147483647, F1: v31, F2: (((v5 * 1.5) / float32((1.5 - 0.1))) >= v5)}
                        _ = v54
                    }
                    v31 = v31
                    _ = v52
                }
                fmt.Println(fi47)
                var v55 bool = (v49 || (float64(float32((1e10 / (3.14159 - 1.5)))) != 1.5))
                _ = v55
            }
            var v56 string = "\n\t"
            var v57 uint64 = 2
            _ = fi48
            _ = v49
            _ = v56
            _ = v57
        }
        if len(v2) > 0 {
            v2[0] = 32767
        }
        _ = fi47
    }
    v5 -= -2.25
    v5 -= float32((0.0 / -2.25))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
