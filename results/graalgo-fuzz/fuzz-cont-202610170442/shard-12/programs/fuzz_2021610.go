package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 uint32 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 int; F1 uint32; F2 string }
func genFn0(p0 uint64) int8 {
    p0 = p0
    if true {
        for fi6 := 0; fi6 < 3; fi6++ {
            if false {
                var v7 uint32 = 8
                v8 := GenS1{F0: 42, F1: int16(fi6), F2: v7}
                _ = v7
                _ = v8
            }
            _ = fi6
        }
        for fi9 := 0; fi9 < 15; fi9++ {
            if ((!true) && (((p0 & p0) != 18446744073709551615) && true)) {
                p0 -= 751
                p0 -= 2
            } else {
                var v10 int32 = -2147483648
                fmt.Println(v10, fi9)
                switch "foo bar" {
                case "a":
                    var v11 int = (^(fi9))
                    _ = v11
                case "foo bar":
                    v12 := []int{128, (fi9 ^ 32767), 4294967295}
                    _ = v12
                case "hello":
                    switch "hello" {
                    case "\n\t":
                        v13 := GenS1{F0: (p0 + p0), F1: 1000, F2: 2147483647}
                        _ = v13
                    case "hello":
                        v14 := []int{-2147483648, (fi9 % ((fi9) | 1)), fi9, fi9}
                        _ = v14
                    case "a":
                        v15 := []int{-782}
                        _ = v15
                    default:
                        v16 := GenS2{F0: (v10 % ((v10) | 1)), F1: 42}
                        _ = v16
                    }
                default:
                    v10 += v10
                }
                _ = v10
            }
            var v17 int32 = 568
            _ = fi9
            _ = v17
        }
    } else {
        p0 -= 18446744073709551615
    }
    _ = p0
    return 127
}

func main() {
    var v1 bool = ((!(341 != 1)) || true)
    var v2 bool = (255 < 649)
    var v3 uint32 = 2
    var v4 uint = 755
    var v5 float32 = 1.5
    switch "a" {
    case "hello":
        for fi18 := 0; fi18 < 16; fi18++ {
            var v19 float32 = v5
            var v20 int16 = -282
            _ = fi18
            _ = v19
            _ = v20
        }
    default:
        var v21 uint16 = 65535
        _ = v21
    }
    v22 := GenS2{F0: -1, F1: 18446744073709551615}
    fmt.Println(v1, v5, v2)
    v23 := GenS3{F0: -767, F1: 4294967295, F2: ("\n\t" + "")}
    for fi24 := 0; fi24 < 5; fi24++ {
        for fi25 := 0; fi25 < 1; fi25++ {
            v4 += 255
            var v26 float64 = ((0.1 * (1.5 - (-2.25 + 1e10))) - 1.5)
            fmt.Println(v3, fi25, fi24)
            _ = fi25
            _ = v26
        }
        var v27 uint32 = (v3 * (v3 + (v3 << ((v4) & 63))))
        v5 -= v5
        _ = fi24
        _ = v27
    }
    v3 = v3
    fmt.Println(v1, v2, v5)
    v28 := []int{-207, 188, -701}
    var v29 uint16 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
