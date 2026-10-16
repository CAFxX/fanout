package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 int }
type GenS2 struct { F0 uint }
func genFn0(p0 int64, p1 int8) int64 {
    p0 = p0
    var v5 string = "a"
    v6 := []int{}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return p0
}
func genFn1() int {
    v7 := GenS1{F0: 2, F1: 42, F2: 1}
    var v8 float64 = (3.14159 + 1e10)
    v9 := []int{}
    var v10 float32 = 1e10
    if true {
        var v11 uint64 = 4294967295
        _ = v11
    }
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return -9223372036854775808
}

func main() {
    var v1 int8 = 100
    var v2 bool = (((1.5 / (-2.25 + 1.5)) != 3.14159) || (65535 < 290))
    var v3 bool = (float64(3.14159) < (1.5 - (-2.25 * 1.5)))
    var v4 uint = 597
    v12 := []int{470, -996, 2147483647, genFn1()}
    var v13 string = "foo bar"
    v13 = (((((v13 + "\n\t") + v13) + v13) + v13) + "a")
    var v14 int32 = -477
    fmt.Println(v14, v2, v13)
    for fi15 := 0; fi15 < 10; fi15++ {
        if v2 {
            var v16 int32 = (v14 / ((43) | 1))
            var v17 int8 = (v1 >> (((v4 / ((189) | 1))) & 63))
            v3 = (!(128 == 1))
            _ = v16
            _ = v17
        } else {
            switch v13 {
            case "\n\t":
                var v18 int8 = v1
                _ = v18
            case "hello":
                var v19 int32 = -845
                _ = v19
            case "a":
                var v20 int32 = v14
                _ = v20
            }
        }
        v21 := GenS1{F0: 4294967295, F1: 9223372036854775807, F2: fi15}
        _ = fi15
        _ = v21
    }
    v22 := GenS2{F0: v4}
    v13 = (v13 + v13)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
