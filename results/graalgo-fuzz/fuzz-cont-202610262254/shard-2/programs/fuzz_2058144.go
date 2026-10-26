package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 int8 }
type GenS2 struct { F0 int32; F1 int64; F2 int8 }
func genFn0() int64 {
    var v5 uint64 = 827
    var v6 bool = (2147483647 > 4294967295)
    v7 := []int{42, 0, 127}
    switch 128 {
    case 0:
        v7 = append(v7, 259)
    case 2147483647:
        v5 = v5
    case -9223372036854775808:
        if v6 {
            var v8 int32 = 42
            v8 = (v8 - v8)
            v6 = v6
            _ = v8
        }
    }
    _ = v5
    _ = v6
    _ = v7
    return 0
}

func main() {
    v1 := []string{"a", "a", "\n\t", "hello"}
    var v2 uint = 18446744073709551615
    v3 := []string{(("" + ((("foo bar" + "foo bar") + "\n\t") + "hello")) + "\n\t"), (("hello" + ("foo bar" + ("\n\t" + "foo bar"))) + "")}
    v4 := make(map[int]uint32)
    v2 += 668
    v9 := GenS1{F0: 1, F1: 124, F2: 100}
    var v10 int64 = 595
    var v11 string = "a"
    v12 := GenS1{F0: 42, F1: 95, F2: 1}
    delete(v4, -153)
    var v13 string = (v11 + v11)
    v14 := GenS2{F0: 2147483647, F1: v10, F2: -110}
    v10 += (v10 / ((v10) | 1))
    for fi15 := 0; fi15 < 12; fi15++ {
        v16 := GenS2{F0: 294, F1: v10, F2: 1}
        _ = fi15
        _ = v16
    }
    v17 := GenS2{F0: -836, F1: 1, F2: 105}
    var v18 uint = 255
    _ = len(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
