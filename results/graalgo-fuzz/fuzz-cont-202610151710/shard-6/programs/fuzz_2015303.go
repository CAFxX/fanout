package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int32 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 bool }
func genFn0(p0 uint8) string {
    switch 42 {
    case 258:
        p0 += p0
    case -785:
        fmt.Println(p0)
    case -2147483648:
        switch "" {
        case "hello":
            v4 := GenS1{F0: 40, F1: 9223372036854775807, F2: 2147483647}
            _ = v4
        default:
            if true {
                v5 := []int{2147483647, 128}
                _ = v5
            } else {
                fmt.Println(p0)
                if false {
                    fmt.Println(p0)
                    var v6 uint64 = 18446744073709551615
                    _ = v6
                } else {
                    p0 = 128
                    for fi7 := 0; fi7 < 19; fi7++ {
                        if (fi7 == (fi7 / ((fi7) | 1))) {
                            fmt.Println(fi7, p0)
                            var v8 uint32 = 2
                            var v9 int16 = 0
                            _ = v8
                            _ = v9
                        }
                        v10 := GenS1{F0: fi7, F1: 42, F2: 1}
                        _ = fi7
                        _ = v10
                    }
                    for fi11 := 0; fi11 < 14; fi11++ {
                        v12 := GenS1{F0: 2147483647, F1: 933, F2: -862}
                        _ = fi11
                        _ = v12
                    }
                }
            }
        }
    }
    v13 := []int{}
    var v14 int = 42
    _ = p0
    _ = v13
    _ = v14
    return ""
}

func main() {
    var v1 uint64 = 909
    var v2 uint = 957
    var v3 float64 = (1.5 * 0.0)
    var v15 uint8 = 234
    v2 = (^(v2))
    var v16 uint16 = 32767
    var v17 uint16 = 42
    v18 := GenS2{F0: (v15 & v15), F1: -1}
    if (false && false) {
        var v19 uint8 = 42
        v20 := GenS3{F0: true}
        v1 = (v1 & v1)
        _ = v19
        _ = v20
    } else {
        for fi21 := 0; fi21 < 20; fi21++ {
            v22 := []int{(fi21 - (fi21 + fi21))}
            _ = fi21
            _ = v22
        }
        v3 = -2.25
        fmt.Println(v16, v17, v16)
    }
    for fi23 := 0; fi23 < 15; fi23++ {
        if (v15 <= 0) {
            var v24 string = "\n\t"
            _ = v24
        }
        _ = fi23
    }
    if (-750 != 618) {
        switch ("hello" + "foo bar") {
        case "hello":
            v17 = (v17 | v16)
        case "":
            if (255 != v15) {
                var v25 uint = (v2 - 255)
                var v26 int64 = 42
                _ = v25
                _ = v26
            }
        case "a":
            v2 = uint(v15)
        }
    }
    v15 = (^(v15))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
