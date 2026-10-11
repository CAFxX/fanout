package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
func genFn0(p0 uint16, p1 uint16) bool {
    var v4 int8 = 51
    v4 = v4
    _ = p0
    _ = p1
    _ = v4
    return (170 != -281)
}
func genFn1() string {
    var v5 uint = 388
    v5 -= v5
    var v6 uint8 = 195
    _ = v5
    _ = v6
    return (("foo bar" + ("\n\t" + "foo bar")) + "\n\t")
}

func main() {
    v1 := []int{-1, 322, 128}
    var v2 uint16 = 327
    var v3 string = "hello"
    fmt.Println(v3)
    var v7 uint16 = v2
    v2 += v7
    var v8 float32 = 0.0
    var v9 int32 = 0
    var v10 int = 1
    fmt.Println(v9, v8)
    if true {
        var v11 int64 = -9223372036854775808
        if false {
            v2 += v2
        } else {
            v3 = v3
        }
        _ = v11
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 int8 = 100
        _ = len(v1)
        v14 := GenS1{F0: 18446744073709551615, F1: v9}
        _ = fi12
        _ = v13
        _ = v14
    }
    v3 = (v3 + "")
    var v15 int8 = 127
    if true {
        var v16 int64 = int64(v7)
        var v17 uint32 = 155
        var v18 bool = (!false)
        _ = v16
        _ = v17
        _ = v18
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
