package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 uint16 }
type GenS2 struct { F0 int; F1 int16; F2 int64 }
type GenS3 struct { F0 int16; F1 uint32; F2 int32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int8; F1 int64; F2 uint32 }
type GenS6 struct { F0 int64; F1 string; F2 uint }
func genFn0(p0 uint32, p1 uint64, p2 uint64) int32 {
    var v6 int64 = 621
    switch ("" + "\n\t") {
    case "foo bar":
        var v7 float64 = ((-2.25 * float64(3.14159)) - float64((1.5 + 1.5)))
        _ = v7
    case "\n\t":
        v8 := []int{-943, 4294967295}
        _ = v8
    case "a":
        if (18446744073709551615 <= 0) {
            v9 := GenS1{F0: ((("hello" + "hello") + "a") + "\n\t"), F1: -1, F2: 452}
            p0 = p0
            _ = v9
        } else {
            if true {
                v6 = (v6 ^ v6)
                if true {
                    var v10 uint32 = (p0 | 829)
                    _ = v10
                }
                var v11 float64 = float64(3.14159)
                _ = v11
            }
            v12 := []int{32767, 546, 32767, 4294967295}
            _ = v12
        }
    default:
        if true {
            p2 = p2
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 1
}

func main() {
    var v1 string = "hello"
    v2 := make(map[uint32]int)
    var v3 bool = (0 <= 2147483647)
    var v4 uint8 = 84
    var v5 uint32 = 2147483648
    fmt.Println(v1, v3)
    v5 -= (v5 + v5)
    v13 := GenS1{F0: "a", F1: 42, F2: 65535}
    var v14 int8 = -128
    v15 := []int{1, -453, -4, 128}
    v1 = "\n\t"
    switch ("a" + v1) {
    case "":
        v3 = ((v4 >> ((18446744073709551615) & 63)) < v4)
    case "\n\t":
        fmt.Println(v1)
    }
    for fi16 := 0; fi16 < 17; fi16++ {
        if v3 {
            _ = len(v2)
            if (((v14 >> ((18446744073709551615) & 63)) < 84) || true) {
                var v17 int16 = 0
                var v18 int64 = -9223372036854775808
                _ = v17
                _ = v18
            }
        } else {
            if (560 <= 0) {
                v4 -= 191
            } else {
                v15 = append(v15, (fi16 % (((-(fi16))) | 1)))
                var v19 int64 = 1
                _ = v19
            }
        }
        fmt.Println(v3, v14, v4)
        v20 := GenS1{F0: v1, F1: (v14 * 107), F2: 2}
        _ = fi16
        _ = v20
    }
    if (461 > 793) {
        var v21 int16 = 534
        var v22 int = 492
        _ = v21
        _ = v22
    } else {
        for fi23 := 0; fi23 < 6; fi23++ {
            var v24 int64 = -9223372036854775808
            v25 := v2[(v5 / (((^(v5))) | 1))]
            _ = fi23
            _ = v24
            _ = v25
        }
        if ((v4 / ((uint8(v5)) | 1)) > (v4 ^ (v4 | (^(v4))))) {
            v3 = v3
            v26 := GenS2{F0: 4294967295, F1: -1, F2: 277}
            v3 = v3
            _ = v26
        } else {
            v14 -= v14
            if v3 {
                v27 := GenS3{F0: -414, F1: 4294967295, F2: 42}
                var v28 uint8 = v4
                _ = v27
                _ = v28
            } else {
                _ = len(v2)
                v29 := GenS4{F0: v1}
                v2[v5] = 9223372036854775807
                _ = v29
            }
            if (v3 && (v5 > v5)) {
                v2[(v5 % (((v5 ^ 2147483648)) | 1))] = 370
                v30 := GenS5{F0: 100, F1: 0, F2: (^(v5))}
                _ = v30
            }
        }
    }
    _ = len(v2)
    switch 127 {
    case 2147483647:
        for fi31 := 0; fi31 < 2; fi31++ {
            if v3 {
                v3 = true
                for fi32 := 0; fi32 < 18; fi32++ {
                    fmt.Println(v4, v3, v1)
                    v15 = append(v15, (^(fi32)))
                    var v33 int = fi32
                    _ = fi32
                    _ = v33
                }
                if len(v15) > 0 {
                    v15[0] = (fi31 & (fi31 - fi31))
                }
            } else {
                if (!v3) {
                    v15 = append(v15, (fi31 - -2147483648))
                }
                var v34 uint16 = 760
                _ = v34
            }
            fmt.Println(fi31, v3)
            _ = fi31
        }
    default:
        switch 579 {
        case 32767:
            v35 := GenS6{F0: -9223372036854775808, F1: v1, F2: 255}
            _ = v35
        case -298:
            switch 4294967295 {
            case 42:
                var v36 uint8 = 42
                _ = v36
            }
        }
    }
    v5 = v5
    var v37 int32 = 42
    fmt.Println(v37, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
}
