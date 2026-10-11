package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint; F1 int16; F2 uint16 }
func genFn0(p0 uint32, p1 int64, p2 int32) bool {
    if (true || false) {
        var v6 uint32 = (^(p0))
        p2 = 368
        var v7 int64 = (p1 / ((0) | 1))
        _ = v6
        _ = v7
    }
    var v8 uint64 = 4294967295
    v8 += (v8 >> ((255) & 63))
    for fi9 := 0; fi9 < 13; fi9++ {
        v10 := GenS1{F0: 2, F1: 1}
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return (2 <= 1)
}

func main() {
    var v1 int = -133
    v2 := make(map[int]uint32)
    var v3 uint32 = 2147483648
    v4 := make(map[uint32]string)
    v5 := make(map[int]bool)
    switch 32767 {
    case 127:
        switch "foo bar" {
        case "a":
            if false {
                v2[v1] = v3
                if (128 > 0) {
                    v3 = (v3 ^ 4294967295)
                    fmt.Println(v1, v1)
                }
            } else {
                for fi11 := 0; fi11 < 13; fi11++ {
                    v3 = v3
                    _ = fi11
                }
                if (!(-1 >= 1)) {
                    switch "foo bar" {
                    case "\n\t":
                        v12 := GenS1{F0: 65535, F1: 144}
                        _ = v12
                    case "foo bar":
                        var v13 float64 = ((1.5 * (-2.25 - (1e10 - (0.0 * 1.5)))) * 1.5)
                        _ = v13
                    }
                    var v14 bool = ((v3 << ((0) & 63)) <= (v3 | (v3 / (((v3 ^ v3)) | 1))))
                    _ = len(v5)
                    _ = v14
                } else {
                    var v15 int16 = 1
                    _ = v15
                }
                v1 = -47
            }
        }
    case -9223372036854775808:
        fmt.Println(v3, v1)
    default:
        if false {
            for fi16 := 0; fi16 < 12; fi16++ {
                var v17 uint = 4294967295
                var v18 uint32 = v3
                _ = fi16
                _ = v17
                _ = v18
            }
            var v19 uint16 = 842
            v19 += v19
            _ = v19
        }
    }
    v3 = v3
    var v20 string = "foo bar"
    v21 := v5[v1]
    var v22 uint64 = 544
    v23 := GenS1{F0: 0, F1: 133}
    var v24 float32 = -2.25
    for fi25 := 0; fi25 < 10; fi25++ {
        v26 := []int{(fi25 / ((4294967295) | 1))}
        _ = fi25
        _ = v26
    }
    var v27 uint32 = v3
    var v28 uint32 = v27
    if v21 {
        var v29 uint = 936
        v30 := GenS2{F0: 675, F1: -32768, F2: 255}
        var v31 int = (v1 / ((1) | 1))
        _ = v29
        _ = v30
        _ = v31
    }
    fmt.Println(v27)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
