package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 string) uint64 {
    fmt.Println(p0)
    v4 := GenS1{F0: 42}
    var v5 bool = ((!true) && true)
    v5 = (3.14159 > 3.14159)
    if v5 {
        v6 := []int{9223372036854775807, -1}
        var v7 int8 = 100
        v8 := GenS1{F0: 42}
        _ = v6
        _ = v7
        _ = v8
    } else {
        for fi9 := 0; fi9 < 8; fi9++ {
            switch "\n\t" {
            case "":
                fmt.Println(p0, p0)
            }
            if v5 {
                v5 = (("" + p0) != "")
                p0 = p0
                p0 = p0
            }
            _ = fi9
        }
        var v10 int64 = 1
        v5 = v5
        _ = v10
    }
    _ = p0
    _ = v4
    _ = v5
    return 42
}
func genFn1(p0 bool, p1 int, p2 int8) bool {
    var v11 string = ("\n\t" + "\n\t")
    v12 := []int{}
    var v13 uint = 4294967295
    _ = len(v12)
    if p0 {
        fmt.Println(p1)
        var v14 uint64 = genFn0(v11)
        if (0.0 == (((((((1.5 * 3.14159) + 0.1) - (1e10 + float64(-2.25))) / (-2.25 + -2.25)) + 1e10) * 0.0) / ((float64(1.5) - 3.14159) / 1.0))) {
            var v15 uint32 = 4294967295
            v16 := GenS1{F0: 255}
            var v17 bool = p0
            _ = v15
            _ = v16
            _ = v17
        } else {
            var v18 int64 = 9223372036854775807
            _ = v18
        }
        _ = v14
    } else {
        p1 -= p1
        var v19 int = p1
        var v20 uint16 = 42
        _ = v19
        _ = v20
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v13
    return false
}

func main() {
    var v1 uint8 = 0
    v2 := make(map[uint32]int)
    var v3 uint8 = (v1 & v1)
    var v21 uint16 = 32767
    v2[0] = 9223372036854775807
    v22 := GenS1{F0: v1}
    for fi23 := 0; fi23 < 13; fi23++ {
        v21 -= (v21 | (^(v21)))
        var v24 int64 = -9223372036854775808
        var v25 uint32 = 577
        _ = fi23
        _ = v24
        _ = v25
    }
    var v26 int16 = 0
    v26 += 0
    fmt.Println(v3, v1, v1)
    v21 = 65535
    v21 = v21
    for fi27 := 0; fi27 < 17; fi27++ {
        var v28 int16 = (^(v26))
        _ = fi27
        _ = v28
    }
    fmt.Println(v21, v21, v26)
    v29 := GenS2{F0: 18446744073709551615}
    v21 = (^(v21))
    var v30 float64 = float64(3.14159)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v26
    _ = v29
    _ = v3
    _ = v30
}
