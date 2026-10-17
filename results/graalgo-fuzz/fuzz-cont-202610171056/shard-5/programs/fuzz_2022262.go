package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 int16 }

func main() {
    var v1 float64 = (float64(-2.25) / 1.5)
    v2 := make(map[string]uint32)
    var v3 int = 4294967295
    var v4 int64 = -978
    v1 = float64(((3.14159 + 1e10) + 3.14159))
    var v5 int8 = 100
    if (1 < 42) {
        fmt.Println(v3)
        if true {
            var v6 uint8 = 127
            var v7 uint8 = 0
            _ = v6
            _ = v7
        } else {
            for fi8 := 0; fi8 < 6; fi8++ {
                var v9 float64 = v1
                _ = fi8
                _ = v9
            }
        }
        v10 := []int{(v3 + (v3 / (((v3 + 0)) | 1))), v3, (v3 / (((^(v3))) | 1)), 2147483647}
        _ = v10
    } else {
        v1 += v1
        v5 += 127
    }
    v11 := []int{(v3 | v3)}
    var v12 float64 = float64((0.0 / (((0.1 / ((0.0 - -2.25) * (3.14159 + 1.5))) - (float32(v1) + 3.14159)) - 0.1)))
    for fi13 := 0; fi13 < 20; fi13++ {
        v14 := GenS1{F0: "hello", F1: 0}
        _ = fi13
        _ = v14
    }
    var v15 uint16 = 0
    if (-1 == 42) {
        for fi16 := 0; fi16 < 3; fi16++ {
            if true {
                v15 = (v15 - 352)
            }
            if (439 <= 42) {
                var v17 uint8 = 255
                for fi18 := 0; fi18 < 10; fi18++ {
                    fmt.Println(v3, fi16, fi16)
                    _ = fi18
                }
                _ = v17
            }
            for fi19 := 0; fi19 < 14; fi19++ {
                var v20 int = v3
                v21 := GenS2{F0: -70}
                if false {
                    v11 = append(v11, (^(v20)))
                    var v22 uint8 = 234
                    _ = v22
                } else {
                    var v23 int64 = 42
                    v11 = append(v11, (v20 >> ((18446744073709551615) & 63)))
                    if (32767 == v20) {
                        var v24 float32 = 1e10
                        v12 = v12
                        v23 = (v23 & (v4 + v4))
                        _ = v24
                    }
                    _ = len(v11)
                    _ = v23
                }
                _ = fi19
                _ = v20
                _ = v21
            }
            _ = fi16
        }
        switch "\n\t" {
        case "\n\t":
            v2["foo bar"] = 65535
        default:
            fmt.Println(v5)
        }
        switch v3 {
        case 42:
            var v25 uint32 = 399
            _ = v25
        }
    } else {
        var v26 int64 = -1
        _ = v26
    }
    var v27 int16 = -1
    v28 := GenS1{F0: ("foo bar" + "foo bar"), F1: v27}
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
