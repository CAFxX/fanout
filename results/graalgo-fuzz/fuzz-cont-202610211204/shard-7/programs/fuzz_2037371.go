package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 bool; F1 uint16 }
func genFn0() uint32 {
    var v4 int32 = 0
    var v5 int32 = (^(v4))
    if false {
        v6 := []int{4294967295, -2147483648}
        _ = v6
    }
    fmt.Println(v4, v4)
    v7 := []int{-1, -489, 0, -346}
    _ = v4
    _ = v5
    _ = v7
    return 42
}

func main() {
    var v1 uint8 = 0
    var v2 uint64 = 1
    var v3 uint64 = 4294967295
    var v8 bool = (false || (-754 >= -190))
    var v9 uint64 = 708
    switch 0 {
    case 44:
        var v10 int32 = -902
        _ = v10
    case -9223372036854775808:
        v11 := GenS1{F0: 0}
        _ = v11
    case 0:
        v12 := GenS1{F0: 42}
        _ = v12
    default:
        var v13 string = "foo bar"
        _ = v13
    }
    var v14 int8 = 0
    v15 := GenS2{F0: 9223372036854775807, F1: 565}
    v16 := GenS2{F0: 1, F1: genFn0()}
    var v17 uint16 = uint16(v9)
    for fi18 := 0; fi18 < 12; fi18++ {
        for fi19 := 0; fi19 < 14; fi19++ {
            var v20 int = (-(fi18))
            _ = fi19
            _ = v20
        }
        switch ("\n\t" + "") {
        case "":
            var v21 int32 = 42
            _ = v21
        }
        if (!v8) {
            v2 = v9
            var v22 uint64 = 0
            var v23 uint = 0
            _ = v22
            _ = v23
        } else {
            fmt.Println(v14)
        }
        _ = fi18
    }
    var v24 int64 = 42
    v9 -= (^(v3))
    v9 += (v9 << ((752) & 63))
    for fi25 := 0; fi25 < 11; fi25++ {
        for fi26 := 0; fi26 < 4; fi26++ {
            v27 := []int{}
            _ = len(v27)
            v28 := GenS1{F0: -2147483648}
            _ = fi26
            _ = v27
            _ = v28
        }
        _ = fi25
    }
    if v8 {
        v29 := []int{}
        _ = v29
    } else {
        v30 := GenS3{F0: (false && (728 == -304)), F1: 2}
        _ = v30
    }
    var v31 uint = 221
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v24
    _ = v3
    _ = v31
    _ = v8
    _ = v9
}
