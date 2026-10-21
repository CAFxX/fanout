package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 uint64; F1 int32; F2 int8 }
func genFn0(p0 string, p1 int64, p2 int) uint8 {
    v7 := []int{128, p2}
    var v8 int16 = 1000
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 2
}
func genFn1(p0 int64) int64 {
    for fi9 := 0; fi9 < 2; fi9++ {
        fmt.Println(fi9, p0)
        var v10 string = "\n\t"
        var v11 uint64 = 108
        _ = fi9
        _ = v10
        _ = v11
    }
    p0 += 9223372036854775807
    v12 := []int{1, 9223372036854775807, 42}
    _ = p0
    _ = v12
    return -1
}

func main() {
    var v1 uint8 = 1
    v2 := make(map[string]int)
    v3 := make(map[string]bool)
    var v4 uint = 42
    var v5 int = -362
    var v6 uint32 = 399
    fmt.Println(v1)
    var v13 int16 = 322
    v2[("hello" + "\n\t")] = v5
    v1 += (v1 ^ (v1 + genFn0("a", 0, (v5 % (((v5 & (v5 | v5))) | 1)))))
    for fi14 := 0; fi14 < 1; fi14++ {
        var v15 uint32 = v6
        _ = fi14
        _ = v15
    }
    v16 := v2[""]
    if true {
        if ((2147483648 < (v6 ^ 265)) || true) {
            var v17 string = ("hello" + "")
            _ = v17
        } else {
            v18 := GenS1{F0: 100, F1: -1, F2: 127}
            _ = v18
        }
    }
    v5 = (v5 * 128)
    v19 := GenS1{F0: 0, F1: v13, F2: (v16 + v5)}
    for fi20 := 0; fi20 < 19; fi20++ {
        var v21 uint8 = (v1 ^ v1)
        fmt.Println(v5)
        var v22 uint16 = 32767
        _ = fi20
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 9; fi23++ {
        v13 -= -2
        fmt.Println(v5, v16)
        if (!((!true) || (!(0 >= (^(v4)))))) {
            fmt.Println(v16, v16)
            switch 86 {
            case 127:
                v3[""] = ((((("foo bar" + "foo bar") + "") + "\n\t") <= "\n\t") || (1.5 == (float32(0.0) * 1e10)))
            }
            v13 -= (^(v13))
        }
        _ = fi23
    }
    fmt.Println(v1, v13, v13)
    v24 := GenS2{F0: -797}
    v25 := GenS2{F0: -962}
    if false {
        if (v6 < (v6 * v6)) {
            for fi26 := 0; fi26 < 17; fi26++ {
                for fi27 := 0; fi27 < 6; fi27++ {
                    switch (("hello" + "a") + "a") {
                    case "a":
                        if ("hello" >= "") {
                            fmt.Println(v1)
                        }
                    default:
                        if true {
                            var v28 int32 = 2147483647
                            v29 := GenS2{F0: 1}
                            if true {
                                if ((0.1 >= ((0.1 * 1.5) + 0.1)) || true) {
                                    v30 := GenS2{F0: (v28 % (((v28 & (v28 & (v28 | v28)))) | 1))}
                                    switch fi26 {
                                    case 32767:
                                        v31 := GenS1{F0: 11, F1: (^(v13)), F2: v5}
                                        _ = v31
                                    case -1:
                                        v32 := []int{(^(v16))}
                                        switch (v5 << (((^(v4))) & 63)) {
                                        case -2147483648:
                                            v32 = append(v32, fi26)
                                        default:
                                            var v33 int8 = -128
                                            _ = v33
                                        }
                                        _ = v32
                                    }
                                    _ = v30
                                }
                            }
                            _ = len(v3)
                            _ = v28
                            _ = v29
                        } else {
                            v5 = fi27
                            v13 = (-(v13))
                            var v34 int32 = 1
                            _ = v34
                        }
                    }
                    switch (("\n\t" + "\n\t") + "hello") {
                    case "":
                        fmt.Println(v16, v13)
                    case "hello":
                        if (!(0 <= 65535)) {
                            v35 := GenS1{F0: 100, F1: -887, F2: (^(v16))}
                            var v36 uint8 = (v1 / (((v1 % ((genFn0("a", 42, (v5 + 9223372036854775807))) | 1))) | 1))
                            _ = v35
                            _ = v36
                        } else {
                            v37 := GenS2{F0: 457}
                            _ = v37
                        }
                    default:
                        if (v6 != 0) {
                            v38 := GenS2{F0: 42}
                            _ = v38
                        } else {
                            switch ("foo bar" + "hello") {
                            case "foo bar":
                                v39 := GenS2{F0: -2147483648}
                                _ = v39
                            default:
                                fmt.Println(v1, fi26)
                            }
                            var v40 uint8 = (v1 / ((v1) | 1))
                            _ = v40
                        }
                    }
                    _ = fi27
                }
                var v41 float64 = (0.1 - (1e10 / -2.25))
                for fi42 := 0; fi42 < 13; fi42++ {
                    v43 := GenS3{F0: v13, F1: true}
                    _ = fi42
                    _ = v43
                }
                _ = fi26
                _ = v41
            }
        } else {
            if false {
                var v44 string = "foo bar"
                _ = v44
            }
            v45 := []int{(v16 * (v16 + v16))}
            for fi46 := 0; fi46 < 17; fi46++ {
                fmt.Println(v5, v6)
                _ = fi46
            }
            _ = v45
        }
        var v47 int64 = 1
        v48 := GenS1{F0: -128, F1: 32767, F2: -702}
        _ = v47
        _ = v48
    } else {
        v49 := GenS4{F0: 417, F1: -2147483648, F2: -1}
        _ = v49
    }
    fmt.Println(v16)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v16
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
