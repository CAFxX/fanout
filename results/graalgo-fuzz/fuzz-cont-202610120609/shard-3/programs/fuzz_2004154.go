package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 int32 = -1
    v2 := []int{}
    var v3 uint16 = 32767
    v4 := make(map[int]int)
    var v5 string = (("a" + ("foo bar" + "hello")) + "foo bar")
    v5 = v5
    if ((!(true && true)) && false) {
        for fi6 := 0; fi6 < 12; fi6++ {
            fmt.Println(v1)
            var v7 float64 = 0.0
            _ = fi6
            _ = v7
        }
        var v8 float64 = 1.5
        _ = v8
    }
    var v9 uint16 = 65535
    var v10 float32 = 0.1
    v11 := GenS1{F0: 2}
    if len(v2) > 0 {
        v2[0] = 32768
    }
    v12 := GenS2{F0: 42}
    var v13 uint64 = 951
    var v14 uint8 = 6
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
