package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32 }
func genFn0(p0 int16, p1 uint8) int16 {
    v7 := GenS1{F0: true}
    fmt.Println(p0, p0)
    var v8 int16 = (p0 & (p0 << ((2) & 63)))
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return p0
}

func main() {
    var v1 float32 = -2.25
    var v2 uint = 1
    var v3 uint16 = 0
    v4 := []int{42, 111, 94}
    var v5 string = (("\n\t" + "a") + "\n\t")
    var v6 bool = false
    if ((float32(3.14159) == (v1 / -2.25)) || ((128 == 2147483647) || true)) {
        if (-2.25 != 3.14159) {
            for fi9 := 0; fi9 < 8; fi9++ {
                switch (v5 + v5) {
                case "":
                    for fi10 := 0; fi10 < 12; fi10++ {
                        v11 := GenS1{F0: false}
                        _ = fi10
                        _ = v11
                    }
                case "\n\t":
                    fmt.Println(v6, v2)
                case "foo bar":
                    var v12 uint16 = v3
                    _ = v12
                default:
                    v13 := GenS2{F0: -962}
                    _ = v13
                }
                _ = fi9
            }
            v5 = "hello"
            v6 = true
        } else {
            for fi14 := 0; fi14 < 3; fi14++ {
                v3 += (v3 + (v3 + (^(v3))))
                v5 = ("" + v5)
                _ = fi14
            }
            _ = len(v4)
            v2 = v2
        }
        for fi15 := 0; fi15 < 14; fi15++ {
            fmt.Println(fi15, v6)
            _ = fi15
        }
        v1 += 3.14159
    } else {
        fmt.Println(v1)
        switch 702 {
        case -1:
            for fi16 := 0; fi16 < 12; fi16++ {
                v4 = append(v4, fi16)
                v3 -= v3
                v17 := GenS2{F0: 1}
                _ = fi16
                _ = v17
            }
        }
        if len(v4) > 0 {
            v4[0] = 1
        }
    }
    var v18 int32 = -1
    var v19 uint16 = (^(v3))
    var v20 int16 = 1000
    v18 += v18
    v4 = append(v4, -2147483648)
    var v21 uint16 = 251
    v19 -= (v21 * (^(v21)))
    v3 = v21
    v3 = v21
    fmt.Println(v20, v19, v18)
    v3 = (v21 | v21)
    for fi22 := 0; fi22 < 11; fi22++ {
        v1 = (v1 * 0.0)
        v20 += int16(v19)
        var v23 float32 = ((v1 - (v1 / ((v1 / v1) + 1e10))) + (3.14159 + ((3.14159 / -2.25) * (1.5 - v1))))
        _ = fi22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
