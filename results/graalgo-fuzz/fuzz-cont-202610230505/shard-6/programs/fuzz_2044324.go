package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int32, p1 uint16) bool {
    var v6 int64 = 42
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = v6
    return (!false)
}

func main() {
    v1 := make(map[string]int)
    var v2 float32 = ((3.14159 * (3.14159 / float32(float64(1e10)))) / (0.0 - (((-2.25 + (0.1 / -2.25)) * (float32(3.14159) * 1e10)) + 3.14159)))
    var v3 float64 = (float64(0.1) / 1.5)
    var v4 bool = false
    var v5 int32 = 1
    switch "a" {
    case "":
        var v7 string = "a"
        _ = v7
    case "foo bar":
        v8 := v1["foo bar"]
        _ = v8
    default:
        var v9 int8 = 103
        _ = v9
    }
    var v10 bool = false
    switch -685 {
    case -1:
        v11 := GenS1{F0: 32767, F1: (false && (540 < 42))}
        _ = v11
    default:
        if ((false && v4) || true) {
            if false {
                switch 0 {
                case 127:
                    if (((32767 <= -32768) && v10) && (float64(v2) <= (0.1 - v3))) {
                        v10 = v4
                    } else {
                        if v10 {
                            v12 := []int{32767, 0, 9223372036854775807}
                            fmt.Println(v3, v4)
                            var v13 uint16 = 255
                            _ = v12
                            _ = v13
                        }
                        if (966 == 216) {
                            v1["hello"] = 42
                            for fi14 := 0; fi14 < 20; fi14++ {
                                _ = len(v1)
                                _ = fi14
                            }
                        }
                    }
                case 9223372036854775807:
                    v1[("" + ("a" + ("foo bar" + (("hello" + ("\n\t" + "\n\t")) + ""))))] = -2147483648
                default:
                    v15 := []int{32768}
                    _ = v15
                }
                v16 := GenS1{F0: 127, F1: (((v2 - (0.1 / ((v2 * v2) - 1.5))) + float32(-2.25)) > 0.0)}
                var v17 int16 = -32768
                _ = v16
                _ = v17
            }
        } else {
            v18 := GenS1{F0: 890, F1: (true || true)}
            _ = v18
        }
    }
    v3 += (0.0 * v3)
    v4 = (((("\n\t" + "foo bar") + "hello") <= "\n\t") || (!v10))
    for fi19 := 0; fi19 < 1; fi19++ {
        if v10 {
            var v20 int8 = 42
            _ = v20
        }
        var v21 uint = 18446744073709551615
        v3 = v3
        _ = fi19
        _ = v21
    }
    v22 := GenS1{F0: 423, F1: v4}
    if (-9223372036854775808 == 9223372036854775807) {
        fmt.Println(v4)
        v5 += v5
        var v23 uint16 = 65535
        _ = v23
    }
    v24 := []int{9223372036854775807, 4294967295, 814, 32767}
    for fi25 := 0; fi25 < 16; fi25++ {
        v26 := GenS1{F0: int(v5), F1: (false && ((2 > 2) || (false || v10)))}
        v3 += ((float64(-2.25) * float64(v2)) + ((v3 + v3) - v3))
        if len(v24) > 0 {
            v24[0] = fi25
        }
        _ = fi25
        _ = v26
    }
    v2 -= v2
    if v4 {
        var v27 bool = v4
        v3 = 3.14159
        _ = v27
    }
    fmt.Println(v5, v10, v4)
    v28 := GenS2{F0: 2}
    v3 = float64(((v2 / v2) - -2.25))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v24
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
