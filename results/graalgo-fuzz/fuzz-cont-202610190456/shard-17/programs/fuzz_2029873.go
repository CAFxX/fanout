package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 uint64; F1 uint64; F2 uint64 }
func genFn0(p0 int32) bool {
    var v5 uint = 42
    var v6 int8 = -1
    _ = p0
    _ = v5
    _ = v6
    return false
}

func main() {
    v1 := []string{((("hello" + (((("a" + "hello") + ("a" + "hello")) + "\n\t") + "foo bar")) + ("" + "a")) + "\n\t"), "foo bar", (("\n\t" + ("\n\t" + "\n\t")) + "foo bar"), "foo bar"}
    var v2 float32 = (((3.14159 - 0.1) - (3.14159 - 0.1)) - 3.14159)
    v3 := make(map[uint32]uint32)
    var v4 int8 = -128
    var v7 float32 = float32(1e10)
    v8 := GenS1{F0: 153, F1: 1}
    v7 += v2
    var v9 float32 = (((0.1 / 1.0) * 1e10) * 1.5)
    fmt.Println(v7, v7)
    delete(v3, 2147483648)
    var v10 int16 = 0
    delete(v3, 681)
    if true {
        v3[0] = 42
        var v11 uint = 4294967295
        _ = v11
    }
    var v12 string = ("foo bar" + "a")
    for fi13 := 0; fi13 < 17; fi13++ {
        v1 = append(v1, v12)
        if (!((262 != 1) && true)) {
            if true {
                var v14 int32 = 1
                v15 := GenS1{F0: 128, F1: 1}
                for fi16 := 0; fi16 < 4; fi16++ {
                    var v17 int64 = 1
                    _ = fi16
                    _ = v17
                }
                _ = v14
                _ = v15
            } else {
                switch (fi13 & (^(fi13))) {
                case 32768:
                    if (!(11 == 1)) {
                        v7 = (1.5 - ((v9 + v7) + (v7 + 0.0)))
                    } else {
                        var v18 int8 = -1
                        fmt.Println(v9, v12, v4)
                        _ = v18
                    }
                case -801:
                    v19 := v3[0]
                    _ = v19
                case 0:
                    switch ("\n\t" + v12) {
                    case "":
                        var v20 float32 = -2.25
                        _ = v20
                    case "foo bar":
                        v10 = 1000
                    }
                }
            }
        } else {
            for fi21 := 0; fi21 < 10; fi21++ {
                fmt.Println(v12, v2)
                _ = fi21
            }
            v9 = (float32((float64(float32(float64(v9))) - 1.5)) * (v2 / float32(-2.25)))
        }
        _ = fi13
    }
    v22 := GenS2{F0: 747, F1: 4294967295, F2: 966}
    v12 = v12
    v7 -= (((3.14159 * float32(1.5)) * (0.1 * v2)) * (1e10 + ((((0.1 - 0.1) * 0.0) - v7) / v7)))
    if (!(255 >= 0)) {
        v10 -= (v10 % ((v10) | 1))
        if len(v1) > 0 {
            v1[0] = v12
        }
    }
    delete(v3, 65535)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
