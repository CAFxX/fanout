package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32; F1 uint32 }

func main() {
    var v1 int8 = 127
    var v2 int16 = 537
    v3 := make(map[int]int)
    if (42 != 2) {
        var v4 string = "foo bar"
        fmt.Println(v1, v4, v4)
        var v5 int8 = (v1 + -44)
        _ = v4
        _ = v5
    } else {
        v6 := []int{4294967295, -2147483648, -1, 9223372036854775807}
        if (1.5 != 0.0) {
            v2 += 680
        } else {
            if true {
                v6 = append(v6, 9223372036854775807)
                v2 -= 106
                v2 = int16(v1)
            } else {
                var v7 bool = (2147483647 > 2)
                _ = v7
            }
            v8 := GenS1{F0: 1}
            _ = v8
        }
        v9 := v3[-1]
        _ = v6
        _ = v9
    }
    fmt.Println(v1, v1)
    v10 := GenS2{F0: 146}
    v11 := GenS3{F0: (("\n\t" + ("" + "hello")) + "foo bar")}
    v1 += 0
    v2 = v2
    if false {
        v12 := GenS3{F0: "\n\t"}
        _ = v12
    } else {
        for fi13 := 0; fi13 < 19; fi13++ {
            for fi14 := 0; fi14 < 20; fi14++ {
                var v15 float32 = ((float32(float64(0.0)) / (1.5 - (0.1 + float32(-2.25)))) + (-2.25 / 1.0))
                v16 := []int{(fi14 << ((657) & 63)), fi14}
                switch fi14 {
                case 9223372036854775807:
                    var v17 string = "\n\t"
                    _ = v17
                }
                _ = fi14
                _ = v15
                _ = v16
            }
            switch (fi13 + (fi13 ^ 32768)) {
            case 32768:
                v3[(fi13 ^ fi13)] = (-(fi13))
            case 1:
                fmt.Println(fi13)
            case -2147483648:
                for fi18 := 0; fi18 < 18; fi18++ {
                    if (1 <= 127) {
                        var v19 int32 = -1
                        _ = v19
                    }
                    var v20 int16 = (v2 / (((v2 - -32768)) | 1))
                    v21 := GenS4{F0: 852, F1: 2}
                    _ = fi18
                    _ = v20
                    _ = v21
                }
            }
            _ = fi13
        }
        var v22 int16 = v2
        var v23 int32 = 938
        _ = v22
        _ = v23
    }
    var v24 int16 = -32768
    v25 := GenS2{F0: 255}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v24
    _ = v25
    _ = v3
}
