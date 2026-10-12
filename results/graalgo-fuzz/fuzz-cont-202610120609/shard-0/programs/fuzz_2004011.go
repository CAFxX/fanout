package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int16 }
type GenS2 struct { F0 uint16; F1 uint8; F2 uint64 }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 int64; F1 int }
func genFn0() int8 {
    v7 := GenS1{F0: 1, F1: 128, F2: 0}
    switch "hello" {
    case "a":
        var v8 uint8 = 255
        _ = v8
    default:
        v9 := []int{0}
        _ = v9
    }
    var v10 uint8 = 255
    _ = v10
    _ = v7
    return 120
}
func genFn1() string {
    switch "\n\t" {
    case "":
        var v11 uint16 = 934
        _ = v11
    default:
        for fi12 := 0; fi12 < 10; fi12++ {
            var v13 uint = 255
            _ = fi12
            _ = v13
        }
    }
    var v14 uint16 = 32767
    _ = v14
    return "\n\t"
}

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{705, 2}
    var v3 uint = 552
    var v4 int64 = 787
    v5 := []uint32{2147483648, uint32(v3), 4294967295, 1}
    var v6 uint32 = 2147483647
    var v15 string = (("foo bar" + "a") + "a")
    fmt.Println(v6, v6)
    if false {
        if len(v2) > 0 {
            v2[0] = v6
        }
        var v16 int16 = -677
        var v17 uint32 = v6
        _ = v16
        _ = v17
    } else {
        v18 := GenS1{F0: 128, F1: 743, F2: -878}
        _ = v18
    }
    v2 = append(v2, v6)
    v4 = (v4 / ((v4) | 1))
    if (float32(float64(((1.5 - 3.14159) + 0.0))) < (3.14159 + 0.1)) {
        var v19 int16 = 1000
        if (!false) {
            v20 := GenS2{F0: 1, F1: 2, F2: 42}
            if (3.14159 <= (3.14159 + 0.0)) {
                for fi21 := 0; fi21 < 12; fi21++ {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v22 := GenS3{F0: (v15 + v15), F1: 18446744073709551615}
                    var v23 uint8 = uint8(v3)
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                fmt.Println(v4, v19, v15)
                v24 := GenS1{F0: 127, F1: 42, F2: v19}
                _ = v24
            }
            _ = v20
        } else {
            if false {
                var v25 uint32 = 65535
                _ = v25
            }
        }
        _ = v19
    }
    var v26 float32 = 1e10
    var v27 uint = v3
    var v28 string = v15
    v29 := GenS2{F0: 0, F1: 0, F2: uint64(v27)}
    v15 = "foo bar"
    var v30 string = "\n\t"
    switch 9223372036854775807 {
    case 2147483647:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case 167:
        v31 := GenS4{F0: v4, F1: -83}
        _ = v31
    case -292:
        var v32 uint64 = 1
        _ = v32
    default:
        for fi33 := 0; fi33 < 9; fi33++ {
            v3 -= v3
            _ = fi33
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
