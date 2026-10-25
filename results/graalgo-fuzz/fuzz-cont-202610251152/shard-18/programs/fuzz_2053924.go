package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64 }
func genFn0(p0 int64, p1 int32) uint {
    for fi7 := 0; fi7 < 18; fi7++ {
        fmt.Println(p1, fi7)
        fmt.Println(p0, p0)
        v8 := []int{fi7, -1, 0}
        _ = fi7
        _ = v8
    }
    switch "\n\t" {
    case "hello":
        p0 += p0
    default:
        v9 := GenS1{F0: 1, F1: 255}
        _ = v9
    }
    var v10 int32 = (-(p1))
    v11 := GenS1{F0: 65535, F1: 42}
    var v12 bool = false
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    return 2
}

func main() {
    var v1 float64 = 0.0
    var v2 uint = 914
    v3 := make(map[string]string)
    v4 := []string{("" + "foo bar"), ("a" + "hello"), "hello"}
    v5 := make(map[uint32]int)
    var v6 uint16 = 789
    for fi13 := 0; fi13 < 7; fi13++ {
        v14 := GenS1{F0: 0, F1: v6}
        _ = fi13
        _ = v14
    }
    v1 -= v1
    delete(v5, 509)
    v15 := GenS2{F0: -977}
    fmt.Println(v6, v1)
    var v16 float64 = (((0.0 / 1.5) + v1) + v1)
    v17 := GenS1{F0: (^(v2)), F1: 724}
    var v18 int16 = 0
    v19 := GenS2{F0: -531}
    if (2 >= 193) {
        fmt.Println(v2)
    }
    if (2 <= 2147483647) {
        var v20 float64 = (v16 + 1.5)
        _ = v20
    }
    v6 = 32767
    v21 := v5[4294967295]
    if true {
        switch v21 {
        case 796:
            if (82 == 35) {
                var v22 uint = 42
                for fi23 := 0; fi23 < 17; fi23++ {
                    for fi24 := 0; fi24 < 2; fi24++ {
                        v25 := GenS3{F0: int64(v18)}
                        v2 = uint(v6)
                        _ = fi24
                        _ = v25
                    }
                    var v26 bool = false
                    _ = fi23
                    _ = v26
                }
                _ = v22
            }
        case 401:
            v6 += v6
        case 0:
            for fi27 := 0; fi27 < 9; fi27++ {
                fmt.Println(v16, fi27, v16)
                for fi28 := 0; fi28 < 3; fi28++ {
                    v2 -= v2
                    v16 = (v16 - (v16 - 0.1))
                    fmt.Println(v1)
                    _ = fi28
                }
                v21 = fi27
                _ = fi27
            }
        default:
            v29 := GenS3{F0: 42}
            _ = v29
        }
    }
    for fi30 := 0; fi30 < 7; fi30++ {
        fmt.Println(v16, v16)
        _ = fi30
    }
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
