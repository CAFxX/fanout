package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint16 = 564
    var v3 string = "hello"
    v2 = v2
    v2 = (v2 ^ 1)
    var v4 uint16 = v2
    if true {
        var v5 uint64 = 4294967295
        v4 = v4
        _ = v5
    }
    delete(v1, v3)
    v1[(((v3 + "\n\t") + ((v3 + ((v3 + "\n\t") + (("foo bar" + v3) + "hello"))) + v3)) + v3)] = 2
    v6 := GenS1{F0: (true || (4294967295 == 42))}
    fmt.Println(v4, v3, v3)
    fmt.Println(v2)
    v7 := GenS2{F0: 1}
    fmt.Println(v4, v2)
    var v8 string = v3
    v2 = v2
    for fi9 := 0; fi9 < 10; fi9++ {
        if (42 >= 1000) {
            var v10 float32 = float32(3.14159)
            _ = v10
        }
        _ = fi9
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
