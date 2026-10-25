package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 uint) int8 {
    var v7 uint16 = 32767
    var v8 string = ""
    _ = p0
    _ = v7
    _ = v8
    return int8(p0)
}

func main() {
    var v1 int64 = -470
    v2 := []string{"\n\t"}
    var v3 int8 = -9
    var v4 bool = (65535 > 42)
    v5 := make(map[int]uint32)
    v6 := make(map[uint32]int)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch "hello" {
    case "hello":
        v9 := GenS1{F0: 1}
        _ = v9
    case "\n\t":
        v3 = (v3 * genFn0(65535))
    }
    var v10 int = -657
    var v11 uint16 = 1
    var v12 uint32 = 2147483647
    switch "\n\t" {
    case "":
        fmt.Println(v3, v11, v4)
    default:
        if (v4 && v4) {
            v3 -= (-(v3))
            v12 = v12
            if v4 {
                v11 = 255
                v12 = 2147483648
            }
        }
    }
    v1 -= 1
    delete(v6, 492)
    var v13 float32 = (0.0 / 1.5)
    if v4 {
        var v14 uint32 = v12
        v12 = v12
        for fi15 := 0; fi15 < 16; fi15++ {
            v10 = (fi15 >> ((371) & 63))
            fmt.Println(v14, v14, v1)
            _ = fi15
        }
        _ = v14
    } else {
        if v4 {
            var v16 float32 = float32(1e10)
            _ = v16
        } else {
            v11 -= v11
            _ = len(v2)
            for fi17 := 0; fi17 < 19; fi17++ {
                var v18 bool = (!((v10 & (fi17 % ((v10) | 1))) >= int(v3)))
                fmt.Println(v13, v12)
                v19 := v6[(v12 ^ 0)]
                _ = fi17
                _ = v18
                _ = v19
            }
        }
        v13 = float32(1.5)
        var v20 int64 = v1
        _ = v20
    }
    switch "\n\t" {
    case "a":
        var v21 uint16 = (v11 << ((255) & 63))
        _ = v21
    case "hello":
        v6[(v12 >> ((2) & 63))] = 127
    }
    if true {
        fmt.Println(v1, v12, v3)
    } else {
        var v22 uint = 649
        fmt.Println(v12, v10, v11)
        _ = v22
    }
    for fi23 := 0; fi23 < 5; fi23++ {
        if (18446744073709551615 < 701) {
            v1 -= (v1 * (v1 << ((0) & 63)))
            v10 += v10
            switch (v10 | fi23) {
            case 209:
                switch "foo bar" {
                case "":
                    for fi24 := 0; fi24 < 5; fi24++ {
                        var v25 int16 = 763
                        _ = fi24
                        _ = v25
                    }
                case "a":
                    v26 := GenS1{F0: 1}
                    _ = v26
                default:
                    fmt.Println(v13, v11)
                }
            }
        }
        v13 -= v13
        v13 = (float32(((-2.25 / 1.5) * (float64(v13) - 0.1))) * v13)
        _ = fi23
    }
    switch ("hello" + "hello") {
    case "a":
        fmt.Println(v3)
    case "\n\t":
        var v27 uint32 = 2
        _ = v27
    default:
        v11 += (v11 / (((^(v11))) | 1))
    }
    if false {
        var v28 uint64 = 2
        _ = v28
    }
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
