package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }

func main() {
    var v1 int8 = -91
    var v2 int8 = -1
    var v3 float32 = 0.0
    v4 := []int{-874}
    var v5 float64 = (0.0 - (1e10 + float64((v3 - (-2.25 + v3)))))
    if false {
        var v6 float64 = (v5 * 1.5)
        switch 32767 {
        case -1:
            var v7 int = -2147483648
            _ = v7
        case 32768:
            switch "foo bar" {
            case "a":
                v8 := GenS1{F0: 0, F1: 2}
                _ = v8
            case "":
                var v9 float32 = (v3 / 0.1)
                _ = v9
            }
        default:
            v6 = v5
        }
        var v10 uint16 = 2
        _ = v10
        _ = v6
    }
    var v11 int8 = 127
    for fi12 := 0; fi12 < 8; fi12++ {
        switch "a" {
        case "\n\t":
            v2 = v11
        }
        if (0 != 1000) {
            v13 := GenS1{F0: 290, F1: 4294967295}
            _ = v13
        }
        fmt.Println(v1, fi12, v11)
        _ = fi12
    }
    v5 = 1e10
    var v14 int8 = -120
    var v15 int32 = 2147483647
    for fi16 := 0; fi16 < 20; fi16++ {
        var v17 float32 = ((float32((v5 * ((3.14159 - (v5 * v5)) * 3.14159))) - v3) / v3)
        _ = fi16
        _ = v17
    }
    v14 += int8(v15)
    var v18 uint8 = 113
    fmt.Println(v18, v11, v2)
    _ = len(v4)
    var v19 int32 = (v15 / (((v15 & v15)) | 1))
    for fi20 := 0; fi20 < 15; fi20++ {
        var v21 int8 = v14
        var v22 int64 = 173
        var v23 uint8 = (v18 - v18)
        _ = fi20
        _ = v21
        _ = v22
        _ = v23
    }
    var v24 uint8 = v18
    v14 = v1
    v25 := GenS1{F0: 4294967295, F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
