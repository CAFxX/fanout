package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 string; F1 uint; F2 int8 }
func genFn0(p0 int16, p1 string) uint8 {
    var v6 int = 2147483647
    fmt.Println(v6, v6, p1)
    _ = p0
    _ = p1
    _ = v6
    return 42
}
func genFn1() uint16 {
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := []int{9223372036854775807, (fi7 + (fi7 & (fi7 + (^(fi7))))), -2147483648}
        var v9 int32 = 42
        if len(v8) > 0 {
            v8[0] = 1
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        var v11 string = ("" + "a")
        fmt.Println(v11, fi10)
        var v12 bool = false
        _ = fi10
        _ = v11
        _ = v12
    }
    if ((1 >= 2) && (42 < 317)) {
        var v13 uint32 = 42
        v13 -= v13
        v14 := []int{-9223372036854775808, -2147483648, -1}
        _ = v13
        _ = v14
    }
    return 255
}

func main() {
    var v1 float32 = float32((3.14159 - 1.5))
    var v2 float32 = 1.5
    var v3 uint32 = 2147483648
    var v4 int32 = 1
    var v5 int8 = 100
    var v15 float64 = (3.14159 - 0.1)
    if true {
        v16 := GenS1{F0: 975, F1: 2147483647}
        if (!true) {
            v4 -= v4
            var v17 float32 = (v2 * ((v1 / (v2 + v1)) * float32((v15 + v15))))
            v18 := []int{-1, -575, 9223372036854775807}
            _ = v17
            _ = v18
        }
        _ = v16
    }
    if true {
        fmt.Println(v5, v15, v2)
    } else {
        var v19 int8 = 127
        var v20 int8 = (v19 >> ((605) & 63))
        if true {
            switch "" {
            case "a":
                v21 := GenS2{F0: 880, F1: -785}
                _ = v21
            }
            v22 := GenS2{F0: 0, F1: int16(v3)}
            var v23 int64 = -923
            _ = v22
            _ = v23
        } else {
            v1 += (float32((3.14159 * v15)) * 1e10)
        }
        _ = v19
        _ = v20
    }
    v24 := []int{42, 108, 42, -1}
    var v25 uint32 = (v3 >> ((65535) & 63))
    switch 42 {
    case 32767:
        for fi26 := 0; fi26 < 20; fi26++ {
            v27 := GenS2{F0: 4294967295, F1: 1}
            _ = fi26
            _ = v27
        }
    default:
        v3 += v25
    }
    var v28 uint16 = 838
    fmt.Println(v2, v1, v3)
    for fi29 := 0; fi29 < 17; fi29++ {
        switch fi29 {
        case -282:
            for fi30 := 0; fi30 < 2; fi30++ {
                var v31 bool = true
                if ((1e10 / v2) == float32(3.14159)) {
                    var v32 string = (("" + (("a" + "a") + "hello")) + "foo bar")
                    v24 = append(v24, (fi30 << ((311) & 63)))
                    fmt.Println(fi30, v1, v15)
                    _ = v32
                } else {
                    var v33 int16 = 32767
                    var v34 int8 = (^(v5))
                    _ = v33
                    _ = v34
                }
                _ = fi30
                _ = v31
            }
        case 130:
            v25 = 2147483648
        case 0:
            for fi35 := 0; fi35 < 3; fi35++ {
                var v36 bool = false
                v5 = v5
                _ = fi35
                _ = v36
            }
        }
        if ((fi29 | (^(fi29))) != (fi29 & 9223372036854775807)) {
            v24 = append(v24, fi29)
            if len(v24) > 0 {
                v24[0] = (^(fi29))
            }
            for fi37 := 0; fi37 < 7; fi37++ {
                if (false && true) {
                    _ = len(v24)
                    v38 := GenS2{F0: 18446744073709551615, F1: -1}
                    _ = v38
                } else {
                    var v39 string = "\n\t"
                    _ = v39
                }
                if len(v24) >= 2 {
                    v24 = v24[1:len(v24)]
                }
                if true {
                    v4 = -2147483648
                }
                _ = fi37
            }
        }
        if ((float32(v15) + v1) >= v2) {
            v40 := GenS1{F0: -9223372036854775808, F1: -369}
            _ = v40
        } else {
            var v41 int = fi29
            v24 = append(v24, (^(fi29)))
            _ = v41
        }
        _ = fi29
    }
    for fi42 := 0; fi42 < 18; fi42++ {
        v43 := GenS1{F0: 1, F1: fi42}
        _ = fi42
        _ = v43
    }
    v44 := GenS3{F0: "\n\t", F1: 909, F2: v5}
    v45 := GenS1{F0: 1, F1: 2147483647}
    var v46 float32 = 0.0
    var v47 uint16 = v28
    for fi48 := 0; fi48 < 19; fi48++ {
        fmt.Println(v3, v3, v1)
        v15 = float64(v46)
        _ = fi48
    }
    if (!(-1 != 252)) {
        var v49 uint = 56
        _ = v49
    }
    if len(v24) >= 2 {
        v24 = v24[1:len(v24)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v5
}
