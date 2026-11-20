package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 bool; F1 uint; F2 uint16 }
func genFn0() int32 {
    v7 := []int{9223372036854775807, 128, -684, 4294967295}
    v8 := GenS1{F0: 2147483647, F1: 0}
    _ = v7
    _ = v8
    return 1
}

func main() {
    v1 := make(map[string]int)
    var v2 uint = 255
    var v3 int = 259
    var v4 uint64 = 18446744073709551615
    var v5 uint = (v2 << (((v2 & v2)) & 63))
    v6 := []int{(v3 * (v3 | -819)), int(v2), v3}
    v9 := GenS2{F0: 2147483647, F1: -1}
    var v10 int = (-(v3))
    v11 := GenS1{F0: v3, F1: 42}
    var v12 float64 = float64((1.5 * 3.14159))
    v10 += v10
    _ = len(v6)
    v6 = append(v6, (^(v10)))
    var v13 bool = false
    if (214 <= 0) {
        var v14 bool = (0.0 > 0.0)
        v15 := GenS3{F0: v14, F1: 42, F2: 1}
        fmt.Println(v13, v12)
        _ = v14
        _ = v15
    }
    fmt.Println(v5, v10)
    for fi16 := 0; fi16 < 12; fi16++ {
        v5 += v5
        var v17 int16 = 432
        _ = fi16
        _ = v17
    }
    var v18 int64 = -1
    var v19 uint32 = uint32(v2)
    if (v18 > (v18 >> (((v5 & (v2 * v2))) & 63))) {
        var v20 uint = 4294967295
        _ = v20
    }
    if v13 {
        var v21 uint16 = 2
        for fi22 := 0; fi22 < 5; fi22++ {
            switch v3 {
            case 42:
                v4 = 343
            case 1:
                for fi23 := 0; fi23 < 15; fi23++ {
                    if false {
                        v24 := GenS2{F0: -2147483648, F1: -547}
                        fmt.Println(v2, v2, v13)
                        _ = v24
                    }
                    v2 = v2
                    _ = fi23
                }
            }
            v12 += ((float64(-2.25) - v12) + v12)
            _ = fi22
        }
        _ = v21
    } else {
        v18 = (v18 & v18)
        var v25 uint8 = 213
        _ = v25
    }
    if true {
        if (true && ((^(v10)) < (v10 & 127))) {
            v12 = (v12 + v12)
            _ = len(v6)
            var v26 float32 = 0.0
            _ = v26
        }
    }
    if ("hello" > "a") {
        v18 = v18
        if true {
            v5 += v2
            _ = len(v6)
        } else {
            var v27 uint = (v5 - uint(v19))
            _ = v27
        }
    } else {
        fmt.Println(v19)
        for fi28 := 0; fi28 < 6; fi28++ {
            if len(v6) > 0 {
                v6[0] = v10
            }
            var v29 uint64 = (v4 - v4)
            _ = fi28
            _ = v29
        }
        v1[(("\n\t" + ("foo bar" + ("a" + ("hello" + "foo bar")))) + "hello")] = v10
    }
    v4 -= (v4 * 0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
