package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint8 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int16 = 1000
    var v3 uint64 = 2
    v4 := make(map[string]int)
    fmt.Println(v3)
    v3 -= 847
    v5 := GenS1{F0: "a", F1: 232, F2: 46}
    for fi6 := 0; fi6 < 16; fi6++ {
        var v7 int8 = 0
        v8 := []int{fi6, fi6, fi6, (fi6 ^ 42)}
        switch -996 {
        case -230:
            var v9 int16 = v2
            _ = v9
        case 9223372036854775807:
            v10 := GenS1{F0: "hello", F1: 65535, F2: 34}
            _ = v10
        case 127:
            var v11 int16 = (^(v2))
            _ = v11
        default:
            var v12 int32 = 0
            _ = v12
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    var v13 bool = true
    v13 = true
    if (42 > 42) {
        v14 := []int{int(v3), -4, -740, -9223372036854775808}
        for fi15 := 0; fi15 < 17; fi15++ {
            var v16 float64 = -2.25
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
            var v17 uint32 = 2
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = v14
    }
    if ("" < "a") {
        v18 := []int{}
        v3 -= v3
        var v19 int8 = 1
        _ = v18
        _ = v19
    }
    var v20 bool = true
    delete(v4, "a")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
