package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 bool }
type GenS2 struct { F0 uint }
func genFn0() string {
    var v4 string = ("hello" + "\n\t")
    fmt.Println(v4)
    _ = v4
    return (("hello" + "hello") + v4)
}

func main() {
    var v1 int64 = -1
    var v2 int = 42
    v3 := []int{-9223372036854775808, -225, (v2 / ((128) | 1))}
    v5 := GenS1{F0: 2, F1: 1, F2: false}
    for fi6 := 0; fi6 < 12; fi6++ {
        v7 := GenS2{F0: 42}
        var v8 int8 = 100
        if len(v3) > 0 {
            v3[0] = -313
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    switch -1 {
    case 169:
        var v9 string = ""
        _ = v9
    case 32768:
        switch -1 {
        case -9223372036854775808:
            v2 = v2
        case -580:
            if false {
                v2 = v2
                v10 := GenS1{F0: 0, F1: 0, F2: true}
                _ = v10
            }
        case 2147483647:
            var v11 int = v2
            _ = v11
        }
    case 0:
        v1 = 141
    }
    fmt.Println(v2, v1)
    var v12 uint16 = 231
    for fi13 := 0; fi13 < 17; fi13++ {
        if len(v3) > 0 {
            v3[0] = (v2 * (-(v2)))
        }
        v12 -= 255
        _ = fi13
    }
    _ = len(v3)
    var v14 float32 = 1.5
    for fi15 := 0; fi15 < 14; fi15++ {
        var v16 uint = 65535
        _ = fi15
        _ = v16
    }
    var v17 uint16 = 65535
    for fi18 := 0; fi18 < 17; fi18++ {
        var v19 uint = 814
        v19 -= v19
        v1 = v1
        _ = fi18
        _ = v19
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v17 = (v17 % (((v17 | (v17 + v17))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v5
}
