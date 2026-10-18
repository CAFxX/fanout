package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string; F1 int16 }
type GenS4 struct { F0 uint16; F1 int32; F2 uint32 }
func genFn0() int32 {
    switch "a" {
    case "foo bar":
        var v4 uint16 = 255
        _ = v4
    case "a":
        if (1e10 > 0.0) {
            v5 := GenS1{F0: 42, F1: 0, F2: "foo bar"}
            _ = v5
        }
    case "\n\t":
        v6 := GenS2{F0: 786}
        _ = v6
    default:
        var v7 int = 32768
        _ = v7
    }
    var v8 int = 127
    v8 = v8
    _ = v8
    return -2147483648
}
func genFn1(p0 int) uint32 {
    p0 += 9223372036854775807
    v9 := []int{4294967295, 127, p0}
    if true {
        var v10 uint8 = 1
        if len(v9) > 0 {
            v9[0] = 4294967295
        }
        _ = v10
    }
    _ = len(v9)
    p0 -= (p0 | (-(p0)))
    _ = p0
    _ = v9
    return 2
}

func main() {
    v1 := []uint32{2}
    var v2 uint8 = 2
    v3 := []uint32{956, 2147483647, 2147483647}
    if (!((183 == -208) && (0 >= v2))) {
        if ((true || true) && true) {
            v2 -= v2
        } else {
            v2 += v2
        }
    } else {
        v2 += 42
        var v11 uint32 = 313
        _ = v11
    }
    var v12 int8 = -1
    v13 := GenS3{F0: ((("\n\t" + "") + "a") + ("\n\t" + "hello")), F1: 1}
    v14 := GenS3{F0: ("hello" + ""), F1: 1000}
    if (2 > 42) {
        if len(v1) > 0 {
            v1[0] = 872
        }
    } else {
        switch 127 {
        case 547:
            v1 = append(v1, 2147483647)
        }
    }
    v15 := GenS4{F0: 0, F1: 0, F2: 42}
    v16 := GenS2{F0: 42}
    v2 = (v2 & v2)
    v12 -= v12
    for fi17 := 0; fi17 < 5; fi17++ {
        fmt.Println(v2, fi17)
        if true {
            var v18 uint64 = 18446744073709551615
            _ = v18
        } else {
            v2 = v2
            v12 = (v12 / (((v12 | v12)) | 1))
        }
        _ = fi17
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
