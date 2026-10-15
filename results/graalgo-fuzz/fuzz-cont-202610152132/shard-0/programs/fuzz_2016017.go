package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 int16 }
type GenS2 struct { F0 string }

func main() {
    var v1 int8 = 119
    v2 := make(map[uint32]int)
    v3 := make(map[string]int)
    v4 := make(map[int]int)
    v1 += -128
    if true {
        fmt.Println(v1)
        fmt.Println(v1)
    }
    v1 = v1
    var v5 int = 128
    for fi6 := 0; fi6 < 7; fi6++ {
        v7 := GenS1{F0: false, F1: 0, F2: 1}
        fmt.Println(v1)
        _ = fi6
        _ = v7
    }
    if true {
        delete(v3, "\n\t")
        v8 := []int{}
        var v9 int = v5
        _ = v8
        _ = v9
    }
    v5 += 9223372036854775807
    fmt.Println(v1, v5)
    var v10 int = (v5 - (v5 - (v5 ^ 0)))
    var v11 uint64 = 42
    v12 := GenS1{F0: (1 <= 0), F1: -480, F2: 42}
    var v13 int = v10
    if ((v11 >> ((255) & 63)) <= (v11 + (v11 | 42))) {
        v14 := GenS2{F0: "\n\t"}
        for fi15 := 0; fi15 < 6; fi15++ {
            switch (^(v10)) {
            case 1:
                for fi16 := 0; fi16 < 14; fi16++ {
                    var v17 int32 = -1
                    _ = fi16
                    _ = v17
                }
            case 9223372036854775807:
                v18 := GenS2{F0: ("foo bar" + "\n\t")}
                _ = v18
            }
            var v19 int32 = 42
            var v20 uint64 = v11
            _ = fi15
            _ = v19
            _ = v20
        }
        fmt.Println(v11, v11, v13)
        _ = v14
    }
    if false {
        v1 = (v1 | (v1 >> ((621) & 63)))
        var v21 int8 = v1
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
