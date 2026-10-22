package main

import (
    "fmt"
)

func genFn0(p0 int32, p1 int, p2 uint8) uint64 {
    p1 = (p1 * (p1 ^ -627))
    var v7 int64 = 1
    var v8 uint8 = (p2 - p2)
    if false {
        switch p1 {
        case 721:
            for fi9 := 0; fi9 < 18; fi9++ {
                fmt.Println(p1)
                fmt.Println(v7, v8)
                _ = fi9
            }
        case 9223372036854775807:
            for fi10 := 0; fi10 < 12; fi10++ {
                var v11 uint16 = 27
                _ = fi10
                _ = v11
            }
        case 32768:
            var v12 uint64 = 18446744073709551615
            _ = v12
        default:
            fmt.Println(p2, p0)
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 4294967295
}
func genFn1(p0 string) uint16 {
    v13 := []int{42, 9223372036854775807, 42, -2147483648}
    switch 264 {
    case 128:
        if true {
            _ = len(v13)
            if true {
                p0 = p0
            }
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
        }
    case 127:
        var v14 int64 = 1000
        _ = v14
    case 42:
        v13 = append(v13, -924)
    default:
        p0 = p0
    }
    switch p0 {
    case "\n\t":
        for fi15 := 0; fi15 < 4; fi15++ {
            var v16 uint8 = 128
            var v17 uint16 = 255
            _ = fi15
            _ = v16
            _ = v17
        }
    case "":
        if len(v13) > 0 {
            v13[0] = 9223372036854775807
        }
    }
    _ = p0
    _ = v13
    return 255
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := []uint32{46}
    v3 := []string{"\n\t"}
    v4 := make(map[uint32]int)
    v5 := make(map[int]string)
    v6 := make(map[string]bool)
    var v18 bool = false
    var v19 int = 2147483647
    v19 += (^(v19))
    _ = len(v2)
    v20 := v5[42]
    var v21 uint = 347
    v1 -= (v1 * v1)
    fmt.Println(v19)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
