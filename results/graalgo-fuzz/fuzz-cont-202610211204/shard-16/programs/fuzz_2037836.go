package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
func genFn0(p0 int, p1 uint16, p2 int16) uint32 {
    var v4 uint = 232
    var v5 int16 = int16(v4)
    var v6 uint = v4
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 2
}

func main() {
    v1 := []int{-9223372036854775808, 316, 128}
    var v2 int16 = -285
    var v3 string = "a"
    switch v3 {
    case "":
        if (false && false) {
            switch 753 {
            case 128:
                fmt.Println(v3, v3)
            case -169:
                var v7 int16 = (v2 >> ((2) & 63))
                _ = v7
            default:
                fmt.Println(v3, v2)
            }
        }
    case "a":
        var v8 uint32 = genFn0(2147483647, 65535, (v2 / ((-1) | 1)))
        _ = v8
    case "hello":
        switch -2147483648 {
        case 987:
            if (330 > 704) {
                var v9 int8 = 127
                _ = v9
            } else {
                var v10 int8 = 96
                fmt.Println(v10, v2)
                _ = v10
            }
        case -2147483648:
            if false {
                v3 = (v3 + v3)
            }
        default:
            var v11 int64 = 0
            _ = v11
        }
    }
    switch "foo bar" {
    case "":
        if (-1 <= -54) {
            if (2147483647 != 4294967295) {
                for fi12 := 0; fi12 < 15; fi12++ {
                    var v13 uint32 = 4294967295
                    _ = fi12
                    _ = v13
                }
            }
            var v14 int64 = 42
            fmt.Println(v3, v3)
            _ = v14
        }
    default:
        var v15 uint = 2
        _ = v15
    }
    fmt.Println(v2, v2)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v16 int32 = 42
    for fi17 := 0; fi17 < 7; fi17++ {
        if len(v1) > 0 {
            v1[0] = 0
        }
        var v18 float64 = (3.14159 + 0.0)
        _ = fi17
        _ = v18
    }
    v3 = v3
    for fi19 := 0; fi19 < 12; fi19++ {
        var v20 string = (v3 + "foo bar")
        _ = fi19
        _ = v20
    }
    v3 = (((v3 + ("" + (v3 + "hello"))) + "") + (v3 + v3))
    fmt.Println(v16)
    if true {
        var v21 string = v3
        _ = v21
    } else {
        v16 += 0
        v22 := GenS1{F0: -9223372036854775808, F1: 4294967295}
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v22
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
}
