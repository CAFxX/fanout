package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 uint16 }
type GenS2 struct { F0 uint32; F1 string; F2 uint16 }
type GenS3 struct { F0 uint32 }
func genFn0() string {
    var v7 uint8 = 255
    v8 := GenS1{F0: 2147483647, F1: (32767 == 1000), F2: 1}
    fmt.Println(v7)
    var v9 int16 = 0
    _ = v7
    _ = v8
    _ = v9
    return "hello"
}

func main() {
    var v1 int64 = 1
    var v2 float32 = (-2.25 * -2.25)
    var v3 uint8 = 1
    var v4 string = (("hello" + "foo bar") + "foo bar")
    v5 := make(map[int]int)
    v6 := make(map[uint32]int)
    var v10 uint32 = 1
    var v11 uint8 = v3
    if ((!(641 > 482)) && ((true && (!((v2 / 3.14159) != (v2 - v2)))) && (!((32767 >= 0) || false)))) {
        fmt.Println(v10, v10, v2)
        v12 := v6[(v10 - 0)]
        _ = v12
    }
    if false {
        var v13 float64 = 0.0
        _ = v13
    } else {
        v14 := GenS2{F0: v10, F1: v4, F2: 65535}
        v4 = v4
        if true {
            if (!false) {
                if false {
                    var v15 uint16 = 65535
                    _ = v15
                } else {
                    var v16 uint64 = 787
                    v16 = 94
                    var v17 float64 = -2.25
                    _ = v16
                    _ = v17
                }
                v18 := GenS3{F0: v10}
                _ = v18
            }
            fmt.Println(v11, v4)
            fmt.Println(v2)
        }
        _ = v14
    }
    if false {
        v19 := GenS2{F0: (^(v10)), F1: v4, F2: uint16(v3)}
        delete(v5, 655)
        var v20 uint16 = 196
        _ = v19
        _ = v20
    } else {
        var v21 uint8 = (v3 ^ 255)
        var v22 int64 = 9223372036854775807
        _ = v21
        _ = v22
    }
    _ = len(v6)
    var v23 int16 = 0
    if true {
        fmt.Println(v11, v1)
        v1 = v1
        v24 := v5[-2147483648]
        _ = v24
    }
    var v25 uint64 = 519
    delete(v6, uint32(v25))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v23
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
