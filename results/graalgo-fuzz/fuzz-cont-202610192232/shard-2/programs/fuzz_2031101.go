package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 int8 }
func genFn0(p0 string, p1 uint32, p2 uint32) uint {
    if (0 > 132) {
        switch (p0 + (p0 + ("\n\t" + "\n\t"))) {
        case "\n\t":
            for fi5 := 0; fi5 < 10; fi5++ {
                for fi6 := 0; fi6 < 1; fi6++ {
                    v7 := GenS1{F0: 44, F1: 42, F2: 1}
                    if (1 <= 551) {
                        var v8 uint = 65535
                        p2 = (p2 - (p1 | (p2 / ((p2) | 1))))
                        _ = v8
                    }
                    fmt.Println(fi5, p2)
                    var v9 uint16 = 0
                    var v10 int16 = 1000
                    var v11 string = (p0 + ("\n\t" + p0))
                    _ = fi6
                    _ = v10
                    _ = v11
                    _ = v7
                    _ = v9
                }
                _ = fi5
            }
        case "a":
            v12 := []int{128, 9223372036854775807, 127, 57}
            _ = v12
        case "":
            for fi13 := 0; fi13 < 19; fi13++ {
                v14 := []int{9223372036854775807}
                _ = fi13
                _ = v14
            }
        default:
            v15 := []int{1, -2147483648, -768}
            _ = v15
        }
    } else {
        var v16 bool = (0.0 <= (((1.5 - float32(1.5)) + (0.0 + 3.14159)) + 1e10))
        _ = v16
    }
    p1 += p1
    var v17 int64 = 9223372036854775807
    v18 := GenS2{F0: 414, F1: true}
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v18
    return uint(p1)
}

func main() {
    v1 := []int{-898, 128, -784, -1}
    var v2 uint32 = 0
    v3 := []uint32{0}
    v4 := []int{4294967295, -2147483648, -653, -747}
    v19 := GenS1{F0: 18446744073709551615, F1: 181, F2: 0}
    fmt.Println(v2)
    switch 541 {
    case -841:
        if ((-2.25 <= (float64((((3.14159 / (((3.14159 + 1e10) / -2.25) / 1e10)) - 1.5) + 3.14159)) + 3.14159)) && false) {
            var v20 int16 = 238
            if true {
                v4 = append(v4, 1)
                if (-128 < 100) {
                    var v21 int32 = -661
                    if len(v1) > 0 {
                        v1[0] = 1
                    }
                    v21 = -1
                    _ = v21
                }
                v2 += (v2 << ((65535) & 63))
            }
            _ = v20
        } else {
            v2 += (v2 ^ v2)
            v3 = append(v3, 0)
            for fi22 := 0; fi22 < 6; fi22++ {
                v2 = (v2 % ((2147483648) | 1))
                v2 = v2
                v23 := GenS1{F0: 0, F1: 1, F2: 0}
                _ = fi22
                _ = v23
            }
        }
    case 223:
        v24 := GenS2{F0: 0, F1: ((v2 & 0) == (v2 ^ 4294967295))}
        _ = v24
    }
    var v25 int16 = 0
    var v26 uint64 = 42
    if false {
        v2 += uint32(v26)
    } else {
        var v27 float32 = -2.25
        if len(v3) > 0 {
            v3[0] = 841
        }
        for fi28 := 0; fi28 < 3; fi28++ {
            for fi29 := 0; fi29 < 16; fi29++ {
                var v30 string = ("hello" + ("\n\t" + "\n\t"))
                v30 = v30
                _ = fi29
                _ = v30
            }
            v25 = v25
            var v31 float64 = (0.1 / 1e10)
            _ = fi28
            _ = v31
        }
        _ = v27
    }
    var v32 float32 = 0.0
    var v33 uint64 = (v26 + v26)
    for fi34 := 0; fi34 < 10; fi34++ {
        fmt.Println(v32)
        if (!(0 > 848)) {
            for fi35 := 0; fi35 < 17; fi35++ {
                if true {
                    v26 += v26
                    v32 = ((float32(0.1) + v32) + 3.14159)
                }
                v36 := GenS2{F0: 42, F1: true}
                _ = fi35
                _ = v36
            }
        }
        _ = fi34
    }
    v37 := GenS3{F0: 90}
    fmt.Println(v2)
    v38 := GenS1{F0: 2, F1: 42, F2: 255}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v32
    _ = v33
    _ = v37
    _ = v38
    _ = v4
}
