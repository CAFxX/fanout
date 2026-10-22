package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 bool, p1 int8) uint8 {
    v5 := []int{}
    var v6 string = "\n\t"
    var v7 int8 = -44
    if len(v5) > 0 {
        v5[0] = -1
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 94
}
func genFn1() int32 {
    var v8 string = "hello"
    v9 := []int{2147483647, 42, 127, 759}
    fmt.Println(v8)
    _ = v8
    _ = v9
    return -2147483648
}

func main() {
    var v1 uint = 846
    var v2 uint32 = 0
    var v3 string = "a"
    var v4 string = ("" + v3)
    v4 = "\n\t"
    v1 += v1
    fmt.Println(v3, v4)
    v10 := []int{-778, 1}
    _ = len(v10)
    var v11 uint = 724
    _ = len(v10)
    v12 := GenS1{F0: v3}
    if len(v10) > 0 {
        v10[0] = -763
    }
    fmt.Println(v2)
    var v13 int32 = -1
    if len(v10) > 0 {
        v10[0] = 0
    }
    var v14 int64 = 1
    if false {
        fmt.Println(v4, v1)
        for fi15 := 0; fi15 < 10; fi15++ {
            v4 = "\n\t"
            fmt.Println(v14)
            for fi16 := 0; fi16 < 4; fi16++ {
                var v17 int8 = 0
                var v18 int16 = 1000
                _ = fi16
                _ = v17
                _ = v18
            }
            _ = fi15
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
