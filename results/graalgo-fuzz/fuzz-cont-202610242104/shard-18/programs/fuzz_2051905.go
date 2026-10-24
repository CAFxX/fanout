package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 uint64; F1 int32 }
func genFn0(p0 int, p1 int16) uint8 {
    p0 -= (p0 / (((p0 + 43)) | 1))
    p1 = p1
    var v7 uint = 1
    _ = p0
    _ = p1
    _ = v7
    return 0
}
func genFn1(p0 int64, p1 string) bool {
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 float32 = ((1.5 / 1.5) - (1.5 + ((0.1 + ((1e10 * 1e10) - -2.25)) - 1e10)))
        switch (fi8 * 4294967295) {
        case -71:
            var v10 uint32 = 856
            _ = v10
        case -9223372036854775808:
            v11 := []int{-1, fi8, (-(fi8)), 1}
            _ = v11
        default:
            p0 = (p0 % (((p0 & (p0 ^ p0))) | 1))
        }
        _ = fi8
        _ = v9
    }
    fmt.Println(p1, p0)
    v12 := GenS1{F0: 2}
    p1 = p1
    _ = p0
    _ = p1
    _ = v12
    return false
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []uint32{4294967295, 2, 2}
    var v3 int = 128
    var v4 int32 = 722
    v5 := make(map[int]int)
    var v6 float32 = float32(float64(((float32(float64((-2.25 / 1.5))) * ((3.14159 * -2.25) * (0.0 + 0.0))) * (-2.25 + (1.5 - 1.5)))))
    v6 = v6
    if false {
        var v13 string = ""
        _ = v13
    } else {
        v1[291] = false
    }
    v3 = (^(v3))
    v14 := GenS2{F0: "", F1: 65535}
    var v15 string = ""
    switch (v3 >> ((42) & 63)) {
    case 2147483647:
        var v16 float32 = v6
        _ = v16
    case -1:
        v17 := GenS3{F0: 320, F1: v4}
        _ = v17
    default:
        var v18 float32 = 1.5
        _ = v18
    }
    switch ((v15 + v15) + "foo bar") {
    case "\n\t":
        v6 += 3.14159
    case "foo bar":
        switch "foo bar" {
        case "a":
            switch (v15 + (v15 + "foo bar")) {
            case "a":
                var v19 uint32 = 444
                _ = v19
            }
        case "foo bar":
            var v20 int = v3
            _ = v20
        case "hello":
            for fi21 := 0; fi21 < 14; fi21++ {
                _ = len(v2)
                v22 := GenS3{F0: 0, F1: v4}
                _ = fi21
                _ = v22
            }
        default:
            delete(v1, 2147483648)
        }
    case "":
        var v23 float64 = (1.5 / 0.1)
        _ = v23
    }
    var v24 int64 = 1
    v25 := GenS3{F0: 472, F1: v4}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
