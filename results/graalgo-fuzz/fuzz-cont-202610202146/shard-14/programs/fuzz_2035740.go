package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 uint16; F2 uint8 }
type GenS4 struct { F0 int32; F1 uint8; F2 int8 }
func genFn0(p0 int16, p1 int32, p2 int64) int64 {
    if ("" == "a") {
        v7 := []int{-669, -1, 9223372036854775807}
        if ((((0.0 / float32((1e10 - 3.14159))) * (float32(1.5) + 0.1)) + 0.0) < 0.1) {
            switch 2147483647 {
            case 2147483647:
                switch ("\n\t" + "foo bar") {
                case "\n\t":
                    p1 = (p1 * 0)
                default:
                    var v8 float64 = 1.5
                    _ = v8
                }
            default:
                fmt.Println(p1, p2, p1)
            }
            p0 -= p0
        }
        _ = v7
    } else {
        p1 = p1
    }
    var v9 float32 = 1.5
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return p2
}

func main() {
    var v1 int64 = 42
    var v2 float32 = ((0.0 - -2.25) - (1.5 / ((3.14159 + -2.25) + 0.0)))
    var v3 int64 = (v1 | (v1 | 9223372036854775807))
    v4 := make(map[int]int)
    var v5 uint8 = 222
    v6 := []string{"foo bar", "hello"}
    _ = len(v4)
    v3 -= genFn0(-95, 714, (v1 % ((-9223372036854775808) | 1)))
    for fi10 := 0; fi10 < 10; fi10++ {
        switch ("foo bar" + (("foo bar" + "foo bar") + ("foo bar" + ""))) {
        case "foo bar":
            v3 += -613
        case "a":
            delete(v4, (fi10 >> ((2) & 63)))
        case "":
            v5 += v5
        default:
            v11 := GenS1{F0: (^(v5))}
            _ = v11
        }
        _ = fi10
    }
    v2 = v2
    v12 := GenS2{F0: -128}
    v6 = append(v6, "foo bar")
    v13 := GenS2{F0: 87}
    v14 := GenS2{F0: 100}
    var v15 uint32 = 166
    v16 := GenS1{F0: (v5 >> ((314) & 63))}
    if len(v6) > 0 {
        v6[0] = "hello"
    }
    v17 := GenS1{F0: 42}
    if len(v6) > 0 {
        v6[0] = ("" + ("hello" + "\n\t"))
    }
    if true {
        fmt.Println(v1, v3)
    } else {
        v5 += v5
    }
    for fi18 := 0; fi18 < 3; fi18++ {
        if true {
            v4[(fi18 << ((42) & 63))] = -1
            if (false || (!true)) {
                for fi19 := 0; fi19 < 19; fi19++ {
                    fmt.Println(v2, v1)
                    v20 := v4[fi19]
                    v4[(fi19 % (((^(fi19))) | 1))] = v20
                    switch (v20 % ((fi19) | 1)) {
                    case -2147483648:
                        v21 := GenS3{F0: 0, F1: 2, F2: (v5 + (v5 % ((v5) | 1)))}
                        _ = v21
                    }
                    fmt.Println(v20)
                    var v22 int = v20
                    v6 = append(v6, ("foo bar" + "hello"))
                    v23 := v4[v22]
                    var v24 float32 = v2
                    var v25 bool = (0.0 >= float64(((v2 + v24) / -2.25)))
                    _ = fi19
                    _ = v20
                    _ = v22
                    _ = v23
                    _ = v24
                    _ = v25
                }
                _ = len(v6)
            } else {
                fmt.Println(v1)
                v5 += 2
            }
        }
        if true {
            for fi26 := 0; fi26 < 9; fi26++ {
                _ = len(v4)
                _ = fi26
            }
        } else {
            v27 := GenS1{F0: (v5 * 255)}
            v28 := GenS4{F0: 88, F1: (v5 - (v5 ^ 249)), F2: 109}
            var v29 int64 = v1
            _ = v27
            _ = v28
            _ = v29
        }
        _ = fi18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
