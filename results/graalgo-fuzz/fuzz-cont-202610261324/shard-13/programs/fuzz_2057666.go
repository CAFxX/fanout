package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint32; F1 string; F2 uint16 }

func main() {
    v1 := make(map[string]string)
    v2 := []string{}
    var v3 int16 = 1
    v4 := make(map[uint32]int)
    v5 := []string{}
    v3 = -32768
    v6 := GenS1{F0: 907}
    v7 := GenS2{F0: false}
    var v8 uint = 474
    v9 := GenS3{F0: -1, F1: 1000}
    switch -9223372036854775808 {
    case 4294967295:
        var v10 string = "foo bar"
        _ = v10
    default:
        for fi11 := 0; fi11 < 10; fi11++ {
            var v12 int64 = -9223372036854775808
            fmt.Println(v8)
            if false {
                if (530 == 1) {
                    switch fi11 {
                    case 42:
                        v13 := GenS4{F0: "foo bar"}
                        _ = v13
                    default:
                        _ = len(v4)
                    }
                    fmt.Println(v12, v3, fi11)
                    fmt.Println(fi11, fi11)
                } else {
                    var v14 bool = true
                    _ = v14
                }
                var v15 int64 = (v12 << (((^(v8))) & 63))
                _ = v15
            }
            _ = fi11
            _ = v12
        }
    }
    v3 += 1000
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch "\n\t" {
    case "a":
        v3 += v3
    case "foo bar":
        var v16 int16 = 32767
        _ = v16
    case "hello":
        for fi17 := 0; fi17 < 1; fi17++ {
            var v18 bool = false
            if v18 {
                var v19 int8 = -1
                v19 += v19
                _ = v19
            }
            v20 := GenS5{F0: 2147483647, F1: "a", F2: 121}
            _ = fi17
            _ = v18
            _ = v20
        }
    }
    v3 -= (v3 / ((v3) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
