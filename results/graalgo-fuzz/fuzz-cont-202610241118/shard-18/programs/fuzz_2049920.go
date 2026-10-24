package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 uint32 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 int8; F1 int64; F2 uint64 }
func genFn0() uint16 {
    var v5 int8 = 42
    if true {
        for fi6 := 0; fi6 < 4; fi6++ {
            switch "\n\t" {
            case "":
                if true {
                    fmt.Println(fi6, v5)
                }
            case "a":
                var v7 float32 = float32(-2.25)
                _ = v7
            default:
                fmt.Println(v5, fi6)
            }
            switch ("" + "") {
            case "":
                v8 := GenS1{F0: 65535, F1: 42, F2: 668}
                _ = v8
            case "a":
                var v9 uint64 = 305
                _ = v9
            case "hello":
                v5 = v5
            default:
                v5 += (v5 % ((v5) | 1))
            }
            fmt.Println(fi6, fi6)
            _ = fi6
        }
        fmt.Println(v5)
        v10 := GenS2{F0: 42, F1: -32768}
        _ = v10
    }
    if (2 < 2) {
        v5 -= 1
        v5 = (-(v5))
    }
    _ = v5
    return 21
}

func main() {
    var v1 uint64 = 749
    v2 := []string{"hello", "\n\t", ((("a" + ("" + "hello")) + ("\n\t" + "")) + "foo bar"), "\n\t"}
    v3 := make(map[uint32]string)
    var v4 float32 = -2.25
    var v11 uint32 = 2147483647
    var v12 int = 32768
    fmt.Println(v1, v4)
    v13 := GenS3{F0: 98, F1: 9223372036854775807, F2: v1}
    var v14 uint64 = v1
    var v15 int16 = 413
    if false {
        fmt.Println(v15, v15, v11)
    }
    v1 += (v14 | v14)
    if true {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v16 uint16 = 333
        _ = v16
    } else {
        _ = len(v3)
        v17 := GenS1{F0: 835, F1: -87, F2: (v11 & 1)}
        _ = v17
    }
    for fi18 := 0; fi18 < 10; fi18++ {
        var v19 uint32 = v11
        _ = fi18
        _ = v19
    }
    var v20 uint64 = (v14 & v14)
    for fi21 := 0; fi21 < 19; fi21++ {
        v22 := GenS1{F0: 32767, F1: 100, F2: (v11 ^ 0)}
        _ = fi21
        _ = v22
    }
    v2 = append(v2, "a")
    v23 := GenS2{F0: 2147483647, F1: (v15 - (v15 + (v15 >> ((474) & 63))))}
    for fi24 := 0; fi24 < 10; fi24++ {
        if true {
            v25 := GenS1{F0: 2, F1: -128, F2: v11}
            v26 := v3[(v11 + v11)]
            if false {
                _ = len(v2)
            } else {
                var v27 uint16 = uint16(v1)
                _ = v27
            }
            _ = v25
            _ = v26
        }
        _ = fi24
    }
    for fi28 := 0; fi28 < 17; fi28++ {
        var v29 int32 = -446
        switch (((("foo bar" + ("a" + (("foo bar" + "hello") + ((("foo bar" + "\n\t") + "hello") + (("foo bar" + ("" + ("\n\t" + ("hello" + "a")))) + ("foo bar" + ((("a" + "hello") + "a") + ("a" + ("a" + "a"))))))))) + ("a" + "foo bar")) + ("\n\t" + ("a" + ("a" + ("\n\t" + "hello"))))) + "") {
        case "\n\t":
            _ = len(v2)
        }
        _ = fi28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
}
