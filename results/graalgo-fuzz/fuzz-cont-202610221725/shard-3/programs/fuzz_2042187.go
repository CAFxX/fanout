package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 int8; F1 int64; F2 int16 }

func main() {
    var v1 int16 = -715
    v2 := []uint32{65535}
    var v3 float32 = ((1e10 / -2.25) - float32((-2.25 - 0.0)))
    var v4 int = 58
    for fi5 := 0; fi5 < 16; fi5++ {
        _ = len(v2)
        fmt.Println(v4)
        v6 := GenS1{F0: true, F1: 0}
        _ = fi5
        _ = v6
    }
    var v7 uint8 = 128
    var v8 uint32 = 0
    switch v4 {
    case -2147483648:
        if len(v2) > 0 {
            v2[0] = 231
        }
    case 916:
        var v9 uint = uint(v7)
        _ = v9
    }
    v10 := GenS2{F0: -248, F1: v8}
    for fi11 := 0; fi11 < 10; fi11++ {
        v3 += v3
        v4 -= fi11
        v12 := GenS3{F0: ("\n\t" + ("hello" + "\n\t")), F1: (v7 | (v7 + (v7 | 127)))}
        _ = fi11
        _ = v12
    }
    fmt.Println(v8, v8)
    var v13 int16 = (v1 ^ 843)
    var v14 int16 = 0
    if (("" + "a") > "foo bar") {
        v2 = append(v2, (v8 << ((18446744073709551615) & 63)))
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        v8 += uint32(v7)
        v15 := GenS4{F0: 42, F1: -1, F2: 0}
        var v16 uint16 = 42
        _ = v15
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
