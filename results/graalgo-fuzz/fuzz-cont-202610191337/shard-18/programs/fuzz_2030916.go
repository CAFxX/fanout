package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8; F1 int16 }
func genFn0() int {
    var v6 bool = false
    if (1 < 865) {
        var v7 int = 4294967295
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 uint32 = 363
            _ = fi8
            _ = v9
        }
        fmt.Println(v7)
        _ = v7
    } else {
        var v10 uint16 = 845
        switch ("hello" + ("" + (("foo bar" + "hello") + "\n\t"))) {
        case "":
            if true {
                switch -1 {
                case 128:
                    fmt.Println(v10, v10)
                case -1:
                    var v11 uint32 = 2147483647
                    _ = v11
                case 9223372036854775807:
                    for fi12 := 0; fi12 < 9; fi12++ {
                        v13 := GenS1{F0: 1000}
                        _ = fi12
                        _ = v13
                    }
                default:
                    v6 = v6
                }
                var v14 int64 = 0
                _ = v14
            }
        default:
            v15 := GenS1{F0: 0}
            _ = v15
        }
        _ = v10
    }
    v16 := []int{-9223372036854775808, 381}
    _ = v16
    _ = v6
    return 397
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]uint32)
    var v3 uint8 = 42
    var v4 float64 = 0.0
    var v5 bool = true
    var v17 uint8 = (^(v3))
    var v18 int32 = 551
    v3 = v17
    v19 := GenS2{F0: -128, F1: -284}
    if false {
        if (1.5 > 1.5) {
            var v20 int = 42
            _ = v20
        } else {
            var v21 uint = 42
            _ = v21
        }
    }
    v18 += (v18 & (v18 % ((v18) | 1)))
    v3 = v17
    fmt.Println(v17, v5)
    switch "" {
    case "a":
        fmt.Println(v5, v4)
    case "foo bar":
        v22 := []int{172, 2147483647, genFn0(), -9223372036854775808}
        _ = v22
    default:
        var v23 string = "hello"
        _ = v23
    }
    if (777 > 1) {
        fmt.Println(v18, v18)
    }
    fmt.Println(v17)
    fmt.Println(v18, v18, v17)
    v24 := v1[("hello" + "")]
    var v25 uint = 1
    var v26 string = "a"
    for fi27 := 0; fi27 < 10; fi27++ {
        var v28 float32 = -2.25
        v25 = uint(v3)
        _ = fi27
        _ = v28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
