package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 string }
type GenS2 struct { F0 string }
func genFn0(p0 int16) int64 {
    v6 := GenS1{F0: 788, F1: 42, F2: "foo bar"}
    var v7 string = "\n\t"
    var v8 uint32 = 65535
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return -9223372036854775808
}
func genFn1(p0 int, p1 uint8, p2 uint32) bool {
    p1 = 128
    if true {
        if (-1 != 279) {
            for fi9 := 0; fi9 < 12; fi9++ {
                if true {
                    v10 := []int{}
                    _ = v10
                } else {
                    p1 += 128
                    p2 = 2147483648
                }
                var v11 uint8 = (p1 & p1)
                var v12 int8 = int8(v11)
                _ = fi9
                _ = v11
                _ = v12
            }
            p2 += (p2 | (p2 - (p2 | 4294967295)))
            v13 := GenS1{F0: 65535, F1: -9223372036854775808, F2: ("a" + "a")}
            _ = v13
        }
    }
    v14 := GenS1{F0: 758, F1: 42, F2: "\n\t"}
    var v15 int8 = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    return true
}

func main() {
    v1 := []string{"hello", (("" + "") + "a"), "a"}
    v2 := make(map[uint32]int)
    var v3 int32 = 1
    v4 := []int{-799, 42, 127}
    var v5 uint64 = 333
    for fi16 := 0; fi16 < 18; fi16++ {
        v2[65535] = (fi16 + fi16)
        var v17 int64 = 0
        var v18 int32 = v3
        _ = fi16
        _ = v17
        _ = v18
    }
    v19 := GenS1{F0: 312, F1: 9223372036854775807, F2: "\n\t"}
    fmt.Println(v5)
    v20 := GenS2{F0: "a"}
    v3 += (^(v3))
    var v21 int = -200
    var v22 uint = 4294967295
    v5 = v5
    delete(v2, 2)
    var v23 int8 = 1
    v23 = (v23 % ((100) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
