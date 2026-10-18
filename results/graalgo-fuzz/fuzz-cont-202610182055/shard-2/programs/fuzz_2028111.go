package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int16 }
type GenS2 struct { F0 uint32 }

func main() {
    v1 := []int{}
    v2 := []string{"\n\t", "foo bar", ("\n\t" + "hello"), ""}
    v3 := []uint32{498}
    var v4 int8 = 1
    v5 := GenS1{F0: 741, F1: 255, F2: 42}
    _ = len(v3)
    if false {
        fmt.Println(v4)
        fmt.Println(v4)
    }
    var v6 int64 = -9223372036854775808
    v4 += 38
    fmt.Println(v6, v6)
    if len(v3) > 0 {
        v3[0] = 2147483648
    }
    if (v6 < v6) {
        if true {
            switch "hello" {
            case "":
                fmt.Println(v4, v6)
            case "a":
                v6 -= (v6 + 42)
            case "hello":
                fmt.Println(v6)
            }
            if (4294967295 > 849) {
                v1 = append(v1, -2147483648)
                var v7 uint16 = 690
                _ = v7
            } else {
                for fi8 := 0; fi8 < 3; fi8++ {
                    switch 42 {
                    case 2147483647:
                        if true {
                            v6 += v6
                            v9 := GenS1{F0: 255, F1: 0, F2: 442}
                            _ = v9
                        }
                    default:
                        v10 := GenS1{F0: 0, F1: 2, F2: 334}
                        _ = v10
                    }
                    _ = fi8
                }
            }
            var v11 int64 = v6
            _ = v11
        } else {
            switch ("a" + "a") {
            case "\n\t":
                _ = len(v2)
            case "a":
                var v12 float64 = 1e10
                _ = v12
            case "":
                if (-1 >= -310) {
                    v4 += (^(v4))
                    var v13 int8 = -128
                    v13 = (v4 ^ (v4 * (v4 - 127)))
                    _ = v13
                } else {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v4 += 1
                }
            }
            for fi14 := 0; fi14 < 16; fi14++ {
                var v15 uint16 = 32767
                v16 := GenS2{F0: 2147483648}
                _ = fi14
                _ = v15
                _ = v16
            }
        }
    }
    fmt.Println(v6, v6)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
