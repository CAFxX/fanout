package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 int64; F1 int32 }
func genFn0(p0 int8, p1 int16) uint64 {
    p1 = (^(p1))
    p1 += (p1 * -932)
    _ = p0
    _ = p1
    return 4294967295
}

func main() {
    var v1 int16 = -492
    var v2 int = -2147483648
    var v3 uint32 = 4294967295
    v4 := []string{("hello" + "a"), "", (("a" + "hello") + "foo bar"), ""}
    var v5 uint = 65535
    var v6 int = v2
    v3 -= 2
    v3 += v3
    if (801 >= 32767) {
        fmt.Println(v1)
        v7 := GenS1{F0: 2147483647, F1: 2147483647}
        _ = v7
    }
    var v8 uint32 = v3
    var v9 int16 = (v1 + v1)
    v8 = (v3 * (v3 & v8))
    var v10 uint8 = 0
    if (v6 >= (v6 << (((^(v5))) & 63))) {
        v1 = v9
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        v12 := GenS1{F0: 42, F1: 2147483647}
        _ = fi11
        _ = v12
    }
    v9 += (-(v9))
    v13 := GenS2{F0: v9, F1: v5}
    for fi14 := 0; fi14 < 12; fi14++ {
        v6 = (v6 << (((^(v5))) & 63))
        var v15 int64 = 897
        _ = fi14
        _ = v15
    }
    var v16 int = (v6 ^ v6)
    v17 := GenS3{F0: -9223372036854775808, F1: 1}
    v6 = 4294967295
    switch "" {
    case "":
        for fi18 := 0; fi18 < 4; fi18++ {
            v19 := GenS1{F0: 0, F1: -12}
            _ = fi18
            _ = v19
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
