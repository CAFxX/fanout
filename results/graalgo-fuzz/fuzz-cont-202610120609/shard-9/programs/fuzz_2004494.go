package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 int16; F2 uint32 }
func genFn0(p0 int8) int32 {
    p0 -= (p0 - (p0 & p0))
    fmt.Println(p0)
    p0 = -38
    var v6 int32 = 42
    _ = p0
    _ = v6
    return v6
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]int)
    var v3 int64 = -939
    var v4 bool = true
    v5 := make(map[int]bool)
    if ((!(!(255 != 188))) && v4) {
        for fi7 := 0; fi7 < 19; fi7++ {
            var v8 bool = v4
            for fi9 := 0; fi9 < 9; fi9++ {
                v3 -= (-(v3))
                v10 := GenS1{F0: 705}
                v11 := []int{(fi9 % ((fi9) | 1)), -1, int(v3)}
                _ = fi9
                _ = v10
                _ = v11
            }
            var v12 bool = v4
            _ = fi7
            _ = v12
            _ = v8
        }
    }
    var v13 float64 = 3.14159
    delete(v5, 2147483647)
    for fi14 := 0; fi14 < 17; fi14++ {
        if v4 {
            for fi15 := 0; fi15 < 4; fi15++ {
                var v16 int8 = -1
                _ = fi15
                _ = v16
            }
            v17 := []int{(fi14 | (fi14 % ((fi14) | 1))), 42, 128}
            if (2147483647 < 2147483647) {
                _ = len(v5)
                for fi18 := 0; fi18 < 14; fi18++ {
                    var v19 int8 = -25
                    v20 := GenS2{F0: (v19 - 125), F1: 1000, F2: 1}
                    if ((false && false) && (!(v19 != 100))) {
                        v13 -= (-2.25 * ((v13 - -2.25) / float64((0.1 + 1e10))))
                    } else {
                        var v21 bool = v4
                        _ = v21
                    }
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                if (!((1 >= 556) || ((0 < 226) || (255 != 0)))) {
                    fmt.Println(v3, v4)
                    for fi22 := 0; fi22 < 12; fi22++ {
                        var v23 uint8 = 169
                        v24 := v1["a"]
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                }
            }
            _ = v17
        }
        _ = fi14
    }
    if (v4 || (true && v4)) {
        for fi25 := 0; fi25 < 3; fi25++ {
            v4 = (!("\n\t" >= ("" + "\n\t")))
            _ = fi25
        }
        fmt.Println(v13, v4)
        v4 = ("a" > "hello")
    }
    for fi26 := 0; fi26 < 7; fi26++ {
        v13 += ((v13 * v13) / 1.0)
        _ = fi26
    }
    var v27 uint32 = 579
    for fi28 := 0; fi28 < 11; fi28++ {
        v5[fi28] = true
        _ = fi28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
