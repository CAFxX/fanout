package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 int32; F2 uint }
func genFn0(p0 uint) uint8 {
    v7 := GenS1{F0: ""}
    p0 = p0
    for fi8 := 0; fi8 < 1; fi8++ {
        var v9 int8 = 42
        var v10 uint = 255
        v11 := GenS2{F0: ""}
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    var v12 uint = (^(p0))
    _ = p0
    _ = v12
    _ = v7
    return 1
}
func genFn1() string {
    var v13 int64 = 42
    var v14 string = "a"
    for fi15 := 0; fi15 < 15; fi15++ {
        v16 := []int{(fi15 * fi15), (fi15 - fi15), -815, fi15}
        _ = fi15
        _ = v16
    }
    _ = v13
    _ = v14
    return v14
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]string)
    v3 := make(map[int]string)
    v4 := []int{42, 42}
    v5 := []int{288, -1, 42}
    var v6 uint64 = 1
    v6 -= 1
    v6 += (v6 | (v6 * v6))
    for fi17 := 0; fi17 < 10; fi17++ {
        fmt.Println(v6, fi17)
        _ = fi17
    }
    fmt.Println(v6)
    v6 += v6
    v6 = (^(v6))
    v6 = v6
    v18 := GenS2{F0: (("a" + "foo bar") + ("hello" + "foo bar"))}
    var v19 uint64 = 369
    if len(v4) > 0 {
        v4[0] = 9223372036854775807
    }
    switch "\n\t" {
    case "\n\t":
        var v20 uint16 = 65535
        _ = v20
    }
    switch 128 {
    case 128:
        var v21 float64 = -2.25
        _ = v21
    case -636:
        if (1000 > 0) {
            if false {
                v19 = 976
            } else {
                v6 -= v19
            }
            var v22 int8 = 127
            _ = v22
        } else {
            v19 = (v6 >> ((1) & 63))
            v19 = (v19 % (((v19 | (v6 << ((496) & 63)))) | 1))
        }
    }
    fmt.Println(v6, v6)
    v23 := GenS3{F0: 32767, F1: 1, F2: 4294967295}
    var v24 int16 = int16(v19)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
