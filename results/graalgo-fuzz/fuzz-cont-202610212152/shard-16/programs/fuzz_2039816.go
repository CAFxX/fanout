package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 bool }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]bool)
    v3 := []uint32{424, 244, 853}
    v4 := []uint32{42}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v5 uint16 = 2
    switch "" {
    case "a":
        if ((!(42 >= -50)) && (!true)) {
            v6 := GenS1{F0: 2, F1: -1, F2: (-955 != -536)}
            _ = v6
        } else {
            if (v5 != (v5 >> ((684) & 63))) {
                v7 := GenS1{F0: 42, F1: -1, F2: (!true)}
                _ = len(v1)
                _ = v7
            } else {
                if ((32767 == 32768) || false) {
                    v8 := GenS1{F0: 188, F1: -1, F2: true}
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = v8
                } else {
                    fmt.Println(v5)
                }
                v5 += 42
                v5 -= v5
            }
            delete(v1, -343)
            v5 -= (v5 % (((v5 % (((v5 & (v5 ^ v5))) | 1))) | 1))
        }
    }
    var v9 int16 = -624
    switch ("" + "hello") {
    case "a":
        v9 = v9
    case "foo bar":
        v9 = v9
    default:
        v5 = 463
    }
    var v10 uint16 = v5
    switch ("hello" + "a") {
    case "":
        v10 += (v5 + v10)
    case "\n\t":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case "foo bar":
        v11 := GenS1{F0: 128, F1: -9223372036854775808, F2: (("foo bar" + "") >= ("hello" + "a"))}
        _ = v11
    default:
        var v12 uint32 = 2147483647
        _ = v12
    }
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
