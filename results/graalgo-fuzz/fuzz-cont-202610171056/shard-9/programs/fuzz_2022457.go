package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32; F1 uint16 }

func main() {
    var v1 float64 = -2.25
    v2 := make(map[string]string)
    var v3 int16 = 0
    var v4 bool = (672 >= v3)
    var v5 uint = 1
    v6 := make(map[int]uint32)
    delete(v2, "foo bar")
    if (true && ((v4 && false) || (v4 && v4))) {
        v7 := GenS1{F0: 1}
        switch (("hello" + "a") + "hello") {
        case "hello":
            var v8 uint8 = 128
            _ = v8
        case "foo bar":
            var v9 uint32 = 2147483647
            _ = v9
        case "a":
            fmt.Println(v3, v4, v3)
        }
        fmt.Println(v4)
        _ = v7
    } else {
        v1 -= (0.0 + ((((v1 - float64(float32(v1))) / -2.25) + v1) * ((v1 - v1) / 1e10)))
        v10 := []int{361, 588, -397}
        _ = v10
    }
    _ = len(v2)
    v11 := GenS1{F0: 512}
    for fi12 := 0; fi12 < 10; fi12++ {
        switch 4294967295 {
        case 58:
            for fi13 := 0; fi13 < 12; fi13++ {
                v14 := []int{fi12, (fi12 * 785), fi13}
                var v15 int = (-(fi12))
                _ = fi13
                _ = v14
                _ = v15
            }
        case -9223372036854775808:
            if v4 {
                if (v4 || v4) {
                    v3 += v3
                    v16 := GenS2{F0: -1, F1: 0}
                    _ = v16
                }
            } else {
                var v17 uint = v5
                _ = v17
            }
        case 1:
            var v18 string = "hello"
            _ = v18
        default:
            var v19 int = (fi12 ^ fi12)
            _ = v19
        }
        var v20 int16 = 298
        var v21 bool = false
        _ = fi12
        _ = v20
        _ = v21
    }
    v22 := v2["foo bar"]
    for fi23 := 0; fi23 < 1; fi23++ {
        v24 := []int{fi23}
        _ = fi23
        _ = v24
    }
    var v25 int8 = 100
    switch -1 {
    case 367:
        v4 = false
    case 0:
        v26 := GenS2{F0: -1, F1: 608}
        _ = v26
    case -339:
        var v27 int16 = v3
        _ = v27
    }
    v28 := []int{-194, 127, -773}
    var v29 uint = (v5 % (((v5 + v5)) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v2
    _ = v22
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
