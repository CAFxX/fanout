package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 int32; F1 int16; F2 uint8 }
type GenS3 struct { F0 int }
func genFn0() int8 {
    var v7 uint64 = 0
    switch ("foo bar" + "") {
    case "foo bar":
        if ("a" == "hello") {
            switch "\n\t" {
            case "foo bar":
                var v8 int16 = -1
                _ = v8
            }
            if true {
                switch ("\n\t" + "") {
                case "a":
                    v9 := []int{2147483647}
                    _ = v9
                case "foo bar":
                    v7 = (v7 ^ 0)
                default:
                    v7 = (v7 & 18446744073709551615)
                }
                fmt.Println(v7)
                v10 := []int{-755}
                _ = v10
            }
            v11 := GenS1{F0: -9223372036854775808, F1: 67}
            _ = v11
        }
    case "hello":
        var v12 uint64 = (v7 & (v7 - v7))
        _ = v12
    }
    v13 := []int{4294967295, -9223372036854775808}
    _ = v13
    _ = v7
    return -1
}

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{42, 2147483647, 4294967295, 396}
    var v3 int = 32767
    var v4 uint8 = 42
    v5 := make(map[string]bool)
    var v6 float32 = float32(1.5)
    v4 += (v4 + 53)
    var v14 bool = true
    if ((^(v3)) >= (v3 << ((1) & 63))) {
        v15 := GenS2{F0: -2147483648, F1: -1, F2: v4}
        v16 := GenS3{F0: v3}
        if (v14 || true) {
            var v17 int8 = -47
            _ = v17
        }
        _ = v15
        _ = v16
    }
    v14 = false
    var v18 float32 = 1e10
    v5["foo bar"] = true
    _ = len(v2)
    fmt.Println(v6, v18, v18)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
