package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 uint32 }

func main() {
    var v1 string = ("foo bar" + ("a" + ("a" + ("a" + "a"))))
    v2 := make(map[int]string)
    var v3 uint16 = 2
    v4 := make(map[uint32]uint32)
    var v5 uint32 = 4294967295
    v1 = v1
    var v6 int = -9223372036854775808
    v2[v6] = v1
    fmt.Println(v3, v3)
    var v7 string = v1
    var v8 uint = 18446744073709551615
    for fi9 := 0; fi9 < 3; fi9++ {
        for fi10 := 0; fi10 < 7; fi10++ {
            fmt.Println(v1)
            switch v6 {
            case -813:
                fmt.Println(v7, v8)
            case 128:
                v11 := GenS1{F0: 42}
                _ = v11
            default:
                var v12 string = v7
                _ = v12
            }
            if false {
                v6 -= -590
                v13 := GenS2{F0: v3, F1: (v5 & (v5 % ((uint32(v8)) | 1)))}
                v14 := []int{v6, (fi10 ^ (fi10 % ((-2147483648) | 1)))}
                _ = v13
                _ = v14
            } else {
                var v15 int8 = 127
                _ = v15
            }
            _ = fi10
        }
        for fi16 := 0; fi16 < 9; fi16++ {
            if true {
                v1 = "a"
                v6 = 127
            } else {
                v17 := GenS2{F0: (v3 << (((v8 ^ v8)) & 63)), F1: (v5 / (((^(v5))) | 1))}
                _ = v17
            }
            var v18 uint16 = 42
            _ = fi16
            _ = v18
        }
        _ = fi9
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
