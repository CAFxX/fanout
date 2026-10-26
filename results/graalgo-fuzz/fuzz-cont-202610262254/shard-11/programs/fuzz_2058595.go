package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64; F1 uint8 }

func main() {
    var v1 uint = 255
    var v2 float32 = 1e10
    v3 := []string{("\n\t" + "\n\t"), "a", ""}
    var v4 uint32 = 137
    var v5 uint16 = 632
    v6 := make(map[uint32]int)
    v4 -= 314
    if false {
        v2 = v2
        v5 += v5
        var v7 float64 = -2.25
        _ = v7
    }
    switch 914 {
    case 9223372036854775807:
        var v8 int = -9223372036854775808
        _ = v8
    case 42:
        if (42 > 0) {
            var v9 string = (("" + "\n\t") + "\n\t")
            var v10 uint16 = 255
            _ = v10
            _ = v9
        }
    default:
        for fi11 := 0; fi11 < 7; fi11++ {
            var v12 int8 = -1
            _ = fi11
            _ = v12
        }
    }
    v13 := GenS1{F0: 42, F1: -884}
    if false {
        v14 := GenS2{F0: 344}
        if len(v3) > 0 {
            v3[0] = "\n\t"
        }
        v2 = (((0.0 - v2) * v2) / v2)
        _ = v14
    } else {
        v15 := GenS1{F0: -1, F1: 0}
        _ = v15
    }
    v5 = 1
    var v16 int64 = 0
    v17 := GenS3{F0: 2, F1: 127}
    var v18 string = "foo bar"
    v4 = (v4 >> ((v1) & 63))
    for fi19 := 0; fi19 < 5; fi19++ {
        if false {
            var v20 uint8 = uint8(v1)
            v4 += v4
            _ = v20
        }
        v5 = v5
        _ = fi19
    }
    fmt.Println(v1)
    if (((0.1 + 0.0) < 3.14159) || true) {
        switch 9223372036854775807 {
        case -9223372036854775808:
            switch "\n\t" {
            case "hello":
                for fi21 := 0; fi21 < 1; fi21++ {
                    if (65535 == 971) {
                        switch (fi21 & 2147483647) {
                        case 707:
                            var v22 uint8 = 127
                            _ = v22
                        default:
                            var v23 uint32 = v4
                            _ = v23
                        }
                        for fi24 := 0; fi24 < 15; fi24++ {
                            var v25 uint8 = 77
                            switch "\n\t" {
                            case "foo bar":
                                v16 = (^(v16))
                            case "a":
                                var v26 float64 = ((1.5 / 1.5) - 0.1)
                                _ = v26
                            case "":
                                switch fi24 {
                                case -819:
                                    fmt.Println(v2, v2)
                                }
                            }
                            _ = fi24
                            _ = v25
                        }
                    }
                    v3 = append(v3, (v18 + v18))
                    _ = fi21
                }
            }
        }
        v27 := GenS3{F0: 1, F1: 8}
        fmt.Println(v1, v16)
        _ = v27
    } else {
        var v28 string = (v18 + v18)
        v1 -= 792
        _ = v28
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi29 := 0; fi29 < 7; fi29++ {
        for fi30 := 0; fi30 < 11; fi30++ {
            var v31 uint32 = (v4 << (((v1 & (v1 * v1))) & 63))
            _ = fi30
            _ = v31
        }
        v1 = (v1 >> ((v1) & 63))
        _ = fi29
    }
    v1 = v1
    var v32 bool = ((int32(v4) <= 42) && (v5 < v5))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
