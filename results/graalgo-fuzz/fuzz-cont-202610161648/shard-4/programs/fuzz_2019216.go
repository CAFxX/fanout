package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }

func main() {
    var v1 uint = 1
    v2 := []uint32{0, 10, 363}
    v3 := make(map[string]bool)
    var v4 string = ("a" + "a")
    fmt.Println(v1, v1)
    for fi5 := 0; fi5 < 3; fi5++ {
        var v6 int16 = 1000
        _ = fi5
        _ = v6
    }
    fmt.Println(v4, v1)
    var v7 int16 = 727
    v4 = v4
    var v8 uint = v1
    delete(v3, v4)
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        var v9 float32 = float32(0.0)
        v9 = v9
        _ = v9
    }
    fmt.Println(v8, v7, v7)
    switch 42 {
    case 127:
        for fi10 := 0; fi10 < 9; fi10++ {
            var v11 int8 = 127
            _ = fi10
            _ = v11
        }
    case -1:
        var v12 float64 = 0.0
        _ = v12
    case 32768:
        var v13 int8 = -91
        _ = v13
    default:
        if len(v2) > 0 {
            v2[0] = 4294967295
        }
    }
    for fi14 := 0; fi14 < 5; fi14++ {
        if ((false && false) && true) {
            v1 += v8
            for fi15 := 0; fi15 < 5; fi15++ {
                var v16 string = v4
                var v17 uint8 = 2
                switch 267 {
                case 128:
                    v17 += (v17 << ((uint(v17)) & 63))
                    v18 := v3["a"]
                    var v19 int64 = -1
                    _ = v18
                    _ = v19
                case -672:
                    v16 = (("hello" + ("\n\t" + (v4 + v16))) + "a")
                case 9223372036854775807:
                    switch fi14 {
                    case 4294967295:
                        v17 = (v17 ^ (v17 % ((v17) | 1)))
                    case 9223372036854775807:
                        fmt.Println(v17)
                    }
                }
                _ = fi15
                _ = v16
                _ = v17
            }
        } else {
            if (false || false) {
                if false {
                    v1 = v1
                    var v20 float32 = 0.1
                    _ = v20
                } else {
                    v8 = 255
                }
                v8 = v8
                v21 := GenS1{F0: 2147483648, F1: 65}
                _ = v21
            }
        }
        _ = fi14
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
