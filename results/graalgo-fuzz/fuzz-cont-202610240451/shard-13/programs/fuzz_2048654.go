package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 int; F2 int64 }
type GenS4 struct { F0 int; F1 uint8 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int; F1 uint64; F2 int64 }

func main() {
    var v1 bool = (false || (("" >= "") && true))
    v2 := []string{}
    var v3 int16 = 42
    v4 := make(map[int]string)
    var v5 int16 = (-(v3))
    switch "\n\t" {
    case "a":
        for fi6 := 0; fi6 < 9; fi6++ {
            v7 := GenS1{F0: 9223372036854775807}
            if (817 >= 4294967295) {
                switch "a" {
                case "hello":
                    switch (^(fi6)) {
                    case 131:
                        fmt.Println(v3)
                    }
                default:
                    v5 += (v3 & v5)
                }
                v1 = true
                var v8 string = "\n\t"
                _ = v8
            } else {
                for fi9 := 0; fi9 < 3; fi9++ {
                    var v10 float32 = (1.5 - (0.1 * (1.5 - -2.25)))
                    v11 := GenS1{F0: -1}
                    v12 := GenS2{F0: "hello"}
                    _ = fi9
                    _ = v10
                    _ = v11
                    _ = v12
                }
            }
            _ = fi6
            _ = v7
        }
    case "foo bar":
        if v1 {
            _ = len(v2)
            v13 := GenS3{F0: 2, F1: -404, F2: 578}
            _ = v13
        }
    case "":
        v3 += (v5 - 32767)
    }
    v14 := GenS4{F0: -530, F1: 42}
    v5 -= v3
    v15 := GenS5{F0: 950}
    var v16 float32 = 1.5
    var v17 uint16 = 2
    v18 := GenS4{F0: 2147483647, F1: 42}
    var v19 int8 = 100
    v20 := GenS6{F0: -966, F1: 4294967295, F2: 408}
    if (128 != 127) {
        fmt.Println(v17)
        v17 += 65535
    }
    v3 += v3
    var v21 uint64 = 1
    v22 := GenS3{F0: 933, F1: -323, F2: -150}
    v1 = (2 == 1)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
