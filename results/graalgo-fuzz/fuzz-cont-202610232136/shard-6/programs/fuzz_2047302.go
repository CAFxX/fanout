package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 uint; F1 uint; F2 int32 }
type GenS4 struct { F0 int32; F1 int64; F2 uint64 }
type GenS5 struct { F0 uint; F1 int32; F2 int32 }
type GenS6 struct { F0 uint8; F1 int32; F2 uint8 }
func genFn0(p0 int16, p1 uint64) uint32 {
    for fi4 := 0; fi4 < 18; fi4++ {
        var v5 bool = true
        _ = fi4
        _ = v5
    }
    p1 = (p1 * (p1 << ((891) & 63)))
    _ = p0
    _ = p1
    return 65535
}
func genFn1(p0 uint64) uint32 {
    switch 4294967295 {
    case 0:
        if true {
            fmt.Println(p0)
            if true {
                if false {
                    v6 := []int{9223372036854775807, 4294967295, -438, -2147483648}
                    for fi7 := 0; fi7 < 14; fi7++ {
                        p0 = 1
                        _ = fi7
                    }
                    _ = v6
                } else {
                    p0 -= (p0 * 36)
                    v8 := GenS1{F0: -2147483648}
                    _ = v8
                }
                if (true && true) {
                    if true {
                        var v9 bool = false
                        fmt.Println(v9, p0)
                        _ = v9
                    } else {
                        var v10 int64 = 9223372036854775807
                        if true {
                            var v11 int16 = -32768
                            for fi12 := 0; fi12 < 8; fi12++ {
                                switch "a" {
                                case "foo bar":
                                    var v13 uint32 = genFn0(0, p0)
                                    _ = v13
                                case "":
                                    if false {
                                        for fi14 := 0; fi14 < 2; fi14++ {
                                            if (!(0.1 > 1e10)) {
                                                var v15 bool = true
                                                _ = v15
                                            } else {
                                                v16 := GenS1{F0: -165}
                                                if ((p0 + (p0 | p0)) < (p0 % (((p0 >> ((1) & 63))) | 1))) {
                                                    var v17 uint = 565
                                                    fmt.Println(fi14, fi12)
                                                    _ = v17
                                                } else {
                                                    v18 := []int{}
                                                    if len(v18) >= 2 {
                                                        v18 = v18[1:len(v18)]
                                                    }
                                                    var v19 uint32 = 2147483648
                                                    _ = v18
                                                    _ = v19
                                                }
                                                _ = v16
                                            }
                                            _ = fi14
                                        }
                                        fmt.Println(v10, fi12)
                                        v20 := []int{-1}
                                        _ = v20
                                    }
                                case "hello":
                                    v21 := GenS2{F0: (^(p0)), F1: 2}
                                    _ = v21
                                }
                                fmt.Println(v11, v10)
                                _ = fi12
                            }
                            _ = v11
                        } else {
                            v22 := []int{256}
                            if len(v22) >= 2 {
                                v22 = v22[1:len(v22)]
                            }
                            v10 = -1
                            _ = v22
                        }
                        var v23 uint = 207
                        _ = v10
                        _ = v23
                    }
                    var v24 int64 = 9223372036854775807
                    _ = v24
                } else {
                    var v25 int64 = 834
                    if true {
                        v25 = v25
                        var v26 int = -9223372036854775808
                        _ = v26
                    }
                    fmt.Println(v25, p0)
                    _ = v25
                }
            }
        } else {
            p0 = p0
        }
    case 128:
        for fi27 := 0; fi27 < 16; fi27++ {
            v28 := GenS1{F0: -2147483648}
            var v29 uint = 65535
            var v30 string = "a"
            _ = fi27
            _ = v28
            _ = v29
            _ = v30
        }
    }
    var v31 uint = 4294967295
    var v32 uint8 = 127
    var v33 float64 = ((0.0 / 1.0) / (1.5 - (((0.0 / (0.1 / 1.5)) * 0.0) * ((1e10 + ((float64(float32(0.1)) / 1.5) / 1e10)) - 0.0))))
    _ = p0
    _ = v31
    _ = v32
    _ = v33
    return 500
}

func main() {
    var v1 float64 = 0.1
    var v2 int32 = -199
    var v3 int = 32767
    v2 -= (v2 / (((v2 % (((v2 | (v2 | v2))) | 1))) | 1))
    v34 := GenS3{F0: 1, F1: 255, F2: (v2 * v2)}
    v1 -= (v1 / (v1 - ((3.14159 * v1) + 0.1)))
    fmt.Println(v3, v1, v1)
    if true {
        var v35 string = "foo bar"
        v3 = v3
        _ = v35
    } else {
        if (42 != 2) {
            var v36 int32 = -1
            for fi37 := 0; fi37 < 3; fi37++ {
                v38 := []int{v3, (^(v3))}
                _ = fi37
                _ = v38
            }
            _ = v36
        }
        for fi39 := 0; fi39 < 13; fi39++ {
            for fi40 := 0; fi40 < 20; fi40++ {
                switch "a" {
                case "foo bar":
                    if false {
                        v1 = -2.25
                        var v41 uint8 = 2
                        _ = v41
                    }
                default:
                    var v42 int32 = (-(v2))
                    _ = v42
                }
                v43 := []int{(fi39 ^ v3)}
                _ = fi40
                _ = v43
            }
            _ = fi39
        }
    }
    v44 := GenS3{F0: 65535, F1: 0, F2: v2}
    v45 := GenS4{F0: v2, F1: 1, F2: 18446744073709551615}
    fmt.Println(v2)
    if false {
        for fi46 := 0; fi46 < 13; fi46++ {
            v1 = (v1 - (0.0 + v1))
            _ = fi46
        }
        var v47 int = v3
        _ = v47
    }
    switch "foo bar" {
    case "hello":
        var v48 int8 = 33
        _ = v48
    case "\n\t":
        v3 += 42
    default:
        if (721 < (v2 % ((v2) | 1))) {
            switch "foo bar" {
            case "a":
                var v49 bool = (true || (61 >= 42))
                _ = v49
            case "foo bar":
                v50 := GenS5{F0: 1, F1: v2, F2: (^(v2))}
                _ = v50
            case "\n\t":
                v51 := []int{v3, v3, (v3 & v3), v3}
                _ = v51
            default:
                var v52 uint32 = 4294967295
                _ = v52
            }
            for fi53 := 0; fi53 < 19; fi53++ {
                fmt.Println(v1, v3)
                var v54 int64 = 42
                _ = fi53
                _ = v54
            }
            v55 := GenS6{F0: 204, F1: 42, F2: 42}
            _ = v55
        }
    }
    if (207 > 65535) {
        v3 = v3
        for fi56 := 0; fi56 < 13; fi56++ {
            if (1 > 0) {
                if true {
                    var v57 uint8 = 4
                    _ = v57
                }
                var v58 int16 = -817
                _ = v58
            } else {
                if (1 != (v2 - 2147483647)) {
                    fmt.Println(fi56)
                    switch fi56 {
                    case -902:
                        fmt.Println(fi56, v3)
                    case 4294967295:
                        var v59 float64 = (v1 - v1)
                        _ = v59
                    case -193:
                        if (211 >= 805) {
                            var v60 int64 = 42
                            _ = v60
                        } else {
                            var v61 uint16 = 961
                            fmt.Println(v1, v2)
                            _ = v61
                        }
                    default:
                        var v62 uint8 = 128
                        _ = v62
                    }
                }
            }
            _ = fi56
        }
    } else {
        var v63 int32 = (v2 * v2)
        _ = v63
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v44
    _ = v45
}
