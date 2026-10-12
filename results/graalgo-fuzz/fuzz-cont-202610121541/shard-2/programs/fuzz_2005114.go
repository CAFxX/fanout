package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 int64 }
func genFn0(p0 int) uint16 {
    for fi7 := 0; fi7 < 18; fi7++ {
        v8 := GenS1{F0: 603}
        _ = fi7
        _ = v8
    }
    var v9 float64 = float64(1.5)
    v10 := GenS1{F0: 42}
    v11 := []int{}
    var v12 float64 = v9
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return 0
}
func genFn1() uint32 {
    v13 := GenS1{F0: 4294967295}
    for fi14 := 0; fi14 < 18; fi14++ {
        var v15 float64 = 0.0
        var v16 uint32 = 4294967295
        fmt.Println(fi14)
        _ = fi14
        _ = v15
        _ = v16
    }
    _ = v13
    return 2147483647
}

func main() {
    var v1 int8 = -1
    v2 := []string{("hello" + "a"), (("\n\t" + "a") + "\n\t"), "\n\t", (("\n\t" + "hello") + "hello")}
    v3 := make(map[int]int)
    var v4 int64 = -316
    v5 := make(map[uint32]bool)
    var v6 int = 128
    var v17 int64 = (-(v4))
    fmt.Println(v4, v1)
    v3[int(v17)] = v6
    v18 := GenS2{F0: -2147483648, F1: -664}
    if false {
        if (1 > 0) {
            delete(v5, 2147483648)
        }
        if (true && true) {
            if false {
                v19 := GenS1{F0: 875}
                _ = len(v2)
                v4 = (v4 * (v17 % (((v4 % ((v17) | 1))) | 1)))
                _ = v19
            }
            v20 := GenS1{F0: 1}
            _ = v20
        } else {
            for fi21 := 0; fi21 < 1; fi21++ {
                if len(v2) > 0 {
                    v2[0] = "a"
                }
                _ = fi21
            }
        }
    }
    fmt.Println(v6)
    for fi22 := 0; fi22 < 10; fi22++ {
        for fi23 := 0; fi23 < 9; fi23++ {
            var v24 int = (fi22 >> ((42) & 63))
            var v25 int = fi22
            v3[(v6 + (v24 << ((743) & 63)))] = -764
            _ = fi23
            _ = v24
            _ = v25
        }
        var v26 uint16 = genFn0(v6)
        _ = fi22
        _ = v26
    }
    var v27 float32 = 1e10
    v28 := GenS2{F0: 1, F1: v4}
    v17 = (v17 * (v17 % (((^(v4))) | 1)))
    var v29 int16 = -810
    v27 -= 3.14159
    v6 = 32767
    for fi30 := 0; fi30 < 18; fi30++ {
        switch fi30 {
        case -1:
            var v31 int8 = (v1 * (v1 % (((v1 << ((65535) & 63))) | 1)))
            _ = v31
        case 127:
            var v32 int16 = 843
            _ = v32
        case 42:
            v1 += (-(v1))
        default:
            fmt.Println(v17, v29, v29)
        }
        _ = fi30
    }
    for fi33 := 0; fi33 < 2; fi33++ {
        v4 -= (v17 ^ v4)
        _ = fi33
    }
    var v34 int = v6
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
