package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 int32 }
type GenS2 struct { F0 bool; F1 uint64 }
func genFn0(p0 int8, p1 uint8, p2 int) int {
    var v5 int32 = -1
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return int(p1)
}

func main() {
    v1 := []int{128, -9223372036854775808}
    var v2 float64 = 1e10
    v3 := make(map[uint32]int)
    var v4 int = 2147483647
    v2 += v2
    var v6 uint = 42
    var v7 int64 = 1
    switch "foo bar" {
    case "\n\t":
        switch ((("" + ("" + "\n\t")) + "\n\t") + "hello") {
        case "a":
            var v8 string = ""
            _ = v8
        default:
            var v9 bool = false
            _ = v9
        }
    }
    v10 := GenS1{F0: 2147483648, F1: (v7 / ((v7) | 1)), F2: -1}
    if (0.1 == 0.0) {
        if true {
            switch "\n\t" {
            case "":
                for fi11 := 0; fi11 < 2; fi11++ {
                    if false {
                        var v12 uint16 = 32767
                        _ = len(v3)
                        var v13 uint64 = 1
                        _ = v12
                        _ = v13
                    } else {
                        v2 += (v2 * 1.5)
                    }
                    _ = fi11
                }
            }
            if false {
                v14 := v3[65535]
                _ = v14
            }
        } else {
            if (v2 > v2) {
                var v15 uint16 = 255
                fmt.Println(v7, v2)
                v16 := GenS1{F0: 297, F1: v7, F2: -1}
                _ = v15
                _ = v16
            } else {
                v17 := GenS2{F0: true, F1: 0}
                _ = v17
            }
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v7, v2)
    v6 = (v6 + v6)
    for fi18 := 0; fi18 < 1; fi18++ {
        var v19 uint32 = 2147483647
        v19 += (v19 * (v19 / ((v19) | 1)))
        fmt.Println(v7)
        _ = fi18
        _ = v19
    }
    if true {
        var v20 uint64 = 0
        var v21 int16 = 1
        fmt.Println(v7, v21, v4)
        _ = v20
        _ = v21
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
