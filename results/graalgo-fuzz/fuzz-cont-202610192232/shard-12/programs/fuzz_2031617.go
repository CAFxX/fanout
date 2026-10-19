package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 int, p1 bool) uint32 {
    var v7 bool = p1
    var v8 float32 = 0.0
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 uint64) string {
    var v9 float64 = float64(1.5)
    v10 := GenS1{F0: 127}
    p0 -= p0
    v9 += 1.5
    var v11 string = ("foo bar" + (("hello" + "") + (("\n\t" + ("" + "hello")) + "")))
    _ = p0
    _ = v10
    _ = v11
    _ = v9
    return v11
}

func main() {
    var v1 int = 2147483647
    var v2 int = v1
    v3 := make(map[string]bool)
    var v4 uint8 = 0
    v5 := []uint32{65535}
    var v6 uint8 = v4
    if (((^(v6)) == v4) || (((!(226 != 155)) || (true || true)) || ((!(0 <= -9223372036854775808)) && true))) {
        for fi12 := 0; fi12 < 16; fi12++ {
            var v13 int32 = 2147483647
            _ = fi12
            _ = v13
        }
        for fi14 := 0; fi14 < 16; fi14++ {
            fmt.Println(fi14)
            _ = fi14
        }
    } else {
        var v15 string = ""
        _ = v15
    }
    switch v1 {
    case 275:
        v6 += (v6 / (((v4 / (((^(v6))) | 1))) | 1))
    default:
        v6 += (v6 >> ((0) & 63))
    }
    var v16 int32 = 858
    var v17 int64 = 937
    var v18 uint8 = v4
    v5 = append(v5, genFn0(32768, (!false)))
    var v19 uint64 = 0
    var v20 int64 = v17
    for fi21 := 0; fi21 < 8; fi21++ {
        for fi22 := 0; fi22 < 4; fi22++ {
            if ((3.14159 * 1.5) > 1e10) {
                v23 := v3[(("\n\t" + ("hello" + "a")) + "\n\t")]
                _ = v23
            } else {
                v3[(("hello" + "hello") + ("foo bar" + ""))] = (0 == -32768)
                if true {
                    if (float32((0.1 * ((0.0 - 1e10) - 1.5))) == 3.14159) {
                        if (!false) {
                            delete(v3, "hello")
                            var v24 float64 = ((((0.1 - float64((float32(3.14159) / -2.25))) * -2.25) * 1.5) + float64(-2.25))
                            _ = v24
                        } else {
                            fmt.Println(v18, v6)
                            delete(v3, "")
                        }
                        var v25 float64 = 1e10
                        _ = len(v3)
                        _ = v25
                    } else {
                        var v26 uint = 1
                        _ = v26
                    }
                    switch "a" {
                    case "hello":
                        v16 += (v16 % ((2147483647) | 1))
                    case "\n\t":
                        var v27 bool = false
                        _ = v27
                    default:
                        fmt.Println(v2)
                    }
                    var v28 int8 = 36
                    v5 = append(v5, 4294967295)
                    _ = v28
                } else {
                    fmt.Println(v6, v4)
                    v6 = v18
                    if (127 >= int8(v16)) {
                        if ((^(v19)) == 241) {
                            if ((false || (0 > int16(fi21))) || (v16 >= v16)) {
                                var v29 int32 = 450
                                _ = v29
                            } else {
                                var v30 int32 = v16
                                _ = v30
                            }
                        }
                        switch fi22 {
                        case 179:
                            if (9223372036854775807 == (v17 + -810)) {
                                var v31 string = "\n\t"
                                _ = v31
                            }
                        default:
                            if true {
                                v17 = (v20 << ((2) & 63))
                            }
                        }
                        v20 -= 469
                    }
                }
                var v32 bool = false
                v32 = v32
                v1 = (fi22 | fi22)
                v20 += v17
                v18 -= v18
                _ = v32
            }
            if (42 >= (v16 * 685)) {
                var v33 int16 = 1
                _ = v33
            } else {
                var v34 bool = (-128 > 118)
                _ = v34
            }
            if (32767 == -462) {
                fmt.Println(fi22, v6)
            }
            _ = fi22
        }
        var v35 uint8 = v18
        _ = fi21
        _ = v35
    }
    v18 = v6
    v2 -= v2
    v36 := GenS1{F0: v4}
    for fi37 := 0; fi37 < 12; fi37++ {
        v38 := GenS1{F0: (v18 / (((v4 * v6)) | 1))}
        v20 += 266
        _ = fi37
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
