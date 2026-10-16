package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8; F1 int16; F2 int32 }
type GenS4 struct { F0 int32; F1 int64; F2 uint }
type GenS5 struct { F0 uint16; F1 int16; F2 uint16 }
type GenS6 struct { F0 uint8; F1 int64 }
func genFn0(p0 int) uint {
    var v7 int64 = 42
    fmt.Println(v7, v7)
    p0 = (p0 - p0)
    for fi8 := 0; fi8 < 7; fi8++ {
        v9 := []int{(^(p0)), -350, 128, (-(p0))}
        if (2 < 130) {
            v9 = append(v9, (p0 % (((-(fi8))) | 1)))
            var v10 int64 = 42
            var v11 int = (fi8 ^ -5)
            _ = v10
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = v7
    return 65535
}

func main() {
    v1 := []uint32{}
    var v2 uint16 = 0
    v3 := make(map[string]int)
    var v4 int8 = -4
    var v5 bool = true
    var v6 int = 128
    if (float64(float32(-2.25)) != 0.0) {
        for fi12 := 0; fi12 < 6; fi12++ {
            var v13 uint8 = 42
            switch ("foo bar" + "\n\t") {
            case "":
                switch ("a" + "hello") {
                case "\n\t":
                    var v14 bool = v5
                    _ = v14
                case "foo bar":
                    v4 -= v4
                }
            case "\n\t":
                for fi15 := 0; fi15 < 19; fi15++ {
                    v4 = (v4 + -61)
                    fmt.Println(v13)
                    if (!v5) {
                        var v16 int64 = 42
                        v17 := GenS1{F0: int32(v2), F1: 1, F2: 0}
                        _ = v16
                        _ = v17
                    } else {
                        v18 := GenS1{F0: 23, F1: 0, F2: 1000}
                        if ((v4 + v4) <= (v4 * -128)) {
                            v19 := GenS2{F0: -9223372036854775808}
                            _ = v19
                        } else {
                            if (!((^(v4)) == 127)) {
                                v13 -= v13
                                var v20 uint8 = 212
                                _ = v20
                            } else {
                                v21 := GenS1{F0: 498, F1: 664, F2: 1000}
                                var v22 uint8 = (v13 * v13)
                                switch (v6 - fi15) {
                                case 1:
                                    var v23 string = "foo bar"
                                    _ = v23
                                case -9223372036854775808:
                                    switch fi15 {
                                    case 128:
                                        var v24 uint64 = 995
                                        _ = v24
                                    case -595:
                                        if (2 >= 1) {
                                            var v25 uint16 = v2
                                            v4 -= (-(v4))
                                            v25 = (v25 + (v2 & 0))
                                            _ = v25
                                        } else {
                                            var v26 int = 32767
                                            _ = v26
                                        }
                                    }
                                default:
                                    v5 = ((!false) && v5)
                                }
                                var v27 int = -618
                                _ = v21
                                _ = v22
                                _ = v27
                            }
                        }
                        switch "\n\t" {
                        case "hello":
                            var v28 int8 = 100
                            _ = v28
                        case "\n\t":
                            var v29 float32 = 1.5
                            _ = v29
                        case "a":
                            v30 := GenS3{F0: (v13 / ((255) | 1)), F1: 32767, F2: -2147483648}
                            _ = v30
                        }
                        _ = v18
                    }
                    var v31 uint8 = 42
                    v2 = (v2 & (v2 ^ v2))
                    _ = fi15
                    _ = v31
                }
            default:
                if v5 {
                    v5 = (v5 || v5)
                    for fi32 := 0; fi32 < 13; fi32++ {
                        if v5 {
                            var v33 int32 = -2147483648
                            fmt.Println(v33)
                            v13 = v13
                            _ = v33
                        }
                        var v34 string = "a"
                        _ = fi32
                        _ = v34
                    }
                } else {
                    var v35 bool = (9223372036854775807 >= 42)
                    fmt.Println(v35, v35)
                    v35 = (genFn0(fi12) >= 0)
                    _ = v35
                }
            }
            _ = fi12
            _ = v13
        }
        v36 := v3["a"]
        _ = v36
    }
    if (-9223372036854775808 < 9223372036854775807) {
        v37 := GenS4{F0: -1, F1: int64(v6), F2: 2}
        v2 = v2
        _ = v37
    } else {
        fmt.Println(v6, v4, v2)
        var v38 uint8 = 128
        _ = v38
    }
    v6 += 128
    if v5 {
        if false {
            if ((v4 & v4) == v4) {
                v39 := GenS4{F0: -680, F1: -9223372036854775808, F2: 18446744073709551615}
                switch -764 {
                case -1:
                    fmt.Println(v4, v2)
                case -364:
                    switch "a" {
                    case "hello":
                        v40 := GenS5{F0: (^(v2)), F1: -1, F2: 0}
                        _ = v40
                    default:
                        v2 += v2
                    }
                default:
                    if (!(((-160 == 9223372036854775807) || v5) && (-1 != -9223372036854775808))) {
                        switch -9223372036854775808 {
                        case 42:
                            if true {
                                var v41 uint64 = 2
                                _ = len(v3)
                                v42 := GenS6{F0: 214, F1: -9223372036854775808}
                                _ = v41
                                _ = v42
                            } else {
                                var v43 uint32 = 35
                                if false {
                                    if v5 {
                                        v1 = append(v1, v43)
                                    }
                                } else {
                                    v43 = 2
                                    v43 = (v43 << ((42) & 63))
                                    var v44 uint32 = (v43 - 2)
                                    _ = v44
                                }
                                _ = v43
                            }
                        case -337:
                            fmt.Println(v2)
                        case -529:
                            v4 = -128
                        }
                        var v45 uint64 = 2
                        for fi46 := 0; fi46 < 10; fi46++ {
                            if v5 {
                                var v47 string = "hello"
                                var v48 uint8 = 244
                                _ = v47
                                _ = v48
                            }
                            v4 = v4
                            _ = fi46
                        }
                        _ = v45
                    } else {
                        fmt.Println(v6, v4)
                    }
                }
                _ = len(v1)
                _ = v39
            }
            switch ("\n\t" + "hello") {
            case "hello":
                v5 = true
            }
        }
        for fi49 := 0; fi49 < 10; fi49++ {
            v4 = (v4 << ((genFn0((v6 / (((v6 >> ((255) & 63))) | 1)))) & 63))
            var v50 int16 = 1000
            switch fi49 {
            case -9223372036854775808:
                var v51 uint8 = 42
                _ = v51
            case 9223372036854775807:
                v5 = v5
            case -1:
                var v52 uint32 = 2147483648
                _ = v52
            }
            _ = fi49
            _ = v50
        }
        v5 = (v6 < v6)
    }
    if true {
        if ((-9223372036854775808 > -9223372036854775808) || true) {
            v6 -= (v6 / ((32768) | 1))
        }
    }
    v53 := GenS1{F0: -2147483648, F1: 594, F2: 32767}
    var v54 int16 = 1
    v55 := GenS3{F0: 3, F1: -349, F2: -1}
    _ = len(v3)
    var v56 uint16 = (v2 % ((732) | 1))
    switch v6 {
    case 304:
        v57 := v3[""]
        _ = v57
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v55
    _ = v56
    _ = v6
}
