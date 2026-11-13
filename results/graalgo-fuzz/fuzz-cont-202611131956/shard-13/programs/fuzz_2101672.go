package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 int8; F1 uint32; F2 uint16 }
func genFn0(p0 int32, p1 uint, p2 uint8) int16 {
    switch 9223372036854775807 {
    case 1:
        p2 += (p2 & p2)
    case 127:
        p0 = (p0 ^ p0)
    default:
        v4 := GenS1{F0: 127, F1: 0}
        _ = v4
    }
    p1 += p1
    if ("hello" > "") {
        p1 -= p1
        v5 := GenS1{F0: -1, F1: -234}
        _ = v5
    } else {
        if (42 > uint16(p2)) {
            v6 := []int{32768}
            fmt.Println(p0, p0)
            _ = v6
        } else {
            for fi7 := 0; fi7 < 12; fi7++ {
                for fi8 := 0; fi8 < 10; fi8++ {
                    var v9 int8 = 1
                    v9 -= 0
                    _ = fi8
                    _ = v9
                }
                p0 = (p0 % ((42) | 1))
                p0 += p0
                _ = fi7
            }
            switch (("a" + "") + "a") {
            case "\n\t":
                v10 := []int{383, 9223372036854775807, 128, 473}
                _ = v10
            case "a":
                p1 += (p1 ^ 4294967295)
            default:
                for fi11 := 0; fi11 < 12; fi11++ {
                    p1 = (p1 >> (((p1 - (p1 ^ (p1 >> (((^(p1))) & 63))))) & 63))
                    if (!(!true)) {
                        for fi12 := 0; fi12 < 5; fi12++ {
                            fmt.Println(p1, fi12)
                            p1 += p1
                            _ = fi12
                        }
                    }
                    if true {
                        v13 := []int{fi11}
                        _ = v13
                    }
                    _ = fi11
                }
            }
            v14 := []int{128, 0, 9223372036854775807, -1}
            _ = v14
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    var v1 float32 = 1e10
    var v2 int = 128
    var v3 float32 = 1e10
    switch "" {
    case "a":
        var v15 float64 = ((0.0 - -2.25) / 1e10)
        _ = v15
    }
    fmt.Println(v3)
    for fi16 := 0; fi16 < 14; fi16++ {
        fmt.Println(v1, v2, v2)
        _ = fi16
    }
    switch 9223372036854775807 {
    case 2147483647:
        fmt.Println(v2, v3)
    case 0:
        v1 += 1.5
    case 42:
        v17 := GenS1{F0: -1, F1: 9223372036854775807}
        _ = v17
    default:
        switch (v2 & (v2 & 128)) {
        case -1:
            if false {
                v3 = v1
                switch v2 {
                case 9223372036854775807:
                    v18 := []int{v2, -9223372036854775808}
                    _ = v18
                case 42:
                    var v19 string = ("" + ("foo bar" + ""))
                    _ = v19
                default:
                    v20 := GenS2{F0: -128, F1: 2147483648, F2: 2}
                    _ = v20
                }
                v21 := []int{v2, v2, v2, (v2 * 4294967295)}
                _ = v21
            } else {
                v1 = v1
                fmt.Println(v3, v3, v3)
            }
        }
    }
    v22 := []int{1}
    var v23 float64 = (1e10 - ((float64((float32(3.14159) + (v1 / 1.0))) / 1.0) + 1.5))
    _ = len(v22)
    var v24 int16 = genFn0(-2147483648, 65535, 131)
    var v25 uint8 = 2
    if false {
        for fi26 := 0; fi26 < 14; fi26++ {
            v27 := GenS2{F0: 127, F1: 950, F2: 1}
            fmt.Println(v25, v23)
            _ = fi26
            _ = v27
        }
        v3 = (v3 * (v1 + 1e10))
        v24 = (v24 << ((1) & 63))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
