package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 uint16 }
func genFn0() int16 {
    var v7 uint64 = 749
    if true {
        if (!(true && true)) {
            if false {
                var v8 uint8 = 2
                var v9 bool = (true || false)
                _ = v8
                _ = v9
            }
            v10 := GenS1{F0: -462}
            _ = v10
        } else {
            switch "\n\t" {
            case "\n\t":
                v11 := []int{}
                _ = v11
            case "a":
                switch "foo bar" {
                case "foo bar":
                    if (295 >= -32768) {
                        var v12 int32 = -2147483648
                        v13 := []int{-317, 32767, -2147483648}
                        var v14 uint16 = 1
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                }
            default:
                var v15 float32 = 1.5
                _ = v15
            }
        }
        if (898 <= 255) {
            for fi16 := 0; fi16 < 13; fi16++ {
                fmt.Println(fi16, fi16)
                _ = fi16
            }
            var v17 uint8 = 127
            v18 := []int{32768}
            _ = v17
            _ = v18
        }
    } else {
        v7 = v7
        v19 := GenS1{F0: 128}
        _ = v19
    }
    v20 := GenS1{F0: 32767}
    v21 := GenS2{F0: 65535, F1: ""}
    v22 := []int{-1, -1, -9223372036854775808, 32768}
    _ = v20
    _ = v21
    _ = v22
    _ = v7
    return -32768
}

func main() {
    var v1 float64 = (1.5 / 0.1)
    var v2 float64 = (1e10 + v1)
    v3 := make(map[int]uint32)
    v4 := []string{("\n\t" + "hello"), "hello", ""}
    v5 := make(map[uint32]uint32)
    v6 := make(map[string]string)
    _ = len(v3)
    var v23 int32 = 1
    v23 += (v23 - (-(v23)))
    switch "\n\t" {
    case "a":
        if (3.14159 < v2) {
            v24 := GenS3{F0: 378}
            _ = v24
        }
    case "hello":
        v2 = -2.25
    }
    var v25 int8 = -1
    v6[("" + ("hello" + "a"))] = ("" + "hello")
    if len(v4) > 0 {
        v4[0] = "a"
    }
    fmt.Println(v2, v1)
    var v26 uint = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
