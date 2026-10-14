package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 int32 }
type GenS2 struct { F0 uint; F1 uint64; F2 bool }

func main() {
    var v1 uint8 = 0
    v2 := []int{42, 128, 1}
    v3 := []string{"hello"}
    v4 := make(map[uint32]bool)
    v5 := []string{("hello" + "\n\t"), ("" + "hello"), "foo bar", ""}
    _ = len(v2)
    var v6 uint8 = (v1 | v1)
    v7 := GenS1{F0: (-538 <= 9223372036854775807), F1: 1, F2: -2147483648}
    v8 := GenS2{F0: 1, F1: 159, F2: (((((-2.25 - 1e10) + float64((3.14159 * 1.5))) - (0.0 - (1.5 + ((1e10 + 0.0) - 3.14159)))) < (float64(-2.25) - (0.0 + 1.5))) || (true && true))}
    delete(v4, 0)
    var v9 uint32 = 1
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi10 := 0; fi10 < 9; fi10++ {
        var v11 uint8 = (^(v1))
        v9 = v9
        _ = fi10
        _ = v11
    }
    var v12 uint32 = 0
    var v13 int64 = -680
    var v14 string = "a"
    v6 = v1
    fmt.Println(v6)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
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
    _ = v9
}
