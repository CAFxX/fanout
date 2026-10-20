package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 int, p1 int, p2 int8) int8 {
    var v4 uint64 = 1
    fmt.Println(p1, v4, p0)
    if (!true) {
        var v5 int8 = p2
        v6 := []int{p1, (p1 / (((p0 & (^(p0)))) | 1)), (p0 / ((p1) | 1)), -511}
        _ = v5
        _ = v6
    } else {
        var v7 uint8 = 2
        v7 = 127
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 0
}

func main() {
    v1 := []string{"hello"}
    var v2 float64 = 1.5
    v3 := make(map[string]uint32)
    v2 -= float64(-2.25)
    _ = len(v1)
    v8 := GenS1{F0: 137}
    var v9 uint = 1
    var v10 uint16 = 1
    v2 += float64(float32(0.0))
    v10 = (v10 | v10)
    var v11 uint8 = 122
    switch "a" {
    case "foo bar":
        var v12 uint8 = v11
        _ = v12
    default:
        if true {
            v9 = v9
            if (0.0 > 0.1) {
                var v13 int32 = -2147483648
                var v14 uint16 = 42
                var v15 uint8 = 246
                _ = v13
                _ = v14
                _ = v15
            } else {
                var v16 uint8 = v11
                _ = v16
            }
            v1 = append(v1, "foo bar")
        } else {
            switch -465 {
            case -9223372036854775808:
                v11 -= v11
            default:
                for fi17 := 0; fi17 < 5; fi17++ {
                    if true {
                        var v18 uint64 = 1
                        v18 = v18
                        switch "" {
                        case "\n\t":
                            for fi19 := 0; fi19 < 13; fi19++ {
                                v20 := v3["a"]
                                fmt.Println(v10, v18)
                                switch "\n\t" {
                                case "foo bar":
                                    v20 -= 4294967295
                                case "\n\t":
                                    v11 += (v11 + v11)
                                }
                                _ = fi19
                                _ = v20
                            }
                        case "":
                            for fi21 := 0; fi21 < 11; fi21++ {
                                fmt.Println(v11)
                                fmt.Println(v2, v9)
                                _ = fi21
                            }
                        }
                        _ = v18
                    }
                    v10 = (v10 & 65535)
                    var v22 int32 = -2147483648
                    _ = fi17
                    _ = v22
                }
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
