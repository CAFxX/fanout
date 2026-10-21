package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 int }
func genFn0(p0 string, p1 bool, p2 int64) uint64 {
    var v6 float64 = 0.0
    p1 = p1
    var v7 int = 9223372036854775807
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 int16 = -1
    var v2 int8 = 0
    var v3 uint32 = 2
    var v4 float64 = (1.5 / 3.14159)
    var v5 bool = false
    switch ("\n\t" + "foo bar") {
    case "foo bar":
        if v5 {
            var v8 int64 = 42
            _ = v8
        }
    default:
        if (32768 != 32767) {
            v4 -= 3.14159
        } else {
            var v9 uint64 = 4294967295
            var v10 int8 = int8(v3)
            _ = v10
            _ = v9
        }
    }
    var v11 int64 = 318
    var v12 float64 = 0.0
    if (0 <= 260) {
        for fi13 := 0; fi13 < 18; fi13++ {
            switch (fi13 & (fi13 / ((fi13) | 1))) {
            case 284:
                switch "hello" {
                case "\n\t":
                    fmt.Println(v11)
                default:
                    switch "" {
                    case "":
                        v3 = (^(v3))
                    case "\n\t":
                        var v14 uint64 = 0
                        _ = v14
                    }
                }
            case -355:
                fmt.Println(v11)
            default:
                v15 := []int{fi13, fi13, (-(fi13))}
                _ = v15
            }
            if v5 {
                v4 += v4
            } else {
                for fi16 := 0; fi16 < 18; fi16++ {
                    var v17 uint8 = 1
                    _ = fi16
                    _ = v17
                }
                switch (fi13 % ((4294967295) | 1)) {
                case -619:
                    fmt.Println(v11)
                case 829:
                    if v5 {
                        v18 := []int{32767}
                        v5 = true
                        v19 := GenS1{F0: 42, F1: 177, F2: (fi13 % ((fi13) | 1))}
                        _ = v18
                        _ = v19
                    }
                case 618:
                    var v20 uint16 = 255
                    _ = v20
                }
                if (42 < -148) {
                    v1 -= v1
                    v4 -= -2.25
                    var v21 uint8 = 2
                    _ = v21
                }
            }
            switch (fi13 - fi13) {
            case -789:
                v2 += v2
            }
            _ = fi13
        }
    } else {
        var v22 int32 = -1
        var v23 uint16 = 376
        _ = v22
        _ = v23
    }
    v12 = v4
    for fi24 := 0; fi24 < 8; fi24++ {
        var v25 float32 = -2.25
        for fi26 := 0; fi26 < 7; fi26++ {
            if (v25 >= (((0.0 * v25) * v25) * v25)) {
                var v27 int32 = 0
                var v28 int = fi24
                _ = v27
                _ = v28
            }
            _ = fi26
        }
        v1 = 32767
        _ = fi24
        _ = v25
    }
    var v29 uint32 = (v3 * (v3 >> ((2) & 63)))
    var v30 uint64 = 1
    var v31 string = "hello"
    v32 := []int{-646}
    if len(v32) > 0 {
        v32[0] = 42
    }
    v32 = append(v32, -1)
    for fi33 := 0; fi33 < 18; fi33++ {
        var v34 int8 = -99
        _ = fi33
        _ = v34
    }
    v4 = (1.5 + v12)
    v3 += v29
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
