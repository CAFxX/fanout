package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }

func main() {
    v1 := make(map[string]bool)
    var v2 uint16 = 97
    v3 := make(map[string]uint32)
    var v4 int32 = 2147483647
    var v5 int32 = v4
    var v6 uint64 = 42
    for fi7 := 0; fi7 < 19; fi7++ {
        fmt.Println(v5, v5, v5)
        var v8 bool = ((^(fi7)) == (fi7 / ((32768) | 1)))
        var v9 uint8 = 128
        _ = fi7
        _ = v8
        _ = v9
    }
    if false {
        var v10 string = "hello"
        _ = v10
    }
    for fi11 := 0; fi11 < 16; fi11++ {
        v5 = (v5 / (((v5 - (^(v4)))) | 1))
        v12 := []int{fi11, -30}
        for fi13 := 0; fi13 < 10; fi13++ {
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
            if len(v12) > 0 {
                v12[0] = (fi11 << ((307) & 63))
            }
            _ = fi13
        }
        _ = fi11
        _ = v12
    }
    var v14 int8 = 42
    fmt.Println(v6)
    fmt.Println(v14, v14, v5)
    v15 := GenS1{F0: v2, F1: v2}
    var v16 int64 = 9223372036854775807
    v4 = int32(v14)
    v2 -= v2
    for fi17 := 0; fi17 < 16; fi17++ {
        for fi18 := 0; fi18 < 3; fi18++ {
            v19 := GenS1{F0: v2, F1: v2}
            v20 := []int{}
            switch "foo bar" {
            case "foo bar":
                var v21 uint = 65535
                _ = v21
            case "":
                v2 = 0
            default:
                v22 := GenS1{F0: (v2 * (v2 | 32767)), F1: (^(v2))}
                _ = v22
            }
            v6 += v6
            _ = fi18
            _ = v19
            _ = v20
        }
        fmt.Println(v4, fi17)
        _ = fi17
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
