package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 uint; F1 int8 }
func genFn0(p0 int, p1 int, p2 int32) uint32 {
    for fi7 := 0; fi7 < 15; fi7++ {
        v8 := GenS1{F0: 1, F1: 2}
        _ = fi7
        _ = v8
    }
    v9 := []int{(p0 & (p0 / ((p0) | 1))), (p0 & p1)}
    if true {
        p1 += -2147483648
        p1 = 107
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 170
}

func main() {
    var v1 float32 = float32(-2.25)
    var v2 int16 = 1
    v3 := []uint32{42}
    v4 := []string{"\n\t", ""}
    var v5 float32 = (v1 * 0.0)
    v6 := make(map[uint32]uint32)
    var v10 uint64 = 2
    for fi11 := 0; fi11 < 4; fi11++ {
        var v12 int8 = 0
        _ = fi11
        _ = v12
    }
    v5 -= 1.5
    v5 += ((((v5 * (((0.0 / -2.25) - v5) * float32((0.1 - (float64(0.0) + 1.5))))) - (1e10 / v5)) + float32(3.14159)) / ((v1 * ((((v5 / 0.1) + (v1 - (v5 * (float32((0.1 + 1e10)) * float32(3.14159))))) / 1.5) / 0.1)) - (v5 + (3.14159 - (v5 + float32(-2.25))))))
    v2 -= v2
    if (!(false && true)) {
        var v13 uint32 = 1
        var v14 string = "foo bar"
        _ = v13
        _ = v14
    }
    switch 1 {
    case 32768:
        for fi15 := 0; fi15 < 12; fi15++ {
            if ((v2 == v2) && (false || false)) {
                if len(v3) > 0 {
                    v3[0] = 2147483647
                }
                var v16 int32 = 0
                _ = v16
            }
            v2 += v2
            for fi17 := 0; fi17 < 8; fi17++ {
                v1 = 0.0
                v18 := v6[32]
                _ = fi17
                _ = v18
            }
            _ = fi15
        }
    }
    v19 := GenS2{F0: 2, F1: 0}
    v10 -= 4294967295
    for fi20 := 0; fi20 < 1; fi20++ {
        if true {
            v1 += float32(((-2.25 - (0.0 + (-2.25 + 1e10))) - float64(1e10)))
            if len(v4) > 0 {
                v4[0] = ("" + "a")
            }
        } else {
            if (65535 <= 608) {
                if (!(!true)) {
                    var v21 bool = true
                    v5 = v1
                    if v21 {
                        if ((859 <= 56) || (v21 || (-100 > 42))) {
                            var v22 uint16 = 2
                            _ = v22
                        }
                        switch "foo bar" {
                        case "hello":
                            v4 = append(v4, "")
                        }
                    }
                    _ = v21
                }
                v1 -= v5
            } else {
                var v23 bool = false
                v5 += v1
                _ = v23
            }
            var v24 uint8 = 127
            _ = v24
        }
        var v25 uint16 = 255
        fmt.Println(fi20)
        _ = fi20
        _ = v25
    }
    v5 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
