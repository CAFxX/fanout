package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 int8) int8 {
    fmt.Println(p0)
    p0 = (-(p0))
    _ = p0
    return p0
}
func genFn1(p0 int16, p1 string) bool {
    switch "\n\t" {
    case "\n\t":
        for fi4 := 0; fi4 < 4; fi4++ {
            v5 := GenS1{F0: (^(fi4))}
            fmt.Println(p0, p1)
            p0 -= (-(p0))
            _ = fi4
            _ = v5
        }
    default:
        var v6 int32 = 42
        _ = v6
    }
    fmt.Println(p1)
    var v7 float32 = -2.25
    _ = p0
    _ = p1
    _ = v7
    return false
}

func main() {
    var v1 uint32 = 4294967295
    v2 := []string{"", "\n\t"}
    var v3 int32 = 2147483647
    fmt.Println(v3, v1)
    v3 = v3
    v8 := GenS1{F0: -1}
    var v9 string = ""
    fmt.Println(v3)
    v3 = v3
    var v10 int64 = -9223372036854775808
    switch (("\n\t" + (("foo bar" + v9) + v9)) + "hello") {
    case "\n\t":
        v1 -= v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
