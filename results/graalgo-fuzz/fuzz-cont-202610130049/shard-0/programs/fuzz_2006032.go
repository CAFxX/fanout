package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8 }

func main() {
    var v1 bool = (!((!true) || (true && false)))
    var v2 uint8 = 4
    var v3 string = ""
    var v4 float32 = 0.0
    var v5 uint8 = (v2 & v2)
    v6 := GenS1{F0: 32767}
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 uint32 = 4294967295
        _ = fi7
        _ = v8
    }
    switch 32768 {
    case 4294967295:
        v9 := []int{32767, int(v2)}
        _ = v9
    }
    v4 -= (v4 / (v4 / float32(float64(v4))))
    v10 := []int{}
    for fi11 := 0; fi11 < 14; fi11++ {
        if len(v10) > 0 {
            v10[0] = fi11
        }
        v2 -= (v5 << ((42) & 63))
        v12 := GenS1{F0: 2}
        _ = fi11
        _ = v12
    }
    v3 = v3
    var v13 uint8 = 42
    var v14 uint8 = v2
    v15 := GenS1{F0: 514}
    for fi16 := 0; fi16 < 4; fi16++ {
        v5 = (v13 + (v5 ^ (v13 * 255)))
        for fi17 := 0; fi17 < 16; fi17++ {
            v18 := GenS2{F0: 18446744073709551615}
            if false {
                if len(v10) > 0 {
                    v10[0] = -970
                }
                v19 := GenS3{F0: -79}
                v20 := GenS3{F0: -128}
                _ = v19
                _ = v20
            } else {
                var v21 string = (v3 + v3)
                var v22 int = (fi16 | fi16)
                _ = v21
                _ = v22
            }
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    switch 233 {
    case 571:
        v23 := GenS1{F0: 65535}
        _ = v23
    case -626:
        if (((((float32(-2.25) - (((v4 + 0.0) * float32(float64(-2.25))) + v4)) * v4) / 1e10) / ((((1.5 + 0.0) + (float32(0.0) - v4)) + ((v4 / float32((1.5 - (-2.25 / 1e10)))) - (float32(((0.0 / 0.1) / 1.0)) * (-2.25 * 0.1)))) - 3.14159)) < v4) {
            v14 = (v2 % ((v13) | 1))
        } else {
            fmt.Println(v2, v4)
            var v24 uint32 = 2147483647
            var v25 int64 = 1
            _ = v24
            _ = v25
        }
    case 631:
        if false {
            v26 := GenS3{F0: -1}
            _ = v26
        } else {
            _ = len(v10)
        }
    default:
        if false {
            var v27 string = ""
            v14 = (v13 % (((v2 >> ((42) & 63))) | 1))
            v10 = append(v10, -1)
            _ = v27
        }
    }
    fmt.Println(v13)
    v10 = append(v10, 128)
    v4 += v4
    v4 = ((1e10 + ((v4 / v4) * (v4 - 0.1))) / 1.0)
    v10 = append(v10, -396)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
