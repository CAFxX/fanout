package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 int }
func genFn0(p0 uint64, p1 int8) string {
    for fi5 := 0; fi5 < 2; fi5++ {
        var v6 int = fi5
        for fi7 := 0; fi7 < 11; fi7++ {
            fmt.Println(fi7, fi5, fi5)
            if false {
                var v8 string = "hello"
                if (!("\n\t" >= "a")) {
                    v9 := []int{fi7, v6, -603}
                    p0 = (^(p0))
                    var v10 float32 = 3.14159
                    _ = v10
                    _ = v9
                }
                switch v8 {
                case "":
                    v11 := GenS1{F0: v8, F1: 1}
                    _ = v11
                case "\n\t":
                    switch v6 {
                    case -1:
                        v12 := []int{fi7, 127}
                        _ = v12
                    default:
                        if true {
                            v6 += fi5
                        }
                    }
                case "foo bar":
                    var v13 int32 = -104
                    _ = v13
                default:
                    fmt.Println(fi7)
                }
                _ = v8
            } else {
                var v14 uint32 = 42
                _ = v14
            }
            _ = fi7
        }
        _ = fi5
        _ = v6
    }
    p0 -= (p0 << ((1) & 63))
    p0 += 470
    _ = p0
    _ = p1
    return "\n\t"
}

func main() {
    var v1 uint64 = 2
    var v2 bool = true
    v3 := make(map[int]bool)
    var v4 bool = true
    var v15 uint8 = 0
    for fi16 := 0; fi16 < 10; fi16++ {
        var v17 int64 = -9223372036854775808
        delete(v3, (fi16 - fi16))
        _ = fi16
        _ = v17
    }
    v18 := GenS2{F0: 0}
    v19 := []int{1, int(v1), 42}
    var v20 int8 = 0
    v15 += v15
    v3[32768] = (v2 || (uint(v1) < 927))
    var v21 int32 = -2147483648
    v20 = (v20 - v20)
    _ = len(v3)
    switch "\n\t" {
    case "a":
        v19 = append(v19, 333)
    case "":
        v1 = (v1 + (v1 / ((v1) | 1)))
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        v23 := GenS2{F0: fi22}
        delete(v3, fi22)
        var v24 uint = 18446744073709551615
        _ = fi22
        _ = v23
        _ = v24
    }
    var v25 uint32 = 65535
    v26 := GenS2{F0: 2147483647}
    var v27 uint8 = 1
    for fi28 := 0; fi28 < 20; fi28++ {
        v25 = v25
        _ = fi28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
