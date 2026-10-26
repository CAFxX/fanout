package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint; F1 int64; F2 bool }
type GenS6 struct { F0 uint32; F1 uint32; F2 int }
type GenS7 struct { F0 int32; F1 bool }
func genFn0(p0 bool, p1 string) uint64 {
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 int = (fi7 / ((-1) | 1))
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 7; fi9++ {
        if p0 {
            if p0 {
                v10 := []int{-377}
                for fi11 := 0; fi11 < 3; fi11++ {
                    if p0 {
                        v12 := GenS1{F0: 42}
                        p1 = p1
                        _ = len(v10)
                        p1 = (p1 + p1)
                        _ = v12
                    } else {
                        if len(v10) >= 2 {
                            v10 = v10[1:len(v10)]
                        }
                        var v13 float64 = 1e10
                        v14 := GenS1{F0: 2}
                        _ = v13
                        _ = v14
                    }
                    if false {
                        if (false && p0) {
                            if (!p0) {
                                switch ("foo bar" + p1) {
                                case "\n\t":
                                    p0 = (42 <= 142)
                                case "foo bar":
                                    if len(v10) >= 2 {
                                        v10 = v10[1:len(v10)]
                                    }
                                }
                                var v15 float64 = 1e10
                                _ = v15
                            } else {
                                if len(v10) > 0 {
                                    v10[0] = (fi9 + 42)
                                }
                            }
                            if (1000 == 1000) {
                                var v16 uint32 = 906
                                p1 = (p1 + ("a" + p1))
                                if ("" <= "foo bar") {
                                    var v17 int = fi9
                                    _ = v17
                                } else {
                                    v18 := GenS2{F0: 1000}
                                    _ = v18
                                }
                                _ = v16
                            } else {
                                var v19 float32 = float32(1e10)
                                v20 := GenS3{F0: -9223372036854775808, F1: 20}
                                _ = v19
                                _ = v20
                            }
                            if true {
                                v10 = append(v10, (fi9 % ((fi11) | 1)))
                            } else {
                                p0 = false
                            }
                        }
                        p0 = p0
                    }
                    _ = fi11
                }
                fmt.Println(p1, fi9, fi9)
                _ = v10
            }
            for fi21 := 0; fi21 < 20; fi21++ {
                var v22 int = (-(fi21))
                switch "foo bar" {
                case "":
                    fmt.Println(fi9, fi9, p1)
                case "hello":
                    switch (v22 - 32767) {
                    case -946:
                        v23 := GenS2{F0: 660}
                        _ = v23
                    case 1:
                        v24 := GenS3{F0: 0, F1: 2}
                        _ = v24
                    }
                }
                var v25 int = fi9
                v26 := []int{fi21, v25}
                if (2 < 42) {
                    fmt.Println(p0, v22, v25)
                } else {
                    var v27 uint64 = 2
                    var v28 uint8 = 190
                    _ = v27
                    _ = v28
                }
                _ = fi21
                _ = v22
                _ = v25
                _ = v26
            }
            p1 = p1
        }
        v29 := GenS2{F0: -292}
        p0 = (-1 < -9223372036854775808)
        _ = fi9
        _ = v29
    }
    fmt.Println(p0, p1)
    fmt.Println(p1, p0)
    fmt.Println(p1)
    _ = p0
    _ = p1
    return 964
}
func genFn1(p0 int8) int32 {
    var v30 uint8 = 2
    v30 += (^(v30))
    _ = p0
    _ = v30
    return 2147483647
}

func main() {
    var v1 uint32 = 42
    var v2 int32 = 42
    v3 := make(map[int]string)
    var v4 int = 4294967295
    var v5 int64 = 144
    v6 := []uint32{v1, 1, v1, v1}
    v6 = append(v6, v1)
    for fi31 := 0; fi31 < 17; fi31++ {
        for fi32 := 0; fi32 < 2; fi32++ {
            v1 -= (v1 | (v1 / ((v1) | 1)))
            _ = fi32
        }
        var v33 int16 = 32767
        _ = fi31
        _ = v33
    }
    _ = len(v6)
    var v34 bool = (true && (18446744073709551615 == 42))
    if (!false) {
        v35 := GenS3{F0: v5, F1: 200}
        v36 := GenS2{F0: 1000}
        _ = v35
        _ = v36
    } else {
        v37 := GenS4{F0: 0}
        if v34 {
            delete(v3, (v4 / (((v4 | v4)) | 1)))
            fmt.Println(v34, v2)
            v38 := GenS5{F0: 4294967295, F1: v5, F2: true}
            _ = v38
        } else {
            if len(v6) > 0 {
                v6[0] = v1
            }
            v4 += (v4 + (v4 ^ (v4 | v4)))
            var v39 int16 = 1000
            _ = v39
        }
        _ = v37
    }
    switch "hello" {
    case "\n\t":
        var v40 int32 = 0
        _ = v40
    case "":
        if v34 {
            v41 := GenS5{F0: 2, F1: (v5 ^ 0), F2: true}
            _ = v41
        } else {
            var v42 int16 = 299
            _ = len(v3)
            _ = v42
        }
    case "hello":
        v43 := GenS4{F0: 100}
        _ = v43
    default:
        if ((v1 / (((v1 & (v1 | (v1 & (v1 % ((4294967295) | 1)))))) | 1)) <= 224) {
            var v44 int8 = 100
            _ = v44
        }
    }
    for fi45 := 0; fi45 < 9; fi45++ {
        if (!false) {
            switch "" {
            case "\n\t":
                v3[fi45] = "foo bar"
            case "hello":
                v34 = (3.14159 >= float32((1e10 / 1e10)))
            case "a":
                v46 := GenS1{F0: genFn0((!false), "a")}
                _ = v46
            }
        } else {
            v5 = 373
            if ((fi45 - fi45) > v4) {
                fmt.Println(v1, v2, v34)
                v47 := GenS6{F0: v1, F1: v1, F2: v4}
                _ = v47
            }
        }
        for fi48 := 0; fi48 < 3; fi48++ {
            if v34 {
                var v49 int64 = v5
                _ = v49
            } else {
                v50 := GenS7{F0: (-(v2)), F1: (!v34)}
                _ = len(v6)
                _ = v50
            }
            v51 := GenS4{F0: -128}
            _ = fi48
            _ = v51
        }
        if (595 < 18446744073709551615) {
            var v52 int16 = 42
            v1 = v1
            v53 := GenS6{F0: 65535, F1: v1, F2: (v4 - fi45)}
            _ = v52
            _ = v53
        } else {
            var v54 uint64 = genFn0(v34, "\n\t")
            _ = v54
        }
        _ = fi45
    }
    v2 = -11
    switch (v4 % ((9223372036854775807) | 1)) {
    case 127:
        fmt.Println(v5)
    default:
        fmt.Println(v34, v34)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
