package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8 }
func genFn0(p0 string) int16 {
    if (true && ((3.14159 * float64(0.0)) > (((1e10 * (3.14159 - -2.25)) + ((((0.0 - 3.14159) / 3.14159) + (1.5 * 1.5)) * 3.14159)) * (1.5 - (1.5 / float64(((0.1 - -2.25) - 3.14159))))))) {
        var v4 int8 = 12
        v4 += v4
        v5 := []int{32768}
        _ = v4
        _ = v5
    }
    p0 = "foo bar"
    p0 = ("\n\t" + (((p0 + p0) + p0) + ("" + p0)))
    _ = p0
    return 0
}
func genFn1() uint32 {
    var v6 uint = 829
    fmt.Println(v6)
    switch "\n\t" {
    case "foo bar":
        for fi7 := 0; fi7 < 4; fi7++ {
            for fi8 := 0; fi8 < 19; fi8++ {
                v9 := GenS1{F0: -48, F1: (718 <= 0)}
                var v10 int8 = -128
                v11 := GenS1{F0: (v10 >> ((1) & 63)), F1: true}
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
            v12 := GenS1{F0: -1, F1: (0 <= -112)}
            switch fi7 {
            case -302:
                v13 := GenS2{F0: -1}
                _ = v13
            case 2147483647:
                v14 := GenS1{F0: -56, F1: false}
                _ = v14
            case 0:
                switch fi7 {
                case -923:
                    var v15 uint32 = 42
                    _ = v15
                default:
                    var v16 bool = true
                    _ = v16
                }
            }
            _ = fi7
            _ = v12
        }
    case "hello":
        v17 := []int{864, 9223372036854775807, 619, -9223372036854775808}
        _ = v17
    case "":
        v6 -= v6
    }
    _ = v6
    return 0
}

func main() {
    var v1 int16 = 266
    var v2 bool = true
    v3 := []int{-743, -988}
    fmt.Println(v1)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v18 uint8 = 42
    var v19 uint32 = 243
    v1 = (v1 & v1)
    fmt.Println(v18, v2, v1)
    if len(v3) > 0 {
        v3[0] = 0
    }
    v2 = false
    switch 127 {
    case 32768:
        var v20 uint16 = 0
        _ = v20
    case 42:
        v18 += v18
    case -323:
        for fi21 := 0; fi21 < 18; fi21++ {
            v19 = (v19 % ((genFn1()) | 1))
            if v2 {
                v19 = genFn1()
            } else {
                var v22 string = "a"
                v1 -= (v1 << ((0) & 63))
                v18 = 1
                _ = v22
            }
            _ = fi21
        }
    default:
        for fi23 := 0; fi23 < 1; fi23++ {
            var v24 string = "foo bar"
            if true {
                v18 += (^(v18))
            } else {
                if v2 {
                    v3 = append(v3, (fi23 ^ fi23))
                    v25 := GenS3{F0: -1}
                    _ = v25
                }
            }
            v1 = (v1 % ((-32768) | 1))
            _ = fi23
            _ = v24
        }
    }
    var v26 int16 = (v1 - (-(v1)))
    if len(v3) > 0 {
        v3[0] = -1
    }
    v1 += v26
    var v27 uint16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v26
    _ = v27
    _ = v3
}
