package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 uint8; F1 int64; F2 uint8 }
type GenS3 struct { F0 int; F1 int }
func genFn0(p0 int16, p1 uint32) uint64 {
    p1 += (^(p1))
    for fi6 := 0; fi6 < 5; fi6++ {
        p0 += (p0 * p0)
        switch 1 {
        case 713:
            fmt.Println(fi6, fi6)
        case 2147483647:
            switch fi6 {
            case -9223372036854775808:
                v7 := []int{fi6}
                _ = v7
            case 209:
                v8 := []int{}
                _ = v8
            }
        }
        _ = fi6
    }
    fmt.Println(p0, p1)
    var v9 uint = 2
    v9 = (v9 * v9)
    _ = p0
    _ = p1
    _ = v9
    return 2
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]string)
    var v3 bool = true
    v4 := []uint32{}
    v5 := make(map[int]string)
    v10 := GenS1{F0: (255 > 42), F1: 0}
    var v11 int8 = 40
    var v12 string = "a"
    if false {
        var v13 int64 = -9223372036854775808
        if (!(!false)) {
            fmt.Println(v3, v12)
        }
        if v3 {
            v4 = append(v4, 2147483648)
            if (783 > 65535) {
                v14 := v5[127]
                v15 := GenS1{F0: (((v11 >> ((4294967295) & 63)) >= -67) || false), F1: 127}
                _ = v14
                _ = v15
            }
        }
        _ = v13
    } else {
        for fi16 := 0; fi16 < 15; fi16++ {
            for fi17 := 0; fi17 < 3; fi17++ {
                var v18 float32 = 3.14159
                _ = fi17
                _ = v18
            }
            _ = fi16
        }
    }
    v19 := GenS2{F0: 2, F1: 42, F2: 1}
    for fi20 := 0; fi20 < 11; fi20++ {
        fmt.Println(v3, v3)
        var v21 float64 = (float64(0.1) + float64(float32((float64(-2.25) * (float64(0.1) - ((1.5 - 3.14159) + 3.14159))))))
        _ = fi20
        _ = v21
    }
    var v22 uint32 = 2
    for fi23 := 0; fi23 < 19; fi23++ {
        switch (fi23 | fi23) {
        case 128:
            v12 = (v12 + ("" + v12))
        case -158:
            var v24 uint16 = 0
            _ = v24
        }
        v12 = v12
        _ = fi23
    }
    v5[698] = (v12 + v12)
    v25 := GenS1{F0: v3, F1: 0}
    switch v12 {
    case "a":
        for fi26 := 0; fi26 < 8; fi26++ {
            var v27 float64 = 0.0
            _ = fi26
            _ = v27
        }
    case "hello":
        for fi28 := 0; fi28 < 17; fi28++ {
            if (42 >= 329) {
                v12 = ((v12 + v12) + v12)
                fmt.Println(fi28, v3, v11)
                var v29 int32 = -2147483648
                _ = v29
            }
            _ = fi28
        }
    default:
        v30 := GenS3{F0: -864, F1: 32768}
        _ = v30
    }
    v22 = v22
    v31 := GenS1{F0: true, F1: 255}
    var v32 bool = (v3 || true)
    if (-469 < 0) {
        var v33 int = 128
        if (((-2.25 * 3.14159) - 1e10) < float32((1e10 + float64(1.5)))) {
            for fi34 := 0; fi34 < 2; fi34++ {
                var v35 uint = 42
                v11 = (v11 * 0)
                _ = fi34
                _ = v35
            }
            var v36 uint = 4294967295
            var v37 float64 = 0.0
            _ = v36
            _ = v37
        } else {
            if v32 {
                v32 = false
                var v38 int16 = -81
                _ = v38
            } else {
                if (42 <= -866) {
                    v39 := GenS2{F0: 34, F1: 333, F2: 2}
                    var v40 int8 = v11
                    for fi41 := 0; fi41 < 7; fi41++ {
                        v42 := GenS1{F0: true, F1: 42}
                        _ = fi41
                        _ = v42
                    }
                    _ = v39
                    _ = v40
                } else {
                    var v43 int32 = 906
                    for fi44 := 0; fi44 < 18; fi44++ {
                        for fi45 := 0; fi45 < 17; fi45++ {
                            v43 -= v43
                            v43 = (v43 + (v43 % ((v43) | 1)))
                            _ = fi45
                        }
                        if (v11 == v11) {
                            v46 := GenS1{F0: (128 != 127), F1: 127}
                            var v47 int = (fi44 << ((4294967295) & 63))
                            _ = v46
                            _ = v47
                        }
                        var v48 float32 = 0.1
                        _ = fi44
                        _ = v48
                    }
                    for fi49 := 0; fi49 < 5; fi49++ {
                        v3 = v3
                        _ = fi49
                    }
                    _ = v43
                }
                v12 = v12
            }
            for fi50 := 0; fi50 < 12; fi50++ {
                var v51 float64 = float64(((0.0 - 0.1) + ((float32(((0.0 * ((-2.25 * -2.25) * ((((0.1 / 0.1) - (3.14159 / 1.0)) * 1e10) * 3.14159))) + 0.1)) - ((1.5 / float32((1e10 + 1.5))) + 0.1)) - float32(1.5))))
                if len(v4) > 0 {
                    v4[0] = 2147483648
                }
                _ = fi50
                _ = v51
            }
        }
        _ = v33
    }
    v52 := v2[v22]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v22
    _ = v25
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v52
}
