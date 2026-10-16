package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 uint, p1 uint32) int16 {
    p1 = (^(p1))
    var v6 uint8 = 91
    for fi7 := 0; fi7 < 7; fi7++ {
        fmt.Println(p0, v6)
        _ = fi7
    }
    v8 := []int{42, -2147483648, 9223372036854775807}
    _ = p0
    _ = p1
    _ = v6
    _ = v8
    return -1
}
func genFn1(p0 uint, p1 int8) int64 {
    v9 := GenS1{F0: (p1 | (p1 >> (((p0 - p0)) & 63)))}
    var v10 int16 = 1000
    fmt.Println(p0, p0, v10)
    for fi11 := 0; fi11 < 8; fi11++ {
        switch fi11 {
        case 645:
            switch (fi11 - 42) {
            case 542:
                fmt.Println(v10, p1, p1)
            case 653:
                var v12 int16 = 1000
                _ = v12
            case -411:
                var v13 int = fi11
                _ = v13
            }
        case 1:
            fmt.Println(v10, p0, p1)
        }
        _ = fi11
    }
    switch "" {
    case "a":
        v10 += v10
    case "":
        for fi14 := 0; fi14 < 4; fi14++ {
            p1 = p1
            for fi15 := 0; fi15 < 14; fi15++ {
                v16 := []int{(fi14 * fi15), 1, 1, fi14}
                _ = fi15
                _ = v16
            }
            _ = fi14
        }
    default:
        var v17 uint = p0
        _ = v17
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 239
}

func main() {
    var v1 bool = (!false)
    var v2 uint32 = 167
    var v3 int16 = 1000
    v4 := []uint32{(v2 - (v2 << ((65535) & 63))), v2}
    var v5 string = ("\n\t" + ("" + ("\n\t" + "hello")))
    v3 += v3
    switch -9223372036854775808 {
    case -1:
        var v18 int32 = 621
        _ = v18
    case 2147483647:
        var v19 string = v5
        _ = v19
    case 4294967295:
        var v20 int = 201
        _ = v20
    default:
        if ("\n\t" >= (("" + (v5 + (("" + v5) + v5))) + (("hello" + v5) + v5))) {
            v21 := GenS1{F0: 1}
            v1 = true
            for fi22 := 0; fi22 < 12; fi22++ {
                switch v5 {
                case "foo bar":
                    switch "\n\t" {
                    case "\n\t":
                        v3 = v3
                    default:
                        var v23 float64 = float64(float32(((((1e10 / float64(1.5)) + 0.1) * (3.14159 / -2.25)) * 3.14159)))
                        _ = v23
                    }
                case "\n\t":
                    fmt.Println(v3, v2, v2)
                }
                fmt.Println(v1)
                _ = fi22
            }
            _ = v21
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    }
    switch -933 {
    case 9223372036854775807:
        v24 := GenS1{F0: 52}
        _ = v24
    default:
        v2 = v2
    }
    fmt.Println(v1)
    v25 := GenS1{F0: -128}
    for fi26 := 0; fi26 < 13; fi26++ {
        v3 += -32768
        v5 = v5
        _ = fi26
    }
    _ = len(v4)
    v27 := GenS1{F0: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
