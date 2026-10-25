package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 uint16 }

func main() {
    var v1 float64 = 1.5
    var v2 uint64 = 4294967295
    var v3 float32 = -2.25
    var v4 uint = 4294967295
    for fi5 := 0; fi5 < 5; fi5++ {
        v6 := GenS1{F0: v4}
        var v7 uint16 = 42
        v8 := []int{1, 32768}
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 797}
    var v10 string = (("hello" + ("" + "")) + ((("hello" + "") + "") + ((("\n\t" + "a") + "\n\t") + ("hello" + "a"))))
    fmt.Println(v10, v4, v1)
    var v11 uint = v4
    v12 := GenS2{F0: 1, F1: 524}
    v10 = ("a" + "\n\t")
    v3 += (v3 + v3)
    var v13 int64 = -600
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
