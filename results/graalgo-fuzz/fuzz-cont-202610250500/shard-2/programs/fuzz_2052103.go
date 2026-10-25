package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 int }
func genFn0() uint16 {
    for fi5 := 0; fi5 < 4; fi5++ {
        for fi6 := 0; fi6 < 14; fi6++ {
            var v7 float64 = float64(((3.14159 / 3.14159) / float32(-2.25)))
            v7 -= v7
            _ = fi6
            _ = v7
        }
        for fi8 := 0; fi8 < 9; fi8++ {
            v9 := []int{fi5, (fi5 + (fi8 | 128)), fi5, (-(fi8))}
            var v10 float32 = (0.0 / 1e10)
            fmt.Println(fi5, fi5)
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi5
    }
    v11 := GenS1{F0: -1, F1: 42, F2: -234}
    v12 := GenS1{F0: -1, F1: 112, F2: 127}
    _ = v11
    _ = v12
    return 714
}
func genFn1() bool {
    for fi13 := 0; fi13 < 6; fi13++ {
        switch (fi13 ^ -156) {
        case 127:
            if true {
                v14 := []int{fi13}
                var v15 int32 = 1
                _ = v14
                _ = v15
            }
        default:
            var v16 float32 = 0.1
            _ = v16
        }
        _ = fi13
    }
    var v17 bool = true
    _ = v17
    return v17
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[string]uint32)
    var v3 float64 = float64(float32(1e10))
    v4 := []string{("\n\t" + "\n\t"), "hello", ("\n\t" + "\n\t")}
    v18 := GenS1{F0: -362, F1: 42, F2: -423}
    v3 -= ((0.0 + v3) - -2.25)
    v19 := GenS1{F0: 128, F1: 1, F2: 32767}
    fmt.Println(v3)
    var v20 bool = false
    for fi21 := 0; fi21 < 2; fi21++ {
        v20 = v20
        if len(v4) > 0 {
            v4[0] = (("a" + "foo bar") + "\n\t")
        }
        _ = fi21
    }
    fmt.Println(v20, v3, v20)
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
