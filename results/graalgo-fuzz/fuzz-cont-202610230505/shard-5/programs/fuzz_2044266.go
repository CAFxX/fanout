package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    v1 := make(map[int]string)
    var v2 bool = (!(!false))
    v3 := []string{("hello" + ("a" + "hello")), "", "foo bar", "\n\t"}
    var v4 uint8 = 87
    var v5 uint = 42
    var v6 float32 = 1.5
    switch "hello" {
    case "a":
        for fi7 := 0; fi7 < 1; fi7++ {
            var v8 uint32 = 42
            v9 := GenS1{F0: 255}
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    if len(v3) > 0 {
        v3[0] = ("foo bar" + "hello")
    }
    v10 := GenS1{F0: 32767}
    for fi11 := 0; fi11 < 8; fi11++ {
        delete(v1, -9223372036854775808)
        _ = fi11
    }
    v5 -= v5
    v5 += v5
    var v12 uint64 = 4294967295
    var v13 int64 = 379
    var v14 int = 74
    v13 = (-(v13))
    var v15 float32 = v6
    var v16 int32 = -2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
