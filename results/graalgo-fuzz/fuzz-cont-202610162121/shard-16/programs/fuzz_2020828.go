package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
func genFn0() string {
    v4 := GenS1{F0: 877, F1: 0}
    var v5 uint8 = 248
    v5 = v5
    _ = v4
    _ = v5
    return ("hello" + ("" + (("hello" + "\n\t") + ("a" + (("" + "foo bar") + ("a" + ("" + "a")))))))
}

func main() {
    var v1 uint32 = 65535
    var v2 uint16 = 2
    var v3 float32 = (float32(0.0) + 0.0)
    fmt.Println(v1)
    fmt.Println(v1, v3)
    var v6 int8 = -1
    v1 += 0
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 uint = 0
        _ = fi7
        _ = v8
    }
    var v9 int64 = 42
    v10 := []int{128}
    for fi11 := 0; fi11 < 4; fi11++ {
        fmt.Println(v1, v1)
        _ = fi11
    }
    v6 = (v6 * 42)
    v3 -= (v3 * v3)
    var v12 int16 = 1000
    v9 = 1
    var v13 int = 32767
    v6 = (v6 % ((v6) | 1))
    var v14 uint16 = v2
    for fi15 := 0; fi15 < 14; fi15++ {
        fmt.Println(v1)
        _ = fi15
    }
    v3 -= ((v3 * v3) / float32(3.14159))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v6
    _ = v9
}
