package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 int16 }
type GenS2 struct { F0 uint16; F1 int8 }
func genFn0() int32 {
    var v4 int8 = 90
    var v5 uint = 4294967295
    _ = v4
    _ = v5
    return 42
}

func main() {
    var v1 int8 = 96
    v2 := []int{1, 131, -2147483648, 242}
    var v3 int32 = 697
    var v6 uint16 = 2
    if (false && (!true)) {
        v6 -= v6
        switch "foo bar" {
        case "a":
            v1 = (v1 ^ v1)
        case "":
            switch -2147483648 {
            case 32768:
                v6 = (^(v6))
            }
        default:
            v1 += int8(v3)
        }
    }
    v3 += (v3 & 1)
    switch "foo bar" {
    case "":
        if true {
            if true {
                v6 = v6
                fmt.Println(v6, v6, v3)
                v3 += 0
            }
        } else {
            v2 = append(v2, 258)
            v3 += (v3 + v3)
        }
    default:
        v3 += (v3 % ((v3) | 1))
    }
    var v7 string = ("" + "foo bar")
    _ = len(v2)
    if true {
        v6 += 988
        v6 += (v6 | v6)
    } else {
        v3 = (v3 ^ -205)
        v1 = v1
        v1 = (v1 ^ int8(v6))
    }
    var v8 uint32 = 1
    var v9 uint16 = 65535
    switch v7 {
    case "":
        switch int(v6) {
        case 42:
            v10 := GenS1{F0: 0, F1: (("a" + "a") + "foo bar"), F2: 163}
            _ = v10
        default:
            if true {
                if false {
                    v3 += (v3 << ((0) & 63))
                    switch 409 {
                    case -2147483648:
                        v11 := GenS2{F0: v9, F1: v1}
                        _ = v11
                    }
                    var v12 uint8 = 50
                    _ = v12
                }
                fmt.Println(v3)
            } else {
                fmt.Println(v3, v8, v9)
                v13 := GenS1{F0: 128, F1: (v7 + v7), F2: 1000}
                fmt.Println(v7, v3)
                _ = v13
            }
        }
    default:
        v3 = 0
    }
    v2 = append(v2, 834)
    fmt.Println(v1)
    fmt.Println(v3, v9)
    v9 = (^(v9))
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
