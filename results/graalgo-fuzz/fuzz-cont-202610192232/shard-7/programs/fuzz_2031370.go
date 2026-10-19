package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 int32 }

func main() {
    var v1 int = 1
    var v2 uint16 = 2
    v3 := make(map[string]bool)
    var v4 int64 = 423
    v5 := make(map[uint32]int)
    v6 := []int{}
    for fi7 := 0; fi7 < 15; fi7++ {
        fmt.Println(v2, v1)
        var v8 uint = 42
        var v9 int32 = 312
        _ = fi7
        _ = v8
        _ = v9
    }
    if (!("" < "foo bar")) {
        v10 := GenS1{F0: 782}
        _ = v10
    }
    switch "" {
    case "hello":
        v2 += 65535
    case "a":
        v6 = append(v6, (v1 % (((v1 % ((32767) | 1))) | 1)))
    }
    if (false || (128 == 127)) {
        for fi11 := 0; fi11 < 20; fi11++ {
            for fi12 := 0; fi12 < 7; fi12++ {
                var v13 bool = ((1 == 1) || false)
                _ = fi12
                _ = v13
            }
            var v14 int = v1
            _ = fi11
            _ = v14
        }
        if false {
            if false {
                v15 := v5[447]
                _ = v15
            }
            var v16 int64 = (v4 - 1)
            _ = v16
        }
        switch v1 {
        case -801:
            v17 := GenS2{F0: v4, F1: -1}
            _ = v17
        default:
            fmt.Println(v4)
        }
    } else {
        v4 += v4
        fmt.Println(v2)
    }
    v6 = append(v6, v1)
    for fi18 := 0; fi18 < 5; fi18++ {
        v4 += v4
        v1 -= v1
        var v19 uint16 = 552
        _ = fi18
        _ = v19
    }
    v4 = v4
    var v20 int32 = 0
    var v21 bool = (!(0.1 != -2.25))
    var v22 float32 = -2.25
    switch "a" {
    case "a":
        if true {
            for fi23 := 0; fi23 < 14; fi23++ {
                fmt.Println(fi23, v2)
                fmt.Println(v4, v20, v21)
                _ = fi23
            }
        }
    case "":
        var v24 bool = (32767 != 32767)
        _ = v24
    case "\n\t":
        var v25 int32 = 42
        _ = v25
    }
    _ = len(v5)
    for fi26 := 0; fi26 < 11; fi26++ {
        fmt.Println(fi26, v2)
        v1 = (-(fi26))
        _ = fi26
    }
    var v27 uint = 437
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
