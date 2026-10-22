package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 int8; F1 int64; F2 int8 }
func genFn0(p0 int8, p1 string, p2 int32) uint8 {
    v7 := []int{0}
    if (!true) {
        p1 = p1
        var v8 uint64 = 4294967295
        var v9 float32 = -2.25
        _ = v8
        _ = v9
    }
    var v10 uint32 = 928
    var v11 string = "hello"
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v7
    return 127
}

func main() {
    var v1 uint64 = 0
    var v2 int8 = 1
    var v3 float32 = float32(3.14159)
    var v4 bool = (-358 > 635)
    var v5 uint32 = 2
    var v6 int16 = -1
    var v12 int16 = (v6 >> ((uint(v5)) & 63))
    var v13 int32 = 633
    var v14 uint32 = v5
    var v15 uint8 = 22
    var v16 bool = v4
    switch int(v2) {
    case 32768:
        var v17 float32 = v3
        _ = v17
    case 1:
        v3 += v3
    default:
        var v18 int16 = v12
        _ = v18
    }
    v19 := GenS1{F0: v1}
    for fi20 := 0; fi20 < 19; fi20++ {
        switch "\n\t" {
        case "foo bar":
            for fi21 := 0; fi21 < 7; fi21++ {
                var v22 float64 = 1.5
                var v23 bool = (false || (false && (-62 >= (v6 >> ((20) & 63)))))
                _ = fi21
                _ = v22
                _ = v23
            }
        }
        _ = fi20
    }
    v2 = (-(v2))
    var v24 float32 = float32(0.1)
    v25 := GenS2{F0: v14, F1: -9223372036854775808}
    v26 := []int{32768, 4294967295, -756}
    if (!(((v2 ^ v2) == v2) && v4)) {
        var v27 int16 = (v6 | (v6 * v12))
        _ = v27
    } else {
        for fi28 := 0; fi28 < 1; fi28++ {
            fmt.Println(v15, fi28)
            var v29 int = (fi28 / ((fi28) | 1))
            _ = fi28
            _ = v29
        }
        v2 -= (v2 | v2)
    }
    if (!(100 >= (v2 >> ((65535) & 63)))) {
        v26 = append(v26, int(v13))
        v5 = v14
        v24 += v24
    }
    switch -9223372036854775808 {
    case 32767:
        if false {
            v3 += (v24 / v3)
            for fi30 := 0; fi30 < 1; fi30++ {
                if ((-(v2)) != (v2 % ((42) | 1))) {
                    var v31 int32 = v13
                    if v4 {
                        if false {
                            if (v6 == 573) {
                                v12 += (v12 / (((v6 >> ((255) & 63))) | 1))
                                var v32 int8 = -1
                                v1 = (v1 % ((0) | 1))
                                _ = v32
                            } else {
                                for fi33 := 0; fi33 < 6; fi33++ {
                                    _ = len(v26)
                                    _ = fi33
                                }
                            }
                            var v34 bool = v16
                            _ = v34
                        }
                    }
                    _ = v31
                }
                for fi35 := 0; fi35 < 5; fi35++ {
                    var v36 float32 = 0.0
                    v37 := GenS1{F0: (v1 ^ v1)}
                    var v38 int32 = int32(v2)
                    _ = fi35
                    _ = v36
                    _ = v37
                    _ = v38
                }
                _ = fi30
            }
            for fi39 := 0; fi39 < 14; fi39++ {
                v14 += 295
                v40 := GenS3{F0: 0, F1: 297, F2: -128}
                v1 += (v1 << ((0) & 63))
                _ = fi39
                _ = v40
            }
        }
    case -126:
        fmt.Println(v4)
    case 601:
        if len(v26) >= 2 {
            v26 = v26[1:len(v26)]
        }
    }
    var v41 int64 = 42
    v24 += v24
    v26 = append(v26, 762)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
