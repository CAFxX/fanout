package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }

func main() {
    var v1 bool = true
    v2 := []string{}
    var v3 uint64 = 382
    var v4 uint8 = 127
    var v5 float64 = 3.14159
    var v6 int16 = -882
    var v7 float32 = 1.5
    v5 = 1.5
    _ = len(v2)
    var v8 int8 = -1
    for fi9 := 0; fi9 < 2; fi9++ {
        if false {
            if v1 {
                var v10 bool = v1
                var v11 int8 = (v8 - (v8 + v8))
                _ = v10
                _ = v11
            }
        }
        _ = fi9
    }
    v12 := GenS1{F0: v1, F1: 42}
    if true {
        if v1 {
            if len(v2) > 0 {
                v2[0] = ""
            }
        }
    } else {
        v13 := GenS1{F0: v1, F1: 2}
        switch 4294967295 {
        case 2147483647:
            _ = len(v2)
        case 1:
            for fi14 := 0; fi14 < 12; fi14++ {
                var v15 uint16 = 130
                _ = fi14
                _ = v15
            }
        }
        v1 = v1
        _ = v13
    }
    switch 1 {
    case -9223372036854775808:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        for fi16 := 0; fi16 < 17; fi16++ {
            var v17 uint16 = 42
            if v1 {
                var v18 int64 = 0
                _ = v18
            }
            _ = fi16
            _ = v17
        }
    }
    if v1 {
        if ((v8 - 0) > -1) {
            if (!v1) {
                v1 = v1
                var v19 uint32 = 42
                _ = v19
            }
            var v20 uint32 = 2
            v3 = v3
            _ = v20
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
