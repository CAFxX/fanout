package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint64; F1 uint }
type GenS4 struct { F0 uint8; F1 int32 }

func main() {
    v1 := []string{"foo bar", ("hello" + "\n\t"), "a"}
    var v2 uint16 = 27
    var v3 uint = 875
    var v4 int64 = 1
    var v5 uint16 = 283
    v6 := make(map[string]uint32)
    var v7 uint64 = 4294967295
    fmt.Println(v4, v3)
    _ = len(v1)
    if true {
        for fi8 := 0; fi8 < 1; fi8++ {
            var v9 float32 = 1e10
            _ = fi8
            _ = v9
        }
        v2 = v5
        v7 -= (v7 + v7)
    } else {
        for fi10 := 0; fi10 < 7; fi10++ {
            switch (fi10 | -9223372036854775808) {
            case 32767:
                for fi11 := 0; fi11 < 15; fi11++ {
                    v5 -= (v2 / (((v5 & 2)) | 1))
                    fmt.Println(v2, v4)
                    _ = fi11
                }
            }
            var v12 int64 = v4
            var v13 bool = (!(fi10 <= 408))
            _ = fi10
            _ = v12
            _ = v13
        }
    }
    v6[((("" + "a") + ("" + ("a" + "\n\t"))) + ("foo bar" + "a"))] = 767
    fmt.Println(v7, v3, v4)
    if false {
        v14 := GenS1{F0: 158, F1: (v5 ^ 0), F2: v7}
        v4 = (v4 / ((v4) | 1))
        var v15 uint32 = 1
        _ = v14
        _ = v15
    }
    if true {
        v3 = (v3 ^ (^(v3)))
        _ = len(v1)
    } else {
        v5 = (v2 | v5)
        v16 := GenS2{F0: v3}
        fmt.Println(v2, v2, v5)
        _ = v16
    }
    switch 0 {
    case -2147483648:
        fmt.Println(v7)
    case 564:
        v17 := GenS3{F0: 4294967295, F1: (v3 - 758)}
        _ = v17
    default:
        v18 := GenS4{F0: 105, F1: -269}
        _ = v18
    }
    if (-1 < -510) {
        v1 = append(v1, "")
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
