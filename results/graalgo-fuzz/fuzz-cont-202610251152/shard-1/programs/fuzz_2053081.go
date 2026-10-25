package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0() int8 {
    var v6 bool = false
    switch "" {
    case "":
        if (v6 && ("hello" <= "a")) {
            switch "\n\t" {
            case "foo bar":
                v6 = true
            case "a":
                v7 := []int{32768, -1, -9223372036854775808, 4294967295}
                _ = v7
            default:
                v6 = (!false)
            }
        } else {
            v6 = false
            for fi8 := 0; fi8 < 11; fi8++ {
                v6 = (((1.5 + 1e10) != (((1e10 / 1.0) - 0.0) / (float64(1.5) * ((3.14159 * (-2.25 / -2.25)) - 0.0)))) || ((v6 || v6) && ("hello" > ("" + "foo bar"))))
                var v9 uint32 = 42
                switch fi8 {
                case -9223372036854775808:
                    for fi10 := 0; fi10 < 20; fi10++ {
                        var v11 uint16 = 255
                        v11 -= 255
                        v12 := []int{1, (fi8 >> ((828) & 63)), fi10, 42}
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                case 845:
                    if (!true) {
                        for fi13 := 0; fi13 < 16; fi13++ {
                            fmt.Println(fi13)
                            _ = fi13
                        }
                    }
                case 1:
                    v9 -= (v9 & v9)
                }
                _ = fi8
                _ = v9
            }
        }
    }
    v6 = v6
    fmt.Println(v6)
    for fi14 := 0; fi14 < 2; fi14++ {
        for fi15 := 0; fi15 < 12; fi15++ {
            v6 = true
            switch "a" {
            case "a":
                v6 = ((float32(-2.25) + (3.14159 - 1e10)) <= 0.1)
            }
            fmt.Println(v6, v6)
            _ = fi15
        }
        _ = fi14
    }
    _ = v6
    return 127
}
func genFn1(p0 uint8, p1 uint, p2 uint64) int {
    if (!true) {
        switch 127 {
        case -9223372036854775808:
            for fi16 := 0; fi16 < 2; fi16++ {
                fmt.Println(p2)
                if (1 > 1) {
                    fmt.Println(p1, p0, p2)
                }
                _ = fi16
            }
        }
        var v17 uint16 = 1
        fmt.Println(p0)
        _ = v17
    } else {
        var v18 uint = 4294967295
        v19 := []int{}
        var v20 uint64 = (p2 << ((95) & 63))
        _ = v18
        _ = v19
        _ = v20
    }
    var v21 uint = 802
    p0 = (p0 - 1)
    v22 := GenS1{F0: -285}
    for fi23 := 0; fi23 < 20; fi23++ {
        fmt.Println(p0, fi23, p1)
        _ = fi23
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    _ = v22
    return 32768
}

func main() {
    var v1 int8 = 100
    var v2 int8 = (v1 + (v1 % (((v1 << ((84) & 63))) | 1)))
    var v3 bool = (((-704 > 593) && true) && (1 > -2147483648))
    v4 := []int{}
    var v5 int = 1
    for fi24 := 0; fi24 < 16; fi24++ {
        var v25 int8 = (v1 % ((20) | 1))
        for fi26 := 0; fi26 < 17; fi26++ {
            fmt.Println(v1, v25)
            _ = fi26
        }
        _ = fi24
        _ = v25
    }
    var v27 int64 = 0
    if (505 > 42) {
        v28 := GenS1{F0: 97}
        var v29 float64 = -2.25
        if len(v4) > 0 {
            v4[0] = (v5 ^ 4294967295)
        }
        _ = v28
        _ = v29
    }
    v3 = (-1 != (^(v1)))
    var v30 int64 = (v27 + (^(v27)))
    v31 := GenS1{F0: 0}
    var v32 float32 = ((-2.25 - 3.14159) / 1.0)
    if (float32((0.1 - ((1e10 - 1e10) / (((((1e10 + float64(float32(0.1))) * (float64(3.14159) - 0.0)) * float64((((v32 * v32) + v32) + (float32(0.0) / v32)))) / 0.1) * (((3.14159 * 1.5) / ((1.5 + 3.14159) / (1e10 * (-2.25 / 1e10)))) / -2.25))))) > ((v32 + (v32 + 1.5)) - 0.1)) {
        v5 = v5
    }
    var v33 bool = v3
    v2 += v2
    switch "" {
    case "\n\t":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case "foo bar":
        fmt.Println(v2)
    }
    var v34 uint = 255
    for fi35 := 0; fi35 < 17; fi35++ {
        for fi36 := 0; fi36 < 9; fi36++ {
            if v3 {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            _ = fi36
        }
        v5 = (-(v5))
        v34 = (v34 / ((2) | 1))
        _ = fi35
    }
    var v37 float32 = 3.14159
    if (true || true) {
        var v38 uint8 = 0
        _ = v38
    }
    if (v33 && true) {
        v2 -= genFn0()
    } else {
        v1 = genFn0()
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v37
    _ = v4
    _ = v5
}
