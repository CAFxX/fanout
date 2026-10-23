package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool; F1 uint8; F2 uint }
type GenS4 struct { F0 int; F1 int; F2 int8 }
type GenS5 struct { F0 int8; F1 uint8; F2 uint64 }

func main() {
    var v1 int16 = 32767
    v2 := make(map[int]string)
    v3 := []uint32{}
    v4 := make(map[int]string)
    if (577 == 42) {
        for fi5 := 0; fi5 < 3; fi5++ {
            var v6 uint8 = 112
            fmt.Println(v6, v1)
            _ = fi5
            _ = v6
        }
        var v7 int16 = (v1 + (-(v1)))
        if (false || false) {
            var v8 int16 = (v1 ^ (v7 << ((42) & 63)))
            if false {
                v9 := GenS1{F0: 1, F1: 255, F2: -1}
                var v10 uint16 = 2
                _ = v10
                _ = v9
            }
            v2[32768] = "hello"
            _ = v8
        }
        _ = v7
    }
    var v11 uint8 = 2
    if (242 < 0) {
        v12 := GenS2{F0: -2147483648}
        if false {
            var v13 uint8 = v11
            _ = len(v3)
            if len(v3) > 0 {
                v3[0] = 65535
            }
            _ = v13
        } else {
            v14 := GenS1{F0: (v11 | v11), F1: v11, F2: 0}
            _ = v14
        }
        v2[-37] = "a"
        _ = v12
    }
    var v15 int16 = 42
    if true {
        var v16 uint8 = v11
        v17 := GenS2{F0: -1}
        _ = v16
        _ = v17
    } else {
        switch (("\n\t" + "hello") + "") {
        case "\n\t":
            var v18 uint32 = 42
            _ = v18
        case "a":
            var v19 uint16 = 0
            _ = v19
        case "foo bar":
            fmt.Println(v15)
        default:
            var v20 uint = 2
            _ = v20
        }
    }
    v21 := GenS3{F0: ((3.14159 > -2.25) || (1.5 < (3.14159 / 1.5))), F1: v11, F2: 42}
    var v22 string = ""
    v23 := GenS3{F0: (false && true), F1: (v11 & (^(v11))), F2: 897}
    v11 = (v11 + v11)
    var v24 uint = 42
    for fi25 := 0; fi25 < 15; fi25++ {
        v15 -= 32767
        _ = fi25
    }
    for fi26 := 0; fi26 < 16; fi26++ {
        v27 := GenS4{F0: fi26, F1: -504, F2: 0}
        switch "a" {
        case "hello":
            v11 -= (^(v11))
        default:
            for fi28 := 0; fi28 < 13; fi28++ {
                v29 := GenS4{F0: fi28, F1: fi26, F2: int8(fi26)}
                v30 := GenS5{F0: -128, F1: v11, F2: 2}
                _ = fi28
                _ = v29
                _ = v30
            }
        }
        var v31 uint32 = 65535
        _ = fi26
        _ = v27
        _ = v31
    }
    var v32 string = ("hello" + v22)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v32
    _ = v4
}
