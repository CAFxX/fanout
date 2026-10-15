package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 int32; F2 uint64 }
type GenS3 struct { F0 int32; F1 int }
func genFn0(p0 int32) uint32 {
    p0 += (-(p0))
    v6 := GenS1{F0: 2}
    p0 -= (-(p0))
    _ = p0
    _ = v6
    return 347
}
func genFn1() uint32 {
    var v7 uint32 = 65535
    var v8 uint16 = 65535
    _ = v7
    _ = v8
    return genFn0(-2147483648)
}

func main() {
    var v1 uint8 = 2
    var v2 float64 = 0.1
    v3 := make(map[string]bool)
    var v4 uint = 42
    var v5 uint = (v4 % ((255) | 1))
    var v9 int64 = -25
    v10 := []int{485, 194, 0}
    if (196 >= 131) {
        v2 += v2
        for fi11 := 0; fi11 < 12; fi11++ {
            v12 := GenS1{F0: 1}
            var v13 uint = (^(v5))
            var v14 uint8 = v1
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        if false {
            fmt.Println(v1, v1)
            fmt.Println(v2)
        } else {
            v15 := GenS1{F0: 830}
            for fi16 := 0; fi16 < 11; fi16++ {
                if (false || ((v4 ^ (v5 / ((v4) | 1))) > (v5 % ((v5) | 1)))) {
                    var v17 int16 = -899
                    v18 := GenS1{F0: 148}
                    _ = v17
                    _ = v18
                } else {
                    v9 -= v9
                    delete(v3, "\n\t")
                    switch (fi16 | (fi16 / (((^(fi16))) | 1))) {
                    case 0:
                        v19 := GenS2{F0: 42, F1: 0, F2: 2}
                        _ = v19
                    case 9223372036854775807:
                        var v20 uint64 = 539
                        _ = v20
                    case 4294967295:
                        delete(v3, "a")
                    }
                }
                _ = fi16
            }
            fmt.Println(v5, v5, v5)
            _ = v15
        }
    } else {
        var v21 int64 = (-(v9))
        v22 := GenS1{F0: 141}
        fmt.Println(v2)
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 10; fi23++ {
        for fi24 := 0; fi24 < 18; fi24++ {
            var v25 uint16 = 0
            var v26 float64 = float64(1.5)
            var v27 int8 = -128
            _ = fi24
            _ = v25
            _ = v26
            _ = v27
        }
        _ = fi23
    }
    v1 = (v1 / ((v1) | 1))
    v28 := GenS2{F0: -360, F1: 42, F2: 1}
    v4 = 18446744073709551615
    v29 := GenS3{F0: 0, F1: 1}
    for fi30 := 0; fi30 < 15; fi30++ {
        for fi31 := 0; fi31 < 20; fi31++ {
            fmt.Println(v4, v9)
            var v32 int = -314
            var v33 uint8 = v1
            var v34 uint = uint(v33)
            var v35 int = 556
            _ = fi31
            _ = v32
            _ = v33
            _ = v34
            _ = v35
        }
        _ = fi30
    }
    var v36 uint64 = 0
    for fi37 := 0; fi37 < 12; fi37++ {
        fmt.Println(v4, v4, v1)
        _ = fi37
    }
    v9 += (v9 / (((v9 & (v9 * v9))) | 1))
    if true {
        var v38 uint32 = 237
        _ = v38
    }
    fmt.Println(v9)
    for fi39 := 0; fi39 < 19; fi39++ {
        if true {
            for fi40 := 0; fi40 < 4; fi40++ {
                var v41 uint8 = (v1 >> ((v5) & 63))
                fmt.Println(fi39, v9, v41)
                var v42 float64 = (3.14159 / 1.5)
                _ = fi40
                _ = v41
                _ = v42
            }
        }
        _ = fi39
    }
    if true {
        v36 += (^(v36))
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v9
}
