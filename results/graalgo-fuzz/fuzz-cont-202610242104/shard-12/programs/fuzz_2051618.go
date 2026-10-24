package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 string }
type GenS2 struct { F0 int; F1 int16; F2 uint32 }
type GenS3 struct { F0 int; F1 string; F2 uint64 }
func genFn0(p0 int64) uint64 {
    var v7 uint8 = 195
    var v8 int8 = -128
    v8 = 109
    v8 += (^(v8))
    _ = p0
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 uint64) uint8 {
    p0 = (p0 << ((110) & 63))
    var v9 int32 = 2147483647
    fmt.Println(v9, v9)
    _ = p0
    _ = v9
    return 1
}

func main() {
    var v1 float64 = (3.14159 - -2.25)
    v2 := make(map[int]string)
    var v3 int16 = 32767
    var v4 float64 = 3.14159
    var v5 float64 = 1e10
    v6 := make(map[string]string)
    v10 := GenS1{F0: 255, F1: v3, F2: ""}
    if (!false) {
        v5 += 1.5
        v5 -= v5
        var v11 string = "foo bar"
        _ = v11
    }
    _ = len(v2)
    if false {
        var v12 float32 = (((1e10 / 3.14159) + 0.0) - 0.1)
        if true {
            switch 32767 {
            case 0:
                v13 := []int{-1, 32767}
                _ = v13
            case 42:
                var v14 int8 = 127
                _ = v14
            default:
                fmt.Println(v1)
            }
            if (2 != 1) {
                var v15 uint8 = 2
                var v16 uint64 = 333
                _ = v15
                _ = v16
            } else {
                v3 = (-(v3))
                v1 = (1.5 * v4)
                var v17 int = 32768
                _ = v17
            }
        }
        for fi18 := 0; fi18 < 13; fi18++ {
            var v19 int8 = -1
            v20 := GenS2{F0: (-(fi18)), F1: (v3 - (v3 << ((65535) & 63))), F2: 804}
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = v12
    }
    _ = len(v6)
    for fi21 := 0; fi21 < 15; fi21++ {
        for fi22 := 0; fi22 < 6; fi22++ {
            var v23 string = ""
            var v24 int32 = 615
            _ = fi22
            _ = v23
            _ = v24
        }
        v25 := GenS3{F0: fi21, F1: "\n\t", F2: genFn0(9223372036854775807)}
        _ = fi21
        _ = v25
    }
    for fi26 := 0; fi26 < 12; fi26++ {
        v4 = v4
        delete(v2, (^(fi26)))
        _ = fi26
    }
    var v27 int = 0
    for fi28 := 0; fi28 < 8; fi28++ {
        fmt.Println(v4)
        _ = fi28
    }
    v29 := GenS3{F0: v27, F1: "hello", F2: 2}
    v3 -= (v3 >> ((18446744073709551615) & 63))
    v5 += (0.0 / 1e10)
    var v30 int8 = -128
    var v31 int = v27
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
