package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 uint16 }
func genFn0(p0 string, p1 uint16, p2 uint16) uint8 {
    fmt.Println(p1, p2)
    fmt.Println(p1)
    var v4 uint32 = 42
    if true {
        p1 += (p1 ^ p2)
        if (42 == 42) {
            switch (("\n\t" + (p0 + p0)) + ((((p0 + "foo bar") + p0) + p0) + (("hello" + (p0 + (p0 + "a"))) + p0))) {
            case "hello":
                var v5 float64 = 1e10
                _ = v5
            case "a":
                p0 = (p0 + (p0 + p0))
            case "":
                v6 := []int{}
                _ = v6
            default:
                p1 = p2
            }
            v7 := GenS1{F0: 42}
            fmt.Println(p0)
            _ = v7
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 96
}

func main() {
    var v1 float64 = 1e10
    var v2 int8 = 1
    v3 := []uint32{4294967295}
    v8 := GenS2{F0: 65535, F1: 0}
    switch "hello" {
    case "foo bar":
        if true {
            var v9 string = ""
            v2 = v2
            fmt.Println(v1, v1)
            _ = v9
        } else {
            if (!(true || true)) {
                var v10 float64 = (0.0 - 0.0)
                for fi11 := 0; fi11 < 13; fi11++ {
                    v10 = 0.0
                    _ = fi11
                }
                _ = v10
            }
        }
    }
    for fi12 := 0; fi12 < 20; fi12++ {
        var v13 int8 = (v2 / (((v2 & v2)) | 1))
        _ = fi12
        _ = v13
    }
    var v14 int32 = 269
    var v15 string = ((("a" + (("\n\t" + "\n\t") + "")) + (((("a" + "foo bar") + ("\n\t" + "")) + "\n\t") + "a")) + ("a" + ((("" + "hello") + ("\n\t" + "foo bar")) + "foo bar")))
    var v16 float64 = (v1 / 3.14159)
    switch -980 {
    case 476:
        var v17 int8 = (v2 >> ((18446744073709551615) & 63))
        _ = v17
    case -780:
        v2 = (v2 ^ -128)
    default:
        var v18 int16 = -484
        _ = v18
    }
    if (950 <= -1) {
        fmt.Println(v1, v14, v14)
        var v19 int8 = (v2 << ((1) & 63))
        var v20 float32 = 3.14159
        _ = v19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v8
}
