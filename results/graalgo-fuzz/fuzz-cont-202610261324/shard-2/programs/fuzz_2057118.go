package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 uint16 }

func main() {
    v1 := []int{42}
    var v2 int64 = -9223372036854775808
    var v3 uint16 = 32767
    var v4 int8 = 68
    var v5 string = ""
    var v6 string = v5
    var v7 uint32 = 42
    fmt.Println(v7)
    for fi8 := 0; fi8 < 16; fi8++ {
        for fi9 := 0; fi9 < 13; fi9++ {
            v10 := GenS1{F0: (v7 * (v7 - v7)), F1: v3}
            _ = len(v1)
            if true {
                switch (-(fi8)) {
                case 42:
                    var v11 uint64 = 161
                    _ = v11
                }
                v12 := GenS2{F0: (^(v3))}
                var v13 uint32 = uint32(v3)
                _ = v12
                _ = v13
            } else {
                if (!true) {
                    v14 := GenS2{F0: 42}
                    var v15 int8 = int8(v7)
                    _ = v14
                    _ = v15
                } else {
                    var v16 string = "a"
                    var v17 int8 = 127
                    var v18 uint8 = 179
                    _ = v16
                    _ = v17
                    _ = v18
                }
            }
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    var v19 uint = 443
    fmt.Println(v7)
    for fi20 := 0; fi20 < 2; fi20++ {
        v21 := GenS2{F0: 32767}
        fmt.Println(v4, v6)
        _ = fi20
        _ = v21
    }
    v7 = v7
    for fi22 := 0; fi22 < 17; fi22++ {
        v4 = -128
        _ = fi22
    }
    var v23 int32 = 1
    v24 := GenS2{F0: 2}
    _ = len(v1)
    v25 := GenS1{F0: 2147483648, F1: 32767}
    switch 42 {
    case 343:
        for fi26 := 0; fi26 < 3; fi26++ {
            fmt.Println(v2, v5)
            fmt.Println(fi26, v5)
            v27 := GenS2{F0: uint16(v7)}
            _ = fi26
            _ = v27
        }
    }
    for fi28 := 0; fi28 < 17; fi28++ {
        var v29 uint8 = uint8(v19)
        v30 := GenS2{F0: v3}
        _ = fi28
        _ = v29
        _ = v30
    }
    v5 = v5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
