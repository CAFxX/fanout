package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 int8; F1 int64; F2 uint64 }
type GenS3 struct { F0 int }
func genFn0() int {
    var v5 uint8 = 128
    if (!(1.5 > 1e10)) {
        var v6 uint = 879
        _ = v6
    } else {
        v5 = 25
        v5 += (v5 << ((239) & 63))
    }
    _ = v5
    return 32767
}

func main() {
    var v1 int32 = 2147483647
    var v2 float32 = (0.0 * -2.25)
    var v3 int32 = 1
    v4 := make(map[uint32]bool)
    v7 := GenS1{F0: 0, F1: 2}
    var v8 uint64 = 4294967295
    v2 = ((v2 + v2) / v2)
    v8 = (v8 & 1)
    switch "hello" {
    case "":
        for fi9 := 0; fi9 < 12; fi9++ {
            var v10 float32 = v2
            var v11 float64 = -2.25
            _ = fi9
            _ = v10
            _ = v11
        }
    case "a":
        fmt.Println(v1, v2)
    }
    if (!false) {
        switch "hello" {
        case "foo bar":
            var v12 int64 = -1
            _ = v12
        case "hello":
            v13 := GenS2{F0: 42, F1: 1, F2: (v8 / ((v8) | 1))}
            _ = v13
        case "":
            v1 -= (v1 ^ (^(v3)))
        default:
            v14 := GenS1{F0: -32768, F1: 1}
            _ = v14
        }
        if false {
            v2 += float32(0.1)
            v15 := GenS3{F0: 127}
            var v16 string = ""
            _ = v15
            _ = v16
        }
    } else {
        for fi17 := 0; fi17 < 9; fi17++ {
            var v18 uint64 = (v8 << ((42) & 63))
            switch "foo bar" {
            case "hello":
                var v19 int = 9223372036854775807
                _ = v19
            case "a":
                var v20 bool = true
                _ = v20
            default:
                v3 += -1
            }
            v18 = v18
            _ = fi17
            _ = v18
        }
    }
    switch ((("foo bar" + "foo bar") + (("foo bar" + "") + "hello")) + "hello") {
    case "":
        v2 += 1.5
    }
    var v21 int16 = -107
    delete(v4, 65535)
    v21 += int16(v8)
    fmt.Println(v1)
    var v22 int32 = (v3 >> ((419) & 63))
    var v23 uint = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
