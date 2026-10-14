package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64; F1 uint16 }
type GenS5 struct { F0 uint; F1 uint8 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]uint32)
    var v3 int16 = 42
    v4 := make(map[uint32]string)
    var v5 int32 = 42
    var v6 int32 = v5
    switch 128 {
    case 688:
        var v7 int64 = 892
        _ = v7
    case 128:
        if false {
            v8 := GenS1{F0: 479, F1: -1}
            v9 := GenS1{F0: 42, F1: 2147483647}
            for fi10 := 0; fi10 < 6; fi10++ {
                var v11 float32 = float32(0.1)
                v12 := []int{}
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = v8
            _ = v9
        }
    default:
        for fi13 := 0; fi13 < 19; fi13++ {
            v14 := GenS2{F0: (v3 ^ v3), F1: (v3 * (v3 << ((156) & 63)))}
            var v15 uint16 = 1
            v5 -= (v5 & v5)
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    var v16 int64 = -346
    if (18446744073709551615 < 0) {
        var v17 string = "hello"
        _ = v17
    } else {
        switch (("hello" + "a") + ((("a" + ((((("" + "foo bar") + ("\n\t" + "\n\t")) + ("" + ("\n\t" + ("foo bar" + "hello")))) + ("\n\t" + "\n\t")) + "hello")) + "foo bar") + "a")) {
        case "":
            for fi18 := 0; fi18 < 16; fi18++ {
                var v19 string = "hello"
                _ = fi18
                _ = v19
            }
        }
    }
    for fi20 := 0; fi20 < 16; fi20++ {
        v3 -= 0
        _ = fi20
    }
    fmt.Println(v16)
    _ = len(v4)
    switch (("foo bar" + "\n\t") + "foo bar") {
    case "foo bar":
        if (("" <= "hello") || true) {
            var v21 uint16 = 255
            switch 451 {
            case 128:
                switch "\n\t" {
                case "\n\t":
                    fmt.Println(v3, v5)
                case "hello":
                    var v22 uint32 = 25
                    _ = v22
                case "a":
                    if (-91 <= -128) {
                        v16 += 207
                        v21 = 65535
                    } else {
                        v23 := GenS3{F0: 2147483647}
                        _ = v23
                    }
                }
            case -847:
                v21 += (^(v21))
            case 32767:
                v24 := GenS4{F0: 601, F1: (^(v21))}
                _ = v24
            default:
                var v25 float64 = 3.14159
                _ = v25
            }
            _ = v21
        }
    }
    v26 := GenS2{F0: v3, F1: v3}
    delete(v2, 0)
    fmt.Println(v3, v16, v16)
    v27 := GenS5{F0: 4294967295, F1: 46}
    v28 := GenS5{F0: 255, F1: 157}
    v29 := GenS4{F0: 1, F1: 969}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
