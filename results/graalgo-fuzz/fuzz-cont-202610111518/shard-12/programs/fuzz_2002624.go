package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 int8; F2 int64 }
func genFn0(p0 int32, p1 int64, p2 uint64) bool {
    v5 := []int{42, 4294967295}
    var v6 int16 = -32768
    fmt.Println(p0)
    if len(v5) > 0 {
        v5[0] = 9223372036854775807
    }
    if (false && ((true || true) || (847 != 32767))) {
        v7 := GenS1{F0: 338}
        _ = v7
    } else {
        var v8 uint16 = 42
        v6 = v6
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return false
}

func main() {
    var v1 float64 = 3.14159
    var v2 int8 = 42
    v3 := make(map[uint32]string)
    var v4 uint8 = 255
    var v9 int8 = v2
    v10 := []int{42, 988, 529}
    v11 := GenS1{F0: 2}
    switch "hello" {
    case "\n\t":
        v12 := GenS2{F0: 286, F1: 0, F2: -9223372036854775808}
        _ = v12
    case "foo bar":
        for fi13 := 0; fi13 < 17; fi13++ {
            v4 = (v4 % (((v4 >> ((559) & 63))) | 1))
            if false {
                v14 := GenS2{F0: 877, F1: v9, F2: 0}
                for fi15 := 0; fi15 < 12; fi15++ {
                    var v16 uint16 = 2
                    var v17 float32 = -2.25
                    v18 := GenS2{F0: 2, F1: v2, F2: 0}
                    _ = fi15
                    _ = v16
                    _ = v17
                    _ = v18
                }
                var v19 bool = (!(-2.25 == (v1 * (v1 + float64(1.5)))))
                _ = v14
                _ = v19
            }
            _ = fi13
        }
    }
    var v20 float32 = (0.1 + 1e10)
    if false {
        var v21 int8 = (v2 - v9)
        if false {
            var v22 uint64 = 0
            v4 += 0
            var v23 int16 = -1
            _ = v22
            _ = v23
        } else {
            v20 = (1.5 * ((v20 / 1.5) - 1.5))
            v9 += 1
        }
        for fi24 := 0; fi24 < 2; fi24++ {
            v1 += 1.5
            _ = len(v3)
            _ = fi24
        }
        _ = v21
    }
    var v25 bool = true
    fmt.Println(v9, v20)
    fmt.Println(v1, v9)
    v25 = true
    v26 := v3[2]
    v26 = v26
    var v27 string = "\n\t"
    var v28 int8 = (v9 - v9)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v9
}
