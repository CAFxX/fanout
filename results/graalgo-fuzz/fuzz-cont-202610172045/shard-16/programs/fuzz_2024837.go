package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }

func main() {
    var v1 uint64 = 42
    v2 := make(map[int]string)
    var v3 bool = (("foo bar" >= (("foo bar" + "a") + "a")) && false)
    var v4 uint8 = 127
    var v5 float32 = float32(1e10)
    var v6 float32 = v5
    if ((v4 * v4) >= v4) {
        var v7 int64 = 1
        v1 = v1
        var v8 bool = ((v1 + 2) <= 42)
        _ = v7
        _ = v8
    }
    v9 := []int{871}
    var v10 string = "a"
    v3 = ((576 < 128) && false)
    v11 := GenS1{F0: 162, F1: (!(v3 || (v3 || ((1e10 != 1e10) && true))))}
    if (0.0 > 3.14159) {
        var v12 int16 = -32768
        v3 = true
        _ = v12
    } else {
        v3 = false
    }
    var v13 uint8 = v4
    v2[9223372036854775807] = (v10 + "foo bar")
    _ = len(v9)
    v14 := v2[127]
    if false {
        for fi15 := 0; fi15 < 10; fi15++ {
            delete(v2, (^(fi15)))
            if v3 {
                _ = len(v9)
                fmt.Println(v13, v6)
                fmt.Println(v5, v10, v5)
            }
            if len(v9) > 0 {
                v9[0] = (^(fi15))
            }
            _ = fi15
        }
        v16 := v2[2147483647]
        _ = v16
    } else {
        v17 := v2[972]
        if len(v9) > 0 {
            v9[0] = 4294967295
        }
        _ = v17
    }
    var v18 int32 = 2147483647
    var v19 uint16 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
