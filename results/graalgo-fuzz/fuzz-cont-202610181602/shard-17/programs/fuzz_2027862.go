package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int8 }

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 1
    v3 := make(map[string]int)
    for fi4 := 0; fi4 < 2; fi4++ {
        var v5 uint = 0
        v6 := GenS1{F0: fi4, F1: 255, F2: 42}
        _ = fi4
        _ = v5
        _ = v6
    }
    if false {
        v7 := GenS1{F0: 127, F1: 2, F2: 100}
        var v8 float32 = (1.5 - 0.1)
        _ = v7
        _ = v8
    } else {
        fmt.Println(v2)
        if true {
            v2 = v2
            switch ("\n\t" + "\n\t") {
            case "":
                if (-1 != 32767) {
                    var v9 uint16 = 552
                    _ = v9
                }
            case "a":
                v2 = (^(v2))
            case "foo bar":
                var v10 float32 = (1.5 - (3.14159 * float32(-2.25)))
                _ = v10
            default:
                for fi11 := 0; fi11 < 7; fi11++ {
                    var v12 int64 = -9223372036854775808
                    v13 := []int{}
                    _ = fi11
                    _ = v12
                    _ = v13
                }
            }
        }
        v14 := []int{}
        _ = v14
    }
    var v15 float64 = 1e10
    var v16 int32 = 0
    var v17 bool = (!(4294967295 >= 2147483647))
    var v18 bool = ("" >= (("hello" + (("\n\t" + (("foo bar" + "hello") + "hello")) + ("foo bar" + "a"))) + "foo bar"))
    v19 := GenS1{F0: 1, F1: 32767, F2: -98}
    for fi20 := 0; fi20 < 18; fi20++ {
        var v21 int32 = v16
        v21 -= (v16 ^ 39)
        v3["hello"] = 9223372036854775807
        _ = fi20
        _ = v21
    }
    fmt.Println(v17)
    for fi22 := 0; fi22 < 1; fi22++ {
        var v23 int8 = -128
        v1[(fi22 / ((-2147483648) | 1))] = (fi22 % ((-2147483648) | 1))
        _ = fi22
        _ = v23
    }
    fmt.Println(v16)
    var v24 int16 = 0
    if v18 {
        fmt.Println(v17, v15)
    }
    var v25 int16 = v24
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v3
}
