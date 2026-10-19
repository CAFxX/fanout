package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
func genFn0(p0 uint, p1 bool) string {
    if (true && (127 < 41)) {
        p1 = (-2.25 <= 0.0)
    } else {
        if p1 {
            var v5 uint = p0
            var v6 uint32 = 2147483648
            _ = v5
            _ = v6
        } else {
            p1 = p1
            p0 -= (p0 + 898)
        }
    }
    v7 := GenS1{F0: (p0 >> ((p0) & 63)), F1: 9223372036854775807}
    v8 := []int{-9223372036854775808, 544, 128}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return ("" + ((("hello" + ("a" + ("" + ""))) + ("" + (("" + "") + "\n\t"))) + ""))
}

func main() {
    var v1 float64 = (1.5 / -2.25)
    v2 := []int{1, 1, 128}
    var v3 uint8 = 32
    v4 := []int{}
    v9 := GenS1{F0: 255, F1: 9223372036854775807}
    var v10 uint = 4294967295
    v3 = (v3 & uint8(v10))
    if ((((v1 - v1) - v1) + ((v1 * (-2.25 - ((v1 + float64(float32(v1))) + -2.25))) + ((1.5 / (float64(3.14159) / 1.0)) * v1))) != float64(float32((-2.25 * v1)))) {
        fmt.Println(v10, v3, v1)
        if (((v3 / ((254) | 1)) <= v3) || true) {
            if (!("\n\t" != ("a" + "hello"))) {
                if true {
                    v3 += v3
                }
            }
            for fi11 := 0; fi11 < 6; fi11++ {
                for fi12 := 0; fi12 < 20; fi12++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    switch ("a" + "hello") {
                    case "\n\t":
                        v10 = v10
                    case "":
                        if len(v2) > 0 {
                            v2[0] = (-(fi11))
                        }
                    default:
                        if len(v2) > 0 {
                            v2[0] = (fi12 & (fi11 - fi11))
                        }
                    }
                    _ = fi12
                }
                _ = fi11
            }
        }
    }
    v10 += 780
    fmt.Println(v10, v1)
    if (!(!false)) {
        var v13 int16 = -1
        _ = v13
    }
    var v14 int64 = -1
    v15 := GenS1{F0: 255, F1: (-(v14))}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
