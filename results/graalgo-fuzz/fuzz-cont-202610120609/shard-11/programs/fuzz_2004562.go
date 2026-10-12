package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 float64 = 3.14159
    v2 := []uint32{2147483647, 466}
    var v3 int8 = -80
    var v4 bool = false
    var v5 float32 = float32((v1 / 0.1))
    if v4 {
        fmt.Println(v1)
    }
    fmt.Println(v1, v4, v3)
    if len(v2) > 0 {
        v2[0] = 42
    }
    for fi6 := 0; fi6 < 7; fi6++ {
        if v4 {
            var v7 int64 = 0
            _ = v7
        } else {
            var v8 int32 = 0
            v8 = (v8 << ((2) & 63))
            _ = v8
        }
        v9 := GenS1{F0: "foo bar"}
        _ = fi6
        _ = v9
    }
    var v10 uint16 = 0
    fmt.Println(v3, v10, v4)
    var v11 int64 = 42
    _ = len(v2)
    v12 := GenS1{F0: (("hello" + "") + ((("" + "\n\t") + "") + ("foo bar" + "")))}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
