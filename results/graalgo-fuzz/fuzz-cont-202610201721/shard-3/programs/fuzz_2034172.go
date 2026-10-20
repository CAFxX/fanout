package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 string = "foo bar"
    var v2 int = 127
    var v3 uint8 = 138
    var v4 int8 = 0
    var v5 uint64 = 80
    v5 += 935
    v6 := GenS1{F0: v4}
    var v7 int = v2
    v8 := GenS1{F0: 0}
    switch 19 {
    case 2147483647:
        v9 := []int{}
        _ = v9
    }
    v10 := GenS1{F0: v4}
    fmt.Println(v2)
    v1 = (("a" + v1) + (v1 + (v1 + "foo bar")))
    v11 := []int{(v2 * -201), (^(v7)), v7}
    fmt.Println(v7, v4)
    v1 = "a"
    for fi12 := 0; fi12 < 3; fi12++ {
        var v13 int8 = 100
        if false {
            var v14 int32 = 42
            var v15 uint = 435
            _ = v14
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
