package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 uint16) int16 {
    p0 = 1
    v6 := []int{-2147483648, 32768}
    _ = p0
    _ = v6
    return -723
}

func main() {
    var v1 int16 = -759
    var v2 uint32 = 42
    v3 := []int{}
    v4 := make(map[string]int)
    var v5 uint16 = 0
    if (!false) {
        var v7 uint8 = 0
        if len(v3) > 0 {
            v3[0] = 2147483647
        }
        _ = v7
    }
    v8 := GenS1{F0: 2147483647}
    v2 = v2
    var v9 float32 = -2.25
    for fi10 := 0; fi10 < 20; fi10++ {
        for fi11 := 0; fi11 < 16; fi11++ {
            fmt.Println(v5, fi10, fi10)
            _ = fi11
        }
        _ = fi10
    }
    _ = len(v4)
    fmt.Println(v9)
    for fi12 := 0; fi12 < 9; fi12++ {
        var v13 float32 = (0.1 / (v9 * v9))
        v9 = (((v13 * v9) - 1.5) - v13)
        var v14 float32 = 0.0
        _ = fi12
        _ = v13
        _ = v14
    }
    v5 = 20
    fmt.Println(v9, v5, v2)
    if false {
        switch "a" {
        case "hello":
            v1 += 0
        case "foo bar":
            for fi15 := 0; fi15 < 20; fi15++ {
                v16 := v4["foo bar"]
                if false {
                    v9 = (0.0 / -2.25)
                    if (300 == 523) {
                        v17 := GenS1{F0: int32(v1)}
                        v18 := GenS1{F0: 2147483647}
                        v19 := GenS1{F0: 0}
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                }
                _ = fi15
                _ = v16
            }
        case "a":
            if (!false) {
                for fi20 := 0; fi20 < 17; fi20++ {
                    v2 += (v2 >> ((0) & 63))
                    v21 := GenS1{F0: 42}
                    switch "hello" {
                    case "\n\t":
                        if (!(0.1 <= ((3.14159 * 3.14159) - 1.5))) {
                            var v22 bool = (!(!(!true)))
                            switch "\n\t" {
                            case "\n\t":
                                fmt.Println(v5, v9, v5)
                            default:
                                v23 := GenS1{F0: 0}
                                _ = v23
                            }
                            v1 += (v1 + v1)
                            _ = v22
                        } else {
                            v24 := GenS1{F0: 0}
                            _ = v24
                        }
                    case "foo bar":
                        v5 = 702
                    case "":
                        if (56 <= 1) {
                            v3 = append(v3, fi20)
                            v25 := v4[(("a" + ((("a" + "a") + (("foo bar" + ("\n\t" + "a")) + "hello")) + ((("foo bar" + (((("foo bar" + "hello") + "foo bar") + "a") + "hello")) + "") + "foo bar"))) + "")]
                            fmt.Println(v2, v25, v5)
                            _ = v25
                        } else {
                            v9 = float32(float64((v9 * (float32(0.1) / 1e10))))
                            v4["foo bar"] = fi20
                            v3 = append(v3, 513)
                        }
                    }
                    _ = fi20
                    _ = v21
                }
                for fi26 := 0; fi26 < 4; fi26++ {
                    for fi27 := 0; fi27 < 4; fi27++ {
                        var v28 uint32 = v2
                        _ = len(v3)
                        _ = fi27
                        _ = v28
                    }
                    for fi29 := 0; fi29 < 15; fi29++ {
                        var v30 uint16 = v5
                        _ = fi29
                        _ = v30
                    }
                    _ = fi26
                }
            } else {
                if len(v3) > 0 {
                    v3[0] = -1
                }
                v4[("hello" + "a")] = 224
            }
        }
        var v31 uint32 = 4294967295
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
