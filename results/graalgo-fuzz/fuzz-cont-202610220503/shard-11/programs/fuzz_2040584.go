package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 bool }
func genFn0(p0 int8, p1 uint64) string {
    v6 := GenS1{F0: (p0 >> ((42) & 63)), F1: 185}
    var v7 string = ("" + ("hello" + "a"))
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return v7
}

func main() {
    var v1 int64 = 42
    var v2 float32 = (1.5 - (-2.25 / 3.14159))
    var v3 float32 = v2
    v4 := make(map[string]uint32)
    v5 := make(map[int]bool)
    var v8 string = (("hello" + "") + "foo bar")
    if (-985 < -1) {
        var v9 uint64 = 18446744073709551615
        _ = v9
    } else {
        var v10 float32 = v3
        _ = v10
    }
    v11 := []int{-1, -495, 127, 42}
    v12 := GenS1{F0: 42, F1: 560}
    for fi13 := 0; fi13 < 4; fi13++ {
        if false {
            var v14 uint16 = 65535
            v11 = append(v11, fi13)
            _ = v14
        } else {
            switch (fi13 ^ 564) {
            case 719:
                v11 = append(v11, (fi13 >> ((767) & 63)))
            case 694:
                v1 = -9223372036854775808
            case 9223372036854775807:
                if false {
                    if true {
                        var v15 int8 = -1
                        v2 -= (((float32((1e10 + 1.5)) / 1.0) + v2) - 1.5)
                        _ = v15
                    }
                }
            }
        }
        var v16 uint64 = 183
        _ = fi13
        _ = v16
    }
    var v17 bool = (127 > 42)
    var v18 int64 = (v1 * (v1 >> ((0) & 63)))
    switch -2147483648 {
    case 0:
        v19 := GenS2{F0: false}
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
