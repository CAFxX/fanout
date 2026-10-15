package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 uint16) uint {
    p0 = p0
    for fi5 := 0; fi5 < 6; fi5++ {
        if false {
            fmt.Println(p0, p0)
            p0 = p0
        } else {
            v6 := []int{32767, (-(fi5)), (fi5 & (fi5 % ((fi5) | 1)))}
            if false {
                fmt.Println(fi5, p0)
                var v7 float64 = 0.1
                fmt.Println(v7, p0)
                _ = v7
            }
            _ = v6
        }
        if true {
            var v8 int64 = 9223372036854775807
            _ = v8
        }
        _ = fi5
    }
    p0 += (p0 + p0)
    fmt.Println(p0)
    p0 = (p0 ^ 0)
    _ = p0
    return 4294967295
}

func main() {
    var v1 string = "\n\t"
    var v2 int16 = -72
    var v3 uint64 = 42
    var v4 float64 = (0.0 + (((0.0 - (((3.14159 + (0.1 + 3.14159)) / float64(((3.14159 / 1e10) / -2.25))) / 1.0)) - -2.25) / 1.5))
    v1 = v1
    switch 4294967295 {
    case 32768:
        v1 = (v1 + v1)
    case 128:
        v9 := GenS1{F0: (((((v1 + (v1 + v1)) + ((v1 + ("foo bar" + v1)) + v1)) + (v1 + "hello")) + v1) + ((v1 + (((("hello" + ("hello" + v1)) + "\n\t") + "foo bar") + (v1 + v1))) + v1))}
        _ = v9
    }
    var v10 float64 = float64(0.1)
    var v11 uint32 = 1
    v11 = uint32(v3)
    v10 += ((v4 * (((((float64(1.5) + (0.1 * (float64((0.0 * 1e10)) * float64(0.0)))) / v10) * v4) - v4) - (float64(1e10) / (v4 - (v10 / (((float64(1e10) + v4) + v10) - (v4 / v10))))))) + v10)
    v3 += (v3 | (v3 >> ((2) & 63)))
    v12 := []int{128, -848}
    v12 = append(v12, 9223372036854775807)
    switch 32767 {
    case -9223372036854775808:
        v13 := GenS1{F0: v1}
        _ = v13
    case -2147483648:
        if (!true) {
            switch (v1 + v1) {
            case "a":
                v4 = (-2.25 + 0.0)
            case "hello":
                if (v11 >= v11) {
                    fmt.Println(v11, v4)
                    switch 32768 {
                    case 1:
                        var v14 float32 = 0.0
                        _ = v14
                    case 933:
                        var v15 int = 32768
                        _ = v15
                    case 863:
                        if ((^(v3)) == 2) {
                            var v16 int16 = 32767
                            var v17 uint8 = 131
                            switch (v1 + "foo bar") {
                            case "foo bar":
                                v1 = (v1 + ("" + v1))
                            case "hello":
                                if len(v12) > 0 {
                                    v12[0] = 32768
                                }
                            default:
                                v18 := GenS1{F0: (v1 + "\n\t")}
                                _ = v18
                            }
                            _ = v16
                            _ = v17
                        } else {
                            v10 += (((0.1 / ((1e10 + 1.5) * (v4 - v4))) / (v4 * (((1.5 / 1e10) - v4) / v4))) * -2.25)
                            var v19 uint16 = 7
                            _ = v19
                        }
                    default:
                        v12 = append(v12, 978)
                    }
                }
            case "\n\t":
                if len(v12) > 0 {
                    v12[0] = 127
                }
            default:
                var v20 uint64 = 2
                _ = v20
            }
        }
    }
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    v1 = v1
    _ = len(v12)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
