package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 string; F2 uint32 }
type GenS3 struct { F0 int16; F1 int; F2 uint32 }
func genFn0() uint8 {
    v7 := GenS1{F0: 765}
    var v8 int16 = -1
    _ = v7
    _ = v8
    return 1
}

func main() {
    v1 := []uint32{362}
    var v2 float64 = float64(0.1)
    var v3 float32 = 0.1
    v4 := make(map[string]string)
    var v5 int64 = 42
    var v6 uint8 = 128
    v3 -= (3.14159 - ((-2.25 + -2.25) / -2.25))
    delete(v4, "")
    v9 := GenS1{F0: 42}
    if false {
        switch 2147483647 {
        case 2147483647:
            for fi10 := 0; fi10 < 5; fi10++ {
                fmt.Println(v6, v5)
                if false {
                    if (true || true) {
                        v11 := GenS2{F0: 18446744073709551615, F1: "a", F2: 2147483648}
                        v2 -= 1.5
                        v3 = v3
                        _ = v11
                    }
                }
                _ = fi10
            }
        case -829:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case -386:
            var v12 int32 = 42
            _ = v12
        }
        var v13 int16 = -711
        _ = v13
    }
    var v14 int64 = v5
    var v15 uint32 = 65535
    if (!((true || true) && (v5 != (v14 >> ((255) & 63))))) {
        var v16 float32 = ((((1e10 + float32(1e10)) / float32(-2.25)) / v3) / float32(0.1))
        v17 := GenS1{F0: 0}
        _ = v16
        _ = v17
    } else {
        var v18 uint32 = (v15 & (^(v15)))
        _ = v18
    }
    switch "\n\t" {
    case "foo bar":
        for fi19 := 0; fi19 < 12; fi19++ {
            v14 -= (v5 >> ((0) & 63))
            var v20 uint16 = 608
            for fi21 := 0; fi21 < 9; fi21++ {
                v2 = (v2 / 1.5)
                v22 := GenS1{F0: 2}
                _ = fi21
                _ = v22
            }
            _ = fi19
            _ = v20
        }
    }
    v23 := GenS3{F0: 1, F1: 2147483647, F2: (v15 | 1)}
    var v24 int8 = -1
    switch 0 {
    case -262:
        delete(v4, (("" + "a") + ((("" + (("foo bar" + (("hello" + ("\n\t" + "hello")) + "hello")) + "a")) + ("foo bar" + "hello")) + "hello")))
    }
    switch 4294967295 {
    case 9223372036854775807:
        fmt.Println(v14, v15)
    }
    v14 = (v5 ^ (v5 >> ((0) & 63)))
    var v25 int64 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
