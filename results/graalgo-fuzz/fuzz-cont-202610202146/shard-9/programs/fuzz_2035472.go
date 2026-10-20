package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 int }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 int64; F1 uint32; F2 bool }
func genFn0(p0 int8) int32 {
    v4 := GenS1{F0: "a", F1: 2147483647, F2: -535}
    for fi5 := 0; fi5 < 1; fi5++ {
        p0 = p0
        var v6 string = "\n\t"
        var v7 uint = 1
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 int16 = -32768
    v9 := GenS1{F0: "\n\t", F1: 128, F2: 42}
    _ = p0
    _ = v4
    _ = v8
    _ = v9
    return -2147483648
}

func main() {
    var v1 uint = 178
    v2 := []string{"\n\t"}
    v3 := make(map[uint32]bool)
    for fi10 := 0; fi10 < 16; fi10++ {
        var v11 uint = (^(v1))
        if (0 <= 63) {
            for fi12 := 0; fi12 < 15; fi12++ {
                if false {
                    var v13 uint = (v1 / ((v11) | 1))
                    var v14 int16 = -189
                    _ = v13
                    _ = v14
                } else {
                    fmt.Println(v11, v11)
                    fmt.Println(fi12, fi12, fi12)
                    v11 += v11
                }
                _ = fi12
            }
            var v15 string = "\n\t"
            _ = v15
        } else {
            v2 = append(v2, "")
        }
        var v16 float64 = 1.5
        _ = fi10
        _ = v11
        _ = v16
    }
    fmt.Println(v1)
    if (42 != 1) {
        v17 := GenS2{F0: 4294967295, F1: "hello"}
        _ = v17
    }
    var v18 int = -9223372036854775808
    v19 := GenS3{F0: -236, F1: 2, F2: true}
    var v20 int64 = -1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v21 := GenS3{F0: (v20 / ((v20) | 1)), F1: 555, F2: true}
    var v22 uint64 = 59
    switch "hello" {
    case "hello":
        v23 := GenS2{F0: v22, F1: "a"}
        _ = v23
    case "":
        var v24 uint = v1
        _ = v24
    default:
        v20 = -372
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
