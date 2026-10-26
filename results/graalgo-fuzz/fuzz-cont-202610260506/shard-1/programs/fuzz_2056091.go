package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64; F1 int32; F2 int64 }
func genFn0(p0 uint64) uint {
    v4 := []int{9223372036854775807}
    fmt.Println(p0)
    p0 = (^(p0))
    v5 := GenS1{F0: 0, F1: ("a" + ("foo bar" + "a"))}
    _ = p0
    _ = v4
    _ = v5
    return 551
}

func main() {
    v1 := make(map[int]string)
    var v2 int64 = 0
    var v3 int64 = -9223372036854775808
    if (0.1 == 3.14159) {
        v3 -= v3
        v6 := v1[0]
        _ = v6
    } else {
        if (!false) {
            v2 = v3
        } else {
            for fi7 := 0; fi7 < 1; fi7++ {
                if true {
                    v8 := GenS1{F0: -117, F1: "foo bar"}
                    for fi9 := 0; fi9 < 20; fi9++ {
                        if true {
                            var v10 int32 = 2147483647
                            _ = v10
                        }
                        v11 := GenS1{F0: 0, F1: (("a" + "a") + "\n\t")}
                        _ = fi9
                        _ = v11
                    }
                    var v12 int32 = -235
                    _ = v12
                    _ = v8
                }
                for fi13 := 0; fi13 < 4; fi13++ {
                    var v14 float32 = (1.5 - 0.0)
                    var v15 uint64 = 18446744073709551615
                    if ((v14 / v14) <= -2.25) {
                        var v16 bool = false
                        var v17 string = "foo bar"
                        v16 = v16
                        _ = v16
                        _ = v17
                    } else {
                        var v18 float64 = 3.14159
                        var v19 float32 = v14
                        _ = v18
                        _ = v19
                    }
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                _ = fi7
            }
            v20 := GenS1{F0: 0, F1: "a"}
            _ = v20
        }
        for fi21 := 0; fi21 < 1; fi21++ {
            var v22 uint64 = 339
            _ = fi21
            _ = v22
        }
        fmt.Println(v3, v3)
    }
    fmt.Println(v2, v3)
    v2 = v2
    v2 -= v2
    for fi23 := 0; fi23 < 9; fi23++ {
        switch (fi23 | 127) {
        case 9223372036854775807:
            v2 -= (v3 | 42)
        case 0:
            v2 = v2
        case -451:
            if false {
                fmt.Println(fi23, v2)
            }
        default:
            v2 = v3
        }
        _ = fi23
    }
    for fi24 := 0; fi24 < 10; fi24++ {
        v3 = (v2 | v3)
        v3 -= v2
        _ = fi24
    }
    var v25 int32 = 1
    switch 244 {
    case 127:
        var v26 int8 = 50
        _ = v26
    case -9223372036854775808:
        for fi27 := 0; fi27 < 7; fi27++ {
            fmt.Println(v2, v2)
            var v28 uint8 = 42
            v3 += v2
            _ = fi27
            _ = v28
        }
    }
    v29 := GenS1{F0: 0, F1: "a"}
    for fi30 := 0; fi30 < 6; fi30++ {
        if ((int(v2) == fi30) && (18446744073709551615 <= 132)) {
            v2 += (v2 * 348)
            for fi31 := 0; fi31 < 14; fi31++ {
                v3 = 0
                _ = fi31
            }
            var v32 uint8 = 0
            _ = v32
        } else {
            fmt.Println(v25)
            var v33 uint16 = 0
            _ = v33
        }
        _ = fi30
    }
    v34 := GenS2{F0: true}
    for fi35 := 0; fi35 < 9; fi35++ {
        v36 := GenS3{F0: (^(v3)), F1: (v25 >> ((18446744073709551615) & 63)), F2: (v3 ^ (v2 - v3))}
        _ = fi35
        _ = v36
    }
    var v37 int64 = 3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v34
    _ = v37
}
