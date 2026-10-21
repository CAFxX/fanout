package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 uint16; F1 string; F2 uint16 }
type GenS3 struct { F0 uint8; F1 int }

func main() {
    v1 := make(map[string]uint32)
    v2 := []uint32{42, 2, 2147483647, 677}
    var v3 uint32 = 2147483648
    v4 := make(map[uint32]int)
    var v5 string = "hello"
    var v6 float64 = 0.0
    switch v5 {
    case "":
        if ("\n\t" <= v5) {
            var v7 uint64 = 2
            v8 := GenS1{F0: 350, F1: 128}
            switch -9223372036854775808 {
            case 2147483647:
                v2 = append(v2, (^(v3)))
            default:
                if true {
                    fmt.Println(v6, v3, v7)
                    var v9 bool = true
                    switch v5 {
                    case "\n\t":
                        v7 = (^(v7))
                    case "hello":
                        fmt.Println(v6, v9)
                    case "a":
                        for fi10 := 0; fi10 < 2; fi10++ {
                            var v11 uint64 = v7
                            _ = fi10
                            _ = v11
                        }
                    default:
                        var v12 uint8 = 146
                        _ = v12
                    }
                    _ = v9
                }
            }
            _ = v7
            _ = v8
        }
    }
    switch 39 {
    case 32767:
        v13 := GenS1{F0: 4294967295, F1: 57}
        _ = v13
    case -2147483648:
        for fi14 := 0; fi14 < 16; fi14++ {
            for fi15 := 0; fi15 < 19; fi15++ {
                v6 -= (v6 - (v6 + (-2.25 / v6)))
                _ = fi15
            }
            var v16 int16 = -32768
            _ = fi14
            _ = v16
        }
    }
    for fi17 := 0; fi17 < 14; fi17++ {
        fmt.Println(v3, v5, fi17)
        _ = fi17
    }
    switch (v5 + "hello") {
    case "\n\t":
        v18 := GenS2{F0: 170, F1: "", F2: 1}
        _ = v18
    }
    v19 := GenS3{F0: 180, F1: 128}
    if len(v2) > 0 {
        v2[0] = (^(v3))
    }
    v3 = v3
    for fi20 := 0; fi20 < 8; fi20++ {
        if false {
            var v21 int64 = 9223372036854775807
            _ = v21
        } else {
            var v22 float64 = (1e10 / v6)
            _ = v22
        }
        _ = fi20
    }
    if false {
        for fi23 := 0; fi23 < 1; fi23++ {
            delete(v4, (v3 & (v3 - (v3 & (v3 + v3)))))
            v6 += -2.25
            v4[v3] = -2147483648
            _ = fi23
        }
        switch (v5 + (v5 + v5)) {
        case "a":
            var v24 string = v5
            _ = v24
        default:
            v3 += v3
        }
        v25 := GenS2{F0: 42, F1: v5, F2: 0}
        _ = v25
    } else {
        v1["foo bar"] = v3
        var v26 int64 = -9223372036854775808
        _ = v26
    }
    v3 = 2
    var v27 int = 128
    var v28 int32 = 1
    var v29 int = 522
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
