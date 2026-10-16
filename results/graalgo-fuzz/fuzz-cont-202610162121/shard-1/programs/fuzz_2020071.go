package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 int16 = 1000
    var v2 int16 = (v1 & (v1 - v1))
    var v3 int8 = -1
    var v4 int = -698
    v5 := make(map[int]uint32)
    v6 := []int{v4, v4, (v4 / (((v4 ^ (v4 * (-(v4))))) | 1)), v4}
    var v7 uint8 = 0
    var v8 string = ""
    for fi9 := 0; fi9 < 13; fi9++ {
        if true {
            fmt.Println(v4, v2, fi9)
            if (-2.25 >= (-2.25 - (1e10 + 0.0))) {
                if (!(!true)) {
                    switch 2147483647 {
                    case -1:
                        var v10 int8 = v3
                        _ = v10
                    case 508:
                        if false {
                            var v11 int64 = 0
                            switch (fi9 >> ((18446744073709551615) & 63)) {
                            case -2147483648:
                                for fi12 := 0; fi12 < 4; fi12++ {
                                    v13 := GenS1{F0: v11, F1: 314, F2: v4}
                                    _ = fi12
                                    _ = v13
                                }
                            case 32768:
                                var v14 uint32 = uint32(v7)
                                _ = v14
                            case 127:
                                var v15 float64 = 0.0
                                _ = v15
                            default:
                                for fi16 := 0; fi16 < 18; fi16++ {
                                    var v17 int32 = -227
                                    switch 32768 {
                                    case 127:
                                        v4 += fi9
                                    case -777:
                                        v18 := GenS2{F0: 563}
                                        _ = v18
                                    case 2147483647:
                                        fmt.Println(v8, v3, v7)
                                    }
                                    v17 += v17
                                    _ = fi16
                                    _ = v17
                                }
                            }
                            v6 = append(v6, v4)
                            _ = v11
                        } else {
                            var v19 int32 = 2147483647
                            v20 := v5[fi9]
                            _ = v19
                            _ = v20
                        }
                    case 32768:
                        var v21 int8 = -1
                        _ = v21
                    }
                    v5[(fi9 << ((255) & 63))] = 957
                }
                var v22 bool = false
                _ = v22
            } else {
                v1 += 918
                var v23 int8 = v3
                _ = v23
            }
            var v24 uint16 = 2
            _ = v24
        }
        var v25 int8 = (v3 | (v3 ^ -100))
        _ = fi9
        _ = v25
    }
    v26 := GenS3{F0: 242}
    v2 += v2
    v6 = append(v6, -9223372036854775808)
    v7 += 61
    for fi27 := 0; fi27 < 8; fi27++ {
        var v28 float64 = 1.5
        var v29 uint8 = v7
        fmt.Println(v8)
        _ = fi27
        _ = v28
        _ = v29
    }
    var v30 string = "foo bar"
    for fi31 := 0; fi31 < 4; fi31++ {
        var v32 int64 = -1
        v6 = append(v6, fi31)
        if (1e10 < float64(0.0)) {
            for fi33 := 0; fi33 < 15; fi33++ {
                fmt.Println(v8, v2)
                _ = fi33
            }
        }
        _ = fi31
        _ = v32
    }
    switch (v4 % (((v4 >> ((18446744073709551615) & 63))) | 1)) {
    case 127:
        for fi34 := 0; fi34 < 10; fi34++ {
            if false {
                v7 -= 0
            }
            _ = fi34
        }
    case -9223372036854775808:
        var v35 float32 = 0.0
        _ = v35
    case -328:
        if true {
            var v36 uint16 = 32767
            fmt.Println(v2, v8, v1)
            _ = v36
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
