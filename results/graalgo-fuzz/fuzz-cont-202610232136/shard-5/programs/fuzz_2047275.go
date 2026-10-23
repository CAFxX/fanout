package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 uint; F1 int8; F2 uint16 }
type GenS4 struct { F0 bool; F1 int8; F2 int }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int64; F1 bool; F2 uint }
func genFn0() uint {
    v5 := GenS1{F0: -2147483648, F1: 208}
    var v6 uint64 = 4294967295
    var v7 float64 = 0.0
    _ = v5
    _ = v6
    _ = v7
    return 18446744073709551615
}

func main() {
    var v1 int16 = 947
    var v2 int64 = -585
    v3 := make(map[int]uint32)
    var v4 float32 = 0.1
    if false {
        v4 -= v4
        var v8 uint64 = 4294967295
        _ = v8
    }
    fmt.Println(v1, v2, v2)
    for fi9 := 0; fi9 < 18; fi9++ {
        v1 += (v1 << ((343) & 63))
        v10 := GenS1{F0: -1, F1: -2147483648}
        _ = fi9
        _ = v10
    }
    v4 = float32(-2.25)
    fmt.Println(v4)
    v2 -= (v2 | v2)
    switch 32768 {
    case -607:
        v1 = 42
    case 32767:
        switch 799 {
        case 2147483647:
            v1 = v1
        default:
            var v11 int = 32768
            _ = v11
        }
    case 0:
        for fi12 := 0; fi12 < 1; fi12++ {
            for fi13 := 0; fi13 < 3; fi13++ {
                v14 := GenS2{F0: 0, F1: 2}
                switch (("foo bar" + "a") + (("\n\t" + ((("foo bar" + ("hello" + "foo bar")) + (("foo bar" + "foo bar") + "a")) + ("\n\t" + "\n\t"))) + (("foo bar" + ("foo bar" + "hello")) + "hello"))) {
                case "":
                    var v15 float32 = v4
                    _ = v15
                case "foo bar":
                    v2 = 9223372036854775807
                default:
                    v16 := GenS3{F0: 1, F1: 112, F2: 304}
                    _ = v16
                }
                _ = fi13
                _ = v14
            }
            if (42 > v1) {
                v17 := GenS4{F0: (1 < 729), F1: 53, F2: (fi12 % (((fi12 | fi12)) | 1))}
                v3[1] = 1
                _ = v17
            }
            _ = fi12
        }
    default:
        v2 += v2
    }
    v18 := GenS5{F0: 62}
    if false {
        v4 -= v4
        if true {
            v19 := GenS1{F0: -868, F1: -2147483648}
            _ = v19
        } else {
            v4 = v4
            v20 := GenS6{F0: v2, F1: false, F2: 0}
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 11; fi21++ {
        v1 = v1
        _ = fi21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
