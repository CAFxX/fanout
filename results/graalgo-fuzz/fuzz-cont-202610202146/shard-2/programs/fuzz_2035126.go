package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 uint32 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 int }

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 42
    v3 := []int{128, 600, 2147483647, -752}
    v4 := make(map[string]string)
    v5 := make(map[string]uint32)
    v6 := make(map[int]string)
    v2 = v2
    var v7 uint8 = (v2 | (v2 << ((329) & 63)))
    var v8 int16 = 0
    if true {
        for fi9 := 0; fi9 < 11; fi9++ {
            v5["a"] = 2
            v2 = (^(v2))
            _ = fi9
        }
        switch "a" {
        case "\n\t":
            fmt.Println(v7)
        default:
            switch "hello" {
            case "\n\t":
                var v10 float64 = 1e10
                _ = v10
            case "a":
                fmt.Println(v2)
            default:
                _ = len(v4)
            }
        }
        var v11 uint16 = 42
        _ = v11
    }
    var v12 string = (("hello" + "foo bar") + "a")
    if false {
        for fi13 := 0; fi13 < 16; fi13++ {
            v2 = 128
            fmt.Println(v8, v12, v8)
            _ = fi13
        }
        for fi14 := 0; fi14 < 14; fi14++ {
            var v15 int16 = (v8 & (v8 * (v8 & (v8 ^ (v8 + v8)))))
            v12 = (v12 + v12)
            v7 -= (v2 >> ((18446744073709551615) & 63))
            _ = fi14
            _ = v15
        }
    } else {
        v16 := GenS1{F0: 533, F1: 254, F2: 679}
        _ = v16
    }
    var v17 uint64 = 0
    v3 = append(v3, 472)
    switch 128 {
    case 2147483647:
        switch 4294967295 {
        case 127:
            var v18 bool = false
            _ = v18
        case -789:
            v19 := GenS2{F0: 1, F1: 1}
            _ = v19
        }
    case 247:
        fmt.Println(v17)
    case 32767:
        v8 = v8
    }
    switch (v12 + v12) {
    case "\n\t":
        for fi20 := 0; fi20 < 9; fi20++ {
            v2 = (^(v2))
            for fi21 := 0; fi21 < 5; fi21++ {
                if (-315 <= v8) {
                    v12 = v12
                } else {
                    v7 = (v7 | 128)
                    v22 := GenS3{F0: 1}
                    _ = v22
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                delete(v4, v12)
                _ = fi21
            }
            var v23 int64 = 0
            _ = fi20
            _ = v23
        }
    }
    var v24 int32 = 2147483647
    v6[-564] = ("\n\t" + ((("" + v12) + v12) + v12))
    _ = len(v3)
    for fi25 := 0; fi25 < 8; fi25++ {
        if len(v3) > 0 {
            v3[0] = (fi25 >> ((43) & 63))
        }
        v8 -= v8
        _ = fi25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
