package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint; F1 int32; F2 int8 }
type GenS4 struct { F0 uint32; F1 uint8; F2 uint64 }
func genFn0(p0 uint64, p1 int32, p2 uint16) uint16 {
    var v6 float32 = 1e10
    p1 = (p1 % (((p1 + (p1 * 42))) | 1))
    v7 := []int{1, 0, 32767}
    var v8 int32 = (p1 << ((42) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return (p2 - p2)
}
func genFn1(p0 int64, p1 uint32, p2 uint32) uint32 {
    p1 -= p1
    switch "" {
    case "hello":
        p1 += 65535
    case "a":
        p1 += p1
    default:
        for fi9 := 0; fi9 < 10; fi9++ {
            p0 = p0
            var v10 string = ""
            for fi11 := 0; fi11 < 4; fi11++ {
                if true {
                    fmt.Println(fi9, p0, v10)
                    var v12 uint16 = 1
                    fmt.Println(fi11, p1, p0)
                    var v13 string = "a"
                    _ = v12
                    _ = v13
                }
                fmt.Println(fi9)
                _ = fi11
            }
            _ = fi9
            _ = v10
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return (p2 * (p2 - p2))
}

func main() {
    v1 := make(map[string]bool)
    v2 := []uint32{65535}
    v3 := make(map[int]uint32)
    var v4 int8 = -1
    v5 := make(map[int]bool)
    v4 = -118
    var v14 uint8 = 127
    v15 := GenS1{F0: ("hello" + "\n\t"), F1: int(v4), F2: 2}
    delete(v1, "")
    v16 := GenS2{F0: (v4 + 42)}
    switch -2147483648 {
    case 9223372036854775807:
        var v17 int16 = 683
        _ = v17
    case 32767:
        var v18 uint16 = 65535
        _ = v18
    }
    v19 := GenS2{F0: (v4 ^ (v4 ^ -92))}
    if true {
        v14 += (v14 % ((v14) | 1))
    }
    v4 = (-(v4))
    _ = len(v2)
    v20 := GenS1{F0: "hello", F1: 0, F2: 4294967295}
    v3[128] = 499
    for fi21 := 0; fi21 < 4; fi21++ {
        v22 := GenS3{F0: 255, F1: 42, F2: (v4 + (v4 - (v4 & -1)))}
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if true {
            switch 93 {
            case 4294967295:
                switch fi21 {
                case -9223372036854775808:
                    v4 = (v4 << ((4294967295) & 63))
                case 4294967295:
                    var v23 string = "foo bar"
                    _ = v23
                case 32767:
                    var v24 int16 = 255
                    _ = v24
                default:
                    fmt.Println(v4)
                }
            default:
                v25 := GenS4{F0: 0, F1: (v14 / ((120) | 1)), F2: 2}
                _ = v25
            }
        } else {
            for fi26 := 0; fi26 < 15; fi26++ {
                v2 = append(v2, 290)
                _ = fi26
            }
            var v27 int64 = -1
            _ = v27
        }
        _ = fi21
        _ = v22
    }
    switch -72 {
    case 32768:
        v4 -= v4
    case -1:
        v4 -= v4
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
