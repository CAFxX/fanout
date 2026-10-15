package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32 }
func genFn0(p0 uint, p1 uint32, p2 int32) bool {
    fmt.Println(p0, p2, p2)
    p2 -= (-(p2))
    p1 = 2147483648
    _ = p0
    _ = p1
    _ = p2
    return (32767 == -1)
}
func genFn1() int8 {
    var v7 string = (("a" + ((("a" + "hello") + (("\n\t" + "a") + "foo bar")) + ("\n\t" + "foo bar"))) + "a")
    var v8 uint64 = 579
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 bool = (!true)
    v2 := make(map[int]bool)
    var v3 uint64 = 42
    v4 := []uint32{}
    v5 := make(map[string]string)
    var v6 uint = 860
    for fi9 := 0; fi9 < 5; fi9++ {
        var v10 bool = (!((v1 || (!(!false))) || (7 < 65535)))
        for fi11 := 0; fi11 < 17; fi11++ {
            v6 = (v6 % ((0) | 1))
            v12 := GenS1{F0: v6}
            _ = fi11
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    var v13 uint32 = 42
    fmt.Println(v13, v13)
    v5[("foo bar" + "foo bar")] = ("foo bar" + "hello")
    for fi14 := 0; fi14 < 5; fi14++ {
        fmt.Println(fi14)
        v1 = v1
        _ = fi14
    }
    var v15 uint = v6
    var v16 float64 = float64(1.5)
    switch "a" {
    case "hello":
        if v1 {
            if (17 > 42) {
                if v1 {
                    v13 -= (^(v13))
                    _ = len(v2)
                } else {
                    v17 := GenS2{F0: -2147483648}
                    if v1 {
                        for fi18 := 0; fi18 < 14; fi18++ {
                            switch fi18 {
                            case 207:
                                var v19 uint16 = 255
                                _ = v19
                            case 42:
                                var v20 uint16 = 1
                                _ = v20
                            }
                            fmt.Println(v15)
                            _ = fi18
                        }
                        v6 += 65535
                        fmt.Println(v15, v15, v3)
                    }
                    _ = v17
                }
            } else {
                var v21 float64 = (-2.25 * v16)
                _ = v21
            }
            v16 = 1.5
        } else {
            var v22 string = "foo bar"
            fmt.Println(v13, v22)
            _ = v22
        }
    case "a":
        for fi23 := 0; fi23 < 5; fi23++ {
            v24 := GenS2{F0: 2147483647}
            v25 := GenS2{F0: 42}
            for fi26 := 0; fi26 < 13; fi26++ {
                v27 := GenS1{F0: (^(v15))}
                var v28 float32 = ((float32(v16) / 0.1) / (1e10 / ((3.14159 / 1.0) / 0.1)))
                _ = fi26
                _ = v27
                _ = v28
            }
            _ = fi23
            _ = v24
            _ = v25
        }
    case "foo bar":
        var v29 float64 = (v16 * float64((((1e10 / (1e10 + (1.5 / 0.1))) + float32((v16 * v16))) / 0.1)))
        _ = v29
    default:
        v16 -= 0.1
    }
    var v30 int64 = 42
    v31 := v2[1]
    var v32 float64 = float64(float32(((v16 / v16) - float64(1.5))))
    v3 -= (v3 - v3)
    if true {
        for fi33 := 0; fi33 < 5; fi33++ {
            v32 += v16
            var v34 uint32 = (^(v13))
            _ = fi33
            _ = v34
        }
        _ = len(v2)
        v35 := GenS2{F0: -1}
        _ = v35
    } else {
        v30 += 0
    }
    v13 -= (v13 & 42)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
