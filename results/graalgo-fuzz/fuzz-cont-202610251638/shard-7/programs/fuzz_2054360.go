package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 int64 }
type GenS2 struct { F0 int8 }
func genFn0() uint8 {
    v4 := []int{32767, -9223372036854775808}
    for fi5 := 0; fi5 < 11; fi5++ {
        var v6 float32 = (0.0 - -2.25)
        v6 = (v6 + (float32(1.5) / ((v6 - 1.5) - (v6 + (float32((1e10 * ((3.14159 / -2.25) / 0.1))) - (v6 / v6))))))
        _ = fi5
        _ = v6
    }
    v4 = append(v4, -2147483648)
    for fi7 := 0; fi7 < 8; fi7++ {
        if len(v4) > 0 {
            v4[0] = (fi7 >> ((65535) & 63))
        }
        if ((89 < 1) || (666 < 2147483648)) {
            var v8 uint16 = 32767
            _ = v8
        } else {
            var v9 string = "hello"
            if ((!false) && false) {
                var v10 float32 = 0.1
                fmt.Println(v10)
                _ = v10
            } else {
                v9 = ((v9 + "hello") + (v9 + ((v9 + ("a" + (v9 + (v9 + v9)))) + v9)))
            }
            _ = v9
        }
        var v11 uint = 65535
        _ = fi7
        _ = v11
    }
    var v12 uint16 = 1
    _ = v12
    _ = v4
    return 42
}
func genFn1(p0 int8, p1 int, p2 uint16) uint {
    var v13 uint32 = 825
    v13 = v13
    v14 := []int{p1, p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return 65535
}

func main() {
    var v1 uint8 = 2
    var v2 string = "foo bar"
    var v3 int16 = 627
    v15 := GenS1{F0: ((v3 % ((v3) | 1)) == (v3 + v3)), F1: 979, F2: -1}
    v16 := []int{42, -9223372036854775808, -848, 9223372036854775807}
    v17 := GenS2{F0: -128}
    v1 = 127
    var v18 float64 = (3.14159 / 1.0)
    fmt.Println(v3, v3)
    v19 := GenS2{F0: -68}
    for fi20 := 0; fi20 < 4; fi20++ {
        switch 4294967295 {
        case 1:
            switch v2 {
            case "":
                for fi21 := 0; fi21 < 12; fi21++ {
                    _ = len(v16)
                    v18 = -2.25
                    var v22 int16 = v3
                    _ = fi21
                    _ = v22
                }
            case "foo bar":
                switch fi20 {
                case -9223372036854775808:
                    v2 = v2
                case -2147483648:
                    for fi23 := 0; fi23 < 14; fi23++ {
                        if len(v16) >= 2 {
                            v16 = v16[1:len(v16)]
                        }
                        _ = fi23
                    }
                }
            }
        case -128:
            v18 = (v18 / v18)
        default:
            switch fi20 {
            case 32768:
                if len(v16) > 0 {
                    v16[0] = (-(fi20))
                }
            case -331:
                if true {
                    for fi24 := 0; fi24 < 13; fi24++ {
                        var v25 uint8 = 0
                        _ = fi24
                        _ = v25
                    }
                }
            case 128:
                for fi26 := 0; fi26 < 7; fi26++ {
                    v1 -= (v1 << ((255) & 63))
                    switch 32768 {
                    case 32767:
                        fmt.Println(v2, v2)
                    }
                    if (663 == -2147483648) {
                        fmt.Println(fi26, v2, v1)
                    }
                    if len(v16) >= 2 {
                        v16 = v16[1:len(v16)]
                    }
                    _ = fi26
                }
            }
        }
        var v27 uint8 = 1
        _ = fi20
        _ = v27
    }
    for fi28 := 0; fi28 < 12; fi28++ {
        var v29 float32 = ((-2.25 - (0.0 + (((float32(0.1) - (0.0 + 3.14159)) * (1.5 - 3.14159)) / (1e10 - (3.14159 + ((3.14159 + (3.14159 - 1.5)) * (0.0 + float32(v18)))))))) + (0.0 - 1.5))
        _ = fi28
        _ = v29
    }
    fmt.Println(v3, v3)
    fmt.Println(v3)
    var v30 uint32 = 65535
    v31 := GenS2{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v31
}
