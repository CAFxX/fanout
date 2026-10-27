package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
func genFn0(p0 string, p1 string, p2 uint16) uint8 {
    if (1 >= 42) {
        for fi6 := 0; fi6 < 12; fi6++ {
            var v7 uint = 65535
            _ = fi6
            _ = v7
        }
        p0 = p1
    }
    var v8 int = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 162
}

func main() {
    var v1 int16 = 0
    var v2 int = -401
    v3 := make(map[string]string)
    var v4 uint64 = 18446744073709551615
    var v5 uint = 2
    for fi9 := 0; fi9 < 6; fi9++ {
        var v10 int32 = 1
        v10 += (v10 % (((v10 >> ((377) & 63))) | 1))
        v1 = v1
        _ = fi9
        _ = v10
    }
    v11 := []int{32767, v2, v2}
    fmt.Println(v4)
    var v12 uint8 = 0
    v1 -= 1
    var v13 float32 = 3.14159
    v14 := GenS1{F0: 958, F1: 204}
    var v15 uint8 = (v12 * 0)
    if (!((^(v4)) > v4)) {
        _ = len(v11)
        v15 += uint8(v4)
    }
    for fi16 := 0; fi16 < 19; fi16++ {
        v1 = v1
        v12 = v12
        delete(v3, "foo bar")
        _ = fi16
    }
    v17 := GenS1{F0: 497, F1: 467}
    if (genFn0("foo bar", (("hello" + ("foo bar" + "foo bar")) + ("\n\t" + "a")), 794) < 0) {
        for fi18 := 0; fi18 < 2; fi18++ {
            if ((-(v1)) <= v1) {
                v1 = -1
                v1 -= v1
            }
            _ = fi18
        }
    }
    v15 = (v12 ^ genFn0("\n\t", ("" + "a"), 42))
    if (!((v4 >> ((v5) & 63)) > uint64(v12))) {
        var v19 float64 = 0.1
        v2 = v2
        _ = v19
    }
    var v20 int16 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
