package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint64 }
func genFn0(p0 int8, p1 int32, p2 int16) int8 {
    p2 = -734
    v6 := []int{9223372036854775807}
    for fi7 := 0; fi7 < 12; fi7++ {
        v6 = append(v6, (fi7 << ((1) & 63)))
        for fi8 := 0; fi8 < 6; fi8++ {
            switch "\n\t" {
            case "a":
                if true {
                    var v9 int64 = -158
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    v9 = -1
                    _ = v9
                }
            default:
                fmt.Println(p2, p2, fi8)
            }
            v10 := GenS1{F0: p2, F1: false, F2: 244}
            v11 := GenS1{F0: (p2 & (p2 % ((42) | 1))), F1: false, F2: 18446744073709551615}
            _ = fi8
            _ = v10
            _ = v11
        }
        if (!(fi7 < fi7)) {
            for fi12 := 0; fi12 < 2; fi12++ {
                var v13 uint32 = 2
                _ = fi12
                _ = v13
            }
            p2 += (-(p2))
            for fi14 := 0; fi14 < 2; fi14++ {
                p2 += p2
                _ = fi14
            }
        } else {
            p1 = -2147483648
            var v15 uint64 = 1
            _ = v15
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p0
}

func main() {
    v1 := make(map[string]string)
    var v2 string = ("" + "hello")
    v3 := []string{""}
    var v4 int64 = 288
    var v5 uint = 42
    var v16 float32 = (float32(1.5) / 3.14159)
    _ = len(v3)
    v1[((((v2 + v2) + ((v2 + v2) + (v2 + ("foo bar" + "\n\t")))) + v2) + v2)] = v2
    var v17 int = -2147483648
    fmt.Println(v4, v17, v16)
    v5 = v5
    v16 = v16
    var v18 uint32 = 42
    var v19 float64 = 3.14159
    delete(v1, (("" + "foo bar") + v2))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
