package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64; F1 string; F2 uint }
func genFn0(p0 uint, p1 int32, p2 int64) int8 {
    p2 -= (p2 >> ((65535) & 63))
    for fi5 := 0; fi5 < 14; fi5++ {
        v6 := GenS1{F0: 60}
        v7 := GenS1{F0: 32767}
        _ = fi5
        _ = v6
        _ = v7
    }
    switch 0 {
    case 4294967295:
        p2 = (-(p2))
    }
    var v8 uint8 = 42
    v8 += v8
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 1
}

func main() {
    var v1 float64 = 1e10
    var v2 float32 = ((float32((3.14159 * 1.5)) + ((1.5 / 1e10) - (float32((-2.25 - ((v1 * ((v1 - v1) - (v1 + v1))) * (v1 + float64(0.0))))) / 3.14159))) * ((0.1 + 3.14159) * (-2.25 * 0.1)))
    var v3 float32 = v2
    var v4 uint32 = 42
    if true {
        if (-32768 > 42) {
            v4 = 1
            v9 := []int{int(v4), 0, 4294967295, -156}
            _ = v9
        } else {
            for fi10 := 0; fi10 < 12; fi10++ {
                if (false || false) {
                    v2 -= v3
                    for fi11 := 0; fi11 < 2; fi11++ {
                        v12 := GenS2{F0: 498, F1: -1}
                        _ = fi11
                        _ = v12
                    }
                    fmt.Println(v4, v1)
                }
                var v13 int = fi10
                v14 := []int{4294967295, -818}
                _ = fi10
                _ = v13
                _ = v14
            }
            v1 += v1
            for fi15 := 0; fi15 < 11; fi15++ {
                v4 -= v4
                if ((v1 * v1) == v1) {
                    v16 := GenS2{F0: 32767, F1: 0}
                    var v17 float32 = -2.25
                    _ = v16
                    _ = v17
                } else {
                    v4 -= (^(v4))
                    fmt.Println(v4, v3, v1)
                    var v18 float32 = 0.0
                    _ = v18
                }
                v1 += v1
                _ = fi15
            }
        }
        var v19 int8 = -1
        for fi20 := 0; fi20 < 13; fi20++ {
            switch (fi20 * 0) {
            case 4294967295:
                v21 := GenS3{F0: -1}
                _ = v21
            }
            _ = fi20
        }
        _ = v19
    }
    v4 = 266
    var v22 int8 = 1
    v2 = v3
    v23 := []int{}
    if false {
        var v24 uint32 = v4
        v24 -= v4
        _ = v24
    } else {
        if (!true) {
            _ = len(v23)
            for fi25 := 0; fi25 < 8; fi25++ {
                var v26 int32 = 42
                if true {
                    v22 = 100
                    var v27 int8 = v22
                    _ = v27
                }
                _ = fi25
                _ = v26
            }
            switch 2147483647 {
            case 341:
                if len(v23) > 0 {
                    v23[0] = 0
                }
            }
        } else {
            if ((v3 > ((v2 - 1e10) + v2)) && (32767 <= 1)) {
                v28 := GenS4{F0: -1, F1: "foo bar", F2: 579}
                _ = v28
            }
        }
        for fi29 := 0; fi29 < 16; fi29++ {
            fmt.Println(fi29, v2, v22)
            var v30 int = -9223372036854775808
            fmt.Println(v4)
            _ = fi29
            _ = v30
        }
    }
    v22 -= (v22 + (v22 | genFn0(0, 1, 42)))
    v2 = 0.0
    v4 = v4
    fmt.Println(v4, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
