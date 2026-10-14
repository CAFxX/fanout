package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint16 }
func genFn0(p0 int32, p1 int16) int {
    var v7 bool = false
    for fi8 := 0; fi8 < 15; fi8++ {
        p1 += p1
        v9 := GenS1{F0: 731, F1: p0, F2: 42}
        _ = fi8
        _ = v9
    }
    var v10 uint16 = 42
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    return 127
}

func main() {
    var v1 uint64 = 1
    var v2 int64 = 42
    v3 := make(map[uint32]bool)
    var v4 string = "hello"
    v5 := []string{}
    v6 := []string{"a", (v4 + v4), v4, v4}
    fmt.Println(v2)
    switch (v4 + v4) {
    case "foo bar":
        var v11 float64 = 1.5
        _ = v11
    case "a":
        switch genFn0(556, 301) {
        case -9223372036854775808:
            v1 -= (v1 ^ v1)
        case 9223372036854775807:
            var v12 uint32 = 2147483647
            _ = v12
        case 896:
            _ = len(v5)
        }
    case "\n\t":
        var v13 uint8 = uint8(v1)
        _ = v13
    }
    var v14 bool = true
    var v15 int = 1
    v14 = (!true)
    var v16 uint = 45
    for fi17 := 0; fi17 < 3; fi17++ {
        for fi18 := 0; fi18 < 3; fi18++ {
            if v14 {
                var v19 uint8 = 255
                _ = v19
            }
            var v20 float32 = (1.5 - 0.0)
            v4 = "a"
            _ = fi18
            _ = v20
        }
        _ = fi17
    }
    if v14 {
        v5 = append(v5, v4)
    }
    v2 -= (v2 | (v2 - (v2 % (((v2 - -266)) | 1))))
    switch (("\n\t" + (v4 + v4)) + "hello") {
    case "hello":
        v3[65535] = v14
    default:
        fmt.Println(v14, v1)
    }
    var v21 uint8 = 255
    var v22 int16 = 0
    for fi23 := 0; fi23 < 17; fi23++ {
        v22 -= int16(v21)
        v22 += v22
        switch v4 {
        case "":
            var v24 uint = uint(v21)
            _ = v24
        case "hello":
            if (4294967295 >= 255) {
                switch fi23 {
                case 555:
                    var v25 int64 = v2
                    _ = v25
                default:
                    var v26 string = ""
                    _ = v26
                }
                if (32767 == 476) {
                    var v27 uint = v16
                    fmt.Println(v4, v14)
                    var v28 uint64 = uint64(v16)
                    _ = v27
                    _ = v28
                }
            }
        case "foo bar":
            delete(v3, 42)
        default:
            v29 := GenS1{F0: 0, F1: 461, F2: 42}
            _ = v29
        }
        _ = fi23
    }
    var v30 int16 = v22
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
