package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string; F1 uint8; F2 int16 }

func main() {
    var v1 uint32 = 42
    v2 := []uint32{}
    v3 := []int{-190, 930, -167, 557}
    var v4 uint16 = 1
    var v5 int = -754
    var v6 int = v5
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v5 -= -691
    if true {
        v1 -= v1
        switch 0 {
        case 4294967295:
            var v7 int8 = -1
            _ = v7
        case 32768:
            for fi8 := 0; fi8 < 10; fi8++ {
                for fi9 := 0; fi9 < 4; fi9++ {
                    switch v6 {
                    case 0:
                        if len(v2) > 0 {
                            v2[0] = (^(v1))
                        }
                    default:
                        v2 = append(v2, (v1 % ((v1) | 1)))
                    }
                    _ = fi9
                }
                var v10 float32 = 0.0
                _ = fi8
                _ = v10
            }
        case 42:
            v4 -= uint16(v1)
        }
    } else {
        if len(v2) > 0 {
            v2[0] = 2147483648
        }
    }
    v5 = v6
    v11 := GenS1{F0: 575, F1: true}
    v12 := GenS2{F0: -9223372036854775808}
    for fi13 := 0; fi13 < 14; fi13++ {
        switch v6 {
        case 9223372036854775807:
            v14 := GenS2{F0: -9223372036854775808}
            _ = v14
        case -9223372036854775808:
            v4 += 32767
        case -443:
            for fi15 := 0; fi15 < 19; fi15++ {
                switch 127 {
                case 127:
                    v16 := GenS2{F0: 9223372036854775807}
                    _ = v16
                case -2147483648:
                    var v17 int64 = 940
                    _ = v17
                case 381:
                    v18 := GenS3{F0: "hello", F1: 2, F2: 1}
                    _ = v18
                }
                _ = fi15
            }
        }
        var v19 uint32 = v1
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi13
        _ = v19
    }
    v4 = v4
    v2 = append(v2, (^(v1)))
    fmt.Println(v1)
    v4 -= v4
    var v20 int8 = 127
    fmt.Println(v1, v5)
    for fi21 := 0; fi21 < 1; fi21++ {
        var v22 float32 = -2.25
        _ = fi21
        _ = v22
    }
    v2 = append(v2, (v1 << ((687) & 63)))
    var v23 string = "foo bar"
    v23 = v23
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
