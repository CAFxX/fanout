package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 uint8; F2 uint64 }
type GenS3 struct { F0 string; F1 uint16 }
func genFn0(p0 string, p1 uint32) uint {
    switch 32767 {
    case 32768:
        fmt.Println(p1)
    case -930:
        v7 := GenS1{F0: 42}
        _ = v7
    case 128:
        var v8 int16 = -32768
        _ = v8
    default:
        v9 := []int{}
        _ = v9
    }
    v10 := GenS1{F0: -128}
    var v11 int8 = 114
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 2
}
func genFn1(p0 uint8, p1 bool) int16 {
    p1 = ((-754 >= 2147483647) || (!(41 > 2)))
    for fi12 := 0; fi12 < 13; fi12++ {
        if (42 >= 827) {
            var v13 int = fi12
            switch fi12 {
            case -227:
                v14 := GenS1{F0: 127}
                _ = v14
            default:
                for fi15 := 0; fi15 < 7; fi15++ {
                    p1 = (fi12 == (fi15 / ((fi15) | 1)))
                    fmt.Println(v13, p0, p0)
                    var v16 uint64 = 0
                    var v17 int = (-(fi15))
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            }
            _ = v13
        }
        v18 := GenS2{F0: "", F1: 0, F2: 2}
        _ = fi12
        _ = v18
    }
    var v19 string = "\n\t"
    _ = p0
    _ = p1
    _ = v19
    return 32767
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]bool)
    var v3 string = "a"
    var v4 uint32 = 2
    var v5 uint32 = (v4 << ((2) & 63))
    var v6 int = 0
    if false {
        var v20 uint64 = 0
        _ = v20
    }
    v2[v4] = true
    fmt.Println(v4, v4, v6)
    v21 := []int{42, v6, (v6 % ((2147483647) | 1))}
    var v22 int = v6
    switch "foo bar" {
    case "hello":
        v1[v4] = true
    default:
        v23 := GenS3{F0: v3, F1: 2}
        _ = v23
    }
    fmt.Println(v22, v5)
    fmt.Println(v6)
    fmt.Println(v5)
    var v24 uint8 = 255
    fmt.Println(v3)
    if len(v21) > 0 {
        v21[0] = 128
    }
    var v25 int64 = -312
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
