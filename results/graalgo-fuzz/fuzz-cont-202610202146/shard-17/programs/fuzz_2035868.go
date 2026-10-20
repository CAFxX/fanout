package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 int8; F1 int32; F2 uint8 }

func main() {
    var v1 uint16 = 65535
    var v2 uint16 = 1
    var v3 int16 = 0
    var v4 int32 = 0
    v5 := make(map[string]int)
    v4 = (v4 % (((^(v4))) | 1))
    var v6 int64 = 395
    var v7 uint8 = 86
    var v8 uint64 = 2
    v1 += v2
    if true {
        var v9 bool = false
        _ = v9
    } else {
        var v10 uint64 = (^(v8))
        _ = v10
    }
    var v11 int = 1
    var v12 int16 = -893
    v3 = v12
    switch "foo bar" {
    case "":
        for fi13 := 0; fi13 < 8; fi13++ {
            switch -9223372036854775808 {
            case 9223372036854775807:
                v5[("foo bar" + "")] = (fi13 << ((706) & 63))
            case 42:
                switch (fi13 ^ (-(fi13))) {
                case 239:
                    var v14 string = "foo bar"
                    _ = v14
                case 32768:
                    var v15 int = 4294967295
                    _ = v15
                }
            case 127:
                var v16 int8 = 25
                _ = v16
            default:
                var v17 uint64 = (^(v8))
                _ = v17
            }
            var v18 uint16 = (v2 - v2)
            switch ("\n\t" + (("a" + ("" + (("a" + "hello") + (("a" + "hello") + "a")))) + ("\n\t" + "\n\t"))) {
            case "hello":
                v1 -= (^(v18))
            case "\n\t":
                switch "\n\t" {
                case "\n\t":
                    v19 := GenS1{F0: (v12 + (v3 + v3)), F1: (fi13 - v11)}
                    _ = v19
                case "hello":
                    v20 := []int{fi13, 1, (v11 & (v11 + v11)), 4294967295}
                    _ = v20
                default:
                    if ((fi13 << ((255) & 63)) <= (v11 % (((fi13 | (v11 + (fi13 >> ((0) & 63))))) | 1))) {
                        v21 := v5[("hello" + "foo bar")]
                        v2 -= 42
                        v18 += (v1 % ((v18) | 1))
                        _ = v21
                    } else {
                        v6 = (v6 & (v6 >> ((0) & 63)))
                        fmt.Println(v8, v4, v18)
                        v2 = v1
                    }
                }
            case "a":
                var v22 int16 = v12
                _ = v22
            }
            _ = fi13
            _ = v18
        }
    default:
        v23 := GenS1{F0: v12, F1: (v11 % (((v11 % ((127) | 1))) | 1))}
        _ = v23
    }
    v24 := GenS2{F0: 42, F1: 1000, F2: v7}
    v25 := []int{v11, v11, int(v6)}
    _ = len(v5)
    fmt.Println(v12, v8)
    fmt.Println(v3, v4, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
