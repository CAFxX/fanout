package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64 }

func main() {
    v1 := make(map[string]string)
    var v2 string = "hello"
    v3 := []int{-879, -1, -1}
    var v4 string = (v2 + v2)
    v5 := make(map[int]int)
    v4 = v4
    var v6 int = -593
    v6 = (v6 * v6)
    switch (((((((v2 + (v2 + (v2 + ((v2 + ("foo bar" + "hello")) + (v2 + "foo bar"))))) + (v2 + "\n\t")) + ((((v4 + "foo bar") + ("foo bar" + (("" + (v2 + ("foo bar" + (v4 + ((((v2 + v2) + ("hello" + v4)) + v4) + "foo bar"))))) + v4))) + "\n\t") + (("foo bar" + v4) + "hello"))) + "") + v2) + (v2 + "")) + "foo bar") {
    case "\n\t":
        v4 = (v4 + v2)
    case "hello":
        var v7 uint = 833
        _ = v7
    case "a":
        v8 := GenS1{F0: 1}
        _ = v8
    default:
        fmt.Println(v6, v2)
    }
    var v9 int = (-(v6))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (4294967295 < 1) {
        if false {
            v9 = v6
        } else {
            var v10 uint = 666
            v11 := GenS2{F0: -9223372036854775808}
            if len(v3) > 0 {
                v3[0] = v6
            }
            _ = v10
            _ = v11
        }
        v12 := v5[v6]
        fmt.Println(v6, v9)
        _ = v12
    } else {
        for fi13 := 0; fi13 < 13; fi13++ {
            fmt.Println(v2)
            _ = fi13
        }
    }
    if false {
        v4 = v4
    }
    _ = len(v3)
    v14 := GenS2{F0: 42}
    var v15 bool = false
    v4 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
