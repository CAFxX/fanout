package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 string }
func genFn0(p0 string) string {
    for fi4 := 0; fi4 < 6; fi4++ {
        v5 := []int{-161, (fi4 | 43)}
        var v6 int8 = 0
        _ = fi4
        _ = v5
        _ = v6
    }
    if (-128 < -18) {
        switch (p0 + p0) {
        case "\n\t":
            v7 := GenS1{F0: p0, F1: 4294967295}
            _ = v7
        }
        for fi8 := 0; fi8 < 1; fi8++ {
            var v9 uint16 = 0
            for fi10 := 0; fi10 < 2; fi10++ {
                var v11 uint16 = v9
                _ = fi10
                _ = v11
            }
            _ = fi8
            _ = v9
        }
        for fi12 := 0; fi12 < 2; fi12++ {
            var v13 float64 = 1.5
            v14 := []int{fi12, (fi12 * (fi12 | fi12)), (fi12 * (fi12 ^ fi12))}
            var v15 int64 = -723
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
    }
    _ = p0
    return (p0 + p0)
}
func genFn1(p0 uint64) uint16 {
    var v16 uint16 = 32767
    fmt.Println(p0)
    fmt.Println(v16, p0)
    v16 += (^(v16))
    var v17 int32 = -2147483648
    _ = p0
    _ = v16
    _ = v17
    return (v16 << ((42) & 63))
}

func main() {
    var v1 int64 = 145
    var v2 float32 = (((1.5 * (1.5 - 0.0)) / 0.1) * ((0.0 + 3.14159) / (3.14159 - 0.0)))
    var v3 bool = (v1 == -1)
    var v18 uint = 65535
    v3 = (v3 && v3)
    for fi19 := 0; fi19 < 19; fi19++ {
        v2 += v2
        v18 -= v18
        _ = fi19
    }
    fmt.Println(v18, v2)
    v20 := []int{}
    v21 := GenS1{F0: ("foo bar" + "hello"), F1: 2}
    for fi22 := 0; fi22 < 3; fi22++ {
        var v23 string = ""
        v24 := GenS2{F0: v23}
        var v25 int32 = 1
        _ = fi22
        _ = v23
        _ = v24
        _ = v25
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
