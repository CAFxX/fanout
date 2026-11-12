package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 int16; F2 int64 }
type GenS3 struct { F0 uint16; F1 uint16; F2 int32 }

func main() {
    var v1 int8 = -113
    var v2 string = ""
    var v3 uint16 = 32767
    v4 := []uint32{787}
    var v5 uint = 18446744073709551615
    var v6 uint = (v5 | v5)
    var v7 float32 = 0.1
    for fi8 := 0; fi8 < 7; fi8++ {
        v6 += (v6 % (((^(v6))) | 1))
        if false {
            switch 822 {
            case 42:
                if len(v4) > 0 {
                    v4[0] = 2147483648
                }
            case 0:
                v9 := GenS1{F0: v3}
                _ = v9
            case -896:
                v10 := GenS1{F0: 255}
                _ = v10
            default:
                v6 -= 4294967295
            }
            var v11 int64 = -1
            _ = v11
        } else {
            v12 := GenS2{F0: "foo bar", F1: 1, F2: 1}
            _ = v12
        }
        _ = fi8
    }
    var v13 string = v2
    var v14 int16 = 0
    v1 = (v1 - -64)
    var v15 int16 = (v14 * (v14 ^ v14))
    v1 = v1
    v16 := GenS3{F0: (v3 | (v3 - v3)), F1: v3, F2: -117}
    v4 = append(v4, 732)
    switch int(v3) {
    case 42:
        var v17 int64 = -643
        _ = v17
    case -9223372036854775808:
        var v18 float64 = 1.5
        _ = v18
    default:
        v3 = (^(v3))
    }
    for fi19 := 0; fi19 < 18; fi19++ {
        v4 = append(v4, 0)
        _ = fi19
    }
    var v20 bool = (v13 > "\n\t")
    var v21 bool = (1 == v3)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
