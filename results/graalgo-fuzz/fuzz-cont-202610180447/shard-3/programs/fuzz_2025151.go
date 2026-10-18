package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
func genFn0(p0 int64) uint32 {
    v5 := []int{}
    v6 := GenS1{F0: (-(p0)), F1: 1}
    var v7 bool = (p0 == 1)
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 457
}
func genFn1() int64 {
    v8 := []int{127, 42}
    v8 = append(v8, -156)
    switch "foo bar" {
    case "a":
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
    default:
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
    }
    _ = v8
    return 42
}

func main() {
    var v1 uint16 = 255
    var v2 int64 = -1
    var v3 uint16 = v1
    var v4 string = "\n\t"
    v9 := GenS1{F0: v2, F1: -1}
    v10 := []int{}
    if true {
        _ = len(v10)
    }
    v11 := GenS1{F0: v2, F1: 42}
    _ = len(v10)
    var v12 uint16 = v1
    for fi13 := 0; fi13 < 15; fi13++ {
        if true {
            fmt.Println(v1, v2, v4)
        } else {
            var v14 int32 = 184
            var v15 uint = 1
            _ = v14
            _ = v15
        }
        v12 += v12
        _ = fi13
    }
    for fi16 := 0; fi16 < 17; fi16++ {
        if (!(false || (!(128 >= 35)))) {
            fmt.Println(v4)
        }
        v10 = append(v10, -1)
        _ = fi16
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
