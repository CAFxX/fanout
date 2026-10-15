package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool; F1 int64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 bool }
func genFn0(p0 uint64, p1 int8) int32 {
    for fi6 := 0; fi6 < 7; fi6++ {
        for fi7 := 0; fi7 < 9; fi7++ {
            p0 = 1
            var v8 uint64 = 42
            _ = fi7
            _ = v8
        }
        var v9 float32 = 0.1
        _ = fi6
        _ = v9
    }
    v10 := []int{-157}
    var v11 uint32 = 42
    var v12 uint16 = 2
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    return int32(v12)
}

func main() {
    var v1 uint64 = 0
    var v2 int16 = 42
    var v3 int64 = 42
    var v4 int32 = -880
    var v5 bool = true
    if (127 == -772) {
        v2 -= 32767
        v13 := GenS1{F0: (^(v3)), F1: v5, F2: 127}
        if v5 {
            if v5 {
                var v14 uint64 = 4294967295
                var v15 int8 = 8
                _ = v14
                _ = v15
            }
            var v16 float64 = (0.0 + (3.14159 / 1.0))
            switch 4294967295 {
            case 9223372036854775807:
                if ((v4 - -254) >= (v4 << ((359) & 63))) {
                    v2 = -203
                } else {
                    if true {
                        for fi17 := 0; fi17 < 2; fi17++ {
                            v2 = 42
                            v18 := []int{(^(fi17)), fi17}
                            v1 = v1
                            _ = fi17
                            _ = v18
                        }
                        v16 -= ((0.1 + 0.0) * (float64((-2.25 + ((-2.25 * 1.5) * ((0.1 - (0.0 / float32(float64(-2.25)))) - 0.1)))) * v16))
                        v19 := GenS2{F0: 18446744073709551615}
                        _ = v19
                    } else {
                        v1 += (v1 * 42)
                    }
                    for fi20 := 0; fi20 < 10; fi20++ {
                        v2 = v2
                        if (100 < 1) {
                            var v21 uint16 = 32767
                            fmt.Println(v21)
                            if v5 {
                                v4 += v4
                                v21 = (v21 | (v21 - (v21 << ((uint(v1)) & 63))))
                            } else {
                                v21 = v21
                            }
                            _ = v21
                        } else {
                            fmt.Println(fi20, v2, v5)
                            var v22 int64 = v3
                            _ = v22
                        }
                        _ = fi20
                    }
                }
            default:
                switch 930 {
                case 127:
                    var v23 uint32 = 65535
                    _ = v23
                case -2147483648:
                    fmt.Println(v16, v3)
                case 42:
                    if (!false) {
                        v24 := []int{9223372036854775807, -9223372036854775808}
                        _ = v24
                    }
                default:
                    var v25 bool = v5
                    _ = v25
                }
            }
            _ = v16
        }
        _ = v13
    } else {
        if v5 {
            v1 = (v1 % ((v1) | 1))
            if v5 {
                fmt.Println(v3)
            } else {
                for fi26 := 0; fi26 < 20; fi26++ {
                    if (-1 != 122) {
                        if v5 {
                            v4 = 1
                            var v27 int16 = -27
                            var v28 uint16 = 211
                            _ = v27
                            _ = v28
                        } else {
                            var v29 bool = true
                            _ = v29
                        }
                        var v30 string = ("" + "\n\t")
                        v3 = v3
                        _ = v30
                    } else {
                        var v31 float64 = (0.1 + ((3.14159 * 1e10) * 0.0))
                        v5 = v5
                        _ = v31
                    }
                    var v32 float32 = (-2.25 * 3.14159)
                    _ = fi26
                    _ = v32
                }
                v3 = -1
            }
        } else {
            fmt.Println(v4, v3)
            for fi33 := 0; fi33 < 10; fi33++ {
                if true {
                    v34 := []int{}
                    _ = v34
                }
                _ = fi33
            }
            for fi35 := 0; fi35 < 20; fi35++ {
                var v36 bool = (0 > v3)
                switch ("foo bar" + "") {
                case "\n\t":
                    var v37 uint8 = 107
                    _ = v37
                case "":
                    v36 = true
                default:
                    for fi38 := 0; fi38 < 9; fi38++ {
                        var v39 float64 = -2.25
                        var v40 uint16 = 967
                        v41 := GenS1{F0: (v3 | v3), F1: false, F2: 0}
                        _ = fi38
                        _ = v39
                        _ = v40
                        _ = v41
                    }
                }
                v42 := GenS1{F0: (v3 - -1), F1: false, F2: 127}
                _ = fi35
                _ = v36
                _ = v42
            }
        }
        switch 32767 {
        case 4294967295:
            v43 := GenS3{F0: (!v5), F1: (v3 ^ (v3 | (-(v3))))}
            _ = v43
        }
        v2 -= v2
    }
    v44 := []int{}
    var v45 uint32 = 4294967295
    var v46 float32 = float32(((1e10 / 3.14159) + float64(3.14159)))
    fmt.Println(v46)
    _ = len(v44)
    v47 := GenS1{F0: 0, F1: (v5 || false), F2: 42}
    var v48 bool = ((-9223372036854775808 > v3) && (v1 > 270))
    v49 := GenS4{F0: v4}
    var v50 uint64 = 1
    v51 := GenS5{F0: v48}
    v1 += v1
    if v48 {
        var v52 int16 = v2
        _ = v52
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v51
}
