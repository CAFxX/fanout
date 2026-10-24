package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 uint64; F1 uint16 }
func genFn0(p0 bool, p1 int64) uint32 {
    if true {
        if (127 <= 1) {
            fmt.Println(p1, p0)
            for fi5 := 0; fi5 < 15; fi5++ {
                for fi6 := 0; fi6 < 13; fi6++ {
                    var v7 float64 = -2.25
                    _ = fi6
                    _ = v7
                }
                _ = fi5
            }
        }
    }
    var v8 uint32 = 336
    v9 := []int{2147483647, 128, -702}
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return (v8 / ((v8) | 1))
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int64 = 9223372036854775807
    v3 := []string{((("" + "\n\t") + (("foo bar" + "") + "foo bar")) + "\n\t"), ("" + ""), "foo bar", "hello"}
    var v4 uint = 4294967295
    v2 = -1
    v10 := GenS1{F0: (v4 - 42), F1: "foo bar"}
    v2 += (v2 & -1)
    if (0.0 >= 0.0) {
        var v11 uint16 = 2
        v12 := GenS1{F0: 705, F1: "a"}
        v2 = v2
        _ = v11
        _ = v12
    }
    v2 -= v2
    v13 := GenS2{F0: 0, F1: 0}
    fmt.Println(v2, v2)
    if (!(false || true)) {
        if true {
            v1[4294967295] = 2
            var v14 int32 = -962
            v3 = append(v3, ((("foo bar" + "foo bar") + "a") + "a"))
            _ = v14
        }
    }
    var v15 uint = (v4 * v4)
    v4 = v4
    var v16 uint32 = uint32(v4)
    var v17 string = ("\n\t" + "hello")
    fmt.Println(v16, v2)
    var v18 bool = true
    var v19 uint32 = v16
    _ = len(v3)
    var v20 int64 = (^(v2))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
