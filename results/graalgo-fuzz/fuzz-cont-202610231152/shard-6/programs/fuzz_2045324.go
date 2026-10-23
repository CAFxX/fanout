package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 int; F2 int8 }
func genFn0(p0 uint32) int {
    v7 := []int{766, 127, -9223372036854775808, 4294967295}
    p0 -= (p0 - p0)
    fmt.Println(p0)
    v8 := GenS1{F0: 100, F1: 1000, F2: 208}
    _ = p0
    _ = v7
    _ = v8
    return -2147483648
}

func main() {
    var v1 string = (("\n\t" + "a") + (("a" + ("foo bar" + "a")) + "a"))
    var v2 int16 = 1000
    var v3 uint8 = 127
    v4 := []uint32{1, 923, 1}
    v5 := make(map[int]int)
    var v6 uint = 18446744073709551615
    var v9 uint8 = v3
    var v10 int64 = -1
    var v11 int = -91
    for fi12 := 0; fi12 < 14; fi12++ {
        v13 := GenS2{F0: 1}
        _ = fi12
        _ = v13
    }
    if (42 >= -1) {
        var v14 uint8 = v3
        _ = v14
    } else {
        var v15 uint16 = 42
        fmt.Println(v2, v11)
        switch v11 {
        case 32767:
            var v16 bool = (uint64(v6) < 637)
            _ = v16
        case 127:
            for fi17 := 0; fi17 < 13; fi17++ {
                v9 -= (v3 >> ((v6) & 63))
                delete(v5, 127)
                _ = fi17
            }
        case -837:
            var v18 uint16 = v15
            _ = v18
        default:
            switch 42 {
            case 127:
                for fi19 := 0; fi19 < 20; fi19++ {
                    fmt.Println(v10, v1, v10)
                    fmt.Println(v1, v1, fi19)
                    _ = fi19
                }
            case 1:
                v20 := GenS2{F0: 362}
                _ = v20
            case 555:
                var v21 int8 = -60
                _ = v21
            default:
                var v22 float64 = 0.0
                _ = v22
            }
        }
        _ = v15
    }
    v23 := GenS2{F0: uint64(v9)}
    fmt.Println(v10)
    if (v1 >= "") {
        v1 = (((("hello" + v1) + v1) + "hello") + v1)
        var v24 int16 = v2
        _ = v24
    }
    v11 -= genFn0(1)
    switch v1 {
    case "\n\t":
        var v25 bool = (false || (1 <= 701))
        _ = v25
    case "":
        if (127 == 100) {
            fmt.Println(v11, v1)
            v9 = 255
            v26 := GenS2{F0: 1}
            _ = v26
        }
    default:
        v9 += (v3 - (v9 | v3))
    }
    for fi27 := 0; fi27 < 3; fi27++ {
        switch v11 {
        case 0:
            if ((65535 < 763) && false) {
                if ((genFn0(4294967295) > (v11 ^ genFn0(2))) && (!(4294967295 > 2))) {
                    if len(v4) > 0 {
                        v4[0] = 2147483648
                    }
                    var v28 uint8 = 68
                    _ = v28
                }
                var v29 float32 = (0.0 * 0.1)
                _ = v29
            }
        case 32768:
            var v30 float32 = 0.1
            _ = v30
        case 747:
            var v31 uint8 = 127
            _ = v31
        }
        _ = fi27
    }
    for fi32 := 0; fi32 < 12; fi32++ {
        v33 := GenS3{F0: uint32(v6), F1: 32768, F2: 0}
        _ = fi32
        _ = v33
    }
    v1 = ("hello" + v1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
