package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint; F1 uint32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint16; F1 uint64; F2 uint32 }
func genFn0(p0 uint32, p1 string, p2 uint16) int {
    p0 += 2147483648
    p2 = (^(p2))
    var v6 float32 = (1.5 + float32(1e10))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 774
}

func main() {
    var v1 int64 = 1
    v2 := []uint32{2147483647}
    v3 := make(map[int]string)
    v4 := make(map[uint32]string)
    var v5 float32 = 0.1
    fmt.Println(v1, v1)
    fmt.Println(v5, v5)
    v7 := GenS1{F0: -128, F1: (255 <= 2), F2: 42}
    v5 -= ((1.5 - v5) / (((v5 - (1e10 - (float32(float64(1e10)) * float32((0.1 / (float64(v5) - 1e10)))))) - v5) / v5))
    v5 -= v5
    fmt.Println(v1)
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 uint32 = 727
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 18; fi10++ {
        switch (fi10 + 32768) {
        case 32768:
            switch fi10 {
            case 981:
                var v11 uint32 = 65535
                _ = v11
            }
        default:
            if false {
                fmt.Println(v5, v1, fi10)
                v5 += v5
            } else {
                v5 = (((v5 / v5) / v5) + float32(1e10))
                for fi12 := 0; fi12 < 17; fi12++ {
                    v1 -= v1
                    switch (^(fi10)) {
                    case 2147483647:
                        switch "" {
                        case "a":
                            var v13 string = "hello"
                            _ = v13
                        case "\n\t":
                            fmt.Println(fi10)
                        }
                    case 128:
                        var v14 bool = false
                        _ = v14
                    case -2147483648:
                        if true {
                            var v15 float32 = (float32(-2.25) - (v5 * float32(3.14159)))
                            v16 := GenS2{F0: 247}
                            if (127 > 1) {
                                fmt.Println(v5, fi10, v5)
                                v17 := GenS3{F0: 1}
                                fmt.Println(v15)
                                v5 += -2.25
                                v18 := GenS2{F0: 2}
                                _ = v17
                                _ = v18
                            }
                            _ = v15
                            _ = v16
                        }
                    default:
                        var v19 int8 = -90
                        _ = v19
                    }
                    _ = fi12
                }
                v5 = 0.1
            }
        }
        v20 := GenS4{F0: 42, F1: 0}
        _ = fi10
        _ = v20
    }
    v5 = float32((1e10 * 1.5))
    fmt.Println(v5, v1)
    for fi21 := 0; fi21 < 18; fi21++ {
        for fi22 := 0; fi22 < 12; fi22++ {
            fmt.Println(fi22, v1)
            switch fi22 {
            case -9223372036854775808:
                switch fi22 {
                case -782:
                    var v23 int16 = 391
                    var v24 int64 = v1
                    _ = v23
                    _ = v24
                case 42:
                    v25 := GenS1{F0: 0, F1: (!false), F2: -128}
                    _ = v25
                }
            }
            _ = fi22
        }
        _ = fi21
    }
    if ((0 < 42) && false) {
        v5 = (v5 / (1e10 / v5))
        v1 = v1
    } else {
        v26 := GenS5{F0: 104}
        for fi27 := 0; fi27 < 13; fi27++ {
            var v28 int64 = 1
            for fi29 := 0; fi29 < 7; fi29++ {
                v30 := GenS6{F0: 4294967295}
                _ = fi29
                _ = v30
            }
            for fi31 := 0; fi31 < 11; fi31++ {
                var v32 uint64 = 42
                v5 = (v5 + float32((float64(1e10) * 0.0)))
                switch "\n\t" {
                case "a":
                    var v33 int = (^(fi27))
                    _ = v33
                case "hello":
                    switch ("a" + "") {
                    case "foo bar":
                        delete(v3, -1)
                    default:
                        v5 += -2.25
                    }
                case "foo bar":
                    _ = len(v2)
                }
                _ = fi31
                _ = v32
            }
            _ = fi27
            _ = v28
        }
        v34 := GenS7{F0: 324, F1: 4294967295, F2: 704}
        _ = v26
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
