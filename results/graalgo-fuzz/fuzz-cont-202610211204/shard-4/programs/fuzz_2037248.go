package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }

func main() {
    var v1 string = "foo bar"
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]int)
    var v4 uint64 = 1
    var v5 uint32 = 837
    var v6 int = 1
    switch "\n\t" {
    case "":
        var v7 int8 = 0
        _ = v7
    default:
        v5 -= v5
    }
    var v8 float64 = (float64(1.5) - 1.5)
    var v9 uint16 = 2
    for fi10 := 0; fi10 < 1; fi10++ {
        v9 -= 970
        if true {
            if ((v1 + v1) < "foo bar") {
                if true {
                    v11 := []int{(fi10 + 2147483647), -929}
                    switch "" {
                    case "":
                        _ = len(v3)
                    case "\n\t":
                        if ((float32(float64(1e10)) * -2.25) >= -2.25) {
                            v6 = fi10
                            v8 = (v8 - v8)
                            var v12 float32 = ((-2.25 - 1.5) * 0.1)
                            _ = v12
                        }
                    case "hello":
                        var v13 int = fi10
                        _ = v13
                    default:
                        switch "a" {
                        case "a":
                            v6 -= v6
                        case "":
                            var v14 float64 = (v8 / (1.5 * -2.25))
                            _ = v14
                        default:
                            v5 = 1
                        }
                    }
                    _ = v11
                }
                v6 += (fi10 ^ v6)
                for fi15 := 0; fi15 < 8; fi15++ {
                    var v16 uint = 65535
                    v17 := GenS1{F0: ((((v6 | 32768) >= (v6 ^ 0)) && (!true)) || (false || true)), F1: -128}
                    var v18 float32 = 1.5
                    _ = fi15
                    _ = v16
                    _ = v17
                    _ = v18
                }
            }
            for fi19 := 0; fi19 < 8; fi19++ {
                v20 := []int{}
                _ = fi19
                _ = v20
            }
            if false {
                for fi21 := 0; fi21 < 19; fi21++ {
                    v3[(v5 * 2147483647)] = v6
                    var v22 bool = (true || true)
                    _ = fi21
                    _ = v22
                }
            } else {
                for fi23 := 0; fi23 < 7; fi23++ {
                    var v24 int64 = 564
                    var v25 uint8 = 244
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            }
        }
        v26 := GenS1{F0: (2 <= 383), F1: 100}
        _ = fi10
        _ = v26
    }
    for fi27 := 0; fi27 < 6; fi27++ {
        v9 = v9
        _ = fi27
    }
    var v28 uint16 = uint16(v4)
    fmt.Println(v1, v9)
    v1 = v1
    fmt.Println(v9, v6, v5)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
