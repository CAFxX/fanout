package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 uint16 }
func genFn0(p0 uint) int16 {
    for fi5 := 0; fi5 < 7; fi5++ {
        switch (fi5 & (fi5 ^ -174)) {
        case 32767:
            v6 := []int{122, fi5}
            _ = v6
        }
        _ = fi5
    }
    v7 := GenS1{F0: (32767 >= 127), F1: 1, F2: 255}
    _ = p0
    _ = v7
    return -995
}
func genFn1() uint16 {
    for fi8 := 0; fi8 < 13; fi8++ {
        var v9 string = "a"
        if (true && false) {
            v9 = v9
        }
        v10 := []int{}
        _ = fi8
        _ = v10
        _ = v9
    }
    if (255 >= 1) {
        v11 := []int{-144, 128, 42, -850}
        _ = len(v11)
        _ = v11
    } else {
        if true {
            var v12 int16 = genFn0(123)
            _ = v12
        }
        var v13 uint32 = 0
        _ = v13
    }
    var v14 bool = true
    _ = v14
    return 255
}

func main() {
    v1 := []uint32{2147483647, 31, 2147483648}
    var v2 int8 = 0
    var v3 int64 = 9223372036854775807
    v4 := make(map[int]bool)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v2, v3)
    var v15 string = "foo bar"
    if (2147483647 > 32768) {
        fmt.Println(v15)
        fmt.Println(v2, v15)
        v16 := GenS1{F0: (!true), F1: (v3 << ((2) & 63)), F2: 32767}
        _ = v16
    } else {
        for fi17 := 0; fi17 < 11; fi17++ {
            var v18 int = fi17
            var v19 int = 9223372036854775807
            v15 = v15
            _ = fi17
            _ = v18
            _ = v19
        }
        var v20 uint32 = 779
        _ = v20
    }
    for fi21 := 0; fi21 < 8; fi21++ {
        _ = len(v1)
        v22 := GenS1{F0: ((209 >= -2147483648) || true), F1: v3, F2: 2}
        _ = fi21
        _ = v22
    }
    var v23 float32 = (0.1 + 3.14159)
    switch 128 {
    case 32768:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        for fi24 := 0; fi24 < 12; fi24++ {
            switch "" {
            case "hello":
                v25 := GenS1{F0: false, F1: (^(v3)), F2: 2}
                _ = v25
            case "a":
                fmt.Println(v3)
            }
            for fi26 := 0; fi26 < 19; fi26++ {
                var v27 uint16 = 42
                _ = len(v4)
                _ = fi26
                _ = v27
            }
            _ = fi24
        }
    }
    if len(v1) > 0 {
        v1[0] = 2
    }
    delete(v4, 42)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
