package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 string }
func genFn0(p0 uint64, p1 int32, p2 uint8) int8 {
    var v4 int = 1
    if true {
        var v5 int32 = 1
        var v6 uint64 = 2
        _ = v5
        _ = v6
    }
    v7 := []int{v4, (v4 | v4), v4, (v4 << ((1) & 63))}
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v7
    return 100
}

func main() {
    var v1 uint16 = 65535
    var v2 int32 = 529
    var v3 bool = (v1 > (v1 / (((v1 + v1)) | 1)))
    var v8 int = -147
    fmt.Println(v8, v1, v1)
    v9 := GenS1{F0: 956}
    var v10 float32 = -2.25
    var v11 uint64 = 2
    v12 := []int{32767, v8, (v8 << ((257) & 63)), v8}
    if false {
        if (!v3) {
            var v13 float64 = float64(3.14159)
            fmt.Println(v8, v8)
            _ = v13
        } else {
            var v14 uint8 = 2
            _ = v14
        }
        fmt.Println(v1, v8, v2)
    }
    fmt.Println(v2)
    for fi15 := 0; fi15 < 11; fi15++ {
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        _ = fi15
    }
    switch "" {
    case "foo bar":
        v16 := GenS2{F0: uint8(v1), F1: ""}
        _ = v16
    case "":
        var v17 string = ("" + "hello")
        _ = v17
    }
    if (((v11 % (((^(v11))) | 1)) > 0) || ((^(v8)) == (v8 - v8))) {
        fmt.Println(v3)
        v2 = v2
    }
    if false {
        if v3 {
            fmt.Println(v1, v1)
            v18 := GenS1{F0: 2}
            _ = v18
        } else {
            var v19 uint16 = (v1 << ((508) & 63))
            _ = v19
        }
        var v20 bool = (1 < -9223372036854775808)
        _ = v20
    }
    var v21 int64 = 593
    v22 := GenS1{F0: 65535}
    var v23 int64 = (v21 ^ v21)
    fmt.Println(v1)
    switch v8 {
    case 122:
        if ((v1 ^ (^(v1))) > 32767) {
            var v24 int64 = v23
            var v25 uint64 = (^(v11))
            fmt.Println(v25, v25, v25)
            _ = v24
            _ = v25
        }
    case 9223372036854775807:
        if len(v12) > 0 {
            v12[0] = (v8 + v8)
        }
    case -800:
        var v26 string = "foo bar"
        _ = v26
    default:
        var v27 uint64 = v11
        _ = v27
    }
    v23 = v21
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v8
    _ = v9
}
