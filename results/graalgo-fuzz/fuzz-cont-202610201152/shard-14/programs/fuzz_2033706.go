package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0() int64 {
    if false {
        var v6 uint8 = 42
        _ = v6
    } else {
        v7 := []int{-2147483648, 4294967295, 32768}
        v8 := GenS1{F0: 2}
        _ = v7
        _ = v8
    }
    v9 := []int{0, -2147483648}
    var v10 int16 = 32767
    var v11 uint = 1
    for fi12 := 0; fi12 < 10; fi12++ {
        v10 += -296
        _ = fi12
    }
    _ = v10
    _ = v11
    _ = v9
    return 1
}

func main() {
    var v1 int32 = -647
    var v2 int32 = v1
    var v3 bool = true
    v4 := []string{("" + ("" + ("" + "a"))), ("foo bar" + "a")}
    v5 := make(map[uint32]string)
    v1 = -1
    var v13 int32 = v1
    v2 = (v13 / ((529) | 1))
    v13 -= v2
    var v14 uint64 = 18446744073709551615
    fmt.Println(v2, v3)
    var v15 string = "\n\t"
    v16 := GenS1{F0: 70}
    if len(v4) > 0 {
        v4[0] = ("hello" + "foo bar")
    }
    v17 := GenS1{F0: 1}
    var v18 int = 42
    var v19 int64 = -259
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
