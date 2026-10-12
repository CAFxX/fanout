package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int }
func genFn0(p0 int64, p1 uint64, p2 int16) uint64 {
    for fi6 := 0; fi6 < 4; fi6++ {
        p0 -= (-(p0))
        v7 := []int{fi6, fi6}
        p0 += (p0 - 1)
        _ = fi6
        _ = v7
    }
    switch ((((("\n\t" + "hello") + "a") + "hello") + "a") + ("\n\t" + ("" + ("hello" + ("hello" + ""))))) {
    case "\n\t":
        v8 := GenS1{F0: 978, F1: false}
        _ = v8
    case "a":
        for fi9 := 0; fi9 < 14; fi9++ {
            p1 += (p1 / (((p1 % ((p1) | 1))) | 1))
            _ = fi9
        }
    case "":
        var v10 uint16 = 1
        _ = v10
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        var v12 uint = 600
        _ = fi11
        _ = v12
    }
    v13 := GenS2{F0: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return (p1 ^ p1)
}

func main() {
    v1 := make(map[int]bool)
    var v2 float32 = (1.5 / 1.0)
    var v3 float64 = 3.14159
    var v4 int = 32768
    v5 := make(map[uint32]bool)
    v14 := GenS2{F0: 18446744073709551615}
    v15 := GenS2{F0: genFn0(-9223372036854775808, 714, 32767)}
    for fi16 := 0; fi16 < 8; fi16++ {
        v17 := GenS3{F0: (v4 / ((-2147483648) | 1))}
        _ = fi16
        _ = v17
    }
    if false {
        for fi18 := 0; fi18 < 12; fi18++ {
            var v19 float64 = 0.0
            _ = fi18
            _ = v19
        }
    }
    fmt.Println(v2, v2)
    if (32767 <= -182) {
        v4 = (v4 ^ (v4 & v4))
    } else {
        v20 := v5[1]
        var v21 uint16 = 255
        _ = v20
        _ = v21
    }
    fmt.Println(v2, v4)
    fmt.Println(v2, v2)
    v22 := []int{(v4 - (v4 << ((953) & 63))), (-(v4)), v4, (v4 * (v4 ^ 32767))}
    var v23 bool = (false || (!true))
    v5[0] = v23
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
