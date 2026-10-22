package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint }

func main() {
    v1 := make(map[int]bool)
    var v2 int16 = 32767
    var v3 int64 = -613
    var v4 int = 0
    var v5 uint32 = 787
    fmt.Println(v4)
    v6 := GenS1{F0: 0, F1: "a", F2: 65535}
    v3 = (^(v3))
    if (!false) {
        v7 := []int{}
        _ = v7
    } else {
        switch "foo bar" {
        case "\n\t":
            fmt.Println(v3, v5, v2)
        case "":
            fmt.Println(v3, v3, v5)
        case "hello":
            v8 := []int{}
            _ = v8
        }
        for fi9 := 0; fi9 < 14; fi9++ {
            if ((^(v2)) <= int16(v3)) {
                v3 -= -151
            } else {
                fmt.Println(v4, v3, v3)
                for fi10 := 0; fi10 < 14; fi10++ {
                    v11 := GenS1{F0: -9223372036854775808, F1: "hello", F2: 1}
                    if true {
                        v2 -= (v2 / (((v2 + v2)) | 1))
                    }
                    _ = fi10
                    _ = v11
                }
            }
            _ = fi9
        }
    }
    v4 = v4
    _ = len(v1)
    v12 := []int{1, (-(v4)), -654}
    var v13 int16 = (-(v2))
    var v14 int64 = 1
    if (true || false) {
        var v15 uint64 = 18446744073709551615
        v15 += (v15 / ((v15) | 1))
        var v16 int16 = (-(v2))
        _ = v15
        _ = v16
    } else {
        v17 := GenS1{F0: (v3 + (-(v3))), F1: ("foo bar" + ("hello" + ("" + ""))), F2: 418}
        v18 := GenS1{F0: v14, F1: ("foo bar" + ("" + "hello")), F2: 0}
        _ = v17
        _ = v18
    }
    fmt.Println(v13, v14)
    v19 := GenS1{F0: v3, F1: "", F2: 255}
    fmt.Println(v3, v13)
    _ = len(v12)
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    var v20 float64 = -2.25
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
