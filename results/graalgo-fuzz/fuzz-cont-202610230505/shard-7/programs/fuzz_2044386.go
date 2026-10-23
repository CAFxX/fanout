package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 uint64; F1 uint64 }
func genFn0(p0 int64, p1 int32) string {
    switch -14 {
    case 128:
        v6 := []int{}
        _ = v6
    default:
        var v7 int16 = 0
        _ = v7
    }
    v8 := []int{779}
    p0 = p0
    if false {
        if (972 < 0) {
            p1 = p1
        } else {
            fmt.Println(p0, p1)
        }
    } else {
        switch "foo bar" {
        case "foo bar":
            if (!true) {
                for fi9 := 0; fi9 < 18; fi9++ {
                    p1 = 1
                    p1 = p1
                    for fi10 := 0; fi10 < 15; fi10++ {
                        v11 := GenS1{F0: -32768, F1: (p1 & (-(p1)))}
                        _ = fi10
                        _ = v11
                    }
                    _ = fi9
                }
                if len(v8) > 0 {
                    v8[0] = 32767
                }
            } else {
                var v12 int32 = (p1 - 42)
                _ = v12
            }
        case "hello":
            if (127 <= 255) {
                p1 += p1
            }
        case "a":
            var v13 uint64 = 758
            _ = v13
        default:
            _ = len(v8)
        }
        switch -2147483648 {
        case -1:
            for fi14 := 0; fi14 < 5; fi14++ {
                fmt.Println(p1)
                switch (^(fi14)) {
                case -1:
                    var v15 int = 9223372036854775807
                    _ = v15
                case -9223372036854775808:
                    v16 := GenS2{F0: 1, F1: (p0 | (p0 / ((p0) | 1)))}
                    _ = v16
                case 1:
                    if len(v8) > 0 {
                        v8[0] = (fi14 * 32768)
                    }
                }
                _ = fi14
            }
        case 937:
            switch -606 {
            case -287:
                if (!false) {
                    if false {
                        fmt.Println(p1)
                        fmt.Println(p1)
                    }
                }
            case 0:
                p0 = (p0 ^ (p0 | 42))
            }
        case -393:
            var v17 int8 = 127
            _ = v17
        }
    }
    _ = p0
    _ = p1
    _ = v8
    return (("hello" + "") + "a")
}
func genFn1(p0 int, p1 bool, p2 int) int8 {
    v18 := GenS3{F0: 1, F1: 2}
    v19 := GenS3{F0: 2, F1: 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v19
    return -24
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint16 = 431
    var v3 int = 766
    v4 := make(map[int]bool)
    v5 := make(map[uint32]int)
    var v20 uint16 = v2
    fmt.Println(v3, v3)
    var v21 float32 = float32(-2.25)
    var v22 int8 = 127
    v23 := []int{662}
    var v24 int16 = -142
    var v25 int = 32767
    var v26 int8 = v22
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
