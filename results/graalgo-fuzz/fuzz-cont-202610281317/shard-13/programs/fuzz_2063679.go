package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 int8; F1 int16; F2 int32 }
type GenS3 struct { F0 int64; F1 uint64; F2 int64 }
type GenS4 struct { F0 int64 }
func genFn0(p0 int16, p1 int16) int32 {
    for fi5 := 0; fi5 < 18; fi5++ {
        p0 = p1
        _ = fi5
    }
    switch -706 {
    case -651:
        for fi6 := 0; fi6 < 15; fi6++ {
            v7 := []int{(fi6 / ((fi6) | 1)), -1, (fi6 + (fi6 | (fi6 - (fi6 & (fi6 ^ 42)))))}
            _ = fi6
            _ = v7
        }
    }
    _ = p0
    _ = p1
    return 1
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]uint32)
    var v3 uint32 = 4294967295
    var v4 float32 = (1.5 / 0.1)
    if false {
        var v8 int64 = -507
        var v9 string = "\n\t"
        _ = v8
        _ = v9
    } else {
        if (true && (false || true)) {
            if (true && (v3 != (v3 * 65535))) {
                if true {
                    if (!((v3 * v3) <= 2)) {
                        var v10 float64 = (0.1 + float64(3.14159))
                        if false {
                            v11 := GenS1{F0: 42, F1: -1}
                            _ = v11
                        }
                        var v12 int8 = 0
                        _ = v10
                        _ = v12
                    }
                } else {
                    fmt.Println(v3)
                    v3 -= (v3 & (v3 | (v3 ^ v3)))
                }
                v3 -= (v3 | v3)
                v13 := []int{1, -792, -784, -576}
                _ = v13
            } else {
                for fi14 := 0; fi14 < 14; fi14++ {
                    v3 -= (v3 & v3)
                    _ = fi14
                }
                for fi15 := 0; fi15 < 3; fi15++ {
                    fmt.Println(v4)
                    v16 := []int{(fi15 | (fi15 >> ((2) & 63))), (^(fi15)), fi15}
                    v17 := GenS1{F0: 0, F1: -251}
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                v18 := GenS1{F0: 127, F1: 0}
                _ = v18
            }
        } else {
            v19 := GenS2{F0: 1, F1: -107, F2: 2147483647}
            _ = v19
        }
    }
    v20 := []int{-2147483648}
    v3 -= v3
    if (!true) {
        v21 := GenS3{F0: 42, F1: 101, F2: -9223372036854775808}
        for fi22 := 0; fi22 < 18; fi22++ {
            for fi23 := 0; fi23 < 9; fi23++ {
                fmt.Println(fi22, v3)
                fmt.Println(fi22, fi22, fi23)
                _ = fi23
            }
            var v24 float32 = 0.0
            _ = fi22
            _ = v24
        }
        _ = v21
    }
    v3 = (v3 + v3)
    var v25 float32 = v4
    if false {
        if false {
            switch 188 {
            case 42:
                var v26 float64 = ((float64(((((-2.25 + ((float32(1e10) + (float32((3.14159 / -2.25)) * -2.25)) * 0.1)) + ((1e10 - (v4 - (1e10 / (0.1 + v25)))) * (v25 * v25))) * v25) + -2.25)) + (0.1 / 1.5)) + 1e10)
                _ = v26
            case 128:
                fmt.Println(v3, v3)
            }
            var v27 float32 = float32(1e10)
            fmt.Println(v3, v27)
            _ = v27
        } else {
            var v28 int8 = 53
            _ = v28
        }
    } else {
        v3 += v3
        v29 := GenS3{F0: 42, F1: 507, F2: -152}
        _ = v29
    }
    var v30 int32 = 2147483647
    delete(v1, 42)
    var v31 uint64 = 2
    switch "foo bar" {
    case "\n\t":
        v32 := GenS2{F0: -10, F1: 1, F2: v30}
        _ = v32
    case "foo bar":
        v2[v3] = v3
    }
    for fi33 := 0; fi33 < 3; fi33++ {
        for fi34 := 0; fi34 < 12; fi34++ {
            v31 = uint64(v3)
            _ = fi34
        }
        v31 = (v31 * uint64(v3))
        _ = fi33
    }
    v35 := GenS4{F0: -9223372036854775808}
    v25 += v4
    v4 += float32(0.1)
    if false {
        fmt.Println(v25, v30)
    }
    if true {
        var v36 int = 0
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v4
}
