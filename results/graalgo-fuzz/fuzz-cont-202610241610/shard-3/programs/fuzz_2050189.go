package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 int8; F2 uint64 }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int32; F1 int16 }
func genFn0() string {
    switch "hello" {
    case "hello":
        v7 := GenS1{F0: ("" + ("foo bar" + "a"))}
        _ = v7
    case "foo bar":
        var v8 float64 = 1.5
        _ = v8
    }
    var v9 uint32 = 9
    _ = v9
    return (("" + ("foo bar" + ("a" + ("foo bar" + ("hello" + "a"))))) + "")
}
func genFn1(p0 bool) uint64 {
    switch 42 {
    case 9223372036854775807:
        for fi10 := 0; fi10 < 17; fi10++ {
            if ((-2.25 > 1.5) || p0) {
                for fi11 := 0; fi11 < 2; fi11++ {
                    v12 := GenS2{F0: 32767, F1: 42, F2: 2}
                    v13 := GenS3{F0: "a", F1: p0}
                    fmt.Println(p0, fi10, p0)
                    var v14 uint64 = 182
                    _ = fi11
                    _ = v12
                    _ = v13
                    _ = v14
                }
            }
            _ = fi10
        }
    }
    v15 := GenS3{F0: "hello", F1: (243 <= 42)}
    v16 := GenS4{F0: 1}
    _ = p0
    _ = v15
    _ = v16
    return 763
}

func main() {
    var v1 int32 = 1
    var v2 uint8 = 117
    v3 := make(map[uint32]int)
    var v4 uint = 65535
    var v5 int64 = -9223372036854775808
    var v6 int = 1
    v17 := GenS5{F0: int32(v6), F1: 1}
    var v18 uint32 = 65535
    fmt.Println(v18, v6)
    if false {
        switch (v6 - int(v4)) {
        case 1:
            _ = len(v3)
        case -947:
            for fi19 := 0; fi19 < 5; fi19++ {
                var v20 int8 = 0
                _ = fi19
                _ = v20
            }
        case 4294967295:
            fmt.Println(v4, v6)
        default:
            fmt.Println(v4, v1, v1)
        }
        v1 += v1
    }
    v1 += 2147483647
    switch "foo bar" {
    case "foo bar":
        v1 = (v1 - 1)
    case "\n\t":
        fmt.Println(v4)
    default:
        v4 = (v4 * 42)
    }
    v2 -= 127
    var v21 int32 = (v1 * (v1 << ((v4) & 63)))
    v6 -= v6
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
