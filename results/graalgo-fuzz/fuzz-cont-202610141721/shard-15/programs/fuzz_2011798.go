package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 int16; F1 int8 }

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[int]string)
    var v3 uint64 = 4294967295
    v4 := make(map[int]string)
    v5 := make(map[string]uint32)
    fmt.Println(v3)
    var v6 string = "a"
    switch "a" {
    case "a":
        var v7 int8 = 0
        _ = v7
    case "foo bar":
        var v8 int = 4294967295
        _ = v8
    case "":
        v9 := GenS1{F0: int8(v3), F1: 42}
        _ = v9
    }
    var v10 int32 = v1
    switch ((("hello" + (v6 + v6)) + v6) + (v6 + v6)) {
    case "hello":
        for fi11 := 0; fi11 < 19; fi11++ {
            var v12 uint16 = 1
            _ = fi11
            _ = v12
        }
    }
    switch "a" {
    case "hello":
        v10 = (v1 | (-(v1)))
    case "":
        for fi13 := 0; fi13 < 10; fi13++ {
            var v14 int16 = -572
            v10 = -97
            v15 := []int{}
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    var v16 int16 = 42
    if ((-641 != 9223372036854775807) && ((86 >= 158) && true)) {
        fmt.Println(v10)
        for fi17 := 0; fi17 < 20; fi17++ {
            for fi18 := 0; fi18 < 7; fi18++ {
                fmt.Println(v3)
                var v19 uint32 = 2147483647
                _ = fi18
                _ = v19
            }
            _ = fi17
        }
    } else {
        for fi20 := 0; fi20 < 12; fi20++ {
            if false {
                v21 := GenS2{F0: 1, F1: 0}
                _ = v21
            } else {
                v16 -= (v16 | -430)
            }
            v22 := []int{-557, fi20, fi20}
            _ = fi20
            _ = v22
        }
    }
    for fi23 := 0; fi23 < 15; fi23++ {
        var v24 bool = false
        v25 := GenS2{F0: v16, F1: -1}
        var v26 uint8 = 215
        _ = fi23
        _ = v24
        _ = v25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
