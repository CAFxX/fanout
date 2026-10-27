package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int8; F1 uint; F2 uint }

func main() {
    v1 := []int{-880, 730}
    var v2 int32 = -205
    var v3 int = 55
    var v4 float32 = float32(0.1)
    var v5 float64 = float64((0.0 / 1.0))
    v6 := []uint32{1, 0, 223}
    if false {
        var v7 bool = (334 <= 18446744073709551615)
        _ = v7
    }
    for fi8 := 0; fi8 < 8; fi8++ {
        fmt.Println(v4)
        _ = fi8
    }
    _ = len(v1)
    switch v3 {
    case 42:
        if false {
            var v9 int64 = -405
            v4 = v4
            _ = v9
        } else {
            for fi10 := 0; fi10 < 3; fi10++ {
                switch "foo bar" {
                case "foo bar":
                    v4 += (v4 - float32(v5))
                case "a":
                    _ = len(v1)
                case "\n\t":
                    if true {
                        if true {
                            if (0.0 >= v4) {
                                _ = len(v6)
                                fmt.Println(v3)
                            }
                        }
                        _ = len(v6)
                    } else {
                        if len(v1) > 0 {
                            v1[0] = v3
                        }
                        v5 += (3.14159 - (v5 + v5))
                    }
                default:
                    fmt.Println(v2, fi10)
                }
                _ = fi10
            }
            var v11 uint8 = 255
            _ = v11
        }
    case 127:
        switch "foo bar" {
        case "":
            for fi12 := 0; fi12 < 4; fi12++ {
                if true {
                    fmt.Println(v2, v2, v2)
                }
                for fi13 := 0; fi13 < 10; fi13++ {
                    fmt.Println(v5, v4, fi13)
                    _ = fi13
                }
                v5 -= -2.25
                _ = fi12
            }
        case "a":
            v4 = 0.0
        case "foo bar":
            v3 += (v3 ^ v3)
        }
    case 0:
        var v14 bool = (((((("a" + "a") + "foo bar") + ("hello" + ("hello" + ""))) + "hello") + "a") < (((("foo bar" + (("" + "hello") + "foo bar")) + ("hello" + "foo bar")) + "hello") + "foo bar"))
        _ = v14
    }
    if false {
        for fi15 := 0; fi15 < 12; fi15++ {
            v4 = (v4 / (v4 - v4))
            _ = fi15
        }
        v16 := GenS1{F0: true, F1: 234}
        _ = v16
    }
    v2 -= 42
    v4 += (v4 / v4)
    fmt.Println(v2, v2, v5)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v5 = 0.0
    if false {
        v1 = append(v1, (v3 | v3))
    } else {
        var v17 float64 = (1.5 + v5)
        if (true || (376 >= 65535)) {
            v18 := GenS2{F0: -1, F1: 255, F2: 65535}
            _ = v18
        }
        _ = v17
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
