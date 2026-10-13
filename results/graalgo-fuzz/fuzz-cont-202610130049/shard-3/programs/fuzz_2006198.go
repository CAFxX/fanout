package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 uint64; F1 string; F2 int8 }
type GenS3 struct { F0 uint64; F1 string; F2 int16 }

func main() {
    var v1 float64 = (-2.25 + -2.25)
    var v2 int16 = -32768
    var v3 float32 = float32(0.0)
    v4 := make(map[string]string)
    v5 := make(map[uint32]bool)
    v2 += v2
    v2 -= (v2 << ((2) & 63))
    fmt.Println(v2)
    for fi6 := 0; fi6 < 20; fi6++ {
        v7 := GenS1{F0: 128, F1: "a"}
        v3 -= 0.1
        v8 := GenS2{F0: 2, F1: ("hello" + (("hello" + "\n\t") + "foo bar")), F2: -128}
        _ = fi6
        _ = v7
        _ = v8
    }
    if false {
        v2 = v2
        v2 = (v2 | 0)
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        v2 = (v2 / ((v2) | 1))
        _ = fi9
    }
    v10 := GenS2{F0: 2, F1: ("" + "foo bar"), F2: 1}
    v11 := GenS2{F0: 42, F1: "hello", F2: 42}
    var v12 uint32 = 189
    for fi13 := 0; fi13 < 9; fi13++ {
        for fi14 := 0; fi14 < 1; fi14++ {
            var v15 float32 = v3
            fmt.Println(v12, v3)
            v16 := GenS1{F0: 42, F1: ""}
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi13
    }
    if true {
        v17 := GenS1{F0: 0, F1: ("a" + "hello")}
        switch "\n\t" {
        case "a":
            switch 206 {
            case -9223372036854775808:
                v18 := GenS2{F0: 1, F1: "foo bar", F2: 127}
                _ = v18
            default:
                v19 := []int{9223372036854775807, 32768}
                _ = v19
            }
        case "":
            var v20 uint = 1
            _ = v20
        case "foo bar":
            var v21 uint = 961
            _ = v21
        default:
            v12 = 1
        }
        var v22 uint8 = 199
        _ = v17
        _ = v22
    } else {
        switch ("\n\t" + "hello") {
        case "":
            delete(v5, (v12 | (v12 << ((18446744073709551615) & 63))))
        case "foo bar":
            v23 := GenS1{F0: 190, F1: "hello"}
            _ = v23
        }
        switch "" {
        case "a":
            var v24 int64 = -859
            _ = v24
        case "\n\t":
            v25 := GenS2{F0: 131, F1: "foo bar", F2: -1}
            _ = v25
        case "":
            delete(v4, "a")
        default:
            fmt.Println(v2, v3)
        }
        fmt.Println(v3, v3)
    }
    v26 := GenS3{F0: 735, F1: ("hello" + ("\n\t" + "hello")), F2: v2}
    var v27 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
