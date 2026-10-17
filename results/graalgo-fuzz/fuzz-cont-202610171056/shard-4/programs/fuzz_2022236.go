package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint32 }
func genFn0() uint {
    var v7 uint16 = 805
    var v8 int16 = -32768
    var v9 uint16 = 42
    fmt.Println(v7, v7, v9)
    v10 := []int{-9223372036854775808, -146, -171, 127}
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 14
}
func genFn1() bool {
    var v11 int64 = 245
    for fi12 := 0; fi12 < 13; fi12++ {
        fmt.Println(fi12)
        var v13 int8 = 100
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 3; fi14++ {
        var v15 float64 = (1e10 + (1e10 + ((1e10 / 3.14159) * (0.1 + (1.5 - 1e10)))))
        if ((0 <= 2) && true) {
            var v16 uint64 = 1
            var v17 string = "\n\t"
            _ = v16
            _ = v17
        } else {
            v18 := []int{(^(fi14)), (-(fi14))}
            _ = v18
        }
        _ = fi14
        _ = v15
    }
    _ = v11
    return false
}

func main() {
    var v1 int8 = -112
    v2 := []string{("\n\t" + "hello"), "hello", ("foo bar" + ("\n\t" + "hello")), "\n\t"}
    var v3 int8 = v1
    var v4 uint8 = 58
    v5 := []string{}
    var v6 int32 = -2147483648
    v6 = 1
    switch 118 {
    case 34:
        v6 = v6
    case 32767:
        if false {
            for fi19 := 0; fi19 < 12; fi19++ {
                for fi20 := 0; fi20 < 5; fi20++ {
                    v4 -= v4
                    _ = fi20
                }
                _ = fi19
            }
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    case -694:
        for fi21 := 0; fi21 < 13; fi21++ {
            var v22 int = fi21
            fmt.Println(v6, v1)
            for fi23 := 0; fi23 < 8; fi23++ {
                fmt.Println(fi23, fi23)
                switch fi21 {
                case 32768:
                    v5 = append(v5, "\n\t")
                case 875:
                    v5 = append(v5, ((("foo bar" + ("foo bar" + "foo bar")) + "hello") + ("\n\t" + ("foo bar" + ("hello" + "\n\t")))))
                default:
                    v22 = 497
                }
                _ = fi23
            }
            _ = fi21
            _ = v22
        }
    }
    v2 = append(v2, ("" + ("foo bar" + "hello")))
    for fi24 := 0; fi24 < 10; fi24++ {
        fmt.Println(v6)
        v25 := GenS1{F0: -9223372036854775808, F1: "", F2: 65535}
        _ = fi24
        _ = v25
    }
    fmt.Println(v1, v1, v3)
    var v26 uint64 = 0
    var v27 int32 = v6
    v28 := GenS1{F0: 42, F1: ("" + ""), F2: 0}
    var v29 float32 = 0.1
    v6 += v6
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
