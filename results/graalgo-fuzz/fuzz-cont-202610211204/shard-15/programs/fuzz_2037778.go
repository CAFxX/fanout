package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 string }
type GenS2 struct { F0 int8; F1 string; F2 int16 }
type GenS3 struct { F0 uint64; F1 int; F2 int32 }
type GenS4 struct { F0 int16; F1 bool }
func genFn0(p0 int) int64 {
    v6 := GenS1{F0: true, F1: false, F2: ("foo bar" + "")}
    var v7 int64 = -419
    v8 := []int{(p0 + -176), (p0 % (((p0 * 42)) | 1)), (p0 >> ((255) & 63)), 32767}
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return (v7 << ((18446744073709551615) & 63))
}
func genFn1() uint32 {
    var v9 uint32 = 871
    v9 += 2
    v10 := []int{-824, 0, -1}
    if false {
        if true {
            v11 := GenS1{F0: (false && false), F1: false, F2: "hello"}
            _ = v11
        } else {
            v9 = 42
        }
    }
    switch ("a" + ("foo bar" + (("foo bar" + ("hello" + "foo bar")) + ("" + ("" + (("hello" + "foo bar") + ((("\n\t" + ("" + "hello")) + "hello") + ("hello" + ("a" + ""))))))))) {
    case "\n\t":
        for fi12 := 0; fi12 < 6; fi12++ {
            var v13 float64 = 1.5
            var v14 int64 = 37
            v14 = 203
            _ = fi12
            _ = v13
            _ = v14
        }
    case "hello":
        var v15 string = ("hello" + "foo bar")
        _ = v15
    default:
        for fi16 := 0; fi16 < 20; fi16++ {
            v9 -= v9
            fmt.Println(v9, v9)
            _ = fi16
        }
    }
    _ = v10
    _ = v9
    return v9
}

func main() {
    var v1 uint8 = 127
    v2 := make(map[uint32]bool)
    var v3 uint64 = 532
    v4 := make(map[uint32]uint32)
    var v5 int = 589
    fmt.Println(v3, v5)
    v1 = 1
    v5 -= v5
    fmt.Println(v5, v5, v3)
    v17 := GenS2{F0: 127, F1: "a", F2: int16(v1)}
    v3 = v3
    v18 := GenS3{F0: (v3 | 42), F1: v5, F2: -1}
    v19 := GenS3{F0: (v3 >> ((255) & 63)), F1: (v5 >> ((2) & 63)), F2: -752}
    v5 += (-(v5))
    v20 := GenS4{F0: 109, F1: (32767 <= 255)}
    v3 -= v3
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
