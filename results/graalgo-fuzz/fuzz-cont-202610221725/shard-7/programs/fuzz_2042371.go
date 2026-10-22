package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
func genFn0() int32 {
    for fi7 := 0; fi7 < 13; fi7++ {
        v8 := []int{(fi7 & (fi7 & 686)), fi7, fi7}
        for fi9 := 0; fi9 < 1; fi9++ {
            if len(v8) > 0 {
                v8[0] = fi9
            }
            if (930 < 65535) {
                _ = len(v8)
            }
            var v10 uint32 = 42
            var v11 int = (fi7 ^ (fi9 ^ -919))
            _ = fi9
            _ = v10
            _ = v11
        }
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = fi7
        _ = v8
    }
    v12 := []int{-641, 127, 0}
    var v13 float32 = 1.5
    _ = len(v12)
    _ = v12
    _ = v13
    return 307
}
func genFn1(p0 string) int32 {
    fmt.Println(p0)
    p0 = (p0 + "foo bar")
    _ = p0
    return -945
}

func main() {
    v1 := []string{}
    v2 := make(map[uint32]int)
    var v3 float64 = 0.0
    v4 := []string{"hello", ("a" + (("a" + "\n\t") + "a")), ""}
    v5 := []uint32{}
    var v6 float64 = (1.5 - (v3 / ((v3 - 3.14159) + ((v3 * v3) * (1.5 * (v3 / (v3 / (v3 / (((-2.25 * (v3 + ((v3 - float64(3.14159)) / (v3 * float64(0.0))))) + (v3 * v3)) - v3)))))))))
    v6 -= v3
    v14 := GenS1{F0: 42, F1: 511}
    if false {
        v15 := GenS1{F0: 2, F1: 489}
        v3 -= float64(-2.25)
        for fi16 := 0; fi16 < 17; fi16++ {
            var v17 int = fi16
            v18 := GenS1{F0: 4294967295, F1: 65535}
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = v15
    } else {
        fmt.Println(v6)
        v6 += v6
    }
    v6 -= v6
    var v19 int64 = 693
    fmt.Println(v19)
    fmt.Println(v19, v19)
    v20 := GenS1{F0: 2, F1: 18446744073709551615}
    v21 := GenS1{F0: 937, F1: 716}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
