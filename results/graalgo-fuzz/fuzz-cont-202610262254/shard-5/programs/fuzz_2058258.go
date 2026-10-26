package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint; F1 uint16; F2 int32 }
type GenS6 struct { F0 string }
func genFn0(p0 int16, p1 int64) uint {
    if (false || ("" <= (("" + "foo bar") + "foo bar"))) {
        fmt.Println(p0)
        v5 := GenS1{F0: true}
        _ = v5
    } else {
        fmt.Println(p1)
        v6 := []int{940, 513, 32768, -294}
        _ = v6
    }
    var v7 uint16 = 65535
    var v8 int16 = p0
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 1
}
func genFn1() uint32 {
    v9 := []int{128, 842}
    v10 := GenS2{F0: 381}
    var v11 uint8 = 42
    if len(v9) > 0 {
        v9[0] = 9223372036854775807
    }
    var v12 string = (("" + "\n\t") + ("hello" + "a"))
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return 2147483647
}

func main() {
    var v1 int64 = -1
    v2 := make(map[int]string)
    v3 := make(map[string]bool)
    var v4 uint = 65535
    v4 = (v4 + 1)
    v4 = (v4 + v4)
    v4 -= v4
    for fi13 := 0; fi13 < 8; fi13++ {
        v4 = 65535
        _ = fi13
    }
    if false {
        var v14 uint32 = 306
        _ = v14
    } else {
        v1 += v1
    }
    switch "a" {
    case "\n\t":
        v1 += 42
    case "":
        var v15 int16 = -967
        _ = v15
    default:
        if (int(v4) > 1) {
            for fi16 := 0; fi16 < 1; fi16++ {
                if (2 >= 255) {
                    v17 := GenS3{F0: fi16}
                    switch "hello" {
                    case "\n\t":
                        v1 += (v1 % (((-(v1))) | 1))
                    case "a":
                        var v18 string = ("" + "foo bar")
                        _ = v18
                    default:
                        switch ("foo bar" + "a") {
                        case "\n\t":
                            v1 += -9223372036854775808
                        case "a":
                            v19 := GenS1{F0: (1 < fi16)}
                            _ = v19
                        default:
                            for fi20 := 0; fi20 < 10; fi20++ {
                                if (((!true) && (0 > genFn1())) || (false || (1 <= 100))) {
                                    var v21 int32 = 1
                                    _ = v21
                                } else {
                                    if (false || true) {
                                        if ((3.14159 == 0.1) && false) {
                                            v22 := GenS4{F0: (v1 << ((v4) & 63))}
                                            var v23 int8 = 42
                                            var v24 uint = genFn0(3, 316)
                                            _ = v22
                                            _ = v23
                                            _ = v24
                                        }
                                        v25 := GenS2{F0: 555}
                                        v1 = v1
                                        if false {
                                            var v26 bool = (227 == 761)
                                            _ = v26
                                        }
                                        v4 -= (v4 + 1)
                                        _ = v25
                                    }
                                }
                                _ = fi20
                            }
                        }
                    }
                    _ = v17
                } else {
                    var v27 uint = (^(v4))
                    if (true && true) {
                        v27 += 1
                    }
                    _ = v27
                }
                v28 := GenS5{F0: v4, F1: 0, F2: -555}
                _ = fi16
                _ = v28
            }
            if (127 <= 2) {
                var v29 float64 = ((1.5 - 0.1) - 0.0)
                _ = v29
            }
        }
    }
    if (255 > 255) {
        for fi30 := 0; fi30 < 19; fi30++ {
            v31 := []int{(fi30 * (fi30 % (((^(fi30))) | 1))), (fi30 >> ((genFn0(32767, (v1 % (((v1 & v1)) | 1)))) & 63)), -485}
            _ = fi30
            _ = v31
        }
        var v32 int32 = -2147483648
        v33 := []int{-318, 32767, 632, 9223372036854775807}
        _ = v32
        _ = v33
    } else {
        var v34 uint32 = 2147483647
        if ((false && (-239 != 32767)) && true) {
            if (false && true) {
                var v35 int32 = -856
                var v36 bool = (true || (((true && false) && false) || true))
                _ = v35
                _ = v36
            }
            for fi37 := 0; fi37 < 9; fi37++ {
                for fi38 := 0; fi38 < 5; fi38++ {
                    var v39 string = "\n\t"
                    var v40 uint = v4
                    _ = fi38
                    _ = v39
                    _ = v40
                }
                _ = fi37
            }
        } else {
            for fi41 := 0; fi41 < 7; fi41++ {
                var v42 int = (fi41 % (((fi41 | (fi41 - (^(fi41))))) | 1))
                v4 = (v4 | (v4 + 65535))
                switch "hello" {
                case "\n\t":
                    if false {
                        fmt.Println(v1, fi41)
                        v42 = v42
                        switch (fi41 + fi41) {
                        case -2147483648:
                            v43 := GenS6{F0: "\n\t"}
                            _ = v43
                        case 42:
                            v44 := GenS5{F0: v4, F1: 65535, F2: 0}
                            _ = v44
                        case 4294967295:
                            var v45 float64 = (((-2.25 * ((-2.25 - 0.1) + 0.0)) * float64((1e10 * (0.1 - (-2.25 / 0.1))))) - (1e10 / 1.5))
                            _ = v45
                        default:
                            fmt.Println(fi41)
                        }
                    } else {
                        v42 += fi41
                    }
                default:
                    if true {
                        v1 += v1
                        for fi46 := 0; fi46 < 18; fi46++ {
                            var v47 uint = genFn0(695, 0)
                            _ = fi46
                            _ = v47
                        }
                    }
                }
                _ = fi41
                _ = v42
            }
            var v48 uint16 = 65535
            _ = v48
        }
        var v49 uint16 = 1
        _ = v34
        _ = v49
    }
    var v50 uint = v4
    fmt.Println(v4, v4, v1)
    switch -513 {
    case -2147483648:
        var v51 uint8 = 128
        _ = v51
    case 2147483647:
        switch 4294967295 {
        case 213:
            v52 := []int{2147483647, -9223372036854775808, 1}
            _ = v52
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v50
}
