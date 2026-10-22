package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
func genFn0(p0 string, p1 int) int16 {
    v5 := []int{}
    var v6 int = p1
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return -1
}
func genFn1(p0 bool, p1 int16) int8 {
    p1 = 1000
    var v7 float64 = (3.14159 / (1.5 * float64(0.1)))
    _ = p0
    _ = p1
    _ = v7
    return 20
}

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[int]string)
    v3 := []string{"foo bar", "foo bar", "", "\n\t"}
    var v4 bool = (("\n\t" + "") <= "\n\t")
    if len(v3) > 0 {
        v3[0] = (("foo bar" + "hello") + "hello")
    }
    for fi8 := 0; fi8 < 9; fi8++ {
        v1 = 3.14159
        v1 = float64((3.14159 + (-2.25 / 1.5)))
        _ = fi8
    }
    v9 := GenS1{F0: 42, F1: 2}
    var v10 uint = 18446744073709551615
    v1 = float64(-2.25)
    v3 = append(v3, "")
    if true {
        delete(v2, 726)
    }
    var v11 int64 = -9223372036854775808
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v12 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
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
