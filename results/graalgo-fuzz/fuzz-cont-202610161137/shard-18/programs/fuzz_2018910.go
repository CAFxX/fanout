package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 int; F1 uint32; F2 int }
type GenS3 struct { F0 int8; F1 uint }

func main() {
    var v1 int8 = -1
    var v2 uint16 = 0
    var v3 uint8 = 1
    var v4 int8 = 1
    v5 := []string{"foo bar"}
    v6 := make(map[int]int)
    v4 -= (v4 | v1)
    if false {
        v2 = v2
        for fi7 := 0; fi7 < 5; fi7++ {
            var v8 int32 = 0
            _ = fi7
            _ = v8
        }
    } else {
        var v9 int64 = 0
        var v10 int8 = 0
        var v11 uint = 255
        _ = v10
        _ = v11
        _ = v9
    }
    v12 := GenS1{F0: 784, F1: 4294967295}
    var v13 bool = false
    for fi14 := 0; fi14 < 16; fi14++ {
        var v15 int16 = 855
        _ = fi14
        _ = v15
    }
    var v16 int = -462
    switch v16 {
    case 127:
        v17 := GenS2{F0: v16, F1: 42, F2: int(v2)}
        _ = v17
    case 2147483647:
        v5 = append(v5, "hello")
    }
    v13 = false
    var v18 uint16 = (^(v2))
    if v13 {
        for fi19 := 0; fi19 < 15; fi19++ {
            for fi20 := 0; fi20 < 12; fi20++ {
                if false {
                    v16 = -9223372036854775808
                    v5 = append(v5, "a")
                } else {
                    v21 := GenS1{F0: 2147483647, F1: 42}
                    fmt.Println(v4, fi20, v13)
                    _ = v21
                }
                _ = fi20
            }
            v16 = fi19
            _ = fi19
        }
    } else {
        var v22 uint64 = 42
        _ = len(v6)
        var v23 uint8 = v3
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 6; fi24++ {
        var v25 int32 = -1
        if len(v5) > 0 {
            v5[0] = "foo bar"
        }
        v26 := GenS3{F0: 0, F1: 65535}
        _ = fi24
        _ = v25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
