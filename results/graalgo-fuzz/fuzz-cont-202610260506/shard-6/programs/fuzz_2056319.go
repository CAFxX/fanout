package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 uint8 }
type GenS2 struct { F0 int32 }

func main() {
    v1 := []int{42, 32767, 0}
    var v2 uint = 18446744073709551615
    var v3 uint16 = 0
    var v4 int64 = 1
    if true {
        fmt.Println(v3)
        var v5 uint32 = uint32(v2)
        if (!false) {
            if true {
                switch 128 {
                case 128:
                    v2 = 254
                case 927:
                    v1 = append(v1, 2147483647)
                default:
                    v4 -= (v4 + (^(v4)))
                }
                var v6 int64 = v4
                _ = v6
            } else {
                var v7 int8 = 1
                _ = v7
            }
        }
        _ = v5
    }
    v4 -= (v4 >> ((1) & 63))
    fmt.Println(v4)
    var v8 uint8 = 127
    if false {
        if (!true) {
            v9 := GenS1{F0: 0, F1: 100, F2: v8}
            v8 += v8
            v10 := GenS1{F0: 99, F1: 109, F2: v8}
            _ = v10
            _ = v9
        } else {
            v2 = (v2 + (v2 - (v2 | (v2 >> ((v2) & 63)))))
            var v11 int8 = -69
            v12 := GenS1{F0: v11, F1: -1, F2: v8}
            _ = v11
            _ = v12
        }
        v4 = (v4 + -1)
        v2 += (^(v2))
    } else {
        v13 := GenS1{F0: 127, F1: 127, F2: 2}
        v14 := GenS1{F0: -128, F1: 0, F2: v8}
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 8; fi15++ {
        v4 = (^(v4))
        v2 = v2
        _ = fi15
    }
    v8 = (v8 + v8)
    v1 = append(v1, 127)
    switch 32767 {
    case -385:
        var v16 int32 = -2147483648
        _ = v16
    case 2147483647:
        if (true && false) {
            var v17 int64 = v4
            v18 := GenS2{F0: 0}
            for fi19 := 0; fi19 < 3; fi19++ {
                var v20 int = (fi19 / (((fi19 >> ((v2) & 63))) | 1))
                var v21 int32 = int32(fi19)
                _ = fi19
                _ = v20
                _ = v21
            }
            _ = v17
            _ = v18
        } else {
            var v22 bool = false
            _ = v22
        }
    }
    var v23 string = (("hello" + "foo bar") + "hello")
    v24 := GenS2{F0: 42}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v8
}
