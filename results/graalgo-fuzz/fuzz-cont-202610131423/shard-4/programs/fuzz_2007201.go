package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 uint8 }
func genFn0(p0 int) uint {
    v4 := []int{4294967295, p0, p0}
    var v5 int32 = 0
    _ = p0
    _ = v4
    _ = v5
    return 775
}

func main() {
    var v1 uint8 = 67
    v2 := make(map[int]int)
    var v3 string = ("foo bar" + (("a" + "a") + ""))
    fmt.Println(v3)
    v3 = v3
    if true {
        for fi6 := 0; fi6 < 15; fi6++ {
            v7 := []int{fi6, fi6, 128, 42}
            var v8 bool = true
            switch (fi6 * (fi6 & (fi6 / (((fi6 | fi6)) | 1)))) {
            case 32768:
                for fi9 := 0; fi9 < 13; fi9++ {
                    var v10 uint = 18446744073709551615
                    var v11 float32 = -2.25
                    var v12 float32 = (v11 - float32(1.5))
                    v12 = 0.0
                    _ = fi9
                    _ = v10
                    _ = v11
                    _ = v12
                }
            case 2147483647:
                for fi13 := 0; fi13 < 20; fi13++ {
                    v1 = (^(v1))
                    var v14 uint16 = 32767
                    var v15 float32 = ((0.0 - (0.1 * float32(0.0))) - (0.1 * 1e10))
                    _ = fi13
                    _ = v14
                    _ = v15
                }
            default:
                for fi16 := 0; fi16 < 7; fi16++ {
                    v17 := GenS1{F0: -918}
                    v7 = append(v7, 128)
                    _ = fi16
                    _ = v17
                }
            }
            _ = fi6
            _ = v7
            _ = v8
        }
        for fi18 := 0; fi18 < 17; fi18++ {
            if (fi18 < (fi18 << ((908) & 63))) {
                v19 := []int{fi18, fi18, int(v1)}
                _ = v19
            }
            v20 := v2[-127]
            _ = fi18
            _ = v20
        }
        v21 := GenS2{F0: 0, F1: (v1 & (v1 * (v1 - v1)))}
        _ = v21
    }
    for fi22 := 0; fi22 < 10; fi22++ {
        v23 := GenS2{F0: 0, F1: v1}
        v24 := []int{fi22, 32768, fi22}
        _ = fi22
        _ = v23
        _ = v24
    }
    if (-67 != -128) {
        var v25 bool = (!false)
        v25 = v25
        v3 = "foo bar"
        _ = v25
    }
    v1 -= v1
    var v26 int16 = 42
    v1 -= 127
    switch v3 {
    case "a":
        var v27 float32 = ((1e10 - -2.25) * 3.14159)
        _ = v27
    }
    var v28 int64 = -1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v26
    _ = v28
    _ = v3
}
