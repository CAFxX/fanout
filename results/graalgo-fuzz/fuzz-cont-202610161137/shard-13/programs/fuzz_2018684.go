package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool }
func genFn0(p0 int32, p1 uint64, p2 uint32) bool {
    var v6 int8 = -61
    var v7 uint16 = 373
    v8 := GenS1{F0: 1, F1: 1000}
    v9 := []int{675, -934}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return (-32768 != -32768)
}

func main() {
    v1 := []uint32{}
    v2 := []int{743, 42, 2147483647}
    var v3 bool = false
    var v4 uint32 = 42
    v5 := make(map[uint32]string)
    for fi10 := 0; fi10 < 3; fi10++ {
        if v3 {
            switch 4294967295 {
            case 42:
                var v11 string = ("\n\t" + "hello")
                _ = v11
            }
            v12 := GenS2{F0: 42}
            _ = v12
        }
        _ = fi10
    }
    switch -758 {
    case 198:
        v13 := GenS3{F0: (v3 && ((false && (true || true)) || v3))}
        _ = v13
    case 1:
        v3 = ((!v3) || v3)
    default:
        v4 = v4
    }
    if v3 {
        switch "foo bar" {
        case "a":
            v2 = append(v2, 9223372036854775807)
        }
        if (-360 != 717) {
            var v14 uint32 = v4
            v3 = (42 >= 0)
            _ = v14
        } else {
            var v15 uint32 = (^(v4))
            var v16 uint64 = 711
            for fi17 := 0; fi17 < 6; fi17++ {
                v18 := GenS1{F0: (fi17 * (fi17 & 247)), F1: -32768}
                v19 := GenS3{F0: (2147483647 >= -868)}
                _ = fi17
                _ = v18
                _ = v19
            }
            _ = v15
            _ = v16
        }
    }
    fmt.Println(v3)
    var v20 uint16 = 394
    switch "a" {
    case "hello":
        v4 += v4
    default:
        switch 364 {
        case 127:
            for fi21 := 0; fi21 < 18; fi21++ {
                v4 = (v4 & (v4 ^ v4))
                fmt.Println(v4)
                var v22 bool = (true || ((0 == 483) || v3))
                _ = fi21
                _ = v22
            }
        }
    }
    switch 128 {
    case -408:
        var v23 int64 = 9223372036854775807
        _ = v23
    default:
        v20 = v20
    }
    fmt.Println(v3, v3, v4)
    for fi24 := 0; fi24 < 16; fi24++ {
        var v25 int64 = 1
        switch 127 {
        case -1:
            v3 = false
        default:
            var v26 string = ("hello" + (((("" + "hello") + "a") + "a") + "a"))
            _ = v26
        }
        var v27 int16 = 1000
        _ = fi24
        _ = v25
        _ = v27
    }
    v4 = uint32(v20)
    v28 := GenS1{F0: 217, F1: -482}
    var v29 bool = (1 != 100)
    v3 = (906 > -1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
