package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 int; F1 int32; F2 uint }
type GenS3 struct { F0 int64; F1 int32 }
func genFn0(p0 uint16, p1 bool) uint16 {
    fmt.Println(p0, p0)
    v6 := []int{241, 128, 32768}
    _ = p0
    _ = p1
    _ = v6
    return (p0 << ((2) & 63))
}
func genFn1() uint16 {
    var v7 string = ("" + "a")
    var v8 uint32 = 467
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 int8 = -47
    v2 := make(map[int]string)
    v3 := make(map[int]int)
    var v4 int = -558
    v5 := []uint32{0, 827, 4294967295, 2147483647}
    var v9 int32 = 480
    var v10 float64 = 1e10
    fmt.Println(v4, v1, v10)
    v9 += v9
    var v11 int16 = 0
    v12 := v2[(v4 - 2147483647)]
    switch "\n\t" {
    case "a":
        v13 := v3[(v4 - v4)]
        _ = v13
    case "hello":
        var v14 bool = (!(genFn1() != 146))
        _ = v14
    case "foo bar":
        v1 = (v1 + (v1 >> ((18446744073709551615) & 63)))
    }
    delete(v2, (v4 / ((int(v1)) | 1)))
    if (!false) {
        for fi15 := 0; fi15 < 16; fi15++ {
            fmt.Println(v12, v9)
            v5 = append(v5, 2147483648)
            _ = fi15
        }
        fmt.Println(v4, v12)
    }
    if false {
        switch 32768 {
        case 128:
            if false {
                v16 := GenS1{F0: v1, F1: 18446744073709551615}
                for fi17 := 0; fi17 < 17; fi17++ {
                    v18 := GenS2{F0: v4, F1: 1, F2: 0}
                    v19 := v2[v4]
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                _ = v16
            } else {
                v20 := GenS2{F0: 2147483647, F1: v9, F2: 518}
                v21 := GenS3{F0: 1, F1: (v9 % (((v9 ^ v9)) | 1))}
                var v22 uint16 = 0
                _ = v20
                _ = v21
                _ = v22
            }
        case 32767:
            _ = len(v2)
        }
    }
    v23 := v2[-1]
    fmt.Println(v9, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
