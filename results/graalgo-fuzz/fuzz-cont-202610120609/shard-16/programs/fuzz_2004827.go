package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 uint, p1 uint32, p2 uint) string {
    var v6 float64 = float64(0.0)
    v6 -= 0.0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return "foo bar"
}
func genFn1() uint {
    v7 := []int{2147483647, -406, 4294967295, 544}
    var v8 string = "a"
    _ = v7
    _ = v8
    return 579
}

func main() {
    v1 := make(map[string]uint32)
    var v2 bool = (-2.25 == 0.1)
    var v3 int8 = -25
    v4 := []int{}
    var v5 int8 = 42
    v9 := v1[("foo bar" + "")]
    v9 += v9
    v2 = false
    _ = len(v1)
    var v10 uint16 = 65535
    var v11 int8 = 100
    fmt.Println(v2, v9)
    fmt.Println(v11)
    v3 = 127
    switch "" {
    case "":
        _ = len(v4)
    case "hello":
        for fi12 := 0; fi12 < 15; fi12++ {
            for fi13 := 0; fi13 < 14; fi13++ {
                if (32767 == 0) {
                    var v14 uint32 = (v9 << ((1) & 63))
                    v15 := GenS1{F0: (v14 % ((v9) | 1))}
                    _ = v14
                    _ = v15
                }
                var v16 uint16 = v10
                v2 = v2
                _ = fi13
                _ = v16
            }
            for fi17 := 0; fi17 < 11; fi17++ {
                v4 = append(v4, fi17)
                _ = fi17
            }
            _ = fi12
        }
    case "\n\t":
        _ = len(v1)
    default:
        v11 = (v11 ^ v5)
    }
    if (255 < v10) {
        var v18 uint16 = (v10 / ((v10) | 1))
        switch 599 {
        case 128:
            fmt.Println(v10, v9, v9)
        case -9223372036854775808:
            switch 171 {
            case 42:
                _ = len(v4)
            case 0:
                if true {
                    _ = len(v1)
                } else {
                    v19 := GenS1{F0: (v9 << ((genFn1()) & 63))}
                    _ = v19
                }
            default:
                _ = len(v4)
            }
        }
        _ = v18
    } else {
        _ = len(v1)
        fmt.Println(v11, v9)
        fmt.Println(v5)
    }
    switch -872 {
    case 128:
        var v20 string = "foo bar"
        _ = v20
    case 9223372036854775807:
        v2 = (("\n\t" > "") || v2)
    default:
        for fi21 := 0; fi21 < 8; fi21++ {
            var v22 int = fi21
            switch v22 {
            case 1:
                for fi23 := 0; fi23 < 13; fi23++ {
                    var v24 uint64 = 18446744073709551615
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi23
                    _ = v24
                }
            default:
                var v25 uint16 = (v10 << ((476) & 63))
                _ = v25
            }
            _ = fi21
            _ = v22
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
