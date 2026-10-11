package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }

func main() {
    v1 := []uint32{62, 0, 2147483647, 2}
    var v2 int16 = 0
    var v3 uint8 = 255
    v4 := make(map[uint32]bool)
    var v5 uint64 = 18446744073709551615
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 int8 = -128
        _ = fi6
        _ = v7
    }
    v4[94] = true
    v3 -= 255
    v8 := GenS1{F0: v2, F1: 42}
    v5 = (v5 & 860)
    switch 127 {
    case -832:
        v9 := GenS1{F0: v2, F1: -2147483648}
        _ = v9
    case 127:
        v10 := GenS1{F0: v2, F1: 42}
        _ = v10
    }
    v11 := GenS1{F0: v2, F1: 32767}
    var v12 float32 = (1.5 - ((0.0 - 3.14159) - 1e10))
    var v13 string = "\n\t"
    fmt.Println(v5)
    if false {
        switch v13 {
        case "foo bar":
            var v14 uint = 299
            _ = v14
        case "":
            v15 := GenS1{F0: v2, F1: -2147483648}
            _ = v15
        }
        for fi16 := 0; fi16 < 16; fi16++ {
            delete(v4, 65535)
            _ = fi16
        }
        _ = len(v1)
    } else {
        v17 := GenS1{F0: v2, F1: 0}
        switch ("a" + (v13 + v13)) {
        case "":
            v3 -= v3
        default:
            switch 128 {
            case 9223372036854775807:
                if true {
                    v5 -= (v5 + v5)
                    v13 = v13
                    fmt.Println(v12, v3, v12)
                }
            case 42:
                fmt.Println(v3)
            case -2147483648:
                switch 128 {
                case 128:
                    var v18 uint16 = 347
                    _ = v18
                }
            }
        }
        _ = v17
    }
    _ = len(v4)
    var v19 float64 = 1.5
    v4[193] = true
    if len(v1) > 0 {
        v1[0] = 852
    }
    v5 = (v5 % (((v5 ^ 18446744073709551615)) | 1))
    v5 = 0
    var v20 int64 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
