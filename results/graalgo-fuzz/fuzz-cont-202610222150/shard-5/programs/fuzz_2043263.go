package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 int16; F1 int; F2 uint16 }
type GenS3 struct { F0 bool }
func genFn0(p0 int32) int32 {
    for fi4 := 0; fi4 < 18; fi4++ {
        for fi5 := 0; fi5 < 18; fi5++ {
            fmt.Println(fi5, fi4)
            p0 = (p0 / (((p0 / (((p0 * p0)) | 1))) | 1))
            var v6 uint32 = 65535
            _ = fi5
            _ = v6
        }
        v7 := GenS1{F0: (fi4 < fi4), F1: ""}
        p0 = int32(fi4)
        _ = fi4
        _ = v7
    }
    v8 := GenS2{F0: 1000, F1: 9223372036854775807, F2: 451}
    p0 += p0
    _ = p0
    _ = v8
    return p0
}
func genFn1(p0 uint64, p1 int64) int64 {
    var v9 float64 = (1e10 + 3.14159)
    var v10 string = ""
    fmt.Println(v9)
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return p1
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]string)
    var v3 uint16 = 88
    if (!true) {
        var v11 uint8 = 0
        _ = v11
    } else {
        v3 -= (v3 ^ v3)
        v3 -= 42
        fmt.Println(v3)
    }
    v3 = v3
    var v12 bool = (!(1 >= 1))
    v13 := v1["foo bar"]
    var v14 float32 = ((1.5 * -2.25) - 3.14159)
    switch 68 {
    case 128:
        v15 := GenS3{F0: ((!v12) && true)}
        _ = v15
    case 4294967295:
        var v16 uint32 = (^(v13))
        _ = v16
    case -1:
        fmt.Println(v13, v3, v3)
    default:
        var v17 int32 = 935
        _ = v17
    }
    v13 += (v13 / (((v13 * (v13 | (v13 * 4294967295)))) | 1))
    fmt.Println(v12, v14, v12)
    fmt.Println(v12)
    fmt.Println(v12)
    var v18 uint = 803
    for fi19 := 0; fi19 < 6; fi19++ {
        if ((!false) || (("hello" + "") <= (("foo bar" + "foo bar") + "hello"))) {
            v20 := []int{fi19, int(v3), (fi19 + fi19), 41}
            _ = v20
        }
        for fi21 := 0; fi21 < 14; fi21++ {
            v22 := GenS2{F0: -1, F1: (fi19 << (((v18 + (^(v18)))) & 63)), F2: (v3 / ((32767) | 1))}
            var v23 int32 = 210
            fmt.Println(v18, fi19)
            _ = fi21
            _ = v22
            _ = v23
        }
        v24 := v1["foo bar"]
        _ = fi19
        _ = v24
    }
    var v25 uint16 = (v3 * v3)
    var v26 int8 = 0
    v3 -= (^(v3))
    var v27 int8 = (^(v26))
    var v28 int = -798
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
}
