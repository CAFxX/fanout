package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int32; F1 string; F2 int32 }
func genFn0(p0 int16) int8 {
    if (0.1 > 1e10) {
        var v6 float32 = 0.1
        _ = v6
    } else {
        for fi7 := 0; fi7 < 20; fi7++ {
            if false {
                p0 = (p0 + (p0 ^ (p0 << ((18446744073709551615) & 63))))
            } else {
                var v8 uint = 903
                _ = v8
            }
            for fi9 := 0; fi9 < 9; fi9++ {
                fmt.Println(p0, p0)
                _ = fi9
            }
            _ = fi7
        }
    }
    var v10 int8 = 42
    if (false && false) {
        switch 2147483647 {
        case 2147483647:
            fmt.Println(p0)
        case 709:
            p0 += p0
        case 419:
            v11 := []int{}
            _ = v11
        }
        for fi12 := 0; fi12 < 13; fi12++ {
            if false {
                var v13 uint8 = 4
                v14 := []int{int(v13)}
                _ = v13
                _ = v14
            } else {
                if (-599 == -305) {
                    v15 := GenS1{F0: 2147483647}
                    v16 := GenS2{F0: true}
                    _ = v15
                    _ = v16
                }
            }
            v10 -= v10
            fmt.Println(fi12, p0, fi12)
            _ = fi12
        }
    }
    v17 := []int{-447}
    _ = p0
    _ = v10
    _ = v17
    return v10
}

func main() {
    var v1 int = -2147483648
    var v2 float32 = 1.5
    var v3 uint32 = 1
    var v4 float64 = ((1e10 * -2.25) * float64((float32((-2.25 - 3.14159)) - 3.14159)))
    var v5 uint = 1
    v18 := []int{(-(v1))}
    fmt.Println(v3)
    var v19 int16 = 0
    v20 := GenS3{F0: int32(v3)}
    for fi21 := 0; fi21 < 1; fi21++ {
        var v22 uint32 = 4294967295
        _ = fi21
        _ = v22
    }
    v23 := GenS4{F0: 0, F1: "", F2: 42}
    for fi24 := 0; fi24 < 2; fi24++ {
        v1 += 128
        v25 := GenS4{F0: 1, F1: ("\n\t" + "\n\t"), F2: -1}
        _ = fi24
        _ = v25
    }
    v4 += v4
    var v26 uint = v5
    if true {
        v4 = (v4 + (((3.14159 * 1.5) * (1e10 / v4)) - 1.5))
    }
    for fi27 := 0; fi27 < 4; fi27++ {
        if len(v18) > 0 {
            v18[0] = (v1 + -685)
        }
        _ = fi27
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
