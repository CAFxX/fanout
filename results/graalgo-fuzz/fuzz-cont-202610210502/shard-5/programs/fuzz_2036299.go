package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 uint }
func genFn0(p0 uint16) int {
    v7 := GenS1{F0: 16, F1: 42}
    v8 := []int{0, 32768, 4294967295}
    _ = len(v8)
    var v9 int64 = 42
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return -1
}

func main() {
    var v1 string = ((("a" + "a") + "") + "a")
    var v2 float32 = 1e10
    v3 := []string{(("hello" + ("" + (((v1 + v1) + "") + (("\n\t" + ("hello" + (v1 + (v1 + ((v1 + v1) + (v1 + v1)))))) + v1)))) + v1), "a", ("\n\t" + "foo bar")}
    var v4 int = -506
    var v5 int64 = -392
    var v6 uint16 = 583
    var v10 uint32 = 4294967295
    var v11 bool = ("hello" <= "foo bar")
    var v12 int8 = 42
    v6 += v6
    var v13 uint32 = uint32(v6)
    v14 := GenS1{F0: 255, F1: (v5 * 9223372036854775807)}
    fmt.Println(v1, v6, v1)
    v15 := GenS2{F0: 961}
    var v16 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
