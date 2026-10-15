package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 int8 }
type GenS2 struct { F0 bool; F1 string; F2 string }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 string; F1 uint64 }
func genFn0() bool {
    v4 := GenS1{F0: 253, F1: "\n\t", F2: -1}
    var v5 int = 0
    _ = v4
    _ = v5
    return ((1.5 / (-2.25 - 3.14159)) == 0.0)
}

func main() {
    var v1 int = 0
    v2 := make(map[uint32]bool)
    v3 := make(map[int]bool)
    var v6 int64 = -9223372036854775808
    v7 := GenS2{F0: (!false), F1: "hello", F2: "foo bar"}
    v8 := []int{}
    delete(v2, 22)
    var v9 float64 = (1e10 * float64(1.5))
    v10 := GenS3{F0: v1, F1: 7}
    fmt.Println(v9, v1)
    for fi11 := 0; fi11 < 4; fi11++ {
        v8 = append(v8, (v1 | 2147483647))
        if false {
            fmt.Println(v6)
            v9 = (float64((float32(v9) / (-2.25 + 1.5))) + ((v9 / v9) * v9))
        }
        var v12 uint64 = 925
        _ = fi11
        _ = v12
    }
    var v13 uint64 = 0
    for fi14 := 0; fi14 < 10; fi14++ {
        v15 := GenS1{F0: v13, F1: "hello", F2: 100}
        _ = fi14
        _ = v15
    }
    v16 := GenS4{F0: ("hello" + "hello"), F1: (v13 % (((^(v13))) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
