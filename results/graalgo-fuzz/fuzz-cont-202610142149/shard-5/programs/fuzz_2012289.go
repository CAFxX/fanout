package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 uint8; F2 bool }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint32; F1 uint64 }
type GenS6 struct { F0 uint; F1 int16 }
type GenS7 struct { F0 uint16; F1 uint8; F2 int8 }
func genFn0(p0 int32, p1 uint64, p2 int8) int {
    var v6 int16 = 1
    p2 = (p2 ^ -89)
    if true {
        v7 := []int{9223372036854775807, -2147483648, int(v6)}
        _ = v7
    }
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 int32 = 42
        v10 := GenS1{F0: 1}
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -1
}
func genFn1(p0 uint16, p1 uint, p2 string) uint64 {
    if (((float32((((((1.5 / (float64(-2.25) * 1e10)) + (0.1 + 1.5)) + 0.0) + (-2.25 / 3.14159)) * -2.25)) + float32(0.0)) / float32(3.14159)) != 3.14159) {
        v11 := GenS2{F0: 1, F1: 1, F2: false}
        _ = v11
    }
    if (p0 >= p0) {
        v12 := GenS2{F0: 1, F1: 0, F2: false}
        fmt.Println(p2, p1)
        _ = v12
    }
    if (!(-9223372036854775808 >= 4294967295)) {
        v13 := GenS2{F0: 0, F1: 24, F2: true}
        _ = v13
    } else {
        var v14 float32 = 0.1
        switch -2147483648 {
        case -2147483648:
            v14 += 0.0
        }
        _ = v14
    }
    var v15 float32 = 0.1
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return 205
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 42
    v3 := []int{-1, -229}
    var v4 bool = false
    var v5 int16 = -633
    v16 := GenS3{F0: -1, F1: v4}
    for fi17 := 0; fi17 < 3; fi17++ {
        for fi18 := 0; fi18 < 20; fi18++ {
            var v19 uint8 = 8
            var v20 uint16 = 203
            _ = fi18
            _ = v19
            _ = v20
        }
        v2 = 562
        var v21 int = (fi17 & 127)
        _ = fi17
        _ = v21
    }
    v22 := GenS4{F0: 0}
    v5 += (v5 | 197)
    v5 -= (-(v5))
    var v23 uint16 = 32767
    if true {
        var v24 string = "foo bar"
        _ = v24
    }
    v4 = ((int64(v2) >= 0) && v4)
    var v25 string = "\n\t"
    v2 -= (v2 & (v2 >> ((2) & 63)))
    fmt.Println(v4)
    var v26 uint = 2
    if (0 > 185) {
        for fi27 := 0; fi27 < 10; fi27++ {
            v25 = v25
            if (genFn0(-1, 1, 127) <= (^(fi27))) {
                for fi28 := 0; fi28 < 3; fi28++ {
                    fmt.Println(v4, v26, v5)
                    v3 = append(v3, (fi27 / ((2147483647) | 1)))
                    var v29 int64 = 1
                    _ = fi28
                    _ = v29
                }
                if v4 {
                    var v30 float32 = 0.1
                    var v31 float64 = ((((0.0 - 0.0) - ((float64(-2.25) * float64(1e10)) - ((1.5 + -2.25) - (3.14159 + 0.1)))) - ((-2.25 - -2.25) + 1.5)) / (1.5 - 3.14159))
                    _ = v30
                    _ = v31
                }
                for fi32 := 0; fi32 < 5; fi32++ {
                    if ((v23 / ((v23) | 1)) >= v23) {
                        v25 = "a"
                        switch (fi27 + 0) {
                        case 2147483647:
                            v33 := GenS5{F0: v2, F1: 2}
                            _ = v33
                        default:
                            v1[(v2 + (v2 * v2))] = false
                        }
                        v34 := v1[v2]
                        v35 := GenS4{F0: 100}
                        switch (fi32 << (((v26 % ((uint(v2)) | 1))) & 63)) {
                        case 1:
                            var v36 string = ""
                            _ = v36
                        default:
                            var v37 uint8 = 0
                            _ = v37
                        }
                        _ = v34
                        _ = v35
                    } else {
                        v38 := GenS2{F0: 18446744073709551615, F1: 127, F2: v4}
                        var v39 uint = 2
                        _ = v38
                        _ = v39
                    }
                    _ = fi32
                }
            } else {
                if v4 {
                    fmt.Println(v4, v26)
                }
            }
            _ = fi27
        }
    } else {
        for fi40 := 0; fi40 < 18; fi40++ {
            switch (v25 + v25) {
            case "":
                v41 := GenS3{F0: 100, F1: v4}
                _ = v41
            case "\n\t":
                if (!v4) {
                    v25 = v25
                }
            }
            for fi42 := 0; fi42 < 13; fi42++ {
                v25 = v25
                if v4 {
                    fmt.Println(fi40, v4)
                    v43 := GenS2{F0: 45, F1: 128, F2: v4}
                    v5 -= (v5 + (v5 - (v5 / ((v5) | 1))))
                    _ = v43
                }
                delete(v1, v2)
                _ = fi42
            }
            _ = fi40
        }
        for fi44 := 0; fi44 < 17; fi44++ {
            for fi45 := 0; fi45 < 17; fi45++ {
                var v46 float32 = (3.14159 * 1e10)
                v46 = float32((0.1 - 0.0))
                var v47 float64 = (float64(v46) + float64(((((0.1 / (v46 / (float32(((3.14159 + (-2.25 * float64((float32(-2.25) / float32(3.14159))))) - (1e10 / (3.14159 - 1.5)))) * (((v46 / v46) + v46) - 3.14159)))) * (1e10 * (-2.25 * 1.5))) / 1e10) * (0.1 + float32(3.14159)))))
                _ = fi45
                _ = v46
                _ = v47
            }
            v48 := GenS6{F0: 18446744073709551615, F1: (v5 ^ (v5 / ((v5) | 1)))}
            var v49 uint8 = 0
            _ = fi44
            _ = v48
            _ = v49
        }
        var v50 int16 = v5
        _ = v50
    }
    v51 := GenS6{F0: v26, F1: v5}
    switch -9223372036854775808 {
    case 351:
        v3 = append(v3, genFn0(-796, 473, -1))
    case 0:
        v4 = true
    default:
        _ = len(v3)
    }
    v25 = ""
    v52 := GenS7{F0: (v23 & v23), F1: 168, F2: 127}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
}
