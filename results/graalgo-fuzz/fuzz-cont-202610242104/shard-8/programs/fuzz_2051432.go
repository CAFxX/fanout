package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 int16; F1 uint64; F2 uint }

func main() {
    var v1 int16 = 525
    var v2 bool = (0 > -2147483648)
    var v3 int8 = 127
    v4 := []int{32767, 0, -948, 9223372036854775807}
    v5 := GenS1{F0: 0, F1: "hello"}
    fmt.Println(v1, v3, v2)
    v6 := GenS1{F0: 15, F1: ("" + ("hello" + "foo bar"))}
    switch ("\n\t" + (((("\n\t" + ("" + ("a" + "hello"))) + "foo bar") + "a") + "hello")) {
    case "\n\t":
        v4 = append(v4, 4294967295)
    case "foo bar":
        if false {
            fmt.Println(v2, v2, v3)
            v7 := GenS1{F0: 127, F1: "hello"}
            var v8 int = 127
            _ = v7
            _ = v8
        } else {
            v3 -= (v3 + (v3 ^ 127))
            for fi9 := 0; fi9 < 18; fi9++ {
                v4 = append(v4, fi9)
                for fi10 := 0; fi10 < 20; fi10++ {
                    v11 := GenS1{F0: 2, F1: "hello"}
                    _ = fi10
                    _ = v11
                }
                if len(v4) > 0 {
                    v4[0] = (fi9 >> ((4294967295) & 63))
                }
                _ = fi9
            }
        }
    case "":
        var v12 uint64 = 42
        _ = v12
    default:
        switch -759 {
        case 142:
            switch "" {
            case "foo bar":
                if len(v4) > 0 {
                    v4[0] = -998
                }
            case "":
                if len(v4) > 0 {
                    v4[0] = -692
                }
            default:
                var v13 float32 = float32(1.5)
                _ = v13
            }
        case -2147483648:
            v14 := GenS2{F0: 2147483647}
            _ = v14
        case -1:
            if v2 {
                v15 := GenS1{F0: 1, F1: "foo bar"}
                v2 = (665 <= -2147483648)
                _ = v15
            } else {
                fmt.Println(v3, v2)
            }
        }
    }
    if len(v4) > 0 {
        v4[0] = 391
    }
    v16 := GenS3{F0: 0, F1: 1}
    var v17 int64 = -9223372036854775808
    var v18 uint = 18446744073709551615
    v19 := GenS4{F0: (v1 + -18), F1: 0, F2: v18}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
