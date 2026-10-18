package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }

func main() {
    var v1 uint16 = 255
    var v2 string = (("hello" + "hello") + ("hello" + ("\n\t" + "")))
    v3 := make(map[int]bool)
    var v4 uint16 = v1
    var v5 uint = 948
    var v6 uint64 = uint64(v4)
    fmt.Println(v6, v1, v1)
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 float32 = -2.25
        v2 = v2
        v9 := GenS1{F0: v1, F1: -2147483648}
        _ = fi7
        _ = v8
        _ = v9
    }
    fmt.Println(v4)
    v10 := GenS1{F0: v1, F1: 953}
    var v11 uint64 = v6
    for fi12 := 0; fi12 < 3; fi12++ {
        v13 := GenS1{F0: v4, F1: -1}
        v14 := []int{}
        _ = fi12
        _ = v13
        _ = v14
    }
    if (!false) {
        v1 = (^(v1))
        _ = len(v3)
    } else {
        v2 = v2
        var v15 int16 = 1
        v16 := GenS1{F0: (^(v4)), F1: -318}
        _ = v15
        _ = v16
    }
    v2 = v2
    v17 := GenS1{F0: (v4 & (^(v1))), F1: int32(v1)}
    v6 = v11
    var v18 bool = ((42 >= 9223372036854775807) || (42 <= -32768))
    var v19 float64 = float64(1e10)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
