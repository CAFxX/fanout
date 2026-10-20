package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 uint, p1 uint8, p2 bool) int16 {
    var v7 bool = false
    v8 := GenS1{F0: "foo bar"}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 1000
}

func main() {
    var v1 int32 = 0
    var v2 int32 = (v1 ^ (v1 << ((2) & 63)))
    var v3 uint32 = 42
    v4 := make(map[int]int)
    var v5 string = ("\n\t" + ("\n\t" + "a"))
    v6 := make(map[int]int)
    switch v5 {
    case "foo bar":
        v5 = "foo bar"
    }
    switch v5 {
    case "":
        v2 += -2147483648
    case "a":
        fmt.Println(v1, v3)
    }
    var v9 uint32 = v3
    var v10 int8 = 127
    var v11 uint32 = v9
    if true {
        fmt.Println(v2)
    } else {
        v12 := GenS1{F0: v5}
        switch v5 {
        case "foo bar":
            fmt.Println(v9, v5, v2)
        }
        _ = v12
    }
    v10 = 100
    v13 := v6[-1]
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
