package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 uint32; F1 string; F2 bool }
func genFn0(p0 int64) uint {
    var v5 float32 = float32((1.5 * (3.14159 + (0.0 / 0.1))))
    v6 := GenS1{F0: (((false || (("hello" + "") <= "a")) && (65535 == 2)) || false), F1: 42}
    p0 += p0
    p0 = p0
    _ = p0
    _ = v5
    _ = v6
    return 494
}

func main() {
    var v1 bool = (255 >= 65535)
    var v2 float32 = ((((1e10 - (0.0 * 1.5)) * 1e10) * 0.1) * (float32(3.14159) - 1e10))
    var v3 float64 = 3.14159
    var v4 uint64 = 42
    v1 = v1
    for fi7 := 0; fi7 < 2; fi7++ {
        fmt.Println(fi7)
        switch "hello" {
        case "hello":
            if false {
                v8 := []int{(fi7 << ((genFn0(0)) & 63)), fi7, (fi7 | -9223372036854775808)}
                _ = v8
            } else {
                var v9 int32 = 42
                if true {
                    switch ("a" + "hello") {
                    case "foo bar":
                        if v1 {
                            if v1 {
                                v2 += 0.1
                            }
                            v10 := []int{}
                            v11 := GenS1{F0: true, F1: v9}
                            _ = v10
                            _ = v11
                        }
                    case "hello":
                        v4 = v4
                    case "a":
                        v12 := []int{}
                        _ = v12
                    default:
                        v3 -= (1.5 / v3)
                    }
                    v13 := []int{fi7, fi7, 4294967295}
                    var v14 int8 = -128
                    _ = v13
                    _ = v14
                }
                var v15 float32 = v2
                _ = v15
                _ = v9
            }
        default:
            for fi16 := 0; fi16 < 20; fi16++ {
                v17 := []int{fi16, fi7, -465, -895}
                if false {
                    if (!(65535 != 2)) {
                        switch fi16 {
                        case 32767:
                            var v18 float32 = 0.0
                            v18 += float32(v3)
                            v19 := GenS1{F0: (207 <= 127), F1: -1}
                            _ = v18
                            _ = v19
                        case -9223372036854775808:
                            if len(v17) > 0 {
                                v17[0] = (fi16 % (((fi16 % ((fi16) | 1))) | 1))
                            }
                        }
                        fmt.Println(fi16, v3, v2)
                        v1 = (!v1)
                    } else {
                        if false {
                            v2 -= float32(((float64(v2) * v3) - (1.5 * v3)))
                        }
                        v3 = (((v3 / 1.0) - (1.5 - 1e10)) / v3)
                    }
                    var v20 float32 = v2
                    var v21 int64 = 9223372036854775807
                    _ = v20
                    _ = v21
                }
                var v22 uint8 = 213
                _ = fi16
                _ = v17
                _ = v22
            }
        }
        _ = fi7
    }
    v2 = float32(v3)
    var v23 uint8 = 2
    v1 = (32767 <= 906)
    v24 := GenS2{F0: 485, F1: ("hello" + (((("foo bar" + ("\n\t" + "\n\t")) + ("foo bar" + ("a" + "\n\t"))) + ("foo bar" + ("a" + ("foo bar" + (("\n\t" + "a") + (("a" + ("\n\t" + "\n\t")) + ("foo bar" + "foo bar"))))))) + ("\n\t" + "\n\t"))), F2: (2147483648 <= 2147483648)}
    var v25 int64 = int64(v4)
    var v26 int8 = 42
    var v27 uint32 = 2147483648
    v28 := GenS1{F0: (v1 || (1 > 2147483647)), F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
