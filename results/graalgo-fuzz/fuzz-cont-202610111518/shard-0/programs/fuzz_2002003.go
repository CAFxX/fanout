package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS2 struct { F0 bool; F1 string; F2 int }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 int32; F1 string; F2 uint }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]bool)
    v3 := make(map[int]string)
    var v4 uint64 = 2
    switch "" {
    case "":
        fmt.Println(v4)
    case "\n\t":
        v4 = 18446744073709551615
    case "foo bar":
        v4 += 813
    }
    var v5 int32 = 2147483647
    var v6 float32 = 0.0
    v6 += v6
    v7 := GenS1{F0: 1, F1: v4, F2: 111}
    fmt.Println(v5, v5, v4)
    fmt.Println(v4, v6)
    v8 := GenS1{F0: 0, F1: v4, F2: 172}
    if (-210 != 42) {
        v4 -= (^(v4))
        v9 := GenS2{F0: false, F1: "hello", F2: 1}
        _ = v9
    } else {
        for fi10 := 0; fi10 < 1; fi10++ {
            v11 := GenS1{F0: 0, F1: (^(v4)), F2: 255}
            if true {
                for fi12 := 0; fi12 < 17; fi12++ {
                    fmt.Println(v5, v4, v4)
                    _ = fi12
                }
            }
            var v13 int8 = -1
            _ = fi10
            _ = v11
            _ = v13
        }
        var v14 int8 = 0
        _ = v14
    }
    for fi15 := 0; fi15 < 20; fi15++ {
        v16 := GenS3{F0: 1, F1: -128}
        for fi17 := 0; fi17 < 10; fi17++ {
            switch fi15 {
            case -298:
                fmt.Println(fi17, fi17, fi15)
            case 869:
                if ((fi17 <= (fi17 + fi17)) && (true && true)) {
                    fmt.Println(fi17, v6, fi15)
                }
            default:
                v5 = v5
            }
            _ = fi17
        }
        _ = fi15
        _ = v16
    }
    var v18 uint = 4294967295
    var v19 uint16 = 32767
    v20 := GenS4{F0: v5, F1: "foo bar", F2: (v18 ^ v18)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
