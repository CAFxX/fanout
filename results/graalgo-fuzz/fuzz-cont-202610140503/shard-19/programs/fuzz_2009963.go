package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
func genFn0(p0 uint8, p1 int32) uint16 {
    p1 = -2147483648
    for fi6 := 0; fi6 < 13; fi6++ {
        fmt.Println(fi6, fi6, fi6)
        v7 := []int{(^(fi6)), 603, 0}
        if len(v7) > 0 {
            v7[0] = (fi6 - fi6)
        }
        _ = fi6
        _ = v7
    }
    var v8 uint32 = 42
    _ = p0
    _ = p1
    _ = v8
    return 1
}
func genFn1(p0 string, p1 int16, p2 int8) uint32 {
    var v9 string = ((p0 + p0) + p0)
    for fi10 := 0; fi10 < 13; fi10++ {
        for fi11 := 0; fi11 < 7; fi11++ {
            p0 = v9
            _ = fi11
        }
        var v12 bool = false
        switch fi10 {
        case 127:
            p1 += p1
        default:
            if v12 {
                v13 := []int{fi10, (fi10 % (((fi10 | (-(fi10)))) | 1)), 4294967295}
                fmt.Println(fi10, p1, p0)
                p1 += (p1 << ((2) & 63))
                _ = v13
            } else {
                v12 = false
                var v14 string = (p0 + "hello")
                _ = v14
            }
        }
        _ = fi10
        _ = v12
    }
    fmt.Println(v9, p2, v9)
    switch p0 {
    case "foo bar":
        var v15 uint16 = genFn0(255, -2147483648)
        _ = v15
    }
    v16 := []int{32768, 32768, -932}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v9
    return 940
}

func main() {
    var v1 uint8 = 103
    var v2 float32 = float32(1.5)
    var v3 int = 32767
    v4 := make(map[int]bool)
    var v5 int64 = -751
    v3 = v3
    var v17 bool = true
    for fi18 := 0; fi18 < 11; fi18++ {
        v19 := []int{v3}
        v3 += (fi18 - fi18)
        _ = fi18
        _ = v19
    }
    var v20 float64 = ((0.0 * 1e10) + -2.25)
    v4[(v3 + 32767)] = true
    var v21 float64 = 0.1
    var v22 int = v3
    if (false && v17) {
        v4[v3] = (!false)
        v23 := GenS1{F0: "hello", F1: 0}
        _ = v23
    }
    v5 = (^(v5))
    var v24 int16 = -1
    var v25 uint64 = 617
    v5 = v5
    v21 = 0.0
    v2 += (v2 / 0.1)
    v26 := []int{v3, (v22 / ((4294967295) | 1)), 32767, (v3 / ((v22) | 1))}
    _ = len(v26)
    v27 := GenS1{F0: ("foo bar" + "\n\t"), F1: 298}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
