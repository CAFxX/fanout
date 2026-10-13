package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 int64 = 506
    v2 := []string{"hello"}
    var v3 string = ""
    var v4 float32 = float32((3.14159 / 1e10))
    if false {
        var v5 uint = 65535
        v2 = append(v2, (("hello" + (v3 + v3)) + ("hello" + "foo bar")))
        fmt.Println(v1)
        _ = v5
    } else {
        var v6 uint8 = 2
        for fi7 := 0; fi7 < 8; fi7++ {
            v8 := GenS1{F0: -385}
            _ = fi7
            _ = v8
        }
        v3 = v3
        _ = v6
    }
    fmt.Println(v3, v1)
    v9 := GenS1{F0: 32767}
    _ = len(v2)
    v3 = ("a" + "foo bar")
    var v10 uint = 833
    var v11 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
