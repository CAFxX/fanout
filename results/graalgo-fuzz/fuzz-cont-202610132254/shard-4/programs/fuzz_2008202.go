package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 uint; F2 int }

func main() {
    var v1 bool = (((!(816 > 1000)) || (!true)) && (4294967295 < 629))
    var v2 string = (("\n\t" + "") + "a")
    v3 := make(map[int]bool)
    var v4 float64 = 0.0
    if (65535 >= 214) {
        var v5 float32 = (0.0 + (((float32(0.1) / float32((float64((1e10 + (3.14159 * -2.25))) / (v4 * v4)))) - (((0.0 - (((3.14159 / 1.5) / 0.1) / (-2.25 - (((1e10 + 0.1) - float32(v4)) / 1.5)))) + -2.25) + (1.5 + 0.1))) - float32(v4)))
        var v6 int = 32768
        _ = v5
        _ = v6
    } else {
        v7 := GenS1{F0: 1}
        if v1 {
            var v8 bool = v1
            for fi9 := 0; fi9 < 9; fi9++ {
                fmt.Println(v8, v1)
                fmt.Println(v2, fi9, v2)
                v10 := GenS1{F0: 320}
                _ = fi9
                _ = v10
            }
            _ = v8
        }
        _ = v7
    }
    v2 = "foo bar"
    var v11 string = v2
    var v12 int16 = -325
    if (1 != 255) {
        if (1 != 997) {
            fmt.Println(v12, v4)
            v13 := []int{}
            _ = v13
        }
    }
    delete(v3, -269)
    v2 = "foo bar"
    var v14 int64 = -9223372036854775808
    v15 := GenS2{F0: 42, F1: 363, F2: 159}
    var v16 uint32 = 266
    var v17 string = ""
    _ = len(v3)
    if v1 {
        for fi18 := 0; fi18 < 19; fi18++ {
            fmt.Println(v4, v4, v11)
            _ = fi18
        }
        for fi19 := 0; fi19 < 13; fi19++ {
            v14 = v14
            for fi20 := 0; fi20 < 8; fi20++ {
                var v21 int64 = v14
                var v22 float32 = ((0.1 * 0.1) - 0.1)
                fmt.Println(v14, v16, v11)
                _ = fi20
                _ = v21
                _ = v22
            }
            for fi23 := 0; fi23 < 7; fi23++ {
                var v24 int16 = -32768
                v25 := v3[(fi19 - (^(fi23)))]
                _ = fi23
                _ = v24
                _ = v25
            }
            _ = fi19
        }
    }
    var v26 int8 = -46
    v26 = (^(v26))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v26
    _ = v3
    _ = v4
}
