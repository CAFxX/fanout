package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 uint64 }
func genFn0(p0 string) int64 {
    v6 := []int{32768, 42, 32767, 2147483647}
    if true {
        v7 := GenS1{F0: 65535, F1: 0, F2: 0}
        fmt.Println(p0)
        _ = v7
    } else {
        v8 := GenS1{F0: 0, F1: 0, F2: 0}
        _ = v8
    }
    _ = p0
    _ = v6
    return 9223372036854775807
}

func main() {
    var v1 int16 = -221
    v2 := make(map[int]bool)
    v3 := make(map[uint32]int)
    v4 := make(map[int]uint32)
    v5 := make(map[string]int)
    v1 = -32768
    switch 42 {
    case -24:
        if (((0.1 / 3.14159) <= 1.5) || false) {
            v1 = v1
            fmt.Println(v1)
        }
    case 94:
        if (2147483647 != 0) {
            for fi9 := 0; fi9 < 17; fi9++ {
                var v10 uint = 255
                v11 := GenS1{F0: 65535, F1: 4294967295, F2: 0}
                _ = fi9
                _ = v10
                _ = v11
            }
            v12 := GenS2{F0: v1}
            var v13 uint16 = 0
            _ = v12
            _ = v13
        } else {
            for fi14 := 0; fi14 < 9; fi14++ {
                v15 := v4[(fi14 + 128)]
                v1 = 1000
                if (999 < v15) {
                    v16 := GenS3{F0: 2, F1: 1}
                    var v17 int64 = int64(v15)
                    _ = v16
                    _ = v17
                }
                _ = fi14
                _ = v15
            }
        }
    case 128:
        for fi18 := 0; fi18 < 2; fi18++ {
            fmt.Println(v1)
            v19 := GenS3{F0: 32767, F1: 1}
            var v20 float64 = 3.14159
            _ = fi18
            _ = v19
            _ = v20
        }
    }
    v1 = (v1 * -162)
    var v21 string = "a"
    var v22 int16 = 1000
    var v23 int = 32767
    if (!false) {
        for fi24 := 0; fi24 < 20; fi24++ {
            var v25 string = "foo bar"
            fmt.Println(fi24, v22)
            fmt.Println(v21, v21, fi24)
            _ = fi24
            _ = v25
        }
    } else {
        v21 = ((((v21 + "foo bar") + (("\n\t" + ("\n\t" + "")) + "")) + "\n\t") + v21)
        v22 = v1
        var v26 int = v23
        _ = v26
    }
    v23 += 669
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
