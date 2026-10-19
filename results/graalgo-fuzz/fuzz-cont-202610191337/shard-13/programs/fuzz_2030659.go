package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16; F1 uint64; F2 int16 }
type GenS4 struct { F0 uint64 }

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 663
    var v3 uint = 978
    for fi4 := 0; fi4 < 5; fi4++ {
        fmt.Println(v2, v3)
        v5 := GenS1{F0: -843, F1: 32767}
        v6 := GenS2{F0: (v2 << ((2) & 63))}
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 float32 = (((float32(0.0) + 1.5) - (float32(float64(-2.25)) + 1e10)) + -2.25)
    var v8 uint16 = v2
    for fi9 := 0; fi9 < 7; fi9++ {
        v10 := GenS3{F0: -1, F1: 0, F2: -32768}
        v3 = v3
        _ = fi9
        _ = v10
    }
    var v11 int64 = 42
    for fi12 := 0; fi12 < 4; fi12++ {
        var v13 float64 = 1e10
        if (!true) {
            switch fi12 {
            case -489:
                var v14 uint64 = 877
                _ = v14
            }
        }
        var v15 float32 = v7
        _ = fi12
        _ = v13
        _ = v15
    }
    for fi16 := 0; fi16 < 3; fi16++ {
        v17 := GenS1{F0: v11, F1: 1}
        _ = fi16
        _ = v17
    }
    var v18 string = "foo bar"
    switch 32768 {
    case -427:
        if (-1 == -128) {
            if true {
                delete(v1, 32768)
            }
            var v19 int8 = 100
            _ = v19
        }
    case 2147483647:
        fmt.Println(v18, v7)
    case -1:
        var v20 float64 = 0.1
        _ = v20
    }
    v11 = v11
    v2 = (^(v2))
    v21 := GenS4{F0: 1}
    v22 := v1[4294967295]
    var v23 int8 = -128
    for fi24 := 0; fi24 < 1; fi24++ {
        v8 = 543
        v11 -= -600
        var v25 bool = (42 != 1)
        _ = fi24
        _ = v25
    }
    v26 := []int{-2147483648, -620, 128}
    v1[-2147483648] = "foo bar"
    if true {
        _ = len(v26)
        var v27 bool = false
        _ = v27
    } else {
        v22 = v18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v26
    _ = v3
    _ = v7
    _ = v8
}
