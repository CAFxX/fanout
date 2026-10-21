package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 int8 }
func genFn0(p0 uint32, p1 int8, p2 uint) uint32 {
    p0 = p0
    p1 = (-(p1))
    v5 := GenS1{F0: 361, F1: (p1 / (((p1 / (((p1 & (p1 / (((p1 & p1)) | 1)))) | 1))) | 1))}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p0
}
func genFn1(p0 uint32) uint16 {
    p0 = 390
    fmt.Println(p0)
    _ = p0
    return 2
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{"foo bar", "\n\t"}
    v3 := []string{}
    v4 := []uint32{42, 2147483647, 1}
    var v6 bool = true
    for fi7 := 0; fi7 < 1; fi7++ {
        fmt.Println(fi7)
        v6 = v6
        v8 := GenS1{F0: genFn1(2), F1: 113}
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 137, F1: 72}
    if true {
        v10 := GenS1{F0: genFn1(1), F1: 0}
        _ = v10
    }
    switch ((("a" + ("a" + "")) + "\n\t") + "hello") {
    case "a":
        if v6 {
            var v11 int32 = -1
            _ = v11
        } else {
            for fi12 := 0; fi12 < 1; fi12++ {
                if ((-621 < 0) || true) {
                    if true {
                        v6 = v6
                    } else {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    for fi13 := 0; fi13 < 13; fi13++ {
                        v6 = (0 >= 42)
                        if false {
                            var v14 int16 = 1000
                            _ = v14
                        } else {
                            v6 = v6
                        }
                        v6 = (true || ((!(fi12 > 2147483647)) && (v6 || (0 > 42))))
                        _ = fi13
                    }
                    v6 = (360 >= 0)
                }
                _ = fi12
            }
            fmt.Println(v6)
            if (v6 && (2 >= 0)) {
                v15 := GenS2{F0: 100}
                _ = v15
            }
        }
    default:
        var v16 int32 = 617
        _ = v16
    }
    fmt.Println(v6)
    fmt.Println(v6)
    v17 := GenS1{F0: 255, F1: -44}
    v6 = (((((0.1 - 0.1) - 3.14159) - 1e10) - 1.5) >= float32((1.5 + (0.1 * (-2.25 / ((0.0 - 1e10) + float64(float32((-2.25 - (float64(1e10) + 0.0))))))))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v9
}
