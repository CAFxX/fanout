package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 uint64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint64 }
func genFn0(p0 int64) int16 {
    fmt.Println(p0)
    v7 := GenS1{F0: 42, F1: 358}
    fmt.Println(p0)
    _ = p0
    _ = v7
    return -32768
}
func genFn1(p0 uint8) uint8 {
    fmt.Println(p0)
    if true {
        switch 128 {
        case -488:
            if ((0.1 / (3.14159 * 3.14159)) < (1.5 + 0.1)) {
                var v8 int64 = 9223372036854775807
                _ = v8
            }
        default:
            if (1 != 42) {
                switch 128 {
                case 300:
                    switch "hello" {
                    case "":
                        v9 := []int{734}
                        _ = v9
                    case "foo bar":
                        var v10 uint16 = 42
                        _ = v10
                    }
                }
            }
        }
        if true {
            var v11 float64 = 0.0
            _ = v11
        } else {
            for fi12 := 0; fi12 < 18; fi12++ {
                v13 := GenS2{F0: uint(p0)}
                for fi14 := 0; fi14 < 4; fi14++ {
                    var v15 float32 = 0.1
                    var v16 int64 = 9223372036854775807
                    var v17 uint16 = 42
                    _ = fi14
                    _ = v15
                    _ = v16
                    _ = v17
                }
                _ = fi12
                _ = v13
            }
            switch "" {
            case "foo bar":
                var v18 int8 = 100
                _ = v18
            case "\n\t":
                switch "" {
                case "foo bar":
                    v19 := []int{}
                    _ = v19
                case "\n\t":
                    for fi20 := 0; fi20 < 3; fi20++ {
                        var v21 uint16 = 42
                        switch (-(fi20)) {
                        case -2147483648:
                            v21 = (v21 << ((2) & 63))
                        }
                        p0 = (p0 / ((42) | 1))
                        _ = fi20
                        _ = v21
                    }
                case "a":
                    v22 := []int{2147483647, -9223372036854775808}
                    _ = v22
                default:
                    fmt.Println(p0)
                }
            case "":
                v23 := []int{4294967295, -2147483648, -930, 1}
                _ = v23
            }
            var v24 uint = 1
            _ = v24
        }
        if (9223372036854775807 != 346) {
            fmt.Println(p0)
        }
    }
    var v25 uint64 = 1
    p0 += (p0 ^ (p0 % (((^(p0))) | 1)))
    _ = p0
    _ = v25
    return 1
}

func main() {
    v1 := []uint32{}
    v2 := []uint32{4294967295, 867, 65535, 590}
    var v3 float32 = (0.1 / -2.25)
    v4 := []int{0, 9223372036854775807, 522}
    v5 := make(map[int]bool)
    v6 := make(map[uint32]string)
    var v26 float64 = 3.14159
    if true {
        v3 = float32(float64((v3 * v3)))
        var v27 int32 = 42
        _ = v27
    }
    for fi28 := 0; fi28 < 1; fi28++ {
        for fi29 := 0; fi29 < 5; fi29++ {
            v30 := GenS3{F0: 2147483647, F1: 4294967295}
            _ = len(v4)
            _ = fi29
            _ = v30
        }
        _ = fi28
    }
    v3 -= v3
    v31 := GenS4{F0: 42}
    var v32 int = 128
    v33 := GenS1{F0: 1, F1: 2}
    var v34 uint8 = 0
    var v35 int = v32
    for fi36 := 0; fi36 < 3; fi36++ {
        for fi37 := 0; fi37 < 6; fi37++ {
            v4 = append(v4, (v32 * (^(v35))))
            _ = fi37
        }
        switch "hello" {
        case "a":
            v38 := GenS3{F0: -1, F1: 0}
            _ = v38
        default:
            for fi39 := 0; fi39 < 4; fi39++ {
                var v40 int32 = -753
                _ = fi39
                _ = v40
            }
        }
        _ = fi36
    }
    for fi41 := 0; fi41 < 17; fi41++ {
        v3 -= (v3 + (float32(float64(((v3 / v3) - 0.0))) + float32(v26)))
        v3 -= (v3 + v3)
        _ = fi41
    }
    if false {
        v42 := GenS5{F0: 0}
        v34 = 127
        var v43 int16 = genFn0(int64(v34))
        _ = v42
        _ = v43
    } else {
        fmt.Println(v3, v34)
    }
    var v44 uint64 = 1
    v2 = append(v2, 1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v44
    _ = v5
    _ = v6
}
