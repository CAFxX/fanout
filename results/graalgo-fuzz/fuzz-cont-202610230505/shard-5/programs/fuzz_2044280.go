package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
func genFn0(p0 uint8) int {
    switch "a" {
    case "hello":
        v5 := GenS1{F0: 1000, F1: p0}
        _ = v5
    }
    var v6 int8 = 15
    for fi7 := 0; fi7 < 12; fi7++ {
        fmt.Println(fi7, v6, v6)
        p0 -= 127
        v8 := GenS1{F0: 0, F1: p0}
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = v6
    return 1
}
func genFn1(p0 int32) uint8 {
    fmt.Println(p0)
    for fi9 := 0; fi9 < 7; fi9++ {
        var v10 bool = (false && true)
        v11 := GenS1{F0: 0, F1: 252}
        p0 = p0
        _ = fi9
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 17; fi12++ {
        p0 = (p0 | p0)
        p0 += p0
        v13 := []int{fi12, int(p0), -2147483648}
        _ = fi12
        _ = v13
    }
    p0 += (p0 * (p0 << ((909) & 63)))
    _ = p0
    return 128
}

func main() {
    var v1 uint64 = 767
    var v2 float64 = 0.1
    var v3 bool = (!(1.5 < 0.1))
    var v4 uint16 = uint16(v1)
    for fi14 := 0; fi14 < 10; fi14++ {
        var v15 uint64 = (v1 * 0)
        v16 := []int{(fi14 - fi14), (fi14 >> ((358) & 63)), (fi14 ^ (fi14 - (fi14 + genFn0(0))))}
        _ = fi14
        _ = v15
        _ = v16
    }
    v17 := []int{}
    v1 = (v1 * (v1 * 18446744073709551615))
    var v18 uint8 = 0
    if (617 >= 42) {
        fmt.Println(v2, v18)
    }
    for fi19 := 0; fi19 < 19; fi19++ {
        v1 = v1
        if len(v17) > 0 {
            v17[0] = genFn0(v18)
        }
        _ = fi19
    }
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v18)
    }
    v17 = append(v17, -388)
    _ = len(v17)
    if (v3 && ((1.5 < 3.14159) || (genFn0((v18 << ((2) & 63))) != 128))) {
        fmt.Println(v2)
    }
    var v20 int = -9223372036854775808
    for fi21 := 0; fi21 < 1; fi21++ {
        var v22 string = "foo bar"
        _ = fi21
        _ = v22
    }
    var v23 uint16 = v4
    for fi24 := 0; fi24 < 10; fi24++ {
        for fi25 := 0; fi25 < 8; fi25++ {
            _ = len(v17)
            var v26 int8 = 100
            _ = fi25
            _ = v26
        }
        for fi27 := 0; fi27 < 7; fi27++ {
            var v28 int = (fi27 / ((fi24) | 1))
            v18 += v18
            _ = fi27
            _ = v28
        }
        _ = fi24
    }
    if v3 {
        v18 = genFn1(2147483647)
        for fi29 := 0; fi29 < 15; fi29++ {
            switch ("" + "a") {
            case "hello":
                v2 += 0.0
            case "\n\t":
                if (false || (42 != 1)) {
                    var v30 uint32 = 42
                    if v3 {
                        v4 = 42
                    } else {
                        var v31 uint = 255
                        var v32 uint32 = (v30 & (v30 | 1))
                        for fi33 := 0; fi33 < 8; fi33++ {
                            var v34 int16 = 1000
                            _ = fi33
                            _ = v34
                        }
                        _ = v31
                        _ = v32
                    }
                    _ = v30
                } else {
                    fmt.Println(v18, v3)
                    v17 = append(v17, genFn0(uint8(v4)))
                    var v35 int16 = 1
                    _ = v35
                }
            }
            fmt.Println(v23, v20)
            v36 := GenS1{F0: -1, F1: genFn1(0)}
            _ = fi29
            _ = v36
        }
    }
    v18 += 2
    v20 += (-(v20))
    v2 = ((float64(1e10) * v2) / ((float64(0.1) * v2) + v2))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
}
