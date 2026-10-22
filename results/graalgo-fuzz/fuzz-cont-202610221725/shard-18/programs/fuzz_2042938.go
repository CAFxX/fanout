package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 bool }
type GenS2 struct { F0 uint; F1 int8; F2 uint8 }
type GenS3 struct { F0 int32; F1 int }

func main() {
    var v1 int32 = 2147483647
    var v2 bool = (714 < 739)
    v3 := []uint32{}
    v4 := GenS1{F0: 1, F1: 0, F2: (!(!v2))}
    var v5 uint8 = 42
    v1 = 42
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (127 > 127) {
        if ("\n\t" != "foo bar") {
            var v6 string = (("\n\t" + ("foo bar" + ((("foo bar" + "foo bar") + ("" + "a")) + "a"))) + "hello")
            if len(v3) > 0 {
                v3[0] = 0
            }
            v7 := GenS2{F0: 883, F1: 0, F2: v5}
            _ = v6
            _ = v7
        }
        fmt.Println(v5, v5)
        var v8 uint32 = 2147483648
        _ = v8
    }
    _ = len(v3)
    v9 := GenS3{F0: 2147483647, F1: -1}
    var v10 int64 = -632
    fmt.Println(v1)
    var v11 int16 = -866
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
