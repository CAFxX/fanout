package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }

func main() {
    var v1 int8 = 127
    var v2 int64 = -589
    var v3 uint32 = 0
    var v4 int = -355
    v5 := []uint32{}
    v3 = v3
    switch 254 {
    case 825:
        _ = len(v5)
    case 127:
        v5 = append(v5, v3)
    }
    v4 -= 0
    var v6 int = (^(v4))
    var v7 float32 = 1.5
    v1 += 42
    fmt.Println(v6, v2, v1)
    v8 := GenS1{F0: 0, F1: v3}
    var v9 float32 = float32(((3.14159 + float64(1e10)) - 0.0))
    switch (-(v6)) {
    case -9223372036854775808:
        if true {
            v9 = (((v9 / float32(-2.25)) + v7) - 0.1)
            var v10 uint32 = v3
            _ = v10
        } else {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            for fi11 := 0; fi11 < 20; fi11++ {
                var v12 int16 = 1
                var v13 int16 = v12
                fmt.Println(v4, v7, v13)
                _ = fi11
                _ = v12
                _ = v13
            }
            var v14 string = ("a" + "foo bar")
            _ = v14
        }
    case 0:
        v15 := GenS1{F0: 2, F1: v3}
        _ = v15
    case 823:
        var v16 float64 = 0.0
        _ = v16
    default:
        var v17 int8 = v1
        _ = v17
    }
    if (977 > -1) {
        fmt.Println(v2)
        var v18 uint8 = uint8(v3)
        _ = v18
    }
    var v19 string = ""
    if ("foo bar" == "foo bar") {
        switch (v6 << ((547) & 63)) {
        case 1:
            v19 = "hello"
        case -795:
            v20 := GenS1{F0: uint16(v3), F1: v3}
            _ = v20
        case 4294967295:
            v5 = append(v5, 804)
        }
        if false {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
    }
    fmt.Println(v19, v9)
    v9 = v9
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
