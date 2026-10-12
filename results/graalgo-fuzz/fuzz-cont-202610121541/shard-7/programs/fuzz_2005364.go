package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
func genFn0(p0 int, p1 uint, p2 bool) string {
    p1 = (p1 % ((p1) | 1))
    var v4 string = ("a" + "\n\t")
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return "\n\t"
}

func main() {
    v1 := make(map[int]bool)
    var v2 string = ("" + "hello")
    v3 := make(map[int]bool)
    v2 = v2
    if false {
        for fi5 := 0; fi5 < 1; fi5++ {
            var v6 bool = true
            v2 = (v2 + v2)
            _ = fi5
            _ = v6
        }
    }
    if ((v2 + v2) > "foo bar") {
        var v7 int32 = -530
        _ = v7
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 bool = (728 <= 1)
        _ = fi8
        _ = v9
    }
    var v10 uint8 = 1
    v2 = ""
    v10 = v10
    switch -9223372036854775808 {
    case 32767:
        for fi11 := 0; fi11 < 11; fi11++ {
            v2 = (("" + ((v2 + v2) + (v2 + (("foo bar" + v2) + ("a" + (v2 + v2)))))) + "\n\t")
            _ = fi11
        }
    case 469:
        fmt.Println(v10, v2)
    case -42:
        v10 = (v10 * v10)
    default:
        for fi12 := 0; fi12 < 2; fi12++ {
            if false {
                v13 := GenS1{F0: (float32(0.0) != (3.14159 - float32((0.1 / float64(0.1))))), F1: 1}
                v14 := []int{fi12, fi12}
                if false {
                    var v15 float32 = 3.14159
                    switch v2 {
                    case "foo bar":
                        v16 := GenS1{F0: (((v15 * (1e10 / (float32(0.0) - v15))) * v15) <= v15), F1: 127}
                        _ = v16
                    case "\n\t":
                        v2 = "a"
                    case "hello":
                        if (-582 < (fi12 & (fi12 & (fi12 ^ fi12)))) {
                            for fi17 := 0; fi17 < 6; fi17++ {
                                v2 = v2
                                var v18 bool = (215 >= 255)
                                _ = fi17
                                _ = v18
                            }
                            fmt.Println(v2, v10, v2)
                            var v19 int8 = 127
                            _ = v19
                        }
                    default:
                        v15 = 1e10
                    }
                    _ = v15
                }
                _ = v13
                _ = v14
            }
            for fi20 := 0; fi20 < 19; fi20++ {
                var v21 int32 = -911
                _ = fi20
                _ = v21
            }
            _ = fi12
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v2
    _ = v3
}
