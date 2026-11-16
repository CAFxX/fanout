package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 bool; F1 uint8 }
func genFn0(p0 int8, p1 bool, p2 uint64) uint8 {
    p2 += p2
    var v4 int16 = 0
    for fi5 := 0; fi5 < 16; fi5++ {
        if p1 {
            v6 := []int{1, 0, fi5}
            _ = v6
        }
        fmt.Println(p0, p2, p2)
        _ = fi5
    }
    switch "a" {
    case "a":
        if false {
            v7 := []int{}
            var v8 string = ""
            p0 = p0
            _ = v7
            _ = v8
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 255
}
func genFn1(p0 int32, p1 uint16) int16 {
    if false {
        v9 := GenS1{F0: 334, F1: -128}
        p1 = (p1 << ((5) & 63))
        switch 2147483647 {
        case 9223372036854775807:
            switch -1 {
            case 2147483647:
                v10 := GenS2{F0: (p1 / (((p1 * p1)) | 1))}
                _ = v10
            }
        }
        _ = v9
    }
    var v11 int = -88
    var v12 float32 = float32(1.5)
    v13 := []int{(v11 % ((v11) | 1)), 636, v11}
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v13
    return -32768
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{"\n\t"}
    v3 := []int{240, -59}
    for fi14 := 0; fi14 < 10; fi14++ {
        var v15 float32 = 1e10
        _ = fi14
        _ = v15
    }
    var v16 string = ""
    v2 = append(v2, v16)
    var v17 float32 = (float32(0.0) * -2.25)
    v16 = v16
    v18 := GenS3{F0: 14, F1: 662}
    fmt.Println(v16)
    v19 := GenS4{F0: -1}
    var v20 uint32 = 2147483648
    switch v16 {
    case "foo bar":
        _ = len(v3)
    case "hello":
        var v21 bool = true
        _ = v21
    default:
        v22 := GenS5{F0: false, F1: 133}
        _ = v22
    }
    v23 := GenS2{F0: 561}
    v16 = v16
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
}
