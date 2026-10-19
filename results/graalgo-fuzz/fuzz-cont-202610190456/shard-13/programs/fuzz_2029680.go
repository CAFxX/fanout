package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
func genFn0(p0 uint64, p1 uint32, p2 int) int16 {
    var v6 int64 = 1
    p1 = (p1 - (p1 * (p1 & 627)))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 379
}

func main() {
    var v1 uint8 = 0
    v2 := make(map[string]bool)
    v3 := []int{32768, 128}
    var v4 uint32 = 1
    v5 := []int{32768}
    switch -2147483648 {
    case 4294967295:
        fmt.Println(v4, v4)
    case 42:
        v7 := GenS1{F0: false, F1: 42}
        _ = v7
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    switch "hello" {
    case "\n\t":
        for fi8 := 0; fi8 < 20; fi8++ {
            var v9 int = 666
            _ = fi8
            _ = v9
        }
    case "a":
        if (uint16(v1) == 1) {
            for fi10 := 0; fi10 < 5; fi10++ {
                v4 -= v4
                switch "foo bar" {
                case "foo bar":
                    v1 -= (v1 - 216)
                default:
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                v2[("a" + "")] = (fi10 == fi10)
                _ = fi10
            }
            var v11 uint16 = 65535
            _ = v11
        } else {
            var v12 float64 = ((1e10 - 0.1) - ((float64(float32(3.14159)) * -2.25) * -2.25))
            v12 += (float64(float32(float64(3.14159))) - 0.0)
            _ = v12
        }
    case "foo bar":
        var v13 int = 32768
        _ = v13
    default:
        var v14 uint = 65535
        _ = v14
    }
    v1 = v1
    v15 := GenS1{F0: true, F1: 42}
    switch 2147483647 {
    case 32767:
        if false {
            v16 := GenS1{F0: false, F1: -1}
            var v17 uint32 = 2147483647
            _ = v16
            _ = v17
        } else {
            for fi18 := 0; fi18 < 20; fi18++ {
                var v19 uint = 18446744073709551615
                v19 = v19
                _ = fi18
                _ = v19
            }
        }
    default:
        for fi20 := 0; fi20 < 3; fi20++ {
            var v21 bool = (((0 <= 255) || true) && (3.14159 >= 0.1))
            v4 -= v4
            _ = fi20
            _ = v21
        }
    }
    v22 := GenS1{F0: true, F1: 1}
    for fi23 := 0; fi23 < 17; fi23++ {
        var v24 string = ("hello" + ("foo bar" + ""))
        switch "foo bar" {
        case "a":
            var v25 uint32 = v4
            _ = v25
        }
        v24 = v24
        _ = fi23
        _ = v24
    }
    v26 := GenS1{F0: false, F1: 2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
