package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 uint64 }
func genFn0(p0 uint16) string {
    var v5 bool = (0.0 == (1e10 - 0.0))
    for fi6 := 0; fi6 < 14; fi6++ {
        p0 = (p0 >> ((214) & 63))
        var v7 uint64 = 2
        _ = fi6
        _ = v7
    }
    v8 := []int{0, 642, -440}
    var v9 int8 = -1
    v5 = v5
    _ = p0
    _ = v5
    _ = v8
    _ = v9
    return "foo bar"
}
func genFn1(p0 uint32, p1 uint8, p2 uint16) uint {
    if (-115 == 100) {
        v10 := GenS1{F0: 42, F1: 42, F2: 32767}
        v11 := GenS1{F0: 0, F1: -1, F2: -1}
        _ = v10
        _ = v11
    }
    switch -935 {
    case 2147483647:
        fmt.Println(p2, p0, p2)
    case 4294967295:
        p1 = 42
    }
    _ = p0
    _ = p1
    _ = p2
    return 18446744073709551615
}

func main() {
    var v1 float32 = (0.0 - ((1.5 + ((float32(0.0) - 1e10) + 1e10)) - 3.14159))
    var v2 int = 698
    var v3 uint = 18446744073709551615
    var v4 bool = false
    var v12 int8 = -104
    switch 1 {
    case 128:
        for fi13 := 0; fi13 < 9; fi13++ {
            v1 = float32((1.5 - (1e10 + 0.0)))
            _ = fi13
        }
    default:
        v14 := GenS2{F0: 9223372036854775807}
        _ = v14
    }
    for fi15 := 0; fi15 < 15; fi15++ {
        fmt.Println(v12)
        _ = fi15
    }
    v2 += 32767
    v16 := []int{v2, -9223372036854775808, v2}
    v1 -= v1
    var v17 float32 = v1
    if true {
        for fi18 := 0; fi18 < 16; fi18++ {
            var v19 uint = v3
            for fi20 := 0; fi20 < 9; fi20++ {
                fmt.Println(fi18, v3)
                switch v2 {
                case -222:
                    var v21 uint64 = 18446744073709551615
                    _ = v21
                case -9223372036854775808:
                    v1 = ((((3.14159 - (-2.25 + (float32(3.14159) / (((((((float32(1.5) / v17) + v1) + v1) + (v17 - (-2.25 - (v1 / 1.0)))) * 1e10) + (float32((float64(v17) * 0.1)) - (1.5 / 3.14159))) + (v17 + (((v1 - float32(0.0)) * v1) * 0.0)))))) / (float32(3.14159) * v17)) + float32(3.14159)) * 3.14159)
                default:
                    switch (^(fi18)) {
                    case 177:
                        v1 = (0.0 - 1.5)
                    case 128:
                        v4 = (2147483647 > -8)
                    case 810:
                        if (416 < 1) {
                            fmt.Println(v19, v19)
                        }
                    }
                }
                _ = fi20
            }
            for fi22 := 0; fi22 < 20; fi22++ {
                fmt.Println(v1, v4, fi22)
                _ = fi22
            }
            _ = fi18
            _ = v19
        }
        if len(v16) >= 2 {
            v16 = v16[1:len(v16)]
        }
    }
    v16 = append(v16, (-(v2)))
    v2 = 4294967295
    var v23 uint32 = 712
    v16 = append(v16, v2)
    v3 += v3
    var v24 int = v2
    var v25 int32 = 2147483647
    if (((!(!v4)) && (!(((("" + "a") + ("\n\t" + "hello")) + "\n\t") < "a"))) && (!("a" <= "hello"))) {
        v23 = (v23 | v23)
    } else {
        v26 := GenS3{F0: 426, F1: 425}
        _ = v26
    }
    v27 := GenS1{F0: 4294967295, F1: -400, F2: 42}
    if (-272 == 513) {
        v17 -= (v17 / v1)
        var v28 uint8 = uint8(v3)
        _ = v28
    } else {
        fmt.Println(v23, v3)
        if ((-494 == 1) && (!v4)) {
            if len(v16) > 0 {
                v16[0] = (^(v2))
            }
        } else {
            fmt.Println(v12, v12)
            var v29 string = ""
            v25 -= (^(v25))
            _ = v29
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v27
    _ = v3
    _ = v4
}
