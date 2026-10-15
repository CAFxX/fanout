package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 uint8; F1 uint; F2 int }

func main() {
    v1 := []int{128, 9223372036854775807}
    var v2 float64 = -2.25
    var v3 float32 = (((0.0 * 1e10) + ((0.0 * 1.5) + 0.1)) + -2.25)
    v4 := make(map[string]int)
    v5 := make(map[int]string)
    v6 := []uint32{2147483647, 2, 4294967295}
    v4[("hello" + "\n\t")] = 42
    if (!(815 < -239)) {
        for fi7 := 0; fi7 < 8; fi7++ {
            fmt.Println(fi7, v2)
            _ = fi7
        }
    }
    fmt.Println(v2, v2)
    for fi8 := 0; fi8 < 2; fi8++ {
        fmt.Println(v3, fi8)
        _ = fi8
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        v2 = 1e10
        _ = fi9
    }
    if (-2.25 < v2) {
        v10 := GenS1{F0: 2147483647, F1: false}
        v3 -= v3
        v2 += v2
        _ = v10
    }
    v2 = v2
    if (42 < 123) {
        for fi11 := 0; fi11 < 10; fi11++ {
            v12 := GenS2{F0: 1, F1: 2, F2: (fi11 >> ((876) & 63))}
            switch (fi11 / ((-9223372036854775808) | 1)) {
            case 2147483647:
                if (("foo bar" + ("a" + "\n\t")) >= "a") {
                    delete(v4, ((("" + (("" + "") + "a")) + ("" + "\n\t")) + "hello"))
                    v13 := v4["foo bar"]
                    _ = v13
                } else {
                    v3 -= ((float32(v2) + v3) / v3)
                    if ((!false) || (fi11 >= (^(fi11)))) {
                        fmt.Println(fi11, v3)
                    } else {
                        v14 := GenS2{F0: 0, F1: 0, F2: fi11}
                        switch "hello" {
                        case "\n\t":
                            var v15 uint = 0
                            _ = v15
                        case "":
                            v16 := GenS2{F0: 128, F1: 1, F2: fi11}
                            _ = v16
                        case "hello":
                            switch "foo bar" {
                            case "a":
                                for fi17 := 0; fi17 < 12; fi17++ {
                                    var v18 int32 = -1
                                    var v19 int = fi11
                                    _ = fi17
                                    _ = v18
                                    _ = v19
                                }
                            }
                        }
                        _ = v14
                    }
                }
            case -2147483648:
                v20 := GenS1{F0: 0, F1: true}
                _ = v20
            }
            v21 := GenS1{F0: 36, F1: ((((2 != 42) && (!(1 == 181))) && (!(-123 >= 127))) || (-1 <= -128))}
            _ = fi11
            _ = v12
            _ = v21
        }
    }
    if (false && (false || false)) {
        _ = len(v6)
        if len(v1) > 0 {
            v1[0] = 1
        }
    } else {
        v3 -= (((v3 - v3) / ((v3 * v3) * v3)) + 3.14159)
    }
    v2 += v2
    var v22 uint16 = 65535
    _ = len(v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
