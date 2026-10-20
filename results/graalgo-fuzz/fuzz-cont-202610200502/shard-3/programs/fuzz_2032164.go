package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 int32 }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 uint; F1 string }

func main() {
    v1 := []uint32{}
    var v2 float32 = (((-2.25 / 0.1) + 3.14159) + 0.1)
    v3 := make(map[uint32]int)
    v4 := make(map[string]bool)
    switch "foo bar" {
    case "a":
        var v5 uint16 = 2
        _ = v5
    case "":
        for fi6 := 0; fi6 < 13; fi6++ {
            switch fi6 {
            case 719:
                switch ("\n\t" + "foo bar") {
                case "foo bar":
                    v2 += v2
                }
            }
            fmt.Println(v2, fi6)
            _ = fi6
        }
    case "foo bar":
        if (!false) {
            var v7 int16 = -750
            var v8 bool = (-1 == 980)
            _ = v7
            _ = v8
        } else {
            for fi9 := 0; fi9 < 20; fi9++ {
                var v10 int32 = 1
                _ = fi9
                _ = v10
            }
        }
    }
    if true {
        if (0 > 10) {
            v11 := v3[2]
            if len(v1) > 0 {
                v1[0] = 942
            }
            _ = v11
        } else {
            v12 := GenS1{F0: 2, F1: false, F2: -2147483648}
            _ = v12
        }
        if true {
            v2 += ((v2 * 3.14159) / v2)
            v2 = v2
        } else {
            for fi13 := 0; fi13 < 7; fi13++ {
                v2 += ((v2 - v2) - (0.0 * ((v2 / 1e10) - v2)))
                var v14 float32 = v2
                _ = fi13
                _ = v14
            }
        }
    } else {
        var v15 int8 = 42
        v16 := GenS2{F0: 0, F1: 4294967295}
        _ = v15
        _ = v16
    }
    var v17 float32 = v2
    v18 := GenS3{F0: 42, F1: 2}
    var v19 uint32 = 2147483647
    v20 := GenS2{F0: 0, F1: 767}
    v17 -= v17
    v19 += (v19 & 502)
    v17 = 1e10
    var v21 int32 = -1
    v21 = (v21 & v21)
    switch ("foo bar" + "\n\t") {
    case "a":
        _ = len(v3)
    case "\n\t":
        v22 := GenS4{F0: 640, F1: "hello"}
        _ = v22
    }
    fmt.Println(v19, v21, v19)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
