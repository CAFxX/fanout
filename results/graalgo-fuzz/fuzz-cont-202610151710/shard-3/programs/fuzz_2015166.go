package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 int32 }
type GenS2 struct { F0 string; F1 int16 }
func genFn0(p0 int, p1 uint, p2 int32) int8 {
    v4 := []int{(^(p0)), p0, (-(p0))}
    var v5 int = p0
    _ = len(v4)
    p2 -= -2147483648
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return -128
}

func main() {
    v1 := []int{-1}
    var v2 int8 = -76
    var v3 uint16 = 42
    var v6 float32 = float32(1e10)
    v7 := GenS1{F0: 42, F1: 255, F2: 1}
    switch "" {
    case "hello":
        v1 = append(v1, -355)
    }
    var v8 uint8 = 128
    v9 := GenS1{F0: 0, F1: (^(v3)), F2: 859}
    _ = len(v1)
    var v10 int16 = 42
    var v11 float64 = 3.14159
    v6 += v6
    if true {
        for fi12 := 0; fi12 < 7; fi12++ {
            var v13 uint32 = 2
            if (-185 == 2147483647) {
                var v14 bool = false
                for fi15 := 0; fi15 < 18; fi15++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v16 int32 = 1
                    _ = fi15
                    _ = v16
                }
                if (!((v11 == ((v11 * v11) - 0.1)) || true)) {
                    v17 := GenS2{F0: "hello", F1: (v10 * 32767)}
                    var v18 int32 = 0
                    _ = v17
                    _ = v18
                } else {
                    for fi19 := 0; fi19 < 14; fi19++ {
                        if ((v10 / (((^(v10))) | 1)) < -1) {
                            v1 = append(v1, (fi19 & (fi12 * (fi19 - 889))))
                            v20 := GenS2{F0: ("\n\t" + (("a" + "a") + "\n\t")), F1: v10}
                            _ = v20
                        }
                        _ = fi19
                    }
                    var v21 uint32 = (^(v13))
                    _ = v21
                }
                _ = v14
            }
            for fi22 := 0; fi22 < 2; fi22++ {
                switch (fi22 >> ((uint(v13)) & 63)) {
                case 32768:
                    v8 = (^(v8))
                case -762:
                    v23 := GenS1{F0: 2147483647, F1: 65535, F2: 674}
                    _ = v23
                case -879:
                    var v24 int8 = genFn0((fi22 << ((583) & 63)), 0, 1)
                    _ = v24
                }
                switch (fi22 / ((fi12) | 1)) {
                case 2147483647:
                    v2 -= int8(fi12)
                }
                v25 := GenS2{F0: "hello", F1: v10}
                fmt.Println(fi22)
                fmt.Println(v13, v13)
                _ = fi22
                _ = v25
            }
            _ = fi12
            _ = v13
        }
    } else {
        v11 = (float64((1e10 + 1.5)) / (v11 * 1.5))
        v2 += genFn0(32767, 382, 700)
    }
    _ = len(v1)
    var v26 uint = 322
    var v27 uint8 = v8
    var v28 int = 2147483647
    v26 -= 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
