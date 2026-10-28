package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 string) int64 {
    v4 := GenS1{F0: 65535}
    fmt.Println(p0)
    p0 = "foo bar"
    _ = p0
    _ = v4
    return 0
}
func genFn1(p0 int8) uint8 {
    p0 += -1
    if false {
        fmt.Println(p0)
        var v5 int64 = 9223372036854775807
        _ = v5
    } else {
        fmt.Println(p0)
    }
    _ = p0
    return 182
}

func main() {
    var v1 uint16 = 42
    v2 := []int{}
    v3 := []string{""}
    if (!((!false) && false)) {
        v1 = (v1 / ((v1) | 1))
        if true {
            var v6 uint32 = 2
            _ = v6
        } else {
            var v7 uint = 65535
            var v8 uint = (v7 % (((v7 ^ (v7 / ((2) | 1)))) | 1))
            fmt.Println(v8, v1)
            _ = v7
            _ = v8
        }
        v1 += v1
    }
    var v9 int32 = int32(v1)
    var v10 float32 = float32(1.5)
    if true {
        switch ("hello" + ("" + "")) {
        case "foo bar":
            for fi11 := 0; fi11 < 19; fi11++ {
                v10 += 1.5
                fmt.Println(v1, v1, v9)
                var v12 float64 = (((-2.25 - -2.25) / 1.5) / 1.0)
                _ = fi11
                _ = v12
            }
        default:
            v13 := GenS1{F0: 38}
            _ = v13
        }
        fmt.Println(v10, v1)
    }
    switch -1 {
    case 20:
        switch 32767 {
        case 42:
            for fi14 := 0; fi14 < 7; fi14++ {
                if (!(!true)) {
                    v10 -= v10
                }
                v15 := GenS1{F0: 2147483648}
                _ = fi14
                _ = v15
            }
        }
    }
    v10 += v10
    fmt.Println(v9)
    var v16 int32 = (v9 & (v9 / (((v9 >> ((0) & 63))) | 1)))
    var v17 uint32 = 4294967295
    var v18 float64 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v9
}
