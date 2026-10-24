package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint8; F1 int }
type GenS7 struct { F0 uint8; F1 int8; F2 bool }
type GenS8 struct { F0 uint8; F1 uint16 }
type GenS9 struct { F0 uint; F1 int }
type GenS10 struct { F0 uint32 }
func genFn0(p0 uint16) int64 {
    fmt.Println(p0)
    v5 := GenS1{F0: 168}
    v6 := []int{int(p0), -2147483648, 1}
    switch 32768 {
    case 32768:
        for fi7 := 0; fi7 < 10; fi7++ {
            p0 = p0
            if ((((!false) || false) || (125 <= -128)) || false) {
                if len(v6) > 0 {
                    v6[0] = 2147483647
                }
            } else {
                for fi8 := 0; fi8 < 15; fi8++ {
                    p0 = (p0 * p0)
                    var v9 int32 = 1
                    _ = fi8
                    _ = v9
                }
                if len(v6) > 0 {
                    v6[0] = fi7
                }
                var v10 int16 = int16(fi7)
                _ = v10
            }
            _ = fi7
        }
    case -1:
        v11 := GenS2{F0: (false && false)}
        _ = v11
    case 752:
        var v12 int = 127
        _ = v12
    default:
        p0 -= p0
    }
    var v13 int32 = 25
    _ = p0
    _ = v13
    _ = v5
    _ = v6
    return -1
}
func genFn1(p0 bool) uint {
    v14 := []int{}
    if ((!p0) || (float64(float32((1.5 * -2.25))) != -2.25)) {
        _ = len(v14)
        v15 := GenS3{F0: 42}
        for fi16 := 0; fi16 < 11; fi16++ {
            if true {
                if (!p0) {
                    if p0 {
                        var v17 int32 = 42
                        _ = v17
                    }
                    var v18 float64 = 3.14159
                    if len(v14) >= 2 {
                        v14 = v14[1:len(v14)]
                    }
                    _ = v18
                }
                v19 := GenS4{F0: 1}
                var v20 int8 = 42
                _ = v19
                _ = v20
            } else {
                for fi21 := 0; fi21 < 8; fi21++ {
                    if (float64(float32(3.14159)) <= -2.25) {
                        var v22 int64 = 3
                        _ = v22
                    }
                    var v23 int = (fi16 | fi16)
                    var v24 float64 = (0.0 + ((((0.0 * (((-2.25 - (1.5 - -2.25)) * 0.1) + -2.25)) / 3.14159) + (-2.25 / (3.14159 - -2.25))) - 1e10))
                    _ = fi21
                    _ = v23
                    _ = v24
                }
                for fi25 := 0; fi25 < 6; fi25++ {
                    _ = len(v14)
                    _ = fi25
                }
            }
            p0 = (float32(3.14159) <= (0.1 / 0.1))
            _ = fi16
        }
        _ = v15
    }
    fmt.Println(p0)
    switch "\n\t" {
    case "\n\t":
        for fi26 := 0; fi26 < 1; fi26++ {
            for fi27 := 0; fi27 < 6; fi27++ {
                p0 = (699 <= 65535)
                v28 := GenS5{F0: 0}
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
                _ = fi27
                _ = v28
            }
            for fi29 := 0; fi29 < 11; fi29++ {
                var v30 int8 = 100
                fmt.Println(v30)
                var v31 int8 = 36
                if (p0 && (1.5 <= -2.25)) {
                    var v32 uint = 1
                    _ = v32
                }
                v33 := GenS6{F0: 2, F1: -9223372036854775808}
                var v34 string = "foo bar"
                fmt.Println(p0, v34)
                v30 -= v31
                _ = fi29
                _ = v30
                _ = v31
                _ = v33
                _ = v34
            }
            _ = fi26
        }
    }
    switch "hello" {
    case "foo bar":
        var v35 uint16 = 42
        _ = v35
    }
    _ = p0
    _ = v14
    return 18446744073709551615
}

func main() {
    var v1 uint32 = 2147483647
    v2 := []uint32{(v1 | (^(v1))), (^(v1)), v1, v1}
    v3 := make(map[string]bool)
    var v4 int64 = -9223372036854775808
    v36 := GenS5{F0: 127}
    v3["a"] = (2 != 128)
    _ = len(v2)
    fmt.Println(v4, v4)
    if false {
        v37 := GenS7{F0: 129, F1: -16, F2: true}
        _ = v37
    }
    v1 = v1
    v38 := v3["\n\t"]
    for fi39 := 0; fi39 < 11; fi39++ {
        for fi40 := 0; fi40 < 13; fi40++ {
            if ((42 > 32767) && true) {
                v2 = append(v2, v1)
                v1 = v1
                v4 += v4
            }
            var v41 uint64 = 530
            _ = fi40
            _ = v41
        }
        _ = fi39
    }
    v38 = v38
    var v42 float32 = 1.5
    var v43 int = 127
    if false {
        fmt.Println(v43, v1)
        v44 := GenS8{F0: 255, F1: 0}
        for fi45 := 0; fi45 < 19; fi45++ {
            if v38 {
                v43 -= v43
                _ = len(v3)
            } else {
                var v46 int = (v43 / (((v43 << ((255) & 63))) | 1))
                var v47 uint = 2
                _ = v46
                _ = v47
            }
            v38 = ((fi45 & 451) > (v43 % ((fi45) | 1)))
            if true {
                if len(v2) > 0 {
                    v2[0] = v1
                }
                for fi48 := 0; fi48 < 4; fi48++ {
                    if (!true) {
                        var v49 int16 = 42
                        _ = v49
                    } else {
                        v2 = append(v2, v1)
                    }
                    var v50 uint32 = v1
                    if len(v2) > 0 {
                        v2[0] = 1
                    }
                    v50 = v50
                    _ = fi48
                    _ = v50
                }
                if v38 {
                    if len(v2) > 0 {
                        v2[0] = (^(v1))
                    }
                } else {
                    v2 = append(v2, 65535)
                    _ = len(v3)
                    fmt.Println(v4, v1, v43)
                }
            }
            _ = fi45
        }
        _ = v44
    } else {
        v38 = (!false)
        if ((18446744073709551615 > 218) && v38) {
            v42 = ((v42 + v42) * 3.14159)
            var v51 int32 = 1
            _ = v51
        }
    }
    delete(v3, "a")
    for fi52 := 0; fi52 < 20; fi52++ {
        v53 := GenS9{F0: 65535, F1: fi52}
        _ = fi52
        _ = v53
    }
    v4 = (v4 % ((int64(v43)) | 1))
    v54 := GenS10{F0: (v1 ^ 167)}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v38
    _ = v4
    _ = v42
    _ = v43
    _ = v54
}
