package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 int16; F1 int32 }
func genFn0(p0 bool) int16 {
    if p0 {
        if p0 {
            fmt.Println(p0)
        } else {
            var v6 int16 = -32768
            _ = v6
        }
    }
    var v7 float64 = 3.14159
    if (!false) {
        var v8 uint16 = 0
        var v9 uint64 = 42
        _ = v8
        _ = v9
    } else {
        v7 = v7
        fmt.Println(p0, p0)
    }
    var v10 string = (("\n\t" + "foo bar") + "a")
    _ = p0
    _ = v10
    _ = v7
    return 1
}

func main() {
    var v1 int64 = 0
    v2 := make(map[int]uint32)
    var v3 int64 = v1
    var v4 float32 = 1e10
    var v5 int64 = v3
    var v11 bool = true
    v1 -= (v5 | 1)
    v1 += v3
    fmt.Println(v5, v11)
    v12 := GenS1{F0: -32768, F1: -81}
    v13 := []int{9223372036854775807}
    fmt.Println(v5, v5, v4)
    if v11 {
        switch "hello" {
        case "\n\t":
            if (3.14159 > (((1e10 / 1.0) * (v4 / 1.5)) / v4)) {
                v11 = v11
                switch -2147483648 {
                case 42:
                    v14 := GenS2{F0: 32767, F1: 0}
                    _ = v14
                case -240:
                    v2[-150] = 2147483647
                default:
                    for fi15 := 0; fi15 < 15; fi15++ {
                        var v16 uint16 = 255
                        _ = fi15
                        _ = v16
                    }
                }
            } else {
                v4 = 0.0
                if true {
                    switch 868 {
                    case 127:
                        var v17 uint32 = 2147483648
                        _ = v17
                    case 32767:
                        if true {
                            v18 := v2[4]
                            for fi19 := 0; fi19 < 6; fi19++ {
                                v20 := v2[fi19]
                                for fi21 := 0; fi21 < 17; fi21++ {
                                    _ = len(v2)
                                    var v22 int32 = -291
                                    _ = fi21
                                    _ = v22
                                }
                                if (737 != 32767) {
                                    v13 = append(v13, fi19)
                                    v23 := v2[32768]
                                    delete(v2, fi19)
                                    _ = v23
                                } else {
                                    fmt.Println(fi19, v20)
                                    for fi24 := 0; fi24 < 20; fi24++ {
                                        v4 -= v4
                                        if v11 {
                                            var v25 uint = 1
                                            var v26 int16 = 1000
                                            _ = v25
                                            _ = v26
                                        }
                                        _ = fi24
                                    }
                                    var v27 string = "\n\t"
                                    _ = v27
                                }
                                _ = fi19
                                _ = v20
                            }
                            _ = v18
                        }
                    }
                    var v28 float64 = float64(float32((((1.5 + 0.0) * 1.5) * -2.25)))
                    var v29 uint64 = 809
                    _ = v28
                    _ = v29
                }
                var v30 int = 32768
                _ = v30
            }
        case "foo bar":
            fmt.Println(v5)
        case "a":
            v4 -= float32(1e10)
        }
        v3 = (v5 >> ((4294967295) & 63))
        for fi31 := 0; fi31 < 19; fi31++ {
            if len(v13) > 0 {
                v13[0] = (fi31 & fi31)
            }
            v11 = (!false)
            _ = fi31
        }
    }
    var v32 int16 = 32767
    var v33 int64 = -9223372036854775808
    fmt.Println(v4, v1, v1)
    v34 := GenS2{F0: (v32 / ((v32) | 1)), F1: -1}
    switch "a" {
    case "a":
        fmt.Println(v3, v4)
    }
    v32 -= (v32 ^ v32)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
