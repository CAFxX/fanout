package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 uint8 }
func genFn0(p0 uint16, p1 int8) uint {
    p0 = (p0 << ((2) & 63))
    if (42 > int64(p1)) {
        if (!(0 <= 626)) {
            p0 += 652
            v7 := GenS1{F0: (p1 >> ((918) & 63))}
            _ = v7
        }
        var v8 uint64 = uint64(p0)
        _ = v8
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        if true {
            var v10 string = (("a" + "") + "hello")
            var v11 int64 = 9223372036854775807
            _ = v10
            _ = v11
        } else {
            var v12 int64 = -1
            v13 := []int{128, (fi9 % (((^(fi9))) | 1)), -2147483648}
            for fi14 := 0; fi14 < 9; fi14++ {
                v15 := GenS2{F0: 42}
                _ = fi14
                _ = v15
            }
            _ = v12
            _ = v13
        }
        _ = fi9
    }
    switch -527 {
    case -163:
        fmt.Println(p1, p1)
    case 127:
        p0 += (p0 * (p0 % ((184) | 1)))
    default:
        var v16 uint64 = 0
        _ = v16
    }
    _ = p0
    _ = p1
    return 18446744073709551615
}

func main() {
    v1 := []int{-1, 9223372036854775807}
    var v2 int16 = 358
    var v3 uint64 = 0
    var v4 uint16 = 133
    var v5 bool = (24 == 128)
    var v6 bool = false
    switch "a" {
    case "\n\t":
        for fi17 := 0; fi17 < 1; fi17++ {
            var v18 int = fi17
            var v19 int = 127
            _ = fi17
            _ = v18
            _ = v19
        }
    case "hello":
        switch 32767 {
        case 32767:
            v20 := GenS1{F0: 42}
            _ = v20
        case 4294967295:
            var v21 uint16 = (v4 ^ 255)
            _ = v21
        case -1:
            for fi22 := 0; fi22 < 4; fi22++ {
                var v23 int8 = 127
                _ = fi22
                _ = v23
            }
        }
    case "":
        for fi24 := 0; fi24 < 19; fi24++ {
            v25 := GenS2{F0: 42}
            v3 = v3
            fmt.Println(fi24, fi24)
            _ = fi24
            _ = v25
        }
    default:
        v26 := GenS3{F0: 2147483647, F1: 174}
        _ = v26
    }
    if true {
        var v27 int64 = int64(v4)
        v2 -= (v2 / ((v2) | 1))
        var v28 float32 = 3.14159
        _ = v27
        _ = v28
    }
    v6 = (0.1 == 0.1)
    var v29 uint8 = 42
    v6 = (1e10 == (-2.25 * (0.0 / 1.0)))
    switch "a" {
    case "\n\t":
        if v6 {
            v5 = (!v6)
            if ((true || true) && true) {
                v3 -= v3
            }
            fmt.Println(v2)
        }
    case "hello":
        switch 32767 {
        case 931:
            switch "" {
            case "a":
                v2 -= v2
            }
        case 9223372036854775807:
            for fi30 := 0; fi30 < 16; fi30++ {
                fmt.Println(v29)
                var v31 int = (fi30 & (fi30 >> ((4294967295) & 63)))
                _ = fi30
                _ = v31
            }
        }
    case "a":
        var v32 uint64 = (^(v3))
        _ = v32
    }
    fmt.Println(v3)
    var v33 uint16 = (v4 << ((genFn0(32767, 56)) & 63))
    switch -236 {
    case -9223372036854775808:
        var v34 int64 = -893
        _ = v34
    default:
        fmt.Println(v5)
    }
    v4 = (v4 + (v4 + (^(v33))))
    v1 = append(v1, -174)
    var v35 int = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v33
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
