package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint }
func genFn0(p0 uint16) uint64 {
    p0 = (^(p0))
    var v6 uint8 = 127
    p0 -= p0
    v7 := GenS1{F0: p0, F1: p0}
    if true {
        switch ("\n\t" + "a") {
        case "foo bar":
            if (!(false && (784 < 0))) {
                if true {
                    v6 = (v6 << ((uint(p0)) & 63))
                } else {
                    v6 = (v6 % (((^(v6))) | 1))
                    p0 += p0
                    var v8 int64 = 41
                    _ = v8
                }
                fmt.Println(v6, p0)
                p0 -= (p0 + p0)
            } else {
                v9 := GenS2{F0: false}
                if false {
                    p0 = (p0 >> ((305) & 63))
                } else {
                    var v10 int64 = 0
                    _ = v10
                }
                _ = v9
            }
        }
    } else {
        var v11 string = "hello"
        _ = v11
    }
    _ = p0
    _ = v6
    _ = v7
    return 18446744073709551615
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int64 = 9223372036854775807
    v3 := make(map[int]bool)
    v4 := make(map[int]uint32)
    v5 := make(map[int]uint32)
    delete(v5, -1)
    if (0 != v2) {
        v12 := GenS3{F0: 1, F1: 42}
        if true {
            var v13 uint = 42
            _ = v13
        }
        v14 := GenS3{F0: 4294967295, F1: -32768}
        _ = v12
        _ = v14
    }
    delete(v1, "\n\t")
    v15 := []int{}
    for fi16 := 0; fi16 < 5; fi16++ {
        var v17 uint16 = 65535
        var v18 int = fi16
        _ = fi16
        _ = v17
        _ = v18
    }
    fmt.Println(v2)
    var v19 float32 = 1e10
    v2 = (v2 & (^(v2)))
    for fi20 := 0; fi20 < 3; fi20++ {
        var v21 int16 = -1
        v22 := GenS3{F0: fi20, F1: (v21 & 1000)}
        switch fi20 {
        case -356:
            for fi23 := 0; fi23 < 16; fi23++ {
                fmt.Println(v19, v2)
                v24 := GenS4{F0: v21}
                _ = fi23
                _ = v24
            }
        case 4294967295:
            var v25 uint32 = 111
            _ = v25
        default:
            v26 := GenS5{F0: 1}
            _ = v26
        }
        _ = fi20
        _ = v21
        _ = v22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
