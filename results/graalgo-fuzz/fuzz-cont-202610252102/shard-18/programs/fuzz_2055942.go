package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 uint8 }
type GenS2 struct { F0 int8; F1 uint64; F2 uint16 }
func genFn0(p0 int) int8 {
    if false {
        p0 = p0
    }
    p0 += 0
    _ = p0
    return 0
}
func genFn1() uint8 {
    v7 := GenS1{F0: 131, F1: 2147483647, F2: 195}
    var v8 float32 = float32(0.0)
    _ = v7
    _ = v8
    return 1
}

func main() {
    v1 := []uint32{}
    var v2 string = "hello"
    v3 := make(map[int]int)
    var v4 float64 = (0.0 + 1.5)
    var v5 int = 4294967295
    var v6 uint8 = 128
    var v9 int8 = 127
    var v10 string = ((("foo bar" + v2) + ("\n\t" + "\n\t")) + v2)
    var v11 int32 = 2147483647
    fmt.Println(v10)
    for fi12 := 0; fi12 < 12; fi12++ {
        _ = len(v3)
        _ = fi12
    }
    v10 = "\n\t"
    var v13 int8 = v9
    v14 := GenS2{F0: 127, F1: 73, F2: 370}
    for fi15 := 0; fi15 < 4; fi15++ {
        if len(v1) > 0 {
            v1[0] = 4294967295
        }
        switch "foo bar" {
        case "foo bar":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        default:
            var v16 int = fi15
            _ = v16
        }
        _ = fi15
    }
    for fi17 := 0; fi17 < 7; fi17++ {
        if (65535 > 622) {
            v9 += (^(v9))
            v1 = append(v1, 2147483648)
            for fi18 := 0; fi18 < 16; fi18++ {
                _ = len(v1)
                if false {
                    fmt.Println(v10, v9, v2)
                } else {
                    v19 := GenS1{F0: 1, F1: v11, F2: (v6 << ((647) & 63))}
                    var v20 string = "foo bar"
                    _ = v19
                    _ = v20
                }
                _ = fi18
            }
        } else {
            var v21 uint = 2
            _ = v21
        }
        _ = fi17
    }
    v22 := v3[-9223372036854775808]
    fmt.Println(v2, v6, v11)
    v4 -= v4
    v5 = (^(v5))
    if false {
        var v23 int64 = 1
        if true {
            v3[v5] = v5
            for fi24 := 0; fi24 < 14; fi24++ {
                v5 = v5
                _ = fi24
            }
        }
        v11 = v11
        _ = v23
    } else {
        var v25 int32 = (^(v11))
        v10 = v10
        _ = v25
    }
    v22 = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
