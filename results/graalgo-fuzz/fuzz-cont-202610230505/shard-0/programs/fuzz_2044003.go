package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 bool; F1 uint8; F2 bool }

func main() {
    var v1 int64 = 9223372036854775807
    v2 := []string{("\n\t" + "foo bar"), "\n\t", ("" + ("hello" + ("hello" + ("a" + (("foo bar" + (("foo bar" + "") + "a")) + ("\n\t" + ("foo bar" + ""))))))), ("hello" + "\n\t")}
    var v3 int64 = -430
    var v4 uint8 = 25
    var v5 int64 = (v3 | (v3 / ((v3) | 1)))
    v6 := GenS1{F0: 65535, F1: -1}
    v4 = 127
    if true {
        v1 = 9223372036854775807
        var v7 int64 = 9223372036854775807
        if (18446744073709551615 != 539) {
            v8 := GenS1{F0: 255, F1: 128}
            fmt.Println(v1, v4, v3)
            _ = v8
        }
        _ = v7
    }
    var v9 uint16 = 255
    if (((!(0.0 <= 1.5)) && ((false || (!false)) && false)) && false) {
        v3 = v5
    }
    switch "hello" {
    case "":
        var v10 float32 = 0.1
        _ = v10
    case "a":
        var v11 uint16 = v9
        _ = v11
    case "hello":
        v12 := GenS2{F0: 400, F1: -1}
        _ = v12
    }
    var v13 int32 = 428
    switch 128 {
    case 1:
        var v14 float64 = (((-2.25 / ((3.14159 / -2.25) + 0.0)) - (0.0 + float64(1.5))) / 1e10)
        _ = v14
    case 2147483647:
        var v15 uint16 = (^(v9))
        _ = v15
    default:
        _ = len(v2)
    }
    for fi16 := 0; fi16 < 17; fi16++ {
        v5 += v3
        if (v4 < 1) {
            v17 := GenS3{F0: ((v13 * (v13 | v13)) <= -618), F1: 127, F2: (true || true)}
            v13 += (-(v13))
            _ = v17
        }
        _ = fi16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
