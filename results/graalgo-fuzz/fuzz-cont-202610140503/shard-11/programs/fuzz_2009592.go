package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 int64; F1 bool; F2 string }
func genFn0(p0 int16) string {
    var v4 string = "hello"
    var v5 int64 = -468
    _ = p0
    _ = v4
    _ = v5
    return v4
}

func main() {
    var v1 uint8 = 128
    var v2 uint = 569
    var v3 int = 9223372036854775807
    v6 := []int{4294967295, (v3 + -900)}
    var v7 int = (v3 / ((2147483647) | 1))
    v2 -= 898
    v8 := GenS1{F0: 65535, F1: -1}
    for fi9 := 0; fi9 < 1; fi9++ {
        fmt.Println(fi9)
        for fi10 := 0; fi10 < 7; fi10++ {
            v11 := GenS1{F0: 65535, F1: 0}
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    var v12 int32 = 42
    v1 = v1
    for fi13 := 0; fi13 < 6; fi13++ {
        for fi14 := 0; fi14 < 15; fi14++ {
            if ((!(!true)) || (-51 < 100)) {
                var v15 int64 = -481
                switch -388 {
                case 0:
                    if (-2.25 < (float32(-2.25) * 0.0)) {
                        v12 = v12
                    }
                }
                _ = v15
            }
            switch v7 {
            case 42:
                v3 += fi14
            case 127:
                var v16 uint64 = 0
                var v17 int64 = -1
                _ = v16
                _ = v17
            default:
                switch fi14 {
                case -2147483648:
                    if true {
                        var v18 int32 = (v12 & v12)
                        v6 = append(v6, fi14)
                        var v19 uint16 = 32767
                        _ = v18
                        _ = v19
                    }
                case 4294967295:
                    fmt.Println(v2, v3, fi13)
                    v7 = (fi13 << ((v2) & 63))
                case 32768:
                    v20 := GenS2{F0: 1, F1: false, F2: "foo bar"}
                    _ = v20
                }
            }
            _ = fi14
        }
        _ = fi13
    }
    v3 = -881
    var v21 int16 = int16(v2)
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v21
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
