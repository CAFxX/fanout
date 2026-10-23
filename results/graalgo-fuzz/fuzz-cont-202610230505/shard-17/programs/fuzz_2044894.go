package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 uint8; F1 string; F2 int32 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]bool)
    v3 := []int{4294967295, -2147483648, 9223372036854775807}
    var v4 int16 = 42
    fmt.Println(v4)
    v4 += 735
    v4 = 500
    v5 := GenS1{F0: 100}
    delete(v1, 330)
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 uint32 = 2147483648
        v8 := v1[0]
        switch fi6 {
        case 0:
            var v9 float64 = 3.14159
            _ = v9
        case -2147483648:
            v10 := GenS1{F0: 42}
            _ = v10
        default:
            v8 = false
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    if false {
        v3 = append(v3, -2147483648)
        for fi11 := 0; fi11 < 6; fi11++ {
            v12 := GenS2{F0: false, F1: 1}
            for fi13 := 0; fi13 < 3; fi13++ {
                v14 := v2[(fi13 | 2147483647)]
                var v15 uint = 2
                _ = fi13
                _ = v14
                _ = v15
            }
            var v16 int32 = 195
            _ = fi11
            _ = v12
            _ = v16
        }
    }
    for fi17 := 0; fi17 < 18; fi17++ {
        v4 = (v4 ^ v4)
        v18 := GenS2{F0: true, F1: 4294967295}
        _ = fi17
        _ = v18
    }
    fmt.Println(v4)
    if false {
        v4 -= v4
        var v19 int32 = 1
        switch 175 {
        case 127:
            v20 := v1[2147483648]
            _ = v20
        default:
            switch 2147483647 {
            case 127:
                v21 := GenS1{F0: 127}
                _ = v21
            case 128:
                var v22 string = "\n\t"
                _ = v22
            case 1:
                var v23 int64 = 42
                _ = v23
            default:
                for fi24 := 0; fi24 < 1; fi24++ {
                    for fi25 := 0; fi25 < 11; fi25++ {
                        var v26 bool = true
                        _ = fi25
                        _ = v26
                    }
                    v4 = (v4 >> ((0) & 63))
                    v4 = 707
                    _ = fi24
                }
            }
        }
        _ = v19
    }
    if true {
        switch "foo bar" {
        case "a":
            for fi27 := 0; fi27 < 14; fi27++ {
                if ((86 > 255) || (true || (fi27 <= (-(fi27))))) {
                    v4 += v4
                    v28 := GenS3{F0: 1, F1: "", F2: 2147483647}
                    _ = v28
                }
                _ = fi27
            }
        default:
            var v29 uint = 930
            _ = v29
        }
        fmt.Println(v4)
    }
    fmt.Println(v4)
    v30 := GenS2{F0: false, F1: 397}
    var v31 uint32 = 542
    fmt.Println(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
