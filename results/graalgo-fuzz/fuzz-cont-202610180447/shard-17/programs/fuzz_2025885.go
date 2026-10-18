package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 int32 }
type GenS2 struct { F0 bool; F1 bool }

func main() {
    var v1 string = ("a" + (("hello" + "foo bar") + "a"))
    var v2 uint16 = 32767
    var v3 float32 = -2.25
    v4 := []int{201, 792}
    var v5 uint64 = 1
    if false {
        var v6 int64 = 914
        v7 := GenS1{F0: -1, F1: -2147483648, F2: -15}
        var v8 int32 = 1
        _ = v6
        _ = v7
        _ = v8
    }
    var v9 uint16 = v2
    for fi10 := 0; fi10 < 2; fi10++ {
        for fi11 := 0; fi11 < 3; fi11++ {
            v1 = ("" + ((("a" + "") + "\n\t") + v1))
            v2 -= (v9 / ((2) | 1))
            _ = fi11
        }
        if len(v4) > 0 {
            v4[0] = fi10
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi10
    }
    if true {
        if len(v4) > 0 {
            v4[0] = 838
        }
        var v12 int8 = -125
        _ = v12
    }
    var v13 uint16 = (v9 * (^(v9)))
    if len(v4) > 0 {
        v4[0] = 629
    }
    if true {
        v3 += v3
        v5 -= 42
        if len(v4) > 0 {
            v4[0] = 128
        }
    }
    switch "foo bar" {
    case "":
        if false {
            var v14 uint = 557
            v13 = v13
            v4 = append(v4, 9223372036854775807)
            _ = v14
        }
    }
    v15 := GenS2{F0: true, F1: true}
    if len(v4) > 0 {
        v4[0] = -869
    }
    var v16 uint16 = 2
    fmt.Println(v16)
    var v17 string = ("\n\t" + v1)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
