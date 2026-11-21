package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 uint }

func main() {
    v1 := make(map[string]uint32)
    v2 := []uint32{}
    var v3 uint8 = 1
    v3 -= (v3 ^ (v3 & (v3 - (v3 | v3))))
    for fi4 := 0; fi4 < 1; fi4++ {
        v3 -= v3
        v3 = v3
        var v5 int32 = 1
        _ = fi4
        _ = v5
    }
    v6 := GenS1{F0: 9223372036854775807, F1: ((-1 <= -106) || true), F2: 86}
    var v7 int = 539
    v3 -= v3
    var v8 int16 = 42
    v1[("a" + (("foo bar" + "\n\t") + "a"))] = 65535
    v8 = v8
    var v9 int8 = -9
    switch "hello" {
    case "a":
        v10 := GenS1{F0: v7, F1: ((0 > -9223372036854775808) && true), F2: 1}
        _ = v10
    case "\n\t":
        var v11 uint8 = 0
        _ = v11
    }
    if false {
        for fi12 := 0; fi12 < 8; fi12++ {
            switch ("" + (("" + "") + ("foo bar" + "foo bar"))) {
            case "hello":
                switch "" {
                case "\n\t":
                    if (-1 > 48) {
                        switch "foo bar" {
                        case "hello":
                            fmt.Println(v9)
                        }
                    }
                case "a":
                    v13 := v1["\n\t"]
                    _ = v13
                default:
                    var v14 bool = false
                    _ = v14
                }
            default:
                var v15 int16 = v8
                _ = v15
            }
            _ = fi12
        }
    } else {
        var v16 uint8 = v3
        var v17 int8 = v9
        if true {
            var v18 int = -509
            v9 = v17
            _ = v18
        }
        _ = v16
        _ = v17
    }
    for fi19 := 0; fi19 < 9; fi19++ {
        for fi20 := 0; fi20 < 2; fi20++ {
            var v21 int32 = -1
            var v22 int16 = v8
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = len(v1)
        _ = fi19
    }
    for fi23 := 0; fi23 < 8; fi23++ {
        switch fi23 {
        case 127:
            var v24 uint8 = v3
            _ = v24
        }
        _ = fi23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
