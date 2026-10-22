package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 int16; F2 uint16 }
type GenS3 struct { F0 uint64; F1 uint; F2 string }
func genFn0(p0 int16) int32 {
    var v5 bool = (117 >= 151)
    v6 := []int{1, -9223372036854775808, 127, 434}
    fmt.Println(v5, v5)
    v7 := GenS1{F0: 255}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 2147483647
}

func main() {
    var v1 int32 = 55
    var v2 uint = 682
    var v3 uint16 = 0
    var v4 uint32 = 4294967295
    v4 -= v4
    v1 = (v1 << ((v2) & 63))
    v3 -= (v3 * v3)
    var v8 bool = (!(false || true))
    v1 = v1
    v9 := GenS1{F0: uint16(v2)}
    switch ("foo bar" + ("\n\t" + ("a" + "a"))) {
    case "":
        fmt.Println(v8)
    case "a":
        v10 := GenS2{F0: 42, F1: 42, F2: v3}
        _ = v10
    case "foo bar":
        v11 := GenS3{F0: uint64(v3), F1: 880, F2: "hello"}
        _ = v11
    }
    var v12 string = "\n\t"
    for fi13 := 0; fi13 < 17; fi13++ {
        var v14 float32 = (0.0 / 1.5)
        _ = fi13
        _ = v14
    }
    fmt.Println(v2)
    if v8 {
        for fi15 := 0; fi15 < 6; fi15++ {
            fmt.Println(fi15)
            for fi16 := 0; fi16 < 20; fi16++ {
                v2 += v2
                switch "\n\t" {
                case "hello":
                    v1 -= v1
                    v8 = (1.5 <= 1.5)
                    var v17 int64 = 422
                    _ = v17
                default:
                    v4 = 2147483648
                }
                _ = fi16
            }
            var v18 int8 = -67
            _ = fi15
            _ = v18
        }
    }
    v1 = (v1 % ((v1) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
