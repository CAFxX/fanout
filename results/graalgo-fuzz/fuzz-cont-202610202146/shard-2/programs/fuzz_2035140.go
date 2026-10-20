package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 string; F1 int8; F2 uint64 }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint64) uint32 {
    for fi5 := 0; fi5 < 14; fi5++ {
        v6 := GenS1{F0: (("" + "hello") + ("foo bar" + "a")), F1: ("hello" + "")}
        var v7 uint16 = 836
        for fi8 := 0; fi8 < 15; fi8++ {
            v7 = 878
            fmt.Println(fi8)
            _ = fi8
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    p0 = 417
    var v9 float64 = 1.5
    v10 := GenS1{F0: "", F1: ""}
    _ = p0
    _ = v10
    _ = v9
    return 4294967295
}
func genFn1() uint8 {
    var v11 int64 = 481
    v12 := GenS2{F0: "foo bar", F1: -1, F2: 115}
    _ = v11
    _ = v12
    return 174
}

func main() {
    v1 := []string{"", ""}
    v2 := []int{-330, -532, -1, 502}
    var v3 uint64 = 4294967295
    v4 := []int{0}
    var v13 int8 = 21
    var v14 int8 = v13
    v15 := GenS2{F0: "hello", F1: -9, F2: (v3 ^ (v3 - (^(v3))))}
    var v16 int = 42
    v17 := GenS3{F0: -32768}
    switch v16 {
    case 127:
        for fi18 := 0; fi18 < 8; fi18++ {
            v19 := GenS2{F0: "", F1: (v13 ^ 42), F2: (v3 / (((v3 ^ 0)) | 1))}
            var v20 uint = 637
            _ = fi18
            _ = v19
            _ = v20
        }
    default:
        for fi21 := 0; fi21 < 17; fi21++ {
            var v22 float32 = -2.25
            _ = fi21
            _ = v22
        }
    }
    v13 = -125
    v13 = 104
    var v23 uint16 = 2
    v16 = v16
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
