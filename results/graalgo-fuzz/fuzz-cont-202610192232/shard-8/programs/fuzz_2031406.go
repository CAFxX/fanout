package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string }

func main() {
    var v1 float64 = ((0.1 * 1e10) + 3.14159)
    var v2 int16 = 580
    var v3 float32 = ((1.5 + (float32((v1 * (float64(((float32(v1) / (1.5 * 3.14159)) * 1e10)) / (v1 + v1)))) * -2.25)) / float32(v1))
    var v4 uint64 = 179
    var v5 int64 = -1
    if false {
        for fi6 := 0; fi6 < 7; fi6++ {
            v7 := GenS1{F0: 9223372036854775807}
            _ = fi6
            _ = v7
        }
        for fi8 := 0; fi8 < 17; fi8++ {
            for fi9 := 0; fi9 < 14; fi9++ {
                v10 := GenS1{F0: v5}
                fmt.Println(v5, v1)
                _ = fi9
                _ = v10
            }
            v5 += (v5 | (v5 - v5))
            _ = fi8
        }
    } else {
        v3 -= v3
        v11 := []int{127, 9223372036854775807, -643, 672}
        _ = v11
    }
    v1 = -2.25
    v12 := []int{int(v5), -9223372036854775808}
    v13 := GenS2{F0: ("a" + "hello")}
    for fi14 := 0; fi14 < 13; fi14++ {
        var v15 bool = true
        _ = fi14
        _ = v15
    }
    v5 += v5
    var v16 int16 = (v2 ^ v2)
    var v17 uint16 = 65535
    switch 0 {
    case 9223372036854775807:
        if false {
            for fi18 := 0; fi18 < 12; fi18++ {
                v5 = (v5 << ((442) & 63))
                fmt.Println(v2)
                for fi19 := 0; fi19 < 5; fi19++ {
                    switch "a" {
                    case "\n\t":
                        v20 := GenS1{F0: 0}
                        _ = v20
                    default:
                        var v21 int64 = (v5 % ((-1) | 1))
                        _ = v21
                    }
                    v17 = v17
                    _ = fi19
                }
                _ = fi18
            }
            for fi22 := 0; fi22 < 17; fi22++ {
                var v23 float32 = (0.0 + (1.5 / v3))
                fmt.Println(v5, v17, v2)
                v24 := GenS1{F0: v5}
                _ = fi22
                _ = v23
                _ = v24
            }
            v5 += v5
        } else {
            v16 = v16
            for fi25 := 0; fi25 < 3; fi25++ {
                v4 = 42
                v16 -= v16
                _ = fi25
            }
        }
    default:
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
