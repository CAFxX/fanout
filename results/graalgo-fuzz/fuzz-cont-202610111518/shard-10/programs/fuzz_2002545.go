package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 int16, p1 uint64) uint32 {
    if ((p0 - 1) < 0) {
        v5 := GenS1{F0: ""}
        _ = v5
    } else {
        fmt.Println(p1, p0)
        if false {
            for fi6 := 0; fi6 < 9; fi6++ {
                v7 := GenS1{F0: "\n\t"}
                _ = fi6
                _ = v7
            }
            var v8 string = "hello"
            _ = v8
        } else {
            fmt.Println(p0)
        }
        fmt.Println(p1, p1)
    }
    for fi9 := 0; fi9 < 15; fi9++ {
        var v10 int64 = 9223372036854775807
        var v11 uint8 = 128
        if false {
            if (1e10 >= ((((0.1 + -2.25) - -2.25) / (1e10 + -2.25)) + ((((float64(0.1) - float64(0.0)) + float64((float32((1.5 + ((((0.1 * 1.5) - 0.1) + (3.14159 - -2.25)) * float64(3.14159)))) * 0.0))) / 1e10) - 0.1))) {
                for fi12 := 0; fi12 < 20; fi12++ {
                    var v13 int16 = (^(p0))
                    if (3.14159 == 1.5) {
                        v14 := []int{32768, int(v10)}
                        _ = v14
                    }
                    var v15 uint8 = (v11 | (v11 & v11))
                    _ = fi12
                    _ = v13
                    _ = v15
                }
                if (false || (2147483647 > 2147483647)) {
                    p0 = (p0 % ((p0) | 1))
                }
            }
            v16 := []int{fi9, fi9}
            _ = v16
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    p0 -= (p0 >> ((4294967295) & 63))
    var v17 int64 = 0
    _ = p0
    _ = p1
    _ = v17
    return 42
}
func genFn1(p0 uint64, p1 uint8) string {
    p0 += p0
    p1 = p1
    _ = p0
    _ = p1
    return "foo bar"
}

func main() {
    var v1 float32 = 3.14159
    var v2 bool = false
    v3 := make(map[string]int)
    v4 := []string{"\n\t", "", ""}
    var v18 int64 = 128
    _ = len(v3)
    if v2 {
        v18 -= (v18 >> ((197) & 63))
        v2 = v2
    }
    v18 -= (-(v18))
    if v2 {
        var v19 uint16 = 333
        var v20 float64 = ((1.5 - ((3.14159 - 0.1) + ((3.14159 + 3.14159) * (1.5 * 1.5)))) / 1.5)
        v4 = append(v4, "hello")
        _ = v19
        _ = v20
    }
    v21 := v3["a"]
    var v22 uint16 = 482
    v21 = v21
    v18 = (v18 - v18)
    for fi23 := 0; fi23 < 15; fi23++ {
        fmt.Println(v1)
        var v24 uint = 562
        v25 := GenS1{F0: "\n\t"}
        _ = fi23
        _ = v24
        _ = v25
    }
    var v26 uint64 = 4294967295
    v1 -= 1e10
    fmt.Println(v2, v2)
    var v27 uint16 = (v22 % ((v22) | 1))
    var v28 uint8 = 1
    v3[""] = (v21 | -430)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
