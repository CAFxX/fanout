package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 uint32; F2 uint8 }
func genFn0() int8 {
    var v4 float64 = -2.25
    var v5 uint64 = 256
    for fi6 := 0; fi6 < 1; fi6++ {
        switch "hello" {
        case "\n\t":
            v7 := []int{}
            _ = v7
        }
        var v8 bool = false
        for fi9 := 0; fi9 < 9; fi9++ {
            var v10 uint64 = v5
            _ = fi9
            _ = v10
        }
        _ = fi6
        _ = v8
    }
    fmt.Println(v4, v5)
    _ = v4
    _ = v5
    return -128
}

func main() {
    var v1 bool = (63 > 34)
    v2 := make(map[string]int)
    v3 := make(map[string]string)
    fmt.Println(v1)
    v11 := GenS1{F0: 1}
    for fi12 := 0; fi12 < 13; fi12++ {
        v1 = v1
        var v13 int64 = 0
        _ = fi12
        _ = v13
    }
    if true {
        v1 = true
        var v14 int = 128
        _ = v14
    }
    if v1 {
        var v15 float64 = 0.1
        for fi16 := 0; fi16 < 17; fi16++ {
            if (!false) {
                fmt.Println(fi16, fi16)
                v1 = v1
            }
            var v17 int32 = -165
            var v18 bool = (42 <= 0)
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = v15
    } else {
        v19 := []int{4294967295}
        _ = v19
    }
    var v20 float64 = ((0.0 / (((1.5 * (1e10 - 0.1)) * 1.5) * 3.14159)) * (float64(1e10) * 1.5))
    var v21 uint16 = 197
    if (!((0.1 - 3.14159) == 1.5)) {
        v22 := GenS2{F0: 65535, F1: 1, F2: 67}
        v23 := GenS1{F0: 1}
        _ = v22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
