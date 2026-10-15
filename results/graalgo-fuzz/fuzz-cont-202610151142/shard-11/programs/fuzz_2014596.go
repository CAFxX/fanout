package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 string }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 int = 128
    var v2 uint = 65535
    v3 := make(map[string]bool)
    fmt.Println(v2)
    _ = len(v3)
    v4 := []int{(v1 / ((v1) | 1)), 1, v1}
    fmt.Println(v2, v2)
    _ = len(v3)
    v5 := GenS1{F0: "", F1: 42, F2: "a"}
    for fi6 := 0; fi6 < 13; fi6++ {
        v2 -= (v2 % (((^(v2))) | 1))
        _ = fi6
    }
    delete(v3, "a")
    var v7 int64 = -1
    var v8 float64 = float64((-2.25 * (1e10 + (-2.25 - 0.1))))
    v9 := GenS1{F0: "\n\t", F1: 230, F2: ("hello" + "foo bar")}
    if ("\n\t" >= "hello") {
        v10 := GenS2{F0: (v7 << (((v2 | (v2 | v2))) & 63)), F1: -128}
        fmt.Println(v8, v8, v7)
        _ = v10
    }
    var v11 int32 = -2147483648
    if (((1e10 * (0.0 - 1e10)) * 1e10) >= 0.0) {
        var v12 int32 = (-(v11))
        v2 = v2
        _ = v12
    } else {
        v13 := v3["foo bar"]
        var v14 uint = 0
        _ = v13
        _ = v14
    }
    v15 := GenS3{F0: 42}
    var v16 uint16 = 32767
    for fi17 := 0; fi17 < 7; fi17++ {
        for fi18 := 0; fi18 < 17; fi18++ {
            v19 := GenS3{F0: 795}
            var v20 int8 = -1
            var v21 uint = 65535
            if true {
                v22 := GenS3{F0: 402}
                _ = v22
            } else {
                v1 -= 0
                v4 = append(v4, v1)
            }
            _ = fi18
            _ = v19
            _ = v20
            _ = v21
        }
        _ = fi17
    }
    switch ("hello" + "a") {
    case "":
        v2 -= v2
    case "hello":
        v23 := GenS2{F0: v7, F1: 1}
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
