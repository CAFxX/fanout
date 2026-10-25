package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 int64; F2 int64 }
func genFn0(p0 int64, p1 int16) bool {
    if false {
        fmt.Println(p0, p0)
    }
    switch "hello" {
    case "a":
        v5 := GenS1{F0: 53}
        _ = v5
    case "":
        v6 := []int{0, -819}
        _ = v6
    case "\n\t":
        switch "hello" {
        case "":
            v7 := GenS1{F0: 65535}
            _ = v7
        }
    default:
        p0 -= 0
    }
    v8 := []int{-23}
    _ = p0
    _ = p1
    _ = v8
    return false
}

func main() {
    var v1 int8 = 0
    var v2 float64 = (0.1 * -2.25)
    var v3 float64 = v2
    var v4 uint = 0
    v1 -= v1
    v2 = 0.0
    var v9 int = -9223372036854775808
    var v10 int64 = -1
    v11 := GenS1{F0: v4}
    for fi12 := 0; fi12 < 14; fi12++ {
        for fi13 := 0; fi13 < 20; fi13++ {
            if ((^(v4)) <= 4294967295) {
                v14 := []int{fi12, v9, (fi12 % ((127) | 1))}
                var v15 uint32 = 2147483647
                _ = v14
                _ = v15
            } else {
                if false {
                    switch fi12 {
                    case 42:
                        if false {
                            fmt.Println(v2, v3, v2)
                            fmt.Println(v9, v3)
                        }
                    case 1:
                        fmt.Println(v1)
                    case -9223372036854775808:
                        var v16 float32 = 0.1
                        _ = v16
                    }
                    v17 := []int{-9223372036854775808, fi13, -9223372036854775808, (fi13 + fi12)}
                    _ = v17
                }
            }
            v2 += v3
            _ = fi13
        }
        var v18 uint = v4
        _ = fi12
        _ = v18
    }
    fmt.Println(v10, v10)
    v4 = 383
    v19 := GenS1{F0: v4}
    fmt.Println(v3, v2)
    if true {
        v20 := []int{(-(v9)), -2147483648, 2147483647, v9}
        _ = v20
    } else {
        if ("foo bar" <= "foo bar") {
            var v21 float64 = (0.0 - (0.0 + (v3 - v3)))
            var v22 float32 = -2.25
            _ = v21
            _ = v22
        }
    }
    fmt.Println(v10, v10, v9)
    v23 := GenS1{F0: 237}
    v24 := GenS2{F0: 2, F1: v10, F2: -9223372036854775808}
    for fi25 := 0; fi25 < 18; fi25++ {
        for fi26 := 0; fi26 < 8; fi26++ {
            v4 = 65535
            var v27 int32 = int32(v4)
            v28 := []int{32768}
            _ = fi26
            _ = v27
            _ = v28
        }
        var v29 int32 = 485
        _ = fi25
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v9
}
