package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 uint64; F2 int64 }
func genFn0(p0 int32, p1 int) int64 {
    v7 := GenS1{F0: -8, F1: 710, F2: true}
    var v8 uint64 = 0
    var v9 string = "foo bar"
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return -9223372036854775808
}

func main() {
    var v1 uint32 = 1
    var v2 uint8 = 0
    v3 := make(map[string]uint32)
    v4 := make(map[string]bool)
    v5 := []string{"\n\t", "a", "\n\t"}
    v6 := make(map[int]bool)
    fmt.Println(v2, v1)
    v2 -= 2
    v10 := GenS1{F0: 42, F1: 42, F2: (v1 >= uint32(v2))}
    v5 = append(v5, ("a" + ""))
    v11 := GenS2{F0: 78}
    for fi12 := 0; fi12 < 1; fi12++ {
        for fi13 := 0; fi13 < 9; fi13++ {
            v1 = v1
            var v14 bool = (128 > (v2 >> ((0) & 63)))
            _ = fi13
            _ = v14
        }
        v15 := GenS1{F0: 0, F1: -726, F2: false}
        _ = fi12
        _ = v15
    }
    _ = len(v4)
    fmt.Println(v1)
    v16 := GenS3{F0: 42, F1: 0, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
