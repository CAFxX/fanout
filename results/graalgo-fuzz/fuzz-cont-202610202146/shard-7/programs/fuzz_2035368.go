package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 string; F1 int64; F2 bool }
type GenS3 struct { F0 uint64; F1 bool; F2 uint64 }
type GenS4 struct { F0 bool; F1 int64; F2 int }
type GenS5 struct { F0 int8; F1 int32; F2 uint }
func genFn0(p0 string) int {
    var v6 int32 = 42
    fmt.Println(v6, p0)
    _ = p0
    _ = v6
    return 1
}

func main() {
    var v1 int8 = -1
    v2 := []uint32{212}
    v3 := []int{0, -1, 128, -9223372036854775808}
    v4 := []string{"", "", "\n\t"}
    v5 := []string{"a", "hello"}
    switch "a" {
    case "hello":
        _ = len(v4)
    case "a":
        switch 1 {
        case -1:
            var v7 float64 = float64(0.0)
            _ = v7
        }
    }
    v1 -= -128
    var v8 uint16 = 255
    v8 -= 42
    v1 += v1
    v9 := GenS1{F0: 127, F1: -873}
    if (42 != (^(v8))) {
        if len(v4) > 0 {
            v4[0] = "foo bar"
        }
    }
    if true {
        v10 := GenS2{F0: "", F1: -9223372036854775808, F2: (127 >= (v1 * (^(v1))))}
        var v11 uint = 42
        var v12 float64 = (0.1 * (1.5 / 1.5))
        _ = v10
        _ = v11
        _ = v12
    } else {
        for fi13 := 0; fi13 < 9; fi13++ {
            var v14 float32 = (float32(0.0) + 1e10)
            var v15 uint8 = 186
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    v16 := GenS3{F0: 2, F1: (255 == 4294967295), F2: 4294967295}
    v17 := GenS4{F0: (-128 < (v1 / ((100) | 1))), F1: -1, F2: 0}
    if false {
        var v18 uint64 = 1
        v2 = append(v2, 1)
        for fi19 := 0; fi19 < 4; fi19++ {
            v20 := GenS5{F0: (v1 - 42), F1: -2147483648, F2: 42}
            _ = fi19
            _ = v20
        }
        _ = v18
    } else {
        v21 := GenS3{F0: 18446744073709551615, F1: ("\n\t" >= ("" + "\n\t")), F2: 0}
        var v22 int64 = -553
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 14; fi23++ {
        switch (fi23 - 558) {
        case 9223372036854775807:
            v24 := GenS5{F0: v1, F1: 0, F2: 2}
            _ = v24
        case 4294967295:
            switch fi23 {
            case 128:
                for fi25 := 0; fi25 < 2; fi25++ {
                    if len(v4) > 0 {
                        v4[0] = ""
                    }
                    v1 = v1
                    v1 = v1
                    _ = fi25
                }
            case 4294967295:
                v1 = (^(v1))
            }
        case -9223372036854775808:
            var v26 int = fi23
            _ = v26
        default:
            v1 = v1
        }
        v2 = append(v2, 2147483647)
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
