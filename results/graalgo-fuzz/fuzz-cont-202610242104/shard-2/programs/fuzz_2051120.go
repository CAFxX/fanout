package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int8 }
func genFn0() int64 {
    var v5 int16 = -32768
    switch "a" {
    case "\n\t":
        if false {
            var v6 float32 = 1.5
            v5 -= (v5 & (v5 - (v5 << ((2) & 63))))
            v5 = (v5 << ((518) & 63))
            _ = v6
        }
    case "foo bar":
        v7 := GenS1{F0: 1, F1: ((("a" + ("foo bar" + "")) + (("foo bar" + ((("hello" + "\n\t") + ("foo bar" + "foo bar")) + ("foo bar" + ("\n\t" + ("" + ("a" + "hello")))))) + "foo bar")) + "hello")}
        _ = v7
    }
    if (false && true) {
        var v8 int8 = -51
        _ = v8
    }
    var v9 string = "foo bar"
    for fi10 := 0; fi10 < 14; fi10++ {
        v9 = "foo bar"
        fmt.Println(v5, fi10)
        var v11 uint16 = 842
        _ = fi10
        _ = v11
    }
    _ = v5
    _ = v9
    return 368
}

func main() {
    var v1 int = -1
    v2 := make(map[int]uint32)
    v3 := []int{}
    v4 := []uint32{1, 468, 2147483647}
    switch (-(v1)) {
    case -2147483648:
        v12 := GenS1{F0: 0, F1: "foo bar"}
        _ = v12
    default:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    switch "hello" {
    case "foo bar":
        v13 := GenS2{F0: 127, F1: 127}
        _ = v13
    case "\n\t":
        _ = len(v2)
    default:
        v14 := GenS3{F0: 127, F1: "a"}
        _ = v14
    }
    if true {
        if false {
            v2[v1] = 1
            v2[(v1 ^ v1)] = 4294967295
            for fi15 := 0; fi15 < 7; fi15++ {
                var v16 float64 = 1e10
                var v17 float32 = (1.5 / -2.25)
                _ = fi15
                _ = v16
                _ = v17
            }
        }
    }
    if (1 <= 255) {
        var v18 float64 = 3.14159
        _ = len(v2)
        _ = v18
    }
    fmt.Println(v1)
    for fi19 := 0; fi19 < 9; fi19++ {
        for fi20 := 0; fi20 < 5; fi20++ {
            v1 = fi19
            var v21 string = "foo bar"
            _ = fi20
            _ = v21
        }
        v1 = (v1 >> ((18446744073709551615) & 63))
        var v22 string = "a"
        _ = fi19
        _ = v22
    }
    var v23 uint32 = 2147483648
    v24 := GenS4{F0: 745}
    var v25 int32 = -1
    v26 := GenS5{F0: 103}
    switch v1 {
    case -797:
        if (18446744073709551615 > 454) {
            v1 = -9223372036854775808
            _ = len(v3)
        }
    case 42:
        for fi27 := 0; fi27 < 1; fi27++ {
            switch fi27 {
            case -892:
                if true {
                    if false {
                        v1 = (-(fi27))
                        var v28 uint32 = 2147483648
                        _ = v28
                    } else {
                        var v29 string = "a"
                        _ = v29
                    }
                }
            default:
                v1 = fi27
            }
            _ = fi27
        }
    default:
        v25 = 0
    }
    delete(v2, 128)
    var v30 uint16 = 281
    var v31 float32 = 0.1
    v25 = 2147483647
    v32 := GenS2{F0: 1, F1: 2}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
