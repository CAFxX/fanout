package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 int8 }
type GenS2 struct { F0 uint32; F1 int16; F2 int64 }
type GenS3 struct { F0 bool; F1 uint16 }
func genFn0(p0 int32, p1 string) int64 {
    var v5 int = -1
    for fi6 := 0; fi6 < 17; fi6++ {
        v7 := GenS1{F0: -1, F1: (-(v5)), F2: 0}
        fmt.Println(p1, v5)
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v5
    return 1
}
func genFn1() int32 {
    if true {
        var v8 uint = 92
        v9 := []int{-950, 797, 0, 9223372036854775807}
        _ = v8
        _ = v9
    } else {
        v10 := GenS2{F0: 429, F1: -32768, F2: 42}
        _ = v10
    }
    var v11 int32 = 0
    v11 -= (v11 % (((-(v11))) | 1))
    _ = v11
    return 42
}

func main() {
    var v1 int16 = 42
    var v2 float64 = (float64(((1e10 - 1e10) * ((0.0 - ((0.0 + 0.1) * 1e10)) * 1.5))) + 1e10)
    var v3 bool = false
    var v4 int32 = -38
    if (!v3) {
        var v12 uint8 = 128
        for fi13 := 0; fi13 < 8; fi13++ {
            var v14 int8 = 100
            _ = fi13
            _ = v14
        }
        _ = v12
    }
    v1 = (v1 ^ (v1 + v1))
    var v15 int16 = 780
    v3 = v3
    var v16 uint64 = 426
    var v17 uint64 = 328
    var v18 float32 = 0.1
    fmt.Println(v17, v15, v4)
    var v19 int8 = 100
    v2 += float64(v18)
    if true {
        if v3 {
            var v20 uint16 = 255
            v21 := []int{int(v4), 414, 4294967295, 32767}
            _ = v20
            _ = v21
        } else {
            if ((100 >= 1) || (1 < -1)) {
                v15 += v1
            } else {
                if (!(4294967295 < 65535)) {
                    if false {
                        var v22 uint64 = (v17 & v17)
                        if true {
                            v2 -= float64(0.0)
                        } else {
                            if (9223372036854775807 < 0) {
                                v23 := []int{}
                                fmt.Println(v1)
                                _ = v23
                            } else {
                                v24 := GenS3{F0: false, F1: 5}
                                _ = v24
                            }
                            if (v2 == v2) {
                                var v25 string = "foo bar"
                                _ = v25
                            } else {
                                var v26 string = "foo bar"
                                _ = v26
                            }
                        }
                        var v27 uint = 0
                        _ = v22
                        _ = v27
                    }
                    v28 := []int{int(v1), 4294967295, 127, -648}
                    _ = v28
                }
                for fi29 := 0; fi29 < 18; fi29++ {
                    var v30 uint64 = (v16 | v17)
                    v31 := []int{}
                    _ = fi29
                    _ = v30
                    _ = v31
                }
                var v32 int16 = -32768
                _ = v32
            }
            v17 = 2
            var v33 int32 = (v4 % ((2147483647) | 1))
            _ = v33
        }
        var v34 bool = (v19 < int8(v4))
        _ = v34
    } else {
        var v35 int32 = (^(v4))
        _ = v35
    }
    v36 := []int{9223372036854775807, 2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v36
    _ = v4
}
