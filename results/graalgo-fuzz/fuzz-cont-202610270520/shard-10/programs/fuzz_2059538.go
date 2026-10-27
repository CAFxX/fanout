package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 int16 }

func main() {
    var v1 uint32 = 617
    var v2 uint16 = 485
    v3 := make(map[int]int)
    var v4 int = -2147483648
    v5 := []int{v4}
    for fi6 := 0; fi6 < 19; fi6++ {
        for fi7 := 0; fi7 < 8; fi7++ {
            v2 = v2
            _ = fi7
        }
        v8 := GenS1{F0: 1}
        _ = fi6
        _ = v8
    }
    switch 378 {
    case 0:
        v9 := GenS1{F0: -32768}
        _ = v9
    }
    if false {
        var v10 uint32 = 1
        var v11 string = ("\n\t" + "hello")
        _ = v10
        _ = v11
    } else {
        v12 := GenS1{F0: 42}
        fmt.Println(v4, v1)
        v13 := GenS1{F0: 1000}
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: 1000}
    v2 = v2
    if ((202 >= v2) && false) {
        fmt.Println(v1)
        v15 := GenS2{F0: 9223372036854775807, F1: (v1 >> ((253) & 63))}
        _ = v15
    }
    if true {
        var v16 string = "a"
        _ = v16
    } else {
        var v17 int16 = 1
        v18 := GenS3{F0: v17}
        _ = v17
        _ = v18
    }
    fmt.Println(v4, v4, v4)
    var v19 bool = (true || (false && true))
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
