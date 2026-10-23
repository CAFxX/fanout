package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 int8 }
type GenS2 struct { F0 int8; F1 int8; F2 int }
func genFn0(p0 string, p1 int32, p2 int64) string {
    v4 := GenS1{F0: 2147483647, F1: -269, F2: 100}
    var v5 uint8 = 32
    switch 645 {
    case -1:
        fmt.Println(p2, p0)
    case 32768:
        fmt.Println(p1, p0, p0)
    case -9223372036854775808:
        fmt.Println(v5, p2)
    }
    switch ("\n\t" + p0) {
    case "a":
        fmt.Println(p2, p0, v5)
    }
    v6 := []int{460}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return (((p0 + p0) + p0) + (p0 + ""))
}

func main() {
    var v1 string = ("a" + "")
    var v2 int16 = 785
    var v3 bool = (!true)
    fmt.Println(v3, v2)
    var v7 uint64 = 4294967295
    var v8 int32 = 2147483647
    v1 = v1
    v9 := []int{}
    for fi10 := 0; fi10 < 3; fi10++ {
        v11 := GenS1{F0: 42, F1: -9223372036854775808, F2: 42}
        var v12 uint32 = 1
        _ = fi10
        _ = v11
        _ = v12
    }
    switch (v1 + v1) {
    case "\n\t":
        v13 := GenS2{F0: -1, F1: 127, F2: -2147483648}
        _ = v13
    case "hello":
        v8 = (v8 & v8)
    }
    v3 = (!false)
    for fi14 := 0; fi14 < 1; fi14++ {
        v3 = (1e10 != (1e10 + ((1e10 - -2.25) * 1.5)))
        if len(v9) > 0 {
            v9[0] = (fi14 & fi14)
        }
        v3 = (128 >= 42)
        _ = fi14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
