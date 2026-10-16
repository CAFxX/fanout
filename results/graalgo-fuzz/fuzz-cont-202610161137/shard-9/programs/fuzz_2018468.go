package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 uint16 }
type GenS2 struct { F0 int8; F1 uint32; F2 uint8 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint32 = 2147483648
    var v3 float32 = 1.5
    v4 := []uint32{v2}
    var v5 int = -934
    v6 := GenS1{F0: 2147483647, F1: 1, F2: 0}
    var v7 bool = ((v2 - (v2 ^ (v2 * v2))) >= 2147483648)
    var v8 uint64 = 2
    var v9 uint64 = v8
    switch v5 {
    case -2147483648:
        if (255 == 100) {
            fmt.Println(v3)
            v9 = (v8 * (v8 + 201))
        } else {
            var v10 int = v5
            v3 -= float32(-2.25)
            _ = v10
        }
    case 2147483647:
        v11 := GenS1{F0: -945, F1: 4294967295, F2: 65535}
        _ = v11
    case -678:
        v8 = 42
    default:
        v7 = ((v7 && true) && (65535 != 255))
    }
    v2 = v2
    if (!(v7 && (true && (v9 < (v8 % ((v9) | 1)))))) {
        var v12 int16 = 32767
        fmt.Println(v7, v3)
        _ = v12
    }
    v13 := GenS2{F0: 0, F1: 674, F2: 220}
    if (v8 >= (^(v9))) {
        if false {
            switch "\n\t" {
            case "":
                v9 = 791
            case "foo bar":
                v3 = float32(3.14159)
            case "\n\t":
                v8 -= 4294967295
            }
        }
    } else {
        var v14 bool = v7
        switch -251 {
        case -40:
            if false {
                if ((v5 | (^(v5))) >= 127) {
                    var v15 float32 = v3
                    v16 := GenS2{F0: int8(v2), F1: (v2 >> ((1) & 63)), F2: 1}
                    _ = v15
                    _ = v16
                }
            } else {
                switch "hello" {
                case "":
                    var v17 int32 = -405
                    _ = v17
                case "foo bar":
                    var v18 string = "hello"
                    _ = v18
                case "\n\t":
                    fmt.Println(v9, v5, v14)
                default:
                    v5 = v5
                }
                var v19 bool = (v14 || v14)
                _ = v19
            }
        }
        _ = v14
    }
    v5 -= (v5 >> ((18446744073709551615) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
