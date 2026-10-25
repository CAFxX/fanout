package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 uint8 }
func genFn0() uint8 {
    v5 := GenS1{F0: 32767, F1: -27, F2: 2}
    var v6 int32 = 2147483647
    var v7 int64 = 1
    switch "a" {
    case "foo bar":
        v8 := GenS1{F0: 42, F1: 100, F2: 2}
        _ = v8
    case "":
        for fi9 := 0; fi9 < 17; fi9++ {
            for fi10 := 0; fi10 < 19; fi10++ {
                v11 := GenS1{F0: 1000, F1: 18, F2: 127}
                var v12 int8 = 1
                fmt.Println(v12, v12)
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = fi9
        }
    }
    _ = v5
    _ = v6
    _ = v7
    return 255
}
func genFn1(p0 bool, p1 uint16, p2 string) string {
    fmt.Println(p2, p1, p2)
    switch p2 {
    case "hello":
        v13 := []int{703, 748, -795}
        _ = v13
    case "":
        p0 = ((!(-2.25 < 1e10)) && false)
    case "\n\t":
        switch 293 {
        case 32768:
            for fi14 := 0; fi14 < 18; fi14++ {
                fmt.Println(p0, p1, p2)
                for fi15 := 0; fi15 < 19; fi15++ {
                    v16 := GenS1{F0: -1, F1: 100, F2: 255}
                    _ = fi15
                    _ = v16
                }
                _ = fi14
            }
        case 140:
            v17 := GenS1{F0: int16(p1), F1: 127, F2: 0}
            _ = v17
        default:
            for fi18 := 0; fi18 < 19; fi18++ {
                var v19 int = fi18
                var v20 float64 = 0.0
                _ = fi18
                _ = v19
                _ = v20
            }
        }
    default:
        v21 := []int{32768, 223, 2147483647, -1}
        _ = v21
    }
    _ = p0
    _ = p1
    _ = p2
    return "a"
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]string)
    var v3 uint8 = 127
    var v4 uint8 = v3
    v22 := GenS1{F0: -545, F1: 42, F2: v3}
    var v23 uint64 = 417
    v23 -= 18446744073709551615
    v23 += (v23 / (((v23 ^ v23)) | 1))
    if (-1 > -1) {
        var v24 int = 9223372036854775807
        var v25 uint64 = 18446744073709551615
        fmt.Println(v23, v24, v23)
        _ = v24
        _ = v25
    } else {
        v26 := []int{4294967295, 127}
        fmt.Println(v4)
        _ = v26
    }
    var v27 uint32 = 65535
    var v28 bool = true
    fmt.Println(v28, v23, v27)
    var v29 int = 2147483647
    var v30 uint = 0
    v28 = v28
    if (true && (0 < -875)) {
        for fi31 := 0; fi31 < 16; fi31++ {
            fmt.Println(v30, v4)
            fmt.Println(v3, v29, v4)
            _ = fi31
        }
    } else {
        _ = len(v2)
    }
    v32 := []int{v29, v29, 2147483647}
    v29 -= (^(v29))
    if v28 {
        if (-2147483648 < -721) {
            _ = len(v32)
            var v33 int = 9223372036854775807
            _ = v33
        } else {
            var v34 float32 = 1.5
            if ((v30 & (^(v30))) > (v30 % (((v30 - (v30 >> ((0) & 63)))) | 1))) {
                var v35 bool = (!(true && true))
                for fi36 := 0; fi36 < 3; fi36++ {
                    v37 := GenS1{F0: 518, F1: -128, F2: genFn0()}
                    v38 := GenS1{F0: 1, F1: 127, F2: uint8(v27)}
                    _ = fi36
                    _ = v37
                    _ = v38
                }
                _ = v35
            }
            _ = v34
        }
        fmt.Println(v29, v27, v29)
        v29 -= 0
    }
    fmt.Println(v30, v28, v23)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v4
}
