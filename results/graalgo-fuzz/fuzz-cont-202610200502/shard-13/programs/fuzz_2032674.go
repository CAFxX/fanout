package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 int }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]uint32)
    var v3 uint32 = 575
    var v4 uint32 = v3
    v5 := []uint32{(^(v3))}
    for fi6 := 0; fi6 < 12; fi6++ {
        var v7 uint = 511
        v8 := GenS1{F0: 0, F1: v7}
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 18446744073709551615, F1: uint(v4)}
    for fi10 := 0; fi10 < 5; fi10++ {
        var v11 bool = (!(false && ((0 > -128) && true)))
        v3 -= (v3 % ((65535) | 1))
        _ = fi10
        _ = v11
    }
    v3 -= 0
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    _ = len(v1)
    var v12 uint8 = 127
    delete(v1, (v3 << ((65535) & 63)))
    fmt.Println(v12, v4, v4)
    var v13 float32 = -2.25
    var v14 float64 = -2.25
    if (42 <= 42) {
        if false {
            switch int(v4) {
            case -189:
                var v15 uint32 = (v4 % ((v3) | 1))
                _ = v15
            default:
                switch (("a" + ("" + "hello")) + (((("a" + "\n\t") + "hello") + ("foo bar" + ("hello" + ((((((("" + "hello") + (("" + "a") + "foo bar")) + "hello") + "\n\t") + "hello") + "a") + "hello")))) + ((("a" + ("\n\t" + "foo bar")) + "foo bar") + "foo bar"))) {
                case "foo bar":
                    var v16 uint8 = 0
                    _ = v16
                default:
                    for fi17 := 0; fi17 < 8; fi17++ {
                        fmt.Println(v4, v4)
                        var v18 uint8 = (v12 >> ((0) & 63))
                        var v19 uint16 = 105
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                }
            }
            v14 -= 0.0
        } else {
            v20 := GenS1{F0: 2, F1: 42}
            _ = v20
        }
    }
    var v21 uint64 = 4294967295
    v12 = v12
    v22 := GenS1{F0: 1, F1: uint(v21)}
    var v23 int = 293
    for fi24 := 0; fi24 < 19; fi24++ {
        for fi25 := 0; fi25 < 7; fi25++ {
            fmt.Println(v12, v4)
            v26 := GenS2{F0: 908}
            _ = fi25
            _ = v26
        }
        _ = fi24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
