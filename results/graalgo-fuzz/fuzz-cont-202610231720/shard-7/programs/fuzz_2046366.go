package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16; F1 uint; F2 int16 }
type GenS4 struct { F0 uint; F1 uint32; F2 int64 }
type GenS5 struct { F0 int64; F1 string }

func main() {
    var v1 string = "hello"
    v2 := make(map[int]bool)
    v3 := make(map[uint32]string)
    var v4 int16 = 42
    var v5 uint64 = 2
    var v6 int32 = 207
    switch -921 {
    case 0:
        fmt.Println(v6, v6)
    case -2147483648:
        fmt.Println(v5, v1)
    }
    switch (v1 + ("\n\t" + v1)) {
    case "":
        v7 := GenS1{F0: (!(!((^(v4)) >= (v4 | v4))))}
        _ = v7
    case "\n\t":
        v1 = "a"
    }
    for fi8 := 0; fi8 < 19; fi8++ {
        v9 := GenS2{F0: 42}
        _ = fi8
        _ = v9
    }
    var v10 int64 = 9223372036854775807
    if true {
        v11 := []int{}
        v4 -= (v4 + v4)
        _ = v11
    } else {
        v6 -= 1
        var v12 int16 = 994
        v4 = int16(v5)
        _ = v12
    }
    v13 := GenS3{F0: (v4 << ((1) & 63)), F1: 1, F2: -1}
    for fi14 := 0; fi14 < 6; fi14++ {
        v10 -= (-(v10))
        if false {
            fmt.Println(fi14, v6)
        }
        _ = fi14
    }
    v15 := []int{int(v6)}
    if true {
        v6 += v6
        var v16 int64 = (-(v10))
        for fi17 := 0; fi17 < 17; fi17++ {
            var v18 string = v1
            if false {
                v19 := v2[int(v6)]
                _ = len(v3)
                var v20 uint32 = 1
                _ = v19
                _ = v20
            }
            v6 += 0
            _ = fi17
            _ = v18
        }
        _ = v16
    }
    var v21 int16 = 32767
    fmt.Println(v21, v1, v10)
    v22 := GenS4{F0: 1, F1: 1, F2: v10}
    v1 = (v1 + "")
    var v23 int64 = (v10 / ((v10) | 1))
    if (4294967295 != 255) {
        switch "\n\t" {
        case "":
            var v24 string = (v1 + v1)
            _ = v24
        }
    } else {
        var v25 uint64 = (^(v5))
        _ = v25
    }
    v1 = v1
    for fi26 := 0; fi26 < 16; fi26++ {
        if (!false) {
            v27 := GenS5{F0: v23, F1: v1}
            var v28 uint16 = 74
            _ = v27
            _ = v28
        }
        _ = fi26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v13
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
