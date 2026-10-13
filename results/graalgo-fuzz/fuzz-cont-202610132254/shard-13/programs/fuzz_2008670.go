package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int32 }
func genFn0(p0 uint32) bool {
    for fi4 := 0; fi4 < 4; fi4++ {
        v5 := GenS1{F0: (p0 << ((18446744073709551615) & 63))}
        v6 := GenS2{F0: ("" + "\n\t")}
        _ = fi4
        _ = v5
        _ = v6
    }
    p0 -= 951
    _ = p0
    return ("foo bar" > (((("a" + "hello") + "hello") + ("hello" + "")) + ""))
}

func main() {
    var v1 int = -402
    v2 := []uint32{}
    v3 := []int{v1, v1, (-(v1))}
    var v7 uint8 = 128
    var v8 int64 = 9223372036854775807
    v9 := GenS1{F0: 4294967295}
    v10 := GenS3{F0: (v8 >> ((255) & 63))}
    for fi11 := 0; fi11 < 10; fi11++ {
        v12 := GenS3{F0: v8}
        _ = fi11
        _ = v12
    }
    var v13 int32 = -830
    v14 := GenS2{F0: ("hello" + "hello")}
    for fi15 := 0; fi15 < 16; fi15++ {
        fmt.Println(fi15)
        _ = fi15
    }
    v8 += (v8 % ((42) | 1))
    switch "hello" {
    case "":
        switch 127 {
        case -839:
            fmt.Println(v8, v8, v8)
        default:
            for fi16 := 0; fi16 < 4; fi16++ {
                v17 := GenS2{F0: "foo bar"}
                _ = fi16
                _ = v17
            }
        }
    case "\n\t":
        v18 := GenS2{F0: (("\n\t" + "\n\t") + "\n\t")}
        _ = v18
    case "a":
        var v19 int32 = v13
        _ = v19
    default:
        for fi20 := 0; fi20 < 4; fi20++ {
            v21 := GenS4{F0: int32(fi20)}
            _ = fi20
            _ = v21
        }
    }
    if len(v3) > 0 {
        v3[0] = 127
    }
    v7 = (v7 % (((v7 >> ((0) & 63))) | 1))
    v8 -= (^(v8))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
