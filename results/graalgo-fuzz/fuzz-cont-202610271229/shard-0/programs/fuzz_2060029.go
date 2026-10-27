package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
func genFn0(p0 uint) uint32 {
    switch 174 {
    case -990:
        var v5 bool = (1 <= 1)
        _ = v5
    case 1:
        v6 := GenS1{F0: 65535, F1: p0}
        _ = v6
    case 0:
        fmt.Println(p0)
    }
    fmt.Println(p0)
    var v7 int32 = -1
    v7 += v7
    _ = p0
    _ = v7
    return 0
}

func main() {
    v1 := []int{-1, 32768, 32767}
    v2 := []int{-366, 721, 0, 0}
    var v3 float64 = float64(3.14159)
    v4 := []int{438, 993, -404}
    switch -2147483648 {
    case 42:
        fmt.Println(v3)
    case 0:
        if ((((1.5 + (0.1 + 0.1)) + 0.0) > (0.0 + 0.0)) && ((1.5 <= 1e10) || false)) {
            if true {
                var v8 int64 = 9223372036854775807
                switch ("hello" + ("hello" + "\n\t")) {
                case "":
                    var v9 uint32 = 1
                    _ = v9
                }
                _ = v8
            } else {
                if ((!true) && true) {
                    v3 += (1.5 * v3)
                }
            }
        }
    }
    var v10 uint8 = 255
    for fi11 := 0; fi11 < 13; fi11++ {
        _ = len(v4)
        _ = fi11
    }
    var v12 float64 = 1e10
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v10 -= v10
    var v13 int8 = 1
    v14 := GenS1{F0: uint32(v10), F1: 709}
    switch "hello" {
    case "a":
        if true {
            v10 = 1
            fmt.Println(v3, v13)
        } else {
            if true {
                _ = len(v2)
            } else {
                if true {
                    if false {
                        v10 = 255
                        v12 += 3.14159
                    }
                } else {
                    v15 := GenS1{F0: 2, F1: 65535}
                    _ = v15
                }
                var v16 int16 = 0
                _ = v16
            }
        }
    default:
        v13 = v13
    }
    v10 = (^(v10))
    _ = len(v1)
    var v17 int64 = -1
    var v18 uint = 255
    fmt.Println(v10, v3, v17)
    v18 = (^(v18))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
