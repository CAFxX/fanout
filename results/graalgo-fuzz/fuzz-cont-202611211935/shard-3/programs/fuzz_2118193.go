package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 int16; F1 int8 }
func genFn0(p0 int32, p1 uint16) uint8 {
    switch 1 {
    case -341:
        switch 99 {
        case 128:
            v7 := GenS1{F0: 2147483648, F1: 66}
            _ = v7
        case 399:
            var v8 int32 = (p0 % ((p0) | 1))
            _ = v8
        default:
            p0 += (p0 & -615)
        }
    case 2147483647:
        p0 += (p0 ^ (p0 | 2147483647))
    case 0:
        for fi9 := 0; fi9 < 2; fi9++ {
            switch "\n\t" {
            case "\n\t":
                v10 := []int{fi9, 755, fi9}
                _ = v10
            default:
                for fi11 := 0; fi11 < 4; fi11++ {
                    var v12 int16 = 1000
                    if true {
                        v13 := GenS1{F0: 469, F1: 56}
                        var v14 string = "foo bar"
                        var v15 string = v14
                        _ = v13
                        _ = v14
                        _ = v15
                    } else {
                        fmt.Println(v12)
                        v16 := GenS2{F0: -32768, F1: 1}
                        v17 := []int{(-(fi9)), (fi9 * (fi11 & fi9)), fi9, (fi11 | (fi11 % (((fi9 << ((1) & 63))) | 1)))}
                        fmt.Println(v12, fi9, fi9)
                        fmt.Println(fi9)
                        v12 -= (v12 + int16(p1))
                        v18 := GenS2{F0: v12, F1: -83}
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                    p0 -= 0
                    p1 += p1
                    _ = fi11
                    _ = v12
                }
            }
            for fi19 := 0; fi19 < 18; fi19++ {
                var v20 bool = true
                if v20 {
                    v20 = (!v20)
                }
                var v21 string = "foo bar"
                var v22 float64 = ((-2.25 + 1.5) + float64(3.14159))
                _ = fi19
                _ = v20
                _ = v21
                _ = v22
            }
            _ = fi9
        }
    }
    var v23 float64 = float64(3.14159)
    v24 := []int{127, 32768}
    _ = p0
    _ = p1
    _ = v23
    _ = v24
    return 128
}

func main() {
    v1 := make(map[string]bool)
    var v2 string = "\n\t"
    var v3 float32 = -2.25
    v4 := make(map[string]bool)
    var v5 int = 128
    var v6 float32 = (3.14159 + v3)
    var v25 int16 = 0
    v3 -= 0.0
    v6 -= (((((1.5 * v3) * 1.5) * v6) / float32(3.14159)) + (((v3 + (((((v3 / v6) / v6) + (v6 / (((((v3 + (v3 - 1e10)) - 1.5) / v6) + v6) * v6))) + 1.5) / ((float32(((3.14159 + (((-2.25 + -2.25) + (((((-2.25 - 1e10) / -2.25) * 0.0) * 3.14159) - 0.0)) / 3.14159)) + -2.25)) * v3) + v3))) + -2.25) + v3))
    v26 := []int{v5, (v5 % ((v5) | 1)), 2147483647}
    for fi27 := 0; fi27 < 20; fi27++ {
        v25 = -32768
        _ = fi27
    }
    for fi28 := 0; fi28 < 18; fi28++ {
        v26 = append(v26, (v5 - v5))
        v6 = v3
        v6 = (v3 - float32(0.1))
        _ = fi28
    }
    v5 -= (v5 / ((42) | 1))
    var v29 bool = true
    if len(v26) >= 2 {
        v26 = v26[1:len(v26)]
    }
    var v30 int8 = 100
    var v31 string = (("" + v2) + v2)
    for fi32 := 0; fi32 < 15; fi32++ {
        var v33 int32 = 2147483647
        if len(v26) > 0 {
            v26[0] = fi32
        }
        v34 := GenS2{F0: (v25 | (v25 * (-(v25)))), F1: 1}
        _ = fi32
        _ = v33
        _ = v34
    }
    v35 := GenS1{F0: 0, F1: 0}
    v36 := GenS1{F0: 2147483647, F1: 2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
