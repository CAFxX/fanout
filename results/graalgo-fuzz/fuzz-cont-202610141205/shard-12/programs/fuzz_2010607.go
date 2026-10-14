package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 uint64 }
func genFn0() uint {
    var v6 int8 = 42
    v6 = (^(v6))
    var v7 int16 = 944
    for fi8 := 0; fi8 < 12; fi8++ {
        if true {
            if false {
                var v9 int = (fi8 - -2147483648)
                var v10 bool = ("" > "a")
                _ = v10
                _ = v9
            }
        }
        _ = fi8
    }
    var v11 int8 = (^(v6))
    _ = v11
    _ = v6
    _ = v7
    return 194
}

func main() {
    var v1 uint32 = 42
    var v2 uint8 = 2
    var v3 uint32 = (v1 / (((^(v1))) | 1))
    var v4 int64 = 42
    var v5 int64 = v4
    for fi12 := 0; fi12 < 1; fi12++ {
        for fi13 := 0; fi13 < 19; fi13++ {
            v2 -= v2
            v14 := []int{-9223372036854775808, fi12}
            v5 += v5
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    var v15 uint = 2
    v3 = 4294967295
    var v16 int8 = 127
    switch "foo bar" {
    case "\n\t":
        var v17 bool = true
        _ = v17
    case "hello":
        v3 += 370
    case "a":
        for fi18 := 0; fi18 < 6; fi18++ {
            var v19 int16 = -423
            v20 := GenS1{F0: fi18, F1: uint32(v2), F2: 1}
            _ = fi18
            _ = v19
            _ = v20
        }
    default:
        if true {
            var v21 int64 = v4
            for fi22 := 0; fi22 < 18; fi22++ {
                v23 := GenS1{F0: fi22, F1: v1, F2: uint64(v2)}
                _ = fi22
                _ = v23
            }
            _ = v21
        } else {
            switch ("\n\t" + "a") {
            case "hello":
                var v24 uint8 = v2
                _ = v24
            case "":
                var v25 int8 = -79
                _ = v25
            }
        }
    }
    v1 = 42
    v26 := []int{394, 42, -713}
    _ = len(v26)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
