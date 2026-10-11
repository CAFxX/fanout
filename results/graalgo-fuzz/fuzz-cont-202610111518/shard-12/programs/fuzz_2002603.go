package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 bool; F1 int64 }

func main() {
    v1 := make(map[uint32]string)
    var v2 string = ("\n\t" + ("" + ""))
    v3 := []uint32{4294967295, 306, 897, 4294967295}
    v4 := make(map[string]string)
    v5 := make(map[string]int)
    if (!(float64(float32(1.5)) <= 0.1)) {
        if (true && true) {
            var v6 bool = false
            _ = v6
        } else {
            var v7 float64 = ((0.1 * 0.0) - 0.1)
            _ = v7
        }
        switch v2 {
        case "a":
            switch 166 {
            case 528:
                v2 = (v2 + "foo bar")
            case 32767:
                v2 = v2
            case 0:
                v8 := GenS1{F0: 774, F1: 42}
                _ = v8
            }
        case "":
            v3 = append(v3, 2)
        case "foo bar":
            switch 0 {
            case 128:
                var v9 float32 = 0.0
                _ = v9
            }
        default:
            var v10 int16 = 1
            _ = v10
        }
    }
    v2 = ""
    fmt.Println(v2)
    switch -203 {
    case 128:
        v11 := GenS2{F0: 0}
        _ = v11
    case 672:
        for fi12 := 0; fi12 < 12; fi12++ {
            for fi13 := 0; fi13 < 11; fi13++ {
                fmt.Println(fi13, fi12)
                if (!false) {
                    v14 := GenS3{F0: 42, F1: -1}
                    fmt.Println(fi13, fi12, fi12)
                    _ = v14
                }
                var v15 uint64 = 4294967295
                _ = fi13
                _ = v15
            }
            _ = fi12
        }
    case 32767:
        v2 = v2
    default:
        v2 = (v2 + ("a" + v2))
    }
    var v16 int32 = 1
    v16 = 2147483647
    v17 := GenS4{F0: false, F1: 1}
    switch -1 {
    case 2147483647:
        v1[0] = v2
    }
    if true {
        v18 := GenS3{F0: -1, F1: -32768}
        var v19 float32 = 1.5
        _ = v18
        _ = v19
    } else {
        v2 = v2
    }
    _ = len(v1)
    var v20 uint32 = 130
    for fi21 := 0; fi21 < 8; fi21++ {
        v1[v20] = (v2 + (v2 + v2))
        _ = fi21
    }
    var v22 uint16 = 1
    v16 = -245
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
