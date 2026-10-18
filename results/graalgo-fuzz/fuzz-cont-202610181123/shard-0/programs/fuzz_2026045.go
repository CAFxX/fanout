package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 string }
func genFn0(p0 uint16, p1 bool) uint8 {
    if p1 {
        v5 := GenS1{F0: 0, F1: 2147483647, F2: ("foo bar" + "\n\t")}
        if (!(-32768 > 75)) {
            v6 := []int{int(p0), 32768, -9223372036854775808, -9223372036854775808}
            p1 = p1
            v7 := GenS1{F0: 2, F1: 32768, F2: ""}
            _ = v6
            _ = v7
        } else {
            var v8 uint = 341
            for fi9 := 0; fi9 < 8; fi9++ {
                var v10 int16 = 0
                _ = fi9
                _ = v10
            }
            if p1 {
                var v11 uint = (v8 << ((v8) & 63))
                _ = v11
            }
            _ = v8
        }
        v12 := []int{0, 2147483647, 2147483647}
        _ = v12
        _ = v5
    } else {
        var v13 int16 = -1
        _ = v13
    }
    for fi14 := 0; fi14 < 7; fi14++ {
        for fi15 := 0; fi15 < 10; fi15++ {
            fmt.Println(fi15, fi14)
            if p1 {
                v16 := []int{fi15, (fi15 - (fi14 * (fi15 * fi14))), fi15, 382}
                _ = v16
            } else {
                v17 := []int{fi14, (fi14 | fi15)}
                if (-559 < -332) {
                    fmt.Println(fi15, fi14)
                    if false {
                        var v18 float32 = 0.1
                        if len(v17) > 0 {
                            v17[0] = 127
                        }
                        var v19 uint32 = 624
                        _ = v18
                        _ = v19
                    } else {
                        var v20 float64 = (1.5 + float64(0.0))
                        var v21 int32 = -120
                        if len(v17) >= 2 {
                            v17 = v17[1:len(v17)]
                        }
                        _ = v20
                        _ = v21
                    }
                    if p1 {
                        if true {
                            v22 := GenS1{F0: 4294967295, F1: fi14, F2: "hello"}
                            _ = v22
                        }
                        p0 = p0
                    }
                }
                _ = v17
            }
            _ = fi15
        }
        _ = fi14
    }
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    return 1
}
func genFn1(p0 int8, p1 int64, p2 uint) uint32 {
    p1 = (p1 / ((int64(p0)) | 1))
    var v23 float32 = float32(0.1)
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    return 42
}

func main() {
    v1 := []uint32{65535, 4294967295, 65535, 2}
    v2 := make(map[string]bool)
    v3 := make(map[int]string)
    var v4 int16 = 1
    v4 += v4
    for fi24 := 0; fi24 < 12; fi24++ {
        fmt.Println(fi24, fi24)
        v4 = (v4 | v4)
        _ = fi24
    }
    fmt.Println(v4)
    var v25 uint16 = 65535
    v4 = (^(v4))
    var v26 uint16 = v25
    if false {
        for fi27 := 0; fi27 < 3; fi27++ {
            v26 += v26
            _ = fi27
        }
    }
    fmt.Println(v4)
    for fi28 := 0; fi28 < 15; fi28++ {
        if true {
            v26 = (v26 << ((255) & 63))
        }
        if len(v1) > 0 {
            v1[0] = 65535
        }
        _ = fi28
    }
    for fi29 := 0; fi29 < 16; fi29++ {
        for fi30 := 0; fi30 < 4; fi30++ {
            if true {
                var v31 int = (fi30 | (^(fi30)))
                var v32 int8 = 69
                v31 += fi29
                _ = v31
                _ = v32
            }
            _ = fi30
        }
        _ = fi29
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
