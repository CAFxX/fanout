package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 int32 = 42
    var v2 int64 = 42
    var v3 uint64 = 42
    v4 := make(map[string]uint32)
    var v5 float64 = 1.5
    var v6 uint16 = 32767
    var v7 string = ("foo bar" + "foo bar")
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 uint32 = 402
        var v10 uint = 42
        for fi11 := 0; fi11 < 17; fi11++ {
            var v12 string = ((v7 + "hello") + v7)
            _ = fi11
            _ = v12
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = len(v4)
    switch 127 {
    case -2147483648:
        fmt.Println(v5)
    }
    v13 := GenS1{F0: 128}
    var v14 int32 = v1
    v15 := GenS1{F0: 255}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
