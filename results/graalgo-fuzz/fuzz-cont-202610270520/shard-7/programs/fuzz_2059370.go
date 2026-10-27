package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool; F1 int; F2 int }

func main() {
    v1 := make(map[int]uint32)
    v2 := []int{128, 4294967295, -463}
    v3 := make(map[uint32]uint32)
    var v4 uint8 = 255
    v5 := GenS1{F0: 894, F1: (v4 ^ v4)}
    var v6 bool = false
    for fi7 := 0; fi7 < 4; fi7++ {
        switch (fi7 << ((525) & 63)) {
        case -9223372036854775808:
            v4 = v4
        }
        _ = fi7
    }
    if v6 {
        for fi8 := 0; fi8 < 17; fi8++ {
            var v9 int64 = 42
            _ = fi8
            _ = v9
        }
        v6 = true
        v10 := GenS1{F0: 1, F1: v4}
        _ = v10
    } else {
        for fi11 := 0; fi11 < 7; fi11++ {
            v6 = (v6 && v6)
            switch fi11 {
            case 128:
                for fi12 := 0; fi12 < 6; fi12++ {
                    var v13 int = fi11
                    v14 := GenS2{F0: 0}
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            case -2147483648:
                var v15 uint64 = 4294967295
                _ = v15
            case -943:
                switch (^(fi11)) {
                case 2147483647:
                    if v6 {
                        for fi16 := 0; fi16 < 16; fi16++ {
                            var v17 int = (fi11 + 2147483647)
                            fmt.Println(fi16, fi16, fi16)
                            v18 := v1[(fi16 | (^(fi11)))]
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                        var v19 int16 = 1000
                        _ = v19
                    }
                case 127:
                    _ = len(v1)
                case 32767:
                    if (!true) {
                        var v20 float64 = -2.25
                        var v21 uint32 = 195
                        for fi22 := 0; fi22 < 13; fi22++ {
                            v23 := GenS2{F0: 36}
                            v24 := GenS1{F0: v21, F1: (^(v4))}
                            _ = fi22
                            _ = v23
                            _ = v24
                        }
                        _ = v20
                        _ = v21
                    }
                default:
                    fmt.Println(v4, fi11, v4)
                }
            }
            _ = len(v1)
            _ = fi11
        }
        if (2147483647 <= -688) {
            var v25 float32 = 1.5
            v26 := GenS2{F0: 663}
            _ = v25
            _ = v26
        } else {
            fmt.Println(v4, v4)
            v27 := GenS3{F0: true, F1: 2147483647, F2: 127}
            var v28 uint64 = 673
            _ = v27
            _ = v28
        }
    }
    _ = len(v3)
    switch 4294967295 {
    case -765:
        v29 := GenS2{F0: -935}
        _ = v29
    case 127:
        var v30 float64 = (1.5 + float64((1e10 * (0.0 / float32(1e10)))))
        _ = v30
    }
    if (65535 < 42) {
        v4 -= v4
        v6 = (4294967295 != 127)
    } else {
        v31 := GenS1{F0: 4294967295, F1: (v4 & v4)}
        var v32 float64 = (float64(-2.25) * 3.14159)
        _ = v31
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
