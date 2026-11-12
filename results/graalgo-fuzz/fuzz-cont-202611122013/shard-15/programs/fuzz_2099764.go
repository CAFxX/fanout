package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint8 }
type GenS2 struct { F0 bool; F1 uint16; F2 uint }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int8; F1 string; F2 uint64 }
func genFn0() int64 {
    v4 := GenS1{F0: 139, F1: 350, F2: 61}
    if (!(261 >= 932)) {
        if false {
            var v5 float32 = 0.0
            v5 = v5
            v6 := []int{32767, 2147483647}
            _ = v5
            _ = v6
        } else {
            if false {
                var v7 uint16 = 32767
                var v8 int32 = 0
                for fi9 := 0; fi9 < 10; fi9++ {
                    var v10 int16 = -32768
                    _ = fi9
                    _ = v10
                }
                _ = v7
                _ = v8
            } else {
                var v11 uint8 = 2
                for fi12 := 0; fi12 < 14; fi12++ {
                    if (!true) {
                        v11 = 42
                        if true {
                            v11 = 127
                            for fi13 := 0; fi13 < 10; fi13++ {
                                if (!false) {
                                    v14 := GenS2{F0: ((!true) && (40 == int32(fi12))), F1: 42, F2: 2}
                                    _ = v14
                                } else {
                                    v11 += v11
                                    if false {
                                        var v15 int32 = -827
                                        v16 := GenS2{F0: ((-(v15)) >= -2147483648), F1: 872, F2: 65535}
                                        v15 = (v15 + v15)
                                        _ = v15
                                        _ = v16
                                    }
                                }
                                v11 = (v11 + v11)
                                switch (fi12 - fi12) {
                                case -114:
                                    if (false && (!false)) {
                                        switch 593 {
                                        case -317:
                                            v17 := []int{-9223372036854775808, 127, (fi12 + -357), fi12}
                                            _ = v17
                                        case -2147483648:
                                            v11 -= v11
                                        default:
                                            switch fi13 {
                                            case -212:
                                                var v18 uint64 = 42
                                                _ = v18
                                            case 233:
                                                v11 += v11
                                            default:
                                                v19 := []int{fi13, (-(fi13)), (fi13 + (fi13 * fi13))}
                                                _ = v19
                                            }
                                        }
                                    }
                                    var v20 int8 = -1
                                    _ = v20
                                }
                                _ = fi13
                            }
                            for fi21 := 0; fi21 < 13; fi21++ {
                                v22 := GenS1{F0: fi21, F1: 0, F2: (v11 + 0)}
                                v23 := GenS2{F0: true, F1: 42, F2: 65535}
                                v11 = (^(v11))
                                _ = fi21
                                _ = v22
                                _ = v23
                            }
                        } else {
                            if (65535 >= 2) {
                                v11 = 1
                                var v24 int64 = 9223372036854775807
                                var v25 uint = 2
                                _ = v24
                                _ = v25
                            } else {
                                if false {
                                    v26 := GenS3{F0: 100}
                                    _ = v26
                                }
                                if (1000 != 42) {
                                    var v27 int16 = 1000
                                    v11 -= (v11 - (v11 << ((4294967295) & 63)))
                                    _ = v27
                                }
                                v11 -= v11
                            }
                            v28 := GenS4{F0: fi12}
                            for fi29 := 0; fi29 < 5; fi29++ {
                                v30 := []int{fi12, fi29, (-(fi29)), -9223372036854775808}
                                _ = fi29
                                _ = v30
                            }
                            _ = v28
                        }
                    }
                    var v31 int16 = 35
                    _ = fi12
                    _ = v31
                }
                _ = v11
            }
            for fi32 := 0; fi32 < 9; fi32++ {
                switch (fi32 % (((fi32 | 317)) | 1)) {
                case 127:
                    var v33 int8 = -128
                    _ = v33
                case 672:
                    v34 := []int{fi32}
                    _ = v34
                }
                _ = fi32
            }
        }
    } else {
        var v35 uint64 = 1
        fmt.Println(v35)
        if true {
            var v36 float64 = 0.1
            _ = v36
        }
        _ = v35
    }
    _ = v4
    return 1
}

func main() {
    var v1 float64 = -2.25
    var v2 int16 = 1000
    v3 := make(map[int]bool)
    for fi37 := 0; fi37 < 15; fi37++ {
        _ = len(v3)
        v38 := GenS3{F0: -34}
        var v39 uint16 = 32767
        _ = fi37
        _ = v38
        _ = v39
    }
    var v40 int8 = -128
    v3[-770] = (4294967295 <= 9223372036854775807)
    fmt.Println(v2, v40)
    if (!(-2.25 >= 1e10)) {
        if (1.5 != 3.14159) {
            v41 := GenS1{F0: -2147483648, F1: 630, F2: 255}
            for fi42 := 0; fi42 < 6; fi42++ {
                fmt.Println(v2, v1)
                switch (fi42 - (fi42 >> ((0) & 63))) {
                case 127:
                    v43 := GenS4{F0: (-(fi42))}
                    _ = v43
                }
                switch fi42 {
                case -9223372036854775808:
                    v2 = v2
                case 127:
                    if false {
                        v44 := GenS3{F0: (v40 / ((v40) | 1))}
                        var v45 int64 = 1
                        _ = v44
                        _ = v45
                    }
                }
                _ = fi42
            }
            _ = v41
        }
        v1 += 1e10
        for fi46 := 0; fi46 < 13; fi46++ {
            v47 := GenS5{F0: v40, F1: "", F2: 4294967295}
            switch ("" + "foo bar") {
            case "\n\t":
                for fi48 := 0; fi48 < 8; fi48++ {
                    v49 := GenS5{F0: -82, F1: "", F2: 18446744073709551615}
                    _ = fi48
                    _ = v49
                }
            }
            for fi50 := 0; fi50 < 2; fi50++ {
                v51 := []int{fi46, (^(fi50)), 232}
                _ = fi50
                _ = v51
            }
            _ = fi46
            _ = v47
        }
    }
    var v52 uint = 18446744073709551615
    var v53 uint8 = 128
    fmt.Println(v40)
    var v54 uint64 = 0
    v55 := []int{42, -1}
    switch (("a" + (("\n\t" + ("" + "")) + "")) + "foo bar") {
    case "foo bar":
        _ = len(v55)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v40:", v40)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("slicelen v55:", len(v55), v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v40
    _ = v52
    _ = v53
    _ = v54
    _ = v55
}
