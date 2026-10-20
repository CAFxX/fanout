package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 uint16; F2 uint32 }
func genFn0() int {
    v4 := GenS1{F0: (("a" + ("\n\t" + (("\n\t" + "foo bar") + "foo bar"))) + "\n\t")}
    var v5 uint16 = 528
    _ = v4
    _ = v5
    return 42
}

func main() {
    var v1 int16 = 716
    var v2 int16 = (-(v1))
    var v3 string = ("foo bar" + ("\n\t" + ""))
    fmt.Println(v2, v3, v1)
    if true {
        v6 := GenS2{F0: 1}
        var v7 int64 = 679
        fmt.Println(v1, v3)
        _ = v6
        _ = v7
    }
    var v8 uint16 = 1
    if true {
        for fi9 := 0; fi9 < 13; fi9++ {
            var v10 uint = 1
            var v11 int32 = 540
            _ = fi9
            _ = v10
            _ = v11
        }
        v12 := []int{128, -781}
        _ = v12
    }
    var v13 int8 = 1
    v8 = 360
    v14 := GenS3{F0: -2147483648, F1: 0, F2: 74}
    for fi15 := 0; fi15 < 9; fi15++ {
        v1 -= (v1 + 468)
        v13 = v13
        _ = fi15
    }
    for fi16 := 0; fi16 < 14; fi16++ {
        fmt.Println(v13, v2, v8)
        _ = fi16
    }
    v17 := GenS2{F0: 2147483647}
    fmt.Println(v13, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v8
}
