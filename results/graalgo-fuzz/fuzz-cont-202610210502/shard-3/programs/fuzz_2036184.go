package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 int8 }

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{2147483647}
    var v3 uint16 = 2
    v4 := make(map[string]string)
    v5 := []string{"foo bar", "\n\t", "hello"}
    v6 := GenS1{F0: -1}
    var v7 string = "a"
    var v8 int = 32767
    v3 -= (v3 * v3)
    var v9 int8 = -27
    var v10 uint = 0
    v11 := GenS1{F0: 0}
    v10 = (^(v10))
    if false {
        for fi12 := 0; fi12 < 2; fi12++ {
            var v13 uint16 = v3
            _ = fi12
            _ = v13
        }
        if false {
            v14 := v1[0]
            v10 = v10
            _ = v14
        } else {
            v15 := GenS2{F0: (!false)}
            _ = v15
        }
    } else {
        var v16 string = (("a" + "a") + v7)
        var v17 float32 = ((1.5 / (1e10 * (-2.25 + 3.14159))) / 1.0)
        switch (v8 | (v8 >> (((^(v10))) & 63))) {
        case -1:
            v8 -= v8
        }
        _ = v16
        _ = v17
    }
    if true {
        var v18 int8 = v9
        _ = v18
    } else {
        v9 += v9
    }
    switch ((v7 + "a") + v7) {
    case "":
        for fi19 := 0; fi19 < 10; fi19++ {
            var v20 uint32 = 42
            _ = fi19
            _ = v20
        }
    default:
        switch (v7 + v7) {
        case "\n\t":
            _ = len(v2)
        case "":
            var v21 int16 = -32768
            _ = v21
        }
    }
    fmt.Println(v3, v9)
    switch (v7 + ((v7 + v7) + v7)) {
    case "foo bar":
        if true {
            v22 := v4[(v7 + ("foo bar" + "a"))]
            _ = v22
        }
    case "":
        v3 -= v3
    case "\n\t":
        v23 := GenS3{F0: 2, F1: (^(v9))}
        _ = v23
    }
    v24 := v4[("hello" + v7)]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
