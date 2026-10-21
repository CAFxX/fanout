package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int32, p1 int8, p2 uint8) int8 {
    var v7 float32 = -2.25
    switch "hello" {
    case "\n\t":
        p0 += p0
    }
    var v8 float32 = v7
    v9 := GenS1{F0: -899}
    v10 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return p1
}

func main() {
    v1 := []int{}
    var v2 int64 = -335
    v3 := make(map[string]int)
    var v4 int32 = 42
    v5 := make(map[int]string)
    v6 := make(map[uint32]int)
    v11 := GenS1{F0: -161}
    for fi12 := 0; fi12 < 12; fi12++ {
        _ = len(v1)
        _ = fi12
    }
    var v13 uint64 = 4294967295
    fmt.Println(v2, v2, v13)
    v14 := GenS1{F0: -1}
    fmt.Println(v13, v2)
    var v15 float32 = 3.14159
    switch ("\n\t" + ("hello" + ("\n\t" + ""))) {
    case "a":
        fmt.Println(v15, v4)
    }
    for fi16 := 0; fi16 < 7; fi16++ {
        v4 = 0
        v17 := GenS1{F0: fi16}
        _ = fi16
        _ = v17
    }
    v18 := GenS2{F0: 255}
    fmt.Println(v2)
    var v19 uint32 = 2147483648
    v4 = -127
    v2 += int64(v13)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
