package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 uint8 = 127
    v2 := []int{-1, -1}
    var v3 uint16 = 32767
    v4 := make(map[string]uint32)
    var v5 uint64 = 374
    if (1.5 > -2.25) {
        fmt.Println(v5)
    } else {
        var v6 int8 = -123
        v3 += v3
        _ = v6
    }
    v7 := GenS1{F0: 42, F1: 42}
    v5 = (v5 ^ 1)
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 uint64 = 595
        v5 -= (v9 + (v9 - (v5 * v5)))
        _ = fi8
        _ = v9
    }
    var v10 int64 = 9223372036854775807
    for fi11 := 0; fi11 < 11; fi11++ {
        if (true && (v3 > uint16(v1))) {
            var v12 uint = 18446744073709551615
            var v13 float64 = 1e10
            switch (("foo bar" + "foo bar") + "foo bar") {
            case "foo bar":
                if len(v2) > 0 {
                    v2[0] = fi11
                }
            case "":
                var v14 int = 32768
                _ = v14
            case "\n\t":
                var v15 int16 = 1
                _ = v15
            }
            _ = v12
            _ = v13
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if true {
            switch "\n\t" {
            case "":
                for fi16 := 0; fi16 < 18; fi16++ {
                    if (!true) {
                        if true {
                            fmt.Println(v1)
                            v1 += 128
                            var v17 int64 = 42
                            _ = v17
                        }
                        v5 += v5
                    }
                    fmt.Println(v10, v5, fi11)
                    _ = fi16
                }
            case "foo bar":
                v4["a"] = 2147483648
            default:
                _ = len(v4)
            }
        }
        _ = fi11
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi18 := 0; fi18 < 7; fi18++ {
        if false {
            v10 += (^(v10))
            switch -2147483648 {
            case -1:
                for fi19 := 0; fi19 < 19; fi19++ {
                    v1 = uint8(v3)
                    v2 = append(v2, fi18)
                    v4["\n\t"] = 585
                    var v20 float32 = (3.14159 * (0.1 - ((0.1 + 1e10) / 0.1)))
                    _ = fi19
                    _ = v20
                }
            }
        }
        _ = fi18
    }
    var v21 int = 94
    for fi22 := 0; fi22 < 18; fi22++ {
        v23 := GenS2{F0: (^(v10))}
        var v24 int16 = -266
        _ = fi22
        _ = v23
        _ = v24
    }
    if false {
        v25 := GenS1{F0: 42, F1: (-(v10))}
        _ = v25
    } else {
        v26 := GenS2{F0: v10}
        _ = v26
    }
    for fi27 := 0; fi27 < 15; fi27++ {
        v2 = append(v2, v21)
        var v28 float32 = 1e10
        switch "" {
        case "foo bar":
            var v29 uint16 = (^(v3))
            _ = v29
        }
        _ = fi27
        _ = v28
    }
    v1 += (v1 | (v1 << ((0) & 63)))
    v3 = (v3 >> ((176) & 63))
    v30 := GenS3{F0: 2}
    var v31 float64 = 0.1
    var v32 uint = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v7
}
