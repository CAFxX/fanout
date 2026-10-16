package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 uint16; F1 int16 }
func genFn0(p0 int32, p1 uint32, p2 uint64) uint {
    if false {
        if ((1.5 * 0.0) == (0.1 * 0.0)) {
            var v4 int32 = 1
            v5 := GenS1{F0: -436}
            _ = v4
            _ = v5
        }
        p1 = (p1 & 2147483648)
    }
    var v6 uint = 170
    if (("foo bar" + "\n\t") == "hello") {
        v7 := GenS2{F0: (false && (true && (32767 <= 807))), F1: 42}
        var v8 int32 = -91
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return v6
}
func genFn1(p0 int64) string {
    v9 := GenS2{F0: false, F1: 42}
    switch -9223372036854775808 {
    case -722:
        v10 := []int{-1, 127}
        _ = v10
    case 42:
        if false {
            p0 = 1
            var v11 string = ""
            _ = v11
        } else {
            var v12 int32 = -131
            _ = v12
        }
    case 32768:
        p0 += p0
    }
    _ = p0
    _ = v9
    return "hello"
}

func main() {
    var v1 int16 = -229
    var v2 int = -1
    var v3 uint16 = 2
    v13 := []int{v2}
    _ = len(v13)
    fmt.Println(v2)
    v2 = (v2 + -343)
    var v14 uint32 = 2147483647
    for fi15 := 0; fi15 < 13; fi15++ {
        fmt.Println(v2, v2)
        _ = fi15
    }
    v3 = (v3 + (v3 / (((v3 / ((42) | 1))) | 1)))
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    if (0.0 == 0.1) {
        v16 := GenS3{F0: 94, F1: -32768}
        v1 += v1
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
}
