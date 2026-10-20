package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8 }
func genFn0(p0 string, p1 int32) string {
    var v6 bool = (1 == 42)
    var v7 float32 = 0.1
    v8 := []int{-889, -9223372036854775808, 2147483647}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return (p0 + p0)
}

func main() {
    var v1 string = "foo bar"
    var v2 uint8 = 0
    var v3 int = -531
    v4 := make(map[string]int)
    var v5 float32 = 1e10
    v3 -= (v3 * v3)
    var v9 uint16 = 177
    v9 += v9
    v10 := GenS1{F0: 42}
    for fi11 := 0; fi11 < 6; fi11++ {
        if true {
            var v12 uint32 = 2147483648
            v3 -= fi11
            _ = v12
        }
        v13 := GenS1{F0: 4294967295}
        _ = fi11
        _ = v13
    }
    v3 += -2147483648
    v14 := []int{479, 0}
    v2 = v2
    fmt.Println(v9, v2, v3)
    if true {
        if (false && false) {
            v5 = 1e10
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
            fmt.Println(v5, v1, v5)
        } else {
            var v15 float64 = 1.5
            v4[v1] = (v3 << ((4294967295) & 63))
            for fi16 := 0; fi16 < 7; fi16++ {
                var v17 bool = false
                v18 := v4["\n\t"]
                var v19 uint8 = v2
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
            }
            _ = v15
        }
    } else {
        delete(v4, ((v1 + v1) + v1))
        v20 := GenS1{F0: 0}
        v5 = (0.1 - 1.5)
        _ = v20
    }
    fmt.Println(v1, v2)
    if (v9 <= v9) {
        v5 = v5
    }
    var v21 int64 = -814
    _ = len(v4)
    for fi22 := 0; fi22 < 3; fi22++ {
        v23 := GenS1{F0: 2}
        _ = fi22
        _ = v23
    }
    if (65535 != 619) {
        v21 -= (v21 >> ((1) & 63))
        v4["a"] = v3
    }
    fmt.Println(v21, v1)
    v24 := GenS2{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
