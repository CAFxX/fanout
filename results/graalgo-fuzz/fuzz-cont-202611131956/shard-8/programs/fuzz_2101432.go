package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 uint64; F2 string }
type GenS3 struct { F0 uint16; F1 int64 }
func genFn0(p0 string, p1 int32, p2 int) int {
    var v7 float32 = 3.14159
    v8 := GenS1{F0: 888}
    p1 += p1
    if (false || false) {
        p2 = 319
        p1 += p1
        fmt.Println(p0)
    } else {
        v9 := []int{p2, 42, (-(p2))}
        if true {
            for fi10 := 0; fi10 < 13; fi10++ {
                fmt.Println(p2, v7)
                _ = fi10
            }
            if ((0.1 / 1.0) > (1.5 / 1e10)) {
                v11 := GenS1{F0: 4294967295}
                if true {
                    v12 := GenS1{F0: 2}
                    var v13 int8 = 127
                    var v14 uint32 = 2147483647
                    _ = v12
                    _ = v13
                    _ = v14
                } else {
                    v15 := GenS1{F0: 1}
                    fmt.Println(p2, p2, v7)
                    p2 = 9223372036854775807
                    _ = v15
                }
                _ = v11
            } else {
                v16 := GenS2{F0: 1, F1: 42, F2: p0}
                v7 = ((v7 * v7) - (v7 * v7))
                _ = v16
            }
        } else {
            v7 = v7
            v17 := GenS2{F0: 534, F1: 3, F2: (p0 + p0)}
            v18 := GenS1{F0: 4294967295}
            _ = v17
            _ = v18
        }
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (p2 & (p2 & (p2 * p2)))
}

func main() {
    var v1 float32 = -2.25
    var v2 float64 = ((0.0 / 1.5) * 0.0)
    var v3 float64 = v2
    var v4 float32 = float32((-2.25 + (0.0 / 1.0)))
    var v5 uint16 = 120
    var v6 int64 = -1
    for fi19 := 0; fi19 < 3; fi19++ {
        v20 := GenS3{F0: (v5 % ((v5) | 1)), F1: 0}
        _ = fi19
        _ = v20
    }
    if (579 > 127) {
        var v21 int8 = 127
        v1 -= ((3.14159 - 0.1) * (-2.25 + -2.25))
        _ = v21
    } else {
        v3 += float64(((v1 + (((v4 - v4) / v4) / (float32((v2 * float64(v1))) + float32(v3)))) + ((0.0 + (((0.0 / v1) - v1) - v1)) / v4)))
        v3 += v3
    }
    var v22 uint = 65535
    v23 := GenS3{F0: 663, F1: -1}
    if ((v5 == (^(v5))) && ((v5 != (^(v5))) || (!false))) {
        v24 := GenS1{F0: v22}
        if (-247 > genFn0((("\n\t" + ("foo bar" + "hello")) + "foo bar"), 913, -9223372036854775808)) {
            v25 := []int{}
            _ = v25
        }
        _ = v24
    }
    v5 = v5
    v22 -= (v22 << ((42) & 63))
    var v26 uint16 = v5
    v1 += v1
    var v27 uint16 = (v26 / ((v26) | 1))
    v2 += v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
