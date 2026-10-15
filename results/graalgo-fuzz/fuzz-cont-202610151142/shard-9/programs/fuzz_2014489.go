package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 uint }

func main() {
    var v1 int8 = -1
    v2 := []string{("foo bar" + "foo bar"), ("foo bar" + ("hello" + "\n\t")), "foo bar", ""}
    v3 := []uint32{755}
    v4 := make(map[uint32]uint32)
    var v5 int64 = 940
    var v6 bool = false
    var v7 string = "a"
    fmt.Println(v1)
    v8 := GenS1{F0: -576, F1: -178}
    v1 -= (v1 | v1)
    for fi9 := 0; fi9 < 12; fi9++ {
        v5 = (v5 + 1)
        for fi10 := 0; fi10 < 14; fi10++ {
            v7 = v7
            var v11 uint = 2
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    v6 = true
    v1 += 83
    for fi12 := 0; fi12 < 1; fi12++ {
        if v6 {
            var v13 int8 = v1
            for fi14 := 0; fi14 < 16; fi14++ {
                if ((true || true) || true) {
                    v4[65535] = 137
                } else {
                    v15 := GenS2{F0: 42}
                    _ = v15
                }
                _ = fi14
            }
            var v16 int64 = v5
            _ = v13
            _ = v16
        } else {
            var v17 uint64 = 4294967295
            var v18 uint16 = 1
            _ = v17
            _ = v18
        }
        var v19 uint64 = 926
        if v6 {
            for fi20 := 0; fi20 < 2; fi20++ {
                v21 := GenS1{F0: 867, F1: 42}
                _ = fi20
                _ = v21
            }
        }
        _ = fi12
        _ = v19
    }
    if v6 {
        var v22 uint8 = 1
        _ = v22
    } else {
        _ = len(v3)
    }
    v7 = "\n\t"
    fmt.Println(v7)
    for fi23 := 0; fi23 < 8; fi23++ {
        if true {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if ((^(v1)) <= -26) {
                v24 := GenS1{F0: 0, F1: 398}
                if true {
                    switch v7 {
                    case "foo bar":
                        var v25 uint64 = 844
                        _ = v25
                    case "hello":
                        switch "a" {
                        case "\n\t":
                            _ = len(v3)
                        default:
                            var v26 string = (v7 + "\n\t")
                            _ = v26
                        }
                    }
                    var v27 float64 = float64(((((((0.1 + 1e10) * 1e10) * 1.5) + float32(1.5)) + float32(0.0)) * (0.1 - (float32(1e10) - (float32(-2.25) - (1.5 - 3.14159))))))
                    _ = len(v4)
                    _ = v27
                }
                _ = v24
            } else {
                var v28 int64 = v5
                v7 = v7
                _ = v28
            }
        }
        _ = fi23
    }
    fmt.Println(v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
