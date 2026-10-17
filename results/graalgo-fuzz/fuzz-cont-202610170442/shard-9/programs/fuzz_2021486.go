package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint }

func main() {
    var v1 uint16 = 526
    var v2 string = "a"
    v3 := make(map[uint32]uint32)
    var v4 uint = 809
    switch 0 {
    case 1:
        if false {
            v4 += v4
        } else {
            v5 := GenS1{F0: 1, F1: 65535}
            v1 = 32767
            _ = v5
        }
    }
    v4 = v4
    var v6 bool = (false && (3.14159 >= (((1.5 - 0.0) - 1e10) - -2.25)))
    v7 := GenS2{F0: 2147483647, F1: 1}
    if ((v4 % (((v4 & (v4 / (((v4 >> ((v4) & 63))) | 1)))) | 1)) > (v4 << ((v4) & 63))) {
        var v8 uint32 = 42
        if (!(!((((v4 >> ((1) & 63)) == v4) || (!true)) && v6))) {
            var v9 int32 = -665
            v1 = v1
            switch -902 {
            case -1:
                switch "a" {
                case "":
                    fmt.Println(v9, v9)
                default:
                    var v10 uint = v4
                    _ = v10
                }
            case 2147483647:
                v4 = 403
            }
            _ = v9
        }
        _ = v8
    } else {
        for fi11 := 0; fi11 < 11; fi11++ {
            _ = len(v3)
            v12 := v3[444]
            v12 += 65535
            _ = fi11
            _ = v12
        }
        v1 = (v1 | v1)
        switch v2 {
        case "a":
            for fi13 := 0; fi13 < 13; fi13++ {
                v1 = v1
                switch (v2 + v2) {
                case "foo bar":
                    v14 := []int{1, 42, (fi13 >> (((v4 - v4)) & 63))}
                    _ = v14
                case "a":
                    v15 := GenS3{F0: 127}
                    _ = v15
                case "":
                    v2 = ((v2 + "") + "")
                default:
                    fmt.Println(v4)
                }
                _ = fi13
            }
        case "\n\t":
            v16 := GenS4{F0: -1}
            _ = v16
        default:
            v17 := []int{127, 1}
            _ = v17
        }
    }
    v6 = v6
    v18 := GenS3{F0: 43}
    fmt.Println(v4, v6, v2)
    v19 := GenS5{F0: v4}
    var v20 int8 = 127
    fmt.Println(v6, v20)
    var v21 int = 518
    v22 := []int{(-(v21)), (v21 | v21)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
