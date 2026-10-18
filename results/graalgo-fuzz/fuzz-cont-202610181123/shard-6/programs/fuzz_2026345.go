package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 int16; F1 uint64; F2 uint8 }
type GenS3 struct { F0 int8; F1 uint64; F2 uint8 }
func genFn0(p0 uint32, p1 uint32) uint64 {
    v4 := []int{32768, 9223372036854775807, -122}
    var v5 float64 = -2.25
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 0
}

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 42
    var v3 uint64 = (v2 >> ((592) & 63))
    fmt.Println(v2, v3)
    fmt.Println(v3, v3)
    for fi6 := 0; fi6 < 6; fi6++ {
        switch ("hello" + ("" + "")) {
        case "foo bar":
            if true {
                if (true || false) {
                    v2 += (v2 & (v2 >> ((243) & 63)))
                    var v7 float64 = (0.0 / ((float64(1.5) + 1.5) / 1.0))
                    _ = v7
                }
                switch "" {
                case "foo bar":
                    v8 := []int{-288, (fi6 << ((984) & 63)), (fi6 - (fi6 - 432)), (-(fi6))}
                    _ = v8
                case "hello":
                    v9 := []int{}
                    _ = v9
                }
                v2 -= 0
            } else {
                v3 = (v3 & (v2 << ((4294967295) & 63)))
                if true {
                    var v10 int16 = 1000
                    if false {
                        fmt.Println(v3, v2, v3)
                    } else {
                        var v11 float64 = float64((float32(0.1) / 1.0))
                        v3 = (^(v2))
                        _ = v11
                    }
                    var v12 float64 = 1e10
                    _ = v10
                    _ = v12
                } else {
                    v3 = v2
                    v2 -= 2
                }
            }
        case "":
            fmt.Println(fi6)
        }
        _ = fi6
    }
    delete(v1, "\n\t")
    fmt.Println(v2)
    v13 := GenS1{F0: 119, F1: 65535}
    var v14 int32 = 2147483647
    if true {
        v1["foo bar"] = -108
        fmt.Println(v3)
    } else {
        v15 := v1[("\n\t" + ("foo bar" + "hello"))]
        var v16 int64 = 905
        _ = v15
        _ = v16
    }
    v14 = -929
    v3 += (v2 ^ v3)
    switch "\n\t" {
    case "foo bar":
        var v17 bool = true
        _ = v17
    case "":
        switch "\n\t" {
        case "\n\t":
            if ((-2147483648 > -2147483648) || (!((v2 / ((42) | 1)) != 42))) {
                fmt.Println(v3, v2, v2)
            }
        case "":
            fmt.Println(v2)
        case "foo bar":
            var v18 int16 = 32767
            _ = v18
        default:
            v19 := GenS2{F0: -1, F1: (v3 % (((v3 - 404)) | 1)), F2: 127}
            _ = v19
        }
    case "\n\t":
        v20 := GenS3{F0: 1, F1: v3, F2: 1}
        _ = v20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
}
