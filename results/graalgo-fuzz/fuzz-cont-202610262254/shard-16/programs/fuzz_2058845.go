package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
func genFn0() string {
    if false {
        for fi5 := 0; fi5 < 18; fi5++ {
            for fi6 := 0; fi6 < 2; fi6++ {
                var v7 uint32 = 153
                _ = fi6
                _ = v7
            }
            _ = fi5
        }
        var v8 float64 = float64(0.0)
        _ = v8
    }
    var v9 int32 = 1
    _ = v9
    return "hello"
}

func main() {
    v1 := make(map[string]string)
    var v2 uint = 1
    var v3 string = "hello"
    var v4 float32 = ((0.0 - 0.1) / -2.25)
    var v10 uint64 = 0
    fmt.Println(v10, v3)
    v4 = v4
    var v11 uint32 = 2
    v3 = (v3 + (v3 + "\n\t"))
    for fi12 := 0; fi12 < 11; fi12++ {
        var v13 uint16 = 2
        v14 := GenS1{F0: -145, F1: 0}
        var v15 int32 = 42
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        var v17 bool = (42 == 669)
        v17 = v17
        if (v2 == 42) {
            v18 := GenS1{F0: 4, F1: -2147483648}
            if true {
                v3 = v3
                v17 = (v17 && true)
            }
            _ = v18
        }
        _ = fi16
        _ = v17
    }
    v11 += (^(v11))
    v19 := GenS1{F0: 0, F1: -15}
    v4 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
