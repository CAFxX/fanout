package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 int64 }
func genFn0(p0 uint16, p1 uint16, p2 uint8) int {
    var v6 float32 = (0.1 - float32((1e10 + ((1e10 - (0.0 + 0.1)) * 0.0))))
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 string = "foo bar"
        _ = fi7
        _ = v8
    }
    switch -1 {
    case -6:
        p2 += p2
    case 419:
        for fi9 := 0; fi9 < 13; fi9++ {
            if (true && (18446744073709551615 != uint64(p2))) {
                p1 = (^(p1))
            }
            p1 = (p0 | (p1 + (p1 / ((708) | 1))))
            if false {
                v10 := GenS1{F0: 32767}
                for fi11 := 0; fi11 < 19; fi11++ {
                    var v12 string = "hello"
                    var v13 uint8 = 255
                    if false {
                        v13 += 42
                    }
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                switch ("foo bar" + "\n\t") {
                case "":
                    for fi14 := 0; fi14 < 1; fi14++ {
                        v15 := []int{(fi9 % (((fi9 ^ -286)) | 1)), fi9, 2147483647}
                        _ = fi14
                        _ = v15
                    }
                }
                _ = v10
            } else {
                p0 += (p1 + 32767)
                v6 += (v6 / (0.1 / ((((((v6 / float32(float64(1e10))) + v6) / v6) - 0.1) + (0.0 - v6)) / 0.1)))
            }
            _ = fi9
        }
    default:
        fmt.Println(p1, p0)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -51
}

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 18446744073709551615
    var v3 int16 = 547
    v4 := make(map[uint32]int)
    var v5 string = ("\n\t" + (("hello" + ("a" + "")) + "foo bar"))
    var v16 uint = 18446744073709551615
    var v17 int32 = -1
    v17 -= 36
    v18 := GenS1{F0: v3}
    var v19 int32 = v17
    for fi20 := 0; fi20 < 15; fi20++ {
        fmt.Println(fi20, v3)
        v3 -= v3
        for fi21 := 0; fi21 < 5; fi21++ {
            fmt.Println(fi21)
            _ = len(v1)
            fmt.Println(v16)
            v22 := GenS1{F0: 0}
            _ = fi21
            _ = v22
        }
        _ = fi20
    }
    v23 := GenS2{F0: 536, F1: 9223372036854775807}
    v24 := v4[1]
    var v25 uint32 = 907
    switch int(v3) {
    case -1:
        for fi26 := 0; fi26 < 12; fi26++ {
            for fi27 := 0; fi27 < 12; fi27++ {
                fmt.Println(v17, v16)
                _ = fi27
            }
            v24 = genFn0(65535, 459, uint8(v25))
            v25 -= (v25 + v25)
            _ = fi26
        }
    case -9223372036854775808:
        v1[v5] = genFn0(uint16(v16), 32767, 255)
    }
    v25 += (^(v25))
    fmt.Println(v24, v25)
    v28 := []int{(-(v24))}
    if false {
        v17 = -2147483648
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
