package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 uint, p1 bool) int64 {
    var v4 uint32 = 2147483647
    for fi5 := 0; fi5 < 3; fi5++ {
        var v6 int16 = 1000
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: "hello"}
    _ = p0
    _ = p1
    _ = v4
    _ = v7
    return -9223372036854775808
}

func main() {
    var v1 string = (("a" + "\n\t") + "a")
    v2 := []string{(v1 + v1), "hello", v1, v1}
    var v3 float64 = 3.14159
    fmt.Println(v3, v3)
    v8 := GenS1{F0: v1}
    if len(v2) > 0 {
        v2[0] = v1
    }
    var v9 float32 = 3.14159
    switch (v1 + v1) {
    case "":
        fmt.Println(v9, v3, v1)
    case "a":
        var v10 int64 = 9223372036854775807
        _ = v10
    }
    var v11 bool = false
    for fi12 := 0; fi12 < 7; fi12++ {
        var v13 int64 = 1
        var v14 uint16 = 255
        v9 = 3.14159
        _ = fi12
        _ = v13
        _ = v14
    }
    v1 = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
