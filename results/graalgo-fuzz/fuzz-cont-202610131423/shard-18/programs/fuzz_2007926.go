package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 int8 }
func genFn0(p0 uint32, p1 uint32, p2 uint) uint32 {
    var v6 int = 122
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 float32 = (1e10 * (3.14159 * (0.0 * -2.25)))
        if (0 >= 1000) {
            var v9 int32 = -2147483648
            var v10 uint = (^(p2))
            for fi11 := 0; fi11 < 8; fi11++ {
                if false {
                    if false {
                        v12 := []int{fi7, -9223372036854775808}
                        switch (fi7 | (^(v6))) {
                        case 724:
                            var v13 int8 = -1
                            _ = v13
                        case 49:
                            var v14 int8 = 42
                            var v15 float64 = 1e10
                            if len(v12) > 0 {
                                v12[0] = (fi7 / ((v6) | 1))
                            }
                            _ = v14
                            _ = v15
                        }
                        _ = v12
                    }
                } else {
                    v16 := GenS1{F0: 1000, F1: v10}
                    if false {
                        if true {
                            var v17 uint16 = 2
                            _ = v17
                        }
                        var v18 int32 = v9
                        v10 = 42
                        _ = v18
                    } else {
                        v19 := []int{32767}
                        _ = v19
                    }
                    v20 := []int{(-(fi11)), fi11}
                    var v21 int64 = 489
                    _ = v16
                    _ = v20
                    _ = v21
                }
                fmt.Println(v8)
                var v22 int64 = 9223372036854775807
                _ = fi11
                _ = v22
            }
            _ = v10
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    p2 += 897
    for fi23 := 0; fi23 < 18; fi23++ {
        var v24 string = "hello"
        p1 = (p1 >> ((255) & 63))
        var v25 int8 = 40
        _ = fi23
        _ = v24
        _ = v25
    }
    for fi26 := 0; fi26 < 4; fi26++ {
        var v27 uint64 = 648
        fmt.Println(p0)
        _ = fi26
        _ = v27
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p0
}
func genFn1() int32 {
    switch -350 {
    case 42:
        v28 := GenS2{F0: 1}
        _ = v28
    }
    v29 := []int{-633, -290, 951}
    var v30 uint16 = 65535
    _ = v29
    _ = v30
    return 195
}

func main() {
    var v1 int = 0
    v2 := make(map[uint32]int)
    v3 := []int{v1, (v1 >> ((147) & 63)), -1}
    var v4 float32 = ((-2.25 + 0.0) - 0.1)
    var v5 uint64 = 42
    v4 += float32(1e10)
    var v31 uint8 = 1
    v2[genFn0(2147483648, genFn0(42, uint32(v5), 1), 2)] = v1
    if false {
        v31 = 142
    }
    for fi32 := 0; fi32 < 2; fi32++ {
        switch fi32 {
        case 717:
            v4 = v4
        case 32767:
            var v33 uint32 = 511
            _ = v33
        }
        fmt.Println(v4, v1, fi32)
        _ = fi32
    }
    v2[1] = 131
    v31 = v31
    v2[11] = 413
    for fi34 := 0; fi34 < 4; fi34++ {
        for fi35 := 0; fi35 < 6; fi35++ {
            v4 -= v4
            if (-128 == 127) {
                var v36 uint16 = 65535
                _ = v36
            }
            _ = fi35
        }
        _ = fi34
    }
    if len(v3) > 0 {
        v3[0] = (v1 & v1)
    }
    v4 -= v4
    var v37 int16 = 42
    v1 -= (-(v1))
    var v38 int64 = int64(v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
