package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint16; F1 string; F2 uint16 }

func main() {
    var v1 string = (("" + "a") + ("\n\t" + (("" + "a") + "a")))
    v2 := make(map[int]bool)
    v3 := make(map[int]string)
    v4 := make(map[uint32]bool)
    var v5 uint16 = 678
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 uint8 = 2
        v8 := GenS1{F0: 4294967295, F1: -9223372036854775808}
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 uint = 0
    var v10 uint16 = v5
    v11 := []int{32768, 128, -1}
    v5 = v10
    v12 := GenS1{F0: (^(v9)), F1: 4294967295}
    v13 := GenS1{F0: (v9 % ((772) | 1)), F1: 32767}
    var v14 bool = ((!(2147483647 != 2147483647)) && (!false))
    for fi15 := 0; fi15 < 15; fi15++ {
        var v16 uint = 4294967295
        var v17 int64 = 1
        v18 := GenS2{F0: (^(v10)), F1: (v1 + "a"), F2: v10}
        _ = fi15
        _ = v16
        _ = v17
        _ = v18
    }
    if len(v11) > 0 {
        v11[0] = -2147483648
    }
    var v19 int32 = -669
    v20 := GenS1{F0: (^(v9)), F1: int(v10)}
    var v21 uint8 = 128
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
