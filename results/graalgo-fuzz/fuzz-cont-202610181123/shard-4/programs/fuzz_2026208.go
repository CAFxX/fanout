package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint32 }
type GenS2 struct { F0 int; F1 string }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint64 = 0
    v3 := make(map[string]string)
    var v4 uint = 641
    var v5 uint16 = 255
    var v6 int32 = 2147483647
    var v7 uint16 = uint16(v4)
    var v8 int16 = 1
    fmt.Println(v2, v7, v7)
    switch 42 {
    case 9223372036854775807:
        v5 = 42
    case 36:
        switch 160 {
        case -1:
            if (false && ((1.5 / 1e10) < 3.14159)) {
                var v9 int = 32767
                v7 = (v5 * v5)
                _ = v9
            }
        }
    case 32768:
        fmt.Println(v5, v8, v5)
    }
    for fi10 := 0; fi10 < 8; fi10++ {
        fmt.Println(v7, v6, v2)
        _ = fi10
    }
    var v11 uint = (v4 << ((uint(v2)) & 63))
    switch "hello" {
    case "a":
        v2 -= (^(v2))
    case "\n\t":
        if (false || (!false)) {
            v12 := GenS1{F0: (v2 & 2), F1: 100, F2: 2147483648}
            v11 -= (v4 >> (((v11 >> ((v4) & 63))) & 63))
            _ = v12
        } else {
            for fi13 := 0; fi13 < 3; fi13++ {
                switch -2147483648 {
                case 9223372036854775807:
                    v6 = 0
                }
                if true {
                    var v14 uint = (v11 - v11)
                    var v15 uint32 = 2147483647
                    fmt.Println(v7, v6, v14)
                    _ = v14
                    _ = v15
                }
                switch (("a" + ("a" + "hello")) + "hello") {
                case "":
                    if false {
                        v16 := []int{(fi13 << ((v11) & 63))}
                        _ = v16
                    }
                case "hello":
                    var v17 int16 = -997
                    _ = v17
                case "a":
                    fmt.Println(fi13, v5, v7)
                default:
                    var v18 int32 = (v6 << ((uint(v2)) & 63))
                    _ = v18
                }
                _ = fi13
            }
        }
    }
    v19 := GenS2{F0: 127, F1: "foo bar"}
    var v20 int16 = int16(v4)
    v21 := []int{}
    switch ("foo bar" + "a") {
    case "a":
        for fi22 := 0; fi22 < 5; fi22++ {
            fmt.Println(v11)
            if true {
                if len(v21) > 0 {
                    v21[0] = 532
                }
                v11 += (v4 + (v11 << ((42) & 63)))
                var v23 int16 = v20
                _ = v23
            }
            _ = fi22
        }
    default:
        v20 = int16(v2)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
