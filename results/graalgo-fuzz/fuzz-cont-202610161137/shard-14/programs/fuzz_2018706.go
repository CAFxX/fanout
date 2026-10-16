package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint, p1 uint32) int64 {
    var v7 uint16 = 957
    for fi8 := 0; fi8 < 11; fi8++ {
        p0 = (p0 % (((p0 | p0)) | 1))
        p0 -= 18446744073709551615
        if true {
            v7 = 1
            for fi9 := 0; fi9 < 10; fi9++ {
                if ((!false) && (!((!false) && false))) {
                    p1 = p1
                } else {
                    if (!(fi8 != fi8)) {
                        if ((false && (("foo bar" != "hello") || true)) || (91 >= 0)) {
                            v10 := GenS1{F0: 998}
                            _ = v10
                        }
                    }
                    var v11 float64 = 1.5
                    _ = v11
                }
                _ = fi9
            }
            p1 = p1
        }
        _ = fi8
    }
    var v12 uint = 659
    switch "a" {
    case "a":
        p1 += (p1 << (((v12 % ((v12) | 1))) & 63))
    case "":
        var v13 uint16 = v7
        _ = v13
    default:
        var v14 string = "a"
        _ = v14
    }
    fmt.Println(v7, v7, p1)
    _ = p0
    _ = p1
    _ = v12
    _ = v7
    return 9223372036854775807
}

func main() {
    var v1 int32 = 42
    var v2 uint8 = 1
    var v3 int = -1
    var v4 int16 = 1000
    v5 := make(map[string]bool)
    v6 := make(map[uint32]uint32)
    for fi15 := 0; fi15 < 9; fi15++ {
        v5["\n\t"] = false
        _ = fi15
    }
    var v16 uint16 = 255
    delete(v5, ("\n\t" + ""))
    var v17 float32 = 0.1
    fmt.Println(v17)
    v18 := GenS2{F0: (v3 % ((127) | 1)), F1: "\n\t"}
    var v19 int8 = -92
    switch v3 {
    case -333:
        switch v3 {
        case 128:
            for fi20 := 0; fi20 < 16; fi20++ {
                var v21 int16 = (v4 >> ((42) & 63))
                fmt.Println(v17, v21, v17)
                fmt.Println(v19, v16, v16)
                _ = fi20
                _ = v21
            }
        case -2147483648:
            v22 := GenS3{F0: (v1 + (v1 * -207))}
            _ = v22
        default:
            var v23 uint16 = v16
            _ = v23
        }
    }
    v24 := []int{32768, (v3 + 128)}
    v4 = v4
    v1 = (v1 / (((v1 / (((v1 ^ (v1 + (v1 ^ (v1 % (((v1 >> ((4294967295) & 63))) | 1)))))) | 1))) | 1))
    for fi25 := 0; fi25 < 20; fi25++ {
        var v26 uint = 1
        _ = fi25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
