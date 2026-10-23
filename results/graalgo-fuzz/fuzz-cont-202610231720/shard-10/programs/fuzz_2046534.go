package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 int; F1 uint16 }
func genFn0() uint64 {
    var v7 uint8 = 0
    v7 -= 0
    _ = v7
    return 0
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint8 = 255
    v3 := make(map[string]bool)
    v4 := make(map[int]string)
    v5 := []int{}
    var v6 float64 = 0.0
    if false {
        if (true && (42 <= 1)) {
            delete(v1, ("a" + ("" + "foo bar")))
            for fi8 := 0; fi8 < 3; fi8++ {
                v6 += float64((3.14159 * -2.25))
                var v9 float32 = float32(v6)
                if false {
                    v9 -= (0.1 * v9)
                }
                _ = fi8
                _ = v9
            }
        } else {
            v6 += float64(float32((v6 / v6)))
        }
        var v10 float32 = (0.1 * (0.1 * 1.5))
        if (!true) {
            if (false && (genFn0() <= genFn0())) {
                v10 -= 3.14159
            } else {
                var v11 uint8 = 2
                var v12 int64 = -9223372036854775808
                _ = v11
                _ = v12
            }
        }
        _ = v10
    } else {
        v13 := v3[(("\n\t" + "") + ("a" + "hello"))]
        v2 = (v2 % (((^(v2))) | 1))
        _ = v13
    }
    var v14 bool = (!(954 != 99))
    var v15 int32 = 861
    if false {
        v16 := GenS1{F0: (v2 % ((1) | 1)), F1: 387}
        _ = v16
    } else {
        v15 = v15
    }
    if v14 {
        v17 := GenS2{F0: -673, F1: 255}
        _ = v17
    } else {
        fmt.Println(v6)
        v14 = v14
        var v18 string = (("a" + "hello") + ("" + "hello"))
        _ = v18
    }
    var v19 uint64 = 0
    var v20 uint = 2
    if len(v5) > 0 {
        v5[0] = 1
    }
    if v14 {
        fmt.Println(v2, v14)
        for fi21 := 0; fi21 < 9; fi21++ {
            v3[(("hello" + "a") + ("foo bar" + "hello"))] = (-156 < fi21)
            for fi22 := 0; fi22 < 9; fi22++ {
                if v14 {
                    if v14 {
                        fmt.Println(v14, v15, v15)
                        v14 = (!(4294967295 == (v20 << ((v20) & 63))))
                    } else {
                        var v23 int = fi21
                        _ = v23
                    }
                    var v24 uint64 = 4294967295
                    _ = v24
                } else {
                    fmt.Println(v6)
                    v6 += 3.14159
                }
                _ = fi22
            }
            _ = fi21
        }
        v20 -= v20
    } else {
        var v25 uint8 = (v2 ^ v2)
        _ = v25
    }
    v14 = (v14 && v14)
    v6 += v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
