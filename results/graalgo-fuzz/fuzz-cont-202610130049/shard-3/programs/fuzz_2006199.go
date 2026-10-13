package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 uint8; F1 uint32 }

func main() {
    var v1 int32 = 22
    v2 := []uint32{1, 2147483647, 1, 2147483648}
    var v3 int64 = -1
    var v4 int64 = 1
    fmt.Println(v4, v4)
    var v5 int32 = v1
    v3 += 42
    for fi6 := 0; fi6 < 7; fi6++ {
        if (32767 <= fi6) {
            switch "a" {
            case "hello":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            default:
                if false {
                    if len(v2) > 0 {
                        v2[0] = 4294967295
                    }
                    v3 -= v4
                    v7 := GenS1{F0: 2147483648, F1: -2147483648}
                    _ = v7
                }
            }
        }
        if true {
            v3 += (v4 & (v3 % (((v3 * -9223372036854775808)) | 1)))
            if (v1 < (v5 + (v5 << ((65535) & 63)))) {
                var v8 int8 = 52
                _ = v8
            }
            v2 = append(v2, 0)
        } else {
            v2 = append(v2, 2147483647)
            var v9 float64 = -2.25
            _ = v9
        }
        v1 += v1
        _ = fi6
    }
    for fi10 := 0; fi10 < 13; fi10++ {
        if (18446744073709551615 == 336) {
            for fi11 := 0; fi11 < 9; fi11++ {
                v12 := GenS2{F0: 102, F1: 42}
                _ = fi11
                _ = v12
            }
            var v13 int32 = (v1 ^ (v5 % ((v5) | 1)))
            var v14 uint8 = 128
            _ = v13
            _ = v14
        } else {
            fmt.Println(fi10)
            v15 := GenS1{F0: 0, F1: fi10}
            fmt.Println(v4)
            _ = v15
        }
        v16 := GenS1{F0: 42, F1: (fi10 & 128)}
        _ = fi10
        _ = v16
    }
    v4 += (v3 + -276)
    v1 += (v5 % ((v5) | 1))
    _ = len(v2)
    var v17 int32 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
