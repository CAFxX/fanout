package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 uint32; F1 uint64; F2 uint8 }
type GenS3 struct { F0 uint16; F1 uint }
func genFn0(p0 uint16) string {
    p0 -= 42
    for fi7 := 0; fi7 < 18; fi7++ {
        p0 -= (p0 / ((p0) | 1))
        v8 := GenS1{F0: 18446744073709551615, F1: 1}
        _ = fi7
        _ = v8
    }
    _ = p0
    return "foo bar"
}
func genFn1() uint {
    var v9 int8 = 100
    var v10 bool = false
    fmt.Println(v10)
    v9 = v9
    _ = v10
    _ = v9
    return 65535
}

func main() {
    var v1 int = 32768
    v2 := make(map[uint32]string)
    v3 := []uint32{65535}
    v4 := []int{(-(v1)), v1, (v1 | v1)}
    var v5 uint = 0
    var v6 bool = true
    v11 := GenS2{F0: 2147483648, F1: 42, F2: 128}
    var v12 uint16 = 890
    var v13 uint32 = 651
    if (v6 || (42 == 42)) {
        v6 = ((9223372036854775807 >= 42) || true)
        v3 = append(v3, (^(v13)))
    }
    switch "" {
    case "foo bar":
        if v6 {
            v13 += 2147483648
            v12 += v12
        }
    case "\n\t":
        if v6 {
            var v14 int = 32767
            for fi15 := 0; fi15 < 12; fi15++ {
                switch "" {
                case "hello":
                    v12 -= v12
                case "foo bar":
                    switch ("\n\t" + "") {
                    case "\n\t":
                        var v16 int64 = -9223372036854775808
                        _ = v16
                    case "hello":
                        for fi17 := 0; fi17 < 4; fi17++ {
                            v14 -= (v14 + 0)
                            v18 := GenS3{F0: v12, F1: (v5 << (((v5 ^ v5)) & 63))}
                            fmt.Println(fi15, v12)
                            _ = fi17
                            _ = v18
                        }
                    }
                case "":
                    var v19 float64 = 0.0
                    _ = v19
                }
                _ = fi15
            }
            _ = v14
        } else {
            var v20 uint16 = (v12 | v12)
            _ = v20
        }
    default:
        var v21 float32 = (3.14159 - 0.1)
        _ = v21
    }
    var v22 string = "a"
    var v23 int8 = 1
    v24 := GenS1{F0: (^(v5)), F1: 1}
    if len(v4) > 0 {
        v4[0] = (v1 >> (((v5 * uint(v12))) & 63))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
