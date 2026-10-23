package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 int8 }
type GenS2 struct { F0 int; F1 int32; F2 uint8 }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 uint64 = 987
    var v2 string = "\n\t"
    var v3 bool = (144 < uint8(v1))
    v4 := make(map[uint32]string)
    v5 := []uint32{2, 42, 1, 283}
    v6 := make(map[string]int)
    _ = len(v5)
    switch 190 {
    case -2147483648:
        if (1 <= 32767) {
            v7 := GenS1{F0: v1, F1: 42, F2: 127}
            _ = v7
        }
    case 4294967295:
        for fi8 := 0; fi8 < 10; fi8++ {
            v2 = ("" + v2)
            _ = fi8
        }
    case 676:
        var v9 uint = 65535
        _ = v9
    default:
        var v10 int8 = -42
        _ = v10
    }
    fmt.Println(v2, v3, v2)
    fmt.Println(v3, v3, v2)
    _ = len(v6)
    for fi11 := 0; fi11 < 10; fi11++ {
        switch fi11 {
        case 0:
            for fi12 := 0; fi12 < 10; fi12++ {
                if false {
                    var v13 int32 = 0
                    _ = v13
                }
                v4[uint32(v1)] = ("hello" + "a")
                if (2 < 0) {
                    v14 := GenS2{F0: -298, F1: 42, F2: 42}
                    if len(v5) > 0 {
                        v5[0] = 323
                    }
                    _ = v14
                }
                _ = fi12
            }
        default:
            v15 := GenS3{F0: 42}
            _ = v15
        }
        for fi16 := 0; fi16 < 12; fi16++ {
            fmt.Println(v2)
            _ = len(v6)
            var v17 uint8 = 42
            _ = fi16
            _ = v17
        }
        v2 = (("foo bar" + v2) + "foo bar")
        _ = fi11
    }
    for fi18 := 0; fi18 < 4; fi18++ {
        v1 = v1
        v1 -= 286
        _ = len(v4)
        _ = fi18
    }
    if (!v3) {
        fmt.Println(v3, v3, v2)
        v3 = v3
    } else {
        v19 := GenS3{F0: 0}
        v1 += (v1 ^ 18446744073709551615)
        if v3 {
            var v20 bool = false
            v2 = "a"
            _ = v20
        }
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
