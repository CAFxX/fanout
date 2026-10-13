package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool }
func genFn0(p0 int, p1 string) uint16 {
    for fi7 := 0; fi7 < 17; fi7++ {
        v8 := GenS1{F0: -567, F1: (3.14159 > 0.1)}
        _ = fi7
        _ = v8
    }
    switch p1 {
    case "\n\t":
        p0 = p0
    case "":
        v9 := GenS2{F0: 1}
        _ = v9
    case "a":
        v10 := GenS1{F0: 462, F1: (255 != 790)}
        _ = v10
    default:
        if (!false) {
            var v11 int = p0
            var v12 int16 = -32768
            switch p1 {
            case "a":
                var v13 int64 = 1
                _ = v13
            case "\n\t":
                for fi14 := 0; fi14 < 14; fi14++ {
                    var v15 int = v11
                    var v16 int8 = -51
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            case "":
                switch "foo bar" {
                case "foo bar":
                    v12 -= v12
                case "hello":
                    var v17 string = "foo bar"
                    _ = v17
                case "":
                    for fi18 := 0; fi18 < 16; fi18++ {
                        p0 += p0
                        fmt.Println(fi18, v11)
                        v12 -= v12
                        _ = fi18
                    }
                }
            default:
                for fi19 := 0; fi19 < 13; fi19++ {
                    v20 := []int{0, p0, (^(fi19)), (p0 ^ -319)}
                    _ = fi19
                    _ = v20
                }
            }
            _ = v11
            _ = v12
        }
    }
    v21 := GenS2{F0: -2147483648}
    v22 := GenS3{F0: p0}
    if (967 < 42) {
        switch p0 {
        case 4294967295:
            p0 += (^(p0))
        }
    }
    _ = p0
    _ = p1
    _ = v21
    _ = v22
    return 295
}
func genFn1(p0 uint, p1 uint8, p2 uint8) int32 {
    fmt.Println(p1, p2)
    v23 := GenS1{F0: 42, F1: true}
    p2 = p1
    var v24 uint64 = 231
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    _ = v24
    return 2147483647
}

func main() {
    v1 := []uint32{}
    var v2 uint = 42
    var v3 uint32 = 122
    v4 := make(map[string]string)
    var v5 bool = (255 >= 255)
    v6 := make(map[string]int)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    _ = len(v6)
    switch -2147483648 {
    case 576:
        fmt.Println(v2)
    case 128:
        fmt.Println(v5, v5, v2)
    default:
        var v25 int16 = 1000
        _ = v25
    }
    v26 := GenS3{F0: -2147483648}
    v27 := GenS2{F0: genFn1((v2 ^ (v2 + (^(v2)))), 255, 192)}
    v28 := GenS4{F0: ((127 < -50) && (((63 < 94) || ((!(!((true && (1.5 == 3.14159)) && true))) || false)) || v5))}
    if len(v1) > 0 {
        v1[0] = v3
    }
    switch 1 {
    case 127:
        fmt.Println(v3, v2, v2)
    }
    fmt.Println(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
