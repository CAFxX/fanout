package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 int; F1 int32 }
type GenS5 struct { F0 int16 }
func genFn0(p0 uint32, p1 bool, p2 int64) int {
    if (4294967295 != 33) {
        if (105 <= p2) {
            v4 := GenS1{F0: 4294967295, F1: 255}
            _ = v4
        } else {
            if (-128 == -82) {
                switch 32767 {
                case -785:
                    if false {
                        p0 -= (p0 % ((p0) | 1))
                        v5 := GenS2{F0: 239, F1: 728}
                        _ = v5
                    } else {
                        p1 = false
                        for fi6 := 0; fi6 < 3; fi6++ {
                            v7 := []int{4294967295, int(p2)}
                            p0 -= p0
                            v8 := GenS1{F0: (fi6 - fi6), F1: 32767}
                            _ = fi6
                            _ = v7
                            _ = v8
                        }
                        v9 := []int{2147483647, 32767, 9223372036854775807}
                        _ = v9
                    }
                case 9223372036854775807:
                    v10 := GenS1{F0: -960, F1: 32767}
                    _ = v10
                }
            } else {
                var v11 int = -9223372036854775808
                _ = v11
            }
            if p1 {
                v12 := GenS3{F0: 42, F1: 1}
                _ = v12
            } else {
                p1 = ((("" + "") >= "hello") || p1)
            }
        }
        v13 := GenS1{F0: -9223372036854775808, F1: 2}
        var v14 int32 = -46
        _ = v13
        _ = v14
    } else {
        p2 -= -94
    }
    v15 := GenS2{F0: 0, F1: 0}
    v16 := GenS1{F0: 1, F1: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    return -9223372036854775808
}

func main() {
    var v1 float32 = (0.1 * 3.14159)
    var v2 int64 = -1
    var v3 int = 48
    switch 32768 {
    case -2147483648:
        if false {
            v2 += 899
        } else {
            v17 := GenS4{F0: (v3 % ((v3) | 1)), F1: 1}
            var v18 uint16 = 255
            _ = v17
            _ = v18
        }
    }
    v3 += v3
    fmt.Println(v1, v2)
    v2 += v2
    var v19 int64 = (v2 / ((v2) | 1))
    var v20 string = "a"
    v20 = (v20 + v20)
    v21 := []int{genFn0(290, (((^(v19)) >= (^(v2))) || (true && (true || true))), (-(v2))), 115, genFn0(1, false, v19), (v3 * v3)}
    v1 -= float32(-2.25)
    for fi22 := 0; fi22 < 7; fi22++ {
        var v23 uint = 2
        _ = fi22
        _ = v23
    }
    switch genFn0(356, (true || false), -9223372036854775808) {
    case -9223372036854775808:
        var v24 uint64 = 868
        _ = v24
    }
    var v25 uint8 = 128
    var v26 float64 = 1.5
    v27 := GenS5{F0: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v26
    _ = v27
    _ = v3
}
