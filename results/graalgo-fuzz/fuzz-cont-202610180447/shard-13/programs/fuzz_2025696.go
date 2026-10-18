package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32; F1 int }
type GenS5 struct { F0 uint8; F1 int64 }
func genFn0(p0 uint64, p1 int64, p2 uint16) bool {
    p1 -= p1
    v4 := GenS1{F0: -32768}
    var v5 bool = true
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return true
}

func main() {
    v1 := []string{}
    var v2 bool = ("a" != "")
    v3 := []uint32{}
    if (97 != 242) {
        var v6 int16 = -1
        for fi7 := 0; fi7 < 2; fi7++ {
            var v8 int32 = 42
            if (711 != 4294967295) {
                var v9 uint8 = 1
                if true {
                    _ = len(v1)
                } else {
                    v2 = v2
                    var v10 int64 = 0
                    var v11 int32 = -1
                    _ = v10
                    _ = v11
                }
                _ = v9
            }
            for fi12 := 0; fi12 < 20; fi12++ {
                v13 := GenS1{F0: 1000}
                var v14 int = (fi12 >> ((0) & 63))
                fmt.Println(v6, v6)
                _ = fi12
                _ = v13
                _ = v14
            }
            _ = fi7
            _ = v8
        }
        _ = v6
    }
    var v15 uint8 = 2
    if (!(!true)) {
        fmt.Println(v2)
    }
    v16 := GenS2{F0: "a"}
    for fi17 := 0; fi17 < 8; fi17++ {
        fmt.Println(fi17, v2)
        _ = fi17
    }
    v2 = v2
    if true {
        v18 := GenS3{F0: (v15 % ((v15) | 1))}
        _ = v18
    }
    var v19 int = 32767
    v19 = (v19 | v19)
    if false {
        v2 = v2
        v20 := GenS3{F0: (v15 & v15)}
        var v21 int32 = 812
        _ = v20
        _ = v21
    }
    switch v19 {
    case 32768:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        var v22 uint = 109
        _ = v22
    }
    var v23 int = v19
    v23 += v19
    v24 := GenS4{F0: 996, F1: v19}
    var v25 uint16 = 42
    fmt.Println(v15)
    if v2 {
        v26 := GenS5{F0: 1, F1: -9223372036854775808}
        _ = v26
    }
    v25 += v25
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
