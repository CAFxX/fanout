package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 int32; F2 string }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 string }

func main() {
    v1 := []int{258}
    var v2 uint32 = 604
    var v3 uint8 = 0
    var v4 float64 = 0.1
    v5 := GenS1{F0: 0}
    if (!("a" >= (("" + "foo bar") + "a"))) {
        v6 := GenS2{F0: 2147483647}
        _ = len(v1)
        switch "foo bar" {
        case "foo bar":
            v7 := GenS3{F0: (v2 | v2), F1: -32, F2: "hello"}
            _ = v7
        default:
            switch -2147483648 {
            case 128:
                if len(v1) > 0 {
                    v1[0] = -1
                }
            default:
                var v8 string = ""
                _ = v8
            }
        }
        _ = v6
    } else {
        for fi9 := 0; fi9 < 15; fi9++ {
            if (-1 < 42) {
                v1 = append(v1, (fi9 % ((-9223372036854775808) | 1)))
                if len(v1) > 0 {
                    v1[0] = fi9
                }
                v3 = (v3 - (v3 << ((2) & 63)))
            } else {
                v3 = v3
                v3 = 127
                var v10 int = fi9
                _ = v10
            }
            var v11 uint32 = (v2 | 1)
            _ = fi9
            _ = v11
        }
        v2 = v2
        if (true && true) {
            v12 := GenS2{F0: -566}
            fmt.Println(v4)
            v2 = (v2 >> ((65535) & 63))
            _ = v12
        } else {
            var v13 float64 = (v4 + v4)
            var v14 bool = (true || false)
            _ = v13
            _ = v14
        }
    }
    fmt.Println(v4, v3)
    v15 := GenS2{F0: -1}
    if len(v1) > 0 {
        v1[0] = 4294967295
    }
    var v16 int8 = -128
    for fi17 := 0; fi17 < 8; fi17++ {
        v2 = (^(v2))
        switch fi17 {
        case 9223372036854775807:
            v18 := GenS4{F0: uint32(v3)}
            _ = v18
        case 42:
            var v19 int64 = -9223372036854775808
            _ = v19
        }
        _ = fi17
    }
    switch "hello" {
    case "foo bar":
        if len(v1) > 0 {
            v1[0] = 127
        }
    case "a":
        v20 := GenS5{F0: ("hello" + "a")}
        _ = v20
    case "\n\t":
        v3 = 128
    default:
        v3 -= v3
    }
    fmt.Println(v16, v3)
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
