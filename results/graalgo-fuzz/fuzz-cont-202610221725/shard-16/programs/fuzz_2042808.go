package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 int64 }
func genFn0(p0 int32) int64 {
    v5 := []int{2147483647, 32767}
    p0 += p0
    p0 = (p0 ^ p0)
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 uint8 = 255
        v7 -= 127
        _ = fi6
        _ = v7
    }
    var v8 string = "hello"
    _ = p0
    _ = v5
    _ = v8
    return 0
}
func genFn1(p0 bool) string {
    p0 = (true && false)
    var v9 string = "\n\t"
    p0 = (true && p0)
    p0 = p0
    _ = p0
    _ = v9
    return ("foo bar" + v9)
}

func main() {
    var v1 int64 = -822
    var v2 int8 = 100
    var v3 int8 = v2
    v4 := []uint32{}
    var v10 int16 = -32768
    fmt.Println(v3, v3)
    v11 := GenS1{F0: 2, F1: ""}
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 int16 = (v10 & -653)
        for fi14 := 0; fi14 < 11; fi14++ {
            switch "hello" {
            case "a":
                var v15 int8 = v3
                _ = v15
            case "\n\t":
                v16 := GenS1{F0: 0, F1: "a"}
                _ = v16
            case "foo bar":
                switch "a" {
                case "hello":
                    fmt.Println(v10, v1, v2)
                case "\n\t":
                    var v17 int = fi14
                    _ = v17
                case "a":
                    v1 += (v1 | v1)
                default:
                    switch fi14 {
                    case -971:
                        v2 = v3
                    default:
                        if len(v4) > 0 {
                            v4[0] = 2147483647
                        }
                    }
                }
            default:
                var v18 int32 = int32(v1)
                _ = v18
            }
            var v19 int32 = -1
            fmt.Println(v19, v19, v1)
            v20 := GenS1{F0: 65535, F1: "hello"}
            var v21 uint32 = 1
            var v22 float64 = 1e10
            _ = fi14
            _ = v19
            _ = v20
            _ = v21
            _ = v22
        }
        _ = fi12
        _ = v13
    }
    fmt.Println(v1, v3)
    _ = len(v4)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v1 -= v1
    v2 += v2
    v23 := GenS2{F0: genFn0(-2147483648)}
    for fi24 := 0; fi24 < 16; fi24++ {
        for fi25 := 0; fi25 < 4; fi25++ {
            if (genFn0(403) < v1) {
                var v26 bool = false
                v1 += genFn0(-2147483648)
                fmt.Println(v1)
                v27 := GenS2{F0: v1}
                _ = v26
                _ = v27
            } else {
                v3 += (v2 | v2)
            }
            var v28 float32 = (((((-2.25 - 0.0) * 0.1) + 3.14159) / float32(0.1)) * 1e10)
            if len(v4) > 0 {
                v4[0] = 65535
            }
            _ = fi25
            _ = v28
        }
        _ = fi24
    }
    v3 = (v3 & 127)
    v4 = append(v4, 810)
    v29 := GenS2{F0: genFn0(-1)}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v4
}
