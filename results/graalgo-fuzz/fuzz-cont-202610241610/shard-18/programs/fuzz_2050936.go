package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 uint8 }
type GenS2 struct { F0 uint; F1 int64; F2 uint32 }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 uint8 = 127
    v2 := make(map[uint32]string)
    var v3 uint8 = 128
    var v4 int8 = -21
    if false {
        fmt.Println(v1)
    } else {
        var v5 int16 = -1
        v6 := []int{2147483647}
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: -386, F1: 32767, F2: 42}
    v8 := []int{int(v3), 423}
    var v9 uint16 = 677
    fmt.Println(v4, v4, v3)
    if (0.1 <= (3.14159 / (0.1 + float64(((1e10 * 0.1) - 0.0))))) {
        v2[823] = "foo bar"
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
    }
    v10 := GenS2{F0: 4294967295, F1: -406, F2: 680}
    switch "" {
    case "\n\t":
        if false {
            for fi11 := 0; fi11 < 19; fi11++ {
                if len(v8) > 0 {
                    v8[0] = 127
                }
                v2[465] = "foo bar"
                v12 := GenS1{F0: 665, F1: 42, F2: (v1 % ((0) | 1))}
                _ = fi11
                _ = v12
            }
            v4 = v4
        } else {
            if true {
                if false {
                    switch "a" {
                    case "":
                        v13 := GenS2{F0: 4294967295, F1: -9223372036854775808, F2: 485}
                        _ = v13
                    default:
                        fmt.Println(v9, v3)
                    }
                    for fi14 := 0; fi14 < 6; fi14++ {
                        v9 = v9
                        for fi15 := 0; fi15 < 7; fi15++ {
                            var v16 int64 = 734
                            v1 += v1
                            fmt.Println(fi14, v4, fi14)
                            _ = fi15
                            _ = v16
                        }
                        if (2 < 255) {
                            _ = len(v2)
                        }
                        _ = fi14
                    }
                    v3 += v1
                } else {
                    v1 += 42
                }
                var v17 uint = 386
                _ = v17
            } else {
                fmt.Println(v3, v1)
                var v18 int8 = 0
                _ = v18
            }
            var v19 int = int(v4)
            _ = v19
        }
    case "hello":
        var v20 float32 = (((1e10 / -2.25) / float32(1e10)) / -2.25)
        _ = v20
    case "":
        for fi21 := 0; fi21 < 5; fi21++ {
            v22 := GenS2{F0: 18446744073709551615, F1: 1, F2: 2}
            _ = fi21
            _ = v22
        }
    default:
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
    }
    if len(v8) > 0 {
        v8[0] = 875
    }
    if (1000 < -1) {
        if ((v9 + v9) == uint16(v3)) {
            v3 -= v3
            v9 -= 32767
            for fi23 := 0; fi23 < 18; fi23++ {
                var v24 uint64 = 682
                _ = fi23
                _ = v24
            }
        }
        v25 := GenS3{F0: v9}
        _ = v25
    }
    if true {
        switch (((("\n\t" + "a") + "") + "") + "a") {
        case "\n\t":
            var v26 uint32 = 2
            _ = v26
        default:
            var v27 uint16 = 42
            _ = v27
        }
        v9 = 150
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
