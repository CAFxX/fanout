package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32; F1 uint32; F2 int32 }
type GenS4 struct { F0 uint32; F1 uint64; F2 uint64 }
func genFn0(p0 int16, p1 int16) bool {
    if (true || false) {
        fmt.Println(p0)
        v6 := GenS1{F0: 0, F1: 0}
        _ = v6
    }
    if true {
        fmt.Println(p0)
        v7 := GenS2{F0: 127}
        for fi8 := 0; fi8 < 19; fi8++ {
            if true {
                v9 := GenS1{F0: 4294967295, F1: int64(p0)}
                var v10 int32 = -501
                _ = v10
                _ = v9
            } else {
                if (false || (!(42 != 127))) {
                    var v11 float64 = 0.1
                    switch "a" {
                    case "":
                        var v12 int16 = p0
                        _ = v12
                    case "hello":
                        fmt.Println(v11, p1)
                    }
                    _ = v11
                }
                if ((1.5 + float64(1e10)) >= -2.25) {
                    var v13 int = fi8
                    var v14 float32 = ((float32(float64((((0.1 * (float32(((1e10 - (3.14159 / (-2.25 - 0.0))) * 1e10)) * (0.1 / 1.0))) - (1.5 * 1.5)) / 1.0))) - 3.14159) + -2.25)
                    _ = v13
                    _ = v14
                }
            }
            _ = fi8
        }
        _ = v7
    } else {
        v15 := []int{219, 298, -48, 4294967295}
        fmt.Println(p1, p0)
        _ = v15
    }
    if false {
        p0 = 1000
        p1 += p0
    } else {
        v16 := GenS1{F0: 311, F1: 42}
        v17 := []int{-1, -418, 2147483647, -2147483648}
        _ = v16
        _ = v17
    }
    p1 = (p0 + (-(p0)))
    _ = p0
    _ = p1
    return (p0 < p1)
}
func genFn1() uint8 {
    var v18 uint64 = 823
    switch 128 {
    case 9223372036854775807:
        for fi19 := 0; fi19 < 15; fi19++ {
            if ((true || false) && (770 >= 1)) {
                fmt.Println(fi19, fi19)
                switch "" {
                case "foo bar":
                    switch fi19 {
                    case 1:
                        v20 := GenS1{F0: 71, F1: -1}
                        _ = v20
                    case 127:
                        v21 := []int{-2147483648, 1}
                        _ = v21
                    case 143:
                        switch "hello" {
                        case "\n\t":
                            v22 := []int{fi19, 18, (fi19 | fi19)}
                            _ = v22
                        case "a":
                            v18 -= v18
                        }
                    default:
                        v23 := GenS2{F0: (-(fi19))}
                        _ = v23
                    }
                default:
                    if true {
                        var v24 int8 = 51
                        _ = v24
                    }
                }
                v25 := GenS3{F0: -2147483648, F1: 42, F2: 42}
                _ = v25
            } else {
                switch fi19 {
                case 32768:
                    var v26 uint16 = 32767
                    _ = v26
                case 32767:
                    fmt.Println(v18, fi19)
                }
            }
            v27 := []int{(fi19 - (^(fi19)))}
            _ = fi19
            _ = v27
        }
    case -693:
        v28 := []int{9223372036854775807, -9223372036854775808, -2147483648, -2147483648}
        _ = v28
    }
    v29 := GenS4{F0: 65535, F1: 4294967295, F2: v18}
    switch 32767 {
    case -2147483648:
        v18 -= (^(v18))
    case 32768:
        switch 9223372036854775807 {
        case 0:
            v18 += 18446744073709551615
        }
    case -9223372036854775808:
        fmt.Println(v18)
    }
    _ = v18
    _ = v29
    return 155
}

func main() {
    var v1 uint = 1
    var v2 uint64 = 108
    var v3 int64 = -1
    var v4 int64 = v3
    v5 := []int{}
    v30 := GenS3{F0: 453, F1: 4294967295, F2: 0}
    for fi31 := 0; fi31 < 16; fi31++ {
        if true {
            if ("foo bar" >= "a") {
                _ = len(v5)
                if ((-(fi31)) == (fi31 & fi31)) {
                    v3 -= int64(fi31)
                    v4 = (v4 / ((v3) | 1))
                    v32 := GenS3{F0: 42, F1: 0, F2: 0}
                    _ = v32
                }
                fmt.Println(v2, v4, v4)
            } else {
                v3 = 0
                v2 = (v2 << (((v1 | v1)) & 63))
                v33 := GenS2{F0: fi31}
                _ = v33
            }
        } else {
            var v34 int8 = -1
            fmt.Println(fi31, v34, v34)
            var v35 float64 = 1.5
            _ = v34
            _ = v35
        }
        _ = fi31
    }
    var v36 int64 = (v4 - (v4 + (v4 | -81)))
    var v37 uint64 = v2
    var v38 bool = false
    var v39 float32 = 0.0
    var v40 uint = v1
    fmt.Println(v39, v37)
    v37 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
