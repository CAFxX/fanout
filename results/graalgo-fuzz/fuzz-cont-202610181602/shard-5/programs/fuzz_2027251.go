package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 int }

func main() {
    v1 := make(map[int]int)
    v2 := []int{2147483647, 88}
    v3 := []int{}
    var v4 float64 = 3.14159
    v5 := make(map[int]string)
    v6 := []string{}
    v4 += -2.25
    if false {
        if (724 > -425) {
            var v7 uint32 = 541
            v4 -= 0.1
            var v8 int = -639
            _ = v7
            _ = v8
        } else {
            v9 := GenS1{F0: -2147483648, F1: 100, F2: 651}
            if true {
                v4 += (v4 - (v4 / float64(-2.25)))
                if (254 == 1) {
                    var v10 int64 = 42
                    v10 += -9223372036854775808
                    _ = v10
                }
                for fi11 := 0; fi11 < 13; fi11++ {
                    if true {
                        var v12 string = ("a" + "hello")
                        _ = v12
                    }
                    _ = fi11
                }
            } else {
                for fi13 := 0; fi13 < 18; fi13++ {
                    v4 = (v4 + (v4 - (1.5 - -2.25)))
                    fmt.Println(fi13)
                    _ = fi13
                }
            }
            _ = v9
        }
        v4 -= ((((1.5 / (((float64(3.14159) / v4) * 0.0) + v4)) - (0.1 - (-2.25 + v4))) - v4) + v4)
    } else {
        v4 = 1e10
        fmt.Println(v4)
    }
    v4 = 3.14159
    v4 = (((v4 - v4) - v4) * v4)
    v4 = v4
    v4 = (v4 * 0.1)
    var v14 uint16 = 32767
    for fi15 := 0; fi15 < 8; fi15++ {
        v6 = append(v6, ("\n\t" + "hello"))
        var v16 uint = 2
        _ = fi15
        _ = v16
    }
    var v17 int64 = 917
    v6 = append(v6, "hello")
    v2 = append(v2, -591)
    v14 = v14
    v4 -= (v4 / (float64(0.1) * float64(1.5)))
    v14 = v14
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
