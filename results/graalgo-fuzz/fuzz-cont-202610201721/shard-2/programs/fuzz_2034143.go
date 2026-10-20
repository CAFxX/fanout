package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 uint8, p1 uint64, p2 int) int64 {
    v4 := GenS1{F0: 42}
    v5 := GenS1{F0: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 0
}

func main() {
    var v1 bool = true
    var v2 uint64 = 18446744073709551615
    var v3 uint8 = 2
    v6 := GenS1{F0: 125}
    v7 := []int{32768, 0}
    fmt.Println(v3, v2, v1)
    var v8 string = ("foo bar" + "hello")
    var v9 string = "foo bar"
    if len(v7) > 0 {
        v7[0] = 32768
    }
    switch 78 {
    case 4294967295:
        fmt.Println(v2, v9)
    }
    _ = len(v7)
    v7 = append(v7, -576)
    if false {
        v2 += 0
        for fi10 := 0; fi10 < 3; fi10++ {
            for fi11 := 0; fi11 < 16; fi11++ {
                v12 := GenS2{F0: v3}
                _ = fi11
                _ = v12
            }
            v8 = "hello"
            v13 := GenS1{F0: 2}
            _ = fi10
            _ = v13
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
