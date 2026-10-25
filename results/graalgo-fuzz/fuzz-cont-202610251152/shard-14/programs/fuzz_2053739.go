package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 int64; F2 string }
type GenS3 struct { F0 int }
func genFn0(p0 int16, p1 uint32) int64 {
    p1 = (p1 >> ((42) & 63))
    for fi6 := 0; fi6 < 16; fi6++ {
        p0 -= p0
        for fi7 := 0; fi7 < 3; fi7++ {
            var v8 uint = 0
            var v9 float32 = float32(0.1)
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    return -1
}

func main() {
    var v1 float32 = (1e10 - 1e10)
    var v2 uint16 = 521
    v3 := make(map[string]bool)
    v4 := []uint32{914, 1, 0, 2147483647}
    var v5 bool = (!false)
    v2 -= (v2 * 584)
    if v5 {
        v10 := GenS1{F0: 127}
        _ = v10
    } else {
        for fi11 := 0; fi11 < 6; fi11++ {
            for fi12 := 0; fi12 < 13; fi12++ {
                switch -494 {
                case -406:
                    fmt.Println(fi12, fi11, v5)
                case 9223372036854775807:
                    var v13 bool = v5
                    _ = v13
                }
                var v14 uint16 = v2
                var v15 float32 = v1
                if v5 {
                    v16 := GenS2{F0: fi12, F1: 726, F2: "a"}
                    if v5 {
                        if v5 {
                            if false {
                                var v17 uint8 = uint8(v2)
                                v18 := v3[("foo bar" + "hello")]
                                _ = v17
                                _ = v18
                            } else {
                                var v19 uint64 = 228
                                var v20 int32 = 2147483647
                                _ = len(v3)
                                _ = v19
                                _ = v20
                            }
                            if true {
                                fmt.Println(fi11, v14)
                                if v5 {
                                    fmt.Println(v5)
                                    fmt.Println(fi11, v14, fi12)
                                    var v21 int16 = -453
                                    _ = v21
                                }
                                var v22 int8 = 1
                                _ = v22
                            }
                            var v23 int8 = 42
                            _ = v23
                        } else {
                            _ = len(v4)
                            v15 = 1.5
                        }
                        var v24 uint8 = 94
                        _ = v24
                    } else {
                        var v25 uint32 = 2147483647
                        v26 := GenS3{F0: (fi12 >> ((4294967295) & 63))}
                        fmt.Println(v25, v5)
                        _ = v25
                        _ = v26
                    }
                    _ = v16
                }
                v3[("" + ("\n\t" + "foo bar"))] = (!false)
                _ = fi12
                _ = v14
                _ = v15
            }
            var v27 string = "a"
            _ = fi11
            _ = v27
        }
        v1 = float32(float64(v1))
        fmt.Println(v1, v5)
    }
    v1 -= v1
    var v28 float32 = 0.1
    v4 = append(v4, 0)
    v28 = 3.14159
    var v29 uint64 = 42
    v1 = float32(1.5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
