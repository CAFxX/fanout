package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
func genFn0(p0 string, p1 int32) int64 {
    if true {
        p1 += (p1 % ((-1) | 1))
        var v4 int64 = 890
        _ = v4
    }
    p1 = p1
    if true {
        fmt.Println(p0, p0)
        for fi5 := 0; fi5 < 8; fi5++ {
            fmt.Println(fi5, fi5, p0)
            switch p0 {
            case "a":
                v6 := GenS1{F0: 1, F1: 0}
                _ = v6
            case "foo bar":
                var v7 int = 32767
                _ = v7
            case "":
                switch "\n\t" {
                case "\n\t":
                    if ((-956 >= -758) && (777 >= 403)) {
                        var v8 uint8 = 1
                        for fi9 := 0; fi9 < 17; fi9++ {
                            v8 = v8
                            _ = fi9
                        }
                        var v10 uint32 = 1
                        _ = v10
                        _ = v8
                    }
                }
            }
            _ = fi5
        }
        var v11 uint = 0
        _ = v11
    } else {
        if true {
            fmt.Println(p0)
        }
    }
    _ = p0
    _ = p1
    return 225
}
func genFn1(p0 string) uint64 {
    fmt.Println(p0)
    switch 4294967295 {
    case 0:
        v12 := GenS1{F0: 954, F1: 42}
        _ = v12
    case -2147483648:
        v13 := GenS1{F0: -289, F1: 255}
        _ = v13
    default:
        v14 := []int{128, 32768, -2147483648}
        _ = v14
    }
    var v15 uint16 = 255
    _ = p0
    _ = v15
    return 90
}

func main() {
    var v1 uint = 65535
    var v2 int8 = 100
    v3 := make(map[int]int)
    v16 := v3[575]
    v3[v16] = (v16 | (v16 & (v16 + v16)))
    var v17 bool = (749 <= 2147483647)
    delete(v3, v16)
    var v18 int16 = 0
    var v19 int = v16
    var v20 uint = v1
    fmt.Println(v20, v19, v1)
    switch "foo bar" {
    case "hello":
        v21 := []int{(v19 ^ (v16 << (((v20 / (((v1 + 2)) | 1))) & 63))), (v16 | (v19 << ((18446744073709551615) & 63)))}
        _ = v21
    default:
        if false {
            v16 -= v16
        }
    }
    var v22 bool = (!((-(v16)) < v16))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v3
}
