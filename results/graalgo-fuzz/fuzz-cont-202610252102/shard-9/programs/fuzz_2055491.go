package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
func genFn0(p0 int8) int {
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 string = "foo bar"
        _ = fi6
        _ = v7
    }
    v8 := []int{-9223372036854775808, 785, 2147483647, -1}
    p0 = p0
    v9 := GenS1{F0: 42, F1: 127}
    _ = p0
    _ = v8
    _ = v9
    return -1
}
func genFn1(p0 int16, p1 uint8, p2 bool) int32 {
    v10 := []int{}
    var v11 uint64 = 42
    for fi12 := 0; fi12 < 13; fi12++ {
        var v13 int8 = int8(fi12)
        if true {
            if false {
                switch "foo bar" {
                case "a":
                    if true {
                        v14 := GenS1{F0: (^(p1)), F1: 140}
                        if false {
                            var v15 uint64 = v11
                            _ = v15
                        }
                        _ = v14
                    }
                case "":
                    for fi16 := 0; fi16 < 10; fi16++ {
                        var v17 float32 = 3.14159
                        _ = fi16
                        _ = v17
                    }
                default:
                    if false {
                        p1 -= (^(p1))
                    }
                }
                var v18 float32 = 0.0
                _ = v18
            }
            if (("" + "a") == "foo bar") {
                v11 = v11
                p2 = p2
                var v19 int32 = -54
                _ = v19
            }
            v13 += v13
        }
        _ = fi12
        _ = v13
    }
    var v20 int32 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v20
    return (^(v20))
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]int)
    var v3 uint64 = 2
    var v4 float32 = (-2.25 / (float32(0.1) - ((0.0 / 1.5) + 0.0)))
    var v5 int64 = -783
    v3 += v3
    switch "a" {
    case "a":
        v5 += 1
    default:
        var v21 int16 = 947
        _ = v21
    }
    var v22 uint = 18446744073709551615
    delete(v1, "a")
    for fi23 := 0; fi23 < 18; fi23++ {
        var v24 bool = false
        _ = fi23
        _ = v24
    }
    fmt.Println(v4, v3, v3)
    fmt.Println(v22, v5, v4)
    if (20 < 0) {
        v22 -= v22
    }
    v25 := GenS1{F0: 205, F1: 2}
    var v26 uint8 = 128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
