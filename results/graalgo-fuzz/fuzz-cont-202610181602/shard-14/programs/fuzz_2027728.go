package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 int8 }

func main() {
    var v1 uint32 = 2147483648
    v2 := []int{-1, 42}
    var v3 string = ""
    var v4 bool = (false || false)
    fmt.Println(v3, v1, v3)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if (255 > 260) {
        var v5 int16 = int16(v1)
        if v4 {
            v4 = true
        }
        v3 = v3
        _ = v5
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    fmt.Println(v1, v1, v4)
    v6 := GenS1{F0: 42, F1: 32767, F2: 0}
    v7 := GenS1{F0: 740, F1: 128, F2: 1}
    switch 2147483647 {
    case 128:
        _ = len(v2)
    }
    var v8 uint = 0
    for fi9 := 0; fi9 < 14; fi9++ {
        var v10 float64 = 0.1
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        v12 := GenS1{F0: v8, F1: (fi11 % ((fi11) | 1)), F2: 1}
        _ = fi11
        _ = v12
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
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
