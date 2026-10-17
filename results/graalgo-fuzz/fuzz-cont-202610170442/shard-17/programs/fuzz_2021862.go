package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
func genFn0(p0 uint32, p1 uint16, p2 string) uint32 {
    if true {
        for fi6 := 0; fi6 < 9; fi6++ {
            v7 := []int{}
            p1 -= (p1 & (^(p1)))
            p2 = ((p2 + p2) + "foo bar")
            _ = fi6
            _ = v7
        }
        var v8 int8 = -101
        _ = v8
    }
    var v9 float32 = (((1.5 + 1.5) * (3.14159 - 1.5)) + float32(((3.14159 + 0.0) - 1.5)))
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return (p0 ^ p0)
}
func genFn1(p0 string, p1 int64) uint8 {
    var v10 uint16 = 0
    v11 := []int{int(v10), 127}
    v10 += v10
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 42
}

func main() {
    var v1 int8 = -60
    var v2 string = "a"
    v3 := make(map[int]int)
    v4 := []string{""}
    var v5 float64 = (((float64(float32((0.0 - (1e10 / 3.14159)))) / 1.0) * 0.1) - (1.5 + 0.0))
    var v12 float32 = (float32(((v5 / v5) - (1.5 + ((v5 * v5) - float64(0.0))))) + (-2.25 - (-2.25 * 1e10)))
    v1 = (v1 % (((v1 - (v1 % ((v1) | 1)))) | 1))
    v12 -= 1.5
    v4 = append(v4, ((v2 + "hello") + (v2 + v2)))
    v3[9223372036854775807] = -9223372036854775808
    if (!false) {
        var v13 string = v2
        var v14 uint8 = 1
        for fi15 := 0; fi15 < 8; fi15++ {
            for fi16 := 0; fi16 < 16; fi16++ {
                if ((true || false) && (0 == 2)) {
                    fmt.Println(v14)
                    var v17 uint = 65535
                    if (!false) {
                        _ = len(v3)
                        v17 -= 1
                        v18 := GenS1{F0: v1, F1: (v1 >> (((v17 << (((v17 | (v17 | (v17 / ((v17) | 1))))) & 63))) & 63))}
                        v17 += v17
                        v1 = 51
                        var v19 string = (v13 + v2)
                        _ = v18
                        _ = v19
                    }
                    _ = v17
                }
                v20 := GenS1{F0: (v1 % (((v1 & 1)) | 1)), F1: 127}
                _ = fi16
                _ = v20
            }
            for fi21 := 0; fi21 < 12; fi21++ {
                var v22 uint32 = 65535
                _ = fi21
                _ = v22
            }
            v13 = (v2 + "\n\t")
            _ = fi15
        }
        _ = v13
        _ = v14
    } else {
        switch 100 {
        case -557:
            _ = len(v4)
        case 9223372036854775807:
            if len(v4) > 0 {
                v4[0] = ""
            }
        case -9223372036854775808:
            v1 += 0
        }
    }
    var v23 bool = true
    v12 -= -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
