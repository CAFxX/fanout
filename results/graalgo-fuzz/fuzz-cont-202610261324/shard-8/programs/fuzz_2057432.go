package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool; F1 int32; F2 bool }
func genFn0() int32 {
    var v6 int32 = -2147483648
    for fi7 := 0; fi7 < 15; fi7++ {
        if (!true) {
            v6 -= 0
            for fi8 := 0; fi8 < 12; fi8++ {
                v6 = v6
                var v9 int = fi8
                if (115 <= 65535) {
                    v10 := GenS1{F0: ("hello" + "a"), F1: ((!false) && false), F2: 267}
                    _ = v10
                } else {
                    v11 := []int{(fi8 / ((fi8) | 1)), (fi7 * (fi8 * 42)), fi7}
                    _ = v11
                }
                if (((true && false) && (-1 != -393)) || (1 == 1)) {
                    fmt.Println(fi8)
                }
                v6 = v6
                fmt.Println(v9, v9, fi7)
                _ = fi8
                _ = v9
            }
        }
        v12 := []int{(fi7 >> ((65535) & 63))}
        _ = fi7
        _ = v12
    }
    _ = v6
    return v6
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]bool)
    v4 := []uint32{492}
    v5 := make(map[string]uint32)
    for fi13 := 0; fi13 < 14; fi13++ {
        v14 := GenS2{F0: false}
        _ = fi13
        _ = v14
    }
    var v15 uint32 = 2147483648
    switch "a" {
    case "foo bar":
        for fi16 := 0; fi16 < 7; fi16++ {
            if len(v4) > 0 {
                v4[0] = (v15 & (v15 % (((v15 * 2147483648)) | 1)))
            }
            _ = fi16
        }
    case "":
        var v17 uint64 = 18446744073709551615
        _ = v17
    }
    if (false && true) {
        v18 := GenS3{F0: true, F1: 2147483647, F2: (true || false)}
        fmt.Println(v15)
        v15 -= (^(v15))
        _ = v18
    } else {
        fmt.Println(v15)
        v15 += 2147483648
    }
    switch ("hello" + "\n\t") {
    case "foo bar":
        v15 += v15
    case "\n\t":
        v15 += (v15 << ((4294967295) & 63))
    default:
        _ = len(v4)
    }
    var v19 int8 = -72
    var v20 int64 = -1
    for fi21 := 0; fi21 < 6; fi21++ {
        v15 = (v15 & (v15 & (v15 | (v15 + 4294967295))))
        _ = fi21
    }
    v22 := v2[2147483648]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
