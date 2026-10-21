package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 uint }
func genFn0() bool {
    var v5 int = 762
    for fi6 := 0; fi6 < 11; fi6++ {
        fmt.Println(v5, fi6)
        _ = fi6
    }
    if false {
        v5 = v5
    } else {
        for fi7 := 0; fi7 < 6; fi7++ {
            v5 = (v5 + (fi7 << ((895) & 63)))
            var v8 uint32 = 2147483647
            _ = fi7
            _ = v8
        }
    }
    _ = v5
    return false
}
func genFn1(p0 int16, p1 uint16) int {
    v9 := []int{}
    p1 -= (p1 ^ p1)
    _ = p0
    _ = p1
    _ = v9
    return 32767
}

func main() {
    v1 := make(map[int]int)
    v2 := []int{-231}
    var v3 uint = 42
    var v4 uint8 = 127
    v4 -= (v4 & 153)
    for fi10 := 0; fi10 < 11; fi10++ {
        if (fi10 > -9223372036854775808) {
            var v11 uint = 735
            for fi12 := 0; fi12 < 16; fi12++ {
                fmt.Println(fi10, v3)
                _ = fi12
            }
            var v13 int = 0
            _ = v11
            _ = v13
        }
        v3 -= v3
        _ = fi10
    }
    v3 = 1
    v3 -= (v3 | v3)
    for fi14 := 0; fi14 < 17; fi14++ {
        delete(v1, (^(fi14)))
        var v15 bool = (!true)
        _ = fi14
        _ = v15
    }
    v16 := GenS1{F0: 1, F1: "hello", F2: (^(v3))}
    if (false || ((float32(0.1) * 1e10) != 1e10)) {
        if true {
            for fi17 := 0; fi17 < 15; fi17++ {
                if true {
                    v3 = 544
                    v18 := GenS1{F0: (fi17 | (-(fi17))), F1: "", F2: (v3 >> ((v3) & 63))}
                    var v19 float64 = -2.25
                    _ = v18
                    _ = v19
                } else {
                    var v20 uint64 = 1
                    _ = v20
                }
                _ = fi17
            }
            if false {
                v3 = v3
            } else {
                fmt.Println(v3, v4)
            }
        }
        v21 := GenS1{F0: 949, F1: "hello", F2: (v3 + (v3 ^ v3))}
        _ = v21
    }
    fmt.Println(v3, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
