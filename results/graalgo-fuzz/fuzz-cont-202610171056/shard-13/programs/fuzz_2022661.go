package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool; F1 bool }

func main() {
    var v1 int64 = 1
    v2 := make(map[uint32]uint32)
    v3 := []string{"a", ("a" + ""), "a", "foo bar"}
    var v4 float64 = 1.5
    v5 := make(map[uint32]uint32)
    v6 := GenS1{F0: 586, F1: (v1 << ((4294967295) & 63)), F2: 1}
    var v7 float64 = (float64(1e10) / 1.5)
    var v8 uint8 = 0
    if false {
        v9 := GenS2{F0: 2}
        _ = v9
    }
    _ = len(v5)
    var v10 uint64 = 42
    switch ("" + ("a" + "")) {
    case "hello":
        var v11 float32 = -2.25
        _ = v11
    default:
        fmt.Println(v1, v4)
    }
    v12 := GenS1{F0: 2, F1: v1, F2: (v10 >> ((794) & 63))}
    v13 := GenS3{F0: 42}
    var v14 uint8 = (v8 & (v8 - (^(v8))))
    var v15 bool = (-204 <= int16(v1))
    v8 = 128
    v16 := GenS3{F0: 2147483647}
    v17 := GenS4{F0: false, F1: (!false)}
    for fi18 := 0; fi18 < 16; fi18++ {
        if (!(v14 <= (v8 % ((v8) | 1)))) {
            var v19 int32 = -2147483648
            switch (("a" + "\n\t") + (("a" + (("hello" + ("hello" + "hello")) + "")) + "hello")) {
            case "foo bar":
                if false {
                    var v20 float32 = (3.14159 - 0.0)
                    var v21 string = (("foo bar" + "\n\t") + ("\n\t" + "\n\t"))
                    _ = v20
                    _ = v21
                }
            case "":
                if len(v3) > 0 {
                    v3[0] = ("foo bar" + "a")
                }
            default:
                v1 += (v1 * v1)
            }
            v22 := GenS2{F0: 2}
            _ = v19
            _ = v22
        }
        var v23 bool = (!(65535 != 1))
        _ = fi18
        _ = v23
    }
    fmt.Println(v7)
    v10 = (v10 / ((v10) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
