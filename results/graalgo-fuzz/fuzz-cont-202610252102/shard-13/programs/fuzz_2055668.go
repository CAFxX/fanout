package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 int64 }

func main() {
    v1 := make(map[int]uint32)
    var v2 bool = (("hello" + "a") < "hello")
    var v3 int64 = 1
    var v4 uint64 = 630
    v1[9223372036854775807] = 0
    var v5 int32 = 0
    var v6 int64 = 722
    v7 := GenS1{F0: 0}
    v6 = 9223372036854775807
    if true {
        var v8 uint16 = 2
        v2 = false
        _ = v8
    } else {
        v2 = true
    }
    if v2 {
        v4 = (^(v4))
    } else {
        var v9 uint = 0
        _ = v9
    }
    if (v4 == v4) {
        var v10 int = 127
        if (false && true) {
            delete(v1, v10)
        } else {
            v11 := v1[(v10 | (v10 % (((v10 * v10)) | 1)))]
            _ = v11
        }
        _ = v10
    }
    v12 := v1[128]
    switch "\n\t" {
    case "":
        var v13 uint8 = 128
        _ = v13
    case "\n\t":
        for fi14 := 0; fi14 < 14; fi14++ {
            if false {
                var v15 uint64 = 42
                switch fi14 {
                case 128:
                    fmt.Println(fi14, v6, v2)
                }
                _ = v15
            }
            for fi16 := 0; fi16 < 2; fi16++ {
                v5 -= (v5 / ((int32(v12)) | 1))
                _ = fi16
            }
            delete(v1, (^(fi14)))
            _ = fi14
        }
    case "foo bar":
        switch "hello" {
        case "":
            fmt.Println(v12)
        case "foo bar":
            v17 := GenS1{F0: 117}
            _ = v17
        }
    }
    if v2 {
        v18 := []int{-412, 32768, -607, -356}
        var v19 uint32 = v12
        _ = v18
        _ = v19
    }
    v20 := GenS2{F0: 1, F1: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
