package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 uint8; F1 uint64; F2 uint }
func genFn0(p0 int8) uint8 {
    for fi4 := 0; fi4 < 12; fi4++ {
        for fi5 := 0; fi5 < 14; fi5++ {
            switch "" {
            case "foo bar":
                v6 := []int{(fi4 ^ (fi5 % (((-(fi4))) | 1))), fi4}
                _ = v6
            }
            fmt.Println(fi5, fi5, p0)
            _ = fi5
        }
        _ = fi4
    }
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 uint8 = 24
        for fi9 := 0; fi9 < 6; fi9++ {
            v10 := GenS1{F0: (!((3.14159 * float32(-2.25)) <= (0.0 + (0.1 + ((-2.25 + 0.0) / (3.14159 / (0.1 * ((((0.0 * 1.5) + 3.14159) * (0.1 / -2.25)) + ((0.0 / 1e10) - 0.1))))))))), F1: 2}
            if ((4294967295 <= 1) || ((1.5 * 0.0) < 3.14159)) {
                if false {
                    v11 := []int{(fi7 - 128), (fi9 * 32768), -362, 4294967295}
                    var v12 int = 4294967295
                    _ = v11
                    _ = v12
                }
                var v13 uint8 = 0
                v14 := GenS2{F0: (v8 << ((0) & 63)), F1: 42, F2: 901}
                _ = v13
                _ = v14
            } else {
                var v15 uint32 = 1
                v16 := []int{-426}
                _ = v15
                _ = v16
            }
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    _ = p0
    return 76
}

func main() {
    v1 := []uint32{4294967295, 2, 776}
    v2 := make(map[uint32]int)
    var v3 float64 = 0.1
    delete(v2, 4294967295)
    var v17 float64 = 3.14159
    fmt.Println(v17, v17)
    v18 := v2[0]
    var v19 int = 788
    var v20 uint16 = 2
    v21 := GenS2{F0: 128, F1: 4294967295, F2: 1}
    for fi22 := 0; fi22 < 1; fi22++ {
        var v23 int8 = 100
        switch "hello" {
        case "":
            if (0 == 42) {
                v19 += -9223372036854775808
            } else {
                var v24 int32 = -1
                _ = v24
            }
        }
        v25 := v2[1]
        _ = fi22
        _ = v23
        _ = v25
    }
    v26 := GenS1{F0: (-32768 >= 42), F1: 101}
    v18 -= 42
    for fi27 := 0; fi27 < 2; fi27++ {
        _ = len(v1)
        if len(v1) > 0 {
            v1[0] = 834
        }
        _ = fi27
    }
    var v28 uint16 = v20
    v17 += v3
    for fi29 := 0; fi29 < 18; fi29++ {
        var v30 int8 = 0
        v20 += v28
        _ = fi29
        _ = v30
    }
    var v31 float32 = 0.1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v28
    _ = v3
    _ = v31
}
