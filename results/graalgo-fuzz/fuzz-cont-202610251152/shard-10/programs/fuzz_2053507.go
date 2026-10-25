package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 uint8; F2 uint }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int64; F1 string }
type GenS5 struct { F0 bool }
func genFn0(p0 int) uint {
    switch (p0 ^ 1) {
    case 128:
        v5 := GenS1{F0: p0}
        _ = v5
    case 127:
        p0 += (p0 & p0)
    default:
        p0 = p0
    }
    for fi6 := 0; fi6 < 19; fi6++ {
        fmt.Println(p0, fi6)
        v7 := GenS1{F0: fi6}
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 13; fi8++ {
        for fi9 := 0; fi9 < 19; fi9++ {
            p0 += (fi8 & -578)
            var v10 uint16 = 713
            var v11 int32 = 0
            if (2147483648 <= 289) {
                if (fi8 >= p0) {
                    var v12 int16 = 1000
                    if true {
                        var v13 int64 = 0
                        _ = v13
                    } else {
                        v14 := GenS1{F0: (fi9 % ((p0) | 1))}
                        var v15 int64 = -9223372036854775808
                        v16 := []int{(-(p0)), fi9, (^(p0)), fi8}
                        _ = v14
                        _ = v15
                        _ = v16
                    }
                    _ = v12
                } else {
                    var v17 float32 = float32((0.1 + (1.5 * 0.1)))
                    _ = v17
                }
                p0 = 0
            }
            _ = fi9
            _ = v10
            _ = v11
        }
        if false {
            for fi18 := 0; fi18 < 8; fi18++ {
                fmt.Println(fi18, p0)
                var v19 uint32 = 42
                if false {
                    switch "a" {
                    case "\n\t":
                        var v20 float32 = 1.5
                        _ = v20
                    case "hello":
                        if (127 == 65) {
                            var v21 int64 = 9223372036854775807
                            _ = v21
                        } else {
                            v22 := GenS1{F0: (p0 >> ((255) & 63))}
                            p0 -= fi18
                            _ = v22
                        }
                    }
                } else {
                    p0 += fi18
                    var v23 uint32 = 1
                    var v24 int8 = 127
                    _ = v23
                    _ = v24
                }
                _ = fi18
                _ = v19
            }
        }
        _ = fi8
    }
    _ = p0
    return 1
}
func genFn1(p0 int16, p1 uint, p2 uint32) int64 {
    if (!(0 != 18446744073709551615)) {
        v25 := []int{}
        _ = v25
    } else {
        p2 -= p2
        v26 := GenS1{F0: -943}
        _ = v26
    }
    if (32767 <= uint16(p2)) {
        if true {
            fmt.Println(p2, p2, p1)
        } else {
            switch "" {
            case "a":
                var v27 uint32 = 65535
                _ = v27
            }
            p0 = p0
            var v28 float32 = 1e10
            _ = v28
        }
    } else {
        v29 := []int{}
        p0 += p0
        if false {
            p0 = p0
            fmt.Println(p1, p0, p0)
            v30 := GenS2{F0: 0, F1: 95, F2: genFn0(127)}
            _ = v30
        }
        _ = v29
    }
    p0 = (p0 - 32767)
    _ = p0
    _ = p1
    _ = p2
    return 876
}

func main() {
    v1 := []uint32{2147483647}
    var v2 float64 = ((float64(1e10) + 0.0) - 1.5)
    var v3 string = "hello"
    var v4 int = 127
    var v31 float32 = 0.1
    v32 := GenS2{F0: 11, F1: 0, F2: 18446744073709551615}
    for fi33 := 0; fi33 < 3; fi33++ {
        fmt.Println(v2, fi33, v4)
        _ = fi33
    }
    v31 -= -2.25
    v34 := GenS3{F0: 2}
    if false {
        var v35 uint = 42
        _ = v35
    }
    v31 += float32(float64(0.1))
    v31 -= 0.1
    var v36 int16 = 1000
    switch -9223372036854775808 {
    case 86:
        v37 := GenS4{F0: 0, F1: v3}
        _ = v37
    case 9223372036854775807:
        fmt.Println(v2, v36, v31)
    case 128:
        var v38 int = v4
        _ = v38
    default:
        v39 := GenS1{F0: v4}
        _ = v39
    }
    if true {
        v1 = append(v1, 42)
    } else {
        fmt.Println(v4, v3)
        if false {
            v4 = (v4 >> ((461) & 63))
        }
        v1 = append(v1, 2147483648)
    }
    switch (v3 + v3) {
    case "hello":
        var v40 int64 = 9223372036854775807
        _ = v40
    case "\n\t":
        v41 := GenS5{F0: (625 <= (v36 >> ((genFn0((v4 & (v4 >> ((255) & 63))))) & 63)))}
        _ = v41
    default:
        v36 = v36
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v34
    _ = v36
    _ = v4
}
