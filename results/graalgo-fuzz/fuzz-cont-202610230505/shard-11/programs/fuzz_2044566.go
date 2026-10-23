package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 bool }
type GenS2 struct { F0 int32; F1 uint16 }

func main() {
    var v1 int = 595
    var v2 int16 = 1
    var v3 int32 = 0
    v4 := make(map[uint32]bool)
    var v5 uint64 = 0
    var v6 uint8 = 1
    v7 := GenS1{F0: "\n\t", F1: 255, F2: true}
    delete(v4, 209)
    for fi8 := 0; fi8 < 15; fi8++ {
        for fi9 := 0; fi9 < 9; fi9++ {
            v2 = -32768
            _ = fi9
        }
        v10 := []int{(^(v1)), fi8}
        _ = fi8
        _ = v10
    }
    v3 = -2147483648
    for fi11 := 0; fi11 < 6; fi11++ {
        switch ("\n\t" + ((("\n\t" + ("a" + "hello")) + "\n\t") + "hello")) {
        case "a":
            var v12 uint8 = v6
            _ = v12
        case "\n\t":
            v3 = (v3 << ((1) & 63))
        case "hello":
            switch "hello" {
            case "hello":
                v1 = -9223372036854775808
            case "\n\t":
                fmt.Println(v2, v3, v6)
            case "":
                if true {
                    switch -1 {
                    case 32767:
                        v1 = fi11
                    default:
                        v2 -= v2
                    }
                    v6 = v6
                    for fi13 := 0; fi13 < 12; fi13++ {
                        v1 = fi11
                        switch "\n\t" {
                        case "hello":
                            fmt.Println(v1)
                        case "a":
                            v14 := []int{-9223372036854775808, (-(fi13))}
                            _ = v14
                        case "foo bar":
                            v15 := GenS1{F0: "\n\t", F1: v6, F2: ((("" + ("foo bar" + "foo bar")) + "foo bar") != ("hello" + ("a" + (("\n\t" + ("a" + ((("\n\t" + "\n\t") + "\n\t") + ""))) + "\n\t"))))}
                            _ = v15
                        }
                        fmt.Println(v3)
                        _ = fi13
                    }
                } else {
                    fmt.Println(v5, v1, v5)
                    switch "a" {
                    case "foo bar":
                        v3 = v3
                    case "hello":
                        _ = len(v4)
                    default:
                        var v16 string = (("hello" + ("" + "hello")) + "")
                        _ = v16
                    }
                }
            }
        }
        if (false && false) {
            fmt.Println(v2, v5, v5)
            v2 -= -643
            v1 = (fi11 | (v1 * (^(v1))))
        } else {
            var v17 uint32 = 65535
            var v18 int = 1
            _ = v17
            _ = v18
        }
        _ = fi11
    }
    if (65535 < 42) {
        fmt.Println(v5)
        v3 -= v3
    }
    v19 := GenS2{F0: v3, F1: 2}
    for fi20 := 0; fi20 < 9; fi20++ {
        v6 = v6
        _ = fi20
    }
    v3 -= -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
