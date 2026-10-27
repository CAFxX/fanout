package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 string; F1 int; F2 uint64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64 }

func main() {
    v1 := []int{-281, -250, 2147483647, -945}
    var v2 bool = (78 < 128)
    var v3 int64 = -9223372036854775808
    v4 := make(map[uint32]uint32)
    v5 := make(map[uint32]bool)
    if (1 == 955) {
        switch "\n\t" {
        case "a":
            switch "hello" {
            case "hello":
                fmt.Println(v3, v3)
            case "foo bar":
                v6 := GenS1{F0: v3, F1: -459}
                _ = v6
            case "a":
                v2 = v2
            default:
                var v7 uint = 0
                _ = v7
            }
        case "foo bar":
            switch ("a" + (("foo bar" + "foo bar") + "")) {
            case "\n\t":
                v2 = false
            case "":
                v2 = (!v2)
            case "a":
                if (!v2) {
                    v8 := GenS2{F0: "", F1: -2147483648, F2: 18446744073709551615}
                    _ = v8
                }
            default:
                v4[2147483648] = 0
            }
        case "":
            v3 += v3
        }
    } else {
        v2 = false
    }
    switch 1 {
    case 1:
        delete(v5, 4294967295)
    }
    if v2 {
        if true {
            for fi9 := 0; fi9 < 20; fi9++ {
                for fi10 := 0; fi10 < 17; fi10++ {
                    var v11 uint = 42
                    var v12 bool = v2
                    v13 := GenS1{F0: (v3 & v3), F1: 0}
                    _ = fi10
                    _ = v11
                    _ = v12
                    _ = v13
                }
                v3 = 9223372036854775807
                v3 += -9223372036854775808
                _ = fi9
            }
            v3 += 1
        }
        v14 := GenS3{F0: 18446744073709551615}
        _ = v14
    }
    v15 := GenS2{F0: (((("" + "") + ("\n\t" + "\n\t")) + "") + "a"), F1: 128, F2: 18446744073709551615}
    var v16 int64 = (^(v3))
    v3 -= (v16 % ((v16) | 1))
    for fi17 := 0; fi17 < 3; fi17++ {
        v18 := GenS4{F0: 84}
        v19 := GenS1{F0: v3, F1: 1}
        for fi20 := 0; fi20 < 7; fi20++ {
            v3 += v16
            _ = fi20
        }
        _ = fi17
        _ = v18
        _ = v19
    }
    delete(v4, 2)
    var v21 int32 = -1
    if v2 {
        var v22 uint = 441
        _ = v22
    }
    _ = len(v5)
    delete(v5, 477)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
