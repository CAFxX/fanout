package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32; F1 int32 }

func main() {
    v1 := make(map[int]bool)
    v2 := []int{}
    v3 := make(map[string]string)
    var v4 uint32 = 2147483647
    for fi5 := 0; fi5 < 9; fi5++ {
        v6 := GenS1{F0: (42 <= 1), F1: 673, F2: 0}
        _ = fi5
        _ = v6
    }
    v4 = 2147483647
    var v7 float32 = 0.0
    if false {
        if false {
            var v8 uint8 = 1
            v9 := GenS2{F0: -1}
            _ = v8
            _ = v9
        } else {
            if len(v2) > 0 {
                v2[0] = 32768
            }
            _ = len(v2)
        }
        if (false && true) {
            if (182 != 919) {
                fmt.Println(v7, v4)
                var v10 int64 = 797
                for fi11 := 0; fi11 < 2; fi11++ {
                    v4 -= v4
                    _ = fi11
                }
                _ = v10
            }
        } else {
            for fi12 := 0; fi12 < 6; fi12++ {
                v13 := GenS2{F0: 1}
                v14 := GenS3{F0: 0, F1: -2147483648}
                v7 -= v7
                _ = fi12
                _ = v13
                _ = v14
            }
        }
        switch -9223372036854775808 {
        case -1:
            v15 := v1[0]
            _ = v15
        case 128:
            var v16 float64 = 0.0
            _ = v16
        case 1:
            var v17 int = -2147483648
            _ = v17
        }
    }
    v18 := v1[128]
    _ = len(v2)
    switch "hello" {
    case "hello":
        var v19 uint64 = 42
        _ = v19
    case "":
        switch 681 {
        case -612:
            v4 = v4
        case -393:
            var v20 float32 = ((v7 / v7) + (v7 - v7))
            _ = v20
        case 346:
            v18 = false
        default:
            if len(v2) > 0 {
                v2[0] = 9223372036854775807
            }
        }
    default:
        _ = len(v3)
    }
    var v21 int8 = -128
    v4 -= (v4 % ((v4) | 1))
    if (false || true) {
        _ = len(v2)
        for fi22 := 0; fi22 < 12; fi22++ {
            var v23 int8 = v21
            v23 -= -1
            _ = fi22
            _ = v23
        }
    } else {
        for fi24 := 0; fi24 < 9; fi24++ {
            var v25 uint8 = 0
            switch (fi24 % ((int(v4)) | 1)) {
            case 91:
                fmt.Println(v7, v7)
            case 42:
                var v26 int = 2147483647
                _ = v26
            }
            _ = fi24
            _ = v25
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v7
}
