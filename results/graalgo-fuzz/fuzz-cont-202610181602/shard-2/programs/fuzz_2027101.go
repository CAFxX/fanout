package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint8 }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 uint }
func genFn0(p0 uint32, p1 int16, p2 uint64) uint {
    v5 := []int{}
    if false {
        _ = len(v5)
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v5 = append(v5, 42)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 1
}

func main() {
    var v1 uint16 = 65535
    var v2 int32 = 42
    var v3 int = -282
    v4 := make(map[int]int)
    v6 := GenS1{F0: 255, F1: true, F2: 127}
    for fi7 := 0; fi7 < 20; fi7++ {
        fmt.Println(fi7, v2)
        var v8 int32 = -1
        for fi9 := 0; fi9 < 12; fi9++ {
            v10 := GenS1{F0: 127, F1: false, F2: 0}
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    var v11 int64 = -1
    for fi12 := 0; fi12 < 2; fi12++ {
        v11 = v11
        fmt.Println(v2, v2, fi12)
        v1 = 1
        _ = fi12
    }
    v3 = v3
    delete(v4, -2147483648)
    switch "" {
    case "a":
        fmt.Println(v3)
    case "foo bar":
        fmt.Println(v3, v2)
    case "\n\t":
        if true {
            switch ("" + "") {
            case "hello":
                v13 := GenS1{F0: 127, F1: false, F2: 128}
                _ = v13
            default:
                var v14 float32 = 3.14159
                _ = v14
            }
            var v15 uint64 = 4294967295
            fmt.Println(v2, v15)
            _ = v15
        } else {
            var v16 bool = (255 <= 128)
            v3 = (v3 | int(v11))
            v2 -= (v2 ^ v2)
            _ = v16
        }
    default:
        fmt.Println(v1)
    }
    for fi17 := 0; fi17 < 17; fi17++ {
        if false {
            v3 = int(v1)
            for fi18 := 0; fi18 < 18; fi18++ {
                v19 := GenS2{F0: 4, F1: 127}
                _ = fi18
                _ = v19
            }
        }
        var v20 uint64 = 4294967295
        _ = fi17
        _ = v20
    }
    v21 := GenS3{F0: genFn0(164, 1, 42)}
    v22 := []int{}
    v23 := GenS2{F0: 248, F1: 127}
    v24 := GenS3{F0: 2}
    v25 := GenS3{F0: 18446744073709551615}
    var v26 uint32 = 2147483648
    fmt.Println(v26, v3)
    v2 += v2
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v6
}
