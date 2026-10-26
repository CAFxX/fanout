package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 uint8 }
type GenS2 struct { F0 string; F1 int32; F2 int }
type GenS3 struct { F0 uint8; F1 bool; F2 int32 }
func genFn0(p0 uint, p1 uint32, p2 uint32) int64 {
    v4 := GenS1{F0: 18446744073709551615, F1: -32768, F2: 1}
    p1 = 65535
    switch ((("a" + "a") + ("hello" + "\n\t")) + "") {
    case "foo bar":
        p0 -= p0
    case "":
        p1 += 191
    default:
        v5 := []int{76}
        _ = v5
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 42
}
func genFn1(p0 string, p1 int32) uint64 {
    p1 = p1
    v6 := []int{}
    _ = p0
    _ = p1
    _ = v6
    return 258
}

func main() {
    v1 := []uint32{}
    var v2 bool = false
    var v3 uint8 = 127
    v2 = true
    if v2 {
        v3 += v3
        var v7 float64 = ((-2.25 * 0.0) - 0.0)
        _ = v7
    }
    switch "" {
    case "a":
        var v8 float64 = 0.0
        _ = v8
    case "foo bar":
        fmt.Println(v3, v3)
    }
    v3 += 128
    var v9 float64 = float64(0.0)
    switch ("" + "\n\t") {
    case "a":
        _ = len(v1)
    default:
        v10 := GenS1{F0: 42, F1: 42, F2: v3}
        _ = v10
    }
    if (-1 < -85) {
        v11 := GenS2{F0: "hello", F1: 2147483647, F2: -661}
        _ = len(v1)
        v2 = (!v2)
        _ = v11
    }
    if ((65535 != 1) && (83 == 1)) {
        v2 = v2
        if (!(42 <= -1)) {
            v2 = (603 < 2)
        } else {
            var v12 int8 = 100
            _ = v12
        }
        var v13 int32 = 42
        _ = v13
    } else {
        fmt.Println(v9, v2)
        var v14 uint32 = 867
        _ = v14
    }
    switch -798 {
    case -1:
        var v15 bool = v2
        _ = v15
    case 2147483647:
        v3 = v3
    default:
        var v16 int32 = 711
        _ = v16
    }
    for fi17 := 0; fi17 < 8; fi17++ {
        var v18 int = fi17
        fmt.Println(v3, v3)
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 10; fi19++ {
        var v20 float32 = 3.14159
        v3 = (v3 ^ (v3 + (v3 * (v3 << ((0) & 63)))))
        _ = fi19
        _ = v20
    }
    fmt.Println(v9)
    v21 := GenS3{F0: 227, F1: (!(721 <= 2)), F2: -2147483648}
    v3 += v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v9
}
