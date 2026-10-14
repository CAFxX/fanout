package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
func genFn0(p0 string, p1 string) int8 {
    var v5 float32 = (-2.25 / 0.1)
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 uint = 4294967295
        if ("foo bar" > "hello") {
            var v8 uint64 = 930
            p0 = ((p1 + (("\n\t" + "") + p1)) + p0)
            _ = v8
        } else {
            for fi9 := 0; fi9 < 3; fi9++ {
                fmt.Println(v7, fi9, v5)
                if false {
                    var v10 float32 = ((v5 * (0.1 * 0.0)) / (v5 + v5))
                    v11 := GenS1{F0: 1, F1: -1}
                    v12 := []int{fi6, (fi9 / ((fi6) | 1)), fi9}
                    _ = v10
                    _ = v11
                    _ = v12
                } else {
                    var v13 int32 = -313
                    v13 = 42
                    _ = v13
                }
                var v14 int64 = 465
                _ = fi9
                _ = v14
            }
            for fi15 := 0; fi15 < 13; fi15++ {
                p0 = ""
                p1 = p1
                _ = fi15
            }
            v16 := []int{-489}
            _ = v16
        }
        _ = fi6
        _ = v7
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v5
    return 0
}

func main() {
    var v1 float64 = 1e10
    v2 := []string{"foo bar"}
    v3 := make(map[uint32]uint32)
    var v4 uint8 = 0
    var v17 string = ("\n\t" + "a")
    fmt.Println(v17, v17)
    for fi18 := 0; fi18 < 7; fi18++ {
        if true {
            _ = len(v2)
            fmt.Println(v4, v17, v17)
        }
        fmt.Println(v17)
        if true {
            v19 := GenS1{F0: -128, F1: 42}
            if false {
                _ = len(v2)
                for fi20 := 0; fi20 < 13; fi20++ {
                    var v21 uint8 = (v4 | (^(v4)))
                    switch v17 {
                    case "\n\t":
                        v4 = (v4 << ((255) & 63))
                    }
                    switch (fi20 * fi18) {
                    case 42:
                        if (42 <= 1) {
                            v22 := GenS1{F0: 42, F1: 778}
                            _ = v22
                        } else {
                            v4 += v4
                            v4 = v4
                            var v23 uint16 = 255
                            _ = v23
                        }
                    }
                    _ = fi20
                    _ = v21
                }
                if (!(18446744073709551615 == 2)) {
                    delete(v3, 4294967295)
                } else {
                    v2 = append(v2, "a")
                }
            }
            v24 := v3[922]
            _ = v19
            _ = v24
        }
        _ = fi18
    }
    v3[215] = 1
    v17 = v17
    for fi25 := 0; fi25 < 3; fi25++ {
        fmt.Println(v4, v17)
        var v26 int64 = 1
        v1 += ((v1 - ((v1 / (v1 + (v1 - v1))) / float64(float32((1e10 / 1e10))))) - v1)
        _ = fi25
        _ = v26
    }
    var v27 uint16 = 634
    if (!(4294967295 > 1)) {
        v1 = ((v1 / float64((0.1 * -2.25))) - float64(-2.25))
    } else {
        v4 -= v4
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v27
    _ = v3
    _ = v4
}
