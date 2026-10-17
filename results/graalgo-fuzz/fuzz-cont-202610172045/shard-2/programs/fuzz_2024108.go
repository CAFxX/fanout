package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 int; F1 uint16; F2 int8 }
type GenS3 struct { F0 int8; F1 bool; F2 string }
func genFn0(p0 int8, p1 uint, p2 uint32) bool {
    v4 := []int{2147483647}
    _ = len(v4)
    _ = len(v4)
    p0 -= (^(p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return true
}

func main() {
    var v1 float64 = 3.14159
    v2 := []int{9223372036854775807, 9223372036854775807}
    var v3 bool = (-9223372036854775808 >= 9223372036854775807)
    v2 = append(v2, 15)
    v3 = (!v3)
    for fi5 := 0; fi5 < 2; fi5++ {
        for fi6 := 0; fi6 < 7; fi6++ {
            if v3 {
                if v3 {
                    v3 = (!(v3 || v3))
                    v1 += v1
                } else {
                    var v7 bool = ("" > "\n\t")
                    v2 = append(v2, fi6)
                    _ = v7
                }
                var v8 int64 = -382
                _ = v8
            } else {
                v9 := GenS1{F0: "\n\t", F1: 152}
                if v3 {
                    _ = len(v2)
                    v1 -= (1e10 - float64((1.5 / -2.25)))
                    v1 = ((v1 / (v1 - (v1 - v1))) * (0.0 / 1.5))
                } else {
                    fmt.Println(v3)
                    fmt.Println(v3, fi6, v3)
                }
                if true {
                    if len(v2) > 0 {
                        v2[0] = (fi6 % ((32767) | 1))
                    }
                } else {
                    v1 -= float64(-2.25)
                    var v10 uint32 = 2147483648
                    v3 = v3
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if true {
                        v11 := GenS2{F0: (fi6 << ((892) & 63)), F1: 65535, F2: 42}
                        fmt.Println(fi6, v1)
                        _ = v11
                    } else {
                        v10 = (v10 % ((v10) | 1))
                        _ = len(v2)
                        fmt.Println(fi6)
                    }
                    fmt.Println(v3, v10)
                    var v12 int8 = 0
                    _ = v10
                    _ = v12
                }
                _ = v9
            }
            _ = fi6
        }
        _ = fi5
    }
    fmt.Println(v1, v1)
    v2 = append(v2, -2147483648)
    v3 = v3
    var v13 uint8 = 128
    v13 += v13
    _ = len(v2)
    _ = len(v2)
    for fi14 := 0; fi14 < 1; fi14++ {
        for fi15 := 0; fi15 < 20; fi15++ {
            var v16 float32 = (3.14159 - (-2.25 / 1e10))
            switch 32767 {
            case 32767:
                v1 = ((0.1 / v1) - float64(v16))
            default:
                v17 := GenS1{F0: "hello", F1: 42}
                _ = v17
            }
            if (!v3) {
                switch ("\n\t" + (("hello" + "a") + (("foo bar" + "") + ("hello" + ("" + ""))))) {
                case "\n\t":
                    v18 := GenS3{F0: -1, F1: v3, F2: "a"}
                    _ = v18
                }
                if (9223372036854775807 == 62) {
                    v13 -= v13
                }
                fmt.Println(fi14, v16)
            } else {
                var v19 uint32 = 0
                v19 = v19
                var v20 uint = 876
                _ = v19
                _ = v20
            }
            _ = fi15
            _ = v16
        }
        _ = fi14
    }
    v3 = true
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
}
