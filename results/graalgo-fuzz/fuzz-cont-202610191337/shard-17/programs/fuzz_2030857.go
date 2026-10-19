package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int64 }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 uint32 = 42
    var v2 bool = true
    var v3 int8 = -128
    var v4 float32 = -2.25
    var v5 uint16 = 255
    switch "\n\t" {
    case "hello":
        var v6 uint = 1
        _ = v6
    case "foo bar":
        v7 := []int{-1, 42, 1}
        _ = v7
    case "a":
        fmt.Println(v1, v5)
    }
    v8 := []int{1, 32767}
    switch 0 {
    case 32767:
        fmt.Println(v3, v4, v5)
    }
    var v9 string = "foo bar"
    fmt.Println(v4)
    fmt.Println(v9, v5)
    for fi10 := 0; fi10 < 10; fi10++ {
        v4 = 1e10
        if false {
            if false {
                var v11 string = "\n\t"
                var v12 string = ""
                v13 := GenS1{F0: 9223372036854775807, F1: 32767, F2: -589}
                _ = v11
                _ = v12
                _ = v13
            } else {
                if ((!(523 < 32767)) || (-387 < (fi10 / ((fi10) | 1)))) {
                    fmt.Println(fi10)
                    v9 = (v9 + v9)
                    fmt.Println(fi10, v4)
                } else {
                    var v14 uint64 = 4294967295
                    _ = v14
                }
                v8 = append(v8, fi10)
                v15 := GenS2{F0: 42}
                _ = v15
            }
            if (1 > 32767) {
                v16 := GenS1{F0: 1, F1: 1, F2: 525}
                _ = v16
            }
        }
        _ = fi10
    }
    if len(v8) > 0 {
        v8[0] = 603
    }
    var v17 int16 = 1
    switch ((v9 + ((v9 + v9) + (((v9 + v9) + "hello") + v9))) + (v9 + ("hello" + v9))) {
    case "":
        if v2 {
            fmt.Println(v5, v3, v1)
            var v18 uint16 = (^(v5))
            _ = v18
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
