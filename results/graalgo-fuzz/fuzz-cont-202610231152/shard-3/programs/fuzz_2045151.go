package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 uint; F1 int8; F2 uint16 }

func main() {
    v1 := make(map[int]uint32)
    var v2 int64 = 836
    v3 := []string{}
    v4 := make(map[uint32]bool)
    v5 := []int{42, 128, 1, 9223372036854775807}
    var v6 int = -83
    fmt.Println(v2, v6)
    _ = len(v3)
    var v7 uint64 = 0
    fmt.Println(v2, v2, v6)
    v8 := GenS1{F0: 179, F1: 480}
    switch v6 {
    case -1:
        var v9 int = (v6 >> ((255) & 63))
        _ = v9
    case 893:
        v7 = (v7 & 42)
    case 32768:
        var v10 uint16 = 246
        _ = v10
    default:
        v6 -= 1
    }
    switch 32767 {
    case 127:
        for fi11 := 0; fi11 < 1; fi11++ {
            switch 1 {
            case 32768:
                if (128 >= 201) {
                    var v12 float32 = -2.25
                    v12 = v12
                    var v13 int64 = (v2 - v2)
                    _ = v12
                    _ = v13
                } else {
                    for fi14 := 0; fi14 < 15; fi14++ {
                        v15 := GenS1{F0: 317, F1: 199}
                        var v16 int = fi11
                        var v17 string = ("a" + (("foo bar" + "a") + "foo bar"))
                        _ = fi14
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                }
            case 487:
                v2 -= 42
            }
            _ = fi11
        }
    case 141:
        for fi18 := 0; fi18 < 3; fi18++ {
            fmt.Println(v6, v2, v6)
            _ = fi18
        }
    case 42:
        v6 -= v6
    }
    for fi19 := 0; fi19 < 2; fi19++ {
        v7 += 18446744073709551615
        v4[65535] = true
        for fi20 := 0; fi20 < 13; fi20++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if (-2.25 >= (1.5 + float64(3.14159))) {
                v2 -= v2
                _ = len(v5)
                _ = len(v5)
                delete(v4, 1)
                if (("hello" + "a") >= "") {
                    switch v6 {
                    case 377:
                        v2 -= 0
                    case 128:
                        var v21 int8 = 1
                        _ = v21
                    }
                    var v22 uint32 = 0
                    v6 = 128
                    _ = v22
                }
            }
            _ = fi20
        }
        _ = fi19
    }
    if (true && (!(42 > 877))) {
        var v23 int = (v6 * (-(v6)))
        _ = v23
    }
    v2 -= -558
    var v24 int = 964
    _ = len(v3)
    var v25 uint32 = 2
    if true {
        v3 = append(v3, "a")
    }
    for fi26 := 0; fi26 < 4; fi26++ {
        var v27 float32 = (1e10 + (0.1 + 3.14159))
        _ = fi26
        _ = v27
    }
    for fi28 := 0; fi28 < 12; fi28++ {
        var v29 int32 = -406
        _ = fi28
        _ = v29
    }
    v30 := GenS2{F0: 1, F1: 42, F2: 602}
    var v31 float32 = float32(1e10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
