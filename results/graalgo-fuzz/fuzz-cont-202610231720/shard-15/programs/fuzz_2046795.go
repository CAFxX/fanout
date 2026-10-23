package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 string; F1 bool; F2 uint32 }
type GenS4 struct { F0 uint16; F1 int64; F2 int8 }
func genFn0() int8 {
    v6 := GenS1{F0: 1}
    v7 := GenS1{F0: 65535}
    _ = v6
    _ = v7
    return -1
}
func genFn1(p0 int, p1 uint16, p2 uint16) int16 {
    if (float32((0.1 - ((0.0 * (-2.25 - 3.14159)) / 1.0))) <= -2.25) {
        v8 := []int{(-(p0)), p0}
        if len(v8) > 0 {
            v8[0] = (p0 << ((1) & 63))
        }
        _ = v8
    }
    var v9 bool = (true || true)
    v10 := []int{p0, p0, (p0 * p0)}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 1000
}

func main() {
    var v1 int32 = -2147483648
    v2 := []int{9223372036854775807, -876}
    var v3 bool = (float64((float32(1.5) + 3.14159)) >= 3.14159)
    var v4 uint = 42
    v5 := make(map[string]bool)
    v2 = append(v2, -1)
    if v3 {
        v11 := GenS1{F0: (v4 & 610)}
        _ = v11
    }
    if ((0.1 * 0.1) > 3.14159) {
        if v3 {
            var v12 int = -9223372036854775808
            v3 = false
            var v13 int16 = 1
            _ = v12
            _ = v13
        } else {
            v4 = v4
        }
        v14 := GenS1{F0: (v4 | (v4 % ((963) | 1)))}
        _ = v14
    }
    var v15 int = 0
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi16 := 0; fi16 < 17; fi16++ {
        var v17 uint32 = 0
        v18 := GenS1{F0: (v4 + v4)}
        if ("hello" == "\n\t") {
            fmt.Println(fi16)
            if (v3 && (0 <= 1)) {
                v17 = v17
                if len(v2) > 0 {
                    v2[0] = fi16
                }
                for fi19 := 0; fi19 < 7; fi19++ {
                    v17 -= v17
                    v1 += v1
                    _ = fi19
                }
            }
            v1 -= 0
        }
        _ = fi16
        _ = v17
        _ = v18
    }
    v1 = (-(v1))
    v20 := GenS1{F0: v4}
    var v21 float32 = 1.5
    var v22 int32 = 0
    v15 = v15
    _ = len(v5)
    switch v15 {
    case -577:
        var v23 uint64 = 891
        _ = v23
    case 42:
        if true {
            for fi24 := 0; fi24 < 3; fi24++ {
                if false {
                    v22 += 2147483647
                    var v25 uint16 = 1
                    v26 := GenS2{F0: 1, F1: (^(v4))}
                    _ = v25
                    _ = v26
                } else {
                    if v3 {
                        v27 := GenS3{F0: (("" + "hello") + "\n\t"), F1: v3, F2: 65535}
                        _ = v27
                    }
                    v1 += 776
                }
                _ = fi24
            }
            if ((v15 % ((300) | 1)) <= v15) {
                for fi28 := 0; fi28 < 3; fi28++ {
                    v29 := v5["hello"]
                    _ = fi28
                    _ = v29
                }
                if ((^(v15)) != (v15 - v15)) {
                    var v30 float32 = ((v21 - 0.1) - (float32(1e10) - float32(-2.25)))
                    if v3 {
                        v22 += v1
                        v31 := GenS4{F0: 32767, F1: 0, F2: -1}
                        _ = v31
                    }
                    v22 += 0
                    _ = v30
                }
                for fi32 := 0; fi32 < 8; fi32++ {
                    var v33 int64 = -9223372036854775808
                    _ = fi32
                    _ = v33
                }
            }
        }
    }
    var v34 uint32 = 836
    v22 = (v1 ^ v22)
    for fi35 := 0; fi35 < 17; fi35++ {
        var v36 int8 = 1
        for fi37 := 0; fi37 < 13; fi37++ {
            v38 := GenS2{F0: 392, F1: (v4 | 353)}
            v3 = v3
            _ = fi37
            _ = v38
        }
        _ = fi35
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
