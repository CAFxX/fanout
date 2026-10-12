package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 float32 = 0.1
    var v2 float64 = float64(((((v1 * float32(0.1)) / (v1 + ((v1 / v1) - (1.5 + (3.14159 / 1e10))))) / float32(float64(((((1e10 * (0.0 + 1.5)) / v1) + (0.0 / (1e10 + ((v1 / -2.25) / v1)))) + 0.0)))) - (v1 / v1)))
    v3 := []string{(("a" + "\n\t") + "\n\t"), "hello"}
    v4 := []string{(("hello" + "foo bar") + "\n\t")}
    var v5 uint32 = 705
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v1 += float32(float64(v1))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v5, v1)
    var v6 int64 = -9223372036854775808
    var v7 uint8 = 143
    v5 = v5
    var v8 uint8 = (^(v7))
    v9 := GenS1{F0: (v5 <= 912)}
    var v10 int32 = 42
    if len(v3) > 0 {
        v3[0] = (("hello" + ("hello" + "a")) + "\n\t")
    }
    v11 := GenS1{F0: (!false)}
    switch -2147483648 {
    case -2147483648:
        fmt.Println(v1, v2, v1)
    case -579:
        v12 := GenS1{F0: ("hello" > ("hello" + "a"))}
        _ = v12
    }
    v10 = -2147483648
    if len(v4) > 0 {
        v4[0] = "hello"
    }
    v13 := GenS1{F0: false}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
