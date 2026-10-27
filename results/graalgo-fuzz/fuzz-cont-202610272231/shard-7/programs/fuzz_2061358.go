package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint }
type GenS2 struct { F0 int32; F1 uint; F2 uint32 }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int; F1 uint32; F2 int32 }
func genFn0(p0 int8) string {
    switch int(p0) {
    case 2147483647:
        v6 := []int{1}
        _ = v6
    case 303:
        var v7 bool = (1 > 1000)
        _ = v7
    case -760:
        var v8 bool = false
        _ = v8
    default:
        var v9 uint = 454
        _ = v9
    }
    var v10 float64 = -2.25
    if (!false) {
        var v11 string = (("foo bar" + "foo bar") + "")
        for fi12 := 0; fi12 < 2; fi12++ {
            fmt.Println(fi12, p0)
            v11 = ("" + (v11 + v11))
            _ = fi12
        }
        _ = v11
    } else {
        p0 = p0
    }
    fmt.Println(v10, p0)
    fmt.Println(p0, p0)
    _ = p0
    _ = v10
    return "\n\t"
}

func main() {
    v1 := []int{32, 2147483647, 9223372036854775807, 127}
    v2 := []uint32{65535, 4294967295}
    var v3 int64 = -9223372036854775808
    v4 := make(map[uint32]int)
    v5 := make(map[uint32]string)
    if false {
        var v13 bool = (-1 <= 42)
        _ = v13
    } else {
        v14 := GenS1{F0: "", F1: false, F2: 53}
        v3 += v3
        if (v3 > (v3 >> ((322) & 63))) {
            fmt.Println(v3)
            switch 32768 {
            case 128:
                if (763 > 42) {
                    if true {
                        var v15 uint8 = 2
                        var v16 int8 = -128
                        _ = v15
                        _ = v16
                    } else {
                        v2 = append(v2, 487)
                    }
                    var v17 string = ""
                    _ = v17
                }
            case -344:
                v18 := GenS1{F0: "\n\t", F1: (18446744073709551615 > 1), F2: 812}
                _ = v18
            default:
                switch 9223372036854775807 {
                case 237:
                    var v19 string = "\n\t"
                    _ = v19
                case 127:
                    v20 := GenS2{F0: 1, F1: 18446744073709551615, F2: 2}
                    _ = v20
                default:
                    if (!(4294967295 != 529)) {
                        if (true && (2147483647 != 9223372036854775807)) {
                            v3 = (v3 / ((v3) | 1))
                        } else {
                            _ = len(v5)
                            switch -100 {
                            case -2147483648:
                                if true {
                                    if (1 >= 298) {
                                        switch "foo bar" {
                                        case "hello":
                                            switch 127 {
                                            case 42:
                                                v21 := v4[44]
                                                _ = v21
                                            case 271:
                                                v22 := GenS3{F0: 255, F1: 2}
                                                _ = v22
                                            case -1:
                                                var v23 uint64 = 4294967295
                                                _ = v23
                                            }
                                        case "":
                                            if (18446744073709551615 <= 42) {
                                                var v24 float32 = (float32((1.5 * ((0.1 - float64(-2.25)) + float64((0.1 / 1.0))))) - float32(3.14159))
                                                v1 = append(v1, 0)
                                                for fi25 := 0; fi25 < 20; fi25++ {
                                                    v26 := GenS3{F0: 65535, F1: 0}
                                                    v27 := GenS3{F0: 65535, F1: 2}
                                                    switch (-(fi25)) {
                                                    case 827:
                                                        fmt.Println(fi25, fi25, v3)
                                                    case 360:
                                                        v28 := v5[2147483648]
                                                        _ = v28
                                                    default:
                                                        _ = len(v1)
                                                    }
                                                    _ = fi25
                                                    _ = v26
                                                    _ = v27
                                                }
                                                _ = v24
                                            }
                                        }
                                        fmt.Println(v3)
                                    } else {
                                        fmt.Println(v3)
                                        var v29 int32 = 42
                                        fmt.Println(v29, v29)
                                        _ = v29
                                    }
                                }
                            }
                            _ = len(v5)
                        }
                    }
                }
            }
            var v30 int32 = 2147483647
            _ = v30
        }
        _ = v14
    }
    v31 := GenS4{F0: 254}
    _ = len(v1)
    v3 = (v3 + (v3 * 95))
    v3 += (v3 % ((v3) | 1))
    if (350 == 2) {
        for fi32 := 0; fi32 < 8; fi32++ {
            for fi33 := 0; fi33 < 10; fi33++ {
                fmt.Println(fi32)
                if true {
                    if true {
                        var v34 int64 = 0
                        var v35 uint = 676
                        _ = v34
                        _ = v35
                    } else {
                        v36 := GenS1{F0: ("foo bar" + ""), F1: (487 != -597), F2: 255}
                        fmt.Println(v3, fi32, fi32)
                        var v37 uint = 2
                        _ = v36
                        _ = v37
                    }
                }
                _ = fi33
            }
            var v38 int64 = 9223372036854775807
            v39 := GenS2{F0: -199, F1: 468, F2: 2147483647}
            _ = fi32
            _ = v38
            _ = v39
        }
        v40 := GenS5{F0: 747, F1: 65535, F2: 42}
        _ = v40
    }
    delete(v5, 222)
    fmt.Println(v3)
    if true {
        fmt.Println(v3)
        var v41 int8 = -112
        _ = v41
    } else {
        v42 := GenS2{F0: 42, F1: 2, F2: 2147483648}
        for fi43 := 0; fi43 < 18; fi43++ {
            for fi44 := 0; fi44 < 19; fi44++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi44
            }
            if len(v2) > 0 {
                v2[0] = 550
            }
            _ = fi43
        }
        v3 += (v3 % ((v3) | 1))
        _ = v42
    }
    fmt.Println(v3)
    fmt.Println(v3)
    for fi45 := 0; fi45 < 1; fi45++ {
        var v46 int32 = 465
        var v47 uint16 = 42
        _ = fi45
        _ = v46
        _ = v47
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
