package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool; F1 uint32; F2 bool }
type GenS5 struct { F0 uint32; F1 int32; F2 string }
func genFn0(p0 int64, p1 int, p2 uint32) uint16 {
    switch p1 {
    case -963:
        fmt.Println(p2, p2, p0)
    case -169:
        var v6 bool = true
        _ = v6
    case 401:
        switch (p1 << ((1) & 63)) {
        case -373:
            p1 = (p1 & 512)
        default:
            v7 := []int{}
            _ = v7
        }
    }
    switch "a" {
    case "a":
        if true {
            p1 -= p1
            for fi8 := 0; fi8 < 3; fi8++ {
                v9 := []int{fi8, 2147483647}
                _ = fi8
                _ = v9
            }
        } else {
            var v10 bool = false
            v11 := GenS1{F0: 46, F1: 1}
            _ = v10
            _ = v11
        }
    default:
        p2 = 2147483647
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}
func genFn1(p0 uint16, p1 int8, p2 uint32) int8 {
    fmt.Println(p0, p2, p2)
    switch 42 {
    case 32768:
        v12 := GenS1{F0: 2, F1: 480}
        _ = v12
    case 735:
        var v13 float64 = (1e10 * (float64(0.0) - -2.25))
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    return p1
}

func main() {
    var v1 bool = (1 <= 9223372036854775807)
    var v2 float64 = -2.25
    var v3 uint64 = 4294967295
    v4 := make(map[string]int)
    var v5 int8 = 44
    v2 += (v2 - -2.25)
    for fi14 := 0; fi14 < 5; fi14++ {
        v3 = 1
        fmt.Println(fi14, v2, fi14)
        _ = fi14
    }
    v2 = v2
    v3 = v3
    for fi15 := 0; fi15 < 17; fi15++ {
        v3 = (v3 & v3)
        var v16 int = (fi15 ^ 9223372036854775807)
        _ = fi15
        _ = v16
    }
    var v17 int = 9223372036854775807
    v18 := GenS2{F0: 0, F1: (v5 & v5)}
    _ = len(v4)
    fmt.Println(v3, v17)
    for fi19 := 0; fi19 < 13; fi19++ {
        var v20 int = v17
        _ = fi19
        _ = v20
    }
    v3 -= 2
    v21 := []int{(v17 ^ 915), v17, (v17 | int(v3))}
    v22 := GenS3{F0: "foo bar"}
    var v23 bool = v1
    if v1 {
        var v24 int = 32767
        if v1 {
            for fi25 := 0; fi25 < 10; fi25++ {
                v21 = append(v21, (fi25 - v17))
                var v26 int64 = -4
                _ = fi25
                _ = v26
            }
            v4["\n\t"] = (v24 << ((uint(v3)) & 63))
        }
        _ = v24
    } else {
        for fi27 := 0; fi27 < 17; fi27++ {
            v3 = (^(v3))
            _ = fi27
        }
        fmt.Println(v1, v5)
        v21 = append(v21, v17)
    }
    switch 671 {
    case 594:
        var v28 int32 = -289
        _ = v28
    case 42:
        v29 := GenS4{F0: true, F1: 0, F2: (3.14159 < (((1.5 * 1e10) / 1e10) + 0.1))}
        _ = v29
    default:
        v30 := GenS5{F0: 937, F1: 0, F2: "foo bar"}
        _ = v30
    }
    for fi31 := 0; fi31 < 19; fi31++ {
        v32 := GenS5{F0: 0, F1: -1, F2: "hello"}
        _ = fi31
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
