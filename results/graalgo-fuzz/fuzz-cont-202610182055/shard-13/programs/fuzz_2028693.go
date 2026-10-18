package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8; F1 int16; F2 uint64 }
func genFn0(p0 uint32, p1 uint) uint32 {
    p0 -= p0
    v5 := GenS1{F0: 319, F1: -2147483648, F2: 42}
    v6 := GenS1{F0: 0, F1: int(p0), F2: -32768}
    var v7 uint64 = 2
    v8 := []int{4294967295, 975, 128}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return p0
}

func main() {
    var v1 uint = 4294967295
    v2 := []string{""}
    var v3 int16 = 423
    var v4 int32 = 1
    var v9 float64 = 1.5
    v1 = 31
    fmt.Println(v9, v3)
    for fi10 := 0; fi10 < 9; fi10++ {
        var v11 float32 = -2.25
        for fi12 := 0; fi12 < 12; fi12++ {
            v13 := GenS1{F0: 42, F1: 42, F2: v3}
            v2 = append(v2, ("foo bar" + "foo bar"))
            if len(v2) > 0 {
                v2[0] = ""
            }
            _ = fi12
            _ = v13
        }
        switch "hello" {
        case "\n\t":
            for fi14 := 0; fi14 < 5; fi14++ {
                var v15 uint16 = 2
                switch (fi10 >> ((255) & 63)) {
                case 128:
                    var v16 uint8 = 128
                    _ = v16
                default:
                    var v17 uint8 = 1
                    _ = v17
                }
                v18 := GenS2{F0: v1}
                _ = fi14
                _ = v15
                _ = v18
            }
        default:
            fmt.Println(v11, fi10)
        }
        _ = fi10
        _ = v11
    }
    for fi19 := 0; fi19 < 8; fi19++ {
        v20 := GenS2{F0: v1}
        _ = fi19
        _ = v20
    }
    v3 += 42
    v3 += (^(v3))
    fmt.Println(v9, v1, v4)
    var v21 uint16 = 42
    var v22 uint8 = 128
    v23 := GenS3{F0: v22, F1: 1, F2: 2}
    switch -349 {
    case 764:
        if true {
            var v24 int = 32767
            v1 = v1
            _ = v24
        } else {
            var v25 int = 127
            _ = v25
        }
    case 9223372036854775807:
        var v26 int16 = 1
        _ = v26
    case -2147483648:
        if len(v2) > 0 {
            v2[0] = ""
        }
    }
    var v27 int = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v27
    _ = v3
    _ = v4
    _ = v9
}
