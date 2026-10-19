package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 uint32 }
type GenS2 struct { F0 int16; F1 uint }
func genFn0() uint16 {
    var v4 int8 = 42
    for fi5 := 0; fi5 < 9; fi5++ {
        v4 = 100
        _ = fi5
    }
    _ = v4
    return 32767
}

func main() {
    var v1 uint64 = 554
    var v2 int32 = 1
    v3 := make(map[uint32]string)
    var v6 int16 = 1000
    v2 -= (v2 * (v2 | v2))
    var v7 int64 = -9223372036854775808
    var v8 uint64 = (v1 + v1)
    fmt.Println(v1, v8)
    v8 -= (v1 >> ((255) & 63))
    if true {
        v6 = v6
        v9 := GenS1{F0: v7, F1: (^(v1)), F2: 42}
        _ = v9
    }
    var v10 uint = 2
    var v11 uint = v10
    v7 = 42
    var v12 bool = ((127 < 93) && true)
    for fi13 := 0; fi13 < 9; fi13++ {
        fmt.Println(v2)
        var v14 uint32 = 0
        v10 = 65535
        _ = fi13
        _ = v14
    }
    v15 := GenS2{F0: 32767, F1: 1}
    if v12 {
        var v16 string = (("" + (("hello" + "a") + (("" + "hello") + "a"))) + "a")
        _ = v16
    }
    var v17 uint8 = 1
    for fi18 := 0; fi18 < 5; fi18++ {
        for fi19 := 0; fi19 < 11; fi19++ {
            var v20 float64 = (-2.25 + 0.0)
            if v12 {
                v3[2147483647] = ""
                fmt.Println(v20, v17)
            } else {
                v7 = (v7 & v7)
                var v21 uint16 = 32767
                v11 = v10
                _ = v21
            }
            var v22 uint64 = 0
            _ = fi19
            _ = v20
            _ = v22
        }
        _ = fi18
    }
    for fi23 := 0; fi23 < 6; fi23++ {
        for fi24 := 0; fi24 < 9; fi24++ {
            if v12 {
                var v25 uint = v11
                v12 = v12
                _ = v25
            } else {
                var v26 float64 = (float64(float32(1.5)) + float64(1e10))
                _ = v26
            }
            var v27 string = ""
            _ = fi24
            _ = v27
        }
        var v28 string = "hello"
        var v29 float32 = 3.14159
        _ = fi23
        _ = v28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
