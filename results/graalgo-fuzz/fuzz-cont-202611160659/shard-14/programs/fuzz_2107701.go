package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 uint16 }
type GenS2 struct { F0 int; F1 uint; F2 uint }
type GenS3 struct { F0 uint; F1 int8; F2 int64 }

func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{42, 65535, 42}
    var v3 bool = (585 <= 0)
    var v4 int = -743
    v5 := GenS1{F0: 255, F1: 909, F2: 643}
    for fi6 := 0; fi6 < 9; fi6++ {
        for fi7 := 0; fi7 < 8; fi7++ {
            var v8 float32 = 0.0
            _ = fi7
            _ = v8
        }
        if v3 {
            fmt.Println(v3, v4, v3)
        } else {
            v3 = (873 == 2147483648)
            var v9 uint = 2
            _ = v9
        }
        _ = fi6
    }
    fmt.Println(v4, v3)
    for fi10 := 0; fi10 < 19; fi10++ {
        fmt.Println(v4)
        _ = fi10
    }
    var v11 uint64 = 42
    if v3 {
        for fi12 := 0; fi12 < 14; fi12++ {
            v3 = false
            switch ("\n\t" + "\n\t") {
            case "a":
                fmt.Println(fi12)
            case "hello":
                v13 := GenS1{F0: 2, F1: v11, F2: 1}
                _ = v13
            case "":
                v11 -= (v11 + v11)
            default:
                v11 += (^(v11))
            }
            _ = fi12
        }
        v1[v4] = 2147483647
    }
    if (-17 > 0) {
        v14 := GenS2{F0: v4, F1: 4294967295, F2: 213}
        _ = v14
    }
    if false {
        switch "foo bar" {
        case "foo bar":
            if (!(-811 <= 0)) {
                if true {
                    delete(v1, (v4 - 42))
                    fmt.Println(v3, v4, v11)
                }
                if true {
                    v11 += (v11 ^ v11)
                } else {
                    v15 := GenS1{F0: 1, F1: (v11 & v11), F2: 266}
                    v11 -= v11
                    var v16 int8 = 1
                    _ = v15
                    _ = v16
                }
                fmt.Println(v11, v3)
            }
        case "":
            switch "a" {
            case "foo bar":
                fmt.Println(v4, v3)
            }
        case "\n\t":
            var v17 string = ""
            _ = v17
        }
        for fi18 := 0; fi18 < 5; fi18++ {
            fmt.Println(fi18, v4)
            var v19 bool = true
            _ = fi18
            _ = v19
        }
    }
    switch "hello" {
    case "\n\t":
        for fi20 := 0; fi20 < 10; fi20++ {
            v21 := GenS3{F0: 901, F1: 42, F2: 42}
            delete(v1, (-(v4)))
            v1[-667] = 712
            _ = fi20
            _ = v21
        }
    }
    v3 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
