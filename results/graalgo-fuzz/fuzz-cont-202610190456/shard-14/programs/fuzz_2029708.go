package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
func genFn0() bool {
    v7 := []int{-9223372036854775808}
    var v8 float32 = (float32(((-2.25 * 0.1) * ((-2.25 * ((float64(0.1) + 0.1) + -2.25)) - (1e10 * 1.5)))) / ((0.1 + 3.14159) / (3.14159 + (-2.25 - 1e10))))
    var v9 int32 = 0
    fmt.Println(v9, v9)
    _ = v7
    _ = v8
    _ = v9
    return (true && true)
}

func main() {
    v1 := make(map[int]int)
    v2 := []int{4294967295, 0, 32768, 2147483647}
    v3 := make(map[string]string)
    v4 := make(map[uint32]string)
    v5 := make(map[int]int)
    v6 := make(map[string]uint32)
    var v10 int32 = -1
    v10 -= (v10 + v10)
    var v11 float32 = float32(0.1)
    v11 += (v11 / float32(float64(v11)))
    if len(v2) > 0 {
        v2[0] = 670
    }
    v11 += (0.1 + ((v11 / float32(float64(v11))) * v11))
    v11 = v11
    v12 := v4[788]
    switch v12 {
    case "\n\t":
        fmt.Println(v12, v10, v11)
    case "hello":
        for fi13 := 0; fi13 < 12; fi13++ {
            if true {
                fmt.Println(fi13, fi13)
                v14 := GenS1{F0: 656, F1: 1}
                _ = v14
            } else {
                v10 = (v10 % ((2147483647) | 1))
                v15 := GenS1{F0: 174, F1: 9223372036854775807}
                _ = v15
            }
            _ = fi13
        }
    case "a":
        for fi16 := 0; fi16 < 11; fi16++ {
            if (!(834 == 1)) {
                if ((-2.25 / 1e10) >= (1.5 - float64((((v11 / 1e10) - (3.14159 * (((((v11 / 0.1) + v11) - (v11 + v11)) - v11) + v11))) + v11)))) {
                    switch (("foo bar" + v12) + v12) {
                    case "":
                        _ = len(v5)
                    }
                } else {
                    for fi17 := 0; fi17 < 6; fi17++ {
                        v12 = "hello"
                        v2 = append(v2, fi17)
                        switch "" {
                        case "a":
                            v12 = v12
                        }
                        _ = fi17
                    }
                }
                v18 := GenS1{F0: 4294967295, F1: 0}
                _ = v18
            }
            _ = fi16
        }
    default:
        v19 := GenS1{F0: 341, F1: 1}
        _ = v19
    }
    v12 = "hello"
    var v20 uint = 737
    v21 := GenS1{F0: 2, F1: -855}
    var v22 uint8 = 255
    var v23 int16 = -1
    if false {
        if (2 < 0) {
            _ = len(v2)
            if (-551 >= (^(v10))) {
                var v24 int = 2147483647
                _ = v24
            }
        } else {
            v20 -= (v20 >> ((v20) & 63))
        }
        var v25 float32 = v11
        _ = v25
    } else {
        v12 = "\n\t"
        v26 := v3[("" + "\n\t")]
        var v27 float32 = v11
        _ = v26
        _ = v27
    }
    v6[v12] = 42
    delete(v3, v12)
    var v28 bool = (!false)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
