package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 uint32; F1 string }

func main() {
    var v1 uint16 = 285
    v2 := []uint32{515, 804}
    v3 := []string{"", "foo bar", "\n\t"}
    var v4 uint8 = 62
    v1 = 510
    fmt.Println(v1)
    switch "" {
    case "":
        var v5 string = ("" + (("a" + "") + ("" + ("a" + (("hello" + "hello") + ("a" + ""))))))
        _ = v5
    default:
        var v6 uint = 18446744073709551615
        _ = v6
    }
    if (-1 > 100) {
        v7 := GenS1{F0: true, F1: -80}
        if (9223372036854775807 > 4294967295) {
            var v8 uint8 = (v4 << ((1) & 63))
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v8
        } else {
            if len(v3) > 0 {
                v3[0] = "a"
            }
            v9 := GenS2{F0: 0, F1: "\n\t"}
            _ = v9
        }
        _ = v7
    } else {
        var v10 uint8 = v4
        fmt.Println(v1, v10, v1)
        _ = v10
    }
    if false {
        var v11 int8 = 0
        _ = v11
    } else {
        v4 += v4
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v12 := GenS2{F0: uint32(v4), F1: (("" + "a") + "")}
    v4 += v4
    if true {
        v1 = 255
        var v13 bool = ((((0.0 * 1.5) - -2.25) < (1.5 * 1e10)) || true)
        var v14 int16 = -32768
        _ = v13
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
