package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 int64; F2 uint16 }
type GenS3 struct { F0 int64; F1 int64; F2 int }
type GenS4 struct { F0 string; F1 uint32; F2 int8 }
func genFn0(p0 uint64, p1 int16, p2 uint64) uint64 {
    fmt.Println(p1, p1, p1)
    var v7 uint64 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (p0 << ((18446744073709551615) & 63))
}
func genFn1(p0 uint64, p1 int32, p2 bool) int {
    if (-58 != 42) {
        v8 := []int{0, 941}
        _ = v8
    }
    v9 := GenS1{F0: 147}
    v10 := GenS1{F0: 127}
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return -110
}

func main() {
    var v1 uint32 = 160
    var v2 int8 = -1
    v3 := make(map[int]int)
    var v4 int = 0
    v5 := make(map[int]bool)
    var v6 float32 = ((((3.14159 * (3.14159 - (-2.25 + 3.14159))) * (-2.25 + (0.0 / 1.5))) - ((1e10 + 1.5) + 3.14159)) + float32(3.14159))
    for fi11 := 0; fi11 < 5; fi11++ {
        var v12 uint16 = 42
        var v13 float32 = (((v6 - v6) - v6) / 1.0)
        _ = fi11
        _ = v12
        _ = v13
    }
    v14 := []int{(v4 & (v4 & v4)), 32767}
    v2 += int8(v4)
    for fi15 := 0; fi15 < 16; fi15++ {
        var v16 uint64 = 42
        if (!(1 >= 4294967295)) {
            v17 := GenS1{F0: 1}
            for fi18 := 0; fi18 < 19; fi18++ {
                var v19 int = (v4 | v4)
                _ = fi18
                _ = v19
            }
            _ = v17
        } else {
            _ = len(v5)
            v20 := GenS2{F0: v16, F1: 9223372036854775807, F2: 255}
            switch "" {
            case "\n\t":
                v2 = v2
            case "":
                if (false && false) {
                    var v21 int64 = 1
                    _ = len(v5)
                    _ = v21
                } else {
                    switch fi15 {
                    case -43:
                        fmt.Println(fi15)
                    }
                }
            case "a":
                _ = len(v14)
            default:
                for fi22 := 0; fi22 < 20; fi22++ {
                    var v23 int = v4
                    if (101 < -1) {
                        v14 = append(v14, (v23 + (v4 ^ (v4 & (fi22 % ((genFn1((v16 / ((v16) | 1)), 2147483647, true)) | 1))))))
                    } else {
                        v24 := GenS3{F0: -1, F1: -9223372036854775808, F2: genFn1(v16, -1, true)}
                        switch int(v2) {
                        case -722:
                            if false {
                                _ = len(v3)
                                var v25 float64 = 0.1
                                _ = v25
                            }
                        case 1:
                            v2 += 100
                        }
                        _ = v24
                    }
                    _ = fi22
                    _ = v23
                }
            }
            _ = v20
        }
        _ = fi15
        _ = v16
    }
    switch 32768 {
    case 1:
        fmt.Println(v2, v4)
    }
    var v26 int32 = 2147483647
    v1 = (v1 >> ((2) & 63))
    v27 := GenS4{F0: "a", F1: v1, F2: (-(v2))}
    var v28 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
