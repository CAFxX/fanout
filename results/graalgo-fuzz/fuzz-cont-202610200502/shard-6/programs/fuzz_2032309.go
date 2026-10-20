package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 uint8) int16 {
    if false {
        fmt.Println(p0)
        p0 = (p0 & (^(p0)))
    }
    var v5 uint16 = 227
    _ = p0
    _ = v5
    return 1
}

func main() {
    var v1 uint64 = 601
    var v2 string = ("foo bar" + ("" + "a"))
    var v3 int8 = -12
    v4 := []int{-9223372036854775808}
    var v6 float64 = 0.0
    if (1 < -766) {
        fmt.Println(v3, v1, v2)
        var v7 int32 = 1
        _ = v7
    }
    var v8 int64 = -1
    switch "foo bar" {
    case "a":
        v1 = 4294967295
    }
    if true {
        if false {
            fmt.Println(v2)
            v3 -= (v3 ^ (v3 / (((v3 & v3)) | 1)))
            if false {
                v6 += ((0.1 / v6) + v6)
            }
        }
        _ = len(v4)
    }
    fmt.Println(v8, v2)
    v9 := GenS1{F0: true}
    v6 -= v6
    v2 = (v2 + (("" + ((v2 + v2) + (v2 + ((("" + v2) + v2) + "a")))) + ("a" + (v2 + v2))))
    v4 = append(v4, 4294967295)
    fmt.Println(v2, v1, v3)
    v8 -= (-(v8))
    if (false || (false || false)) {
        var v10 int = 32767
        var v11 int = -2147483648
        for fi12 := 0; fi12 < 8; fi12++ {
            var v13 uint = 294
            _ = fi12
            _ = v13
        }
        _ = v10
        _ = v11
    } else {
        var v14 float32 = 1e10
        var v15 uint = 0
        if ((!(!true)) && false) {
            v16 := GenS1{F0: ((v3 | 118) != v3)}
            _ = v16
        }
        _ = v14
        _ = v15
    }
    var v17 uint32 = 2
    var v18 int32 = -2147483648
    if len(v4) > 0 {
        v4[0] = -9223372036854775808
    }
    v1 -= (v1 + 1)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v8
    _ = v9
}
