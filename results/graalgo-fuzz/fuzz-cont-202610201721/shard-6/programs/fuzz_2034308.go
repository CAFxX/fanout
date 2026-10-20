package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16 }

func main() {
    v1 := []uint32{2}
    var v2 float32 = 1.5
    var v3 int64 = 266
    var v4 uint32 = 2
    v5 := make(map[int]uint32)
    var v6 float64 = ((0.1 + 3.14159) + -2.25)
    for fi7 := 0; fi7 < 9; fi7++ {
        v4 += (v4 % ((v4) | 1))
        for fi8 := 0; fi8 < 5; fi8++ {
            v9 := GenS1{F0: "\n\t"}
            v6 += 1e10
            var v10 float32 = v2
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    if true {
        v4 -= (v4 & v4)
    } else {
        if false {
            fmt.Println(v3, v6)
            v6 = (float64(v2) / 1.5)
        } else {
            v1 = append(v1, v4)
            v6 = 3.14159
            switch "foo bar" {
            case "\n\t":
                var v11 uint8 = 143
                _ = v11
            case "hello":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case "":
                if (v2 <= (float32((float64(((v2 * float32((float64(((0.1 * ((float32(float64((float32(v6) / v2))) / float32((float64(1.5) + 0.1))) + v2)) + (v2 / 1.5))) + v6))) + v2)) + v6)) * (float32((float64(((v2 + ((1.5 + (((v2 - v2) / v2) - 0.0)) + 0.1)) + v2)) - 3.14159)) - v2))) {
                    v6 = v6
                    var v12 uint8 = 127
                    v6 += 1.5
                    _ = v12
                } else {
                    _ = len(v5)
                }
            }
        }
        v13 := v5[4294967295]
        v14 := GenS1{F0: ("hello" + "")}
        _ = v13
        _ = v14
    }
    v15 := GenS2{F0: -944}
    v3 -= v3
    fmt.Println(v2, v3)
    for fi16 := 0; fi16 < 5; fi16++ {
        v4 = (v4 + (v4 + (v4 | (v4 + (v4 + 0)))))
        _ = fi16
    }
    v2 -= 3.14159
    var v17 int = 1
    switch v17 {
    case 32767:
        v2 = float32(1.5)
    case 128:
        v18 := GenS2{F0: 0}
        _ = v18
    default:
        v19 := GenS1{F0: "\n\t"}
        _ = v19
    }
    for fi20 := 0; fi20 < 3; fi20++ {
        for fi21 := 0; fi21 < 20; fi21++ {
            if true {
                var v22 uint16 = 1
                if len(v1) > 0 {
                    v1[0] = v4
                }
                _ = v22
            } else {
                v1 = append(v1, v4)
                var v23 int8 = -128
                var v24 float64 = v6
                _ = v23
                _ = v24
            }
            v17 = v17
            v1 = append(v1, 65535)
            _ = fi21
        }
        var v25 int16 = 645
        _ = fi20
        _ = v25
    }
    var v26 int16 = int16(v4)
    var v27 uint16 = 1
    switch "" {
    case "":
        v17 += (v17 / (((v17 >> ((1) & 63))) | 1))
    case "hello":
        for fi28 := 0; fi28 < 4; fi28++ {
            switch 1 {
            case -2147483648:
                var v29 bool = true
                _ = v29
            case -589:
                v30 := GenS2{F0: v26}
                _ = v30
            case 42:
                if false {
                    v3 -= (v3 >> ((65535) & 63))
                }
            default:
                var v31 float64 = ((1e10 + 1e10) * (1.5 / (v6 / 1e10)))
                _ = v31
            }
            _ = fi28
        }
    case "foo bar":
        fmt.Println(v4, v27)
    default:
        v2 -= v2
    }
    var v32 uint16 = (^(v27))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
