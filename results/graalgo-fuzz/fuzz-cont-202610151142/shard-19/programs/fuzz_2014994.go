package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 int8; F2 uint64 }
type GenS3 struct { F0 int16; F1 uint8; F2 bool }
type GenS4 struct { F0 int64; F1 bool }

func main() {
    var v1 int = -1
    v2 := make(map[uint32]string)
    v3 := make(map[string]string)
    var v4 int32 = 646
    v5 := make(map[int]string)
    switch "a" {
    case "a":
        var v6 int = v1
        _ = v6
    }
    switch ("" + "hello") {
    case "hello":
        v7 := GenS1{F0: (("" + ("" + (("hello" + ("a" + ((((("" + "a") + "") + "") + ("foo bar" + "foo bar")) + ("\n\t" + "")))) + "\n\t"))) + ("hello" + "hello"))}
        _ = v7
    case "\n\t":
        var v8 int16 = 0
        _ = v8
    default:
        v9 := GenS2{F0: 18446744073709551615, F1: 42, F2: 998}
        _ = v9
    }
    v1 = (v1 - v1)
    switch v1 {
    case 0:
        v10 := GenS3{F0: -32768, F1: 96, F2: false}
        _ = v10
    case 514:
        var v11 float32 = 3.14159
        _ = v11
    case 4294967295:
        if true {
            v12 := GenS4{F0: 9223372036854775807, F1: (false || ((v1 * v1) < (v1 ^ (-(v1)))))}
            fmt.Println(v1)
            var v13 uint32 = 1
            _ = v12
            _ = v13
        } else {
            v4 -= -453
        }
    default:
        switch v1 {
        case -9223372036854775808:
            v14 := []int{(v1 / (((v1 % ((v1) | 1))) | 1))}
            _ = v14
        }
    }
    v4 = (-(v4))
    for fi15 := 0; fi15 < 19; fi15++ {
        var v16 bool = true
        _ = fi15
        _ = v16
    }
    v4 -= (v4 / ((v4) | 1))
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
