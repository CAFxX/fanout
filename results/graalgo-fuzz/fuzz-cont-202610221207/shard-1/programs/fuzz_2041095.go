package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]int)
    v3 := make(map[string]bool)
    v2[0] = 470
    var v4 int32 = 513
    var v5 string = "hello"
    for fi6 := 0; fi6 < 11; fi6++ {
        var v7 uint8 = 17
        if false {
            v5 = v5
        } else {
            var v8 string = "a"
            v9 := GenS1{F0: 123, F1: 42}
            if (!(!false)) {
                var v10 uint16 = 719
                var v11 uint64 = 0
                v12 := GenS1{F0: 999, F1: 42}
                _ = v10
                _ = v11
                _ = v12
            }
            _ = v8
            _ = v9
        }
        var v13 float32 = -2.25
        _ = fi6
        _ = v13
        _ = v7
    }
    var v14 float64 = 1.5
    if true {
        for fi15 := 0; fi15 < 15; fi15++ {
            v16 := GenS1{F0: 418, F1: -103}
            _ = fi15
            _ = v16
        }
        v17 := GenS1{F0: 4294967295, F1: 90}
        _ = v17
    } else {
        for fi18 := 0; fi18 < 20; fi18++ {
            var v19 float64 = v14
            v20 := []int{}
            if len(v20) >= 2 {
                v20 = v20[1:len(v20)]
            }
            _ = fi18
            _ = v19
            _ = v20
        }
    }
    v14 = v14
    v5 = "\n\t"
    if (false || (!false)) {
        var v21 int64 = 1
        v22 := []int{32768, -2147483648, 222}
        _ = v21
        _ = v22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
