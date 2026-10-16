package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }

func main() {
    var v1 float32 = 1.5
    var v2 int8 = 0
    v3 := make(map[uint32]string)
    var v4 uint32 = 2147483648
    var v5 float32 = -2.25
    fmt.Println(v1)
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 int8 = v2
        _ = fi6
        _ = v7
    }
    if true {
        v1 += v5
        var v8 string = "hello"
        _ = v8
    } else {
        if (!false) {
            v9 := GenS1{F0: (-(v2)), F1: v2}
            fmt.Println(v4, v5, v2)
            _ = v9
        }
        var v10 uint8 = 73
        _ = v10
    }
    fmt.Println(v1)
    if false {
        if true {
            var v11 string = ((("foo bar" + "foo bar") + "a") + "a")
            _ = v11
        }
        v1 = v5
    } else {
        v12 := []int{-9223372036854775808}
        _ = v12
    }
    switch "\n\t" {
    case "foo bar":
        var v13 int = 32767
        _ = v13
    case "a":
        v14 := GenS1{F0: (v2 / ((v2) | 1)), F1: 127}
        _ = v14
    }
    if false {
        v2 -= (v2 | v2)
        v15 := GenS1{F0: (v2 & v2), F1: (^(v2))}
        _ = v15
    } else {
        v5 = (v1 / 3.14159)
        v16 := GenS1{F0: (^(v2)), F1: v2}
        delete(v3, v4)
        _ = v16
    }
    var v17 int64 = -1
    var v18 uint8 = 42
    v1 += float32(-2.25)
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
