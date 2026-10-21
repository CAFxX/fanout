package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 int8 }
func genFn0(p0 uint32) int8 {
    if ((0.0 * (1e10 * (((0.0 / (1.5 + -2.25)) / 0.1) * 0.0))) == 1.5) {
        switch 0 {
        case 42:
            var v6 int = -609
            _ = v6
        case 130:
            v7 := GenS1{F0: -1, F1: 0}
            _ = v7
        case 32768:
            p0 -= (p0 / (((p0 >> ((682) & 63))) | 1))
        }
        var v8 bool = (!true)
        _ = v8
    }
    v9 := []int{32767, 128}
    v10 := GenS1{F0: 2147483647, F1: -24}
    _ = p0
    _ = v10
    _ = v9
    return 0
}

func main() {
    var v1 float32 = 0.0
    v2 := make(map[uint32]bool)
    v3 := make(map[int]string)
    var v4 int8 = 1
    var v5 uint = 65535
    v3[127] = ("\n\t" + "foo bar")
    var v11 int16 = 0
    var v12 int = 32767
    v13 := GenS2{F0: (v4 | v4)}
    if false {
        var v14 uint = v5
        if false {
            fmt.Println(v11, v1)
            if true {
                v1 += 1.5
                var v15 uint32 = 2147483648
                if false {
                    for fi16 := 0; fi16 < 14; fi16++ {
                        var v17 uint64 = 861
                        v18 := []int{-9223372036854775808}
                        if (254 == 2) {
                            if len(v18) > 0 {
                                v18[0] = fi16
                            }
                            var v19 uint64 = (v17 * (v17 + v17))
                            _ = v19
                        } else {
                            fmt.Println(v1, v17, v15)
                        }
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                } else {
                    v20 := GenS2{F0: int8(v15)}
                    _ = v20
                }
                _ = v15
            }
            for fi21 := 0; fi21 < 16; fi21++ {
                switch "hello" {
                case "\n\t":
                    v22 := []int{fi21, fi21, v12, (fi21 + v12)}
                    _ = v22
                case "hello":
                    fmt.Println(v12, v5, v12)
                case "foo bar":
                    var v23 bool = (1 == 1)
                    _ = v23
                default:
                    v12 = (fi21 / ((v12) | 1))
                }
                _ = fi21
            }
        } else {
            var v24 uint8 = 193
            _ = v24
        }
        _ = v14
    }
    var v25 int = (v12 + v12)
    v2[65535] = (true && (true && (523 == 750)))
    v26 := GenS1{F0: int32(v25), F1: (v4 - (v4 ^ v4))}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
