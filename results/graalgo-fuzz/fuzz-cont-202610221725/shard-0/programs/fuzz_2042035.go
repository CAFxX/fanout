package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int8; F1 int16; F2 int8 }
func genFn0(p0 uint16, p1 int) uint32 {
    v7 := GenS1{F0: 0}
    if (!true) {
        p1 = p1
    } else {
        for fi8 := 0; fi8 < 18; fi8++ {
            var v9 uint16 = 0
            _ = fi8
            _ = v9
        }
        p0 = 754
        fmt.Println(p0)
    }
    v10 := []int{(-(p1)), -605, (^(p1)), p1}
    var v11 int64 = 9223372036854775807
    var v12 uint16 = (p0 * p0)
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    _ = v7
    return 2147483648
}
func genFn1(p0 uint64) uint64 {
    if (1.5 <= 3.14159) {
        var v13 uint16 = 1
        _ = v13
    } else {
        var v14 float32 = 1e10
        _ = v14
    }
    v15 := []int{2147483647, -999}
    _ = p0
    _ = v15
    return 264
}

func main() {
    v1 := []string{}
    v2 := []uint32{}
    v3 := make(map[string]uint32)
    v4 := make(map[int]bool)
    var v5 int = 2147483647
    var v6 int64 = 42
    var v16 uint = 776
    var v17 float32 = 1.5
    var v18 float32 = v17
    fmt.Println(v16, v17, v16)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch "\n\t" {
    case "foo bar":
        for fi19 := 0; fi19 < 1; fi19++ {
            v16 += (v16 - v16)
            var v20 uint64 = 1
            _ = fi19
            _ = v20
        }
    case "\n\t":
        v21 := GenS2{F0: v5, F1: -11}
        _ = v21
    case "":
        for fi22 := 0; fi22 < 5; fi22++ {
            v6 = (v6 * 9223372036854775807)
            _ = fi22
        }
    default:
        v6 = (v6 & (v6 + (v6 << ((2) & 63))))
    }
    v5 = (v5 >> ((v16) & 63))
    var v23 uint16 = 65535
    if (255 >= (^(v16))) {
        if (true && true) {
            switch (^(v5)) {
            case 2147483647:
                v6 += -9223372036854775808
            case -489:
                v5 -= -9223372036854775808
            default:
                switch (v5 | v5) {
                case -242:
                    for fi24 := 0; fi24 < 5; fi24++ {
                        v25 := GenS3{F0: 128}
                        _ = fi24
                        _ = v25
                    }
                case 127:
                    v26 := GenS4{F0: 128}
                    _ = v26
                }
            }
        }
        v27 := v4[v5]
        var v28 int32 = int32(v5)
        _ = v27
        _ = v28
    } else {
        v2 = append(v2, 932)
    }
    for fi29 := 0; fi29 < 16; fi29++ {
        if (48 <= 100) {
            fmt.Println(v16, v6, v17)
            switch 128 {
            case 127:
                fmt.Println(v23)
            }
            if false {
                v5 -= 32767
                fmt.Println(fi29, v16)
            } else {
                var v30 bool = (true && false)
                _ = v30
            }
        } else {
            v18 = v17
            v31 := GenS5{F0: 42, F1: 42, F2: 127}
            if (42 < 932) {
                var v32 int = (v5 / ((127) | 1))
                v5 = (v5 & v32)
                _ = v32
            } else {
                switch v5 {
                case -429:
                    v5 = 128
                case -485:
                    v5 -= (fi29 - int(v23))
                default:
                    var v33 float32 = v17
                    _ = v33
                }
                var v34 bool = true
                _ = v34
            }
            _ = v31
        }
        v23 -= (^(v23))
        var v35 uint64 = 2
        _ = fi29
        _ = v35
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
