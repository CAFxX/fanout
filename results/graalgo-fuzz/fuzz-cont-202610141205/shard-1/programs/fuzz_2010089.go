package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 uint8; F1 int; F2 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint64; F1 int32 }
type GenS5 struct { F0 bool }

func main() {
    var v1 float32 = -2.25
    var v2 uint16 = 430
    v3 := make(map[int]string)
    v4 := make(map[string]string)
    v5 := []string{(("a" + "foo bar") + "\n\t"), "\n\t"}
    v6 := []uint32{2147483648, 65535, 42, 330}
    if (v1 != v1) {
        for fi7 := 0; fi7 < 3; fi7++ {
            v8 := GenS1{F0: 0, F1: -901}
            for fi9 := 0; fi9 < 6; fi9++ {
                if true {
                    fmt.Println(v1, fi7)
                }
                _ = fi9
            }
            _ = fi7
            _ = v8
        }
    }
    v10 := GenS2{F0: 182, F1: -1, F2: "hello"}
    if false {
        var v11 int16 = 0
        _ = v11
    }
    v12 := GenS2{F0: 42, F1: 9223372036854775807, F2: "a"}
    for fi13 := 0; fi13 < 8; fi13++ {
        var v14 float32 = (v1 - 3.14159)
        for fi15 := 0; fi15 < 12; fi15++ {
            fmt.Println(fi15, v14)
            _ = fi15
        }
        _ = fi13
        _ = v14
    }
    v16 := GenS3{F0: 591}
    for fi17 := 0; fi17 < 7; fi17++ {
        for fi18 := 0; fi18 < 11; fi18++ {
            delete(v3, 9223372036854775807)
            var v19 int32 = 2147483647
            v1 = v1
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if len(v6) > 0 {
        v6[0] = 912
    }
    v1 += 0.1
    for fi20 := 0; fi20 < 5; fi20++ {
        v21 := GenS4{F0: 1, F1: 2147483647}
        v22 := GenS4{F0: 129, F1: -287}
        _ = fi20
        _ = v21
        _ = v22
    }
    switch 9223372036854775807 {
    case 4294967295:
        fmt.Println(v2, v1)
    }
    v23 := GenS5{F0: false}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v12
    _ = v16
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
