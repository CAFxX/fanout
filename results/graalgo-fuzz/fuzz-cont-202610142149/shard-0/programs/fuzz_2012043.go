package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 int64; F1 uint16; F2 uint32 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]uint32)
    v3 := []string{("\n\t" + ("\n\t" + "hello")), "", ("" + (("" + "\n\t") + "hello"))}
    switch -604 {
    case 376:
        v4 := GenS1{F0: "a", F1: -1}
        _ = v4
    default:
        v5 := GenS1{F0: "hello", F1: 127}
        _ = v5
    }
    switch "" {
    case "\n\t":
        var v6 int64 = 42
        _ = v6
    }
    v3 = append(v3, ("" + "\n\t"))
    v3 = append(v3, "\n\t")
    var v7 uint64 = 300
    if false {
        if ((2 == 255) || ((!false) || (false && (false && false)))) {
            if (true && true) {
                for fi8 := 0; fi8 < 5; fi8++ {
                    var v9 int32 = 2147483647
                    _ = fi8
                    _ = v9
                }
                var v10 int = -276
                _ = v10
            }
            v7 += 0
            var v11 int16 = 628
            _ = v11
        }
    } else {
        v12 := GenS2{F0: 0, F1: 0}
        v7 -= v7
        v7 = 0
        _ = v12
    }
    for fi13 := 0; fi13 < 6; fi13++ {
        var v14 uint8 = 255
        _ = fi13
        _ = v14
    }
    var v15 uint64 = (v7 - (^(v7)))
    var v16 int8 = 102
    for fi17 := 0; fi17 < 6; fi17++ {
        v18 := GenS2{F0: (^(v16)), F1: -1}
        _ = fi17
        _ = v18
    }
    v19 := GenS1{F0: "a", F1: -128}
    for fi20 := 0; fi20 < 15; fi20++ {
        var v21 int = (^(fi20))
        v22 := GenS3{F0: 0, F1: 507, F2: 796}
        _ = fi20
        _ = v21
        _ = v22
    }
    var v23 int32 = -1
    var v24 float32 = 0.1
    _ = len(v3)
    v25 := GenS3{F0: 9223372036854775807, F1: 42, F2: 42}
    var v26 string = (("a" + "") + "hello")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v7
}
