package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 uint64; F1 int64; F2 uint64 }
type GenS5 struct { F0 int16; F1 uint; F2 int }

func main() {
    var v1 int = -9223372036854775808
    var v2 float64 = 0.0
    v3 := make(map[int]int)
    var v4 int32 = -2147483648
    _ = len(v3)
    v5 := GenS1{F0: 827}
    if ((!((0 != 0) || true)) || (((-801 <= 42) && false) || (!false))) {
        v1 -= int(v4)
    }
    if (!true) {
        v2 -= (-2.25 / ((v2 + (v2 + float64(float32((v2 * 3.14159))))) + (1.5 * v2)))
        v1 = v1
        fmt.Println(v1)
    } else {
        v6 := GenS2{F0: 2}
        for fi7 := 0; fi7 < 12; fi7++ {
            v8 := []int{}
            _ = fi7
            _ = v8
        }
        var v9 bool = ((1 > -1) || (-1 < 100))
        _ = v6
        _ = v9
    }
    v4 = (v4 - v4)
    switch v1 {
    case 1:
        fmt.Println(v2)
    default:
        v2 = 1e10
    }
    if (-375 == (v4 >> ((2) & 63))) {
        delete(v3, 42)
        v10 := GenS1{F0: 0}
        _ = v10
    }
    var v11 int16 = 42
    _ = len(v3)
    v12 := v3[(v1 ^ (v1 | v1))]
    if (false || false) {
        if false {
            for fi13 := 0; fi13 < 13; fi13++ {
                v4 -= (v4 + (v4 * (v4 << ((18446744073709551615) & 63))))
                if ((true || (true && true)) || true) {
                    for fi14 := 0; fi14 < 18; fi14++ {
                        v15 := GenS3{F0: 4294967295, F1: 182}
                        _ = fi14
                        _ = v15
                    }
                    var v16 int32 = (v4 * (v4 | 42))
                    var v17 string = "\n\t"
                    _ = v16
                    _ = v17
                } else {
                    var v18 int = (^(v1))
                    switch (("a" + ("foo bar" + ("" + (("foo bar" + "") + "foo bar")))) + "\n\t") {
                    case "foo bar":
                        if (-9223372036854775808 >= 42) {
                            fmt.Println(v11)
                            var v19 int16 = -32768
                            var v20 uint32 = 612
                            _ = v19
                            _ = v20
                        } else {
                            var v21 float32 = 1e10
                            for fi22 := 0; fi22 < 13; fi22++ {
                                v23 := GenS2{F0: 145}
                                if true {
                                    v24 := []int{9223372036854775807}
                                    v21 += ((v21 * (1.5 * (v21 - v21))) + v21)
                                    _ = v24
                                } else {
                                    v4 += v4
                                }
                                _ = fi22
                                _ = v23
                            }
                            _ = v21
                        }
                    }
                    var v25 string = "hello"
                    _ = v18
                    _ = v25
                }
                _ = fi13
            }
            var v26 uint16 = 378
            if false {
                if ((-(v4)) <= (v4 ^ (v4 >> ((4294967295) & 63)))) {
                    var v27 string = ("hello" + "hello")
                    var v28 string = "\n\t"
                    _ = v27
                    _ = v28
                } else {
                    delete(v3, -44)
                    var v29 uint32 = 189
                    _ = v29
                }
                for fi30 := 0; fi30 < 18; fi30++ {
                    v4 += v4
                    v2 -= v2
                    var v31 int8 = 53
                    _ = fi30
                    _ = v31
                }
                v32 := []int{4294967295, v12, v1}
                _ = v32
            } else {
                v33 := v3[32767]
                v34 := GenS2{F0: v26}
                var v35 int8 = 0
                _ = v33
                _ = v34
                _ = v35
            }
            _ = v26
        } else {
            var v36 uint16 = 2
            _ = v36
        }
        if true {
            v4 += (v4 & 131)
            var v37 uint = 0
            _ = v37
        } else {
            var v38 uint8 = 0
            _ = v38
        }
    }
    for fi39 := 0; fi39 < 13; fi39++ {
        var v40 int8 = 37
        _ = fi39
        _ = v40
    }
    var v41 uint16 = 32767
    v42 := GenS3{F0: 2, F1: 228}
    v43 := v3[(v12 ^ (v1 - -9223372036854775808))]
    for fi44 := 0; fi44 < 1; fi44++ {
        if true {
            var v45 int16 = 0
            _ = v45
        } else {
            var v46 string = "a"
            _ = v46
        }
        v47 := GenS4{F0: 4294967295, F1: 1, F2: 0}
        v43 -= (v12 % ((v43) | 1))
        _ = fi44
        _ = v47
    }
    for fi48 := 0; fi48 < 1; fi48++ {
        v49 := []int{fi48}
        v50 := GenS5{F0: v11, F1: 18446744073709551615, F2: 32767}
        _ = fi48
        _ = v49
        _ = v50
    }
    switch v1 {
    case 32768:
        var v51 uint8 = 5
        _ = v51
    case 648:
        v12 = v1
    case -2147483648:
        var v52 string = "a"
        _ = v52
    default:
        var v53 uint32 = 858
        _ = v53
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
}
