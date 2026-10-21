package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint32 }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 bool; F1 string; F2 int16 }

func main() {
    v1 := []int{157, 128, 127, 42}
    var v2 int8 = -1
    var v3 int16 = 1
    v4 := make(map[uint32]int)
    var v5 float32 = -2.25
    v6 := []string{"", (("a" + "\n\t") + ("foo bar" + ("foo bar" + ""))), ("a" + "\n\t")}
    v7 := GenS1{F0: v3, F1: int(v2), F2: 65535}
    var v8 uint8 = 216
    switch 9223372036854775807 {
    case 128:
        switch -753 {
        case 1:
            v9 := GenS2{F0: (v2 ^ (^(v2))), F1: false}
            _ = v9
        default:
            v10 := GenS3{F0: 32767, F1: -892}
            _ = v10
        }
    default:
        switch -1 {
        case 9223372036854775807:
            if true {
                var v11 int = 352
                var v12 uint32 = 1
                if true {
                    v11 -= (v11 % (((v11 >> ((255) & 63))) | 1))
                }
                _ = v11
                _ = v12
            } else {
                v3 = v3
                v2 = 116
                fmt.Println(v3)
            }
        case -629:
            v13 := GenS4{F0: (false && (!true)), F1: "a", F2: v3}
            _ = v13
        }
    }
    v3 -= v3
    fmt.Println(v3, v8)
    fmt.Println(v5, v2)
    v6 = append(v6, "hello")
    switch "foo bar" {
    case "":
        var v14 float32 = (3.14159 * ((v5 * v5) + v5))
        _ = v14
    case "hello":
        v8 = v8
    case "a":
        _ = len(v4)
    }
    var v15 uint32 = 65535
    var v16 int32 = -1
    var v17 int = -2147483648
    for fi18 := 0; fi18 < 10; fi18++ {
        var v19 int8 = v2
        for fi20 := 0; fi20 < 6; fi20++ {
            if false {
                var v21 int64 = -9223372036854775808
                var v22 int16 = (v3 & v3)
                _ = len(v4)
                _ = v21
                _ = v22
            } else {
                _ = len(v4)
                fmt.Println(v15, v16)
            }
            _ = fi20
        }
        _ = fi18
        _ = v19
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
