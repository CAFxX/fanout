package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    v1 := []string{"a", ("a" + "hello"), "a", ("" + ("hello" + (((("foo bar" + "\n\t") + "a") + "") + ("a" + ("a" + ("a" + "a"))))))}
    v2 := []int{2147483647}
    v3 := make(map[string]string)
    v4 := make(map[uint32]int)
    v5 := make(map[int]bool)
    var v6 uint8 = 100
    fmt.Println(v6)
    var v7 bool = false
    v7 = v7
    v8 := GenS1{F0: 65535}
    v7 = v7
    for fi9 := 0; fi9 < 6; fi9++ {
        v7 = (!(42 < 0))
        for fi10 := 0; fi10 < 10; fi10++ {
            v7 = (1 > (^(fi10)))
            _ = fi10
        }
        _ = fi9
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        for fi12 := 0; fi12 < 4; fi12++ {
            var v13 float32 = 0.1
            var v14 int64 = -9223372036854775808
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = fi11
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
