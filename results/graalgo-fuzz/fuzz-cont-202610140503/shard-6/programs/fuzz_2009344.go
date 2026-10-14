package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 uint8 }
func genFn0(p0 uint32, p1 uint16) bool {
    p0 = p0
    p0 = (p0 >> ((1) & 63))
    _ = p0
    _ = p1
    return true
}
func genFn1(p0 bool, p1 int16, p2 uint32) uint64 {
    switch "" {
    case "":
        p2 = 1
    default:
        for fi5 := 0; fi5 < 4; fi5++ {
            p2 -= p2
            if p0 {
                var v6 int = fi5
                v6 = fi5
                _ = v6
            }
            _ = fi5
        }
    }
    v7 := GenS1{F0: 42, F1: 32768, F2: 202}
    var v8 int = 32767
    var v9 uint32 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 1
}

func main() {
    v1 := make(map[string]string)
    var v2 uint32 = 0
    v3 := make(map[int]int)
    var v4 bool = (2 < 18446744073709551615)
    var v10 uint32 = v2
    if (!false) {
        switch "\n\t" {
        case "":
            if ("" == "\n\t") {
                var v11 uint32 = (^(v10))
                var v12 int16 = 676
                _ = v11
                _ = v12
            } else {
                var v13 int = 32768
                var v14 uint32 = v10
                for fi15 := 0; fi15 < 13; fi15++ {
                    v4 = ((1.5 - 0.1) == 0.0)
                    _ = fi15
                }
                _ = v13
                _ = v14
            }
        }
    }
    v16 := []int{}
    fmt.Println(v4, v4, v2)
    v17 := GenS1{F0: 78, F1: -2147483648, F2: 2}
    if (!false) {
        delete(v3, -974)
        fmt.Println(v10, v10, v2)
    }
    if (!(255 != 255)) {
        switch "hello" {
        case "a":
            v2 -= 704
        default:
            v18 := GenS1{F0: 1, F1: 128, F2: 255}
            _ = v18
        }
        var v19 string = ("\n\t" + "hello")
        _ = v19
    } else {
        var v20 int16 = -152
        for fi21 := 0; fi21 < 2; fi21++ {
            var v22 int32 = 580
            fmt.Println(fi21, v22, v22)
            _ = fi21
            _ = v22
        }
        _ = v20
    }
    var v23 string = "a"
    for fi24 := 0; fi24 < 4; fi24++ {
        var v25 uint8 = 0
        _ = fi24
        _ = v25
    }
    var v26 uint64 = 4294967295
    v27 := GenS1{F0: 255, F1: 4294967295, F2: 155}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
