package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 uint32 = 2147483648
    v2 := make(map[uint32]uint32)
    var v3 uint = 4294967295
    var v4 float64 = 1e10
    v4 = v4
    v3 = v3
    v3 = uint(v1)
    v5 := GenS1{F0: 605}
    var v6 string = ((("foo bar" + ("" + ("a" + "foo bar"))) + "") + "foo bar")
    v7 := GenS1{F0: 1}
    v1 = v1
    var v8 uint32 = v1
    switch v6 {
    case "foo bar":
        for fi9 := 0; fi9 < 6; fi9++ {
            var v10 int32 = -1
            var v11 uint8 = 2
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    var v12 int8 = 100
    var v13 int = -271
    var v14 int16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
