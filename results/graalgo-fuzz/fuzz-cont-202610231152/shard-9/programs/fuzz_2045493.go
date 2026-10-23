package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32; F1 int8 }

func main() {
    v1 := []int{}
    var v2 bool = false
    var v3 float32 = -2.25
    var v4 bool = (128 == 4294967295)
    _ = len(v1)
    for fi5 := 0; fi5 < 12; fi5++ {
        switch ("\n\t" + "hello") {
        case "":
            v2 = (!v2)
        case "a":
            var v6 float32 = float32(0.0)
            _ = v6
        }
        if (!(1 < 1)) {
            if (-612 == 0) {
                var v7 uint8 = 255
                var v8 int32 = -991
                fmt.Println(v4, v3, v8)
                _ = v7
                _ = v8
            }
            v4 = true
            v9 := GenS1{F0: v2, F1: 42}
            _ = v9
        }
        _ = fi5
    }
    v10 := GenS2{F0: 42}
    if len(v1) > 0 {
        v1[0] = 42
    }
    if v2 {
        fmt.Println(v4)
        v2 = ((!v4) || ((65535 != 1) && (v3 == v3)))
    }
    v11 := GenS3{F0: 42, F1: -81}
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
}
