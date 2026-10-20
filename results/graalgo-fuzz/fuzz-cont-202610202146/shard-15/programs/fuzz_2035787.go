package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int8 }

func main() {
    var v1 uint32 = 905
    var v2 int = -2147483648
    var v3 uint8 = 205
    var v4 float32 = -2.25
    v5 := GenS1{F0: "a", F1: 42, F2: -1}
    v6 := []int{}
    var v7 uint32 = v1
    v8 := GenS1{F0: (((("hello" + "a") + ("" + (("foo bar" + "\n\t") + ("hello" + "a")))) + "\n\t") + "\n\t"), F1: 127, F2: int8(v2)}
    v9 := GenS1{F0: "a", F1: -122, F2: 36}
    var v10 float64 = float64(v4)
    fmt.Println(v3, v4, v4)
    for fi11 := 0; fi11 < 2; fi11++ {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = fi11
    }
    v12 := GenS1{F0: "\n\t", F1: 127, F2: int8(v3)}
    var v13 int64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v12
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
