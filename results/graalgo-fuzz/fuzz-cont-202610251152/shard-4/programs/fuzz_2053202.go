package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 int64; F1 int8; F2 int32 }
func genFn0(p0 int) int8 {
    v5 := []int{p0, 127, p0}
    if false {
        switch p0 {
        case 663:
            switch (p0 ^ 0) {
            case 42:
                var v6 int64 = 604
                _ = v6
            case 2147483647:
                p0 = 2147483647
            case -2147483648:
                fmt.Println(p0)
            default:
                var v7 int8 = 0
                _ = v7
            }
        case -1:
            if true {
                var v8 uint8 = 1
                _ = v8
            } else {
                v9 := GenS1{F0: 2147483647, F1: 42}
                _ = v9
            }
        case 249:
            var v10 int16 = 42
            _ = v10
        default:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        var v11 string = "\n\t"
        var v12 uint = 255
        _ = v11
        _ = v12
    } else {
        var v13 float64 = 0.0
        _ = len(v5)
        fmt.Println(v13)
        _ = v13
    }
    var v14 int16 = -1
    _ = p0
    _ = v14
    _ = v5
    return -128
}

func main() {
    v1 := []uint32{42, 65535}
    var v2 int = -9223372036854775808
    var v3 string = "a"
    v4 := make(map[string]string)
    fmt.Println(v2, v2)
    v15 := GenS1{F0: 2147483647, F1: 42}
    v3 = ("hello" + "\n\t")
    for fi16 := 0; fi16 < 19; fi16++ {
        v3 = v3
        _ = fi16
    }
    v17 := GenS2{F0: 1, F1: 0, F2: 42}
    for fi18 := 0; fi18 < 6; fi18++ {
        for fi19 := 0; fi19 < 16; fi19++ {
            v2 = (fi18 + fi19)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            fmt.Println(fi19, fi18)
            _ = fi19
        }
        _ = fi18
    }
    fmt.Println(v3, v2)
    v2 -= v2
    for fi20 := 0; fi20 < 8; fi20++ {
        if (true || (!true)) {
            _ = len(v1)
            switch "a" {
            case "foo bar":
                v21 := GenS2{F0: -1, F1: 0, F2: 1}
                _ = v21
            case "\n\t":
                var v22 bool = (139 > 219)
                _ = v22
            default:
                v3 = "a"
            }
        } else {
            var v23 bool = (42 < -156)
            _ = v23
        }
        v2 = (fi20 + v2)
        v3 = "foo bar"
        _ = fi20
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
