package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 uint }
func genFn0() int32 {
    var v7 int64 = 42
    if true {
        for fi8 := 0; fi8 < 17; fi8++ {
            for fi9 := 0; fi9 < 19; fi9++ {
                fmt.Println(fi8, v7, v7)
                var v10 int16 = 1
                if true {
                    v7 = (v7 / ((-104) | 1))
                }
                _ = fi9
                _ = v10
            }
            v11 := GenS1{F0: 1}
            _ = fi8
            _ = v11
        }
        v12 := []int{-9223372036854775808, 200, -9223372036854775808}
        if (1.5 != -2.25) {
            if true {
                var v13 uint64 = 2
                _ = v13
            }
        } else {
            fmt.Println(v7)
            if len(v12) > 0 {
                v12[0] = 127
            }
            if false {
                var v14 float64 = float64(0.1)
                _ = v14
            }
        }
        _ = v12
    } else {
        v7 += -9223372036854775808
    }
    v7 = (v7 ^ (-(v7)))
    v15 := GenS1{F0: 42}
    var v16 int64 = (v7 % ((42) | 1))
    _ = v15
    _ = v16
    _ = v7
    return -744
}

func main() {
    var v1 int32 = -266
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]int)
    var v5 int = -1
    var v6 int32 = (v1 << ((42) & 63))
    var v17 bool = (true || false)
    if (0 == 2) {
        v6 += (v6 >> ((255) & 63))
        v18 := []int{(v5 ^ (^(v5))), (v5 % ((v5) | 1)), (v5 & v5)}
        var v19 int64 = -410
        _ = v18
        _ = v19
    } else {
        v6 -= v6
        v20 := GenS1{F0: 0}
        _ = v20
    }
    v21 := GenS2{F0: 2147483647, F1: v17}
    var v22 uint = 1
    var v23 int64 = 9223372036854775807
    v23 = v23
    if true {
        v6 -= v6
        if ((259 > 100) || (((3.14159 * 0.0) * (-2.25 * (((3.14159 * (0.1 / (1e10 / 1.5))) * 0.0) * 0.1))) >= ((0.0 / 0.1) * 3.14159))) {
            for fi24 := 0; fi24 < 16; fi24++ {
                for fi25 := 0; fi25 < 4; fi25++ {
                    var v26 float64 = float64(0.0)
                    if (fi25 > (v5 + -2147483648)) {
                        v27 := []int{}
                        v28 := GenS2{F0: 2147483647, F1: false}
                        var v29 uint = (v22 | 18446744073709551615)
                        var v30 float64 = v26
                        _ = v27
                        _ = v28
                        _ = v29
                        _ = v30
                    } else {
                        v31 := v2["foo bar"]
                        fmt.Println(v1)
                        _ = len(v3)
                        v32 := []int{v5, fi25, (fi24 << ((v22) & 63)), (fi24 / ((fi25) | 1))}
                        fmt.Println(v5, fi25, v31)
                        v26 -= float64(float32(float64((3.14159 * (1e10 / 0.1)))))
                        _ = v31
                        _ = v32
                    }
                    v22 -= (v22 / (((v22 | 938)) | 1))
                    _ = fi25
                    _ = v26
                }
                var v33 int16 = 42
                _ = fi24
                _ = v33
            }
            for fi34 := 0; fi34 < 17; fi34++ {
                v23 += (v23 << ((729) & 63))
                _ = fi34
            }
        }
        switch v5 {
        case 632:
            var v35 int16 = 0
            _ = v35
        }
    }
    v6 -= (^(v6))
    v36 := GenS3{F0: v22}
    v5 += (v5 % ((v5) | 1))
    v37 := []int{784}
    if v17 {
        v38 := GenS2{F0: 65535, F1: v17}
        v5 -= (v5 | 42)
        var v39 int = -198
        _ = v38
        _ = v39
    }
    v17 = (0.0 >= 0.0)
    if false {
        var v40 uint16 = 253
        var v41 uint = v22
        _ = v40
        _ = v41
    }
    for fi42 := 0; fi42 < 9; fi42++ {
        if v17 {
            v23 += 42
            var v43 float64 = 1.5
            _ = v43
        }
        switch "\n\t" {
        case "a":
            _ = len(v2)
        case "hello":
            _ = len(v37)
        case "\n\t":
            switch v5 {
            case 460:
                v23 -= (-(v23))
            default:
                var v44 int = fi42
                _ = v44
            }
        default:
            v17 = true
        }
        _ = len(v37)
        _ = fi42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
