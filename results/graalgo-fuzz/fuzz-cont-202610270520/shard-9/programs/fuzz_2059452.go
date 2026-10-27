package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
func genFn0(p0 uint64, p1 string, p2 int64) uint32 {
    fmt.Println(p1)
    if (true && (p2 <= (p2 >> ((42) & 63)))) {
        v5 := []int{127, 453, 127, -9223372036854775808}
        for fi6 := 0; fi6 < 16; fi6++ {
            v7 := GenS1{F0: 127, F1: 0}
            _ = fi6
            _ = v7
        }
        _ = v5
    } else {
        for fi8 := 0; fi8 < 6; fi8++ {
            p2 = 9223372036854775807
            p2 -= 9223372036854775807
            v9 := GenS1{F0: -1, F1: 42}
            _ = fi8
            _ = v9
        }
        p1 = p1
        var v10 float64 = 0.1
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}
func genFn1(p0 int16) uint64 {
    p0 = 595
    for fi11 := 0; fi11 < 18; fi11++ {
        v12 := []int{(fi11 * fi11), fi11, fi11, fi11}
        var v13 bool = false
        var v14 uint16 = 27
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    p0 -= p0
    _ = p0
    return 2
}

func main() {
    var v1 bool = true
    var v2 uint8 = 42
    var v3 uint64 = 190
    var v4 float32 = (0.0 - 1e10)
    for fi15 := 0; fi15 < 14; fi15++ {
        var v16 int16 = 865
        v17 := []int{125}
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 float64 = (0.0 * float64(3.14159))
    v4 += 1e10
    switch "\n\t" {
    case "a":
        for fi19 := 0; fi19 < 14; fi19++ {
            if v1 {
                for fi20 := 0; fi20 < 7; fi20++ {
                    v18 += 0.1
                    var v21 bool = (42 > 4294967295)
                    v3 = (v3 | v3)
                    _ = fi20
                    _ = v21
                }
                for fi22 := 0; fi22 < 18; fi22++ {
                    var v23 int16 = -829
                    v3 += 1
                    _ = fi22
                    _ = v23
                }
            }
            v24 := []int{fi19, (fi19 - 2147483647), 9223372036854775807, fi19}
            _ = fi19
            _ = v24
        }
    case "hello":
        var v25 int16 = 1000
        _ = v25
    case "":
        v1 = v1
    }
    var v26 int16 = -1
    var v27 uint = 65535
    v3 = (v3 >> ((v27) & 63))
    v28 := []int{9223372036854775807, 32768}
    v28 = append(v28, 59)
    var v29 float64 = v18
    var v30 int32 = 3
    v30 = (-(v30))
    var v31 float32 = float32(((v18 * v18) + float64((v4 * 0.0))))
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
