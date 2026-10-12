package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 string) bool {
    p0 = p0
    v6 := GenS1{F0: -386}
    switch p0 {
    case "":
        v7 := GenS1{F0: -2147483648}
        _ = v7
    }
    _ = p0
    _ = v6
    return false
}

func main() {
    var v1 uint8 = 127
    var v2 string = "foo bar"
    var v3 int16 = -32768
    var v4 bool = (42 > 777)
    v5 := []int{4294967295, 311}
    var v8 float32 = float32(0.0)
    v2 = v2
    v9 := GenS2{F0: 255}
    var v10 int32 = 2147483647
    v10 = 0
    fmt.Println(v4, v3)
    v8 += -2.25
    var v11 uint32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
