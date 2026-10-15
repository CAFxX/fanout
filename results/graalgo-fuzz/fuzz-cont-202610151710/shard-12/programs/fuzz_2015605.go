package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 uint32; F1 uint64 }
func genFn0(p0 string) uint8 {
    v7 := GenS1{F0: false, F1: ((127 > 651) || (!(!((float32((3.14159 - (float64(3.14159) * 0.0))) * 1e10) >= (0.0 / 1e10)))))}
    var v8 int32 = 0
    _ = p0
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 int8 = -1
    var v2 bool = false
    var v3 string = (("" + ("foo bar" + "a")) + "\n\t")
    var v4 int64 = 9223372036854775807
    v5 := []uint32{65535, 0}
    var v6 uint64 = 0
    var v9 int8 = -15
    v2 = (2 >= 127)
    _ = len(v5)
    var v10 uint8 = genFn0(v3)
    if v2 {
        var v11 bool = v2
        _ = v11
    }
    var v12 float64 = (((3.14159 / -2.25) - (0.0 / float64(-2.25))) * 1e10)
    var v13 float64 = 3.14159
    if (2 != 4294967295) {
        _ = len(v5)
        for fi14 := 0; fi14 < 15; fi14++ {
            for fi15 := 0; fi15 < 12; fi15++ {
                if (true && v2) {
                    v5 = append(v5, 1)
                } else {
                    switch (fi15 % ((fi14) | 1)) {
                    case 32767:
                        v9 -= v1
                    case -576:
                        var v16 float64 = v12
                        _ = v16
                    }
                    if v2 {
                        var v17 uint16 = 42
                        v12 += v12
                        _ = v17
                    } else {
                        _ = len(v5)
                    }
                }
                _ = fi15
            }
            v4 += (v4 + 0)
            var v18 float32 = 0.0
            _ = fi14
            _ = v18
        }
        v19 := GenS1{F0: (-1 < (v1 << ((4294967295) & 63))), F1: false}
        _ = v19
    }
    v20 := GenS2{F0: 2147483647, F1: 0}
    switch ("\n\t" + (v3 + (v3 + v3))) {
    case "\n\t":
        if v2 {
            for fi21 := 0; fi21 < 9; fi21++ {
                var v22 uint32 = 289
                _ = fi21
                _ = v22
            }
            switch -301 {
            case 448:
                var v23 uint32 = 0
                _ = v23
            case 32767:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            default:
                v24 := GenS2{F0: 4294967295, F1: v6}
                _ = v24
            }
            if len(v5) > 0 {
                v5[0] = 410
            }
        }
    case "":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case "a":
        switch 4294967295 {
        case 443:
            var v25 uint = uint(v10)
            _ = v25
        case -593:
            v13 -= v12
        case 128:
            var v26 int16 = 336
            _ = v26
        }
    }
    var v27 int = 151
    v6 = (v6 * v6)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
