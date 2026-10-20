package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 uint8 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]bool)
    var v3 uint32 = 4294967295
    v4 := make(map[int]bool)
    var v5 uint64 = 2
    var v6 string = ("a" + (("a" + ("foo bar" + "hello")) + "hello"))
    if ((v3 % ((v3) | 1)) != (v3 % (((v3 + 4294967295)) | 1))) {
        v4[-1] = false
        var v7 int32 = 42
        if (1 >= 172) {
            var v8 int32 = 242
            var v9 string = v6
            _ = v8
            _ = v9
        }
        _ = v7
    } else {
        var v10 string = ((v6 + v6) + v6)
        v10 = "a"
        switch 265 {
        case 787:
            var v11 uint16 = 1
            _ = v11
        case 32767:
            for fi12 := 0; fi12 < 13; fi12++ {
                switch fi12 {
                case 1:
                    var v13 int16 = 1000
                    _ = v13
                case -835:
                    v3 += v3
                }
                _ = fi12
            }
        default:
            v14 := []int{1, -2147483648, 9223372036854775807, -9223372036854775808}
            _ = v14
        }
        _ = v10
    }
    if (!true) {
        v6 = ("a" + "\n\t")
    }
    v6 = ""
    if (32768 >= 128) {
        var v15 uint32 = 1
        switch -1 {
        case 9223372036854775807:
            v5 -= (v5 * v5)
        case 42:
            fmt.Println(v15)
        default:
            var v16 uint32 = (v15 >> ((310) & 63))
            _ = v16
        }
        _ = v15
    } else {
        v5 = (v5 | (v5 - v5))
    }
    v17 := []int{4294967295}
    var v18 int16 = -553
    if (v5 < (v5 * v5)) {
        var v19 string = "\n\t"
        v20 := GenS1{F0: 9223372036854775807, F1: 771, F2: 161}
        v6 = ("" + "\n\t")
        _ = v19
        _ = v20
    } else {
        v2[v6] = true
        v5 += uint64(v3)
        for fi21 := 0; fi21 < 10; fi21++ {
            var v22 int16 = v18
            for fi23 := 0; fi23 < 4; fi23++ {
                switch 128 {
                case -861:
                    var v24 bool = false
                    _ = v24
                case 1:
                    v22 -= v22
                case 128:
                    v3 = (v3 / ((645) | 1))
                    var v25 uint64 = v5
                    v25 += (v25 | 2)
                    _ = v25
                default:
                    var v26 string = v6
                    _ = v26
                }
                _ = fi23
            }
            var v27 uint32 = v3
            _ = fi21
            _ = v22
            _ = v27
        }
    }
    for fi28 := 0; fi28 < 10; fi28++ {
        var v29 float64 = (0.1 - (1e10 / 0.1))
        var v30 int = fi28
        _ = fi28
        _ = v29
        _ = v30
    }
    v17 = append(v17, 1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
