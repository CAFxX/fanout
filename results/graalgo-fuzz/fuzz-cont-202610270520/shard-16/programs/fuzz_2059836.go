package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 int64; F1 uint; F2 uint }
type GenS3 struct { F0 int8; F1 uint8 }
func genFn0() string {
    v4 := []int{}
    v5 := GenS1{F0: 2, F1: 399}
    _ = v4
    _ = v5
    return (("foo bar" + "") + "a")
}

func main() {
    var v1 uint = 42
    v2 := make(map[uint32]int)
    v3 := make(map[int]int)
    fmt.Println(v1)
    fmt.Println(v1)
    for fi6 := 0; fi6 < 1; fi6++ {
        switch ("a" + "") {
        case "hello":
            v7 := GenS2{F0: -1, F1: 1, F2: (^(v1))}
            _ = v7
        default:
            var v8 int = (fi6 / ((fi6) | 1))
            _ = v8
        }
        v9 := []int{fi6}
        var v10 uint64 = 0
        _ = fi6
        _ = v10
        _ = v9
    }
    if (0.0 != 0.0) {
        var v11 float32 = -2.25
        _ = v11
    } else {
        for fi12 := 0; fi12 < 8; fi12++ {
            for fi13 := 0; fi13 < 9; fi13++ {
                v14 := GenS3{F0: 100, F1: 68}
                var v15 float32 = 1e10
                var v16 int16 = 42
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
            }
            var v17 bool = false
            _ = fi12
            _ = v17
        }
        for fi18 := 0; fi18 < 1; fi18++ {
            v1 = (v1 ^ 4294967295)
            _ = fi18
        }
        v1 -= v1
    }
    v1 += v1
    var v19 uint = 18446744073709551615
    v20 := GenS1{F0: 506, F1: 128}
    for fi21 := 0; fi21 < 18; fi21++ {
        delete(v3, fi21)
        _ = fi21
    }
    fmt.Println(v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
