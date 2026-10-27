package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 string; F1 int32; F2 uint16 }
type GenS4 struct { F0 uint32 }
func genFn0(p0 bool, p1 int8) uint64 {
    var v7 float32 = 0.0
    v7 = ((1e10 * 0.0) / ((3.14159 + 0.1) / v7))
    var v8 float64 = 0.0
    if (0.0 <= v7) {
        if false {
            fmt.Println(p1)
            if (int32(p1) != 42) {
                switch 449 {
                case 127:
                    v9 := GenS1{F0: 127}
                    _ = v9
                default:
                    var v10 float32 = ((v7 - v7) / (v7 - (v7 + 3.14159)))
                    _ = v10
                }
                var v11 bool = p0
                if v11 {
                    v7 -= v7
                    for fi12 := 0; fi12 < 6; fi12++ {
                        var v13 uint32 = 0
                        var v14 uint = 534
                        if (v7 >= ((1.5 * v7) * float32(v8))) {
                            p1 += 42
                            for fi15 := 0; fi15 < 10; fi15++ {
                                v16 := []int{}
                                _ = fi15
                                _ = v16
                            }
                        } else {
                            if v11 {
                                for fi17 := 0; fi17 < 18; fi17++ {
                                    var v18 uint64 = 111
                                    var v19 uint8 = 195
                                    v20 := []int{-160, -1, 9223372036854775807}
                                    _ = fi17
                                    _ = v18
                                    _ = v19
                                    _ = v20
                                }
                                v14 -= (v14 - 575)
                                v14 += (v14 >> ((uint(v13)) & 63))
                            }
                        }
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                } else {
                    p1 = (p1 >> ((367) & 63))
                }
                _ = v11
            }
            fmt.Println(v7)
        }
        if true {
            v21 := GenS2{F0: 2, F1: 9223372036854775807}
            _ = v21
        }
    } else {
        var v22 uint = 255
        var v23 float32 = v7
        _ = v22
        _ = v23
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 447
}

func main() {
    v1 := []uint32{}
    var v2 uint8 = 1
    var v3 string = "a"
    var v4 uint = 1
    var v5 bool = true
    var v6 uint8 = 1
    v1 = append(v1, 4294967295)
    fmt.Println(v6)
    v5 = (v5 && v5)
    var v24 int64 = 9223372036854775807
    v25 := GenS1{F0: (v2 & 145)}
    var v26 uint8 = 1
    v27 := GenS3{F0: v3, F1: -1, F2: 1}
    if v5 {
        v4 = v4
        if false {
            fmt.Println(v24, v26)
            var v28 uint64 = 1
            var v29 float32 = float32(-2.25)
            _ = v28
            _ = v29
        } else {
            switch 9223372036854775807 {
            case 4294967295:
                for fi30 := 0; fi30 < 12; fi30++ {
                    var v31 uint = (v4 ^ 65535)
                    var v32 int = 482
                    v33 := GenS4{F0: 2147483648}
                    _ = fi30
                    _ = v31
                    _ = v32
                    _ = v33
                }
            }
            v34 := GenS4{F0: 482}
            _ = v34
        }
        var v35 uint64 = 1
        _ = v35
    } else {
        if (!((("a" + v3) + ("\n\t" + v3)) <= v3)) {
            var v36 float64 = float64(0.0)
            v1 = append(v1, 4294967295)
            var v37 uint = 65535
            _ = v36
            _ = v37
        } else {
            if len(v1) > 0 {
                v1[0] = 0
            }
        }
        _ = len(v1)
        var v38 uint = 65535
        _ = v38
    }
    v2 = v2
    var v39 int64 = v24
    v6 -= (v2 << ((v4) & 63))
    v6 += (^(v2))
    if (56 >= 42) {
        _ = len(v1)
        fmt.Println(v26, v5, v5)
    }
    for fi40 := 0; fi40 < 17; fi40++ {
        for fi41 := 0; fi41 < 2; fi41++ {
            v4 = 0
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if (true || (v5 || (true || true))) {
                var v42 uint32 = 104
                _ = v42
            }
            _ = fi41
        }
        _ = fi40
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
