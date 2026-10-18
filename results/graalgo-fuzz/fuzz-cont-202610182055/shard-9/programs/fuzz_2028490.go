package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 uint16; F1 uint32; F2 uint16 }

func main() {
    var v1 float64 = 1.5
    var v2 int = 823
    var v3 float64 = float64(float32(v1))
    v4 := []string{("" + "a")}
    var v5 int64 = 42
    var v6 float32 = (0.0 - (3.14159 - ((0.0 * 0.0) + ((0.0 - (((-2.25 + 1.5) * 1e10) * 3.14159)) / -2.25))))
    var v7 int32 = -1
    switch v2 {
    case -9223372036854775808:
        for fi8 := 0; fi8 < 15; fi8++ {
            var v9 bool = ((310 == 255) || (false || ("foo bar" < "hello")))
            _ = fi8
            _ = v9
        }
    case 4294967295:
        var v10 uint = 28
        _ = v10
    case -634:
        if ((-32768 != 0) && false) {
            for fi11 := 0; fi11 < 5; fi11++ {
                fmt.Println(fi11, v3)
                _ = fi11
            }
            var v12 float32 = 1e10
            _ = v12
        }
    default:
        if len(v4) > 0 {
            v4[0] = "\n\t"
        }
    }
    _ = len(v4)
    v7 = v7
    if (-1 >= -32768) {
        v6 += 1e10
        v1 = v3
        for fi13 := 0; fi13 < 13; fi13++ {
            var v14 int32 = (-(v7))
            var v15 float64 = float64(v6)
            v4 = append(v4, "a")
            _ = fi13
            _ = v14
            _ = v15
        }
    } else {
        for fi16 := 0; fi16 < 15; fi16++ {
            switch (^(v2)) {
            case 163:
                v6 = 0.0
            }
            _ = fi16
        }
        fmt.Println(v5)
        for fi17 := 0; fi17 < 10; fi17++ {
            var v18 uint16 = 796
            _ = fi17
            _ = v18
        }
    }
    switch (-(v2)) {
    case -14:
        var v19 int16 = 609
        _ = v19
    case 128:
        var v20 uint64 = 18446744073709551615
        _ = v20
    case 0:
        v21 := GenS1{F0: false, F1: 4294967295}
        _ = v21
    default:
        for fi22 := 0; fi22 < 14; fi22++ {
            v23 := GenS2{F0: 65535, F1: 723, F2: 2}
            _ = fi22
            _ = v23
        }
    }
    if true {
        switch 1 {
        case 127:
            fmt.Println(v2)
        }
        v3 = v3
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        switch (v2 * v2) {
        case 9223372036854775807:
            for fi24 := 0; fi24 < 16; fi24++ {
                fmt.Println(v1, fi24, v7)
                _ = len(v4)
                var v25 string = "foo bar"
                _ = fi24
                _ = v25
            }
        default:
            for fi26 := 0; fi26 < 12; fi26++ {
                var v27 float32 = v6
                v1 += v1
                fmt.Println(v7)
                _ = fi26
                _ = v27
            }
        }
    }
    var v28 uint64 = 608
    v29 := GenS2{F0: 2, F1: 1, F2: 1}
    v4 = append(v4, "foo bar")
    v3 += 1e10
    var v30 string = ("" + "hello")
    if (float32(1e10) >= ((3.14159 - v6) + v6)) {
        fmt.Println(v3, v28, v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
