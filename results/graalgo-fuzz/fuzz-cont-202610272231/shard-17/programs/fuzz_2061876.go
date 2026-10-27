package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0() uint64 {
    v7 := []int{128}
    var v8 int = -1
    v8 += -575
    v8 += v8
    for fi9 := 0; fi9 < 14; fi9++ {
        var v10 int8 = 127
        _ = fi9
        _ = v10
    }
    _ = v7
    _ = v8
    return 222
}
func genFn1() uint {
    var v11 int64 = -1
    for fi12 := 0; fi12 < 7; fi12++ {
        var v13 uint32 = 42
        _ = fi12
        _ = v13
    }
    v11 += v11
    v11 -= 179
    _ = v11
    return 205
}

func main() {
    v1 := []string{}
    v2 := []string{""}
    var v3 int8 = -128
    var v4 uint = 0
    var v5 uint16 = 871
    var v6 uint8 = 55
    var v14 uint16 = v5
    v2 = append(v2, ("a" + "hello"))
    if true {
        v15 := GenS1{F0: 42}
        _ = v15
    }
    for fi16 := 0; fi16 < 7; fi16++ {
        v17 := GenS1{F0: v14}
        var v18 float32 = 0.1
        _ = fi16
        _ = v17
        _ = v18
    }
    v3 = 17
    v19 := GenS1{F0: 65535}
    v2 = append(v2, "")
    v5 = (^(v14))
    var v20 int8 = 127
    v3 = -115
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
