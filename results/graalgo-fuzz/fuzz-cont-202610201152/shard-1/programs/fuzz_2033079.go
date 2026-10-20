package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 bool }
type GenS2 struct { F0 int32 }
func genFn0(p0 int) string {
    p0 = p0
    var v5 bool = true
    v6 := []int{(p0 - (p0 | (-(p0))))}
    fmt.Println(v5)
    _ = p0
    _ = v5
    _ = v6
    return ""
}

func main() {
    var v1 bool = true
    var v2 uint16 = 1
    var v3 bool = v1
    var v4 int16 = 1
    v4 = 0
    v1 = ((1 > -364) || v1)
    var v7 int32 = 0
    v8 := GenS1{F0: (v2 | v2), F1: 556, F2: (!(v1 && ((v4 + -32768) < (^(v4)))))}
    v4 -= (^(v4))
    switch "hello" {
    case "":
        v9 := []int{}
        _ = v9
    }
    if v1 {
        if (0 < -55) {
            var v10 float64 = (1e10 * float64(((0.0 + (0.1 - (-2.25 / (-2.25 + (0.1 * ((1.5 * (0.0 + float32(((0.1 * 3.14159) * ((-2.25 + 3.14159) + 1e10))))) - 3.14159)))))) + float32(((1e10 / 1.0) + 0.0)))))
            v2 = v2
            if true {
                for fi11 := 0; fi11 < 16; fi11++ {
                    v12 := GenS1{F0: (v2 % (((v2 / (((v2 % (((v2 >> ((0) & 63))) | 1))) | 1))) | 1)), F1: 2, F2: (("\n\t" > "foo bar") && (4294967295 == 4294967295))}
                    var v13 uint32 = 387
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                v14 := GenS2{F0: v7}
                for fi15 := 0; fi15 < 5; fi15++ {
                    var v16 bool = v3
                    _ = fi15
                    _ = v16
                }
                _ = v14
            } else {
                v1 = ((!v1) && (!v1))
                fmt.Println(v7, v4)
            }
            _ = v10
        }
        v17 := GenS1{F0: 32767, F1: 4294967295, F2: false}
        _ = v17
    } else {
        v7 -= v7
        var v18 int16 = 32767
        _ = v18
    }
    var v19 int8 = 42
    var v20 int8 = v19
    v4 += v4
    v21 := []int{32768, 128, 128}
    for fi22 := 0; fi22 < 8; fi22++ {
        v23 := GenS2{F0: (v7 | -176)}
        _ = fi22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
