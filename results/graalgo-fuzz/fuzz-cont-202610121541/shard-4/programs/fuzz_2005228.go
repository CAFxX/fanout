package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 uint32; F1 uint64 }
func genFn0(p0 bool, p1 int8) uint16 {
    var v7 bool = p0
    v8 := GenS1{F0: 42, F1: -56}
    p1 -= (p1 + p1)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 65535
}
func genFn1(p0 int) string {
    for fi9 := 0; fi9 < 18; fi9++ {
        v10 := GenS1{F0: -1, F1: 1}
        _ = fi9
        _ = v10
    }
    fmt.Println(p0)
    p0 = (-(p0))
    _ = p0
    return ("\n\t" + "hello")
}

func main() {
    v1 := make(map[uint32]int)
    var v2 float64 = (3.14159 + float64(((float32(-2.25) * (0.0 * 0.0)) * (1.5 - 3.14159))))
    v3 := []uint32{2147483648, 2147483648, 0}
    v4 := make(map[int]uint32)
    var v5 int = 32767
    var v6 uint32 = 643
    v3 = append(v3, (v6 | v6))
    v6 = (v6 + v6)
    v5 -= (-(v5))
    if true {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v2 = v2
    } else {
        if false {
            switch (v5 & v5) {
            case 2147483647:
                if false {
                    switch "foo bar" {
                    case "":
                        v11 := GenS1{F0: -14, F1: 1}
                        _ = v11
                    case "hello":
                        for fi12 := 0; fi12 < 15; fi12++ {
                            v5 = 128
                            v13 := GenS1{F0: -106, F1: 100}
                            var v14 int8 = 9
                            _ = fi12
                            _ = v13
                            _ = v14
                        }
                    case "a":
                        v15 := GenS1{F0: 71, F1: 42}
                        _ = v15
                    default:
                        v6 = v6
                    }
                } else {
                    fmt.Println(v2)
                }
            case -9223372036854775808:
                delete(v1, 42)
            case -463:
                v2 += v2
            }
        } else {
            v2 -= (1.5 / float64(((float32(0.0) - ((3.14159 - (0.1 / 3.14159)) + (3.14159 + (3.14159 * 0.0)))) + (1.5 * 0.0))))
        }
    }
    v16 := GenS2{F0: 1, F1: 255}
    v5 += v5
    v17 := GenS3{F0: 42, F1: uint64(v6)}
    for fi18 := 0; fi18 < 5; fi18++ {
        v2 = v2
        _ = fi18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
