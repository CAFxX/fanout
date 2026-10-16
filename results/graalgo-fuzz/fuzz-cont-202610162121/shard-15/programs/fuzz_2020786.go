package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 bool }
type GenS2 struct { F0 uint8; F1 int16; F2 int }

func main() {
    var v1 int16 = 0
    v2 := []uint32{326, 1, 921, 0}
    var v3 int32 = 0
    var v4 int16 = (v1 << ((18446744073709551615) & 63))
    v5 := []uint32{896, 1, 855}
    fmt.Println(v3)
    var v6 uint32 = 793
    var v7 string = "a"
    v8 := GenS1{F0: 571, F1: ("foo bar" + v7), F2: (-1 <= 0)}
    if true {
        var v9 int32 = 233
        if false {
            _ = len(v5)
            v4 += 1
            v7 = ("foo bar" + "hello")
        } else {
            _ = len(v5)
            fmt.Println(v7, v4, v3)
        }
        for fi10 := 0; fi10 < 8; fi10++ {
            var v11 string = (("foo bar" + "") + (v7 + v7))
            var v12 uint16 = 255
            for fi13 := 0; fi13 < 9; fi13++ {
                if true {
                    v14 := GenS2{F0: 1, F1: -32768, F2: 0}
                    v4 -= (v1 >> ((84) & 63))
                    var v15 int32 = (v3 % ((-2147483648) | 1))
                    _ = v14
                    _ = v15
                }
                switch 137 {
                case 42:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case -2147483648:
                    var v16 uint64 = 42
                    _ = v16
                }
                var v17 uint16 = v12
                _ = fi13
                _ = v17
            }
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = v9
    }
    var v18 uint64 = 42
    var v19 uint = 491
    v3 += (^(v3))
    var v20 int64 = -9223372036854775808
    v6 = v6
    fmt.Println(v6)
    v4 -= (v4 & (v4 ^ v1))
    v7 = v7
    v6 = 2
    var v21 uint16 = 255
    switch 128 {
    case 960:
        v6 -= v6
    case 128:
        for fi22 := 0; fi22 < 14; fi22++ {
            var v23 float32 = ((1.5 / float32(0.1)) / (float32((3.14159 + 0.1)) + 0.0))
            _ = fi22
            _ = v23
        }
    case -2147483648:
        v24 := GenS2{F0: 255, F1: 0, F2: -852}
        _ = v24
    }
    fmt.Println(v7)
    switch 482 {
    case 672:
        switch 32768 {
        case -2147483648:
            v25 := GenS2{F0: 42, F1: v1, F2: -114}
            _ = v25
        }
    case 128:
        v1 -= (v4 + (v1 % ((v1) | 1)))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
