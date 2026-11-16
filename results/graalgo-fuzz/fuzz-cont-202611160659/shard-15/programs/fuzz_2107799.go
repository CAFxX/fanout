package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 bool }
type GenS2 struct { F0 uint64; F1 string; F2 int64 }

func main() {
    var v1 float64 = 1.5
    v2 := []string{(("" + "\n\t") + "foo bar")}
    var v3 bool = ((!((true || true) && true)) && (float64(0.1) <= 1.5))
    v4 := GenS1{F0: 1, F1: -1, F2: (((v1 + v1) * ((0.0 - (3.14159 - (float64(float32(v1)) + (((v1 / float64(3.14159)) * 1e10) / 0.1)))) - 3.14159)) > 1.5)}
    var v5 uint64 = 1
    var v6 float64 = (((v1 - ((1.5 + float64((((float32((1.5 / v1)) * -2.25) + -2.25) / 3.14159))) / v1)) / ((v1 * ((v1 * float64((1.5 / 1.0))) - v1)) / ((v1 / v1) / (float64(0.1) - (1.5 * ((float64(-2.25) + (v1 / (v1 * (((-2.25 / (v1 - 1e10)) - (3.14159 * 3.14159)) - v1)))) * (1.5 / v1))))))) + 1e10)
    fmt.Println(v1, v6, v5)
    v1 += (v6 / 1.5)
    v3 = v3
    v6 += (1e10 + (v6 / 1.5))
    for fi7 := 0; fi7 < 4; fi7++ {
        v1 -= v6
        switch ("a" + "\n\t") {
        case "\n\t":
            switch (fi7 & (-(fi7))) {
            case -1:
                for fi8 := 0; fi8 < 20; fi8++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v9 := GenS2{F0: 42, F1: "hello", F2: -976}
                    _ = fi8
                    _ = v9
                }
            case 661:
                var v10 int16 = 0
                _ = v10
            case 0:
                v11 := GenS1{F0: 0, F1: 100, F2: (v1 >= (1e10 / 1e10))}
                _ = v11
            default:
                v6 = v6
            }
        }
        _ = len(v2)
        _ = fi7
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
